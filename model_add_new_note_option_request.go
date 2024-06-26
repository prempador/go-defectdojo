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

// checks if the AddNewNoteOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddNewNoteOptionRequest{}

// AddNewNoteOptionRequest struct for AddNewNoteOptionRequest
type AddNewNoteOptionRequest struct {
	Entry string `json:"entry"`
	Private *bool `json:"private,omitempty"`
	NoteType NullableInt32 `json:"note_type,omitempty"`
}

type _AddNewNoteOptionRequest AddNewNoteOptionRequest

// NewAddNewNoteOptionRequest instantiates a new AddNewNoteOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNewNoteOptionRequest(entry string) *AddNewNoteOptionRequest {
	this := AddNewNoteOptionRequest{}
	this.Entry = entry
	return &this
}

// NewAddNewNoteOptionRequestWithDefaults instantiates a new AddNewNoteOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNewNoteOptionRequestWithDefaults() *AddNewNoteOptionRequest {
	this := AddNewNoteOptionRequest{}
	return &this
}

// GetEntry returns the Entry field value
func (o *AddNewNoteOptionRequest) GetEntry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entry
}

// GetEntryOk returns a tuple with the Entry field value
// and a boolean to check if the value has been set.
func (o *AddNewNoteOptionRequest) GetEntryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entry, true
}

// SetEntry sets field value
func (o *AddNewNoteOptionRequest) SetEntry(v string) {
	o.Entry = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *AddNewNoteOptionRequest) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddNewNoteOptionRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *AddNewNoteOptionRequest) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *AddNewNoteOptionRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetNoteType returns the NoteType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddNewNoteOptionRequest) GetNoteType() int32 {
	if o == nil || IsNil(o.NoteType.Get()) {
		var ret int32
		return ret
	}
	return *o.NoteType.Get()
}

// GetNoteTypeOk returns a tuple with the NoteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddNewNoteOptionRequest) GetNoteTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoteType.Get(), o.NoteType.IsSet()
}

// HasNoteType returns a boolean if a field has been set.
func (o *AddNewNoteOptionRequest) HasNoteType() bool {
	if o != nil && o.NoteType.IsSet() {
		return true
	}

	return false
}

// SetNoteType gets a reference to the given NullableInt32 and assigns it to the NoteType field.
func (o *AddNewNoteOptionRequest) SetNoteType(v int32) {
	o.NoteType.Set(&v)
}
// SetNoteTypeNil sets the value for NoteType to be an explicit nil
func (o *AddNewNoteOptionRequest) SetNoteTypeNil() {
	o.NoteType.Set(nil)
}

// UnsetNoteType ensures that no value is present for NoteType, not even an explicit nil
func (o *AddNewNoteOptionRequest) UnsetNoteType() {
	o.NoteType.Unset()
}

func (o AddNewNoteOptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddNewNoteOptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry"] = o.Entry
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.NoteType.IsSet() {
		toSerialize["note_type"] = o.NoteType.Get()
	}
	return toSerialize, nil
}

func (o *AddNewNoteOptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entry",
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

	varAddNewNoteOptionRequest := _AddNewNoteOptionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddNewNoteOptionRequest)

	if err != nil {
		return err
	}

	*o = AddNewNoteOptionRequest(varAddNewNoteOptionRequest)

	return err
}

type NullableAddNewNoteOptionRequest struct {
	value *AddNewNoteOptionRequest
	isSet bool
}

func (v NullableAddNewNoteOptionRequest) Get() *AddNewNoteOptionRequest {
	return v.value
}

func (v *NullableAddNewNoteOptionRequest) Set(val *AddNewNoteOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNewNoteOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNewNoteOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNewNoteOptionRequest(val *AddNewNoteOptionRequest) *NullableAddNewNoteOptionRequest {
	return &NullableAddNewNoteOptionRequest{value: val, isSet: true}
}

func (v NullableAddNewNoteOptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNewNoteOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


