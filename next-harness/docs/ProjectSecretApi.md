# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectScopedSecret**](ProjectSecretApi.md#CreateProjectScopedSecret) | **Post** /v1/orgs/{org}/projects/{project}/secrets | Create a secret
[**DeleteProjectScopedSecret**](ProjectSecretApi.md#DeleteProjectScopedSecret) | **Delete** /v1/orgs/{org}/projects/{project}/secrets/{secret} | Delete a secret
[**GetProjectScopedSecret**](ProjectSecretApi.md#GetProjectScopedSecret) | **Get** /v1/orgs/{org}/projects/{project}/secrets/{secret} | Retrieve a secret
[**GetProjectScopedSecrets**](ProjectSecretApi.md#GetProjectScopedSecrets) | **Get** /v1/orgs/{org}/projects/{project}/secrets | List secrets
[**UpdateProjectScopedSecret**](ProjectSecretApi.md#UpdateProjectScopedSecret) | **Put** /v1/orgs/{org}/projects/{project}/secrets/{secret} | Update a secret
[**ValidateProjectSecretRef**](ProjectSecretApi.md#ValidateProjectSecretRef) | **Post** /v1/orgs/{org}/projects/{project}/secrets/validate-secret-ref | Validate secret reference

# **CreateProjectScopedSecret**
> SecretResponse CreateProjectScopedSecret(ctx, body, spec, file, org, project, optional)
Create a secret

Creates a new secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequest**](SecretRequest.md)|  | 
  **spec** | [**SecretRequest**](.md)|  | 
  **file** | ***os.File*****os.File**|  | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectSecretApiCreateProjectScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSecretApiCreateProjectScopedSecretOpts struct
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

# **DeleteProjectScopedSecret**
> SecretResponse DeleteProjectScopedSecret(ctx, org, project, secret, optional)
Delete a secret

Deletes the information of the secret with the matching secret identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **secret** | **string**| Identifier field of the secret | 
 **optional** | ***ProjectSecretApiDeleteProjectScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSecretApiDeleteProjectScopedSecretOpts struct
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

# **GetProjectScopedSecret**
> SecretResponse GetProjectScopedSecret(ctx, org, project, secret, optional)
Retrieve a secret

Retrieves the information of the secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **secret** | **string**| Identifier field of the secret | 
 **optional** | ***ProjectSecretApiGetProjectScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSecretApiGetProjectScopedSecretOpts struct
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

# **GetProjectScopedSecrets**
> []SecretResponse GetProjectScopedSecrets(ctx, org, project, optional)
List secrets

Retrieves the information of the secrets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectSecretApiGetProjectScopedSecretsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSecretApiGetProjectScopedSecretsOpts struct
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

# **UpdateProjectScopedSecret**
> SecretResponse UpdateProjectScopedSecret(ctx, body, spec, file, org, project, secret, optional)
Update a secret

Updates the information of the secret with the matching secret identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SecretRequest**](SecretRequest.md)|  | 
  **spec** | [**SecretRequest**](.md)|  | 
  **file** | ***os.File*****os.File**|  | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **secret** | **string**| Identifier field of the secret | 
 **optional** | ***ProjectSecretApiUpdateProjectScopedSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSecretApiUpdateProjectScopedSecretOpts struct
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

# **ValidateProjectSecretRef**
> SecretValidationResponse ValidateProjectSecretRef(ctx, org, project, optional)
Validate secret reference

Validates if the secret at the secretManager path can be referenced

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project to which the resource is scoped. | 
 **optional** | ***ProjectSecretApiValidateProjectSecretRefOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectSecretApiValidateProjectSecretRefOpts struct
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

