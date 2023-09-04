package nextgen

// Response body for InputSet save.
type InputSetSaveResponseBody struct {
	Identifier         string               `json:"identifier,omitempty"`
	GovernanceMetadata []GovernanceMetadata `json:"governance_metadata,omitempty"`
}