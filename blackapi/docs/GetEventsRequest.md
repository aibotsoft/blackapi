# GetEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludePrices** | **bool** |  | 
**AllEvents** | **bool** |  | 
**AllHcaps** | Pointer to **bool** |  | [optional] 
**EventIds** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewGetEventsRequest

`func NewGetEventsRequest(includePrices bool, allEvents bool, ) *GetEventsRequest`

NewGetEventsRequest instantiates a new GetEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsRequestWithDefaults

`func NewGetEventsRequestWithDefaults() *GetEventsRequest`

NewGetEventsRequestWithDefaults instantiates a new GetEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludePrices

`func (o *GetEventsRequest) GetIncludePrices() bool`

GetIncludePrices returns the IncludePrices field if non-nil, zero value otherwise.

### GetIncludePricesOk

`func (o *GetEventsRequest) GetIncludePricesOk() (*bool, bool)`

GetIncludePricesOk returns a tuple with the IncludePrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrices

`func (o *GetEventsRequest) SetIncludePrices(v bool)`

SetIncludePrices sets IncludePrices field to given value.


### GetAllEvents

`func (o *GetEventsRequest) GetAllEvents() bool`

GetAllEvents returns the AllEvents field if non-nil, zero value otherwise.

### GetAllEventsOk

`func (o *GetEventsRequest) GetAllEventsOk() (*bool, bool)`

GetAllEventsOk returns a tuple with the AllEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllEvents

`func (o *GetEventsRequest) SetAllEvents(v bool)`

SetAllEvents sets AllEvents field to given value.


### GetAllHcaps

`func (o *GetEventsRequest) GetAllHcaps() bool`

GetAllHcaps returns the AllHcaps field if non-nil, zero value otherwise.

### GetAllHcapsOk

`func (o *GetEventsRequest) GetAllHcapsOk() (*bool, bool)`

GetAllHcapsOk returns a tuple with the AllHcaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllHcaps

`func (o *GetEventsRequest) SetAllHcaps(v bool)`

SetAllHcaps sets AllHcaps field to given value.

### HasAllHcaps

`func (o *GetEventsRequest) HasAllHcaps() bool`

HasAllHcaps returns a boolean if a field has been set.

### GetEventIds

`func (o *GetEventsRequest) GetEventIds() [][]string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *GetEventsRequest) GetEventIdsOk() (*[][]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *GetEventsRequest) SetEventIds(v [][]string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *GetEventsRequest) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


