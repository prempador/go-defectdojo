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

// checks if the PatchedNoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedNoteRequest{}

// PatchedNoteRequest struct for PatchedNoteRequest
type PatchedNoteRequest struct {
	Entry *string `json:"entry,omitempty"`
	Private *bool `json:"private,omitempty"`
	Edited *bool `json:"edited,omitempty"`
}

// NewPatchedNoteRequest instantiates a new PatchedNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNoteRequest() *PatchedNoteRequest {
	this := PatchedNoteRequest{}
	return &this
}

// NewPatchedNoteRequestWithDefaults instantiates a new PatchedNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNoteRequestWithDefaults() *PatchedNoteRequest {
	this := PatchedNoteRequest{}
	return &this
}

// GetEntry returns the Entry field value if set, zero value otherwise.
func (o *PatchedNoteRequest) GetEntry() string {
	if o == nil || IsNil(o.Entry) {
		var ret string
		return ret
	}
	return *o.Entry
}

// GetEntryOk returns a tuple with the Entry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNoteRequest) GetEntryOk() (*string, bool) {
	if o == nil || IsNil(o.Entry) {
		return nil, false
	}
	return o.Entry, true
}

// HasEntry returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasEntry() bool {
	if o != nil && !IsNil(o.Entry) {
		return true
	}

	return false
}

// SetEntry gets a reference to the given string and assigns it to the Entry field.
func (o *PatchedNoteRequest) SetEntry(v string) {
	o.Entry = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *PatchedNoteRequest) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNoteRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *PatchedNoteRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetEdited returns the Edited field value if set, zero value otherwise.
func (o *PatchedNoteRequest) GetEdited() bool {
	if o == nil || IsNil(o.Edited) {
		var ret bool
		return ret
	}
	return *o.Edited
}

// GetEditedOk returns a tuple with the Edited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNoteRequest) GetEditedOk() (*bool, bool) {
	if o == nil || IsNil(o.Edited) {
		return nil, false
	}
	return o.Edited, true
}

// HasEdited returns a boolean if a field has been set.
func (o *PatchedNoteRequest) HasEdited() bool {
	if o != nil && !IsNil(o.Edited) {
		return true
	}

	return false
}

// SetEdited gets a reference to the given bool and assigns it to the Edited field.
func (o *PatchedNoteRequest) SetEdited(v bool) {
	o.Edited = &v
}

func (o PatchedNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entry) {
		toSerialize["entry"] = o.Entry
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Edited) {
		toSerialize["edited"] = o.Edited
	}
	return toSerialize, nil
}

type NullablePatchedNoteRequest struct {
	value *PatchedNoteRequest
	isSet bool
}

func (v NullablePatchedNoteRequest) Get() *PatchedNoteRequest {
	return v.value
}

func (v *NullablePatchedNoteRequest) Set(val *PatchedNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNoteRequest(val *PatchedNoteRequest) *NullablePatchedNoteRequest {
	return &NullablePatchedNoteRequest{value: val, isSet: true}
}

func (v NullablePatchedNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


