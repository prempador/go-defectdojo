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

// checks if the ProductTypeMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTypeMemberRequest{}

// ProductTypeMemberRequest struct for ProductTypeMemberRequest
type ProductTypeMemberRequest struct {
	ProductType int32 `json:"product_type"`
	User int32 `json:"user"`
	Role int32 `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _ProductTypeMemberRequest ProductTypeMemberRequest

// NewProductTypeMemberRequest instantiates a new ProductTypeMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTypeMemberRequest(productType int32, user int32, role int32) *ProductTypeMemberRequest {
	this := ProductTypeMemberRequest{}
	this.ProductType = productType
	this.User = user
	this.Role = role
	return &this
}

// NewProductTypeMemberRequestWithDefaults instantiates a new ProductTypeMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeMemberRequestWithDefaults() *ProductTypeMemberRequest {
	this := ProductTypeMemberRequest{}
	return &this
}

// GetProductType returns the ProductType field value
func (o *ProductTypeMemberRequest) GetProductType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMemberRequest) GetProductTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ProductTypeMemberRequest) SetProductType(v int32) {
	o.ProductType = v
}


// GetUser returns the User field value
func (o *ProductTypeMemberRequest) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMemberRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ProductTypeMemberRequest) SetUser(v int32) {
	o.User = v
}


// GetRole returns the Role field value
func (o *ProductTypeMemberRequest) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMemberRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProductTypeMemberRequest) SetRole(v int32) {
	o.Role = v
}


func (o ProductTypeMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductTypeMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_type"] = o.ProductType
	toSerialize["user"] = o.User
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductTypeMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_type",
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
	varProductTypeMemberRequest := _ProductTypeMemberRequest{}

	err = json.Unmarshal(data, &varProductTypeMemberRequest)

	if err != nil {
		return err
	}

	*o = ProductTypeMemberRequest(varProductTypeMemberRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductTypeMemberRequest struct {
	value *ProductTypeMemberRequest
	isSet bool
}

func (v NullableProductTypeMemberRequest) Get() *ProductTypeMemberRequest {
	return v.value
}

func (v *NullableProductTypeMemberRequest) Set(val *ProductTypeMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTypeMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTypeMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTypeMemberRequest(val *ProductTypeMemberRequest) *NullableProductTypeMemberRequest {
	return &NullableProductTypeMemberRequest{value: val, isSet: true}
}

func (v NullableProductTypeMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTypeMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


