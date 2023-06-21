/*
 * CoinMarketCap Cryptocurrency API Documentation
 *
 * CoinMarketCap API
 *
 * API version: 1.26.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cmc

type CryptocurrencyListingsLatestResponseModel struct {
	// Array of cryptocurrency objects matching the list options.
	Data   []CryptocurrencyListingsLatestCryptocurrencyObject `json:"data"`
	Status *ApiStatusObject                                   `json:"status"`
}
