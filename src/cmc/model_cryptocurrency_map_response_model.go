/*
 * CoinMarketCap Cryptocurrency API Documentation
 *
 * CoinMarketCap API
 *
 * API version: 1.26.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cmc

type CryptocurrencyMapResponseModel struct {
	// Array of cryptocurrency object results.
	Data   []CryptocurrencyMapCryptocurrencyObject `json:"data"`
	Status *ApiStatusObject                        `json:"status"`
}