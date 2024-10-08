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

// checks if the PatchedStaticGroupAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedStaticGroupAssociationRequest{}

// PatchedStaticGroupAssociationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedStaticGroupAssociationRequest struct {
	AssociatedObjectType *string `json:"associated_object_type,omitempty"`
	AssociatedObjectId *string `json:"associated_object_id,omitempty"`
	DynamicGroup *BulkWritableCableRequestStatus `json:"dynamic_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedStaticGroupAssociationRequest PatchedStaticGroupAssociationRequest

// NewPatchedStaticGroupAssociationRequest instantiates a new PatchedStaticGroupAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedStaticGroupAssociationRequest() *PatchedStaticGroupAssociationRequest {
	this := PatchedStaticGroupAssociationRequest{}
	return &this
}

// NewPatchedStaticGroupAssociationRequestWithDefaults instantiates a new PatchedStaticGroupAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedStaticGroupAssociationRequestWithDefaults() *PatchedStaticGroupAssociationRequest {
	this := PatchedStaticGroupAssociationRequest{}
	return &this
}

// GetAssociatedObjectType returns the AssociatedObjectType field value if set, zero value otherwise.
func (o *PatchedStaticGroupAssociationRequest) GetAssociatedObjectType() string {
	if o == nil || IsNil(o.AssociatedObjectType) {
		var ret string
		return ret
	}
	return *o.AssociatedObjectType
}

// GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStaticGroupAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedObjectType) {
		return nil, false
	}
	return o.AssociatedObjectType, true
}

// HasAssociatedObjectType returns a boolean if a field has been set.
func (o *PatchedStaticGroupAssociationRequest) HasAssociatedObjectType() bool {
	if o != nil && !IsNil(o.AssociatedObjectType) {
		return true
	}

	return false
}

// SetAssociatedObjectType gets a reference to the given string and assigns it to the AssociatedObjectType field.
func (o *PatchedStaticGroupAssociationRequest) SetAssociatedObjectType(v string) {
	o.AssociatedObjectType = &v
}

// GetAssociatedObjectId returns the AssociatedObjectId field value if set, zero value otherwise.
func (o *PatchedStaticGroupAssociationRequest) GetAssociatedObjectId() string {
	if o == nil || IsNil(o.AssociatedObjectId) {
		var ret string
		return ret
	}
	return *o.AssociatedObjectId
}

// GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStaticGroupAssociationRequest) GetAssociatedObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedObjectId) {
		return nil, false
	}
	return o.AssociatedObjectId, true
}

// HasAssociatedObjectId returns a boolean if a field has been set.
func (o *PatchedStaticGroupAssociationRequest) HasAssociatedObjectId() bool {
	if o != nil && !IsNil(o.AssociatedObjectId) {
		return true
	}

	return false
}

// SetAssociatedObjectId gets a reference to the given string and assigns it to the AssociatedObjectId field.
func (o *PatchedStaticGroupAssociationRequest) SetAssociatedObjectId(v string) {
	o.AssociatedObjectId = &v
}

// GetDynamicGroup returns the DynamicGroup field value if set, zero value otherwise.
func (o *PatchedStaticGroupAssociationRequest) GetDynamicGroup() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.DynamicGroup) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.DynamicGroup
}

// GetDynamicGroupOk returns a tuple with the DynamicGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStaticGroupAssociationRequest) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.DynamicGroup) {
		return nil, false
	}
	return o.DynamicGroup, true
}

// HasDynamicGroup returns a boolean if a field has been set.
func (o *PatchedStaticGroupAssociationRequest) HasDynamicGroup() bool {
	if o != nil && !IsNil(o.DynamicGroup) {
		return true
	}

	return false
}

// SetDynamicGroup gets a reference to the given BulkWritableCableRequestStatus and assigns it to the DynamicGroup field.
func (o *PatchedStaticGroupAssociationRequest) SetDynamicGroup(v BulkWritableCableRequestStatus) {
	o.DynamicGroup = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedStaticGroupAssociationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStaticGroupAssociationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedStaticGroupAssociationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedStaticGroupAssociationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedStaticGroupAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStaticGroupAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedStaticGroupAssociationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedStaticGroupAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedStaticGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedStaticGroupAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedObjectType) {
		toSerialize["associated_object_type"] = o.AssociatedObjectType
	}
	if !IsNil(o.AssociatedObjectId) {
		toSerialize["associated_object_id"] = o.AssociatedObjectId
	}
	if !IsNil(o.DynamicGroup) {
		toSerialize["dynamic_group"] = o.DynamicGroup
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

func (o *PatchedStaticGroupAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedStaticGroupAssociationRequest := _PatchedStaticGroupAssociationRequest{}

	err = json.Unmarshal(data, &varPatchedStaticGroupAssociationRequest)

	if err != nil {
		return err
	}

	*o = PatchedStaticGroupAssociationRequest(varPatchedStaticGroupAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associated_object_type")
		delete(additionalProperties, "associated_object_id")
		delete(additionalProperties, "dynamic_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedStaticGroupAssociationRequest struct {
	value *PatchedStaticGroupAssociationRequest
	isSet bool
}

func (v NullablePatchedStaticGroupAssociationRequest) Get() *PatchedStaticGroupAssociationRequest {
	return v.value
}

func (v *NullablePatchedStaticGroupAssociationRequest) Set(val *PatchedStaticGroupAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedStaticGroupAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedStaticGroupAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedStaticGroupAssociationRequest(val *PatchedStaticGroupAssociationRequest) *NullablePatchedStaticGroupAssociationRequest {
	return &NullablePatchedStaticGroupAssociationRequest{value: val, isSet: true}
}

func (v NullablePatchedStaticGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedStaticGroupAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


