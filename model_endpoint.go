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

// checks if the Endpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Endpoint{}

// Endpoint struct for Endpoint
type Endpoint struct {
	Id int32 `json:"id"`
	Tags []string `json:"tags,omitempty"`
	// The communication protocol/scheme such as 'http', 'ftp', 'dns', etc.
	Protocol NullableString `json:"protocol,omitempty"`
	// User info as 'alice', 'bob', etc.
	Userinfo NullableString `json:"userinfo,omitempty"`
	// The host name or IP address. It must not include the port number. For example '127.0.0.1', 'localhost', 'yourdomain.com'.
	Host NullableString `json:"host,omitempty"`
	// The network port associated with the endpoint.
	Port NullableInt32 `json:"port,omitempty"`
	// The location of the resource, it must not start with a '/'. For example endpoint/420/edit
	Path NullableString `json:"path,omitempty"`
	// The query string, the question mark should be omitted.For example 'group=4&team=8'
	Query NullableString `json:"query,omitempty"`
	// The fragment identifier which follows the hash mark. The hash mark should be omitted. For example 'section-13', 'paragraph-2'.
	Fragment NullableString `json:"fragment,omitempty"`
	Product NullableInt32 `json:"product,omitempty"`
	EndpointParams []int32 `json:"endpoint_params"`
	Findings []int32 `json:"findings"`
	AdditionalProperties map[string]interface{}
}

type _Endpoint Endpoint

// NewEndpoint instantiates a new Endpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpoint(id int32, endpointParams []int32, findings []int32) *Endpoint {
	this := Endpoint{}
	this.Id = id
	this.EndpointParams = endpointParams
	this.Findings = findings
	return &this
}

// NewEndpointWithDefaults instantiates a new Endpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointWithDefaults() *Endpoint {
	this := Endpoint{}
	return &this
}

// GetId returns the Id field value
func (o *Endpoint) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Endpoint) SetId(v int32) {
	o.Id = v
}


// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Endpoint) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Endpoint) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Endpoint) SetTags(v []string) {
	o.Tags = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *Endpoint) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *Endpoint) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *Endpoint) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *Endpoint) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetUserinfo returns the Userinfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetUserinfo() string {
	if o == nil || IsNil(o.Userinfo.Get()) {
		var ret string
		return ret
	}
	return *o.Userinfo.Get()
}

// GetUserinfoOk returns a tuple with the Userinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetUserinfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Userinfo.Get(), o.Userinfo.IsSet()
}

// HasUserinfo returns a boolean if a field has been set.
func (o *Endpoint) HasUserinfo() bool {
	if o != nil && o.Userinfo.IsSet() {
		return true
	}

	return false
}

// SetUserinfo gets a reference to the given NullableString and assigns it to the Userinfo field.
func (o *Endpoint) SetUserinfo(v string) {
	o.Userinfo.Set(&v)
}
// SetUserinfoNil sets the value for Userinfo to be an explicit nil
func (o *Endpoint) SetUserinfoNil() {
	o.Userinfo.Set(nil)
}

// UnsetUserinfo ensures that no value is present for Userinfo, not even an explicit nil
func (o *Endpoint) UnsetUserinfo() {
	o.Userinfo.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *Endpoint) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *Endpoint) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *Endpoint) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *Endpoint) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *Endpoint) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *Endpoint) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *Endpoint) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *Endpoint) UnsetPort() {
	o.Port.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *Endpoint) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *Endpoint) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *Endpoint) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *Endpoint) UnsetPath() {
	o.Path.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetQuery() string {
	if o == nil || IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *Endpoint) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *Endpoint) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *Endpoint) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *Endpoint) UnsetQuery() {
	o.Query.Unset()
}

// GetFragment returns the Fragment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetFragment() string {
	if o == nil || IsNil(o.Fragment.Get()) {
		var ret string
		return ret
	}
	return *o.Fragment.Get()
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetFragmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fragment.Get(), o.Fragment.IsSet()
}

// HasFragment returns a boolean if a field has been set.
func (o *Endpoint) HasFragment() bool {
	if o != nil && o.Fragment.IsSet() {
		return true
	}

	return false
}

// SetFragment gets a reference to the given NullableString and assigns it to the Fragment field.
func (o *Endpoint) SetFragment(v string) {
	o.Fragment.Set(&v)
}
// SetFragmentNil sets the value for Fragment to be an explicit nil
func (o *Endpoint) SetFragmentNil() {
	o.Fragment.Set(nil)
}

// UnsetFragment ensures that no value is present for Fragment, not even an explicit nil
func (o *Endpoint) UnsetFragment() {
	o.Fragment.Unset()
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Endpoint) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Endpoint) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *Endpoint) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *Endpoint) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *Endpoint) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *Endpoint) UnsetProduct() {
	o.Product.Unset()
}

// GetEndpointParams returns the EndpointParams field value
func (o *Endpoint) GetEndpointParams() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.EndpointParams
}

// GetEndpointParamsOk returns a tuple with the EndpointParams field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetEndpointParamsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndpointParams, true
}

// SetEndpointParams sets field value
func (o *Endpoint) SetEndpointParams(v []int32) {
	o.EndpointParams = v
}


// GetFindings returns the Findings field value
func (o *Endpoint) GetFindings() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetFindingsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Findings, true
}

// SetFindings sets field value
func (o *Endpoint) SetFindings(v []int32) {
	o.Findings = v
}


func (o Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Endpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	if o.Userinfo.IsSet() {
		toSerialize["userinfo"] = o.Userinfo.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Fragment.IsSet() {
		toSerialize["fragment"] = o.Fragment.Get()
	}
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	toSerialize["endpoint_params"] = o.EndpointParams
	toSerialize["findings"] = o.Findings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Endpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"endpoint_params",
		"findings",
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
	varEndpoint := _Endpoint{}

	err = json.Unmarshal(data, &varEndpoint)

	if err != nil {
		return err
	}

	*o = Endpoint(varEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "userinfo")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "path")
		delete(additionalProperties, "query")
		delete(additionalProperties, "fragment")
		delete(additionalProperties, "product")
		delete(additionalProperties, "endpoint_params")
		delete(additionalProperties, "findings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndpoint struct {
	value *Endpoint
	isSet bool
}

func (v NullableEndpoint) Get() *Endpoint {
	return v.value
}

func (v *NullableEndpoint) Set(val *Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpoint(val *Endpoint) *NullableEndpoint {
	return &NullableEndpoint{value: val, isSet: true}
}

func (v NullableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


