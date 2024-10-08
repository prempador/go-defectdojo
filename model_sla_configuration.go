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

// checks if the SLAConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SLAConfiguration{}

// SLAConfiguration struct for SLAConfiguration
type SLAConfiguration struct {
	Id int32 `json:"id"`
	// A unique name for the set of SLAs.
	Name string `json:"name"`
	Description NullableString `json:"description,omitempty"`
	// The number of days to remediate a critical finding.
	Critical *int32 `json:"critical,omitempty"`
	// When enabled, critical findings will be assigned an SLA expiration date based on the critical finding SLA days within this SLA configuration.
	EnforceCritical *bool `json:"enforce_critical,omitempty"`
	// The number of days to remediate a high finding.
	High *int32 `json:"high,omitempty"`
	// When enabled, high findings will be assigned an SLA expiration date based on the high finding SLA days within this SLA configuration.
	EnforceHigh *bool `json:"enforce_high,omitempty"`
	// The number of days to remediate a medium finding.
	Medium *int32 `json:"medium,omitempty"`
	// When enabled, medium findings will be assigned an SLA expiration date based on the medium finding SLA days within this SLA configuration.
	EnforceMedium *bool `json:"enforce_medium,omitempty"`
	// The number of days to remediate a low finding.
	Low *int32 `json:"low,omitempty"`
	// When enabled, low findings will be assigned an SLA expiration date based on the low finding SLA days within this SLA configuration.
	EnforceLow *bool `json:"enforce_low,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SLAConfiguration SLAConfiguration

// NewSLAConfiguration instantiates a new SLAConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLAConfiguration(id int32, name string) *SLAConfiguration {
	this := SLAConfiguration{}
	this.Id = id
	this.Name = name
	return &this
}

// NewSLAConfigurationWithDefaults instantiates a new SLAConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLAConfigurationWithDefaults() *SLAConfiguration {
	this := SLAConfiguration{}
	return &this
}

// GetId returns the Id field value
func (o *SLAConfiguration) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SLAConfiguration) SetId(v int32) {
	o.Id = v
}


// GetName returns the Name field value
func (o *SLAConfiguration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SLAConfiguration) SetName(v string) {
	o.Name = v
}


// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLAConfiguration) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SLAConfiguration) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SLAConfiguration) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SLAConfiguration) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SLAConfiguration) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SLAConfiguration) UnsetDescription() {
	o.Description.Unset()
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *SLAConfiguration) GetCritical() int32 {
	if o == nil || IsNil(o.Critical) {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetCriticalOk() (*int32, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *SLAConfiguration) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *SLAConfiguration) SetCritical(v int32) {
	o.Critical = &v
}

// GetEnforceCritical returns the EnforceCritical field value if set, zero value otherwise.
func (o *SLAConfiguration) GetEnforceCritical() bool {
	if o == nil || IsNil(o.EnforceCritical) {
		var ret bool
		return ret
	}
	return *o.EnforceCritical
}

// GetEnforceCriticalOk returns a tuple with the EnforceCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetEnforceCriticalOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceCritical) {
		return nil, false
	}
	return o.EnforceCritical, true
}

// HasEnforceCritical returns a boolean if a field has been set.
func (o *SLAConfiguration) HasEnforceCritical() bool {
	if o != nil && !IsNil(o.EnforceCritical) {
		return true
	}

	return false
}

// SetEnforceCritical gets a reference to the given bool and assigns it to the EnforceCritical field.
func (o *SLAConfiguration) SetEnforceCritical(v bool) {
	o.EnforceCritical = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *SLAConfiguration) GetHigh() int32 {
	if o == nil || IsNil(o.High) {
		var ret int32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetHighOk() (*int32, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *SLAConfiguration) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given int32 and assigns it to the High field.
func (o *SLAConfiguration) SetHigh(v int32) {
	o.High = &v
}

// GetEnforceHigh returns the EnforceHigh field value if set, zero value otherwise.
func (o *SLAConfiguration) GetEnforceHigh() bool {
	if o == nil || IsNil(o.EnforceHigh) {
		var ret bool
		return ret
	}
	return *o.EnforceHigh
}

// GetEnforceHighOk returns a tuple with the EnforceHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetEnforceHighOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceHigh) {
		return nil, false
	}
	return o.EnforceHigh, true
}

// HasEnforceHigh returns a boolean if a field has been set.
func (o *SLAConfiguration) HasEnforceHigh() bool {
	if o != nil && !IsNil(o.EnforceHigh) {
		return true
	}

	return false
}

// SetEnforceHigh gets a reference to the given bool and assigns it to the EnforceHigh field.
func (o *SLAConfiguration) SetEnforceHigh(v bool) {
	o.EnforceHigh = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *SLAConfiguration) GetMedium() int32 {
	if o == nil || IsNil(o.Medium) {
		var ret int32
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetMediumOk() (*int32, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *SLAConfiguration) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given int32 and assigns it to the Medium field.
func (o *SLAConfiguration) SetMedium(v int32) {
	o.Medium = &v
}

// GetEnforceMedium returns the EnforceMedium field value if set, zero value otherwise.
func (o *SLAConfiguration) GetEnforceMedium() bool {
	if o == nil || IsNil(o.EnforceMedium) {
		var ret bool
		return ret
	}
	return *o.EnforceMedium
}

// GetEnforceMediumOk returns a tuple with the EnforceMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetEnforceMediumOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceMedium) {
		return nil, false
	}
	return o.EnforceMedium, true
}

// HasEnforceMedium returns a boolean if a field has been set.
func (o *SLAConfiguration) HasEnforceMedium() bool {
	if o != nil && !IsNil(o.EnforceMedium) {
		return true
	}

	return false
}

// SetEnforceMedium gets a reference to the given bool and assigns it to the EnforceMedium field.
func (o *SLAConfiguration) SetEnforceMedium(v bool) {
	o.EnforceMedium = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *SLAConfiguration) GetLow() int32 {
	if o == nil || IsNil(o.Low) {
		var ret int32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetLowOk() (*int32, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *SLAConfiguration) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given int32 and assigns it to the Low field.
func (o *SLAConfiguration) SetLow(v int32) {
	o.Low = &v
}

// GetEnforceLow returns the EnforceLow field value if set, zero value otherwise.
func (o *SLAConfiguration) GetEnforceLow() bool {
	if o == nil || IsNil(o.EnforceLow) {
		var ret bool
		return ret
	}
	return *o.EnforceLow
}

// GetEnforceLowOk returns a tuple with the EnforceLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfiguration) GetEnforceLowOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceLow) {
		return nil, false
	}
	return o.EnforceLow, true
}

// HasEnforceLow returns a boolean if a field has been set.
func (o *SLAConfiguration) HasEnforceLow() bool {
	if o != nil && !IsNil(o.EnforceLow) {
		return true
	}

	return false
}

// SetEnforceLow gets a reference to the given bool and assigns it to the EnforceLow field.
func (o *SLAConfiguration) SetEnforceLow(v bool) {
	o.EnforceLow = &v
}

func (o SLAConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SLAConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	if !IsNil(o.EnforceCritical) {
		toSerialize["enforce_critical"] = o.EnforceCritical
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.EnforceHigh) {
		toSerialize["enforce_high"] = o.EnforceHigh
	}
	if !IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !IsNil(o.EnforceMedium) {
		toSerialize["enforce_medium"] = o.EnforceMedium
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.EnforceLow) {
		toSerialize["enforce_low"] = o.EnforceLow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SLAConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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
	varSLAConfiguration := _SLAConfiguration{}

	err = json.Unmarshal(data, &varSLAConfiguration)

	if err != nil {
		return err
	}

	*o = SLAConfiguration(varSLAConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "critical")
		delete(additionalProperties, "enforce_critical")
		delete(additionalProperties, "high")
		delete(additionalProperties, "enforce_high")
		delete(additionalProperties, "medium")
		delete(additionalProperties, "enforce_medium")
		delete(additionalProperties, "low")
		delete(additionalProperties, "enforce_low")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSLAConfiguration struct {
	value *SLAConfiguration
	isSet bool
}

func (v NullableSLAConfiguration) Get() *SLAConfiguration {
	return v.value
}

func (v *NullableSLAConfiguration) Set(val *SLAConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableSLAConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableSLAConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLAConfiguration(val *SLAConfiguration) *NullableSLAConfiguration {
	return &NullableSLAConfiguration{value: val, isSet: true}
}

func (v NullableSLAConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLAConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


