# BetSlipData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetslipId** | **string** |  | 
**Sport** | **string** |  | 
**EventId** | **string** |  | 
**BetType** | **string** |  | 
**BetTypeDescription** | Pointer to **string** |  | [optional] 
**BetTypeTemplate** | Pointer to **string** |  | [optional] 
**EquivalentBets** | Pointer to **bool** |  | [optional] 
**MultipleAccounts** | Pointer to **bool** |  | [optional] 
**IsOpen** | Pointer to **bool** |  | [optional] 
**ExpiryTs** | **float64** |  | 
**BookiesWithOffers** | Pointer to **[]string** |  | [optional] 
**EquivalentBetsBookies** | Pointer to **[]string** |  | [optional] 
**CloseReason** | Pointer to **string** |  | [optional] 
**CustomerUsername** | Pointer to **string** |  | [optional] 
**CustomerCcy** | Pointer to **string** |  | [optional] 
**InvalidAccounts** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WantBookies** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]AccountItem**](AccountItem.md) |  | [optional] 

## Methods

### NewBetSlipData

`func NewBetSlipData(betslipId string, sport string, eventId string, betType string, expiryTs float64, ) *BetSlipData`

NewBetSlipData instantiates a new BetSlipData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetSlipDataWithDefaults

`func NewBetSlipDataWithDefaults() *BetSlipData`

NewBetSlipDataWithDefaults instantiates a new BetSlipData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetslipId

`func (o *BetSlipData) GetBetslipId() string`

GetBetslipId returns the BetslipId field if non-nil, zero value otherwise.

### GetBetslipIdOk

`func (o *BetSlipData) GetBetslipIdOk() (*string, bool)`

GetBetslipIdOk returns a tuple with the BetslipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetslipId

`func (o *BetSlipData) SetBetslipId(v string)`

SetBetslipId sets BetslipId field to given value.


### GetSport

`func (o *BetSlipData) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *BetSlipData) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *BetSlipData) SetSport(v string)`

SetSport sets Sport field to given value.


### GetEventId

`func (o *BetSlipData) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BetSlipData) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BetSlipData) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetBetType

`func (o *BetSlipData) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *BetSlipData) GetBetTypeOk() (*string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetType

`func (o *BetSlipData) SetBetType(v string)`

SetBetType sets BetType field to given value.


### GetBetTypeDescription

`func (o *BetSlipData) GetBetTypeDescription() string`

GetBetTypeDescription returns the BetTypeDescription field if non-nil, zero value otherwise.

### GetBetTypeDescriptionOk

`func (o *BetSlipData) GetBetTypeDescriptionOk() (*string, bool)`

GetBetTypeDescriptionOk returns a tuple with the BetTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetTypeDescription

`func (o *BetSlipData) SetBetTypeDescription(v string)`

SetBetTypeDescription sets BetTypeDescription field to given value.

### HasBetTypeDescription

`func (o *BetSlipData) HasBetTypeDescription() bool`

HasBetTypeDescription returns a boolean if a field has been set.

### GetBetTypeTemplate

`func (o *BetSlipData) GetBetTypeTemplate() string`

GetBetTypeTemplate returns the BetTypeTemplate field if non-nil, zero value otherwise.

### GetBetTypeTemplateOk

`func (o *BetSlipData) GetBetTypeTemplateOk() (*string, bool)`

GetBetTypeTemplateOk returns a tuple with the BetTypeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetTypeTemplate

`func (o *BetSlipData) SetBetTypeTemplate(v string)`

SetBetTypeTemplate sets BetTypeTemplate field to given value.

### HasBetTypeTemplate

`func (o *BetSlipData) HasBetTypeTemplate() bool`

HasBetTypeTemplate returns a boolean if a field has been set.

### GetEquivalentBets

`func (o *BetSlipData) GetEquivalentBets() bool`

GetEquivalentBets returns the EquivalentBets field if non-nil, zero value otherwise.

### GetEquivalentBetsOk

`func (o *BetSlipData) GetEquivalentBetsOk() (*bool, bool)`

GetEquivalentBetsOk returns a tuple with the EquivalentBets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquivalentBets

`func (o *BetSlipData) SetEquivalentBets(v bool)`

SetEquivalentBets sets EquivalentBets field to given value.

### HasEquivalentBets

`func (o *BetSlipData) HasEquivalentBets() bool`

HasEquivalentBets returns a boolean if a field has been set.

### GetMultipleAccounts

`func (o *BetSlipData) GetMultipleAccounts() bool`

GetMultipleAccounts returns the MultipleAccounts field if non-nil, zero value otherwise.

### GetMultipleAccountsOk

`func (o *BetSlipData) GetMultipleAccountsOk() (*bool, bool)`

GetMultipleAccountsOk returns a tuple with the MultipleAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAccounts

`func (o *BetSlipData) SetMultipleAccounts(v bool)`

SetMultipleAccounts sets MultipleAccounts field to given value.

### HasMultipleAccounts

`func (o *BetSlipData) HasMultipleAccounts() bool`

HasMultipleAccounts returns a boolean if a field has been set.

### GetIsOpen

`func (o *BetSlipData) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *BetSlipData) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *BetSlipData) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *BetSlipData) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### GetExpiryTs

`func (o *BetSlipData) GetExpiryTs() float64`

GetExpiryTs returns the ExpiryTs field if non-nil, zero value otherwise.

### GetExpiryTsOk

`func (o *BetSlipData) GetExpiryTsOk() (*float64, bool)`

GetExpiryTsOk returns a tuple with the ExpiryTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTs

`func (o *BetSlipData) SetExpiryTs(v float64)`

SetExpiryTs sets ExpiryTs field to given value.


### GetBookiesWithOffers

`func (o *BetSlipData) GetBookiesWithOffers() []string`

GetBookiesWithOffers returns the BookiesWithOffers field if non-nil, zero value otherwise.

### GetBookiesWithOffersOk

`func (o *BetSlipData) GetBookiesWithOffersOk() (*[]string, bool)`

GetBookiesWithOffersOk returns a tuple with the BookiesWithOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookiesWithOffers

`func (o *BetSlipData) SetBookiesWithOffers(v []string)`

SetBookiesWithOffers sets BookiesWithOffers field to given value.

### HasBookiesWithOffers

`func (o *BetSlipData) HasBookiesWithOffers() bool`

HasBookiesWithOffers returns a boolean if a field has been set.

### GetEquivalentBetsBookies

`func (o *BetSlipData) GetEquivalentBetsBookies() []string`

GetEquivalentBetsBookies returns the EquivalentBetsBookies field if non-nil, zero value otherwise.

### GetEquivalentBetsBookiesOk

`func (o *BetSlipData) GetEquivalentBetsBookiesOk() (*[]string, bool)`

GetEquivalentBetsBookiesOk returns a tuple with the EquivalentBetsBookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquivalentBetsBookies

`func (o *BetSlipData) SetEquivalentBetsBookies(v []string)`

SetEquivalentBetsBookies sets EquivalentBetsBookies field to given value.

### HasEquivalentBetsBookies

`func (o *BetSlipData) HasEquivalentBetsBookies() bool`

HasEquivalentBetsBookies returns a boolean if a field has been set.

### GetCloseReason

`func (o *BetSlipData) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *BetSlipData) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *BetSlipData) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *BetSlipData) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetCustomerUsername

`func (o *BetSlipData) GetCustomerUsername() string`

GetCustomerUsername returns the CustomerUsername field if non-nil, zero value otherwise.

### GetCustomerUsernameOk

`func (o *BetSlipData) GetCustomerUsernameOk() (*string, bool)`

GetCustomerUsernameOk returns a tuple with the CustomerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUsername

`func (o *BetSlipData) SetCustomerUsername(v string)`

SetCustomerUsername sets CustomerUsername field to given value.

### HasCustomerUsername

`func (o *BetSlipData) HasCustomerUsername() bool`

HasCustomerUsername returns a boolean if a field has been set.

### GetCustomerCcy

`func (o *BetSlipData) GetCustomerCcy() string`

GetCustomerCcy returns the CustomerCcy field if non-nil, zero value otherwise.

### GetCustomerCcyOk

`func (o *BetSlipData) GetCustomerCcyOk() (*string, bool)`

GetCustomerCcyOk returns a tuple with the CustomerCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerCcy

`func (o *BetSlipData) SetCustomerCcy(v string)`

SetCustomerCcy sets CustomerCcy field to given value.

### HasCustomerCcy

`func (o *BetSlipData) HasCustomerCcy() bool`

HasCustomerCcy returns a boolean if a field has been set.

### GetInvalidAccounts

`func (o *BetSlipData) GetInvalidAccounts() map[string]map[string]interface{}`

GetInvalidAccounts returns the InvalidAccounts field if non-nil, zero value otherwise.

### GetInvalidAccountsOk

`func (o *BetSlipData) GetInvalidAccountsOk() (*map[string]map[string]interface{}, bool)`

GetInvalidAccountsOk returns a tuple with the InvalidAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidAccounts

`func (o *BetSlipData) SetInvalidAccounts(v map[string]map[string]interface{})`

SetInvalidAccounts sets InvalidAccounts field to given value.

### HasInvalidAccounts

`func (o *BetSlipData) HasInvalidAccounts() bool`

HasInvalidAccounts returns a boolean if a field has been set.

### GetWantBookies

`func (o *BetSlipData) GetWantBookies() string`

GetWantBookies returns the WantBookies field if non-nil, zero value otherwise.

### GetWantBookiesOk

`func (o *BetSlipData) GetWantBookiesOk() (*string, bool)`

GetWantBookiesOk returns a tuple with the WantBookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantBookies

`func (o *BetSlipData) SetWantBookies(v string)`

SetWantBookies sets WantBookies field to given value.

### HasWantBookies

`func (o *BetSlipData) HasWantBookies() bool`

HasWantBookies returns a boolean if a field has been set.

### GetAccounts

`func (o *BetSlipData) GetAccounts() []AccountItem`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *BetSlipData) GetAccountsOk() (*[]AccountItem, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *BetSlipData) SetAccounts(v []AccountItem)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *BetSlipData) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


