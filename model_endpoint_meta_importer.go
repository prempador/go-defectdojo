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

// checks if the EndpointMetaImporter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointMetaImporter{}

// EndpointMetaImporter struct for EndpointMetaImporter
type EndpointMetaImporter struct {
	File string `json:"file"`
	CreateEndpoints *bool `json:"create_endpoints,omitempty"`
	CreateTags *bool `json:"create_tags,omitempty"`
	CreateDojoMeta *bool `json:"create_dojo_meta,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	Product *int32 `json:"product,omitempty"`
	ProductId int32 `json:"product_id"`
	AdditionalProperties map[string]interface{}
}

type _EndpointMetaImporter EndpointMetaImporter

// NewEndpointMetaImporter instantiates a new EndpointMetaImporter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointMetaImporter(file string, productId int32) *EndpointMetaImporter {
	this := EndpointMetaImporter{}
	this.File = file
	var createEndpoints bool = true
	this.CreateEndpoints = &createEndpoints
	var createTags bool = true
	this.CreateTags = &createTags
	var createDojoMeta bool = false
	this.CreateDojoMeta = &createDojoMeta
	this.ProductId = productId
	return &this
}

// NewEndpointMetaImporterWithDefaults instantiates a new EndpointMetaImporter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointMetaImporterWithDefaults() *EndpointMetaImporter {
	this := EndpointMetaImporter{}
	var createEndpoints bool = true
	this.CreateEndpoints = &createEndpoints
	var createTags bool = true
	this.CreateTags = &createTags
	var createDojoMeta bool = false
	this.CreateDojoMeta = &createDojoMeta
	return &this
}

// GetFile returns the File field value
func (o *EndpointMetaImporter) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *EndpointMetaImporter) SetFile(v string) {
	o.File = v
}


// GetCreateEndpoints returns the CreateEndpoints field value if set, zero value otherwise.
func (o *EndpointMetaImporter) GetCreateEndpoints() bool {
	if o == nil || IsNil(o.CreateEndpoints) {
		var ret bool
		return ret
	}
	return *o.CreateEndpoints
}

// GetCreateEndpointsOk returns a tuple with the CreateEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetCreateEndpointsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateEndpoints) {
		return nil, false
	}
	return o.CreateEndpoints, true
}

// HasCreateEndpoints returns a boolean if a field has been set.
func (o *EndpointMetaImporter) HasCreateEndpoints() bool {
	if o != nil && !IsNil(o.CreateEndpoints) {
		return true
	}

	return false
}

// SetCreateEndpoints gets a reference to the given bool and assigns it to the CreateEndpoints field.
func (o *EndpointMetaImporter) SetCreateEndpoints(v bool) {
	o.CreateEndpoints = &v
}

// GetCreateTags returns the CreateTags field value if set, zero value otherwise.
func (o *EndpointMetaImporter) GetCreateTags() bool {
	if o == nil || IsNil(o.CreateTags) {
		var ret bool
		return ret
	}
	return *o.CreateTags
}

// GetCreateTagsOk returns a tuple with the CreateTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetCreateTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateTags) {
		return nil, false
	}
	return o.CreateTags, true
}

// HasCreateTags returns a boolean if a field has been set.
func (o *EndpointMetaImporter) HasCreateTags() bool {
	if o != nil && !IsNil(o.CreateTags) {
		return true
	}

	return false
}

// SetCreateTags gets a reference to the given bool and assigns it to the CreateTags field.
func (o *EndpointMetaImporter) SetCreateTags(v bool) {
	o.CreateTags = &v
}

// GetCreateDojoMeta returns the CreateDojoMeta field value if set, zero value otherwise.
func (o *EndpointMetaImporter) GetCreateDojoMeta() bool {
	if o == nil || IsNil(o.CreateDojoMeta) {
		var ret bool
		return ret
	}
	return *o.CreateDojoMeta
}

// GetCreateDojoMetaOk returns a tuple with the CreateDojoMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetCreateDojoMetaOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateDojoMeta) {
		return nil, false
	}
	return o.CreateDojoMeta, true
}

// HasCreateDojoMeta returns a boolean if a field has been set.
func (o *EndpointMetaImporter) HasCreateDojoMeta() bool {
	if o != nil && !IsNil(o.CreateDojoMeta) {
		return true
	}

	return false
}

// SetCreateDojoMeta gets a reference to the given bool and assigns it to the CreateDojoMeta field.
func (o *EndpointMetaImporter) SetCreateDojoMeta(v bool) {
	o.CreateDojoMeta = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *EndpointMetaImporter) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *EndpointMetaImporter) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *EndpointMetaImporter) SetProductName(v string) {
	o.ProductName = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *EndpointMetaImporter) GetProduct() int32 {
	if o == nil || IsNil(o.Product) {
		var ret int32
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetProductOk() (*int32, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *EndpointMetaImporter) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given int32 and assigns it to the Product field.
func (o *EndpointMetaImporter) SetProduct(v int32) {
	o.Product = &v
}

// GetProductId returns the ProductId field value
func (o *EndpointMetaImporter) GetProductId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *EndpointMetaImporter) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *EndpointMetaImporter) SetProductId(v int32) {
	o.ProductId = v
}


func (o EndpointMetaImporter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointMetaImporter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	if !IsNil(o.CreateEndpoints) {
		toSerialize["create_endpoints"] = o.CreateEndpoints
	}
	if !IsNil(o.CreateTags) {
		toSerialize["create_tags"] = o.CreateTags
	}
	if !IsNil(o.CreateDojoMeta) {
		toSerialize["create_dojo_meta"] = o.CreateDojoMeta
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	toSerialize["product_id"] = o.ProductId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EndpointMetaImporter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
		"product_id",
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
	varEndpointMetaImporter := _EndpointMetaImporter{}

	err = json.Unmarshal(data, &varEndpointMetaImporter)

	if err != nil {
		return err
	}

	*o = EndpointMetaImporter(varEndpointMetaImporter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "create_endpoints")
		delete(additionalProperties, "create_tags")
		delete(additionalProperties, "create_dojo_meta")
		delete(additionalProperties, "product_name")
		delete(additionalProperties, "product")
		delete(additionalProperties, "product_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndpointMetaImporter struct {
	value *EndpointMetaImporter
	isSet bool
}

func (v NullableEndpointMetaImporter) Get() *EndpointMetaImporter {
	return v.value
}

func (v *NullableEndpointMetaImporter) Set(val *EndpointMetaImporter) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointMetaImporter) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointMetaImporter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointMetaImporter(val *EndpointMetaImporter) *NullableEndpointMetaImporter {
	return &NullableEndpointMetaImporter{value: val, isSet: true}
}

func (v NullableEndpointMetaImporter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointMetaImporter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


