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

// checks if the BurpRawRequestResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BurpRawRequestResponseRequest{}

// BurpRawRequestResponseRequest struct for BurpRawRequestResponseRequest
type BurpRawRequestResponseRequest struct {
	ReqResp []map[string]string `json:"req_resp"`
	AdditionalProperties map[string]interface{}
}

type _BurpRawRequestResponseRequest BurpRawRequestResponseRequest

// NewBurpRawRequestResponseRequest instantiates a new BurpRawRequestResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurpRawRequestResponseRequest(reqResp []map[string]string) *BurpRawRequestResponseRequest {
	this := BurpRawRequestResponseRequest{}
	this.ReqResp = reqResp
	return &this
}

// NewBurpRawRequestResponseRequestWithDefaults instantiates a new BurpRawRequestResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurpRawRequestResponseRequestWithDefaults() *BurpRawRequestResponseRequest {
	this := BurpRawRequestResponseRequest{}
	return &this
}

// GetReqResp returns the ReqResp field value
func (o *BurpRawRequestResponseRequest) GetReqResp() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.ReqResp
}

// GetReqRespOk returns a tuple with the ReqResp field value
// and a boolean to check if the value has been set.
func (o *BurpRawRequestResponseRequest) GetReqRespOk() ([]map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReqResp, true
}

// SetReqResp sets field value
func (o *BurpRawRequestResponseRequest) SetReqResp(v []map[string]string) {
	o.ReqResp = v
}


func (o BurpRawRequestResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurpRawRequestResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["req_resp"] = o.ReqResp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BurpRawRequestResponseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"req_resp",
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
	varBurpRawRequestResponseRequest := _BurpRawRequestResponseRequest{}

	err = json.Unmarshal(data, &varBurpRawRequestResponseRequest)

	if err != nil {
		return err
	}

	*o = BurpRawRequestResponseRequest(varBurpRawRequestResponseRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "req_resp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBurpRawRequestResponseRequest struct {
	value *BurpRawRequestResponseRequest
	isSet bool
}

func (v NullableBurpRawRequestResponseRequest) Get() *BurpRawRequestResponseRequest {
	return v.value
}

func (v *NullableBurpRawRequestResponseRequest) Set(val *BurpRawRequestResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBurpRawRequestResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBurpRawRequestResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurpRawRequestResponseRequest(val *BurpRawRequestResponseRequest) *NullableBurpRawRequestResponseRequest {
	return &NullableBurpRawRequestResponseRequest{value: val, isSet: true}
}

func (v NullableBurpRawRequestResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurpRawRequestResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


