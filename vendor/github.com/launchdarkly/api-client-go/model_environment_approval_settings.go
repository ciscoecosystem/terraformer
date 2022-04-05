/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.3.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

// Approval settings for an environment. Only appears if the approvals feature is enabled.
type EnvironmentApprovalSettings struct {
	// The approvals system used.
	ServiceKind string `json:"serviceKind,omitempty"`
	// Whether any changes to flags in this environment will require approval. You may only set required or requiredApprovalTags, not both.
	Required bool `json:"required,omitempty"`
	// Whether requesters can approve or decline their own request. They may always comment.
	CanReviewOwnRequest bool `json:"canReviewOwnRequest,omitempty"`
	// The number of approvals required before an approval request can be applied.
	MinNumApprovals int64 `json:"minNumApprovals,omitempty"`
	// Whether changes can be applied as long as minNumApprovals is met, regardless of if any reviewers have declined a request.
	CanApplyDeclinedChanges bool `json:"canApplyDeclinedChanges,omitempty"`
	// An array of tags used to specify which flags with those tags require approval. You may only set requiredApprovalTags or required, not both.
	RequiredApprovalTags []string `json:"requiredApprovalTags,omitempty"`
}
