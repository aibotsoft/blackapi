# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Away** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **string** |  | [optional] 
**IrStatus** | Pointer to **string** |  | [optional] 
**Sports** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAway

`func (o *Event) GetAway() string`

GetAway returns the Away field if non-nil, zero value otherwise.

### GetAwayOk

`func (o *Event) GetAwayOk() (*string, bool)`

GetAwayOk returns a tuple with the Away field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAway

`func (o *Event) SetAway(v string)`

SetAway sets Away field to given value.

### HasAway

`func (o *Event) HasAway() bool`

HasAway returns a boolean if a field has been set.

### GetHome

`func (o *Event) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *Event) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *Event) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *Event) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetIrStatus

`func (o *Event) GetIrStatus() string`

GetIrStatus returns the IrStatus field if non-nil, zero value otherwise.

### GetIrStatusOk

`func (o *Event) GetIrStatusOk() (*string, bool)`

GetIrStatusOk returns a tuple with the IrStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrStatus

`func (o *Event) SetIrStatus(v string)`

SetIrStatus sets IrStatus field to given value.

### HasIrStatus

`func (o *Event) HasIrStatus() bool`

HasIrStatus returns a boolean if a field has been set.

### GetSports

`func (o *Event) GetSports() []string`

GetSports returns the Sports field if non-nil, zero value otherwise.

### GetSportsOk

`func (o *Event) GetSportsOk() (*[]string, bool)`

GetSportsOk returns a tuple with the Sports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSports

`func (o *Event) SetSports(v []string)`

SetSports sets Sports field to given value.

### HasSports

`func (o *Event) HasSports() bool`

HasSports returns a boolean if a field has been set.

### GetStartTime

`func (o *Event) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Event) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Event) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Event) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


