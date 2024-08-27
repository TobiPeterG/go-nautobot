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

// checks if the VRFPrefixAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VRFPrefixAssignmentRequest{}

// VRFPrefixAssignmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type VRFPrefixAssignmentRequest struct {
	Vrf BulkWritableCableRequestStatus `json:"vrf"`
	Prefix BulkWritableCableRequestStatus `json:"prefix"`
	AdditionalProperties map[string]interface{}
}

type _VRFPrefixAssignmentRequest VRFPrefixAssignmentRequest

// NewVRFPrefixAssignmentRequest instantiates a new VRFPrefixAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVRFPrefixAssignmentRequest(vrf BulkWritableCableRequestStatus, prefix BulkWritableCableRequestStatus) *VRFPrefixAssignmentRequest {
	this := VRFPrefixAssignmentRequest{}
	this.Vrf = vrf
	this.Prefix = prefix
	return &this
}

// NewVRFPrefixAssignmentRequestWithDefaults instantiates a new VRFPrefixAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVRFPrefixAssignmentRequestWithDefaults() *VRFPrefixAssignmentRequest {
	this := VRFPrefixAssignmentRequest{}
	return &this
}

// GetVrf returns the Vrf field value
func (o *VRFPrefixAssignmentRequest) GetVrf() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Vrf
}

// GetVrfOk returns a tuple with the Vrf field value
// and a boolean to check if the value has been set.
func (o *VRFPrefixAssignmentRequest) GetVrfOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrf, true
}

// SetVrf sets field value
func (o *VRFPrefixAssignmentRequest) SetVrf(v BulkWritableCableRequestStatus) {
	o.Vrf = v
}

// GetPrefix returns the Prefix field value
func (o *VRFPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *VRFPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *VRFPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus) {
	o.Prefix = v
}

func (o VRFPrefixAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VRFPrefixAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vrf"] = o.Vrf
	toSerialize["prefix"] = o.Prefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VRFPrefixAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vrf",
		"prefix",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVRFPrefixAssignmentRequest := _VRFPrefixAssignmentRequest{}

	err = json.Unmarshal(data, &varVRFPrefixAssignmentRequest)

	if err != nil {
		return err
	}

	*o = VRFPrefixAssignmentRequest(varVRFPrefixAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vrf")
		delete(additionalProperties, "prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVRFPrefixAssignmentRequest struct {
	value *VRFPrefixAssignmentRequest
	isSet bool
}

func (v NullableVRFPrefixAssignmentRequest) Get() *VRFPrefixAssignmentRequest {
	return v.value
}

func (v *NullableVRFPrefixAssignmentRequest) Set(val *VRFPrefixAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVRFPrefixAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVRFPrefixAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVRFPrefixAssignmentRequest(val *VRFPrefixAssignmentRequest) *NullableVRFPrefixAssignmentRequest {
	return &NullableVRFPrefixAssignmentRequest{value: val, isSet: true}
}

func (v NullableVRFPrefixAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVRFPrefixAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


