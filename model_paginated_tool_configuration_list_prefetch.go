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

// checks if the PaginatedToolConfigurationListPrefetch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedToolConfigurationListPrefetch{}

// PaginatedToolConfigurationListPrefetch struct for PaginatedToolConfigurationListPrefetch
type PaginatedToolConfigurationListPrefetch struct {
	ToolType *map[string]ToolType `json:"tool_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedToolConfigurationListPrefetch PaginatedToolConfigurationListPrefetch

// NewPaginatedToolConfigurationListPrefetch instantiates a new PaginatedToolConfigurationListPrefetch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedToolConfigurationListPrefetch() *PaginatedToolConfigurationListPrefetch {
	this := PaginatedToolConfigurationListPrefetch{}
	return &this
}

// NewPaginatedToolConfigurationListPrefetchWithDefaults instantiates a new PaginatedToolConfigurationListPrefetch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedToolConfigurationListPrefetchWithDefaults() *PaginatedToolConfigurationListPrefetch {
	this := PaginatedToolConfigurationListPrefetch{}
	return &this
}

// GetToolType returns the ToolType field value if set, zero value otherwise.
func (o *PaginatedToolConfigurationListPrefetch) GetToolType() map[string]ToolType {
	if o == nil || IsNil(o.ToolType) {
		var ret map[string]ToolType
		return ret
	}
	return *o.ToolType
}

// GetToolTypeOk returns a tuple with the ToolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedToolConfigurationListPrefetch) GetToolTypeOk() (*map[string]ToolType, bool) {
	if o == nil || IsNil(o.ToolType) {
		return nil, false
	}
	return o.ToolType, true
}

// HasToolType returns a boolean if a field has been set.
func (o *PaginatedToolConfigurationListPrefetch) HasToolType() bool {
	if o != nil && !IsNil(o.ToolType) {
		return true
	}

	return false
}

// SetToolType gets a reference to the given map[string]ToolType and assigns it to the ToolType field.
func (o *PaginatedToolConfigurationListPrefetch) SetToolType(v map[string]ToolType) {
	o.ToolType = &v
}

func (o PaginatedToolConfigurationListPrefetch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedToolConfigurationListPrefetch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToolType) {
		toSerialize["tool_type"] = o.ToolType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedToolConfigurationListPrefetch) UnmarshalJSON(data []byte) (err error) {
	varPaginatedToolConfigurationListPrefetch := _PaginatedToolConfigurationListPrefetch{}

	err = json.Unmarshal(data, &varPaginatedToolConfigurationListPrefetch)

	if err != nil {
		return err
	}

	*o = PaginatedToolConfigurationListPrefetch(varPaginatedToolConfigurationListPrefetch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tool_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedToolConfigurationListPrefetch struct {
	value *PaginatedToolConfigurationListPrefetch
	isSet bool
}

func (v NullablePaginatedToolConfigurationListPrefetch) Get() *PaginatedToolConfigurationListPrefetch {
	return v.value
}

func (v *NullablePaginatedToolConfigurationListPrefetch) Set(val *PaginatedToolConfigurationListPrefetch) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedToolConfigurationListPrefetch) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedToolConfigurationListPrefetch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedToolConfigurationListPrefetch(val *PaginatedToolConfigurationListPrefetch) *NullablePaginatedToolConfigurationListPrefetch {
	return &NullablePaginatedToolConfigurationListPrefetch{value: val, isSet: true}
}

func (v NullablePaginatedToolConfigurationListPrefetch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedToolConfigurationListPrefetch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


