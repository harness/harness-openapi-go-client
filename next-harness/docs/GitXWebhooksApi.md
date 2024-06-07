# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGitxWebhook**](GitXWebhooksApi.md#CreateGitxWebhook) | **Post** /v1/gitx-webhooks | Create a GitX webhook at Account level
[**DeleteGitxWebhook**](GitXWebhooksApi.md#DeleteGitxWebhook) | **Delete** /v1/gitx-webhooks/{gitx-webhook} | Deletes a GitX Webhook at Account level
[**GetGitxWebhook**](GitXWebhooksApi.md#GetGitxWebhook) | **Get** /v1/gitx-webhooks/{gitx-webhook} | Fetch GitX Webhook at Account level
[**ListGitxWebhookEvents**](GitXWebhooksApi.md#ListGitxWebhookEvents) | **Get** /v1/gitx-webhook-events | Lists all the GitX Webhooks events at Account level
[**ListGitxWebhooks**](GitXWebhooksApi.md#ListGitxWebhooks) | **Get** /v1/gitx-webhooks | Lists all the GitX Webhooks at Account level
[**UpdateGitxWebhook**](GitXWebhooksApi.md#UpdateGitxWebhook) | **Put** /v1/gitx-webhooks/{gitx-webhook} | Updates a GitX Webhook at Account level

# **CreateGitxWebhook**
> CreateGitXWebhookResponse CreateGitxWebhook(ctx, optional)
Create a GitX webhook at Account level

Create GitXWebhook at account level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitXWebhooksApiCreateGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksApiCreateGitxWebhookOpts struct
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

# **DeleteGitxWebhook**
> DeleteGitXWebhookResponse DeleteGitxWebhook(ctx, gitxWebhook, optional)
Deletes a GitX Webhook at Account level

Deletes a gitx webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***GitXWebhooksApiDeleteGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksApiDeleteGitxWebhookOpts struct
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

# **GetGitxWebhook**
> GitXWebhookResponse GetGitxWebhook(ctx, gitxWebhook, optional)
Fetch GitX Webhook at Account level

Fetch a gitx webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***GitXWebhooksApiGetGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksApiGetGitxWebhookOpts struct
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

# **ListGitxWebhookEvents**
> []GitXWebhookEventResponse ListGitxWebhookEvents(ctx, optional)
Lists all the GitX Webhooks events at Account level

List GitX webhook Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitXWebhooksApiListGitxWebhookEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksApiListGitxWebhookEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **webhookIdentifier** | **optional.String**|  | 
 **eventStartTime** | **optional.Int64**|  | 
 **eventEndTime** | **optional.Int64**|  | 
 **repoName** | **optional.String**|  | 
 **filePath** | **optional.String**|  | 
 **eventIdentifier** | **optional.String**|  | 
 **eventStatus** | [**optional.Interface of []string**](string.md)|  | 
 **connectorRef** | **optional.String**|  | 
 **includeParentScope** | **optional.Bool**|  | 
 **commitId** | **optional.String**|  | 

### Return type

[**[]GitXWebhookEventResponse**](GitXWebhookEventResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGitxWebhooks**
> []GitXWebhookResponse ListGitxWebhooks(ctx, optional)
Lists all the GitX Webhooks at Account level

List GitX webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GitXWebhooksApiListGitxWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksApiListGitxWebhooksOpts struct
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

# **UpdateGitxWebhook**
> UpdateGitXWebhookResponse UpdateGitxWebhook(ctx, gitxWebhook, optional)
Updates a GitX Webhook at Account level

Update a Gitx webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gitxWebhook** | **string**| GitX webhook identifier | 
 **optional** | ***GitXWebhooksApiUpdateGitxWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksApiUpdateGitxWebhookOpts struct
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

