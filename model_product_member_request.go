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

// checks if the ProductMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMemberRequest{}

// ProductMemberRequest struct for ProductMemberRequest
type ProductMemberRequest struct {
	Product int32 `json:"product"`
	User int32 `json:"user"`
	Role int32 `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _ProductMemberRequest ProductMemberRequest

// NewProductMemberRequest instantiates a new ProductMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMemberRequest(product int32, user int32, role int32) *ProductMemberRequest {
	this := ProductMemberRequest{}
	this.Product = product
	this.User = user
	this.Role = role
	return &this
}

// NewProductMemberRequestWithDefaults instantiates a new ProductMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMemberRequestWithDefaults() *ProductMemberRequest {
	this := ProductMemberRequest{}
	return &this
}

// GetProduct returns the Product field value
func (o *ProductMemberRequest) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ProductMemberRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ProductMemberRequest) SetProduct(v int32) {
	o.Product = v
}


// GetUser returns the User field value
func (o *ProductMemberRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ProductMemberRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ProductMemberRequest) SetUser(v int32) {
	o.User = v
}


// GetRole returns the Role field value
func (o *ProductMemberRequest) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProductMemberRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProductMemberRequest) SetRole(v int32) {
	o.Role = v
}


func (o ProductMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["user"] = o.User
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product",
		"user",
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
	varProductMemberRequest := _ProductMemberRequest{}

	err = json.Unmarshal(data, &varProductMemberRequest)

	if err != nil {
		return err
	}

	*o = ProductMemberRequest(varProductMemberRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "user")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductMemberRequest struct {
	value *ProductMemberRequest
	isSet bool
}

func (v NullableProductMemberRequest) Get() *ProductMemberRequest {
	return v.value
}

func (v *NullableProductMemberRequest) Set(val *ProductMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMemberRequest(val *ProductMemberRequest) *NullableProductMemberRequest {
	return &NullableProductMemberRequest{value: val, isSet: true}
}

func (v NullableProductMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


