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

// checks if the BulkWritableObjectPermissionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableObjectPermissionRequest{}

// BulkWritableObjectPermissionRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableObjectPermissionRequest struct {
	Id string `json:"id"`
	ObjectTypes []string `json:"object_types"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// The list of actions granted by this permission
	Actions interface{} `json:"actions"`
	// Queryset filter matching the applicable objects of the selected type(s)
	Constraints interface{} `json:"constraints,omitempty"`
	Groups []BulkWritableCableRequestStatus `json:"groups,omitempty"`
	Users []BulkWritableCableRequestStatus `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableObjectPermissionRequest BulkWritableObjectPermissionRequest

// NewBulkWritableObjectPermissionRequest instantiates a new BulkWritableObjectPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableObjectPermissionRequest(id string, objectTypes []string, name string, actions interface{}) *BulkWritableObjectPermissionRequest {
	this := BulkWritableObjectPermissionRequest{}
	this.Id = id
	this.ObjectTypes = objectTypes
	this.Name = name
	this.Actions = actions
	return &this
}

// NewBulkWritableObjectPermissionRequestWithDefaults instantiates a new BulkWritableObjectPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableObjectPermissionRequestWithDefaults() *BulkWritableObjectPermissionRequest {
	this := BulkWritableObjectPermissionRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableObjectPermissionRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableObjectPermissionRequest) SetId(v string) {
	o.Id = v
}

// GetObjectTypes returns the ObjectTypes field value
func (o *BulkWritableObjectPermissionRequest) GetObjectTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ObjectTypes
}

// GetObjectTypesOk returns a tuple with the ObjectTypes field value
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetObjectTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectTypes, true
}

// SetObjectTypes sets field value
func (o *BulkWritableObjectPermissionRequest) SetObjectTypes(v []string) {
	o.ObjectTypes = v
}

// GetName returns the Name field value
func (o *BulkWritableObjectPermissionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableObjectPermissionRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableObjectPermissionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableObjectPermissionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableObjectPermissionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BulkWritableObjectPermissionRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BulkWritableObjectPermissionRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BulkWritableObjectPermissionRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetActions returns the Actions field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BulkWritableObjectPermissionRequest) GetActions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableObjectPermissionRequest) GetActionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value
func (o *BulkWritableObjectPermissionRequest) SetActions(v interface{}) {
	o.Actions = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableObjectPermissionRequest) GetConstraints() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableObjectPermissionRequest) GetConstraintsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return &o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *BulkWritableObjectPermissionRequest) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given interface{} and assigns it to the Constraints field.
func (o *BulkWritableObjectPermissionRequest) SetConstraints(v interface{}) {
	o.Constraints = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *BulkWritableObjectPermissionRequest) GetGroups() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Groups) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetGroupsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *BulkWritableObjectPermissionRequest) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Groups field.
func (o *BulkWritableObjectPermissionRequest) SetGroups(v []BulkWritableCableRequestStatus) {
	o.Groups = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *BulkWritableObjectPermissionRequest) GetUsers() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Users) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableObjectPermissionRequest) GetUsersOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BulkWritableObjectPermissionRequest) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Users field.
func (o *BulkWritableObjectPermissionRequest) SetUsers(v []BulkWritableCableRequestStatus) {
	o.Users = v
}

func (o BulkWritableObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableObjectPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_types"] = o.ObjectTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableObjectPermissionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_types",
		"name",
		"actions",
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

	varBulkWritableObjectPermissionRequest := _BulkWritableObjectPermissionRequest{}

	err = json.Unmarshal(data, &varBulkWritableObjectPermissionRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableObjectPermissionRequest(varBulkWritableObjectPermissionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "actions")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableObjectPermissionRequest struct {
	value *BulkWritableObjectPermissionRequest
	isSet bool
}

func (v NullableBulkWritableObjectPermissionRequest) Get() *BulkWritableObjectPermissionRequest {
	return v.value
}

func (v *NullableBulkWritableObjectPermissionRequest) Set(val *BulkWritableObjectPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableObjectPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableObjectPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableObjectPermissionRequest(val *BulkWritableObjectPermissionRequest) *NullableBulkWritableObjectPermissionRequest {
	return &NullableBulkWritableObjectPermissionRequest{value: val, isSet: true}
}

func (v NullableBulkWritableObjectPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableObjectPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


