# MetaMainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | [**[]MetadataRequest**](MetadataRequest.md) |  | 

## Methods

### NewMetaMainRequest

`func NewMetaMainRequest(metadata []MetadataRequest, ) *MetaMainRequest`

NewMetaMainRequest instantiates a new MetaMainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaMainRequestWithDefaults

`func NewMetaMainRequestWithDefaults() *MetaMainRequest`

NewMetaMainRequestWithDefaults instantiates a new MetaMainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *MetaMainRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *MetaMainRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *MetaMainRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *MetaMainRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *MetaMainRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *MetaMainRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetEndpoint

`func (o *MetaMainRequest) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetaMainRequest) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetaMainRequest) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MetaMainRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *MetaMainRequest) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *MetaMainRequest) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetFinding

`func (o *MetaMainRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *MetaMainRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *MetaMainRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *MetaMainRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *MetaMainRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *MetaMainRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetMetadata

`func (o *MetaMainRequest) GetMetadata() []MetadataRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetaMainRequest) GetMetadataOk() (*[]MetadataRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetaMainRequest) SetMetadata(v []MetadataRequest)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


