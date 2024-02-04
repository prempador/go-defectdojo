# ProductMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Product** | **int32** |  | 
**User** | **int32** |  | 
**Role** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedProductMemberListPrefetch**](PaginatedProductMemberListPrefetch.md) |  | [optional] 

## Methods

### NewProductMember

`func NewProductMember(id int32, product int32, user int32, role int32, ) *ProductMember`

NewProductMember instantiates a new ProductMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMemberWithDefaults

`func NewProductMemberWithDefaults() *ProductMember`

NewProductMemberWithDefaults instantiates a new ProductMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductMember) SetId(v int32)`

SetId sets Id field to given value.


### GetProduct

`func (o *ProductMember) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductMember) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductMember) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetUser

`func (o *ProductMember) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProductMember) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProductMember) SetUser(v int32)`

SetUser sets User field to given value.


### GetRole

`func (o *ProductMember) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProductMember) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProductMember) SetRole(v int32)`

SetRole sets Role field to given value.


### GetPrefetch

`func (o *ProductMember) GetPrefetch() PaginatedProductMemberListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ProductMember) GetPrefetchOk() (*PaginatedProductMemberListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ProductMember) SetPrefetch(v PaginatedProductMemberListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ProductMember) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


