/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// PatchedWritableRackRequestOuterUnit - struct for PatchedWritableRackRequestOuterUnit
type PatchedWritableRackRequestOuterUnit struct {
	BlankEnum *BlankEnum
	OuterUnitEnum *OuterUnitEnum
}

// BlankEnumAsPatchedWritableRackRequestOuterUnit is a convenience function that returns BlankEnum wrapped in PatchedWritableRackRequestOuterUnit
func BlankEnumAsPatchedWritableRackRequestOuterUnit(v *BlankEnum) PatchedWritableRackRequestOuterUnit {
	return PatchedWritableRackRequestOuterUnit{
		BlankEnum: v,
	}
}

// OuterUnitEnumAsPatchedWritableRackRequestOuterUnit is a convenience function that returns OuterUnitEnum wrapped in PatchedWritableRackRequestOuterUnit
func OuterUnitEnumAsPatchedWritableRackRequestOuterUnit(v *OuterUnitEnum) PatchedWritableRackRequestOuterUnit {
	return PatchedWritableRackRequestOuterUnit{
		OuterUnitEnum: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PatchedWritableRackRequestOuterUnit) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into OuterUnitEnum
	err = newStrictDecoder(data).Decode(&dst.OuterUnitEnum)
	if err == nil {
		jsonOuterUnitEnum, _ := json.Marshal(dst.OuterUnitEnum)
		if string(jsonOuterUnitEnum) == "{}" { // empty struct
			dst.OuterUnitEnum = nil
		} else {
			if err = validator.Validate(dst.OuterUnitEnum); err != nil {
				dst.OuterUnitEnum = nil
			} else {
				match++
			}
		}
	} else {
		dst.OuterUnitEnum = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.OuterUnitEnum = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PatchedWritableRackRequestOuterUnit)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PatchedWritableRackRequestOuterUnit)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PatchedWritableRackRequestOuterUnit) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.OuterUnitEnum != nil {
		return json.Marshal(&src.OuterUnitEnum)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PatchedWritableRackRequestOuterUnit) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.OuterUnitEnum != nil {
		return obj.OuterUnitEnum
	}

	// all schemas are nil
	return nil
}

type NullablePatchedWritableRackRequestOuterUnit struct {
	value *PatchedWritableRackRequestOuterUnit
	isSet bool
}

func (v NullablePatchedWritableRackRequestOuterUnit) Get() *PatchedWritableRackRequestOuterUnit {
	return v.value
}

func (v *NullablePatchedWritableRackRequestOuterUnit) Set(val *PatchedWritableRackRequestOuterUnit) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableRackRequestOuterUnit) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableRackRequestOuterUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableRackRequestOuterUnit(val *PatchedWritableRackRequestOuterUnit) *NullablePatchedWritableRackRequestOuterUnit {
	return &NullablePatchedWritableRackRequestOuterUnit{value: val, isSet: true}
}

func (v NullablePatchedWritableRackRequestOuterUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableRackRequestOuterUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


