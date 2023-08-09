/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Pipeline Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Contains the UUID of the Pipeline Validation Event that's started
type PipelineValidationUuidResponseBody struct {
	// UUID of the Pipeline Validation Event
	Uuid string `json:"uuid,omitempty"`
}