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

// JobExecutionTypeIntervalChoices the model 'JobExecutionTypeIntervalChoices'
type JobExecutionTypeIntervalChoices string

// List of JobExecutionTypeIntervalChoices
const (
	JOBEXECUTIONTYPEINTERVALCHOICES_IMMEDIATELY JobExecutionTypeIntervalChoices = "immediately"
	JOBEXECUTIONTYPEINTERVALCHOICES_FUTURE JobExecutionTypeIntervalChoices = "future"
	JOBEXECUTIONTYPEINTERVALCHOICES_HOURLY JobExecutionTypeIntervalChoices = "hourly"
	JOBEXECUTIONTYPEINTERVALCHOICES_DAILY JobExecutionTypeIntervalChoices = "daily"
	JOBEXECUTIONTYPEINTERVALCHOICES_WEEKLY JobExecutionTypeIntervalChoices = "weekly"
	JOBEXECUTIONTYPEINTERVALCHOICES_CUSTOM JobExecutionTypeIntervalChoices = "custom"
)

// All allowed values of JobExecutionTypeIntervalChoices enum
var AllowedJobExecutionTypeIntervalChoicesEnumValues = []JobExecutionTypeIntervalChoices{
	"immediately",
	"future",
	"hourly",
	"daily",
	"weekly",
	"custom",
}

func (v *JobExecutionTypeIntervalChoices) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobExecutionTypeIntervalChoices(value)
	for _, existing := range AllowedJobExecutionTypeIntervalChoicesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobExecutionTypeIntervalChoices", value)
}

// NewJobExecutionTypeIntervalChoicesFromValue returns a pointer to a valid JobExecutionTypeIntervalChoices
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobExecutionTypeIntervalChoicesFromValue(v string) (*JobExecutionTypeIntervalChoices, error) {
	ev := JobExecutionTypeIntervalChoices(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobExecutionTypeIntervalChoices: valid values are %v", v, AllowedJobExecutionTypeIntervalChoicesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobExecutionTypeIntervalChoices) IsValid() bool {
	for _, existing := range AllowedJobExecutionTypeIntervalChoicesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobExecutionTypeIntervalChoices value
func (v JobExecutionTypeIntervalChoices) Ptr() *JobExecutionTypeIntervalChoices {
	return &v
}

type NullableJobExecutionTypeIntervalChoices struct {
	value *JobExecutionTypeIntervalChoices
	isSet bool
}

func (v NullableJobExecutionTypeIntervalChoices) Get() *JobExecutionTypeIntervalChoices {
	return v.value
}

func (v *NullableJobExecutionTypeIntervalChoices) Set(val *JobExecutionTypeIntervalChoices) {
	v.value = val
	v.isSet = true
}

func (v NullableJobExecutionTypeIntervalChoices) IsSet() bool {
	return v.isSet
}

func (v *NullableJobExecutionTypeIntervalChoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobExecutionTypeIntervalChoices(val *JobExecutionTypeIntervalChoices) *NullableJobExecutionTypeIntervalChoices {
	return &NullableJobExecutionTypeIntervalChoices{value: val, isSet: true}
}

func (v NullableJobExecutionTypeIntervalChoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobExecutionTypeIntervalChoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

