# Endpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to **NullableString** | The communication protocol/scheme such as &#39;http&#39;, &#39;ftp&#39;, &#39;dns&#39;, etc. | [optional] 
**Userinfo** | Pointer to **NullableString** | User info as &#39;alice&#39;, &#39;bob&#39;, etc. | [optional] 
**Host** | Pointer to **NullableString** | The host name or IP address. It must not include the port number. For example &#39;127.0.0.1&#39;, &#39;localhost&#39;, &#39;yourdomain.com&#39;. | [optional] 
**Port** | Pointer to **NullableInt32** | The network port associated with the endpoint. | [optional] 
**Path** | Pointer to **NullableString** | The location of the resource, it must not start with a &#39;/&#39;. For example endpoint/420/edit | [optional] 
**Query** | Pointer to **NullableString** | The query string, the question mark should be omitted.For example &#39;group&#x3D;4&amp;team&#x3D;8&#39; | [optional] 
**Fragment** | Pointer to **NullableString** | The fragment identifier which follows the hash mark. The hash mark should be omitted. For example &#39;section-13&#39;, &#39;paragraph-2&#39;. | [optional] 
**Product** | Pointer to **NullableInt32** |  | [optional] 
**EndpointParams** | **[]int32** |  | [readonly] 
**Findings** | **[]int32** |  | [readonly] 

## Methods

### NewEndpoint

`func NewEndpoint(id int32, endpointParams []int32, findings []int32, ) *Endpoint`

NewEndpoint instantiates a new Endpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointWithDefaults

`func NewEndpointWithDefaults() *Endpoint`

NewEndpointWithDefaults instantiates a new Endpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Endpoint) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Endpoint) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Endpoint) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *Endpoint) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Endpoint) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Endpoint) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Endpoint) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProtocol

`func (o *Endpoint) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Endpoint) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Endpoint) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Endpoint) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *Endpoint) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *Endpoint) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetUserinfo

`func (o *Endpoint) GetUserinfo() string`

GetUserinfo returns the Userinfo field if non-nil, zero value otherwise.

### GetUserinfoOk

`func (o *Endpoint) GetUserinfoOk() (*string, bool)`

GetUserinfoOk returns a tuple with the Userinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfo

`func (o *Endpoint) SetUserinfo(v string)`

SetUserinfo sets Userinfo field to given value.

### HasUserinfo

`func (o *Endpoint) HasUserinfo() bool`

HasUserinfo returns a boolean if a field has been set.

### SetUserinfoNil

`func (o *Endpoint) SetUserinfoNil(b bool)`

 SetUserinfoNil sets the value for Userinfo to be an explicit nil

### UnsetUserinfo
`func (o *Endpoint) UnsetUserinfo()`

UnsetUserinfo ensures that no value is present for Userinfo, not even an explicit nil
### GetHost

`func (o *Endpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Endpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Endpoint) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Endpoint) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *Endpoint) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Endpoint) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *Endpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Endpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Endpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Endpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *Endpoint) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Endpoint) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPath

`func (o *Endpoint) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Endpoint) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Endpoint) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Endpoint) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *Endpoint) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *Endpoint) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetQuery

`func (o *Endpoint) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Endpoint) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Endpoint) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Endpoint) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *Endpoint) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *Endpoint) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetFragment

`func (o *Endpoint) GetFragment() string`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *Endpoint) GetFragmentOk() (*string, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *Endpoint) SetFragment(v string)`

SetFragment sets Fragment field to given value.

### HasFragment

`func (o *Endpoint) HasFragment() bool`

HasFragment returns a boolean if a field has been set.

### SetFragmentNil

`func (o *Endpoint) SetFragmentNil(b bool)`

 SetFragmentNil sets the value for Fragment to be an explicit nil

### UnsetFragment
`func (o *Endpoint) UnsetFragment()`

UnsetFragment ensures that no value is present for Fragment, not even an explicit nil
### GetProduct

`func (o *Endpoint) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Endpoint) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Endpoint) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Endpoint) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *Endpoint) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *Endpoint) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetEndpointParams

`func (o *Endpoint) GetEndpointParams() []int32`

GetEndpointParams returns the EndpointParams field if non-nil, zero value otherwise.

### GetEndpointParamsOk

`func (o *Endpoint) GetEndpointParamsOk() (*[]int32, bool)`

GetEndpointParamsOk returns a tuple with the EndpointParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointParams

`func (o *Endpoint) SetEndpointParams(v []int32)`

SetEndpointParams sets EndpointParams field to given value.


### GetFindings

`func (o *Endpoint) GetFindings() []int32`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *Endpoint) GetFindingsOk() (*[]int32, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *Endpoint) SetFindings(v []int32)`

SetFindings sets Findings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


