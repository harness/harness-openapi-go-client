# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedRoleAssignments**](AccountRoleAssignmentsApi.md#CreateAccountScopedRoleAssignments) | **Post** /v1/roleassignments | Create a role assignment
[**DeleteAccountScopedRoleAssignment**](AccountRoleAssignmentsApi.md#DeleteAccountScopedRoleAssignment) | **Delete** /v1/roleassignments/{roleassignment} | Delete a role assignment
[**GetAccountScopedRoleAssignment**](AccountRoleAssignmentsApi.md#GetAccountScopedRoleAssignment) | **Get** /v1/roleassignments/{roleassignment} | Retrieve a role assignment
[**GetAccountScopedRoleAssignments**](AccountRoleAssignmentsApi.md#GetAccountScopedRoleAssignments) | **Get** /v1/roleassignments | List role assignments

# **CreateAccountScopedRoleAssignments**
> RoleAssignmentResponse CreateAccountScopedRoleAssignments(ctx, body, optional)
Create a role assignment

Create a role assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleAssignment**](RoleAssignment.md)| Role Request body | 
 **optional** | ***AccountRoleAssignmentsApiCreateAccountScopedRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRoleAssignmentsApiCreateAccountScopedRoleAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RoleAssignmentResponse**](RoleAssignmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountScopedRoleAssignment**
> RoleAssignmentResponse DeleteAccountScopedRoleAssignment(ctx, roleassignment, optional)
Delete a role assignment

Deletes the information of the role assignment with the matching role assignment slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleassignment** | **string**| Role assignment slug | 
 **optional** | ***AccountRoleAssignmentsApiDeleteAccountScopedRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRoleAssignmentsApiDeleteAccountScopedRoleAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RoleAssignmentResponse**](RoleAssignmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedRoleAssignment**
> RoleAssignmentResponse GetAccountScopedRoleAssignment(ctx, roleassignment, optional)
Retrieve a role assignment

Retrieves the information of the role assignment with the matching role assignment slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleassignment** | **string**| Role assignment slug | 
 **optional** | ***AccountRoleAssignmentsApiGetAccountScopedRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRoleAssignmentsApiGetAccountScopedRoleAssignmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**RoleAssignmentResponse**](RoleAssignmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedRoleAssignments**
> []RoleAssignmentResponse GetAccountScopedRoleAssignments(ctx, optional)
List role assignments

Retrieves the information of the role assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountRoleAssignmentsApiGetAccountScopedRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRoleAssignmentsApiGetAccountScopedRoleAssignmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items on each page. | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return. | [default to 30]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]RoleAssignmentResponse**](RoleAssignmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

