# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInfrastructure**](ProjectInfrastructuresApi.md#CreateInfrastructure) | **Post** /v1/orgs/{org}/projects/{project}/environments/{environment}/infrastructures | Creates an Infrastructure
[**DeleteInfrastructure**](ProjectInfrastructuresApi.md#DeleteInfrastructure) | **Delete** /v1/orgs/{org}/projects/{project}/environments/{environment}/infrastructures/{infrastructure-definition} | Delete a Infrastructure
[**GetInfrastructure**](ProjectInfrastructuresApi.md#GetInfrastructure) | **Get** /v1/orgs/{org}/projects/{project}/environments/{environment}/infrastructures/{infrastructure-definition} | Retrieve a infrastructure
[**GetInfrastructures**](ProjectInfrastructuresApi.md#GetInfrastructures) | **Get** /v1/orgs/{org}/projects/{project}/environments/{environment}/infrastructures | List Infrastructures
[**UpdateInfrastructure**](ProjectInfrastructuresApi.md#UpdateInfrastructure) | **Put** /v1/orgs/{org}/projects/{project}/environments/{environment}/infrastructures/{infrastructure-definition} | Update Infrastructure

# **CreateInfrastructure**
> InfrastructureResponse CreateInfrastructure(ctx, body, org, project, environment, optional)
Creates an Infrastructure

Creates an Infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureCreateRequest**](InfrastructureCreateRequest.md)| Create Infrastructure request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***ProjectInfrastructuresApiCreateInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectInfrastructuresApiCreateInfrastructureOpts struct
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

# **DeleteInfrastructure**
> DeleteInfrastructure(ctx, org, project, environment, infrastructureDefinition, optional)
Delete a Infrastructure

Deletes the requested environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***ProjectInfrastructuresApiDeleteInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectInfrastructuresApiDeleteInfrastructureOpts struct
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

# **GetInfrastructure**
> InfrastructureResponse GetInfrastructure(ctx, org, project, environment, infrastructureDefinition, optional)
Retrieve a infrastructure

Retrieves the specified infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***ProjectInfrastructuresApiGetInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectInfrastructuresApiGetInfrastructureOpts struct
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

# **GetInfrastructures**
> []InfrastructureResponse GetInfrastructures(ctx, org, project, environment, optional)
List Infrastructures

Returns a list of the infrastructure for which you have view permissions in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
 **optional** | ***ProjectInfrastructuresApiGetInfrastructuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectInfrastructuresApiGetInfrastructuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **infraIds** | [**optional.Interface of []string**](string.md)| List of Infrastructure Identifiers | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **isAccessList** | **optional.Bool**| Specifies whether the list is an access list. An access list is a list of infrastructure that you are permitted to use in a pipeline. | 
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

# **UpdateInfrastructure**
> InfrastructureResponse UpdateInfrastructure(ctx, body, org, project, environment, infrastructureDefinition, optional)
Update Infrastructure

Updates the specified infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InfrastructureUpdateRequest**](InfrastructureUpdateRequest.md)| Update Infrastructure request body | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***ProjectInfrastructuresApiUpdateInfrastructureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectInfrastructuresApiUpdateInfrastructureOpts struct
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

