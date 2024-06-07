# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNgLdapSettings**](LdapSettingsApi.md#CreateNgLdapSettings) | **Post** /v1/ldap-settings | Create Account LDAP Settings
[**DeleteLdapSettings**](LdapSettingsApi.md#DeleteLdapSettings) | **Delete** /v1/ldap-settings/{ldap-settings-id} | Delete Account Ldap Settings
[**GetAccountLdapSettings**](LdapSettingsApi.md#GetAccountLdapSettings) | **Get** /v1/ldap-settings | Get ldap settings for the account
[**GetLdapSettings**](LdapSettingsApi.md#GetLdapSettings) | **Get** /v1/ldap-settings/{ldap-settings-id} | Get Account Ldap Settings by identifier
[**LdapSettingsIterations**](LdapSettingsApi.md#LdapSettingsIterations) | **Post** /v1/ldap-settings/iterations | Ldap Settings iterations
[**LdapSettingsLdapLoginTest**](LdapSettingsApi.md#LdapSettingsLdapLoginTest) | **Post** /v1/ldap-settings/ldap-login-test | Ldap login test
[**LdapSettingsSearchGroup**](LdapSettingsApi.md#LdapSettingsSearchGroup) | **Get** /v1/ldap-settings/search-group/{group-id} | Search Ldap groups with matching name
[**LdapSettingsSyncGroupWithId**](LdapSettingsApi.md#LdapSettingsSyncGroupWithId) | **Get** /v1/ldap-settings/sync-group/{group-id} | Ldap Sync Group with user group id
[**LdapSettingsSyncGroups**](LdapSettingsApi.md#LdapSettingsSyncGroups) | **Put** /v1/ldap-settings/sync-groups | Ldap Settings Sync User Groups
[**LinkLdapSettings**](LdapSettingsApi.md#LinkLdapSettings) | **Post** /v1/ldap-settings/link/{group-id} | Link ldap settings to userGroup
[**LinkLdapSettingsOrg**](LdapSettingsApi.md#LinkLdapSettingsOrg) | **Post** /v1/ldap-settings/link/orgs/{org}/{group-id} | Link dap settings
[**LinkLdapSettingsProj**](LdapSettingsApi.md#LinkLdapSettingsProj) | **Post** /v1/ldap-settings/link/orgs/{org}/projects/{project}/{group-id} | Link ldap-settings
[**UnlinkLdapSettings**](LdapSettingsApi.md#UnlinkLdapSettings) | **Post** /v1/ldap-settings/unlink/{group-id} | Unlink group from ldap settings
[**UnlinkLdapSettingsOrg**](LdapSettingsApi.md#UnlinkLdapSettingsOrg) | **Post** /v1/ldap-settings/unlink/orgs/{org}/{group-id} | Unlink SSO Group from ldap-settings
[**UnlinkLdapSettingsProj**](LdapSettingsApi.md#UnlinkLdapSettingsProj) | **Post** /v1/ldap-settings/unlink/orgs/{org}/projects/{project}/{group-id} | Unlink SSO Group from ldap-settings
[**UpdateLdapSettings**](LdapSettingsApi.md#UpdateLdapSettings) | **Put** /v1/ldap-settings/{ldap-settings-id} | Update Account Ldap Settings
[**ValidateConnectionSettings**](LdapSettingsApi.md#ValidateConnectionSettings) | **Post** /v1/ldap-settings/validate/connection-settings | Validate Ldap Connection Settings
[**ValidateGroupSettings**](LdapSettingsApi.md#ValidateGroupSettings) | **Post** /v1/ldap-settings/validate/group-settings | Validate Ldap Group Settings
[**ValidateUserSettings**](LdapSettingsApi.md#ValidateUserSettings) | **Post** /v1/ldap-settings/validate/user-settings | Validate Ldap User Settings

# **CreateNgLdapSettings**
> LdapSettingsResponse CreateNgLdapSettings(ctx, optional)
Create Account LDAP Settings

Create NG Ldap setttings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiCreateNgLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiCreateNgLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapSettingsRequest**](LdapSettingsRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapSettingsResponse**](LdapSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLdapSettings**
> DeleteLdapSettings(ctx, ldapSettingsId, optional)
Delete Account Ldap Settings

Delete Account Ldap settings 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapSettingsId** | **string**|  | 
 **optional** | ***LdapSettingsApiDeleteLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiDeleteLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountLdapSettings**
> LdapSettingsResponse GetAccountLdapSettings(ctx, optional)
Get ldap settings for the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiGetAccountLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiGetAccountLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapSettingsResponse**](LdapSettingsResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLdapSettings**
> LdapSettingsResponse GetLdapSettings(ctx, ldapSettingsId, optional)
Get Account Ldap Settings by identifier

Get Account NG Ldap settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapSettingsId** | **string**|  | 
 **optional** | ***LdapSettingsApiGetLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiGetLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapSettingsResponse**](LdapSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapSettingsIterations**
> LdapSettingsIterations LdapSettingsIterations(ctx, optional)
Ldap Settings iterations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiLdapSettingsIterationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLdapSettingsIterationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CronExpressionRequestDto**](CronExpressionRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapSettingsIterations**](LdapSettingsIterations.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapSettingsLdapLoginTest**
> LdapTestLoginResponseDto LdapSettingsLdapLoginTest(ctx, optional)
Ldap login test

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiLdapSettingsLdapLoginTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLdapSettingsLdapLoginTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapTestLoginRequestDto**](LdapTestLoginRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapTestLoginResponseDto**](LdapTestLoginResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapSettingsSearchGroup**
> LdapGroupResponseDto LdapSettingsSearchGroup(ctx, groupId, optional)
Search Ldap groups with matching name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiLdapSettingsSearchGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLdapSettingsSearchGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapGroupResponseDto**](LdapGroupResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapSettingsSyncGroupWithId**
> LdapSettingsSyncGroupWithId(ctx, groupId, optional)
Ldap Sync Group with user group id

Ldap Sync Group with user group id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiLdapSettingsSyncGroupWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLdapSettingsSyncGroupWithIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapSettingsSyncGroups**
> LdapSettingsSyncGroups(ctx, optional)
Ldap Settings Sync User Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiLdapSettingsSyncGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLdapSettingsSyncGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkLdapSettings**
> UserGroupDto LinkLdapSettings(ctx, groupId, optional)
Link ldap settings to userGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiLinkLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLinkLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of LinkSsoGroupRequestDto**](LinkSsoGroupRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UserGroupDto**](UserGroupDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkLdapSettingsOrg**
> UserGroupDto LinkLdapSettingsOrg(ctx, org, groupId, optional)
Link dap settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiLinkLdapSettingsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLinkLdapSettingsOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of LinkSsoGroupRequestDto**](LinkSsoGroupRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UserGroupDto**](UserGroupDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkLdapSettingsProj**
> UserGroupDto LinkLdapSettingsProj(ctx, org, project, groupId, optional)
Link ldap-settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiLinkLdapSettingsProjOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiLinkLdapSettingsProjOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of LinkSsoGroupRequestDto**](LinkSsoGroupRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UserGroupDto**](UserGroupDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkLdapSettings**
> UserGroupDto UnlinkLdapSettings(ctx, groupId, optional)
Unlink group from ldap settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiUnlinkLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiUnlinkLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UnlinkSsoGroupRequestDto**](UnlinkSsoGroupRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UserGroupDto**](UserGroupDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkLdapSettingsOrg**
> UserGroupDto UnlinkLdapSettingsOrg(ctx, org, groupId, optional)
Unlink SSO Group from ldap-settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiUnlinkLdapSettingsOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiUnlinkLdapSettingsOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UnlinkSsoGroupRequestDto**](UnlinkSsoGroupRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UserGroupDto**](UserGroupDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkLdapSettingsProj**
> UserGroupDto UnlinkLdapSettingsProj(ctx, org, project, groupId, optional)
Unlink SSO Group from ldap-settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **groupId** | **string**| UserGroup Identifier | 
 **optional** | ***LdapSettingsApiUnlinkLdapSettingsProjOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiUnlinkLdapSettingsProjOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UnlinkSsoGroupRequestDto**](UnlinkSsoGroupRequestDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**UserGroupDto**](UserGroupDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLdapSettings**
> LdapSettingsResponse UpdateLdapSettings(ctx, ldapSettingsId, optional)
Update Account Ldap Settings

Update NG Ldap settings for the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ldapSettingsId** | **string**|  | 
 **optional** | ***LdapSettingsApiUpdateLdapSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiUpdateLdapSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of LdapSettingsRequest**](LdapSettingsRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapSettingsResponse**](LdapSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateConnectionSettings**
> LdapValidateResponseDto ValidateConnectionSettings(ctx, optional)
Validate Ldap Connection Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiValidateConnectionSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiValidateConnectionSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapSettingsRequest**](LdapSettingsRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapValidateResponseDto**](LdapValidateResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateGroupSettings**
> LdapValidateResponseDto ValidateGroupSettings(ctx, optional)
Validate Ldap Group Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiValidateGroupSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiValidateGroupSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapSettingsRequest**](LdapSettingsRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapValidateResponseDto**](LdapValidateResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateUserSettings**
> LdapValidateResponseDto ValidateUserSettings(ctx, optional)
Validate Ldap User Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapSettingsApiValidateUserSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapSettingsApiValidateUserSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapSettingsRequest**](LdapSettingsRequest.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**LdapValidateResponseDto**](LdapValidateResponseDTO.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

