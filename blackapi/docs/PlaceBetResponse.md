# PlaceBetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | [**OrderItem**](OrderItem.md) |  | 

## Methods

### NewPlaceBetResponse

`func NewPlaceBetResponse(status string, data OrderItem, ) *PlaceBetResponse`

NewPlaceBetResponse instantiates a new PlaceBetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceBetResponseWithDefaults

`func NewPlaceBetResponseWithDefaults() *PlaceBetResponse`

NewPlaceBetResponseWithDefaults instantiates a new PlaceBetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PlaceBetResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlaceBetResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlaceBetResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *PlaceBetResponse) GetData() OrderItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlaceBetResponse) GetDataOk() (*OrderItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlaceBetResponse) SetData(v OrderItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


