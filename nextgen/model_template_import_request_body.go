package nextgen

// Templates import request body
type TemplatesImportRequestBody struct {
	GitImportInfo         *GitImportInfo            `json:"git_import_info,omitempty"`
	TemplatesImportRequest *TemplatesImportRequestDto `json:"template_import_request,omitempty"`
}