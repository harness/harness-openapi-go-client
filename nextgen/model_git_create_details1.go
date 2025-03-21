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

// Contains parameters related to creating an Entity for Git Experience.
type GitCreateDetails1 struct {
	// Name of the branch.
	BranchName string `json:"branch_name,omitempty"`
	// File path of the Entity in the repository.
	FilePath string `json:"file_path,omitempty"`
	// Commit message used for the merge commit.
	CommitMessage string `json:"commit_message,omitempty"`
	// Name of the default branch (this checks out a new branch titled by branch_name).
	BaseBranch string `json:"base_branch,omitempty"`
	// Identifier of the Harness Connector used for CRUD operations on the Entity.
	ConnectorRef string `json:"connector_ref,omitempty"`
	// Specifies whether the Entity is to be stored in Git or not.
	StoreType string `json:"store_type,omitempty"`
	// Name of the repository.
	RepoName string `json:"repo_name,omitempty"`
	// Is harness code repo enabled
	IsHarnessCodeRepo bool `json:"is_harness_code_repo,omitempty"`
}
