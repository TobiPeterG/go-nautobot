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

// checks if the Tenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenant{}

// Tenant Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type Tenant struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	CircuitCount *int32 `json:"circuit_count,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	IpaddressCount *int32 `json:"ipaddress_count,omitempty"`
	PrefixCount *int32 `json:"prefix_count,omitempty"`
	RackCount *int32 `json:"rack_count,omitempty"`
	VirtualmachineCount *int32 `json:"virtualmachine_count,omitempty"`
	VlanCount *int32 `json:"vlan_count,omitempty"`
	VrfCount *int32 `json:"vrf_count,omitempty"`
	ClusterCount *int32 `json:"cluster_count,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	TenantGroup NullableBulkWritableCircuitRequestTenant `json:"tenant_group,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Tenant Tenant

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant(id string, objectType string, display string, url string, naturalSlug string, name string, created NullableTime, lastUpdated NullableTime, notesUrl string) *Tenant {
	this := Tenant{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Name = name
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value
func (o *Tenant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tenant) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *Tenant) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *Tenant) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *Tenant) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Tenant) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *Tenant) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Tenant) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *Tenant) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *Tenant) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetCircuitCount returns the CircuitCount field value if set, zero value otherwise.
func (o *Tenant) GetCircuitCount() int32 {
	if o == nil || IsNil(o.CircuitCount) {
		var ret int32
		return ret
	}
	return *o.CircuitCount
}

// GetCircuitCountOk returns a tuple with the CircuitCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCircuitCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CircuitCount) {
		return nil, false
	}
	return o.CircuitCount, true
}

// HasCircuitCount returns a boolean if a field has been set.
func (o *Tenant) HasCircuitCount() bool {
	if o != nil && !IsNil(o.CircuitCount) {
		return true
	}

	return false
}

// SetCircuitCount gets a reference to the given int32 and assigns it to the CircuitCount field.
func (o *Tenant) SetCircuitCount(v int32) {
	o.CircuitCount = &v
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *Tenant) GetDeviceCount() int32 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDeviceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *Tenant) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *Tenant) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetIpaddressCount returns the IpaddressCount field value if set, zero value otherwise.
func (o *Tenant) GetIpaddressCount() int32 {
	if o == nil || IsNil(o.IpaddressCount) {
		var ret int32
		return ret
	}
	return *o.IpaddressCount
}

// GetIpaddressCountOk returns a tuple with the IpaddressCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetIpaddressCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IpaddressCount) {
		return nil, false
	}
	return o.IpaddressCount, true
}

// HasIpaddressCount returns a boolean if a field has been set.
func (o *Tenant) HasIpaddressCount() bool {
	if o != nil && !IsNil(o.IpaddressCount) {
		return true
	}

	return false
}

// SetIpaddressCount gets a reference to the given int32 and assigns it to the IpaddressCount field.
func (o *Tenant) SetIpaddressCount(v int32) {
	o.IpaddressCount = &v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *Tenant) GetPrefixCount() int32 {
	if o == nil || IsNil(o.PrefixCount) {
		var ret int32
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetPrefixCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrefixCount) {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *Tenant) HasPrefixCount() bool {
	if o != nil && !IsNil(o.PrefixCount) {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int32 and assigns it to the PrefixCount field.
func (o *Tenant) SetPrefixCount(v int32) {
	o.PrefixCount = &v
}

// GetRackCount returns the RackCount field value if set, zero value otherwise.
func (o *Tenant) GetRackCount() int32 {
	if o == nil || IsNil(o.RackCount) {
		var ret int32
		return ret
	}
	return *o.RackCount
}

// GetRackCountOk returns a tuple with the RackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetRackCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RackCount) {
		return nil, false
	}
	return o.RackCount, true
}

// HasRackCount returns a boolean if a field has been set.
func (o *Tenant) HasRackCount() bool {
	if o != nil && !IsNil(o.RackCount) {
		return true
	}

	return false
}

// SetRackCount gets a reference to the given int32 and assigns it to the RackCount field.
func (o *Tenant) SetRackCount(v int32) {
	o.RackCount = &v
}

// GetVirtualmachineCount returns the VirtualmachineCount field value if set, zero value otherwise.
func (o *Tenant) GetVirtualmachineCount() int32 {
	if o == nil || IsNil(o.VirtualmachineCount) {
		var ret int32
		return ret
	}
	return *o.VirtualmachineCount
}

// GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetVirtualmachineCountOk() (*int32, bool) {
	if o == nil || IsNil(o.VirtualmachineCount) {
		return nil, false
	}
	return o.VirtualmachineCount, true
}

// HasVirtualmachineCount returns a boolean if a field has been set.
func (o *Tenant) HasVirtualmachineCount() bool {
	if o != nil && !IsNil(o.VirtualmachineCount) {
		return true
	}

	return false
}

// SetVirtualmachineCount gets a reference to the given int32 and assigns it to the VirtualmachineCount field.
func (o *Tenant) SetVirtualmachineCount(v int32) {
	o.VirtualmachineCount = &v
}

// GetVlanCount returns the VlanCount field value if set, zero value otherwise.
func (o *Tenant) GetVlanCount() int32 {
	if o == nil || IsNil(o.VlanCount) {
		var ret int32
		return ret
	}
	return *o.VlanCount
}

// GetVlanCountOk returns a tuple with the VlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetVlanCountOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanCount) {
		return nil, false
	}
	return o.VlanCount, true
}

// HasVlanCount returns a boolean if a field has been set.
func (o *Tenant) HasVlanCount() bool {
	if o != nil && !IsNil(o.VlanCount) {
		return true
	}

	return false
}

// SetVlanCount gets a reference to the given int32 and assigns it to the VlanCount field.
func (o *Tenant) SetVlanCount(v int32) {
	o.VlanCount = &v
}

// GetVrfCount returns the VrfCount field value if set, zero value otherwise.
func (o *Tenant) GetVrfCount() int32 {
	if o == nil || IsNil(o.VrfCount) {
		var ret int32
		return ret
	}
	return *o.VrfCount
}

// GetVrfCountOk returns a tuple with the VrfCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetVrfCountOk() (*int32, bool) {
	if o == nil || IsNil(o.VrfCount) {
		return nil, false
	}
	return o.VrfCount, true
}

// HasVrfCount returns a boolean if a field has been set.
func (o *Tenant) HasVrfCount() bool {
	if o != nil && !IsNil(o.VrfCount) {
		return true
	}

	return false
}

// SetVrfCount gets a reference to the given int32 and assigns it to the VrfCount field.
func (o *Tenant) SetVrfCount(v int32) {
	o.VrfCount = &v
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *Tenant) GetClusterCount() int32 {
	if o == nil || IsNil(o.ClusterCount) {
		var ret int32
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetClusterCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ClusterCount) {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *Tenant) HasClusterCount() bool {
	if o != nil && !IsNil(o.ClusterCount) {
		return true
	}

	return false
}

// SetClusterCount gets a reference to the given int32 and assigns it to the ClusterCount field.
func (o *Tenant) SetClusterCount(v int32) {
	o.ClusterCount = &v
}

// GetName returns the Name field value
func (o *Tenant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tenant) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tenant) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tenant) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tenant) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Tenant) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Tenant) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Tenant) SetComments(v string) {
	o.Comments = &v
}

// GetTenantGroup returns the TenantGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tenant) GetTenantGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.TenantGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.TenantGroup.Get()
}

// GetTenantGroupOk returns a tuple with the TenantGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetTenantGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.TenantGroup.Get(), o.TenantGroup.IsSet()
}

// HasTenantGroup returns a boolean if a field has been set.
func (o *Tenant) HasTenantGroup() bool {
	if o != nil && o.TenantGroup.IsSet() {
		return true
	}

	return false
}

// SetTenantGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the TenantGroup field.
func (o *Tenant) SetTenantGroup(v BulkWritableCircuitRequestTenant) {
	o.TenantGroup.Set(&v)
}
// SetTenantGroupNil sets the value for TenantGroup to be an explicit nil
func (o *Tenant) SetTenantGroupNil() {
	o.TenantGroup.Set(nil)
}

// UnsetTenantGroup ensures that no value is present for TenantGroup, not even an explicit nil
func (o *Tenant) UnsetTenantGroup() {
	o.TenantGroup.Unset()
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tenant) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Tenant) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tenant) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tenant) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Tenant) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Tenant) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Tenant) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *Tenant) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetNotesUrl returns the NotesUrl field value
func (o *Tenant) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *Tenant) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *Tenant) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Tenant) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Tenant) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Tenant) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	if !IsNil(o.CircuitCount) {
		toSerialize["circuit_count"] = o.CircuitCount
	}
	if !IsNil(o.DeviceCount) {
		toSerialize["device_count"] = o.DeviceCount
	}
	if !IsNil(o.IpaddressCount) {
		toSerialize["ipaddress_count"] = o.IpaddressCount
	}
	if !IsNil(o.PrefixCount) {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	if !IsNil(o.RackCount) {
		toSerialize["rack_count"] = o.RackCount
	}
	if !IsNil(o.VirtualmachineCount) {
		toSerialize["virtualmachine_count"] = o.VirtualmachineCount
	}
	if !IsNil(o.VlanCount) {
		toSerialize["vlan_count"] = o.VlanCount
	}
	if !IsNil(o.VrfCount) {
		toSerialize["vrf_count"] = o.VrfCount
	}
	if !IsNil(o.ClusterCount) {
		toSerialize["cluster_count"] = o.ClusterCount
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.TenantGroup.IsSet() {
		toSerialize["tenant_group"] = o.TenantGroup.Get()
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tenant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"name",
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

	varTenant := _Tenant{}

	err = json.Unmarshal(data, &varTenant)

	if err != nil {
		return err
	}

	*o = Tenant(varTenant)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "circuit_count")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "ipaddress_count")
		delete(additionalProperties, "prefix_count")
		delete(additionalProperties, "rack_count")
		delete(additionalProperties, "virtualmachine_count")
		delete(additionalProperties, "vlan_count")
		delete(additionalProperties, "vrf_count")
		delete(additionalProperties, "cluster_count")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tenant_group")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenant struct {
	value *Tenant
	isSet bool
}

func (v NullableTenant) Get() *Tenant {
	return v.value
}

func (v *NullableTenant) Set(val *Tenant) {
	v.value = val
	v.isSet = true
}

func (v NullableTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenant(val *Tenant) *NullableTenant {
	return &NullableTenant{value: val, isSet: true}
}

func (v NullableTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


