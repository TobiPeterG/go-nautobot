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

// checks if the PatchedBulkWritableNoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableNoteRequest{}

// PatchedBulkWritableNoteRequest This base serializer implements common fields and logic for all ModelSerializers.  Namely, it:  - defines the `display` field which exposes a human friendly value for the given object. - ensures that `id` field is always present on the serializer as well. - ensures that `created` and `last_updated` fields are always present if applicable to this model and serializer. - ensures that `object_type` field is always present on the serializer which represents the content-type of this   serializer's associated model (e.g. \"dcim.device\"). This is required as the OpenAPI schema, using the   PolymorphicProxySerializer class defined below, relies upon this field as a way to identify to the client   which of several possible serializers are in use for a given attribute. - supports `?depth` query parameter. It is passed in as `nested_depth` to the `build_nested_field()` function   to enable the dynamic generation of nested serializers.
type PatchedBulkWritableNoteRequest struct {
	Id string `json:"id"`
	AssignedObjectType *string `json:"assigned_object_type,omitempty"`
	AssignedObjectId *string `json:"assigned_object_id,omitempty"`
	Note *string `json:"note,omitempty"`
	User NullableBulkWritableCircuitRequestTenant `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableNoteRequest PatchedBulkWritableNoteRequest

// NewPatchedBulkWritableNoteRequest instantiates a new PatchedBulkWritableNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableNoteRequest(id string) *PatchedBulkWritableNoteRequest {
	this := PatchedBulkWritableNoteRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableNoteRequestWithDefaults instantiates a new PatchedBulkWritableNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableNoteRequestWithDefaults() *PatchedBulkWritableNoteRequest {
	this := PatchedBulkWritableNoteRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableNoteRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableNoteRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableNoteRequest) SetId(v string) {
	o.Id = v
}

// GetAssignedObjectType returns the AssignedObjectType field value if set, zero value otherwise.
func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectType() string {
	if o == nil || IsNil(o.AssignedObjectType) {
		var ret string
		return ret
	}
	return *o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedObjectType) {
		return nil, false
	}
	return o.AssignedObjectType, true
}

// HasAssignedObjectType returns a boolean if a field has been set.
func (o *PatchedBulkWritableNoteRequest) HasAssignedObjectType() bool {
	if o != nil && !IsNil(o.AssignedObjectType) {
		return true
	}

	return false
}

// SetAssignedObjectType gets a reference to the given string and assigns it to the AssignedObjectType field.
func (o *PatchedBulkWritableNoteRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = &v
}

// GetAssignedObjectId returns the AssignedObjectId field value if set, zero value otherwise.
func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectId() string {
	if o == nil || IsNil(o.AssignedObjectId) {
		var ret string
		return ret
	}
	return *o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedObjectId) {
		return nil, false
	}
	return o.AssignedObjectId, true
}

// HasAssignedObjectId returns a boolean if a field has been set.
func (o *PatchedBulkWritableNoteRequest) HasAssignedObjectId() bool {
	if o != nil && !IsNil(o.AssignedObjectId) {
		return true
	}

	return false
}

// SetAssignedObjectId gets a reference to the given string and assigns it to the AssignedObjectId field.
func (o *PatchedBulkWritableNoteRequest) SetAssignedObjectId(v string) {
	o.AssignedObjectId = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *PatchedBulkWritableNoteRequest) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableNoteRequest) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *PatchedBulkWritableNoteRequest) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *PatchedBulkWritableNoteRequest) SetNote(v string) {
	o.Note = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableNoteRequest) GetUser() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.User.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableNoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedBulkWritableNoteRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the User field.
func (o *PatchedBulkWritableNoteRequest) SetUser(v BulkWritableCircuitRequestTenant) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *PatchedBulkWritableNoteRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PatchedBulkWritableNoteRequest) UnsetUser() {
	o.User.Unset()
}

func (o PatchedBulkWritableNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.AssignedObjectType) {
		toSerialize["assigned_object_type"] = o.AssignedObjectType
	}
	if !IsNil(o.AssignedObjectId) {
		toSerialize["assigned_object_id"] = o.AssignedObjectId
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableNoteRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableNoteRequest := _PatchedBulkWritableNoteRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableNoteRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableNoteRequest(varPatchedBulkWritableNoteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "note")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableNoteRequest struct {
	value *PatchedBulkWritableNoteRequest
	isSet bool
}

func (v NullablePatchedBulkWritableNoteRequest) Get() *PatchedBulkWritableNoteRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableNoteRequest) Set(val *PatchedBulkWritableNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableNoteRequest(val *PatchedBulkWritableNoteRequest) *NullablePatchedBulkWritableNoteRequest {
	return &NullablePatchedBulkWritableNoteRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


