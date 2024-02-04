# MetaPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to [**map[string]Endpoint**](Endpoint.md) |  | [optional] [readonly] 
**Finding** | Pointer to [**map[string]Finding**](Finding.md) |  | [optional] [readonly] 
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 

## Methods

### NewMetaPrefetch

`func NewMetaPrefetch() *MetaPrefetch`

NewMetaPrefetch instantiates a new MetaPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaPrefetchWithDefaults

`func NewMetaPrefetchWithDefaults() *MetaPrefetch`

NewMetaPrefetchWithDefaults instantiates a new MetaPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *MetaPrefetch) GetEndpoint() map[string]Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetaPrefetch) GetEndpointOk() (*map[string]Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetaPrefetch) SetEndpoint(v map[string]Endpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MetaPrefetch) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetFinding

`func (o *MetaPrefetch) GetFinding() map[string]Finding`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *MetaPrefetch) GetFindingOk() (*map[string]Finding, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *MetaPrefetch) SetFinding(v map[string]Finding)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *MetaPrefetch) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### GetProduct

`func (o *MetaPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *MetaPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *MetaPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *MetaPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


