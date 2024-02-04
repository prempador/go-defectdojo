# ProductGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Product** | **int32** |  | 
**Group** | **int32** |  | 
**Role** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedProductGroupListPrefetch**](PaginatedProductGroupListPrefetch.md) |  | [optional] 

## Methods

### NewProductGroup

`func NewProductGroup(id int32, product int32, group int32, role int32, ) *ProductGroup`

NewProductGroup instantiates a new ProductGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductGroupWithDefaults

`func NewProductGroupWithDefaults() *ProductGroup`

NewProductGroupWithDefaults instantiates a new ProductGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetProduct

`func (o *ProductGroup) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductGroup) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductGroup) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetGroup

`func (o *ProductGroup) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ProductGroup) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ProductGroup) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetRole

`func (o *ProductGroup) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProductGroup) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProductGroup) SetRole(v int32)`

SetRole sets Role field to given value.


### GetPrefetch

`func (o *ProductGroup) GetPrefetch() PaginatedProductGroupListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ProductGroup) GetPrefetchOk() (*PaginatedProductGroupListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ProductGroup) SetPrefetch(v PaginatedProductGroupListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ProductGroup) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


