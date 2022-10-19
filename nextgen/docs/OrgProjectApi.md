# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedProject**](OrgProjectApi.md#CreateOrgScopedProject) | **Post** /v1/orgs/{org}/projects | Creates a project
[**DeleteOrgScopedProject**](OrgProjectApi.md#DeleteOrgScopedProject) | **Delete** /v1/orgs/{org}/projects/{project} | Delete a project
[**GetOrgScopedProject**](OrgProjectApi.md#GetOrgScopedProject) | **Get** /v1/orgs/{org}/projects/{project} | Retrieve a project
[**GetOrgScopedProjects**](OrgProjectApi.md#GetOrgScopedProjects) | **Get** /v1/orgs/{org}/projects | List projects
[**UpdateOrgScopedProject**](OrgProjectApi.md#UpdateOrgScopedProject) | **Put** /v1/orgs/{org}/projects/{project} | Update a project

# **CreateOrgScopedProject**
> ProjectResponse CreateOrgScopedProject(ctx, body, org, optional)
Creates a project

Creates a new project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProjectRequest**](CreateProjectRequest.md)| Post the necessary fields for the API to create a project. | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
 **optional** | ***OrgProjectApiCreateOrgScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgProjectApiCreateOrgScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgScopedProject**
> ProjectResponse DeleteOrgScopedProject(ctx, org, project, optional)
Delete a project

Deletes the information of the project with the matching project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***OrgProjectApiDeleteOrgScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgProjectApiDeleteOrgScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedProject**
> ProjectResponse GetOrgScopedProject(ctx, org, project, optional)
Retrieve a project

Retrieves the information of the project with the matching project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***OrgProjectApiGetOrgScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgProjectApiGetOrgScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedProjects**
> []ProjectResponse GetOrgScopedProjects(ctx, org, optional)
List projects

Retrieves the information of the projects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
 **optional** | ***OrgProjectApiGetOrgScopedProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgProjectApiGetOrgScopedProjectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**optional.Interface of []string**](string.md)| Slug field of the projects the resource is scoped to | 
 **hasModule** | **optional.Bool**| This boolean specifies whether to filter projects which has the module of type passed in the moduleType parameter or not | [default to true]
 **moduleType** | **optional.String**| Project&#x27;s module type | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedProject**
> ProjectResponse UpdateOrgScopedProject(ctx, body, org, project, optional)
Update a project

Updates the information of the project with the matching project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProjectRequest**](UpdateProjectRequest.md)| Put the necessary fields for the API to update a Project. | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***OrgProjectApiUpdateOrgScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgProjectApiUpdateOrgScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

