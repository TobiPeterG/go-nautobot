/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedJobButtonRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedJobButtonRequest{}

// PatchedJobButtonRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedJobButtonRequest struct {
	ContentTypes []string `json:"content_types,omitempty"`
	Name *string `json:"name,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// Jinja2 template code for button text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>. Buttons which render as empty text will not be displayed.
	Text *string `json:"text,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	// Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group.
	GroupName *string `json:"group_name,omitempty"`
	ButtonClass *ButtonClassEnum `json:"button_class,omitempty"`
	// Enable confirmation pop-up box. <span class='text-danger'>WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!</span>
	Confirmation *bool `json:"confirmation,omitempty"`
	Job *BulkWritableJobButtonRequestJob `json:"job,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedJobButtonRequest PatchedJobButtonRequest

// NewPatchedJobButtonRequest instantiates a new PatchedJobButtonRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedJobButtonRequest() *PatchedJobButtonRequest {
	this := PatchedJobButtonRequest{}
	return &this
}

// NewPatchedJobButtonRequestWithDefaults instantiates a new PatchedJobButtonRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedJobButtonRequestWithDefaults() *PatchedJobButtonRequest {
	this := PatchedJobButtonRequest{}
	return &this
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetContentTypes() []string {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []string
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []string and assigns it to the ContentTypes field.
func (o *PatchedJobButtonRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedJobButtonRequest) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchedJobButtonRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PatchedJobButtonRequest) SetText(v string) {
	o.Text = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PatchedJobButtonRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *PatchedJobButtonRequest) SetGroupName(v string) {
	o.GroupName = &v
}

// GetButtonClass returns the ButtonClass field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetButtonClass() ButtonClassEnum {
	if o == nil || IsNil(o.ButtonClass) {
		var ret ButtonClassEnum
		return ret
	}
	return *o.ButtonClass
}

// GetButtonClassOk returns a tuple with the ButtonClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetButtonClassOk() (*ButtonClassEnum, bool) {
	if o == nil || IsNil(o.ButtonClass) {
		return nil, false
	}
	return o.ButtonClass, true
}

// HasButtonClass returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasButtonClass() bool {
	if o != nil && !IsNil(o.ButtonClass) {
		return true
	}

	return false
}

// SetButtonClass gets a reference to the given ButtonClassEnum and assigns it to the ButtonClass field.
func (o *PatchedJobButtonRequest) SetButtonClass(v ButtonClassEnum) {
	o.ButtonClass = &v
}

// GetConfirmation returns the Confirmation field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetConfirmation() bool {
	if o == nil || IsNil(o.Confirmation) {
		var ret bool
		return ret
	}
	return *o.Confirmation
}

// GetConfirmationOk returns a tuple with the Confirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetConfirmationOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmation) {
		return nil, false
	}
	return o.Confirmation, true
}

// HasConfirmation returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasConfirmation() bool {
	if o != nil && !IsNil(o.Confirmation) {
		return true
	}

	return false
}

// SetConfirmation gets a reference to the given bool and assigns it to the Confirmation field.
func (o *PatchedJobButtonRequest) SetConfirmation(v bool) {
	o.Confirmation = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *PatchedJobButtonRequest) GetJob() BulkWritableJobButtonRequestJob {
	if o == nil || IsNil(o.Job) {
		var ret BulkWritableJobButtonRequestJob
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedJobButtonRequest) GetJobOk() (*BulkWritableJobButtonRequestJob, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *PatchedJobButtonRequest) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given BulkWritableJobButtonRequestJob and assigns it to the Job field.
func (o *PatchedJobButtonRequest) SetJob(v BulkWritableJobButtonRequestJob) {
	o.Job = &v
}

func (o PatchedJobButtonRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedJobButtonRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentTypes) {
		toSerialize["content_types"] = o.ContentTypes
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if !IsNil(o.ButtonClass) {
		toSerialize["button_class"] = o.ButtonClass
	}
	if !IsNil(o.Confirmation) {
		toSerialize["confirmation"] = o.Confirmation
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedJobButtonRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedJobButtonRequest := _PatchedJobButtonRequest{}

	err = json.Unmarshal(data, &varPatchedJobButtonRequest)

	if err != nil {
		return err
	}

	*o = PatchedJobButtonRequest(varPatchedJobButtonRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "text")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "group_name")
		delete(additionalProperties, "button_class")
		delete(additionalProperties, "confirmation")
		delete(additionalProperties, "job")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedJobButtonRequest struct {
	value *PatchedJobButtonRequest
	isSet bool
}

func (v NullablePatchedJobButtonRequest) Get() *PatchedJobButtonRequest {
	return v.value
}

func (v *NullablePatchedJobButtonRequest) Set(val *PatchedJobButtonRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedJobButtonRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedJobButtonRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedJobButtonRequest(val *PatchedJobButtonRequest) *NullablePatchedJobButtonRequest {
	return &NullablePatchedJobButtonRequest{value: val, isSet: true}
}

func (v NullablePatchedJobButtonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedJobButtonRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


