# \CreateInstallmentPlanApi

All URIs are relative to *https://webapi.production.splitit.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstallmentPlanGet**](CreateInstallmentPlanApi.md#CreateInstallmentPlanGet) | **Get** /api/CreateInstallmentPlan | 



## CreateInstallmentPlanGet

> CreateInstallmentPlanLegacyResponse CreateInstallmentPlanGet(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateInstallmentPlanGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInstallmentPlanGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amountBeforeFees** | **optional.Float32**|  | 
 **apiKey** | **optional.String**|  | 
 **avsAddress** | **optional.String**|  | 
 **avsZip** | **optional.String**|  | 
 **cardCvv** | **optional.String**|  | 
 **cardExpMonth** | **optional.String**|  | 
 **cardExpYear** | **optional.String**|  | 
 **cardHolder** | **optional.String**|  | 
 **cardNumber** | **optional.String**|  | 
 **cardTypeId** | **optional.Int32**|  | 
 **consumerFullName** | **optional.String**|  | 
 **countryId** | **optional.Int32**|  | 
 **email** | **optional.String**|  | 
 **installmentNumber** | **optional.Int32**|  | 
 **paramX** | **optional.String**|  | 
 **sessionId** | **optional.String**|  | 

### Return type

[**CreateInstallmentPlanLegacyResponse**](CreateInstallmentPlanLegacyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

