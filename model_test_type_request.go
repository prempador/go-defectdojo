/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TestTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestTypeRequest{}

// TestTypeRequest struct for TestTypeRequest
type TestTypeRequest struct {
	Tags []string `json:"tags,omitempty"`
	Name string `json:"name"`
	StaticTool *bool `json:"static_tool,omitempty"`
	DynamicTool *bool `json:"dynamic_tool,omitempty"`
	Active *bool `json:"active,omitempty"`
}

type _TestTypeRequest TestTypeRequest

// NewTestTypeRequest instantiates a new TestTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestTypeRequest(name string) *TestTypeRequest {
	this := TestTypeRequest{}
	this.Name = name
	return &this
}

// NewTestTypeRequestWithDefaults instantiates a new TestTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestTypeRequestWithDefaults() *TestTypeRequest {
	this := TestTypeRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TestTypeRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestTypeRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TestTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TestTypeRequest) SetTags(v []string) {
	o.Tags = v
}

// GetName returns the Name field value
func (o *TestTypeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestTypeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestTypeRequest) SetName(v string) {
	o.Name = v
}

// GetStaticTool returns the StaticTool field value if set, zero value otherwise.
func (o *TestTypeRequest) GetStaticTool() bool {
	if o == nil || IsNil(o.StaticTool) {
		var ret bool
		return ret
	}
	return *o.StaticTool
}

// GetStaticToolOk returns a tuple with the StaticTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestTypeRequest) GetStaticToolOk() (*bool, bool) {
	if o == nil || IsNil(o.StaticTool) {
		return nil, false
	}
	return o.StaticTool, true
}

// HasStaticTool returns a boolean if a field has been set.
func (o *TestTypeRequest) HasStaticTool() bool {
	if o != nil && !IsNil(o.StaticTool) {
		return true
	}

	return false
}

// SetStaticTool gets a reference to the given bool and assigns it to the StaticTool field.
func (o *TestTypeRequest) SetStaticTool(v bool) {
	o.StaticTool = &v
}

// GetDynamicTool returns the DynamicTool field value if set, zero value otherwise.
func (o *TestTypeRequest) GetDynamicTool() bool {
	if o == nil || IsNil(o.DynamicTool) {
		var ret bool
		return ret
	}
	return *o.DynamicTool
}

// GetDynamicToolOk returns a tuple with the DynamicTool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestTypeRequest) GetDynamicToolOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicTool) {
		return nil, false
	}
	return o.DynamicTool, true
}

// HasDynamicTool returns a boolean if a field has been set.
func (o *TestTypeRequest) HasDynamicTool() bool {
	if o != nil && !IsNil(o.DynamicTool) {
		return true
	}

	return false
}

// SetDynamicTool gets a reference to the given bool and assigns it to the DynamicTool field.
func (o *TestTypeRequest) SetDynamicTool(v bool) {
	o.DynamicTool = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TestTypeRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestTypeRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TestTypeRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TestTypeRequest) SetActive(v bool) {
	o.Active = &v
}

func (o TestTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.StaticTool) {
		toSerialize["static_tool"] = o.StaticTool
	}
	if !IsNil(o.DynamicTool) {
		toSerialize["dynamic_tool"] = o.DynamicTool
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

func (o *TestTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varTestTypeRequest := _TestTypeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestTypeRequest)

	if err != nil {
		return err
	}

	*o = TestTypeRequest(varTestTypeRequest)

	return err
}

type NullableTestTypeRequest struct {
	value *TestTypeRequest
	isSet bool
}

func (v NullableTestTypeRequest) Get() *TestTypeRequest {
	return v.value
}

func (v *NullableTestTypeRequest) Set(val *TestTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestTypeRequest(val *TestTypeRequest) *NullableTestTypeRequest {
	return &NullableTestTypeRequest{value: val, isSet: true}
}

func (v NullableTestTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


