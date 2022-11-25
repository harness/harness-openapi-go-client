/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

import (
	"fmt"

	"github.com/harness/harness-openapi-go-client/harness"
	"github.com/harness/harness-openapi-go-client/helpers"
	retryablehttp "github.com/hashicorp/go-retryablehttp"

	"github.com/harness/harness-openapi-go-client/logging"
	"github.com/harness/harness-openapi-go-client/utils"
	log "github.com/sirupsen/logrus"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	AccountId     string            `json:"accountId,omitempty"`
	ApiKey        string            `json:"apiKey,omitempty"`
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *retryablehttp.Client
	Logger        *log.Logger
	DebugLogging  bool
}

func NewConfiguration() *Configuration {
	logger := logging.NewLogger()
	if helpers.EnvVars.TfLog.Get() == "DEBUG" {
		logger.SetLevel(log.DebugLevel)
	}

	cfg := &Configuration{
		AccountId:     helpers.EnvVars.AccountId.Get(),
		ApiKey:        helpers.EnvVars.PlatformApiKey.Get(),
		BasePath:      helpers.EnvVars.Endpoint.GetWithDefault(utils.BaseUrl),
		DefaultHeader: make(map[string]string),
		HTTPClient:    utils.GetDefaultHttpClient(logger),
		Logger:        logger,
		UserAgent:     fmt.Sprintf("%s-%s", harness.SDKName, harness.SDKVersion),
	}

	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
