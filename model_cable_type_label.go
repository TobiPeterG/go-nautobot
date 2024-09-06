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

// CableTypeLabel the model 'CableTypeLabel'
type CableTypeLabel string

// List of Cable_type_label
const (
	CABLETYPELABEL_CAT3 CableTypeLabel = "CAT3"
	CABLETYPELABEL_CAT5 CableTypeLabel = "CAT5"
	CABLETYPELABEL_CAT5E CableTypeLabel = "CAT5e"
	CABLETYPELABEL_CAT6 CableTypeLabel = "CAT6"
	CABLETYPELABEL_CAT6A CableTypeLabel = "CAT6a"
	CABLETYPELABEL_CAT7 CableTypeLabel = "CAT7"
	CABLETYPELABEL_CAT7A CableTypeLabel = "CAT7a"
	CABLETYPELABEL_CAT8 CableTypeLabel = "CAT8"
	CABLETYPELABEL_DIRECT_ATTACH_COPPER__ACTIVE CableTypeLabel = "Direct Attach Copper (Active)"
	CABLETYPELABEL_DIRECT_ATTACH_COPPER__PASSIVE CableTypeLabel = "Direct Attach Copper (Passive)"
	CABLETYPELABEL_MRJ21_TRUNK CableTypeLabel = "MRJ21 Trunk"
	CABLETYPELABEL_COAXIAL CableTypeLabel = "Coaxial"
	CABLETYPELABEL_MULTIMODE_FIBER CableTypeLabel = "Multimode Fiber"
	CABLETYPELABEL_MULTIMODE_FIBER__OM1 CableTypeLabel = "Multimode Fiber (OM1)"
	CABLETYPELABEL_MULTIMODE_FIBER__OM2 CableTypeLabel = "Multimode Fiber (OM2)"
	CABLETYPELABEL_MULTIMODE_FIBER__OM3 CableTypeLabel = "Multimode Fiber (OM3)"
	CABLETYPELABEL_MULTIMODE_FIBER__OM4 CableTypeLabel = "Multimode Fiber (OM4)"
	CABLETYPELABEL_MULTIMODE_FIBER__OM5 CableTypeLabel = "Multimode Fiber (OM5)"
	CABLETYPELABEL_SINGLEMODE_FIBER CableTypeLabel = "Singlemode Fiber"
	CABLETYPELABEL_SINGLEMODE_FIBER__OS1 CableTypeLabel = "Singlemode Fiber (OS1)"
	CABLETYPELABEL_SINGLEMODE_FIBER__OS2 CableTypeLabel = "Singlemode Fiber (OS2)"
	CABLETYPELABEL_ACTIVE_OPTICAL_CABLING__AOC CableTypeLabel = "Active Optical Cabling (AOC)"
	CABLETYPELABEL_POWER CableTypeLabel = "Power"
	CABLETYPELABEL_OTHER CableTypeLabel = "Other"
)

// All allowed values of CableTypeLabel enum
var AllowedCableTypeLabelEnumValues = []CableTypeLabel{
	"CAT3",
	"CAT5",
	"CAT5e",
	"CAT6",
	"CAT6a",
	"CAT7",
	"CAT7a",
	"CAT8",
	"Direct Attach Copper (Active)",
	"Direct Attach Copper (Passive)",
	"MRJ21 Trunk",
	"Coaxial",
	"Multimode Fiber",
	"Multimode Fiber (OM1)",
	"Multimode Fiber (OM2)",
	"Multimode Fiber (OM3)",
	"Multimode Fiber (OM4)",
	"Multimode Fiber (OM5)",
	"Singlemode Fiber",
	"Singlemode Fiber (OS1)",
	"Singlemode Fiber (OS2)",
	"Active Optical Cabling (AOC)",
	"Power",
	"Other",
}

func (v *CableTypeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableTypeLabel(value)
	for _, existing := range AllowedCableTypeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableTypeLabel", value)
}

// NewCableTypeLabelFromValue returns a pointer to a valid CableTypeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableTypeLabelFromValue(v string) (*CableTypeLabel, error) {
	ev := CableTypeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableTypeLabel: valid values are %v", v, AllowedCableTypeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableTypeLabel) IsValid() bool {
	for _, existing := range AllowedCableTypeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cable_type_label value
func (v CableTypeLabel) Ptr() *CableTypeLabel {
	return &v
}

type NullableCableTypeLabel struct {
	value *CableTypeLabel
	isSet bool
}

func (v NullableCableTypeLabel) Get() *CableTypeLabel {
	return v.value
}

func (v *NullableCableTypeLabel) Set(val *CableTypeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCableTypeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCableTypeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableTypeLabel(val *CableTypeLabel) *NullableCableTypeLabel {
	return &NullableCableTypeLabel{value: val, isSet: true}
}

func (v NullableCableTypeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableTypeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

