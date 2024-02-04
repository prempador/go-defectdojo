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

// checks if the DeltaStatisticsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeltaStatisticsRequest{}

// DeltaStatisticsRequest struct for DeltaStatisticsRequest
type DeltaStatisticsRequest struct {
	Created SeverityStatusStatisticsRequest `json:"created"`
	Closed SeverityStatusStatisticsRequest `json:"closed"`
	Reactivated SeverityStatusStatisticsRequest `json:"reactivated"`
	LeftUntouched SeverityStatusStatisticsRequest `json:"left untouched"`
}

type _DeltaStatisticsRequest DeltaStatisticsRequest

// NewDeltaStatisticsRequest instantiates a new DeltaStatisticsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeltaStatisticsRequest(created SeverityStatusStatisticsRequest, closed SeverityStatusStatisticsRequest, reactivated SeverityStatusStatisticsRequest, leftUntouched SeverityStatusStatisticsRequest) *DeltaStatisticsRequest {
	this := DeltaStatisticsRequest{}
	this.Created = created
	this.Closed = closed
	this.Reactivated = reactivated
	this.LeftUntouched = leftUntouched
	return &this
}

// NewDeltaStatisticsRequestWithDefaults instantiates a new DeltaStatisticsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeltaStatisticsRequestWithDefaults() *DeltaStatisticsRequest {
	this := DeltaStatisticsRequest{}
	return &this
}

// GetCreated returns the Created field value
func (o *DeltaStatisticsRequest) GetCreated() SeverityStatusStatisticsRequest {
	if o == nil {
		var ret SeverityStatusStatisticsRequest
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DeltaStatisticsRequest) GetCreatedOk() (*SeverityStatusStatisticsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DeltaStatisticsRequest) SetCreated(v SeverityStatusStatisticsRequest) {
	o.Created = v
}

// GetClosed returns the Closed field value
func (o *DeltaStatisticsRequest) GetClosed() SeverityStatusStatisticsRequest {
	if o == nil {
		var ret SeverityStatusStatisticsRequest
		return ret
	}

	return o.Closed
}

// GetClosedOk returns a tuple with the Closed field value
// and a boolean to check if the value has been set.
func (o *DeltaStatisticsRequest) GetClosedOk() (*SeverityStatusStatisticsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Closed, true
}

// SetClosed sets field value
func (o *DeltaStatisticsRequest) SetClosed(v SeverityStatusStatisticsRequest) {
	o.Closed = v
}

// GetReactivated returns the Reactivated field value
func (o *DeltaStatisticsRequest) GetReactivated() SeverityStatusStatisticsRequest {
	if o == nil {
		var ret SeverityStatusStatisticsRequest
		return ret
	}

	return o.Reactivated
}

// GetReactivatedOk returns a tuple with the Reactivated field value
// and a boolean to check if the value has been set.
func (o *DeltaStatisticsRequest) GetReactivatedOk() (*SeverityStatusStatisticsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reactivated, true
}

// SetReactivated sets field value
func (o *DeltaStatisticsRequest) SetReactivated(v SeverityStatusStatisticsRequest) {
	o.Reactivated = v
}

// GetLeftUntouched returns the LeftUntouched field value
func (o *DeltaStatisticsRequest) GetLeftUntouched() SeverityStatusStatisticsRequest {
	if o == nil {
		var ret SeverityStatusStatisticsRequest
		return ret
	}

	return o.LeftUntouched
}

// GetLeftUntouchedOk returns a tuple with the LeftUntouched field value
// and a boolean to check if the value has been set.
func (o *DeltaStatisticsRequest) GetLeftUntouchedOk() (*SeverityStatusStatisticsRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeftUntouched, true
}

// SetLeftUntouched sets field value
func (o *DeltaStatisticsRequest) SetLeftUntouched(v SeverityStatusStatisticsRequest) {
	o.LeftUntouched = v
}

func (o DeltaStatisticsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeltaStatisticsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["closed"] = o.Closed
	toSerialize["reactivated"] = o.Reactivated
	toSerialize["left untouched"] = o.LeftUntouched
	return toSerialize, nil
}

func (o *DeltaStatisticsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"closed",
		"reactivated",
		"left untouched",
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

	varDeltaStatisticsRequest := _DeltaStatisticsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeltaStatisticsRequest)

	if err != nil {
		return err
	}

	*o = DeltaStatisticsRequest(varDeltaStatisticsRequest)

	return err
}

type NullableDeltaStatisticsRequest struct {
	value *DeltaStatisticsRequest
	isSet bool
}

func (v NullableDeltaStatisticsRequest) Get() *DeltaStatisticsRequest {
	return v.value
}

func (v *NullableDeltaStatisticsRequest) Set(val *DeltaStatisticsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeltaStatisticsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeltaStatisticsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeltaStatisticsRequest(val *DeltaStatisticsRequest) *NullableDeltaStatisticsRequest {
	return &NullableDeltaStatisticsRequest{value: val, isSet: true}
}

func (v NullableDeltaStatisticsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeltaStatisticsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

