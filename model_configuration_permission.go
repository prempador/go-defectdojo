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

// checks if the ConfigurationPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPermission{}

// ConfigurationPermission struct for ConfigurationPermission
type ConfigurationPermission struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Codename string `json:"codename"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationPermission ConfigurationPermission

// NewConfigurationPermission instantiates a new ConfigurationPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationPermission(id int32, name string, codename string) *ConfigurationPermission {
	this := ConfigurationPermission{}
	this.Id = id
	this.Name = name
	this.Codename = codename
	return &this
}

// NewConfigurationPermissionWithDefaults instantiates a new ConfigurationPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPermissionWithDefaults() *ConfigurationPermission {
	this := ConfigurationPermission{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigurationPermission) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPermission) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigurationPermission) SetId(v int32) {
	o.Id = v
}


// GetName returns the Name field value
func (o *ConfigurationPermission) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPermission) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigurationPermission) SetName(v string) {
	o.Name = v
}


// GetCodename returns the Codename field value
func (o *ConfigurationPermission) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *ConfigurationPermission) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *ConfigurationPermission) SetCodename(v string) {
	o.Codename = v
}


func (o ConfigurationPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["codename"] = o.Codename

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationPermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"codename",
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
	varConfigurationPermission := _ConfigurationPermission{}

	err = json.Unmarshal(data, &varConfigurationPermission)

	if err != nil {
		return err
	}

	*o = ConfigurationPermission(varConfigurationPermission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "codename")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigurationPermission struct {
	value *ConfigurationPermission
	isSet bool
}

func (v NullableConfigurationPermission) Get() *ConfigurationPermission {
	return v.value
}

func (v *NullableConfigurationPermission) Set(val *ConfigurationPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationPermission(val *ConfigurationPermission) *NullableConfigurationPermission {
	return &NullableConfigurationPermission{value: val, isSet: true}
}

func (v NullableConfigurationPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


