/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// CreateInstallmentsPlanResponse struct for CreateInstallmentsPlanResponse
type CreateInstallmentsPlanResponse struct {
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"` 
	InstallmentPlan *InstallmentPlan `json:"InstallmentPlan,omitempty"` 
	ApprovalUrl string `json:"ApprovalUrl,omitempty"` 
}
