# PaginatedProductMemberListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**Role** | Pointer to [**map[string]Role**](Role.md) |  | [optional] [readonly] 
**User** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedProductMemberListPrefetch

`func NewPaginatedProductMemberListPrefetch() *PaginatedProductMemberListPrefetch`

NewPaginatedProductMemberListPrefetch instantiates a new PaginatedProductMemberListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductMemberListPrefetchWithDefaults

`func NewPaginatedProductMemberListPrefetchWithDefaults() *PaginatedProductMemberListPrefetch`

NewPaginatedProductMemberListPrefetchWithDefaults instantiates a new PaginatedProductMemberListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *PaginatedProductMemberListPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PaginatedProductMemberListPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PaginatedProductMemberListPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PaginatedProductMemberListPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRole

`func (o *PaginatedProductMemberListPrefetch) GetRole() map[string]Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PaginatedProductMemberListPrefetch) GetRoleOk() (*map[string]Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PaginatedProductMemberListPrefetch) SetRole(v map[string]Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *PaginatedProductMemberListPrefetch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *PaginatedProductMemberListPrefetch) GetUser() map[string]UserStub`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PaginatedProductMemberListPrefetch) GetUserOk() (*map[string]UserStub, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PaginatedProductMemberListPrefetch) SetUser(v map[string]UserStub)`

SetUser sets User field to given value.

### HasUser

`func (o *PaginatedProductMemberListPrefetch) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


