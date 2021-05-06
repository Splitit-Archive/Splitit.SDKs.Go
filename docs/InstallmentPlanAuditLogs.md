# InstallmentPlanAuditLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**ActivityExecutionDate** | [**time.Time**](time.Time.md) |  | 
**UserId** | **int64** |  | [optional] 
**UserUniqueId** | **string** |  | [optional] 
**BusinessActivity** | [**BusinessActivity**](BusinessActivity.md) |  | 
**TraceId** | **string** |  | [optional] 
**Result** | **bool** |  | 
**VersionedTouchPoint** | [**VersionedTouchPoints**](VersionedTouchPoints.md) |  | [optional] 
**VersionedTouchPointId** | **int64** |  | [optional] 
**Discriminator** | **string** |  | [optional] 
**AdditionalInfo** | **string** |  | [optional] 
**UserType** | [**UserType**](UserType.md) |  | [optional] 
**InstallmentPlanId** | **int64** |  | [optional] 
**CreatingIPAddress** | **string** |  | [optional] 
**CardId** | **int64** |  | [optional] 
**Card** | [**Cards**](Cards.md) |  | [optional] 
**InstallmentPlan** | [**InstallmentPlans**](InstallmentPlans.md) |  | [optional] 
**ErrorCode** | [**PisErrorCodes**](PisErrorCodes.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


