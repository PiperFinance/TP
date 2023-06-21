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
	"time"
)

// A market quote in the currency conversion option.
type CryptocurrencyListingsLatestQuoteObject struct {
	// Price in the specified currency for this historical.
	Price float64 `json:"price,omitempty"`
	// Rolling 24 hour adjusted volume in the specified currency.
	Volume24h float64 `json:"volume_24h,omitempty"`
	// Rolling 24 hour reported volume in the specified currency. *This field is only returned if requested through the `aux` request parameter.*
	Volume24hReported float64 `json:"volume_24h_reported,omitempty"`
	// Rolling 7 day adjusted volume in the specified currency. *This field is only returned if requested through the `aux` request parameter.*
	Volume7d float64 `json:"volume_7d,omitempty"`
	// Rolling 7 day reported volume in the specified currency. *This field is only returned if requested through the `aux` request parameter.*
	Volume7dReported float64 `json:"volume_7d_reported,omitempty"`
	// Rolling 30 day adjusted volume in the specified currency. *This field is only returned if requested through the `aux` request parameter.*
	Volume30d float64 `json:"volume_30d,omitempty"`
	// Rolling 30 day reported volume in the specified currency. *This field is only returned if requested through the `aux` request parameter.*
	Volume30dReported float64 `json:"volume_30d_reported,omitempty"`
	// Market cap in the specified currency.
	MarketCap float64 `json:"market_cap,omitempty"`
	// 1 hour change in the specified currency.
	PercentChange1h float64 `json:"percent_change_1h,omitempty"`
	// 24 hour change in the specified currency.
	PercentChange24h float64 `json:"percent_change_24h,omitempty"`
	// 7 day change in the specified currency.
	PercentChange7d float64 `json:"percent_change_7d,omitempty"`
	// Timestamp (ISO 8601) of when the conversion currency's current value was referenced.
	LastUpdated time.Time `json:"last_updated,omitempty"`
}