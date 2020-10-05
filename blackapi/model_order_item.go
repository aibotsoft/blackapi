/*
 * Black API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package blackapi

import (
	"encoding/json"
	"time"
)

// OrderItem struct for OrderItem
type OrderItem struct {
	OrderId int64 `json:"order_id"`
	OrderType string `json:"order_type"`
	BetType string `json:"bet_type"`
	BetTypeDescription *string `json:"bet_type_description,omitempty"`
	BetTypeTemplate *string `json:"bet_type_template,omitempty"`
	Sport *string `json:"sport,omitempty"`
	Placer *string `json:"placer,omitempty"`
	WantPrice *float64 `json:"want_price,omitempty"`
	WantStake *[]interface{} `json:"want_stake,omitempty"`
	CcyRate *float64 `json:"ccy_rate,omitempty"`
	PlacementTime *time.Time `json:"placement_time,omitempty"`
	ExpiryTime *time.Time `json:"expiry_time,omitempty"`
	Closed *bool `json:"closed,omitempty"`
	CloseReason *string `json:"close_reason,omitempty"`
	Status *string `json:"status,omitempty"`
	UserData *string `json:"user_data,omitempty"`
	TakeStartingPrice *bool `json:"take_starting_price,omitempty"`
	KeepOpenIr *bool `json:"keep_open_ir,omitempty"`
	EventInfo *EventInfo `json:"event_info,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Stake *[]interface{} `json:"stake,omitempty"`
	ProfitLoss *[]interface{} `json:"profit_loss,omitempty"`
	BetBarValues *BetBar `json:"bet_bar_values,omitempty"`
	Bets *[]BetItem `json:"bets,omitempty"`
}

// NewOrderItem instantiates a new OrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItem(orderId int64, orderType string, betType string, ) *OrderItem {
	this := OrderItem{}
	this.OrderId = orderId
	this.OrderType = orderType
	this.BetType = betType
	return &this
}

// NewOrderItemWithDefaults instantiates a new OrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemWithDefaults() *OrderItem {
	this := OrderItem{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *OrderItem) GetOrderId() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetOrderIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *OrderItem) SetOrderId(v int64) {
	o.OrderId = v
}

// GetOrderType returns the OrderType field value
func (o *OrderItem) GetOrderType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetOrderTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OrderType, true
}

// SetOrderType sets field value
func (o *OrderItem) SetOrderType(v string) {
	o.OrderType = v
}

// GetBetType returns the BetType field value
func (o *OrderItem) GetBetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetBetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetType, true
}

// SetBetType sets field value
func (o *OrderItem) SetBetType(v string) {
	o.BetType = v
}

// GetBetTypeDescription returns the BetTypeDescription field value if set, zero value otherwise.
func (o *OrderItem) GetBetTypeDescription() string {
	if o == nil || o.BetTypeDescription == nil {
		var ret string
		return ret
	}
	return *o.BetTypeDescription
}

// GetBetTypeDescriptionOk returns a tuple with the BetTypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetBetTypeDescriptionOk() (*string, bool) {
	if o == nil || o.BetTypeDescription == nil {
		return nil, false
	}
	return o.BetTypeDescription, true
}

// HasBetTypeDescription returns a boolean if a field has been set.
func (o *OrderItem) HasBetTypeDescription() bool {
	if o != nil && o.BetTypeDescription != nil {
		return true
	}

	return false
}

// SetBetTypeDescription gets a reference to the given string and assigns it to the BetTypeDescription field.
func (o *OrderItem) SetBetTypeDescription(v string) {
	o.BetTypeDescription = &v
}

// GetBetTypeTemplate returns the BetTypeTemplate field value if set, zero value otherwise.
func (o *OrderItem) GetBetTypeTemplate() string {
	if o == nil || o.BetTypeTemplate == nil {
		var ret string
		return ret
	}
	return *o.BetTypeTemplate
}

// GetBetTypeTemplateOk returns a tuple with the BetTypeTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetBetTypeTemplateOk() (*string, bool) {
	if o == nil || o.BetTypeTemplate == nil {
		return nil, false
	}
	return o.BetTypeTemplate, true
}

// HasBetTypeTemplate returns a boolean if a field has been set.
func (o *OrderItem) HasBetTypeTemplate() bool {
	if o != nil && o.BetTypeTemplate != nil {
		return true
	}

	return false
}

// SetBetTypeTemplate gets a reference to the given string and assigns it to the BetTypeTemplate field.
func (o *OrderItem) SetBetTypeTemplate(v string) {
	o.BetTypeTemplate = &v
}

// GetSport returns the Sport field value if set, zero value otherwise.
func (o *OrderItem) GetSport() string {
	if o == nil || o.Sport == nil {
		var ret string
		return ret
	}
	return *o.Sport
}

// GetSportOk returns a tuple with the Sport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetSportOk() (*string, bool) {
	if o == nil || o.Sport == nil {
		return nil, false
	}
	return o.Sport, true
}

// HasSport returns a boolean if a field has been set.
func (o *OrderItem) HasSport() bool {
	if o != nil && o.Sport != nil {
		return true
	}

	return false
}

// SetSport gets a reference to the given string and assigns it to the Sport field.
func (o *OrderItem) SetSport(v string) {
	o.Sport = &v
}

// GetPlacer returns the Placer field value if set, zero value otherwise.
func (o *OrderItem) GetPlacer() string {
	if o == nil || o.Placer == nil {
		var ret string
		return ret
	}
	return *o.Placer
}

// GetPlacerOk returns a tuple with the Placer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetPlacerOk() (*string, bool) {
	if o == nil || o.Placer == nil {
		return nil, false
	}
	return o.Placer, true
}

// HasPlacer returns a boolean if a field has been set.
func (o *OrderItem) HasPlacer() bool {
	if o != nil && o.Placer != nil {
		return true
	}

	return false
}

// SetPlacer gets a reference to the given string and assigns it to the Placer field.
func (o *OrderItem) SetPlacer(v string) {
	o.Placer = &v
}

// GetWantPrice returns the WantPrice field value if set, zero value otherwise.
func (o *OrderItem) GetWantPrice() float64 {
	if o == nil || o.WantPrice == nil {
		var ret float64
		return ret
	}
	return *o.WantPrice
}

// GetWantPriceOk returns a tuple with the WantPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetWantPriceOk() (*float64, bool) {
	if o == nil || o.WantPrice == nil {
		return nil, false
	}
	return o.WantPrice, true
}

// HasWantPrice returns a boolean if a field has been set.
func (o *OrderItem) HasWantPrice() bool {
	if o != nil && o.WantPrice != nil {
		return true
	}

	return false
}

// SetWantPrice gets a reference to the given float64 and assigns it to the WantPrice field.
func (o *OrderItem) SetWantPrice(v float64) {
	o.WantPrice = &v
}

// GetWantStake returns the WantStake field value if set, zero value otherwise.
func (o *OrderItem) GetWantStake() []interface{} {
	if o == nil || o.WantStake == nil {
		var ret []interface{}
		return ret
	}
	return *o.WantStake
}

// GetWantStakeOk returns a tuple with the WantStake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetWantStakeOk() (*[]interface{}, bool) {
	if o == nil || o.WantStake == nil {
		return nil, false
	}
	return o.WantStake, true
}

// HasWantStake returns a boolean if a field has been set.
func (o *OrderItem) HasWantStake() bool {
	if o != nil && o.WantStake != nil {
		return true
	}

	return false
}

// SetWantStake gets a reference to the given []interface{} and assigns it to the WantStake field.
func (o *OrderItem) SetWantStake(v []interface{}) {
	o.WantStake = &v
}

// GetCcyRate returns the CcyRate field value if set, zero value otherwise.
func (o *OrderItem) GetCcyRate() float64 {
	if o == nil || o.CcyRate == nil {
		var ret float64
		return ret
	}
	return *o.CcyRate
}

// GetCcyRateOk returns a tuple with the CcyRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetCcyRateOk() (*float64, bool) {
	if o == nil || o.CcyRate == nil {
		return nil, false
	}
	return o.CcyRate, true
}

// HasCcyRate returns a boolean if a field has been set.
func (o *OrderItem) HasCcyRate() bool {
	if o != nil && o.CcyRate != nil {
		return true
	}

	return false
}

// SetCcyRate gets a reference to the given float64 and assigns it to the CcyRate field.
func (o *OrderItem) SetCcyRate(v float64) {
	o.CcyRate = &v
}

// GetPlacementTime returns the PlacementTime field value if set, zero value otherwise.
func (o *OrderItem) GetPlacementTime() time.Time {
	if o == nil || o.PlacementTime == nil {
		var ret time.Time
		return ret
	}
	return *o.PlacementTime
}

// GetPlacementTimeOk returns a tuple with the PlacementTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetPlacementTimeOk() (*time.Time, bool) {
	if o == nil || o.PlacementTime == nil {
		return nil, false
	}
	return o.PlacementTime, true
}

// HasPlacementTime returns a boolean if a field has been set.
func (o *OrderItem) HasPlacementTime() bool {
	if o != nil && o.PlacementTime != nil {
		return true
	}

	return false
}

// SetPlacementTime gets a reference to the given time.Time and assigns it to the PlacementTime field.
func (o *OrderItem) SetPlacementTime(v time.Time) {
	o.PlacementTime = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *OrderItem) GetExpiryTime() time.Time {
	if o == nil || o.ExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *OrderItem) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *OrderItem) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *OrderItem) GetClosed() bool {
	if o == nil || o.Closed == nil {
		var ret bool
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetClosedOk() (*bool, bool) {
	if o == nil || o.Closed == nil {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *OrderItem) HasClosed() bool {
	if o != nil && o.Closed != nil {
		return true
	}

	return false
}

// SetClosed gets a reference to the given bool and assigns it to the Closed field.
func (o *OrderItem) SetClosed(v bool) {
	o.Closed = &v
}

// GetCloseReason returns the CloseReason field value if set, zero value otherwise.
func (o *OrderItem) GetCloseReason() string {
	if o == nil || o.CloseReason == nil {
		var ret string
		return ret
	}
	return *o.CloseReason
}

// GetCloseReasonOk returns a tuple with the CloseReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetCloseReasonOk() (*string, bool) {
	if o == nil || o.CloseReason == nil {
		return nil, false
	}
	return o.CloseReason, true
}

// HasCloseReason returns a boolean if a field has been set.
func (o *OrderItem) HasCloseReason() bool {
	if o != nil && o.CloseReason != nil {
		return true
	}

	return false
}

// SetCloseReason gets a reference to the given string and assigns it to the CloseReason field.
func (o *OrderItem) SetCloseReason(v string) {
	o.CloseReason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderItem) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderItem) SetStatus(v string) {
	o.Status = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *OrderItem) GetUserData() string {
	if o == nil || o.UserData == nil {
		var ret string
		return ret
	}
	return *o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetUserDataOk() (*string, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *OrderItem) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *OrderItem) SetUserData(v string) {
	o.UserData = &v
}

// GetTakeStartingPrice returns the TakeStartingPrice field value if set, zero value otherwise.
func (o *OrderItem) GetTakeStartingPrice() bool {
	if o == nil || o.TakeStartingPrice == nil {
		var ret bool
		return ret
	}
	return *o.TakeStartingPrice
}

// GetTakeStartingPriceOk returns a tuple with the TakeStartingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetTakeStartingPriceOk() (*bool, bool) {
	if o == nil || o.TakeStartingPrice == nil {
		return nil, false
	}
	return o.TakeStartingPrice, true
}

// HasTakeStartingPrice returns a boolean if a field has been set.
func (o *OrderItem) HasTakeStartingPrice() bool {
	if o != nil && o.TakeStartingPrice != nil {
		return true
	}

	return false
}

// SetTakeStartingPrice gets a reference to the given bool and assigns it to the TakeStartingPrice field.
func (o *OrderItem) SetTakeStartingPrice(v bool) {
	o.TakeStartingPrice = &v
}

// GetKeepOpenIr returns the KeepOpenIr field value if set, zero value otherwise.
func (o *OrderItem) GetKeepOpenIr() bool {
	if o == nil || o.KeepOpenIr == nil {
		var ret bool
		return ret
	}
	return *o.KeepOpenIr
}

// GetKeepOpenIrOk returns a tuple with the KeepOpenIr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetKeepOpenIrOk() (*bool, bool) {
	if o == nil || o.KeepOpenIr == nil {
		return nil, false
	}
	return o.KeepOpenIr, true
}

// HasKeepOpenIr returns a boolean if a field has been set.
func (o *OrderItem) HasKeepOpenIr() bool {
	if o != nil && o.KeepOpenIr != nil {
		return true
	}

	return false
}

// SetKeepOpenIr gets a reference to the given bool and assigns it to the KeepOpenIr field.
func (o *OrderItem) SetKeepOpenIr(v bool) {
	o.KeepOpenIr = &v
}

// GetEventInfo returns the EventInfo field value if set, zero value otherwise.
func (o *OrderItem) GetEventInfo() EventInfo {
	if o == nil || o.EventInfo == nil {
		var ret EventInfo
		return ret
	}
	return *o.EventInfo
}

// GetEventInfoOk returns a tuple with the EventInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetEventInfoOk() (*EventInfo, bool) {
	if o == nil || o.EventInfo == nil {
		return nil, false
	}
	return o.EventInfo, true
}

// HasEventInfo returns a boolean if a field has been set.
func (o *OrderItem) HasEventInfo() bool {
	if o != nil && o.EventInfo != nil {
		return true
	}

	return false
}

// SetEventInfo gets a reference to the given EventInfo and assigns it to the EventInfo field.
func (o *OrderItem) SetEventInfo(v EventInfo) {
	o.EventInfo = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *OrderItem) GetPrice() float64 {
	if o == nil || o.Price == nil {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetPriceOk() (*float64, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *OrderItem) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *OrderItem) SetPrice(v float64) {
	o.Price = &v
}

// GetStake returns the Stake field value if set, zero value otherwise.
func (o *OrderItem) GetStake() []interface{} {
	if o == nil || o.Stake == nil {
		var ret []interface{}
		return ret
	}
	return *o.Stake
}

// GetStakeOk returns a tuple with the Stake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetStakeOk() (*[]interface{}, bool) {
	if o == nil || o.Stake == nil {
		return nil, false
	}
	return o.Stake, true
}

// HasStake returns a boolean if a field has been set.
func (o *OrderItem) HasStake() bool {
	if o != nil && o.Stake != nil {
		return true
	}

	return false
}

// SetStake gets a reference to the given []interface{} and assigns it to the Stake field.
func (o *OrderItem) SetStake(v []interface{}) {
	o.Stake = &v
}

// GetProfitLoss returns the ProfitLoss field value if set, zero value otherwise.
func (o *OrderItem) GetProfitLoss() []interface{} {
	if o == nil || o.ProfitLoss == nil {
		var ret []interface{}
		return ret
	}
	return *o.ProfitLoss
}

// GetProfitLossOk returns a tuple with the ProfitLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetProfitLossOk() (*[]interface{}, bool) {
	if o == nil || o.ProfitLoss == nil {
		return nil, false
	}
	return o.ProfitLoss, true
}

// HasProfitLoss returns a boolean if a field has been set.
func (o *OrderItem) HasProfitLoss() bool {
	if o != nil && o.ProfitLoss != nil {
		return true
	}

	return false
}

// SetProfitLoss gets a reference to the given []interface{} and assigns it to the ProfitLoss field.
func (o *OrderItem) SetProfitLoss(v []interface{}) {
	o.ProfitLoss = &v
}

// GetBetBarValues returns the BetBarValues field value if set, zero value otherwise.
func (o *OrderItem) GetBetBarValues() BetBar {
	if o == nil || o.BetBarValues == nil {
		var ret BetBar
		return ret
	}
	return *o.BetBarValues
}

// GetBetBarValuesOk returns a tuple with the BetBarValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetBetBarValuesOk() (*BetBar, bool) {
	if o == nil || o.BetBarValues == nil {
		return nil, false
	}
	return o.BetBarValues, true
}

// HasBetBarValues returns a boolean if a field has been set.
func (o *OrderItem) HasBetBarValues() bool {
	if o != nil && o.BetBarValues != nil {
		return true
	}

	return false
}

// SetBetBarValues gets a reference to the given BetBar and assigns it to the BetBarValues field.
func (o *OrderItem) SetBetBarValues(v BetBar) {
	o.BetBarValues = &v
}

// GetBets returns the Bets field value if set, zero value otherwise.
func (o *OrderItem) GetBets() []BetItem {
	if o == nil || o.Bets == nil {
		var ret []BetItem
		return ret
	}
	return *o.Bets
}

// GetBetsOk returns a tuple with the Bets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetBetsOk() (*[]BetItem, bool) {
	if o == nil || o.Bets == nil {
		return nil, false
	}
	return o.Bets, true
}

// HasBets returns a boolean if a field has been set.
func (o *OrderItem) HasBets() bool {
	if o != nil && o.Bets != nil {
		return true
	}

	return false
}

// SetBets gets a reference to the given []BetItem and assigns it to the Bets field.
func (o *OrderItem) SetBets(v []BetItem) {
	o.Bets = &v
}

func (o OrderItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if true {
		toSerialize["order_type"] = o.OrderType
	}
	if true {
		toSerialize["bet_type"] = o.BetType
	}
	if o.BetTypeDescription != nil {
		toSerialize["bet_type_description"] = o.BetTypeDescription
	}
	if o.BetTypeTemplate != nil {
		toSerialize["bet_type_template"] = o.BetTypeTemplate
	}
	if o.Sport != nil {
		toSerialize["sport"] = o.Sport
	}
	if o.Placer != nil {
		toSerialize["placer"] = o.Placer
	}
	if o.WantPrice != nil {
		toSerialize["want_price"] = o.WantPrice
	}
	if o.WantStake != nil {
		toSerialize["want_stake"] = o.WantStake
	}
	if o.CcyRate != nil {
		toSerialize["ccy_rate"] = o.CcyRate
	}
	if o.PlacementTime != nil {
		toSerialize["placement_time"] = o.PlacementTime
	}
	if o.ExpiryTime != nil {
		toSerialize["expiry_time"] = o.ExpiryTime
	}
	if o.Closed != nil {
		toSerialize["closed"] = o.Closed
	}
	if o.CloseReason != nil {
		toSerialize["close_reason"] = o.CloseReason
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserData != nil {
		toSerialize["user_data"] = o.UserData
	}
	if o.TakeStartingPrice != nil {
		toSerialize["take_starting_price"] = o.TakeStartingPrice
	}
	if o.KeepOpenIr != nil {
		toSerialize["keep_open_ir"] = o.KeepOpenIr
	}
	if o.EventInfo != nil {
		toSerialize["event_info"] = o.EventInfo
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Stake != nil {
		toSerialize["stake"] = o.Stake
	}
	if o.ProfitLoss != nil {
		toSerialize["profit_loss"] = o.ProfitLoss
	}
	if o.BetBarValues != nil {
		toSerialize["bet_bar_values"] = o.BetBarValues
	}
	if o.Bets != nil {
		toSerialize["bets"] = o.Bets
	}
	return json.Marshal(toSerialize)
}

type NullableOrderItem struct {
	value *OrderItem
	isSet bool
}

func (v NullableOrderItem) Get() *OrderItem {
	return v.value
}

func (v *NullableOrderItem) Set(val *OrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItem(val *OrderItem) *NullableOrderItem {
	return &NullableOrderItem{value: val, isSet: true}
}

func (v NullableOrderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
