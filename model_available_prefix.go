/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the AvailablePrefix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailablePrefix{}

// AvailablePrefix Representation of a prefix which does not exist in the database.  Response serializer for a GET to /api/ipam/prefixes/<id>/available-prefixes/.
type AvailablePrefix struct {
	IpVersion int32 `json:"ip_version"`
	Prefix string `json:"prefix"`
	AdditionalProperties map[string]interface{}
}

type _AvailablePrefix AvailablePrefix

// NewAvailablePrefix instantiates a new AvailablePrefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePrefix(ipVersion int32, prefix string) *AvailablePrefix {
	this := AvailablePrefix{}
	this.IpVersion = ipVersion
	this.Prefix = prefix
	return &this
}

// NewAvailablePrefixWithDefaults instantiates a new AvailablePrefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePrefixWithDefaults() *AvailablePrefix {
	this := AvailablePrefix{}
	return &this
}

// GetIpVersion returns the IpVersion field value
func (o *AvailablePrefix) GetIpVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *AvailablePrefix) GetIpVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *AvailablePrefix) SetIpVersion(v int32) {
	o.IpVersion = v
}

// GetPrefix returns the Prefix field value
func (o *AvailablePrefix) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *AvailablePrefix) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *AvailablePrefix) SetPrefix(v string) {
	o.Prefix = v
}

func (o AvailablePrefix) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailablePrefix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip_version"] = o.IpVersion
	toSerialize["prefix"] = o.Prefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailablePrefix) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip_version",
		"prefix",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAvailablePrefix := _AvailablePrefix{}

	err = json.Unmarshal(data, &varAvailablePrefix)

	if err != nil {
		return err
	}

	*o = AvailablePrefix(varAvailablePrefix)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_version")
		delete(additionalProperties, "prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailablePrefix struct {
	value *AvailablePrefix
	isSet bool
}

func (v NullableAvailablePrefix) Get() *AvailablePrefix {
	return v.value
}

func (v *NullableAvailablePrefix) Set(val *AvailablePrefix) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePrefix) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePrefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePrefix(val *AvailablePrefix) *NullableAvailablePrefix {
	return &NullableAvailablePrefix{value: val, isSet: true}
}

func (v NullableAvailablePrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePrefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


