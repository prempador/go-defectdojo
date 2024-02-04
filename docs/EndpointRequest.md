# EndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to **NullableString** | The communication protocol/scheme such as &#39;http&#39;, &#39;ftp&#39;, &#39;dns&#39;, etc. | [optional] 
**Userinfo** | Pointer to **NullableString** | User info as &#39;alice&#39;, &#39;bob&#39;, etc. | [optional] 
**Host** | Pointer to **NullableString** | The host name or IP address. It must not include the port number. For example &#39;127.0.0.1&#39;, &#39;localhost&#39;, &#39;yourdomain.com&#39;. | [optional] 
**Port** | Pointer to **NullableInt32** | The network port associated with the endpoint. | [optional] 
**Path** | Pointer to **NullableString** | The location of the resource, it must not start with a &#39;/&#39;. For example endpoint/420/edit | [optional] 
**Query** | Pointer to **NullableString** | The query string, the question mark should be omitted.For example &#39;group&#x3D;4&amp;team&#x3D;8&#39; | [optional] 
**Fragment** | Pointer to **NullableString** | The fragment identifier which follows the hash mark. The hash mark should be omitted. For example &#39;section-13&#39;, &#39;paragraph-2&#39;. | [optional] 
**Product** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewEndpointRequest

`func NewEndpointRequest() *EndpointRequest`

NewEndpointRequest instantiates a new EndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointRequestWithDefaults

`func NewEndpointRequestWithDefaults() *EndpointRequest`

NewEndpointRequestWithDefaults instantiates a new EndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *EndpointRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EndpointRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EndpointRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EndpointRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProtocol

`func (o *EndpointRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *EndpointRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *EndpointRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *EndpointRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *EndpointRequest) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *EndpointRequest) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetUserinfo

`func (o *EndpointRequest) GetUserinfo() string`

GetUserinfo returns the Userinfo field if non-nil, zero value otherwise.

### GetUserinfoOk

`func (o *EndpointRequest) GetUserinfoOk() (*string, bool)`

GetUserinfoOk returns a tuple with the Userinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfo

`func (o *EndpointRequest) SetUserinfo(v string)`

SetUserinfo sets Userinfo field to given value.

### HasUserinfo

`func (o *EndpointRequest) HasUserinfo() bool`

HasUserinfo returns a boolean if a field has been set.

### SetUserinfoNil

`func (o *EndpointRequest) SetUserinfoNil(b bool)`

 SetUserinfoNil sets the value for Userinfo to be an explicit nil

### UnsetUserinfo
`func (o *EndpointRequest) UnsetUserinfo()`

UnsetUserinfo ensures that no value is present for Userinfo, not even an explicit nil
### GetHost

`func (o *EndpointRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EndpointRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EndpointRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EndpointRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *EndpointRequest) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *EndpointRequest) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *EndpointRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EndpointRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EndpointRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EndpointRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *EndpointRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *EndpointRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPath

`func (o *EndpointRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EndpointRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EndpointRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EndpointRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *EndpointRequest) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *EndpointRequest) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQuery

`func (o *EndpointRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EndpointRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EndpointRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EndpointRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *EndpointRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *EndpointRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetFragment

`func (o *EndpointRequest) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *EndpointRequest) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *EndpointRequest) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *EndpointRequest) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### SetFragmentNil

`func (o *EndpointRequest) SetFragmentNil(b bool)`

 SetFragmentNil sets the value for Fragment to be an explicit nil

### UnsetFragment
`func (o *EndpointRequest) UnsetFragment()`

UnsetFragment ensures that no value is present for Fragment, not even an explicit nil
### GetProduct

`func (o *EndpointRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EndpointRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EndpointRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *EndpointRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *EndpointRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *EndpointRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


