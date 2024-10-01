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

// checks if the PatchedDojoGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDojoGroupRequest{}

// PatchedDojoGroupRequest struct for PatchedDojoGroupRequest
type PatchedDojoGroupRequest struct {
	ConfigurationPermissions []*int32 `json:"configuration_permissions,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// Group imported from a social provider.  * `AzureAD` - AzureAD * `Remote` - Remote
	SocialProvider NullableString `json:"social_provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedDojoGroupRequest PatchedDojoGroupRequest

// NewPatchedDojoGroupRequest instantiates a new PatchedDojoGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDojoGroupRequest() *PatchedDojoGroupRequest {
	this := PatchedDojoGroupRequest{}
	return &this
}

// NewPatchedDojoGroupRequestWithDefaults instantiates a new PatchedDojoGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDojoGroupRequestWithDefaults() *PatchedDojoGroupRequest {
	this := PatchedDojoGroupRequest{}
	return &this
}

// GetConfigurationPermissions returns the ConfigurationPermissions field value if set, zero value otherwise.
func (o *PatchedDojoGroupRequest) GetConfigurationPermissions() []*int32 {
	if o == nil || IsNil(o.ConfigurationPermissions) {
		var ret []*int32
		return ret
	}
	return o.ConfigurationPermissions
}

// GetConfigurationPermissionsOk returns a tuple with the ConfigurationPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDojoGroupRequest) GetConfigurationPermissionsOk() ([]*int32, bool) {
	if o == nil || IsNil(o.ConfigurationPermissions) {
		return nil, false
	}
	return o.ConfigurationPermissions, true
}

// HasConfigurationPermissions returns a boolean if a field has been set.
func (o *PatchedDojoGroupRequest) HasConfigurationPermissions() bool {
	if o != nil && !IsNil(o.ConfigurationPermissions) {
		return true
	}

	return false
}

// SetConfigurationPermissions gets a reference to the given []*int32 and assigns it to the ConfigurationPermissions field.
func (o *PatchedDojoGroupRequest) SetConfigurationPermissions(v []*int32) {
	o.ConfigurationPermissions = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDojoGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDojoGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDojoGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDojoGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDojoGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDojoGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedDojoGroupRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedDojoGroupRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedDojoGroupRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedDojoGroupRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetSocialProvider returns the SocialProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDojoGroupRequest) GetSocialProvider() string {
	if o == nil || IsNil(o.SocialProvider.Get()) {
		var ret string
		return ret
	}
	return *o.SocialProvider.Get()
}

// GetSocialProviderOk returns a tuple with the SocialProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDojoGroupRequest) GetSocialProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialProvider.Get(), o.SocialProvider.IsSet()
}

// HasSocialProvider returns a boolean if a field has been set.
func (o *PatchedDojoGroupRequest) HasSocialProvider() bool {
	if o != nil && o.SocialProvider.IsSet() {
		return true
	}

	return false
}

// SetSocialProvider gets a reference to the given NullableString and assigns it to the SocialProvider field.
func (o *PatchedDojoGroupRequest) SetSocialProvider(v string) {
	o.SocialProvider.Set(&v)
}
// SetSocialProviderNil sets the value for SocialProvider to be an explicit nil
func (o *PatchedDojoGroupRequest) SetSocialProviderNil() {
	o.SocialProvider.Set(nil)
}

// UnsetSocialProvider ensures that no value is present for SocialProvider, not even an explicit nil
func (o *PatchedDojoGroupRequest) UnsetSocialProvider() {
	o.SocialProvider.Unset()
}

func (o PatchedDojoGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDojoGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigurationPermissions) {
		toSerialize["configuration_permissions"] = o.ConfigurationPermissions
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.SocialProvider.IsSet() {
		toSerialize["social_provider"] = o.SocialProvider.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedDojoGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedDojoGroupRequest := _PatchedDojoGroupRequest{}

	err = json.Unmarshal(data, &varPatchedDojoGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedDojoGroupRequest(varPatchedDojoGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration_permissions")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "social_provider")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedDojoGroupRequest struct {
	value *PatchedDojoGroupRequest
	isSet bool
}

func (v NullablePatchedDojoGroupRequest) Get() *PatchedDojoGroupRequest {
	return v.value
}

func (v *NullablePatchedDojoGroupRequest) Set(val *PatchedDojoGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDojoGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDojoGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDojoGroupRequest(val *PatchedDojoGroupRequest) *NullablePatchedDojoGroupRequest {
	return &NullablePatchedDojoGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedDojoGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDojoGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


