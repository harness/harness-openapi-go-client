# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectScopedRoleAssignments**](ProjectRoleAssignmentsApi.md#CreateProjectScopedRoleAssignments) | **Post** /v1/orgs/{org}/projects/{project}/roleassignments | Create a role assignment
[**DeleteProjectScopedRoleAssignment**](ProjectRoleAssignmentsApi.md#DeleteProjectScopedRoleAssignment) | **Delete** /v1/orgs/{org}/projects/{project}/roleassignments/{roleassignment} | Delete a role assignment
[**GetProjectScopedRoleAssignment**](ProjectRoleAssignmentsApi.md#GetProjectScopedRoleAssignment) | **Get** /v1/orgs/{org}/projects/{project}/roleassignments/{roleassignment} | Retrieve a role assignment
[**GetProjectScopedRoleAssignments**](ProjectRoleAssignmentsApi.md#GetProjectScopedRoleAssignments) | **Get** /v1/orgs/{org}/projects/{project}/roleassignments | List role assignments

# **CreateProjectScopedRoleAssignments**
> RoleAssignmentResponse CreateProjectScopedRoleAssignments(ctx, body, org, project, optional)
Create a role assignment

Create a role assignment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoleAssignment**](RoleAssignment.md)| Role Request body | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectRoleAssignmentsApiCreateProjectScopedRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRoleAssignmentsApiCreateProjectScopedRoleAssignmentsOpts struct
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

# **DeleteProjectScopedRoleAssignment**
> RoleAssignmentResponse DeleteProjectScopedRoleAssignment(ctx, roleassignment, org, project, optional)
Delete a role assignment

Deletes the information of the role assignment with the matching role assignment slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleassignment** | **string**| Role assignment slug | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectRoleAssignmentsApiDeleteProjectScopedRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRoleAssignmentsApiDeleteProjectScopedRoleAssignmentOpts struct
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

# **GetProjectScopedRoleAssignment**
> RoleAssignmentResponse GetProjectScopedRoleAssignment(ctx, roleassignment, org, project, optional)
Retrieve a role assignment

Retrieves the information of the role assignment with the matching role assignment slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleassignment** | **string**| Role assignment slug | 
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentOpts struct
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

# **GetProjectScopedRoleAssignments**
> []RoleAssignmentResponse GetProjectScopedRoleAssignments(ctx, org, project, optional)
List role assignments

Retrieves the information of the role assignments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
  **project** | **string**| Project identifier | 
 **optional** | ***ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRoleAssignmentsApiGetProjectScopedRoleAssignmentsOpts struct
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

