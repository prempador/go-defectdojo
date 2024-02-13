/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedToolProductSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedToolProductSettingsRequest{}

// PatchedToolProductSettingsRequest struct for PatchedToolProductSettingsRequest
type PatchedToolProductSettingsRequest struct {
	SettingUrl *string `json:"setting_url,omitempty"`
	Product *int32 `json:"product,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Url NullableString `json:"url,omitempty"`
	ToolProjectId NullableString `json:"tool_project_id,omitempty"`
	ToolConfiguration *int32 `json:"tool_configuration,omitempty"`
}

// NewPatchedToolProductSettingsRequest instantiates a new PatchedToolProductSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedToolProductSettingsRequest() *PatchedToolProductSettingsRequest {
	this := PatchedToolProductSettingsRequest{}
	return &this
}

// NewPatchedToolProductSettingsRequestWithDefaults instantiates a new PatchedToolProductSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedToolProductSettingsRequestWithDefaults() *PatchedToolProductSettingsRequest {
	this := PatchedToolProductSettingsRequest{}
	return &this
}

// GetSettingUrl returns the SettingUrl field value if set, zero value otherwise.
func (o *PatchedToolProductSettingsRequest) GetSettingUrl() string {
	if o == nil || IsNil(o.SettingUrl) {
		var ret string
		return ret
	}
	return *o.SettingUrl
}

// GetSettingUrlOk returns a tuple with the SettingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToolProductSettingsRequest) GetSettingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SettingUrl) {
		return nil, false
	}
	return o.SettingUrl, true
}

// HasSettingUrl returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasSettingUrl() bool {
	if o != nil && !IsNil(o.SettingUrl) {
		return true
	}

	return false
}

// SetSettingUrl gets a reference to the given string and assigns it to the SettingUrl field.
func (o *PatchedToolProductSettingsRequest) SetSettingUrl(v string) {
	o.SettingUrl = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PatchedToolProductSettingsRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product) {
		var ret int32
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToolProductSettingsRequest) GetProductOk() (*int32, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given int32 and assigns it to the Product field.
func (o *PatchedToolProductSettingsRequest) SetProduct(v int32) {
	o.Product = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedToolProductSettingsRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToolProductSettingsRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedToolProductSettingsRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedToolProductSettingsRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedToolProductSettingsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedToolProductSettingsRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedToolProductSettingsRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedToolProductSettingsRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedToolProductSettingsRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedToolProductSettingsRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PatchedToolProductSettingsRequest) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PatchedToolProductSettingsRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PatchedToolProductSettingsRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetToolProjectId returns the ToolProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedToolProductSettingsRequest) GetToolProjectId() string {
	if o == nil || IsNil(o.ToolProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ToolProjectId.Get()
}

// GetToolProjectIdOk returns a tuple with the ToolProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedToolProductSettingsRequest) GetToolProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolProjectId.Get(), o.ToolProjectId.IsSet()
}

// HasToolProjectId returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasToolProjectId() bool {
	if o != nil && o.ToolProjectId.IsSet() {
		return true
	}

	return false
}

// SetToolProjectId gets a reference to the given NullableString and assigns it to the ToolProjectId field.
func (o *PatchedToolProductSettingsRequest) SetToolProjectId(v string) {
	o.ToolProjectId.Set(&v)
}
// SetToolProjectIdNil sets the value for ToolProjectId to be an explicit nil
func (o *PatchedToolProductSettingsRequest) SetToolProjectIdNil() {
	o.ToolProjectId.Set(nil)
}

// UnsetToolProjectId ensures that no value is present for ToolProjectId, not even an explicit nil
func (o *PatchedToolProductSettingsRequest) UnsetToolProjectId() {
	o.ToolProjectId.Unset()
}

// GetToolConfiguration returns the ToolConfiguration field value if set, zero value otherwise.
func (o *PatchedToolProductSettingsRequest) GetToolConfiguration() int32 {
	if o == nil || IsNil(o.ToolConfiguration) {
		var ret int32
		return ret
	}
	return *o.ToolConfiguration
}

// GetToolConfigurationOk returns a tuple with the ToolConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedToolProductSettingsRequest) GetToolConfigurationOk() (*int32, bool) {
	if o == nil || IsNil(o.ToolConfiguration) {
		return nil, false
	}
	return o.ToolConfiguration, true
}

// HasToolConfiguration returns a boolean if a field has been set.
func (o *PatchedToolProductSettingsRequest) HasToolConfiguration() bool {
	if o != nil && !IsNil(o.ToolConfiguration) {
		return true
	}

	return false
}

// SetToolConfiguration gets a reference to the given int32 and assigns it to the ToolConfiguration field.
func (o *PatchedToolProductSettingsRequest) SetToolConfiguration(v int32) {
	o.ToolConfiguration = &v
}

func (o PatchedToolProductSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedToolProductSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SettingUrl) {
		toSerialize["setting_url"] = o.SettingUrl
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.ToolProjectId.IsSet() {
		toSerialize["tool_project_id"] = o.ToolProjectId.Get()
	}
	if !IsNil(o.ToolConfiguration) {
		toSerialize["tool_configuration"] = o.ToolConfiguration
	}
	return toSerialize, nil
}

type NullablePatchedToolProductSettingsRequest struct {
	value *PatchedToolProductSettingsRequest
	isSet bool
}

func (v NullablePatchedToolProductSettingsRequest) Get() *PatchedToolProductSettingsRequest {
	return v.value
}

func (v *NullablePatchedToolProductSettingsRequest) Set(val *PatchedToolProductSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedToolProductSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedToolProductSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedToolProductSettingsRequest(val *PatchedToolProductSettingsRequest) *NullablePatchedToolProductSettingsRequest {
	return &NullablePatchedToolProductSettingsRequest{value: val, isSet: true}
}

func (v NullablePatchedToolProductSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedToolProductSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


