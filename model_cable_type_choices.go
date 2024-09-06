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

// CableTypeChoices the model 'CableTypeChoices'
type CableTypeChoices string

// List of CableTypeChoices
const (
	CABLETYPECHOICES_CAT3 CableTypeChoices = "cat3"
	CABLETYPECHOICES_CAT5 CableTypeChoices = "cat5"
	CABLETYPECHOICES_CAT5E CableTypeChoices = "cat5e"
	CABLETYPECHOICES_CAT6 CableTypeChoices = "cat6"
	CABLETYPECHOICES_CAT6A CableTypeChoices = "cat6a"
	CABLETYPECHOICES_CAT7 CableTypeChoices = "cat7"
	CABLETYPECHOICES_CAT7A CableTypeChoices = "cat7a"
	CABLETYPECHOICES_CAT8 CableTypeChoices = "cat8"
	CABLETYPECHOICES_DAC_ACTIVE CableTypeChoices = "dac-active"
	CABLETYPECHOICES_DAC_PASSIVE CableTypeChoices = "dac-passive"
	CABLETYPECHOICES_MRJ21_TRUNK CableTypeChoices = "mrj21-trunk"
	CABLETYPECHOICES_COAXIAL CableTypeChoices = "coaxial"
	CABLETYPECHOICES_MMF CableTypeChoices = "mmf"
	CABLETYPECHOICES_MMF_OM1 CableTypeChoices = "mmf-om1"
	CABLETYPECHOICES_MMF_OM2 CableTypeChoices = "mmf-om2"
	CABLETYPECHOICES_MMF_OM3 CableTypeChoices = "mmf-om3"
	CABLETYPECHOICES_MMF_OM4 CableTypeChoices = "mmf-om4"
	CABLETYPECHOICES_MMF_OM5 CableTypeChoices = "mmf-om5"
	CABLETYPECHOICES_SMF CableTypeChoices = "smf"
	CABLETYPECHOICES_SMF_OS1 CableTypeChoices = "smf-os1"
	CABLETYPECHOICES_SMF_OS2 CableTypeChoices = "smf-os2"
	CABLETYPECHOICES_AOC CableTypeChoices = "aoc"
	CABLETYPECHOICES_POWER CableTypeChoices = "power"
	CABLETYPECHOICES_OTHER CableTypeChoices = "other"
)

// All allowed values of CableTypeChoices enum
var AllowedCableTypeChoicesEnumValues = []CableTypeChoices{
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

func (v *CableTypeChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableTypeChoices(value)
	for _, existing := range AllowedCableTypeChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableTypeChoices", value)
}

// NewCableTypeChoicesFromValue returns a pointer to a valid CableTypeChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableTypeChoicesFromValue(v string) (*CableTypeChoices, error) {
	ev := CableTypeChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableTypeChoices: valid values are %v", v, AllowedCableTypeChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableTypeChoices) IsValid() bool {
	for _, existing := range AllowedCableTypeChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CableTypeChoices value
func (v CableTypeChoices) Ptr() *CableTypeChoices {
	return &v
}

type NullableCableTypeChoices struct {
	value *CableTypeChoices
	isSet bool
}

func (v NullableCableTypeChoices) Get() *CableTypeChoices {
	return v.value
}

func (v *NullableCableTypeChoices) Set(val *CableTypeChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableCableTypeChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableCableTypeChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableTypeChoices(val *CableTypeChoices) *NullableCableTypeChoices {
	return &NullableCableTypeChoices{value: val, isSet: true}
}

func (v NullableCableTypeChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableTypeChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

