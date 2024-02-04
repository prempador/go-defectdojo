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

// checks if the Credential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credential{}

// Credential struct for Credential
type Credential struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Role string `json:"role"`
	// * `Form` - Form Authentication * `SSO` - SSO Redirect
	Authentication *string `json:"authentication,omitempty"`
	// * `Basic` - Basic * `NTLM` - NTLM
	HttpAuthentication NullableString `json:"http_authentication,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Url string `json:"url"`
	LoginRegex NullableString `json:"login_regex,omitempty"`
	LogoutRegex NullableString `json:"logout_regex,omitempty"`
	IsValid *bool `json:"is_valid,omitempty"`
	Environment int32 `json:"environment"`
	Notes []int32 `json:"notes"`
}

type _Credential Credential

// NewCredential instantiates a new Credential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredential(id int32, name string, username string, role string, url string, environment int32, notes []int32) *Credential {
	this := Credential{}
	this.Id = id
	this.Name = name
	this.Username = username
	this.Role = role
	this.Url = url
	this.Environment = environment
	this.Notes = notes
	return &this
}

// NewCredentialWithDefaults instantiates a new Credential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialWithDefaults() *Credential {
	this := Credential{}
	return &this
}

// GetId returns the Id field value
func (o *Credential) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Credential) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Credential) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Credential) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Credential) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Credential) SetName(v string) {
	o.Name = v
}

// GetUsername returns the Username field value
func (o *Credential) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Credential) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Credential) SetUsername(v string) {
	o.Username = v
}

// GetRole returns the Role field value
func (o *Credential) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Credential) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *Credential) SetRole(v string) {
	o.Role = v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *Credential) GetAuthentication() string {
	if o == nil || IsNil(o.Authentication) {
		var ret string
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetAuthenticationOk() (*string, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *Credential) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given string and assigns it to the Authentication field.
func (o *Credential) SetAuthentication(v string) {
	o.Authentication = &v
}

// GetHttpAuthentication returns the HttpAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetHttpAuthentication() string {
	if o == nil || IsNil(o.HttpAuthentication.Get()) {
		var ret string
		return ret
	}
	return *o.HttpAuthentication.Get()
}

// GetHttpAuthenticationOk returns a tuple with the HttpAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetHttpAuthenticationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpAuthentication.Get(), o.HttpAuthentication.IsSet()
}

// HasHttpAuthentication returns a boolean if a field has been set.
func (o *Credential) HasHttpAuthentication() bool {
	if o != nil && o.HttpAuthentication.IsSet() {
		return true
	}

	return false
}

// SetHttpAuthentication gets a reference to the given NullableString and assigns it to the HttpAuthentication field.
func (o *Credential) SetHttpAuthentication(v string) {
	o.HttpAuthentication.Set(&v)
}
// SetHttpAuthenticationNil sets the value for HttpAuthentication to be an explicit nil
func (o *Credential) SetHttpAuthenticationNil() {
	o.HttpAuthentication.Set(nil)
}

// UnsetHttpAuthentication ensures that no value is present for HttpAuthentication, not even an explicit nil
func (o *Credential) UnsetHttpAuthentication() {
	o.HttpAuthentication.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Credential) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Credential) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Credential) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Credential) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value
func (o *Credential) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Credential) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Credential) SetUrl(v string) {
	o.Url = v
}

// GetLoginRegex returns the LoginRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetLoginRegex() string {
	if o == nil || IsNil(o.LoginRegex.Get()) {
		var ret string
		return ret
	}
	return *o.LoginRegex.Get()
}

// GetLoginRegexOk returns a tuple with the LoginRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetLoginRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginRegex.Get(), o.LoginRegex.IsSet()
}

// HasLoginRegex returns a boolean if a field has been set.
func (o *Credential) HasLoginRegex() bool {
	if o != nil && o.LoginRegex.IsSet() {
		return true
	}

	return false
}

// SetLoginRegex gets a reference to the given NullableString and assigns it to the LoginRegex field.
func (o *Credential) SetLoginRegex(v string) {
	o.LoginRegex.Set(&v)
}
// SetLoginRegexNil sets the value for LoginRegex to be an explicit nil
func (o *Credential) SetLoginRegexNil() {
	o.LoginRegex.Set(nil)
}

// UnsetLoginRegex ensures that no value is present for LoginRegex, not even an explicit nil
func (o *Credential) UnsetLoginRegex() {
	o.LoginRegex.Unset()
}

// GetLogoutRegex returns the LogoutRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credential) GetLogoutRegex() string {
	if o == nil || IsNil(o.LogoutRegex.Get()) {
		var ret string
		return ret
	}
	return *o.LogoutRegex.Get()
}

// GetLogoutRegexOk returns a tuple with the LogoutRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credential) GetLogoutRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoutRegex.Get(), o.LogoutRegex.IsSet()
}

// HasLogoutRegex returns a boolean if a field has been set.
func (o *Credential) HasLogoutRegex() bool {
	if o != nil && o.LogoutRegex.IsSet() {
		return true
	}

	return false
}

// SetLogoutRegex gets a reference to the given NullableString and assigns it to the LogoutRegex field.
func (o *Credential) SetLogoutRegex(v string) {
	o.LogoutRegex.Set(&v)
}
// SetLogoutRegexNil sets the value for LogoutRegex to be an explicit nil
func (o *Credential) SetLogoutRegexNil() {
	o.LogoutRegex.Set(nil)
}

// UnsetLogoutRegex ensures that no value is present for LogoutRegex, not even an explicit nil
func (o *Credential) UnsetLogoutRegex() {
	o.LogoutRegex.Unset()
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *Credential) GetIsValid() bool {
	if o == nil || IsNil(o.IsValid) {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credential) GetIsValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValid) {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *Credential) HasIsValid() bool {
	if o != nil && !IsNil(o.IsValid) {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *Credential) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetEnvironment returns the Environment field value
func (o *Credential) GetEnvironment() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Credential) GetEnvironmentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *Credential) SetEnvironment(v int32) {
	o.Environment = v
}

// GetNotes returns the Notes field value
func (o *Credential) GetNotes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *Credential) GetNotesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *Credential) SetNotes(v []int32) {
	o.Notes = v
}

func (o Credential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["username"] = o.Username
	toSerialize["role"] = o.Role
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if o.HttpAuthentication.IsSet() {
		toSerialize["http_authentication"] = o.HttpAuthentication.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["url"] = o.Url
	if o.LoginRegex.IsSet() {
		toSerialize["login_regex"] = o.LoginRegex.Get()
	}
	if o.LogoutRegex.IsSet() {
		toSerialize["logout_regex"] = o.LogoutRegex.Get()
	}
	if !IsNil(o.IsValid) {
		toSerialize["is_valid"] = o.IsValid
	}
	toSerialize["environment"] = o.Environment
	toSerialize["notes"] = o.Notes
	return toSerialize, nil
}

func (o *Credential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"username",
		"role",
		"url",
		"environment",
		"notes",
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

	varCredential := _Credential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCredential)

	if err != nil {
		return err
	}

	*o = Credential(varCredential)

	return err
}

type NullableCredential struct {
	value *Credential
	isSet bool
}

func (v NullableCredential) Get() *Credential {
	return v.value
}

func (v *NullableCredential) Set(val *Credential) {
	v.value = val
	v.isSet = true
}

func (v NullableCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredential(val *Credential) *NullableCredential {
	return &NullableCredential{value: val, isSet: true}
}

func (v NullableCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


