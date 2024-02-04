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

// checks if the SonarqubeIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SonarqubeIssueRequest{}

// SonarqubeIssueRequest struct for SonarqubeIssueRequest
type SonarqubeIssueRequest struct {
	// SonarQube issue key
	Key string `json:"key"`
	// SonarQube issue status
	Status string `json:"status"`
	// SonarQube issue type
	Type string `json:"type"`
}

type _SonarqubeIssueRequest SonarqubeIssueRequest

// NewSonarqubeIssueRequest instantiates a new SonarqubeIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarqubeIssueRequest(key string, status string, type_ string) *SonarqubeIssueRequest {
	this := SonarqubeIssueRequest{}
	this.Key = key
	this.Status = status
	this.Type = type_
	return &this
}

// NewSonarqubeIssueRequestWithDefaults instantiates a new SonarqubeIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarqubeIssueRequestWithDefaults() *SonarqubeIssueRequest {
	this := SonarqubeIssueRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *SonarqubeIssueRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SonarqubeIssueRequest) SetKey(v string) {
	o.Key = v
}

// GetStatus returns the Status field value
func (o *SonarqubeIssueRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SonarqubeIssueRequest) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *SonarqubeIssueRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SonarqubeIssueRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SonarqubeIssueRequest) SetType(v string) {
	o.Type = v
}

func (o SonarqubeIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SonarqubeIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *SonarqubeIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"status",
		"type",
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

	varSonarqubeIssueRequest := _SonarqubeIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSonarqubeIssueRequest)

	if err != nil {
		return err
	}

	*o = SonarqubeIssueRequest(varSonarqubeIssueRequest)

	return err
}

type NullableSonarqubeIssueRequest struct {
	value *SonarqubeIssueRequest
	isSet bool
}

func (v NullableSonarqubeIssueRequest) Get() *SonarqubeIssueRequest {
	return v.value
}

func (v *NullableSonarqubeIssueRequest) Set(val *SonarqubeIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarqubeIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarqubeIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarqubeIssueRequest(val *SonarqubeIssueRequest) *NullableSonarqubeIssueRequest {
	return &NullableSonarqubeIssueRequest{value: val, isSet: true}
}

func (v NullableSonarqubeIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarqubeIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


