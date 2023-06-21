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

// Standardized status object for API calls.
type ApiStatusObject struct {
	// Current timestamp (ISO 8601) on the server.
	Timestamp time.Time `json:"timestamp"`
	// An internal error code for the current error. If a unique platform error code is not available the HTTP status code is returned. `null` is returned if there is no error.
	ErrorCode int32 `json:"error_code"`
	// An error message to go along with the error code.
	ErrorMessage string `json:"error_message"`
	// Number of milliseconds taken to generate this response.
	Elapsed int32 `json:"elapsed"`
	// Number of API call credits that were used for this call.
	CreditCount int32 `json:"credit_count"`
}
