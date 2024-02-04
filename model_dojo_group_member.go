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

// checks if the DojoGroupMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DojoGroupMember{}

// DojoGroupMember struct for DojoGroupMember
type DojoGroupMember struct {
	Id int32 `json:"id"`
	Group int32 `json:"group"`
	User int32 `json:"user"`
	// This role determines the permissions of the user to manage the group.
	Role int32 `json:"role"`
	Prefetch *DojoGroupMemberPrefetch `json:"prefetch,omitempty"`
}

type _DojoGroupMember DojoGroupMember

// NewDojoGroupMember instantiates a new DojoGroupMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDojoGroupMember(id int32, group int32, user int32, role int32) *DojoGroupMember {
	this := DojoGroupMember{}
	this.Id = id
	this.Group = group
	this.User = user
	this.Role = role
	return &this
}

// NewDojoGroupMemberWithDefaults instantiates a new DojoGroupMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDojoGroupMemberWithDefaults() *DojoGroupMember {
	this := DojoGroupMember{}
	return &this
}

// GetId returns the Id field value
func (o *DojoGroupMember) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DojoGroupMember) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DojoGroupMember) SetId(v int32) {
	o.Id = v
}

// GetGroup returns the Group field value
func (o *DojoGroupMember) GetGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *DojoGroupMember) GetGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *DojoGroupMember) SetGroup(v int32) {
	o.Group = v
}

// GetUser returns the User field value
func (o *DojoGroupMember) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *DojoGroupMember) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *DojoGroupMember) SetUser(v int32) {
	o.User = v
}

// GetRole returns the Role field value
func (o *DojoGroupMember) GetRole() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *DojoGroupMember) GetRoleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *DojoGroupMember) SetRole(v int32) {
	o.Role = v
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *DojoGroupMember) GetPrefetch() DojoGroupMemberPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret DojoGroupMemberPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DojoGroupMember) GetPrefetchOk() (*DojoGroupMemberPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *DojoGroupMember) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given DojoGroupMemberPrefetch and assigns it to the Prefetch field.
func (o *DojoGroupMember) SetPrefetch(v DojoGroupMemberPrefetch) {
	o.Prefetch = &v
}

func (o DojoGroupMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DojoGroupMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["group"] = o.Group
	toSerialize["user"] = o.User
	toSerialize["role"] = o.Role
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}
	return toSerialize, nil
}

func (o *DojoGroupMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"group",
		"user",
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

	varDojoGroupMember := _DojoGroupMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDojoGroupMember)

	if err != nil {
		return err
	}

	*o = DojoGroupMember(varDojoGroupMember)

	return err
}

type NullableDojoGroupMember struct {
	value *DojoGroupMember
	isSet bool
}

func (v NullableDojoGroupMember) Get() *DojoGroupMember {
	return v.value
}

func (v *NullableDojoGroupMember) Set(val *DojoGroupMember) {
	v.value = val
	v.isSet = true
}

func (v NullableDojoGroupMember) IsSet() bool {
	return v.isSet
}

func (v *NullableDojoGroupMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDojoGroupMember(val *DojoGroupMember) *NullableDojoGroupMember {
	return &NullableDojoGroupMember{value: val, isSet: true}
}

func (v NullableDojoGroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDojoGroupMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


