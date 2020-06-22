/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// InitiateInstallmentsPlanResponse struct for InitiateInstallmentsPlanResponse
type InitiateInstallmentsPlanResponse struct {
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"`
	InstallmentPlan *InstallmentPlan `json:"InstallmentPlan,omitempty"`
	CheckoutUrl string `json:"CheckoutUrl,omitempty"`
	ApprovalUrl string `json:"ApprovalUrl,omitempty"`
	TermsAndConditionsUrl string `json:"TermsAndConditionsUrl,omitempty"`
	PrivacyPolicyUrl string `json:"PrivacyPolicyUrl,omitempty"`
	InstallmentPlanInfoUrl string `json:"InstallmentPlanInfoUrl,omitempty"`
	PublicToken string `json:"PublicToken,omitempty"`
	LearnMoreUrl string `json:"LearnMoreUrl,omitempty"`
}