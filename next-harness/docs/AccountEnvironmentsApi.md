# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedEnvironment**](AccountEnvironmentsApi.md#CreateAccountScopedEnvironment) | **Post** /v1/environments | Creates an account scoped Environment
[**DeleteAccountScopedEnvironment**](AccountEnvironmentsApi.md#DeleteAccountScopedEnvironment) | **Delete** /v1/environments/{environment} | Delete an account scoped Environment
[**GetAccountScopedEnvironment**](AccountEnvironmentsApi.md#GetAccountScopedEnvironment) | **Get** /v1/environments/{environment} | Retrieve an account scoped environment
[**GetAccountScopedEnvironments**](AccountEnvironmentsApi.md#GetAccountScopedEnvironments) | **Get** /v1/environments | List account scoped Environments
[**UpdateAccountScopedEnvironment**](AccountEnvironmentsApi.md#UpdateAccountScopedEnvironment) | **Put** /v1/environments/{environment} | Update account scoped Environment

# **CreateAccountScopedEnvironment**
> EnvironmentResponse CreateAccountScopedEnvironment(ctx, body, optional)
Creates an account scoped Environment

Creates an account scoped Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentCreateRequest**](EnvironmentCreateRequest.md)| Update Environment request body | 
 **optional** | ***AccountEnvironmentsApiCreateAccountScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountEnvironmentsApiCreateAccountScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountScopedEnvironment**
> DeleteAccountScopedEnvironment(ctx, environment, optional)
Delete an account scoped Environment

Deletes the requested account scoped environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***AccountEnvironmentsApiDeleteAccountScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountEnvironmentsApiDeleteAccountScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **forceDelete** | **optional.Bool**| Enable this field to force delete the entity | [default to false]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedEnvironment**
> EnvironmentResponse GetAccountScopedEnvironment(ctx, environment, optional)
Retrieve an account scoped environment

Retrieves the specified account scoped environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***AccountEnvironmentsApiGetAccountScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountEnvironmentsApiGetAccountScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedEnvironments**
> []EnvironmentResponse GetAccountScopedEnvironments(ctx, optional)
List account scoped Environments

Returns a list of the account scoped environments for which you have view permissions in the given account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountEnvironmentsApiGetAccountScopedEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountEnvironmentsApiGetAccountScopedEnvironmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **environmentIds** | [**optional.Interface of []string**](string.md)| List of Environment Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list account scoped of environments that you are permitted to use in a pipeline. | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountScopedEnvironment**
> EnvironmentResponse UpdateAccountScopedEnvironment(ctx, body, environment, optional)
Update account scoped Environment

Updates the specified account scoped environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentUpdateRequest**](EnvironmentUpdateRequest.md)| Update Environment request body | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***AccountEnvironmentsApiUpdateAccountScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountEnvironmentsApiUpdateAccountScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

