/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// PowerFeedType the model 'PowerFeedType'
type PowerFeedType string

// List of PowerFeed_type
const (
	POWERFEEDTYPE_PRIMARY PowerFeedType = "primary"
	POWERFEEDTYPE_REDUNDANT PowerFeedType = "redundant"
)

// All allowed values of PowerFeedType enum
var AllowedPowerFeedTypeEnumValues = []PowerFeedType{
	"primary",
	"redundant",
}

func (v *PowerFeedType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedType(value)
	for _, existing := range AllowedPowerFeedTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedType", value)
}

// NewPowerFeedTypeFromValue returns a pointer to a valid PowerFeedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedTypeFromValue(v string) (*PowerFeedType, error) {
	ev := PowerFeedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedType: valid values are %v", v, AllowedPowerFeedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedType) IsValid() bool {
	for _, existing := range AllowedPowerFeedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_type value
func (v PowerFeedType) Ptr() *PowerFeedType {
	return &v
}

type NullablePowerFeedType struct {
	value *PowerFeedType
	isSet bool
}

func (v NullablePowerFeedType) Get() *PowerFeedType {
	return v.value
}

func (v *NullablePowerFeedType) Set(val *PowerFeedType) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedType) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedType(val *PowerFeedType) *NullablePowerFeedType {
	return &NullablePowerFeedType{value: val, isSet: true}
}

func (v NullablePowerFeedType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

