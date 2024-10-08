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

// checks if the ModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleBayRequest{}

// ModuleBayRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type ModuleBayRequest struct {
	Name string `json:"name"`
	// The position of the module bay within the parent device/module
	Position *string `json:"position,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	ParentDevice NullableBulkWritableCircuitRequestTenant `json:"parent_device,omitempty"`
	ParentModule NullableBulkWritableCircuitRequestTenant `json:"parent_module,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModuleBayRequest ModuleBayRequest

// NewModuleBayRequest instantiates a new ModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleBayRequest(name string) *ModuleBayRequest {
	this := ModuleBayRequest{}
	this.Name = name
	return &this
}

// NewModuleBayRequestWithDefaults instantiates a new ModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleBayRequestWithDefaults() *ModuleBayRequest {
	this := ModuleBayRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ModuleBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModuleBayRequest) SetName(v string) {
	o.Name = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ModuleBayRequest) SetPosition(v string) {
	o.Position = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ModuleBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModuleBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetParentDevice returns the ParentDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleBayRequest) GetParentDevice() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ParentDevice.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ParentDevice.Get()
}

// GetParentDeviceOk returns a tuple with the ParentDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleBayRequest) GetParentDeviceOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentDevice.Get(), o.ParentDevice.IsSet()
}

// HasParentDevice returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasParentDevice() bool {
	if o != nil && o.ParentDevice.IsSet() {
		return true
	}

	return false
}

// SetParentDevice gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ParentDevice field.
func (o *ModuleBayRequest) SetParentDevice(v BulkWritableCircuitRequestTenant) {
	o.ParentDevice.Set(&v)
}
// SetParentDeviceNil sets the value for ParentDevice to be an explicit nil
func (o *ModuleBayRequest) SetParentDeviceNil() {
	o.ParentDevice.Set(nil)
}

// UnsetParentDevice ensures that no value is present for ParentDevice, not even an explicit nil
func (o *ModuleBayRequest) UnsetParentDevice() {
	o.ParentDevice.Unset()
}

// GetParentModule returns the ParentModule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleBayRequest) GetParentModule() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ParentModule.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ParentModule.Get()
}

// GetParentModuleOk returns a tuple with the ParentModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModuleBayRequest) GetParentModuleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentModule.Get(), o.ParentModule.IsSet()
}

// HasParentModule returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasParentModule() bool {
	if o != nil && o.ParentModule.IsSet() {
		return true
	}

	return false
}

// SetParentModule gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ParentModule field.
func (o *ModuleBayRequest) SetParentModule(v BulkWritableCircuitRequestTenant) {
	o.ParentModule.Set(&v)
}
// SetParentModuleNil sets the value for ParentModule to be an explicit nil
func (o *ModuleBayRequest) SetParentModuleNil() {
	o.ParentModule.Set(nil)
}

// UnsetParentModule ensures that no value is present for ParentModule, not even an explicit nil
func (o *ModuleBayRequest) UnsetParentModule() {
	o.ParentModule.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ModuleBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *ModuleBayRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModuleBayRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleBayRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModuleBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *ModuleBayRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o ModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ParentDevice.IsSet() {
		toSerialize["parent_device"] = o.ParentDevice.Get()
	}
	if o.ParentModule.IsSet() {
		toSerialize["parent_module"] = o.ParentModule.Get()
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

func (o *ModuleBayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varModuleBayRequest := _ModuleBayRequest{}

	err = json.Unmarshal(data, &varModuleBayRequest)

	if err != nil {
		return err
	}

	*o = ModuleBayRequest(varModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "position")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "parent_device")
		delete(additionalProperties, "parent_module")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModuleBayRequest struct {
	value *ModuleBayRequest
	isSet bool
}

func (v NullableModuleBayRequest) Get() *ModuleBayRequest {
	return v.value
}

func (v *NullableModuleBayRequest) Set(val *ModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleBayRequest(val *ModuleBayRequest) *NullableModuleBayRequest {
	return &NullableModuleBayRequest{value: val, isSet: true}
}

func (v NullableModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


