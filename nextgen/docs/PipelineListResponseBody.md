# PipelineListResponseBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | Pipeline identifier | [optional] [default to null]
**Name** | **string** | Pipeline name | [optional] [default to null]
**Description** | **string** | Pipeline description | [optional] [default to null]
**Tags** | **map[string]string** | Pipeline tags | [optional] [default to null]
**Created** | **int64** | Creation timestamp for Pipeline. | [optional] [default to null]
**Updated** | **int64** | Last modification timestamp for Pipeline. | [optional] [default to null]
**Modules** | **[]string** | Modules utilised in the Pipeline. | [optional] [default to null]
**ExecutionSummary** | [***ExecutionSummary**](ExecutionSummary.md) |  | [optional] [default to null]
**RecentExecutionInfo** | [**[]RecentExecutionInfo**](RecentExecutionInfo.md) | Array of recent Execution information | [optional] [default to null]
**StoreType** | **string** | Specifies whether the Entity is to be stored in Git or not (for Git Experience). | [optional] [default to null]
**ConnectorRef** | **string** | Identifier of the Harness Connector used for CRUD operations on the Entity (for Git Experience). | [optional] [default to null]
**Valid** | **bool** | Specifies whether Pipeline is a valid or not. | [optional] [default to null]
**GitDetails** | [***GitDetails**](GitDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

