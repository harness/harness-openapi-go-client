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

// Contains information about the GitX webhook creation request
type CreateGitXWebhookRequest struct {
	WebhookIdentifier string `json:"webhook_identifier,omitempty"`
	RepoName string `json:"repo_name,omitempty"`
	ConnectorRef string `json:"connector_ref,omitempty"`
	FolderPaths []string `json:"folder_paths,omitempty"`
	WebhookName string `json:"webhook_name,omitempty"`
}
