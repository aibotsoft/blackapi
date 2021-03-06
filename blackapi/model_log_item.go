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
	"time"
)

// LogItem struct for LogItem
type LogItem struct {
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Message *string `json:"message,omitempty"`
	Category *string `json:"category,omitempty"`
}

// NewLogItem instantiates a new LogItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogItem() *LogItem {
	this := LogItem{}
	return &this
}

// NewLogItemWithDefaults instantiates a new LogItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogItemWithDefaults() *LogItem {
	this := LogItem{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *LogItem) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *LogItem) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *LogItem) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogItem) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogItem) SetMessage(v string) {
	o.Message = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *LogItem) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogItem) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *LogItem) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *LogItem) SetCategory(v string) {
	o.Category = &v
}

func (o LogItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	return json.Marshal(toSerialize)
}

type NullableLogItem struct {
	value *LogItem
	isSet bool
}

func (v NullableLogItem) Get() *LogItem {
	return v.value
}

func (v *NullableLogItem) Set(val *LogItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLogItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLogItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogItem(val *LogItem) *NullableLogItem {
	return &NullableLogItem{value: val, isSet: true}
}

func (v NullableLogItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


