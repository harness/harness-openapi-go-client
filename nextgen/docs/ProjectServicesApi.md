# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ProjectServicesApi.md#CreateService) | **Post** /v1/orgs/{org}/projects/{project}/services | Create a Service
[**DeleteService**](ProjectServicesApi.md#DeleteService) | **Delete** /v1/orgs/{org}/projects/{project}/services/{service} | Delete a Service
[**GetService**](ProjectServicesApi.md#GetService) | **Get** /v1/orgs/{org}/projects/{project}/services/{service} | Retrieve a Service
[**GetServices**](ProjectServicesApi.md#GetServices) | **Get** /v1/orgs/{org}/projects/{project}/services | List Services
[**UpdateService**](ProjectServicesApi.md#UpdateService) | **Put** /v1/orgs/{org}/projects/{project}/services/{service} | Update Service

# **CreateService**
> ServiceResponse CreateService(ctx, body, org, project, optional)
Create a Service

Creates a Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceRequest**](ServiceRequest.md)| Create Service request body | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***ProjectServicesApiCreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServicesApiCreateServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteService**
> ServiceResponse DeleteService(ctx, org, project, service, optional)
Delete a Service

Deletes the requested service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectServicesApiDeleteServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServicesApiDeleteServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetService**
> ServiceResponse GetService(ctx, org, project, service, optional)
Retrieve a Service

Retrieves specified Service 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectServicesApiGetServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServicesApiGetServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServices**
> []ServiceResponse GetServices(ctx, org, project, optional)
List Services

Returns all services that the current user has view permissions for given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***ProjectServicesApiGetServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServicesApiGetServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **serviceIds** | [**optional.Interface of []string**](string.md)| List of Service Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specify whether list an access list or not. Access List refers to list of services that current user has permission to use in the pipeline. | 
 **type_** | **optional.String**| Service Definition type | 
 **gitOpsEnabled** | **optional.Bool**| Enables use of this service in Harness GitOps PR Pipelines | 
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateService**
> ServiceResponse UpdateService(ctx, body, org, project, service, optional)
Update Service

Update Services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceRequest**](ServiceRequest.md)| Create Service request body | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **service** | **string**| Slug field of the service the resource is scoped to | 
 **optional** | ***ProjectServicesApiUpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServicesApiUpdateServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

