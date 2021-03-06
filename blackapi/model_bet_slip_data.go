/*
 * Black API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blackapi

import (
	"encoding/json"
)

// BetSlipData struct for BetSlipData
type BetSlipData struct {
	BetslipId string `json:"betslip_id"`
	Sport string `json:"sport"`
	EventId string `json:"event_id"`
	BetType string `json:"bet_type"`
	BetTypeDescription *string `json:"bet_type_description,omitempty"`
	BetTypeTemplate *string `json:"bet_type_template,omitempty"`
	EquivalentBets *bool `json:"equivalent_bets,omitempty"`
	MultipleAccounts *bool `json:"multiple_accounts,omitempty"`
	IsOpen *bool `json:"is_open,omitempty"`
	ExpiryTs float64 `json:"expiry_ts"`
	BookiesWithOffers *[]string `json:"bookies_with_offers,omitempty"`
	WantBookies *[]string `json:"want_bookies,omitempty"`
	EquivalentBetsBookies *[]string `json:"equivalent_bets_bookies,omitempty"`
	CloseReason *string `json:"close_reason,omitempty"`
	CustomerUsername *string `json:"customer_username,omitempty"`
	CustomerCcy *string `json:"customer_ccy,omitempty"`
	InvalidAccounts *map[string]map[string]interface{} `json:"invalid_accounts,omitempty"`
	Accounts *[]AccountItem `json:"accounts,omitempty"`
}

// NewBetSlipData instantiates a new BetSlipData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetSlipData(betslipId string, sport string, eventId string, betType string, expiryTs float64, ) *BetSlipData {
	this := BetSlipData{}
	this.BetslipId = betslipId
	this.Sport = sport
	this.EventId = eventId
	this.BetType = betType
	this.ExpiryTs = expiryTs
	return &this
}

// NewBetSlipDataWithDefaults instantiates a new BetSlipData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetSlipDataWithDefaults() *BetSlipData {
	this := BetSlipData{}
	return &this
}

// GetBetslipId returns the BetslipId field value
func (o *BetSlipData) GetBetslipId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetslipId
}

// GetBetslipIdOk returns a tuple with the BetslipId field value
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetBetslipIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetslipId, true
}

// SetBetslipId sets field value
func (o *BetSlipData) SetBetslipId(v string) {
	o.BetslipId = v
}

// GetSport returns the Sport field value
func (o *BetSlipData) GetSport() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Sport
}

// GetSportOk returns a tuple with the Sport field value
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetSportOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sport, true
}

// SetSport sets field value
func (o *BetSlipData) SetSport(v string) {
	o.Sport = v
}

// GetEventId returns the EventId field value
func (o *BetSlipData) GetEventId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetEventIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *BetSlipData) SetEventId(v string) {
	o.EventId = v
}

// GetBetType returns the BetType field value
func (o *BetSlipData) GetBetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BetType
}

// GetBetTypeOk returns a tuple with the BetType field value
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetBetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetType, true
}

// SetBetType sets field value
func (o *BetSlipData) SetBetType(v string) {
	o.BetType = v
}

// GetBetTypeDescription returns the BetTypeDescription field value if set, zero value otherwise.
func (o *BetSlipData) GetBetTypeDescription() string {
	if o == nil || o.BetTypeDescription == nil {
		var ret string
		return ret
	}
	return *o.BetTypeDescription
}

// GetBetTypeDescriptionOk returns a tuple with the BetTypeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetBetTypeDescriptionOk() (*string, bool) {
	if o == nil || o.BetTypeDescription == nil {
		return nil, false
	}
	return o.BetTypeDescription, true
}

// HasBetTypeDescription returns a boolean if a field has been set.
func (o *BetSlipData) HasBetTypeDescription() bool {
	if o != nil && o.BetTypeDescription != nil {
		return true
	}

	return false
}

// SetBetTypeDescription gets a reference to the given string and assigns it to the BetTypeDescription field.
func (o *BetSlipData) SetBetTypeDescription(v string) {
	o.BetTypeDescription = &v
}

// GetBetTypeTemplate returns the BetTypeTemplate field value if set, zero value otherwise.
func (o *BetSlipData) GetBetTypeTemplate() string {
	if o == nil || o.BetTypeTemplate == nil {
		var ret string
		return ret
	}
	return *o.BetTypeTemplate
}

// GetBetTypeTemplateOk returns a tuple with the BetTypeTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetBetTypeTemplateOk() (*string, bool) {
	if o == nil || o.BetTypeTemplate == nil {
		return nil, false
	}
	return o.BetTypeTemplate, true
}

// HasBetTypeTemplate returns a boolean if a field has been set.
func (o *BetSlipData) HasBetTypeTemplate() bool {
	if o != nil && o.BetTypeTemplate != nil {
		return true
	}

	return false
}

// SetBetTypeTemplate gets a reference to the given string and assigns it to the BetTypeTemplate field.
func (o *BetSlipData) SetBetTypeTemplate(v string) {
	o.BetTypeTemplate = &v
}

// GetEquivalentBets returns the EquivalentBets field value if set, zero value otherwise.
func (o *BetSlipData) GetEquivalentBets() bool {
	if o == nil || o.EquivalentBets == nil {
		var ret bool
		return ret
	}
	return *o.EquivalentBets
}

// GetEquivalentBetsOk returns a tuple with the EquivalentBets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetEquivalentBetsOk() (*bool, bool) {
	if o == nil || o.EquivalentBets == nil {
		return nil, false
	}
	return o.EquivalentBets, true
}

// HasEquivalentBets returns a boolean if a field has been set.
func (o *BetSlipData) HasEquivalentBets() bool {
	if o != nil && o.EquivalentBets != nil {
		return true
	}

	return false
}

// SetEquivalentBets gets a reference to the given bool and assigns it to the EquivalentBets field.
func (o *BetSlipData) SetEquivalentBets(v bool) {
	o.EquivalentBets = &v
}

// GetMultipleAccounts returns the MultipleAccounts field value if set, zero value otherwise.
func (o *BetSlipData) GetMultipleAccounts() bool {
	if o == nil || o.MultipleAccounts == nil {
		var ret bool
		return ret
	}
	return *o.MultipleAccounts
}

// GetMultipleAccountsOk returns a tuple with the MultipleAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetMultipleAccountsOk() (*bool, bool) {
	if o == nil || o.MultipleAccounts == nil {
		return nil, false
	}
	return o.MultipleAccounts, true
}

// HasMultipleAccounts returns a boolean if a field has been set.
func (o *BetSlipData) HasMultipleAccounts() bool {
	if o != nil && o.MultipleAccounts != nil {
		return true
	}

	return false
}

// SetMultipleAccounts gets a reference to the given bool and assigns it to the MultipleAccounts field.
func (o *BetSlipData) SetMultipleAccounts(v bool) {
	o.MultipleAccounts = &v
}

// GetIsOpen returns the IsOpen field value if set, zero value otherwise.
func (o *BetSlipData) GetIsOpen() bool {
	if o == nil || o.IsOpen == nil {
		var ret bool
		return ret
	}
	return *o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetIsOpenOk() (*bool, bool) {
	if o == nil || o.IsOpen == nil {
		return nil, false
	}
	return o.IsOpen, true
}

// HasIsOpen returns a boolean if a field has been set.
func (o *BetSlipData) HasIsOpen() bool {
	if o != nil && o.IsOpen != nil {
		return true
	}

	return false
}

// SetIsOpen gets a reference to the given bool and assigns it to the IsOpen field.
func (o *BetSlipData) SetIsOpen(v bool) {
	o.IsOpen = &v
}

// GetExpiryTs returns the ExpiryTs field value
func (o *BetSlipData) GetExpiryTs() float64 {
	if o == nil  {
		var ret float64
		return ret
	}

	return o.ExpiryTs
}

// GetExpiryTsOk returns a tuple with the ExpiryTs field value
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetExpiryTsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpiryTs, true
}

// SetExpiryTs sets field value
func (o *BetSlipData) SetExpiryTs(v float64) {
	o.ExpiryTs = v
}

// GetBookiesWithOffers returns the BookiesWithOffers field value if set, zero value otherwise.
func (o *BetSlipData) GetBookiesWithOffers() []string {
	if o == nil || o.BookiesWithOffers == nil {
		var ret []string
		return ret
	}
	return *o.BookiesWithOffers
}

// GetBookiesWithOffersOk returns a tuple with the BookiesWithOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetBookiesWithOffersOk() (*[]string, bool) {
	if o == nil || o.BookiesWithOffers == nil {
		return nil, false
	}
	return o.BookiesWithOffers, true
}

// HasBookiesWithOffers returns a boolean if a field has been set.
func (o *BetSlipData) HasBookiesWithOffers() bool {
	if o != nil && o.BookiesWithOffers != nil {
		return true
	}

	return false
}

// SetBookiesWithOffers gets a reference to the given []string and assigns it to the BookiesWithOffers field.
func (o *BetSlipData) SetBookiesWithOffers(v []string) {
	o.BookiesWithOffers = &v
}

// GetWantBookies returns the WantBookies field value if set, zero value otherwise.
func (o *BetSlipData) GetWantBookies() []string {
	if o == nil || o.WantBookies == nil {
		var ret []string
		return ret
	}
	return *o.WantBookies
}

// GetWantBookiesOk returns a tuple with the WantBookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetWantBookiesOk() (*[]string, bool) {
	if o == nil || o.WantBookies == nil {
		return nil, false
	}
	return o.WantBookies, true
}

// HasWantBookies returns a boolean if a field has been set.
func (o *BetSlipData) HasWantBookies() bool {
	if o != nil && o.WantBookies != nil {
		return true
	}

	return false
}

// SetWantBookies gets a reference to the given []string and assigns it to the WantBookies field.
func (o *BetSlipData) SetWantBookies(v []string) {
	o.WantBookies = &v
}

// GetEquivalentBetsBookies returns the EquivalentBetsBookies field value if set, zero value otherwise.
func (o *BetSlipData) GetEquivalentBetsBookies() []string {
	if o == nil || o.EquivalentBetsBookies == nil {
		var ret []string
		return ret
	}
	return *o.EquivalentBetsBookies
}

// GetEquivalentBetsBookiesOk returns a tuple with the EquivalentBetsBookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetEquivalentBetsBookiesOk() (*[]string, bool) {
	if o == nil || o.EquivalentBetsBookies == nil {
		return nil, false
	}
	return o.EquivalentBetsBookies, true
}

// HasEquivalentBetsBookies returns a boolean if a field has been set.
func (o *BetSlipData) HasEquivalentBetsBookies() bool {
	if o != nil && o.EquivalentBetsBookies != nil {
		return true
	}

	return false
}

// SetEquivalentBetsBookies gets a reference to the given []string and assigns it to the EquivalentBetsBookies field.
func (o *BetSlipData) SetEquivalentBetsBookies(v []string) {
	o.EquivalentBetsBookies = &v
}

// GetCloseReason returns the CloseReason field value if set, zero value otherwise.
func (o *BetSlipData) GetCloseReason() string {
	if o == nil || o.CloseReason == nil {
		var ret string
		return ret
	}
	return *o.CloseReason
}

// GetCloseReasonOk returns a tuple with the CloseReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetCloseReasonOk() (*string, bool) {
	if o == nil || o.CloseReason == nil {
		return nil, false
	}
	return o.CloseReason, true
}

// HasCloseReason returns a boolean if a field has been set.
func (o *BetSlipData) HasCloseReason() bool {
	if o != nil && o.CloseReason != nil {
		return true
	}

	return false
}

// SetCloseReason gets a reference to the given string and assigns it to the CloseReason field.
func (o *BetSlipData) SetCloseReason(v string) {
	o.CloseReason = &v
}

// GetCustomerUsername returns the CustomerUsername field value if set, zero value otherwise.
func (o *BetSlipData) GetCustomerUsername() string {
	if o == nil || o.CustomerUsername == nil {
		var ret string
		return ret
	}
	return *o.CustomerUsername
}

// GetCustomerUsernameOk returns a tuple with the CustomerUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetCustomerUsernameOk() (*string, bool) {
	if o == nil || o.CustomerUsername == nil {
		return nil, false
	}
	return o.CustomerUsername, true
}

// HasCustomerUsername returns a boolean if a field has been set.
func (o *BetSlipData) HasCustomerUsername() bool {
	if o != nil && o.CustomerUsername != nil {
		return true
	}

	return false
}

// SetCustomerUsername gets a reference to the given string and assigns it to the CustomerUsername field.
func (o *BetSlipData) SetCustomerUsername(v string) {
	o.CustomerUsername = &v
}

// GetCustomerCcy returns the CustomerCcy field value if set, zero value otherwise.
func (o *BetSlipData) GetCustomerCcy() string {
	if o == nil || o.CustomerCcy == nil {
		var ret string
		return ret
	}
	return *o.CustomerCcy
}

// GetCustomerCcyOk returns a tuple with the CustomerCcy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetCustomerCcyOk() (*string, bool) {
	if o == nil || o.CustomerCcy == nil {
		return nil, false
	}
	return o.CustomerCcy, true
}

// HasCustomerCcy returns a boolean if a field has been set.
func (o *BetSlipData) HasCustomerCcy() bool {
	if o != nil && o.CustomerCcy != nil {
		return true
	}

	return false
}

// SetCustomerCcy gets a reference to the given string and assigns it to the CustomerCcy field.
func (o *BetSlipData) SetCustomerCcy(v string) {
	o.CustomerCcy = &v
}

// GetInvalidAccounts returns the InvalidAccounts field value if set, zero value otherwise.
func (o *BetSlipData) GetInvalidAccounts() map[string]map[string]interface{} {
	if o == nil || o.InvalidAccounts == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.InvalidAccounts
}

// GetInvalidAccountsOk returns a tuple with the InvalidAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetInvalidAccountsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.InvalidAccounts == nil {
		return nil, false
	}
	return o.InvalidAccounts, true
}

// HasInvalidAccounts returns a boolean if a field has been set.
func (o *BetSlipData) HasInvalidAccounts() bool {
	if o != nil && o.InvalidAccounts != nil {
		return true
	}

	return false
}

// SetInvalidAccounts gets a reference to the given map[string]map[string]interface{} and assigns it to the InvalidAccounts field.
func (o *BetSlipData) SetInvalidAccounts(v map[string]map[string]interface{}) {
	o.InvalidAccounts = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *BetSlipData) GetAccounts() []AccountItem {
	if o == nil || o.Accounts == nil {
		var ret []AccountItem
		return ret
	}
	return *o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetSlipData) GetAccountsOk() (*[]AccountItem, bool) {
	if o == nil || o.Accounts == nil {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *BetSlipData) HasAccounts() bool {
	if o != nil && o.Accounts != nil {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []AccountItem and assigns it to the Accounts field.
func (o *BetSlipData) SetAccounts(v []AccountItem) {
	o.Accounts = &v
}

func (o BetSlipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["betslip_id"] = o.BetslipId
	}
	if true {
		toSerialize["sport"] = o.Sport
	}
	if true {
		toSerialize["event_id"] = o.EventId
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
	if o.EquivalentBets != nil {
		toSerialize["equivalent_bets"] = o.EquivalentBets
	}
	if o.MultipleAccounts != nil {
		toSerialize["multiple_accounts"] = o.MultipleAccounts
	}
	if o.IsOpen != nil {
		toSerialize["is_open"] = o.IsOpen
	}
	if true {
		toSerialize["expiry_ts"] = o.ExpiryTs
	}
	if o.BookiesWithOffers != nil {
		toSerialize["bookies_with_offers"] = o.BookiesWithOffers
	}
	if o.WantBookies != nil {
		toSerialize["want_bookies"] = o.WantBookies
	}
	if o.EquivalentBetsBookies != nil {
		toSerialize["equivalent_bets_bookies"] = o.EquivalentBetsBookies
	}
	if o.CloseReason != nil {
		toSerialize["close_reason"] = o.CloseReason
	}
	if o.CustomerUsername != nil {
		toSerialize["customer_username"] = o.CustomerUsername
	}
	if o.CustomerCcy != nil {
		toSerialize["customer_ccy"] = o.CustomerCcy
	}
	if o.InvalidAccounts != nil {
		toSerialize["invalid_accounts"] = o.InvalidAccounts
	}
	if o.Accounts != nil {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableBetSlipData struct {
	value *BetSlipData
	isSet bool
}

func (v NullableBetSlipData) Get() *BetSlipData {
	return v.value
}

func (v *NullableBetSlipData) Set(val *BetSlipData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetSlipData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetSlipData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetSlipData(val *BetSlipData) *NullableBetSlipData {
	return &NullableBetSlipData{value: val, isSet: true}
}

func (v NullableBetSlipData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetSlipData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


