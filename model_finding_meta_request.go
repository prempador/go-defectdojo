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

// checks if the FindingMetaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindingMetaRequest{}

// FindingMetaRequest struct for FindingMetaRequest
type FindingMetaRequest struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _FindingMetaRequest FindingMetaRequest

// NewFindingMetaRequest instantiates a new FindingMetaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindingMetaRequest(name string, value string) *FindingMetaRequest {
	this := FindingMetaRequest{}
	this.Name = name
	this.Value = value
	return &this
}

// NewFindingMetaRequestWithDefaults instantiates a new FindingMetaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindingMetaRequestWithDefaults() *FindingMetaRequest {
	this := FindingMetaRequest{}
	return &this
}

// GetName returns the Name field value
func (o *FindingMetaRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FindingMetaRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FindingMetaRequest) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *FindingMetaRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FindingMetaRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FindingMetaRequest) SetValue(v string) {
	o.Value = v
}

func (o FindingMetaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindingMetaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *FindingMetaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varFindingMetaRequest := _FindingMetaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFindingMetaRequest)

	if err != nil {
		return err
	}

	*o = FindingMetaRequest(varFindingMetaRequest)

	return err
}

type NullableFindingMetaRequest struct {
	value *FindingMetaRequest
	isSet bool
}

func (v NullableFindingMetaRequest) Get() *FindingMetaRequest {
	return v.value
}

func (v *NullableFindingMetaRequest) Set(val *FindingMetaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFindingMetaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFindingMetaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindingMetaRequest(val *FindingMetaRequest) *NullableFindingMetaRequest {
	return &NullableFindingMetaRequest{value: val, isSet: true}
}

func (v NullableFindingMetaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindingMetaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


