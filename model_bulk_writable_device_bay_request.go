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

// checks if the BulkWritableDeviceBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableDeviceBayRequest{}

// BulkWritableDeviceBayRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableDeviceBayRequest struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Device BulkWritableCableRequestStatus `json:"device"`
	InstalledDevice NullableBulkWritableCircuitRequestTenant `json:"installed_device,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableDeviceBayRequest BulkWritableDeviceBayRequest

// NewBulkWritableDeviceBayRequest instantiates a new BulkWritableDeviceBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableDeviceBayRequest(id string, name string, device BulkWritableCableRequestStatus) *BulkWritableDeviceBayRequest {
	this := BulkWritableDeviceBayRequest{}
	this.Id = id
	this.Name = name
	this.Device = device
	return &this
}

// NewBulkWritableDeviceBayRequestWithDefaults instantiates a new BulkWritableDeviceBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableDeviceBayRequestWithDefaults() *BulkWritableDeviceBayRequest {
	this := BulkWritableDeviceBayRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableDeviceBayRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableDeviceBayRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BulkWritableDeviceBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableDeviceBayRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritableDeviceBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritableDeviceBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BulkWritableDeviceBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableDeviceBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableDeviceBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableDeviceBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDevice returns the Device field value
func (o *BulkWritableDeviceBayRequest) GetDevice() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *BulkWritableDeviceBayRequest) SetDevice(v BulkWritableCableRequestStatus) {
	o.Device = v
}

// GetInstalledDevice returns the InstalledDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableDeviceBayRequest) GetInstalledDevice() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.InstalledDevice.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.InstalledDevice.Get()
}

// GetInstalledDeviceOk returns a tuple with the InstalledDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableDeviceBayRequest) GetInstalledDeviceOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledDevice.Get(), o.InstalledDevice.IsSet()
}

// HasInstalledDevice returns a boolean if a field has been set.
func (o *BulkWritableDeviceBayRequest) HasInstalledDevice() bool {
	if o != nil && o.InstalledDevice.IsSet() {
		return true
	}

	return false
}

// SetInstalledDevice gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the InstalledDevice field.
func (o *BulkWritableDeviceBayRequest) SetInstalledDevice(v BulkWritableCircuitRequestTenant) {
	o.InstalledDevice.Set(&v)
}
// SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil
func (o *BulkWritableDeviceBayRequest) SetInstalledDeviceNil() {
	o.InstalledDevice.Set(nil)
}

// UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
func (o *BulkWritableDeviceBayRequest) UnsetInstalledDevice() {
	o.InstalledDevice.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableDeviceBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableDeviceBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableDeviceBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableDeviceBayRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableDeviceBayRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableDeviceBayRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableDeviceBayRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableDeviceBayRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableDeviceBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableDeviceBayRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableDeviceBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableDeviceBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["device"] = o.Device
	if o.InstalledDevice.IsSet() {
		toSerialize["installed_device"] = o.InstalledDevice.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableDeviceBayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"device",
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

	varBulkWritableDeviceBayRequest := _BulkWritableDeviceBayRequest{}

	err = json.Unmarshal(data, &varBulkWritableDeviceBayRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableDeviceBayRequest(varBulkWritableDeviceBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device")
		delete(additionalProperties, "installed_device")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableDeviceBayRequest struct {
	value *BulkWritableDeviceBayRequest
	isSet bool
}

func (v NullableBulkWritableDeviceBayRequest) Get() *BulkWritableDeviceBayRequest {
	return v.value
}

func (v *NullableBulkWritableDeviceBayRequest) Set(val *BulkWritableDeviceBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableDeviceBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableDeviceBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableDeviceBayRequest(val *BulkWritableDeviceBayRequest) *NullableBulkWritableDeviceBayRequest {
	return &NullableBulkWritableDeviceBayRequest{value: val, isSet: true}
}

func (v NullableBulkWritableDeviceBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableDeviceBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


