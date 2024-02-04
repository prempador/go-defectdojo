/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ProductGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductGroup{}

// ProductGroup struct for ProductGroup
type ProductGroup struct {
	Id int32 `json:"id"`
	Product int32 `json:"product"`
	Group int32 `json:"group"`
	Role int32 `json:"role"`
	Prefetch *PaginatedProductGroupListPrefetch `json:"prefetch,omitempty"`
}

type _ProductGroup ProductGroup

// NewProductGroup instantiates a new ProductGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductGroup(id int32, product int32, group int32, role int32) *ProductGroup {
	this := ProductGroup{}
	this.Id = id
	this.Product = product
	this.Group = group
	this.Role = role
	return &this
}

// NewProductGroupWithDefaults instantiates a new ProductGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductGroupWithDefaults() *ProductGroup {
	this := ProductGroup{}
	return &this
}

// GetId returns the Id field value
func (o *ProductGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductGroup) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductGroup) SetId(v int32) {
	o.Id = v
}

// GetProduct returns the Product field value
func (o *ProductGroup) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ProductGroup) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ProductGroup) SetProduct(v int32) {
	o.Product = v
}

// GetGroup returns the Group field value
func (o *ProductGroup) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ProductGroup) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ProductGroup) SetGroup(v int32) {
	o.Group = v
}

// GetRole returns the Role field value
func (o *ProductGroup) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ProductGroup) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ProductGroup) SetRole(v int32) {
	o.Role = v
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *ProductGroup) GetPrefetch() PaginatedProductGroupListPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret PaginatedProductGroupListPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductGroup) GetPrefetchOk() (*PaginatedProductGroupListPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *ProductGroup) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given PaginatedProductGroupListPrefetch and assigns it to the Prefetch field.
func (o *ProductGroup) SetPrefetch(v PaginatedProductGroupListPrefetch) {
	o.Prefetch = &v
}

func (o ProductGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["product"] = o.Product
	toSerialize["group"] = o.Group
	toSerialize["role"] = o.Role
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}
	return toSerialize, nil
}

func (o *ProductGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"product",
		"group",
		"role",
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

	varProductGroup := _ProductGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductGroup)

	if err != nil {
		return err
	}

	*o = ProductGroup(varProductGroup)

	return err
}

type NullableProductGroup struct {
	value *ProductGroup
	isSet bool
}

func (v NullableProductGroup) Get() *ProductGroup {
	return v.value
}

func (v *NullableProductGroup) Set(val *ProductGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableProductGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableProductGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductGroup(val *ProductGroup) *NullableProductGroup {
	return &NullableProductGroup{value: val, isSet: true}
}

func (v NullableProductGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


