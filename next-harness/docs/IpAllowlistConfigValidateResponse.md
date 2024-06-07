# IpAllowlistConfigValidateResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedForCustomBlock** | **bool** | This indicates if given IP Address lies in range of custom IP block or not. | [optional] [default to null]
**AllowlistedConfigs** | [**[]IpAllowlistConfigResponse**](IPAllowlistConfigResponse.md) | This is the list of IP configs configured in Harness from which IP address is allowed. This is empty in case of custom IP address block.  | [optional] [default to null]
**AllowedForUi** | **bool** | This indicates if a given IP is allowlisted in Harness for UI requests | [optional] [default to null]
**AllowedForApi** | **bool** | This indicates if a given IP is allowlisted in Harness for API requests | [optional] [default to null]
**DisabledAllowlistedConfigs** | [**[]IpAllowlistConfigResponse**](IPAllowlistConfigResponse.md) | This is the list of IP configs configured in Harness from which IP address is allowed but the config is disabled. This is empty in case of custom IP address block.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

