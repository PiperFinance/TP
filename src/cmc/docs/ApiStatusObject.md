# ApiStatusObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**time.Time**](time.Time.md) | Current timestamp (ISO 8601) on the server. | [default to null]
**ErrorCode** | **int32** | An internal error code for the current error. If a unique platform error code is not available the HTTP status code is returned. &#x60;null&#x60; is returned if there is no error. | [default to null]
**ErrorMessage** | **string** | An error message to go along with the error code. | [default to null]
**Elapsed** | **int32** | Number of milliseconds taken to generate this response. | [default to null]
**CreditCount** | **int32** | Number of API call credits that were used for this call. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

