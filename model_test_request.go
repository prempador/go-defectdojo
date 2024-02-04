/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRequest{}

// TestRequest struct for TestRequest
type TestRequest struct {
	Tags []string `json:"tags,omitempty"`
	ScanType NullableString `json:"scan_type,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	TargetStart time.Time `json:"target_start"`
	TargetEnd time.Time `json:"target_end"`
	PercentComplete NullableInt32 `json:"percent_complete,omitempty"`
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
}

type _TestRequest TestRequest

// NewTestRequest instantiates a new TestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRequest(targetStart time.Time, targetEnd time.Time, testType int32) *TestRequest {
	this := TestRequest{}
	this.TargetStart = targetStart
	this.TargetEnd = targetEnd
	this.TestType = testType
	return &this
}

// NewTestRequestWithDefaults instantiates a new TestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRequestWithDefaults() *TestRequest {
	this := TestRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TestRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TestRequest) SetTags(v []string) {
	o.Tags = v
}

// GetScanType returns the ScanType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetScanType() string {
	if o == nil || IsNil(o.ScanType.Get()) {
		var ret string
		return ret
	}
	return *o.ScanType.Get()
}

// GetScanTypeOk returns a tuple with the ScanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScanType.Get(), o.ScanType.IsSet()
}

// HasScanType returns a boolean if a field has been set.
func (o *TestRequest) HasScanType() bool {
	if o != nil && o.ScanType.IsSet() {
		return true
	}

	return false
}

// SetScanType gets a reference to the given NullableString and assigns it to the ScanType field.
func (o *TestRequest) SetScanType(v string) {
	o.ScanType.Set(&v)
}
// SetScanTypeNil sets the value for ScanType to be an explicit nil
func (o *TestRequest) SetScanTypeNil() {
	o.ScanType.Set(nil)
}

// UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
func (o *TestRequest) UnsetScanType() {
	o.ScanType.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TestRequest) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TestRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TestRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TestRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TestRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TestRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TestRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TestRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetTargetStart returns the TargetStart field value
func (o *TestRequest) GetTargetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TargetStart
}

// GetTargetStartOk returns a tuple with the TargetStart field value
// and a boolean to check if the value has been set.
func (o *TestRequest) GetTargetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetStart, true
}

// SetTargetStart sets field value
func (o *TestRequest) SetTargetStart(v time.Time) {
	o.TargetStart = v
}

// GetTargetEnd returns the TargetEnd field value
func (o *TestRequest) GetTargetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TargetEnd
}

// GetTargetEndOk returns a tuple with the TargetEnd field value
// and a boolean to check if the value has been set.
func (o *TestRequest) GetTargetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEnd, true
}

// SetTargetEnd sets field value
func (o *TestRequest) SetTargetEnd(v time.Time) {
	o.TargetEnd = v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetPercentComplete() int32 {
	if o == nil || IsNil(o.PercentComplete.Get()) {
		var ret int32
		return ret
	}
	return *o.PercentComplete.Get()
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetPercentCompleteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentComplete.Get(), o.PercentComplete.IsSet()
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *TestRequest) HasPercentComplete() bool {
	if o != nil && o.PercentComplete.IsSet() {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given NullableInt32 and assigns it to the PercentComplete field.
func (o *TestRequest) SetPercentComplete(v int32) {
	o.PercentComplete.Set(&v)
}
// SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil
func (o *TestRequest) SetPercentCompleteNil() {
	o.PercentComplete.Set(nil)
}

// UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
func (o *TestRequest) UnsetPercentComplete() {
	o.PercentComplete.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *TestRequest) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *TestRequest) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *TestRequest) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *TestRequest) UnsetVersion() {
	o.Version.Unset()
}

// GetBuildId returns the BuildId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetBuildId() string {
	if o == nil || IsNil(o.BuildId.Get()) {
		var ret string
		return ret
	}
	return *o.BuildId.Get()
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetBuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildId.Get(), o.BuildId.IsSet()
}

// HasBuildId returns a boolean if a field has been set.
func (o *TestRequest) HasBuildId() bool {
	if o != nil && o.BuildId.IsSet() {
		return true
	}

	return false
}

// SetBuildId gets a reference to the given NullableString and assigns it to the BuildId field.
func (o *TestRequest) SetBuildId(v string) {
	o.BuildId.Set(&v)
}
// SetBuildIdNil sets the value for BuildId to be an explicit nil
func (o *TestRequest) SetBuildIdNil() {
	o.BuildId.Set(nil)
}

// UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
func (o *TestRequest) UnsetBuildId() {
	o.BuildId.Unset()
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash.Get()) {
		var ret string
		return ret
	}
	return *o.CommitHash.Get()
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitHash.Get(), o.CommitHash.IsSet()
}

// HasCommitHash returns a boolean if a field has been set.
func (o *TestRequest) HasCommitHash() bool {
	if o != nil && o.CommitHash.IsSet() {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given NullableString and assigns it to the CommitHash field.
func (o *TestRequest) SetCommitHash(v string) {
	o.CommitHash.Set(&v)
}
// SetCommitHashNil sets the value for CommitHash to be an explicit nil
func (o *TestRequest) SetCommitHashNil() {
	o.CommitHash.Set(nil)
}

// UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
func (o *TestRequest) UnsetCommitHash() {
	o.CommitHash.Unset()
}

// GetBranchTag returns the BranchTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetBranchTag() string {
	if o == nil || IsNil(o.BranchTag.Get()) {
		var ret string
		return ret
	}
	return *o.BranchTag.Get()
}

// GetBranchTagOk returns a tuple with the BranchTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetBranchTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BranchTag.Get(), o.BranchTag.IsSet()
}

// HasBranchTag returns a boolean if a field has been set.
func (o *TestRequest) HasBranchTag() bool {
	if o != nil && o.BranchTag.IsSet() {
		return true
	}

	return false
}

// SetBranchTag gets a reference to the given NullableString and assigns it to the BranchTag field.
func (o *TestRequest) SetBranchTag(v string) {
	o.BranchTag.Set(&v)
}
// SetBranchTagNil sets the value for BranchTag to be an explicit nil
func (o *TestRequest) SetBranchTagNil() {
	o.BranchTag.Set(nil)
}

// UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
func (o *TestRequest) UnsetBranchTag() {
	o.BranchTag.Unset()
}

// GetLead returns the Lead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetLead() int32 {
	if o == nil || IsNil(o.Lead.Get()) {
		var ret int32
		return ret
	}
	return *o.Lead.Get()
}

// GetLeadOk returns a tuple with the Lead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetLeadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lead.Get(), o.Lead.IsSet()
}

// HasLead returns a boolean if a field has been set.
func (o *TestRequest) HasLead() bool {
	if o != nil && o.Lead.IsSet() {
		return true
	}

	return false
}

// SetLead gets a reference to the given NullableInt32 and assigns it to the Lead field.
func (o *TestRequest) SetLead(v int32) {
	o.Lead.Set(&v)
}
// SetLeadNil sets the value for Lead to be an explicit nil
func (o *TestRequest) SetLeadNil() {
	o.Lead.Set(nil)
}

// UnsetLead ensures that no value is present for Lead, not even an explicit nil
func (o *TestRequest) UnsetLead() {
	o.Lead.Unset()
}

// GetTestType returns the TestType field value
func (o *TestRequest) GetTestType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value
// and a boolean to check if the value has been set.
func (o *TestRequest) GetTestTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestType, true
}

// SetTestType sets field value
func (o *TestRequest) SetTestType(v int32) {
	o.TestType = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetEnvironment() int32 {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret int32
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetEnvironmentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *TestRequest) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableInt32 and assigns it to the Environment field.
func (o *TestRequest) SetEnvironment(v int32) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *TestRequest) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *TestRequest) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetApiScanConfiguration returns the ApiScanConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestRequest) GetApiScanConfiguration() int32 {
	if o == nil || IsNil(o.ApiScanConfiguration.Get()) {
		var ret int32
		return ret
	}
	return *o.ApiScanConfiguration.Get()
}

// GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestRequest) GetApiScanConfigurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiScanConfiguration.Get(), o.ApiScanConfiguration.IsSet()
}

// HasApiScanConfiguration returns a boolean if a field has been set.
func (o *TestRequest) HasApiScanConfiguration() bool {
	if o != nil && o.ApiScanConfiguration.IsSet() {
		return true
	}

	return false
}

// SetApiScanConfiguration gets a reference to the given NullableInt32 and assigns it to the ApiScanConfiguration field.
func (o *TestRequest) SetApiScanConfiguration(v int32) {
	o.ApiScanConfiguration.Set(&v)
}
// SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil
func (o *TestRequest) SetApiScanConfigurationNil() {
	o.ApiScanConfiguration.Set(nil)
}

// UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil
func (o *TestRequest) UnsetApiScanConfiguration() {
	o.ApiScanConfiguration.Unset()
}

func (o TestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if o.PercentComplete.IsSet() {
		toSerialize["percent_complete"] = o.PercentComplete.Get()
	}
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
	return toSerialize, nil
}

func (o *TestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"target_start",
		"target_end",
		"test_type",
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

	varTestRequest := _TestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestRequest)

	if err != nil {
		return err
	}

	*o = TestRequest(varTestRequest)

	return err
}

type NullableTestRequest struct {
	value *TestRequest
	isSet bool
}

func (v NullableTestRequest) Get() *TestRequest {
	return v.value
}

func (v *NullableTestRequest) Set(val *TestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRequest(val *TestRequest) *NullableTestRequest {
	return &NullableTestRequest{value: val, isSet: true}
}

func (v NullableTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


