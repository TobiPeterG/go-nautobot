/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedCloudNetworkPrefixAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCloudNetworkPrefixAssignmentRequest{}

// PatchedCloudNetworkPrefixAssignmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedCloudNetworkPrefixAssignmentRequest struct {
	CloudNetwork *BulkWritableCableRequestStatus `json:"cloud_network,omitempty"`
	Prefix *BulkWritableCableRequestStatus `json:"prefix,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedCloudNetworkPrefixAssignmentRequest PatchedCloudNetworkPrefixAssignmentRequest

// NewPatchedCloudNetworkPrefixAssignmentRequest instantiates a new PatchedCloudNetworkPrefixAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCloudNetworkPrefixAssignmentRequest() *PatchedCloudNetworkPrefixAssignmentRequest {
	this := PatchedCloudNetworkPrefixAssignmentRequest{}
	return &this
}

// NewPatchedCloudNetworkPrefixAssignmentRequestWithDefaults instantiates a new PatchedCloudNetworkPrefixAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCloudNetworkPrefixAssignmentRequestWithDefaults() *PatchedCloudNetworkPrefixAssignmentRequest {
	this := PatchedCloudNetworkPrefixAssignmentRequest{}
	return &this
}

// GetCloudNetwork returns the CloudNetwork field value if set, zero value otherwise.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) GetCloudNetwork() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.CloudNetwork) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.CloudNetwork
}

// GetCloudNetworkOk returns a tuple with the CloudNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) GetCloudNetworkOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.CloudNetwork) {
		return nil, false
	}
	return o.CloudNetwork, true
}

// HasCloudNetwork returns a boolean if a field has been set.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) HasCloudNetwork() bool {
	if o != nil && !IsNil(o.CloudNetwork) {
		return true
	}

	return false
}

// SetCloudNetwork gets a reference to the given BulkWritableCableRequestStatus and assigns it to the CloudNetwork field.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) SetCloudNetwork(v BulkWritableCableRequestStatus) {
	o.CloudNetwork = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) GetPrefix() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Prefix) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) GetPrefixOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Prefix field.
func (o *PatchedCloudNetworkPrefixAssignmentRequest) SetPrefix(v BulkWritableCableRequestStatus) {
	o.Prefix = &v
}

func (o PatchedCloudNetworkPrefixAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCloudNetworkPrefixAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudNetwork) {
		toSerialize["cloud_network"] = o.CloudNetwork
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedCloudNetworkPrefixAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedCloudNetworkPrefixAssignmentRequest := _PatchedCloudNetworkPrefixAssignmentRequest{}

	err = json.Unmarshal(data, &varPatchedCloudNetworkPrefixAssignmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedCloudNetworkPrefixAssignmentRequest(varPatchedCloudNetworkPrefixAssignmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cloud_network")
		delete(additionalProperties, "prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedCloudNetworkPrefixAssignmentRequest struct {
	value *PatchedCloudNetworkPrefixAssignmentRequest
	isSet bool
}

func (v NullablePatchedCloudNetworkPrefixAssignmentRequest) Get() *PatchedCloudNetworkPrefixAssignmentRequest {
	return v.value
}

func (v *NullablePatchedCloudNetworkPrefixAssignmentRequest) Set(val *PatchedCloudNetworkPrefixAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCloudNetworkPrefixAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCloudNetworkPrefixAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCloudNetworkPrefixAssignmentRequest(val *PatchedCloudNetworkPrefixAssignmentRequest) *NullablePatchedCloudNetworkPrefixAssignmentRequest {
	return &NullablePatchedCloudNetworkPrefixAssignmentRequest{value: val, isSet: true}
}

func (v NullablePatchedCloudNetworkPrefixAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCloudNetworkPrefixAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


