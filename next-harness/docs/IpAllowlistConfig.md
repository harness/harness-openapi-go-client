# IpAllowlistConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the IP Config defined in Harness | [default to null]
**Identifier** | **string** | Identifier of the IP Config | [default to null]
**Description** | **string** | Description of the entity | [optional] [default to null]
**Enabled** | **bool** | If true, it will allow all the IPs that are part of the config and block others. | [optional] [default to false]
**Tags** | **map[string]string** | IP Allowlist tags | [optional] [default to null]
**AllowedSourceType** | [**[]AllowedSourceType**](AllowedSourceType.md) |  | [optional] [default to null]
**IpAddress** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

