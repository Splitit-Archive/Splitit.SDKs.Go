/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PaymentWizardData struct for PaymentWizardData
type PaymentWizardData struct {
	RequestedNumberOfInstallments string `json:"RequestedNumberOfInstallments,omitempty"` 
	SuccessExitURL string `json:"SuccessExitURL,omitempty"` 
	ErrorExitURL string `json:"ErrorExitURL,omitempty"` 
	CancelExitURL string `json:"CancelExitURL,omitempty"` 
	SuccessAsyncUrl string `json:"SuccessAsyncUrl,omitempty"` 
	ViewName string `json:"ViewName,omitempty"` 
	IsOpenedInIframe bool `json:"IsOpenedInIframe"` 
	Is3dSecureInPopup bool `json:"Is3dSecureInPopup,omitempty"` 
	PaymentFormMessage string `json:"PaymentFormMessage,omitempty"` 
	SetShortUrl bool `json:"SetShortUrl"` 
}
