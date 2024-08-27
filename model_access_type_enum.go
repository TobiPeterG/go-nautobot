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

// AccessTypeEnum the model 'AccessTypeEnum'
type AccessTypeEnum string

// List of AccessTypeEnum
const (
	ACCESSTYPEENUM_GENERIC AccessTypeEnum = "Generic"
	ACCESSTYPEENUM_CONSOLE AccessTypeEnum = "Console"
	ACCESSTYPEENUM_G_NMI AccessTypeEnum = "gNMI"
	ACCESSTYPEENUM_HTTP_S AccessTypeEnum = "HTTP(S)"
	ACCESSTYPEENUM_NETCONF AccessTypeEnum = "NETCONF"
	ACCESSTYPEENUM_REST AccessTypeEnum = "REST"
	ACCESSTYPEENUM_RESTCONF AccessTypeEnum = "RESTCONF"
	ACCESSTYPEENUM_SNMP AccessTypeEnum = "SNMP"
	ACCESSTYPEENUM_SSH AccessTypeEnum = "SSH"
)

// All allowed values of AccessTypeEnum enum
var AllowedAccessTypeEnumEnumValues = []AccessTypeEnum{
	"Generic",
	"Console",
	"gNMI",
	"HTTP(S)",
	"NETCONF",
	"REST",
	"RESTCONF",
	"SNMP",
	"SSH",
}

func (v *AccessTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessTypeEnum(value)
	for _, existing := range AllowedAccessTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessTypeEnum", value)
}

// NewAccessTypeEnumFromValue returns a pointer to a valid AccessTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessTypeEnumFromValue(v string) (*AccessTypeEnum, error) {
	ev := AccessTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessTypeEnum: valid values are %v", v, AllowedAccessTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessTypeEnum) IsValid() bool {
	for _, existing := range AllowedAccessTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessTypeEnum value
func (v AccessTypeEnum) Ptr() *AccessTypeEnum {
	return &v
}

type NullableAccessTypeEnum struct {
	value *AccessTypeEnum
	isSet bool
}

func (v NullableAccessTypeEnum) Get() *AccessTypeEnum {
	return v.value
}

func (v *NullableAccessTypeEnum) Set(val *AccessTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTypeEnum(val *AccessTypeEnum) *NullableAccessTypeEnum {
	return &NullableAccessTypeEnum{value: val, isSet: true}
}

func (v NullableAccessTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

