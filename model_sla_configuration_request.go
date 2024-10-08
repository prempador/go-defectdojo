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

// checks if the SLAConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SLAConfigurationRequest{}

// SLAConfigurationRequest struct for SLAConfigurationRequest
type SLAConfigurationRequest struct {
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

type _SLAConfigurationRequest SLAConfigurationRequest

// NewSLAConfigurationRequest instantiates a new SLAConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLAConfigurationRequest(name string) *SLAConfigurationRequest {
	this := SLAConfigurationRequest{}
	this.Name = name
	return &this
}

// NewSLAConfigurationRequestWithDefaults instantiates a new SLAConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLAConfigurationRequestWithDefaults() *SLAConfigurationRequest {
	this := SLAConfigurationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SLAConfigurationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SLAConfigurationRequest) SetName(v string) {
	o.Name = v
}


// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLAConfigurationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SLAConfigurationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SLAConfigurationRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SLAConfigurationRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SLAConfigurationRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetCritical() int32 {
	if o == nil || IsNil(o.Critical) {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetCriticalOk() (*int32, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *SLAConfigurationRequest) SetCritical(v int32) {
	o.Critical = &v
}

// GetEnforceCritical returns the EnforceCritical field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetEnforceCritical() bool {
	if o == nil || IsNil(o.EnforceCritical) {
		var ret bool
		return ret
	}
	return *o.EnforceCritical
}

// GetEnforceCriticalOk returns a tuple with the EnforceCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetEnforceCriticalOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceCritical) {
		return nil, false
	}
	return o.EnforceCritical, true
}

// HasEnforceCritical returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasEnforceCritical() bool {
	if o != nil && !IsNil(o.EnforceCritical) {
		return true
	}

	return false
}

// SetEnforceCritical gets a reference to the given bool and assigns it to the EnforceCritical field.
func (o *SLAConfigurationRequest) SetEnforceCritical(v bool) {
	o.EnforceCritical = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetHigh() int32 {
	if o == nil || IsNil(o.High) {
		var ret int32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetHighOk() (*int32, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given int32 and assigns it to the High field.
func (o *SLAConfigurationRequest) SetHigh(v int32) {
	o.High = &v
}

// GetEnforceHigh returns the EnforceHigh field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetEnforceHigh() bool {
	if o == nil || IsNil(o.EnforceHigh) {
		var ret bool
		return ret
	}
	return *o.EnforceHigh
}

// GetEnforceHighOk returns a tuple with the EnforceHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetEnforceHighOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceHigh) {
		return nil, false
	}
	return o.EnforceHigh, true
}

// HasEnforceHigh returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasEnforceHigh() bool {
	if o != nil && !IsNil(o.EnforceHigh) {
		return true
	}

	return false
}

// SetEnforceHigh gets a reference to the given bool and assigns it to the EnforceHigh field.
func (o *SLAConfigurationRequest) SetEnforceHigh(v bool) {
	o.EnforceHigh = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetMedium() int32 {
	if o == nil || IsNil(o.Medium) {
		var ret int32
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetMediumOk() (*int32, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given int32 and assigns it to the Medium field.
func (o *SLAConfigurationRequest) SetMedium(v int32) {
	o.Medium = &v
}

// GetEnforceMedium returns the EnforceMedium field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetEnforceMedium() bool {
	if o == nil || IsNil(o.EnforceMedium) {
		var ret bool
		return ret
	}
	return *o.EnforceMedium
}

// GetEnforceMediumOk returns a tuple with the EnforceMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetEnforceMediumOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceMedium) {
		return nil, false
	}
	return o.EnforceMedium, true
}

// HasEnforceMedium returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasEnforceMedium() bool {
	if o != nil && !IsNil(o.EnforceMedium) {
		return true
	}

	return false
}

// SetEnforceMedium gets a reference to the given bool and assigns it to the EnforceMedium field.
func (o *SLAConfigurationRequest) SetEnforceMedium(v bool) {
	o.EnforceMedium = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetLow() int32 {
	if o == nil || IsNil(o.Low) {
		var ret int32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetLowOk() (*int32, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given int32 and assigns it to the Low field.
func (o *SLAConfigurationRequest) SetLow(v int32) {
	o.Low = &v
}

// GetEnforceLow returns the EnforceLow field value if set, zero value otherwise.
func (o *SLAConfigurationRequest) GetEnforceLow() bool {
	if o == nil || IsNil(o.EnforceLow) {
		var ret bool
		return ret
	}
	return *o.EnforceLow
}

// GetEnforceLowOk returns a tuple with the EnforceLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLAConfigurationRequest) GetEnforceLowOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceLow) {
		return nil, false
	}
	return o.EnforceLow, true
}

// HasEnforceLow returns a boolean if a field has been set.
func (o *SLAConfigurationRequest) HasEnforceLow() bool {
	if o != nil && !IsNil(o.EnforceLow) {
		return true
	}

	return false
}

// SetEnforceLow gets a reference to the given bool and assigns it to the EnforceLow field.
func (o *SLAConfigurationRequest) SetEnforceLow(v bool) {
	o.EnforceLow = &v
}

func (o SLAConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SLAConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *SLAConfigurationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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
	varSLAConfigurationRequest := _SLAConfigurationRequest{}

	err = json.Unmarshal(data, &varSLAConfigurationRequest)

	if err != nil {
		return err
	}

	*o = SLAConfigurationRequest(varSLAConfigurationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

type NullableSLAConfigurationRequest struct {
	value *SLAConfigurationRequest
	isSet bool
}

func (v NullableSLAConfigurationRequest) Get() *SLAConfigurationRequest {
	return v.value
}

func (v *NullableSLAConfigurationRequest) Set(val *SLAConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSLAConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSLAConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLAConfigurationRequest(val *SLAConfigurationRequest) *NullableSLAConfigurationRequest {
	return &NullableSLAConfigurationRequest{value: val, isSet: true}
}

func (v NullableSLAConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLAConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


