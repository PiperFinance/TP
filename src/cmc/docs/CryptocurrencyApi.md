# {{classname}}

All URIs are relative to *https://pro-api.coinmarketcap.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CryptocurrencyMap**](CryptocurrencyApi.md#GetV1CryptocurrencyMap) | **Get** /v1/cryptocurrency/map | CoinMarketCap ID Map
[**GetV1CryptocurrencyQuotesLatest**](CryptocurrencyApi.md#GetV1CryptocurrencyQuotesLatest) | **Get** /v1/cryptocurrency/quotes/latest | Quotes Latest

# **GetV1CryptocurrencyMap**
> CryptocurrencyMapResponseModel GetV1CryptocurrencyMap(ctx, optional)
CoinMarketCap ID Map

Returns a mapping of all cryptocurrencies to unique CoinMarketCap `id`s. Per our <a href=\"#section/Best-Practices\" target=\"_blank\">Best Practices</a> we recommend utilizing CMC ID instead of cryptocurrency symbols to securely identify cryptocurrencies with our other endpoints and in your own application logic.  Each cryptocurrency returned includes typical identifiers such as `name`, `symbol`, and `token_address` for flexible mapping to `id`.    By default this endpoint returns cryptocurrencies that have actively tracked markets on supported exchanges. You may receive a map of all inactive cryptocurrencies by passing `listing_status=inactive`. You may also receive a map of registered cryptocurrency projects that are listed but do not yet meet methodology requirements to have tracked markets via `listing_status=untracked`. Please review our <a target=\"_blank\" href=\"https://coinmarketcap.com/methodology/\">methodology documentation</a> for additional details on listing states.    Cryptocurrencies returned include `first_historical_data` and `last_historical_data` timestamps to conveniently reference historical date ranges available to query with historical time-series data endpoints. You may also use the `aux` parameter to only include properties you require to slim down the payload if calling this endpoint frequently.    **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Mapping data is updated only as needed, every 30 seconds. **Plan credit use:** 1 API call credit per request no matter query size. **CMC equivalent pages:** No equivalent, this data is only available via API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listingStatus** | **optional.String**| Only active cryptocurrencies are returned by default. Pass &#x60;inactive&#x60; to get a list of cryptocurrencies that are no longer active. Pass &#x60;untracked&#x60; to get a list of cryptocurrencies that are listed but do not yet meet methodology requirements to have tracked markets available. You may pass one or more comma-separated values. | [default to active]
 **start** | **optional.Int32**| Optionally offset the start (1-based index) of the paginated list of items to return. | 
 **limit** | **optional.Int32**| Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size. | 
 **sort** | **optional.String**| What field to sort the list of cryptocurrencies by. | 
 **symbol** | **optional.String**| Optionally pass a comma-separated list of cryptocurrency symbols to return CoinMarketCap IDs for. If this option is passed, other options will be ignored. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;platform,first_historical_data,last_historical_data,is_active,status&#x60; to include all auxiliary fields. | 

### Return type

[**CryptocurrencyMapResponseModel**](cryptocurrency-map-response-model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetV1CryptocurrencyQuotesLatest**
> CryptocurrencyQuotesLatestResponseModel GetV1CryptocurrencyQuotesLatest(ctx, optional)
Quotes Latest

Returns the latest market quote for 1 or more cryptocurrencies. Use the \"convert\" option to return market values in multiple fiat and cryptocurrency conversions in the same call.  **This endpoint is available on the following <a href=\"https://coinmarketcap.com/api/features\" target=\"_blank\">API plans</a>:** - Basic - Startup - Hobbyist - Standard - Professional - Enterprise  **Cache / Update frequency:** Every 60 seconds. **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per `convert` option beyond the first. **CMC equivalent pages:** Latest market data pages for specific cryptocurrencies like [coinmarketcap.com/currencies/bitcoin/](https://coinmarketcap.com/currencies/bitcoin/). ***NOTE:** Use this endpoint to request the latest quote for specific cryptocurrencies. If you need to request all cryptocurrencies use [/v1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CryptocurrencyApiGetV1CryptocurrencyQuotesLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptocurrencyApiGetV1CryptocurrencyQuotesLatestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| One or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2 | 
 **slug** | **optional.String**| Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot; | 
 **symbol** | **optional.String**| Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request. | 
 **convert** | **optional.String**| Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object. | 
 **convertId** | **optional.String**| Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used. | 
 **aux** | **optional.String**| Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_active,is_fiat&#x60; to include all auxiliary fields. | [default to num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,is_active,is_fiat]
 **skipInvalid** | **optional.Bool**| Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if no match is found for 1 or more requested cryptocurrencies. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned. | 

### Return type

[**CryptocurrencyQuotesLatestResponseModel**](cryptocurrency-quotes-latest-response-model.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

