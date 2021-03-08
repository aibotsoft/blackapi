# PlaceBetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetslipId** | **string** |  | 
**Price** | **float64** |  | 
**Stake** | **[]interface{}** |  | 
**IgnoreSystemMaintenance** | **bool** |  | 
**NoPutOfferExchange** | **bool** |  | 
**AdaptiveBookies** | Pointer to **[]string** |  | [optional] 
**RequestUuid** | **string** |  | 
**Duration** | **int64** |  | 
**Accounts** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewPlaceBetRequest

`func NewPlaceBetRequest(betslipId string, price float64, stake []interface{}, ignoreSystemMaintenance bool, noPutOfferExchange bool, requestUuid string, duration int64, ) *PlaceBetRequest`

NewPlaceBetRequest instantiates a new PlaceBetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceBetRequestWithDefaults

`func NewPlaceBetRequestWithDefaults() *PlaceBetRequest`

NewPlaceBetRequestWithDefaults instantiates a new PlaceBetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetslipId

`func (o *PlaceBetRequest) GetBetslipId() string`

GetBetslipId returns the BetslipId field if non-nil, zero value otherwise.

### GetBetslipIdOk

`func (o *PlaceBetRequest) GetBetslipIdOk() (*string, bool)`

GetBetslipIdOk returns a tuple with the BetslipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetslipId

`func (o *PlaceBetRequest) SetBetslipId(v string)`

SetBetslipId sets BetslipId field to given value.


### GetPrice

`func (o *PlaceBetRequest) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceBetRequest) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceBetRequest) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetStake

`func (o *PlaceBetRequest) GetStake() []interface{}`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *PlaceBetRequest) GetStakeOk() (*[]interface{}, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *PlaceBetRequest) SetStake(v []interface{})`

SetStake sets Stake field to given value.


### GetIgnoreSystemMaintenance

`func (o *PlaceBetRequest) GetIgnoreSystemMaintenance() bool`

GetIgnoreSystemMaintenance returns the IgnoreSystemMaintenance field if non-nil, zero value otherwise.

### GetIgnoreSystemMaintenanceOk

`func (o *PlaceBetRequest) GetIgnoreSystemMaintenanceOk() (*bool, bool)`

GetIgnoreSystemMaintenanceOk returns a tuple with the IgnoreSystemMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSystemMaintenance

`func (o *PlaceBetRequest) SetIgnoreSystemMaintenance(v bool)`

SetIgnoreSystemMaintenance sets IgnoreSystemMaintenance field to given value.


### GetNoPutOfferExchange

`func (o *PlaceBetRequest) GetNoPutOfferExchange() bool`

GetNoPutOfferExchange returns the NoPutOfferExchange field if non-nil, zero value otherwise.

### GetNoPutOfferExchangeOk

`func (o *PlaceBetRequest) GetNoPutOfferExchangeOk() (*bool, bool)`

GetNoPutOfferExchangeOk returns a tuple with the NoPutOfferExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPutOfferExchange

`func (o *PlaceBetRequest) SetNoPutOfferExchange(v bool)`

SetNoPutOfferExchange sets NoPutOfferExchange field to given value.


### GetAdaptiveBookies

`func (o *PlaceBetRequest) GetAdaptiveBookies() []string`

GetAdaptiveBookies returns the AdaptiveBookies field if non-nil, zero value otherwise.

### GetAdaptiveBookiesOk

`func (o *PlaceBetRequest) GetAdaptiveBookiesOk() (*[]string, bool)`

GetAdaptiveBookiesOk returns a tuple with the AdaptiveBookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveBookies

`func (o *PlaceBetRequest) SetAdaptiveBookies(v []string)`

SetAdaptiveBookies sets AdaptiveBookies field to given value.

### HasAdaptiveBookies

`func (o *PlaceBetRequest) HasAdaptiveBookies() bool`

HasAdaptiveBookies returns a boolean if a field has been set.

### GetRequestUuid

`func (o *PlaceBetRequest) GetRequestUuid() string`

GetRequestUuid returns the RequestUuid field if non-nil, zero value otherwise.

### GetRequestUuidOk

`func (o *PlaceBetRequest) GetRequestUuidOk() (*string, bool)`

GetRequestUuidOk returns a tuple with the RequestUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUuid

`func (o *PlaceBetRequest) SetRequestUuid(v string)`

SetRequestUuid sets RequestUuid field to given value.


### GetDuration

`func (o *PlaceBetRequest) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PlaceBetRequest) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PlaceBetRequest) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### GetAccounts

`func (o *PlaceBetRequest) GetAccounts() [][]string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PlaceBetRequest) GetAccountsOk() (*[][]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PlaceBetRequest) SetAccounts(v [][]string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PlaceBetRequest) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


