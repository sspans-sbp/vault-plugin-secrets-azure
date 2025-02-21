## Unreleased

## v0.17.1

BUG FIXES:
* Add nil check for response when unassigning roles [[GH-191]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/191)

## v0.17.0

IMPROVEMENTS:
* Update dependencies [[GH-176]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/176)
  * github.com/Azure/azure-sdk-for-go/sdk/azcore v1.9.0 -> v1.9.1
  * github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.4.0 -> v1.5.1
  * github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2 v2.1.1 -> v2.2.0
  * github.com/google/uuid v1.3.1 -> v1.6.0
  * github.com/hashicorp/go-hclog v1.5.0 -> v1.6.2
  * github.com/hashicorp/vault/api v1.10.0 -> v1.11.0
  * github.com/hashicorp/vault/sdk v0.10.0 -> v0.10.2
  * github.com/microsoftgraph/msgraph-sdk-go v1.22.0 -> v1.32.0
  * github.com/microsoftgraph/msgraph-sdk-go-core v1.0.0 -> v1.0.1

## v0.16.3

IMPROVEMENTS:
* Add sign_in_audience and tags fields to application registration [GH-174](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/174)
* Prevent write-ahead-log data from being replicated to performance secondaries [GH-164](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/164)
* Update dependencies [[GH-161]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/161)
  * github.com/Azure/azure-sdk-for-go v68.0.0
* Update dependencies [[GH-162]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/162)
  * golang.org/x/crypto v0.13.0
  * golang.org/x/net v0.15.0
  * golang.org/x/sys v0.12.0
  * golang.org/x/text v0.13.0

## v0.16.2

IMPROVEMENTS:
* Update dependencies [[GH-160]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/160)
  * github.com/hashicorp/vault/api v1.9.1 -> v1.10.0
  * github.com/hashicorp/vault/sdk v0.9.0 -> v0.10.0

## v0.16.1

BUG FIXES:
* Fix intermittent 401s by preventing performance secondary clusters from rotating root credentials [[GH-150]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/150)

## v0.16.0

IMPROVEMENTS:

* permanently delete app during WAL rollback [GH-138](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/138)
* enable plugin multiplexing [GH-134](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/134)
* add display attributes for OpenAPI OperationID's [GH-141](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/141)
* update dependencies
  * `github.com/hashicorp/vault/api` v1.9.1 [GH-145](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/145)
  * `github.com/hashicorp/vault/sdk` v0.9.0 [GH-141](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/141)
  * `github.com/hashicorp/go-hclog` v1.5.0 [GH-140](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/140)
  * `github.com/Azure/go-autorest/autorest` v0.11.29 [GH-144](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/144)

## v0.15.1

BUG FIXES:

* Fix intermittent 401s by preventing performance secondary clusters from rotating root credentials [[GH-150]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/150)

## v0.15.0

CHANGES:

* Changes user-agent header value to use correct Vault version information and include
  the plugin type and name in the comment section [[GH-123]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/123)

FEATURES:

* Adds ability to persist an application for the lifetime of a role [[GH-98]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/98)

IMPROVEMENTS:

* Updated dependencies [[GH-109](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/109)]
    * `github.com/Azure/azure-sdk-for-go v67.0.0+incompatible`
    * `github.com/Azure/go-autorest/autorest v0.11.28`
    * `github.com/Azure/go-autorest/autorest/azure/auth v0.5.11`
    * `github.com/hashicorp/go-hclog v1.3.1`
    * `github.com/hashicorp/go-uuid v1.0.3`
    * `github.com/hashicorp/vault/api v1.8.2`
    * `github.com/hashicorp/vault/sdk v0.6.1`
    * `github.com/mitchellh/mapstructure v1.5.0`
* Upgraded to go 1.19 [[GH-109](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/109)]

## v0.14.2

BUG FIXES:

* Fix intermittent 401s by preventing performance secondary clusters from rotating root credentials [[GH-150]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/150)

## v0.14.1

BUG FIXES:

* Adds WAL rollback mechanism to clean up Role Assignments during partial failure [[GH-110]](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/110)

## v0.14.0

IMPROVEMENTS:

* Add option to permanently delete AzureAD objects [[GH-104](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/104)]

CHANGES:

* Remove deprecated AAD graph code [[GH-101](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/101)]
* Remove partner ID from user agent string [[GH-95](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/95)]

## v0.11.4

CHANGES:

* Sets `use_microsoft_graph_api` to true by default [[GH-90](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/90)]

BUG FIXES:

* Fixes environment not being used when using MS Graph [[GH-87](https://github.com/hashicorp/vault-plugin-secrets-azure/pull/87)]
