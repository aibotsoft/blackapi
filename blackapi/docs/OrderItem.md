# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int64** |  | 
**OrderType** | **string** |  | 
**BetType** | **string** |  | 
**BetTypeDescription** | Pointer to **string** |  | [optional] 
**BetTypeTemplate** | Pointer to **string** |  | [optional] 
**Sport** | Pointer to **string** |  | [optional] 
**Placer** | Pointer to **string** |  | [optional] 
**WantPrice** | Pointer to **float64** |  | [optional] 
**WantStake** | Pointer to **[]interface{}** |  | [optional] 
**CcyRate** | Pointer to **float64** |  | [optional] 
**PlacementTime** | Pointer to **time.Time** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**Closed** | Pointer to **bool** |  | [optional] 
**CloseReason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserData** | Pointer to **string** |  | [optional] 
**TakeStartingPrice** | Pointer to **bool** |  | [optional] 
**KeepOpenIr** | Pointer to **bool** |  | [optional] 
**EventInfo** | Pointer to [**EventInfo**](EventInfo.md) |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Stake** | Pointer to **[]interface{}** |  | [optional] 
**ProfitLoss** | Pointer to **[]interface{}** |  | [optional] 
**BetBarValues** | Pointer to [**BetBar**](BetBar.md) |  | [optional] 
**Bets** | Pointer to [**[]BetItem**](BetItem.md) |  | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem(orderId int64, orderType string, betType string, ) *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderItem) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderItem) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderItem) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.


### GetOrderType

`func (o *OrderItem) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderItem) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderItem) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.


### GetBetType

`func (o *OrderItem) GetBetType() string`

GetBetType returns the BetType field if non-nil, zero value otherwise.

### GetBetTypeOk

`func (o *OrderItem) GetBetTypeOk() (*string, bool)`

GetBetTypeOk returns a tuple with the BetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetType

`func (o *OrderItem) SetBetType(v string)`

SetBetType sets BetType field to given value.


### GetBetTypeDescription

`func (o *OrderItem) GetBetTypeDescription() string`

GetBetTypeDescription returns the BetTypeDescription field if non-nil, zero value otherwise.

### GetBetTypeDescriptionOk

`func (o *OrderItem) GetBetTypeDescriptionOk() (*string, bool)`

GetBetTypeDescriptionOk returns a tuple with the BetTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetTypeDescription

`func (o *OrderItem) SetBetTypeDescription(v string)`

SetBetTypeDescription sets BetTypeDescription field to given value.

### HasBetTypeDescription

`func (o *OrderItem) HasBetTypeDescription() bool`

HasBetTypeDescription returns a boolean if a field has been set.

### GetBetTypeTemplate

`func (o *OrderItem) GetBetTypeTemplate() string`

GetBetTypeTemplate returns the BetTypeTemplate field if non-nil, zero value otherwise.

### GetBetTypeTemplateOk

`func (o *OrderItem) GetBetTypeTemplateOk() (*string, bool)`

GetBetTypeTemplateOk returns a tuple with the BetTypeTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetTypeTemplate

`func (o *OrderItem) SetBetTypeTemplate(v string)`

SetBetTypeTemplate sets BetTypeTemplate field to given value.

### HasBetTypeTemplate

`func (o *OrderItem) HasBetTypeTemplate() bool`

HasBetTypeTemplate returns a boolean if a field has been set.

### GetSport

`func (o *OrderItem) GetSport() string`

GetSport returns the Sport field if non-nil, zero value otherwise.

### GetSportOk

`func (o *OrderItem) GetSportOk() (*string, bool)`

GetSportOk returns a tuple with the Sport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSport

`func (o *OrderItem) SetSport(v string)`

SetSport sets Sport field to given value.

### HasSport

`func (o *OrderItem) HasSport() bool`

HasSport returns a boolean if a field has been set.

### GetPlacer

`func (o *OrderItem) GetPlacer() string`

GetPlacer returns the Placer field if non-nil, zero value otherwise.

### GetPlacerOk

`func (o *OrderItem) GetPlacerOk() (*string, bool)`

GetPlacerOk returns a tuple with the Placer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacer

`func (o *OrderItem) SetPlacer(v string)`

SetPlacer sets Placer field to given value.

### HasPlacer

`func (o *OrderItem) HasPlacer() bool`

HasPlacer returns a boolean if a field has been set.

### GetWantPrice

`func (o *OrderItem) GetWantPrice() float64`

GetWantPrice returns the WantPrice field if non-nil, zero value otherwise.

### GetWantPriceOk

`func (o *OrderItem) GetWantPriceOk() (*float64, bool)`

GetWantPriceOk returns a tuple with the WantPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantPrice

`func (o *OrderItem) SetWantPrice(v float64)`

SetWantPrice sets WantPrice field to given value.

### HasWantPrice

`func (o *OrderItem) HasWantPrice() bool`

HasWantPrice returns a boolean if a field has been set.

### GetWantStake

`func (o *OrderItem) GetWantStake() []interface{}`

GetWantStake returns the WantStake field if non-nil, zero value otherwise.

### GetWantStakeOk

`func (o *OrderItem) GetWantStakeOk() (*[]interface{}, bool)`

GetWantStakeOk returns a tuple with the WantStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantStake

`func (o *OrderItem) SetWantStake(v []interface{})`

SetWantStake sets WantStake field to given value.

### HasWantStake

`func (o *OrderItem) HasWantStake() bool`

HasWantStake returns a boolean if a field has been set.

### GetCcyRate

`func (o *OrderItem) GetCcyRate() float64`

GetCcyRate returns the CcyRate field if non-nil, zero value otherwise.

### GetCcyRateOk

`func (o *OrderItem) GetCcyRateOk() (*float64, bool)`

GetCcyRateOk returns a tuple with the CcyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcyRate

`func (o *OrderItem) SetCcyRate(v float64)`

SetCcyRate sets CcyRate field to given value.

### HasCcyRate

`func (o *OrderItem) HasCcyRate() bool`

HasCcyRate returns a boolean if a field has been set.

### GetPlacementTime

`func (o *OrderItem) GetPlacementTime() time.Time`

GetPlacementTime returns the PlacementTime field if non-nil, zero value otherwise.

### GetPlacementTimeOk

`func (o *OrderItem) GetPlacementTimeOk() (*time.Time, bool)`

GetPlacementTimeOk returns a tuple with the PlacementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementTime

`func (o *OrderItem) SetPlacementTime(v time.Time)`

SetPlacementTime sets PlacementTime field to given value.

### HasPlacementTime

`func (o *OrderItem) HasPlacementTime() bool`

HasPlacementTime returns a boolean if a field has been set.

### GetExpiryTime

`func (o *OrderItem) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *OrderItem) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *OrderItem) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *OrderItem) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetClosed

`func (o *OrderItem) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *OrderItem) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *OrderItem) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *OrderItem) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetCloseReason

`func (o *OrderItem) GetCloseReason() string`

GetCloseReason returns the CloseReason field if non-nil, zero value otherwise.

### GetCloseReasonOk

`func (o *OrderItem) GetCloseReasonOk() (*string, bool)`

GetCloseReasonOk returns a tuple with the CloseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseReason

`func (o *OrderItem) SetCloseReason(v string)`

SetCloseReason sets CloseReason field to given value.

### HasCloseReason

`func (o *OrderItem) HasCloseReason() bool`

HasCloseReason returns a boolean if a field has been set.

### GetStatus

`func (o *OrderItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserData

`func (o *OrderItem) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *OrderItem) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *OrderItem) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *OrderItem) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetTakeStartingPrice

`func (o *OrderItem) GetTakeStartingPrice() bool`

GetTakeStartingPrice returns the TakeStartingPrice field if non-nil, zero value otherwise.

### GetTakeStartingPriceOk

`func (o *OrderItem) GetTakeStartingPriceOk() (*bool, bool)`

GetTakeStartingPriceOk returns a tuple with the TakeStartingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeStartingPrice

`func (o *OrderItem) SetTakeStartingPrice(v bool)`

SetTakeStartingPrice sets TakeStartingPrice field to given value.

### HasTakeStartingPrice

`func (o *OrderItem) HasTakeStartingPrice() bool`

HasTakeStartingPrice returns a boolean if a field has been set.

### GetKeepOpenIr

`func (o *OrderItem) GetKeepOpenIr() bool`

GetKeepOpenIr returns the KeepOpenIr field if non-nil, zero value otherwise.

### GetKeepOpenIrOk

`func (o *OrderItem) GetKeepOpenIrOk() (*bool, bool)`

GetKeepOpenIrOk returns a tuple with the KeepOpenIr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOpenIr

`func (o *OrderItem) SetKeepOpenIr(v bool)`

SetKeepOpenIr sets KeepOpenIr field to given value.

### HasKeepOpenIr

`func (o *OrderItem) HasKeepOpenIr() bool`

HasKeepOpenIr returns a boolean if a field has been set.

### GetEventInfo

`func (o *OrderItem) GetEventInfo() EventInfo`

GetEventInfo returns the EventInfo field if non-nil, zero value otherwise.

### GetEventInfoOk

`func (o *OrderItem) GetEventInfoOk() (*EventInfo, bool)`

GetEventInfoOk returns a tuple with the EventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventInfo

`func (o *OrderItem) SetEventInfo(v EventInfo)`

SetEventInfo sets EventInfo field to given value.

### HasEventInfo

`func (o *OrderItem) HasEventInfo() bool`

HasEventInfo returns a boolean if a field has been set.

### GetPrice

`func (o *OrderItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItem) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStake

`func (o *OrderItem) GetStake() []interface{}`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *OrderItem) GetStakeOk() (*[]interface{}, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *OrderItem) SetStake(v []interface{})`

SetStake sets Stake field to given value.

### HasStake

`func (o *OrderItem) HasStake() bool`

HasStake returns a boolean if a field has been set.

### GetProfitLoss

`func (o *OrderItem) GetProfitLoss() []interface{}`

GetProfitLoss returns the ProfitLoss field if non-nil, zero value otherwise.

### GetProfitLossOk

`func (o *OrderItem) GetProfitLossOk() (*[]interface{}, bool)`

GetProfitLossOk returns a tuple with the ProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitLoss

`func (o *OrderItem) SetProfitLoss(v []interface{})`

SetProfitLoss sets ProfitLoss field to given value.

### HasProfitLoss

`func (o *OrderItem) HasProfitLoss() bool`

HasProfitLoss returns a boolean if a field has been set.

### GetBetBarValues

`func (o *OrderItem) GetBetBarValues() BetBar`

GetBetBarValues returns the BetBarValues field if non-nil, zero value otherwise.

### GetBetBarValuesOk

`func (o *OrderItem) GetBetBarValuesOk() (*BetBar, bool)`

GetBetBarValuesOk returns a tuple with the BetBarValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetBarValues

`func (o *OrderItem) SetBetBarValues(v BetBar)`

SetBetBarValues sets BetBarValues field to given value.

### HasBetBarValues

`func (o *OrderItem) HasBetBarValues() bool`

HasBetBarValues returns a boolean if a field has been set.

### GetBets

`func (o *OrderItem) GetBets() []BetItem`

GetBets returns the Bets field if non-nil, zero value otherwise.

### GetBetsOk

`func (o *OrderItem) GetBetsOk() (*[]BetItem, bool)`

GetBetsOk returns a tuple with the Bets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBets

`func (o *OrderItem) SetBets(v []BetItem)`

SetBets sets Bets field to given value.

### HasBets

`func (o *OrderItem) HasBets() bool`

HasBets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


