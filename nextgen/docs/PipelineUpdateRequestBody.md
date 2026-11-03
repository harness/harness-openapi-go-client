# PipelineUpdateRequestBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineYaml** | **string** | Pipeline YAML (to be passed as a String). | [default to null]
**Identifier** | **string** | Pipeline identifier | [default to null]
**Version** | **string** | Pipeline yaml version, should be one of &#39;0&#39; or &#39;1&#39; | [optional] [default to null]
**Name** | **string** | Pipeline name | [default to null]
**Description** | **string** | Pipeline description | [optional] [default to null]
**AllowDynamicExecutions** | **bool** | Allow dynamic executions of Pipeline | [optional] [default to false]
**Tags** | **map[string]string** | Pipeline tags | [optional] [default to null]
**GitDetails** | [***GitUpdateDetails**](GitUpdateDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

