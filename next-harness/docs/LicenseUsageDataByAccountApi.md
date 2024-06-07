# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LicenseUsageActivity**](LicenseUsageDataByAccountApi.md#LicenseUsageActivity) | **Post** /v1/licenseUsageActivity/{accountIdentifier} | Get License Usage Data for an account for a given time range

# **LicenseUsageActivity**
> []LicenseUsageActivity LicenseUsageActivity(ctx, body, accountIdentifier, startTime, endTime, rollup)
Get License Usage Data for an account for a given time range

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LicenseUsageActivityFilterPropertiesDto**](LicenseUsageActivityFilterPropertiesDto.md)| Details of the filters applied | 
  **accountIdentifier** | **string**| Account Identifier for the Entity. | 
  **startTime** | **int64**| Start Time of the Interval for the Entity. | 
  **endTime** | **int64**| End Time of the Interval for the Entity. | 
  **rollup** | **bool**| Rollup all credit usages in the specified timestamp to a single value. | 

### Return type

[**[]LicenseUsageActivity**](LicenseUsageActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

