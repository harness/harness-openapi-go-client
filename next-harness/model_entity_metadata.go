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

// Entity metadata
type EntityMetadata struct {
	PipelineIdentifier string `json:"pipeline_identifier,omitempty"`
	TemplateVersionLabel string `json:"template_version_label,omitempty"`
	EnvironmentIdentifier string `json:"environment_identifier,omitempty"`
	OverrideType string `json:"override_type,omitempty"`
}
