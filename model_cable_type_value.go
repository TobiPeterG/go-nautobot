/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// CableTypeValue the model 'CableTypeValue'
type CableTypeValue string

// List of Cable_type_value
const (
	CABLETYPEVALUE_CAT3 CableTypeValue = "cat3"
	CABLETYPEVALUE_CAT5 CableTypeValue = "cat5"
	CABLETYPEVALUE_CAT5E CableTypeValue = "cat5e"
	CABLETYPEVALUE_CAT6 CableTypeValue = "cat6"
	CABLETYPEVALUE_CAT6A CableTypeValue = "cat6a"
	CABLETYPEVALUE_CAT7 CableTypeValue = "cat7"
	CABLETYPEVALUE_CAT7A CableTypeValue = "cat7a"
	CABLETYPEVALUE_CAT8 CableTypeValue = "cat8"
	CABLETYPEVALUE_DAC_ACTIVE CableTypeValue = "dac-active"
	CABLETYPEVALUE_DAC_PASSIVE CableTypeValue = "dac-passive"
	CABLETYPEVALUE_MRJ21_TRUNK CableTypeValue = "mrj21-trunk"
	CABLETYPEVALUE_COAXIAL CableTypeValue = "coaxial"
	CABLETYPEVALUE_MMF CableTypeValue = "mmf"
	CABLETYPEVALUE_MMF_OM1 CableTypeValue = "mmf-om1"
	CABLETYPEVALUE_MMF_OM2 CableTypeValue = "mmf-om2"
	CABLETYPEVALUE_MMF_OM3 CableTypeValue = "mmf-om3"
	CABLETYPEVALUE_MMF_OM4 CableTypeValue = "mmf-om4"
	CABLETYPEVALUE_MMF_OM5 CableTypeValue = "mmf-om5"
	CABLETYPEVALUE_SMF CableTypeValue = "smf"
	CABLETYPEVALUE_SMF_OS1 CableTypeValue = "smf-os1"
	CABLETYPEVALUE_SMF_OS2 CableTypeValue = "smf-os2"
	CABLETYPEVALUE_AOC CableTypeValue = "aoc"
	CABLETYPEVALUE_POWER CableTypeValue = "power"
	CABLETYPEVALUE_OTHER CableTypeValue = "other"
)

// All allowed values of CableTypeValue enum
var AllowedCableTypeValueEnumValues = []CableTypeValue{
	"cat3",
	"cat5",
	"cat5e",
	"cat6",
	"cat6a",
	"cat7",
	"cat7a",
	"cat8",
	"dac-active",
	"dac-passive",
	"mrj21-trunk",
	"coaxial",
	"mmf",
	"mmf-om1",
	"mmf-om2",
	"mmf-om3",
	"mmf-om4",
	"mmf-om5",
	"smf",
	"smf-os1",
	"smf-os2",
	"aoc",
	"power",
	"other",
}

func (v *CableTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableTypeValue(value)
	for _, existing := range AllowedCableTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableTypeValue", value)
}

// NewCableTypeValueFromValue returns a pointer to a valid CableTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableTypeValueFromValue(v string) (*CableTypeValue, error) {
	ev := CableTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableTypeValue: valid values are %v", v, AllowedCableTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableTypeValue) IsValid() bool {
	for _, existing := range AllowedCableTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cable_type_value value
func (v CableTypeValue) Ptr() *CableTypeValue {
	return &v
}

type NullableCableTypeValue struct {
	value *CableTypeValue
	isSet bool
}

func (v NullableCableTypeValue) Get() *CableTypeValue {
	return v.value
}

func (v *NullableCableTypeValue) Set(val *CableTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCableTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCableTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableTypeValue(val *CableTypeValue) *NullableCableTypeValue {
	return &NullableCableTypeValue{value: val, isSet: true}
}

func (v NullableCableTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

