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

// checks if the DeletePreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePreview{}

// DeletePreview struct for DeletePreview
type DeletePreview struct {
	Model string `json:"model"`
	Id NullableInt32 `json:"id"`
	Name string `json:"name"`
}

type _DeletePreview DeletePreview

// NewDeletePreview instantiates a new DeletePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePreview(model string, id NullableInt32, name string) *DeletePreview {
	this := DeletePreview{}
	this.Model = model
	this.Id = id
	this.Name = name
	return &this
}

// NewDeletePreviewWithDefaults instantiates a new DeletePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePreviewWithDefaults() *DeletePreview {
	this := DeletePreview{}
	return &this
}

// GetModel returns the Model field value
func (o *DeletePreview) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *DeletePreview) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *DeletePreview) SetModel(v string) {
	o.Model = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DeletePreview) GetId() int32 {
	if o == nil || o.Id.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletePreview) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *DeletePreview) SetId(v int32) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
func (o *DeletePreview) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeletePreview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeletePreview) SetName(v string) {
	o.Name = v
}

func (o DeletePreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *DeletePreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"model",
		"id",
		"name",
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

	varDeletePreview := _DeletePreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeletePreview)

	if err != nil {
		return err
	}

	*o = DeletePreview(varDeletePreview)

	return err
}

type NullableDeletePreview struct {
	value *DeletePreview
	isSet bool
}

func (v NullableDeletePreview) Get() *DeletePreview {
	return v.value
}

func (v *NullableDeletePreview) Set(val *DeletePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePreview(val *DeletePreview) *NullableDeletePreview {
	return &NullableDeletePreview{value: val, isSet: true}
}

func (v NullableDeletePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


