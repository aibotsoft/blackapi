# GetEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**map[string]League**](League.md) |  | [optional] 

## Methods

### NewGetEventsResponse

`func NewGetEventsResponse() *GetEventsResponse`

NewGetEventsResponse instantiates a new GetEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsResponseWithDefaults

`func NewGetEventsResponseWithDefaults() *GetEventsResponse`

NewGetEventsResponseWithDefaults instantiates a new GetEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetEventsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEventsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEventsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEventsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GetEventsResponse) GetData() map[string]League`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventsResponse) GetDataOk() (*map[string]League, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventsResponse) SetData(v map[string]League)`

SetData sets Data field to given value.

### HasData

`func (o *GetEventsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


