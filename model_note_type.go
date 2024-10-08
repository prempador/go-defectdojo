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

// checks if the NoteType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteType{}

// NoteType struct for NoteType
type NoteType struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	IsSingle *bool `json:"is_single,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	IsMandatory *bool `json:"is_mandatory,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NoteType NoteType

// NewNoteType instantiates a new NoteType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteType(id int32, name string, description string) *NoteType {
	this := NoteType{}
	this.Id = id
	this.Name = name
	this.Description = description
	return &this
}

// NewNoteTypeWithDefaults instantiates a new NoteType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteTypeWithDefaults() *NoteType {
	this := NoteType{}
	return &this
}

// GetId returns the Id field value
func (o *NoteType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NoteType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NoteType) SetId(v int32) {
	o.Id = v
}


// GetName returns the Name field value
func (o *NoteType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NoteType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NoteType) SetName(v string) {
	o.Name = v
}


// GetDescription returns the Description field value
func (o *NoteType) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NoteType) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NoteType) SetDescription(v string) {
	o.Description = v
}


// GetIsSingle returns the IsSingle field value if set, zero value otherwise.
func (o *NoteType) GetIsSingle() bool {
	if o == nil || IsNil(o.IsSingle) {
		var ret bool
		return ret
	}
	return *o.IsSingle
}

// GetIsSingleOk returns a tuple with the IsSingle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteType) GetIsSingleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSingle) {
		return nil, false
	}
	return o.IsSingle, true
}

// HasIsSingle returns a boolean if a field has been set.
func (o *NoteType) HasIsSingle() bool {
	if o != nil && !IsNil(o.IsSingle) {
		return true
	}

	return false
}

// SetIsSingle gets a reference to the given bool and assigns it to the IsSingle field.
func (o *NoteType) SetIsSingle(v bool) {
	o.IsSingle = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *NoteType) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteType) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *NoteType) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *NoteType) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsMandatory returns the IsMandatory field value if set, zero value otherwise.
func (o *NoteType) GetIsMandatory() bool {
	if o == nil || IsNil(o.IsMandatory) {
		var ret bool
		return ret
	}
	return *o.IsMandatory
}

// GetIsMandatoryOk returns a tuple with the IsMandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoteType) GetIsMandatoryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMandatory) {
		return nil, false
	}
	return o.IsMandatory, true
}

// HasIsMandatory returns a boolean if a field has been set.
func (o *NoteType) HasIsMandatory() bool {
	if o != nil && !IsNil(o.IsMandatory) {
		return true
	}

	return false
}

// SetIsMandatory gets a reference to the given bool and assigns it to the IsMandatory field.
func (o *NoteType) SetIsMandatory(v bool) {
	o.IsMandatory = &v
}

func (o NoteType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.IsSingle) {
		toSerialize["is_single"] = o.IsSingle
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsMandatory) {
		toSerialize["is_mandatory"] = o.IsMandatory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NoteType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
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
	varNoteType := _NoteType{}

	err = json.Unmarshal(data, &varNoteType)

	if err != nil {
		return err
	}

	*o = NoteType(varNoteType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "is_single")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_mandatory")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNoteType struct {
	value *NoteType
	isSet bool
}

func (v NullableNoteType) Get() *NoteType {
	return v.value
}

func (v *NullableNoteType) Set(val *NoteType) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteType) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteType(val *NoteType) *NullableNoteType {
	return &NullableNoteType{value: val, isSet: true}
}

func (v NullableNoteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


