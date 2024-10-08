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

// checks if the Test type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Test{}

// Test struct for Test
type Test struct {
	Id int32 `json:"id"`
	Tags []string `json:"tags,omitempty"`
	TestTypeName string `json:"test_type_name"`
	FindingGroups []FindingGroup `json:"finding_groups"`
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
	Engagement int32 `json:"engagement"`
	Lead NullableInt32 `json:"lead,omitempty"`
	TestType int32 `json:"test_type"`
	Environment NullableInt32 `json:"environment,omitempty"`
	ApiScanConfiguration NullableInt32 `json:"api_scan_configuration,omitempty"`
	Notes []Note `json:"notes"`
	Files []File `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _Test Test

// NewTest instantiates a new Test object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTest(id int32, testTypeName string, findingGroups []FindingGroup, targetStart time.Time, targetEnd time.Time, estimatedTime NullableString, actualTime NullableString, updated NullableTime, created NullableTime, engagement int32, testType int32, notes []Note, files []File) *Test {
	this := Test{}
	this.Id = id
	this.TestTypeName = testTypeName
	this.FindingGroups = findingGroups
	this.TargetStart = targetStart
	this.TargetEnd = targetEnd
	this.EstimatedTime = estimatedTime
	this.ActualTime = actualTime
	this.Updated = updated
	this.Created = created
	this.Engagement = engagement
	this.TestType = testType
	this.Notes = notes
	this.Files = files
	return &this
}

// NewTestWithDefaults instantiates a new Test object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestWithDefaults() *Test {
	this := Test{}
	return &this
}

// GetId returns the Id field value
func (o *Test) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Test) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Test) SetId(v int32) {
	o.Id = v
}


// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Test) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Test) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Test) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Test) SetTags(v []string) {
	o.Tags = v
}

// GetTestTypeName returns the TestTypeName field value
func (o *Test) GetTestTypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestTypeName
}

// GetTestTypeNameOk returns a tuple with the TestTypeName field value
// and a boolean to check if the value has been set.
func (o *Test) GetTestTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestTypeName, true
}

// SetTestTypeName sets field value
func (o *Test) SetTestTypeName(v string) {
	o.TestTypeName = v
}


// GetFindingGroups returns the FindingGroups field value
func (o *Test) GetFindingGroups() []FindingGroup {
	if o == nil {
		var ret []FindingGroup
		return ret
	}

	return o.FindingGroups
}

// GetFindingGroupsOk returns a tuple with the FindingGroups field value
// and a boolean to check if the value has been set.
func (o *Test) GetFindingGroupsOk() ([]FindingGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.FindingGroups, true
}

// SetFindingGroups sets field value
func (o *Test) SetFindingGroups(v []FindingGroup) {
	o.FindingGroups = v
}


// GetScanType returns the ScanType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetScanType() string {
	if o == nil || IsNil(o.ScanType.Get()) {
		var ret string
		return ret
	}
	return *o.ScanType.Get()
}

// GetScanTypeOk returns a tuple with the ScanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScanType.Get(), o.ScanType.IsSet()
}

// HasScanType returns a boolean if a field has been set.
func (o *Test) HasScanType() bool {
	if o != nil && o.ScanType.IsSet() {
		return true
	}

	return false
}

// SetScanType gets a reference to the given NullableString and assigns it to the ScanType field.
func (o *Test) SetScanType(v string) {
	o.ScanType.Set(&v)
}
// SetScanTypeNil sets the value for ScanType to be an explicit nil
func (o *Test) SetScanTypeNil() {
	o.ScanType.Set(nil)
}

// UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
func (o *Test) UnsetScanType() {
	o.ScanType.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Test) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Test) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Test) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Test) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Test) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Test) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Test) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Test) UnsetDescription() {
	o.Description.Unset()
}

// GetTargetStart returns the TargetStart field value
func (o *Test) GetTargetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TargetStart
}

// GetTargetStartOk returns a tuple with the TargetStart field value
// and a boolean to check if the value has been set.
func (o *Test) GetTargetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetStart, true
}

// SetTargetStart sets field value
func (o *Test) SetTargetStart(v time.Time) {
	o.TargetStart = v
}


// GetTargetEnd returns the TargetEnd field value
func (o *Test) GetTargetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TargetEnd
}

// GetTargetEndOk returns a tuple with the TargetEnd field value
// and a boolean to check if the value has been set.
func (o *Test) GetTargetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnd, true
}

// SetTargetEnd sets field value
func (o *Test) SetTargetEnd(v time.Time) {
	o.TargetEnd = v
}


// GetEstimatedTime returns the EstimatedTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Test) GetEstimatedTime() string {
	if o == nil || o.EstimatedTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.EstimatedTime.Get()
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetEstimatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedTime.Get(), o.EstimatedTime.IsSet()
}

// SetEstimatedTime sets field value
func (o *Test) SetEstimatedTime(v string) {
	o.EstimatedTime.Set(&v)
}


// GetActualTime returns the ActualTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Test) GetActualTime() string {
	if o == nil || o.ActualTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActualTime.Get()
}

// GetActualTimeOk returns a tuple with the ActualTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetActualTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualTime.Get(), o.ActualTime.IsSet()
}

// SetActualTime sets field value
func (o *Test) SetActualTime(v string) {
	o.ActualTime.Set(&v)
}


// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetPercentComplete() int32 {
	if o == nil || IsNil(o.PercentComplete.Get()) {
		var ret int32
		return ret
	}
	return *o.PercentComplete.Get()
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetPercentCompleteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentComplete.Get(), o.PercentComplete.IsSet()
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *Test) HasPercentComplete() bool {
	if o != nil && o.PercentComplete.IsSet() {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given NullableInt32 and assigns it to the PercentComplete field.
func (o *Test) SetPercentComplete(v int32) {
	o.PercentComplete.Set(&v)
}
// SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil
func (o *Test) SetPercentCompleteNil() {
	o.PercentComplete.Set(nil)
}

// UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
func (o *Test) UnsetPercentComplete() {
	o.PercentComplete.Unset()
}

// GetUpdated returns the Updated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Test) GetUpdated() time.Time {
	if o == nil || o.Updated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// SetUpdated sets field value
func (o *Test) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}


// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Test) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Test) SetCreated(v time.Time) {
	o.Created.Set(&v)
}


// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *Test) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *Test) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *Test) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *Test) UnsetVersion() {
	o.Version.Unset()
}

// GetBuildId returns the BuildId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetBuildId() string {
	if o == nil || IsNil(o.BuildId.Get()) {
		var ret string
		return ret
	}
	return *o.BuildId.Get()
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetBuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildId.Get(), o.BuildId.IsSet()
}

// HasBuildId returns a boolean if a field has been set.
func (o *Test) HasBuildId() bool {
	if o != nil && o.BuildId.IsSet() {
		return true
	}

	return false
}

// SetBuildId gets a reference to the given NullableString and assigns it to the BuildId field.
func (o *Test) SetBuildId(v string) {
	o.BuildId.Set(&v)
}
// SetBuildIdNil sets the value for BuildId to be an explicit nil
func (o *Test) SetBuildIdNil() {
	o.BuildId.Set(nil)
}

// UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
func (o *Test) UnsetBuildId() {
	o.BuildId.Unset()
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash.Get()) {
		var ret string
		return ret
	}
	return *o.CommitHash.Get()
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitHash.Get(), o.CommitHash.IsSet()
}

// HasCommitHash returns a boolean if a field has been set.
func (o *Test) HasCommitHash() bool {
	if o != nil && o.CommitHash.IsSet() {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullableString and assigns it to the CommitHash field.
func (o *Test) SetCommitHash(v string) {
	o.CommitHash.Set(&v)
}
// SetCommitHashNil sets the value for CommitHash to be an explicit nil
func (o *Test) SetCommitHashNil() {
	o.CommitHash.Set(nil)
}

// UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
func (o *Test) UnsetCommitHash() {
	o.CommitHash.Unset()
}

// GetBranchTag returns the BranchTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetBranchTag() string {
	if o == nil || IsNil(o.BranchTag.Get()) {
		var ret string
		return ret
	}
	return *o.BranchTag.Get()
}

// GetBranchTagOk returns a tuple with the BranchTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetBranchTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchTag.Get(), o.BranchTag.IsSet()
}

// HasBranchTag returns a boolean if a field has been set.
func (o *Test) HasBranchTag() bool {
	if o != nil && o.BranchTag.IsSet() {
		return true
	}

	return false
}

// SetBranchTag gets a reference to the given NullableString and assigns it to the BranchTag field.
func (o *Test) SetBranchTag(v string) {
	o.BranchTag.Set(&v)
}
// SetBranchTagNil sets the value for BranchTag to be an explicit nil
func (o *Test) SetBranchTagNil() {
	o.BranchTag.Set(nil)
}

// UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
func (o *Test) UnsetBranchTag() {
	o.BranchTag.Unset()
}

// GetEngagement returns the Engagement field value
func (o *Test) GetEngagement() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Engagement
}

// GetEngagementOk returns a tuple with the Engagement field value
// and a boolean to check if the value has been set.
func (o *Test) GetEngagementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engagement, true
}

// SetEngagement sets field value
func (o *Test) SetEngagement(v int32) {
	o.Engagement = v
}


// GetLead returns the Lead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetLead() int32 {
	if o == nil || IsNil(o.Lead.Get()) {
		var ret int32
		return ret
	}
	return *o.Lead.Get()
}

// GetLeadOk returns a tuple with the Lead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetLeadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lead.Get(), o.Lead.IsSet()
}

// HasLead returns a boolean if a field has been set.
func (o *Test) HasLead() bool {
	if o != nil && o.Lead.IsSet() {
		return true
	}

	return false
}

// SetLead gets a reference to the given NullableInt32 and assigns it to the Lead field.
func (o *Test) SetLead(v int32) {
	o.Lead.Set(&v)
}
// SetLeadNil sets the value for Lead to be an explicit nil
func (o *Test) SetLeadNil() {
	o.Lead.Set(nil)
}

// UnsetLead ensures that no value is present for Lead, not even an explicit nil
func (o *Test) UnsetLead() {
	o.Lead.Unset()
}

// GetTestType returns the TestType field value
func (o *Test) GetTestType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value
// and a boolean to check if the value has been set.
func (o *Test) GetTestTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestType, true
}

// SetTestType sets field value
func (o *Test) SetTestType(v int32) {
	o.TestType = v
}


// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetEnvironment() int32 {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret int32
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetEnvironmentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Test) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableInt32 and assigns it to the Environment field.
func (o *Test) SetEnvironment(v int32) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *Test) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *Test) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetApiScanConfiguration returns the ApiScanConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Test) GetApiScanConfiguration() int32 {
	if o == nil || IsNil(o.ApiScanConfiguration.Get()) {
		var ret int32
		return ret
	}
	return *o.ApiScanConfiguration.Get()
}

// GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Test) GetApiScanConfigurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiScanConfiguration.Get(), o.ApiScanConfiguration.IsSet()
}

// HasApiScanConfiguration returns a boolean if a field has been set.
func (o *Test) HasApiScanConfiguration() bool {
	if o != nil && o.ApiScanConfiguration.IsSet() {
		return true
	}

	return false
}

// SetApiScanConfiguration gets a reference to the given NullableInt32 and assigns it to the ApiScanConfiguration field.
func (o *Test) SetApiScanConfiguration(v int32) {
	o.ApiScanConfiguration.Set(&v)
}
// SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil
func (o *Test) SetApiScanConfigurationNil() {
	o.ApiScanConfiguration.Set(nil)
}

// UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil
func (o *Test) UnsetApiScanConfiguration() {
	o.ApiScanConfiguration.Unset()
}

// GetNotes returns the Notes field value
func (o *Test) GetNotes() []Note {
	if o == nil {
		var ret []Note
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *Test) GetNotesOk() ([]Note, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *Test) SetNotes(v []Note) {
	o.Notes = v
}


// GetFiles returns the Files field value
func (o *Test) GetFiles() []File {
	if o == nil {
		var ret []File
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *Test) GetFilesOk() ([]File, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *Test) SetFiles(v []File) {
	o.Files = v
}


func (o Test) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Test) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["test_type_name"] = o.TestTypeName
	toSerialize["finding_groups"] = o.FindingGroups
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
	toSerialize["engagement"] = o.Engagement
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
	toSerialize["notes"] = o.Notes
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Test) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"test_type_name",
		"finding_groups",
		"target_start",
		"target_end",
		"estimated_time",
		"actual_time",
		"updated",
		"created",
		"engagement",
		"test_type",
		"notes",
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
	varTest := _Test{}

	err = json.Unmarshal(data, &varTest)

	if err != nil {
		return err
	}

	*o = Test(varTest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "test_type_name")
		delete(additionalProperties, "finding_groups")
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
		delete(additionalProperties, "engagement")
		delete(additionalProperties, "lead")
		delete(additionalProperties, "test_type")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "api_scan_configuration")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTest struct {
	value *Test
	isSet bool
}

func (v NullableTest) Get() *Test {
	return v.value
}

func (v *NullableTest) Set(val *Test) {
	v.value = val
	v.isSet = true
}

func (v NullableTest) IsSet() bool {
	return v.isSet
}

func (v *NullableTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTest(val *Test) *NullableTest {
	return &NullableTest{value: val, isSet: true}
}

func (v NullableTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


