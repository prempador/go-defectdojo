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

// checks if the LanguageType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguageType{}

// LanguageType struct for LanguageType
type LanguageType struct {
	Id int32 `json:"id"`
	Language string `json:"language"`
	Color NullableString `json:"color,omitempty"`
}

type _LanguageType LanguageType

// NewLanguageType instantiates a new LanguageType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageType(id int32, language string) *LanguageType {
	this := LanguageType{}
	this.Id = id
	this.Language = language
	return &this
}

// NewLanguageTypeWithDefaults instantiates a new LanguageType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageTypeWithDefaults() *LanguageType {
	this := LanguageType{}
	return &this
}

// GetId returns the Id field value
func (o *LanguageType) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LanguageType) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LanguageType) SetId(v int32) {
	o.Id = v
}

// GetLanguage returns the Language field value
func (o *LanguageType) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *LanguageType) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *LanguageType) SetLanguage(v string) {
	o.Language = v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LanguageType) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LanguageType) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *LanguageType) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *LanguageType) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *LanguageType) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *LanguageType) UnsetColor() {
	o.Color.Unset()
}

func (o LanguageType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguageType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["language"] = o.Language
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	return toSerialize, nil
}

func (o *LanguageType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"language",
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

	varLanguageType := _LanguageType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLanguageType)

	if err != nil {
		return err
	}

	*o = LanguageType(varLanguageType)

	return err
}

type NullableLanguageType struct {
	value *LanguageType
	isSet bool
}

func (v NullableLanguageType) Get() *LanguageType {
	return v.value
}

func (v *NullableLanguageType) Set(val *LanguageType) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageType) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageType(val *LanguageType) *NullableLanguageType {
	return &NullableLanguageType{value: val, isSet: true}
}

func (v NullableLanguageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


