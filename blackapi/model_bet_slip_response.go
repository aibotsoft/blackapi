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

// BetSlipResponse struct for BetSlipResponse
type BetSlipResponse struct {
	Data BetSlipData `json:"data"`
	Status string `json:"status"`
}

// NewBetSlipResponse instantiates a new BetSlipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetSlipResponse(data BetSlipData, status string, ) *BetSlipResponse {
	this := BetSlipResponse{}
	this.Data = data
	this.Status = status
	return &this
}

// NewBetSlipResponseWithDefaults instantiates a new BetSlipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetSlipResponseWithDefaults() *BetSlipResponse {
	this := BetSlipResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetSlipResponse) GetData() BetSlipData {
	if o == nil  {
		var ret BetSlipData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetSlipResponse) GetDataOk() (*BetSlipData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetSlipResponse) SetData(v BetSlipData) {
	o.Data = v
}

// GetStatus returns the Status field value
func (o *BetSlipResponse) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BetSlipResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BetSlipResponse) SetStatus(v string) {
	o.Status = v
}

func (o BetSlipResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBetSlipResponse struct {
	value *BetSlipResponse
	isSet bool
}

func (v NullableBetSlipResponse) Get() *BetSlipResponse {
	return v.value
}

func (v *NullableBetSlipResponse) Set(val *BetSlipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetSlipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetSlipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetSlipResponse(val *BetSlipResponse) *NullableBetSlipResponse {
	return &NullableBetSlipResponse{value: val, isSet: true}
}

func (v NullableBetSlipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetSlipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
