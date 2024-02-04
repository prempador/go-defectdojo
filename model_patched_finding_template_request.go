/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchedFindingTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedFindingTemplateRequest{}

// PatchedFindingTemplateRequest struct for PatchedFindingTemplateRequest
type PatchedFindingTemplateRequest struct {
	Tags []string `json:"tags,omitempty"`
	VulnerabilityIds []VulnerabilityIdTemplateRequest `json:"vulnerability_ids,omitempty"`
	Title *string `json:"title,omitempty"`
	Cwe NullableInt32 `json:"cwe,omitempty"`
	Cvssv3 NullableString `json:"cvssv3,omitempty"`
	Severity NullableString `json:"severity,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Mitigation NullableString `json:"mitigation,omitempty"`
	Impact NullableString `json:"impact,omitempty"`
	References NullableString `json:"references,omitempty"`
	// Enables this template for matching remediation advice. Match will be applied to all active, verified findings by CWE.
	TemplateMatch *bool `json:"template_match,omitempty"`
	// Matches by title text (contains search) and CWE.
	TemplateMatchTitle *bool `json:"template_match_title,omitempty"`
}

// NewPatchedFindingTemplateRequest instantiates a new PatchedFindingTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFindingTemplateRequest() *PatchedFindingTemplateRequest {
	this := PatchedFindingTemplateRequest{}
	return &this
}

// NewPatchedFindingTemplateRequestWithDefaults instantiates a new PatchedFindingTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFindingTemplateRequestWithDefaults() *PatchedFindingTemplateRequest {
	this := PatchedFindingTemplateRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedFindingTemplateRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFindingTemplateRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PatchedFindingTemplateRequest) SetTags(v []string) {
	o.Tags = v
}

// GetVulnerabilityIds returns the VulnerabilityIds field value if set, zero value otherwise.
func (o *PatchedFindingTemplateRequest) GetVulnerabilityIds() []VulnerabilityIdTemplateRequest {
	if o == nil || IsNil(o.VulnerabilityIds) {
		var ret []VulnerabilityIdTemplateRequest
		return ret
	}
	return o.VulnerabilityIds
}

// GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFindingTemplateRequest) GetVulnerabilityIdsOk() ([]VulnerabilityIdTemplateRequest, bool) {
	if o == nil || IsNil(o.VulnerabilityIds) {
		return nil, false
	}
	return o.VulnerabilityIds, true
}

// HasVulnerabilityIds returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasVulnerabilityIds() bool {
	if o != nil && !IsNil(o.VulnerabilityIds) {
		return true
	}

	return false
}

// SetVulnerabilityIds gets a reference to the given []VulnerabilityIdTemplateRequest and assigns it to the VulnerabilityIds field.
func (o *PatchedFindingTemplateRequest) SetVulnerabilityIds(v []VulnerabilityIdTemplateRequest) {
	o.VulnerabilityIds = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PatchedFindingTemplateRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFindingTemplateRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PatchedFindingTemplateRequest) SetTitle(v string) {
	o.Title = &v
}

// GetCwe returns the Cwe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetCwe() int32 {
	if o == nil || IsNil(o.Cwe.Get()) {
		var ret int32
		return ret
	}
	return *o.Cwe.Get()
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetCweOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cwe.Get(), o.Cwe.IsSet()
}

// HasCwe returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasCwe() bool {
	if o != nil && o.Cwe.IsSet() {
		return true
	}

	return false
}

// SetCwe gets a reference to the given NullableInt32 and assigns it to the Cwe field.
func (o *PatchedFindingTemplateRequest) SetCwe(v int32) {
	o.Cwe.Set(&v)
}
// SetCweNil sets the value for Cwe to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetCweNil() {
	o.Cwe.Set(nil)
}

// UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetCwe() {
	o.Cwe.Unset()
}

// GetCvssv3 returns the Cvssv3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetCvssv3() string {
	if o == nil || IsNil(o.Cvssv3.Get()) {
		var ret string
		return ret
	}
	return *o.Cvssv3.Get()
}

// GetCvssv3Ok returns a tuple with the Cvssv3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetCvssv3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cvssv3.Get(), o.Cvssv3.IsSet()
}

// HasCvssv3 returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasCvssv3() bool {
	if o != nil && o.Cvssv3.IsSet() {
		return true
	}

	return false
}

// SetCvssv3 gets a reference to the given NullableString and assigns it to the Cvssv3 field.
func (o *PatchedFindingTemplateRequest) SetCvssv3(v string) {
	o.Cvssv3.Set(&v)
}
// SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetCvssv3Nil() {
	o.Cvssv3.Set(nil)
}

// UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetCvssv3() {
	o.Cvssv3.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetSeverity() string {
	if o == nil || IsNil(o.Severity.Get()) {
		var ret string
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasSeverity() bool {
	if o != nil && o.Severity.IsSet() {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given NullableString and assigns it to the Severity field.
func (o *PatchedFindingTemplateRequest) SetSeverity(v string) {
	o.Severity.Set(&v)
}
// SetSeverityNil sets the value for Severity to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetSeverity() {
	o.Severity.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedFindingTemplateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetMitigation returns the Mitigation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetMitigation() string {
	if o == nil || IsNil(o.Mitigation.Get()) {
		var ret string
		return ret
	}
	return *o.Mitigation.Get()
}

// GetMitigationOk returns a tuple with the Mitigation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetMitigationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mitigation.Get(), o.Mitigation.IsSet()
}

// HasMitigation returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasMitigation() bool {
	if o != nil && o.Mitigation.IsSet() {
		return true
	}

	return false
}

// SetMitigation gets a reference to the given NullableString and assigns it to the Mitigation field.
func (o *PatchedFindingTemplateRequest) SetMitigation(v string) {
	o.Mitigation.Set(&v)
}
// SetMitigationNil sets the value for Mitigation to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetMitigationNil() {
	o.Mitigation.Set(nil)
}

// UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetMitigation() {
	o.Mitigation.Unset()
}

// GetImpact returns the Impact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetImpact() string {
	if o == nil || IsNil(o.Impact.Get()) {
		var ret string
		return ret
	}
	return *o.Impact.Get()
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Impact.Get(), o.Impact.IsSet()
}

// HasImpact returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasImpact() bool {
	if o != nil && o.Impact.IsSet() {
		return true
	}

	return false
}

// SetImpact gets a reference to the given NullableString and assigns it to the Impact field.
func (o *PatchedFindingTemplateRequest) SetImpact(v string) {
	o.Impact.Set(&v)
}
// SetImpactNil sets the value for Impact to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetImpactNil() {
	o.Impact.Set(nil)
}

// UnsetImpact ensures that no value is present for Impact, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetImpact() {
	o.Impact.Unset()
}

// GetReferences returns the References field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedFindingTemplateRequest) GetReferences() string {
	if o == nil || IsNil(o.References.Get()) {
		var ret string
		return ret
	}
	return *o.References.Get()
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedFindingTemplateRequest) GetReferencesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.References.Get(), o.References.IsSet()
}

// HasReferences returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasReferences() bool {
	if o != nil && o.References.IsSet() {
		return true
	}

	return false
}

// SetReferences gets a reference to the given NullableString and assigns it to the References field.
func (o *PatchedFindingTemplateRequest) SetReferences(v string) {
	o.References.Set(&v)
}
// SetReferencesNil sets the value for References to be an explicit nil
func (o *PatchedFindingTemplateRequest) SetReferencesNil() {
	o.References.Set(nil)
}

// UnsetReferences ensures that no value is present for References, not even an explicit nil
func (o *PatchedFindingTemplateRequest) UnsetReferences() {
	o.References.Unset()
}

// GetTemplateMatch returns the TemplateMatch field value if set, zero value otherwise.
func (o *PatchedFindingTemplateRequest) GetTemplateMatch() bool {
	if o == nil || IsNil(o.TemplateMatch) {
		var ret bool
		return ret
	}
	return *o.TemplateMatch
}

// GetTemplateMatchOk returns a tuple with the TemplateMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFindingTemplateRequest) GetTemplateMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.TemplateMatch) {
		return nil, false
	}
	return o.TemplateMatch, true
}

// HasTemplateMatch returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasTemplateMatch() bool {
	if o != nil && !IsNil(o.TemplateMatch) {
		return true
	}

	return false
}

// SetTemplateMatch gets a reference to the given bool and assigns it to the TemplateMatch field.
func (o *PatchedFindingTemplateRequest) SetTemplateMatch(v bool) {
	o.TemplateMatch = &v
}

// GetTemplateMatchTitle returns the TemplateMatchTitle field value if set, zero value otherwise.
func (o *PatchedFindingTemplateRequest) GetTemplateMatchTitle() bool {
	if o == nil || IsNil(o.TemplateMatchTitle) {
		var ret bool
		return ret
	}
	return *o.TemplateMatchTitle
}

// GetTemplateMatchTitleOk returns a tuple with the TemplateMatchTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFindingTemplateRequest) GetTemplateMatchTitleOk() (*bool, bool) {
	if o == nil || IsNil(o.TemplateMatchTitle) {
		return nil, false
	}
	return o.TemplateMatchTitle, true
}

// HasTemplateMatchTitle returns a boolean if a field has been set.
func (o *PatchedFindingTemplateRequest) HasTemplateMatchTitle() bool {
	if o != nil && !IsNil(o.TemplateMatchTitle) {
		return true
	}

	return false
}

// SetTemplateMatchTitle gets a reference to the given bool and assigns it to the TemplateMatchTitle field.
func (o *PatchedFindingTemplateRequest) SetTemplateMatchTitle(v bool) {
	o.TemplateMatchTitle = &v
}

func (o PatchedFindingTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedFindingTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.VulnerabilityIds) {
		toSerialize["vulnerability_ids"] = o.VulnerabilityIds
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Cwe.IsSet() {
		toSerialize["cwe"] = o.Cwe.Get()
	}
	if o.Cvssv3.IsSet() {
		toSerialize["cvssv3"] = o.Cvssv3.Get()
	}
	if o.Severity.IsSet() {
		toSerialize["severity"] = o.Severity.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Mitigation.IsSet() {
		toSerialize["mitigation"] = o.Mitigation.Get()
	}
	if o.Impact.IsSet() {
		toSerialize["impact"] = o.Impact.Get()
	}
	if o.References.IsSet() {
		toSerialize["references"] = o.References.Get()
	}
	if !IsNil(o.TemplateMatch) {
		toSerialize["template_match"] = o.TemplateMatch
	}
	if !IsNil(o.TemplateMatchTitle) {
		toSerialize["template_match_title"] = o.TemplateMatchTitle
	}
	return toSerialize, nil
}

type NullablePatchedFindingTemplateRequest struct {
	value *PatchedFindingTemplateRequest
	isSet bool
}

func (v NullablePatchedFindingTemplateRequest) Get() *PatchedFindingTemplateRequest {
	return v.value
}

func (v *NullablePatchedFindingTemplateRequest) Set(val *PatchedFindingTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFindingTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFindingTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFindingTemplateRequest(val *PatchedFindingTemplateRequest) *NullablePatchedFindingTemplateRequest {
	return &NullablePatchedFindingTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedFindingTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFindingTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

