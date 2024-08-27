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

// checks if the PatchedBulkWritableCloudServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableCloudServiceRequest{}

// PatchedBulkWritableCloudServiceRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableCloudServiceRequest struct {
	Id string `json:"id"`
	ExtraConfig interface{} `json:"extra_config,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CloudResourceType *BulkWritableCableRequestStatus `json:"cloud_resource_type,omitempty"`
	CloudAccount NullableBulkWritableCircuitRequestTenant `json:"cloud_account,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableCloudServiceRequest PatchedBulkWritableCloudServiceRequest

// NewPatchedBulkWritableCloudServiceRequest instantiates a new PatchedBulkWritableCloudServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableCloudServiceRequest(id string) *PatchedBulkWritableCloudServiceRequest {
	this := PatchedBulkWritableCloudServiceRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableCloudServiceRequestWithDefaults instantiates a new PatchedBulkWritableCloudServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableCloudServiceRequestWithDefaults() *PatchedBulkWritableCloudServiceRequest {
	this := PatchedBulkWritableCloudServiceRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableCloudServiceRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableCloudServiceRequest) SetId(v string) {
	o.Id = v
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCloudServiceRequest) GetExtraConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCloudServiceRequest) GetExtraConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExtraConfig) {
		return nil, false
	}
	return &o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasExtraConfig() bool {
	if o != nil && !IsNil(o.ExtraConfig) {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given interface{} and assigns it to the ExtraConfig field.
func (o *PatchedBulkWritableCloudServiceRequest) SetExtraConfig(v interface{}) {
	o.ExtraConfig = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudServiceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableCloudServiceRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudServiceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableCloudServiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCloudResourceType returns the CloudResourceType field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudServiceRequest) GetCloudResourceType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.CloudResourceType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.CloudResourceType
}

// GetCloudResourceTypeOk returns a tuple with the CloudResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetCloudResourceTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.CloudResourceType) {
		return nil, false
	}
	return o.CloudResourceType, true
}

// HasCloudResourceType returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasCloudResourceType() bool {
	if o != nil && !IsNil(o.CloudResourceType) {
		return true
	}

	return false
}

// SetCloudResourceType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the CloudResourceType field.
func (o *PatchedBulkWritableCloudServiceRequest) SetCloudResourceType(v BulkWritableCableRequestStatus) {
	o.CloudResourceType = &v
}

// GetCloudAccount returns the CloudAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableCloudServiceRequest) GetCloudAccount() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.CloudAccount.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.CloudAccount.Get()
}

// GetCloudAccountOk returns a tuple with the CloudAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableCloudServiceRequest) GetCloudAccountOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudAccount.Get(), o.CloudAccount.IsSet()
}

// HasCloudAccount returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasCloudAccount() bool {
	if o != nil && o.CloudAccount.IsSet() {
		return true
	}

	return false
}

// SetCloudAccount gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the CloudAccount field.
func (o *PatchedBulkWritableCloudServiceRequest) SetCloudAccount(v BulkWritableCircuitRequestTenant) {
	o.CloudAccount.Set(&v)
}
// SetCloudAccountNil sets the value for CloudAccount to be an explicit nil
func (o *PatchedBulkWritableCloudServiceRequest) SetCloudAccountNil() {
	o.CloudAccount.Set(nil)
}

// UnsetCloudAccount ensures that no value is present for CloudAccount, not even an explicit nil
func (o *PatchedBulkWritableCloudServiceRequest) UnsetCloudAccount() {
	o.CloudAccount.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudServiceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableCloudServiceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudServiceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableCloudServiceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableCloudServiceRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableCloudServiceRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableCloudServiceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableCloudServiceRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableCloudServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableCloudServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.ExtraConfig != nil {
		toSerialize["extra_config"] = o.ExtraConfig
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CloudResourceType) {
		toSerialize["cloud_resource_type"] = o.CloudResourceType
	}
	if o.CloudAccount.IsSet() {
		toSerialize["cloud_account"] = o.CloudAccount.Get()
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

func (o *PatchedBulkWritableCloudServiceRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableCloudServiceRequest := _PatchedBulkWritableCloudServiceRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableCloudServiceRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableCloudServiceRequest(varPatchedBulkWritableCloudServiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "extra_config")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cloud_resource_type")
		delete(additionalProperties, "cloud_account")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableCloudServiceRequest struct {
	value *PatchedBulkWritableCloudServiceRequest
	isSet bool
}

func (v NullablePatchedBulkWritableCloudServiceRequest) Get() *PatchedBulkWritableCloudServiceRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableCloudServiceRequest) Set(val *PatchedBulkWritableCloudServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableCloudServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableCloudServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableCloudServiceRequest(val *PatchedBulkWritableCloudServiceRequest) *NullablePatchedBulkWritableCloudServiceRequest {
	return &NullablePatchedBulkWritableCloudServiceRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableCloudServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableCloudServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


