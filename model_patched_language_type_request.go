/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedLanguageTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedLanguageTypeRequest{}

// PatchedLanguageTypeRequest struct for PatchedLanguageTypeRequest
type PatchedLanguageTypeRequest struct {
	Language *string `json:"language,omitempty"`
	Color NullableString `json:"color,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedLanguageTypeRequest PatchedLanguageTypeRequest

// NewPatchedLanguageTypeRequest instantiates a new PatchedLanguageTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLanguageTypeRequest() *PatchedLanguageTypeRequest {
	this := PatchedLanguageTypeRequest{}
	return &this
}

// NewPatchedLanguageTypeRequestWithDefaults instantiates a new PatchedLanguageTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLanguageTypeRequestWithDefaults() *PatchedLanguageTypeRequest {
	this := PatchedLanguageTypeRequest{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PatchedLanguageTypeRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLanguageTypeRequest) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PatchedLanguageTypeRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PatchedLanguageTypeRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLanguageTypeRequest) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLanguageTypeRequest) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *PatchedLanguageTypeRequest) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *PatchedLanguageTypeRequest) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *PatchedLanguageTypeRequest) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *PatchedLanguageTypeRequest) UnsetColor() {
	o.Color.Unset()
}

func (o PatchedLanguageTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedLanguageTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedLanguageTypeRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedLanguageTypeRequest := _PatchedLanguageTypeRequest{}

	err = json.Unmarshal(data, &varPatchedLanguageTypeRequest)

	if err != nil {
		return err
	}

	*o = PatchedLanguageTypeRequest(varPatchedLanguageTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "language")
		delete(additionalProperties, "color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedLanguageTypeRequest struct {
	value *PatchedLanguageTypeRequest
	isSet bool
}

func (v NullablePatchedLanguageTypeRequest) Get() *PatchedLanguageTypeRequest {
	return v.value
}

func (v *NullablePatchedLanguageTypeRequest) Set(val *PatchedLanguageTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLanguageTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLanguageTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLanguageTypeRequest(val *PatchedLanguageTypeRequest) *NullablePatchedLanguageTypeRequest {
	return &NullablePatchedLanguageTypeRequest{value: val, isSet: true}
}

func (v NullablePatchedLanguageTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLanguageTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


