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

// Contains parameters related to Fetching an Entity for Git Experience.
type GitFindDetails struct {
	// Name of the branch.
	BranchName string `json:"branch_name,omitempty"`
	// Connector ref of parent template if its remote
	ParentEntityConnectorRef string `json:"parent_entity_connector_ref,omitempty"`
	// Repo name of parent template if its remote
	ParentEntityRepoName string `json:"parent_entity_repo_name,omitempty"`
	// Account name of parent template if its remote
	ParentEntityAccountId string `json:"parent_entity_account_id,omitempty"`
	// Organization name of parent template if its remote
	ParentEntityOrgId string `json:"parent_entity_org_id,omitempty"`
	// Project name of parent entity if its remote
	ParentEntityProjectId string `json:"parent_entity_project_id,omitempty"`
	// Specify comment with respect to changes
	Comments string `json:"comments,omitempty"`
	// Is harness code repo enabled
	IsHarnessCodeRepo bool `json:"is_harness_code_repo,omitempty"`
}
