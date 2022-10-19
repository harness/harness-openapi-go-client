# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilterResourceGroups**](FilterResourceGroupsApi.md#FilterResourceGroups) | **Post** /v1/resource-groups/filter | Filter Resource Groups

# **FilterResourceGroups**
> []ResourceGroupsResponse FilterResourceGroups(ctx, body, optional)
Filter Resource Groups

Returns a list of Resource Groups based on filter criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceGroupFilterRequestBody**](ResourceGroupFilterRequestBody.md)| Filter Resource Group request body | 
 **optional** | ***FilterResourceGroupsApiFilterResourceGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilterResourceGroupsApiFilterResourceGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.**| Pagination: Number of items to return. | [default to 30]
 **sort** | **optional.**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.**| Order on the basis of which sorting is done. | 

### Return type

[**[]ResourceGroupsResponse**](ResourceGroupsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

