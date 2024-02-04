/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfile{}

// UserProfile struct for UserProfile
type UserProfile struct {
	User User `json:"user"`
	UserContactInfo UserContactInfo `json:"user_contact_info"`
	GlobalRole GlobalRole `json:"global_role"`
	DojoGroupMember []DojoGroupMember `json:"dojo_group_member"`
	ProductTypeMember []ProductTypeMember `json:"product_type_member"`
	ProductMember []ProductMember `json:"product_member"`
}

type _UserProfile UserProfile

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile(user User, userContactInfo UserContactInfo, globalRole GlobalRole, dojoGroupMember []DojoGroupMember, productTypeMember []ProductTypeMember, productMember []ProductMember) *UserProfile {
	this := UserProfile{}
	this.User = user
	this.UserContactInfo = userContactInfo
	this.GlobalRole = globalRole
	this.DojoGroupMember = dojoGroupMember
	this.ProductTypeMember = productTypeMember
	this.ProductMember = productMember
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetUser returns the User field value
func (o *UserProfile) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserProfile) SetUser(v User) {
	o.User = v
}

// GetUserContactInfo returns the UserContactInfo field value
func (o *UserProfile) GetUserContactInfo() UserContactInfo {
	if o == nil {
		var ret UserContactInfo
		return ret
	}

	return o.UserContactInfo
}

// GetUserContactInfoOk returns a tuple with the UserContactInfo field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserContactInfoOk() (*UserContactInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserContactInfo, true
}

// SetUserContactInfo sets field value
func (o *UserProfile) SetUserContactInfo(v UserContactInfo) {
	o.UserContactInfo = v
}

// GetGlobalRole returns the GlobalRole field value
func (o *UserProfile) GetGlobalRole() GlobalRole {
	if o == nil {
		var ret GlobalRole
		return ret
	}

	return o.GlobalRole
}

// GetGlobalRoleOk returns a tuple with the GlobalRole field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetGlobalRoleOk() (*GlobalRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalRole, true
}

// SetGlobalRole sets field value
func (o *UserProfile) SetGlobalRole(v GlobalRole) {
	o.GlobalRole = v
}

// GetDojoGroupMember returns the DojoGroupMember field value
func (o *UserProfile) GetDojoGroupMember() []DojoGroupMember {
	if o == nil {
		var ret []DojoGroupMember
		return ret
	}

	return o.DojoGroupMember
}

// GetDojoGroupMemberOk returns a tuple with the DojoGroupMember field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDojoGroupMemberOk() ([]DojoGroupMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.DojoGroupMember, true
}

// SetDojoGroupMember sets field value
func (o *UserProfile) SetDojoGroupMember(v []DojoGroupMember) {
	o.DojoGroupMember = v
}

// GetProductTypeMember returns the ProductTypeMember field value
func (o *UserProfile) GetProductTypeMember() []ProductTypeMember {
	if o == nil {
		var ret []ProductTypeMember
		return ret
	}

	return o.ProductTypeMember
}

// GetProductTypeMemberOk returns a tuple with the ProductTypeMember field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetProductTypeMemberOk() ([]ProductTypeMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductTypeMember, true
}

// SetProductTypeMember sets field value
func (o *UserProfile) SetProductTypeMember(v []ProductTypeMember) {
	o.ProductTypeMember = v
}

// GetProductMember returns the ProductMember field value
func (o *UserProfile) GetProductMember() []ProductMember {
	if o == nil {
		var ret []ProductMember
		return ret
	}

	return o.ProductMember
}

// GetProductMemberOk returns a tuple with the ProductMember field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetProductMemberOk() ([]ProductMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductMember, true
}

// SetProductMember sets field value
func (o *UserProfile) SetProductMember(v []ProductMember) {
	o.ProductMember = v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["user_contact_info"] = o.UserContactInfo
	toSerialize["global_role"] = o.GlobalRole
	toSerialize["dojo_group_member"] = o.DojoGroupMember
	toSerialize["product_type_member"] = o.ProductTypeMember
	toSerialize["product_member"] = o.ProductMember
	return toSerialize, nil
}

func (o *UserProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"user_contact_info",
		"global_role",
		"dojo_group_member",
		"product_type_member",
		"product_member",
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

	varUserProfile := _UserProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserProfile)

	if err != nil {
		return err
	}

	*o = UserProfile(varUserProfile)

	return err
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


