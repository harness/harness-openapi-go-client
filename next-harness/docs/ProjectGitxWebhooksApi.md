# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectGitxWebhook**](ProjectGitxWebhooksApi.md#CreateProjectGitxWebhook) | **Post** /v1/orgs/{org}/projects/{project}/gitx-webhooks | Create Project level GitX webhook
[**DeleteProjectGitxWebhook**](ProjectGitxWebhooksApi.md#DeleteProjectGitxWebhook) | **Delete** /v1/orgs/{org}/projects/{project}/gitx-webhooks/{gitx-webhook} | Deletes a GitX Webhook at project level
[**GetProjectGitxWebhook**](ProjectGitxWebhooksApi.md#GetProjectGitxWebhook) | **Get** /v1/orgs/{org}/projects/{project}/gitx-webhooks/{gitx-webhook} | Fetch GitX Webhook at project level
[**ListProjectGitxWebhook**](ProjectGitxWebhooksApi.md#ListProjectGitxWebhook) | **Get** /v1/orgs/{org}/projects/{project}/gitx-webhooks | Lists all the GitX Webhooks at project level
[**UpdateProjectGitxWebhook**](ProjectGitxWebhooksApi.md#UpdateProjectGitxWebhook) | **Put** /v1/orgs/{org}/projects/{project}/gitx-webhooks/{gitx-webhook} | Updates a GitX Webhook at project level

# **CreateProjectGitxWebhook**
> CreateGitXWebhookResponse CreateProjectGitxWebhook(ctx, org, project, optional)
Create Project level GitX webhook

Create GitXWebhook at project level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectGitxWebhooksApiCreateProjectGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectGitxWebhooksApiCreateProjectGitxWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CreateGitXWebhookRequest**](CreateGitXWebhookRequest.md)| Create GitX webhook request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**CreateGitXWebhookResponse**](CreateGitXWebhookResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectGitxWebhook**
> DeleteGitXWebhookResponse DeleteProjectGitxWebhook(ctx, org, project, gitxWebhook, optional)
Deletes a GitX Webhook at project level

Deletes a project level gitx webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***ProjectGitxWebhooksApiDeleteProjectGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectGitxWebhooksApiDeleteProjectGitxWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**DeleteGitXWebhookResponse**](DeleteGitXWebhookResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectGitxWebhook**
> GitXWebhookResponse GetProjectGitxWebhook(ctx, org, project, gitxWebhook, optional)
Fetch GitX Webhook at project level

Fetch a project level gitx webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***ProjectGitxWebhooksApiGetProjectGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectGitxWebhooksApiGetProjectGitxWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**GitXWebhookResponse**](GitXWebhookResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjectGitxWebhook**
> []GitXWebhookResponse ListProjectGitxWebhook(ctx, org, project, optional)
Lists all the GitX Webhooks at project level

List project level GitX webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectGitxWebhooksApiListProjectGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectGitxWebhooksApiListProjectGitxWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **webhookIdentifier** | **optional.String**|  | 

### Return type

[**[]GitXWebhookResponse**](GitXWebhookResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectGitxWebhook**
> UpdateGitXWebhookResponse UpdateProjectGitxWebhook(ctx, org, project, gitxWebhook, optional)
Updates a GitX Webhook at project level

Update a project level Gitx webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***ProjectGitxWebhooksApiUpdateProjectGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectGitxWebhooksApiUpdateProjectGitxWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UpdateGitXWebhookRequest**](UpdateGitXWebhookRequest.md)| Update GitX webhook request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UpdateGitXWebhookResponse**](UpdateGitXWebhookResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

