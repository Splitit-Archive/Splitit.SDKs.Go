# PaymentGatewayTransactionLogsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvsMessageMessageCode** | **string** |  | [optional] 
**AvsMessageMessageText** | **string** |  | [optional] 
**CvvMessageMessageCode** | **string** |  | [optional] 
**CvvMessageMessageText** | **string** |  | [optional] 
**ExpiryDateMessageMessageCode** | **string** |  | [optional] 
**ExpiryDateMessageMessageText** | **string** |  | [optional] 
**ResultMessageMessageCode** | **string** |  | [optional] 
**ResultMessageMessageText** | **string** |  | [optional] 
**Result** | **bool** |  | 
**TransactionId** | **string** |  | [optional] 
**CompleteResponseXml** | **string** |  | [optional] 
**AdditionalData** | **string** |  | [optional] 
**RequestedCurrencyCode** | **string** |  | [optional] 
**TerminalGatewayDataId** | **int64** |  | 
**ReferencePaymentGatewayTransactionLogId** | **int64** |  | [optional] 
**ProcessedAmountAmount** | **float32** |  | 
**RequestedAmountAmount** | **float32** |  | 
**InstallmentPlanId** | **int64** |  | [optional] 
**IsChargeback** | **bool** |  | 
**IsAsync** | **bool** |  | [optional] 
**TransferId** | **string** |  | [optional] 
**ReferencePaymentGatewayTransactionLog** | [**PaymentGatewayTransactionLogs**](PaymentGatewayTransactionLogs.md) |  | [optional] 
**TerminalGatewayData** | [**TerminalGatewayDatas**](TerminalGatewayDatas.md) |  | [optional] 
**FraudDetectionLogs** | [**[]FraudDetectionLogs**](FraudDetectionLogs.md) |  | [optional] 
**InstallmentPaymentGatewayTransactionLogs** | [**[]InstallmentPaymentGatewayTransactionLogs**](InstallmentPaymentGatewayTransactionLogs.md) |  | [optional] 
**InverseReferencePaymentGatewayTransactionLog** | [**[]PaymentGatewayTransactionLogs**](PaymentGatewayTransactionLogs.md) |  | [optional] 
**PaymentGatewayTransactionAsyncLogs** | [**[]PaymentGatewayTransactionAsyncLogs**](PaymentGatewayTransactionAsyncLogs.md) |  | [optional] 
**ReAuthorizationPaymentGatewayTransactionLogs** | [**[]ReAuthorizationPaymentGatewayTransactionLogs**](ReAuthorizationPaymentGatewayTransactionLogs.md) |  | [optional] 
**PublicId** | **string** |  | [optional] 
**TraceId** | **string** |  | [optional] 
**Type** | [**OperationType**](OperationType.md) |  | 
**AvsMessage** | [**PaymentGatewayMessage**](PaymentGatewayMessage.md) |  | [optional] 
**CvvMessage** | [**PaymentGatewayMessage**](PaymentGatewayMessage.md) |  | [optional] 
**ExpiryDateMessage** | [**PaymentGatewayMessage**](PaymentGatewayMessage.md) |  | [optional] 
**ResultMessage** | [**PaymentGatewayMessage**](PaymentGatewayMessage.md) |  | [optional] 
**ProcessedAmount** | [**Money2**](Money2.md) |  | [optional] 
**RequestedAmount** | [**Money2**](Money2.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


