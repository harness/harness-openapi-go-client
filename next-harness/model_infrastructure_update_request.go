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

// Infrastructure Update Request Body 
type InfrastructureUpdateRequest struct {
	// Identifier of the Infrastructure
	Identifier string `json:"identifier"`
	// Name of the Infrastructure
	Name string `json:"name"`
	// Description of the Infrastructure
	Description string `json:"description,omitempty"`
	// Infrastructure tags
	Tags map[string]string `json:"tags,omitempty"`
	Type_ *InfrastructureType `json:"type"`
	// version of harness yaml
	HarnessVersion string `json:"harness_version,omitempty"`
	// YAML for the Infrastructure Request
	Yaml string `json:"yaml"`
}
