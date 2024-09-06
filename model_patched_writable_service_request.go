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

// checks if the PatchedWritableServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableServiceRequest{}

// PatchedWritableServiceRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedWritableServiceRequest struct {
	Ports []int32 `json:"ports,omitempty"`
	Name *string `json:"name,omitempty"`
	Protocol *ServiceProtocolChoices `json:"protocol,omitempty"`
	Description *string `json:"description,omitempty"`
	Device NullableBulkWritableServiceRequestDevice `json:"device,omitempty"`
	VirtualMachine NullableBulkWritableServiceRequestVirtualMachine `json:"virtual_machine,omitempty"`
	IpAddresses []IPAddresses `json:"ip_addresses,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableServiceRequest PatchedWritableServiceRequest

// NewPatchedWritableServiceRequest instantiates a new PatchedWritableServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableServiceRequest() *PatchedWritableServiceRequest {
	this := PatchedWritableServiceRequest{}
	return &this
}

// NewPatchedWritableServiceRequestWithDefaults instantiates a new PatchedWritableServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableServiceRequestWithDefaults() *PatchedWritableServiceRequest {
	this := PatchedWritableServiceRequest{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetPorts() []int32 {
	if o == nil || IsNil(o.Ports) {
		var ret []int32
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetPortsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []int32 and assigns it to the Ports field.
func (o *PatchedWritableServiceRequest) SetPorts(v []int32) {
	o.Ports = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableServiceRequest) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetProtocol() ServiceProtocolChoices {
	if o == nil || IsNil(o.Protocol) {
		var ret ServiceProtocolChoices
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetProtocolOk() (*ServiceProtocolChoices, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given ServiceProtocolChoices and assigns it to the Protocol field.
func (o *PatchedWritableServiceRequest) SetProtocol(v ServiceProtocolChoices) {
	o.Protocol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableServiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableServiceRequest) GetDevice() BulkWritableServiceRequestDevice {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BulkWritableServiceRequestDevice
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableServiceRequest) GetDeviceOk() (*BulkWritableServiceRequestDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBulkWritableServiceRequestDevice and assigns it to the Device field.
func (o *PatchedWritableServiceRequest) SetDevice(v BulkWritableServiceRequestDevice) {
	o.Device.Set(&v)
}
// SetDeviceNil sets the value for Device to be an explicit nil
func (o *PatchedWritableServiceRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *PatchedWritableServiceRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableServiceRequest) GetVirtualMachine() BulkWritableServiceRequestVirtualMachine {
	if o == nil || IsNil(o.VirtualMachine.Get()) {
		var ret BulkWritableServiceRequestVirtualMachine
		return ret
	}
	return *o.VirtualMachine.Get()
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableServiceRequest) GetVirtualMachineOk() (*BulkWritableServiceRequestVirtualMachine, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualMachine.Get(), o.VirtualMachine.IsSet()
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine.IsSet() {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given NullableBulkWritableServiceRequestVirtualMachine and assigns it to the VirtualMachine field.
func (o *PatchedWritableServiceRequest) SetVirtualMachine(v BulkWritableServiceRequestVirtualMachine) {
	o.VirtualMachine.Set(&v)
}
// SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil
func (o *PatchedWritableServiceRequest) SetVirtualMachineNil() {
	o.VirtualMachine.Set(nil)
}

// UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil
func (o *PatchedWritableServiceRequest) UnsetVirtualMachine() {
	o.VirtualMachine.Unset()
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetIpAddresses() []IPAddresses {
	if o == nil || IsNil(o.IpAddresses) {
		var ret []IPAddresses
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetIpAddressesOk() ([]IPAddresses, bool) {
	if o == nil || IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasIpAddresses() bool {
	if o != nil && !IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []IPAddresses and assigns it to the IpAddresses field.
func (o *PatchedWritableServiceRequest) SetIpAddresses(v []IPAddresses) {
	o.IpAddresses = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedWritableServiceRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableServiceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedWritableServiceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableServiceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedWritableServiceRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedWritableServiceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedWritableServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.VirtualMachine.IsSet() {
		toSerialize["virtual_machine"] = o.VirtualMachine.Get()
	}
	if !IsNil(o.IpAddresses) {
		toSerialize["ip_addresses"] = o.IpAddresses
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

func (o *PatchedWritableServiceRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableServiceRequest := _PatchedWritableServiceRequest{}

	err = json.Unmarshal(data, &varPatchedWritableServiceRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableServiceRequest(varPatchedWritableServiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ports")
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device")
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "ip_addresses")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableServiceRequest struct {
	value *PatchedWritableServiceRequest
	isSet bool
}

func (v NullablePatchedWritableServiceRequest) Get() *PatchedWritableServiceRequest {
	return v.value
}

func (v *NullablePatchedWritableServiceRequest) Set(val *PatchedWritableServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableServiceRequest(val *PatchedWritableServiceRequest) *NullablePatchedWritableServiceRequest {
	return &NullablePatchedWritableServiceRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


