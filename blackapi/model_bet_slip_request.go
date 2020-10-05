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
)

// BetSlipRequest struct for BetSlipRequest
type BetSlipRequest struct {
	Sport string `json:"sport"`
	EventId string `json:"event_id"`
	BetType string `json:"bet_type"`
	EquivalentBets bool `json:"equivalent_bets"`
	MultipleAccounts bool `json:"multiple_accounts"`
}

// NewBetSlipRequest instantiates a new BetSlipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetSlipRequest(sport string, eventId string, betType string, equivalentBets bool, multipleAccounts bool, ) *BetSlipRequest {
	this := BetSlipRequest{}
	this.Sport = sport
	this.EventId = eventId
	this.BetType = betType
	this.EquivalentBets = equivalentBets
	this.MultipleAccounts = multipleAccounts
	return &this
}

// NewBetSlipRequestWithDefaults instantiates a new BetSlipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetSlipRequestWithDefaults() *BetSlipRequest {
	this := BetSlipRequest{}
	return &this
}

// GetSport returns the Sport field value
func (o *BetSlipRequest) GetSport() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Sport
}

// GetSportOk returns a tuple with the Sport field value
// and a boolean to check if the value has been set.
func (o *BetSlipRequest) GetSportOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sport, true
}

// SetSport sets field value
func (o *BetSlipRequest) SetSport(v string) {
	o.Sport = v
}

// GetEventId returns the EventId field value
func (o *BetSlipRequest) GetEventId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *BetSlipRequest) GetEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *BetSlipRequest) SetEventId(v string) {
	o.EventId = v
}

// GetBetType returns the BetType field value
func (o *BetSlipRequest) GetBetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value
// and a boolean to check if the value has been set.
func (o *BetSlipRequest) GetBetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetType, true
}

// SetBetType sets field value
func (o *BetSlipRequest) SetBetType(v string) {
	o.BetType = v
}

// GetEquivalentBets returns the EquivalentBets field value
func (o *BetSlipRequest) GetEquivalentBets() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.EquivalentBets
}

// GetEquivalentBetsOk returns a tuple with the EquivalentBets field value
// and a boolean to check if the value has been set.
func (o *BetSlipRequest) GetEquivalentBetsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EquivalentBets, true
}

// SetEquivalentBets sets field value
func (o *BetSlipRequest) SetEquivalentBets(v bool) {
	o.EquivalentBets = v
}

// GetMultipleAccounts returns the MultipleAccounts field value
func (o *BetSlipRequest) GetMultipleAccounts() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.MultipleAccounts
}

// GetMultipleAccountsOk returns a tuple with the MultipleAccounts field value
// and a boolean to check if the value has been set.
func (o *BetSlipRequest) GetMultipleAccountsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MultipleAccounts, true
}

// SetMultipleAccounts sets field value
func (o *BetSlipRequest) SetMultipleAccounts(v bool) {
	o.MultipleAccounts = v
}

func (o BetSlipRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sport"] = o.Sport
	}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if true {
		toSerialize["bet_type"] = o.BetType
	}
	if true {
		toSerialize["equivalent_bets"] = o.EquivalentBets
	}
	if true {
		toSerialize["multiple_accounts"] = o.MultipleAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableBetSlipRequest struct {
	value *BetSlipRequest
	isSet bool
}

func (v NullableBetSlipRequest) Get() *BetSlipRequest {
	return v.value
}

func (v *NullableBetSlipRequest) Set(val *BetSlipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetSlipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetSlipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetSlipRequest(val *BetSlipRequest) *NullableBetSlipRequest {
	return &NullableBetSlipRequest{value: val, isSet: true}
}

func (v NullableBetSlipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetSlipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
