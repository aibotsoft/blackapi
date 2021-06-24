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

// DeleteBetSlipResponse struct for DeleteBetSlipResponse
type DeleteBetSlipResponse struct {
	Data *string `json:"data,omitempty"`
	Status string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DeleteBetSlipResponse DeleteBetSlipResponse

// NewDeleteBetSlipResponse instantiates a new DeleteBetSlipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteBetSlipResponse(status string, ) *DeleteBetSlipResponse {
	this := DeleteBetSlipResponse{}
	this.Status = status
	return &this
}

// NewDeleteBetSlipResponseWithDefaults instantiates a new DeleteBetSlipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteBetSlipResponseWithDefaults() *DeleteBetSlipResponse {
	this := DeleteBetSlipResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteBetSlipResponse) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteBetSlipResponse) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteBetSlipResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *DeleteBetSlipResponse) SetData(v string) {
	o.Data = &v
}

// GetStatus returns the Status field value
func (o *DeleteBetSlipResponse) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteBetSlipResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteBetSlipResponse) SetStatus(v string) {
	o.Status = v
}

func (o DeleteBetSlipResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeleteBetSlipResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDeleteBetSlipResponse := _DeleteBetSlipResponse{}

	if err = json.Unmarshal(bytes, &varDeleteBetSlipResponse); err == nil {
		*o = DeleteBetSlipResponse(varDeleteBetSlipResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteBetSlipResponse struct {
	value *DeleteBetSlipResponse
	isSet bool
}

func (v NullableDeleteBetSlipResponse) Get() *DeleteBetSlipResponse {
	return v.value
}

func (v *NullableDeleteBetSlipResponse) Set(val *DeleteBetSlipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteBetSlipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteBetSlipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteBetSlipResponse(val *DeleteBetSlipResponse) *NullableDeleteBetSlipResponse {
	return &NullableDeleteBetSlipResponse{value: val, isSet: true}
}

func (v NullableDeleteBetSlipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteBetSlipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


