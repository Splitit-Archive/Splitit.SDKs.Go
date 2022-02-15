/*
 * splitit-web-api-public-sdk
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package splitit
// PaymentWizardDataResponseAllOf struct for PaymentWizardDataResponseAllOf
type PaymentWizardDataResponseAllOf struct {
	ShowAddressElements string `json:"ShowAddressElements,omitempty"` 
	CurrencyDisplay *ExtendedCurrency `json:"CurrencyDisplay,omitempty"` 
	ForceDisplayImportantNotes bool `json:"ForceDisplayImportantNotes"` 
	ShowShopperDetailsExpendedOnStart bool `json:"ShowShopperDetailsExpendedOnStart"` 
	ShowPaymentScheduleRequiredCredit bool `json:"ShowPaymentScheduleRequiredCredit"` 
	IsShopperEmailMandatory bool `json:"IsShopperEmailMandatory"` 
	IsShopperPhoneMandatory bool `json:"IsShopperPhoneMandatory"` 
	NumberOfInstallmentsSelectionsOption string `json:"NumberOfInstallmentsSelectionsOption,omitempty"` 
	Is3ds2Supported bool `json:"Is3ds2Supported"` 
	ProcessorName string `json:"ProcessorName,omitempty"` 
	AddressIsReadonly bool `json:"AddressIsReadonly"` 
	PhoneIsReadOnly bool `json:"PhoneIsReadOnly"` 
	EmailIsReadOnly bool `json:"EmailIsReadOnly"` 
	ShowLearnMore bool `json:"ShowLearnMore"` 
	ShowMobilePhone bool `json:"ShowMobilePhone"` 
	ShowCloseDialogBeforeAbandon bool `json:"ShowCloseDialogBeforeAbandon"` 
	LogoURL string `json:"LogoURL,omitempty"` 
	DefaultNumOfInstallments int32 `json:"DefaultNumOfInstallments"` 
	PrivacyPolicyUrl string `json:"PrivacyPolicyUrl,omitempty"` 
	TermsAndConditionsUrl string `json:"TermsAndConditionsUrl,omitempty"` 
	LearnMoreUrl string `json:"LearnMoreUrl,omitempty"` 
	PotentialCardTypes []CardType `json:"PotentialCardTypes,omitempty"` 
	PotentialCardBrands []CardBrand `json:"PotentialCardBrands,omitempty"` 
	PaymentFormMessages []PaymentFormMessage `json:"PaymentFormMessages,omitempty"` 
	DisplayProperties map[string]string `json:"DisplayProperties,omitempty"` 
	TermsAndConditions *TermsAndConditions `json:"TermsAndConditions,omitempty"` 
	PaymentMethods []PaymentMethods `json:"PaymentMethods,omitempty"` 
	Status InstallmentPlanStatus `json:"Status"` 
	IsAttempt3Dsecure bool `json:"IsAttempt3Dsecure"` 
}
