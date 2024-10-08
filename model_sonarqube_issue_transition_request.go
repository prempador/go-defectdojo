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

// checks if the SonarqubeIssueTransitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SonarqubeIssueTransitionRequest{}

// SonarqubeIssueTransitionRequest struct for SonarqubeIssueTransitionRequest
type SonarqubeIssueTransitionRequest struct {
	FindingStatus string `json:"finding_status"`
	SonarqubeStatus string `json:"sonarqube_status"`
	Transitions string `json:"transitions"`
	SonarqubeIssue int32 `json:"sonarqube_issue"`
	AdditionalProperties map[string]interface{}
}

type _SonarqubeIssueTransitionRequest SonarqubeIssueTransitionRequest

// NewSonarqubeIssueTransitionRequest instantiates a new SonarqubeIssueTransitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarqubeIssueTransitionRequest(findingStatus string, sonarqubeStatus string, transitions string, sonarqubeIssue int32) *SonarqubeIssueTransitionRequest {
	this := SonarqubeIssueTransitionRequest{}
	this.FindingStatus = findingStatus
	this.SonarqubeStatus = sonarqubeStatus
	this.Transitions = transitions
	this.SonarqubeIssue = sonarqubeIssue
	return &this
}

// NewSonarqubeIssueTransitionRequestWithDefaults instantiates a new SonarqubeIssueTransitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarqubeIssueTransitionRequestWithDefaults() *SonarqubeIssueTransitionRequest {
	this := SonarqubeIssueTransitionRequest{}
	return &this
}

// GetFindingStatus returns the FindingStatus field value
func (o *SonarqubeIssueTransitionRequest) GetFindingStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FindingStatus
}

// GetFindingStatusOk returns a tuple with the FindingStatus field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueTransitionRequest) GetFindingStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FindingStatus, true
}

// SetFindingStatus sets field value
func (o *SonarqubeIssueTransitionRequest) SetFindingStatus(v string) {
	o.FindingStatus = v
}


// GetSonarqubeStatus returns the SonarqubeStatus field value
func (o *SonarqubeIssueTransitionRequest) GetSonarqubeStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SonarqubeStatus
}

// GetSonarqubeStatusOk returns a tuple with the SonarqubeStatus field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueTransitionRequest) GetSonarqubeStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SonarqubeStatus, true
}

// SetSonarqubeStatus sets field value
func (o *SonarqubeIssueTransitionRequest) SetSonarqubeStatus(v string) {
	o.SonarqubeStatus = v
}


// GetTransitions returns the Transitions field value
func (o *SonarqubeIssueTransitionRequest) GetTransitions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transitions
}

// GetTransitionsOk returns a tuple with the Transitions field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueTransitionRequest) GetTransitionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transitions, true
}

// SetTransitions sets field value
func (o *SonarqubeIssueTransitionRequest) SetTransitions(v string) {
	o.Transitions = v
}


// GetSonarqubeIssue returns the SonarqubeIssue field value
func (o *SonarqubeIssueTransitionRequest) GetSonarqubeIssue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SonarqubeIssue
}

// GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueTransitionRequest) GetSonarqubeIssueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SonarqubeIssue, true
}

// SetSonarqubeIssue sets field value
func (o *SonarqubeIssueTransitionRequest) SetSonarqubeIssue(v int32) {
	o.SonarqubeIssue = v
}


func (o SonarqubeIssueTransitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SonarqubeIssueTransitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["finding_status"] = o.FindingStatus
	toSerialize["sonarqube_status"] = o.SonarqubeStatus
	toSerialize["transitions"] = o.Transitions
	toSerialize["sonarqube_issue"] = o.SonarqubeIssue

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SonarqubeIssueTransitionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"finding_status",
		"sonarqube_status",
		"transitions",
		"sonarqube_issue",
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
	varSonarqubeIssueTransitionRequest := _SonarqubeIssueTransitionRequest{}

	err = json.Unmarshal(data, &varSonarqubeIssueTransitionRequest)

	if err != nil {
		return err
	}

	*o = SonarqubeIssueTransitionRequest(varSonarqubeIssueTransitionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "finding_status")
		delete(additionalProperties, "sonarqube_status")
		delete(additionalProperties, "transitions")
		delete(additionalProperties, "sonarqube_issue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSonarqubeIssueTransitionRequest struct {
	value *SonarqubeIssueTransitionRequest
	isSet bool
}

func (v NullableSonarqubeIssueTransitionRequest) Get() *SonarqubeIssueTransitionRequest {
	return v.value
}

func (v *NullableSonarqubeIssueTransitionRequest) Set(val *SonarqubeIssueTransitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarqubeIssueTransitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarqubeIssueTransitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarqubeIssueTransitionRequest(val *SonarqubeIssueTransitionRequest) *NullableSonarqubeIssueTransitionRequest {
	return &NullableSonarqubeIssueTransitionRequest{value: val, isSet: true}
}

func (v NullableSonarqubeIssueTransitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarqubeIssueTransitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


