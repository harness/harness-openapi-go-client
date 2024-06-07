# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReleaseDetails**](K8sReleaseServiceMappingApi.md#GetV1ReleaseDetails) | **Get** /v1/kubernetes/releases/service-mapping | List service and environment details using namespace and releasename

# **GetV1ReleaseDetails**
> [][]ReleaseDetailsResponse GetV1ReleaseDetails(ctx, optional)
List service and environment details using namespace and releasename

Return details of service and environment mapped to pods namespace and release-name for a given account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***K8sReleaseServiceMappingApiGetV1ReleaseDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a K8sReleaseServiceMappingApiGetV1ReleaseDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []ReleaseDetailsRequest**](ReleaseDetailsRequest.md)| Batch Release Details request body | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**[][]ReleaseDetailsResponse**](array.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

