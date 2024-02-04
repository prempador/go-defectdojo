/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the NoteHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteHistory{}

// NoteHistory struct for NoteHistory
type NoteHistory struct {
	Id int32 `json:"id"`
	CurrentEditor UserStub `json:"current_editor"`
	NoteType NoteType `json:"note_type"`
	Data string `json:"data"`
	Time NullableTime `json:"time"`
}

type _NoteHistory NoteHistory

// NewNoteHistory instantiates a new NoteHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteHistory(id int32, currentEditor UserStub, noteType NoteType, data string, time NullableTime) *NoteHistory {
	this := NoteHistory{}
	this.Id = id
	this.CurrentEditor = currentEditor
	this.NoteType = noteType
	this.Data = data
	this.Time = time
	return &this
}

// NewNoteHistoryWithDefaults instantiates a new NoteHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteHistoryWithDefaults() *NoteHistory {
	this := NoteHistory{}
	return &this
}

// GetId returns the Id field value
func (o *NoteHistory) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NoteHistory) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NoteHistory) SetId(v int32) {
	o.Id = v
}

// GetCurrentEditor returns the CurrentEditor field value
func (o *NoteHistory) GetCurrentEditor() UserStub {
	if o == nil {
		var ret UserStub
		return ret
	}

	return o.CurrentEditor
}

// GetCurrentEditorOk returns a tuple with the CurrentEditor field value
// and a boolean to check if the value has been set.
func (o *NoteHistory) GetCurrentEditorOk() (*UserStub, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentEditor, true
}

// SetCurrentEditor sets field value
func (o *NoteHistory) SetCurrentEditor(v UserStub) {
	o.CurrentEditor = v
}

// GetNoteType returns the NoteType field value
func (o *NoteHistory) GetNoteType() NoteType {
	if o == nil {
		var ret NoteType
		return ret
	}

	return o.NoteType
}

// GetNoteTypeOk returns a tuple with the NoteType field value
// and a boolean to check if the value has been set.
func (o *NoteHistory) GetNoteTypeOk() (*NoteType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoteType, true
}

// SetNoteType sets field value
func (o *NoteHistory) SetNoteType(v NoteType) {
	o.NoteType = v
}

// GetData returns the Data field value
func (o *NoteHistory) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NoteHistory) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NoteHistory) SetData(v string) {
	o.Data = v
}

// GetTime returns the Time field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *NoteHistory) GetTime() time.Time {
	if o == nil || o.Time.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NoteHistory) GetTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// SetTime sets field value
func (o *NoteHistory) SetTime(v time.Time) {
	o.Time.Set(&v)
}

func (o NoteHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["current_editor"] = o.CurrentEditor
	toSerialize["note_type"] = o.NoteType
	toSerialize["data"] = o.Data
	toSerialize["time"] = o.Time.Get()
	return toSerialize, nil
}

func (o *NoteHistory) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"current_editor",
		"note_type",
		"data",
		"time",
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

	varNoteHistory := _NoteHistory{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNoteHistory)

	if err != nil {
		return err
	}

	*o = NoteHistory(varNoteHistory)

	return err
}

type NullableNoteHistory struct {
	value *NoteHistory
	isSet bool
}

func (v NullableNoteHistory) Get() *NoteHistory {
	return v.value
}

func (v *NullableNoteHistory) Set(val *NoteHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteHistory(val *NoteHistory) *NullableNoteHistory {
	return &NullableNoteHistory{value: val, isSet: true}
}

func (v NullableNoteHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


