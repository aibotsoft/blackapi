# UnauthorizedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Code** | **string** |  | 
**Data** | Pointer to [**UnauthorizedErrorData**](UnauthorizedError_data.md) |  | [optional] 

## Methods

### NewUnauthorizedError

`func NewUnauthorizedError(status string, code string, ) *UnauthorizedError`

NewUnauthorizedError instantiates a new UnauthorizedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorWithDefaults

`func NewUnauthorizedErrorWithDefaults() *UnauthorizedError`

NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UnauthorizedError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnauthorizedError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnauthorizedError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *UnauthorizedError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthorizedError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthorizedError) SetCode(v string)`

SetCode sets Code field to given value.


### GetData

`func (o *UnauthorizedError) GetData() UnauthorizedErrorData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnauthorizedError) GetDataOk() (*UnauthorizedErrorData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnauthorizedError) SetData(v UnauthorizedErrorData)`

SetData sets Data field to given value.

### HasData

`func (o *UnauthorizedError) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


