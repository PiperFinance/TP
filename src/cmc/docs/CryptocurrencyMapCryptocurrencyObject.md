# CryptocurrencyMapCryptocurrencyObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique cryptocurrency ID for this cryptocurrency. | [default to null]
**Name** | **string** | The name of this cryptocurrency. | [default to null]
**Symbol** | **string** | The ticker symbol for this cryptocurrency, always in all caps. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of this cryptocurrency name. | [default to null]
**IsActive** | **int32** | 1 if this cryptocurrency has at least 1 active market currently being tracked by the platform, otherwise 0. A value of 1 is analogous with &#x60;listing_status&#x3D;active&#x60;. | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**FirstHistoricalData** | [**time.Time**](time.Time.md) | Timestamp (ISO 8601) of the date this cryptocurrency was first available on the platform. | [optional] [default to null]
**LastHistoricalData** | [**time.Time**](time.Time.md) | Timestamp (ISO 8601) of the last time this cryptocurrency&#x27;s market data was updated. | [optional] [default to null]
**Platform** | [***Platform**](platform.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

