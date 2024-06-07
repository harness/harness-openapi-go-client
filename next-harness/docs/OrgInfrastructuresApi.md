# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedInfrastructure**](OrgInfrastructuresApi.md#CreateOrgScopedInfrastructure) | **Post** /v1/orgs/{org}/environments/{environment}/infrastructures | Creates an org scoped Infrastructure
[**DeleteOrgScopedInfrastructure**](OrgInfrastructuresApi.md#DeleteOrgScopedInfrastructure) | **Delete** /v1/orgs/{org}/environments/{environment}/infrastructures/{infrastructure-definition} | Delete an org scoped Infrastructure
[**GetOrgScopedInfrastructure**](OrgInfrastructuresApi.md#GetOrgScopedInfrastructure) | **Get** /v1/orgs/{org}/environments/{environment}/infrastructures/{infrastructure-definition} | Retrieve an org scoped Infrastructure
[**GetOrgScopedInfrastructures**](OrgInfrastructuresApi.md#GetOrgScopedInfrastructures) | **Get** /v1/orgs/{org}/environments/{environment}/infrastructures | List org scoped Infrastructures
[**UpdateOrgScopedInfrastructure**](OrgInfrastructuresApi.md#UpdateOrgScopedInfrastructure) | **Put** /v1/orgs/{org}/environments/{environment}/infrastructures/{infrastructure-definition} | Update org scoped Infrastructure

# **CreateOrgScopedInfrastructure**
> InfrastructureResponse CreateOrgScopedInfrastructure(ctx, body, org, environment, optional)
Creates an org scoped Infrastructure

Creates an org scoped Infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureCreateRequest**](InfrastructureCreateRequest.md)| Create Infrastructure request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***OrgInfrastructuresApiCreateOrgScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgInfrastructuresApiCreateOrgScopedInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**InfrastructureResponse**](InfrastructureResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgScopedInfrastructure**
> DeleteOrgScopedInfrastructure(ctx, org, environment, infrastructureDefinition, optional)
Delete an org scoped Infrastructure

Deletes the requested org scoped infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***OrgInfrastructuresApiDeleteOrgScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgInfrastructuresApiDeleteOrgScopedInfrastructureOpts struct
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

# **GetOrgScopedInfrastructure**
> InfrastructureResponse GetOrgScopedInfrastructure(ctx, org, environment, infrastructureDefinition, optional)
Retrieve an org scoped Infrastructure

Retrieves the specified org scoped infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***OrgInfrastructuresApiGetOrgScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgInfrastructuresApiGetOrgScopedInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**InfrastructureResponse**](InfrastructureResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedInfrastructures**
> []InfrastructureResponse GetOrgScopedInfrastructures(ctx, org, environment, optional)
List org scoped Infrastructures

Returns a list of the org scoped infrastructures for which you have view permissions in the given org.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***OrgInfrastructuresApiGetOrgScopedInfrastructuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgInfrastructuresApiGetOrgScopedInfrastructuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **infraIds** | [**optional.Interface of []string**](string.md)| List of Infrastructure Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list org scoped of infrastructure that you are permitted to use in a pipeline. | 
 **serviceRefs** | [**optional.Interface of []string**](string.md)| Specifies services to which scoped infra are to be fetched. | 
 **templateIdentifier** | **optional.String**| The Identifier of deployment template if infrastructure is of type custom deployment. | 
 **templateVersion** | **optional.String**| The version label of deployment template if infrastructure is of type custom deployment(deployment template). | 
 **deploymentType** | **optional.String**| Service Definition Type | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]InfrastructureResponse**](InfrastructureResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedInfrastructure**
> InfrastructureResponse UpdateOrgScopedInfrastructure(ctx, body, org, environment, infrastructureDefinition, optional)
Update org scoped Infrastructure

Updates the specified org scoped infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureUpdateRequest**](InfrastructureUpdateRequest.md)| Update Infrastructure request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***OrgInfrastructuresApiUpdateOrgScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgInfrastructuresApiUpdateOrgScopedInfrastructureOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**InfrastructureResponse**](InfrastructureResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

