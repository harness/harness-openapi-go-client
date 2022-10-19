# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectScopedConnector**](ProjectConnectorApi.md#CreateProjectScopedConnector) | **Post** /v1/orgs/{org}/projects/{project}/connectors | Create a Connector
[**DeleteProjectScopedConnector**](ProjectConnectorApi.md#DeleteProjectScopedConnector) | **Delete** /v1/orgs/{org}/projects/{project}/connectors/{connector} | Delete a connector
[**GetProjectScopedConnector**](ProjectConnectorApi.md#GetProjectScopedConnector) | **Get** /v1/orgs/{org}/projects/{project}/connectors/{connector} | Retrieve a connector
[**GetProjectScopedConnectors**](ProjectConnectorApi.md#GetProjectScopedConnectors) | **Get** /v1/orgs/{org}/projects/{project}/connectors | List connectors
[**TestProjectScopedConnector**](ProjectConnectorApi.md#TestProjectScopedConnector) | **Get** /v1/orgs/{org}/projects/{project}/connectors/{connector}/test-connection | Test a connector
[**UpdateProjectScopedConnector**](ProjectConnectorApi.md#UpdateProjectScopedConnector) | **Put** /v1/orgs/{org}/projects/{project}/connectors/{connector} | Update a connector

# **CreateProjectScopedConnector**
> ConnectorResponse CreateProjectScopedConnector(ctx, body, org, project, optional)
Create a Connector

Creates a new connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRequest**](ConnectorRequest.md)|  | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***ProjectConnectorApiCreateProjectScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectConnectorApiCreateProjectScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ConnectorResponse**](ConnectorResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectScopedConnector**
> DeleteProjectScopedConnector(ctx, org, project, connector, optional)
Delete a connector

Deletes the information of the connector with the matching connector slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***ProjectConnectorApiDeleteProjectScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectConnectorApiDeleteProjectScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectScopedConnector**
> ConnectorResponse GetProjectScopedConnector(ctx, org, project, connector, optional)
Retrieve a connector

Retrieves the information of the connector with the matching connector slug. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***ProjectConnectorApiGetProjectScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectConnectorApiGetProjectScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ConnectorResponse**](ConnectorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectScopedConnectors**
> []ConnectorResponse GetProjectScopedConnectors(ctx, org, project, optional)
List connectors

Retrieves the information of the connectors.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
 **optional** | ***ProjectConnectorApiGetProjectScopedConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectConnectorApiGetProjectScopedConnectorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recursive** | **optional.Bool**| Expand current scope to include all child scopes  | [default to false]
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Pagination: Number of items to return | [default to 30]
 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**[]ConnectorResponse**](ConnectorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestProjectScopedConnector**
> ConnectorTestConnectionResponse TestProjectScopedConnector(ctx, org, project, connector, optional)
Test a connector

Tests connection of the connector with the matching connector slug. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***ProjectConnectorApiTestProjectScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectConnectorApiTestProjectScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ConnectorTestConnectionResponse**](ConnectorTestConnectionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectScopedConnector**
> ConnectorResponse UpdateProjectScopedConnector(ctx, body, org, project, connector, optional)
Update a connector

Updates the information of the secret with the matching secret slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRequest**](ConnectorRequest.md)|  | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **project** | **string**| Slug field of the project the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***ProjectConnectorApiUpdateProjectScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectConnectorApiUpdateProjectScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **harnessAccount** | **optional.**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ConnectorResponse**](ConnectorResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

