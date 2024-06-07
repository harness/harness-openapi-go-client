# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedSecret**](AccountSecretApi.md#CreateAccountScopedSecret) | **Post** /v1/secrets | Create a secret
[**DeleteAccountScopedSecret**](AccountSecretApi.md#DeleteAccountScopedSecret) | **Delete** /v1/secrets/{secret} | Deletes a secret
[**GetAccountScopedSecret**](AccountSecretApi.md#GetAccountScopedSecret) | **Get** /v1/secrets/{secret} | Retrieve a secret
[**GetAccountScopedSecrets**](AccountSecretApi.md#GetAccountScopedSecrets) | **Get** /v1/secrets | List secrets
[**UpdateAccountScopedSecret**](AccountSecretApi.md#UpdateAccountScopedSecret) | **Put** /v1/secrets/{secret} | Update a secret
[**ValidateAccountSecretRef**](AccountSecretApi.md#ValidateAccountSecretRef) | **Post** /v1/secrets/validate-secret-ref | Validate secret reference

# **CreateAccountScopedSecret**
> SecretResponse CreateAccountScopedSecret(ctx, body, spec, file, optional)
Create a secret

Creates a new secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequest**](SecretRequest.md)|  | 
  **spec** | [**SecretRequest**](.md)|  | 
  **file** | ***os.File*****os.File**|  | 
 **optional** | ***AccountSecretApiCreateAccountScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSecretApiCreateAccountScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **privateSecret** | **optional.**| This would be used to define secret as private. | [default to false]
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountScopedSecret**
> SecretResponse DeleteAccountScopedSecret(ctx, secret, optional)
Deletes a secret

Deletes the information of the secret with the matching secret identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **secret** | **string**| Identifier field of the secret | 
 **optional** | ***AccountSecretApiDeleteAccountScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSecretApiDeleteAccountScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedSecret**
> SecretResponse GetAccountScopedSecret(ctx, secret, optional)
Retrieve a secret

Retrieves the information of the secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **secret** | **string**| Identifier field of the secret | 
 **optional** | ***AccountSecretApiGetAccountScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSecretApiGetAccountScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedSecrets**
> []SecretResponse GetAccountScopedSecrets(ctx, optional)
List secrets

Retrieves the information of the secrets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountSecretApiGetAccountScopedSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSecretApiGetAccountScopedSecretsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | [**optional.Interface of []string**](string.md)| Identifier field of secrets | 
 **type_** | [**optional.Interface of []string**](string.md)| Secret types on which the filter will be applied | 
 **recursive** | **optional.Bool**| Expand current scope to include all child scopes  | [default to false]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountScopedSecret**
> SecretResponse UpdateAccountScopedSecret(ctx, body, spec, file, secret, optional)
Update a secret

Updates the information of the secret with the matching secret identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequest**](SecretRequest.md)|  | 
  **spec** | [**SecretRequest**](.md)|  | 
  **file** | ***os.File*****os.File**|  | 
  **secret** | **string**| Identifier field of the secret | 
 **optional** | ***AccountSecretApiUpdateAccountScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSecretApiUpdateAccountScopedSecretOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**SecretResponse**](SecretResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, multipart/form-data
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateAccountSecretRef**
> SecretValidationResponse ValidateAccountSecretRef(ctx, optional)
Validate secret reference

Validates if the secret at the secretManager path can be referenced

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountSecretApiValidateAccountSecretRefOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountSecretApiValidateAccountSecretRefOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SecretRequest**](SecretRequest.md)| Details of the secret reference | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**SecretValidationResponse**](SecretValidationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

