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

// checks if the ToolConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolConfiguration{}

// ToolConfiguration struct for ToolConfiguration
type ToolConfiguration struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Url NullableString `json:"url,omitempty"`
	// * `API` - API Key * `Password` - Username/Password * `SSH` - SSH
	AuthenticationType NullableString `json:"authentication_type,omitempty"`
	// Additional definitions that will be consumed by scanner
	Extras NullableString `json:"extras,omitempty"`
	Username NullableString `json:"username,omitempty"`
	AuthTitle NullableString `json:"auth_title,omitempty"`
	ToolType int32 `json:"tool_type"`
}

type _ToolConfiguration ToolConfiguration

// NewToolConfiguration instantiates a new ToolConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolConfiguration(id int32, name string, toolType int32) *ToolConfiguration {
	this := ToolConfiguration{}
	this.Id = id
	this.Name = name
	this.ToolType = toolType
	return &this
}

// NewToolConfigurationWithDefaults instantiates a new ToolConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolConfigurationWithDefaults() *ToolConfiguration {
	this := ToolConfiguration{}
	return &this
}

// GetId returns the Id field value
func (o *ToolConfiguration) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolConfiguration) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolConfiguration) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ToolConfiguration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolConfiguration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolConfiguration) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolConfiguration) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ToolConfiguration) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ToolConfiguration) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ToolConfiguration) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ToolConfiguration) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolConfiguration) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolConfiguration) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ToolConfiguration) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ToolConfiguration) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ToolConfiguration) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ToolConfiguration) UnsetUrl() {
	o.Url.Unset()
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolConfiguration) GetAuthenticationType() string {
	if o == nil || IsNil(o.AuthenticationType.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationType.Get()
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolConfiguration) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationType.Get(), o.AuthenticationType.IsSet()
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *ToolConfiguration) HasAuthenticationType() bool {
	if o != nil && o.AuthenticationType.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given NullableString and assigns it to the AuthenticationType field.
func (o *ToolConfiguration) SetAuthenticationType(v string) {
	o.AuthenticationType.Set(&v)
}
// SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil
func (o *ToolConfiguration) SetAuthenticationTypeNil() {
	o.AuthenticationType.Set(nil)
}

// UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
func (o *ToolConfiguration) UnsetAuthenticationType() {
	o.AuthenticationType.Unset()
}

// GetExtras returns the Extras field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolConfiguration) GetExtras() string {
	if o == nil || IsNil(o.Extras.Get()) {
		var ret string
		return ret
	}
	return *o.Extras.Get()
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolConfiguration) GetExtrasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extras.Get(), o.Extras.IsSet()
}

// HasExtras returns a boolean if a field has been set.
func (o *ToolConfiguration) HasExtras() bool {
	if o != nil && o.Extras.IsSet() {
		return true
	}

	return false
}

// SetExtras gets a reference to the given NullableString and assigns it to the Extras field.
func (o *ToolConfiguration) SetExtras(v string) {
	o.Extras.Set(&v)
}
// SetExtrasNil sets the value for Extras to be an explicit nil
func (o *ToolConfiguration) SetExtrasNil() {
	o.Extras.Set(nil)
}

// UnsetExtras ensures that no value is present for Extras, not even an explicit nil
func (o *ToolConfiguration) UnsetExtras() {
	o.Extras.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolConfiguration) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolConfiguration) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *ToolConfiguration) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *ToolConfiguration) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *ToolConfiguration) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *ToolConfiguration) UnsetUsername() {
	o.Username.Unset()
}

// GetAuthTitle returns the AuthTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolConfiguration) GetAuthTitle() string {
	if o == nil || IsNil(o.AuthTitle.Get()) {
		var ret string
		return ret
	}
	return *o.AuthTitle.Get()
}

// GetAuthTitleOk returns a tuple with the AuthTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolConfiguration) GetAuthTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthTitle.Get(), o.AuthTitle.IsSet()
}

// HasAuthTitle returns a boolean if a field has been set.
func (o *ToolConfiguration) HasAuthTitle() bool {
	if o != nil && o.AuthTitle.IsSet() {
		return true
	}

	return false
}

// SetAuthTitle gets a reference to the given NullableString and assigns it to the AuthTitle field.
func (o *ToolConfiguration) SetAuthTitle(v string) {
	o.AuthTitle.Set(&v)
}
// SetAuthTitleNil sets the value for AuthTitle to be an explicit nil
func (o *ToolConfiguration) SetAuthTitleNil() {
	o.AuthTitle.Set(nil)
}

// UnsetAuthTitle ensures that no value is present for AuthTitle, not even an explicit nil
func (o *ToolConfiguration) UnsetAuthTitle() {
	o.AuthTitle.Unset()
}

// GetToolType returns the ToolType field value
func (o *ToolConfiguration) GetToolType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToolType
}

// GetToolTypeOk returns a tuple with the ToolType field value
// and a boolean to check if the value has been set.
func (o *ToolConfiguration) GetToolTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolType, true
}

// SetToolType sets field value
func (o *ToolConfiguration) SetToolType(v int32) {
	o.ToolType = v
}

func (o ToolConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.AuthenticationType.IsSet() {
		toSerialize["authentication_type"] = o.AuthenticationType.Get()
	}
	if o.Extras.IsSet() {
		toSerialize["extras"] = o.Extras.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.AuthTitle.IsSet() {
		toSerialize["auth_title"] = o.AuthTitle.Get()
	}
	toSerialize["tool_type"] = o.ToolType
	return toSerialize, nil
}

func (o *ToolConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"tool_type",
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

	varToolConfiguration := _ToolConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varToolConfiguration)

	if err != nil {
		return err
	}

	*o = ToolConfiguration(varToolConfiguration)

	return err
}

type NullableToolConfiguration struct {
	value *ToolConfiguration
	isSet bool
}

func (v NullableToolConfiguration) Get() *ToolConfiguration {
	return v.value
}

func (v *NullableToolConfiguration) Set(val *ToolConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableToolConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableToolConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolConfiguration(val *ToolConfiguration) *NullableToolConfiguration {
	return &NullableToolConfiguration{value: val, isSet: true}
}

func (v NullableToolConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


