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

// checks if the PaginatedProductTypeGroupListPrefetch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedProductTypeGroupListPrefetch{}

// PaginatedProductTypeGroupListPrefetch struct for PaginatedProductTypeGroupListPrefetch
type PaginatedProductTypeGroupListPrefetch struct {
	Group *map[string]DojoGroup `json:"group,omitempty"`
	ProductType *map[string]ProductType `json:"product_type,omitempty"`
	Role *map[string]Role `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedProductTypeGroupListPrefetch PaginatedProductTypeGroupListPrefetch

// NewPaginatedProductTypeGroupListPrefetch instantiates a new PaginatedProductTypeGroupListPrefetch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProductTypeGroupListPrefetch() *PaginatedProductTypeGroupListPrefetch {
	this := PaginatedProductTypeGroupListPrefetch{}
	return &this
}

// NewPaginatedProductTypeGroupListPrefetchWithDefaults instantiates a new PaginatedProductTypeGroupListPrefetch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProductTypeGroupListPrefetchWithDefaults() *PaginatedProductTypeGroupListPrefetch {
	this := PaginatedProductTypeGroupListPrefetch{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PaginatedProductTypeGroupListPrefetch) GetGroup() map[string]DojoGroup {
	if o == nil || IsNil(o.Group) {
		var ret map[string]DojoGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductTypeGroupListPrefetch) GetGroupOk() (*map[string]DojoGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PaginatedProductTypeGroupListPrefetch) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given map[string]DojoGroup and assigns it to the Group field.
func (o *PaginatedProductTypeGroupListPrefetch) SetGroup(v map[string]DojoGroup) {
	o.Group = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *PaginatedProductTypeGroupListPrefetch) GetProductType() map[string]ProductType {
	if o == nil || IsNil(o.ProductType) {
		var ret map[string]ProductType
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductTypeGroupListPrefetch) GetProductTypeOk() (*map[string]ProductType, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *PaginatedProductTypeGroupListPrefetch) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given map[string]ProductType and assigns it to the ProductType field.
func (o *PaginatedProductTypeGroupListPrefetch) SetProductType(v map[string]ProductType) {
	o.ProductType = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PaginatedProductTypeGroupListPrefetch) GetRole() map[string]Role {
	if o == nil || IsNil(o.Role) {
		var ret map[string]Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductTypeGroupListPrefetch) GetRoleOk() (*map[string]Role, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PaginatedProductTypeGroupListPrefetch) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given map[string]Role and assigns it to the Role field.
func (o *PaginatedProductTypeGroupListPrefetch) SetRole(v map[string]Role) {
	o.Role = &v
}

func (o PaginatedProductTypeGroupListPrefetch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedProductTypeGroupListPrefetch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.ProductType) {
		toSerialize["product_type"] = o.ProductType
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedProductTypeGroupListPrefetch) UnmarshalJSON(data []byte) (err error) {
	varPaginatedProductTypeGroupListPrefetch := _PaginatedProductTypeGroupListPrefetch{}

	err = json.Unmarshal(data, &varPaginatedProductTypeGroupListPrefetch)

	if err != nil {
		return err
	}

	*o = PaginatedProductTypeGroupListPrefetch(varPaginatedProductTypeGroupListPrefetch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "product_type")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedProductTypeGroupListPrefetch struct {
	value *PaginatedProductTypeGroupListPrefetch
	isSet bool
}

func (v NullablePaginatedProductTypeGroupListPrefetch) Get() *PaginatedProductTypeGroupListPrefetch {
	return v.value
}

func (v *NullablePaginatedProductTypeGroupListPrefetch) Set(val *PaginatedProductTypeGroupListPrefetch) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProductTypeGroupListPrefetch) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProductTypeGroupListPrefetch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProductTypeGroupListPrefetch(val *PaginatedProductTypeGroupListPrefetch) *NullablePaginatedProductTypeGroupListPrefetch {
	return &NullablePaginatedProductTypeGroupListPrefetch{value: val, isSet: true}
}

func (v NullablePaginatedProductTypeGroupListPrefetch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProductTypeGroupListPrefetch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


