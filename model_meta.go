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

// checks if the Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Meta{}

// Meta struct for Meta
type Meta struct {
	Id int32 `json:"id"`
	Product NullableInt32 `json:"product,omitempty"`
	Endpoint NullableInt32 `json:"endpoint,omitempty"`
	Finding NullableInt32 `json:"finding,omitempty"`
	Name string `json:"name"`
	Value string `json:"value"`
	Prefetch *MetaPrefetch `json:"prefetch,omitempty"`
}

type _Meta Meta

// NewMeta instantiates a new Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeta(id int32, name string, value string) *Meta {
	this := Meta{}
	this.Id = id
	this.Name = name
	this.Value = value
	return &this
}

// NewMetaWithDefaults instantiates a new Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaWithDefaults() *Meta {
	this := Meta{}
	return &this
}

// GetId returns the Id field value
func (o *Meta) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Meta) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Meta) SetId(v int32) {
	o.Id = v
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Meta) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Meta) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *Meta) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *Meta) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *Meta) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *Meta) UnsetProduct() {
	o.Product.Unset()
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Meta) GetEndpoint() int32 {
	if o == nil || IsNil(o.Endpoint.Get()) {
		var ret int32
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Meta) GetEndpointOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *Meta) HasEndpoint() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableInt32 and assigns it to the Endpoint field.
func (o *Meta) SetEndpoint(v int32) {
	o.Endpoint.Set(&v)
}
// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *Meta) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *Meta) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetFinding returns the Finding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Meta) GetFinding() int32 {
	if o == nil || IsNil(o.Finding.Get()) {
		var ret int32
		return ret
	}
	return *o.Finding.Get()
}

// GetFindingOk returns a tuple with the Finding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Meta) GetFindingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Finding.Get(), o.Finding.IsSet()
}

// HasFinding returns a boolean if a field has been set.
func (o *Meta) HasFinding() bool {
	if o != nil && o.Finding.IsSet() {
		return true
	}

	return false
}

// SetFinding gets a reference to the given NullableInt32 and assigns it to the Finding field.
func (o *Meta) SetFinding(v int32) {
	o.Finding.Set(&v)
}
// SetFindingNil sets the value for Finding to be an explicit nil
func (o *Meta) SetFindingNil() {
	o.Finding.Set(nil)
}

// UnsetFinding ensures that no value is present for Finding, not even an explicit nil
func (o *Meta) UnsetFinding() {
	o.Finding.Unset()
}

// GetName returns the Name field value
func (o *Meta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Meta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Meta) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *Meta) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Meta) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Meta) SetValue(v string) {
	o.Value = v
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *Meta) GetPrefetch() MetaPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret MetaPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meta) GetPrefetchOk() (*MetaPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *Meta) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given MetaPrefetch and assigns it to the Prefetch field.
func (o *Meta) SetPrefetch(v MetaPrefetch) {
	o.Prefetch = &v
}

func (o Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if o.Finding.IsSet() {
		toSerialize["finding"] = o.Finding.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}
	return toSerialize, nil
}

func (o *Meta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"value",
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

	varMeta := _Meta{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeta)

	if err != nil {
		return err
	}

	*o = Meta(varMeta)

	return err
}

type NullableMeta struct {
	value *Meta
	isSet bool
}

func (v NullableMeta) Get() *Meta {
	return v.value
}

func (v *NullableMeta) Set(val *Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeta(val *Meta) *NullableMeta {
	return &NullableMeta{value: val, isSet: true}
}

func (v NullableMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

