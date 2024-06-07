# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedInfrastructure**](AccountInfrastructuresApi.md#CreateAccountScopedInfrastructure) | **Post** /v1/environments/{environment}/infrastructures | Creates an account scoped Infrastructure
[**DeleteAccountScopedInfrastructure**](AccountInfrastructuresApi.md#DeleteAccountScopedInfrastructure) | **Delete** /v1/environments/{environment}/infrastructures/{infrastructure-definition} | Delete an account scoped Infrastructure
[**GetAccountScopedInfrastructure**](AccountInfrastructuresApi.md#GetAccountScopedInfrastructure) | **Get** /v1/environments/{environment}/infrastructures/{infrastructure-definition} | Retrieve an account scoped Infrastructure
[**GetAccountScopedInfrastructures**](AccountInfrastructuresApi.md#GetAccountScopedInfrastructures) | **Get** /v1/environments/{environment}/infrastructures | List account scoped Infrastructures
[**UpdateAccountScopedInfrastructure**](AccountInfrastructuresApi.md#UpdateAccountScopedInfrastructure) | **Put** /v1/environments/{environment}/infrastructures/{infrastructure-definition} | Update account scoped Infrastructure

# **CreateAccountScopedInfrastructure**
> InfrastructureResponse CreateAccountScopedInfrastructure(ctx, body, environment, optional)
Creates an account scoped Infrastructure

Creates an account scoped Infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureCreateRequest**](InfrastructureCreateRequest.md)| Create Infrastructure request body | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***AccountInfrastructuresApiCreateAccountScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountInfrastructuresApiCreateAccountScopedInfrastructureOpts struct
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

# **DeleteAccountScopedInfrastructure**
> DeleteAccountScopedInfrastructure(ctx, environment, infrastructureDefinition, optional)
Delete an account scoped Infrastructure

Deletes the requested account scoped infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***AccountInfrastructuresApiDeleteAccountScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountInfrastructuresApiDeleteAccountScopedInfrastructureOpts struct
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

# **GetAccountScopedInfrastructure**
> InfrastructureResponse GetAccountScopedInfrastructure(ctx, environment, infrastructureDefinition, optional)
Retrieve an account scoped Infrastructure

Retrieves the specified account scoped infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***AccountInfrastructuresApiGetAccountScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountInfrastructuresApiGetAccountScopedInfrastructureOpts struct
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

# **GetAccountScopedInfrastructures**
> []InfrastructureResponse GetAccountScopedInfrastructures(ctx, environment, optional)
List account scoped Infrastructures

Returns a list of the account scoped infrastructures for which you have view permissions in the given account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***AccountInfrastructuresApiGetAccountScopedInfrastructuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountInfrastructuresApiGetAccountScopedInfrastructuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **infraIds** | [**optional.Interface of []string**](string.md)| List of Infrastructure Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list account scoped of infrastructures that you are permitted to use in a pipeline. | 
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

# **UpdateAccountScopedInfrastructure**
> InfrastructureResponse UpdateAccountScopedInfrastructure(ctx, body, environment, infrastructureDefinition, optional)
Update account scoped Infrastructure

Updates the specified account scoped infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureUpdateRequest**](InfrastructureUpdateRequest.md)| Update Infrastructure request body | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***AccountInfrastructuresApiUpdateAccountScopedInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountInfrastructuresApiUpdateAccountScopedInfrastructureOpts struct
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

