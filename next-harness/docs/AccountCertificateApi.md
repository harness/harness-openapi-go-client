# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountScopedCertificates**](AccountCertificateApi.md#CreateAccountScopedCertificates) | **Post** /v1/certificates | Creates a new Certificate in account scope
[**DeleteAccountScopedCertificate**](AccountCertificateApi.md#DeleteAccountScopedCertificate) | **Delete** /v1/certificates/{certificate} | Deletes the specified Certificate from account scope
[**GetAccountScopedCertificate**](AccountCertificateApi.md#GetAccountScopedCertificate) | **Get** /v1/certificates/{certificate} | Retrieves the specified Certificate in account scope
[**GetAccountScopedCertificates**](AccountCertificateApi.md#GetAccountScopedCertificates) | **Get** /v1/certificates | Retrieves the list of Certificates in account scope
[**UpdateAccountScopedCertificate**](AccountCertificateApi.md#UpdateAccountScopedCertificate) | **Put** /v1/certificates/{certificate} | Updates the specified Certificate in account scope
[**ValidateAccountScopedCertificateIdentifier**](AccountCertificateApi.md#ValidateAccountScopedCertificateIdentifier) | **Get** /v1/certificates/validate-unique-identifier/{certificate} | Validate if the specified Certificate identifier is available for use in account scope

# **CreateAccountScopedCertificates**
> CertificateResponseDto CreateAccountScopedCertificates(ctx, optional)
Creates a new Certificate in account scope

Creates a new certificate in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountCertificateApiCreateAccountScopedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCertificateApiCreateAccountScopedCertificatesOpts struct
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

# **DeleteAccountScopedCertificate**
> DeleteAccountScopedCertificate(ctx, certificate, optional)
Deletes the specified Certificate from account scope

Deletes the specified Certificate from account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***AccountCertificateApiDeleteAccountScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCertificateApiDeleteAccountScopedCertificateOpts struct
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

# **GetAccountScopedCertificate**
> CertificateResponseDto GetAccountScopedCertificate(ctx, certificate, optional)
Retrieves the specified Certificate in account scope

Retrieves the specified certificate in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***AccountCertificateApiGetAccountScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCertificateApiGetAccountScopedCertificateOpts struct
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

# **GetAccountScopedCertificates**
> []CertificateResponseDto GetAccountScopedCertificates(ctx, optional)
Retrieves the list of Certificates in account scope

Retrieves the list of certificates in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountCertificateApiGetAccountScopedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCertificateApiGetAccountScopedCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **searchTerm** | **optional.String**| This would be used to filter resources having attributes matching with search term. | 
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

# **UpdateAccountScopedCertificate**
> CertificateResponseDto UpdateAccountScopedCertificate(ctx, certificate, optional)
Updates the specified Certificate in account scope

Updates the specified Certificate in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***AccountCertificateApiUpdateAccountScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCertificateApiUpdateAccountScopedCertificateOpts struct
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

# **ValidateAccountScopedCertificateIdentifier**
> bool ValidateAccountScopedCertificateIdentifier(ctx, certificate, optional)
Validate if the specified Certificate identifier is available for use in account scope

Validate if the specified certificate identifier is available for use in account scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***AccountCertificateApiValidateAccountScopedCertificateIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountCertificateApiValidateAccountScopedCertificateIdentifierOpts struct
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

