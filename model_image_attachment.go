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

// checks if the ImageAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageAttachment{}

// ImageAttachment Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ImageAttachment struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	ContentType string `json:"content_type"`
	ObjectId string `json:"object_id"`
	Image string `json:"image"`
	ImageHeight int32 `json:"image_height"`
	ImageWidth int32 `json:"image_width"`
	Name *string `json:"name,omitempty"`
	Created time.Time `json:"created"`
	AdditionalProperties map[string]interface{}
}

type _ImageAttachment ImageAttachment

// NewImageAttachment instantiates a new ImageAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAttachment(id string, objectType string, display string, url string, naturalSlug string, contentType string, objectId string, image string, imageHeight int32, imageWidth int32, created time.Time) *ImageAttachment {
	this := ImageAttachment{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.ContentType = contentType
	this.ObjectId = objectId
	this.Image = image
	this.ImageHeight = imageHeight
	this.ImageWidth = imageWidth
	this.Created = created
	return &this
}

// NewImageAttachmentWithDefaults instantiates a new ImageAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAttachmentWithDefaults() *ImageAttachment {
	this := ImageAttachment{}
	return &this
}

// GetId returns the Id field value
func (o *ImageAttachment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImageAttachment) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *ImageAttachment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ImageAttachment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *ImageAttachment) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ImageAttachment) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *ImageAttachment) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ImageAttachment) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *ImageAttachment) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *ImageAttachment) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetContentType returns the ContentType field value
func (o *ImageAttachment) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ImageAttachment) SetContentType(v string) {
	o.ContentType = v
}

// GetObjectId returns the ObjectId field value
func (o *ImageAttachment) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ImageAttachment) SetObjectId(v string) {
	o.ObjectId = v
}

// GetImage returns the Image field value
func (o *ImageAttachment) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ImageAttachment) SetImage(v string) {
	o.Image = v
}

// GetImageHeight returns the ImageHeight field value
func (o *ImageAttachment) GetImageHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetImageHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageHeight, true
}

// SetImageHeight sets field value
func (o *ImageAttachment) SetImageHeight(v int32) {
	o.ImageHeight = v
}

// GetImageWidth returns the ImageWidth field value
func (o *ImageAttachment) GetImageWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetImageWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageWidth, true
}

// SetImageWidth sets field value
func (o *ImageAttachment) SetImageWidth(v int32) {
	o.ImageWidth = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageAttachment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageAttachment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageAttachment) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value
func (o *ImageAttachment) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ImageAttachment) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ImageAttachment) SetCreated(v time.Time) {
	o.Created = v
}

func (o ImageAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["content_type"] = o.ContentType
	toSerialize["object_id"] = o.ObjectId
	toSerialize["image"] = o.Image
	toSerialize["image_height"] = o.ImageHeight
	toSerialize["image_width"] = o.ImageWidth
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["created"] = o.Created

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImageAttachment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"content_type",
		"object_id",
		"image",
		"image_height",
		"image_width",
		"created",
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

	varImageAttachment := _ImageAttachment{}

	err = json.Unmarshal(data, &varImageAttachment)

	if err != nil {
		return err
	}

	*o = ImageAttachment(varImageAttachment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "image")
		delete(additionalProperties, "image_height")
		delete(additionalProperties, "image_width")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImageAttachment struct {
	value *ImageAttachment
	isSet bool
}

func (v NullableImageAttachment) Get() *ImageAttachment {
	return v.value
}

func (v *NullableImageAttachment) Set(val *ImageAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAttachment(val *ImageAttachment) *NullableImageAttachment {
	return &NullableImageAttachment{value: val, isSet: true}
}

func (v NullableImageAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


