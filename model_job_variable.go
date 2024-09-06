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

// checks if the JobVariable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobVariable{}

// JobVariable Serializer used for responses from the JobModelViewSet.variables() detail endpoint.
type JobVariable struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Label string `json:"label"`
	HelpText string `json:"help_text"`
	Default interface{} `json:"default"`
	Required bool `json:"required"`
	MinLength int32 `json:"min_length"`
	MaxLength int32 `json:"max_length"`
	MinValue int32 `json:"min_value"`
	MaxValue int32 `json:"max_value"`
	Choices interface{} `json:"choices"`
	Model string `json:"model"`
	AdditionalProperties map[string]interface{}
}

type _JobVariable JobVariable

// NewJobVariable instantiates a new JobVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobVariable(name string, type_ string, label string, helpText string, default_ interface{}, required bool, minLength int32, maxLength int32, minValue int32, maxValue int32, choices interface{}, model string) *JobVariable {
	this := JobVariable{}
	this.Name = name
	this.Type = type_
	this.Label = label
	this.HelpText = helpText
	this.Default = default_
	this.Required = required
	this.MinLength = minLength
	this.MaxLength = maxLength
	this.MinValue = minValue
	this.MaxValue = maxValue
	this.Choices = choices
	this.Model = model
	return &this
}

// NewJobVariableWithDefaults instantiates a new JobVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobVariableWithDefaults() *JobVariable {
	this := JobVariable{}
	return &this
}

// GetName returns the Name field value
func (o *JobVariable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobVariable) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *JobVariable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobVariable) SetType(v string) {
	o.Type = v
}

// GetLabel returns the Label field value
func (o *JobVariable) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *JobVariable) SetLabel(v string) {
	o.Label = v
}

// GetHelpText returns the HelpText field value
func (o *JobVariable) GetHelpText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetHelpTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HelpText, true
}

// SetHelpText sets field value
func (o *JobVariable) SetHelpText(v string) {
	o.HelpText = v
}

// GetDefault returns the Default field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JobVariable) GetDefault() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobVariable) GetDefaultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *JobVariable) SetDefault(v interface{}) {
	o.Default = v
}

// GetRequired returns the Required field value
func (o *JobVariable) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *JobVariable) SetRequired(v bool) {
	o.Required = v
}

// GetMinLength returns the MinLength field value
func (o *JobVariable) GetMinLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetMinLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinLength, true
}

// SetMinLength sets field value
func (o *JobVariable) SetMinLength(v int32) {
	o.MinLength = v
}

// GetMaxLength returns the MaxLength field value
func (o *JobVariable) GetMaxLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetMaxLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLength, true
}

// SetMaxLength sets field value
func (o *JobVariable) SetMaxLength(v int32) {
	o.MaxLength = v
}

// GetMinValue returns the MinValue field value
func (o *JobVariable) GetMinValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetMinValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinValue, true
}

// SetMinValue sets field value
func (o *JobVariable) SetMinValue(v int32) {
	o.MinValue = v
}

// GetMaxValue returns the MaxValue field value
func (o *JobVariable) GetMaxValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetMaxValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxValue, true
}

// SetMaxValue sets field value
func (o *JobVariable) SetMaxValue(v int32) {
	o.MaxValue = v
}

// GetChoices returns the Choices field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JobVariable) GetChoices() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobVariable) GetChoicesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return &o.Choices, true
}

// SetChoices sets field value
func (o *JobVariable) SetChoices(v interface{}) {
	o.Choices = v
}

// GetModel returns the Model field value
func (o *JobVariable) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *JobVariable) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *JobVariable) SetModel(v string) {
	o.Model = v
}

func (o JobVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["label"] = o.Label
	toSerialize["help_text"] = o.HelpText
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	toSerialize["required"] = o.Required
	toSerialize["min_length"] = o.MinLength
	toSerialize["max_length"] = o.MaxLength
	toSerialize["min_value"] = o.MinValue
	toSerialize["max_value"] = o.MaxValue
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	toSerialize["model"] = o.Model

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobVariable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"label",
		"help_text",
		"default",
		"required",
		"min_length",
		"max_length",
		"min_value",
		"max_value",
		"choices",
		"model",
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

	varJobVariable := _JobVariable{}

	err = json.Unmarshal(data, &varJobVariable)

	if err != nil {
		return err
	}

	*o = JobVariable(varJobVariable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "label")
		delete(additionalProperties, "help_text")
		delete(additionalProperties, "default")
		delete(additionalProperties, "required")
		delete(additionalProperties, "min_length")
		delete(additionalProperties, "max_length")
		delete(additionalProperties, "min_value")
		delete(additionalProperties, "max_value")
		delete(additionalProperties, "choices")
		delete(additionalProperties, "model")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobVariable struct {
	value *JobVariable
	isSet bool
}

func (v NullableJobVariable) Get() *JobVariable {
	return v.value
}

func (v *NullableJobVariable) Set(val *JobVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableJobVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableJobVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobVariable(val *JobVariable) *NullableJobVariable {
	return &NullableJobVariable{value: val, isSet: true}
}

func (v NullableJobVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


