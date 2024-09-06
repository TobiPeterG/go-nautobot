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

// FrontPortTypeValue the model 'FrontPortTypeValue'
type FrontPortTypeValue string

// List of FrontPort_type_value
const (
	FRONTPORTTYPEVALUE__8P8C FrontPortTypeValue = "8p8c"
	FRONTPORTTYPEVALUE__8P6C FrontPortTypeValue = "8p6c"
	FRONTPORTTYPEVALUE__8P4C FrontPortTypeValue = "8p4c"
	FRONTPORTTYPEVALUE__8P2C FrontPortTypeValue = "8p2c"
	FRONTPORTTYPEVALUE__6P6C FrontPortTypeValue = "6p6c"
	FRONTPORTTYPEVALUE__6P4C FrontPortTypeValue = "6p4c"
	FRONTPORTTYPEVALUE__6P2C FrontPortTypeValue = "6p2c"
	FRONTPORTTYPEVALUE__4P4C FrontPortTypeValue = "4p4c"
	FRONTPORTTYPEVALUE__4P2C FrontPortTypeValue = "4p2c"
	FRONTPORTTYPEVALUE_GG45 FrontPortTypeValue = "gg45"
	FRONTPORTTYPEVALUE_TERA_4P FrontPortTypeValue = "tera-4p"
	FRONTPORTTYPEVALUE_TERA_2P FrontPortTypeValue = "tera-2p"
	FRONTPORTTYPEVALUE_TERA_1P FrontPortTypeValue = "tera-1p"
	FRONTPORTTYPEVALUE__110_PUNCH FrontPortTypeValue = "110-punch"
	FRONTPORTTYPEVALUE_BNC FrontPortTypeValue = "bnc"
	FRONTPORTTYPEVALUE_F FrontPortTypeValue = "f"
	FRONTPORTTYPEVALUE_N FrontPortTypeValue = "n"
	FRONTPORTTYPEVALUE_MRJ21 FrontPortTypeValue = "mrj21"
	FRONTPORTTYPEVALUE_FC FrontPortTypeValue = "fc"
	FRONTPORTTYPEVALUE_LC FrontPortTypeValue = "lc"
	FRONTPORTTYPEVALUE_LC_PC FrontPortTypeValue = "lc-pc"
	FRONTPORTTYPEVALUE_LC_UPC FrontPortTypeValue = "lc-upc"
	FRONTPORTTYPEVALUE_LC_APC FrontPortTypeValue = "lc-apc"
	FRONTPORTTYPEVALUE_LSH FrontPortTypeValue = "lsh"
	FRONTPORTTYPEVALUE_LSH_PC FrontPortTypeValue = "lsh-pc"
	FRONTPORTTYPEVALUE_LSH_UPC FrontPortTypeValue = "lsh-upc"
	FRONTPORTTYPEVALUE_LSH_APC FrontPortTypeValue = "lsh-apc"
	FRONTPORTTYPEVALUE_LX5 FrontPortTypeValue = "lx5"
	FRONTPORTTYPEVALUE_LX5_PC FrontPortTypeValue = "lx5-pc"
	FRONTPORTTYPEVALUE_LX5_UPC FrontPortTypeValue = "lx5-upc"
	FRONTPORTTYPEVALUE_LX5_APC FrontPortTypeValue = "lx5-apc"
	FRONTPORTTYPEVALUE_MPO FrontPortTypeValue = "mpo"
	FRONTPORTTYPEVALUE_MTRJ FrontPortTypeValue = "mtrj"
	FRONTPORTTYPEVALUE_SC FrontPortTypeValue = "sc"
	FRONTPORTTYPEVALUE_SC_PC FrontPortTypeValue = "sc-pc"
	FRONTPORTTYPEVALUE_SC_UPC FrontPortTypeValue = "sc-upc"
	FRONTPORTTYPEVALUE_SC_APC FrontPortTypeValue = "sc-apc"
	FRONTPORTTYPEVALUE_ST FrontPortTypeValue = "st"
	FRONTPORTTYPEVALUE_CS FrontPortTypeValue = "cs"
	FRONTPORTTYPEVALUE_SN FrontPortTypeValue = "sn"
	FRONTPORTTYPEVALUE_SMA_905 FrontPortTypeValue = "sma-905"
	FRONTPORTTYPEVALUE_SMA_906 FrontPortTypeValue = "sma-906"
	FRONTPORTTYPEVALUE_URM_P2 FrontPortTypeValue = "urm-p2"
	FRONTPORTTYPEVALUE_URM_P4 FrontPortTypeValue = "urm-p4"
	FRONTPORTTYPEVALUE_URM_P8 FrontPortTypeValue = "urm-p8"
	FRONTPORTTYPEVALUE_SPLICE FrontPortTypeValue = "splice"
	FRONTPORTTYPEVALUE_OTHER FrontPortTypeValue = "other"
)

// All allowed values of FrontPortTypeValue enum
var AllowedFrontPortTypeValueEnumValues = []FrontPortTypeValue{
	"8p8c",
	"8p6c",
	"8p4c",
	"8p2c",
	"6p6c",
	"6p4c",
	"6p2c",
	"4p4c",
	"4p2c",
	"gg45",
	"tera-4p",
	"tera-2p",
	"tera-1p",
	"110-punch",
	"bnc",
	"f",
	"n",
	"mrj21",
	"fc",
	"lc",
	"lc-pc",
	"lc-upc",
	"lc-apc",
	"lsh",
	"lsh-pc",
	"lsh-upc",
	"lsh-apc",
	"lx5",
	"lx5-pc",
	"lx5-upc",
	"lx5-apc",
	"mpo",
	"mtrj",
	"sc",
	"sc-pc",
	"sc-upc",
	"sc-apc",
	"st",
	"cs",
	"sn",
	"sma-905",
	"sma-906",
	"urm-p2",
	"urm-p4",
	"urm-p8",
	"splice",
	"other",
}

func (v *FrontPortTypeValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FrontPortTypeValue(value)
	for _, existing := range AllowedFrontPortTypeValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FrontPortTypeValue", value)
}

// NewFrontPortTypeValueFromValue returns a pointer to a valid FrontPortTypeValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFrontPortTypeValueFromValue(v string) (*FrontPortTypeValue, error) {
	ev := FrontPortTypeValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FrontPortTypeValue: valid values are %v", v, AllowedFrontPortTypeValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FrontPortTypeValue) IsValid() bool {
	for _, existing := range AllowedFrontPortTypeValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FrontPort_type_value value
func (v FrontPortTypeValue) Ptr() *FrontPortTypeValue {
	return &v
}

type NullableFrontPortTypeValue struct {
	value *FrontPortTypeValue
	isSet bool
}

func (v NullableFrontPortTypeValue) Get() *FrontPortTypeValue {
	return v.value
}

func (v *NullableFrontPortTypeValue) Set(val *FrontPortTypeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortTypeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortTypeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortTypeValue(val *FrontPortTypeValue) *NullableFrontPortTypeValue {
	return &NullableFrontPortTypeValue{value: val, isSet: true}
}

func (v NullableFrontPortTypeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortTypeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

