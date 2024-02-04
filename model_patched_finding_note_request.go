/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchedFindingNoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedFindingNoteRequest{}

// PatchedFindingNoteRequest struct for PatchedFindingNoteRequest
type PatchedFindingNoteRequest struct {
	NoteId *int32 `json:"note_id,omitempty"`
}

// NewPatchedFindingNoteRequest instantiates a new PatchedFindingNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedFindingNoteRequest() *PatchedFindingNoteRequest {
	this := PatchedFindingNoteRequest{}
	return &this
}

// NewPatchedFindingNoteRequestWithDefaults instantiates a new PatchedFindingNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedFindingNoteRequestWithDefaults() *PatchedFindingNoteRequest {
	this := PatchedFindingNoteRequest{}
	return &this
}

// GetNoteId returns the NoteId field value if set, zero value otherwise.
func (o *PatchedFindingNoteRequest) GetNoteId() int32 {
	if o == nil || IsNil(o.NoteId) {
		var ret int32
		return ret
	}
	return *o.NoteId
}

// GetNoteIdOk returns a tuple with the NoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedFindingNoteRequest) GetNoteIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NoteId) {
		return nil, false
	}
	return o.NoteId, true
}

// HasNoteId returns a boolean if a field has been set.
func (o *PatchedFindingNoteRequest) HasNoteId() bool {
	if o != nil && !IsNil(o.NoteId) {
		return true
	}

	return false
}

// SetNoteId gets a reference to the given int32 and assigns it to the NoteId field.
func (o *PatchedFindingNoteRequest) SetNoteId(v int32) {
	o.NoteId = &v
}

func (o PatchedFindingNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedFindingNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NoteId) {
		toSerialize["note_id"] = o.NoteId
	}
	return toSerialize, nil
}

type NullablePatchedFindingNoteRequest struct {
	value *PatchedFindingNoteRequest
	isSet bool
}

func (v NullablePatchedFindingNoteRequest) Get() *PatchedFindingNoteRequest {
	return v.value
}

func (v *NullablePatchedFindingNoteRequest) Set(val *PatchedFindingNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedFindingNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedFindingNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedFindingNoteRequest(val *PatchedFindingNoteRequest) *NullablePatchedFindingNoteRequest {
	return &NullablePatchedFindingNoteRequest{value: val, isSet: true}
}

func (v NullablePatchedFindingNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedFindingNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

