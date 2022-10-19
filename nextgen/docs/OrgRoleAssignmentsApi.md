# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedRoleAssignments**](OrgRoleAssignmentsApi.md#CreateOrgScopedRoleAssignments) | **Post** /v1/orgs/{org}/roleassignments | Create a role assignment
[**DeleteOrgScopedRoleAssignment**](OrgRoleAssignmentsApi.md#DeleteOrgScopedRoleAssignment) | **Delete** /v1/orgs/{org}/roleassignments/{roleassignment} | Delete a role assignment
[**GetOrgScopedRoleAssignment**](OrgRoleAssignmentsApi.md#GetOrgScopedRoleAssignment) | **Get** /v1/orgs/{org}/roleassignments/{roleassignment} | Retrieve a role assignment
[**GetOrgScopedRoleAssignments**](OrgRoleAssignmentsApi.md#GetOrgScopedRoleAssignments) | **Get** /v1/orgs/{org}/roleassignments | List role assignments

# **CreateOrgScopedRoleAssignments**
> RoleAssignmentResponse CreateOrgScopedRoleAssignments(ctx, body, org, optional)
Create a role assignment

Create a role assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleAssignment**](RoleAssignment.md)| Role Request body | 
  **org** | **string**| Organization identifier | 
 **optional** | ***OrgRoleAssignmentsApiCreateOrgScopedRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgRoleAssignmentsApiCreateOrgScopedRoleAssignmentsOpts struct
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

# **DeleteOrgScopedRoleAssignment**
> RoleAssignmentResponse DeleteOrgScopedRoleAssignment(ctx, roleassignment, org, optional)
Delete a role assignment

Deletes the information of the role assignment with the matching role assignment slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleassignment** | **string**| Role assignment slug | 
  **org** | **string**| Organization identifier | 
 **optional** | ***OrgRoleAssignmentsApiDeleteOrgScopedRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgRoleAssignmentsApiDeleteOrgScopedRoleAssignmentOpts struct
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

# **GetOrgScopedRoleAssignment**
> RoleAssignmentResponse GetOrgScopedRoleAssignment(ctx, roleassignment, org, optional)
Retrieve a role assignment

Retrieves the information of the role assignment with the matching role assignment slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleassignment** | **string**| Role assignment slug | 
  **org** | **string**| Organization identifier | 
 **optional** | ***OrgRoleAssignmentsApiGetOrgScopedRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgRoleAssignmentsApiGetOrgScopedRoleAssignmentOpts struct
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

# **GetOrgScopedRoleAssignments**
> []RoleAssignmentResponse GetOrgScopedRoleAssignments(ctx, org, optional)
List role assignments

Retrieves the information of the role assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
 **optional** | ***OrgRoleAssignmentsApiGetOrgScopedRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgRoleAssignmentsApiGetOrgScopedRoleAssignmentsOpts struct
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

