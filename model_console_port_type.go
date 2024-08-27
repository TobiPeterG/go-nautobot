/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the ConsolePortType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsolePortType{}

// ConsolePortType struct for ConsolePortType
type ConsolePortType struct {
	Value *ConsolePortTypeValue `json:"value,omitempty"`
	Label *ConsolePortTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConsolePortType ConsolePortType

// NewConsolePortType instantiates a new ConsolePortType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsolePortType() *ConsolePortType {
	this := ConsolePortType{}
	return &this
}

// NewConsolePortTypeWithDefaults instantiates a new ConsolePortType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsolePortTypeWithDefaults() *ConsolePortType {
	this := ConsolePortType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConsolePortType) GetValue() ConsolePortTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret ConsolePortTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortType) GetValueOk() (*ConsolePortTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConsolePortType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given ConsolePortTypeValue and assigns it to the Value field.
func (o *ConsolePortType) SetValue(v ConsolePortTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConsolePortType) GetLabel() ConsolePortTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret ConsolePortTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsolePortType) GetLabelOk() (*ConsolePortTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConsolePortType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given ConsolePortTypeLabel and assigns it to the Label field.
func (o *ConsolePortType) SetLabel(v ConsolePortTypeLabel) {
	o.Label = &v
}

func (o ConsolePortType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsolePortType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConsolePortType) UnmarshalJSON(data []byte) (err error) {
	varConsolePortType := _ConsolePortType{}

	err = json.Unmarshal(data, &varConsolePortType)

	if err != nil {
		return err
	}

	*o = ConsolePortType(varConsolePortType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConsolePortType struct {
	value *ConsolePortType
	isSet bool
}

func (v NullableConsolePortType) Get() *ConsolePortType {
	return v.value
}

func (v *NullableConsolePortType) Set(val *ConsolePortType) {
	v.value = val
	v.isSet = true
}

func (v NullableConsolePortType) IsSet() bool {
	return v.isSet
}

func (v *NullableConsolePortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsolePortType(val *ConsolePortType) *NullableConsolePortType {
	return &NullableConsolePortType{value: val, isSet: true}
}

func (v NullableConsolePortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsolePortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


