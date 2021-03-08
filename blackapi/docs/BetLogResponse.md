# BetLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | [**[]LogItem**](LogItem.md) |  | 

## Methods

### NewBetLogResponse

`func NewBetLogResponse(status string, data []LogItem, ) *BetLogResponse`

NewBetLogResponse instantiates a new BetLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetLogResponseWithDefaults

`func NewBetLogResponseWithDefaults() *BetLogResponse`

NewBetLogResponseWithDefaults instantiates a new BetLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BetLogResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BetLogResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BetLogResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *BetLogResponse) GetData() []LogItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetLogResponse) GetDataOk() (*[]LogItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetLogResponse) SetData(v []LogItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


