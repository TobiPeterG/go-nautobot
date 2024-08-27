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

// checks if the BulkWritableExternalIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableExternalIntegrationRequest{}

// BulkWritableExternalIntegrationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableExternalIntegrationRequest struct {
	Id string `json:"id"`
	Name string `json:"name"`
	RemoteUrl string `json:"remote_url"`
	// Verify SSL certificates when connecting to the remote system
	VerifySsl *bool `json:"verify_ssl,omitempty"`
	// Number of seconds to wait for a response
	Timeout *int32 `json:"timeout,omitempty"`
	// Optional user-defined JSON data for this integration
	ExtraConfig interface{} `json:"extra_config,omitempty"`
	HttpMethod *BulkWritableExternalIntegrationRequestHttpMethod `json:"http_method,omitempty"`
	// Headers for the HTTP request
	Headers interface{} `json:"headers,omitempty"`
	CaFilePath *string `json:"ca_file_path,omitempty"`
	SecretsGroup NullableBulkWritableExternalIntegrationRequestSecretsGroup `json:"secrets_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableExternalIntegrationRequest BulkWritableExternalIntegrationRequest

// NewBulkWritableExternalIntegrationRequest instantiates a new BulkWritableExternalIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableExternalIntegrationRequest(id string, name string, remoteUrl string) *BulkWritableExternalIntegrationRequest {
	this := BulkWritableExternalIntegrationRequest{}
	this.Id = id
	this.Name = name
	this.RemoteUrl = remoteUrl
	return &this
}

// NewBulkWritableExternalIntegrationRequestWithDefaults instantiates a new BulkWritableExternalIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableExternalIntegrationRequestWithDefaults() *BulkWritableExternalIntegrationRequest {
	this := BulkWritableExternalIntegrationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableExternalIntegrationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableExternalIntegrationRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BulkWritableExternalIntegrationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableExternalIntegrationRequest) SetName(v string) {
	o.Name = v
}

// GetRemoteUrl returns the RemoteUrl field value
func (o *BulkWritableExternalIntegrationRequest) GetRemoteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteUrl
}

// GetRemoteUrlOk returns a tuple with the RemoteUrl field value
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetRemoteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteUrl, true
}

// SetRemoteUrl sets field value
func (o *BulkWritableExternalIntegrationRequest) SetRemoteUrl(v string) {
	o.RemoteUrl = v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *BulkWritableExternalIntegrationRequest) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *BulkWritableExternalIntegrationRequest) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *BulkWritableExternalIntegrationRequest) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *BulkWritableExternalIntegrationRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableExternalIntegrationRequest) GetExtraConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableExternalIntegrationRequest) GetExtraConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExtraConfig) {
		return nil, false
	}
	return &o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasExtraConfig() bool {
	if o != nil && !IsNil(o.ExtraConfig) {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given interface{} and assigns it to the ExtraConfig field.
func (o *BulkWritableExternalIntegrationRequest) SetExtraConfig(v interface{}) {
	o.ExtraConfig = v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *BulkWritableExternalIntegrationRequest) GetHttpMethod() BulkWritableExternalIntegrationRequestHttpMethod {
	if o == nil || IsNil(o.HttpMethod) {
		var ret BulkWritableExternalIntegrationRequestHttpMethod
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetHttpMethodOk() (*BulkWritableExternalIntegrationRequestHttpMethod, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given BulkWritableExternalIntegrationRequestHttpMethod and assigns it to the HttpMethod field.
func (o *BulkWritableExternalIntegrationRequest) SetHttpMethod(v BulkWritableExternalIntegrationRequestHttpMethod) {
	o.HttpMethod = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableExternalIntegrationRequest) GetHeaders() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableExternalIntegrationRequest) GetHeadersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given interface{} and assigns it to the Headers field.
func (o *BulkWritableExternalIntegrationRequest) SetHeaders(v interface{}) {
	o.Headers = v
}

// GetCaFilePath returns the CaFilePath field value if set, zero value otherwise.
func (o *BulkWritableExternalIntegrationRequest) GetCaFilePath() string {
	if o == nil || IsNil(o.CaFilePath) {
		var ret string
		return ret
	}
	return *o.CaFilePath
}

// GetCaFilePathOk returns a tuple with the CaFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetCaFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.CaFilePath) {
		return nil, false
	}
	return o.CaFilePath, true
}

// HasCaFilePath returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasCaFilePath() bool {
	if o != nil && !IsNil(o.CaFilePath) {
		return true
	}

	return false
}

// SetCaFilePath gets a reference to the given string and assigns it to the CaFilePath field.
func (o *BulkWritableExternalIntegrationRequest) SetCaFilePath(v string) {
	o.CaFilePath = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableExternalIntegrationRequest) GetSecretsGroup() BulkWritableExternalIntegrationRequestSecretsGroup {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableExternalIntegrationRequestSecretsGroup
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableExternalIntegrationRequest) GetSecretsGroupOk() (*BulkWritableExternalIntegrationRequestSecretsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableExternalIntegrationRequestSecretsGroup and assigns it to the SecretsGroup field.
func (o *BulkWritableExternalIntegrationRequest) SetSecretsGroup(v BulkWritableExternalIntegrationRequestSecretsGroup) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *BulkWritableExternalIntegrationRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *BulkWritableExternalIntegrationRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableExternalIntegrationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableExternalIntegrationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableExternalIntegrationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableExternalIntegrationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableExternalIntegrationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableExternalIntegrationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableExternalIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableExternalIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["remote_url"] = o.RemoteUrl
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if o.ExtraConfig != nil {
		toSerialize["extra_config"] = o.ExtraConfig
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["http_method"] = o.HttpMethod
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.CaFilePath) {
		toSerialize["ca_file_path"] = o.CaFilePath
	}
	if o.SecretsGroup.IsSet() {
		toSerialize["secrets_group"] = o.SecretsGroup.Get()
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

func (o *BulkWritableExternalIntegrationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"remote_url",
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

	varBulkWritableExternalIntegrationRequest := _BulkWritableExternalIntegrationRequest{}

	err = json.Unmarshal(data, &varBulkWritableExternalIntegrationRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableExternalIntegrationRequest(varBulkWritableExternalIntegrationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "remote_url")
		delete(additionalProperties, "verify_ssl")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "extra_config")
		delete(additionalProperties, "http_method")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "ca_file_path")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableExternalIntegrationRequest struct {
	value *BulkWritableExternalIntegrationRequest
	isSet bool
}

func (v NullableBulkWritableExternalIntegrationRequest) Get() *BulkWritableExternalIntegrationRequest {
	return v.value
}

func (v *NullableBulkWritableExternalIntegrationRequest) Set(val *BulkWritableExternalIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableExternalIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableExternalIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableExternalIntegrationRequest(val *BulkWritableExternalIntegrationRequest) *NullableBulkWritableExternalIntegrationRequest {
	return &NullableBulkWritableExternalIntegrationRequest{value: val, isSet: true}
}

func (v NullableBulkWritableExternalIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableExternalIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


