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

// checks if the GraphQLAPIRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphQLAPIRequest{}

// GraphQLAPIRequest struct for GraphQLAPIRequest
type GraphQLAPIRequest struct {
	// GraphQL query
	Query string `json:"query"`
	// Variables in JSON Format
	Variables interface{} `json:"variables,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GraphQLAPIRequest GraphQLAPIRequest

// NewGraphQLAPIRequest instantiates a new GraphQLAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphQLAPIRequest(query string) *GraphQLAPIRequest {
	this := GraphQLAPIRequest{}
	this.Query = query
	return &this
}

// NewGraphQLAPIRequestWithDefaults instantiates a new GraphQLAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphQLAPIRequestWithDefaults() *GraphQLAPIRequest {
	this := GraphQLAPIRequest{}
	return &this
}

// GetQuery returns the Query field value
func (o *GraphQLAPIRequest) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GraphQLAPIRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *GraphQLAPIRequest) SetQuery(v string) {
	o.Query = v
}

// GetVariables returns the Variables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphQLAPIRequest) GetVariables() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphQLAPIRequest) GetVariablesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *GraphQLAPIRequest) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given interface{} and assigns it to the Variables field.
func (o *GraphQLAPIRequest) SetVariables(v interface{}) {
	o.Variables = v
}

func (o GraphQLAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphQLAPIRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GraphQLAPIRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"query",
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

	varGraphQLAPIRequest := _GraphQLAPIRequest{}

	err = json.Unmarshal(data, &varGraphQLAPIRequest)

	if err != nil {
		return err
	}

	*o = GraphQLAPIRequest(varGraphQLAPIRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		delete(additionalProperties, "variables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGraphQLAPIRequest struct {
	value *GraphQLAPIRequest
	isSet bool
}

func (v NullableGraphQLAPIRequest) Get() *GraphQLAPIRequest {
	return v.value
}

func (v *NullableGraphQLAPIRequest) Set(val *GraphQLAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphQLAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphQLAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphQLAPIRequest(val *GraphQLAPIRequest) *NullableGraphQLAPIRequest {
	return &NullableGraphQLAPIRequest{value: val, isSet: true}
}

func (v NullableGraphQLAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphQLAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


