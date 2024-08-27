/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the BulkWritablePowerFeedRequestSupply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritablePowerFeedRequestSupply{}

// BulkWritablePowerFeedRequestSupply struct for BulkWritablePowerFeedRequestSupply
type BulkWritablePowerFeedRequestSupply struct {
	Value *BulkWritablePowerFeedRequestSupplyValue `json:"value,omitempty"`
	Label *BulkWritablePowerFeedRequestSupplyLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritablePowerFeedRequestSupply BulkWritablePowerFeedRequestSupply

// NewBulkWritablePowerFeedRequestSupply instantiates a new BulkWritablePowerFeedRequestSupply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritablePowerFeedRequestSupply() *BulkWritablePowerFeedRequestSupply {
	this := BulkWritablePowerFeedRequestSupply{}
	return &this
}

// NewBulkWritablePowerFeedRequestSupplyWithDefaults instantiates a new BulkWritablePowerFeedRequestSupply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritablePowerFeedRequestSupplyWithDefaults() *BulkWritablePowerFeedRequestSupply {
	this := BulkWritablePowerFeedRequestSupply{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BulkWritablePowerFeedRequestSupply) GetValue() BulkWritablePowerFeedRequestSupplyValue {
	if o == nil || IsNil(o.Value) {
		var ret BulkWritablePowerFeedRequestSupplyValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerFeedRequestSupply) GetValueOk() (*BulkWritablePowerFeedRequestSupplyValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BulkWritablePowerFeedRequestSupply) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given BulkWritablePowerFeedRequestSupplyValue and assigns it to the Value field.
func (o *BulkWritablePowerFeedRequestSupply) SetValue(v BulkWritablePowerFeedRequestSupplyValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritablePowerFeedRequestSupply) GetLabel() BulkWritablePowerFeedRequestSupplyLabel {
	if o == nil || IsNil(o.Label) {
		var ret BulkWritablePowerFeedRequestSupplyLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerFeedRequestSupply) GetLabelOk() (*BulkWritablePowerFeedRequestSupplyLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritablePowerFeedRequestSupply) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given BulkWritablePowerFeedRequestSupplyLabel and assigns it to the Label field.
func (o *BulkWritablePowerFeedRequestSupply) SetLabel(v BulkWritablePowerFeedRequestSupplyLabel) {
	o.Label = &v
}

func (o BulkWritablePowerFeedRequestSupply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritablePowerFeedRequestSupply) ToMap() (map[string]interface{}, error) {
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

func (o *BulkWritablePowerFeedRequestSupply) UnmarshalJSON(data []byte) (err error) {
	varBulkWritablePowerFeedRequestSupply := _BulkWritablePowerFeedRequestSupply{}

	err = json.Unmarshal(data, &varBulkWritablePowerFeedRequestSupply)

	if err != nil {
		return err
	}

	*o = BulkWritablePowerFeedRequestSupply(varBulkWritablePowerFeedRequestSupply)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritablePowerFeedRequestSupply struct {
	value *BulkWritablePowerFeedRequestSupply
	isSet bool
}

func (v NullableBulkWritablePowerFeedRequestSupply) Get() *BulkWritablePowerFeedRequestSupply {
	return v.value
}

func (v *NullableBulkWritablePowerFeedRequestSupply) Set(val *BulkWritablePowerFeedRequestSupply) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritablePowerFeedRequestSupply) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritablePowerFeedRequestSupply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritablePowerFeedRequestSupply(val *BulkWritablePowerFeedRequestSupply) *NullableBulkWritablePowerFeedRequestSupply {
	return &NullableBulkWritablePowerFeedRequestSupply{value: val, isSet: true}
}

func (v NullableBulkWritablePowerFeedRequestSupply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritablePowerFeedRequestSupply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


