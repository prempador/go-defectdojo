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

// checks if the TestToNotes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestToNotes{}

// TestToNotes struct for TestToNotes
type TestToNotes struct {
	TestId NullableInt32 `json:"test_id"`
	Notes []Note `json:"notes"`
	AdditionalProperties map[string]interface{}
}

type _TestToNotes TestToNotes

// NewTestToNotes instantiates a new TestToNotes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestToNotes(testId NullableInt32, notes []Note) *TestToNotes {
	this := TestToNotes{}
	this.TestId = testId
	this.Notes = notes
	return &this
}

// NewTestToNotesWithDefaults instantiates a new TestToNotes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestToNotesWithDefaults() *TestToNotes {
	this := TestToNotes{}
	return &this
}

// GetTestId returns the TestId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TestToNotes) GetTestId() int32 {
	if o == nil || o.TestId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TestId.Get()
}

// GetTestIdOk returns a tuple with the TestId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestToNotes) GetTestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestId.Get(), o.TestId.IsSet()
}

// SetTestId sets field value
func (o *TestToNotes) SetTestId(v int32) {
	o.TestId.Set(&v)
}


// GetNotes returns the Notes field value
func (o *TestToNotes) GetNotes() []Note {
	if o == nil {
		var ret []Note
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *TestToNotes) GetNotesOk() ([]Note, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *TestToNotes) SetNotes(v []Note) {
	o.Notes = v
}


func (o TestToNotes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestToNotes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["test_id"] = o.TestId.Get()
	toSerialize["notes"] = o.Notes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestToNotes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"test_id",
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
	varTestToNotes := _TestToNotes{}

	err = json.Unmarshal(data, &varTestToNotes)

	if err != nil {
		return err
	}

	*o = TestToNotes(varTestToNotes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "test_id")
		delete(additionalProperties, "notes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestToNotes struct {
	value *TestToNotes
	isSet bool
}

func (v NullableTestToNotes) Get() *TestToNotes {
	return v.value
}

func (v *NullableTestToNotes) Set(val *TestToNotes) {
	v.value = val
	v.isSet = true
}

func (v NullableTestToNotes) IsSet() bool {
	return v.isSet
}

func (v *NullableTestToNotes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestToNotes(val *TestToNotes) *NullableTestToNotes {
	return &NullableTestToNotes{value: val, isSet: true}
}

func (v NullableTestToNotes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestToNotes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


