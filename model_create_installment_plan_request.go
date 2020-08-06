/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// CreateInstallmentPlanRequest struct for CreateInstallmentPlanRequest
type CreateInstallmentPlanRequest struct {
	InstallmentPlanNumber string `json:"InstallmentPlanNumber,omitempty"` 
	PlanData *PlanData `json:"PlanData,omitempty"` 
	CartData *CartData `json:"CartData,omitempty"` 
	ConsumerData *ConsumerData `json:"ConsumerData,omitempty"` 
	BillingAddress *AddressData `json:"BillingAddress,omitempty"` 
	CreditCardDetails *CardData `json:"CreditCardDetails,omitempty"` 
	PaymentToken *PaymentToken `json:"PaymentToken,omitempty"` 
	PlanApprovalEvidence *PlanApprovalEvidence `json:"PlanApprovalEvidence,omitempty"` 
	RedirectUrls *RedirectUrls `json:"RedirectUrls,omitempty"` 
	EventsEndpoints *EventsEndpoints `json:"EventsEndpoints,omitempty"` 
	ExternalAuth *ExternalAuth `json:"ExternalAuth,omitempty"` 
}
