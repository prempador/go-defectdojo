/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StubFindingCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StubFindingCreateRequest{}

// StubFindingCreateRequest struct for StubFindingCreateRequest
type StubFindingCreateRequest struct {
	Test int32 `json:"test"`
	Title string `json:"title"`
	Date *string `json:"date,omitempty"`
	Severity NullableString `json:"severity,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

type _StubFindingCreateRequest StubFindingCreateRequest

// NewStubFindingCreateRequest instantiates a new StubFindingCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStubFindingCreateRequest(test int32, title string) *StubFindingCreateRequest {
	this := StubFindingCreateRequest{}
	this.Test = test
	this.Title = title
	return &this
}

// NewStubFindingCreateRequestWithDefaults instantiates a new StubFindingCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStubFindingCreateRequestWithDefaults() *StubFindingCreateRequest {
	this := StubFindingCreateRequest{}
	return &this
}

// GetTest returns the Test field value
func (o *StubFindingCreateRequest) GetTest() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *StubFindingCreateRequest) GetTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *StubFindingCreateRequest) SetTest(v int32) {
	o.Test = v
}

// GetTitle returns the Title field value
func (o *StubFindingCreateRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *StubFindingCreateRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *StubFindingCreateRequest) SetTitle(v string) {
	o.Title = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *StubFindingCreateRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StubFindingCreateRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *StubFindingCreateRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *StubFindingCreateRequest) SetDate(v string) {
	o.Date = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StubFindingCreateRequest) GetSeverity() string {
	if o == nil || IsNil(o.Severity.Get()) {
		var ret string
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StubFindingCreateRequest) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *StubFindingCreateRequest) HasSeverity() bool {
	if o != nil && o.Severity.IsSet() {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given NullableString and assigns it to the Severity field.
func (o *StubFindingCreateRequest) SetSeverity(v string) {
	o.Severity.Set(&v)
}
// SetSeverityNil sets the value for Severity to be an explicit nil
func (o *StubFindingCreateRequest) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
func (o *StubFindingCreateRequest) UnsetSeverity() {
	o.Severity.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StubFindingCreateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StubFindingCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StubFindingCreateRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *StubFindingCreateRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StubFindingCreateRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StubFindingCreateRequest) UnsetDescription() {
	o.Description.Unset()
}

func (o StubFindingCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StubFindingCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["test"] = o.Test
	toSerialize["title"] = o.Title
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

func (o *StubFindingCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"test",
		"title",
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

	varStubFindingCreateRequest := _StubFindingCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStubFindingCreateRequest)

	if err != nil {
		return err
	}

	*o = StubFindingCreateRequest(varStubFindingCreateRequest)

	return err
}

type NullableStubFindingCreateRequest struct {
	value *StubFindingCreateRequest
	isSet bool
}

func (v NullableStubFindingCreateRequest) Get() *StubFindingCreateRequest {
	return v.value
}

func (v *NullableStubFindingCreateRequest) Set(val *StubFindingCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStubFindingCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStubFindingCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStubFindingCreateRequest(val *StubFindingCreateRequest) *NullableStubFindingCreateRequest {
	return &NullableStubFindingCreateRequest{value: val, isSet: true}
}

func (v NullableStubFindingCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStubFindingCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


