# BalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | [**[]BalanceItem**](BalanceItem.md) |  | 

## Methods

### NewBalanceResponse

`func NewBalanceResponse(status string, data []BalanceItem, ) *BalanceResponse`

NewBalanceResponse instantiates a new BalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceResponseWithDefaults

`func NewBalanceResponseWithDefaults() *BalanceResponse`

NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BalanceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BalanceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BalanceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *BalanceResponse) GetData() []BalanceItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BalanceResponse) GetDataOk() (*[]BalanceItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BalanceResponse) SetData(v []BalanceItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


