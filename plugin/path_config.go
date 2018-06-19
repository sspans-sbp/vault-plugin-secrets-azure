package azuresecrets

import (
	"context"
	"errors"
	"time"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/errwrap"
	multierror "github.com/hashicorp/go-multierror"
	uuid "github.com/hashicorp/go-uuid"
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

const (
	configPath = "config"
)

type azureConfig struct {
	SubscriptionID string        `json:"subscription_id"`
	TenantID       string        `json:"tenant_id"`
	ClientID       string        `json:"client_id"`
	ClientSecret   string        `json:"client_secret"`
	DefaultTTL     time.Duration `json:"ttl"`
	MaxTTL         time.Duration `json:"max_ttl"`
	Resource       string        `json:"resource"`
	Environment    string        `json:"environment"`
}

func pathConfig(b *azureSecretBackend) *framework.Path {
	return &framework.Path{
		Pattern: "config",
		Fields: map[string]*framework.FieldSchema{
			"subscription_id": &framework.FieldSchema{
				Type: framework.TypeString,
				Description: `The subscription id for the Azure Active Directory.
				This value can also be provided with the AZURE_SUBSCREIPTION_ID environment variable.`,
			},
			"tenant_id": &framework.FieldSchema{
				Type: framework.TypeString,
				Description: `The tenant id for the Azure Active Directory. This value can also
				be provided with the AZURE_TENANT_ID environment variable.`,
			},
			"environment": &framework.FieldSchema{
				Type: framework.TypeString,
				Description: `The Azure environment name. If not provided, AzurePublicCloud is used.
				This value can also be provided with the AZURE_ENVIRONMENT environment variable.`,
			},
			"client_id": &framework.FieldSchema{
				Type: framework.TypeString,
				Description: `The OAuth2 client id to connection to Azure.
				This value can also be provided with the AZURE_CLIENT_ID environment variable.`,
			},
			"client_secret": &framework.FieldSchema{
				Type: framework.TypeString,
				Description: `The OAuth2 client secret to connection to Azure.
				This value can also be provided with the AZURE_CLIENT_SECRET environment variable.`,
			},
			"resource": &framework.FieldSchema{
				Type: framework.TypeString,
				Description: `The resource URL for the vault application in Azure Active Directory.
				This value can also be provided with the AZURE_AD_RESOURCE environment variable.`,
			},
			"ttl": {
				Type:        framework.TypeDurationSecond,
				Description: "Default lease for generated credentials. If <= 0, will use system default.",
			},
			"max_ttl": {
				Type:        framework.TypeDurationSecond,
				Description: "Maximum time a service principal. If <= 0, will use system default.",
			},
		},
		Callbacks: map[logical.Operation]framework.OperationFunc{
			logical.ReadOperation:   b.pathConfigRead,
			logical.CreateOperation: b.pathConfigWrite,
			logical.UpdateOperation: b.pathConfigWrite,
		},
		ExistenceCheck:  b.pathConfigExistenceCheck,
		HelpSynopsis:    confHelpSyn,
		HelpDescription: confHelpDesc,
	}
}

func (b *azureSecretBackend) pathConfigExistenceCheck(ctx context.Context, req *logical.Request, data *framework.FieldData) (bool, error) {
	config, err := b.getConfig(ctx, req.Storage)
	if err != nil {
		return false, err
	}
	return config != nil, nil
}

func (b *azureSecretBackend) pathConfigWrite(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	var merr *multierror.Error

	b.cfgLock.Lock()
	defer b.cfgLock.Unlock()

	config, err := b.getConfig(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	if config == nil {
		config = new(azureConfig)
	}

	if subscriptionID, ok := data.GetOk("subscription_id"); ok {
		s := subscriptionID.(string)
		if _, err := uuid.ParseUUID(s); err != nil {
			merr = multierror.Append(merr, errwrap.Wrapf("subscription_id format error: {{err}}", err))
		} else {
			config.SubscriptionID = s
		}
	}

	if tenantID, ok := data.GetOk("tenant_id"); ok {
		t := tenantID.(string)
		if _, err := uuid.ParseUUID(t); err != nil {
			merr = multierror.Append(merr, errwrap.Wrapf("tenant_id format error: {{err}}", err))
		} else {
			config.TenantID = t
		}
	}

	if environment, ok := data.GetOk("environment"); ok {
		e := environment.(string)
		if _, err := azure.EnvironmentFromName(e); err != nil {
			merr = multierror.Append(merr, err)
		} else {
			config.Environment = e
		}
	}

	if resource, ok := data.GetOk("resource"); ok {
		config.Resource = resource.(string)
	}

	if clientID, ok := data.GetOk("client_id"); ok {
		config.ClientID = clientID.(string)
	}

	if clientSecret, ok := data.GetOk("client_secret"); ok {
		config.ClientSecret = clientSecret.(string)
	}

	if ttlRaw, ok := data.GetOk("ttl"); ok {
		config.DefaultTTL = time.Duration(ttlRaw.(int)) * time.Second
	}

	if maxTTLRaw, ok := data.GetOk("max_ttl"); ok {
		config.MaxTTL = time.Duration(maxTTLRaw.(int)) * time.Second
	}

	// validate ttl constraints
	if config.DefaultTTL < 0 {
		merr = multierror.Append(merr, errors.New("ttl < 0"))
	}
	if config.MaxTTL < 0 {
		merr = multierror.Append(merr, errors.New("max_ttl < 0"))
	}
	if config.DefaultTTL > config.MaxTTL && config.MaxTTL != 0 {
		merr = multierror.Append(merr, errors.New("ttl > max_ttl"))
	}

	if merr.ErrorOrNil() != nil {
		return logical.ErrorResponse(merr.Error()), nil
	}

	err = b.saveConfig(ctx, config, req.Storage)
	if err == nil {
		b.reset()
	}

	return nil, err
}

func (b *azureSecretBackend) pathConfigRead(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	config, err := b.getConfig(ctx, req.Storage)
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, nil
	}

	resp := &logical.Response{
		Data: map[string]interface{}{
			"subscription_id": config.SubscriptionID,
			"tenant_id":       config.TenantID,
			"environment":     config.Environment,
			"client_id":       config.ClientID,
			"ttl":             int64(config.DefaultTTL / time.Second),
			"max_ttl":         int64(config.MaxTTL / time.Second),
		},
	}
	return resp, nil
}

func (b *azureSecretBackend) getConfig(ctx context.Context, s logical.Storage) (*azureConfig, error) {
	config := new(azureConfig)
	entry, err := s.Get(ctx, configPath)
	if err != nil {
		return nil, err
	}
	if entry == nil {
		return config, nil
	}

	if err := entry.DecodeJSON(config); err != nil {
		return nil, err
	}
	return config, nil
}

func (b *azureSecretBackend) saveConfig(ctx context.Context, cfg *azureConfig, s logical.Storage) error {
	entry, err := logical.StorageEntryJSON(configPath, cfg)
	if err != nil {
		return err
	}
	err = s.Put(ctx, entry)

	return nil
}

const confHelpSyn = `Configures the Azure secret backend.`
const confHelpDesc = `
The Azure secret backend ............... validates the login JWTs using the
configured credentials.  In order to validate machine information, the
OAuth2 client id and secret are used to query the Azure API.  The OAuth2
credentials require Microsoft.Compute/virtualMachines/read permission on
the resource requesting credentials.
`
