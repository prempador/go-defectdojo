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
	"bytes"
	"fmt"
)

// checks if the EndpointStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointStatus{}

// EndpointStatus struct for EndpointStatus
type EndpointStatus struct {
	Id int32 `json:"id"`
	Date *string `json:"date,omitempty"`
	LastModified NullableTime `json:"last_modified"`
	Mitigated *bool `json:"mitigated,omitempty"`
	MitigatedTime NullableTime `json:"mitigated_time"`
	FalsePositive *bool `json:"false_positive,omitempty"`
	OutOfScope *bool `json:"out_of_scope,omitempty"`
	RiskAccepted *bool `json:"risk_accepted,omitempty"`
	MitigatedBy NullableInt32 `json:"mitigated_by,omitempty"`
	Endpoint int32 `json:"endpoint"`
	Finding int32 `json:"finding"`
}

type _EndpointStatus EndpointStatus

// NewEndpointStatus instantiates a new EndpointStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointStatus(id int32, lastModified NullableTime, mitigatedTime NullableTime, endpoint int32, finding int32) *EndpointStatus {
	this := EndpointStatus{}
	this.Id = id
	this.LastModified = lastModified
	this.MitigatedTime = mitigatedTime
	this.Endpoint = endpoint
	this.Finding = finding
	return &this
}

// NewEndpointStatusWithDefaults instantiates a new EndpointStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointStatusWithDefaults() *EndpointStatus {
	this := EndpointStatus{}
	return &this
}

// GetId returns the Id field value
func (o *EndpointStatus) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndpointStatus) SetId(v int32) {
	o.Id = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *EndpointStatus) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *EndpointStatus) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *EndpointStatus) SetDate(v string) {
	o.Date = &v
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EndpointStatus) GetLastModified() time.Time {
	if o == nil || o.LastModified.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointStatus) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// SetLastModified sets field value
func (o *EndpointStatus) SetLastModified(v time.Time) {
	o.LastModified.Set(&v)
}

// GetMitigated returns the Mitigated field value if set, zero value otherwise.
func (o *EndpointStatus) GetMitigated() bool {
	if o == nil || IsNil(o.Mitigated) {
		var ret bool
		return ret
	}
	return *o.Mitigated
}

// GetMitigatedOk returns a tuple with the Mitigated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetMitigatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Mitigated) {
		return nil, false
	}
	return o.Mitigated, true
}

// HasMitigated returns a boolean if a field has been set.
func (o *EndpointStatus) HasMitigated() bool {
	if o != nil && !IsNil(o.Mitigated) {
		return true
	}

	return false
}

// SetMitigated gets a reference to the given bool and assigns it to the Mitigated field.
func (o *EndpointStatus) SetMitigated(v bool) {
	o.Mitigated = &v
}

// GetMitigatedTime returns the MitigatedTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EndpointStatus) GetMitigatedTime() time.Time {
	if o == nil || o.MitigatedTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.MitigatedTime.Get()
}

// GetMitigatedTimeOk returns a tuple with the MitigatedTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointStatus) GetMitigatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitigatedTime.Get(), o.MitigatedTime.IsSet()
}

// SetMitigatedTime sets field value
func (o *EndpointStatus) SetMitigatedTime(v time.Time) {
	o.MitigatedTime.Set(&v)
}

// GetFalsePositive returns the FalsePositive field value if set, zero value otherwise.
func (o *EndpointStatus) GetFalsePositive() bool {
	if o == nil || IsNil(o.FalsePositive) {
		var ret bool
		return ret
	}
	return *o.FalsePositive
}

// GetFalsePositiveOk returns a tuple with the FalsePositive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetFalsePositiveOk() (*bool, bool) {
	if o == nil || IsNil(o.FalsePositive) {
		return nil, false
	}
	return o.FalsePositive, true
}

// HasFalsePositive returns a boolean if a field has been set.
func (o *EndpointStatus) HasFalsePositive() bool {
	if o != nil && !IsNil(o.FalsePositive) {
		return true
	}

	return false
}

// SetFalsePositive gets a reference to the given bool and assigns it to the FalsePositive field.
func (o *EndpointStatus) SetFalsePositive(v bool) {
	o.FalsePositive = &v
}

// GetOutOfScope returns the OutOfScope field value if set, zero value otherwise.
func (o *EndpointStatus) GetOutOfScope() bool {
	if o == nil || IsNil(o.OutOfScope) {
		var ret bool
		return ret
	}
	return *o.OutOfScope
}

// GetOutOfScopeOk returns a tuple with the OutOfScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetOutOfScopeOk() (*bool, bool) {
	if o == nil || IsNil(o.OutOfScope) {
		return nil, false
	}
	return o.OutOfScope, true
}

// HasOutOfScope returns a boolean if a field has been set.
func (o *EndpointStatus) HasOutOfScope() bool {
	if o != nil && !IsNil(o.OutOfScope) {
		return true
	}

	return false
}

// SetOutOfScope gets a reference to the given bool and assigns it to the OutOfScope field.
func (o *EndpointStatus) SetOutOfScope(v bool) {
	o.OutOfScope = &v
}

// GetRiskAccepted returns the RiskAccepted field value if set, zero value otherwise.
func (o *EndpointStatus) GetRiskAccepted() bool {
	if o == nil || IsNil(o.RiskAccepted) {
		var ret bool
		return ret
	}
	return *o.RiskAccepted
}

// GetRiskAcceptedOk returns a tuple with the RiskAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetRiskAcceptedOk() (*bool, bool) {
	if o == nil || IsNil(o.RiskAccepted) {
		return nil, false
	}
	return o.RiskAccepted, true
}

// HasRiskAccepted returns a boolean if a field has been set.
func (o *EndpointStatus) HasRiskAccepted() bool {
	if o != nil && !IsNil(o.RiskAccepted) {
		return true
	}

	return false
}

// SetRiskAccepted gets a reference to the given bool and assigns it to the RiskAccepted field.
func (o *EndpointStatus) SetRiskAccepted(v bool) {
	o.RiskAccepted = &v
}

// GetMitigatedBy returns the MitigatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointStatus) GetMitigatedBy() int32 {
	if o == nil || IsNil(o.MitigatedBy.Get()) {
		var ret int32
		return ret
	}
	return *o.MitigatedBy.Get()
}

// GetMitigatedByOk returns a tuple with the MitigatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointStatus) GetMitigatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MitigatedBy.Get(), o.MitigatedBy.IsSet()
}

// HasMitigatedBy returns a boolean if a field has been set.
func (o *EndpointStatus) HasMitigatedBy() bool {
	if o != nil && o.MitigatedBy.IsSet() {
		return true
	}

	return false
}

// SetMitigatedBy gets a reference to the given NullableInt32 and assigns it to the MitigatedBy field.
func (o *EndpointStatus) SetMitigatedBy(v int32) {
	o.MitigatedBy.Set(&v)
}
// SetMitigatedByNil sets the value for MitigatedBy to be an explicit nil
func (o *EndpointStatus) SetMitigatedByNil() {
	o.MitigatedBy.Set(nil)
}

// UnsetMitigatedBy ensures that no value is present for MitigatedBy, not even an explicit nil
func (o *EndpointStatus) UnsetMitigatedBy() {
	o.MitigatedBy.Unset()
}

// GetEndpoint returns the Endpoint field value
func (o *EndpointStatus) GetEndpoint() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetEndpointOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *EndpointStatus) SetEndpoint(v int32) {
	o.Endpoint = v
}

// GetFinding returns the Finding field value
func (o *EndpointStatus) GetFinding() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Finding
}

// GetFindingOk returns a tuple with the Finding field value
// and a boolean to check if the value has been set.
func (o *EndpointStatus) GetFindingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finding, true
}

// SetFinding sets field value
func (o *EndpointStatus) SetFinding(v int32) {
	o.Finding = v
}

func (o EndpointStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	toSerialize["last_modified"] = o.LastModified.Get()
	if !IsNil(o.Mitigated) {
		toSerialize["mitigated"] = o.Mitigated
	}
	toSerialize["mitigated_time"] = o.MitigatedTime.Get()
	if !IsNil(o.FalsePositive) {
		toSerialize["false_positive"] = o.FalsePositive
	}
	if !IsNil(o.OutOfScope) {
		toSerialize["out_of_scope"] = o.OutOfScope
	}
	if !IsNil(o.RiskAccepted) {
		toSerialize["risk_accepted"] = o.RiskAccepted
	}
	if o.MitigatedBy.IsSet() {
		toSerialize["mitigated_by"] = o.MitigatedBy.Get()
	}
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["finding"] = o.Finding
	return toSerialize, nil
}

func (o *EndpointStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"last_modified",
		"mitigated_time",
		"endpoint",
		"finding",
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

	varEndpointStatus := _EndpointStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointStatus)

	if err != nil {
		return err
	}

	*o = EndpointStatus(varEndpointStatus)

	return err
}

type NullableEndpointStatus struct {
	value *EndpointStatus
	isSet bool
}

func (v NullableEndpointStatus) Get() *EndpointStatus {
	return v.value
}

func (v *NullableEndpointStatus) Set(val *EndpointStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointStatus(val *EndpointStatus) *NullableEndpointStatus {
	return &NullableEndpointStatus{value: val, isSet: true}
}

func (v NullableEndpointStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


