/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the EngagementCheckListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EngagementCheckListRequest{}

// EngagementCheckListRequest struct for EngagementCheckListRequest
type EngagementCheckListRequest struct {
	SessionManagement *string `json:"session_management,omitempty"`
	EncryptionCrypto *string `json:"encryption_crypto,omitempty"`
	ConfigurationManagement *string `json:"configuration_management,omitempty"`
	Authentication *string `json:"authentication,omitempty"`
	AuthorizationAndAccessControl *string `json:"authorization_and_access_control,omitempty"`
	DataInputSanitizationValidation *string `json:"data_input_sanitization_validation,omitempty"`
	SensitiveData *string `json:"sensitive_data,omitempty"`
	Other *string `json:"other,omitempty"`
	SessionIssues []int32 `json:"session_issues,omitempty"`
	CryptoIssues []int32 `json:"crypto_issues,omitempty"`
	ConfigIssues []int32 `json:"config_issues,omitempty"`
	AuthIssues []int32 `json:"auth_issues,omitempty"`
	AuthorIssues []int32 `json:"author_issues,omitempty"`
	DataIssues []int32 `json:"data_issues,omitempty"`
	SensitiveIssues []int32 `json:"sensitive_issues,omitempty"`
	OtherIssues []int32 `json:"other_issues,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EngagementCheckListRequest EngagementCheckListRequest

// NewEngagementCheckListRequest instantiates a new EngagementCheckListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngagementCheckListRequest() *EngagementCheckListRequest {
	this := EngagementCheckListRequest{}
	return &this
}

// NewEngagementCheckListRequestWithDefaults instantiates a new EngagementCheckListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngagementCheckListRequestWithDefaults() *EngagementCheckListRequest {
	this := EngagementCheckListRequest{}
	return &this
}

// GetSessionManagement returns the SessionManagement field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetSessionManagement() string {
	if o == nil || IsNil(o.SessionManagement) {
		var ret string
		return ret
	}
	return *o.SessionManagement
}

// GetSessionManagementOk returns a tuple with the SessionManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetSessionManagementOk() (*string, bool) {
	if o == nil || IsNil(o.SessionManagement) {
		return nil, false
	}
	return o.SessionManagement, true
}

// HasSessionManagement returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasSessionManagement() bool {
	if o != nil && !IsNil(o.SessionManagement) {
		return true
	}

	return false
}

// SetSessionManagement gets a reference to the given string and assigns it to the SessionManagement field.
func (o *EngagementCheckListRequest) SetSessionManagement(v string) {
	o.SessionManagement = &v
}

// GetEncryptionCrypto returns the EncryptionCrypto field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetEncryptionCrypto() string {
	if o == nil || IsNil(o.EncryptionCrypto) {
		var ret string
		return ret
	}
	return *o.EncryptionCrypto
}

// GetEncryptionCryptoOk returns a tuple with the EncryptionCrypto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetEncryptionCryptoOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionCrypto) {
		return nil, false
	}
	return o.EncryptionCrypto, true
}

// HasEncryptionCrypto returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasEncryptionCrypto() bool {
	if o != nil && !IsNil(o.EncryptionCrypto) {
		return true
	}

	return false
}

// SetEncryptionCrypto gets a reference to the given string and assigns it to the EncryptionCrypto field.
func (o *EngagementCheckListRequest) SetEncryptionCrypto(v string) {
	o.EncryptionCrypto = &v
}

// GetConfigurationManagement returns the ConfigurationManagement field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetConfigurationManagement() string {
	if o == nil || IsNil(o.ConfigurationManagement) {
		var ret string
		return ret
	}
	return *o.ConfigurationManagement
}

// GetConfigurationManagementOk returns a tuple with the ConfigurationManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetConfigurationManagementOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationManagement) {
		return nil, false
	}
	return o.ConfigurationManagement, true
}

// HasConfigurationManagement returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasConfigurationManagement() bool {
	if o != nil && !IsNil(o.ConfigurationManagement) {
		return true
	}

	return false
}

// SetConfigurationManagement gets a reference to the given string and assigns it to the ConfigurationManagement field.
func (o *EngagementCheckListRequest) SetConfigurationManagement(v string) {
	o.ConfigurationManagement = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetAuthentication() string {
	if o == nil || IsNil(o.Authentication) {
		var ret string
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetAuthenticationOk() (*string, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given string and assigns it to the Authentication field.
func (o *EngagementCheckListRequest) SetAuthentication(v string) {
	o.Authentication = &v
}

// GetAuthorizationAndAccessControl returns the AuthorizationAndAccessControl field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetAuthorizationAndAccessControl() string {
	if o == nil || IsNil(o.AuthorizationAndAccessControl) {
		var ret string
		return ret
	}
	return *o.AuthorizationAndAccessControl
}

// GetAuthorizationAndAccessControlOk returns a tuple with the AuthorizationAndAccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetAuthorizationAndAccessControlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationAndAccessControl) {
		return nil, false
	}
	return o.AuthorizationAndAccessControl, true
}

// HasAuthorizationAndAccessControl returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasAuthorizationAndAccessControl() bool {
	if o != nil && !IsNil(o.AuthorizationAndAccessControl) {
		return true
	}

	return false
}

// SetAuthorizationAndAccessControl gets a reference to the given string and assigns it to the AuthorizationAndAccessControl field.
func (o *EngagementCheckListRequest) SetAuthorizationAndAccessControl(v string) {
	o.AuthorizationAndAccessControl = &v
}

// GetDataInputSanitizationValidation returns the DataInputSanitizationValidation field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetDataInputSanitizationValidation() string {
	if o == nil || IsNil(o.DataInputSanitizationValidation) {
		var ret string
		return ret
	}
	return *o.DataInputSanitizationValidation
}

// GetDataInputSanitizationValidationOk returns a tuple with the DataInputSanitizationValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetDataInputSanitizationValidationOk() (*string, bool) {
	if o == nil || IsNil(o.DataInputSanitizationValidation) {
		return nil, false
	}
	return o.DataInputSanitizationValidation, true
}

// HasDataInputSanitizationValidation returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasDataInputSanitizationValidation() bool {
	if o != nil && !IsNil(o.DataInputSanitizationValidation) {
		return true
	}

	return false
}

// SetDataInputSanitizationValidation gets a reference to the given string and assigns it to the DataInputSanitizationValidation field.
func (o *EngagementCheckListRequest) SetDataInputSanitizationValidation(v string) {
	o.DataInputSanitizationValidation = &v
}

// GetSensitiveData returns the SensitiveData field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetSensitiveData() string {
	if o == nil || IsNil(o.SensitiveData) {
		var ret string
		return ret
	}
	return *o.SensitiveData
}

// GetSensitiveDataOk returns a tuple with the SensitiveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetSensitiveDataOk() (*string, bool) {
	if o == nil || IsNil(o.SensitiveData) {
		return nil, false
	}
	return o.SensitiveData, true
}

// HasSensitiveData returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasSensitiveData() bool {
	if o != nil && !IsNil(o.SensitiveData) {
		return true
	}

	return false
}

// SetSensitiveData gets a reference to the given string and assigns it to the SensitiveData field.
func (o *EngagementCheckListRequest) SetSensitiveData(v string) {
	o.SensitiveData = &v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetOther() string {
	if o == nil || IsNil(o.Other) {
		var ret string
		return ret
	}
	return *o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetOtherOk() (*string, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given string and assigns it to the Other field.
func (o *EngagementCheckListRequest) SetOther(v string) {
	o.Other = &v
}

// GetSessionIssues returns the SessionIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetSessionIssues() []int32 {
	if o == nil || IsNil(o.SessionIssues) {
		var ret []int32
		return ret
	}
	return o.SessionIssues
}

// GetSessionIssuesOk returns a tuple with the SessionIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetSessionIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.SessionIssues) {
		return nil, false
	}
	return o.SessionIssues, true
}

// HasSessionIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasSessionIssues() bool {
	if o != nil && !IsNil(o.SessionIssues) {
		return true
	}

	return false
}

// SetSessionIssues gets a reference to the given []int32 and assigns it to the SessionIssues field.
func (o *EngagementCheckListRequest) SetSessionIssues(v []int32) {
	o.SessionIssues = v
}

// GetCryptoIssues returns the CryptoIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetCryptoIssues() []int32 {
	if o == nil || IsNil(o.CryptoIssues) {
		var ret []int32
		return ret
	}
	return o.CryptoIssues
}

// GetCryptoIssuesOk returns a tuple with the CryptoIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetCryptoIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.CryptoIssues) {
		return nil, false
	}
	return o.CryptoIssues, true
}

// HasCryptoIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasCryptoIssues() bool {
	if o != nil && !IsNil(o.CryptoIssues) {
		return true
	}

	return false
}

// SetCryptoIssues gets a reference to the given []int32 and assigns it to the CryptoIssues field.
func (o *EngagementCheckListRequest) SetCryptoIssues(v []int32) {
	o.CryptoIssues = v
}

// GetConfigIssues returns the ConfigIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetConfigIssues() []int32 {
	if o == nil || IsNil(o.ConfigIssues) {
		var ret []int32
		return ret
	}
	return o.ConfigIssues
}

// GetConfigIssuesOk returns a tuple with the ConfigIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetConfigIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.ConfigIssues) {
		return nil, false
	}
	return o.ConfigIssues, true
}

// HasConfigIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasConfigIssues() bool {
	if o != nil && !IsNil(o.ConfigIssues) {
		return true
	}

	return false
}

// SetConfigIssues gets a reference to the given []int32 and assigns it to the ConfigIssues field.
func (o *EngagementCheckListRequest) SetConfigIssues(v []int32) {
	o.ConfigIssues = v
}

// GetAuthIssues returns the AuthIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetAuthIssues() []int32 {
	if o == nil || IsNil(o.AuthIssues) {
		var ret []int32
		return ret
	}
	return o.AuthIssues
}

// GetAuthIssuesOk returns a tuple with the AuthIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetAuthIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.AuthIssues) {
		return nil, false
	}
	return o.AuthIssues, true
}

// HasAuthIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasAuthIssues() bool {
	if o != nil && !IsNil(o.AuthIssues) {
		return true
	}

	return false
}

// SetAuthIssues gets a reference to the given []int32 and assigns it to the AuthIssues field.
func (o *EngagementCheckListRequest) SetAuthIssues(v []int32) {
	o.AuthIssues = v
}

// GetAuthorIssues returns the AuthorIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetAuthorIssues() []int32 {
	if o == nil || IsNil(o.AuthorIssues) {
		var ret []int32
		return ret
	}
	return o.AuthorIssues
}

// GetAuthorIssuesOk returns a tuple with the AuthorIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetAuthorIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.AuthorIssues) {
		return nil, false
	}
	return o.AuthorIssues, true
}

// HasAuthorIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasAuthorIssues() bool {
	if o != nil && !IsNil(o.AuthorIssues) {
		return true
	}

	return false
}

// SetAuthorIssues gets a reference to the given []int32 and assigns it to the AuthorIssues field.
func (o *EngagementCheckListRequest) SetAuthorIssues(v []int32) {
	o.AuthorIssues = v
}

// GetDataIssues returns the DataIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetDataIssues() []int32 {
	if o == nil || IsNil(o.DataIssues) {
		var ret []int32
		return ret
	}
	return o.DataIssues
}

// GetDataIssuesOk returns a tuple with the DataIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetDataIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.DataIssues) {
		return nil, false
	}
	return o.DataIssues, true
}

// HasDataIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasDataIssues() bool {
	if o != nil && !IsNil(o.DataIssues) {
		return true
	}

	return false
}

// SetDataIssues gets a reference to the given []int32 and assigns it to the DataIssues field.
func (o *EngagementCheckListRequest) SetDataIssues(v []int32) {
	o.DataIssues = v
}

// GetSensitiveIssues returns the SensitiveIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetSensitiveIssues() []int32 {
	if o == nil || IsNil(o.SensitiveIssues) {
		var ret []int32
		return ret
	}
	return o.SensitiveIssues
}

// GetSensitiveIssuesOk returns a tuple with the SensitiveIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetSensitiveIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.SensitiveIssues) {
		return nil, false
	}
	return o.SensitiveIssues, true
}

// HasSensitiveIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasSensitiveIssues() bool {
	if o != nil && !IsNil(o.SensitiveIssues) {
		return true
	}

	return false
}

// SetSensitiveIssues gets a reference to the given []int32 and assigns it to the SensitiveIssues field.
func (o *EngagementCheckListRequest) SetSensitiveIssues(v []int32) {
	o.SensitiveIssues = v
}

// GetOtherIssues returns the OtherIssues field value if set, zero value otherwise.
func (o *EngagementCheckListRequest) GetOtherIssues() []int32 {
	if o == nil || IsNil(o.OtherIssues) {
		var ret []int32
		return ret
	}
	return o.OtherIssues
}

// GetOtherIssuesOk returns a tuple with the OtherIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementCheckListRequest) GetOtherIssuesOk() ([]int32, bool) {
	if o == nil || IsNil(o.OtherIssues) {
		return nil, false
	}
	return o.OtherIssues, true
}

// HasOtherIssues returns a boolean if a field has been set.
func (o *EngagementCheckListRequest) HasOtherIssues() bool {
	if o != nil && !IsNil(o.OtherIssues) {
		return true
	}

	return false
}

// SetOtherIssues gets a reference to the given []int32 and assigns it to the OtherIssues field.
func (o *EngagementCheckListRequest) SetOtherIssues(v []int32) {
	o.OtherIssues = v
}

func (o EngagementCheckListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EngagementCheckListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SessionManagement) {
		toSerialize["session_management"] = o.SessionManagement
	}
	if !IsNil(o.EncryptionCrypto) {
		toSerialize["encryption_crypto"] = o.EncryptionCrypto
	}
	if !IsNil(o.ConfigurationManagement) {
		toSerialize["configuration_management"] = o.ConfigurationManagement
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !IsNil(o.AuthorizationAndAccessControl) {
		toSerialize["authorization_and_access_control"] = o.AuthorizationAndAccessControl
	}
	if !IsNil(o.DataInputSanitizationValidation) {
		toSerialize["data_input_sanitization_validation"] = o.DataInputSanitizationValidation
	}
	if !IsNil(o.SensitiveData) {
		toSerialize["sensitive_data"] = o.SensitiveData
	}
	if !IsNil(o.Other) {
		toSerialize["other"] = o.Other
	}
	if !IsNil(o.SessionIssues) {
		toSerialize["session_issues"] = o.SessionIssues
	}
	if !IsNil(o.CryptoIssues) {
		toSerialize["crypto_issues"] = o.CryptoIssues
	}
	if !IsNil(o.ConfigIssues) {
		toSerialize["config_issues"] = o.ConfigIssues
	}
	if !IsNil(o.AuthIssues) {
		toSerialize["auth_issues"] = o.AuthIssues
	}
	if !IsNil(o.AuthorIssues) {
		toSerialize["author_issues"] = o.AuthorIssues
	}
	if !IsNil(o.DataIssues) {
		toSerialize["data_issues"] = o.DataIssues
	}
	if !IsNil(o.SensitiveIssues) {
		toSerialize["sensitive_issues"] = o.SensitiveIssues
	}
	if !IsNil(o.OtherIssues) {
		toSerialize["other_issues"] = o.OtherIssues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EngagementCheckListRequest) UnmarshalJSON(data []byte) (err error) {
	varEngagementCheckListRequest := _EngagementCheckListRequest{}

	err = json.Unmarshal(data, &varEngagementCheckListRequest)

	if err != nil {
		return err
	}

	*o = EngagementCheckListRequest(varEngagementCheckListRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "session_management")
		delete(additionalProperties, "encryption_crypto")
		delete(additionalProperties, "configuration_management")
		delete(additionalProperties, "authentication")
		delete(additionalProperties, "authorization_and_access_control")
		delete(additionalProperties, "data_input_sanitization_validation")
		delete(additionalProperties, "sensitive_data")
		delete(additionalProperties, "other")
		delete(additionalProperties, "session_issues")
		delete(additionalProperties, "crypto_issues")
		delete(additionalProperties, "config_issues")
		delete(additionalProperties, "auth_issues")
		delete(additionalProperties, "author_issues")
		delete(additionalProperties, "data_issues")
		delete(additionalProperties, "sensitive_issues")
		delete(additionalProperties, "other_issues")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEngagementCheckListRequest struct {
	value *EngagementCheckListRequest
	isSet bool
}

func (v NullableEngagementCheckListRequest) Get() *EngagementCheckListRequest {
	return v.value
}

func (v *NullableEngagementCheckListRequest) Set(val *EngagementCheckListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEngagementCheckListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEngagementCheckListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngagementCheckListRequest(val *EngagementCheckListRequest) *NullableEngagementCheckListRequest {
	return &NullableEngagementCheckListRequest{value: val, isSet: true}
}

func (v NullableEngagementCheckListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngagementCheckListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


