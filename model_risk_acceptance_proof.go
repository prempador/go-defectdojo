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

// checks if the RiskAcceptanceProof type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskAcceptanceProof{}

// RiskAcceptanceProof struct for RiskAcceptanceProof
type RiskAcceptanceProof struct {
	Path string `json:"path"`
	AdditionalProperties map[string]interface{}
}

type _RiskAcceptanceProof RiskAcceptanceProof

// NewRiskAcceptanceProof instantiates a new RiskAcceptanceProof object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAcceptanceProof(path string) *RiskAcceptanceProof {
	this := RiskAcceptanceProof{}
	this.Path = path
	return &this
}

// NewRiskAcceptanceProofWithDefaults instantiates a new RiskAcceptanceProof object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAcceptanceProofWithDefaults() *RiskAcceptanceProof {
	this := RiskAcceptanceProof{}
	return &this
}

// GetPath returns the Path field value
func (o *RiskAcceptanceProof) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RiskAcceptanceProof) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RiskAcceptanceProof) SetPath(v string) {
	o.Path = v
}


func (o RiskAcceptanceProof) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskAcceptanceProof) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskAcceptanceProof) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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
	varRiskAcceptanceProof := _RiskAcceptanceProof{}

	err = json.Unmarshal(data, &varRiskAcceptanceProof)

	if err != nil {
		return err
	}

	*o = RiskAcceptanceProof(varRiskAcceptanceProof)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskAcceptanceProof struct {
	value *RiskAcceptanceProof
	isSet bool
}

func (v NullableRiskAcceptanceProof) Get() *RiskAcceptanceProof {
	return v.value
}

func (v *NullableRiskAcceptanceProof) Set(val *RiskAcceptanceProof) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAcceptanceProof) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAcceptanceProof) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAcceptanceProof(val *RiskAcceptanceProof) *NullableRiskAcceptanceProof {
	return &NullableRiskAcceptanceProof{value: val, isSet: true}
}

func (v NullableRiskAcceptanceProof) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAcceptanceProof) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


