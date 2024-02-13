/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FindingGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindingGroupRequest{}

// FindingGroupRequest struct for FindingGroupRequest
type FindingGroupRequest struct {
	Name string `json:"name"`
	Test int32 `json:"test"`
}

type _FindingGroupRequest FindingGroupRequest

// NewFindingGroupRequest instantiates a new FindingGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindingGroupRequest(name string, test int32) *FindingGroupRequest {
	this := FindingGroupRequest{}
	this.Name = name
	this.Test = test
	return &this
}

// NewFindingGroupRequestWithDefaults instantiates a new FindingGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindingGroupRequestWithDefaults() *FindingGroupRequest {
	this := FindingGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *FindingGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FindingGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FindingGroupRequest) SetName(v string) {
	o.Name = v
}

// GetTest returns the Test field value
func (o *FindingGroupRequest) GetTest() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *FindingGroupRequest) GetTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *FindingGroupRequest) SetTest(v int32) {
	o.Test = v
}

func (o FindingGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindingGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["test"] = o.Test
	return toSerialize, nil
}

func (o *FindingGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"test",
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

	varFindingGroupRequest := _FindingGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFindingGroupRequest)

	if err != nil {
		return err
	}

	*o = FindingGroupRequest(varFindingGroupRequest)

	return err
}

type NullableFindingGroupRequest struct {
	value *FindingGroupRequest
	isSet bool
}

func (v NullableFindingGroupRequest) Get() *FindingGroupRequest {
	return v.value
}

func (v *NullableFindingGroupRequest) Set(val *FindingGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFindingGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFindingGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindingGroupRequest(val *FindingGroupRequest) *NullableFindingGroupRequest {
	return &NullableFindingGroupRequest{value: val, isSet: true}
}

func (v NullableFindingGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindingGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


