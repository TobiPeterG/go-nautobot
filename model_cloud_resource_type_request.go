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

// checks if the CloudResourceTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudResourceTypeRequest{}

// CloudResourceTypeRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type CloudResourceTypeRequest struct {
	ContentTypes []string `json:"content_types"`
	// Type of cloud objects
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ConfigSchema interface{} `json:"config_schema,omitempty"`
	Provider BulkWritableCloudAccountRequestProvider `json:"provider"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudResourceTypeRequest CloudResourceTypeRequest

// NewCloudResourceTypeRequest instantiates a new CloudResourceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudResourceTypeRequest(contentTypes []string, name string, provider BulkWritableCloudAccountRequestProvider) *CloudResourceTypeRequest {
	this := CloudResourceTypeRequest{}
	this.ContentTypes = contentTypes
	this.Name = name
	this.Provider = provider
	return &this
}

// NewCloudResourceTypeRequestWithDefaults instantiates a new CloudResourceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudResourceTypeRequestWithDefaults() *CloudResourceTypeRequest {
	this := CloudResourceTypeRequest{}
	return &this
}

// GetContentTypes returns the ContentTypes field value
func (o *CloudResourceTypeRequest) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *CloudResourceTypeRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *CloudResourceTypeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CloudResourceTypeRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudResourceTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudResourceTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudResourceTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetConfigSchema returns the ConfigSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudResourceTypeRequest) GetConfigSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ConfigSchema
}

// GetConfigSchemaOk returns a tuple with the ConfigSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudResourceTypeRequest) GetConfigSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ConfigSchema) {
		return nil, false
	}
	return &o.ConfigSchema, true
}

// HasConfigSchema returns a boolean if a field has been set.
func (o *CloudResourceTypeRequest) HasConfigSchema() bool {
	if o != nil && !IsNil(o.ConfigSchema) {
		return true
	}

	return false
}

// SetConfigSchema gets a reference to the given interface{} and assigns it to the ConfigSchema field.
func (o *CloudResourceTypeRequest) SetConfigSchema(v interface{}) {
	o.ConfigSchema = v
}

// GetProvider returns the Provider field value
func (o *CloudResourceTypeRequest) GetProvider() BulkWritableCloudAccountRequestProvider {
	if o == nil {
		var ret BulkWritableCloudAccountRequestProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetProviderOk() (*BulkWritableCloudAccountRequestProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CloudResourceTypeRequest) SetProvider(v BulkWritableCloudAccountRequestProvider) {
	o.Provider = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CloudResourceTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CloudResourceTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CloudResourceTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CloudResourceTypeRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CloudResourceTypeRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *CloudResourceTypeRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CloudResourceTypeRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudResourceTypeRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CloudResourceTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *CloudResourceTypeRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o CloudResourceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudResourceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ConfigSchema != nil {
		toSerialize["config_schema"] = o.ConfigSchema
	}
	toSerialize["provider"] = o.Provider
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

func (o *CloudResourceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content_types",
		"name",
		"provider",
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

	varCloudResourceTypeRequest := _CloudResourceTypeRequest{}

	err = json.Unmarshal(data, &varCloudResourceTypeRequest)

	if err != nil {
		return err
	}

	*o = CloudResourceTypeRequest(varCloudResourceTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "config_schema")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudResourceTypeRequest struct {
	value *CloudResourceTypeRequest
	isSet bool
}

func (v NullableCloudResourceTypeRequest) Get() *CloudResourceTypeRequest {
	return v.value
}

func (v *NullableCloudResourceTypeRequest) Set(val *CloudResourceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudResourceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudResourceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudResourceTypeRequest(val *CloudResourceTypeRequest) *NullableCloudResourceTypeRequest {
	return &NullableCloudResourceTypeRequest{value: val, isSet: true}
}

func (v NullableCloudResourceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudResourceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


