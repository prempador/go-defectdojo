# EndpointMetaImporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** |  | 
**CreateEndpoints** | Pointer to **bool** |  | [optional] [default to true]
**CreateTags** | Pointer to **bool** |  | [optional] [default to true]
**CreateDojoMeta** | Pointer to **bool** |  | [optional] [default to false]
**ProductName** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **int32** |  | [optional] 
**ProductId** | **int32** |  | [readonly] 

## Methods

### NewEndpointMetaImporter

`func NewEndpointMetaImporter(file string, productId int32, ) *EndpointMetaImporter`

NewEndpointMetaImporter instantiates a new EndpointMetaImporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointMetaImporterWithDefaults

`func NewEndpointMetaImporterWithDefaults() *EndpointMetaImporter`

NewEndpointMetaImporterWithDefaults instantiates a new EndpointMetaImporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *EndpointMetaImporter) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *EndpointMetaImporter) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *EndpointMetaImporter) SetFile(v string)`

SetFile sets File field to given value.


### GetCreateEndpoints

`func (o *EndpointMetaImporter) GetCreateEndpoints() bool`

GetCreateEndpoints returns the CreateEndpoints field if non-nil, zero value otherwise.

### GetCreateEndpointsOk

`func (o *EndpointMetaImporter) GetCreateEndpointsOk() (*bool, bool)`

GetCreateEndpointsOk returns a tuple with the CreateEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEndpoints

`func (o *EndpointMetaImporter) SetCreateEndpoints(v bool)`

SetCreateEndpoints sets CreateEndpoints field to given value.

### HasCreateEndpoints

`func (o *EndpointMetaImporter) HasCreateEndpoints() bool`

HasCreateEndpoints returns a boolean if a field has been set.

### GetCreateTags

`func (o *EndpointMetaImporter) GetCreateTags() bool`

GetCreateTags returns the CreateTags field if non-nil, zero value otherwise.

### GetCreateTagsOk

`func (o *EndpointMetaImporter) GetCreateTagsOk() (*bool, bool)`

GetCreateTagsOk returns a tuple with the CreateTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTags

`func (o *EndpointMetaImporter) SetCreateTags(v bool)`

SetCreateTags sets CreateTags field to given value.

### HasCreateTags

`func (o *EndpointMetaImporter) HasCreateTags() bool`

HasCreateTags returns a boolean if a field has been set.

### GetCreateDojoMeta

`func (o *EndpointMetaImporter) GetCreateDojoMeta() bool`

GetCreateDojoMeta returns the CreateDojoMeta field if non-nil, zero value otherwise.

### GetCreateDojoMetaOk

`func (o *EndpointMetaImporter) GetCreateDojoMetaOk() (*bool, bool)`

GetCreateDojoMetaOk returns a tuple with the CreateDojoMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDojoMeta

`func (o *EndpointMetaImporter) SetCreateDojoMeta(v bool)`

SetCreateDojoMeta sets CreateDojoMeta field to given value.

### HasCreateDojoMeta

`func (o *EndpointMetaImporter) HasCreateDojoMeta() bool`

HasCreateDojoMeta returns a boolean if a field has been set.

### GetProductName

`func (o *EndpointMetaImporter) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *EndpointMetaImporter) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *EndpointMetaImporter) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *EndpointMetaImporter) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProduct

`func (o *EndpointMetaImporter) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EndpointMetaImporter) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EndpointMetaImporter) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *EndpointMetaImporter) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductId

`func (o *EndpointMetaImporter) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *EndpointMetaImporter) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *EndpointMetaImporter) SetProductId(v int32)`

SetProductId sets ProductId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


