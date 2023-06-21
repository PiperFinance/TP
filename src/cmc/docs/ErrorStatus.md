# ErrorStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | [**time.Time**](time.Time.md) | Current ISO 8601 timestamp on the server. | [optional] [default to null]
**ErrorCode** | **int32** | An internal error code string for the current error. If a unique platform error code is not available the HTTP status code is returned. | [optional] [default to null]
**ErrorMessage** | **string** | An error message to go along with the error code. | [optional] [default to null]
**Elapsed** | **int32** | Number of milliseconds taken to generate this response | [optional] [default to null]
**CreditCount** | **int32** | Number of API call credits required for this call. Always 0 for errors. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

