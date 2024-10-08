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

// checks if the AcceptedRiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcceptedRiskRequest{}

// AcceptedRiskRequest struct for AcceptedRiskRequest
type AcceptedRiskRequest struct {
	// An id of a vulnerability in a security advisory associated with this finding. Can be a Common Vulnerabilities and Exposure (CVE) or from other sources.
	VulnerabilityId string `json:"vulnerability_id"`
	// Justification for accepting findings with this vulnerability id
	Justification string `json:"justification"`
	// Name or email of person who accepts the risk
	AcceptedBy string `json:"accepted_by"`
	AdditionalProperties map[string]interface{}
}

type _AcceptedRiskRequest AcceptedRiskRequest

// NewAcceptedRiskRequest instantiates a new AcceptedRiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcceptedRiskRequest(vulnerabilityId string, justification string, acceptedBy string) *AcceptedRiskRequest {
	this := AcceptedRiskRequest{}
	this.VulnerabilityId = vulnerabilityId
	this.Justification = justification
	this.AcceptedBy = acceptedBy
	return &this
}

// NewAcceptedRiskRequestWithDefaults instantiates a new AcceptedRiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedRiskRequestWithDefaults() *AcceptedRiskRequest {
	this := AcceptedRiskRequest{}
	return &this
}

// GetVulnerabilityId returns the VulnerabilityId field value
func (o *AcceptedRiskRequest) GetVulnerabilityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VulnerabilityId
}

// GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field value
// and a boolean to check if the value has been set.
func (o *AcceptedRiskRequest) GetVulnerabilityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityId, true
}

// SetVulnerabilityId sets field value
func (o *AcceptedRiskRequest) SetVulnerabilityId(v string) {
	o.VulnerabilityId = v
}


// GetJustification returns the Justification field value
func (o *AcceptedRiskRequest) GetJustification() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value
// and a boolean to check if the value has been set.
func (o *AcceptedRiskRequest) GetJustificationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Justification, true
}

// SetJustification sets field value
func (o *AcceptedRiskRequest) SetJustification(v string) {
	o.Justification = v
}


// GetAcceptedBy returns the AcceptedBy field value
func (o *AcceptedRiskRequest) GetAcceptedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcceptedBy
}

// GetAcceptedByOk returns a tuple with the AcceptedBy field value
// and a boolean to check if the value has been set.
func (o *AcceptedRiskRequest) GetAcceptedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedBy, true
}

// SetAcceptedBy sets field value
func (o *AcceptedRiskRequest) SetAcceptedBy(v string) {
	o.AcceptedBy = v
}


func (o AcceptedRiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcceptedRiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vulnerability_id"] = o.VulnerabilityId
	toSerialize["justification"] = o.Justification
	toSerialize["accepted_by"] = o.AcceptedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcceptedRiskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vulnerability_id",
		"justification",
		"accepted_by",
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
	varAcceptedRiskRequest := _AcceptedRiskRequest{}

	err = json.Unmarshal(data, &varAcceptedRiskRequest)

	if err != nil {
		return err
	}

	*o = AcceptedRiskRequest(varAcceptedRiskRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vulnerability_id")
		delete(additionalProperties, "justification")
		delete(additionalProperties, "accepted_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcceptedRiskRequest struct {
	value *AcceptedRiskRequest
	isSet bool
}

func (v NullableAcceptedRiskRequest) Get() *AcceptedRiskRequest {
	return v.value
}

func (v *NullableAcceptedRiskRequest) Set(val *AcceptedRiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAcceptedRiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAcceptedRiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcceptedRiskRequest(val *AcceptedRiskRequest) *NullableAcceptedRiskRequest {
	return &NullableAcceptedRiskRequest{value: val, isSet: true}
}

func (v NullableAcceptedRiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcceptedRiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


