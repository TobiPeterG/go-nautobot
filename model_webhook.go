/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Webhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Webhook{}

// Webhook Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Webhook struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	ContentTypes []string `json:"content_types"`
	Name string `json:"name"`
	// Call this webhook when a matching object is created.
	TypeCreate *bool `json:"type_create,omitempty"`
	// Call this webhook when a matching object is updated.
	TypeUpdate *bool `json:"type_update,omitempty"`
	// Call this webhook when a matching object is deleted.
	TypeDelete *bool `json:"type_delete,omitempty"`
	// A POST will be sent to this URL when the webhook is called.
	PayloadUrl string `json:"payload_url"`
	Enabled *bool `json:"enabled,omitempty"`
	HttpMethod *HttpMethodEnum `json:"http_method,omitempty"`
	// The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.
	HttpContentType *string `json:"http_content_type,omitempty"`
	// User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).
	AdditionalHeaders *string `json:"additional_headers,omitempty"`
	// Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.
	BodyTemplate *string `json:"body_template,omitempty"`
	// When provided, the request will include a 'X-Hook-Signature' header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.
	Secret *string `json:"secret,omitempty"`
	// Enable SSL certificate verification. Disable with caution!
	SslVerification *bool `json:"ssl_verification,omitempty"`
	// The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.
	CaFilePath *string `json:"ca_file_path,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	AdditionalProperties map[string]interface{}
}

type _Webhook Webhook

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook(id string, objectType string, display string, url string, naturalSlug string, contentTypes []string, name string, payloadUrl string, created NullableTime, lastUpdated NullableTime, notesUrl string) *Webhook {
	this := Webhook{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.ContentTypes = contentTypes
	this.Name = name
	this.PayloadUrl = payloadUrl
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetId returns the Id field value
func (o *Webhook) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Webhook) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *Webhook) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *Webhook) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *Webhook) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Webhook) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *Webhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Webhook) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *Webhook) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *Webhook) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetContentTypes returns the ContentTypes field value
func (o *Webhook) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *Webhook) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *Webhook) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Webhook) SetName(v string) {
	o.Name = v
}

// GetTypeCreate returns the TypeCreate field value if set, zero value otherwise.
func (o *Webhook) GetTypeCreate() bool {
	if o == nil || IsNil(o.TypeCreate) {
		var ret bool
		return ret
	}
	return *o.TypeCreate
}

// GetTypeCreateOk returns a tuple with the TypeCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetTypeCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeCreate) {
		return nil, false
	}
	return o.TypeCreate, true
}

// HasTypeCreate returns a boolean if a field has been set.
func (o *Webhook) HasTypeCreate() bool {
	if o != nil && !IsNil(o.TypeCreate) {
		return true
	}

	return false
}

// SetTypeCreate gets a reference to the given bool and assigns it to the TypeCreate field.
func (o *Webhook) SetTypeCreate(v bool) {
	o.TypeCreate = &v
}

// GetTypeUpdate returns the TypeUpdate field value if set, zero value otherwise.
func (o *Webhook) GetTypeUpdate() bool {
	if o == nil || IsNil(o.TypeUpdate) {
		var ret bool
		return ret
	}
	return *o.TypeUpdate
}

// GetTypeUpdateOk returns a tuple with the TypeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetTypeUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeUpdate) {
		return nil, false
	}
	return o.TypeUpdate, true
}

// HasTypeUpdate returns a boolean if a field has been set.
func (o *Webhook) HasTypeUpdate() bool {
	if o != nil && !IsNil(o.TypeUpdate) {
		return true
	}

	return false
}

// SetTypeUpdate gets a reference to the given bool and assigns it to the TypeUpdate field.
func (o *Webhook) SetTypeUpdate(v bool) {
	o.TypeUpdate = &v
}

// GetTypeDelete returns the TypeDelete field value if set, zero value otherwise.
func (o *Webhook) GetTypeDelete() bool {
	if o == nil || IsNil(o.TypeDelete) {
		var ret bool
		return ret
	}
	return *o.TypeDelete
}

// GetTypeDeleteOk returns a tuple with the TypeDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetTypeDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeDelete) {
		return nil, false
	}
	return o.TypeDelete, true
}

// HasTypeDelete returns a boolean if a field has been set.
func (o *Webhook) HasTypeDelete() bool {
	if o != nil && !IsNil(o.TypeDelete) {
		return true
	}

	return false
}

// SetTypeDelete gets a reference to the given bool and assigns it to the TypeDelete field.
func (o *Webhook) SetTypeDelete(v bool) {
	o.TypeDelete = &v
}

// GetPayloadUrl returns the PayloadUrl field value
func (o *Webhook) GetPayloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadUrl
}

// GetPayloadUrlOk returns a tuple with the PayloadUrl field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetPayloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadUrl, true
}

// SetPayloadUrl sets field value
func (o *Webhook) SetPayloadUrl(v string) {
	o.PayloadUrl = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Webhook) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Webhook) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Webhook) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *Webhook) GetHttpMethod() HttpMethodEnum {
	if o == nil || IsNil(o.HttpMethod) {
		var ret HttpMethodEnum
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHttpMethodOk() (*HttpMethodEnum, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *Webhook) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given HttpMethodEnum and assigns it to the HttpMethod field.
func (o *Webhook) SetHttpMethod(v HttpMethodEnum) {
	o.HttpMethod = &v
}

// GetHttpContentType returns the HttpContentType field value if set, zero value otherwise.
func (o *Webhook) GetHttpContentType() string {
	if o == nil || IsNil(o.HttpContentType) {
		var ret string
		return ret
	}
	return *o.HttpContentType
}

// GetHttpContentTypeOk returns a tuple with the HttpContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHttpContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpContentType) {
		return nil, false
	}
	return o.HttpContentType, true
}

// HasHttpContentType returns a boolean if a field has been set.
func (o *Webhook) HasHttpContentType() bool {
	if o != nil && !IsNil(o.HttpContentType) {
		return true
	}

	return false
}

// SetHttpContentType gets a reference to the given string and assigns it to the HttpContentType field.
func (o *Webhook) SetHttpContentType(v string) {
	o.HttpContentType = &v
}

// GetAdditionalHeaders returns the AdditionalHeaders field value if set, zero value otherwise.
func (o *Webhook) GetAdditionalHeaders() string {
	if o == nil || IsNil(o.AdditionalHeaders) {
		var ret string
		return ret
	}
	return *o.AdditionalHeaders
}

// GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetAdditionalHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalHeaders) {
		return nil, false
	}
	return o.AdditionalHeaders, true
}

// HasAdditionalHeaders returns a boolean if a field has been set.
func (o *Webhook) HasAdditionalHeaders() bool {
	if o != nil && !IsNil(o.AdditionalHeaders) {
		return true
	}

	return false
}

// SetAdditionalHeaders gets a reference to the given string and assigns it to the AdditionalHeaders field.
func (o *Webhook) SetAdditionalHeaders(v string) {
	o.AdditionalHeaders = &v
}

// GetBodyTemplate returns the BodyTemplate field value if set, zero value otherwise.
func (o *Webhook) GetBodyTemplate() string {
	if o == nil || IsNil(o.BodyTemplate) {
		var ret string
		return ret
	}
	return *o.BodyTemplate
}

// GetBodyTemplateOk returns a tuple with the BodyTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetBodyTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.BodyTemplate) {
		return nil, false
	}
	return o.BodyTemplate, true
}

// HasBodyTemplate returns a boolean if a field has been set.
func (o *Webhook) HasBodyTemplate() bool {
	if o != nil && !IsNil(o.BodyTemplate) {
		return true
	}

	return false
}

// SetBodyTemplate gets a reference to the given string and assigns it to the BodyTemplate field.
func (o *Webhook) SetBodyTemplate(v string) {
	o.BodyTemplate = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *Webhook) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *Webhook) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *Webhook) SetSecret(v string) {
	o.Secret = &v
}

// GetSslVerification returns the SslVerification field value if set, zero value otherwise.
func (o *Webhook) GetSslVerification() bool {
	if o == nil || IsNil(o.SslVerification) {
		var ret bool
		return ret
	}
	return *o.SslVerification
}

// GetSslVerificationOk returns a tuple with the SslVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetSslVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SslVerification) {
		return nil, false
	}
	return o.SslVerification, true
}

// HasSslVerification returns a boolean if a field has been set.
func (o *Webhook) HasSslVerification() bool {
	if o != nil && !IsNil(o.SslVerification) {
		return true
	}

	return false
}

// SetSslVerification gets a reference to the given bool and assigns it to the SslVerification field.
func (o *Webhook) SetSslVerification(v bool) {
	o.SslVerification = &v
}

// GetCaFilePath returns the CaFilePath field value if set, zero value otherwise.
func (o *Webhook) GetCaFilePath() string {
	if o == nil || IsNil(o.CaFilePath) {
		var ret string
		return ret
	}
	return *o.CaFilePath
}

// GetCaFilePathOk returns a tuple with the CaFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetCaFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.CaFilePath) {
		return nil, false
	}
	return o.CaFilePath, true
}

// HasCaFilePath returns a boolean if a field has been set.
func (o *Webhook) HasCaFilePath() bool {
	if o != nil && !IsNil(o.CaFilePath) {
		return true
	}

	return false
}

// SetCaFilePath gets a reference to the given string and assigns it to the CaFilePath field.
func (o *Webhook) SetCaFilePath(v string) {
	o.CaFilePath = &v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Webhook) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webhook) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Webhook) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Webhook) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Webhook) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Webhook) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetNotesUrl returns the NotesUrl field value
func (o *Webhook) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *Webhook) SetNotesUrl(v string) {
	o.NotesUrl = v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.TypeCreate) {
		toSerialize["type_create"] = o.TypeCreate
	}
	if !IsNil(o.TypeUpdate) {
		toSerialize["type_update"] = o.TypeUpdate
	}
	if !IsNil(o.TypeDelete) {
		toSerialize["type_delete"] = o.TypeDelete
	}
	toSerialize["payload_url"] = o.PayloadUrl
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["http_method"] = o.HttpMethod
	}
	if !IsNil(o.HttpContentType) {
		toSerialize["http_content_type"] = o.HttpContentType
	}
	if !IsNil(o.AdditionalHeaders) {
		toSerialize["additional_headers"] = o.AdditionalHeaders
	}
	if !IsNil(o.BodyTemplate) {
		toSerialize["body_template"] = o.BodyTemplate
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.SslVerification) {
		toSerialize["ssl_verification"] = o.SslVerification
	}
	if !IsNil(o.CaFilePath) {
		toSerialize["ca_file_path"] = o.CaFilePath
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	toSerialize["notes_url"] = o.NotesUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Webhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"content_types",
		"name",
		"payload_url",
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

	varWebhook := _Webhook{}

	err = json.Unmarshal(data, &varWebhook)

	if err != nil {
		return err
	}

	*o = Webhook(varWebhook)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type_create")
		delete(additionalProperties, "type_update")
		delete(additionalProperties, "type_delete")
		delete(additionalProperties, "payload_url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "http_method")
		delete(additionalProperties, "http_content_type")
		delete(additionalProperties, "additional_headers")
		delete(additionalProperties, "body_template")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "ssl_verification")
		delete(additionalProperties, "ca_file_path")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "notes_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


