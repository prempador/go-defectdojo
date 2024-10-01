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

// checks if the ProductGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductGroupRequest{}

// ProductGroupRequest struct for ProductGroupRequest
type ProductGroupRequest struct {
	Product int32 `json:"product"`
	Group int32 `json:"group"`
	Role int32 `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _ProductGroupRequest ProductGroupRequest

// NewProductGroupRequest instantiates a new ProductGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGroupRequest(product int32, group int32, role int32) *ProductGroupRequest {
	this := ProductGroupRequest{}
	this.Product = product
	this.Group = group
	this.Role = role
	return &this
}

// NewProductGroupRequestWithDefaults instantiates a new ProductGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGroupRequestWithDefaults() *ProductGroupRequest {
	this := ProductGroupRequest{}
	return &this
}

// GetProduct returns the Product field value
func (o *ProductGroupRequest) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ProductGroupRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ProductGroupRequest) SetProduct(v int32) {
	o.Product = v
}


// GetGroup returns the Group field value
func (o *ProductGroupRequest) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ProductGroupRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ProductGroupRequest) SetGroup(v int32) {
	o.Group = v
}


// GetRole returns the Role field value
func (o *ProductGroupRequest) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProductGroupRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProductGroupRequest) SetRole(v int32) {
	o.Role = v
}


func (o ProductGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["group"] = o.Group
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
		"group",
		"role",
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
	varProductGroupRequest := _ProductGroupRequest{}

	err = json.Unmarshal(data, &varProductGroupRequest)

	if err != nil {
		return err
	}

	*o = ProductGroupRequest(varProductGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "group")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductGroupRequest struct {
	value *ProductGroupRequest
	isSet bool
}

func (v NullableProductGroupRequest) Get() *ProductGroupRequest {
	return v.value
}

func (v *NullableProductGroupRequest) Set(val *ProductGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGroupRequest(val *ProductGroupRequest) *NullableProductGroupRequest {
	return &NullableProductGroupRequest{value: val, isSet: true}
}

func (v NullableProductGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


