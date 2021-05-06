# InstallmentPlansAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TerminalId** | **int64** |  | 
**NumberOfInstallments** | **int64** |  | 
**Amount** | **float32** |  | 
**AmountBeforeFees** | **float32** |  | 
**Eula** | **bool** |  | 
**InstallmentPlanNumber** | **string** |  | [optional] 
**IsFullCaptured** | **bool** |  | 
**BeginLockTime** | [**time.Time**](time.Time.md) |  | [optional] 
**MerchantFinancedDate** | [**time.Time**](time.Time.md) |  | [optional] 
**MerchantReserveReturnedDate** | [**time.Time**](time.Time.md) |  | [optional] 
**PlannedFullCapture** | [**time.Time**](time.Time.md) |  | [optional] 
**InstallmentPlanTerminalDataId** | **int64** |  | 
**CurrencyId** | **int64** |  | 
**RefOrderNumber** | **string** |  | [optional] 
**ActiveCardId** | **int64** |  | [optional] 
**Base64PngSignature** | **string** |  | [optional] 
**IsChargedBack** | **bool** |  | 
**AreChargesHeld** | **bool** |  | 
**AutoRetry** | **bool** |  | 
**PisMemberUniqueId** | **string** |  | [optional] 
**PisMemberId** | **int64** |  | [optional] 
**Subtotal** | **float32** |  | 
**Tax** | **float32** |  | 
**Shipping** | **float32** |  | 
**SetupContextId** | **int64** |  | 
**OriginalAmount** | **float32** |  | 
**RefundAmount** | **float32** |  | 
**ReservePoolAmount** | **float32** |  | 
**ShopperApprovalDateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**CancellationDateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**ActivatedDate** | [**time.Time**](time.Time.md) |  | [optional] 
**DeActivatedDate** | [**time.Time**](time.Time.md) |  | [optional] 
**AmountForFunding** | **float32** |  | 
**AmountForFundingCurrency** | [**Currencies**](Currencies.md) |  | [optional] 
**AmountForFundingExchangeRate** | **float32** |  | 
**AmountForFundingTransactionCurrency** | **float32** |  | 
**ActiveCard** | [**Cards**](Cards.md) |  | [optional] 
**Currency** | [**Currencies**](Currencies.md) |  | [optional] 
**InstallmentPlanTerminalData** | [**InstallmentPlanTerminalDatas**](InstallmentPlanTerminalDatas.md) |  | [optional] 
**PisMember** | [**PisMembers**](PisMembers.md) |  | [optional] 
**SetupContext** | [**InstallmentPlanSetupContexts**](InstallmentPlanSetupContexts.md) |  | [optional] 
**Terminal** | [**Terminals**](Terminals.md) |  | [optional] 
**CartItems** | [**[]CartItems**](CartItems.md) |  | [optional] 
**EmailAuditLogs** | [**[]EmailAuditLogs**](EmailAuditLogs.md) |  | [optional] 
**FeesDocuments** | [**[]FeesDocuments**](FeesDocuments.md) |  | [optional] 
**FraudDetectionLogs** | [**[]FraudDetectionLogs**](FraudDetectionLogs.md) |  | [optional] 
**FundingCollectDocumentDetails** | [**[]FundingCollectDocumentDetails**](FundingCollectDocumentDetails.md) |  | [optional] 
**InstallmentPlanAuditLogs** | [**[]InstallmentPlanAuditLogs**](InstallmentPlanAuditLogs.md) |  | [optional] 
**InstallmentPlanEvents** | [**[]InstallmentPlanEvents**](InstallmentPlanEvents.md) |  | [optional] 
**InstallmentPlanStatusLogEntries** | [**[]InstallmentPlanStatusLogEntries**](InstallmentPlanStatusLogEntries.md) |  | [optional] 
**Installments** | [**[]Installments**](Installments.md) |  | [optional] 
**ReAuthorizations** | [**[]ReAuthorizations**](ReAuthorizations.md) |  | [optional] 
**RefundLogEntries** | [**[]RefundLogEntries**](RefundLogEntries.md) |  | [optional] 
**TransferDocumentDetails** | [**[]TransferDocumentDetails**](TransferDocumentDetails.md) |  | [optional] 
**ActiveTerminalDataId** | **int64** |  | [optional] 
**ActiveTerminalData** | [**TerminalGatewayDatas**](TerminalGatewayDatas.md) |  | [optional] 
**DelayResolution** | [**DelayResolution**](DelayResolution.md) |  | [optional] 
**TestMode** | [**TestModes**](TestModes.md) |  | 
**FundingTypesId** | [**MoneyFlows**](MoneyFlows.md) |  | 
**Strategy** | [**PlanStrategy**](PlanStrategy.md) |  | 
**Status** | [**InstallmentPlanStatus**](InstallmentPlanStatus.md) |  | 
**PurchaseMethod** | [**PurchaseMethod**](PurchaseMethod.md) |  | 
**ExtendedParamsSerializedData** | **string** |  | [optional] 
**ExtendedParams** | **map[string]string** |  | [optional] 
**AutoCapture** | **bool** |  | 
**Items** | [**[]CartItems**](CartItems.md) |  | [optional] 
**AmountDetails** | [**AmountDetails2**](AmountDetails2.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


