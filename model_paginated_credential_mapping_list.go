/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginatedCredentialMappingList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCredentialMappingList{}

// PaginatedCredentialMappingList struct for PaginatedCredentialMappingList
type PaginatedCredentialMappingList struct {
	Count int32 `json:"count"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []CredentialMapping `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedCredentialMappingList PaginatedCredentialMappingList

// NewPaginatedCredentialMappingList instantiates a new PaginatedCredentialMappingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCredentialMappingList(count int32, results []CredentialMapping) *PaginatedCredentialMappingList {
	this := PaginatedCredentialMappingList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedCredentialMappingListWithDefaults instantiates a new PaginatedCredentialMappingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCredentialMappingListWithDefaults() *PaginatedCredentialMappingList {
	this := PaginatedCredentialMappingList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedCredentialMappingList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedCredentialMappingList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedCredentialMappingList) SetCount(v int32) {
	o.Count = v
}


// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedCredentialMappingList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedCredentialMappingList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedCredentialMappingList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedCredentialMappingList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedCredentialMappingList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedCredentialMappingList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedCredentialMappingList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedCredentialMappingList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedCredentialMappingList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedCredentialMappingList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedCredentialMappingList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedCredentialMappingList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedCredentialMappingList) GetResults() []CredentialMapping {
	if o == nil {
		var ret []CredentialMapping
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedCredentialMappingList) GetResultsOk() ([]CredentialMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedCredentialMappingList) SetResults(v []CredentialMapping) {
	o.Results = v
}


func (o PaginatedCredentialMappingList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedCredentialMappingList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedCredentialMappingList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"results",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varPaginatedCredentialMappingList := _PaginatedCredentialMappingList{}

	err = json.Unmarshal(data, &varPaginatedCredentialMappingList)

	if err != nil {
		return err
	}

	*o = PaginatedCredentialMappingList(varPaginatedCredentialMappingList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedCredentialMappingList struct {
	value *PaginatedCredentialMappingList
	isSet bool
}

func (v NullablePaginatedCredentialMappingList) Get() *PaginatedCredentialMappingList {
	return v.value
}

func (v *NullablePaginatedCredentialMappingList) Set(val *PaginatedCredentialMappingList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCredentialMappingList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCredentialMappingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCredentialMappingList(val *PaginatedCredentialMappingList) *NullablePaginatedCredentialMappingList {
	return &NullablePaginatedCredentialMappingList{value: val, isSet: true}
}

func (v NullablePaginatedCredentialMappingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCredentialMappingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


