# ProductTypeGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ProductType** | **int32** |  | 
**Group** | **int32** |  | 
**Role** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedProductTypeGroupListPrefetch**](PaginatedProductTypeGroupListPrefetch.md) |  | [optional] 

## Methods

### NewProductTypeGroup

`func NewProductTypeGroup(id int32, productType int32, group int32, role int32, ) *ProductTypeGroup`

NewProductTypeGroup instantiates a new ProductTypeGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeGroupWithDefaults

`func NewProductTypeGroupWithDefaults() *ProductTypeGroup`

NewProductTypeGroupWithDefaults instantiates a new ProductTypeGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductTypeGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductTypeGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductTypeGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetProductType

`func (o *ProductTypeGroup) GetProductType() int32`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductTypeGroup) GetProductTypeOk() (*int32, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductTypeGroup) SetProductType(v int32)`

SetProductType sets ProductType field to given value.


### GetGroup

`func (o *ProductTypeGroup) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ProductTypeGroup) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ProductTypeGroup) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetRole

`func (o *ProductTypeGroup) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProductTypeGroup) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProductTypeGroup) SetRole(v int32)`

SetRole sets Role field to given value.


### GetPrefetch

`func (o *ProductTypeGroup) GetPrefetch() PaginatedProductTypeGroupListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ProductTypeGroup) GetPrefetchOk() (*PaginatedProductTypeGroupListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ProductTypeGroup) SetPrefetch(v PaginatedProductTypeGroupListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ProductTypeGroup) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


