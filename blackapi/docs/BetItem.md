# BetItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetId** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Sport** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**Bookie** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**BetType** | Pointer to **string** |  | [optional] 
**CcyRate** | Pointer to **float64** |  | [optional] 
**Reconciled** | Pointer to **bool** |  | [optional] 
**BookieBetId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BetItemStatus**](BetItem_status.md) |  | [optional] 
**WantPrice** | Pointer to **float64** |  | [optional] 
**GotPrice** | Pointer to **float64** |  | [optional] 
**WantStake** | Pointer to **[]interface{}** |  | [optional] 
**GotStake** | Pointer to **[]interface{}** |  | [optional] 
**ProfitLoss** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewBetItem

`func NewBetItem() *BetItem`

NewBetItem instantiates a new BetItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetItemWithDefaults

`func NewBetItemWithDefaults() *BetItem`

NewBetItemWithDefaults instantiates a new BetItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetId

`func (o *BetItem) GetBetId() int64`

GetBetId returns the BetId field if non-nil, zero value otherwise.

### GetBetIdOk

`func (o *BetItem) GetBetIdOk() (*int64, bool)`

GetBetIdOk returns a tuple with the BetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetId

`func (o *BetItem) SetBetId(v int64)`

SetBetId sets BetId field to given value.

### HasBetId

`func (o *BetItem) HasBetId() bool`

HasBetId returns a boolean if a field has been set.

### GetOrderId

`func (o *BetItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BetItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BetItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BetItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSport

`func (o *BetItem) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *BetItem) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *BetItem) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *BetItem) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetEventId

`func (o *BetItem) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BetItem) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BetItem) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *BetItem) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetBookie

`func (o *BetItem) GetBookie() string`

GetBookie returns the Bookie field if non-nil, zero value otherwise.

### GetBookieOk

`func (o *BetItem) GetBookieOk() (*string, bool)`

GetBookieOk returns a tuple with the Bookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookie

`func (o *BetItem) SetBookie(v string)`

SetBookie sets Bookie field to given value.

### HasBookie

`func (o *BetItem) HasBookie() bool`

HasBookie returns a boolean if a field has been set.

### GetUsername

`func (o *BetItem) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BetItem) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BetItem) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BetItem) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetBetType

`func (o *BetItem) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *BetItem) GetBetTypeOk() (*string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetType

`func (o *BetItem) SetBetType(v string)`

SetBetType sets BetType field to given value.

### HasBetType

`func (o *BetItem) HasBetType() bool`

HasBetType returns a boolean if a field has been set.

### GetCcyRate

`func (o *BetItem) GetCcyRate() float64`

GetCcyRate returns the CcyRate field if non-nil, zero value otherwise.

### GetCcyRateOk

`func (o *BetItem) GetCcyRateOk() (*float64, bool)`

GetCcyRateOk returns a tuple with the CcyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcyRate

`func (o *BetItem) SetCcyRate(v float64)`

SetCcyRate sets CcyRate field to given value.

### HasCcyRate

`func (o *BetItem) HasCcyRate() bool`

HasCcyRate returns a boolean if a field has been set.

### GetReconciled

`func (o *BetItem) GetReconciled() bool`

GetReconciled returns the Reconciled field if non-nil, zero value otherwise.

### GetReconciledOk

`func (o *BetItem) GetReconciledOk() (*bool, bool)`

GetReconciledOk returns a tuple with the Reconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciled

`func (o *BetItem) SetReconciled(v bool)`

SetReconciled sets Reconciled field to given value.

### HasReconciled

`func (o *BetItem) HasReconciled() bool`

HasReconciled returns a boolean if a field has been set.

### GetBookieBetId

`func (o *BetItem) GetBookieBetId() string`

GetBookieBetId returns the BookieBetId field if non-nil, zero value otherwise.

### GetBookieBetIdOk

`func (o *BetItem) GetBookieBetIdOk() (*string, bool)`

GetBookieBetIdOk returns a tuple with the BookieBetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookieBetId

`func (o *BetItem) SetBookieBetId(v string)`

SetBookieBetId sets BookieBetId field to given value.

### HasBookieBetId

`func (o *BetItem) HasBookieBetId() bool`

HasBookieBetId returns a boolean if a field has been set.

### GetStatus

`func (o *BetItem) GetStatus() BetItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BetItem) GetStatusOk() (*BetItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BetItem) SetStatus(v BetItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BetItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWantPrice

`func (o *BetItem) GetWantPrice() float64`

GetWantPrice returns the WantPrice field if non-nil, zero value otherwise.

### GetWantPriceOk

`func (o *BetItem) GetWantPriceOk() (*float64, bool)`

GetWantPriceOk returns a tuple with the WantPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantPrice

`func (o *BetItem) SetWantPrice(v float64)`

SetWantPrice sets WantPrice field to given value.

### HasWantPrice

`func (o *BetItem) HasWantPrice() bool`

HasWantPrice returns a boolean if a field has been set.

### GetGotPrice

`func (o *BetItem) GetGotPrice() float64`

GetGotPrice returns the GotPrice field if non-nil, zero value otherwise.

### GetGotPriceOk

`func (o *BetItem) GetGotPriceOk() (*float64, bool)`

GetGotPriceOk returns a tuple with the GotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotPrice

`func (o *BetItem) SetGotPrice(v float64)`

SetGotPrice sets GotPrice field to given value.

### HasGotPrice

`func (o *BetItem) HasGotPrice() bool`

HasGotPrice returns a boolean if a field has been set.

### GetWantStake

`func (o *BetItem) GetWantStake() []interface{}`

GetWantStake returns the WantStake field if non-nil, zero value otherwise.

### GetWantStakeOk

`func (o *BetItem) GetWantStakeOk() (*[]interface{}, bool)`

GetWantStakeOk returns a tuple with the WantStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantStake

`func (o *BetItem) SetWantStake(v []interface{})`

SetWantStake sets WantStake field to given value.

### HasWantStake

`func (o *BetItem) HasWantStake() bool`

HasWantStake returns a boolean if a field has been set.

### GetGotStake

`func (o *BetItem) GetGotStake() []interface{}`

GetGotStake returns the GotStake field if non-nil, zero value otherwise.

### GetGotStakeOk

`func (o *BetItem) GetGotStakeOk() (*[]interface{}, bool)`

GetGotStakeOk returns a tuple with the GotStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGotStake

`func (o *BetItem) SetGotStake(v []interface{})`

SetGotStake sets GotStake field to given value.

### HasGotStake

`func (o *BetItem) HasGotStake() bool`

HasGotStake returns a boolean if a field has been set.

### GetProfitLoss

`func (o *BetItem) GetProfitLoss() []interface{}`

GetProfitLoss returns the ProfitLoss field if non-nil, zero value otherwise.

### GetProfitLossOk

`func (o *BetItem) GetProfitLossOk() (*[]interface{}, bool)`

GetProfitLossOk returns a tuple with the ProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitLoss

`func (o *BetItem) SetProfitLoss(v []interface{})`

SetProfitLoss sets ProfitLoss field to given value.

### HasProfitLoss

`func (o *BetItem) HasProfitLoss() bool`

HasProfitLoss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


