# MetaMain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | [**[]Metadata**](Metadata.md) |  | 

## Methods

### NewMetaMain

`func NewMetaMain(id int32, metadata []Metadata, ) *MetaMain`

NewMetaMain instantiates a new MetaMain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaMainWithDefaults

`func NewMetaMainWithDefaults() *MetaMain`

NewMetaMainWithDefaults instantiates a new MetaMain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetaMain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetaMain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetaMain) SetId(v int32)`

SetId sets Id field to given value.


### GetProduct

`func (o *MetaMain) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *MetaMain) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *MetaMain) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *MetaMain) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *MetaMain) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *MetaMain) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetEndpoint

`func (o *MetaMain) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetaMain) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetaMain) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MetaMain) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *MetaMain) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *MetaMain) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetFinding

`func (o *MetaMain) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *MetaMain) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *MetaMain) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *MetaMain) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *MetaMain) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *MetaMain) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetMetadata

`func (o *MetaMain) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetaMain) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetaMain) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


