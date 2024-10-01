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

// checks if the PaginatedProductTypeListPrefetch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedProductTypeListPrefetch{}

// PaginatedProductTypeListPrefetch struct for PaginatedProductTypeListPrefetch
type PaginatedProductTypeListPrefetch struct {
	AuthorizationGroups *map[string]DojoGroup `json:"authorization_groups,omitempty"`
	Members *map[string]UserStub `json:"members,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedProductTypeListPrefetch PaginatedProductTypeListPrefetch

// NewPaginatedProductTypeListPrefetch instantiates a new PaginatedProductTypeListPrefetch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProductTypeListPrefetch() *PaginatedProductTypeListPrefetch {
	this := PaginatedProductTypeListPrefetch{}
	return &this
}

// NewPaginatedProductTypeListPrefetchWithDefaults instantiates a new PaginatedProductTypeListPrefetch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProductTypeListPrefetchWithDefaults() *PaginatedProductTypeListPrefetch {
	this := PaginatedProductTypeListPrefetch{}
	return &this
}

// GetAuthorizationGroups returns the AuthorizationGroups field value if set, zero value otherwise.
func (o *PaginatedProductTypeListPrefetch) GetAuthorizationGroups() map[string]DojoGroup {
	if o == nil || IsNil(o.AuthorizationGroups) {
		var ret map[string]DojoGroup
		return ret
	}
	return *o.AuthorizationGroups
}

// GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductTypeListPrefetch) GetAuthorizationGroupsOk() (*map[string]DojoGroup, bool) {
	if o == nil || IsNil(o.AuthorizationGroups) {
		return nil, false
	}
	return o.AuthorizationGroups, true
}

// HasAuthorizationGroups returns a boolean if a field has been set.
func (o *PaginatedProductTypeListPrefetch) HasAuthorizationGroups() bool {
	if o != nil && !IsNil(o.AuthorizationGroups) {
		return true
	}

	return false
}

// SetAuthorizationGroups gets a reference to the given map[string]DojoGroup and assigns it to the AuthorizationGroups field.
func (o *PaginatedProductTypeListPrefetch) SetAuthorizationGroups(v map[string]DojoGroup) {
	o.AuthorizationGroups = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *PaginatedProductTypeListPrefetch) GetMembers() map[string]UserStub {
	if o == nil || IsNil(o.Members) {
		var ret map[string]UserStub
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductTypeListPrefetch) GetMembersOk() (*map[string]UserStub, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *PaginatedProductTypeListPrefetch) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given map[string]UserStub and assigns it to the Members field.
func (o *PaginatedProductTypeListPrefetch) SetMembers(v map[string]UserStub) {
	o.Members = &v
}

func (o PaginatedProductTypeListPrefetch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedProductTypeListPrefetch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationGroups) {
		toSerialize["authorization_groups"] = o.AuthorizationGroups
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedProductTypeListPrefetch) UnmarshalJSON(data []byte) (err error) {
	varPaginatedProductTypeListPrefetch := _PaginatedProductTypeListPrefetch{}

	err = json.Unmarshal(data, &varPaginatedProductTypeListPrefetch)

	if err != nil {
		return err
	}

	*o = PaginatedProductTypeListPrefetch(varPaginatedProductTypeListPrefetch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorization_groups")
		delete(additionalProperties, "members")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedProductTypeListPrefetch struct {
	value *PaginatedProductTypeListPrefetch
	isSet bool
}

func (v NullablePaginatedProductTypeListPrefetch) Get() *PaginatedProductTypeListPrefetch {
	return v.value
}

func (v *NullablePaginatedProductTypeListPrefetch) Set(val *PaginatedProductTypeListPrefetch) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProductTypeListPrefetch) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProductTypeListPrefetch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProductTypeListPrefetch(val *PaginatedProductTypeListPrefetch) *NullablePaginatedProductTypeListPrefetch {
	return &NullablePaginatedProductTypeListPrefetch{value: val, isSet: true}
}

func (v NullablePaginatedProductTypeListPrefetch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProductTypeListPrefetch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


