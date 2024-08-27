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

// checks if the BulkWritableModuleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableModuleRequest{}

// BulkWritableModuleRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableModuleRequest struct {
	Id string `json:"id"`
	Serial NullableString `json:"serial,omitempty"`
	// A unique tag used to identify this module
	AssetTag NullableString `json:"asset_tag,omitempty"`
	ModuleType BulkWritableCableRequestStatus `json:"module_type"`
	ParentModuleBay NullableBulkWritableCircuitRequestTenant `json:"parent_module_bay,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Location NullableBulkWritableCircuitRequestTenant `json:"location,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableModuleRequest BulkWritableModuleRequest

// NewBulkWritableModuleRequest instantiates a new BulkWritableModuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableModuleRequest(id string, moduleType BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus) *BulkWritableModuleRequest {
	this := BulkWritableModuleRequest{}
	this.Id = id
	this.ModuleType = moduleType
	this.Status = status
	return &this
}

// NewBulkWritableModuleRequestWithDefaults instantiates a new BulkWritableModuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableModuleRequestWithDefaults() *BulkWritableModuleRequest {
	this := BulkWritableModuleRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableModuleRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableModuleRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableModuleRequest) SetId(v string) {
	o.Id = v
}

// GetSerial returns the Serial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableModuleRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial.Get()) {
		var ret string
		return ret
	}
	return *o.Serial.Get()
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableModuleRequest) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serial.Get(), o.Serial.IsSet()
}

// HasSerial returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasSerial() bool {
	if o != nil && o.Serial.IsSet() {
		return true
	}

	return false
}

// SetSerial gets a reference to the given NullableString and assigns it to the Serial field.
func (o *BulkWritableModuleRequest) SetSerial(v string) {
	o.Serial.Set(&v)
}
// SetSerialNil sets the value for Serial to be an explicit nil
func (o *BulkWritableModuleRequest) SetSerialNil() {
	o.Serial.Set(nil)
}

// UnsetSerial ensures that no value is present for Serial, not even an explicit nil
func (o *BulkWritableModuleRequest) UnsetSerial() {
	o.Serial.Unset()
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableModuleRequest) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableModuleRequest) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *BulkWritableModuleRequest) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}
// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *BulkWritableModuleRequest) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *BulkWritableModuleRequest) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetModuleType returns the ModuleType field value
func (o *BulkWritableModuleRequest) GetModuleType() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.ModuleType
}

// GetModuleTypeOk returns a tuple with the ModuleType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableModuleRequest) GetModuleTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleType, true
}

// SetModuleType sets field value
func (o *BulkWritableModuleRequest) SetModuleType(v BulkWritableCableRequestStatus) {
	o.ModuleType = v
}

// GetParentModuleBay returns the ParentModuleBay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableModuleRequest) GetParentModuleBay() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ParentModuleBay.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ParentModuleBay.Get()
}

// GetParentModuleBayOk returns a tuple with the ParentModuleBay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableModuleRequest) GetParentModuleBayOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentModuleBay.Get(), o.ParentModuleBay.IsSet()
}

// HasParentModuleBay returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasParentModuleBay() bool {
	if o != nil && o.ParentModuleBay.IsSet() {
		return true
	}

	return false
}

// SetParentModuleBay gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ParentModuleBay field.
func (o *BulkWritableModuleRequest) SetParentModuleBay(v BulkWritableCircuitRequestTenant) {
	o.ParentModuleBay.Set(&v)
}
// SetParentModuleBayNil sets the value for ParentModuleBay to be an explicit nil
func (o *BulkWritableModuleRequest) SetParentModuleBayNil() {
	o.ParentModuleBay.Set(nil)
}

// UnsetParentModuleBay ensures that no value is present for ParentModuleBay, not even an explicit nil
func (o *BulkWritableModuleRequest) UnsetParentModuleBay() {
	o.ParentModuleBay.Unset()
}

// GetStatus returns the Status field value
func (o *BulkWritableModuleRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableModuleRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableModuleRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableModuleRequest) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableModuleRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *BulkWritableModuleRequest) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *BulkWritableModuleRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *BulkWritableModuleRequest) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableModuleRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableModuleRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *BulkWritableModuleRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *BulkWritableModuleRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *BulkWritableModuleRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableModuleRequest) GetLocation() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Location.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableModuleRequest) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Location field.
func (o *BulkWritableModuleRequest) SetLocation(v BulkWritableCircuitRequestTenant) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *BulkWritableModuleRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *BulkWritableModuleRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableModuleRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableModuleRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableModuleRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableModuleRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableModuleRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableModuleRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableModuleRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableModuleRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableModuleRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableModuleRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableModuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableModuleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Serial.IsSet() {
		toSerialize["serial"] = o.Serial.Get()
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	toSerialize["module_type"] = o.ModuleType
	if o.ParentModuleBay.IsSet() {
		toSerialize["parent_module_bay"] = o.ParentModuleBay.Get()
	}
	toSerialize["status"] = o.Status
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
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

func (o *BulkWritableModuleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"module_type",
		"status",
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

	varBulkWritableModuleRequest := _BulkWritableModuleRequest{}

	err = json.Unmarshal(data, &varBulkWritableModuleRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableModuleRequest(varBulkWritableModuleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "parent_module_bay")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "location")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableModuleRequest struct {
	value *BulkWritableModuleRequest
	isSet bool
}

func (v NullableBulkWritableModuleRequest) Get() *BulkWritableModuleRequest {
	return v.value
}

func (v *NullableBulkWritableModuleRequest) Set(val *BulkWritableModuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableModuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableModuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableModuleRequest(val *BulkWritableModuleRequest) *NullableBulkWritableModuleRequest {
	return &NullableBulkWritableModuleRequest{value: val, isSet: true}
}

func (v NullableBulkWritableModuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableModuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


