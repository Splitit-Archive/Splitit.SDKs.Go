/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// UpdateInstallmentsPlanResponse struct for UpdateInstallmentsPlanResponse
type UpdateInstallmentsPlanResponse struct {
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"` 
	InstallmentPlan *InstallmentPlan `json:"InstallmentPlan,omitempty"` 
	GatewayTransactionResults []TransactionResult `json:"GatewayTransactionResults,omitempty"` 
	ApprovalUrl string `json:"ApprovalUrl,omitempty"` 
}
