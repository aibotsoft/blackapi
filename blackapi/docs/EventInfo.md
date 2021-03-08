# EventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**HomeId** | **int64** |  | 
**HomeTeam** | **string** |  | 
**AwayId** | **int64** |  | 
**AwayTeam** | **string** |  | 
**CompetitionId** | **int64** |  | 
**CompetitionName** | **string** |  | 
**CompetitionCountry** | **string** |  | 
**StartTime** | **time.Time** |  | 
**Date** | **string** |  | 

## Methods

### NewEventInfo

`func NewEventInfo(eventId string, homeId int64, homeTeam string, awayId int64, awayTeam string, competitionId int64, competitionName string, competitionCountry string, startTime time.Time, date string, ) *EventInfo`

NewEventInfo instantiates a new EventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInfoWithDefaults

`func NewEventInfoWithDefaults() *EventInfo`

NewEventInfoWithDefaults instantiates a new EventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *EventInfo) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventInfo) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventInfo) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetHomeId

`func (o *EventInfo) GetHomeId() int64`

GetHomeId returns the HomeId field if non-nil, zero value otherwise.

### GetHomeIdOk

`func (o *EventInfo) GetHomeIdOk() (*int64, bool)`

GetHomeIdOk returns a tuple with the HomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeId

`func (o *EventInfo) SetHomeId(v int64)`

SetHomeId sets HomeId field to given value.


### GetHomeTeam

`func (o *EventInfo) GetHomeTeam() string`

GetHomeTeam returns the HomeTeam field if non-nil, zero value otherwise.

### GetHomeTeamOk

`func (o *EventInfo) GetHomeTeamOk() (*string, bool)`

GetHomeTeamOk returns a tuple with the HomeTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTeam

`func (o *EventInfo) SetHomeTeam(v string)`

SetHomeTeam sets HomeTeam field to given value.


### GetAwayId

`func (o *EventInfo) GetAwayId() int64`

GetAwayId returns the AwayId field if non-nil, zero value otherwise.

### GetAwayIdOk

`func (o *EventInfo) GetAwayIdOk() (*int64, bool)`

GetAwayIdOk returns a tuple with the AwayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwayId

`func (o *EventInfo) SetAwayId(v int64)`

SetAwayId sets AwayId field to given value.


### GetAwayTeam

`func (o *EventInfo) GetAwayTeam() string`

GetAwayTeam returns the AwayTeam field if non-nil, zero value otherwise.

### GetAwayTeamOk

`func (o *EventInfo) GetAwayTeamOk() (*string, bool)`

GetAwayTeamOk returns a tuple with the AwayTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwayTeam

`func (o *EventInfo) SetAwayTeam(v string)`

SetAwayTeam sets AwayTeam field to given value.


### GetCompetitionId

`func (o *EventInfo) GetCompetitionId() int64`

GetCompetitionId returns the CompetitionId field if non-nil, zero value otherwise.

### GetCompetitionIdOk

`func (o *EventInfo) GetCompetitionIdOk() (*int64, bool)`

GetCompetitionIdOk returns a tuple with the CompetitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionId

`func (o *EventInfo) SetCompetitionId(v int64)`

SetCompetitionId sets CompetitionId field to given value.


### GetCompetitionName

`func (o *EventInfo) GetCompetitionName() string`

GetCompetitionName returns the CompetitionName field if non-nil, zero value otherwise.

### GetCompetitionNameOk

`func (o *EventInfo) GetCompetitionNameOk() (*string, bool)`

GetCompetitionNameOk returns a tuple with the CompetitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionName

`func (o *EventInfo) SetCompetitionName(v string)`

SetCompetitionName sets CompetitionName field to given value.


### GetCompetitionCountry

`func (o *EventInfo) GetCompetitionCountry() string`

GetCompetitionCountry returns the CompetitionCountry field if non-nil, zero value otherwise.

### GetCompetitionCountryOk

`func (o *EventInfo) GetCompetitionCountryOk() (*string, bool)`

GetCompetitionCountryOk returns a tuple with the CompetitionCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionCountry

`func (o *EventInfo) SetCompetitionCountry(v string)`

SetCompetitionCountry sets CompetitionCountry field to given value.


### GetStartTime

`func (o *EventInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EventInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EventInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetDate

`func (o *EventInfo) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EventInfo) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EventInfo) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


