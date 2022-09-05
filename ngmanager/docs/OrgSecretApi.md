# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedSecret**](OrgSecretApi.md#CreateOrgScopedSecret) | **Post** /v1/orgs/{org}/secrets | Create a secret
[**DeleteOrgScopedSecret**](OrgSecretApi.md#DeleteOrgScopedSecret) | **Delete** /v1/org/{org}/secrets/{secret} | Delete a secret
[**GetOrgScopedSecret**](OrgSecretApi.md#GetOrgScopedSecret) | **Get** /v1/org/{org}/secrets/{secret} | Retrieve a secret
[**GetOrgScopedSecrets**](OrgSecretApi.md#GetOrgScopedSecrets) | **Get** /v1/orgs/{org}/secrets | List secrets
[**UpdateOrgScopedSecret**](OrgSecretApi.md#UpdateOrgScopedSecret) | **Put** /v1/org/{org}/secrets/{secret} | Update a secret
[**ValidateUniqueOrgScopedSecretSlug**](OrgSecretApi.md#ValidateUniqueOrgScopedSecretSlug) | **Head** /v1/org/{org}/secrets/{secret} | Validate unique secret slug

# **CreateOrgScopedSecret**
> SecretResponse CreateOrgScopedSecret(ctx, body, spec, file, org, optional)
Create a secret

Creates a new secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequest**](SecretRequest.md)|  | 
  **spec** | [**SecretRequest**](.md)|  | 
  **file** | ***os.File*****os.File**|  | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
 **optional** | ***OrgSecretApiCreateOrgScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgSecretApiCreateOrgScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **account** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 
 **privateSecret** | **optional.**| This would be used to define secret as private. | [default to false]

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgScopedSecret**
> SecretResponse DeleteOrgScopedSecret(ctx, org, secret, optional)
Delete a secret

Deletes the information of the secret with the matching secret slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **secret** | **string**| Slug field of the secret | 
 **optional** | ***OrgSecretApiDeleteOrgScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgSecretApiDeleteOrgScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedSecret**
> SecretResponse GetOrgScopedSecret(ctx, org, secret, optional)
Retrieve a secret

Retrieves the information of the secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **secret** | **string**| Slug field of the secret | 
 **optional** | ***OrgSecretApiGetOrgScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgSecretApiGetOrgScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedSecrets**
> []SecretResponse GetOrgScopedSecrets(ctx, org, optional)
List secrets

Retrieves the information of the secrets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
 **optional** | ***OrgSecretApiGetOrgScopedSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgSecretApiGetOrgScopedSecretsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 
 **project** | **optional.String**| Slug field of the project the resource is scoped to | 
 **secret** | [**optional.Interface of []string**](string.md)| Slug field of secrets | 
 **type_** | [**optional.Interface of []string**](string.md)| Secret types on which the filter will be applied | 
 **recursive** | **optional.Bool**| Expand current scope to include all child scopes  | [default to false]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]

### Return type

[**[]SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedSecret**
> SecretResponse UpdateOrgScopedSecret(ctx, body, spec, file, org, secret, optional)
Update a secret

Updates the information of the secret with the matching secret slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequest**](SecretRequest.md)|  | 
  **spec** | [**SecretRequest**](.md)|  | 
  **file** | ***os.File*****os.File**|  | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **secret** | **string**| Slug field of the secret | 
 **optional** | ***OrgSecretApiUpdateOrgScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgSecretApiUpdateOrgScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **account** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateUniqueOrgScopedSecretSlug**
> ValidateSecretSlugResponse ValidateUniqueOrgScopedSecretSlug(ctx, org, secret, optional)
Validate unique secret slug

Validates org scoped secret slug is unique

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **secret** | **string**| Slug field of the secret | 
 **optional** | ***OrgSecretApiValidateUniqueOrgScopedSecretSlugOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgSecretApiValidateUniqueOrgScopedSecretSlugOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**ValidateSecretSlugResponse**](ValidateSecretSlugResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

