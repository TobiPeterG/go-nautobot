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

// BulkWritablePowerFeedRequestSupplyValue the model 'BulkWritablePowerFeedRequestSupplyValue'
type BulkWritablePowerFeedRequestSupplyValue string

// List of BulkWritablePowerFeedRequest_supply_value
const (
	BULKWRITABLEPOWERFEEDREQUESTSUPPLYVALUE_AC BulkWritablePowerFeedRequestSupplyValue = "ac"
	BULKWRITABLEPOWERFEEDREQUESTSUPPLYVALUE_DC BulkWritablePowerFeedRequestSupplyValue = "dc"
)

// All allowed values of BulkWritablePowerFeedRequestSupplyValue enum
var AllowedBulkWritablePowerFeedRequestSupplyValueEnumValues = []BulkWritablePowerFeedRequestSupplyValue{
	"ac",
	"dc",
}

func (v *BulkWritablePowerFeedRequestSupplyValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BulkWritablePowerFeedRequestSupplyValue(value)
	for _, existing := range AllowedBulkWritablePowerFeedRequestSupplyValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BulkWritablePowerFeedRequestSupplyValue", value)
}

// NewBulkWritablePowerFeedRequestSupplyValueFromValue returns a pointer to a valid BulkWritablePowerFeedRequestSupplyValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBulkWritablePowerFeedRequestSupplyValueFromValue(v string) (*BulkWritablePowerFeedRequestSupplyValue, error) {
	ev := BulkWritablePowerFeedRequestSupplyValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BulkWritablePowerFeedRequestSupplyValue: valid values are %v", v, AllowedBulkWritablePowerFeedRequestSupplyValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BulkWritablePowerFeedRequestSupplyValue) IsValid() bool {
	for _, existing := range AllowedBulkWritablePowerFeedRequestSupplyValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BulkWritablePowerFeedRequest_supply_value value
func (v BulkWritablePowerFeedRequestSupplyValue) Ptr() *BulkWritablePowerFeedRequestSupplyValue {
	return &v
}

type NullableBulkWritablePowerFeedRequestSupplyValue struct {
	value *BulkWritablePowerFeedRequestSupplyValue
	isSet bool
}

func (v NullableBulkWritablePowerFeedRequestSupplyValue) Get() *BulkWritablePowerFeedRequestSupplyValue {
	return v.value
}

func (v *NullableBulkWritablePowerFeedRequestSupplyValue) Set(val *BulkWritablePowerFeedRequestSupplyValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritablePowerFeedRequestSupplyValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritablePowerFeedRequestSupplyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritablePowerFeedRequestSupplyValue(val *BulkWritablePowerFeedRequestSupplyValue) *NullableBulkWritablePowerFeedRequestSupplyValue {
	return &NullableBulkWritablePowerFeedRequestSupplyValue{value: val, isSet: true}
}

func (v NullableBulkWritablePowerFeedRequestSupplyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritablePowerFeedRequestSupplyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

