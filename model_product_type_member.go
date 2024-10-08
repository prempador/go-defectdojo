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

// checks if the ProductTypeMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductTypeMember{}

// ProductTypeMember struct for ProductTypeMember
type ProductTypeMember struct {
	Id int32 `json:"id"`
	ProductType int32 `json:"product_type"`
	User int32 `json:"user"`
	Role int32 `json:"role"`
	Prefetch *PaginatedProductTypeMemberListPrefetch `json:"prefetch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductTypeMember ProductTypeMember

// NewProductTypeMember instantiates a new ProductTypeMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductTypeMember(id int32, productType int32, user int32, role int32) *ProductTypeMember {
	this := ProductTypeMember{}
	this.Id = id
	this.ProductType = productType
	this.User = user
	this.Role = role
	return &this
}

// NewProductTypeMemberWithDefaults instantiates a new ProductTypeMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductTypeMemberWithDefaults() *ProductTypeMember {
	this := ProductTypeMember{}
	return &this
}

// GetId returns the Id field value
func (o *ProductTypeMember) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMember) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductTypeMember) SetId(v int32) {
	o.Id = v
}


// GetProductType returns the ProductType field value
func (o *ProductTypeMember) GetProductType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMember) GetProductTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ProductTypeMember) SetProductType(v int32) {
	o.ProductType = v
}


// GetUser returns the User field value
func (o *ProductTypeMember) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMember) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ProductTypeMember) SetUser(v int32) {
	o.User = v
}


// GetRole returns the Role field value
func (o *ProductTypeMember) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProductTypeMember) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProductTypeMember) SetRole(v int32) {
	o.Role = v
}


// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *ProductTypeMember) GetPrefetch() PaginatedProductTypeMemberListPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret PaginatedProductTypeMemberListPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductTypeMember) GetPrefetchOk() (*PaginatedProductTypeMemberListPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *ProductTypeMember) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given PaginatedProductTypeMemberListPrefetch and assigns it to the Prefetch field.
func (o *ProductTypeMember) SetPrefetch(v PaginatedProductTypeMemberListPrefetch) {
	o.Prefetch = &v
}

func (o ProductTypeMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductTypeMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["product_type"] = o.ProductType
	toSerialize["user"] = o.User
	toSerialize["role"] = o.Role
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductTypeMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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
	varProductTypeMember := _ProductTypeMember{}

	err = json.Unmarshal(data, &varProductTypeMember)

	if err != nil {
		return err
	}

	*o = ProductTypeMember(varProductTypeMember)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "role")
		delete(additionalProperties, "prefetch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductTypeMember struct {
	value *ProductTypeMember
	isSet bool
}

func (v NullableProductTypeMember) Get() *ProductTypeMember {
	return v.value
}

func (v *NullableProductTypeMember) Set(val *ProductTypeMember) {
	v.value = val
	v.isSet = true
}

func (v NullableProductTypeMember) IsSet() bool {
	return v.isSet
}

func (v *NullableProductTypeMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductTypeMember(val *ProductTypeMember) *NullableProductTypeMember {
	return &NullableProductTypeMember{value: val, isSet: true}
}

func (v NullableProductTypeMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductTypeMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


