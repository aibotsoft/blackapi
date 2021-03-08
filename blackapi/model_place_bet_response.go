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

// PlaceBetResponse struct for PlaceBetResponse
type PlaceBetResponse struct {
	Status string `json:"status"`
	Data OrderItem `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _PlaceBetResponse PlaceBetResponse

// NewPlaceBetResponse instantiates a new PlaceBetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceBetResponse(status string, data OrderItem, ) *PlaceBetResponse {
	this := PlaceBetResponse{}
	this.Status = status
	this.Data = data
	return &this
}

// NewPlaceBetResponseWithDefaults instantiates a new PlaceBetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceBetResponseWithDefaults() *PlaceBetResponse {
	this := PlaceBetResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *PlaceBetResponse) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PlaceBetResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PlaceBetResponse) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value
func (o *PlaceBetResponse) GetData() OrderItem {
	if o == nil  {
		var ret OrderItem
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PlaceBetResponse) GetDataOk() (*OrderItem, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PlaceBetResponse) SetData(v OrderItem) {
	o.Data = v
}

func (o PlaceBetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlaceBetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPlaceBetResponse := _PlaceBetResponse{}

	if err = json.Unmarshal(bytes, &varPlaceBetResponse); err == nil {
		*o = PlaceBetResponse(varPlaceBetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceBetResponse struct {
	value *PlaceBetResponse
	isSet bool
}

func (v NullablePlaceBetResponse) Get() *PlaceBetResponse {
	return v.value
}

func (v *NullablePlaceBetResponse) Set(val *PlaceBetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceBetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceBetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceBetResponse(val *PlaceBetResponse) *NullablePlaceBetResponse {
	return &NullablePlaceBetResponse{value: val, isSet: true}
}

func (v NullablePlaceBetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceBetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

