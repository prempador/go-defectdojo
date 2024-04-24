/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PaginatedProductAPIScanConfigurationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedProductAPIScanConfigurationList{}

// PaginatedProductAPIScanConfigurationList struct for PaginatedProductAPIScanConfigurationList
type PaginatedProductAPIScanConfigurationList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []ProductAPIScanConfiguration `json:"results,omitempty"`
}

// NewPaginatedProductAPIScanConfigurationList instantiates a new PaginatedProductAPIScanConfigurationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedProductAPIScanConfigurationList() *PaginatedProductAPIScanConfigurationList {
	this := PaginatedProductAPIScanConfigurationList{}
	return &this
}

// NewPaginatedProductAPIScanConfigurationListWithDefaults instantiates a new PaginatedProductAPIScanConfigurationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedProductAPIScanConfigurationListWithDefaults() *PaginatedProductAPIScanConfigurationList {
	this := PaginatedProductAPIScanConfigurationList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedProductAPIScanConfigurationList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductAPIScanConfigurationList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedProductAPIScanConfigurationList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedProductAPIScanConfigurationList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedProductAPIScanConfigurationList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedProductAPIScanConfigurationList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedProductAPIScanConfigurationList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedProductAPIScanConfigurationList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedProductAPIScanConfigurationList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedProductAPIScanConfigurationList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedProductAPIScanConfigurationList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedProductAPIScanConfigurationList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedProductAPIScanConfigurationList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedProductAPIScanConfigurationList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedProductAPIScanConfigurationList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedProductAPIScanConfigurationList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedProductAPIScanConfigurationList) GetResults() []ProductAPIScanConfiguration {
	if o == nil || IsNil(o.Results) {
		var ret []ProductAPIScanConfiguration
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedProductAPIScanConfigurationList) GetResultsOk() ([]ProductAPIScanConfiguration, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedProductAPIScanConfigurationList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProductAPIScanConfiguration and assigns it to the Results field.
func (o *PaginatedProductAPIScanConfigurationList) SetResults(v []ProductAPIScanConfiguration) {
	o.Results = v
}

func (o PaginatedProductAPIScanConfigurationList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedProductAPIScanConfigurationList) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullablePaginatedProductAPIScanConfigurationList struct {
	value *PaginatedProductAPIScanConfigurationList
	isSet bool
}

func (v NullablePaginatedProductAPIScanConfigurationList) Get() *PaginatedProductAPIScanConfigurationList {
	return v.value
}

func (v *NullablePaginatedProductAPIScanConfigurationList) Set(val *PaginatedProductAPIScanConfigurationList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedProductAPIScanConfigurationList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedProductAPIScanConfigurationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedProductAPIScanConfigurationList(val *PaginatedProductAPIScanConfigurationList) *NullablePaginatedProductAPIScanConfigurationList {
	return &NullablePaginatedProductAPIScanConfigurationList{value: val, isSet: true}
}

func (v NullablePaginatedProductAPIScanConfigurationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedProductAPIScanConfigurationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


