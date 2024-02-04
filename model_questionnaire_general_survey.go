/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the QuestionnaireGeneralSurvey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireGeneralSurvey{}

// QuestionnaireGeneralSurvey struct for QuestionnaireGeneralSurvey
type QuestionnaireGeneralSurvey struct {
	Id int32 `json:"id"`
	Survey QuestionnaireEngagementSurvey `json:"survey"`
	NumResponses *int32 `json:"num_responses,omitempty"`
	Generated NullableTime `json:"generated"`
	Expiration time.Time `json:"expiration"`
}

type _QuestionnaireGeneralSurvey QuestionnaireGeneralSurvey

// NewQuestionnaireGeneralSurvey instantiates a new QuestionnaireGeneralSurvey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireGeneralSurvey(id int32, survey QuestionnaireEngagementSurvey, generated NullableTime, expiration time.Time) *QuestionnaireGeneralSurvey {
	this := QuestionnaireGeneralSurvey{}
	this.Id = id
	this.Survey = survey
	this.Generated = generated
	this.Expiration = expiration
	return &this
}

// NewQuestionnaireGeneralSurveyWithDefaults instantiates a new QuestionnaireGeneralSurvey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireGeneralSurveyWithDefaults() *QuestionnaireGeneralSurvey {
	this := QuestionnaireGeneralSurvey{}
	return &this
}

// GetId returns the Id field value
func (o *QuestionnaireGeneralSurvey) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireGeneralSurvey) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuestionnaireGeneralSurvey) SetId(v int32) {
	o.Id = v
}

// GetSurvey returns the Survey field value
func (o *QuestionnaireGeneralSurvey) GetSurvey() QuestionnaireEngagementSurvey {
	if o == nil {
		var ret QuestionnaireEngagementSurvey
		return ret
	}

	return o.Survey
}

// GetSurveyOk returns a tuple with the Survey field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireGeneralSurvey) GetSurveyOk() (*QuestionnaireEngagementSurvey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Survey, true
}

// SetSurvey sets field value
func (o *QuestionnaireGeneralSurvey) SetSurvey(v QuestionnaireEngagementSurvey) {
	o.Survey = v
}

// GetNumResponses returns the NumResponses field value if set, zero value otherwise.
func (o *QuestionnaireGeneralSurvey) GetNumResponses() int32 {
	if o == nil || IsNil(o.NumResponses) {
		var ret int32
		return ret
	}
	return *o.NumResponses
}

// GetNumResponsesOk returns a tuple with the NumResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireGeneralSurvey) GetNumResponsesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumResponses) {
		return nil, false
	}
	return o.NumResponses, true
}

// HasNumResponses returns a boolean if a field has been set.
func (o *QuestionnaireGeneralSurvey) HasNumResponses() bool {
	if o != nil && !IsNil(o.NumResponses) {
		return true
	}

	return false
}

// SetNumResponses gets a reference to the given int32 and assigns it to the NumResponses field.
func (o *QuestionnaireGeneralSurvey) SetNumResponses(v int32) {
	o.NumResponses = &v
}

// GetGenerated returns the Generated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *QuestionnaireGeneralSurvey) GetGenerated() time.Time {
	if o == nil || o.Generated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Generated.Get()
}

// GetGeneratedOk returns a tuple with the Generated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuestionnaireGeneralSurvey) GetGeneratedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Generated.Get(), o.Generated.IsSet()
}

// SetGenerated sets field value
func (o *QuestionnaireGeneralSurvey) SetGenerated(v time.Time) {
	o.Generated.Set(&v)
}

// GetExpiration returns the Expiration field value
func (o *QuestionnaireGeneralSurvey) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireGeneralSurvey) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *QuestionnaireGeneralSurvey) SetExpiration(v time.Time) {
	o.Expiration = v
}

func (o QuestionnaireGeneralSurvey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireGeneralSurvey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["survey"] = o.Survey
	if !IsNil(o.NumResponses) {
		toSerialize["num_responses"] = o.NumResponses
	}
	toSerialize["generated"] = o.Generated.Get()
	toSerialize["expiration"] = o.Expiration
	return toSerialize, nil
}

func (o *QuestionnaireGeneralSurvey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"survey",
		"generated",
		"expiration",
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

	varQuestionnaireGeneralSurvey := _QuestionnaireGeneralSurvey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuestionnaireGeneralSurvey)

	if err != nil {
		return err
	}

	*o = QuestionnaireGeneralSurvey(varQuestionnaireGeneralSurvey)

	return err
}

type NullableQuestionnaireGeneralSurvey struct {
	value *QuestionnaireGeneralSurvey
	isSet bool
}

func (v NullableQuestionnaireGeneralSurvey) Get() *QuestionnaireGeneralSurvey {
	return v.value
}

func (v *NullableQuestionnaireGeneralSurvey) Set(val *QuestionnaireGeneralSurvey) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireGeneralSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireGeneralSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireGeneralSurvey(val *QuestionnaireGeneralSurvey) *NullableQuestionnaireGeneralSurvey {
	return &NullableQuestionnaireGeneralSurvey{value: val, isSet: true}
}

func (v NullableQuestionnaireGeneralSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireGeneralSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


