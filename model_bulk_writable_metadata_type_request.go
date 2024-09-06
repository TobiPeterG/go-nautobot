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

// checks if the BulkWritableMetadataTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableMetadataTypeRequest{}

// BulkWritableMetadataTypeRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableMetadataTypeRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// The type of data allowed for any Metadata of this type.
	DataType DataTypeEnum `json:"data_type"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableMetadataTypeRequest BulkWritableMetadataTypeRequest

// NewBulkWritableMetadataTypeRequest instantiates a new BulkWritableMetadataTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableMetadataTypeRequest(id string, contentTypes []string, name string, dataType DataTypeEnum) *BulkWritableMetadataTypeRequest {
	this := BulkWritableMetadataTypeRequest{}
	this.Id = id
	this.ContentTypes = contentTypes
	this.Name = name
	this.DataType = dataType
	return &this
}

// NewBulkWritableMetadataTypeRequestWithDefaults instantiates a new BulkWritableMetadataTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableMetadataTypeRequestWithDefaults() *BulkWritableMetadataTypeRequest {
	this := BulkWritableMetadataTypeRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableMetadataTypeRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableMetadataTypeRequest) SetId(v string) {
	o.Id = v
}

// GetContentTypes returns the ContentTypes field value
func (o *BulkWritableMetadataTypeRequest) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *BulkWritableMetadataTypeRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *BulkWritableMetadataTypeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableMetadataTypeRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableMetadataTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableMetadataTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableMetadataTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDataType returns the DataType field value
func (o *BulkWritableMetadataTypeRequest) GetDataType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetDataTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *BulkWritableMetadataTypeRequest) SetDataType(v DataTypeEnum) {
	o.DataType = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableMetadataTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableMetadataTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableMetadataTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableMetadataTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableMetadataTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableMetadataTypeRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableMetadataTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableMetadataTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableMetadataTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["data_type"] = o.DataType
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

func (o *BulkWritableMetadataTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"content_types",
		"name",
		"data_type",
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

	varBulkWritableMetadataTypeRequest := _BulkWritableMetadataTypeRequest{}

	err = json.Unmarshal(data, &varBulkWritableMetadataTypeRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableMetadataTypeRequest(varBulkWritableMetadataTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "data_type")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableMetadataTypeRequest struct {
	value *BulkWritableMetadataTypeRequest
	isSet bool
}

func (v NullableBulkWritableMetadataTypeRequest) Get() *BulkWritableMetadataTypeRequest {
	return v.value
}

func (v *NullableBulkWritableMetadataTypeRequest) Set(val *BulkWritableMetadataTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableMetadataTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableMetadataTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableMetadataTypeRequest(val *BulkWritableMetadataTypeRequest) *NullableBulkWritableMetadataTypeRequest {
	return &NullableBulkWritableMetadataTypeRequest{value: val, isSet: true}
}

func (v NullableBulkWritableMetadataTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableMetadataTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


