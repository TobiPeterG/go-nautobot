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

// checks if the PatchedBulkWritableJobHookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableJobHookRequest{}

// PatchedBulkWritableJobHookRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableJobHookRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Name *string `json:"name,omitempty"`
	// Call this job hook when a matching object is created.
	TypeCreate *bool `json:"type_create,omitempty"`
	// Call this job hook when a matching object is deleted.
	TypeDelete *bool `json:"type_delete,omitempty"`
	// Call this job hook when a matching object is updated.
	TypeUpdate *bool `json:"type_update,omitempty"`
	Job *BulkWritableJobHookRequestJob `json:"job,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableJobHookRequest PatchedBulkWritableJobHookRequest

// NewPatchedBulkWritableJobHookRequest instantiates a new PatchedBulkWritableJobHookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableJobHookRequest(id string) *PatchedBulkWritableJobHookRequest {
	this := PatchedBulkWritableJobHookRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableJobHookRequestWithDefaults instantiates a new PatchedBulkWritableJobHookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableJobHookRequestWithDefaults() *PatchedBulkWritableJobHookRequest {
	this := PatchedBulkWritableJobHookRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableJobHookRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableJobHookRequest) SetId(v string) {
	o.Id = v
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *PatchedBulkWritableJobHookRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedBulkWritableJobHookRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableJobHookRequest) SetName(v string) {
	o.Name = &v
}

// GetTypeCreate returns the TypeCreate field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetTypeCreate() bool {
	if o == nil || IsNil(o.TypeCreate) {
		var ret bool
		return ret
	}
	return *o.TypeCreate
}

// GetTypeCreateOk returns a tuple with the TypeCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetTypeCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeCreate) {
		return nil, false
	}
	return o.TypeCreate, true
}

// HasTypeCreate returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasTypeCreate() bool {
	if o != nil && !IsNil(o.TypeCreate) {
		return true
	}

	return false
}

// SetTypeCreate gets a reference to the given bool and assigns it to the TypeCreate field.
func (o *PatchedBulkWritableJobHookRequest) SetTypeCreate(v bool) {
	o.TypeCreate = &v
}

// GetTypeDelete returns the TypeDelete field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetTypeDelete() bool {
	if o == nil || IsNil(o.TypeDelete) {
		var ret bool
		return ret
	}
	return *o.TypeDelete
}

// GetTypeDeleteOk returns a tuple with the TypeDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetTypeDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeDelete) {
		return nil, false
	}
	return o.TypeDelete, true
}

// HasTypeDelete returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasTypeDelete() bool {
	if o != nil && !IsNil(o.TypeDelete) {
		return true
	}

	return false
}

// SetTypeDelete gets a reference to the given bool and assigns it to the TypeDelete field.
func (o *PatchedBulkWritableJobHookRequest) SetTypeDelete(v bool) {
	o.TypeDelete = &v
}

// GetTypeUpdate returns the TypeUpdate field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetTypeUpdate() bool {
	if o == nil || IsNil(o.TypeUpdate) {
		var ret bool
		return ret
	}
	return *o.TypeUpdate
}

// GetTypeUpdateOk returns a tuple with the TypeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetTypeUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeUpdate) {
		return nil, false
	}
	return o.TypeUpdate, true
}

// HasTypeUpdate returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasTypeUpdate() bool {
	if o != nil && !IsNil(o.TypeUpdate) {
		return true
	}

	return false
}

// SetTypeUpdate gets a reference to the given bool and assigns it to the TypeUpdate field.
func (o *PatchedBulkWritableJobHookRequest) SetTypeUpdate(v bool) {
	o.TypeUpdate = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetJob() BulkWritableJobHookRequestJob {
	if o == nil || IsNil(o.Job) {
		var ret BulkWritableJobHookRequestJob
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetJobOk() (*BulkWritableJobHookRequestJob, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given BulkWritableJobHookRequestJob and assigns it to the Job field.
func (o *PatchedBulkWritableJobHookRequest) SetJob(v BulkWritableJobHookRequestJob) {
	o.Job = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableJobHookRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableJobHookRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableJobHookRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableJobHookRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableJobHookRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritableJobHookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableJobHookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ContentTypes) {
		toSerialize["content_types"] = o.ContentTypes
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TypeCreate) {
		toSerialize["type_create"] = o.TypeCreate
	}
	if !IsNil(o.TypeDelete) {
		toSerialize["type_delete"] = o.TypeDelete
	}
	if !IsNil(o.TypeUpdate) {
		toSerialize["type_update"] = o.TypeUpdate
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
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

func (o *PatchedBulkWritableJobHookRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableJobHookRequest := _PatchedBulkWritableJobHookRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableJobHookRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableJobHookRequest(varPatchedBulkWritableJobHookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type_create")
		delete(additionalProperties, "type_delete")
		delete(additionalProperties, "type_update")
		delete(additionalProperties, "job")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableJobHookRequest struct {
	value *PatchedBulkWritableJobHookRequest
	isSet bool
}

func (v NullablePatchedBulkWritableJobHookRequest) Get() *PatchedBulkWritableJobHookRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableJobHookRequest) Set(val *PatchedBulkWritableJobHookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableJobHookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableJobHookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableJobHookRequest(val *PatchedBulkWritableJobHookRequest) *NullablePatchedBulkWritableJobHookRequest {
	return &NullablePatchedBulkWritableJobHookRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableJobHookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableJobHookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


