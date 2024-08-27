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

// checks if the BulkWritableCustomLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableCustomLinkRequest{}

// BulkWritableCustomLinkRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableCustomLinkRequest struct {
	Id string `json:"id"`
	ContentType string `json:"content_type"`
	Name string `json:"name"`
	// Jinja2 template code for link text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>. Links which render as empty text will not be displayed.
	Text string `json:"text"`
	// Jinja2 template code for link URL. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>.
	TargetUrl string `json:"target_url"`
	Weight *int32 `json:"weight,omitempty"`
	// Links with the same group will appear as a dropdown menu
	GroupName *string `json:"group_name,omitempty"`
	// The class of the first link in a group will be used for the dropdown button
	ButtonClass *ButtonClassEnum `json:"button_class,omitempty"`
	// Force link to open in a new window
	NewWindow bool `json:"new_window"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableCustomLinkRequest BulkWritableCustomLinkRequest

// NewBulkWritableCustomLinkRequest instantiates a new BulkWritableCustomLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableCustomLinkRequest(id string, contentType string, name string, text string, targetUrl string, newWindow bool) *BulkWritableCustomLinkRequest {
	this := BulkWritableCustomLinkRequest{}
	this.Id = id
	this.ContentType = contentType
	this.Name = name
	this.Text = text
	this.TargetUrl = targetUrl
	this.NewWindow = newWindow
	return &this
}

// NewBulkWritableCustomLinkRequestWithDefaults instantiates a new BulkWritableCustomLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableCustomLinkRequestWithDefaults() *BulkWritableCustomLinkRequest {
	this := BulkWritableCustomLinkRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableCustomLinkRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableCustomLinkRequest) SetId(v string) {
	o.Id = v
}

// GetContentType returns the ContentType field value
func (o *BulkWritableCustomLinkRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *BulkWritableCustomLinkRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetName returns the Name field value
func (o *BulkWritableCustomLinkRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableCustomLinkRequest) SetName(v string) {
	o.Name = v
}

// GetText returns the Text field value
func (o *BulkWritableCustomLinkRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *BulkWritableCustomLinkRequest) SetText(v string) {
	o.Text = v
}

// GetTargetUrl returns the TargetUrl field value
func (o *BulkWritableCustomLinkRequest) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *BulkWritableCustomLinkRequest) SetTargetUrl(v string) {
	o.TargetUrl = v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *BulkWritableCustomLinkRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *BulkWritableCustomLinkRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *BulkWritableCustomLinkRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *BulkWritableCustomLinkRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *BulkWritableCustomLinkRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *BulkWritableCustomLinkRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetButtonClass returns the ButtonClass field value if set, zero value otherwise.
func (o *BulkWritableCustomLinkRequest) GetButtonClass() ButtonClassEnum {
	if o == nil || IsNil(o.ButtonClass) {
		var ret ButtonClassEnum
		return ret
	}
	return *o.ButtonClass
}

// GetButtonClassOk returns a tuple with the ButtonClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetButtonClassOk() (*ButtonClassEnum, bool) {
	if o == nil || IsNil(o.ButtonClass) {
		return nil, false
	}
	return o.ButtonClass, true
}

// HasButtonClass returns a boolean if a field has been set.
func (o *BulkWritableCustomLinkRequest) HasButtonClass() bool {
	if o != nil && !IsNil(o.ButtonClass) {
		return true
	}

	return false
}

// SetButtonClass gets a reference to the given ButtonClassEnum and assigns it to the ButtonClass field.
func (o *BulkWritableCustomLinkRequest) SetButtonClass(v ButtonClassEnum) {
	o.ButtonClass = &v
}

// GetNewWindow returns the NewWindow field value
func (o *BulkWritableCustomLinkRequest) GetNewWindow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NewWindow
}

// GetNewWindowOk returns a tuple with the NewWindow field value
// and a boolean to check if the value has been set.
func (o *BulkWritableCustomLinkRequest) GetNewWindowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewWindow, true
}

// SetNewWindow sets field value
func (o *BulkWritableCustomLinkRequest) SetNewWindow(v bool) {
	o.NewWindow = v
}

func (o BulkWritableCustomLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableCustomLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["content_type"] = o.ContentType
	toSerialize["name"] = o.Name
	toSerialize["text"] = o.Text
	toSerialize["target_url"] = o.TargetUrl
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.ButtonClass) {
		toSerialize["button_class"] = o.ButtonClass
	}
	toSerialize["new_window"] = o.NewWindow

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableCustomLinkRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"content_type",
		"name",
		"text",
		"target_url",
		"new_window",
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

	varBulkWritableCustomLinkRequest := _BulkWritableCustomLinkRequest{}

	err = json.Unmarshal(data, &varBulkWritableCustomLinkRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableCustomLinkRequest(varBulkWritableCustomLinkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "text")
		delete(additionalProperties, "target_url")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "button_class")
		delete(additionalProperties, "new_window")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableCustomLinkRequest struct {
	value *BulkWritableCustomLinkRequest
	isSet bool
}

func (v NullableBulkWritableCustomLinkRequest) Get() *BulkWritableCustomLinkRequest {
	return v.value
}

func (v *NullableBulkWritableCustomLinkRequest) Set(val *BulkWritableCustomLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableCustomLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableCustomLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableCustomLinkRequest(val *BulkWritableCustomLinkRequest) *NullableBulkWritableCustomLinkRequest {
	return &NullableBulkWritableCustomLinkRequest{value: val, isSet: true}
}

func (v NullableBulkWritableCustomLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableCustomLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


