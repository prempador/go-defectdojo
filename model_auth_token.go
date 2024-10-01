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

// checks if the AuthToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthToken{}

// AuthToken struct for AuthToken
type AuthToken struct {
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _AuthToken AuthToken

// NewAuthToken instantiates a new AuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthToken(token string) *AuthToken {
	this := AuthToken{}
	this.Token = token
	return &this
}

// NewAuthTokenWithDefaults instantiates a new AuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenWithDefaults() *AuthToken {
	this := AuthToken{}
	return &this
}

// GetToken returns the Token field value
func (o *AuthToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AuthToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AuthToken) SetToken(v string) {
	o.Token = v
}


func (o AuthToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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
	varAuthToken := _AuthToken{}

	err = json.Unmarshal(data, &varAuthToken)

	if err != nil {
		return err
	}

	*o = AuthToken(varAuthToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthToken struct {
	value *AuthToken
	isSet bool
}

func (v NullableAuthToken) Get() *AuthToken {
	return v.value
}

func (v *NullableAuthToken) Set(val *AuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthToken(val *AuthToken) *NullableAuthToken {
	return &NullableAuthToken{value: val, isSet: true}
}

func (v NullableAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


