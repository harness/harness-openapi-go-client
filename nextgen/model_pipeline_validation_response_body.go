/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Pipeline Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Pipeline Validation Response Body
type PipelineValidationResponseBody struct {
	// Status of the Pipeline Validation Event
	Status string `json:"status,omitempty"`
	// Result of Policy Evaluations on the Pipeline
	PolicyEval *interface{} `json:"policy_eval,omitempty"`
	// Start time of the Evaluation
	StartTs int64 `json:"start_ts,omitempty"`
	// End time of the Evaluation
	EndTs                      int64                           `json:"end_ts,omitempty"`
	TemplateValidationResponse *TemplateValidationResponseBody `json:"template_validation_response,omitempty"`
}