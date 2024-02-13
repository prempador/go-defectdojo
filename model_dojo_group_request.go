/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DojoGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DojoGroupRequest{}

// DojoGroupRequest struct for DojoGroupRequest
type DojoGroupRequest struct {
	ConfigurationPermissions []*int32 `json:"configuration_permissions,omitempty"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	// Group imported from a social provider.  * `AzureAD` - AzureAD
	SocialProvider NullableString `json:"social_provider,omitempty"`
}

type _DojoGroupRequest DojoGroupRequest

// NewDojoGroupRequest instantiates a new DojoGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDojoGroupRequest(name string) *DojoGroupRequest {
	this := DojoGroupRequest{}
	this.Name = name
	return &this
}

// NewDojoGroupRequestWithDefaults instantiates a new DojoGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDojoGroupRequestWithDefaults() *DojoGroupRequest {
	this := DojoGroupRequest{}
	return &this
}

// GetConfigurationPermissions returns the ConfigurationPermissions field value if set, zero value otherwise.
func (o *DojoGroupRequest) GetConfigurationPermissions() []*int32 {
	if o == nil || IsNil(o.ConfigurationPermissions) {
		var ret []*int32
		return ret
	}
	return o.ConfigurationPermissions
}

// GetConfigurationPermissionsOk returns a tuple with the ConfigurationPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DojoGroupRequest) GetConfigurationPermissionsOk() ([]*int32, bool) {
	if o == nil || IsNil(o.ConfigurationPermissions) {
		return nil, false
	}
	return o.ConfigurationPermissions, true
}

// HasConfigurationPermissions returns a boolean if a field has been set.
func (o *DojoGroupRequest) HasConfigurationPermissions() bool {
	if o != nil && !IsNil(o.ConfigurationPermissions) {
		return true
	}

	return false
}

// SetConfigurationPermissions gets a reference to the given []*int32 and assigns it to the ConfigurationPermissions field.
func (o *DojoGroupRequest) SetConfigurationPermissions(v []*int32) {
	o.ConfigurationPermissions = v
}

// GetName returns the Name field value
func (o *DojoGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DojoGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DojoGroupRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DojoGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DojoGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DojoGroupRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DojoGroupRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DojoGroupRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DojoGroupRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetSocialProvider returns the SocialProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DojoGroupRequest) GetSocialProvider() string {
	if o == nil || IsNil(o.SocialProvider.Get()) {
		var ret string
		return ret
	}
	return *o.SocialProvider.Get()
}

// GetSocialProviderOk returns a tuple with the SocialProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DojoGroupRequest) GetSocialProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialProvider.Get(), o.SocialProvider.IsSet()
}

// HasSocialProvider returns a boolean if a field has been set.
func (o *DojoGroupRequest) HasSocialProvider() bool {
	if o != nil && o.SocialProvider.IsSet() {
		return true
	}

	return false
}

// SetSocialProvider gets a reference to the given NullableString and assigns it to the SocialProvider field.
func (o *DojoGroupRequest) SetSocialProvider(v string) {
	o.SocialProvider.Set(&v)
}
// SetSocialProviderNil sets the value for SocialProvider to be an explicit nil
func (o *DojoGroupRequest) SetSocialProviderNil() {
	o.SocialProvider.Set(nil)
}

// UnsetSocialProvider ensures that no value is present for SocialProvider, not even an explicit nil
func (o *DojoGroupRequest) UnsetSocialProvider() {
	o.SocialProvider.Unset()
}

func (o DojoGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DojoGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigurationPermissions) {
		toSerialize["configuration_permissions"] = o.ConfigurationPermissions
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.SocialProvider.IsSet() {
		toSerialize["social_provider"] = o.SocialProvider.Get()
	}
	return toSerialize, nil
}

func (o *DojoGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varDojoGroupRequest := _DojoGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDojoGroupRequest)

	if err != nil {
		return err
	}

	*o = DojoGroupRequest(varDojoGroupRequest)

	return err
}

type NullableDojoGroupRequest struct {
	value *DojoGroupRequest
	isSet bool
}

func (v NullableDojoGroupRequest) Get() *DojoGroupRequest {
	return v.value
}

func (v *NullableDojoGroupRequest) Set(val *DojoGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDojoGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDojoGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDojoGroupRequest(val *DojoGroupRequest) *NullableDojoGroupRequest {
	return &NullableDojoGroupRequest{value: val, isSet: true}
}

func (v NullableDojoGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDojoGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


