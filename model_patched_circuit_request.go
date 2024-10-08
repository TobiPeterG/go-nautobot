/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedCircuitRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCircuitRequest{}

// PatchedCircuitRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedCircuitRequest struct {
	Cid *string `json:"cid,omitempty"`
	InstallDate NullableString `json:"install_date,omitempty"`
	CommitRate NullableInt32 `json:"commit_rate,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Provider *BulkWritableCableRequestStatus `json:"provider,omitempty"`
	CircuitType *BulkWritableCableRequestStatus `json:"circuit_type,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedCircuitRequest PatchedCircuitRequest

// NewPatchedCircuitRequest instantiates a new PatchedCircuitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCircuitRequest() *PatchedCircuitRequest {
	this := PatchedCircuitRequest{}
	return &this
}

// NewPatchedCircuitRequestWithDefaults instantiates a new PatchedCircuitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCircuitRequestWithDefaults() *PatchedCircuitRequest {
	this := PatchedCircuitRequest{}
	return &this
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetCid() string {
	if o == nil || IsNil(o.Cid) {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetCidOk() (*string, bool) {
	if o == nil || IsNil(o.Cid) {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasCid() bool {
	if o != nil && !IsNil(o.Cid) {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *PatchedCircuitRequest) SetCid(v string) {
	o.Cid = &v
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitRequest) GetInstallDate() string {
	if o == nil || IsNil(o.InstallDate.Get()) {
		var ret string
		return ret
	}
	return *o.InstallDate.Get()
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitRequest) GetInstallDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallDate.Get(), o.InstallDate.IsSet()
}

// HasInstallDate returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasInstallDate() bool {
	if o != nil && o.InstallDate.IsSet() {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given NullableString and assigns it to the InstallDate field.
func (o *PatchedCircuitRequest) SetInstallDate(v string) {
	o.InstallDate.Set(&v)
}
// SetInstallDateNil sets the value for InstallDate to be an explicit nil
func (o *PatchedCircuitRequest) SetInstallDateNil() {
	o.InstallDate.Set(nil)
}

// UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
func (o *PatchedCircuitRequest) UnsetInstallDate() {
	o.InstallDate.Unset()
}

// GetCommitRate returns the CommitRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitRequest) GetCommitRate() int32 {
	if o == nil || IsNil(o.CommitRate.Get()) {
		var ret int32
		return ret
	}
	return *o.CommitRate.Get()
}

// GetCommitRateOk returns a tuple with the CommitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitRequest) GetCommitRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitRate.Get(), o.CommitRate.IsSet()
}

// HasCommitRate returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasCommitRate() bool {
	if o != nil && o.CommitRate.IsSet() {
		return true
	}

	return false
}

// SetCommitRate gets a reference to the given NullableInt32 and assigns it to the CommitRate field.
func (o *PatchedCircuitRequest) SetCommitRate(v int32) {
	o.CommitRate.Set(&v)
}
// SetCommitRateNil sets the value for CommitRate to be an explicit nil
func (o *PatchedCircuitRequest) SetCommitRateNil() {
	o.CommitRate.Set(nil)
}

// UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
func (o *PatchedCircuitRequest) UnsetCommitRate() {
	o.CommitRate.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedCircuitRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedCircuitRequest) SetComments(v string) {
	o.Comments = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedCircuitRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetProvider() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Provider) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetProviderOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Provider field.
func (o *PatchedCircuitRequest) SetProvider(v BulkWritableCableRequestStatus) {
	o.Provider = &v
}

// GetCircuitType returns the CircuitType field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetCircuitType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.CircuitType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.CircuitType
}

// GetCircuitTypeOk returns a tuple with the CircuitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetCircuitTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.CircuitType) {
		return nil, false
	}
	return o.CircuitType, true
}

// HasCircuitType returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasCircuitType() bool {
	if o != nil && !IsNil(o.CircuitType) {
		return true
	}

	return false
}

// SetCircuitType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the CircuitType field.
func (o *PatchedCircuitRequest) SetCircuitType(v BulkWritableCableRequestStatus) {
	o.CircuitType = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCircuitRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCircuitRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *PatchedCircuitRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedCircuitRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedCircuitRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedCircuitRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedCircuitRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedCircuitRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCircuitRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedCircuitRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedCircuitRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedCircuitRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCircuitRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cid) {
		toSerialize["cid"] = o.Cid
	}
	if o.InstallDate.IsSet() {
		toSerialize["install_date"] = o.InstallDate.Get()
	}
	if o.CommitRate.IsSet() {
		toSerialize["commit_rate"] = o.CommitRate.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.CircuitType) {
		toSerialize["circuit_type"] = o.CircuitType
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedCircuitRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedCircuitRequest := _PatchedCircuitRequest{}

	err = json.Unmarshal(data, &varPatchedCircuitRequest)

	if err != nil {
		return err
	}

	*o = PatchedCircuitRequest(varPatchedCircuitRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cid")
		delete(additionalProperties, "install_date")
		delete(additionalProperties, "commit_rate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "status")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "circuit_type")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedCircuitRequest struct {
	value *PatchedCircuitRequest
	isSet bool
}

func (v NullablePatchedCircuitRequest) Get() *PatchedCircuitRequest {
	return v.value
}

func (v *NullablePatchedCircuitRequest) Set(val *PatchedCircuitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCircuitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCircuitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCircuitRequest(val *PatchedCircuitRequest) *NullablePatchedCircuitRequest {
	return &NullablePatchedCircuitRequest{value: val, isSet: true}
}

func (v NullablePatchedCircuitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCircuitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


