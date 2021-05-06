# CollectDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Amount** | **float32** |  | 
**CurrencyId** | **int64** |  | 
**BusinessUnitId** | **int64** |  | 
**ApprovalDate** | [**time.Time**](time.Time.md) |  | [optional] 
**Discriminator** | **string** |  | [optional] 
**LenderId** | **int64** |  | [optional] 
**ExcludeFromEmails** | **bool** |  | [optional] 
**BusinessUnit** | [**BusinessUnits**](BusinessUnits.md) |  | [optional] 
**Currency** | [**Currencies**](Currencies.md) |  | [optional] 
**Lender** | [**ScpProviders**](ScpProviders.md) |  | [optional] 
**FundingCollectDocumentDetails** | [**[]FundingCollectDocumentDetails**](FundingCollectDocumentDetails.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


