# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleAcc**](AccountRolesApi.md#CreateRoleAcc) | **Post** /v1/roles | Create a Role
[**DeleteRoleAcc**](AccountRolesApi.md#DeleteRoleAcc) | **Delete** /v1/roles/{role} | Delete a Role
[**GetRoleAcc**](AccountRolesApi.md#GetRoleAcc) | **Get** /v1/roles/{role} | Retrieve a Role
[**ListRolesAcc**](AccountRolesApi.md#ListRolesAcc) | **Get** /v1/roles | List Roles
[**UpdateRoleAcc**](AccountRolesApi.md#UpdateRoleAcc) | **Put** /v1/roles/{role} | Update a Role

# **CreateRoleAcc**
> RolesResponse CreateRoleAcc(ctx, body, optional)
Create a Role

Creates a custom Role in the Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleRequest**](CreateRoleRequest.md)| Role Request body | 
 **optional** | ***AccountRolesApiCreateRoleAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRolesApiCreateRoleAccOpts struct
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

# **DeleteRoleAcc**
> RolesResponse DeleteRoleAcc(ctx, role, optional)
Delete a Role

Deletes a custom Role from Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| Role identifier | 
 **optional** | ***AccountRolesApiDeleteRoleAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRolesApiDeleteRoleAccOpts struct
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

# **GetRoleAcc**
> RolesResponse GetRoleAcc(ctx, role, optional)
Retrieve a Role

Retrieves a Role from Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **role** | **string**| Role identifier | 
 **optional** | ***AccountRolesApiGetRoleAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRolesApiGetRoleAccOpts struct
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

# **ListRolesAcc**
> []RolesResponse ListRolesAcc(ctx, optional)
List Roles

Returns a list of Roles present in the Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountRolesApiListRolesAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRolesApiListRolesAccOpts struct
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

# **UpdateRoleAcc**
> RolesResponse UpdateRoleAcc(ctx, body, role, optional)
Update a Role

Updates a Role from Account scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRoleRequest**](CreateRoleRequest.md)| Role Request body | 
  **role** | **string**| Role identifier | 
 **optional** | ***AccountRolesApiUpdateRoleAccOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRolesApiUpdateRoleAccOpts struct
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

