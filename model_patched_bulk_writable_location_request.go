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

// checks if the PatchedBulkWritableLocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableLocationRequest{}

// PatchedBulkWritableLocationRequest Add a `tree_depth` field to non-nested model serializers based on django-tree-queries.
type PatchedBulkWritableLocationRequest struct {
	Id string `json:"id"`
	TimeZone NullableString `json:"time_zone,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Local facility ID or description
	Facility *string `json:"facility,omitempty"`
	// 32-bit autonomous system number
	Asn NullableInt64 `json:"asn,omitempty"`
	PhysicalAddress *string `json:"physical_address,omitempty"`
	ShippingAddress *string `json:"shipping_address,omitempty"`
	// GPS coordinate (latitude)
	Latitude NullableFloat64 `json:"latitude,omitempty" validate:"regexp=^-?\\\\d{0,2}(?:\\\\.\\\\d{0,6})?$"`
	// GPS coordinate (longitude)
	Longitude NullableFloat64 `json:"longitude,omitempty" validate:"regexp=^-?\\\\d{0,3}(?:\\\\.\\\\d{0,6})?$"`
	ContactName *string `json:"contact_name,omitempty"`
	ContactPhone *string `json:"contact_phone,omitempty"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Parent NullableBulkWritableCircuitRequestTenant `json:"parent,omitempty"`
	LocationType *BulkWritableCableRequestStatus `json:"location_type,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableLocationRequest PatchedBulkWritableLocationRequest

// NewPatchedBulkWritableLocationRequest instantiates a new PatchedBulkWritableLocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableLocationRequest(id string) *PatchedBulkWritableLocationRequest {
	this := PatchedBulkWritableLocationRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableLocationRequestWithDefaults instantiates a new PatchedBulkWritableLocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableLocationRequestWithDefaults() *PatchedBulkWritableLocationRequest {
	this := PatchedBulkWritableLocationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableLocationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableLocationRequest) SetId(v string) {
	o.Id = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableLocationRequest) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone.Get()) {
		var ret string
		return ret
	}
	return *o.TimeZone.Get()
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableLocationRequest) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeZone.Get(), o.TimeZone.IsSet()
}

// HasTimeZone returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasTimeZone() bool {
	if o != nil && o.TimeZone.IsSet() {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given NullableString and assigns it to the TimeZone field.
func (o *PatchedBulkWritableLocationRequest) SetTimeZone(v string) {
	o.TimeZone.Set(&v)
}
// SetTimeZoneNil sets the value for TimeZone to be an explicit nil
func (o *PatchedBulkWritableLocationRequest) SetTimeZoneNil() {
	o.TimeZone.Set(nil)
}

// UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
func (o *PatchedBulkWritableLocationRequest) UnsetTimeZone() {
	o.TimeZone.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableLocationRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritableLocationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetFacility() string {
	if o == nil || IsNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetFacilityOk() (*string, bool) {
	if o == nil || IsNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasFacility() bool {
	if o != nil && !IsNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *PatchedBulkWritableLocationRequest) SetFacility(v string) {
	o.Facility = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableLocationRequest) GetAsn() int64 {
	if o == nil || IsNil(o.Asn.Get()) {
		var ret int64
		return ret
	}
	return *o.Asn.Get()
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableLocationRequest) GetAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asn.Get(), o.Asn.IsSet()
}

// HasAsn returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasAsn() bool {
	if o != nil && o.Asn.IsSet() {
		return true
	}

	return false
}

// SetAsn gets a reference to the given NullableInt64 and assigns it to the Asn field.
func (o *PatchedBulkWritableLocationRequest) SetAsn(v int64) {
	o.Asn.Set(&v)
}
// SetAsnNil sets the value for Asn to be an explicit nil
func (o *PatchedBulkWritableLocationRequest) SetAsnNil() {
	o.Asn.Set(nil)
}

// UnsetAsn ensures that no value is present for Asn, not even an explicit nil
func (o *PatchedBulkWritableLocationRequest) UnsetAsn() {
	o.Asn.Unset()
}

// GetPhysicalAddress returns the PhysicalAddress field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetPhysicalAddress() string {
	if o == nil || IsNil(o.PhysicalAddress) {
		var ret string
		return ret
	}
	return *o.PhysicalAddress
}

// GetPhysicalAddressOk returns a tuple with the PhysicalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetPhysicalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PhysicalAddress) {
		return nil, false
	}
	return o.PhysicalAddress, true
}

// HasPhysicalAddress returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasPhysicalAddress() bool {
	if o != nil && !IsNil(o.PhysicalAddress) {
		return true
	}

	return false
}

// SetPhysicalAddress gets a reference to the given string and assigns it to the PhysicalAddress field.
func (o *PatchedBulkWritableLocationRequest) SetPhysicalAddress(v string) {
	o.PhysicalAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetShippingAddress() string {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret string
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetShippingAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given string and assigns it to the ShippingAddress field.
func (o *PatchedBulkWritableLocationRequest) SetShippingAddress(v string) {
	o.ShippingAddress = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableLocationRequest) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableLocationRequest) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *PatchedBulkWritableLocationRequest) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *PatchedBulkWritableLocationRequest) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *PatchedBulkWritableLocationRequest) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableLocationRequest) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableLocationRequest) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *PatchedBulkWritableLocationRequest) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *PatchedBulkWritableLocationRequest) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *PatchedBulkWritableLocationRequest) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *PatchedBulkWritableLocationRequest) SetContactName(v string) {
	o.ContactName = &v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetContactPhone() string {
	if o == nil || IsNil(o.ContactPhone) {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ContactPhone) {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasContactPhone() bool {
	if o != nil && !IsNil(o.ContactPhone) {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *PatchedBulkWritableLocationRequest) SetContactPhone(v string) {
	o.ContactPhone = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *PatchedBulkWritableLocationRequest) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedBulkWritableLocationRequest) SetComments(v string) {
	o.Comments = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableLocationRequest) GetParent() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableLocationRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Parent field.
func (o *PatchedBulkWritableLocationRequest) SetParent(v BulkWritableCircuitRequestTenant) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedBulkWritableLocationRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedBulkWritableLocationRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetLocationType returns the LocationType field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetLocationType() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.LocationType) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetLocationTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.LocationType) {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasLocationType() bool {
	if o != nil && !IsNil(o.LocationType) {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given BulkWritableCableRequestStatus and assigns it to the LocationType field.
func (o *PatchedBulkWritableLocationRequest) SetLocationType(v BulkWritableCableRequestStatus) {
	o.LocationType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedBulkWritableLocationRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableLocationRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableLocationRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *PatchedBulkWritableLocationRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedBulkWritableLocationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedBulkWritableLocationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableLocationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableLocationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableLocationRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableLocationRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableLocationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableLocationRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableLocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableLocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.TimeZone.IsSet() {
		toSerialize["time_zone"] = o.TimeZone.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if o.Asn.IsSet() {
		toSerialize["asn"] = o.Asn.Get()
	}
	if !IsNil(o.PhysicalAddress) {
		toSerialize["physical_address"] = o.PhysicalAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if !IsNil(o.ContactName) {
		toSerialize["contact_name"] = o.ContactName
	}
	if !IsNil(o.ContactPhone) {
		toSerialize["contact_phone"] = o.ContactPhone
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.LocationType) {
		toSerialize["location_type"] = o.LocationType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
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

func (o *PatchedBulkWritableLocationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varPatchedBulkWritableLocationRequest := _PatchedBulkWritableLocationRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableLocationRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableLocationRequest(varPatchedBulkWritableLocationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "time_zone")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "facility")
		delete(additionalProperties, "asn")
		delete(additionalProperties, "physical_address")
		delete(additionalProperties, "shipping_address")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "contact_name")
		delete(additionalProperties, "contact_phone")
		delete(additionalProperties, "contact_email")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "location_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableLocationRequest struct {
	value *PatchedBulkWritableLocationRequest
	isSet bool
}

func (v NullablePatchedBulkWritableLocationRequest) Get() *PatchedBulkWritableLocationRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableLocationRequest) Set(val *PatchedBulkWritableLocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableLocationRequest(val *PatchedBulkWritableLocationRequest) *NullablePatchedBulkWritableLocationRequest {
	return &NullablePatchedBulkWritableLocationRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


