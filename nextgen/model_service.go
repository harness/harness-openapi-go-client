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

// This is the Service entity defined in Harness
type Service struct {
	// Account Identifier
	Account string `json:"account,omitempty"`
	// Identifier of the Service Request.
	Identifier string `json:"identifier"`
	// Organization Identifier for the Entity.
	Org string `json:"org,omitempty"`
	// Project Identifier for the Entity.
	Project string `json:"project,omitempty"`
	// Name of the Service Request.
	Name string `json:"name"`
	// Description of the entity
	Description string `json:"description,omitempty"`
	// Service tags
	Tags map[string]string `json:"tags,omitempty"`
	// Yaml related to service
	Yaml string `json:"yaml,omitempty"`
}
