/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PaginatedProductMemberList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedProductMemberList{}

// PaginatedProductMemberList struct for PaginatedProductMemberList
type PaginatedProductMemberList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []ProductMember `json:"results,omitempty"`
	Prefetch *PaginatedProductMemberListPrefetch `json:"prefetch,omitempty"`
}

// NewPaginatedProductMemberList instantiates a new PaginatedProductMemberList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProductMemberList() *PaginatedProductMemberList {
	this := PaginatedProductMemberList{}
	return &this
}

// NewPaginatedProductMemberListWithDefaults instantiates a new PaginatedProductMemberList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProductMemberListWithDefaults() *PaginatedProductMemberList {
	this := PaginatedProductMemberList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedProductMemberList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductMemberList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedProductMemberList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedProductMemberList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedProductMemberList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedProductMemberList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedProductMemberList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedProductMemberList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedProductMemberList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedProductMemberList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedProductMemberList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedProductMemberList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedProductMemberList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedProductMemberList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedProductMemberList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedProductMemberList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedProductMemberList) GetResults() []ProductMember {
	if o == nil || IsNil(o.Results) {
		var ret []ProductMember
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductMemberList) GetResultsOk() ([]ProductMember, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedProductMemberList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProductMember and assigns it to the Results field.
func (o *PaginatedProductMemberList) SetResults(v []ProductMember) {
	o.Results = v
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *PaginatedProductMemberList) GetPrefetch() PaginatedProductMemberListPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret PaginatedProductMemberListPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductMemberList) GetPrefetchOk() (*PaginatedProductMemberListPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *PaginatedProductMemberList) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given PaginatedProductMemberListPrefetch and assigns it to the Prefetch field.
func (o *PaginatedProductMemberList) SetPrefetch(v PaginatedProductMemberListPrefetch) {
	o.Prefetch = &v
}

func (o PaginatedProductMemberList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedProductMemberList) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedProductMemberList struct {
	value *PaginatedProductMemberList
	isSet bool
}

func (v NullablePaginatedProductMemberList) Get() *PaginatedProductMemberList {
	return v.value
}

func (v *NullablePaginatedProductMemberList) Set(val *PaginatedProductMemberList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProductMemberList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProductMemberList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProductMemberList(val *PaginatedProductMemberList) *NullablePaginatedProductMemberList {
	return &NullablePaginatedProductMemberList{value: val, isSet: true}
}

func (v NullablePaginatedProductMemberList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProductMemberList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


