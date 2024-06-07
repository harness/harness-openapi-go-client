# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgGitxWebhook**](OrgGitxWebhooksApi.md#CreateOrgGitxWebhook) | **Post** /v1/orgs/{org}/gitx-webhooks | Create Org Level GitX webhook
[**DeleteOrgGitxWebhook**](OrgGitxWebhooksApi.md#DeleteOrgGitxWebhook) | **Delete** /v1/orgs/{org}/gitx-webhooks/{gitx-webhook} | Deletes a GitX Webhook at org level
[**GetOrgGitxWebhook**](OrgGitxWebhooksApi.md#GetOrgGitxWebhook) | **Get** /v1/orgs/{org}/gitx-webhooks/{gitx-webhook} | Fetch GitX Webhook at org level
[**ListOrgGitxWebhooks**](OrgGitxWebhooksApi.md#ListOrgGitxWebhooks) | **Get** /v1/orgs/{org}/gitx-webhooks | Lists all the GitX Webhooks at Org level
[**UpdateOrgGitxWebhook**](OrgGitxWebhooksApi.md#UpdateOrgGitxWebhook) | **Put** /v1/orgs/{org}/gitx-webhooks/{gitx-webhook} | Updates a GitX Webhook at org level

# **CreateOrgGitxWebhook**
> CreateGitXWebhookResponse CreateOrgGitxWebhook(ctx, org, optional)
Create Org Level GitX webhook

Create GitXWebhook at org level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgGitxWebhooksApiCreateOrgGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgGitxWebhooksApiCreateOrgGitxWebhookOpts struct
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

# **DeleteOrgGitxWebhook**
> DeleteGitXWebhookResponse DeleteOrgGitxWebhook(ctx, org, gitxWebhook, optional)
Deletes a GitX Webhook at org level

Deletes a org level gitx webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***OrgGitxWebhooksApiDeleteOrgGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgGitxWebhooksApiDeleteOrgGitxWebhookOpts struct
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

# **GetOrgGitxWebhook**
> GitXWebhookResponse GetOrgGitxWebhook(ctx, org, gitxWebhook, optional)
Fetch GitX Webhook at org level

Fetch a org level gitx webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***OrgGitxWebhooksApiGetOrgGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgGitxWebhooksApiGetOrgGitxWebhookOpts struct
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

# **ListOrgGitxWebhooks**
> []GitXWebhookResponse ListOrgGitxWebhooks(ctx, org, optional)
Lists all the GitX Webhooks at Org level

List org level GitX webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgGitxWebhooksApiListOrgGitxWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgGitxWebhooksApiListOrgGitxWebhooksOpts struct
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

# **UpdateOrgGitxWebhook**
> UpdateGitXWebhookResponse UpdateOrgGitxWebhook(ctx, org, gitxWebhook, optional)
Updates a GitX Webhook at org level

Update a org level Gitx webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***OrgGitxWebhooksApiUpdateOrgGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgGitxWebhooksApiUpdateOrgGitxWebhookOpts struct
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

