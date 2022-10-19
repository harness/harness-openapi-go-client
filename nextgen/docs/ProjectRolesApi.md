# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleProject**](ProjectRolesApi.md#CreateRoleProject) | **Post** /v1/orgs/{org}/projects/{project}/roles | Create a Role
[**DeleteRoleProject**](ProjectRolesApi.md#DeleteRoleProject) | **Delete** /v1/orgs/{org}/projects/{project}/roles/{role} | Delete a Role
[**GetRoleProject**](ProjectRolesApi.md#GetRoleProject) | **Get** /v1/orgs/{org}/projects/{project}/roles/{role} | Retrieve a Role
[**ListRolesProject**](ProjectRolesApi.md#ListRolesProject) | **Get** /v1/orgs/{org}/projects/{project}/roles | List Roles
[**UpdateRoleProject**](ProjectRolesApi.md#UpdateRoleProject) | **Put** /v1/orgs/{org}/projects/{project}/roles/{role} | Update a Role

# **CreateRoleProject**
> RolesResponse CreateRoleProject(ctx, body, org, project, optional)
Create a Role

Creates a custom Role in the Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleRequest**](CreateRoleRequest.md)| Role Request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectRolesApiCreateRoleProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRolesApiCreateRoleProjectOpts struct
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

# **DeleteRoleProject**
> RolesResponse DeleteRoleProject(ctx, org, project, role, optional)
Delete a Role

Deletes a custom Role from Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **role** | **string**| Role identifier | 
 **optional** | ***ProjectRolesApiDeleteRoleProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRolesApiDeleteRoleProjectOpts struct
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

# **GetRoleProject**
> RolesResponse GetRoleProject(ctx, org, project, role, optional)
Retrieve a Role

Retrieves a Role from Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **role** | **string**| Role identifier | 
 **optional** | ***ProjectRolesApiGetRoleProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRolesApiGetRoleProjectOpts struct
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

# **ListRolesProject**
> []RolesResponse ListRolesProject(ctx, org, project, optional)
List Roles

Returns a list of Roles present in the Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectRolesApiListRolesProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRolesApiListRolesProjectOpts struct
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

# **UpdateRoleProject**
> RolesResponse UpdateRoleProject(ctx, body, org, project, role, optional)
Update a Role

Updates a Role from Project scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleRequest**](CreateRoleRequest.md)| Role Request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
  **role** | **string**| Role identifier | 
 **optional** | ***ProjectRolesApiUpdateRoleProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRolesApiUpdateRoleProjectOpts struct
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

