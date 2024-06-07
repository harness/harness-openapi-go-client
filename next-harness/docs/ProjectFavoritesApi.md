# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectFavorite**](ProjectFavoritesApi.md#CreateProjectFavorite) | **Post** /v1/orgs/{org}/projects/{project}/favorites | Create a favorite at project level
[**DeleteProjectFavorite**](ProjectFavoritesApi.md#DeleteProjectFavorite) | **Delete** /v1/orgs/{org}/projects/{project}/favorites/{user-id} | Delete a project level favorite
[**GetProjectFavorites**](ProjectFavoritesApi.md#GetProjectFavorites) | **Get** /v1/orgs/{org}/projects/{project}/favorites/{user-id} | Retrieve a favorite resources of a user at project level

# **CreateProjectFavorite**
> FavoriteResponse CreateProjectFavorite(ctx, org, project, optional)
Create a favorite at project level

Creates favorite resource of a user at project level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
 **optional** | ***ProjectFavoritesApiCreateProjectFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectFavoritesApiCreateProjectFavoriteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of FavoriteDto**](FavoriteDto.md)|  | 
 **harnessAccount** | **optional.**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 

### Return type

[**FavoriteResponse**](FavoriteResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectFavorite**
> DeleteProjectFavorite(ctx, org, project, userId, optional)
Delete a project level favorite

Deletes the requested favorite at project level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **userId** | **string**| This would be used to fetch the Favorites of the user. | 
 **optional** | ***ProjectFavoritesApiDeleteProjectFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectFavoritesApiDeleteProjectFavoriteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **resourceType** | **optional.String**| Determines the type of favorite entity requested. | 
 **resourceId** | **optional.String**| This would be used to do operations on the favorite entity. | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectFavorites**
> []FavoriteResponse GetProjectFavorites(ctx, org, project, userId, optional)
Retrieve a favorite resources of a user at project level

Retrieves the favorite resource of a user at project level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **org** | **string**| Identifier field of the organization the resource is scoped to | 
  **project** | **string**| Identifier field of the project the resource is scoped to | 
  **userId** | **string**| This would be used to fetch the Favorites of the user. | 
 **optional** | ***ProjectFavoritesApiGetProjectFavoritesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectFavoritesApiGetProjectFavoritesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **harnessAccount** | **optional.String**| Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped. | 
 **resourceType** | **optional.String**| Determines the type of favorite entity requested. | 

### Return type

[**[]FavoriteResponse**](FavoriteResponse.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

