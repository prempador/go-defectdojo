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

// checks if the ImportLanguages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportLanguages{}

// ImportLanguages struct for ImportLanguages
type ImportLanguages struct {
	Product int32 `json:"product"`
	File string `json:"file"`
}

type _ImportLanguages ImportLanguages

// NewImportLanguages instantiates a new ImportLanguages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportLanguages(product int32, file string) *ImportLanguages {
	this := ImportLanguages{}
	this.Product = product
	this.File = file
	return &this
}

// NewImportLanguagesWithDefaults instantiates a new ImportLanguages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportLanguagesWithDefaults() *ImportLanguages {
	this := ImportLanguages{}
	return &this
}

// GetProduct returns the Product field value
func (o *ImportLanguages) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ImportLanguages) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ImportLanguages) SetProduct(v int32) {
	o.Product = v
}

// GetFile returns the File field value
func (o *ImportLanguages) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ImportLanguages) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ImportLanguages) SetFile(v string) {
	o.File = v
}

func (o ImportLanguages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportLanguages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["file"] = o.File
	return toSerialize, nil
}

func (o *ImportLanguages) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
		"file",
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

	varImportLanguages := _ImportLanguages{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportLanguages)

	if err != nil {
		return err
	}

	*o = ImportLanguages(varImportLanguages)

	return err
}

type NullableImportLanguages struct {
	value *ImportLanguages
	isSet bool
}

func (v NullableImportLanguages) Get() *ImportLanguages {
	return v.value
}

func (v *NullableImportLanguages) Set(val *ImportLanguages) {
	v.value = val
	v.isSet = true
}

func (v NullableImportLanguages) IsSet() bool {
	return v.isSet
}

func (v *NullableImportLanguages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportLanguages(val *ImportLanguages) *NullableImportLanguages {
	return &NullableImportLanguages{value: val, isSet: true}
}

func (v NullableImportLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportLanguages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


