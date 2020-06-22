# \InfrastructureApi

All URIs are relative to *https://webapi.production.splitit.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfrastructureGetResources**](InfrastructureApi.md#InfrastructureGetResources) | **Post** /api/Infrastructure/GetResources | 
[**InfrastructureGetResources2**](InfrastructureApi.md#InfrastructureGetResources2) | **Get** /api/Infrastructure/GetResources | 



## InfrastructureGetResources

> GetResourcesResponse InfrastructureGetResources(ctx, request)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**request** | [**GetResourcesRequest**](GetResourcesRequest.md)|  | 

### Return type

[**GetResourcesResponse**](GetResourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InfrastructureGetResources2

> GetResourcesResponse InfrastructureGetResources2(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InfrastructureGetResources2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InfrastructureGetResources2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **optional.String**|  | 
 **sessionId** | **optional.String**|  | 
 **merchantCode** | **optional.String**|  | 
 **cultureName** | **optional.String**|  | 
 **touchPointCode** | **optional.String**|  | 
 **systemTextCategories** | [**optional.Interface of []SystemTextCategory**](SystemTextCategory.md)|  | 

### Return type

[**GetResourcesResponse**](GetResourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

