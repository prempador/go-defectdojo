/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedProductAPIScanConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedProductAPIScanConfigurationRequest{}

// PatchedProductAPIScanConfigurationRequest struct for PatchedProductAPIScanConfigurationRequest
type PatchedProductAPIScanConfigurationRequest struct {
	ServiceKey1 NullableString `json:"service_key_1,omitempty"`
	ServiceKey2 NullableString `json:"service_key_2,omitempty"`
	ServiceKey3 NullableString `json:"service_key_3,omitempty"`
	Product *int32 `json:"product,omitempty"`
	ToolConfiguration *int32 `json:"tool_configuration,omitempty"`
}

// NewPatchedProductAPIScanConfigurationRequest instantiates a new PatchedProductAPIScanConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedProductAPIScanConfigurationRequest() *PatchedProductAPIScanConfigurationRequest {
	this := PatchedProductAPIScanConfigurationRequest{}
	return &this
}

// NewPatchedProductAPIScanConfigurationRequestWithDefaults instantiates a new PatchedProductAPIScanConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedProductAPIScanConfigurationRequestWithDefaults() *PatchedProductAPIScanConfigurationRequest {
	this := PatchedProductAPIScanConfigurationRequest{}
	return &this
}

// GetServiceKey1 returns the ServiceKey1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProductAPIScanConfigurationRequest) GetServiceKey1() string {
	if o == nil || IsNil(o.ServiceKey1.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceKey1.Get()
}

// GetServiceKey1Ok returns a tuple with the ServiceKey1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProductAPIScanConfigurationRequest) GetServiceKey1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceKey1.Get(), o.ServiceKey1.IsSet()
}

// HasServiceKey1 returns a boolean if a field has been set.
func (o *PatchedProductAPIScanConfigurationRequest) HasServiceKey1() bool {
	if o != nil && o.ServiceKey1.IsSet() {
		return true
	}

	return false
}

// SetServiceKey1 gets a reference to the given NullableString and assigns it to the ServiceKey1 field.
func (o *PatchedProductAPIScanConfigurationRequest) SetServiceKey1(v string) {
	o.ServiceKey1.Set(&v)
}
// SetServiceKey1Nil sets the value for ServiceKey1 to be an explicit nil
func (o *PatchedProductAPIScanConfigurationRequest) SetServiceKey1Nil() {
	o.ServiceKey1.Set(nil)
}

// UnsetServiceKey1 ensures that no value is present for ServiceKey1, not even an explicit nil
func (o *PatchedProductAPIScanConfigurationRequest) UnsetServiceKey1() {
	o.ServiceKey1.Unset()
}

// GetServiceKey2 returns the ServiceKey2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProductAPIScanConfigurationRequest) GetServiceKey2() string {
	if o == nil || IsNil(o.ServiceKey2.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceKey2.Get()
}

// GetServiceKey2Ok returns a tuple with the ServiceKey2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProductAPIScanConfigurationRequest) GetServiceKey2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceKey2.Get(), o.ServiceKey2.IsSet()
}

// HasServiceKey2 returns a boolean if a field has been set.
func (o *PatchedProductAPIScanConfigurationRequest) HasServiceKey2() bool {
	if o != nil && o.ServiceKey2.IsSet() {
		return true
	}

	return false
}

// SetServiceKey2 gets a reference to the given NullableString and assigns it to the ServiceKey2 field.
func (o *PatchedProductAPIScanConfigurationRequest) SetServiceKey2(v string) {
	o.ServiceKey2.Set(&v)
}
// SetServiceKey2Nil sets the value for ServiceKey2 to be an explicit nil
func (o *PatchedProductAPIScanConfigurationRequest) SetServiceKey2Nil() {
	o.ServiceKey2.Set(nil)
}

// UnsetServiceKey2 ensures that no value is present for ServiceKey2, not even an explicit nil
func (o *PatchedProductAPIScanConfigurationRequest) UnsetServiceKey2() {
	o.ServiceKey2.Unset()
}

// GetServiceKey3 returns the ServiceKey3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedProductAPIScanConfigurationRequest) GetServiceKey3() string {
	if o == nil || IsNil(o.ServiceKey3.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceKey3.Get()
}

// GetServiceKey3Ok returns a tuple with the ServiceKey3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedProductAPIScanConfigurationRequest) GetServiceKey3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceKey3.Get(), o.ServiceKey3.IsSet()
}

// HasServiceKey3 returns a boolean if a field has been set.
func (o *PatchedProductAPIScanConfigurationRequest) HasServiceKey3() bool {
	if o != nil && o.ServiceKey3.IsSet() {
		return true
	}

	return false
}

// SetServiceKey3 gets a reference to the given NullableString and assigns it to the ServiceKey3 field.
func (o *PatchedProductAPIScanConfigurationRequest) SetServiceKey3(v string) {
	o.ServiceKey3.Set(&v)
}
// SetServiceKey3Nil sets the value for ServiceKey3 to be an explicit nil
func (o *PatchedProductAPIScanConfigurationRequest) SetServiceKey3Nil() {
	o.ServiceKey3.Set(nil)
}

// UnsetServiceKey3 ensures that no value is present for ServiceKey3, not even an explicit nil
func (o *PatchedProductAPIScanConfigurationRequest) UnsetServiceKey3() {
	o.ServiceKey3.Unset()
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PatchedProductAPIScanConfigurationRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product) {
		var ret int32
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProductAPIScanConfigurationRequest) GetProductOk() (*int32, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PatchedProductAPIScanConfigurationRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given int32 and assigns it to the Product field.
func (o *PatchedProductAPIScanConfigurationRequest) SetProduct(v int32) {
	o.Product = &v
}

// GetToolConfiguration returns the ToolConfiguration field value if set, zero value otherwise.
func (o *PatchedProductAPIScanConfigurationRequest) GetToolConfiguration() int32 {
	if o == nil || IsNil(o.ToolConfiguration) {
		var ret int32
		return ret
	}
	return *o.ToolConfiguration
}

// GetToolConfigurationOk returns a tuple with the ToolConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProductAPIScanConfigurationRequest) GetToolConfigurationOk() (*int32, bool) {
	if o == nil || IsNil(o.ToolConfiguration) {
		return nil, false
	}
	return o.ToolConfiguration, true
}

// HasToolConfiguration returns a boolean if a field has been set.
func (o *PatchedProductAPIScanConfigurationRequest) HasToolConfiguration() bool {
	if o != nil && !IsNil(o.ToolConfiguration) {
		return true
	}

	return false
}

// SetToolConfiguration gets a reference to the given int32 and assigns it to the ToolConfiguration field.
func (o *PatchedProductAPIScanConfigurationRequest) SetToolConfiguration(v int32) {
	o.ToolConfiguration = &v
}

func (o PatchedProductAPIScanConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedProductAPIScanConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceKey1.IsSet() {
		toSerialize["service_key_1"] = o.ServiceKey1.Get()
	}
	if o.ServiceKey2.IsSet() {
		toSerialize["service_key_2"] = o.ServiceKey2.Get()
	}
	if o.ServiceKey3.IsSet() {
		toSerialize["service_key_3"] = o.ServiceKey3.Get()
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.ToolConfiguration) {
		toSerialize["tool_configuration"] = o.ToolConfiguration
	}
	return toSerialize, nil
}

type NullablePatchedProductAPIScanConfigurationRequest struct {
	value *PatchedProductAPIScanConfigurationRequest
	isSet bool
}

func (v NullablePatchedProductAPIScanConfigurationRequest) Get() *PatchedProductAPIScanConfigurationRequest {
	return v.value
}

func (v *NullablePatchedProductAPIScanConfigurationRequest) Set(val *PatchedProductAPIScanConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedProductAPIScanConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedProductAPIScanConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedProductAPIScanConfigurationRequest(val *PatchedProductAPIScanConfigurationRequest) *NullablePatchedProductAPIScanConfigurationRequest {
	return &NullablePatchedProductAPIScanConfigurationRequest{value: val, isSet: true}
}

func (v NullablePatchedProductAPIScanConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedProductAPIScanConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


