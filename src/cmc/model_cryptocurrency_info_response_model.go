/*
 * CoinMarketCap Cryptocurrency API Documentation
 *
 * CoinMarketCap API
 *
 * API version: 1.26.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cmc

type CryptocurrencyInfoResponseModel struct {
	Data   *map[string]CryptocurrenciesInfoCryptocurrencyObject `json:"data,omitempty"`
	Status *ApiStatusObject                                     `json:"status,omitempty"`
}