# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceGroupProject**](ProjectResourceGroupsApi.md#CreateResourceGroupProject) | **Post** /v1/orgs/{org}/projects/{project}/resource-groups | Create a Resource Group
[**DeleteResourceGroupProject**](ProjectResourceGroupsApi.md#DeleteResourceGroupProject) | **Delete** /v1/orgs/{org}/projects/{project}/resource-groups/{resource-group} | Delete a Resource Group
[**GetResourceGroupProject**](ProjectResourceGroupsApi.md#GetResourceGroupProject) | **Get** /v1/orgs/{org}/projects/{project}/resource-groups/{resource-group} | Retrieve a Resource Group
[**ListResourceGroupsProject**](ProjectResourceGroupsApi.md#ListResourceGroupsProject) | **Get** /v1/orgs/{org}/projects/{project}/resource-groups | List Resource Groups
[**UpdateResourceGroupProject**](ProjectResourceGroupsApi.md#UpdateResourceGroupProject) | **Put** /v1/orgs/{org}/projects/{project}/resource-groups/{resource-group} | Update a Resource Group

# **CreateResourceGroupProject**
> ResourceGroupsResponse CreateResourceGroupProject(ctx, body, org, project, optional)
Create a Resource Group

Creates a custom Resource Group in the Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceGroupRequest**](CreateResourceGroupRequest.md)| Resource Group request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectResourceGroupsApiCreateResourceGroupProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectResourceGroupsApiCreateResourceGroupProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ResourceGroupsResponse**](ResourceGroupsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResourceGroupProject**
> ResourceGroupsResponse DeleteResourceGroupProject(ctx, org, project, resourceGroup, optional)
Delete a Resource Group

Deletes a custom Resource Group from Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***ProjectResourceGroupsApiDeleteResourceGroupProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectResourceGroupsApiDeleteResourceGroupProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ResourceGroupsResponse**](ResourceGroupsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourceGroupProject**
> ResourceGroupsResponse GetResourceGroupProject(ctx, org, project, resourceGroup, optional)
Retrieve a Resource Group

Retrieves a Resource Group from Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***ProjectResourceGroupsApiGetResourceGroupProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectResourceGroupsApiGetResourceGroupProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ResourceGroupsResponse**](ResourceGroupsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListResourceGroupsProject**
> []ResourceGroupsResponse ListResourceGroupsProject(ctx, org, project, optional)
List Resource Groups

Returns a list of Resource Groups present in the Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectResourceGroupsApiListResourceGroupsProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectResourceGroupsApiListResourceGroupsProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return. | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]ResourceGroupsResponse**](ResourceGroupsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceGroupProject**
> ResourceGroupsResponse UpdateResourceGroupProject(ctx, body, org, project, resourceGroup, optional)
Update a Resource Group

Updates a Resource Group from Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceGroupRequest**](CreateResourceGroupRequest.md)| Resource Group request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***ProjectResourceGroupsApiUpdateResourceGroupProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectResourceGroupsApiUpdateResourceGroupProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ResourceGroupsResponse**](ResourceGroupsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

