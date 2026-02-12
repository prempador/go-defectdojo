# PatchedMetaMainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to [**[]MetadataRequest**](MetadataRequest.md) |  | [optional] 

## Methods

### NewPatchedMetaMainRequest

`func NewPatchedMetaMainRequest() *PatchedMetaMainRequest`

NewPatchedMetaMainRequest instantiates a new PatchedMetaMainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedMetaMainRequestWithDefaults

`func NewPatchedMetaMainRequestWithDefaults() *PatchedMetaMainRequest`

NewPatchedMetaMainRequestWithDefaults instantiates a new PatchedMetaMainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *PatchedMetaMainRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedMetaMainRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedMetaMainRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedMetaMainRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *PatchedMetaMainRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *PatchedMetaMainRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetEndpoint

`func (o *PatchedMetaMainRequest) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PatchedMetaMainRequest) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PatchedMetaMainRequest) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PatchedMetaMainRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *PatchedMetaMainRequest) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *PatchedMetaMainRequest) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetFinding

`func (o *PatchedMetaMainRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *PatchedMetaMainRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *PatchedMetaMainRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *PatchedMetaMainRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *PatchedMetaMainRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *PatchedMetaMainRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetMetadata

`func (o *PatchedMetaMainRequest) GetMetadata() []MetadataRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchedMetaMainRequest) GetMetadataOk() (*[]MetadataRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchedMetaMainRequest) SetMetadata(v []MetadataRequest)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchedMetaMainRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


