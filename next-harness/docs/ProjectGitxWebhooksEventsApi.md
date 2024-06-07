# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProjectGitxWebhookEvents**](ProjectGitxWebhooksEventsApi.md#ListProjectGitxWebhookEvents) | **Get** /v1/orgs/{org}/projects/{project}/gitx-webhook-events | Lists all the GitX Webhooks events at project level

# **ListProjectGitxWebhookEvents**
> []GitXWebhookEventResponse ListProjectGitxWebhookEvents(ctx, org, project, optional)
Lists all the GitX Webhooks events at project level

List GitX webhook Events at project level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectGitxWebhooksEventsApiListProjectGitxWebhookEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectGitxWebhooksEventsApiListProjectGitxWebhookEventsOpts struct
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

