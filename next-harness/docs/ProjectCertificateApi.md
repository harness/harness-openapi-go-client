# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectScopedCertificates**](ProjectCertificateApi.md#CreateProjectScopedCertificates) | **Post** /v1/orgs/{org}/projects/{project}/certificates | Creates a new Certificate in project scope
[**DeleteProjectScopedCertificate**](ProjectCertificateApi.md#DeleteProjectScopedCertificate) | **Delete** /v1/orgs/{org}/projects/{project}/certificates/{certificate} | Deletes the specified Certificate from project scope
[**GetProjectScopedCertificate**](ProjectCertificateApi.md#GetProjectScopedCertificate) | **Get** /v1/orgs/{org}/projects/{project}/certificates/{certificate} | Retrieves the specified Certificate in project scope
[**GetProjectScopedCertificates**](ProjectCertificateApi.md#GetProjectScopedCertificates) | **Get** /v1/orgs/{org}/projects/{project}/certificates | Retrieves the list of Certificates in project scope
[**UpdateProjectScopedCertificate**](ProjectCertificateApi.md#UpdateProjectScopedCertificate) | **Put** /v1/orgs/{org}/projects/{project}/certificates/{certificate} | Updates the specified Certificate in project scope
[**ValidateProjectScopedCertificateIdentifier**](ProjectCertificateApi.md#ValidateProjectScopedCertificateIdentifier) | **Get** /v1/orgs/{org}/projects/{project}/certificates/validate-unique-identifier/{certificate} | Validate if the specified Certificate identifier is available for use in project scope

# **CreateProjectScopedCertificates**
> CertificateResponseDto CreateProjectScopedCertificates(ctx, org, project, optional)
Creates a new Certificate in project scope

Creates a new certificate in project scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectCertificateApiCreateProjectScopedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCertificateApiCreateProjectScopedCertificatesOpts struct
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

# **DeleteProjectScopedCertificate**
> DeleteProjectScopedCertificate(ctx, certificate, org, project, optional)
Deletes the specified Certificate from project scope

Deletes the specified Certificate from project scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectCertificateApiDeleteProjectScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCertificateApiDeleteProjectScopedCertificateOpts struct
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

# **GetProjectScopedCertificate**
> CertificateResponseDto GetProjectScopedCertificate(ctx, certificate, org, project, optional)
Retrieves the specified Certificate in project scope

Retrieves the specified certificate in project scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectCertificateApiGetProjectScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCertificateApiGetProjectScopedCertificateOpts struct
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

# **GetProjectScopedCertificates**
> []CertificateResponseDto GetProjectScopedCertificates(ctx, org, project, optional)
Retrieves the list of Certificates in project scope

Retrieves the list of certificates in project scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectCertificateApiGetProjectScopedCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCertificateApiGetProjectScopedCertificatesOpts struct
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

# **UpdateProjectScopedCertificate**
> CertificateResponseDto UpdateProjectScopedCertificate(ctx, certificate, org, project, optional)
Updates the specified Certificate in project scope

Updates the specified Certificate in project scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Identifier field of the certificate | 
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectCertificateApiUpdateProjectScopedCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCertificateApiUpdateProjectScopedCertificateOpts struct
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

# **ValidateProjectScopedCertificateIdentifier**
> bool ValidateProjectScopedCertificateIdentifier(ctx, org, project, certificate, optional)
Validate if the specified Certificate identifier is available for use in project scope

Validate if the specified certificate identifier is available for use in project scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **certificate** | **string**| Identifier field of the certificate | 
 **optional** | ***ProjectCertificateApiValidateProjectScopedCertificateIdentifierOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectCertificateApiValidateProjectScopedCertificateIdentifierOpts struct
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

