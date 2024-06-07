# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GitxWebhookEventFileValidation**](GitXWebhooksEventFileValidationApi.md#GitxWebhookEventFileValidation) | **Get** /v1/gitx-webhook-events/{gitx-webhook-event}/validation-info | Get file validations for a given event identifier

# **GitxWebhookEventFileValidation**
> []GetGitXWebhookEventFileValidationResponse GitxWebhookEventFileValidation(ctx, gitxWebhookEvent, optional)
Get file validations for a given event identifier

Get file validations for a given event identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gitxWebhookEvent** | **string**|  | 
 **optional** | ***GitXWebhooksEventFileValidationApiGitxWebhookEventFileValidationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GitXWebhooksEventFileValidationApiGitxWebhookEventFileValidationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]

### Return type

[**[]GetGitXWebhookEventFileValidationResponse**](GetGitXWebhookEventFileValidationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

