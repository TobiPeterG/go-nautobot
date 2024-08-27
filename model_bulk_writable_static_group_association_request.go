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

// checks if the BulkWritableStaticGroupAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableStaticGroupAssociationRequest{}

// BulkWritableStaticGroupAssociationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableStaticGroupAssociationRequest struct {
	Id string `json:"id"`
	AssociatedObjectType string `json:"associated_object_type"`
	AssociatedObjectId string `json:"associated_object_id"`
	DynamicGroup BulkWritableCableRequestStatus `json:"dynamic_group"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableStaticGroupAssociationRequest BulkWritableStaticGroupAssociationRequest

// NewBulkWritableStaticGroupAssociationRequest instantiates a new BulkWritableStaticGroupAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableStaticGroupAssociationRequest(id string, associatedObjectType string, associatedObjectId string, dynamicGroup BulkWritableCableRequestStatus) *BulkWritableStaticGroupAssociationRequest {
	this := BulkWritableStaticGroupAssociationRequest{}
	this.Id = id
	this.AssociatedObjectType = associatedObjectType
	this.AssociatedObjectId = associatedObjectId
	this.DynamicGroup = dynamicGroup
	return &this
}

// NewBulkWritableStaticGroupAssociationRequestWithDefaults instantiates a new BulkWritableStaticGroupAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableStaticGroupAssociationRequestWithDefaults() *BulkWritableStaticGroupAssociationRequest {
	this := BulkWritableStaticGroupAssociationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableStaticGroupAssociationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableStaticGroupAssociationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableStaticGroupAssociationRequest) SetId(v string) {
	o.Id = v
}

// GetAssociatedObjectType returns the AssociatedObjectType field value
func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociatedObjectType
}

// GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociatedObjectType, true
}

// SetAssociatedObjectType sets field value
func (o *BulkWritableStaticGroupAssociationRequest) SetAssociatedObjectType(v string) {
	o.AssociatedObjectType = v
}

// GetAssociatedObjectId returns the AssociatedObjectId field value
func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssociatedObjectId
}

// GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field value
// and a boolean to check if the value has been set.
func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociatedObjectId, true
}

// SetAssociatedObjectId sets field value
func (o *BulkWritableStaticGroupAssociationRequest) SetAssociatedObjectId(v string) {
	o.AssociatedObjectId = v
}

// GetDynamicGroup returns the DynamicGroup field value
func (o *BulkWritableStaticGroupAssociationRequest) GetDynamicGroup() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.DynamicGroup
}

// GetDynamicGroupOk returns a tuple with the DynamicGroup field value
// and a boolean to check if the value has been set.
func (o *BulkWritableStaticGroupAssociationRequest) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DynamicGroup, true
}

// SetDynamicGroup sets field value
func (o *BulkWritableStaticGroupAssociationRequest) SetDynamicGroup(v BulkWritableCableRequestStatus) {
	o.DynamicGroup = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableStaticGroupAssociationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableStaticGroupAssociationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableStaticGroupAssociationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableStaticGroupAssociationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableStaticGroupAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableStaticGroupAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableStaticGroupAssociationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableStaticGroupAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableStaticGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableStaticGroupAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["associated_object_type"] = o.AssociatedObjectType
	toSerialize["associated_object_id"] = o.AssociatedObjectId
	toSerialize["dynamic_group"] = o.DynamicGroup
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

func (o *BulkWritableStaticGroupAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"associated_object_type",
		"associated_object_id",
		"dynamic_group",
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

	varBulkWritableStaticGroupAssociationRequest := _BulkWritableStaticGroupAssociationRequest{}

	err = json.Unmarshal(data, &varBulkWritableStaticGroupAssociationRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableStaticGroupAssociationRequest(varBulkWritableStaticGroupAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "associated_object_type")
		delete(additionalProperties, "associated_object_id")
		delete(additionalProperties, "dynamic_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableStaticGroupAssociationRequest struct {
	value *BulkWritableStaticGroupAssociationRequest
	isSet bool
}

func (v NullableBulkWritableStaticGroupAssociationRequest) Get() *BulkWritableStaticGroupAssociationRequest {
	return v.value
}

func (v *NullableBulkWritableStaticGroupAssociationRequest) Set(val *BulkWritableStaticGroupAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableStaticGroupAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableStaticGroupAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableStaticGroupAssociationRequest(val *BulkWritableStaticGroupAssociationRequest) *NullableBulkWritableStaticGroupAssociationRequest {
	return &NullableBulkWritableStaticGroupAssociationRequest{value: val, isSet: true}
}

func (v NullableBulkWritableStaticGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableStaticGroupAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


