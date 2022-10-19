# GitHttpConnectorSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | This specifies the type of connector | [default to null]
**Url** | **string** | Git repo url | [default to null]
**Branch** | **string** | branch name | [optional] [default to null]
**ConnectionType** | **string** |  | [default to null]
**Username** | **string** | git username | [default to null]
**PasswordRef** | **string** | Reference to encrypted Harness secret for git password | [default to null]
**ValidationRepo** | **string** | validation repo | [optional] [default to null]
**DelegateSelectors** | **[]string** | List of unique delegate selectors | [optional] [default to null]
**ExecuteOnDelegate** | **bool** | execute on delegate | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

