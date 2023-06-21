/*
 * CoinMarketCap Cryptocurrency API Documentation
 *
 * CoinMarketCap API
 *
 * API version: 1.26.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cmc

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type CryptocurrencyApiService service

/*
CryptocurrencyApiService CoinMarketCap ID Map
Returns a mapping of all cryptocurrencies to unique CoinMarketCap &#x60;id&#x60;s. Per our &lt;a href&#x3D;\&quot;#section/Best-Practices\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Best Practices&lt;/a&gt; we recommend utilizing CMC ID instead of cryptocurrency symbols to securely identify cryptocurrencies with our other endpoints and in your own application logic.  Each cryptocurrency returned includes typical identifiers such as &#x60;name&#x60;, &#x60;symbol&#x60;, and &#x60;token_address&#x60; for flexible mapping to &#x60;id&#x60;.    By default this endpoint returns cryptocurrencies that have actively tracked markets on supported exchanges. You may receive a map of all inactive cryptocurrencies by passing &#x60;listing_status&#x3D;inactive&#x60;. You may also receive a map of registered cryptocurrency projects that are listed but do not yet meet methodology requirements to have tracked markets via &#x60;listing_status&#x3D;untracked&#x60;. Please review our &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;\&quot;https://coinmarketcap.com/methodology/\&quot;&gt;methodology documentation&lt;/a&gt; for additional details on listing states.    Cryptocurrencies returned include &#x60;first_historical_data&#x60; and &#x60;last_historical_data&#x60; timestamps to conveniently reference historical date ranges available to query with historical time-series data endpoints. You may also use the &#x60;aux&#x60; parameter to only include properties you require to slim down the payload if calling this endpoint frequently.    **This endpoint is available on the following &lt;a href&#x3D;\&quot;https://coinmarketcap.com/api/features\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API plans&lt;/a&gt;:**   - Basic   - Hobbyist   - Startup   - Standard   - Professional   - Enterprise  **Cache / Update frequency:** Mapping data is updated only as needed, every 30 seconds. **Plan credit use:** 1 API call credit per request no matter query size. **CMC equivalent pages:** No equivalent, this data is only available via API.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CryptocurrencyApiGetV1CryptocurrencyMapOpts - Optional Parameters:
     * @param "ListingStatus" (optional.String) -  Only active cryptocurrencies are returned by default. Pass &#x60;inactive&#x60; to get a list of cryptocurrencies that are no longer active. Pass &#x60;untracked&#x60; to get a list of cryptocurrencies that are listed but do not yet meet methodology requirements to have tracked markets available. You may pass one or more comma-separated values.
     * @param "Start" (optional.Int32) -  Optionally offset the start (1-based index) of the paginated list of items to return.
     * @param "Limit" (optional.Int32) -  Optionally specify the number of results to return. Use this parameter and the \&quot;start\&quot; parameter to determine your own pagination size.
     * @param "Sort" (optional.String) -  What field to sort the list of cryptocurrencies by.
     * @param "Symbol" (optional.String) -  Optionally pass a comma-separated list of cryptocurrency symbols to return CoinMarketCap IDs for. If this option is passed, other options will be ignored.
     * @param "Aux" (optional.String) -  Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;platform,first_historical_data,last_historical_data,is_active,status&#x60; to include all auxiliary fields.
@return CryptocurrencyMapResponseModel
*/

type CryptocurrencyApiGetV1CryptocurrencyMapOpts struct {
	ListingStatus optional.String
	Start         optional.Int32
	Limit         optional.Int32
	Sort          optional.String
	Symbol        optional.String
	Aux           optional.String
}

func (a *CryptocurrencyApiService) GetV1CryptocurrencyMap(ctx context.Context, localVarOptionals *CryptocurrencyApiGetV1CryptocurrencyMapOpts) (CryptocurrencyMapResponseModel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue CryptocurrencyMapResponseModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/cryptocurrency/map"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ListingStatus.IsSet() {
		localVarQueryParams.Add("listing_status", parameterToString(localVarOptionals.ListingStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Symbol.IsSet() {
		localVarQueryParams.Add("symbol", parameterToString(localVarOptionals.Symbol.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Aux.IsSet() {
		localVarQueryParams.Add("aux", parameterToString(localVarOptionals.Aux.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CryptocurrencyMapResponseModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v HttpStatus400ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v HttpStatus401ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v HttpStatus403ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v HttpStatus429ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v HttpStatus500ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CryptocurrencyApiService Quotes Latest
Returns the latest market quote for 1 or more cryptocurrencies. Use the \&quot;convert\&quot; option to return market values in multiple fiat and cryptocurrency conversions in the same call.  **This endpoint is available on the following &lt;a href&#x3D;\&quot;https://coinmarketcap.com/api/features\&quot; target&#x3D;\&quot;_blank\&quot;&gt;API plans&lt;/a&gt;:** - Basic - Startup - Hobbyist - Standard - Professional - Enterprise  **Cache / Update frequency:** Every 60 seconds. **Plan credit use:** 1 call credit per 100 cryptocurrencies returned (rounded up) and 1 call credit per &#x60;convert&#x60; option beyond the first. **CMC equivalent pages:** Latest market data pages for specific cryptocurrencies like [coinmarketcap.com/currencies/bitcoin/](https://coinmarketcap.com/currencies/bitcoin/). ***NOTE:** Use this endpoint to request the latest quote for specific cryptocurrencies. If you need to request all cryptocurrencies use [/v1/cryptocurrency/listings/latest](#operation/getV1CryptocurrencyListingsLatest) which is optimized for that purpose. The response data between these endpoints is otherwise the same.*
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CryptocurrencyApiGetV1CryptocurrencyQuotesLatestOpts - Optional Parameters:
     * @param "Id" (optional.String) -  One or more comma-separated cryptocurrency CoinMarketCap IDs. Example: 1,2
     * @param "Slug" (optional.String) -  Alternatively pass a comma-separated list of cryptocurrency slugs. Example: \&quot;bitcoin,ethereum\&quot;
     * @param "Symbol" (optional.String) -  Alternatively pass one or more comma-separated cryptocurrency symbols. Example: \&quot;BTC,ETH\&quot;. At least one \&quot;id\&quot; *or* \&quot;slug\&quot; *or* \&quot;symbol\&quot; is required for this request.
     * @param "Convert" (optional.String) -  Optionally calculate market quotes in up to 120 currencies at once by passing a comma-separated list of cryptocurrency or fiat currency symbols. Each additional convert option beyond the first requires an additional call credit. A list of supported fiat options can be found [here](#section/Standards-and-Conventions). Each conversion is returned in its own \&quot;quote\&quot; object.
     * @param "ConvertId" (optional.String) -  Optionally calculate market quotes by CoinMarketCap ID instead of symbol. This option is identical to &#x60;convert&#x60; outside of ID format. Ex: convert_id&#x3D;1,2781 would replace convert&#x3D;BTC,USD in your query. This parameter cannot be used when &#x60;convert&#x60; is used.
     * @param "Aux" (optional.String) -  Optionally specify a comma-separated list of supplemental data fields to return. Pass &#x60;num_market_pairs,cmc_rank,date_added,tags,platform,max_supply,circulating_supply,total_supply,market_cap_by_total_supply,volume_24h_reported,volume_7d,volume_7d_reported,volume_30d,volume_30d_reported,is_active,is_fiat&#x60; to include all auxiliary fields.
     * @param "SkipInvalid" (optional.Bool) -  Pass &#x60;true&#x60; to relax request validation rules. When requesting records on multiple cryptocurrencies an error is returned if no match is found for 1 or more requested cryptocurrencies. If set to true, invalid lookups will be skipped allowing valid cryptocurrencies to still be returned.
@return CryptocurrencyQuotesLatestResponseModel
*/

type CryptocurrencyApiGetV1CryptocurrencyQuotesLatestOpts struct {
	Id          optional.String
	Slug        optional.String
	Symbol      optional.String
	Convert     optional.String
	ConvertId   optional.String
	Aux         optional.String
	SkipInvalid optional.Bool
}

func (a *CryptocurrencyApiService) GetV1CryptocurrencyQuotesLatest(ctx context.Context, localVarOptionals *CryptocurrencyApiGetV1CryptocurrencyQuotesLatestOpts) (CryptocurrencyQuotesLatestResponseModel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue CryptocurrencyQuotesLatestResponseModel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v1/cryptocurrency/quotes/latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Slug.IsSet() {
		localVarQueryParams.Add("slug", parameterToString(localVarOptionals.Slug.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Symbol.IsSet() {
		localVarQueryParams.Add("symbol", parameterToString(localVarOptionals.Symbol.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Convert.IsSet() {
		localVarQueryParams.Add("convert", parameterToString(localVarOptionals.Convert.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConvertId.IsSet() {
		localVarQueryParams.Add("convert_id", parameterToString(localVarOptionals.ConvertId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Aux.IsSet() {
		localVarQueryParams.Add("aux", parameterToString(localVarOptionals.Aux.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipInvalid.IsSet() {
		localVarQueryParams.Add("skip_invalid", parameterToString(localVarOptionals.SkipInvalid.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CryptocurrencyQuotesLatestResponseModel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v HttpStatus400ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v HttpStatus401ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 403 {
			var v HttpStatus403ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 429 {
			var v HttpStatus429ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v HttpStatus500ErrorObject
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}