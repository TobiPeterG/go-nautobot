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

// checks if the SoftwareImageFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareImageFiles{}

// SoftwareImageFiles struct for SoftwareImageFiles
type SoftwareImageFiles struct {
	Id *BulkWritableCableRequestStatusId `json:"id,omitempty"`
	ObjectType *string `json:"object_type,omitempty" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareImageFiles SoftwareImageFiles

// NewSoftwareImageFiles instantiates a new SoftwareImageFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareImageFiles() *SoftwareImageFiles {
	this := SoftwareImageFiles{}
	return &this
}

// NewSoftwareImageFilesWithDefaults instantiates a new SoftwareImageFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareImageFilesWithDefaults() *SoftwareImageFiles {
	this := SoftwareImageFiles{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SoftwareImageFiles) GetId() BulkWritableCableRequestStatusId {
	if o == nil || IsNil(o.Id) {
		var ret BulkWritableCableRequestStatusId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageFiles) GetIdOk() (*BulkWritableCableRequestStatusId, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SoftwareImageFiles) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given BulkWritableCableRequestStatusId and assigns it to the Id field.
func (o *SoftwareImageFiles) SetId(v BulkWritableCableRequestStatusId) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *SoftwareImageFiles) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageFiles) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *SoftwareImageFiles) HasObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *SoftwareImageFiles) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SoftwareImageFiles) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareImageFiles) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SoftwareImageFiles) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SoftwareImageFiles) SetUrl(v string) {
	o.Url = &v
}

func (o SoftwareImageFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareImageFiles) ToMap() (map[string]interface{}, error) {
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

func (o *SoftwareImageFiles) UnmarshalJSON(data []byte) (err error) {
	varSoftwareImageFiles := _SoftwareImageFiles{}

	err = json.Unmarshal(data, &varSoftwareImageFiles)

	if err != nil {
		return err
	}

	*o = SoftwareImageFiles(varSoftwareImageFiles)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareImageFiles struct {
	value *SoftwareImageFiles
	isSet bool
}

func (v NullableSoftwareImageFiles) Get() *SoftwareImageFiles {
	return v.value
}

func (v *NullableSoftwareImageFiles) Set(val *SoftwareImageFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareImageFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareImageFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareImageFiles(val *SoftwareImageFiles) *NullableSoftwareImageFiles {
	return &NullableSoftwareImageFiles{value: val, isSet: true}
}

func (v NullableSoftwareImageFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareImageFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


