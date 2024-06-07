# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllowedIpAddress**](IPAllowlistApi.md#AllowedIpAddress) | **Get** /v1/ip-allowlist/allowed/ip-address | Check if IP address is allowed or not
[**CreateIpAllowlistConfig**](IPAllowlistApi.md#CreateIpAllowlistConfig) | **Post** /v1/ip-allowlist | Create a IP Allowlist config
[**DeleteIpAllowlistConfig**](IPAllowlistApi.md#DeleteIpAllowlistConfig) | **Delete** /v1/ip-allowlist/{ip-config-identifier} | Delete an IP Allowlist config
[**GetIpAllowlistConfig**](IPAllowlistApi.md#GetIpAllowlistConfig) | **Get** /v1/ip-allowlist/{ip-config-identifier} | Retrieve a IP Allowlist config
[**GetIpAllowlistConfigs**](IPAllowlistApi.md#GetIpAllowlistConfigs) | **Get** /v1/ip-allowlist | List IP Allowlist Configs
[**UpdateIpAllowlistConfig**](IPAllowlistApi.md#UpdateIpAllowlistConfig) | **Put** /v1/ip-allowlist/{ip-config-identifier} | Update IP Allowlist config
[**ValidateIpAddressAllowlistedOrNot**](IPAllowlistApi.md#ValidateIpAddressAllowlistedOrNot) | **Get** /v1/ip-allowlist/validate/ip-address | Validate IP address lies in a specified range or not
[**ValidateUniqueIpAllowlistConfigIdentifier**](IPAllowlistApi.md#ValidateUniqueIpAllowlistConfigIdentifier) | **Get** /v1/ip-allowlist/validate-unique-identifier/{ip-config-identifier} | Validate unique IP Allowlist config identifier

# **AllowedIpAddress**
> IpAllowlistConfigValidateResponse AllowedIpAddress(ctx, ipAddress, optional)
Check if IP address is allowed or not

Checks whether the IP address is allowed or not.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**| This is the IP address that needs to be checked if allowed or not | 
 **optional** | ***IPAllowlistApiAllowedIpAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiAllowedIpAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**IpAllowlistConfigValidateResponse**](IPAllowlistConfigValidateResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpAllowlistConfig**
> IpAllowlistConfigResponse CreateIpAllowlistConfig(ctx, optional)
Create a IP Allowlist config

Creates a new IP Allowlist config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAllowlistApiCreateIpAllowlistConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiCreateIpAllowlistConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of IpAllowlistConfigRequest**](IpAllowlistConfigRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**IpAllowlistConfigResponse**](IPAllowlistConfigResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpAllowlistConfig**
> DeleteIpAllowlistConfig(ctx, ipConfigIdentifier, optional)
Delete an IP Allowlist config

Deletes the specified IP Allowlist config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipConfigIdentifier** | **string**|  | 
 **optional** | ***IPAllowlistApiDeleteIpAllowlistConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiDeleteIpAllowlistConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpAllowlistConfig**
> IpAllowlistConfigResponse GetIpAllowlistConfig(ctx, ipConfigIdentifier, optional)
Retrieve a IP Allowlist config

Retrieves the specified IP Allowlist config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipConfigIdentifier** | **string**|  | 
 **optional** | ***IPAllowlistApiGetIpAllowlistConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiGetIpAllowlistConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**IpAllowlistConfigResponse**](IPAllowlistConfigResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpAllowlistConfigs**
> []IpAllowlistConfigResponse GetIpAllowlistConfigs(ctx, optional)
List IP Allowlist Configs

Retrieves the information of the IP Allowlist Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IPAllowlistApiGetIpAllowlistConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiGetIpAllowlistConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **allowedSourceType** | **optional.String**| This is to filter IP allowlist configs only blocked from UI or API | 

### Return type

[**[]IpAllowlistConfigResponse**](IPAllowlistConfigResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpAllowlistConfig**
> IpAllowlistConfigResponse UpdateIpAllowlistConfig(ctx, ipConfigIdentifier, optional)
Update IP Allowlist config

Updates the specified IP Allowlist config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipConfigIdentifier** | **string**|  | 
 **optional** | ***IPAllowlistApiUpdateIpAllowlistConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiUpdateIpAllowlistConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IpAllowlistConfigRequest**](IpAllowlistConfigRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**IpAllowlistConfigResponse**](IPAllowlistConfigResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateIpAddressAllowlistedOrNot**
> IpAllowlistConfigValidateResponse ValidateIpAddressAllowlistedOrNot(ctx, ipAddress, optional)
Validate IP address lies in a specified range or not

Checks whether the IP address is allowed or not. It also supports checking against a specific IP block range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**| This is the IP address that needs to be checked if allowed or not | 
 **optional** | ***IPAllowlistApiValidateIpAddressAllowlistedOrNotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiValidateIpAddressAllowlistedOrNotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **customIpAddressBlock** | **optional.String**| This is the IP address or block of IP address against which we need to verify if a given IP address is allowed or not. If not passed we do the validation against the IP configs within Harness. | 
 **includeDisabledConfigs** | **optional.Bool**| This setting controls the visibility of IP allowlist configurations. When set to &#x27;true&#x27;, it displays both enabled and disabled configurations. When set to &#x27;false&#x27; or left unset, it displays only the enabled configurations. | [default to false]

### Return type

[**IpAllowlistConfigValidateResponse**](IPAllowlistConfigValidateResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateUniqueIpAllowlistConfigIdentifier**
> bool ValidateUniqueIpAllowlistConfigIdentifier(ctx, ipConfigIdentifier, optional)
Validate unique IP Allowlist config identifier

Checks whether the IP Allowlist config identifier is unique or not

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipConfigIdentifier** | **string**|  | 
 **optional** | ***IPAllowlistApiValidateUniqueIpAllowlistConfigIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IPAllowlistApiValidateUniqueIpAllowlistConfigIdentifierOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

**bool**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

