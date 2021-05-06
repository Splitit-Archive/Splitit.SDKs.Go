# Installments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**InstallmentPlanId** | **int64** |  | 
**InstallmentNumber** | **int64** |  | 
**ProcessorId** | **int64** |  | 
**OriginUtcOffset** | **float64** |  | 
**CreatedDateUtc** | [**time.Time**](time.Time.md) |  | 
**Amount** | **float32** |  | 
**ProcessDateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**ProcessDateTimeUtc** | [**time.Time**](time.Time.md) |  | [optional] 
**IsRefund** | **bool** |  | 
**CardStateId** | **int64** |  | [optional] 
**OriginalAmount** | **float32** |  | 
**RefundAmount** | **float32** |  | 
**IsFullCapture** | **bool** |  | 
**AuthorizedAmount** | **float32** |  | 
**IsPreAuthorized** | **bool** |  | 
**AmountForFunding** | **float32** |  | 
**CardState** | [**CardStateLogs**](CardStateLogs.md) |  | [optional] 
**InstallmentPlan** | [**InstallmentPlans**](InstallmentPlans.md) |  | [optional] 
**Processor** | [**Processors**](Processors.md) |  | [optional] 
**FeesDocuments** | [**[]FeesDocuments**](FeesDocuments.md) |  | [optional] 
**FundingCollectDocumentDetails** | [**[]FundingCollectDocumentDetails**](FundingCollectDocumentDetails.md) |  | [optional] 
**InstallmentPaymentGatewayTransactionLogs** | [**[]InstallmentPaymentGatewayTransactionLogs**](InstallmentPaymentGatewayTransactionLogs.md) |  | [optional] 
**TransferDocumentDetails** | [**[]TransferDocumentDetails**](TransferDocumentDetails.md) |  | [optional] 
**MerchantAccountType** | [**MerchantAccountType**](MerchantAccountType.md) |  | 
**Status** | [**InstallmentStatus**](InstallmentStatus.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


