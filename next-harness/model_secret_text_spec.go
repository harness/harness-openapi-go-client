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

// This is the SSH key authentication details defined in Harness.
type SecretTextSpec struct {
	// This specifies the type of secret
	Type_ string `json:"type"`
	// Identifier of the secret manager used to manage the secret
	SecretManagerIdentifier string `json:"secret_manager_identifier"`
	// This has details to specify if the secret value is inline or referenced
	ValueType string `json:"value_type"`
	// Value of the Secret
	Value string `json:"value,omitempty"`
	AdditionalMetadata *SecretTextSpecAdditionalMetadata `json:"additional_metadata,omitempty"`
}
