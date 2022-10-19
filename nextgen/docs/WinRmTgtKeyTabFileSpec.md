# WinRmTgtKeyTabFileSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | This specifies the type of secret | [default to null]
**Port** | **int32** | WinRm port | [optional] [default to 5986]
**Principal** | **string** | Kerberos principal | [optional] [default to null]
**Realm** | **string** | Kerberos realm | [optional] [default to null]
**KeyPath** | **string** | Keytab file path | [optional] [default to null]
**UseSsl** | **bool** | This is the Kerberos either to use SSL/https | [optional] [default to true]
**SkipCertChecks** | **bool** | This is the Kerberos either to skip certificate checks | [optional] [default to true]
**UseNoProfile** | **bool** | This is the Kerberos powershell runs without loading profile | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

