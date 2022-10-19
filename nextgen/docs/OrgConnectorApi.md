# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedConnector**](OrgConnectorApi.md#CreateOrgScopedConnector) | **Post** /v1/orgs/{org}/connectors | Create a Connector
[**DeleteOrgScopedConnector**](OrgConnectorApi.md#DeleteOrgScopedConnector) | **Delete** /v1/orgs/{org}/connectors/{connector} | Delete a connector
[**GetOrgScopedConnector**](OrgConnectorApi.md#GetOrgScopedConnector) | **Get** /v1/orgs/{org}/connectors/{connector} | Retrieve a connector
[**GetOrgScopedConnectors**](OrgConnectorApi.md#GetOrgScopedConnectors) | **Get** /v1/orgs/{org}/connectors | List connectors
[**TestOrgScopedConnector**](OrgConnectorApi.md#TestOrgScopedConnector) | **Get** /v1/orgs/{org}/connectors/{connector}/test-connection | Test a connector
[**UpdateOrgScopedConnector**](OrgConnectorApi.md#UpdateOrgScopedConnector) | **Put** /v1/orgs/{org}/connectors/{connector} | Update a connector

# **CreateOrgScopedConnector**
> ConnectorResponse CreateOrgScopedConnector(ctx, body, org, optional)
Create a Connector

Creates a new connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRequest**](ConnectorRequest.md)|  | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
 **optional** | ***OrgConnectorApiCreateOrgScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgConnectorApiCreateOrgScopedConnectorOpts struct
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

# **DeleteOrgScopedConnector**
> DeleteOrgScopedConnector(ctx, org, connector, optional)
Delete a connector

Deletes the information of the connector with the matching connector slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***OrgConnectorApiDeleteOrgScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgConnectorApiDeleteOrgScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedConnector**
> ConnectorResponse GetOrgScopedConnector(ctx, org, connector, optional)
Retrieve a connector

Retrieves the information of the connector with the matching connector slug. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***OrgConnectorApiGetOrgScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgConnectorApiGetOrgScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ConnectorResponse**](ConnectorResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedConnectors**
> []ConnectorResponse GetOrgScopedConnectors(ctx, org, optional)
List connectors

Retrieves the information of the connectors.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
 **optional** | ***OrgConnectorApiGetOrgScopedConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgConnectorApiGetOrgScopedConnectorsOpts struct
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

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestOrgScopedConnector**
> ConnectorTestConnectionResponse TestOrgScopedConnector(ctx, org, connector, optional)
Test a connector

Tests connection of the connector with the matching connector slug. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***OrgConnectorApiTestOrgScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgConnectorApiTestOrgScopedConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Slug field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**ConnectorTestConnectionResponse**](ConnectorTestConnectionResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedConnector**
> ConnectorResponse UpdateOrgScopedConnector(ctx, body, org, connector, optional)
Update a connector

Updates the information of the secret with the matching secret slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRequest**](ConnectorRequest.md)|  | 
  **org** | **string**| Slug field of the organization the resource is scoped to | 
  **connector** | **string**| Connector slug | 
 **optional** | ***OrgConnectorApiUpdateOrgScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgConnectorApiUpdateOrgScopedConnectorOpts struct
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

