# PaymentWizardDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedNumberOfInstallments** | **string** |  | [optional] 
**SuccessExitURL** | **string** |  | [optional] 
**ErrorExitURL** | **string** |  | [optional] 
**CancelExitURL** | **string** |  | [optional] 
**SuccessAsyncUrl** | **string** |  | [optional] 
**ViewName** | **string** |  | [optional] 
**IsOpenedInIframe** | **bool** |  | 
**Is3dSecureInPopup** | **bool** |  | [optional] 
**PaymentFormMessage** | **string** |  | [optional] 
**SetShortUrl** | **bool** |  | 
**ShowAddressElements** | **string** |  | [optional] 
**CurrencyDisplay** | [**ExtendedCurrency**](ExtendedCurrency.md) |  | [optional] 
**ForceDisplayImportantNotes** | **bool** |  | 
**ShowShopperDetailsExpendedOnStart** | **bool** |  | 
**ShowPaymentScheduleRequiredCredit** | **bool** |  | 
**IsShopperEmailMandatory** | **bool** |  | 
**IsShopperPhoneMandatory** | **bool** |  | 
**NumberOfInstallmentsSelectionsOption** | **string** |  | [optional] 
**Is3ds2Supported** | **bool** |  | 
**ProcessorName** | **string** |  | [optional] 
**AddressIsReadonly** | **bool** |  | 
**PhoneIsReadOnly** | **bool** |  | 
**EmailIsReadOnly** | **bool** |  | 
**ShowLearnMore** | **bool** |  | 
**ShowMobilePhone** | **bool** |  | 
**ShowCloseDialogBeforeAbandon** | **bool** |  | 
**LogoURL** | **string** |  | [optional] 
**DefaultNumOfInstallments** | **int32** |  | 
**PrivacyPolicyUrl** | **string** |  | [optional] 
**TermsAndConditionsUrl** | **string** |  | [optional] 
**LearnMoreUrl** | **string** |  | [optional] 
**PotentialCardTypes** | [**[]CardType**](CardType.md) |  | [optional] 
**PotentialCardBrands** | [**[]CardBrand**](CardBrand.md) |  | [optional] 
**PaymentFormMessages** | [**[]PaymentFormMessage**](PaymentFormMessage.md) |  | [optional] 
**DisplayProperties** | **map[string]string** |  | [optional] 
**TermsAndConditions** | [**TermsAndConditions**](TermsAndConditions.md) |  | [optional] 
**PaymentMethods** | [**[]PaymentMethods**](PaymentMethods.md) |  | [optional] 
**Status** | [**InstallmentPlanStatus**](InstallmentPlanStatus.md) |  | 
**IsAttempt3Dsecure** | **bool** |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


