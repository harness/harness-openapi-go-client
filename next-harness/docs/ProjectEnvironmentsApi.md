# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](ProjectEnvironmentsApi.md#CreateEnvironment) | **Post** /v1/orgs/{org}/projects/{project}/environments | Creates an Environment
[**DeleteEnvironment**](ProjectEnvironmentsApi.md#DeleteEnvironment) | **Delete** /v1/orgs/{org}/projects/{project}/environments/{environment} | Delete a Environment
[**GetEnvironment**](ProjectEnvironmentsApi.md#GetEnvironment) | **Get** /v1/orgs/{org}/projects/{project}/environments/{environment} | Retrieve a environment
[**GetEnvironments**](ProjectEnvironmentsApi.md#GetEnvironments) | **Get** /v1/orgs/{org}/projects/{project}/environments | List Environments
[**UpdateEnvironment**](ProjectEnvironmentsApi.md#UpdateEnvironment) | **Put** /v1/orgs/{org}/projects/{project}/environments/{environment} | Update Environment

# **CreateEnvironment**
> EnvironmentResponse CreateEnvironment(ctx, body, org, project, optional)
Creates an Environment

Creates an Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentCreateRequest**](EnvironmentCreateRequest.md)| Update Environment request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiCreateEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiCreateEnvironmentOpts struct
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

# **DeleteEnvironment**
> DeleteEnvironment(ctx, org, project, environment, optional)
Delete a Environment

Deletes the requested environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***ProjectEnvironmentsApiDeleteEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiDeleteEnvironmentOpts struct
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

# **GetEnvironment**
> EnvironmentResponse GetEnvironment(ctx, org, project, environment, optional)
Retrieve a environment

Retrieves the specified environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***ProjectEnvironmentsApiGetEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiGetEnvironmentOpts struct
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

# **GetEnvironments**
> []EnvironmentResponse GetEnvironments(ctx, org, project, optional)
List Environments

Returns a list of the enviornments for which you have view permissions in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectEnvironmentsApiGetEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiGetEnvironmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **environmentIds** | [**optional.Interface of []string**](string.md)| List of Environment Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list of environments that you are permitted to use in a pipeline. | 
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

# **UpdateEnvironment**
> EnvironmentResponse UpdateEnvironment(ctx, body, org, project, environment, optional)
Update Environment

Updates the specified environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentUpdateRequest**](EnvironmentUpdateRequest.md)| Update Environment request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***ProjectEnvironmentsApiUpdateEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectEnvironmentsApiUpdateEnvironmentOpts struct
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

