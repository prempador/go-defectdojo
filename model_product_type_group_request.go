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

// checks if the ProductTypeGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTypeGroupRequest{}

// ProductTypeGroupRequest struct for ProductTypeGroupRequest
type ProductTypeGroupRequest struct {
	ProductType int32 `json:"product_type"`
	Group int32 `json:"group"`
	Role int32 `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _ProductTypeGroupRequest ProductTypeGroupRequest

// NewProductTypeGroupRequest instantiates a new ProductTypeGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTypeGroupRequest(productType int32, group int32, role int32) *ProductTypeGroupRequest {
	this := ProductTypeGroupRequest{}
	this.ProductType = productType
	this.Group = group
	this.Role = role
	return &this
}

// NewProductTypeGroupRequestWithDefaults instantiates a new ProductTypeGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeGroupRequestWithDefaults() *ProductTypeGroupRequest {
	this := ProductTypeGroupRequest{}
	return &this
}

// GetProductType returns the ProductType field value
func (o *ProductTypeGroupRequest) GetProductType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ProductTypeGroupRequest) GetProductTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ProductTypeGroupRequest) SetProductType(v int32) {
	o.ProductType = v
}


// GetGroup returns the Group field value
func (o *ProductTypeGroupRequest) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ProductTypeGroupRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ProductTypeGroupRequest) SetGroup(v int32) {
	o.Group = v
}


// GetRole returns the Role field value
func (o *ProductTypeGroupRequest) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProductTypeGroupRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProductTypeGroupRequest) SetRole(v int32) {
	o.Role = v
}


func (o ProductTypeGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductTypeGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_type"] = o.ProductType
	toSerialize["group"] = o.Group
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductTypeGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_type",
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
	varProductTypeGroupRequest := _ProductTypeGroupRequest{}

	err = json.Unmarshal(data, &varProductTypeGroupRequest)

	if err != nil {
		return err
	}

	*o = ProductTypeGroupRequest(varProductTypeGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "group")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductTypeGroupRequest struct {
	value *ProductTypeGroupRequest
	isSet bool
}

func (v NullableProductTypeGroupRequest) Get() *ProductTypeGroupRequest {
	return v.value
}

func (v *NullableProductTypeGroupRequest) Set(val *ProductTypeGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTypeGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTypeGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTypeGroupRequest(val *ProductTypeGroupRequest) *NullableProductTypeGroupRequest {
	return &NullableProductTypeGroupRequest{value: val, isSet: true}
}

func (v NullableProductTypeGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTypeGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


