# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceGroupOrg**](OrganizationResourceGroupsApi.md#CreateResourceGroupOrg) | **Post** /v1/orgs/{org}/resource-groups | Create a Resource Group
[**DeleteResourceGroupOrg**](OrganizationResourceGroupsApi.md#DeleteResourceGroupOrg) | **Delete** /v1/orgs/{org}/resource-groups/{resource-group} | Delete a Resource Group
[**GetResourceGroupOrg**](OrganizationResourceGroupsApi.md#GetResourceGroupOrg) | **Get** /v1/orgs/{org}/resource-groups/{resource-group} | Retrieve a Resource Group
[**ListResourceGroupsOrg**](OrganizationResourceGroupsApi.md#ListResourceGroupsOrg) | **Get** /v1/orgs/{org}/resource-groups | List Resource Groups
[**UpdateResourceGroupOrg**](OrganizationResourceGroupsApi.md#UpdateResourceGroupOrg) | **Put** /v1/orgs/{org}/resource-groups/{resource-group} | Update a Resource Group

# **CreateResourceGroupOrg**
> ResourceGroupsResponse CreateResourceGroupOrg(ctx, body, org, optional)
Create a Resource Group

Creates a custom Resource Group in the Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceGroupRequest**](CreateResourceGroupRequest.md)| Resource Group request body | 
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationResourceGroupsApiCreateResourceGroupOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationResourceGroupsApiCreateResourceGroupOrgOpts struct
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

# **DeleteResourceGroupOrg**
> ResourceGroupsResponse DeleteResourceGroupOrg(ctx, org, resourceGroup, optional)
Delete a Resource Group

Deletes a custom Resource Group from Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***OrganizationResourceGroupsApiDeleteResourceGroupOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationResourceGroupsApiDeleteResourceGroupOrgOpts struct
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

# **GetResourceGroupOrg**
> ResourceGroupsResponse GetResourceGroupOrg(ctx, org, resourceGroup, optional)
Retrieve a Resource Group

Retrieves a Resource Group from Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***OrganizationResourceGroupsApiGetResourceGroupOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationResourceGroupsApiGetResourceGroupOrgOpts struct
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

# **ListResourceGroupsOrg**
> []ResourceGroupsResponse ListResourceGroupsOrg(ctx, org, optional)
List Resource Groups

Returns a list of Resource Groups present in the Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationResourceGroupsApiListResourceGroupsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationResourceGroupsApiListResourceGroupsOrgOpts struct
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

# **UpdateResourceGroupOrg**
> ResourceGroupsResponse UpdateResourceGroupOrg(ctx, body, org, resourceGroup, optional)
Update a Resource Group

Updates a Resource Group from Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceGroupRequest**](CreateResourceGroupRequest.md)| Resource Group request body | 
  **org** | **string**| Organization identifier | 
  **resourceGroup** | **string**| Resource Group identifier | 
 **optional** | ***OrganizationResourceGroupsApiUpdateResourceGroupOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationResourceGroupsApiUpdateResourceGroupOrgOpts struct
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

