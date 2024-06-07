# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgScopedCertificates**](OrgCertificateApi.md#CreateOrgScopedCertificates) | **Post** /v1/orgs/{org}/certificates | Creates a new Certificate in organization scope
[**DeleteOrgScopedCertificate**](OrgCertificateApi.md#DeleteOrgScopedCertificate) | **Delete** /v1/orgs/{org}/certificates/{certificate} | Deletes the specified Certificate from org scope
[**GetOrgScopedCertificate**](OrgCertificateApi.md#GetOrgScopedCertificate) | **Get** /v1/orgs/{org}/certificates/{certificate} | Retrieves the specified Certificate in organization scope
[**GetOrgScopedCertificates**](OrgCertificateApi.md#GetOrgScopedCertificates) | **Get** /v1/orgs/{org}/certificates | Retrieves the list of Certificates in organization scope
[**UpdateOrgScopedCertificate**](OrgCertificateApi.md#UpdateOrgScopedCertificate) | **Put** /v1/orgs/{org}/certificates/{certificate} | Updates the specified Certificate in org scope
[**ValidateOrgScopedCertificateIdentifier**](OrgCertificateApi.md#ValidateOrgScopedCertificateIdentifier) | **Get** /v1/orgs/{org}/certificates/validate-unique-identifier/{certificate} | Validate if the specified Certificate identifier is available for use in organization scope

# **CreateOrgScopedCertificates**
> CertificateResponseDto CreateOrgScopedCertificates(ctx, org, optional)
Creates a new Certificate in organization scope

Creates a new certificate in organization scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgCertificateApiCreateOrgScopedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgCertificateApiCreateOrgScopedCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CertificateDto**](CertificateDto.md)|  | 
 **spec** | [**optional.Interface of CertificateDto**](.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**CertificateResponseDto**](CertificateResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgScopedCertificate**
> DeleteOrgScopedCertificate(ctx, org, certificate, optional)
Deletes the specified Certificate from org scope

Deletes the specified Certificate from org scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***OrgCertificateApiDeleteOrgScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgCertificateApiDeleteOrgScopedCertificateOpts struct
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
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedCertificate**
> CertificateResponseDto GetOrgScopedCertificate(ctx, org, certificate, optional)
Retrieves the specified Certificate in organization scope

Retrieves the specified certificate in org scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***OrgCertificateApiGetOrgScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgCertificateApiGetOrgScopedCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**CertificateResponseDto**](CertificateResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgScopedCertificates**
> []CertificateResponseDto GetOrgScopedCertificates(ctx, org, optional)
Retrieves the list of Certificates in organization scope

Retrieves the list of certificates in organization scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
 **optional** | ***OrgCertificateApiGetOrgScopedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgCertificateApiGetOrgScopedCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **page** | **optional.Int32**| Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page  | [default to 0]
 **limit** | **optional.Int32**| Number of items to return per page. | [default to 20]
 **sort** | **optional.String**| Parameter on the basis of which sorting is done. | 
 **order** | **optional.String**| Order on the basis of which sorting is done. | 
 **includeAllAvailableAtScope** | **optional.Bool**| Include all entities accessible at current scope | [default to false]

### Return type

[**[]CertificateResponseDto**](CertificateResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgScopedCertificate**
> CertificateResponseDto UpdateOrgScopedCertificate(ctx, org, certificate, optional)
Updates the specified Certificate in org scope

Updates the specified Certificate in org scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***OrgCertificateApiUpdateOrgScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgCertificateApiUpdateOrgScopedCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of CertificateDto**](CertificateDto.md)|  | 
 **spec** | [**optional.Interface of CertificateDto**](.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**CertificateResponseDto**](CertificateResponseDTO.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateOrgScopedCertificateIdentifier**
> bool ValidateOrgScopedCertificateIdentifier(ctx, org, certificate, optional)
Validate if the specified Certificate identifier is available for use in organization scope

Validate if the specified certificate identifier is available for use in organization scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***OrgCertificateApiValidateOrgScopedCertificateIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgCertificateApiValidateOrgScopedCertificateIdentifierOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

**bool**

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

