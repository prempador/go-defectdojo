/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the EndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointRequest{}

// EndpointRequest struct for EndpointRequest
type EndpointRequest struct {
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
}

// NewEndpointRequest instantiates a new EndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointRequest() *EndpointRequest {
	this := EndpointRequest{}
	return &this
}

// NewEndpointRequestWithDefaults instantiates a new EndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointRequestWithDefaults() *EndpointRequest {
	this := EndpointRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EndpointRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EndpointRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *EndpointRequest) SetTags(v []string) {
	o.Tags = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *EndpointRequest) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *EndpointRequest) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *EndpointRequest) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *EndpointRequest) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetUserinfo returns the Userinfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetUserinfo() string {
	if o == nil || IsNil(o.Userinfo.Get()) {
		var ret string
		return ret
	}
	return *o.Userinfo.Get()
}

// GetUserinfoOk returns a tuple with the Userinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetUserinfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Userinfo.Get(), o.Userinfo.IsSet()
}

// HasUserinfo returns a boolean if a field has been set.
func (o *EndpointRequest) HasUserinfo() bool {
	if o != nil && o.Userinfo.IsSet() {
		return true
	}

	return false
}

// SetUserinfo gets a reference to the given NullableString and assigns it to the Userinfo field.
func (o *EndpointRequest) SetUserinfo(v string) {
	o.Userinfo.Set(&v)
}
// SetUserinfoNil sets the value for Userinfo to be an explicit nil
func (o *EndpointRequest) SetUserinfoNil() {
	o.Userinfo.Set(nil)
}

// UnsetUserinfo ensures that no value is present for Userinfo, not even an explicit nil
func (o *EndpointRequest) UnsetUserinfo() {
	o.Userinfo.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *EndpointRequest) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *EndpointRequest) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *EndpointRequest) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *EndpointRequest) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *EndpointRequest) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *EndpointRequest) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *EndpointRequest) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *EndpointRequest) UnsetPort() {
	o.Port.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *EndpointRequest) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *EndpointRequest) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *EndpointRequest) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *EndpointRequest) UnsetPath() {
	o.Path.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetQuery() string {
	if o == nil || IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *EndpointRequest) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *EndpointRequest) SetQuery(v string) {
	o.Query.Set(&v)
}
// SetQueryNil sets the value for Query to be an explicit nil
func (o *EndpointRequest) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *EndpointRequest) UnsetQuery() {
	o.Query.Unset()
}

// GetFragment returns the Fragment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetFragment() string {
	if o == nil || IsNil(o.Fragment.Get()) {
		var ret string
		return ret
	}
	return *o.Fragment.Get()
}

// GetFragmentOk returns a tuple with the Fragment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetFragmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fragment.Get(), o.Fragment.IsSet()
}

// HasFragment returns a boolean if a field has been set.
func (o *EndpointRequest) HasFragment() bool {
	if o != nil && o.Fragment.IsSet() {
		return true
	}

	return false
}

// SetFragment gets a reference to the given NullableString and assigns it to the Fragment field.
func (o *EndpointRequest) SetFragment(v string) {
	o.Fragment.Set(&v)
}
// SetFragmentNil sets the value for Fragment to be an explicit nil
func (o *EndpointRequest) SetFragmentNil() {
	o.Fragment.Set(nil)
}

// UnsetFragment ensures that no value is present for Fragment, not even an explicit nil
func (o *EndpointRequest) UnsetFragment() {
	o.Fragment.Unset()
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EndpointRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EndpointRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *EndpointRequest) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *EndpointRequest) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *EndpointRequest) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *EndpointRequest) UnsetProduct() {
	o.Product.Unset()
}

func (o EndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableEndpointRequest struct {
	value *EndpointRequest
	isSet bool
}

func (v NullableEndpointRequest) Get() *EndpointRequest {
	return v.value
}

func (v *NullableEndpointRequest) Set(val *EndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointRequest(val *EndpointRequest) *NullableEndpointRequest {
	return &NullableEndpointRequest{value: val, isSet: true}
}

func (v NullableEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


