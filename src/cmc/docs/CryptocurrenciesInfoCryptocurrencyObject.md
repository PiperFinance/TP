# CryptocurrenciesInfoCryptocurrencyObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique CoinMarketCap ID for this cryptocurrency. | [default to null]
**Name** | **string** | The name of this cryptocurrency. | [default to null]
**Symbol** | **string** | The ticker symbol for this cryptocurrency. | [default to null]
**Category** | **string** | The category for this cryptocurrency. | [default to null]
**Slug** | **string** | The web URL friendly shorthand version of this cryptocurrency name. | [default to null]
**Logo** | **string** | Link to a CoinMarketCap hosted logo png for this cryptocurrency. 64px is default size returned. Replace \&quot;64x64\&quot; in the image path with these alternative sizes: 16, 32, 64, 128, 200 | [default to null]
**Description** | **string** | A CoinMarketCap supplied brief description of this cryptocurrency. This field will return null if a description is not available. | [default to null]
**DateAdded** | [**time.Time**](time.Time.md) | Timestamp (ISO 8601) of when this cryptocurrency was added to CoinMarketCap. | [default to null]
**Notice** | **string** | A [Markdown](https://commonmark.org/help/) formatted notice that may highlight a significant event or condition that is impacting the cryptocurrency or how it is displayed, otherwise null. A notice may highlight a recent or upcoming mainnet swap, symbol change, exploit event, or known issue with a particular exchange or market, for example. If present, this notice is also displayed in an alert banner at the top of the cryptocurrency&#x27;s page on coinmarketcap.com. | [default to null]
**Tags** | [***[]string**](array.md) |  | [default to null]
**Platform** | [***Platform**](platform.md) |  | [default to null]
**Urls** | [***CryptocurrenciesInfoUrlsObject**](cryptocurrencies-info-urls-object.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

