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

// checks if the FrontPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrontPort{}

// FrontPort Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type FrontPort struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	CablePeerType NullableString `json:"cable_peer_type"`
	CablePeer NullableCableTermination `json:"cable_peer"`
	Type FrontPortType `json:"type"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	Device NullableBulkWritableCircuitRequestTenant `json:"device,omitempty"`
	Module NullableBulkWritableCircuitRequestTenant `json:"module,omitempty"`
	Cable NullableCircuitCircuitTerminationA `json:"cable"`
	RearPort BulkWritableCableRequestStatus `json:"rear_port"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FrontPort FrontPort

// NewFrontPort instantiates a new FrontPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPort(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, type_ FrontPortType, name string, cable NullableCircuitCircuitTerminationA, rearPort BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string) *FrontPort {
	this := FrontPort{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.CablePeerType = cablePeerType
	this.CablePeer = cablePeer
	this.Type = type_
	this.Name = name
	this.Cable = cable
	this.RearPort = rearPort
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewFrontPortWithDefaults instantiates a new FrontPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortWithDefaults() *FrontPort {
	this := FrontPort{}
	return &this
}

// GetId returns the Id field value
func (o *FrontPort) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FrontPort) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *FrontPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FrontPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *FrontPort) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *FrontPort) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *FrontPort) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FrontPort) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *FrontPort) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *FrontPort) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetCablePeerType returns the CablePeerType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FrontPort) GetCablePeerType() string {
	if o == nil || o.CablePeerType.Get() == nil {
		var ret string
		return ret
	}

	return *o.CablePeerType.Get()
}

// GetCablePeerTypeOk returns a tuple with the CablePeerType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetCablePeerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeerType.Get(), o.CablePeerType.IsSet()
}

// SetCablePeerType sets field value
func (o *FrontPort) SetCablePeerType(v string) {
	o.CablePeerType.Set(&v)
}

// GetCablePeer returns the CablePeer field value
// If the value is explicit nil, the zero value for CableTermination will be returned
func (o *FrontPort) GetCablePeer() CableTermination {
	if o == nil || o.CablePeer.Get() == nil {
		var ret CableTermination
		return ret
	}

	return *o.CablePeer.Get()
}

// GetCablePeerOk returns a tuple with the CablePeer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetCablePeerOk() (*CableTermination, bool) {
	if o == nil {
		return nil, false
	}
	return o.CablePeer.Get(), o.CablePeer.IsSet()
}

// SetCablePeer sets field value
func (o *FrontPort) SetCablePeer(v CableTermination) {
	o.CablePeer.Set(&v)
}

// GetType returns the Type field value
func (o *FrontPort) GetType() FrontPortType {
	if o == nil {
		var ret FrontPortType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetTypeOk() (*FrontPortType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrontPort) SetType(v FrontPortType) {
	o.Type = v
}

// GetName returns the Name field value
func (o *FrontPort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPort) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPort) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPort) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPort) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPort) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPort) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPort) SetDescription(v string) {
	o.Description = &v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *FrontPort) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *FrontPort) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *FrontPort) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPort) GetDevice() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *FrontPort) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Device field.
func (o *FrontPort) SetDevice(v BulkWritableCircuitRequestTenant) {
	o.Device.Set(&v)
}
// SetDeviceNil sets the value for Device to be an explicit nil
func (o *FrontPort) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *FrontPort) UnsetDevice() {
	o.Device.Unset()
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPort) GetModule() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetModuleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *FrontPort) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Module field.
func (o *FrontPort) SetModule(v BulkWritableCircuitRequestTenant) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *FrontPort) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *FrontPort) UnsetModule() {
	o.Module.Unset()
}

// GetCable returns the Cable field value
// If the value is explicit nil, the zero value for CircuitCircuitTerminationA will be returned
func (o *FrontPort) GetCable() CircuitCircuitTerminationA {
	if o == nil || o.Cable.Get() == nil {
		var ret CircuitCircuitTerminationA
		return ret
	}

	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetCableOk() (*CircuitCircuitTerminationA, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// SetCable sets field value
func (o *FrontPort) SetCable(v CircuitCircuitTerminationA) {
	o.Cable.Set(&v)
}

// GetRearPort returns the RearPort field value
func (o *FrontPort) GetRearPort() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetRearPortOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPort, true
}

// SetRearPort sets field value
func (o *FrontPort) SetRearPort(v BulkWritableCableRequestStatus) {
	o.RearPort = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FrontPort) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *FrontPort) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *FrontPort) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPort) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *FrontPort) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *FrontPort) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *FrontPort) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *FrontPort) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *FrontPort) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *FrontPort) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *FrontPort) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FrontPort) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPort) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FrontPort) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *FrontPort) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o FrontPort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrontPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["cable_peer_type"] = o.CablePeerType.Get()
	toSerialize["cable_peer"] = o.CablePeer.Get()
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
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	toSerialize["cable"] = o.Cable.Get()
	toSerialize["rear_port"] = o.RearPort
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrontPort) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"cable_peer_type",
		"cable_peer",
		"type",
		"name",
		"cable",
		"rear_port",
		"created",
		"last_updated",
		"notes_url",
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

	varFrontPort := _FrontPort{}

	err = json.Unmarshal(data, &varFrontPort)

	if err != nil {
		return err
	}

	*o = FrontPort(varFrontPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "cable_peer_type")
		delete(additionalProperties, "cable_peer")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "rear_port")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFrontPort struct {
	value *FrontPort
	isSet bool
}

func (v NullableFrontPort) Get() *FrontPort {
	return v.value
}

func (v *NullableFrontPort) Set(val *FrontPort) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPort) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPort(val *FrontPort) *NullableFrontPort {
	return &NullableFrontPort{value: val, isSet: true}
}

func (v NullableFrontPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


