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

// checks if the ProductAPIScanConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAPIScanConfiguration{}

// ProductAPIScanConfiguration struct for ProductAPIScanConfiguration
type ProductAPIScanConfiguration struct {
	Id int32 `json:"id"`
	ServiceKey1 NullableString `json:"service_key_1,omitempty"`
	ServiceKey2 NullableString `json:"service_key_2,omitempty"`
	ServiceKey3 NullableString `json:"service_key_3,omitempty"`
	Product int32 `json:"product"`
	ToolConfiguration int32 `json:"tool_configuration"`
}

type _ProductAPIScanConfiguration ProductAPIScanConfiguration

// NewProductAPIScanConfiguration instantiates a new ProductAPIScanConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAPIScanConfiguration(id int32, product int32, toolConfiguration int32) *ProductAPIScanConfiguration {
	this := ProductAPIScanConfiguration{}
	this.Id = id
	this.Product = product
	this.ToolConfiguration = toolConfiguration
	return &this
}

// NewProductAPIScanConfigurationWithDefaults instantiates a new ProductAPIScanConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAPIScanConfigurationWithDefaults() *ProductAPIScanConfiguration {
	this := ProductAPIScanConfiguration{}
	return &this
}

// GetId returns the Id field value
func (o *ProductAPIScanConfiguration) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductAPIScanConfiguration) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductAPIScanConfiguration) SetId(v int32) {
	o.Id = v
}

// GetServiceKey1 returns the ServiceKey1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAPIScanConfiguration) GetServiceKey1() string {
	if o == nil || IsNil(o.ServiceKey1.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceKey1.Get()
}

// GetServiceKey1Ok returns a tuple with the ServiceKey1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAPIScanConfiguration) GetServiceKey1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceKey1.Get(), o.ServiceKey1.IsSet()
}

// HasServiceKey1 returns a boolean if a field has been set.
func (o *ProductAPIScanConfiguration) HasServiceKey1() bool {
	if o != nil && o.ServiceKey1.IsSet() {
		return true
	}

	return false
}

// SetServiceKey1 gets a reference to the given NullableString and assigns it to the ServiceKey1 field.
func (o *ProductAPIScanConfiguration) SetServiceKey1(v string) {
	o.ServiceKey1.Set(&v)
}
// SetServiceKey1Nil sets the value for ServiceKey1 to be an explicit nil
func (o *ProductAPIScanConfiguration) SetServiceKey1Nil() {
	o.ServiceKey1.Set(nil)
}

// UnsetServiceKey1 ensures that no value is present for ServiceKey1, not even an explicit nil
func (o *ProductAPIScanConfiguration) UnsetServiceKey1() {
	o.ServiceKey1.Unset()
}

// GetServiceKey2 returns the ServiceKey2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAPIScanConfiguration) GetServiceKey2() string {
	if o == nil || IsNil(o.ServiceKey2.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceKey2.Get()
}

// GetServiceKey2Ok returns a tuple with the ServiceKey2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAPIScanConfiguration) GetServiceKey2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceKey2.Get(), o.ServiceKey2.IsSet()
}

// HasServiceKey2 returns a boolean if a field has been set.
func (o *ProductAPIScanConfiguration) HasServiceKey2() bool {
	if o != nil && o.ServiceKey2.IsSet() {
		return true
	}

	return false
}

// SetServiceKey2 gets a reference to the given NullableString and assigns it to the ServiceKey2 field.
func (o *ProductAPIScanConfiguration) SetServiceKey2(v string) {
	o.ServiceKey2.Set(&v)
}
// SetServiceKey2Nil sets the value for ServiceKey2 to be an explicit nil
func (o *ProductAPIScanConfiguration) SetServiceKey2Nil() {
	o.ServiceKey2.Set(nil)
}

// UnsetServiceKey2 ensures that no value is present for ServiceKey2, not even an explicit nil
func (o *ProductAPIScanConfiguration) UnsetServiceKey2() {
	o.ServiceKey2.Unset()
}

// GetServiceKey3 returns the ServiceKey3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductAPIScanConfiguration) GetServiceKey3() string {
	if o == nil || IsNil(o.ServiceKey3.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceKey3.Get()
}

// GetServiceKey3Ok returns a tuple with the ServiceKey3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductAPIScanConfiguration) GetServiceKey3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceKey3.Get(), o.ServiceKey3.IsSet()
}

// HasServiceKey3 returns a boolean if a field has been set.
func (o *ProductAPIScanConfiguration) HasServiceKey3() bool {
	if o != nil && o.ServiceKey3.IsSet() {
		return true
	}

	return false
}

// SetServiceKey3 gets a reference to the given NullableString and assigns it to the ServiceKey3 field.
func (o *ProductAPIScanConfiguration) SetServiceKey3(v string) {
	o.ServiceKey3.Set(&v)
}
// SetServiceKey3Nil sets the value for ServiceKey3 to be an explicit nil
func (o *ProductAPIScanConfiguration) SetServiceKey3Nil() {
	o.ServiceKey3.Set(nil)
}

// UnsetServiceKey3 ensures that no value is present for ServiceKey3, not even an explicit nil
func (o *ProductAPIScanConfiguration) UnsetServiceKey3() {
	o.ServiceKey3.Unset()
}

// GetProduct returns the Product field value
func (o *ProductAPIScanConfiguration) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ProductAPIScanConfiguration) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ProductAPIScanConfiguration) SetProduct(v int32) {
	o.Product = v
}

// GetToolConfiguration returns the ToolConfiguration field value
func (o *ProductAPIScanConfiguration) GetToolConfiguration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToolConfiguration
}

// GetToolConfigurationOk returns a tuple with the ToolConfiguration field value
// and a boolean to check if the value has been set.
func (o *ProductAPIScanConfiguration) GetToolConfigurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolConfiguration, true
}

// SetToolConfiguration sets field value
func (o *ProductAPIScanConfiguration) SetToolConfiguration(v int32) {
	o.ToolConfiguration = v
}

func (o ProductAPIScanConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductAPIScanConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.ServiceKey1.IsSet() {
		toSerialize["service_key_1"] = o.ServiceKey1.Get()
	}
	if o.ServiceKey2.IsSet() {
		toSerialize["service_key_2"] = o.ServiceKey2.Get()
	}
	if o.ServiceKey3.IsSet() {
		toSerialize["service_key_3"] = o.ServiceKey3.Get()
	}
	toSerialize["product"] = o.Product
	toSerialize["tool_configuration"] = o.ToolConfiguration
	return toSerialize, nil
}

func (o *ProductAPIScanConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"product",
		"tool_configuration",
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

	varProductAPIScanConfiguration := _ProductAPIScanConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductAPIScanConfiguration)

	if err != nil {
		return err
	}

	*o = ProductAPIScanConfiguration(varProductAPIScanConfiguration)

	return err
}

type NullableProductAPIScanConfiguration struct {
	value *ProductAPIScanConfiguration
	isSet bool
}

func (v NullableProductAPIScanConfiguration) Get() *ProductAPIScanConfiguration {
	return v.value
}

func (v *NullableProductAPIScanConfiguration) Set(val *ProductAPIScanConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAPIScanConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAPIScanConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAPIScanConfiguration(val *ProductAPIScanConfiguration) *NullableProductAPIScanConfiguration {
	return &NullableProductAPIScanConfiguration{value: val, isSet: true}
}

func (v NullableProductAPIScanConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductAPIScanConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


