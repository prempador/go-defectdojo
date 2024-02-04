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

// checks if the QuestionnaireQuestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuestionnaireQuestion{}

// QuestionnaireQuestion struct for QuestionnaireQuestion
type QuestionnaireQuestion struct {
	Id int32 `json:"id"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	// The render order
	Order *int32 `json:"order,omitempty"`
	// If selected, user doesn't have to answer this question
	Optional *bool `json:"optional,omitempty"`
	// The question text
	Text *string `json:"text,omitempty"`
}

type _QuestionnaireQuestion QuestionnaireQuestion

// NewQuestionnaireQuestion instantiates a new QuestionnaireQuestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireQuestion(id int32, created time.Time, modified time.Time) *QuestionnaireQuestion {
	this := QuestionnaireQuestion{}
	this.Id = id
	this.Created = created
	this.Modified = modified
	return &this
}

// NewQuestionnaireQuestionWithDefaults instantiates a new QuestionnaireQuestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireQuestionWithDefaults() *QuestionnaireQuestion {
	this := QuestionnaireQuestion{}
	return &this
}

// GetId returns the Id field value
func (o *QuestionnaireQuestion) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestion) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuestionnaireQuestion) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *QuestionnaireQuestion) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestion) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *QuestionnaireQuestion) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *QuestionnaireQuestion) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestion) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *QuestionnaireQuestion) SetModified(v time.Time) {
	o.Modified = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QuestionnaireQuestion) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestion) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QuestionnaireQuestion) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *QuestionnaireQuestion) SetOrder(v int32) {
	o.Order = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *QuestionnaireQuestion) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestion) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *QuestionnaireQuestion) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *QuestionnaireQuestion) SetOptional(v bool) {
	o.Optional = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *QuestionnaireQuestion) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestion) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *QuestionnaireQuestion) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *QuestionnaireQuestion) SetText(v string) {
	o.Text = &v
}

func (o QuestionnaireQuestion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuestionnaireQuestion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

func (o *QuestionnaireQuestion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created",
		"modified",
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

	varQuestionnaireQuestion := _QuestionnaireQuestion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuestionnaireQuestion)

	if err != nil {
		return err
	}

	*o = QuestionnaireQuestion(varQuestionnaireQuestion)

	return err
}

type NullableQuestionnaireQuestion struct {
	value *QuestionnaireQuestion
	isSet bool
}

func (v NullableQuestionnaireQuestion) Get() *QuestionnaireQuestion {
	return v.value
}

func (v *NullableQuestionnaireQuestion) Set(val *QuestionnaireQuestion) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireQuestion) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireQuestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireQuestion(val *QuestionnaireQuestion) *NullableQuestionnaireQuestion {
	return &NullableQuestionnaireQuestion{value: val, isSet: true}
}

func (v NullableQuestionnaireQuestion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireQuestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

