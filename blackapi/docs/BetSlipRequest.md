# BetSlipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sport** | **string** |  | 
**EventId** | **string** |  | 
**BetType** | **string** |  | 
**EquivalentBets** | **bool** |  | 
**MultipleAccounts** | **bool** |  | 

## Methods

### NewBetSlipRequest

`func NewBetSlipRequest(sport string, eventId string, betType string, equivalentBets bool, multipleAccounts bool, ) *BetSlipRequest`

NewBetSlipRequest instantiates a new BetSlipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetSlipRequestWithDefaults

`func NewBetSlipRequestWithDefaults() *BetSlipRequest`

NewBetSlipRequestWithDefaults instantiates a new BetSlipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSport

`func (o *BetSlipRequest) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *BetSlipRequest) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *BetSlipRequest) SetSport(v string)`

SetSport sets Sport field to given value.


### GetEventId

`func (o *BetSlipRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BetSlipRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BetSlipRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetBetType

`func (o *BetSlipRequest) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *BetSlipRequest) GetBetTypeOk() (*string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetType

`func (o *BetSlipRequest) SetBetType(v string)`

SetBetType sets BetType field to given value.


### GetEquivalentBets

`func (o *BetSlipRequest) GetEquivalentBets() bool`

GetEquivalentBets returns the EquivalentBets field if non-nil, zero value otherwise.

### GetEquivalentBetsOk

`func (o *BetSlipRequest) GetEquivalentBetsOk() (*bool, bool)`

GetEquivalentBetsOk returns a tuple with the EquivalentBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquivalentBets

`func (o *BetSlipRequest) SetEquivalentBets(v bool)`

SetEquivalentBets sets EquivalentBets field to given value.


### GetMultipleAccounts

`func (o *BetSlipRequest) GetMultipleAccounts() bool`

GetMultipleAccounts returns the MultipleAccounts field if non-nil, zero value otherwise.

### GetMultipleAccountsOk

`func (o *BetSlipRequest) GetMultipleAccountsOk() (*bool, bool)`

GetMultipleAccountsOk returns a tuple with the MultipleAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAccounts

`func (o *BetSlipRequest) SetMultipleAccounts(v bool)`

SetMultipleAccounts sets MultipleAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


