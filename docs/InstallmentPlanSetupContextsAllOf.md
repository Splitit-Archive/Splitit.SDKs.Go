# InstallmentPlanSetupContextsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicInstallmentPlanSession** | **string** |  | [optional] 
**BillingAddressAddressLine** | **string** |  | [optional] 
**BillingAddressZip** | **string** |  | [optional] 
**BillingAddressAddressLine2** | **string** |  | [optional] 
**BillingAddressCity** | **string** |  | [optional] 
**RequestedNumberOfInstallments** | **string** |  | [optional] 
**SuccessExitUrl** | **string** |  | [optional] 
**ErrorExitUrl** | **string** |  | [optional] 
**CancelExitUrl** | **string** |  | [optional] 
**AbTestQueryString** | **string** |  | [optional] 
**RequestedFirstInstallmentAmount** | **float32** |  | 
**RequestedFirstScheduledInstallmentDate** | [**time.Time**](time.Time.md) |  | [optional] 
**SendShopperApprovalRequest** | [**time.Time**](time.Time.md) |  | [optional] 
**SendShopperPaymentRequest** | [**time.Time**](time.Time.md) |  | [optional] 
**RequestedCaptureOnCreation** | **bool** |  | [optional] 
**BillingAddressCountryId** | **int64** |  | [optional] 
**SuggestedPaymentTPId** | **int64** |  | [optional] 
**ABTestingSessionId** | **int64** |  | [optional] 
**DefaultNumOfInstallments** | **int32** |  | 
**BillingAddressStateId** | **int64** |  | [optional] 
**BillingAddressLegacyFullAddressLine** | **string** |  | [optional] 
**RequestedFunding** | **bool** |  | [optional] 
**SuccessAsyncUrl** | **string** |  | [optional] 
**Attempt3Dsecure** | **bool** |  | [optional] 
**ViewName** | **string** |  | [optional] 
**IsOpenedInIframe** | **bool** |  | 
**Is3dSecureInPopup** | **bool** |  | [optional] 
**ExternalAuthUniqueGatewayId** | **string** |  | [optional] 
**ExternalAmountDetails** | **float32** |  | 
**ExternalAuthDate** | [**time.Time**](time.Time.md) |  | 
**ExternalAuthTransactionLog** | **string** |  | [optional] 
**PaymentApprovalPhone** | **string** |  | [optional] 
**PaymentApprovalEmail** | **string** |  | [optional] 
**PaymentMessage** | **string** |  | [optional] 
**TermsAndConditionsKey** | **string** |  | [optional] 
**FingerPrintId** | **string** |  | [optional] 
**RequestUpdateCardPhone** | **string** |  | [optional] 
**RequestUpdateCardEmail** | **string** |  | [optional] 
**RequestUpdateCardTime** | [**time.Time**](time.Time.md) |  | [optional] 
**BillingAddressFirstTimePopulatedBy** | **string** |  | [optional] 
**Country** | [**Countries**](Countries.md) |  | [optional] 
**State** | [**CountrySubdivisions**](CountrySubdivisions.md) |  | [optional] 
**ABTestingSession** | [**PaymentFormTpabTestingDefinition**](PaymentFormTPABTestingDefinition.md) |  | [optional] 
**SuggestedPaymentTP** | [**VersionedTouchPoints**](VersionedTouchPoints.md) |  | [optional] 
**InstallmentPlans** | [**[]InstallmentPlans**](InstallmentPlans.md) |  | [optional] 
**RequestedStrategy** | [**PlanStrategy**](PlanStrategy.md) |  | [optional] 
**BillingAddress** | [**AddressData2**](AddressData2.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


