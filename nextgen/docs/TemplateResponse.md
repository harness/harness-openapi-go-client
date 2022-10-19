# TemplateResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Account identifier | [default to null]
**Org** | **string** | Organization identifier | [optional] [default to null]
**Project** | **string** | Project identifier | [optional] [default to null]
**Slug** | **string** | Template identifier | [default to null]
**Name** | **string** | Template Name | [default to null]
**Description** | **string** | Template description | [optional] [default to null]
**Tags** | **map[string]string** | Template tags | [optional] [default to null]
**Yaml** | **string** | Yaml related to template | [default to null]
**VersionLabel** | **string** | Version label of template | [optional] [default to null]
**EntityType** | **string** | Type of Template  | [optional] [default to null]
**ChildType** | **string** | Defines child template type | [optional] [default to null]
**Scope** | **string** | Scope of template | [optional] [default to null]
**Version** | **int64** | Version of template | [optional] [default to null]
**GitDetails** | [***EntityGitDetails**](EntityGitDetails.md) |  | [optional] [default to null]
**Updated** | **int64** | Last modification timestamp for Service.  | [optional] [default to null]
**StoreType** | **string** | Specifies whether the Entity is to be stored in Git or not (for Git Experience). | [optional] [default to null]
**ConnectorRef** | **string** | Identifier of the Harness Connector used for CRUD operations on the Entity (for Git Experience). | [optional] [default to null]
**StableTemplate** | **bool** | True if this version is stable version of Template | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

