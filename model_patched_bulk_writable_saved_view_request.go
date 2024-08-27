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

// checks if the PatchedBulkWritableSavedViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableSavedViewRequest{}

// PatchedBulkWritableSavedViewRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedBulkWritableSavedViewRequest struct {
	Id string `json:"id"`
	// The name of this view
	Name *string `json:"name,omitempty"`
	// The name of the list view that the saved view is derived from, e.g. dcim:device_list
	View *string `json:"view,omitempty"`
	// Saved Configuration on this view
	Config interface{} `json:"config,omitempty"`
	IsGlobalDefault *bool `json:"is_global_default,omitempty"`
	IsShared *bool `json:"is_shared,omitempty"`
	Owner *BulkWritableSavedViewRequestOwner `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableSavedViewRequest PatchedBulkWritableSavedViewRequest

// NewPatchedBulkWritableSavedViewRequest instantiates a new PatchedBulkWritableSavedViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableSavedViewRequest(id string) *PatchedBulkWritableSavedViewRequest {
	this := PatchedBulkWritableSavedViewRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableSavedViewRequestWithDefaults instantiates a new PatchedBulkWritableSavedViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableSavedViewRequestWithDefaults() *PatchedBulkWritableSavedViewRequest {
	this := PatchedBulkWritableSavedViewRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableSavedViewRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSavedViewRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableSavedViewRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableSavedViewRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSavedViewRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableSavedViewRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableSavedViewRequest) SetName(v string) {
	o.Name = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *PatchedBulkWritableSavedViewRequest) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSavedViewRequest) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *PatchedBulkWritableSavedViewRequest) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *PatchedBulkWritableSavedViewRequest) SetView(v string) {
	o.View = &v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableSavedViewRequest) GetConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableSavedViewRequest) GetConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PatchedBulkWritableSavedViewRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *PatchedBulkWritableSavedViewRequest) SetConfig(v interface{}) {
	o.Config = v
}

// GetIsGlobalDefault returns the IsGlobalDefault field value if set, zero value otherwise.
func (o *PatchedBulkWritableSavedViewRequest) GetIsGlobalDefault() bool {
	if o == nil || IsNil(o.IsGlobalDefault) {
		var ret bool
		return ret
	}
	return *o.IsGlobalDefault
}

// GetIsGlobalDefaultOk returns a tuple with the IsGlobalDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSavedViewRequest) GetIsGlobalDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalDefault) {
		return nil, false
	}
	return o.IsGlobalDefault, true
}

// HasIsGlobalDefault returns a boolean if a field has been set.
func (o *PatchedBulkWritableSavedViewRequest) HasIsGlobalDefault() bool {
	if o != nil && !IsNil(o.IsGlobalDefault) {
		return true
	}

	return false
}

// SetIsGlobalDefault gets a reference to the given bool and assigns it to the IsGlobalDefault field.
func (o *PatchedBulkWritableSavedViewRequest) SetIsGlobalDefault(v bool) {
	o.IsGlobalDefault = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *PatchedBulkWritableSavedViewRequest) GetIsShared() bool {
	if o == nil || IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSavedViewRequest) GetIsSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *PatchedBulkWritableSavedViewRequest) HasIsShared() bool {
	if o != nil && !IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *PatchedBulkWritableSavedViewRequest) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PatchedBulkWritableSavedViewRequest) GetOwner() BulkWritableSavedViewRequestOwner {
	if o == nil || IsNil(o.Owner) {
		var ret BulkWritableSavedViewRequestOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSavedViewRequest) GetOwnerOk() (*BulkWritableSavedViewRequestOwner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PatchedBulkWritableSavedViewRequest) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BulkWritableSavedViewRequestOwner and assigns it to the Owner field.
func (o *PatchedBulkWritableSavedViewRequest) SetOwner(v BulkWritableSavedViewRequestOwner) {
	o.Owner = &v
}

func (o PatchedBulkWritableSavedViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableSavedViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.IsGlobalDefault) {
		toSerialize["is_global_default"] = o.IsGlobalDefault
	}
	if !IsNil(o.IsShared) {
		toSerialize["is_shared"] = o.IsShared
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableSavedViewRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableSavedViewRequest := _PatchedBulkWritableSavedViewRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableSavedViewRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableSavedViewRequest(varPatchedBulkWritableSavedViewRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "view")
		delete(additionalProperties, "config")
		delete(additionalProperties, "is_global_default")
		delete(additionalProperties, "is_shared")
		delete(additionalProperties, "owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableSavedViewRequest struct {
	value *PatchedBulkWritableSavedViewRequest
	isSet bool
}

func (v NullablePatchedBulkWritableSavedViewRequest) Get() *PatchedBulkWritableSavedViewRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableSavedViewRequest) Set(val *PatchedBulkWritableSavedViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableSavedViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableSavedViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableSavedViewRequest(val *PatchedBulkWritableSavedViewRequest) *NullablePatchedBulkWritableSavedViewRequest {
	return &NullablePatchedBulkWritableSavedViewRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableSavedViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableSavedViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


