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

// checks if the ImportStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportStatistics{}

// ImportStatistics struct for ImportStatistics
type ImportStatistics struct {
	// Finding statistics as stored in Defect Dojo before the import
	Before *SeverityStatusStatistics `json:"before,omitempty"`
	// Finding statistics of modifications made by the reimport. Only available when TRACK_IMPORT_HISTORY hass not disabled.
	Delta *DeltaStatistics `json:"delta,omitempty"`
	// Finding statistics as stored in Defect Dojo after the import
	After SeverityStatusStatistics `json:"after"`
	AdditionalProperties map[string]interface{}
}

type _ImportStatistics ImportStatistics

// NewImportStatistics instantiates a new ImportStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportStatistics(after SeverityStatusStatistics) *ImportStatistics {
	this := ImportStatistics{}
	this.After = after
	return &this
}

// NewImportStatisticsWithDefaults instantiates a new ImportStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportStatisticsWithDefaults() *ImportStatistics {
	this := ImportStatistics{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *ImportStatistics) GetBefore() SeverityStatusStatistics {
	if o == nil || IsNil(o.Before) {
		var ret SeverityStatusStatistics
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStatistics) GetBeforeOk() (*SeverityStatusStatistics, bool) {
	if o == nil || IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *ImportStatistics) HasBefore() bool {
	if o != nil && !IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given SeverityStatusStatistics and assigns it to the Before field.
func (o *ImportStatistics) SetBefore(v SeverityStatusStatistics) {
	o.Before = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *ImportStatistics) GetDelta() DeltaStatistics {
	if o == nil || IsNil(o.Delta) {
		var ret DeltaStatistics
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStatistics) GetDeltaOk() (*DeltaStatistics, bool) {
	if o == nil || IsNil(o.Delta) {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *ImportStatistics) HasDelta() bool {
	if o != nil && !IsNil(o.Delta) {
		return true
	}

	return false
}

// SetDelta gets a reference to the given DeltaStatistics and assigns it to the Delta field.
func (o *ImportStatistics) SetDelta(v DeltaStatistics) {
	o.Delta = &v
}

// GetAfter returns the After field value
func (o *ImportStatistics) GetAfter() SeverityStatusStatistics {
	if o == nil {
		var ret SeverityStatusStatistics
		return ret
	}

	return o.After
}

// GetAfterOk returns a tuple with the After field value
// and a boolean to check if the value has been set.
func (o *ImportStatistics) GetAfterOk() (*SeverityStatusStatistics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.After, true
}

// SetAfter sets field value
func (o *ImportStatistics) SetAfter(v SeverityStatusStatistics) {
	o.After = v
}


func (o ImportStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Before) {
		toSerialize["before"] = o.Before
	}
	if !IsNil(o.Delta) {
		toSerialize["delta"] = o.Delta
	}
	toSerialize["after"] = o.After

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportStatistics) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"after",
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
	varImportStatistics := _ImportStatistics{}

	err = json.Unmarshal(data, &varImportStatistics)

	if err != nil {
		return err
	}

	*o = ImportStatistics(varImportStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "before")
		delete(additionalProperties, "delta")
		delete(additionalProperties, "after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportStatistics struct {
	value *ImportStatistics
	isSet bool
}

func (v NullableImportStatistics) Get() *ImportStatistics {
	return v.value
}

func (v *NullableImportStatistics) Set(val *ImportStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableImportStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableImportStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportStatistics(val *ImportStatistics) *NullableImportStatistics {
	return &NullableImportStatistics{value: val, isSet: true}
}

func (v NullableImportStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


