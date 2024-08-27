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

// PowerFeedPhase the model 'PowerFeedPhase'
type PowerFeedPhase string

// List of PowerFeed_phase
const (
	POWERFEEDPHASE_SINGLE_PHASE PowerFeedPhase = "single-phase"
	POWERFEEDPHASE_THREE_PHASE PowerFeedPhase = "three-phase"
)

// All allowed values of PowerFeedPhase enum
var AllowedPowerFeedPhaseEnumValues = []PowerFeedPhase{
	"single-phase",
	"three-phase",
}

func (v *PowerFeedPhase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PowerFeedPhase(value)
	for _, existing := range AllowedPowerFeedPhaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PowerFeedPhase", value)
}

// NewPowerFeedPhaseFromValue returns a pointer to a valid PowerFeedPhase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPowerFeedPhaseFromValue(v string) (*PowerFeedPhase, error) {
	ev := PowerFeedPhase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PowerFeedPhase: valid values are %v", v, AllowedPowerFeedPhaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PowerFeedPhase) IsValid() bool {
	for _, existing := range AllowedPowerFeedPhaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PowerFeed_phase value
func (v PowerFeedPhase) Ptr() *PowerFeedPhase {
	return &v
}

type NullablePowerFeedPhase struct {
	value *PowerFeedPhase
	isSet bool
}

func (v NullablePowerFeedPhase) Get() *PowerFeedPhase {
	return v.value
}

func (v *NullablePowerFeedPhase) Set(val *PowerFeedPhase) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFeedPhase) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFeedPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFeedPhase(val *PowerFeedPhase) *NullablePowerFeedPhase {
	return &NullablePowerFeedPhase{value: val, isSet: true}
}

func (v NullablePowerFeedPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFeedPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

