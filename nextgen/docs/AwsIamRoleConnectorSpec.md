# AwsIamRoleConnectorSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | This specifies the type of connector | [default to null]
**CrossAccountRoleArn** | **string** | If you want to use one AWS account for the connection, but you want to deploy or build in a different AWS account. In this scenario, the AWS account used for AWS access in Credentials will assume the IAM role you specify in Cross-account role ARN setting. This option uses the AWS Security Token Service (STS) feature. | [optional] [default to null]
**ExternalId** | **string** | If the administrator of the account to which the role belongs provided you with an external ID, then enter that value. | [optional] [default to null]
**TestRegion** | **string** | By default, Harness uses the us-east-1 region to test the credentials for this Connector. If you want to use an AWS GovCloud account for this Connector, select it in Test Region. GovCloud is used by organizations such as government agencies at the federal, state, and local level, as well as contractors, educational institutions. It is also used for regulatory compliance with these organizations. | [optional] [default to null]
**DelegateSelectors** | **[]string** | List of unique delegate selectors | [optional] [default to null]
**ExecuteOnDelegate** | **bool** | execute on delegate | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

