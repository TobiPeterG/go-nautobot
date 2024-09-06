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

// RedundancyProtocol - struct for RedundancyProtocol
type RedundancyProtocol struct {
	BlankEnum *BlankEnum
	InterfaceRedundancyGroupProtocolChoices *InterfaceRedundancyGroupProtocolChoices
}

// BlankEnumAsRedundancyProtocol is a convenience function that returns BlankEnum wrapped in RedundancyProtocol
func BlankEnumAsRedundancyProtocol(v *BlankEnum) RedundancyProtocol {
	return RedundancyProtocol{
		BlankEnum: v,
	}
}

// InterfaceRedundancyGroupProtocolChoicesAsRedundancyProtocol is a convenience function that returns InterfaceRedundancyGroupProtocolChoices wrapped in RedundancyProtocol
func InterfaceRedundancyGroupProtocolChoicesAsRedundancyProtocol(v *InterfaceRedundancyGroupProtocolChoices) RedundancyProtocol {
	return RedundancyProtocol{
		InterfaceRedundancyGroupProtocolChoices: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RedundancyProtocol) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into InterfaceRedundancyGroupProtocolChoices
	err = newStrictDecoder(data).Decode(&dst.InterfaceRedundancyGroupProtocolChoices)
	if err == nil {
		jsonInterfaceRedundancyGroupProtocolChoices, _ := json.Marshal(dst.InterfaceRedundancyGroupProtocolChoices)
		if string(jsonInterfaceRedundancyGroupProtocolChoices) == "{}" { // empty struct
			dst.InterfaceRedundancyGroupProtocolChoices = nil
		} else {
			if err = validator.Validate(dst.InterfaceRedundancyGroupProtocolChoices); err != nil {
				dst.InterfaceRedundancyGroupProtocolChoices = nil
			} else {
				match++
			}
		}
	} else {
		dst.InterfaceRedundancyGroupProtocolChoices = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlankEnum = nil
		dst.InterfaceRedundancyGroupProtocolChoices = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RedundancyProtocol)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RedundancyProtocol)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RedundancyProtocol) MarshalJSON() ([]byte, error) {
	if src.BlankEnum != nil {
		return json.Marshal(&src.BlankEnum)
	}

	if src.InterfaceRedundancyGroupProtocolChoices != nil {
		return json.Marshal(&src.InterfaceRedundancyGroupProtocolChoices)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RedundancyProtocol) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlankEnum != nil {
		return obj.BlankEnum
	}

	if obj.InterfaceRedundancyGroupProtocolChoices != nil {
		return obj.InterfaceRedundancyGroupProtocolChoices
	}

	// all schemas are nil
	return nil
}

type NullableRedundancyProtocol struct {
	value *RedundancyProtocol
	isSet bool
}

func (v NullableRedundancyProtocol) Get() *RedundancyProtocol {
	return v.value
}

func (v *NullableRedundancyProtocol) Set(val *RedundancyProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableRedundancyProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableRedundancyProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedundancyProtocol(val *RedundancyProtocol) *NullableRedundancyProtocol {
	return &NullableRedundancyProtocol{value: val, isSet: true}
}

func (v NullableRedundancyProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedundancyProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


