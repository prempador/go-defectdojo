/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the FindingClose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindingClose{}

// FindingClose struct for FindingClose
type FindingClose struct {
	IsMitigated *bool `json:"is_mitigated,omitempty"`
	Mitigated *time.Time `json:"mitigated,omitempty"`
	FalseP *bool `json:"false_p,omitempty"`
	OutOfScope *bool `json:"out_of_scope,omitempty"`
	Duplicate *bool `json:"duplicate,omitempty"`
}

// NewFindingClose instantiates a new FindingClose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindingClose() *FindingClose {
	this := FindingClose{}
	return &this
}

// NewFindingCloseWithDefaults instantiates a new FindingClose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindingCloseWithDefaults() *FindingClose {
	this := FindingClose{}
	return &this
}

// GetIsMitigated returns the IsMitigated field value if set, zero value otherwise.
func (o *FindingClose) GetIsMitigated() bool {
	if o == nil || IsNil(o.IsMitigated) {
		var ret bool
		return ret
	}
	return *o.IsMitigated
}

// GetIsMitigatedOk returns a tuple with the IsMitigated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingClose) GetIsMitigatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMitigated) {
		return nil, false
	}
	return o.IsMitigated, true
}

// HasIsMitigated returns a boolean if a field has been set.
func (o *FindingClose) HasIsMitigated() bool {
	if o != nil && !IsNil(o.IsMitigated) {
		return true
	}

	return false
}

// SetIsMitigated gets a reference to the given bool and assigns it to the IsMitigated field.
func (o *FindingClose) SetIsMitigated(v bool) {
	o.IsMitigated = &v
}

// GetMitigated returns the Mitigated field value if set, zero value otherwise.
func (o *FindingClose) GetMitigated() time.Time {
	if o == nil || IsNil(o.Mitigated) {
		var ret time.Time
		return ret
	}
	return *o.Mitigated
}

// GetMitigatedOk returns a tuple with the Mitigated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingClose) GetMitigatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Mitigated) {
		return nil, false
	}
	return o.Mitigated, true
}

// HasMitigated returns a boolean if a field has been set.
func (o *FindingClose) HasMitigated() bool {
	if o != nil && !IsNil(o.Mitigated) {
		return true
	}

	return false
}

// SetMitigated gets a reference to the given time.Time and assigns it to the Mitigated field.
func (o *FindingClose) SetMitigated(v time.Time) {
	o.Mitigated = &v
}

// GetFalseP returns the FalseP field value if set, zero value otherwise.
func (o *FindingClose) GetFalseP() bool {
	if o == nil || IsNil(o.FalseP) {
		var ret bool
		return ret
	}
	return *o.FalseP
}

// GetFalsePOk returns a tuple with the FalseP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingClose) GetFalsePOk() (*bool, bool) {
	if o == nil || IsNil(o.FalseP) {
		return nil, false
	}
	return o.FalseP, true
}

// HasFalseP returns a boolean if a field has been set.
func (o *FindingClose) HasFalseP() bool {
	if o != nil && !IsNil(o.FalseP) {
		return true
	}

	return false
}

// SetFalseP gets a reference to the given bool and assigns it to the FalseP field.
func (o *FindingClose) SetFalseP(v bool) {
	o.FalseP = &v
}

// GetOutOfScope returns the OutOfScope field value if set, zero value otherwise.
func (o *FindingClose) GetOutOfScope() bool {
	if o == nil || IsNil(o.OutOfScope) {
		var ret bool
		return ret
	}
	return *o.OutOfScope
}

// GetOutOfScopeOk returns a tuple with the OutOfScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingClose) GetOutOfScopeOk() (*bool, bool) {
	if o == nil || IsNil(o.OutOfScope) {
		return nil, false
	}
	return o.OutOfScope, true
}

// HasOutOfScope returns a boolean if a field has been set.
func (o *FindingClose) HasOutOfScope() bool {
	if o != nil && !IsNil(o.OutOfScope) {
		return true
	}

	return false
}

// SetOutOfScope gets a reference to the given bool and assigns it to the OutOfScope field.
func (o *FindingClose) SetOutOfScope(v bool) {
	o.OutOfScope = &v
}

// GetDuplicate returns the Duplicate field value if set, zero value otherwise.
func (o *FindingClose) GetDuplicate() bool {
	if o == nil || IsNil(o.Duplicate) {
		var ret bool
		return ret
	}
	return *o.Duplicate
}

// GetDuplicateOk returns a tuple with the Duplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingClose) GetDuplicateOk() (*bool, bool) {
	if o == nil || IsNil(o.Duplicate) {
		return nil, false
	}
	return o.Duplicate, true
}

// HasDuplicate returns a boolean if a field has been set.
func (o *FindingClose) HasDuplicate() bool {
	if o != nil && !IsNil(o.Duplicate) {
		return true
	}

	return false
}

// SetDuplicate gets a reference to the given bool and assigns it to the Duplicate field.
func (o *FindingClose) SetDuplicate(v bool) {
	o.Duplicate = &v
}

func (o FindingClose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindingClose) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsMitigated) {
		toSerialize["is_mitigated"] = o.IsMitigated
	}
	if !IsNil(o.Mitigated) {
		toSerialize["mitigated"] = o.Mitigated
	}
	if !IsNil(o.FalseP) {
		toSerialize["false_p"] = o.FalseP
	}
	if !IsNil(o.OutOfScope) {
		toSerialize["out_of_scope"] = o.OutOfScope
	}
	if !IsNil(o.Duplicate) {
		toSerialize["duplicate"] = o.Duplicate
	}
	return toSerialize, nil
}

type NullableFindingClose struct {
	value *FindingClose
	isSet bool
}

func (v NullableFindingClose) Get() *FindingClose {
	return v.value
}

func (v *NullableFindingClose) Set(val *FindingClose) {
	v.value = val
	v.isSet = true
}

func (v NullableFindingClose) IsSet() bool {
	return v.isSet
}

func (v *NullableFindingClose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindingClose(val *FindingClose) *NullableFindingClose {
	return &NullableFindingClose{value: val, isSet: true}
}

func (v NullableFindingClose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindingClose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


