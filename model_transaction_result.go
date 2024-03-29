/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
import (
	
)
// TransactionResult struct for TransactionResult
type TransactionResult struct {
	GatewayTransactionId string `json:"GatewayTransactionId,omitempty"` 
	SplititTransactionId int64 `json:"SplititTransactionId"` 
	GatewayResultCode string `json:"GatewayResultCode,omitempty"` 
	GatewayResultMessage string `json:"GatewayResultMessage,omitempty"` 
	OperationType *ReferenceEntityBase `json:"OperationType,omitempty"` 
	GatewayResult bool `json:"GatewayResult"` 
	GatewayTransactionDate *SplititTime `json:"GatewayTransactionDate"` 
	IsChargeback bool `json:"IsChargeback"` 
	AVSResult *CardResult `json:"AVSResult,omitempty"` 
	CVCResult *CardResult `json:"CVCResult,omitempty"` 
	IsInDispute bool `json:"IsInDispute,omitempty"` 
	DisputeStatus DisputeStatus `json:"DisputeStatus,omitempty"` 
}
