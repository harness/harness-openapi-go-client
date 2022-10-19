# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceGroupAcc**](AccountResourceGroupsApi.md#CreateResourceGroupAcc) | **Post** /v1/resource-groups | Create a Resource Group
[**DeleteResourceGroupAcc**](AccountResourceGroupsApi.md#DeleteResourceGroupAcc) | **Delete** /v1/resource-groups/{resource-group} | Delete a Resource Group
[**GetResourceGroupAcc**](AccountResourceGroupsApi.md#GetResourceGroupAcc) | **Get** /v1/resource-groups/{resource-group} | Retrieve a Resource Group
[**ListResourceGroupsAcc**](AccountResourceGroupsApi.md#ListResourceGroupsAcc) | **Get** /v1/resource-groups | List Resource Groups
[**UpdateResourceGroupAcc**](AccountResourceGroupsApi.md#UpdateResourceGroupAcc) | **Put** /v1/resource-groups/{resource-group} | Update a Resource Group

# **CreateResourceGroupAcc**
> ResourceGroupsResponse CreateResourceGroupAcc(ctx, body, optional)
Create a Resource Group

Creates a custom Resource Group in the Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceGroupRequest**](CreateResourceGroupRequest.md)| Resource Group request body | 
 **optional** | ***AccountResourceGroupsApiCreateResourceGroupAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountResourceGroupsApiCreateResourceGroupAccOpts struct
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

# **DeleteResourceGroupAcc**
> ResourceGroupsResponse DeleteResourceGroupAcc(ctx, resourceGroup, optional)
Delete a Resource Group

Deletes a custom Resource Group from Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***AccountResourceGroupsApiDeleteResourceGroupAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountResourceGroupsApiDeleteResourceGroupAccOpts struct
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

# **GetResourceGroupAcc**
> ResourceGroupsResponse GetResourceGroupAcc(ctx, resourceGroup, optional)
Retrieve a Resource Group

Retrieves a Resource Group from Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***AccountResourceGroupsApiGetResourceGroupAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountResourceGroupsApiGetResourceGroupAccOpts struct
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

# **ListResourceGroupsAcc**
> []ResourceGroupsResponse ListResourceGroupsAcc(ctx, optional)
List Resource Groups

Returns a list of Resource Groups present in the Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountResourceGroupsApiListResourceGroupsAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountResourceGroupsApiListResourceGroupsAccOpts struct
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

# **UpdateResourceGroupAcc**
> ResourceGroupsResponse UpdateResourceGroupAcc(ctx, body, resourceGroup, optional)
Update a Resource Group

Updates a Resource Group from Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceGroupRequest**](CreateResourceGroupRequest.md)| Resource Group request body | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***AccountResourceGroupsApiUpdateResourceGroupAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountResourceGroupsApiUpdateResourceGroupAccOpts struct
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

