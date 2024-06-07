# nextgen{{classname}}

All URIs are relative to *https://app.harness.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountFavorite**](AccountFavoritesApi.md#CreateAccountFavorite) | **Post** /v1/favorites | Create a favorite at account level
[**DeleteAccountFavorite**](AccountFavoritesApi.md#DeleteAccountFavorite) | **Delete** /v1/favorites/{user-id} | Delete a account level favorite
[**GetAccountFavorites**](AccountFavoritesApi.md#GetAccountFavorites) | **Get** /v1/favorites/{user-id} | Retrieve a favorite resources of a user at account level

# **CreateAccountFavorite**
> FavoriteResponse CreateAccountFavorite(ctx, optional)
Create a favorite at account level

Creates favorite resource of a user at account level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountFavoritesApiCreateAccountFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountFavoritesApiCreateAccountFavoriteOpts struct
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

# **DeleteAccountFavorite**
> DeleteAccountFavorite(ctx, userId, optional)
Delete a account level favorite

Deletes the requested favorite at account level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This would be used to fetch the Favorites of the user. | 
 **optional** | ***AccountFavoritesApiDeleteAccountFavoriteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountFavoritesApiDeleteAccountFavoriteOpts struct
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

# **GetAccountFavorites**
> []FavoriteResponse GetAccountFavorites(ctx, userId, optional)
Retrieve a favorite resources of a user at account level

Retrieves the favorite resource of a user at account level

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| This would be used to fetch the Favorites of the user. | 
 **optional** | ***AccountFavoritesApiGetAccountFavoritesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountFavoritesApiGetAccountFavoritesOpts struct
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

