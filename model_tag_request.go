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

// checks if the TagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagRequest{}

// TagRequest struct for TagRequest
type TagRequest struct {
	Tags []string `json:"tags"`
}

type _TagRequest TagRequest

// NewTagRequest instantiates a new TagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagRequest(tags []string) *TagRequest {
	this := TagRequest{}
	this.Tags = tags
	return &this
}

// NewTagRequestWithDefaults instantiates a new TagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagRequestWithDefaults() *TagRequest {
	this := TagRequest{}
	return &this
}

// GetTags returns the Tags field value
func (o *TagRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagRequest) SetTags(v []string) {
	o.Tags = v
}

func (o TagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *TagRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
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

	varTagRequest := _TagRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagRequest)

	if err != nil {
		return err
	}

	*o = TagRequest(varTagRequest)

	return err
}

type NullableTagRequest struct {
	value *TagRequest
	isSet bool
}

func (v NullableTagRequest) Get() *TagRequest {
	return v.value
}

func (v *NullableTagRequest) Set(val *TagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRequest(val *TagRequest) *NullableTagRequest {
	return &NullableTagRequest{value: val, isSet: true}
}

func (v NullableTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


