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

// checks if the FrontPortTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrontPortTemplate{}

// FrontPortTemplate Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type FrontPortTemplate struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	NotesUrl string `json:"notes_url"`
	Type FrontPortType `json:"type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	DeviceType NullableBulkWritableCircuitRequestTenant `json:"device_type,omitempty"`
	ModuleType NullableBulkWritableCircuitRequestTenant `json:"module_type,omitempty"`
	RearPortTemplate BulkWritableCableRequestStatus `json:"rear_port_template"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FrontPortTemplate FrontPortTemplate

// NewFrontPortTemplate instantiates a new FrontPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPortTemplate(id string, objectType string, display string, url string, naturalSlug string, notesUrl string, type_ FrontPortType, name string, rearPortTemplate BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime) *FrontPortTemplate {
	this := FrontPortTemplate{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.NotesUrl = notesUrl
	this.Type = type_
	this.Name = name
	this.RearPortTemplate = rearPortTemplate
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewFrontPortTemplateWithDefaults instantiates a new FrontPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortTemplateWithDefaults() *FrontPortTemplate {
	this := FrontPortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *FrontPortTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FrontPortTemplate) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *FrontPortTemplate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FrontPortTemplate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *FrontPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FrontPortTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *FrontPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FrontPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *FrontPortTemplate) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *FrontPortTemplate) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetNotesUrl returns the NotesUrl field value
func (o *FrontPortTemplate) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *FrontPortTemplate) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetType returns the Type field value
func (o *FrontPortTemplate) GetType() FrontPortType {
	if o == nil {
		var ret FrontPortType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetTypeOk() (*FrontPortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrontPortTemplate) SetType(v FrontPortType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *FrontPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPortTemplate) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPortTemplate) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPortTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *FrontPortTemplate) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPortTemplate) GetDeviceType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetDeviceTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the DeviceType field.
func (o *FrontPortTemplate) SetDeviceType(v BulkWritableCircuitRequestTenant) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *FrontPortTemplate) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *FrontPortTemplate) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPortTemplate) GetModuleType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetModuleTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ModuleType field.
func (o *FrontPortTemplate) SetModuleType(v BulkWritableCircuitRequestTenant) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *FrontPortTemplate) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *FrontPortTemplate) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetRearPortTemplate returns the RearPortTemplate field value
func (o *FrontPortTemplate) GetRearPortTemplate() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.RearPortTemplate
}

// GetRearPortTemplateOk returns a tuple with the RearPortTemplate field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetRearPortTemplateOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPortTemplate, true
}

// SetRearPortTemplate sets field value
func (o *FrontPortTemplate) SetRearPortTemplate(v BulkWritableCableRequestStatus) {
	o.RearPortTemplate = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FrontPortTemplate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *FrontPortTemplate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FrontPortTemplate) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplate) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *FrontPortTemplate) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FrontPortTemplate) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplate) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FrontPortTemplate) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *FrontPortTemplate) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o FrontPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrontPortTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["notes_url"] = o.NotesUrl
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RearPortPosition) {
		toSerialize["rear_port_position"] = o.RearPortPosition
	}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["rear_port_template"] = o.RearPortTemplate
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrontPortTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"notes_url",
		"type",
		"name",
		"rear_port_template",
		"created",
		"last_updated",
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

	varFrontPortTemplate := _FrontPortTemplate{}

	err = json.Unmarshal(data, &varFrontPortTemplate)

	if err != nil {
		return err
	}

	*o = FrontPortTemplate(varFrontPortTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "rear_port_template")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFrontPortTemplate struct {
	value *FrontPortTemplate
	isSet bool
}

func (v NullableFrontPortTemplate) Get() *FrontPortTemplate {
	return v.value
}

func (v *NullableFrontPortTemplate) Set(val *FrontPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortTemplate(val *FrontPortTemplate) *NullableFrontPortTemplate {
	return &NullableFrontPortTemplate{value: val, isSet: true}
}

func (v NullableFrontPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


