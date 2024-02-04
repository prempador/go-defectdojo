# ProductTypeMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ProductType** | **int32** |  | 
**User** | **int32** |  | 
**Role** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedProductTypeMemberListPrefetch**](PaginatedProductTypeMemberListPrefetch.md) |  | [optional] 

## Methods

### NewProductTypeMember

`func NewProductTypeMember(id int32, productType int32, user int32, role int32, ) *ProductTypeMember`

NewProductTypeMember instantiates a new ProductTypeMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeMemberWithDefaults

`func NewProductTypeMemberWithDefaults() *ProductTypeMember`

NewProductTypeMemberWithDefaults instantiates a new ProductTypeMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductTypeMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductTypeMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductTypeMember) SetId(v int32)`

SetId sets Id field to given value.


### GetProductType

`func (o *ProductTypeMember) GetProductType() int32`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductTypeMember) GetProductTypeOk() (*int32, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductTypeMember) SetProductType(v int32)`

SetProductType sets ProductType field to given value.


### GetUser

`func (o *ProductTypeMember) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProductTypeMember) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProductTypeMember) SetUser(v int32)`

SetUser sets User field to given value.


### GetRole

`func (o *ProductTypeMember) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProductTypeMember) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProductTypeMember) SetRole(v int32)`

SetRole sets Role field to given value.


### GetPrefetch

`func (o *ProductTypeMember) GetPrefetch() PaginatedProductTypeMemberListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ProductTypeMember) GetPrefetchOk() (*PaginatedProductTypeMemberListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ProductTypeMember) SetPrefetch(v PaginatedProductTypeMemberListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ProductTypeMember) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


