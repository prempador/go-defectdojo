/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the TestCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestCreate{}

// TestCreate struct for TestCreate
type TestCreate struct {
	Id int32 `json:"id"`
	Engagement int32 `json:"engagement"`
	Notes []*int32 `json:"notes,omitempty"`
	Tags []string `json:"tags,omitempty"`
	ScanType NullableString `json:"scan_type,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	TargetStart time.Time `json:"target_start"`
	TargetEnd time.Time `json:"target_end"`
	EstimatedTime NullableString `json:"estimated_time"`
	ActualTime NullableString `json:"actual_time"`
	PercentComplete NullableInt32 `json:"percent_complete,omitempty"`
	Updated NullableTime `json:"updated"`
	Created NullableTime `json:"created"`
	Version NullableString `json:"version,omitempty"`
	// Build ID that was tested, a reimport may update this field.
	BuildId NullableString `json:"build_id,omitempty"`
	// Commit hash tested, a reimport may update this field.
	CommitHash NullableString `json:"commit_hash,omitempty"`
	// Tag or branch that was tested, a reimport may update this field.
	BranchTag NullableString `json:"branch_tag,omitempty"`
	Lead NullableInt32 `json:"lead,omitempty"`
	TestType int32 `json:"test_type"`
	Environment NullableInt32 `json:"environment,omitempty"`
	ApiScanConfiguration NullableInt32 `json:"api_scan_configuration,omitempty"`
	Files []int32 `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _TestCreate TestCreate

// NewTestCreate instantiates a new TestCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestCreate(id int32, engagement int32, targetStart time.Time, targetEnd time.Time, estimatedTime NullableString, actualTime NullableString, updated NullableTime, created NullableTime, testType int32, files []int32) *TestCreate {
	this := TestCreate{}
	this.Id = id
	this.Engagement = engagement
	this.TargetStart = targetStart
	this.TargetEnd = targetEnd
	this.EstimatedTime = estimatedTime
	this.ActualTime = actualTime
	this.Updated = updated
	this.Created = created
	this.TestType = testType
	this.Files = files
	return &this
}

// NewTestCreateWithDefaults instantiates a new TestCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestCreateWithDefaults() *TestCreate {
	this := TestCreate{}
	return &this
}

// GetId returns the Id field value
func (o *TestCreate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestCreate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestCreate) SetId(v int32) {
	o.Id = v
}


// GetEngagement returns the Engagement field value
func (o *TestCreate) GetEngagement() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Engagement
}

// GetEngagementOk returns a tuple with the Engagement field value
// and a boolean to check if the value has been set.
func (o *TestCreate) GetEngagementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engagement, true
}

// SetEngagement sets field value
func (o *TestCreate) SetEngagement(v int32) {
	o.Engagement = v
}


// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *TestCreate) GetNotes() []*int32 {
	if o == nil || IsNil(o.Notes) {
		var ret []*int32
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCreate) GetNotesOk() ([]*int32, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *TestCreate) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []*int32 and assigns it to the Notes field.
func (o *TestCreate) SetNotes(v []*int32) {
	o.Notes = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TestCreate) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestCreate) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestCreate) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TestCreate) SetTags(v []string) {
	o.Tags = v
}

// GetScanType returns the ScanType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetScanType() string {
	if o == nil || IsNil(o.ScanType.Get()) {
		var ret string
		return ret
	}
	return *o.ScanType.Get()
}

// GetScanTypeOk returns a tuple with the ScanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScanType.Get(), o.ScanType.IsSet()
}

// HasScanType returns a boolean if a field has been set.
func (o *TestCreate) HasScanType() bool {
	if o != nil && o.ScanType.IsSet() {
		return true
	}

	return false
}

// SetScanType gets a reference to the given NullableString and assigns it to the ScanType field.
func (o *TestCreate) SetScanType(v string) {
	o.ScanType.Set(&v)
}
// SetScanTypeNil sets the value for ScanType to be an explicit nil
func (o *TestCreate) SetScanTypeNil() {
	o.ScanType.Set(nil)
}

// UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
func (o *TestCreate) UnsetScanType() {
	o.ScanType.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TestCreate) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TestCreate) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TestCreate) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TestCreate) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestCreate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestCreate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestCreate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestCreate) UnsetDescription() {
	o.Description.Unset()
}

// GetTargetStart returns the TargetStart field value
func (o *TestCreate) GetTargetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TargetStart
}

// GetTargetStartOk returns a tuple with the TargetStart field value
// and a boolean to check if the value has been set.
func (o *TestCreate) GetTargetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetStart, true
}

// SetTargetStart sets field value
func (o *TestCreate) SetTargetStart(v time.Time) {
	o.TargetStart = v
}


// GetTargetEnd returns the TargetEnd field value
func (o *TestCreate) GetTargetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TargetEnd
}

// GetTargetEndOk returns a tuple with the TargetEnd field value
// and a boolean to check if the value has been set.
func (o *TestCreate) GetTargetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnd, true
}

// SetTargetEnd sets field value
func (o *TestCreate) SetTargetEnd(v time.Time) {
	o.TargetEnd = v
}


// GetEstimatedTime returns the EstimatedTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TestCreate) GetEstimatedTime() string {
	if o == nil || o.EstimatedTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.EstimatedTime.Get()
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetEstimatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedTime.Get(), o.EstimatedTime.IsSet()
}

// SetEstimatedTime sets field value
func (o *TestCreate) SetEstimatedTime(v string) {
	o.EstimatedTime.Set(&v)
}


// GetActualTime returns the ActualTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TestCreate) GetActualTime() string {
	if o == nil || o.ActualTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActualTime.Get()
}

// GetActualTimeOk returns a tuple with the ActualTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetActualTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualTime.Get(), o.ActualTime.IsSet()
}

// SetActualTime sets field value
func (o *TestCreate) SetActualTime(v string) {
	o.ActualTime.Set(&v)
}


// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetPercentComplete() int32 {
	if o == nil || IsNil(o.PercentComplete.Get()) {
		var ret int32
		return ret
	}
	return *o.PercentComplete.Get()
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetPercentCompleteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentComplete.Get(), o.PercentComplete.IsSet()
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *TestCreate) HasPercentComplete() bool {
	if o != nil && o.PercentComplete.IsSet() {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given NullableInt32 and assigns it to the PercentComplete field.
func (o *TestCreate) SetPercentComplete(v int32) {
	o.PercentComplete.Set(&v)
}
// SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil
func (o *TestCreate) SetPercentCompleteNil() {
	o.PercentComplete.Set(nil)
}

// UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
func (o *TestCreate) UnsetPercentComplete() {
	o.PercentComplete.Unset()
}

// GetUpdated returns the Updated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TestCreate) GetUpdated() time.Time {
	if o == nil || o.Updated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// SetUpdated sets field value
func (o *TestCreate) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}


// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *TestCreate) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *TestCreate) SetCreated(v time.Time) {
	o.Created.Set(&v)
}


// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *TestCreate) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *TestCreate) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *TestCreate) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *TestCreate) UnsetVersion() {
	o.Version.Unset()
}

// GetBuildId returns the BuildId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetBuildId() string {
	if o == nil || IsNil(o.BuildId.Get()) {
		var ret string
		return ret
	}
	return *o.BuildId.Get()
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetBuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildId.Get(), o.BuildId.IsSet()
}

// HasBuildId returns a boolean if a field has been set.
func (o *TestCreate) HasBuildId() bool {
	if o != nil && o.BuildId.IsSet() {
		return true
	}

	return false
}

// SetBuildId gets a reference to the given NullableString and assigns it to the BuildId field.
func (o *TestCreate) SetBuildId(v string) {
	o.BuildId.Set(&v)
}
// SetBuildIdNil sets the value for BuildId to be an explicit nil
func (o *TestCreate) SetBuildIdNil() {
	o.BuildId.Set(nil)
}

// UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
func (o *TestCreate) UnsetBuildId() {
	o.BuildId.Unset()
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash.Get()) {
		var ret string
		return ret
	}
	return *o.CommitHash.Get()
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitHash.Get(), o.CommitHash.IsSet()
}

// HasCommitHash returns a boolean if a field has been set.
func (o *TestCreate) HasCommitHash() bool {
	if o != nil && o.CommitHash.IsSet() {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullableString and assigns it to the CommitHash field.
func (o *TestCreate) SetCommitHash(v string) {
	o.CommitHash.Set(&v)
}
// SetCommitHashNil sets the value for CommitHash to be an explicit nil
func (o *TestCreate) SetCommitHashNil() {
	o.CommitHash.Set(nil)
}

// UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
func (o *TestCreate) UnsetCommitHash() {
	o.CommitHash.Unset()
}

// GetBranchTag returns the BranchTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetBranchTag() string {
	if o == nil || IsNil(o.BranchTag.Get()) {
		var ret string
		return ret
	}
	return *o.BranchTag.Get()
}

// GetBranchTagOk returns a tuple with the BranchTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetBranchTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchTag.Get(), o.BranchTag.IsSet()
}

// HasBranchTag returns a boolean if a field has been set.
func (o *TestCreate) HasBranchTag() bool {
	if o != nil && o.BranchTag.IsSet() {
		return true
	}

	return false
}

// SetBranchTag gets a reference to the given NullableString and assigns it to the BranchTag field.
func (o *TestCreate) SetBranchTag(v string) {
	o.BranchTag.Set(&v)
}
// SetBranchTagNil sets the value for BranchTag to be an explicit nil
func (o *TestCreate) SetBranchTagNil() {
	o.BranchTag.Set(nil)
}

// UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
func (o *TestCreate) UnsetBranchTag() {
	o.BranchTag.Unset()
}

// GetLead returns the Lead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetLead() int32 {
	if o == nil || IsNil(o.Lead.Get()) {
		var ret int32
		return ret
	}
	return *o.Lead.Get()
}

// GetLeadOk returns a tuple with the Lead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetLeadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lead.Get(), o.Lead.IsSet()
}

// HasLead returns a boolean if a field has been set.
func (o *TestCreate) HasLead() bool {
	if o != nil && o.Lead.IsSet() {
		return true
	}

	return false
}

// SetLead gets a reference to the given NullableInt32 and assigns it to the Lead field.
func (o *TestCreate) SetLead(v int32) {
	o.Lead.Set(&v)
}
// SetLeadNil sets the value for Lead to be an explicit nil
func (o *TestCreate) SetLeadNil() {
	o.Lead.Set(nil)
}

// UnsetLead ensures that no value is present for Lead, not even an explicit nil
func (o *TestCreate) UnsetLead() {
	o.Lead.Unset()
}

// GetTestType returns the TestType field value
func (o *TestCreate) GetTestType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value
// and a boolean to check if the value has been set.
func (o *TestCreate) GetTestTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestType, true
}

// SetTestType sets field value
func (o *TestCreate) SetTestType(v int32) {
	o.TestType = v
}


// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetEnvironment() int32 {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret int32
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetEnvironmentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *TestCreate) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableInt32 and assigns it to the Environment field.
func (o *TestCreate) SetEnvironment(v int32) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *TestCreate) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *TestCreate) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetApiScanConfiguration returns the ApiScanConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestCreate) GetApiScanConfiguration() int32 {
	if o == nil || IsNil(o.ApiScanConfiguration.Get()) {
		var ret int32
		return ret
	}
	return *o.ApiScanConfiguration.Get()
}

// GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestCreate) GetApiScanConfigurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiScanConfiguration.Get(), o.ApiScanConfiguration.IsSet()
}

// HasApiScanConfiguration returns a boolean if a field has been set.
func (o *TestCreate) HasApiScanConfiguration() bool {
	if o != nil && o.ApiScanConfiguration.IsSet() {
		return true
	}

	return false
}

// SetApiScanConfiguration gets a reference to the given NullableInt32 and assigns it to the ApiScanConfiguration field.
func (o *TestCreate) SetApiScanConfiguration(v int32) {
	o.ApiScanConfiguration.Set(&v)
}
// SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil
func (o *TestCreate) SetApiScanConfigurationNil() {
	o.ApiScanConfiguration.Set(nil)
}

// UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil
func (o *TestCreate) UnsetApiScanConfiguration() {
	o.ApiScanConfiguration.Unset()
}

// GetFiles returns the Files field value
func (o *TestCreate) GetFiles() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *TestCreate) GetFilesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *TestCreate) SetFiles(v []int32) {
	o.Files = v
}


func (o TestCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["engagement"] = o.Engagement
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.ScanType.IsSet() {
		toSerialize["scan_type"] = o.ScanType.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["target_start"] = o.TargetStart
	toSerialize["target_end"] = o.TargetEnd
	toSerialize["estimated_time"] = o.EstimatedTime.Get()
	toSerialize["actual_time"] = o.ActualTime.Get()
	if o.PercentComplete.IsSet() {
		toSerialize["percent_complete"] = o.PercentComplete.Get()
	}
	toSerialize["updated"] = o.Updated.Get()
	toSerialize["created"] = o.Created.Get()
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.BuildId.IsSet() {
		toSerialize["build_id"] = o.BuildId.Get()
	}
	if o.CommitHash.IsSet() {
		toSerialize["commit_hash"] = o.CommitHash.Get()
	}
	if o.BranchTag.IsSet() {
		toSerialize["branch_tag"] = o.BranchTag.Get()
	}
	if o.Lead.IsSet() {
		toSerialize["lead"] = o.Lead.Get()
	}
	toSerialize["test_type"] = o.TestType
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.ApiScanConfiguration.IsSet() {
		toSerialize["api_scan_configuration"] = o.ApiScanConfiguration.Get()
	}
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"engagement",
		"target_start",
		"target_end",
		"estimated_time",
		"actual_time",
		"updated",
		"created",
		"test_type",
		"files",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varTestCreate := _TestCreate{}

	err = json.Unmarshal(data, &varTestCreate)

	if err != nil {
		return err
	}

	*o = TestCreate(varTestCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "engagement")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "scan_type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "target_start")
		delete(additionalProperties, "target_end")
		delete(additionalProperties, "estimated_time")
		delete(additionalProperties, "actual_time")
		delete(additionalProperties, "percent_complete")
		delete(additionalProperties, "updated")
		delete(additionalProperties, "created")
		delete(additionalProperties, "version")
		delete(additionalProperties, "build_id")
		delete(additionalProperties, "commit_hash")
		delete(additionalProperties, "branch_tag")
		delete(additionalProperties, "lead")
		delete(additionalProperties, "test_type")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "api_scan_configuration")
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestCreate struct {
	value *TestCreate
	isSet bool
}

func (v NullableTestCreate) Get() *TestCreate {
	return v.value
}

func (v *NullableTestCreate) Set(val *TestCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTestCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTestCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestCreate(val *TestCreate) *NullableTestCreate {
	return &NullableTestCreate{value: val, isSet: true}
}

func (v NullableTestCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


