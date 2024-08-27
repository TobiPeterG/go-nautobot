/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedCloudNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCloudNetworkRequest{}

// PatchedCloudNetworkRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedCloudNetworkRequest struct {
	ExtraConfig interface{} `json:"extra_config,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CloudResourceType *BulkWritableCableRequestStatus `json:"cloud_resource_type,omitempty"`
	CloudAccount *BulkWritableCableRequestStatus `json:"cloud_account,omitempty"`
	Parent NullableBulkWritableCircuitRequestTenant `json:"parent,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedCloudNetworkRequest PatchedCloudNetworkRequest

// NewPatchedCloudNetworkRequest instantiates a new PatchedCloudNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCloudNetworkRequest() *PatchedCloudNetworkRequest {
	this := PatchedCloudNetworkRequest{}
	return &this
}

// NewPatchedCloudNetworkRequestWithDefaults instantiates a new PatchedCloudNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCloudNetworkRequestWithDefaults() *PatchedCloudNetworkRequest {
	this := PatchedCloudNetworkRequest{}
	return &this
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCloudNetworkRequest) GetExtraConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCloudNetworkRequest) GetExtraConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExtraConfig) {
		return nil, false
	}
	return &o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasExtraConfig() bool {
	if o != nil && !IsNil(o.ExtraConfig) {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given interface{} and assigns it to the ExtraConfig field.
func (o *PatchedCloudNetworkRequest) SetExtraConfig(v interface{}) {
	o.ExtraConfig = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCloudNetworkRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedCloudNetworkRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCloudResourceType returns the CloudResourceType field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetCloudResourceType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.CloudResourceType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.CloudResourceType
}

// GetCloudResourceTypeOk returns a tuple with the CloudResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetCloudResourceTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.CloudResourceType) {
		return nil, false
	}
	return o.CloudResourceType, true
}

// HasCloudResourceType returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasCloudResourceType() bool {
	if o != nil && !IsNil(o.CloudResourceType) {
		return true
	}

	return false
}

// SetCloudResourceType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the CloudResourceType field.
func (o *PatchedCloudNetworkRequest) SetCloudResourceType(v BulkWritableCableRequestStatus) {
	o.CloudResourceType = &v
}

// GetCloudAccount returns the CloudAccount field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetCloudAccount() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.CloudAccount) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.CloudAccount
}

// GetCloudAccountOk returns a tuple with the CloudAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetCloudAccountOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.CloudAccount) {
		return nil, false
	}
	return o.CloudAccount, true
}

// HasCloudAccount returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasCloudAccount() bool {
	if o != nil && !IsNil(o.CloudAccount) {
		return true
	}

	return false
}

// SetCloudAccount gets a reference to the given BulkWritableCableRequestStatus and assigns it to the CloudAccount field.
func (o *PatchedCloudNetworkRequest) SetCloudAccount(v BulkWritableCableRequestStatus) {
	o.CloudAccount = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCloudNetworkRequest) GetParent() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCloudNetworkRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Parent field.
func (o *PatchedCloudNetworkRequest) SetParent(v BulkWritableCircuitRequestTenant) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedCloudNetworkRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedCloudNetworkRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedCloudNetworkRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedCloudNetworkRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedCloudNetworkRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCloudNetworkRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedCloudNetworkRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedCloudNetworkRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedCloudNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCloudNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.CloudAccount) {
		toSerialize["cloud_account"] = o.CloudAccount
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
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

func (o *PatchedCloudNetworkRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedCloudNetworkRequest := _PatchedCloudNetworkRequest{}

	err = json.Unmarshal(data, &varPatchedCloudNetworkRequest)

	if err != nil {
		return err
	}

	*o = PatchedCloudNetworkRequest(varPatchedCloudNetworkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extra_config")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cloud_resource_type")
		delete(additionalProperties, "cloud_account")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedCloudNetworkRequest struct {
	value *PatchedCloudNetworkRequest
	isSet bool
}

func (v NullablePatchedCloudNetworkRequest) Get() *PatchedCloudNetworkRequest {
	return v.value
}

func (v *NullablePatchedCloudNetworkRequest) Set(val *PatchedCloudNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCloudNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCloudNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCloudNetworkRequest(val *PatchedCloudNetworkRequest) *NullablePatchedCloudNetworkRequest {
	return &NullablePatchedCloudNetworkRequest{value: val, isSet: true}
}

func (v NullablePatchedCloudNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCloudNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


