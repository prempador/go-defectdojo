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

// checks if the PaginatedQuestionnaireGeneralSurveyList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedQuestionnaireGeneralSurveyList{}

// PaginatedQuestionnaireGeneralSurveyList struct for PaginatedQuestionnaireGeneralSurveyList
type PaginatedQuestionnaireGeneralSurveyList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []QuestionnaireGeneralSurvey `json:"results,omitempty"`
}

// NewPaginatedQuestionnaireGeneralSurveyList instantiates a new PaginatedQuestionnaireGeneralSurveyList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedQuestionnaireGeneralSurveyList() *PaginatedQuestionnaireGeneralSurveyList {
	this := PaginatedQuestionnaireGeneralSurveyList{}
	return &this
}

// NewPaginatedQuestionnaireGeneralSurveyListWithDefaults instantiates a new PaginatedQuestionnaireGeneralSurveyList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedQuestionnaireGeneralSurveyListWithDefaults() *PaginatedQuestionnaireGeneralSurveyList {
	this := PaginatedQuestionnaireGeneralSurveyList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedQuestionnaireGeneralSurveyList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedQuestionnaireGeneralSurveyList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireGeneralSurveyList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedQuestionnaireGeneralSurveyList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedQuestionnaireGeneralSurveyList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedQuestionnaireGeneralSurveyList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireGeneralSurveyList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedQuestionnaireGeneralSurveyList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedQuestionnaireGeneralSurveyList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedQuestionnaireGeneralSurveyList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedQuestionnaireGeneralSurveyList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedQuestionnaireGeneralSurveyList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireGeneralSurveyList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedQuestionnaireGeneralSurveyList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedQuestionnaireGeneralSurveyList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedQuestionnaireGeneralSurveyList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedQuestionnaireGeneralSurveyList) GetResults() []QuestionnaireGeneralSurvey {
	if o == nil || IsNil(o.Results) {
		var ret []QuestionnaireGeneralSurvey
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedQuestionnaireGeneralSurveyList) GetResultsOk() ([]QuestionnaireGeneralSurvey, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireGeneralSurveyList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []QuestionnaireGeneralSurvey and assigns it to the Results field.
func (o *PaginatedQuestionnaireGeneralSurveyList) SetResults(v []QuestionnaireGeneralSurvey) {
	o.Results = v
}

func (o PaginatedQuestionnaireGeneralSurveyList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedQuestionnaireGeneralSurveyList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedQuestionnaireGeneralSurveyList struct {
	value *PaginatedQuestionnaireGeneralSurveyList
	isSet bool
}

func (v NullablePaginatedQuestionnaireGeneralSurveyList) Get() *PaginatedQuestionnaireGeneralSurveyList {
	return v.value
}

func (v *NullablePaginatedQuestionnaireGeneralSurveyList) Set(val *PaginatedQuestionnaireGeneralSurveyList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedQuestionnaireGeneralSurveyList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedQuestionnaireGeneralSurveyList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedQuestionnaireGeneralSurveyList(val *PaginatedQuestionnaireGeneralSurveyList) *NullablePaginatedQuestionnaireGeneralSurveyList {
	return &NullablePaginatedQuestionnaireGeneralSurveyList{value: val, isSet: true}
}

func (v NullablePaginatedQuestionnaireGeneralSurveyList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedQuestionnaireGeneralSurveyList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


