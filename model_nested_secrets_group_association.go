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

// checks if the NestedSecretsGroupAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedSecretsGroupAssociation{}

// NestedSecretsGroupAssociation Serializer for `SecretsGroupAssociation` objects.
type NestedSecretsGroupAssociation struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _NestedSecretsGroupAssociation NestedSecretsGroupAssociation

// NewNestedSecretsGroupAssociation instantiates a new NestedSecretsGroupAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedSecretsGroupAssociation(id string, objectType string, url string) *NestedSecretsGroupAssociation {
	this := NestedSecretsGroupAssociation{}
	this.Id = id
	this.ObjectType = objectType
	this.Url = url
	return &this
}

// NewNestedSecretsGroupAssociationWithDefaults instantiates a new NestedSecretsGroupAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedSecretsGroupAssociationWithDefaults() *NestedSecretsGroupAssociation {
	this := NestedSecretsGroupAssociation{}
	return &this
}

// GetId returns the Id field value
func (o *NestedSecretsGroupAssociation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NestedSecretsGroupAssociation) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *NestedSecretsGroupAssociation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NestedSecretsGroupAssociation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *NestedSecretsGroupAssociation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NestedSecretsGroupAssociation) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NestedSecretsGroupAssociation) SetUrl(v string) {
	o.Url = v
}

func (o NestedSecretsGroupAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedSecretsGroupAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedSecretsGroupAssociation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"url",
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

	varNestedSecretsGroupAssociation := _NestedSecretsGroupAssociation{}

	err = json.Unmarshal(data, &varNestedSecretsGroupAssociation)

	if err != nil {
		return err
	}

	*o = NestedSecretsGroupAssociation(varNestedSecretsGroupAssociation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedSecretsGroupAssociation struct {
	value *NestedSecretsGroupAssociation
	isSet bool
}

func (v NullableNestedSecretsGroupAssociation) Get() *NestedSecretsGroupAssociation {
	return v.value
}

func (v *NullableNestedSecretsGroupAssociation) Set(val *NestedSecretsGroupAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedSecretsGroupAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedSecretsGroupAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedSecretsGroupAssociation(val *NestedSecretsGroupAssociation) *NullableNestedSecretsGroupAssociation {
	return &NullableNestedSecretsGroupAssociation{value: val, isSet: true}
}

func (v NullableNestedSecretsGroupAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedSecretsGroupAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


