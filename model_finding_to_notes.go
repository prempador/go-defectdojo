/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"fmt"
)

// checks if the FindingToNotes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindingToNotes{}

// FindingToNotes struct for FindingToNotes
type FindingToNotes struct {
	FindingId NullableInt32 `json:"finding_id"`
	Notes []Note `json:"notes"`
	AdditionalProperties map[string]interface{}
}

type _FindingToNotes FindingToNotes

// NewFindingToNotes instantiates a new FindingToNotes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindingToNotes(findingId NullableInt32, notes []Note) *FindingToNotes {
	this := FindingToNotes{}
	this.FindingId = findingId
	this.Notes = notes
	return &this
}

// NewFindingToNotesWithDefaults instantiates a new FindingToNotes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindingToNotesWithDefaults() *FindingToNotes {
	this := FindingToNotes{}
	return &this
}

// GetFindingId returns the FindingId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *FindingToNotes) GetFindingId() int32 {
	if o == nil || o.FindingId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.FindingId.Get()
}

// GetFindingIdOk returns a tuple with the FindingId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FindingToNotes) GetFindingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FindingId.Get(), o.FindingId.IsSet()
}

// SetFindingId sets field value
func (o *FindingToNotes) SetFindingId(v int32) {
	o.FindingId.Set(&v)
}


// GetNotes returns the Notes field value
func (o *FindingToNotes) GetNotes() []Note {
	if o == nil {
		var ret []Note
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *FindingToNotes) GetNotesOk() ([]Note, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *FindingToNotes) SetNotes(v []Note) {
	o.Notes = v
}


func (o FindingToNotes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindingToNotes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["finding_id"] = o.FindingId.Get()
	toSerialize["notes"] = o.Notes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FindingToNotes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"finding_id",
		"notes",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varFindingToNotes := _FindingToNotes{}

	err = json.Unmarshal(data, &varFindingToNotes)

	if err != nil {
		return err
	}

	*o = FindingToNotes(varFindingToNotes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "finding_id")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFindingToNotes struct {
	value *FindingToNotes
	isSet bool
}

func (v NullableFindingToNotes) Get() *FindingToNotes {
	return v.value
}

func (v *NullableFindingToNotes) Set(val *FindingToNotes) {
	v.value = val
	v.isSet = true
}

func (v NullableFindingToNotes) IsSet() bool {
	return v.isSet
}

func (v *NullableFindingToNotes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindingToNotes(val *FindingToNotes) *NullableFindingToNotes {
	return &NullableFindingToNotes{value: val, isSet: true}
}

func (v NullableFindingToNotes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindingToNotes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


