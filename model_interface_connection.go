/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the InterfaceConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceConnection{}

// InterfaceConnection Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InterfaceConnection struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	InterfaceA Interface `json:"interface_a"`
	InterfaceB Interface `json:"interface_b"`
	ConnectedEndpointReachable NullableBool `json:"connected_endpoint_reachable"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceConnection InterfaceConnection

// NewInterfaceConnection instantiates a new InterfaceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceConnection(id string, objectType string, display string, interfaceA Interface, interfaceB Interface, connectedEndpointReachable NullableBool, created NullableTime, lastUpdated NullableTime) *InterfaceConnection {
	this := InterfaceConnection{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.InterfaceA = interfaceA
	this.InterfaceB = interfaceB
	this.ConnectedEndpointReachable = connectedEndpointReachable
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewInterfaceConnectionWithDefaults instantiates a new InterfaceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceConnectionWithDefaults() *InterfaceConnection {
	this := InterfaceConnection{}
	return &this
}

// GetId returns the Id field value
func (o *InterfaceConnection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InterfaceConnection) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *InterfaceConnection) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InterfaceConnection) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *InterfaceConnection) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *InterfaceConnection) SetDisplay(v string) {
	o.Display = v
}

// GetInterfaceA returns the InterfaceA field value
func (o *InterfaceConnection) GetInterfaceA() Interface {
	if o == nil {
		var ret Interface
		return ret
	}

	return o.InterfaceA
}

// GetInterfaceAOk returns a tuple with the InterfaceA field value
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetInterfaceAOk() (*Interface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceA, true
}

// SetInterfaceA sets field value
func (o *InterfaceConnection) SetInterfaceA(v Interface) {
	o.InterfaceA = v
}

// GetInterfaceB returns the InterfaceB field value
func (o *InterfaceConnection) GetInterfaceB() Interface {
	if o == nil {
		var ret Interface
		return ret
	}

	return o.InterfaceB
}

// GetInterfaceBOk returns a tuple with the InterfaceB field value
// and a boolean to check if the value has been set.
func (o *InterfaceConnection) GetInterfaceBOk() (*Interface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceB, true
}

// SetInterfaceB sets field value
func (o *InterfaceConnection) SetInterfaceB(v Interface) {
	o.InterfaceB = v
}

// GetConnectedEndpointReachable returns the ConnectedEndpointReachable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *InterfaceConnection) GetConnectedEndpointReachable() bool {
	if o == nil || o.ConnectedEndpointReachable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ConnectedEndpointReachable.Get()
}

// GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceConnection) GetConnectedEndpointReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectedEndpointReachable.Get(), o.ConnectedEndpointReachable.IsSet()
}

// SetConnectedEndpointReachable sets field value
func (o *InterfaceConnection) SetConnectedEndpointReachable(v bool) {
	o.ConnectedEndpointReachable.Set(&v)
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InterfaceConnection) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *InterfaceConnection) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InterfaceConnection) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceConnection) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *InterfaceConnection) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o InterfaceConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["interface_a"] = o.InterfaceA
	toSerialize["interface_b"] = o.InterfaceB
	toSerialize["connected_endpoint_reachable"] = o.ConnectedEndpointReachable.Get()
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"interface_a",
		"interface_b",
		"connected_endpoint_reachable",
		"created",
		"last_updated",
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

	varInterfaceConnection := _InterfaceConnection{}

	err = json.Unmarshal(data, &varInterfaceConnection)

	if err != nil {
		return err
	}

	*o = InterfaceConnection(varInterfaceConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "interface_a")
		delete(additionalProperties, "interface_b")
		delete(additionalProperties, "connected_endpoint_reachable")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceConnection struct {
	value *InterfaceConnection
	isSet bool
}

func (v NullableInterfaceConnection) Get() *InterfaceConnection {
	return v.value
}

func (v *NullableInterfaceConnection) Set(val *InterfaceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceConnection(val *InterfaceConnection) *NullableInterfaceConnection {
	return &NullableInterfaceConnection{value: val, isSet: true}
}

func (v NullableInterfaceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


