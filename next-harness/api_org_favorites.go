
/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type OrgFavoritesApiService service
/*
OrgFavoritesApiService Create a favorite at org level
Creates favorite resource of a user at organisation level
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Identifier field of the organization the resource is scoped to
 * @param optional nil or *OrgFavoritesApiCreateOrgFavoriteOpts - Optional Parameters:
     * @param "Body" (optional.Interface of FavoriteDto) - 
     * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
@return FavoriteResponse
*/

type OrgFavoritesApiCreateOrgFavoriteOpts struct {
    Body optional.Interface
    HarnessAccount optional.String
}

func (a *OrgFavoritesApiService) CreateOrgFavorite(ctx context.Context, org string, localVarOptionals *OrgFavoritesApiCreateOrgFavoriteOpts) (FavoriteResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue FavoriteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/favorites"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v FavoriteResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
OrgFavoritesApiService Delete a org level favorite
Deletes the requested favorite at organization level.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Identifier field of the organization the resource is scoped to
 * @param userId This would be used to fetch the Favorites of the user.
 * @param optional nil or *OrgFavoritesApiDeleteOrgFavoriteOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "ResourceType" (optional.String) -  Determines the type of favorite entity requested.
     * @param "ResourceId" (optional.String) -  This would be used to do operations on the favorite entity.

*/

type OrgFavoritesApiDeleteOrgFavoriteOpts struct {
    HarnessAccount optional.String
    ResourceType optional.String
    ResourceId optional.String
}

func (a *OrgFavoritesApiService) DeleteOrgFavorite(ctx context.Context, org string, userId string, localVarOptionals *OrgFavoritesApiDeleteOrgFavoriteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/favorites/{user-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resource_type", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceId.IsSet() {
		localVarQueryParams.Add("resource_id", parameterToString(localVarOptionals.ResourceId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
OrgFavoritesApiService Retrieve a favorite resources of a user at organization level
Retrieves the favorite resource of a user at organization level
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Identifier field of the organization the resource is scoped to
 * @param userId This would be used to fetch the Favorites of the user.
 * @param optional nil or *OrgFavoritesApiGetOrgFavoritesOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "ResourceType" (optional.String) -  Determines the type of favorite entity requested.
@return []FavoriteResponse
*/

type OrgFavoritesApiGetOrgFavoritesOpts struct {
    HarnessAccount optional.String
    ResourceType optional.String
}

func (a *OrgFavoritesApiService) GetOrgFavorites(ctx context.Context, org string, userId string, localVarOptionals *OrgFavoritesApiGetOrgFavoritesOpts) ([]FavoriteResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []FavoriteResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/favorites/{user-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user-id"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resource_type", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.HarnessAccount.IsSet() {
		localVarHeaderParams["Harness-Account"] = parameterToString(localVarOptionals.HarnessAccount.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []FavoriteResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
