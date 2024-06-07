# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedService**](OrgServicesApi.md#CreateOrgScopedService) | **Post** /v1/orgs/{org}/services | Create a service
[**DeleteOrgScopedService**](OrgServicesApi.md#DeleteOrgScopedService) | **Delete** /v1/orgs/{org}/services/{service} | Delete a service
[**GetOrgScopedPrimaryManifests**](OrgServicesApi.md#GetOrgScopedPrimaryManifests) | **Get** /v1/orgs/{org}/services/{service}/primary-manifests | List primary manifests of a org scoped service
[**GetOrgScopedService**](OrgServicesApi.md#GetOrgScopedService) | **Get** /v1/orgs/{org}/services/{service} | Retrieve a service
[**GetOrgScopedServices**](OrgServicesApi.md#GetOrgScopedServices) | **Get** /v1/orgs/{org}/services | List Services
[**UpdateOrgScopedService**](OrgServicesApi.md#UpdateOrgScopedService) | **Put** /v1/orgs/{org}/services/{service} | Update Service

# **CreateOrgScopedService**
> ServiceResponse CreateOrgScopedService(ctx, body, org, optional)
Create a service

Creates a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceCreateRequest**](ServiceCreateRequest.md)| Create Service request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgServicesApiCreateOrgScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgServicesApiCreateOrgScopedServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgScopedService**
> ServiceResponse DeleteOrgScopedService(ctx, org, service, optional)
Delete a service

Deletes the requested service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***OrgServicesApiDeleteOrgScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgServicesApiDeleteOrgScopedServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **forceDelete** | **optional.Bool**| Enable this field to force delete the entity | [default to false]

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedPrimaryManifests**
> ManifestsResponseDto GetOrgScopedPrimaryManifests(ctx, service, org, optional)
List primary manifests of a org scoped service

Returns a list of eligible primary manifests from all configured manifest in an org-scoped service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Identifier field of the service the resource is scoped to | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgServicesApiGetOrgScopedPrimaryManifestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgServicesApiGetOrgScopedPrimaryManifestsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ManifestsResponseDto**](ManifestsResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedService**
> ServiceResponse GetOrgScopedService(ctx, org, service, optional)
Retrieve a service

Retrieves the specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***OrgServicesApiGetOrgScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgServicesApiGetOrgScopedServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedServices**
> []ServiceResponse GetOrgScopedServices(ctx, org, optional)
List Services

Returns a list of the services for which you have view permissions in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgServicesApiGetOrgScopedServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgServicesApiGetOrgScopedServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **serviceIds** | [**optional.Interface of []string**](string.md)| List of Service Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list of services that you are permitted to use in a pipeline. | 
 **deploymentType** | **optional.String**| Service Definition Type | 
 **gitOpsEnabled** | **optional.Bool**| Enables you to use the service in Harness GitOps PR pipelines. | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedService**
> ServiceResponse UpdateOrgScopedService(ctx, body, org, service, optional)
Update Service

Updates the specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceUpdateRequest**](ServiceUpdateRequest.md)| Update Service request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***OrgServicesApiUpdateOrgScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgServicesApiUpdateOrgScopedServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ServiceResponse**](ServiceResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

