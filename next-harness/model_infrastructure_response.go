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

// Default response when a infrastructure is returned
type InfrastructureResponse struct {
	Infrastructure *Infrastructure `json:"infrastructure,omitempty"`
	// Creation timestamp for infrastructure.
	Created int64 `json:"created,omitempty"`
	// Last modification timestamp for infrastructure.
	Updated int64 `json:"updated,omitempty"`
}
