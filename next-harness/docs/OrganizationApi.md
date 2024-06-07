# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganization**](OrganizationApi.md#CreateOrganization) | **Post** /v1/orgs | Create an organization [Beta]
[**DeleteOrganization**](OrganizationApi.md#DeleteOrganization) | **Delete** /v1/orgs/{org} | Delete an organization [Beta]
[**GetOrganization**](OrganizationApi.md#GetOrganization) | **Get** /v1/orgs/{org} | Retrieve an organization [Beta]
[**GetOrganizations**](OrganizationApi.md#GetOrganizations) | **Get** /v1/orgs | List organizations [Beta]
[**UpdateOrganization**](OrganizationApi.md#UpdateOrganization) | **Put** /v1/orgs/{org} | Update an organization [Beta]

# **CreateOrganization**
> OrganizationResponse CreateOrganization(ctx, body, optional)
Create an organization [Beta]

Creates a new organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md)| Post the necessary fields for the API to create an organization. | 
 **optional** | ***OrganizationApiCreateOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiCreateOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrganization**
> OrganizationResponse DeleteOrganization(ctx, org, optional)
Delete an organization [Beta]

Deletes the information of the organization with the matching organization identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationApiDeleteOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiDeleteOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganization**
> OrganizationResponse GetOrganization(ctx, org, optional)
Retrieve an organization [Beta]

Retrieves the information of the organization with the matching organization identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationApiGetOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiGetOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizations**
> []OrganizationResponse GetOrganizations(ctx, optional)
List organizations [Beta]

Retrieves the information of the organizations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OrganizationApiGetOrganizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiGetOrganizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | [**optional.Interface of []string**](string.md)| Identifier field of the organizations the resource is scoped to | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 

### Return type

[**[]OrganizationResponse**](OrganizationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrganization**
> OrganizationResponse UpdateOrganization(ctx, body, org, optional)
Update an organization [Beta]

Updates the information of the organization with the matching organization identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrganizationRequest**](UpdateOrganizationRequest.md)| Put the necessary fields for the API to update a organization. | 
  **org** | **string**| Organization identifier | 
 **optional** | ***OrganizationApiUpdateOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationApiUpdateOrganizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**OrganizationResponse**](OrganizationResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

