# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedService**](AccountServicesApi.md#CreateAccountScopedService) | **Post** /v1/services | Create a service
[**DeleteAccountScopedService**](AccountServicesApi.md#DeleteAccountScopedService) | **Delete** /v1/services/{service} | Delete a service
[**GetAccountScopedPrimaryManifests**](AccountServicesApi.md#GetAccountScopedPrimaryManifests) | **Get** /v1/services/{service}/primary-manifests | List primary manifests of a account scoped service
[**GetAccountScopedService**](AccountServicesApi.md#GetAccountScopedService) | **Get** /v1/services/{service} | Retrieve a service
[**GetAccountScopedServices**](AccountServicesApi.md#GetAccountScopedServices) | **Get** /v1/services | List services
[**UpdateAccountScopedService**](AccountServicesApi.md#UpdateAccountScopedService) | **Put** /v1/services/{service} | Update service

# **CreateAccountScopedService**
> ServiceResponse CreateAccountScopedService(ctx, body, optional)
Create a service

Creates a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceCreateRequest**](ServiceCreateRequest.md)| Create Service request body | 
 **optional** | ***AccountServicesApiCreateAccountScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountServicesApiCreateAccountScopedServiceOpts struct
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

# **DeleteAccountScopedService**
> ServiceResponse DeleteAccountScopedService(ctx, service, optional)
Delete a service

Deletes the requested service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***AccountServicesApiDeleteAccountScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountServicesApiDeleteAccountScopedServiceOpts struct
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

# **GetAccountScopedPrimaryManifests**
> ManifestsResponseDto GetAccountScopedPrimaryManifests(ctx, service, optional)
List primary manifests of a account scoped service

Returns a list of eligible primary manifests from all configured manifest in an account-scoped service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***AccountServicesApiGetAccountScopedPrimaryManifestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountServicesApiGetAccountScopedPrimaryManifestsOpts struct
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

# **GetAccountScopedService**
> ServiceResponse GetAccountScopedService(ctx, service, optional)
Retrieve a service

Retrieves the specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***AccountServicesApiGetAccountScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountServicesApiGetAccountScopedServiceOpts struct
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

# **GetAccountScopedServices**
> []ServiceResponse GetAccountScopedServices(ctx, optional)
List services

Returns a list of the services for which you have view permissions in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountServicesApiGetAccountScopedServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountServicesApiGetAccountScopedServicesOpts struct
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

# **UpdateAccountScopedService**
> ServiceResponse UpdateAccountScopedService(ctx, body, service, optional)
Update service

Updates the specified service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceUpdateRequest**](ServiceUpdateRequest.md)| Update Service request body | 
  **service** | **string**| Identifier field of the service the resource is scoped to | 
 **optional** | ***AccountServicesApiUpdateAccountScopedServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountServicesApiUpdateAccountScopedServiceOpts struct
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

