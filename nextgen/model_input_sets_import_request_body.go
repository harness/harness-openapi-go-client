package nextgen

// InputSets import request body
type InputSetsImportRequestBody struct {
	GitImportInfo         *GitImportInfo            `json:"git_import_info,omitempty"`
	InputSetsImportRequest *InputSetsImportRequestDto `json:"input_sets_import_request,omitempty"`
}