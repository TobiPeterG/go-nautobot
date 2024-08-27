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

// checks if the InterfaceRedundancyGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceRedundancyGroup{}

// InterfaceRedundancyGroup InterfaceRedundancyGroup Serializer.
type InterfaceRedundancyGroup struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Protocol InterfaceRedundancyGroupProtocol `json:"protocol"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ProtocolGroupId *string `json:"protocol_group_id,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	VirtualIp NullableBulkWritableCircuitRequestTenant `json:"virtual_ip,omitempty"`
	Interfaces []BulkWritableCableRequestStatus `json:"interfaces"`
	Created time.Time `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceRedundancyGroup InterfaceRedundancyGroup

// NewInterfaceRedundancyGroup instantiates a new InterfaceRedundancyGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceRedundancyGroup(id string, objectType string, display string, url string, naturalSlug string, protocol InterfaceRedundancyGroupProtocol, name string, status BulkWritableCableRequestStatus, interfaces []BulkWritableCableRequestStatus, created time.Time, lastUpdated NullableTime, notesUrl string) *InterfaceRedundancyGroup {
	this := InterfaceRedundancyGroup{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Protocol = protocol
	this.Name = name
	this.Status = status
	this.Interfaces = interfaces
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewInterfaceRedundancyGroupWithDefaults instantiates a new InterfaceRedundancyGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceRedundancyGroupWithDefaults() *InterfaceRedundancyGroup {
	this := InterfaceRedundancyGroup{}
	return &this
}

// GetId returns the Id field value
func (o *InterfaceRedundancyGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InterfaceRedundancyGroup) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *InterfaceRedundancyGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InterfaceRedundancyGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *InterfaceRedundancyGroup) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InterfaceRedundancyGroup) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *InterfaceRedundancyGroup) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InterfaceRedundancyGroup) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *InterfaceRedundancyGroup) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *InterfaceRedundancyGroup) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetProtocol returns the Protocol field value
func (o *InterfaceRedundancyGroup) GetProtocol() InterfaceRedundancyGroupProtocol {
	if o == nil {
		var ret InterfaceRedundancyGroupProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetProtocolOk() (*InterfaceRedundancyGroupProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *InterfaceRedundancyGroup) SetProtocol(v InterfaceRedundancyGroupProtocol) {
	o.Protocol = v
}

// GetName returns the Name field value
func (o *InterfaceRedundancyGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InterfaceRedundancyGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceRedundancyGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceRedundancyGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceRedundancyGroup) SetDescription(v string) {
	o.Description = &v
}

// GetProtocolGroupId returns the ProtocolGroupId field value if set, zero value otherwise.
func (o *InterfaceRedundancyGroup) GetProtocolGroupId() string {
	if o == nil || IsNil(o.ProtocolGroupId) {
		var ret string
		return ret
	}
	return *o.ProtocolGroupId
}

// GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetProtocolGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolGroupId) {
		return nil, false
	}
	return o.ProtocolGroupId, true
}

// HasProtocolGroupId returns a boolean if a field has been set.
func (o *InterfaceRedundancyGroup) HasProtocolGroupId() bool {
	if o != nil && !IsNil(o.ProtocolGroupId) {
		return true
	}

	return false
}

// SetProtocolGroupId gets a reference to the given string and assigns it to the ProtocolGroupId field.
func (o *InterfaceRedundancyGroup) SetProtocolGroupId(v string) {
	o.ProtocolGroupId = &v
}

// GetStatus returns the Status field value
func (o *InterfaceRedundancyGroup) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InterfaceRedundancyGroup) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceRedundancyGroup) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceRedundancyGroup) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *InterfaceRedundancyGroup) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *InterfaceRedundancyGroup) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *InterfaceRedundancyGroup) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *InterfaceRedundancyGroup) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetVirtualIp returns the VirtualIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceRedundancyGroup) GetVirtualIp() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.VirtualIp.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.VirtualIp.Get()
}

// GetVirtualIpOk returns a tuple with the VirtualIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceRedundancyGroup) GetVirtualIpOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualIp.Get(), o.VirtualIp.IsSet()
}

// HasVirtualIp returns a boolean if a field has been set.
func (o *InterfaceRedundancyGroup) HasVirtualIp() bool {
	if o != nil && o.VirtualIp.IsSet() {
		return true
	}

	return false
}

// SetVirtualIp gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the VirtualIp field.
func (o *InterfaceRedundancyGroup) SetVirtualIp(v BulkWritableCircuitRequestTenant) {
	o.VirtualIp.Set(&v)
}
// SetVirtualIpNil sets the value for VirtualIp to be an explicit nil
func (o *InterfaceRedundancyGroup) SetVirtualIpNil() {
	o.VirtualIp.Set(nil)
}

// UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil
func (o *InterfaceRedundancyGroup) UnsetVirtualIp() {
	o.VirtualIp.Unset()
}

// GetInterfaces returns the Interfaces field value
func (o *InterfaceRedundancyGroup) GetInterfaces() []BulkWritableCableRequestStatus {
	if o == nil {
		var ret []BulkWritableCableRequestStatus
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetInterfacesOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// SetInterfaces sets field value
func (o *InterfaceRedundancyGroup) SetInterfaces(v []BulkWritableCableRequestStatus) {
	o.Interfaces = v
}

// GetCreated returns the Created field value
func (o *InterfaceRedundancyGroup) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *InterfaceRedundancyGroup) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InterfaceRedundancyGroup) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceRedundancyGroup) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *InterfaceRedundancyGroup) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *InterfaceRedundancyGroup) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *InterfaceRedundancyGroup) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InterfaceRedundancyGroup) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InterfaceRedundancyGroup) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *InterfaceRedundancyGroup) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InterfaceRedundancyGroup) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceRedundancyGroup) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InterfaceRedundancyGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *InterfaceRedundancyGroup) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o InterfaceRedundancyGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceRedundancyGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["protocol"] = o.Protocol
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProtocolGroupId) {
		toSerialize["protocol_group_id"] = o.ProtocolGroupId
	}
	toSerialize["status"] = o.Status
	if o.SecretsGroup.IsSet() {
		toSerialize["secrets_group"] = o.SecretsGroup.Get()
	}
	if o.VirtualIp.IsSet() {
		toSerialize["virtual_ip"] = o.VirtualIp.Get()
	}
	toSerialize["interfaces"] = o.Interfaces
	toSerialize["created"] = o.Created
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

func (o *InterfaceRedundancyGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"protocol",
		"name",
		"status",
		"interfaces",
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

	varInterfaceRedundancyGroup := _InterfaceRedundancyGroup{}

	err = json.Unmarshal(data, &varInterfaceRedundancyGroup)

	if err != nil {
		return err
	}

	*o = InterfaceRedundancyGroup(varInterfaceRedundancyGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "protocol_group_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "virtual_ip")
		delete(additionalProperties, "interfaces")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceRedundancyGroup struct {
	value *InterfaceRedundancyGroup
	isSet bool
}

func (v NullableInterfaceRedundancyGroup) Get() *InterfaceRedundancyGroup {
	return v.value
}

func (v *NullableInterfaceRedundancyGroup) Set(val *InterfaceRedundancyGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRedundancyGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRedundancyGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRedundancyGroup(val *InterfaceRedundancyGroup) *NullableInterfaceRedundancyGroup {
	return &NullableInterfaceRedundancyGroup{value: val, isSet: true}
}

func (v NullableInterfaceRedundancyGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRedundancyGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


