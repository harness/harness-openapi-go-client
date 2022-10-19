# {{classname}}

All URIs are relative to *http://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedConnector**](AccountConnectorApi.md#CreateAccountScopedConnector) | **Post** /v1/connectors | Create a Connector
[**DeleteAccountScopedConnector**](AccountConnectorApi.md#DeleteAccountScopedConnector) | **Delete** /v1/connectors/{connector} | Delete a connector
[**GetAccountScopedConnector**](AccountConnectorApi.md#GetAccountScopedConnector) | **Get** /v1/connectors/{connector} | Retrieve a connector
[**GetAccountScopedConnectors**](AccountConnectorApi.md#GetAccountScopedConnectors) | **Get** /v1/connectors | List connectors
[**TestAccountScopedConnector**](AccountConnectorApi.md#TestAccountScopedConnector) | **Get** /v1/connectors/{connector}/test-connection | Test a connector
[**UpdateAccountScopedConnector**](AccountConnectorApi.md#UpdateAccountScopedConnector) | **Put** /v1/connectors/{connector} | Update a connector

# **CreateAccountScopedConnector**
> ConnectorResponse CreateAccountScopedConnector(ctx, body, optional)
Create a Connector

Creates a new connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRequest**](ConnectorRequest.md)|  | 
 **optional** | ***AccountConnectorApiCreateAccountScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountConnectorApiCreateAccountScopedConnectorOpts struct
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

# **DeleteAccountScopedConnector**
> DeleteAccountScopedConnector(ctx, connector, optional)
Delete a connector

Deletes the information of the connector with the matching connector slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| Connector slug | 
 **optional** | ***AccountConnectorApiDeleteAccountScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountConnectorApiDeleteAccountScopedConnectorOpts struct
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

# **GetAccountScopedConnector**
> ConnectorResponse GetAccountScopedConnector(ctx, connector, optional)
Retrieve a connector

Retrieves the information of the connector with the matching connector slug. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| Connector slug | 
 **optional** | ***AccountConnectorApiGetAccountScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountConnectorApiGetAccountScopedConnectorOpts struct
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

# **GetAccountScopedConnectors**
> []ConnectorResponse GetAccountScopedConnectors(ctx, optional)
List connectors

Retrieves the information of the connectors.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountConnectorApiGetAccountScopedConnectorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountConnectorApiGetAccountScopedConnectorsOpts struct
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

# **TestAccountScopedConnector**
> ConnectorTestConnectionResponse TestAccountScopedConnector(ctx, connector, optional)
Test a connector

Tests connection of the connector with the matching connector slug. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | **string**| Connector slug | 
 **optional** | ***AccountConnectorApiTestAccountScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountConnectorApiTestAccountScopedConnectorOpts struct
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

# **UpdateAccountScopedConnector**
> ConnectorResponse UpdateAccountScopedConnector(ctx, body, connector, optional)
Update a connector

Updates the information of the secret with the matching secret slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRequest**](ConnectorRequest.md)|  | 
  **connector** | **string**| Connector slug | 
 **optional** | ***AccountConnectorApiUpdateAccountScopedConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountConnectorApiUpdateAccountScopedConnectorOpts struct
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

