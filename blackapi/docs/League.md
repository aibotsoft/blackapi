# League

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompBookie** | Pointer to **bool** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rank** | Pointer to **int64** |  | [optional] 
**Events** | Pointer to [**map[string]Event**](Event.md) |  | [optional] 

## Methods

### NewLeague

`func NewLeague() *League`

NewLeague instantiates a new League object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeagueWithDefaults

`func NewLeagueWithDefaults() *League`

NewLeagueWithDefaults instantiates a new League object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompBookie

`func (o *League) GetCompBookie() bool`

GetCompBookie returns the CompBookie field if non-nil, zero value otherwise.

### GetCompBookieOk

`func (o *League) GetCompBookieOk() (*bool, bool)`

GetCompBookieOk returns a tuple with the CompBookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompBookie

`func (o *League) SetCompBookie(v bool)`

SetCompBookie sets CompBookie field to given value.

### HasCompBookie

`func (o *League) HasCompBookie() bool`

HasCompBookie returns a boolean if a field has been set.

### GetCountry

`func (o *League) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *League) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *League) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *League) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetName

`func (o *League) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *League) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *League) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *League) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRank

`func (o *League) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *League) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *League) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *League) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetEvents

`func (o *League) GetEvents() map[string]Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *League) GetEventsOk() (*map[string]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *League) SetEvents(v map[string]Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *League) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


