# RolesResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** | Role Identifier | [default to null]
**Name** | **string** | Role Name | [default to null]
**Permissions** | **[]string** | Permissions for this Role. | [optional] [default to null]
**AllowedScopeLevels** | **[]string** | The Scope levels at which this Role can be used. | [optional] [default to null]
**Description** | **string** | Role description | [optional] [default to null]
**Tags** | **map[string]string** | Role tags | [optional] [default to null]
**Scope** | [***RoleScope**](RoleScope.md) |  | [optional] [default to null]
**Created** | **int64** | Creation timestamp for Role. | [optional] [default to null]
**Updated** | **int64** | Last modification timestamp for Role. | [optional] [default to null]
**HarnessManaged** | **bool** | This indicates if this Role is managed by Harness or not. If true, Harness can manage and modify this Role. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

