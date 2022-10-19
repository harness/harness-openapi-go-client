# AwsSecretManagerAccessKeyConnectorSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | This specifies the type of connector | [default to null]
**Region** | **string** | AWS Region for kms | [default to null]
**Default_** | **bool** | Boolean value to indicate if the Secret Manager is your default Secret Manager | [optional] [default to null]
**AccessKey** | **string** | Access Key for AWS authentication | [default to null]
**SecretKey** | **string** | Secret Key for AWS authentication | [default to null]
**SecretNamePrefix** | **string** | Text that is prepended to the Secret name as a prefix | [optional] [default to null]
**DelegateSelectors** | **[]string** | List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

