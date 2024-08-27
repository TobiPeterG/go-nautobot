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

// checks if the ScheduledJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledJob{}

// ScheduledJob This base serializer implements common fields and logic for all ModelSerializers.  Namely, it:  - defines the `display` field which exposes a human friendly value for the given object. - ensures that `id` field is always present on the serializer as well. - ensures that `created` and `last_updated` fields are always present if applicable to this model and serializer. - ensures that `object_type` field is always present on the serializer which represents the content-type of this   serializer's associated model (e.g. \"dcim.device\"). This is required as the OpenAPI schema, using the   PolymorphicProxySerializer class defined below, relies upon this field as a way to identify to the client   which of several possible serializers are in use for a given attribute. - supports `?depth` query parameter. It is passed in as `nested_depth` to the `build_nested_field()` function   to enable the dynamic generation of nested serializers.
type ScheduledJob struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	// Human-readable description of this scheduled task
	Name string `json:"name"`
	// The name of the Celery task that should be run. (Example: \"proj.tasks.import_contacts\")
	Task string `json:"task"`
	Interval JobExecutionTypeIntervalChoices `json:"interval"`
	Args interface{} `json:"args,omitempty"`
	Kwargs interface{} `json:"kwargs,omitempty"`
	CeleryKwargs interface{} `json:"celery_kwargs,omitempty"`
	// Queue defined in CELERY_TASK_QUEUES. Leave empty for default queuing.
	Queue *string `json:"queue,omitempty"`
	// If True, the schedule will only run the task a single time
	OneOff *bool `json:"one_off,omitempty"`
	// Datetime when the schedule should begin triggering the task to run
	StartTime time.Time `json:"start_time"`
	// Set to False to disable the schedule
	Enabled *bool `json:"enabled,omitempty"`
	// Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False.
	LastRunAt NullableTime `json:"last_run_at"`
	// Running count of how many times the schedule has triggered the task
	TotalRunCount int32 `json:"total_run_count"`
	// Datetime that this scheduled job was last modified
	DateChanged time.Time `json:"date_changed"`
	// Detailed description about the details of this scheduled job
	Description *string `json:"description,omitempty"`
	ApprovalRequired *bool `json:"approval_required,omitempty"`
	// Datetime that the schedule was approved
	ApprovedAt NullableTime `json:"approved_at"`
	// Cronjob syntax string for custom scheduling
	Crontab *string `json:"crontab,omitempty"`
	JobModel NullableBulkWritableCircuitRequestTenant `json:"job_model,omitempty"`
	User NullableScheduledJobUser `json:"user,omitempty"`
	ApprovedByUser NullableScheduledJobApprovedByUser `json:"approved_by_user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledJob ScheduledJob

// NewScheduledJob instantiates a new ScheduledJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledJob(id string, objectType string, display string, url string, naturalSlug string, name string, task string, interval JobExecutionTypeIntervalChoices, startTime time.Time, lastRunAt NullableTime, totalRunCount int32, dateChanged time.Time, approvedAt NullableTime) *ScheduledJob {
	this := ScheduledJob{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Name = name
	this.Task = task
	this.Interval = interval
	this.StartTime = startTime
	this.LastRunAt = lastRunAt
	this.TotalRunCount = totalRunCount
	this.DateChanged = dateChanged
	this.ApprovedAt = approvedAt
	return &this
}

// NewScheduledJobWithDefaults instantiates a new ScheduledJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledJobWithDefaults() *ScheduledJob {
	this := ScheduledJob{}
	return &this
}

// GetId returns the Id field value
func (o *ScheduledJob) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScheduledJob) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *ScheduledJob) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ScheduledJob) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *ScheduledJob) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ScheduledJob) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *ScheduledJob) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ScheduledJob) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *ScheduledJob) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *ScheduledJob) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetName returns the Name field value
func (o *ScheduledJob) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScheduledJob) SetName(v string) {
	o.Name = v
}

// GetTask returns the Task field value
func (o *ScheduledJob) GetTask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Task
}

// GetTaskOk returns a tuple with the Task field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Task, true
}

// SetTask sets field value
func (o *ScheduledJob) SetTask(v string) {
	o.Task = v
}

// GetInterval returns the Interval field value
func (o *ScheduledJob) GetInterval() JobExecutionTypeIntervalChoices {
	if o == nil {
		var ret JobExecutionTypeIntervalChoices
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetIntervalOk() (*JobExecutionTypeIntervalChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ScheduledJob) SetInterval(v JobExecutionTypeIntervalChoices) {
	o.Interval = v
}

// GetArgs returns the Args field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledJob) GetArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return &o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ScheduledJob) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given interface{} and assigns it to the Args field.
func (o *ScheduledJob) SetArgs(v interface{}) {
	o.Args = v
}

// GetKwargs returns the Kwargs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledJob) GetKwargs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Kwargs
}

// GetKwargsOk returns a tuple with the Kwargs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetKwargsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Kwargs) {
		return nil, false
	}
	return &o.Kwargs, true
}

// HasKwargs returns a boolean if a field has been set.
func (o *ScheduledJob) HasKwargs() bool {
	if o != nil && !IsNil(o.Kwargs) {
		return true
	}

	return false
}

// SetKwargs gets a reference to the given interface{} and assigns it to the Kwargs field.
func (o *ScheduledJob) SetKwargs(v interface{}) {
	o.Kwargs = v
}

// GetCeleryKwargs returns the CeleryKwargs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledJob) GetCeleryKwargs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CeleryKwargs
}

// GetCeleryKwargsOk returns a tuple with the CeleryKwargs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetCeleryKwargsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CeleryKwargs) {
		return nil, false
	}
	return &o.CeleryKwargs, true
}

// HasCeleryKwargs returns a boolean if a field has been set.
func (o *ScheduledJob) HasCeleryKwargs() bool {
	if o != nil && !IsNil(o.CeleryKwargs) {
		return true
	}

	return false
}

// SetCeleryKwargs gets a reference to the given interface{} and assigns it to the CeleryKwargs field.
func (o *ScheduledJob) SetCeleryKwargs(v interface{}) {
	o.CeleryKwargs = v
}

// GetQueue returns the Queue field value if set, zero value otherwise.
func (o *ScheduledJob) GetQueue() string {
	if o == nil || IsNil(o.Queue) {
		var ret string
		return ret
	}
	return *o.Queue
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetQueueOk() (*string, bool) {
	if o == nil || IsNil(o.Queue) {
		return nil, false
	}
	return o.Queue, true
}

// HasQueue returns a boolean if a field has been set.
func (o *ScheduledJob) HasQueue() bool {
	if o != nil && !IsNil(o.Queue) {
		return true
	}

	return false
}

// SetQueue gets a reference to the given string and assigns it to the Queue field.
func (o *ScheduledJob) SetQueue(v string) {
	o.Queue = &v
}

// GetOneOff returns the OneOff field value if set, zero value otherwise.
func (o *ScheduledJob) GetOneOff() bool {
	if o == nil || IsNil(o.OneOff) {
		var ret bool
		return ret
	}
	return *o.OneOff
}

// GetOneOffOk returns a tuple with the OneOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetOneOffOk() (*bool, bool) {
	if o == nil || IsNil(o.OneOff) {
		return nil, false
	}
	return o.OneOff, true
}

// HasOneOff returns a boolean if a field has been set.
func (o *ScheduledJob) HasOneOff() bool {
	if o != nil && !IsNil(o.OneOff) {
		return true
	}

	return false
}

// SetOneOff gets a reference to the given bool and assigns it to the OneOff field.
func (o *ScheduledJob) SetOneOff(v bool) {
	o.OneOff = &v
}

// GetStartTime returns the StartTime field value
func (o *ScheduledJob) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ScheduledJob) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScheduledJob) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScheduledJob) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScheduledJob) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLastRunAt returns the LastRunAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ScheduledJob) GetLastRunAt() time.Time {
	if o == nil || o.LastRunAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastRunAt.Get()
}

// GetLastRunAtOk returns a tuple with the LastRunAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetLastRunAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRunAt.Get(), o.LastRunAt.IsSet()
}

// SetLastRunAt sets field value
func (o *ScheduledJob) SetLastRunAt(v time.Time) {
	o.LastRunAt.Set(&v)
}

// GetTotalRunCount returns the TotalRunCount field value
func (o *ScheduledJob) GetTotalRunCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalRunCount
}

// GetTotalRunCountOk returns a tuple with the TotalRunCount field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetTotalRunCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRunCount, true
}

// SetTotalRunCount sets field value
func (o *ScheduledJob) SetTotalRunCount(v int32) {
	o.TotalRunCount = v
}

// GetDateChanged returns the DateChanged field value
func (o *ScheduledJob) GetDateChanged() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DateChanged
}

// GetDateChangedOk returns a tuple with the DateChanged field value
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetDateChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateChanged, true
}

// SetDateChanged sets field value
func (o *ScheduledJob) SetDateChanged(v time.Time) {
	o.DateChanged = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScheduledJob) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScheduledJob) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScheduledJob) SetDescription(v string) {
	o.Description = &v
}

// GetApprovalRequired returns the ApprovalRequired field value if set, zero value otherwise.
func (o *ScheduledJob) GetApprovalRequired() bool {
	if o == nil || IsNil(o.ApprovalRequired) {
		var ret bool
		return ret
	}
	return *o.ApprovalRequired
}

// GetApprovalRequiredOk returns a tuple with the ApprovalRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetApprovalRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ApprovalRequired) {
		return nil, false
	}
	return o.ApprovalRequired, true
}

// HasApprovalRequired returns a boolean if a field has been set.
func (o *ScheduledJob) HasApprovalRequired() bool {
	if o != nil && !IsNil(o.ApprovalRequired) {
		return true
	}

	return false
}

// SetApprovalRequired gets a reference to the given bool and assigns it to the ApprovalRequired field.
func (o *ScheduledJob) SetApprovalRequired(v bool) {
	o.ApprovalRequired = &v
}

// GetApprovedAt returns the ApprovedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ScheduledJob) GetApprovedAt() time.Time {
	if o == nil || o.ApprovedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ApprovedAt.Get()
}

// GetApprovedAtOk returns a tuple with the ApprovedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetApprovedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovedAt.Get(), o.ApprovedAt.IsSet()
}

// SetApprovedAt sets field value
func (o *ScheduledJob) SetApprovedAt(v time.Time) {
	o.ApprovedAt.Set(&v)
}

// GetCrontab returns the Crontab field value if set, zero value otherwise.
func (o *ScheduledJob) GetCrontab() string {
	if o == nil || IsNil(o.Crontab) {
		var ret string
		return ret
	}
	return *o.Crontab
}

// GetCrontabOk returns a tuple with the Crontab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledJob) GetCrontabOk() (*string, bool) {
	if o == nil || IsNil(o.Crontab) {
		return nil, false
	}
	return o.Crontab, true
}

// HasCrontab returns a boolean if a field has been set.
func (o *ScheduledJob) HasCrontab() bool {
	if o != nil && !IsNil(o.Crontab) {
		return true
	}

	return false
}

// SetCrontab gets a reference to the given string and assigns it to the Crontab field.
func (o *ScheduledJob) SetCrontab(v string) {
	o.Crontab = &v
}

// GetJobModel returns the JobModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledJob) GetJobModel() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.JobModel.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.JobModel.Get()
}

// GetJobModelOk returns a tuple with the JobModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetJobModelOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.JobModel.Get(), o.JobModel.IsSet()
}

// HasJobModel returns a boolean if a field has been set.
func (o *ScheduledJob) HasJobModel() bool {
	if o != nil && o.JobModel.IsSet() {
		return true
	}

	return false
}

// SetJobModel gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the JobModel field.
func (o *ScheduledJob) SetJobModel(v BulkWritableCircuitRequestTenant) {
	o.JobModel.Set(&v)
}
// SetJobModelNil sets the value for JobModel to be an explicit nil
func (o *ScheduledJob) SetJobModelNil() {
	o.JobModel.Set(nil)
}

// UnsetJobModel ensures that no value is present for JobModel, not even an explicit nil
func (o *ScheduledJob) UnsetJobModel() {
	o.JobModel.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledJob) GetUser() ScheduledJobUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret ScheduledJobUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetUserOk() (*ScheduledJobUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ScheduledJob) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableScheduledJobUser and assigns it to the User field.
func (o *ScheduledJob) SetUser(v ScheduledJobUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *ScheduledJob) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ScheduledJob) UnsetUser() {
	o.User.Unset()
}

// GetApprovedByUser returns the ApprovedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledJob) GetApprovedByUser() ScheduledJobApprovedByUser {
	if o == nil || IsNil(o.ApprovedByUser.Get()) {
		var ret ScheduledJobApprovedByUser
		return ret
	}
	return *o.ApprovedByUser.Get()
}

// GetApprovedByUserOk returns a tuple with the ApprovedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledJob) GetApprovedByUserOk() (*ScheduledJobApprovedByUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApprovedByUser.Get(), o.ApprovedByUser.IsSet()
}

// HasApprovedByUser returns a boolean if a field has been set.
func (o *ScheduledJob) HasApprovedByUser() bool {
	if o != nil && o.ApprovedByUser.IsSet() {
		return true
	}

	return false
}

// SetApprovedByUser gets a reference to the given NullableScheduledJobApprovedByUser and assigns it to the ApprovedByUser field.
func (o *ScheduledJob) SetApprovedByUser(v ScheduledJobApprovedByUser) {
	o.ApprovedByUser.Set(&v)
}
// SetApprovedByUserNil sets the value for ApprovedByUser to be an explicit nil
func (o *ScheduledJob) SetApprovedByUserNil() {
	o.ApprovedByUser.Set(nil)
}

// UnsetApprovedByUser ensures that no value is present for ApprovedByUser, not even an explicit nil
func (o *ScheduledJob) UnsetApprovedByUser() {
	o.ApprovedByUser.Unset()
}

func (o ScheduledJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["name"] = o.Name
	toSerialize["task"] = o.Task
	toSerialize["interval"] = o.Interval
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if o.Kwargs != nil {
		toSerialize["kwargs"] = o.Kwargs
	}
	if o.CeleryKwargs != nil {
		toSerialize["celery_kwargs"] = o.CeleryKwargs
	}
	if !IsNil(o.Queue) {
		toSerialize["queue"] = o.Queue
	}
	if !IsNil(o.OneOff) {
		toSerialize["one_off"] = o.OneOff
	}
	toSerialize["start_time"] = o.StartTime
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["last_run_at"] = o.LastRunAt.Get()
	toSerialize["total_run_count"] = o.TotalRunCount
	toSerialize["date_changed"] = o.DateChanged
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ApprovalRequired) {
		toSerialize["approval_required"] = o.ApprovalRequired
	}
	toSerialize["approved_at"] = o.ApprovedAt.Get()
	if !IsNil(o.Crontab) {
		toSerialize["crontab"] = o.Crontab
	}
	if o.JobModel.IsSet() {
		toSerialize["job_model"] = o.JobModel.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.ApprovedByUser.IsSet() {
		toSerialize["approved_by_user"] = o.ApprovedByUser.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduledJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"name",
		"task",
		"interval",
		"start_time",
		"last_run_at",
		"total_run_count",
		"date_changed",
		"approved_at",
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

	varScheduledJob := _ScheduledJob{}

	err = json.Unmarshal(data, &varScheduledJob)

	if err != nil {
		return err
	}

	*o = ScheduledJob(varScheduledJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "name")
		delete(additionalProperties, "task")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "args")
		delete(additionalProperties, "kwargs")
		delete(additionalProperties, "celery_kwargs")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "one_off")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "last_run_at")
		delete(additionalProperties, "total_run_count")
		delete(additionalProperties, "date_changed")
		delete(additionalProperties, "description")
		delete(additionalProperties, "approval_required")
		delete(additionalProperties, "approved_at")
		delete(additionalProperties, "crontab")
		delete(additionalProperties, "job_model")
		delete(additionalProperties, "user")
		delete(additionalProperties, "approved_by_user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledJob struct {
	value *ScheduledJob
	isSet bool
}

func (v NullableScheduledJob) Get() *ScheduledJob {
	return v.value
}

func (v *NullableScheduledJob) Set(val *ScheduledJob) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledJob) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledJob(val *ScheduledJob) *NullableScheduledJob {
	return &NullableScheduledJob{value: val, isSet: true}
}

func (v NullableScheduledJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


