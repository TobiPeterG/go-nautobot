/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the CircuitCircuitTerminationA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitCircuitTerminationA{}

// CircuitCircuitTerminationA struct for CircuitCircuitTerminationA
type CircuitCircuitTerminationA struct {
	Id *BulkWritableCableRequestStatusId `json:"id,omitempty"`
	ObjectType *string `json:"object_type,omitempty" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CircuitCircuitTerminationA CircuitCircuitTerminationA

// NewCircuitCircuitTerminationA instantiates a new CircuitCircuitTerminationA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitCircuitTerminationA() *CircuitCircuitTerminationA {
	this := CircuitCircuitTerminationA{}
	return &this
}

// NewCircuitCircuitTerminationAWithDefaults instantiates a new CircuitCircuitTerminationA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitCircuitTerminationAWithDefaults() *CircuitCircuitTerminationA {
	this := CircuitCircuitTerminationA{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationA) GetId() BulkWritableCableRequestStatusId {
	if o == nil || IsNil(o.Id) {
		var ret BulkWritableCableRequestStatusId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationA) GetIdOk() (*BulkWritableCableRequestStatusId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationA) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given BulkWritableCableRequestStatusId and assigns it to the Id field.
func (o *CircuitCircuitTerminationA) SetId(v BulkWritableCableRequestStatusId) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationA) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationA) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationA) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *CircuitCircuitTerminationA) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CircuitCircuitTerminationA) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitCircuitTerminationA) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CircuitCircuitTerminationA) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CircuitCircuitTerminationA) SetUrl(v string) {
	o.Url = &v
}

func (o CircuitCircuitTerminationA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitCircuitTerminationA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ObjectType) {
		toSerialize["object_type"] = o.ObjectType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CircuitCircuitTerminationA) UnmarshalJSON(data []byte) (err error) {
	varCircuitCircuitTerminationA := _CircuitCircuitTerminationA{}

	err = json.Unmarshal(data, &varCircuitCircuitTerminationA)

	if err != nil {
		return err
	}

	*o = CircuitCircuitTerminationA(varCircuitCircuitTerminationA)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitCircuitTerminationA struct {
	value *CircuitCircuitTerminationA
	isSet bool
}

func (v NullableCircuitCircuitTerminationA) Get() *CircuitCircuitTerminationA {
	return v.value
}

func (v *NullableCircuitCircuitTerminationA) Set(val *CircuitCircuitTerminationA) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitCircuitTerminationA) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitCircuitTerminationA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitCircuitTerminationA(val *CircuitCircuitTerminationA) *NullableCircuitCircuitTerminationA {
	return &NullableCircuitCircuitTerminationA{value: val, isSet: true}
}

func (v NullableCircuitCircuitTerminationA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitCircuitTerminationA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


