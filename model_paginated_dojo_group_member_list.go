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

// checks if the PaginatedDojoGroupMemberList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedDojoGroupMemberList{}

// PaginatedDojoGroupMemberList struct for PaginatedDojoGroupMemberList
type PaginatedDojoGroupMemberList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []DojoGroupMember `json:"results,omitempty"`
	Prefetch *DojoGroupMemberPrefetch `json:"prefetch,omitempty"`
}

// NewPaginatedDojoGroupMemberList instantiates a new PaginatedDojoGroupMemberList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedDojoGroupMemberList() *PaginatedDojoGroupMemberList {
	this := PaginatedDojoGroupMemberList{}
	return &this
}

// NewPaginatedDojoGroupMemberListWithDefaults instantiates a new PaginatedDojoGroupMemberList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedDojoGroupMemberListWithDefaults() *PaginatedDojoGroupMemberList {
	this := PaginatedDojoGroupMemberList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedDojoGroupMemberList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDojoGroupMemberList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedDojoGroupMemberList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedDojoGroupMemberList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedDojoGroupMemberList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDojoGroupMemberList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedDojoGroupMemberList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedDojoGroupMemberList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedDojoGroupMemberList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedDojoGroupMemberList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedDojoGroupMemberList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedDojoGroupMemberList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedDojoGroupMemberList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedDojoGroupMemberList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedDojoGroupMemberList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedDojoGroupMemberList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedDojoGroupMemberList) GetResults() []DojoGroupMember {
	if o == nil || IsNil(o.Results) {
		var ret []DojoGroupMember
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDojoGroupMemberList) GetResultsOk() ([]DojoGroupMember, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedDojoGroupMemberList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DojoGroupMember and assigns it to the Results field.
func (o *PaginatedDojoGroupMemberList) SetResults(v []DojoGroupMember) {
	o.Results = v
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *PaginatedDojoGroupMemberList) GetPrefetch() DojoGroupMemberPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret DojoGroupMemberPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedDojoGroupMemberList) GetPrefetchOk() (*DojoGroupMemberPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *PaginatedDojoGroupMemberList) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given DojoGroupMemberPrefetch and assigns it to the Prefetch field.
func (o *PaginatedDojoGroupMemberList) SetPrefetch(v DojoGroupMemberPrefetch) {
	o.Prefetch = &v
}

func (o PaginatedDojoGroupMemberList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedDojoGroupMemberList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}
	return toSerialize, nil
}

type NullablePaginatedDojoGroupMemberList struct {
	value *PaginatedDojoGroupMemberList
	isSet bool
}

func (v NullablePaginatedDojoGroupMemberList) Get() *PaginatedDojoGroupMemberList {
	return v.value
}

func (v *NullablePaginatedDojoGroupMemberList) Set(val *PaginatedDojoGroupMemberList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedDojoGroupMemberList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedDojoGroupMemberList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedDojoGroupMemberList(val *PaginatedDojoGroupMemberList) *NullablePaginatedDojoGroupMemberList {
	return &NullablePaginatedDojoGroupMemberList{value: val, isSet: true}
}

func (v NullablePaginatedDojoGroupMemberList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedDojoGroupMemberList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

