# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountScopedRancherClustersUsingConnector**](AccountRancherInfrastructureApi.md#ListAccountScopedRancherClustersUsingConnector) | **Get** /v1/rancher/connectors/{connector}/clusters | List rancher clusters using account level connector
[**ListAccountScopedRancherClustersUsingEnvAndInfra**](AccountRancherInfrastructureApi.md#ListAccountScopedRancherClustersUsingEnvAndInfra) | **Get** /v1/rancher/environments/{environment}/infrastructure-definitions/{infrastructure-definition}/clusters | List rancher clusters using account level env and infra def

# **ListAccountScopedRancherClustersUsingConnector**
> []string ListAccountScopedRancherClustersUsingConnector(ctx, connector, optional)
List rancher clusters using account level connector

List rancher clusters using the given account level rancher connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| Identifier field of the scoped connector entity to be used for this operation. | 
 **optional** | ***AccountRancherInfrastructureApiListAccountScopedRancherClustersUsingConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRancherInfrastructureApiListAccountScopedRancherClustersUsingConnectorOpts struct
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

# **ListAccountScopedRancherClustersUsingEnvAndInfra**
> []string ListAccountScopedRancherClustersUsingEnvAndInfra(ctx, environment, infrastructureDefinition, optional)
List rancher clusters using account level env and infra def

List rancher clusters using the given account level environment and infrastructure definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environment** | **string**| Identifier field of the scoped environment entity to be used for the selected operation. | 
  **infrastructureDefinition** | **string**| Identifier field of the scoped infrastructure definition entity to be used in the selected operation. | 
 **optional** | ***AccountRancherInfrastructureApiListAccountScopedRancherClustersUsingEnvAndInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountRancherInfrastructureApiListAccountScopedRancherClustersUsingEnvAndInfraOpts struct
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

