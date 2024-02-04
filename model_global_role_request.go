/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GlobalRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRoleRequest{}

// GlobalRoleRequest struct for GlobalRoleRequest
type GlobalRoleRequest struct {
	User NullableInt32 `json:"user,omitempty"`
	Group NullableInt32 `json:"group,omitempty"`
	// The global role will be applied to all product types and products.
	Role NullableInt32 `json:"role,omitempty"`
}

// NewGlobalRoleRequest instantiates a new GlobalRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRoleRequest() *GlobalRoleRequest {
	this := GlobalRoleRequest{}
	return &this
}

// NewGlobalRoleRequestWithDefaults instantiates a new GlobalRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRoleRequestWithDefaults() *GlobalRoleRequest {
	this := GlobalRoleRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalRoleRequest) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalRoleRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *GlobalRoleRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *GlobalRoleRequest) SetUser(v int32) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *GlobalRoleRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *GlobalRoleRequest) UnsetUser() {
	o.User.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalRoleRequest) GetGroup() int32 {
	if o == nil || IsNil(o.Group.Get()) {
		var ret int32
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalRoleRequest) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *GlobalRoleRequest) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableInt32 and assigns it to the Group field.
func (o *GlobalRoleRequest) SetGroup(v int32) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *GlobalRoleRequest) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *GlobalRoleRequest) UnsetGroup() {
	o.Group.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GlobalRoleRequest) GetRole() int32 {
	if o == nil || IsNil(o.Role.Get()) {
		var ret int32
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GlobalRoleRequest) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *GlobalRoleRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableInt32 and assigns it to the Role field.
func (o *GlobalRoleRequest) SetRole(v int32) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *GlobalRoleRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *GlobalRoleRequest) UnsetRole() {
	o.Role.Unset()
}

func (o GlobalRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	return toSerialize, nil
}

type NullableGlobalRoleRequest struct {
	value *GlobalRoleRequest
	isSet bool
}

func (v NullableGlobalRoleRequest) Get() *GlobalRoleRequest {
	return v.value
}

func (v *NullableGlobalRoleRequest) Set(val *GlobalRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRoleRequest(val *GlobalRoleRequest) *NullableGlobalRoleRequest {
	return &NullableGlobalRoleRequest{value: val, isSet: true}
}

func (v NullableGlobalRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


