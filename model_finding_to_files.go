/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FindingToFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindingToFiles{}

// FindingToFiles struct for FindingToFiles
type FindingToFiles struct {
	FindingId NullableInt32 `json:"finding_id"`
	Files []File `json:"files"`
}

type _FindingToFiles FindingToFiles

// NewFindingToFiles instantiates a new FindingToFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindingToFiles(findingId NullableInt32, files []File) *FindingToFiles {
	this := FindingToFiles{}
	this.FindingId = findingId
	this.Files = files
	return &this
}

// NewFindingToFilesWithDefaults instantiates a new FindingToFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindingToFilesWithDefaults() *FindingToFiles {
	this := FindingToFiles{}
	return &this
}

// GetFindingId returns the FindingId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FindingToFiles) GetFindingId() int32 {
	if o == nil || o.FindingId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FindingId.Get()
}

// GetFindingIdOk returns a tuple with the FindingId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FindingToFiles) GetFindingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FindingId.Get(), o.FindingId.IsSet()
}

// SetFindingId sets field value
func (o *FindingToFiles) SetFindingId(v int32) {
	o.FindingId.Set(&v)
}

// GetFiles returns the Files field value
func (o *FindingToFiles) GetFiles() []File {
	if o == nil {
		var ret []File
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *FindingToFiles) GetFilesOk() ([]File, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *FindingToFiles) SetFiles(v []File) {
	o.Files = v
}

func (o FindingToFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindingToFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["finding_id"] = o.FindingId.Get()
	toSerialize["files"] = o.Files
	return toSerialize, nil
}

func (o *FindingToFiles) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"finding_id",
		"files",
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

	varFindingToFiles := _FindingToFiles{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFindingToFiles)

	if err != nil {
		return err
	}

	*o = FindingToFiles(varFindingToFiles)

	return err
}

type NullableFindingToFiles struct {
	value *FindingToFiles
	isSet bool
}

func (v NullableFindingToFiles) Get() *FindingToFiles {
	return v.value
}

func (v *NullableFindingToFiles) Set(val *FindingToFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableFindingToFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableFindingToFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindingToFiles(val *FindingToFiles) *NullableFindingToFiles {
	return &NullableFindingToFiles{value: val, isSet: true}
}

func (v NullableFindingToFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindingToFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


