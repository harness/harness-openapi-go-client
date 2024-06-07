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

// Default response when a environment is returned
type EnvironmentResponse struct {
	Environment *Environment `json:"environment,omitempty"`
	// Creation timestamp for environment.
	Created int64 `json:"created,omitempty"`
	// Last modification timestamp for environment.
	Updated int64 `json:"updated,omitempty"`
}
