/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// CreateInstallmentPlanLegacyResponse struct for CreateInstallmentPlanLegacyResponse
type CreateInstallmentPlanLegacyResponse struct {
	ApiKey string `json:"ApiKey,omitempty"` 
	InstallmentPlanStatus int32 `json:"InstallmentPlanStatus"` 
	Result int32 `json:"Result"` 
	PaymentGateway string `json:"PaymentGateway,omitempty"` 
	Email string `json:"Email,omitempty"` 
	ConsumerFullName string `json:"ConsumerFullName,omitempty"` 
	ParamX string `json:"ParamX,omitempty"` 
	InstallmentNumber int32 `json:"InstallmentNumber"` 
	Amount float64 `json:"Amount"` 
	CurrencyName string `json:"CurrencyName,omitempty"` 
	CurrencySymbol string `json:"CurrencySymbol,omitempty"` 
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"` 
	ResponseHeader *ResponseHeader `json:"ResponseHeader,omitempty"` 
}
