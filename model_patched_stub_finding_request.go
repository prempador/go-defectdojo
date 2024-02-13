/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedStubFindingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedStubFindingRequest{}

// PatchedStubFindingRequest struct for PatchedStubFindingRequest
type PatchedStubFindingRequest struct {
	Title *string `json:"title,omitempty"`
	Date *string `json:"date,omitempty"`
	Severity NullableString `json:"severity,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewPatchedStubFindingRequest instantiates a new PatchedStubFindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedStubFindingRequest() *PatchedStubFindingRequest {
	this := PatchedStubFindingRequest{}
	return &this
}

// NewPatchedStubFindingRequestWithDefaults instantiates a new PatchedStubFindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedStubFindingRequestWithDefaults() *PatchedStubFindingRequest {
	this := PatchedStubFindingRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PatchedStubFindingRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStubFindingRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PatchedStubFindingRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PatchedStubFindingRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PatchedStubFindingRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedStubFindingRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PatchedStubFindingRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *PatchedStubFindingRequest) SetDate(v string) {
	o.Date = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedStubFindingRequest) GetSeverity() string {
	if o == nil || IsNil(o.Severity.Get()) {
		var ret string
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedStubFindingRequest) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *PatchedStubFindingRequest) HasSeverity() bool {
	if o != nil && o.Severity.IsSet() {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given NullableString and assigns it to the Severity field.
func (o *PatchedStubFindingRequest) SetSeverity(v string) {
	o.Severity.Set(&v)
}
// SetSeverityNil sets the value for Severity to be an explicit nil
func (o *PatchedStubFindingRequest) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
func (o *PatchedStubFindingRequest) UnsetSeverity() {
	o.Severity.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedStubFindingRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedStubFindingRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedStubFindingRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedStubFindingRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedStubFindingRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedStubFindingRequest) UnsetDescription() {
	o.Description.Unset()
}

func (o PatchedStubFindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedStubFindingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.Severity.IsSet() {
		toSerialize["severity"] = o.Severity.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullablePatchedStubFindingRequest struct {
	value *PatchedStubFindingRequest
	isSet bool
}

func (v NullablePatchedStubFindingRequest) Get() *PatchedStubFindingRequest {
	return v.value
}

func (v *NullablePatchedStubFindingRequest) Set(val *PatchedStubFindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedStubFindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedStubFindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedStubFindingRequest(val *PatchedStubFindingRequest) *NullablePatchedStubFindingRequest {
	return &NullablePatchedStubFindingRequest{value: val, isSet: true}
}

func (v NullablePatchedStubFindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedStubFindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


