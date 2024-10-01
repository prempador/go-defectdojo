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

// checks if the ToolProductSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ToolProductSettings{}

// ToolProductSettings struct for ToolProductSettings
type ToolProductSettings struct {
	Id int32 `json:"id"`
	SettingUrl string `json:"setting_url"`
	Product int32 `json:"product"`
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	Url NullableString `json:"url,omitempty"`
	ToolProjectId NullableString `json:"tool_project_id,omitempty"`
	ToolConfiguration int32 `json:"tool_configuration"`
	Notes []int32 `json:"notes"`
	Prefetch *PaginatedToolProductSettingsListPrefetch `json:"prefetch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ToolProductSettings ToolProductSettings

// NewToolProductSettings instantiates a new ToolProductSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToolProductSettings(id int32, settingUrl string, product int32, name string, toolConfiguration int32, notes []int32) *ToolProductSettings {
	this := ToolProductSettings{}
	this.Id = id
	this.SettingUrl = settingUrl
	this.Product = product
	this.Name = name
	this.ToolConfiguration = toolConfiguration
	this.Notes = notes
	return &this
}

// NewToolProductSettingsWithDefaults instantiates a new ToolProductSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewToolProductSettingsWithDefaults() *ToolProductSettings {
	this := ToolProductSettings{}
	return &this
}

// GetId returns the Id field value
func (o *ToolProductSettings) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ToolProductSettings) SetId(v int32) {
	o.Id = v
}


// GetSettingUrl returns the SettingUrl field value
func (o *ToolProductSettings) GetSettingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettingUrl
}

// GetSettingUrlOk returns a tuple with the SettingUrl field value
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetSettingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettingUrl, true
}

// SetSettingUrl sets field value
func (o *ToolProductSettings) SetSettingUrl(v string) {
	o.SettingUrl = v
}


// GetProduct returns the Product field value
func (o *ToolProductSettings) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ToolProductSettings) SetProduct(v int32) {
	o.Product = v
}


// GetName returns the Name field value
func (o *ToolProductSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ToolProductSettings) SetName(v string) {
	o.Name = v
}


// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolProductSettings) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolProductSettings) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ToolProductSettings) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ToolProductSettings) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ToolProductSettings) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ToolProductSettings) UnsetDescription() {
	o.Description.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolProductSettings) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolProductSettings) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ToolProductSettings) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ToolProductSettings) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ToolProductSettings) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ToolProductSettings) UnsetUrl() {
	o.Url.Unset()
}

// GetToolProjectId returns the ToolProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ToolProductSettings) GetToolProjectId() string {
	if o == nil || IsNil(o.ToolProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.ToolProjectId.Get()
}

// GetToolProjectIdOk returns a tuple with the ToolProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ToolProductSettings) GetToolProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolProjectId.Get(), o.ToolProjectId.IsSet()
}

// HasToolProjectId returns a boolean if a field has been set.
func (o *ToolProductSettings) HasToolProjectId() bool {
	if o != nil && o.ToolProjectId.IsSet() {
		return true
	}

	return false
}

// SetToolProjectId gets a reference to the given NullableString and assigns it to the ToolProjectId field.
func (o *ToolProductSettings) SetToolProjectId(v string) {
	o.ToolProjectId.Set(&v)
}
// SetToolProjectIdNil sets the value for ToolProjectId to be an explicit nil
func (o *ToolProductSettings) SetToolProjectIdNil() {
	o.ToolProjectId.Set(nil)
}

// UnsetToolProjectId ensures that no value is present for ToolProjectId, not even an explicit nil
func (o *ToolProductSettings) UnsetToolProjectId() {
	o.ToolProjectId.Unset()
}

// GetToolConfiguration returns the ToolConfiguration field value
func (o *ToolProductSettings) GetToolConfiguration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToolConfiguration
}

// GetToolConfigurationOk returns a tuple with the ToolConfiguration field value
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetToolConfigurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolConfiguration, true
}

// SetToolConfiguration sets field value
func (o *ToolProductSettings) SetToolConfiguration(v int32) {
	o.ToolConfiguration = v
}


// GetNotes returns the Notes field value
func (o *ToolProductSettings) GetNotes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetNotesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes, true
}

// SetNotes sets field value
func (o *ToolProductSettings) SetNotes(v []int32) {
	o.Notes = v
}


// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *ToolProductSettings) GetPrefetch() PaginatedToolProductSettingsListPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret PaginatedToolProductSettingsListPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ToolProductSettings) GetPrefetchOk() (*PaginatedToolProductSettingsListPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *ToolProductSettings) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given PaginatedToolProductSettingsListPrefetch and assigns it to the Prefetch field.
func (o *ToolProductSettings) SetPrefetch(v PaginatedToolProductSettingsListPrefetch) {
	o.Prefetch = &v
}

func (o ToolProductSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ToolProductSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["setting_url"] = o.SettingUrl
	toSerialize["product"] = o.Product
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.ToolProjectId.IsSet() {
		toSerialize["tool_project_id"] = o.ToolProjectId.Get()
	}
	toSerialize["tool_configuration"] = o.ToolConfiguration
	toSerialize["notes"] = o.Notes
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ToolProductSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"setting_url",
		"product",
		"name",
		"tool_configuration",
		"notes",
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
	varToolProductSettings := _ToolProductSettings{}

	err = json.Unmarshal(data, &varToolProductSettings)

	if err != nil {
		return err
	}

	*o = ToolProductSettings(varToolProductSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "setting_url")
		delete(additionalProperties, "product")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "url")
		delete(additionalProperties, "tool_project_id")
		delete(additionalProperties, "tool_configuration")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "prefetch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableToolProductSettings struct {
	value *ToolProductSettings
	isSet bool
}

func (v NullableToolProductSettings) Get() *ToolProductSettings {
	return v.value
}

func (v *NullableToolProductSettings) Set(val *ToolProductSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableToolProductSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableToolProductSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToolProductSettings(val *ToolProductSettings) *NullableToolProductSettings {
	return &NullableToolProductSettings{value: val, isSet: true}
}

func (v NullableToolProductSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToolProductSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


