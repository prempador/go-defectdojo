# MetaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewMetaRequest

`func NewMetaRequest(name string, value string, ) *MetaRequest`

NewMetaRequest instantiates a new MetaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaRequestWithDefaults

`func NewMetaRequestWithDefaults() *MetaRequest`

NewMetaRequestWithDefaults instantiates a new MetaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *MetaRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *MetaRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *MetaRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *MetaRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *MetaRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *MetaRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetEndpoint

`func (o *MetaRequest) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetaRequest) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetaRequest) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MetaRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *MetaRequest) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *MetaRequest) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetFinding

`func (o *MetaRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *MetaRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *MetaRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *MetaRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *MetaRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *MetaRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetName

`func (o *MetaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *MetaRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetaRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetaRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


