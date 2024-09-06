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

// checks if the DeviceType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceType{}

// DeviceType Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type DeviceType struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	SubdeviceRole *DeviceTypeSubdeviceRole `json:"subdevice_role,omitempty"`
	FrontImage NullableString `json:"front_image,omitempty"`
	RearImage NullableString `json:"rear_image,omitempty"`
	DeviceCount *int32 `json:"device_count,omitempty"`
	Model string `json:"model"`
	// Discrete part number (optional)
	PartNumber *string `json:"part_number,omitempty"`
	UHeight *int32 `json:"u_height,omitempty"`
	// Device consumes both front and rear rack faces
	IsFullDepth *bool `json:"is_full_depth,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Manufacturer BulkWritableCableRequestStatus `json:"manufacturer"`
	DeviceFamily NullableBulkWritableCircuitRequestTenant `json:"device_family,omitempty"`
	SoftwareImageFiles []SoftwareImageFiles `json:"software_image_files"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceType DeviceType

// NewDeviceType instantiates a new DeviceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceType(id string, objectType string, display string, url string, naturalSlug string, model string, manufacturer BulkWritableCableRequestStatus, softwareImageFiles []SoftwareImageFiles, created NullableTime, lastUpdated NullableTime, notesUrl string) *DeviceType {
	this := DeviceType{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Model = model
	this.Manufacturer = manufacturer
	this.SoftwareImageFiles = softwareImageFiles
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewDeviceTypeWithDefaults instantiates a new DeviceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceTypeWithDefaults() *DeviceType {
	this := DeviceType{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceType) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceType) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *DeviceType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *DeviceType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *DeviceType) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DeviceType) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *DeviceType) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeviceType) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *DeviceType) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *DeviceType) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetSubdeviceRole returns the SubdeviceRole field value if set, zero value otherwise.
func (o *DeviceType) GetSubdeviceRole() DeviceTypeSubdeviceRole {
	if o == nil || IsNil(o.SubdeviceRole) {
		var ret DeviceTypeSubdeviceRole
		return ret
	}
	return *o.SubdeviceRole
}

// GetSubdeviceRoleOk returns a tuple with the SubdeviceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetSubdeviceRoleOk() (*DeviceTypeSubdeviceRole, bool) {
	if o == nil || IsNil(o.SubdeviceRole) {
		return nil, false
	}
	return o.SubdeviceRole, true
}

// HasSubdeviceRole returns a boolean if a field has been set.
func (o *DeviceType) HasSubdeviceRole() bool {
	if o != nil && !IsNil(o.SubdeviceRole) {
		return true
	}

	return false
}

// SetSubdeviceRole gets a reference to the given DeviceTypeSubdeviceRole and assigns it to the SubdeviceRole field.
func (o *DeviceType) SetSubdeviceRole(v DeviceTypeSubdeviceRole) {
	o.SubdeviceRole = &v
}

// GetFrontImage returns the FrontImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetFrontImage() string {
	if o == nil || IsNil(o.FrontImage.Get()) {
		var ret string
		return ret
	}
	return *o.FrontImage.Get()
}

// GetFrontImageOk returns a tuple with the FrontImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetFrontImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrontImage.Get(), o.FrontImage.IsSet()
}

// HasFrontImage returns a boolean if a field has been set.
func (o *DeviceType) HasFrontImage() bool {
	if o != nil && o.FrontImage.IsSet() {
		return true
	}

	return false
}

// SetFrontImage gets a reference to the given NullableString and assigns it to the FrontImage field.
func (o *DeviceType) SetFrontImage(v string) {
	o.FrontImage.Set(&v)
}
// SetFrontImageNil sets the value for FrontImage to be an explicit nil
func (o *DeviceType) SetFrontImageNil() {
	o.FrontImage.Set(nil)
}

// UnsetFrontImage ensures that no value is present for FrontImage, not even an explicit nil
func (o *DeviceType) UnsetFrontImage() {
	o.FrontImage.Unset()
}

// GetRearImage returns the RearImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetRearImage() string {
	if o == nil || IsNil(o.RearImage.Get()) {
		var ret string
		return ret
	}
	return *o.RearImage.Get()
}

// GetRearImageOk returns a tuple with the RearImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetRearImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RearImage.Get(), o.RearImage.IsSet()
}

// HasRearImage returns a boolean if a field has been set.
func (o *DeviceType) HasRearImage() bool {
	if o != nil && o.RearImage.IsSet() {
		return true
	}

	return false
}

// SetRearImage gets a reference to the given NullableString and assigns it to the RearImage field.
func (o *DeviceType) SetRearImage(v string) {
	o.RearImage.Set(&v)
}
// SetRearImageNil sets the value for RearImage to be an explicit nil
func (o *DeviceType) SetRearImageNil() {
	o.RearImage.Set(nil)
}

// UnsetRearImage ensures that no value is present for RearImage, not even an explicit nil
func (o *DeviceType) UnsetRearImage() {
	o.RearImage.Unset()
}

// GetDeviceCount returns the DeviceCount field value if set, zero value otherwise.
func (o *DeviceType) GetDeviceCount() int32 {
	if o == nil || IsNil(o.DeviceCount) {
		var ret int32
		return ret
	}
	return *o.DeviceCount
}

// GetDeviceCountOk returns a tuple with the DeviceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetDeviceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DeviceCount) {
		return nil, false
	}
	return o.DeviceCount, true
}

// HasDeviceCount returns a boolean if a field has been set.
func (o *DeviceType) HasDeviceCount() bool {
	if o != nil && !IsNil(o.DeviceCount) {
		return true
	}

	return false
}

// SetDeviceCount gets a reference to the given int32 and assigns it to the DeviceCount field.
func (o *DeviceType) SetDeviceCount(v int32) {
	o.DeviceCount = &v
}

// GetModel returns the Model field value
func (o *DeviceType) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeviceType) SetModel(v string) {
	o.Model = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *DeviceType) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *DeviceType) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *DeviceType) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetUHeight returns the UHeight field value if set, zero value otherwise.
func (o *DeviceType) GetUHeight() int32 {
	if o == nil || IsNil(o.UHeight) {
		var ret int32
		return ret
	}
	return *o.UHeight
}

// GetUHeightOk returns a tuple with the UHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetUHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.UHeight) {
		return nil, false
	}
	return o.UHeight, true
}

// HasUHeight returns a boolean if a field has been set.
func (o *DeviceType) HasUHeight() bool {
	if o != nil && !IsNil(o.UHeight) {
		return true
	}

	return false
}

// SetUHeight gets a reference to the given int32 and assigns it to the UHeight field.
func (o *DeviceType) SetUHeight(v int32) {
	o.UHeight = &v
}

// GetIsFullDepth returns the IsFullDepth field value if set, zero value otherwise.
func (o *DeviceType) GetIsFullDepth() bool {
	if o == nil || IsNil(o.IsFullDepth) {
		var ret bool
		return ret
	}
	return *o.IsFullDepth
}

// GetIsFullDepthOk returns a tuple with the IsFullDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetIsFullDepthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFullDepth) {
		return nil, false
	}
	return o.IsFullDepth, true
}

// HasIsFullDepth returns a boolean if a field has been set.
func (o *DeviceType) HasIsFullDepth() bool {
	if o != nil && !IsNil(o.IsFullDepth) {
		return true
	}

	return false
}

// SetIsFullDepth gets a reference to the given bool and assigns it to the IsFullDepth field.
func (o *DeviceType) SetIsFullDepth(v bool) {
	o.IsFullDepth = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *DeviceType) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *DeviceType) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *DeviceType) SetComments(v string) {
	o.Comments = &v
}

// GetManufacturer returns the Manufacturer field value
func (o *DeviceType) GetManufacturer() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetManufacturerOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *DeviceType) SetManufacturer(v BulkWritableCableRequestStatus) {
	o.Manufacturer = v
}

// GetDeviceFamily returns the DeviceFamily field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceType) GetDeviceFamily() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.DeviceFamily.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.DeviceFamily.Get()
}

// GetDeviceFamilyOk returns a tuple with the DeviceFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetDeviceFamilyOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceFamily.Get(), o.DeviceFamily.IsSet()
}

// HasDeviceFamily returns a boolean if a field has been set.
func (o *DeviceType) HasDeviceFamily() bool {
	if o != nil && o.DeviceFamily.IsSet() {
		return true
	}

	return false
}

// SetDeviceFamily gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the DeviceFamily field.
func (o *DeviceType) SetDeviceFamily(v BulkWritableCircuitRequestTenant) {
	o.DeviceFamily.Set(&v)
}
// SetDeviceFamilyNil sets the value for DeviceFamily to be an explicit nil
func (o *DeviceType) SetDeviceFamilyNil() {
	o.DeviceFamily.Set(nil)
}

// UnsetDeviceFamily ensures that no value is present for DeviceFamily, not even an explicit nil
func (o *DeviceType) UnsetDeviceFamily() {
	o.DeviceFamily.Unset()
}

// GetSoftwareImageFiles returns the SoftwareImageFiles field value
func (o *DeviceType) GetSoftwareImageFiles() []SoftwareImageFiles {
	if o == nil {
		var ret []SoftwareImageFiles
		return ret
	}

	return o.SoftwareImageFiles
}

// GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetSoftwareImageFilesOk() ([]SoftwareImageFiles, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoftwareImageFiles, true
}

// SetSoftwareImageFiles sets field value
func (o *DeviceType) SetSoftwareImageFiles(v []SoftwareImageFiles) {
	o.SoftwareImageFiles = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceType) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *DeviceType) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DeviceType) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceType) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *DeviceType) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *DeviceType) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *DeviceType) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *DeviceType) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceType) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceType) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceType) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceType) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceType) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceType) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *DeviceType) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o DeviceType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	if !IsNil(o.SubdeviceRole) {
		toSerialize["subdevice_role"] = o.SubdeviceRole
	}
	if o.FrontImage.IsSet() {
		toSerialize["front_image"] = o.FrontImage.Get()
	}
	if o.RearImage.IsSet() {
		toSerialize["rear_image"] = o.RearImage.Get()
	}
	if !IsNil(o.DeviceCount) {
		toSerialize["device_count"] = o.DeviceCount
	}
	toSerialize["model"] = o.Model
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.UHeight) {
		toSerialize["u_height"] = o.UHeight
	}
	if !IsNil(o.IsFullDepth) {
		toSerialize["is_full_depth"] = o.IsFullDepth
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["manufacturer"] = o.Manufacturer
	if o.DeviceFamily.IsSet() {
		toSerialize["device_family"] = o.DeviceFamily.Get()
	}
	toSerialize["software_image_files"] = o.SoftwareImageFiles
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"model",
		"manufacturer",
		"software_image_files",
		"created",
		"last_updated",
		"notes_url",
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

	varDeviceType := _DeviceType{}

	err = json.Unmarshal(data, &varDeviceType)

	if err != nil {
		return err
	}

	*o = DeviceType(varDeviceType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "subdevice_role")
		delete(additionalProperties, "front_image")
		delete(additionalProperties, "rear_image")
		delete(additionalProperties, "device_count")
		delete(additionalProperties, "model")
		delete(additionalProperties, "part_number")
		delete(additionalProperties, "u_height")
		delete(additionalProperties, "is_full_depth")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "device_family")
		delete(additionalProperties, "software_image_files")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceType struct {
	value *DeviceType
	isSet bool
}

func (v NullableDeviceType) Get() *DeviceType {
	return v.value
}

func (v *NullableDeviceType) Set(val *DeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceType(val *DeviceType) *NullableDeviceType {
	return &NullableDeviceType{value: val, isSet: true}
}

func (v NullableDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


