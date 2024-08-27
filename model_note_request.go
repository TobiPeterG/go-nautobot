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

// checks if the NoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteRequest{}

// NoteRequest This base serializer implements common fields and logic for all ModelSerializers.  Namely, it:  - defines the `display` field which exposes a human friendly value for the given object. - ensures that `id` field is always present on the serializer as well. - ensures that `created` and `last_updated` fields are always present if applicable to this model and serializer. - ensures that `object_type` field is always present on the serializer which represents the content-type of this   serializer's associated model (e.g. \"dcim.device\"). This is required as the OpenAPI schema, using the   PolymorphicProxySerializer class defined below, relies upon this field as a way to identify to the client   which of several possible serializers are in use for a given attribute. - supports `?depth` query parameter. It is passed in as `nested_depth` to the `build_nested_field()` function   to enable the dynamic generation of nested serializers.
type NoteRequest struct {
	AssignedObjectType string `json:"assigned_object_type"`
	AssignedObjectId string `json:"assigned_object_id"`
	Note string `json:"note"`
	User NullableBulkWritableCircuitRequestTenant `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NoteRequest NoteRequest

// NewNoteRequest instantiates a new NoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteRequest(assignedObjectType string, assignedObjectId string, note string) *NoteRequest {
	this := NoteRequest{}
	this.AssignedObjectType = assignedObjectType
	this.AssignedObjectId = assignedObjectId
	this.Note = note
	return &this
}

// NewNoteRequestWithDefaults instantiates a new NoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteRequestWithDefaults() *NoteRequest {
	this := NoteRequest{}
	return &this
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *NoteRequest) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *NoteRequest) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *NoteRequest) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *NoteRequest) GetAssignedObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *NoteRequest) GetAssignedObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *NoteRequest) SetAssignedObjectId(v string) {
	o.AssignedObjectId = v
}

// GetNote returns the Note field value
func (o *NoteRequest) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *NoteRequest) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *NoteRequest) SetNote(v string) {
	o.Note = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NoteRequest) GetUser() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.User.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *NoteRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the User field.
func (o *NoteRequest) SetUser(v BulkWritableCircuitRequestTenant) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *NoteRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *NoteRequest) UnsetUser() {
	o.User.Unset()
}

func (o NoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object_id"] = o.AssignedObjectId
	toSerialize["note"] = o.Note
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NoteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assigned_object_type",
		"assigned_object_id",
		"note",
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

	varNoteRequest := _NoteRequest{}

	err = json.Unmarshal(data, &varNoteRequest)

	if err != nil {
		return err
	}

	*o = NoteRequest(varNoteRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "note")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNoteRequest struct {
	value *NoteRequest
	isSet bool
}

func (v NullableNoteRequest) Get() *NoteRequest {
	return v.value
}

func (v *NullableNoteRequest) Set(val *NoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteRequest(val *NoteRequest) *NullableNoteRequest {
	return &NullableNoteRequest{value: val, isSet: true}
}

func (v NullableNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


