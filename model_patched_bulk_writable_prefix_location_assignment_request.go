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

// checks if the PatchedBulkWritablePrefixLocationAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritablePrefixLocationAssignmentRequest{}

// PatchedBulkWritablePrefixLocationAssignmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritablePrefixLocationAssignmentRequest struct {
	Id string `json:"id"`
	Prefix *BulkWritableCableRequestStatus `json:"prefix,omitempty"`
	Location *BulkWritableCableRequestStatus `json:"location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritablePrefixLocationAssignmentRequest PatchedBulkWritablePrefixLocationAssignmentRequest

// NewPatchedBulkWritablePrefixLocationAssignmentRequest instantiates a new PatchedBulkWritablePrefixLocationAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritablePrefixLocationAssignmentRequest(id string) *PatchedBulkWritablePrefixLocationAssignmentRequest {
	this := PatchedBulkWritablePrefixLocationAssignmentRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritablePrefixLocationAssignmentRequestWithDefaults instantiates a new PatchedBulkWritablePrefixLocationAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritablePrefixLocationAssignmentRequestWithDefaults() *PatchedBulkWritablePrefixLocationAssignmentRequest {
	this := PatchedBulkWritablePrefixLocationAssignmentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) SetId(v string) {
	o.Id = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Prefix) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Prefix field.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus) {
	o.Prefix = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetLocation() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Location) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Location field.
func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = &v
}

func (o PatchedBulkWritablePrefixLocationAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritablePrefixLocationAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritablePrefixLocationAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPatchedBulkWritablePrefixLocationAssignmentRequest := _PatchedBulkWritablePrefixLocationAssignmentRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritablePrefixLocationAssignmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritablePrefixLocationAssignmentRequest(varPatchedBulkWritablePrefixLocationAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritablePrefixLocationAssignmentRequest struct {
	value *PatchedBulkWritablePrefixLocationAssignmentRequest
	isSet bool
}

func (v NullablePatchedBulkWritablePrefixLocationAssignmentRequest) Get() *PatchedBulkWritablePrefixLocationAssignmentRequest {
	return v.value
}

func (v *NullablePatchedBulkWritablePrefixLocationAssignmentRequest) Set(val *PatchedBulkWritablePrefixLocationAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritablePrefixLocationAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritablePrefixLocationAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritablePrefixLocationAssignmentRequest(val *PatchedBulkWritablePrefixLocationAssignmentRequest) *NullablePatchedBulkWritablePrefixLocationAssignmentRequest {
	return &NullablePatchedBulkWritablePrefixLocationAssignmentRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritablePrefixLocationAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritablePrefixLocationAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


