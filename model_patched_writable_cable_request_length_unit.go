/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// PatchedWritableCableRequestLengthUnit - struct for PatchedWritableCableRequestLengthUnit
type PatchedWritableCableRequestLengthUnit struct {
	BlankEnum *BlankEnum
	LengthUnitEnum *LengthUnitEnum
}

// BlankEnumAsPatchedWritableCableRequestLengthUnit is a convenience function that returns BlankEnum wrapped in PatchedWritableCableRequestLengthUnit
func BlankEnumAsPatchedWritableCableRequestLengthUnit(v *BlankEnum) PatchedWritableCableRequestLengthUnit {
	return PatchedWritableCableRequestLengthUnit{
		BlankEnum: v,
	}
}

// LengthUnitEnumAsPatchedWritableCableRequestLengthUnit is a convenience function that returns LengthUnitEnum wrapped in PatchedWritableCableRequestLengthUnit
func LengthUnitEnumAsPatchedWritableCableRequestLengthUnit(v *LengthUnitEnum) PatchedWritableCableRequestLengthUnit {
	return PatchedWritableCableRequestLengthUnit{
		LengthUnitEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableCableRequestLengthUnit) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlankEnum
	err = newStrictDecoder(data).Decode(&dst.BlankEnum)
	if err == nil {
		jsonBlankEnum, _ := json.Marshal(dst.BlankEnum)
		if string(jsonBlankEnum) == "{}" { // empty struct
			dst.BlankEnum = nil
		} else {
			if err = validator.Validate(dst.BlankEnum); err != nil {
				dst.BlankEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlankEnum = nil
	}

	// try to unmarshal data into LengthUnitEnum
	err = newStrictDecoder(data).Decode(&dst.LengthUnitEnum)
	if err == nil {
		jsonLengthUnitEnum, _ := json.Marshal(dst.LengthUnitEnum)
		if string(jsonLengthUnitEnum) == "{}" { // empty struct
			dst.LengthUnitEnum = nil
		} else {
			if err = validator.Validate(dst.LengthUnitEnum); err != nil {
				dst.LengthUnitEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.LengthUnitEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.LengthUnitEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableCableRequestLengthUnit)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableCableRequestLengthUnit)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableCableRequestLengthUnit) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.LengthUnitEnum != nil {
		return json.Marshal(&src.LengthUnitEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableCableRequestLengthUnit) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.LengthUnitEnum != nil {
		return obj.LengthUnitEnum
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableCableRequestLengthUnit struct {
	value *PatchedWritableCableRequestLengthUnit
	isSet bool
}

func (v NullablePatchedWritableCableRequestLengthUnit) Get() *PatchedWritableCableRequestLengthUnit {
	return v.value
}

func (v *NullablePatchedWritableCableRequestLengthUnit) Set(val *PatchedWritableCableRequestLengthUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableCableRequestLengthUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableCableRequestLengthUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableCableRequestLengthUnit(val *PatchedWritableCableRequestLengthUnit) *NullablePatchedWritableCableRequestLengthUnit {
	return &NullablePatchedWritableCableRequestLengthUnit{value: val, isSet: true}
}

func (v NullablePatchedWritableCableRequestLengthUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableCableRequestLengthUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


