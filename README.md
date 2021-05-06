# Splitit SDK for Go language

This is Splitit Web API SDK source code for Go applications. For other languages, please visit [Splitit.SDKs](https://github.com/Splitit/Splitit.SDKs).

## Overview

- API version: 1.0.0
- Package version: 1.6.7

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get github.com/fatih/structs
go get golang.org/x/net/context
go get github.com/antihax/optional
go get github.com/btubbs/datetime
```

Install the Splitit SDK package:

```shell
go get github.com/splitit/splitit.sdks.go
```

## Getting Started

Replace _YOUR_API_KEY_, _YOUR_USERNAME_ and _YOUR_PASSWORD_ placeholders with your corresponding credentials.
Please install required packages as described above and then run the following:

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/splitit/splitit.sdks.go"
)

func main() {
	// Recommended to have one shared client
	// Use splitit.NewSandboxAPIClient for development and splitit.NewAPIClient in production
	client := splitit.NewSandboxAPIClient(
		"_YOUR_API_KEY_",
		"_YOUR_USERNAME_",
		"_YOUR_PASSWORD_",
		// Supply optional configuration
		// splitit.Debug(), // print out full request and response
		splitit.DefaultCulture("en-US"),
	)

	ctx := context.TODO()

	// Create initiate request
	initReq := splitit.InitiateInstallmentPlanRequest{
		PlanData: &splitit.PlanData{
			Amount: &splitit.MoneyWithCurrencyCode{300, "USD"},
			NumberOfInstallments: 3,
		},
		BillingAddress: &splitit.AddressData{
			AddressLine:  "260 Madison Avenue.",
			AddressLine2: "Apartment 1",
			City:         "New York",
			State:        "NY",
			Country:      "USA",
			Zip:          "10016",
		},
		ConsumerData: &splitit.ConsumerData{
			FullName:    "John Smith",
			Email:       "JohnGo@splitit.com",
			PhoneNumber: "1-415-775-4848",
			CultureName: "en-us",
		},
	}

	initResponse, _, err := client.InstallmentPlanApi.InstallmentPlanInitiate(
		// DefaultCulture could be overridden on per-request basis:
		splitit.WithCulture(ctx, "en-GB"),
		initReq,
	)

	if err != nil {
		fmt.Printf("Error during initiate: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Initiate success: %t\n", initResponse.ResponseHeader.Succeeded)
	}

	createReq := splitit.CreateInstallmentPlanRequest{
		CreditCardDetails: &splitit.CardData{
			CardNumber:         "411111111111111",
			CardCvv:            "111",
			CardHolderFullName: "John Smith",
			CardExpMonth:       "12",
			CardExpYear:        "2022",
		},
		InstallmentPlanNumber: initResponse.InstallmentPlan.InstallmentPlanNumber,
	}

	createResponse, _, err := client.InstallmentPlanApi.InstallmentPlanCreate(ctx, createReq)

	if err != nil {
		fmt.Printf("Create error: %s\n", err)
	} else {
		fmt.Printf("Create success: %t\t", createResponse.ResponseHeader.Succeeded)
	}
}

```

## Flex Fields

Common usage for Splitit PHP SDK is in making necessary server-side requests as part of FlexFields product integration.
The code below is an example of how SDK wrappers can be used to simplify acquiring public token and verifying payment.
For more information, please visit [FlexFields documentation](https://flex-fields.production.splitit.com/#php).

Server-side code consists of two parts: acquiring public token which needs to be passed to FlexFields JS library
and verifying payment before order is finalized and shipped.

### Getting public token
```go
func get_flexfields_public_token() string {
	ctx := context.TODO()

	client := splitit.NewSandboxAPIClient(
		"_YOUR_API_KEY_",
		"_YOUR_USERNAME_",
		"_YOUR_PASSWORD_",
		splitit.DefaultCulture("en-US"),
	)

	ff := splitit.NewFlexFields(client)
	token, _ := ff.GetPublicToken(ctx, 350, "USD")
	return token
}
```

### Payment verification
```go
func verify_payment(planNumber string, originalAmount float64) {
	ctx := context.TODO()

	client := splitit.NewSandboxAPIClient(
		"_YOUR_API_KEY_",
		"_YOUR_USERNAME_",
		"_YOUR_PASSWORD_",
	)

	ff := splitit.NewFlexFields(client)
	isVerified, _ := ff.VerifyPayment(ctx, planNumber, originalAmount)

	if !isVerified {
		// Respond to potential fraud attempt.
	}
}
```

For detailed information on request and response procedures, please visit [Splitit Web API documentation](https://documenter.getpostman.com/view/795699/RWaNQSJH?version=latest)

## Documentation for API Endpoints

All URIs are relative to *https://webapi.production.splitit.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InfoApi* | [**InfoGetLearnMoreDetails**](docs/InfoApi.md#infogetlearnmoredetails) | **Post** /api/Merchant/GetLearnMoreDetails | 
*InfrastructureApi* | [**InfrastructureGetResources**](docs/InfrastructureApi.md#infrastructuregetresources) | **Post** /api/Infrastructure/GetResources | 
*InfrastructureApi* | [**InfrastructureGetResources2**](docs/InfrastructureApi.md#infrastructuregetresources2) | **Get** /api/Infrastructure/GetResources | 
*InstallmentPlanApi* | [**InstallmentPlanApprove**](docs/InstallmentPlanApi.md#installmentplanapprove) | **Post** /api/InstallmentPlan/Approve | 
*InstallmentPlanApi* | [**InstallmentPlanCancel**](docs/InstallmentPlanApi.md#installmentplancancel) | **Post** /api/InstallmentPlan/Cancel | 
*InstallmentPlanApi* | [**InstallmentPlanChargeBack**](docs/InstallmentPlanApi.md#installmentplanchargeback) | **Post** /api/InstallmentPlan/ChargeBack | 
*InstallmentPlanApi* | [**InstallmentPlanCreate**](docs/InstallmentPlanApi.md#installmentplancreate) | **Post** /api/InstallmentPlan/Create | 
*InstallmentPlanApi* | [**InstallmentPlanGet**](docs/InstallmentPlanApi.md#installmentplanget) | **Post** /api/InstallmentPlan/Get | 
*InstallmentPlanApi* | [**InstallmentPlanGet3DSecureParameters**](docs/InstallmentPlanApi.md#installmentplanget3dsecureparameters) | **Post** /api/InstallmentPlan/Get3DSecureParameters | 
*InstallmentPlanApi* | [**InstallmentPlanGetExtended**](docs/InstallmentPlanApi.md#installmentplangetextended) | **Post** /api/InstallmentPlan/GetExtended | 
*InstallmentPlanApi* | [**InstallmentPlanGetFraudStatusDisplay**](docs/InstallmentPlanApi.md#installmentplangetfraudstatusdisplay) | **Post** /api/InstallmentPlan/GetFraudStatusDisplay | 
*InstallmentPlanApi* | [**InstallmentPlanGetInitiatedInstallmentPlanRequest**](docs/InstallmentPlanApi.md#installmentplangetinitiatedinstallmentplanrequest) | **Post** /api/InstallmentPlan/GetInitiatedInstallmentPlanRequest | 
*InstallmentPlanApi* | [**InstallmentPlanGetInitiatedUpdatePaymentData**](docs/InstallmentPlanApi.md#installmentplangetinitiatedupdatepaymentdata) | **Get** /api/InstallmentPlan/GetInitiatedUpdatePaymentData | 
*InstallmentPlanApi* | [**InstallmentPlanGetLearnMoreDetails**](docs/InstallmentPlanApi.md#installmentplangetlearnmoredetails) | **Post** /api/InstallmentPlan/GetLearnMoreDetails | 
*InstallmentPlanApi* | [**InstallmentPlanGetPGTL**](docs/InstallmentPlanApi.md#installmentplangetpgtl) | **Post** /api/InstallmentPlan/GetPGTL | 
*InstallmentPlanApi* | [**InstallmentPlanGetSchedules**](docs/InstallmentPlanApi.md#installmentplangetschedules) | **Post** /api/InstallmentPlan/GetSchedules | 
*InstallmentPlanApi* | [**InstallmentPlanInitiate**](docs/InstallmentPlanApi.md#installmentplaninitiate) | **Post** /api/InstallmentPlan/Initiate | 
*InstallmentPlanApi* | [**InstallmentPlanRefund**](docs/InstallmentPlanApi.md#installmentplanrefund) | **Post** /api/InstallmentPlan/Refund | 
*InstallmentPlanApi* | [**InstallmentPlanStartInstallments**](docs/InstallmentPlanApi.md#installmentplanstartinstallments) | **Post** /api/InstallmentPlan/StartInstallments | 
*InstallmentPlanApi* | [**InstallmentPlanTermsAndConditions**](docs/InstallmentPlanApi.md#installmentplantermsandconditions) | **Post** /api/InstallmentPlan/TermsAndConditions | 
*InstallmentPlanApi* | [**InstallmentPlanTestCard**](docs/InstallmentPlanApi.md#installmentplantestcard) | **Post** /api/InstallmentPlan/TestCard | 
*InstallmentPlanApi* | [**InstallmentPlanUpdate**](docs/InstallmentPlanApi.md#installmentplanupdate) | **Post** /api/InstallmentPlan/Update | 
*InstallmentPlanApi* | [**InstallmentPlanVerifyPayment**](docs/InstallmentPlanApi.md#installmentplanverifypayment) | **Post** /api/InstallmentPlan/Get/VerifyPayment | 
*LoginApi* | [**LoginPost**](docs/LoginApi.md#loginpost) | **Post** /api/Login | 


## Documentation For Models

 - [AccountUpdaterPendingCards](docs/AccountUpdaterPendingCards.md)
 - [AccountUpdaterPendingCardsAllOf](docs/AccountUpdaterPendingCardsAllOf.md)
 - [AccountUpdaterProvider](docs/AccountUpdaterProvider.md)
 - [AccountUpdaterSubscribersScopes](docs/AccountUpdaterSubscribersScopes.md)
 - [AccountUpdaterSubscribersScopesAllOf](docs/AccountUpdaterSubscribersScopesAllOf.md)
 - [AccountingParty](docs/AccountingParty.md)
 - [AddressData](docs/AddressData.md)
 - [AddressData2](docs/AddressData2.md)
 - [AgentFeeType](docs/AgentFeeType.md)
 - [Agents](docs/Agents.md)
 - [AgentsAllOf](docs/AgentsAllOf.md)
 - [AmountDetails](docs/AmountDetails.md)
 - [AmountDetails2](docs/AmountDetails2.md)
 - [ApiUserPasswordHistories](docs/ApiUserPasswordHistories.md)
 - [ApiUserPasswordHistoriesAllOf](docs/ApiUserPasswordHistoriesAllOf.md)
 - [ApiUsers](docs/ApiUsers.md)
 - [ApiUsersAllOf](docs/ApiUsersAllOf.md)
 - [ApplicativeUser](docs/ApplicativeUser.md)
 - [ApplicativeUserAllOf](docs/ApplicativeUserAllOf.md)
 - [ApplicativeUserPasswordHistory](docs/ApplicativeUserPasswordHistory.md)
 - [ApplicativeUserPasswordHistoryAllOf](docs/ApplicativeUserPasswordHistoryAllOf.md)
 - [ApproveInstallmentPlanRequest](docs/ApproveInstallmentPlanRequest.md)
 - [AuditLogs](docs/AuditLogs.md)
 - [AuditLogsAllOf](docs/AuditLogsAllOf.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [BankDetails](docs/BankDetails.md)
 - [BinData](docs/BinData.md)
 - [BinDataItems](docs/BinDataItems.md)
 - [BinDataItemsAllOf](docs/BinDataItemsAllOf.md)
 - [BusinessActivity](docs/BusinessActivity.md)
 - [BusinessParty](docs/BusinessParty.md)
 - [BusinessUnits](docs/BusinessUnits.md)
 - [BusinessUnitsAllOf](docs/BusinessUnitsAllOf.md)
 - [CancelInstallmentPlanRequest](docs/CancelInstallmentPlanRequest.md)
 - [CardBrand](docs/CardBrand.md)
 - [CardData](docs/CardData.md)
 - [CardResult](docs/CardResult.md)
 - [CardStateLogs](docs/CardStateLogs.md)
 - [CardStateLogsAllOf](docs/CardStateLogsAllOf.md)
 - [CardType](docs/CardType.md)
 - [Cards](docs/Cards.md)
 - [CardsAllOf](docs/CardsAllOf.md)
 - [CartData](docs/CartData.md)
 - [CartItems](docs/CartItems.md)
 - [CartItemsAllOf](docs/CartItemsAllOf.md)
 - [ChargebackRequest](docs/ChargebackRequest.md)
 - [ChbDefaultAction](docs/ChbDefaultAction.md)
 - [CollectDocuments](docs/CollectDocuments.md)
 - [CollectDocumentsAllOf](docs/CollectDocumentsAllOf.md)
 - [ConfigKeys](docs/ConfigKeys.md)
 - [ConfigKeysAllOf](docs/ConfigKeysAllOf.md)
 - [ConfigValues](docs/ConfigValues.md)
 - [ConfigValuesAllOf](docs/ConfigValuesAllOf.md)
 - [ConsumerData](docs/ConsumerData.md)
 - [Countries](docs/Countries.md)
 - [CountriesAllOf](docs/CountriesAllOf.md)
 - [CountryState](docs/CountryState.md)
 - [CountrySubdivisions](docs/CountrySubdivisions.md)
 - [CountrySubdivisionsAllOf](docs/CountrySubdivisionsAllOf.md)
 - [CreateInstallmentPlanLegacyResponse](docs/CreateInstallmentPlanLegacyResponse.md)
 - [CreateInstallmentPlanRequest](docs/CreateInstallmentPlanRequest.md)
 - [CreateInstallmentsPlanResponse](docs/CreateInstallmentsPlanResponse.md)
 - [CreateInstallmentsPlanResponseAllOf](docs/CreateInstallmentsPlanResponseAllOf.md)
 - [Currencies](docs/Currencies.md)
 - [CurrenciesAllOf](docs/CurrenciesAllOf.md)
 - [Currency](docs/Currency.md)
 - [CurrencyAllOf](docs/CurrencyAllOf.md)
 - [CurrencyCountries](docs/CurrencyCountries.md)
 - [DekData](docs/DekData.md)
 - [DekDataAllOf](docs/DekDataAllOf.md)
 - [DelayResolution](docs/DelayResolution.md)
 - [EmailAuditLogs](docs/EmailAuditLogs.md)
 - [EmailAuditLogsAllOf](docs/EmailAuditLogsAllOf.md)
 - [EntityBase](docs/EntityBase.md)
 - [Error](docs/Error.md)
 - [EventsEndpoints](docs/EventsEndpoints.md)
 - [ExtendedCurrency](docs/ExtendedCurrency.md)
 - [ExtendedCurrencyAllOf](docs/ExtendedCurrencyAllOf.md)
 - [ExternalAuth](docs/ExternalAuth.md)
 - [FailureUnderFrozenInstallmentsPlan](docs/FailureUnderFrozenInstallmentsPlan.md)
 - [FeeRateCollections](docs/FeeRateCollections.md)
 - [FeeRateCollectionsAllOf](docs/FeeRateCollectionsAllOf.md)
 - [FeeRates](docs/FeeRates.md)
 - [FeeRatesAllOf](docs/FeeRatesAllOf.md)
 - [FeesDocuments](docs/FeesDocuments.md)
 - [FeesDocumentsAllOf](docs/FeesDocumentsAllOf.md)
 - [FeesRuleDatas](docs/FeesRuleDatas.md)
 - [FeesRuleDatasAllOf](docs/FeesRuleDatasAllOf.md)
 - [FeesTypes](docs/FeesTypes.md)
 - [FraudCheck](docs/FraudCheck.md)
 - [FraudCheckResult](docs/FraudCheckResult.md)
 - [FraudDetectionLogs](docs/FraudDetectionLogs.md)
 - [FraudDetectionLogsAllOf](docs/FraudDetectionLogsAllOf.md)
 - [FundingCollectDocumentDetails](docs/FundingCollectDocumentDetails.md)
 - [FundingCollectDocumentDetailsAllOf](docs/FundingCollectDocumentDetailsAllOf.md)
 - [FundingPayDocumentDetails](docs/FundingPayDocumentDetails.md)
 - [FundingPayDocumentDetailsAllOf](docs/FundingPayDocumentDetailsAllOf.md)
 - [FundingRuleDataScpProviderSetting](docs/FundingRuleDataScpProviderSetting.md)
 - [FundingRuleDataScpProviderSettingAllOf](docs/FundingRuleDataScpProviderSettingAllOf.md)
 - [FundingTransferReason](docs/FundingTransferReason.md)
 - [Get3DSecureParametersRequest](docs/Get3DSecureParametersRequest.md)
 - [Get3DSecureParametersResponse](docs/Get3DSecureParametersResponse.md)
 - [GetFraudStatusDisplayRequest](docs/GetFraudStatusDisplayRequest.md)
 - [GetFraudStatusDisplayResponse](docs/GetFraudStatusDisplayResponse.md)
 - [GetInitiatedInstallmentPlanRequest](docs/GetInitiatedInstallmentPlanRequest.md)
 - [GetInitiatedInstallmentPlanResponse](docs/GetInitiatedInstallmentPlanResponse.md)
 - [GetInitiatedUpdatePaymentDataResponse](docs/GetInitiatedUpdatePaymentDataResponse.md)
 - [GetInstallmentSchedulesRequest](docs/GetInstallmentSchedulesRequest.md)
 - [GetInstallmentsPlanExtendedResponse](docs/GetInstallmentsPlanExtendedResponse.md)
 - [GetInstallmentsPlanResponse](docs/GetInstallmentsPlanResponse.md)
 - [GetInstallmentsPlanSearchCriteriaRequest](docs/GetInstallmentsPlanSearchCriteriaRequest.md)
 - [GetInstallmentsScheduleResponse](docs/GetInstallmentsScheduleResponse.md)
 - [GetPgtlRequest](docs/GetPgtlRequest.md)
 - [GetPgtlResponse](docs/GetPgtlResponse.md)
 - [GetResourcesRequest](docs/GetResourcesRequest.md)
 - [GetResourcesRequestContext](docs/GetResourcesRequestContext.md)
 - [GetResourcesResponse](docs/GetResourcesResponse.md)
 - [IUserPasswordHistory](docs/IUserPasswordHistory.md)
 - [InitiateInstallmentPlanRequest](docs/InitiateInstallmentPlanRequest.md)
 - [InitiateInstallmentsPlanResponse](docs/InitiateInstallmentsPlanResponse.md)
 - [InitiateInstallmentsPlanResponseAllOf](docs/InitiateInstallmentsPlanResponseAllOf.md)
 - [Installment](docs/Installment.md)
 - [Installment2](docs/Installment2.md)
 - [InstallmentPaymentGatewayTransactionLogs](docs/InstallmentPaymentGatewayTransactionLogs.md)
 - [InstallmentPlan](docs/InstallmentPlan.md)
 - [InstallmentPlanActivityStatus](docs/InstallmentPlanActivityStatus.md)
 - [InstallmentPlanAuditLogs](docs/InstallmentPlanAuditLogs.md)
 - [InstallmentPlanAuditLogsAllOf](docs/InstallmentPlanAuditLogsAllOf.md)
 - [InstallmentPlanCancelationReason](docs/InstallmentPlanCancelationReason.md)
 - [InstallmentPlanDateInfo](docs/InstallmentPlanDateInfo.md)
 - [InstallmentPlanEventSubscriberRecordSendLogs](docs/InstallmentPlanEventSubscriberRecordSendLogs.md)
 - [InstallmentPlanEventSubscriberRecordSendLogsAllOf](docs/InstallmentPlanEventSubscriberRecordSendLogsAllOf.md)
 - [InstallmentPlanEventSubscriberRecords](docs/InstallmentPlanEventSubscriberRecords.md)
 - [InstallmentPlanEventSubscriberRecordsAllOf](docs/InstallmentPlanEventSubscriberRecordsAllOf.md)
 - [InstallmentPlanEventType](docs/InstallmentPlanEventType.md)
 - [InstallmentPlanEvents](docs/InstallmentPlanEvents.md)
 - [InstallmentPlanEventsAllOf](docs/InstallmentPlanEventsAllOf.md)
 - [InstallmentPlanEventsSubscriptionDatas](docs/InstallmentPlanEventsSubscriptionDatas.md)
 - [InstallmentPlanEventsSubscriptionDatasAllOf](docs/InstallmentPlanEventsSubscriptionDatasAllOf.md)
 - [InstallmentPlanInitiatedStatuses](docs/InstallmentPlanInitiatedStatuses.md)
 - [InstallmentPlanQueryCriteria](docs/InstallmentPlanQueryCriteria.md)
 - [InstallmentPlanResponse](docs/InstallmentPlanResponse.md)
 - [InstallmentPlanSetupContexts](docs/InstallmentPlanSetupContexts.md)
 - [InstallmentPlanSetupContextsAllOf](docs/InstallmentPlanSetupContextsAllOf.md)
 - [InstallmentPlanStatus](docs/InstallmentPlanStatus.md)
 - [InstallmentPlanStatusLogEntries](docs/InstallmentPlanStatusLogEntries.md)
 - [InstallmentPlanStatusLogEntriesAllOf](docs/InstallmentPlanStatusLogEntriesAllOf.md)
 - [InstallmentPlanTerminalDatas](docs/InstallmentPlanTerminalDatas.md)
 - [InstallmentPlanTerminalDatasAllOf](docs/InstallmentPlanTerminalDatasAllOf.md)
 - [InstallmentPlans](docs/InstallmentPlans.md)
 - [InstallmentPlansAllOf](docs/InstallmentPlansAllOf.md)
 - [InstallmentStatus](docs/InstallmentStatus.md)
 - [Installments](docs/Installments.md)
 - [InstallmentsAllOf](docs/InstallmentsAllOf.md)
 - [InstallmentsPlanDateType](docs/InstallmentsPlanDateType.md)
 - [IntegrationType](docs/IntegrationType.md)
 - [ItemData](docs/ItemData.md)
 - [Keks](docs/Keks.md)
 - [KeksAllOf](docs/KeksAllOf.md)
 - [LearnMoreDetailsRequest](docs/LearnMoreDetailsRequest.md)
 - [LearnMoreDetailsResponse](docs/LearnMoreDetailsResponse.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [MerchantAccountType](docs/MerchantAccountType.md)
 - [MerchantRef](docs/MerchantRef.md)
 - [MerchantRefAllOf](docs/MerchantRefAllOf.md)
 - [MerchantStatus](docs/MerchantStatus.md)
 - [MerchantVertical](docs/MerchantVertical.md)
 - [Merchants](docs/Merchants.md)
 - [MerchantsAllOf](docs/MerchantsAllOf.md)
 - [Money](docs/Money.md)
 - [Money2](docs/Money2.md)
 - [MoneyFlows](docs/MoneyFlows.md)
 - [MoneyWithCurrencyCode](docs/MoneyWithCurrencyCode.md)
 - [OnBoardingMethod](docs/OnBoardingMethod.md)
 - [OperationType](docs/OperationType.md)
 - [PagingRequestHeader](docs/PagingRequestHeader.md)
 - [PagingResponseHeader](docs/PagingResponseHeader.md)
 - [ParameterGroups](docs/ParameterGroups.md)
 - [ParameterGroupsAllOf](docs/ParameterGroupsAllOf.md)
 - [Parameters](docs/Parameters.md)
 - [ParametersAllOf](docs/ParametersAllOf.md)
 - [PayDocuments](docs/PayDocuments.md)
 - [PayDocumentsAllOf](docs/PayDocumentsAllOf.md)
 - [PaymentFormMessage](docs/PaymentFormMessage.md)
 - [PaymentFormMessageType](docs/PaymentFormMessageType.md)
 - [PaymentFormTpabTestingDefinition](docs/PaymentFormTpabTestingDefinition.md)
 - [PaymentFormTpabTestingDefinitionAllOf](docs/PaymentFormTpabTestingDefinitionAllOf.md)
 - [PaymentGatewayMessage](docs/PaymentGatewayMessage.md)
 - [PaymentGatewayTransactionAsyncLogs](docs/PaymentGatewayTransactionAsyncLogs.md)
 - [PaymentGatewayTransactionAsyncLogsAllOf](docs/PaymentGatewayTransactionAsyncLogsAllOf.md)
 - [PaymentGatewayTransactionLogs](docs/PaymentGatewayTransactionLogs.md)
 - [PaymentGatewayTransactionLogsAllOf](docs/PaymentGatewayTransactionLogsAllOf.md)
 - [PaymentToken](docs/PaymentToken.md)
 - [PaymentWizardData](docs/PaymentWizardData.md)
 - [PaymentWizardDataResponse](docs/PaymentWizardDataResponse.md)
 - [PaymentWizardDataResponseAllOf](docs/PaymentWizardDataResponseAllOf.md)
 - [PisErrorCodes](docs/PisErrorCodes.md)
 - [PisMemberPasswordHistories](docs/PisMemberPasswordHistories.md)
 - [PisMemberPasswordHistoriesAllOf](docs/PisMemberPasswordHistoriesAllOf.md)
 - [PisMembers](docs/PisMembers.md)
 - [PisMembersAllOf](docs/PisMembersAllOf.md)
 - [PisSessions](docs/PisSessions.md)
 - [PisSessionsAllOf](docs/PisSessionsAllOf.md)
 - [PisUserBusinessUnits](docs/PisUserBusinessUnits.md)
 - [PisUserPasswordHistories](docs/PisUserPasswordHistories.md)
 - [PisUserPasswordHistoriesAllOf](docs/PisUserPasswordHistoriesAllOf.md)
 - [PisUsers](docs/PisUsers.md)
 - [PisUsersAllOf](docs/PisUsersAllOf.md)
 - [PlanApprovalEvidence](docs/PlanApprovalEvidence.md)
 - [PlanData](docs/PlanData.md)
 - [PlanStrategy](docs/PlanStrategy.md)
 - [ProcessorAuthenticationParameters](docs/ProcessorAuthenticationParameters.md)
 - [ProcessorAuthenticationParametersAllOf](docs/ProcessorAuthenticationParametersAllOf.md)
 - [ProcessorType](docs/ProcessorType.md)
 - [Processors](docs/Processors.md)
 - [ProcessorsAllOf](docs/ProcessorsAllOf.md)
 - [PurchaseMethod](docs/PurchaseMethod.md)
 - [RangeType](docs/RangeType.md)
 - [ReAuthorization](docs/ReAuthorization.md)
 - [ReAuthorizationPaymentGatewayTransactionLogs](docs/ReAuthorizationPaymentGatewayTransactionLogs.md)
 - [ReAuthorizations](docs/ReAuthorizations.md)
 - [ReAuthorizationsAllOf](docs/ReAuthorizationsAllOf.md)
 - [ReconciliationReportItems](docs/ReconciliationReportItems.md)
 - [ReconciliationReportItemsAllOf](docs/ReconciliationReportItemsAllOf.md)
 - [ReconciliationReports](docs/ReconciliationReports.md)
 - [ReconciliationReportsAllOf](docs/ReconciliationReportsAllOf.md)
 - [RedirectUrls](docs/RedirectUrls.md)
 - [ReferenceEntityBase](docs/ReferenceEntityBase.md)
 - [RefundInstallmentPlanResponse](docs/RefundInstallmentPlanResponse.md)
 - [RefundInstallmentPlanResponseAllOf](docs/RefundInstallmentPlanResponseAllOf.md)
 - [RefundLogEntries](docs/RefundLogEntries.md)
 - [RefundLogEntriesAllOf](docs/RefundLogEntriesAllOf.md)
 - [RefundPlanRequest](docs/RefundPlanRequest.md)
 - [RefundStrategy](docs/RefundStrategy.md)
 - [RefundUnderCancelation](docs/RefundUnderCancelation.md)
 - [RelationsLoad](docs/RelationsLoad.md)
 - [RequestHeader](docs/RequestHeader.md)
 - [ResponseHeader](docs/ResponseHeader.md)
 - [Roles](docs/Roles.md)
 - [RolesAllOf](docs/RolesAllOf.md)
 - [SalesAssociates](docs/SalesAssociates.md)
 - [SalesAssociatesAllOf](docs/SalesAssociatesAllOf.md)
 - [Schedule](docs/Schedule.md)
 - [ScheduleElements](docs/ScheduleElements.md)
 - [ScpProviders](docs/ScpProviders.md)
 - [ScpProvidersAllOf](docs/ScpProvidersAllOf.md)
 - [ScpProvidersCurrencies](docs/ScpProvidersCurrencies.md)
 - [SessionAvailibility](docs/SessionAvailibility.md)
 - [SplititEntity](docs/SplititEntity.md)
 - [SplititJobDefinitions](docs/SplititJobDefinitions.md)
 - [SplititJobDefinitionsAllOf](docs/SplititJobDefinitionsAllOf.md)
 - [StartInstallmentsRequest](docs/StartInstallmentsRequest.md)
 - [StateLimitRuleDatas](docs/StateLimitRuleDatas.md)
 - [StateLimitRuleDatasAllOf](docs/StateLimitRuleDatasAllOf.md)
 - [SystemEmailsTypes](docs/SystemEmailsTypes.md)
 - [SystemTextCategory](docs/SystemTextCategory.md)
 - [TerminalCountries](docs/TerminalCountries.md)
 - [TerminalGatewayDatas](docs/TerminalGatewayDatas.md)
 - [TerminalGatewayDatasAllOf](docs/TerminalGatewayDatasAllOf.md)
 - [TerminalKvps](docs/TerminalKvps.md)
 - [TerminalKvpsAllOf](docs/TerminalKvpsAllOf.md)
 - [Terminals](docs/Terminals.md)
 - [TerminalsAllOf](docs/TerminalsAllOf.md)
 - [TermsAndConditions](docs/TermsAndConditions.md)
 - [TermsAndConditionsGetRequest](docs/TermsAndConditionsGetRequest.md)
 - [TermsAndConditionsGetResponse](docs/TermsAndConditionsGetResponse.md)
 - [TestCardRequest](docs/TestCardRequest.md)
 - [TestModes](docs/TestModes.md)
 - [Tokens](docs/Tokens.md)
 - [TokensAllOf](docs/TokensAllOf.md)
 - [TouchPoint](docs/TouchPoint.md)
 - [TouchPointColorValues](docs/TouchPointColorValues.md)
 - [TouchPointColorValuesAllOf](docs/TouchPointColorValuesAllOf.md)
 - [TouchPoints](docs/TouchPoints.md)
 - [TouchPointsAllOf](docs/TouchPointsAllOf.md)
 - [TransactionInfo](docs/TransactionInfo.md)
 - [TransactionResult](docs/TransactionResult.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionType](docs/TransactionType.md)
 - [TransferDocumentDetails](docs/TransferDocumentDetails.md)
 - [TransferDocumentDetailsAllOf](docs/TransferDocumentDetailsAllOf.md)
 - [TransferDocuments](docs/TransferDocuments.md)
 - [TransferDocumentsAllOf](docs/TransferDocumentsAllOf.md)
 - [UpdateInstallmentPlanRequest](docs/UpdateInstallmentPlanRequest.md)
 - [UpdateInstallmentPlanRequestAllOf](docs/UpdateInstallmentPlanRequestAllOf.md)
 - [UpdateInstallmentsPlanResponse](docs/UpdateInstallmentsPlanResponse.md)
 - [UpdateInstallmentsPlanResponseAllOf](docs/UpdateInstallmentsPlanResponseAllOf.md)
 - [User](docs/User.md)
 - [User2](docs/User2.md)
 - [User2AllOf](docs/User2AllOf.md)
 - [UserJobSubscriptions](docs/UserJobSubscriptions.md)
 - [UserJobSubscriptionsAllOf](docs/UserJobSubscriptionsAllOf.md)
 - [UserPasswordHistory](docs/UserPasswordHistory.md)
 - [UserPasswordHistoryAllOf](docs/UserPasswordHistoryAllOf.md)
 - [UserPermissionLevel](docs/UserPermissionLevel.md)
 - [UserStatus](docs/UserStatus.md)
 - [UserType](docs/UserType.md)
 - [UserWithActionItem](docs/UserWithActionItem.md)
 - [UserWithActionItemAllOf](docs/UserWithActionItemAllOf.md)
 - [VerifyPaymentRequest](docs/VerifyPaymentRequest.md)
 - [VerifyPaymentResponse](docs/VerifyPaymentResponse.md)
 - [VersionedTouchPoints](docs/VersionedTouchPoints.md)
 - [VersionedTouchPointsAllOf](docs/VersionedTouchPointsAllOf.md)
 - [ZipAddressDetails](docs/ZipAddressDetails.md)
 - [ZipAddressDetailsAllOf](docs/ZipAddressDetailsAllOf.md)


## Documentation For Authorization

Usually, login request is performed server-side, sessionId is acquired and publicToken is obtained by calling InstallmentPlanApi.InstallmentPlanInitiate.
For that particular transaction, publicToken should be used on client side for maximum security.

## Author

[Splitit](https://www.splitit.com/developers/integrations/hosted-solution/)
