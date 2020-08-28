# \InstallmentPlanApi

All URIs are relative to *https://webapi.production.splitit.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstallmentPlanApprove**](InstallmentPlanApi.md#InstallmentPlanApprove) | **Post** /api/InstallmentPlan/Approve | 
[**InstallmentPlanCancel**](InstallmentPlanApi.md#InstallmentPlanCancel) | **Post** /api/InstallmentPlan/Cancel | 
[**InstallmentPlanChargeBack**](InstallmentPlanApi.md#InstallmentPlanChargeBack) | **Post** /api/InstallmentPlan/ChargeBack | 
[**InstallmentPlanCreate**](InstallmentPlanApi.md#InstallmentPlanCreate) | **Post** /api/InstallmentPlan/Create | 
[**InstallmentPlanGet**](InstallmentPlanApi.md#InstallmentPlanGet) | **Post** /api/InstallmentPlan/Get | 
[**InstallmentPlanGet3DSecureParameters**](InstallmentPlanApi.md#InstallmentPlanGet3DSecureParameters) | **Post** /api/InstallmentPlan/Get3DSecureParameters | 
[**InstallmentPlanGetExtended**](InstallmentPlanApi.md#InstallmentPlanGetExtended) | **Post** /api/InstallmentPlan/GetExtended | 
[**InstallmentPlanGetFraudStatusDisplay**](InstallmentPlanApi.md#InstallmentPlanGetFraudStatusDisplay) | **Post** /api/InstallmentPlan/GetFraudStatusDisplay | 
[**InstallmentPlanGetInitiatedInstallmentPlanRequest**](InstallmentPlanApi.md#InstallmentPlanGetInitiatedInstallmentPlanRequest) | **Post** /api/InstallmentPlan/GetInitiatedInstallmentPlanRequest | 
[**InstallmentPlanGetLearnMoreDetails**](InstallmentPlanApi.md#InstallmentPlanGetLearnMoreDetails) | **Post** /api/InstallmentPlan/GetLearnMoreDetails | 
[**InstallmentPlanGetSchedules**](InstallmentPlanApi.md#InstallmentPlanGetSchedules) | **Post** /api/InstallmentPlan/GetSchedules | 
[**InstallmentPlanInitiate**](InstallmentPlanApi.md#InstallmentPlanInitiate) | **Post** /api/InstallmentPlan/Initiate | 
[**InstallmentPlanRefund**](InstallmentPlanApi.md#InstallmentPlanRefund) | **Post** /api/InstallmentPlan/Refund | 
[**InstallmentPlanStartInstallments**](InstallmentPlanApi.md#InstallmentPlanStartInstallments) | **Post** /api/InstallmentPlan/StartInstallments | 
[**InstallmentPlanTermsAndConditions**](InstallmentPlanApi.md#InstallmentPlanTermsAndConditions) | **Post** /api/InstallmentPlan/TermsAndConditions | 
[**InstallmentPlanUpdate**](InstallmentPlanApi.md#InstallmentPlanUpdate) | **Post** /api/InstallmentPlan/Update | 
[**InstallmentPlanVerifyPayment**](InstallmentPlanApi.md#InstallmentPlanVerifyPayment) | **Post** /api/InstallmentPlan/Get/VerifyPayment | 



## InstallmentPlanApprove

> InstallmentPlanResponse InstallmentPlanApprove(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**ApproveInstallmentPlanRequest**](ApproveInstallmentPlanRequest.md)|  | 

### Return type

[**InstallmentPlanResponse**](InstallmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanCancel

> InstallmentPlanResponse InstallmentPlanCancel(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**CancelInstallmentPlanRequest**](CancelInstallmentPlanRequest.md)|  | 

### Return type

[**InstallmentPlanResponse**](InstallmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanChargeBack

> InstallmentPlanResponse InstallmentPlanChargeBack(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**ChargebackRequest**](ChargebackRequest.md)|  | 

### Return type

[**InstallmentPlanResponse**](InstallmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanCreate

> CreateInstallmentsPlanResponse InstallmentPlanCreate(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**CreateInstallmentPlanRequest**](CreateInstallmentPlanRequest.md)|  | 

### Return type

[**CreateInstallmentsPlanResponse**](CreateInstallmentsPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGet

> GetInstallmentsPlanResponse InstallmentPlanGet(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**GetInstallmentsPlanSearchCriteriaRequest**](GetInstallmentsPlanSearchCriteriaRequest.md)|  | 

### Return type

[**GetInstallmentsPlanResponse**](GetInstallmentsPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGet3DSecureParameters

> Get3DSecureParametersResponse InstallmentPlanGet3DSecureParameters(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**Get3DSecureParametersRequest**](Get3DSecureParametersRequest.md)|  | 

### Return type

[**Get3DSecureParametersResponse**](Get3DSecureParametersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGetExtended

> GetInstallmentsPlanExtendedResponse InstallmentPlanGetExtended(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**GetInstallmentsPlanSearchCriteriaRequest**](GetInstallmentsPlanSearchCriteriaRequest.md)|  | 

### Return type

[**GetInstallmentsPlanExtendedResponse**](GetInstallmentsPlanExtendedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGetFraudStatusDisplay

> GetFraudStatusDisplayResponse InstallmentPlanGetFraudStatusDisplay(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**GetFraudStatusDisplayRequest**](GetFraudStatusDisplayRequest.md)|  | 

### Return type

[**GetFraudStatusDisplayResponse**](GetFraudStatusDisplayResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGetInitiatedInstallmentPlanRequest

> GetInitiatedInstallmentPlanResponse InstallmentPlanGetInitiatedInstallmentPlanRequest(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**GetInitiatedInstallmentPlanRequest**](GetInitiatedInstallmentPlanRequest.md)|  | 

### Return type

[**GetInitiatedInstallmentPlanResponse**](GetInitiatedInstallmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGetLearnMoreDetails

> LearnMoreDetailsResponse InstallmentPlanGetLearnMoreDetails(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**LearnMoreDetailsRequest**](LearnMoreDetailsRequest.md)|  | 

### Return type

[**LearnMoreDetailsResponse**](LearnMoreDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanGetSchedules

> GetInstallmentsScheduleResponse InstallmentPlanGetSchedules(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**GetInstallmentSchedulesRequest**](GetInstallmentSchedulesRequest.md)|  | 

### Return type

[**GetInstallmentsScheduleResponse**](GetInstallmentsScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanInitiate

> InitiateInstallmentsPlanResponse InstallmentPlanInitiate(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**InitiateInstallmentPlanRequest**](InitiateInstallmentPlanRequest.md)|  | 

### Return type

[**InitiateInstallmentsPlanResponse**](InitiateInstallmentsPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanRefund

> RefundInstallmentPlanResponse InstallmentPlanRefund(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**RefundPlanRequest**](RefundPlanRequest.md)|  | 

### Return type

[**RefundInstallmentPlanResponse**](RefundInstallmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanStartInstallments

> InstallmentPlanResponse InstallmentPlanStartInstallments(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**StartInstallmentsRequest**](StartInstallmentsRequest.md)|  | 

### Return type

[**InstallmentPlanResponse**](InstallmentPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanTermsAndConditions

> TermsAndConditionsGetResponse InstallmentPlanTermsAndConditions(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**TermsAndConditionsGetRequest**](TermsAndConditionsGetRequest.md)|  | 

### Return type

[**TermsAndConditionsGetResponse**](TermsAndConditionsGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanUpdate

> UpdateInstallmentsPlanResponse InstallmentPlanUpdate(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**UpdateInstallmentPlanRequest**](UpdateInstallmentPlanRequest.md)|  | 

### Return type

[**UpdateInstallmentsPlanResponse**](UpdateInstallmentsPlanResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallmentPlanVerifyPayment

> VerifyPaymentResponse InstallmentPlanVerifyPayment(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**VerifyPaymentRequest**](VerifyPaymentRequest.md)|  | 

### Return type

[**VerifyPaymentResponse**](VerifyPaymentResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

