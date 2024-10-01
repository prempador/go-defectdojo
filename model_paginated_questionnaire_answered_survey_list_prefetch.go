/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PaginatedQuestionnaireAnsweredSurveyListPrefetch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedQuestionnaireAnsweredSurveyListPrefetch{}

// PaginatedQuestionnaireAnsweredSurveyListPrefetch struct for PaginatedQuestionnaireAnsweredSurveyListPrefetch
type PaginatedQuestionnaireAnsweredSurveyListPrefetch struct {
	Assignee *map[string]UserStub `json:"assignee,omitempty"`
	Engagement *map[string]FindingEngagement `json:"engagement,omitempty"`
	Responder *map[string]UserStub `json:"responder,omitempty"`
	Survey *map[string]QuestionnaireEngagementSurvey `json:"survey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedQuestionnaireAnsweredSurveyListPrefetch PaginatedQuestionnaireAnsweredSurveyListPrefetch

// NewPaginatedQuestionnaireAnsweredSurveyListPrefetch instantiates a new PaginatedQuestionnaireAnsweredSurveyListPrefetch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedQuestionnaireAnsweredSurveyListPrefetch() *PaginatedQuestionnaireAnsweredSurveyListPrefetch {
	this := PaginatedQuestionnaireAnsweredSurveyListPrefetch{}
	return &this
}

// NewPaginatedQuestionnaireAnsweredSurveyListPrefetchWithDefaults instantiates a new PaginatedQuestionnaireAnsweredSurveyListPrefetch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedQuestionnaireAnsweredSurveyListPrefetchWithDefaults() *PaginatedQuestionnaireAnsweredSurveyListPrefetch {
	this := PaginatedQuestionnaireAnsweredSurveyListPrefetch{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetAssignee() map[string]UserStub {
	if o == nil || IsNil(o.Assignee) {
		var ret map[string]UserStub
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetAssigneeOk() (*map[string]UserStub, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given map[string]UserStub and assigns it to the Assignee field.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetAssignee(v map[string]UserStub) {
	o.Assignee = &v
}

// GetEngagement returns the Engagement field value if set, zero value otherwise.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetEngagement() map[string]FindingEngagement {
	if o == nil || IsNil(o.Engagement) {
		var ret map[string]FindingEngagement
		return ret
	}
	return *o.Engagement
}

// GetEngagementOk returns a tuple with the Engagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetEngagementOk() (*map[string]FindingEngagement, bool) {
	if o == nil || IsNil(o.Engagement) {
		return nil, false
	}
	return o.Engagement, true
}

// HasEngagement returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasEngagement() bool {
	if o != nil && !IsNil(o.Engagement) {
		return true
	}

	return false
}

// SetEngagement gets a reference to the given map[string]FindingEngagement and assigns it to the Engagement field.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetEngagement(v map[string]FindingEngagement) {
	o.Engagement = &v
}

// GetResponder returns the Responder field value if set, zero value otherwise.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetResponder() map[string]UserStub {
	if o == nil || IsNil(o.Responder) {
		var ret map[string]UserStub
		return ret
	}
	return *o.Responder
}

// GetResponderOk returns a tuple with the Responder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetResponderOk() (*map[string]UserStub, bool) {
	if o == nil || IsNil(o.Responder) {
		return nil, false
	}
	return o.Responder, true
}

// HasResponder returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasResponder() bool {
	if o != nil && !IsNil(o.Responder) {
		return true
	}

	return false
}

// SetResponder gets a reference to the given map[string]UserStub and assigns it to the Responder field.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetResponder(v map[string]UserStub) {
	o.Responder = &v
}

// GetSurvey returns the Survey field value if set, zero value otherwise.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetSurvey() map[string]QuestionnaireEngagementSurvey {
	if o == nil || IsNil(o.Survey) {
		var ret map[string]QuestionnaireEngagementSurvey
		return ret
	}
	return *o.Survey
}

// GetSurveyOk returns a tuple with the Survey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetSurveyOk() (*map[string]QuestionnaireEngagementSurvey, bool) {
	if o == nil || IsNil(o.Survey) {
		return nil, false
	}
	return o.Survey, true
}

// HasSurvey returns a boolean if a field has been set.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasSurvey() bool {
	if o != nil && !IsNil(o.Survey) {
		return true
	}

	return false
}

// SetSurvey gets a reference to the given map[string]QuestionnaireEngagementSurvey and assigns it to the Survey field.
func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetSurvey(v map[string]QuestionnaireEngagementSurvey) {
	o.Survey = &v
}

func (o PaginatedQuestionnaireAnsweredSurveyListPrefetch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedQuestionnaireAnsweredSurveyListPrefetch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Engagement) {
		toSerialize["engagement"] = o.Engagement
	}
	if !IsNil(o.Responder) {
		toSerialize["responder"] = o.Responder
	}
	if !IsNil(o.Survey) {
		toSerialize["survey"] = o.Survey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) UnmarshalJSON(data []byte) (err error) {
	varPaginatedQuestionnaireAnsweredSurveyListPrefetch := _PaginatedQuestionnaireAnsweredSurveyListPrefetch{}

	err = json.Unmarshal(data, &varPaginatedQuestionnaireAnsweredSurveyListPrefetch)

	if err != nil {
		return err
	}

	*o = PaginatedQuestionnaireAnsweredSurveyListPrefetch(varPaginatedQuestionnaireAnsweredSurveyListPrefetch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignee")
		delete(additionalProperties, "engagement")
		delete(additionalProperties, "responder")
		delete(additionalProperties, "survey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch struct {
	value *PaginatedQuestionnaireAnsweredSurveyListPrefetch
	isSet bool
}

func (v NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch) Get() *PaginatedQuestionnaireAnsweredSurveyListPrefetch {
	return v.value
}

func (v *NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch) Set(val *PaginatedQuestionnaireAnsweredSurveyListPrefetch) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedQuestionnaireAnsweredSurveyListPrefetch(val *PaginatedQuestionnaireAnsweredSurveyListPrefetch) *NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch {
	return &NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch{value: val, isSet: true}
}

func (v NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedQuestionnaireAnsweredSurveyListPrefetch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


