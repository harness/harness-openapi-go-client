# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchGitxWebhookEventsGitxWebhookEvent**](GitXWebhookEventApi.md#PatchGitxWebhookEventsGitxWebhookEvent) | **Patch** /v1/gitx-webhook-events/{gitx-webhook-event} | Update the Webhook Event status

# **PatchGitxWebhookEventsGitxWebhookEvent**
> GitXWebhookEventResponse PatchGitxWebhookEventsGitxWebhookEvent(ctx, gitxWebhookEvent, optional)
Update the Webhook Event status

Update the webhook event Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gitxWebhookEvent** | **string**| GitX Webhook Event Identifier | 
 **optional** | ***GitXWebhookEventApiPatchGitxWebhookEventsGitxWebhookEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhookEventApiPatchGitxWebhookEventsGitxWebhookEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateGitXWebhookEventRequest**](UpdateGitXWebhookEventRequest.md)| Update GitX Webhook Event Request | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**GitXWebhookEventResponse**](GitXWebhookEventResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

