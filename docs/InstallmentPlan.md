# InstallmentPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallmentPlanNumber** | **string** |  | [optional] 
**InstallmentPlanStatus** | [**ReferenceEntityBase**](ReferenceEntityBase.md) |  | [optional] 
**Amount** | [**Money**](Money.md) |  | [optional] 
**OutstandingAmount** | [**Money**](Money.md) |  | [optional] 
**NumberOfInstallments** | **int32** |  | 
**NumberOfProcessedInstallments** | **int32** |  | 
**OriginalAmount** | [**Money**](Money.md) |  | [optional] 
**RefundAmount** | [**Money**](Money.md) |  | [optional] 
**Consumer** | [**ConsumerData**](ConsumerData.md) |  | [optional] 
**ActiveCard** | [**CardData**](CardData.md) |  | [optional] 
**FraudCheck** | [**FraudCheck**](FraudCheck.md) |  | [optional] 
**Merchant** | [**MerchantRef**](MerchantRef.md) |  | [optional] 
**RefOrderNumber** | **string** |  | [optional] 
**PurchaseMethod** | [**ReferenceEntityBase**](ReferenceEntityBase.md) |  | [optional] 
**Strategy** | [**ReferenceEntityBase**](ReferenceEntityBase.md) |  | [optional] 
**DelayResolution** | [**ReferenceEntityBase**](ReferenceEntityBase.md) |  | [optional] 
**ExtendedParams** | **map[string]string** |  | [optional] 
**IsFullCaptured** | **bool** |  | 
**IsChargedBack** | **bool** |  | 
**ArePaymentsOnHold** | **bool** |  | 
**ScpFundingPercent** | **float32** |  | 
**IsFunded** | **bool** |  | 
**TestMode** | [**TestModes**](TestModes.md) |  | 
**CreationDateTime** | [**time.Time**](time.Time.md) |  | 
**Installments** | [**[]Installment2**](Installment2.md) |  | [optional] 
**SecureAuthorizations** | [**[]ReAuthorization**](ReAuthorization.md) |  | [optional] 
**LogoUrl** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


