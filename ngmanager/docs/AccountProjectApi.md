# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedProject**](AccountProjectApi.md#CreateAccountScopedProject) | **Post** /v1/projects | Create a project
[**DeleteAccountScopedProject**](AccountProjectApi.md#DeleteAccountScopedProject) | **Delete** /v1/projects/{project} | Delete a project
[**GetAccountScopedProject**](AccountProjectApi.md#GetAccountScopedProject) | **Get** /v1/projects/{project} | Retrieve a project
[**GetAccountScopedProjects**](AccountProjectApi.md#GetAccountScopedProjects) | **Get** /v1/projects | List projects
[**UpdateAccountScopedProject**](AccountProjectApi.md#UpdateAccountScopedProject) | **Put** /v1/projects/{project} | Update a project

# **CreateAccountScopedProject**
> ProjectResponse CreateAccountScopedProject(ctx, body, optional)
Create a project

Creates a new project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateProjectRequest**](CreateProjectRequest.md)| Post the necessary fields for the API to create a project. | 
 **optional** | ***AccountProjectApiCreateAccountScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountProjectApiCreateAccountScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountScopedProject**
> ProjectResponse DeleteAccountScopedProject(ctx, project, optional)
Delete a project

Deletes the information of the project with the matching project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project slug | 
 **optional** | ***AccountProjectApiDeleteAccountScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountProjectApiDeleteAccountScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedProject**
> ProjectResponse GetAccountScopedProject(ctx, project, optional)
Retrieve a project

Retrieves the information of the project with the matching project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**| Project slug | 
 **optional** | ***AccountProjectApiGetAccountScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountProjectApiGetAccountScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountScopedProjects**
> []ProjectResponse GetAccountScopedProjects(ctx, optional)
List projects

Retrieves the information of the projects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountProjectApiGetAccountScopedProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountProjectApiGetAccountScopedProjectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 
 **org** | [**optional.Interface of []string**](string.md)| Slug field of the organizations the resource is scoped to | 
 **project** | [**optional.Interface of []string**](string.md)| Slug field of the projects the resource is scoped to | 
 **hasModule** | **optional.Bool**| This boolean specifies whether to filter projects which has the module of type passed in the moduleType parameter or not | [default to true]
 **moduleType** | **optional.String**| Project&#x27;s module type | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]

### Return type

[**[]ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountScopedProject**
> ProjectResponse UpdateAccountScopedProject(ctx, body, project, optional)
Update a project

Updates the information of the project with the matching project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateProjectRequest**](UpdateProjectRequest.md)| Put the necessary fields for the API to update a Project. | 
  **project** | **string**| Project slug | 
 **optional** | ***AccountProjectApiUpdateAccountScopedProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountProjectApiUpdateAccountScopedProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **account** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization method other than x-api-key header. If you are using x-api-key header this can be skipped. | 

### Return type

[**ProjectResponse**](ProjectResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

