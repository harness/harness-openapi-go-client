# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProjectScopedRancherClustersUsingConnector**](ProjectRancherInfrastructureApi.md#ListProjectScopedRancherClustersUsingConnector) | **Get** /v1/orgs/{org}/projects/{project}/rancher/connectors/{connector}/clusters | List rancher clusters using project level connector
[**ListProjectScopedRancherClustersUsingEnvAndInfra**](ProjectRancherInfrastructureApi.md#ListProjectScopedRancherClustersUsingEnvAndInfra) | **Get** /v1/orgs/{org}/projects/{project}/rancher/environments/{environment}/infrastructure-definitions/{infrastructure-definition}/clusters | List rancher clusters using project level env and infra def

# **ListProjectScopedRancherClustersUsingConnector**
> []string ListProjectScopedRancherClustersUsingConnector(ctx, org, project, connector, optional)
List rancher clusters using project level connector

List rancher clusters using the given project level rancher connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **connector** | **string**| Identifier field of the scoped connector entity to be used for this operation. | 
 **optional** | ***ProjectRancherInfrastructureApiListProjectScopedRancherClustersUsingConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRancherInfrastructureApiListProjectScopedRancherClustersUsingConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

**[]string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjectScopedRancherClustersUsingEnvAndInfra**
> []string ListProjectScopedRancherClustersUsingEnvAndInfra(ctx, org, project, environment, infrastructureDefinition, optional)
List rancher clusters using project level env and infra def

List rancher clusters using the given project level environment and infrastructure definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***ProjectRancherInfrastructureApiListProjectScopedRancherClustersUsingEnvAndInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectRancherInfrastructureApiListProjectScopedRancherClustersUsingEnvAndInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

**[]string**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

