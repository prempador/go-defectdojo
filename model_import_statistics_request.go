/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ImportStatisticsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportStatisticsRequest{}

// ImportStatisticsRequest struct for ImportStatisticsRequest
type ImportStatisticsRequest struct {
	Before *SeverityStatusStatisticsRequest `json:"before,omitempty"`
	Delta *DeltaStatisticsRequest `json:"delta,omitempty"`
	After SeverityStatusStatisticsRequest `json:"after"`
}

type _ImportStatisticsRequest ImportStatisticsRequest

// NewImportStatisticsRequest instantiates a new ImportStatisticsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportStatisticsRequest(after SeverityStatusStatisticsRequest) *ImportStatisticsRequest {
	this := ImportStatisticsRequest{}
	this.After = after
	return &this
}

// NewImportStatisticsRequestWithDefaults instantiates a new ImportStatisticsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportStatisticsRequestWithDefaults() *ImportStatisticsRequest {
	this := ImportStatisticsRequest{}
	return &this
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *ImportStatisticsRequest) GetBefore() SeverityStatusStatisticsRequest {
	if o == nil || IsNil(o.Before) {
		var ret SeverityStatusStatisticsRequest
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStatisticsRequest) GetBeforeOk() (*SeverityStatusStatisticsRequest, bool) {
	if o == nil || IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *ImportStatisticsRequest) HasBefore() bool {
	if o != nil && !IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given SeverityStatusStatisticsRequest and assigns it to the Before field.
func (o *ImportStatisticsRequest) SetBefore(v SeverityStatusStatisticsRequest) {
	o.Before = &v
}

// GetDelta returns the Delta field value if set, zero value otherwise.
func (o *ImportStatisticsRequest) GetDelta() DeltaStatisticsRequest {
	if o == nil || IsNil(o.Delta) {
		var ret DeltaStatisticsRequest
		return ret
	}
	return *o.Delta
}

// GetDeltaOk returns a tuple with the Delta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportStatisticsRequest) GetDeltaOk() (*DeltaStatisticsRequest, bool) {
	if o == nil || IsNil(o.Delta) {
		return nil, false
	}
	return o.Delta, true
}

// HasDelta returns a boolean if a field has been set.
func (o *ImportStatisticsRequest) HasDelta() bool {
	if o != nil && !IsNil(o.Delta) {
		return true
	}

	return false
}

// SetDelta gets a reference to the given DeltaStatisticsRequest and assigns it to the Delta field.
func (o *ImportStatisticsRequest) SetDelta(v DeltaStatisticsRequest) {
	o.Delta = &v
}

// GetAfter returns the After field value
func (o *ImportStatisticsRequest) GetAfter() SeverityStatusStatisticsRequest {
	if o == nil {
		var ret SeverityStatusStatisticsRequest
		return ret
	}

	return o.After
}

// GetAfterOk returns a tuple with the After field value
// and a boolean to check if the value has been set.
func (o *ImportStatisticsRequest) GetAfterOk() (*SeverityStatusStatisticsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.After, true
}

// SetAfter sets field value
func (o *ImportStatisticsRequest) SetAfter(v SeverityStatusStatisticsRequest) {
	o.After = v
}

func (o ImportStatisticsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportStatisticsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Before) {
		toSerialize["before"] = o.Before
	}
	if !IsNil(o.Delta) {
		toSerialize["delta"] = o.Delta
	}
	toSerialize["after"] = o.After
	return toSerialize, nil
}

func (o *ImportStatisticsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"after",
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

	varImportStatisticsRequest := _ImportStatisticsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportStatisticsRequest)

	if err != nil {
		return err
	}

	*o = ImportStatisticsRequest(varImportStatisticsRequest)

	return err
}

type NullableImportStatisticsRequest struct {
	value *ImportStatisticsRequest
	isSet bool
}

func (v NullableImportStatisticsRequest) Get() *ImportStatisticsRequest {
	return v.value
}

func (v *NullableImportStatisticsRequest) Set(val *ImportStatisticsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportStatisticsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportStatisticsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportStatisticsRequest(val *ImportStatisticsRequest) *NullableImportStatisticsRequest {
	return &NullableImportStatisticsRequest{value: val, isSet: true}
}

func (v NullableImportStatisticsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportStatisticsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

