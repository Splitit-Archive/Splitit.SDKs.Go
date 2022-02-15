# TransactionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayTransactionId** | **string** |  | [optional] 
**SplititTransactionId** | **int64** |  | 
**GatewayResultCode** | **string** |  | [optional] 
**GatewayResultMessage** | **string** |  | [optional] 
**OperationType** | [**ReferenceEntityBase**](ReferenceEntityBase.md) |  | [optional] 
**GatewayResult** | **bool** |  | 
**GatewayTransactionDate** | [**time.Time**](time.Time.md) |  | 
**IsChargeback** | **bool** |  | 
**AVSResult** | [**CardResult**](CardResult.md) |  | [optional] 
**CVCResult** | [**CardResult**](CardResult.md) |  | [optional] 
**IsInDispute** | **bool** |  | [optional] 
**DisputeStatus** | [**DisputeStatus**](DisputeStatus.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


