# Splitit SDK for Go language

This is Splitit Web API SDK source code for Go applications. For other languages, please visit [Splitit.SDKs](https://github.com/Splitit/Splitit.SDKs).

## Overview

- API version: 1.0.0
- Package version: 1.5.33

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
*InstallmentPlanApi* | [**InstallmentPlanGetLearnMoreDetails**](docs/InstallmentPlanApi.md#installmentplangetlearnmoredetails) | **Post** /api/InstallmentPlan/GetLearnMoreDetails | 
*InstallmentPlanApi* | [**InstallmentPlanGetSchedules**](docs/InstallmentPlanApi.md#installmentplangetschedules) | **Post** /api/InstallmentPlan/GetSchedules | 
*InstallmentPlanApi* | [**InstallmentPlanInitiate**](docs/InstallmentPlanApi.md#installmentplaninitiate) | **Post** /api/InstallmentPlan/Initiate | 
*InstallmentPlanApi* | [**InstallmentPlanRefund**](docs/InstallmentPlanApi.md#installmentplanrefund) | **Post** /api/InstallmentPlan/Refund | 
*InstallmentPlanApi* | [**InstallmentPlanStartInstallments**](docs/InstallmentPlanApi.md#installmentplanstartinstallments) | **Post** /api/InstallmentPlan/StartInstallments | 
*InstallmentPlanApi* | [**InstallmentPlanTermsAndConditions**](docs/InstallmentPlanApi.md#installmentplantermsandconditions) | **Post** /api/InstallmentPlan/TermsAndConditions | 
*InstallmentPlanApi* | [**InstallmentPlanUpdate**](docs/InstallmentPlanApi.md#installmentplanupdate) | **Post** /api/InstallmentPlan/Update | 
*InstallmentPlanApi* | [**InstallmentPlanVerifyPayment**](docs/InstallmentPlanApi.md#installmentplanverifypayment) | **Post** /api/InstallmentPlan/Get/VerifyPayment | 
*LoginApi* | [**LoginPost**](docs/LoginApi.md#loginpost) | **Post** /api/login | 
*LogoutApi* | [**LogoutPost**](docs/LogoutApi.md#logoutpost) | **Post** /api/logout | 


## Documentation For Models

 - [AddressData](docs/AddressData.md)
 - [AddressData2](docs/AddressData2.md)
 - [AmountDetails](docs/AmountDetails.md)
 - [AmountDetails2](docs/AmountDetails2.md)
 - [ApproveInstallmentPlanRequest](docs/ApproveInstallmentPlanRequest.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [CancelInstallmentPlanRequest](docs/CancelInstallmentPlanRequest.md)
 - [CardBrand](docs/CardBrand.md)
 - [CardData](docs/CardData.md)
 - [CardResult](docs/CardResult.md)
 - [CardType](docs/CardType.md)
 - [CartData](docs/CartData.md)
 - [ChargebackRequest](docs/ChargebackRequest.md)
 - [ConsumerData](docs/ConsumerData.md)
 - [CreateInstallmentPlanLegacyResponse](docs/CreateInstallmentPlanLegacyResponse.md)
 - [CreateInstallmentPlanRequest](docs/CreateInstallmentPlanRequest.md)
 - [CreateInstallmentsPlanResponse](docs/CreateInstallmentsPlanResponse.md)
 - [CreateInstallmentsPlanResponseAllOf](docs/CreateInstallmentsPlanResponseAllOf.md)
 - [Currency](docs/Currency.md)
 - [CurrencyAllOf](docs/CurrencyAllOf.md)
 - [DelayResolution](docs/DelayResolution.md)
 - [Error](docs/Error.md)
 - [EventsEndpoints](docs/EventsEndpoints.md)
 - [ExtendedCurrency](docs/ExtendedCurrency.md)
 - [ExtendedCurrencyAllOf](docs/ExtendedCurrencyAllOf.md)
 - [ExternalAuth](docs/ExternalAuth.md)
 - [FraudCheck](docs/FraudCheck.md)
 - [FraudCheckResult](docs/FraudCheckResult.md)
 - [Get3DSecureParametersRequest](docs/Get3DSecureParametersRequest.md)
 - [Get3DSecureParametersResponse](docs/Get3DSecureParametersResponse.md)
 - [GetFraudStatusDisplayRequest](docs/GetFraudStatusDisplayRequest.md)
 - [GetFraudStatusDisplayResponse](docs/GetFraudStatusDisplayResponse.md)
 - [GetInitiatedInstallmentPlanRequest](docs/GetInitiatedInstallmentPlanRequest.md)
 - [GetInitiatedInstallmentPlanResponse](docs/GetInitiatedInstallmentPlanResponse.md)
 - [GetInstallmentSchedulesRequest](docs/GetInstallmentSchedulesRequest.md)
 - [GetInstallmentsPlanExtendedResponse](docs/GetInstallmentsPlanExtendedResponse.md)
 - [GetInstallmentsPlanResponse](docs/GetInstallmentsPlanResponse.md)
 - [GetInstallmentsPlanSearchCriteriaRequest](docs/GetInstallmentsPlanSearchCriteriaRequest.md)
 - [GetInstallmentsScheduleResponse](docs/GetInstallmentsScheduleResponse.md)
 - [GetResourcesRequest](docs/GetResourcesRequest.md)
 - [GetResourcesRequestContext](docs/GetResourcesRequestContext.md)
 - [GetResourcesResponse](docs/GetResourcesResponse.md)
 - [InitiateInstallmentPlanRequest](docs/InitiateInstallmentPlanRequest.md)
 - [InitiateInstallmentsPlanResponse](docs/InitiateInstallmentsPlanResponse.md)
 - [InitiateInstallmentsPlanResponseAllOf](docs/InitiateInstallmentsPlanResponseAllOf.md)
 - [Installment](docs/Installment.md)
 - [Installment2](docs/Installment2.md)
 - [InstallmentPlan](docs/InstallmentPlan.md)
 - [InstallmentPlanActivityStatus](docs/InstallmentPlanActivityStatus.md)
 - [InstallmentPlanCancelationReason](docs/InstallmentPlanCancelationReason.md)
 - [InstallmentPlanDateInfo](docs/InstallmentPlanDateInfo.md)
 - [InstallmentPlanInitiatedStatuses](docs/InstallmentPlanInitiatedStatuses.md)
 - [InstallmentPlanQueryCriteria](docs/InstallmentPlanQueryCriteria.md)
 - [InstallmentPlanResponse](docs/InstallmentPlanResponse.md)
 - [InstallmentPlanStatus](docs/InstallmentPlanStatus.md)
 - [InstallmentsPlanDateType](docs/InstallmentsPlanDateType.md)
 - [ItemData](docs/ItemData.md)
 - [LearnMoreDetailsRequest](docs/LearnMoreDetailsRequest.md)
 - [LearnMoreDetailsResponse](docs/LearnMoreDetailsResponse.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [LoginResponse](docs/LoginResponse.md)
 - [LogoutResponse](docs/LogoutResponse.md)
 - [MerchantRef](docs/MerchantRef.md)
 - [MerchantRefAllOf](docs/MerchantRefAllOf.md)
 - [Money](docs/Money.md)
 - [MoneyWithCurrencyCode](docs/MoneyWithCurrencyCode.md)
 - [PagingRequestHeader](docs/PagingRequestHeader.md)
 - [PagingResponseHeader](docs/PagingResponseHeader.md)
 - [PaymentFormMessage](docs/PaymentFormMessage.md)
 - [PaymentFormMessageType](docs/PaymentFormMessageType.md)
 - [PaymentToken](docs/PaymentToken.md)
 - [PaymentWizardData](docs/PaymentWizardData.md)
 - [PaymentWizardDataResponse](docs/PaymentWizardDataResponse.md)
 - [PaymentWizardDataResponseAllOf](docs/PaymentWizardDataResponseAllOf.md)
 - [PlanApprovalEvidence](docs/PlanApprovalEvidence.md)
 - [PlanData](docs/PlanData.md)
 - [PlanStrategy](docs/PlanStrategy.md)
 - [PurchaseMethod](docs/PurchaseMethod.md)
 - [ReAuthorization](docs/ReAuthorization.md)
 - [RedirectUrls](docs/RedirectUrls.md)
 - [ReferenceEntityBase](docs/ReferenceEntityBase.md)
 - [RefundInstallmentPlanResponse](docs/RefundInstallmentPlanResponse.md)
 - [RefundInstallmentPlanResponseAllOf](docs/RefundInstallmentPlanResponseAllOf.md)
 - [RefundPlanRequest](docs/RefundPlanRequest.md)
 - [RefundStrategy](docs/RefundStrategy.md)
 - [RefundUnderCancelation](docs/RefundUnderCancelation.md)
 - [RelationsLoad](docs/RelationsLoad.md)
 - [RequestHeader](docs/RequestHeader.md)
 - [ResponseHeader](docs/ResponseHeader.md)
 - [Schedule](docs/Schedule.md)
 - [ScheduleElements](docs/ScheduleElements.md)
 - [StartInstallmentsRequest](docs/StartInstallmentsRequest.md)
 - [SystemTextCategory](docs/SystemTextCategory.md)
 - [TermsAndConditions](docs/TermsAndConditions.md)
 - [TermsAndConditionsGetRequest](docs/TermsAndConditionsGetRequest.md)
 - [TermsAndConditionsGetResponse](docs/TermsAndConditionsGetResponse.md)
 - [TestModes](docs/TestModes.md)
 - [TouchPoint](docs/TouchPoint.md)
 - [TransactionInfo](docs/TransactionInfo.md)
 - [TransactionResult](docs/TransactionResult.md)
 - [TransactionStatus](docs/TransactionStatus.md)
 - [TransactionType](docs/TransactionType.md)
 - [UpdateInstallmentPlanRequest](docs/UpdateInstallmentPlanRequest.md)
 - [UpdateInstallmentPlanRequestAllOf](docs/UpdateInstallmentPlanRequestAllOf.md)
 - [UpdateInstallmentsPlanResponse](docs/UpdateInstallmentsPlanResponse.md)
 - [User](docs/User.md)
 - [VerifyPaymentRequest](docs/VerifyPaymentRequest.md)
 - [VerifyPaymentResponse](docs/VerifyPaymentResponse.md)


## Documentation For Authorization

Usually, login request is performed server-side, sessionId is acquired and publicToken is obtained by calling InstallmentPlanApi.InstallmentPlanInitiate.
For that particular transaction, publicToken should be used on client side for maximum security.

## Author

[Splitit](https://www.splitit.com/developers/integrations/hosted-solution/)
