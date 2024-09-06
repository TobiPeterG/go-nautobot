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

// DcimRacksElevationListFaceParameter the model 'DcimRacksElevationListFaceParameter'
type DcimRacksElevationListFaceParameter string

// List of dcim_racks_elevation_list_face_parameter
const (
	DCIMRACKSELEVATIONLISTFACEPARAMETER_FRONT DcimRacksElevationListFaceParameter = "front"
	DCIMRACKSELEVATIONLISTFACEPARAMETER_REAR DcimRacksElevationListFaceParameter = "rear"
)

// All allowed values of DcimRacksElevationListFaceParameter enum
var AllowedDcimRacksElevationListFaceParameterEnumValues = []DcimRacksElevationListFaceParameter{
	"front",
	"rear",
}

func (v *DcimRacksElevationListFaceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimRacksElevationListFaceParameter(value)
	for _, existing := range AllowedDcimRacksElevationListFaceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimRacksElevationListFaceParameter", value)
}

// NewDcimRacksElevationListFaceParameterFromValue returns a pointer to a valid DcimRacksElevationListFaceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimRacksElevationListFaceParameterFromValue(v string) (*DcimRacksElevationListFaceParameter, error) {
	ev := DcimRacksElevationListFaceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimRacksElevationListFaceParameter: valid values are %v", v, AllowedDcimRacksElevationListFaceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimRacksElevationListFaceParameter) IsValid() bool {
	for _, existing := range AllowedDcimRacksElevationListFaceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_racks_elevation_list_face_parameter value
func (v DcimRacksElevationListFaceParameter) Ptr() *DcimRacksElevationListFaceParameter {
	return &v
}

type NullableDcimRacksElevationListFaceParameter struct {
	value *DcimRacksElevationListFaceParameter
	isSet bool
}

func (v NullableDcimRacksElevationListFaceParameter) Get() *DcimRacksElevationListFaceParameter {
	return v.value
}

func (v *NullableDcimRacksElevationListFaceParameter) Set(val *DcimRacksElevationListFaceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimRacksElevationListFaceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimRacksElevationListFaceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimRacksElevationListFaceParameter(val *DcimRacksElevationListFaceParameter) *NullableDcimRacksElevationListFaceParameter {
	return &NullableDcimRacksElevationListFaceParameter{value: val, isSet: true}
}

func (v NullableDcimRacksElevationListFaceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimRacksElevationListFaceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

