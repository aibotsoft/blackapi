# ErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationErrors** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewErrorData

`func NewErrorData() *ErrorData`

NewErrorData instantiates a new ErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDataWithDefaults

`func NewErrorDataWithDefaults() *ErrorData`

NewErrorDataWithDefaults instantiates a new ErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationErrors

`func (o *ErrorData) GetValidationErrors() map[string][]string`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ErrorData) GetValidationErrorsOk() (*map[string][]string, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ErrorData) SetValidationErrors(v map[string][]string)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ErrorData) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


