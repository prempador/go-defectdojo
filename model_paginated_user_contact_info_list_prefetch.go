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

// checks if the PaginatedUserContactInfoListPrefetch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedUserContactInfoListPrefetch{}

// PaginatedUserContactInfoListPrefetch struct for PaginatedUserContactInfoListPrefetch
type PaginatedUserContactInfoListPrefetch struct {
	User *map[string]UserStub `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedUserContactInfoListPrefetch PaginatedUserContactInfoListPrefetch

// NewPaginatedUserContactInfoListPrefetch instantiates a new PaginatedUserContactInfoListPrefetch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedUserContactInfoListPrefetch() *PaginatedUserContactInfoListPrefetch {
	this := PaginatedUserContactInfoListPrefetch{}
	return &this
}

// NewPaginatedUserContactInfoListPrefetchWithDefaults instantiates a new PaginatedUserContactInfoListPrefetch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedUserContactInfoListPrefetchWithDefaults() *PaginatedUserContactInfoListPrefetch {
	this := PaginatedUserContactInfoListPrefetch{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *PaginatedUserContactInfoListPrefetch) GetUser() map[string]UserStub {
	if o == nil || IsNil(o.User) {
		var ret map[string]UserStub
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedUserContactInfoListPrefetch) GetUserOk() (*map[string]UserStub, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *PaginatedUserContactInfoListPrefetch) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given map[string]UserStub and assigns it to the User field.
func (o *PaginatedUserContactInfoListPrefetch) SetUser(v map[string]UserStub) {
	o.User = &v
}

func (o PaginatedUserContactInfoListPrefetch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedUserContactInfoListPrefetch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedUserContactInfoListPrefetch) UnmarshalJSON(data []byte) (err error) {
	varPaginatedUserContactInfoListPrefetch := _PaginatedUserContactInfoListPrefetch{}

	err = json.Unmarshal(data, &varPaginatedUserContactInfoListPrefetch)

	if err != nil {
		return err
	}

	*o = PaginatedUserContactInfoListPrefetch(varPaginatedUserContactInfoListPrefetch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedUserContactInfoListPrefetch struct {
	value *PaginatedUserContactInfoListPrefetch
	isSet bool
}

func (v NullablePaginatedUserContactInfoListPrefetch) Get() *PaginatedUserContactInfoListPrefetch {
	return v.value
}

func (v *NullablePaginatedUserContactInfoListPrefetch) Set(val *PaginatedUserContactInfoListPrefetch) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedUserContactInfoListPrefetch) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedUserContactInfoListPrefetch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedUserContactInfoListPrefetch(val *PaginatedUserContactInfoListPrefetch) *NullablePaginatedUserContactInfoListPrefetch {
	return &NullablePaginatedUserContactInfoListPrefetch{value: val, isSet: true}
}

func (v NullablePaginatedUserContactInfoListPrefetch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedUserContactInfoListPrefetch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


