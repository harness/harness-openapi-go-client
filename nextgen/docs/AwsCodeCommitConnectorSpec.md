# AwsCodeCommitConnectorSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | This specifies the type of connector | [default to null]
**Url** | **string** | AWS codecommit repository url | [default to null]
**UrlType** | **string** | AWS codecommit repository url type | [default to null]
**AccessKey** | **string** | AWS access key | [optional] [default to null]
**AccessKeyRef** | **string** | Reference to encrypted Harness secret for AWS access key | [optional] [default to null]
**SecretKeyRef** | **string** | Reference to encrypted Harness secret for AWS secret key | [default to null]
**DelegateSelectors** | **[]string** | List of unique delegate selectors | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

