/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AuthTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthTokenRequest{}

// AuthTokenRequest struct for AuthTokenRequest
type AuthTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type _AuthTokenRequest AuthTokenRequest

// NewAuthTokenRequest instantiates a new AuthTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthTokenRequest(username string, password string) *AuthTokenRequest {
	this := AuthTokenRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewAuthTokenRequestWithDefaults instantiates a new AuthTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthTokenRequestWithDefaults() *AuthTokenRequest {
	this := AuthTokenRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *AuthTokenRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthTokenRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *AuthTokenRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthTokenRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthTokenRequest) SetPassword(v string) {
	o.Password = v
}

func (o AuthTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *AuthTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
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

	varAuthTokenRequest := _AuthTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthTokenRequest)

	if err != nil {
		return err
	}

	*o = AuthTokenRequest(varAuthTokenRequest)

	return err
}

type NullableAuthTokenRequest struct {
	value *AuthTokenRequest
	isSet bool
}

func (v NullableAuthTokenRequest) Get() *AuthTokenRequest {
	return v.value
}

func (v *NullableAuthTokenRequest) Set(val *AuthTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthTokenRequest(val *AuthTokenRequest) *NullableAuthTokenRequest {
	return &NullableAuthTokenRequest{value: val, isSet: true}
}

func (v NullableAuthTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


