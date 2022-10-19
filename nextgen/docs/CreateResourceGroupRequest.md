# CreateResourceGroupRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | Resource Group identifier | [default to null]
**Name** | **string** | Resource Group name | [default to null]
**Color** | **string** | Color associated with the Resource Group. | [optional] [default to null]
**Tags** | **map[string]string** | Resource Group tags | [optional] [default to null]
**Description** | **string** | Resource Group description | [optional] [default to null]
**IncludedScope** | [**[]ResourceGroupScope**](ResourceGroupScope.md) | Included scopes for the resources belonging to the Resource Group. | [optional] [default to null]
**ResourceFilter** | [**[]ResourceFilter**](ResourceFilter.md) | Specifies the actual resources present in the Resource Group. | [optional] [default to null]
**IncludeAllResources** | **bool** | Boolean value for including all resources in Resource Group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

