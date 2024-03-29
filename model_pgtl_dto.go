/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PgtlDto struct for PgtlDto
type PgtlDto struct {
	Id int64 `json:"Id"` 
	Result bool `json:"Result"` 
	TraceId string `json:"TraceId,omitempty"` 
	CaptureId string `json:"CaptureId,omitempty"` 
	IsChargeback bool `json:"IsChargeback"` 
	CreatedDate string `json:"CreatedDate,omitempty"` 
	TransactionId string `json:"TransactionId,omitempty"` 
	InstallmentPlanId int64 `json:"InstallmentPlanId,omitempty"` 
	CompleteResponseXml string `json:"CompleteResponseXml,omitempty"` 
	TerminalGatewayDataId int64 `json:"TerminalGatewayDataId"` 
	AvsMessageMessageCode string `json:"AvsMessageMessageCode,omitempty"` 
	AvsMessageMessageText string `json:"AvsMessageMessageText,omitempty"` 
	CvvMessageMessageCode string `json:"CvvMessageMessageCode,omitempty"` 
	CvvMessageMessageText string `json:"CvvMessageMessageText,omitempty"` 
	RequestedCurrencyCode string `json:"RequestedCurrencyCode,omitempty"` 
	ProcessedAmountAmount float64 `json:"ProcessedAmountAmount"` 
	RequestedAmountAmount float64 `json:"RequestedAmountAmount"` 
	ResultMessageMessageCode string `json:"ResultMessageMessageCode,omitempty"` 
	ResultMessageMessageText string `json:"ResultMessageMessageText,omitempty"` 
	Type OperationType `json:"Type"` 
	ReferencePaymentGatewayTransactionLogId int64 `json:"ReferencePaymentGatewayTransactionLogId,omitempty"` 
}
