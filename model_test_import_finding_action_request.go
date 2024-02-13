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

// checks if the TestImportFindingActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestImportFindingActionRequest{}

// TestImportFindingActionRequest struct for TestImportFindingActionRequest
type TestImportFindingActionRequest struct {
	// * `N` - created * `C` - closed * `R` - reactivated * `U` - left untouched
	Action NullableString `json:"action,omitempty"`
}

// NewTestImportFindingActionRequest instantiates a new TestImportFindingActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestImportFindingActionRequest() *TestImportFindingActionRequest {
	this := TestImportFindingActionRequest{}
	return &this
}

// NewTestImportFindingActionRequestWithDefaults instantiates a new TestImportFindingActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestImportFindingActionRequestWithDefaults() *TestImportFindingActionRequest {
	this := TestImportFindingActionRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestImportFindingActionRequest) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestImportFindingActionRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *TestImportFindingActionRequest) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *TestImportFindingActionRequest) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *TestImportFindingActionRequest) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *TestImportFindingActionRequest) UnsetAction() {
	o.Action.Unset()
}

func (o TestImportFindingActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestImportFindingActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	return toSerialize, nil
}

type NullableTestImportFindingActionRequest struct {
	value *TestImportFindingActionRequest
	isSet bool
}

func (v NullableTestImportFindingActionRequest) Get() *TestImportFindingActionRequest {
	return v.value
}

func (v *NullableTestImportFindingActionRequest) Set(val *TestImportFindingActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestImportFindingActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestImportFindingActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestImportFindingActionRequest(val *TestImportFindingActionRequest) *NullableTestImportFindingActionRequest {
	return &NullableTestImportFindingActionRequest{value: val, isSet: true}
}

func (v NullableTestImportFindingActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestImportFindingActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


