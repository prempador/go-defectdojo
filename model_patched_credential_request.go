/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchedCredentialRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCredentialRequest{}

// PatchedCredentialRequest struct for PatchedCredentialRequest
type PatchedCredentialRequest struct {
	Name *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Role *string `json:"role,omitempty"`
	// * `Form` - Form Authentication * `SSO` - SSO Redirect
	Authentication *string `json:"authentication,omitempty"`
	// * `Basic` - Basic * `NTLM` - NTLM
	HttpAuthentication NullableString `json:"http_authentication,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	LoginRegex NullableString `json:"login_regex,omitempty"`
	LogoutRegex NullableString `json:"logout_regex,omitempty"`
	IsValid *bool `json:"is_valid,omitempty"`
	Environment *int32 `json:"environment,omitempty"`
}

// NewPatchedCredentialRequest instantiates a new PatchedCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCredentialRequest() *PatchedCredentialRequest {
	this := PatchedCredentialRequest{}
	return &this
}

// NewPatchedCredentialRequestWithDefaults instantiates a new PatchedCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCredentialRequestWithDefaults() *PatchedCredentialRequest {
	this := PatchedCredentialRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedCredentialRequest) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedCredentialRequest) SetUsername(v string) {
	o.Username = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PatchedCredentialRequest) SetRole(v string) {
	o.Role = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetAuthentication() string {
	if o == nil || IsNil(o.Authentication) {
		var ret string
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetAuthenticationOk() (*string, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given string and assigns it to the Authentication field.
func (o *PatchedCredentialRequest) SetAuthentication(v string) {
	o.Authentication = &v
}

// GetHttpAuthentication returns the HttpAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialRequest) GetHttpAuthentication() string {
	if o == nil || IsNil(o.HttpAuthentication.Get()) {
		var ret string
		return ret
	}
	return *o.HttpAuthentication.Get()
}

// GetHttpAuthenticationOk returns a tuple with the HttpAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialRequest) GetHttpAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpAuthentication.Get(), o.HttpAuthentication.IsSet()
}

// HasHttpAuthentication returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasHttpAuthentication() bool {
	if o != nil && o.HttpAuthentication.IsSet() {
		return true
	}

	return false
}

// SetHttpAuthentication gets a reference to the given NullableString and assigns it to the HttpAuthentication field.
func (o *PatchedCredentialRequest) SetHttpAuthentication(v string) {
	o.HttpAuthentication.Set(&v)
}
// SetHttpAuthenticationNil sets the value for HttpAuthentication to be an explicit nil
func (o *PatchedCredentialRequest) SetHttpAuthenticationNil() {
	o.HttpAuthentication.Set(nil)
}

// UnsetHttpAuthentication ensures that no value is present for HttpAuthentication, not even an explicit nil
func (o *PatchedCredentialRequest) UnsetHttpAuthentication() {
	o.HttpAuthentication.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedCredentialRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedCredentialRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedCredentialRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PatchedCredentialRequest) SetUrl(v string) {
	o.Url = &v
}

// GetLoginRegex returns the LoginRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialRequest) GetLoginRegex() string {
	if o == nil || IsNil(o.LoginRegex.Get()) {
		var ret string
		return ret
	}
	return *o.LoginRegex.Get()
}

// GetLoginRegexOk returns a tuple with the LoginRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialRequest) GetLoginRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginRegex.Get(), o.LoginRegex.IsSet()
}

// HasLoginRegex returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasLoginRegex() bool {
	if o != nil && o.LoginRegex.IsSet() {
		return true
	}

	return false
}

// SetLoginRegex gets a reference to the given NullableString and assigns it to the LoginRegex field.
func (o *PatchedCredentialRequest) SetLoginRegex(v string) {
	o.LoginRegex.Set(&v)
}
// SetLoginRegexNil sets the value for LoginRegex to be an explicit nil
func (o *PatchedCredentialRequest) SetLoginRegexNil() {
	o.LoginRegex.Set(nil)
}

// UnsetLoginRegex ensures that no value is present for LoginRegex, not even an explicit nil
func (o *PatchedCredentialRequest) UnsetLoginRegex() {
	o.LoginRegex.Unset()
}

// GetLogoutRegex returns the LogoutRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialRequest) GetLogoutRegex() string {
	if o == nil || IsNil(o.LogoutRegex.Get()) {
		var ret string
		return ret
	}
	return *o.LogoutRegex.Get()
}

// GetLogoutRegexOk returns a tuple with the LogoutRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialRequest) GetLogoutRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoutRegex.Get(), o.LogoutRegex.IsSet()
}

// HasLogoutRegex returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasLogoutRegex() bool {
	if o != nil && o.LogoutRegex.IsSet() {
		return true
	}

	return false
}

// SetLogoutRegex gets a reference to the given NullableString and assigns it to the LogoutRegex field.
func (o *PatchedCredentialRequest) SetLogoutRegex(v string) {
	o.LogoutRegex.Set(&v)
}
// SetLogoutRegexNil sets the value for LogoutRegex to be an explicit nil
func (o *PatchedCredentialRequest) SetLogoutRegexNil() {
	o.LogoutRegex.Set(nil)
}

// UnsetLogoutRegex ensures that no value is present for LogoutRegex, not even an explicit nil
func (o *PatchedCredentialRequest) UnsetLogoutRegex() {
	o.LogoutRegex.Unset()
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *PatchedCredentialRequest) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *PatchedCredentialRequest) GetEnvironment() int32 {
	if o == nil || IsNil(o.Environment) {
		var ret int32
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialRequest) GetEnvironmentOk() (*int32, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PatchedCredentialRequest) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given int32 and assigns it to the Environment field.
func (o *PatchedCredentialRequest) SetEnvironment(v int32) {
	o.Environment = &v
}

func (o PatchedCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if o.HttpAuthentication.IsSet() {
		toSerialize["http_authentication"] = o.HttpAuthentication.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if o.LoginRegex.IsSet() {
		toSerialize["login_regex"] = o.LoginRegex.Get()
	}
	if o.LogoutRegex.IsSet() {
		toSerialize["logout_regex"] = o.LogoutRegex.Get()
	}
	if !IsNil(o.IsValid) {
		toSerialize["is_valid"] = o.IsValid
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullablePatchedCredentialRequest struct {
	value *PatchedCredentialRequest
	isSet bool
}

func (v NullablePatchedCredentialRequest) Get() *PatchedCredentialRequest {
	return v.value
}

func (v *NullablePatchedCredentialRequest) Set(val *PatchedCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCredentialRequest(val *PatchedCredentialRequest) *NullablePatchedCredentialRequest {
	return &NullablePatchedCredentialRequest{value: val, isSet: true}
}

func (v NullablePatchedCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

