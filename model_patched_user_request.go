/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUserRequest{}

// PatchedUserRequest struct for PatchedUserRequest
type PatchedUserRequest struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email *string `json:"email,omitempty"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive *bool `json:"is_active,omitempty"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser *bool `json:"is_superuser,omitempty"`
	Password *string `json:"password,omitempty"`
	ConfigurationPermissions []*int32 `json:"configuration_permissions,omitempty"`
}

// NewPatchedUserRequest instantiates a new PatchedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUserRequest() *PatchedUserRequest {
	this := PatchedUserRequest{}
	return &this
}

// NewPatchedUserRequestWithDefaults instantiates a new PatchedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUserRequestWithDefaults() *PatchedUserRequest {
	this := PatchedUserRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PatchedUserRequest) SetUsername(v string) {
	o.Username = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PatchedUserRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PatchedUserRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PatchedUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PatchedUserRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsSuperuser returns the IsSuperuser field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetIsSuperuser() bool {
	if o == nil || IsNil(o.IsSuperuser) {
		var ret bool
		return ret
	}
	return *o.IsSuperuser
}

// GetIsSuperuserOk returns a tuple with the IsSuperuser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetIsSuperuserOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSuperuser) {
		return nil, false
	}
	return o.IsSuperuser, true
}

// HasIsSuperuser returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasIsSuperuser() bool {
	if o != nil && !IsNil(o.IsSuperuser) {
		return true
	}

	return false
}

// SetIsSuperuser gets a reference to the given bool and assigns it to the IsSuperuser field.
func (o *PatchedUserRequest) SetIsSuperuser(v bool) {
	o.IsSuperuser = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PatchedUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetConfigurationPermissions returns the ConfigurationPermissions field value if set, zero value otherwise.
func (o *PatchedUserRequest) GetConfigurationPermissions() []*int32 {
	if o == nil || IsNil(o.ConfigurationPermissions) {
		var ret []*int32
		return ret
	}
	return o.ConfigurationPermissions
}

// GetConfigurationPermissionsOk returns a tuple with the ConfigurationPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUserRequest) GetConfigurationPermissionsOk() ([]*int32, bool) {
	if o == nil || IsNil(o.ConfigurationPermissions) {
		return nil, false
	}
	return o.ConfigurationPermissions, true
}

// HasConfigurationPermissions returns a boolean if a field has been set.
func (o *PatchedUserRequest) HasConfigurationPermissions() bool {
	if o != nil && !IsNil(o.ConfigurationPermissions) {
		return true
	}

	return false
}

// SetConfigurationPermissions gets a reference to the given []*int32 and assigns it to the ConfigurationPermissions field.
func (o *PatchedUserRequest) SetConfigurationPermissions(v []*int32) {
	o.ConfigurationPermissions = v
}

func (o PatchedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsSuperuser) {
		toSerialize["is_superuser"] = o.IsSuperuser
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ConfigurationPermissions) {
		toSerialize["configuration_permissions"] = o.ConfigurationPermissions
	}
	return toSerialize, nil
}

type NullablePatchedUserRequest struct {
	value *PatchedUserRequest
	isSet bool
}

func (v NullablePatchedUserRequest) Get() *PatchedUserRequest {
	return v.value
}

func (v *NullablePatchedUserRequest) Set(val *PatchedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUserRequest(val *PatchedUserRequest) *NullablePatchedUserRequest {
	return &NullablePatchedUserRequest{value: val, isSet: true}
}

func (v NullablePatchedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


