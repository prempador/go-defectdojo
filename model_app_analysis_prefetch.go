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

// checks if the AppAnalysisPrefetch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAnalysisPrefetch{}

// AppAnalysisPrefetch struct for AppAnalysisPrefetch
type AppAnalysisPrefetch struct {
	Product *map[string]Product `json:"product,omitempty"`
	User *map[string]UserStub `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAnalysisPrefetch AppAnalysisPrefetch

// NewAppAnalysisPrefetch instantiates a new AppAnalysisPrefetch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAnalysisPrefetch() *AppAnalysisPrefetch {
	this := AppAnalysisPrefetch{}
	return &this
}

// NewAppAnalysisPrefetchWithDefaults instantiates a new AppAnalysisPrefetch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAnalysisPrefetchWithDefaults() *AppAnalysisPrefetch {
	this := AppAnalysisPrefetch{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AppAnalysisPrefetch) GetProduct() map[string]Product {
	if o == nil || IsNil(o.Product) {
		var ret map[string]Product
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAnalysisPrefetch) GetProductOk() (*map[string]Product, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AppAnalysisPrefetch) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given map[string]Product and assigns it to the Product field.
func (o *AppAnalysisPrefetch) SetProduct(v map[string]Product) {
	o.Product = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AppAnalysisPrefetch) GetUser() map[string]UserStub {
	if o == nil || IsNil(o.User) {
		var ret map[string]UserStub
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAnalysisPrefetch) GetUserOk() (*map[string]UserStub, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AppAnalysisPrefetch) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given map[string]UserStub and assigns it to the User field.
func (o *AppAnalysisPrefetch) SetUser(v map[string]UserStub) {
	o.User = &v
}

func (o AppAnalysisPrefetch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAnalysisPrefetch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAnalysisPrefetch) UnmarshalJSON(data []byte) (err error) {
	varAppAnalysisPrefetch := _AppAnalysisPrefetch{}

	err = json.Unmarshal(data, &varAppAnalysisPrefetch)

	if err != nil {
		return err
	}

	*o = AppAnalysisPrefetch(varAppAnalysisPrefetch)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAnalysisPrefetch struct {
	value *AppAnalysisPrefetch
	isSet bool
}

func (v NullableAppAnalysisPrefetch) Get() *AppAnalysisPrefetch {
	return v.value
}

func (v *NullableAppAnalysisPrefetch) Set(val *AppAnalysisPrefetch) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAnalysisPrefetch) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAnalysisPrefetch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAnalysisPrefetch(val *AppAnalysisPrefetch) *NullableAppAnalysisPrefetch {
	return &NullableAppAnalysisPrefetch{value: val, isSet: true}
}

func (v NullableAppAnalysisPrefetch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAnalysisPrefetch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


