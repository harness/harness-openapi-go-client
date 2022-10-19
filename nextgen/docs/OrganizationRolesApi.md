# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleOrg**](OrganizationRolesApi.md#CreateRoleOrg) | **Post** /v1/orgs/{org}/roles | Create a Role
[**DeleteRoleOrg**](OrganizationRolesApi.md#DeleteRoleOrg) | **Delete** /v1/orgs/{org}/roles/{role} | Delete a Role
[**GetRoleOrg**](OrganizationRolesApi.md#GetRoleOrg) | **Get** /v1/orgs/{org}/roles/{role} | Retrieve a Role
[**ListRolesOrg**](OrganizationRolesApi.md#ListRolesOrg) | **Get** /v1/orgs/{org}/roles | List Roles
[**UpdateRoleOrg**](OrganizationRolesApi.md#UpdateRoleOrg) | **Put** /v1/orgs/{org}/roles/{role} | Update a Role

# **CreateRoleOrg**
> RolesResponse CreateRoleOrg(ctx, body, org, optional)
Create a Role

Creates a custom Role in the Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleRequest**](CreateRoleRequest.md)| Role Request body | 
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationRolesApiCreateRoleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationRolesApiCreateRoleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleOrg**
> RolesResponse DeleteRoleOrg(ctx, org, role, optional)
Delete a Role

Deletes a custom Role from Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **role** | **string**| Role identifier | 
 **optional** | ***OrganizationRolesApiDeleteRoleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationRolesApiDeleteRoleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoleOrg**
> RolesResponse GetRoleOrg(ctx, org, role, optional)
Retrieve a Role

Retrieves a Role from Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **role** | **string**| Role identifier | 
 **optional** | ***OrganizationRolesApiGetRoleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationRolesApiGetRoleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRolesOrg**
> []RolesResponse ListRolesOrg(ctx, org, optional)
List Roles

Returns a list of Roles present in the Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationRolesApiListRolesOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationRolesApiListRolesOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return. | [default to 30]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching the search term. | 
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]RolesResponse**](RolesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoleOrg**
> RolesResponse UpdateRoleOrg(ctx, body, org, role, optional)
Update a Role

Updates a Role from Organization scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleRequest**](CreateRoleRequest.md)| Role Request body | 
  **org** | **string**| Organization identifier | 
  **role** | **string**| Role identifier | 
 **optional** | ***OrganizationRolesApiUpdateRoleOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationRolesApiUpdateRoleOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RolesResponse**](RolesResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

