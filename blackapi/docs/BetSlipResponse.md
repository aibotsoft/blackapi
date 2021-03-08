# BetSlipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**BetSlipData**](BetSlipData.md) |  | 
**Status** | **string** |  | 

## Methods

### NewBetSlipResponse

`func NewBetSlipResponse(data BetSlipData, status string, ) *BetSlipResponse`

NewBetSlipResponse instantiates a new BetSlipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetSlipResponseWithDefaults

`func NewBetSlipResponseWithDefaults() *BetSlipResponse`

NewBetSlipResponseWithDefaults instantiates a new BetSlipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetSlipResponse) GetData() BetSlipData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetSlipResponse) GetDataOk() (*BetSlipData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetSlipResponse) SetData(v BetSlipData)`

SetData sets Data field to given value.


### GetStatus

`func (o *BetSlipResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BetSlipResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BetSlipResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


