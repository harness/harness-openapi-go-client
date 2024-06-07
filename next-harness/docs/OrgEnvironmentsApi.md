# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedEnvironment**](OrgEnvironmentsApi.md#CreateOrgScopedEnvironment) | **Post** /v1/orgs/{org}/environments | Creates an org scoped Environment
[**DeleteOrgScopedEnvironment**](OrgEnvironmentsApi.md#DeleteOrgScopedEnvironment) | **Delete** /v1/orgs/{org}/environments/{environment} | Delete an org scoped Environment
[**GetOrgScopedEnvironment**](OrgEnvironmentsApi.md#GetOrgScopedEnvironment) | **Get** /v1/orgs/{org}/environments/{environment} | Retrieve an org scoped environment
[**GetOrgScopedEnvironments**](OrgEnvironmentsApi.md#GetOrgScopedEnvironments) | **Get** /v1/orgs/{org}/environments | List org scoped Environments
[**UpdateOrgScopedEnvironment**](OrgEnvironmentsApi.md#UpdateOrgScopedEnvironment) | **Put** /v1/orgs/{org}/environments/{environment} | Update org scoped Environment

# **CreateOrgScopedEnvironment**
> EnvironmentResponse CreateOrgScopedEnvironment(ctx, body, org, optional)
Creates an org scoped Environment

Creates an org scoped Environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentCreateRequest**](EnvironmentCreateRequest.md)| Update Environment request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgEnvironmentsApiCreateOrgScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgEnvironmentsApiCreateOrgScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgScopedEnvironment**
> DeleteOrgScopedEnvironment(ctx, org, environment, optional)
Delete an org scoped Environment

Deletes the requested org scoped environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***OrgEnvironmentsApiDeleteOrgScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgEnvironmentsApiDeleteOrgScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **forceDelete** | **optional.Bool**| Enable this field to force delete the entity | [default to false]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedEnvironment**
> EnvironmentResponse GetOrgScopedEnvironment(ctx, org, environment, optional)
Retrieve an org scoped environment

Retrieves the specified org scoped environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***OrgEnvironmentsApiGetOrgScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgEnvironmentsApiGetOrgScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedEnvironments**
> []EnvironmentResponse GetOrgScopedEnvironments(ctx, org, optional)
List org scoped Environments

Returns a list of the org scoped environments for which you have view permissions in the given org.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgEnvironmentsApiGetOrgScopedEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgEnvironmentsApiGetOrgScopedEnvironmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **environmentIds** | [**optional.Interface of []string**](string.md)| List of Environment Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list org scoped of environments that you are permitted to use in a pipeline. | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedEnvironment**
> EnvironmentResponse UpdateOrgScopedEnvironment(ctx, body, org, environment, optional)
Update org scoped Environment

Updates the specified org scoped environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnvironmentUpdateRequest**](EnvironmentUpdateRequest.md)| Update Environment request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***OrgEnvironmentsApiUpdateOrgScopedEnvironmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgEnvironmentsApiUpdateOrgScopedEnvironmentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

