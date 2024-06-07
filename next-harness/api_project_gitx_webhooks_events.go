
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

type ProjectGitxWebhooksEventsApiService service
/*
ProjectGitxWebhooksEventsApiService Lists all the GitX Webhooks events at project level
List GitX webhook Events at project level
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param org Identifier field of the organization the resource is scoped to
 * @param project Identifier field of the project the resource is scoped to
 * @param optional nil or *ProjectGitxWebhooksEventsApiListProjectGitxWebhookEventsOpts - Optional Parameters:
     * @param "HarnessAccount" (optional.String) -  Identifier field of the account the resource is scoped to. This is required for Authorization methods other than the x-api-key header. If you are using the x-api-key header, this can be skipped.
     * @param "Page" (optional.Int32) -  Pagination page number strategy: Specify the page number within the paginated collection related to the number of items in each page 
     * @param "Limit" (optional.Int32) -  Number of items to return per page.
     * @param "WebhookIdentifier" (optional.String) - 
     * @param "EventStartTime" (optional.Int64) - 
     * @param "EventEndTime" (optional.Int64) - 
     * @param "RepoName" (optional.String) - 
     * @param "FilePath" (optional.String) - 
     * @param "EventIdentifier" (optional.String) - 
     * @param "EventStatus" (optional.Interface of []string) - 
     * @param "ConnectorRef" (optional.String) - 
     * @param "IncludeParentScope" (optional.Bool) - 
     * @param "CommitId" (optional.String) - 
@return []GitXWebhookEventResponse
*/

type ProjectGitxWebhooksEventsApiListProjectGitxWebhookEventsOpts struct {
    HarnessAccount optional.String
    Page optional.Int32
    Limit optional.Int32
    WebhookIdentifier optional.String
    EventStartTime optional.Int64
    EventEndTime optional.Int64
    RepoName optional.String
    FilePath optional.String
    EventIdentifier optional.String
    EventStatus optional.Interface
    ConnectorRef optional.String
    IncludeParentScope optional.Bool
    CommitId optional.String
}

func (a *ProjectGitxWebhooksEventsApiService) ListProjectGitxWebhookEvents(ctx context.Context, org string, project string, localVarOptionals *ProjectGitxWebhooksEventsApiListProjectGitxWebhookEventsOpts) ([]GitXWebhookEventResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []GitXWebhookEventResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/orgs/{org}/projects/{project}/gitx-webhook-events"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", fmt.Sprintf("%v", org), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project"+"}", fmt.Sprintf("%v", project), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WebhookIdentifier.IsSet() {
		localVarQueryParams.Add("webhook_identifier", parameterToString(localVarOptionals.WebhookIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventStartTime.IsSet() {
		localVarQueryParams.Add("event_start_time", parameterToString(localVarOptionals.EventStartTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventEndTime.IsSet() {
		localVarQueryParams.Add("event_end_time", parameterToString(localVarOptionals.EventEndTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoName.IsSet() {
		localVarQueryParams.Add("repo_name", parameterToString(localVarOptionals.RepoName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FilePath.IsSet() {
		localVarQueryParams.Add("file_path", parameterToString(localVarOptionals.FilePath.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventIdentifier.IsSet() {
		localVarQueryParams.Add("event_identifier", parameterToString(localVarOptionals.EventIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventStatus.IsSet() {
		localVarQueryParams.Add("event_status", parameterToString(localVarOptionals.EventStatus.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ConnectorRef.IsSet() {
		localVarQueryParams.Add("connector_ref", parameterToString(localVarOptionals.ConnectorRef.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeParentScope.IsSet() {
		localVarQueryParams.Add("include_parent_scope", parameterToString(localVarOptionals.IncludeParentScope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CommitId.IsSet() {
		localVarQueryParams.Add("commit_id", parameterToString(localVarOptionals.CommitId.Value(), ""))
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
			var v []GitXWebhookEventResponse
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
