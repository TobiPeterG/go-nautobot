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

// checks if the JobResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobResult{}

// JobResult Extends ModelSerializer to render any CustomFields and their values associated with an object.
type JobResult struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Status JobResultStatus `json:"status"`
	Name string `json:"name"`
	// Registered name of the Celery task for this job. Internal use only.
	TaskName NullableString `json:"task_name,omitempty"`
	DateCreated time.Time `json:"date_created"`
	DateDone NullableTime `json:"date_done,omitempty"`
	// The data returned by the task
	Result interface{} `json:"result"`
	Worker NullableString `json:"worker,omitempty"`
	TaskArgs interface{} `json:"task_args,omitempty"`
	TaskKwargs interface{} `json:"task_kwargs,omitempty"`
	CeleryKwargs interface{} `json:"celery_kwargs,omitempty"`
	Traceback NullableString `json:"traceback,omitempty"`
	Meta interface{} `json:"meta"`
	JobModel NullableBulkWritableCircuitRequestTenant `json:"job_model,omitempty"`
	User NullableBulkWritableCircuitRequestTenant `json:"user,omitempty"`
	ScheduledJob NullableBulkWritableCircuitRequestTenant `json:"scheduled_job,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Files []BulkWritableCableRequestStatus `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _JobResult JobResult

// NewJobResult instantiates a new JobResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobResult(id string, objectType string, display string, url string, naturalSlug string, status JobResultStatus, name string, dateCreated time.Time, result interface{}, meta interface{}, files []BulkWritableCableRequestStatus) *JobResult {
	this := JobResult{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Status = status
	this.Name = name
	this.DateCreated = dateCreated
	this.Result = result
	this.Meta = meta
	this.Files = files
	return &this
}

// NewJobResultWithDefaults instantiates a new JobResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobResultWithDefaults() *JobResult {
	this := JobResult{}
	return &this
}

// GetId returns the Id field value
func (o *JobResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobResult) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *JobResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *JobResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *JobResult) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *JobResult) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *JobResult) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *JobResult) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *JobResult) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *JobResult) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetStatus returns the Status field value
func (o *JobResult) GetStatus() JobResultStatus {
	if o == nil {
		var ret JobResultStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetStatusOk() (*JobResultStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobResult) SetStatus(v JobResultStatus) {
	o.Status = v
}

// GetName returns the Name field value
func (o *JobResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobResult) SetName(v string) {
	o.Name = v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetTaskName() string {
	if o == nil || IsNil(o.TaskName.Get()) {
		var ret string
		return ret
	}
	return *o.TaskName.Get()
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskName.Get(), o.TaskName.IsSet()
}

// HasTaskName returns a boolean if a field has been set.
func (o *JobResult) HasTaskName() bool {
	if o != nil && o.TaskName.IsSet() {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given NullableString and assigns it to the TaskName field.
func (o *JobResult) SetTaskName(v string) {
	o.TaskName.Set(&v)
}
// SetTaskNameNil sets the value for TaskName to be an explicit nil
func (o *JobResult) SetTaskNameNil() {
	o.TaskName.Set(nil)
}

// UnsetTaskName ensures that no value is present for TaskName, not even an explicit nil
func (o *JobResult) UnsetTaskName() {
	o.TaskName.Unset()
}

// GetDateCreated returns the DateCreated field value
func (o *JobResult) GetDateCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *JobResult) SetDateCreated(v time.Time) {
	o.DateCreated = v
}

// GetDateDone returns the DateDone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetDateDone() time.Time {
	if o == nil || IsNil(o.DateDone.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateDone.Get()
}

// GetDateDoneOk returns a tuple with the DateDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetDateDoneOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateDone.Get(), o.DateDone.IsSet()
}

// HasDateDone returns a boolean if a field has been set.
func (o *JobResult) HasDateDone() bool {
	if o != nil && o.DateDone.IsSet() {
		return true
	}

	return false
}

// SetDateDone gets a reference to the given NullableTime and assigns it to the DateDone field.
func (o *JobResult) SetDateDone(v time.Time) {
	o.DateDone.Set(&v)
}
// SetDateDoneNil sets the value for DateDone to be an explicit nil
func (o *JobResult) SetDateDoneNil() {
	o.DateDone.Set(nil)
}

// UnsetDateDone ensures that no value is present for DateDone, not even an explicit nil
func (o *JobResult) UnsetDateDone() {
	o.DateDone.Unset()
}

// GetResult returns the Result field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JobResult) GetResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *JobResult) SetResult(v interface{}) {
	o.Result = v
}

// GetWorker returns the Worker field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetWorker() string {
	if o == nil || IsNil(o.Worker.Get()) {
		var ret string
		return ret
	}
	return *o.Worker.Get()
}

// GetWorkerOk returns a tuple with the Worker field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetWorkerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Worker.Get(), o.Worker.IsSet()
}

// HasWorker returns a boolean if a field has been set.
func (o *JobResult) HasWorker() bool {
	if o != nil && o.Worker.IsSet() {
		return true
	}

	return false
}

// SetWorker gets a reference to the given NullableString and assigns it to the Worker field.
func (o *JobResult) SetWorker(v string) {
	o.Worker.Set(&v)
}
// SetWorkerNil sets the value for Worker to be an explicit nil
func (o *JobResult) SetWorkerNil() {
	o.Worker.Set(nil)
}

// UnsetWorker ensures that no value is present for Worker, not even an explicit nil
func (o *JobResult) UnsetWorker() {
	o.Worker.Unset()
}

// GetTaskArgs returns the TaskArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetTaskArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaskArgs
}

// GetTaskArgsOk returns a tuple with the TaskArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetTaskArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TaskArgs) {
		return nil, false
	}
	return &o.TaskArgs, true
}

// HasTaskArgs returns a boolean if a field has been set.
func (o *JobResult) HasTaskArgs() bool {
	if o != nil && !IsNil(o.TaskArgs) {
		return true
	}

	return false
}

// SetTaskArgs gets a reference to the given interface{} and assigns it to the TaskArgs field.
func (o *JobResult) SetTaskArgs(v interface{}) {
	o.TaskArgs = v
}

// GetTaskKwargs returns the TaskKwargs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetTaskKwargs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaskKwargs
}

// GetTaskKwargsOk returns a tuple with the TaskKwargs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetTaskKwargsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TaskKwargs) {
		return nil, false
	}
	return &o.TaskKwargs, true
}

// HasTaskKwargs returns a boolean if a field has been set.
func (o *JobResult) HasTaskKwargs() bool {
	if o != nil && !IsNil(o.TaskKwargs) {
		return true
	}

	return false
}

// SetTaskKwargs gets a reference to the given interface{} and assigns it to the TaskKwargs field.
func (o *JobResult) SetTaskKwargs(v interface{}) {
	o.TaskKwargs = v
}

// GetCeleryKwargs returns the CeleryKwargs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetCeleryKwargs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CeleryKwargs
}

// GetCeleryKwargsOk returns a tuple with the CeleryKwargs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetCeleryKwargsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CeleryKwargs) {
		return nil, false
	}
	return &o.CeleryKwargs, true
}

// HasCeleryKwargs returns a boolean if a field has been set.
func (o *JobResult) HasCeleryKwargs() bool {
	if o != nil && !IsNil(o.CeleryKwargs) {
		return true
	}

	return false
}

// SetCeleryKwargs gets a reference to the given interface{} and assigns it to the CeleryKwargs field.
func (o *JobResult) SetCeleryKwargs(v interface{}) {
	o.CeleryKwargs = v
}

// GetTraceback returns the Traceback field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetTraceback() string {
	if o == nil || IsNil(o.Traceback.Get()) {
		var ret string
		return ret
	}
	return *o.Traceback.Get()
}

// GetTracebackOk returns a tuple with the Traceback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetTracebackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traceback.Get(), o.Traceback.IsSet()
}

// HasTraceback returns a boolean if a field has been set.
func (o *JobResult) HasTraceback() bool {
	if o != nil && o.Traceback.IsSet() {
		return true
	}

	return false
}

// SetTraceback gets a reference to the given NullableString and assigns it to the Traceback field.
func (o *JobResult) SetTraceback(v string) {
	o.Traceback.Set(&v)
}
// SetTracebackNil sets the value for Traceback to be an explicit nil
func (o *JobResult) SetTracebackNil() {
	o.Traceback.Set(nil)
}

// UnsetTraceback ensures that no value is present for Traceback, not even an explicit nil
func (o *JobResult) UnsetTraceback() {
	o.Traceback.Unset()
}

// GetMeta returns the Meta field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JobResult) GetMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetMetaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *JobResult) SetMeta(v interface{}) {
	o.Meta = v
}

// GetJobModel returns the JobModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetJobModel() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.JobModel.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.JobModel.Get()
}

// GetJobModelOk returns a tuple with the JobModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetJobModelOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobModel.Get(), o.JobModel.IsSet()
}

// HasJobModel returns a boolean if a field has been set.
func (o *JobResult) HasJobModel() bool {
	if o != nil && o.JobModel.IsSet() {
		return true
	}

	return false
}

// SetJobModel gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the JobModel field.
func (o *JobResult) SetJobModel(v BulkWritableCircuitRequestTenant) {
	o.JobModel.Set(&v)
}
// SetJobModelNil sets the value for JobModel to be an explicit nil
func (o *JobResult) SetJobModelNil() {
	o.JobModel.Set(nil)
}

// UnsetJobModel ensures that no value is present for JobModel, not even an explicit nil
func (o *JobResult) UnsetJobModel() {
	o.JobModel.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetUser() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.User.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetUserOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *JobResult) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the User field.
func (o *JobResult) SetUser(v BulkWritableCircuitRequestTenant) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *JobResult) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *JobResult) UnsetUser() {
	o.User.Unset()
}

// GetScheduledJob returns the ScheduledJob field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JobResult) GetScheduledJob() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ScheduledJob.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ScheduledJob.Get()
}

// GetScheduledJobOk returns a tuple with the ScheduledJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobResult) GetScheduledJobOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduledJob.Get(), o.ScheduledJob.IsSet()
}

// HasScheduledJob returns a boolean if a field has been set.
func (o *JobResult) HasScheduledJob() bool {
	if o != nil && o.ScheduledJob.IsSet() {
		return true
	}

	return false
}

// SetScheduledJob gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ScheduledJob field.
func (o *JobResult) SetScheduledJob(v BulkWritableCircuitRequestTenant) {
	o.ScheduledJob.Set(&v)
}
// SetScheduledJobNil sets the value for ScheduledJob to be an explicit nil
func (o *JobResult) SetScheduledJobNil() {
	o.ScheduledJob.Set(nil)
}

// UnsetScheduledJob ensures that no value is present for ScheduledJob, not even an explicit nil
func (o *JobResult) UnsetScheduledJob() {
	o.ScheduledJob.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *JobResult) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobResult) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *JobResult) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *JobResult) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetFiles returns the Files field value
func (o *JobResult) GetFiles() []BulkWritableCableRequestStatus {
	if o == nil {
		var ret []BulkWritableCableRequestStatus
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *JobResult) GetFilesOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *JobResult) SetFiles(v []BulkWritableCableRequestStatus) {
	o.Files = v
}

func (o JobResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["status"] = o.Status
	toSerialize["name"] = o.Name
	if o.TaskName.IsSet() {
		toSerialize["task_name"] = o.TaskName.Get()
	}
	toSerialize["date_created"] = o.DateCreated
	if o.DateDone.IsSet() {
		toSerialize["date_done"] = o.DateDone.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Worker.IsSet() {
		toSerialize["worker"] = o.Worker.Get()
	}
	if o.TaskArgs != nil {
		toSerialize["task_args"] = o.TaskArgs
	}
	if o.TaskKwargs != nil {
		toSerialize["task_kwargs"] = o.TaskKwargs
	}
	if o.CeleryKwargs != nil {
		toSerialize["celery_kwargs"] = o.CeleryKwargs
	}
	if o.Traceback.IsSet() {
		toSerialize["traceback"] = o.Traceback.Get()
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.JobModel.IsSet() {
		toSerialize["job_model"] = o.JobModel.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.ScheduledJob.IsSet() {
		toSerialize["scheduled_job"] = o.ScheduledJob.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JobResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"status",
		"name",
		"date_created",
		"result",
		"meta",
		"files",
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

	varJobResult := _JobResult{}

	err = json.Unmarshal(data, &varJobResult)

	if err != nil {
		return err
	}

	*o = JobResult(varJobResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "status")
		delete(additionalProperties, "name")
		delete(additionalProperties, "task_name")
		delete(additionalProperties, "date_created")
		delete(additionalProperties, "date_done")
		delete(additionalProperties, "result")
		delete(additionalProperties, "worker")
		delete(additionalProperties, "task_args")
		delete(additionalProperties, "task_kwargs")
		delete(additionalProperties, "celery_kwargs")
		delete(additionalProperties, "traceback")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "job_model")
		delete(additionalProperties, "user")
		delete(additionalProperties, "scheduled_job")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJobResult struct {
	value *JobResult
	isSet bool
}

func (v NullableJobResult) Get() *JobResult {
	return v.value
}

func (v *NullableJobResult) Set(val *JobResult) {
	v.value = val
	v.isSet = true
}

func (v NullableJobResult) IsSet() bool {
	return v.isSet
}

func (v *NullableJobResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobResult(val *JobResult) *NullableJobResult {
	return &NullableJobResult{value: val, isSet: true}
}

func (v NullableJobResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


