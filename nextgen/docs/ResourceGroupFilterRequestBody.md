# ResourceGroupFilterRequestBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Filter by Account identifier. | [optional] [default to null]
**Org** | **string** | Filter by Organization identifier. | [optional] [default to null]
**Project** | **string** | Filter by Project identifier. | [optional] [default to null]
**SearchTerm** | **string** | Filter Resource Group matching by identifier/name. | [optional] [default to null]
**IdentifierFilter** | **[]string** | Filter by Resource Group identifiers | [optional] [default to null]
**ResourceSelectorFilter** | [**[]ResourceSelectorFilter**](ResourceSelectorFilter.md) | Filter based on whether it has a particular Resource. | [optional] [default to null]
**ManagedFilter** | **string** | Filter based on whether the Resource Group is Harness Managed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

