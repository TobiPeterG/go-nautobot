/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the PatchedBulkWritableTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableTokenRequest{}

// PatchedBulkWritableTokenRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableTokenRequest struct {
	Id string `json:"id"`
	Key *string `json:"key,omitempty"`
	Expires NullableTime `json:"expires,omitempty"`
	// Permit create/update/delete operations using this key
	WriteEnabled *bool `json:"write_enabled,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableTokenRequest PatchedBulkWritableTokenRequest

// NewPatchedBulkWritableTokenRequest instantiates a new PatchedBulkWritableTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableTokenRequest(id string) *PatchedBulkWritableTokenRequest {
	this := PatchedBulkWritableTokenRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableTokenRequestWithDefaults instantiates a new PatchedBulkWritableTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableTokenRequestWithDefaults() *PatchedBulkWritableTokenRequest {
	this := PatchedBulkWritableTokenRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableTokenRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableTokenRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableTokenRequest) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PatchedBulkWritableTokenRequest) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableTokenRequest) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PatchedBulkWritableTokenRequest) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PatchedBulkWritableTokenRequest) SetKey(v string) {
	o.Key = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableTokenRequest) GetExpires() time.Time {
	if o == nil || IsNil(o.Expires.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableTokenRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *PatchedBulkWritableTokenRequest) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *PatchedBulkWritableTokenRequest) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}
// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *PatchedBulkWritableTokenRequest) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *PatchedBulkWritableTokenRequest) UnsetExpires() {
	o.Expires.Unset()
}

// GetWriteEnabled returns the WriteEnabled field value if set, zero value otherwise.
func (o *PatchedBulkWritableTokenRequest) GetWriteEnabled() bool {
	if o == nil || IsNil(o.WriteEnabled) {
		var ret bool
		return ret
	}
	return *o.WriteEnabled
}

// GetWriteEnabledOk returns a tuple with the WriteEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableTokenRequest) GetWriteEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WriteEnabled) {
		return nil, false
	}
	return o.WriteEnabled, true
}

// HasWriteEnabled returns a boolean if a field has been set.
func (o *PatchedBulkWritableTokenRequest) HasWriteEnabled() bool {
	if o != nil && !IsNil(o.WriteEnabled) {
		return true
	}

	return false
}

// SetWriteEnabled gets a reference to the given bool and assigns it to the WriteEnabled field.
func (o *PatchedBulkWritableTokenRequest) SetWriteEnabled(v bool) {
	o.WriteEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableTokenRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableTokenRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableTokenRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableTokenRequest) SetDescription(v string) {
	o.Description = &v
}

func (o PatchedBulkWritableTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if !IsNil(o.WriteEnabled) {
		toSerialize["write_enabled"] = o.WriteEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableTokenRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableTokenRequest := _PatchedBulkWritableTokenRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableTokenRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableTokenRequest(varPatchedBulkWritableTokenRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "expires")
		delete(additionalProperties, "write_enabled")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableTokenRequest struct {
	value *PatchedBulkWritableTokenRequest
	isSet bool
}

func (v NullablePatchedBulkWritableTokenRequest) Get() *PatchedBulkWritableTokenRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableTokenRequest) Set(val *PatchedBulkWritableTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableTokenRequest(val *PatchedBulkWritableTokenRequest) *NullablePatchedBulkWritableTokenRequest {
	return &NullablePatchedBulkWritableTokenRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


