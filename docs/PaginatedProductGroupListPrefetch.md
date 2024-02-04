# PaginatedProductGroupListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**map[string]DojoGroup**](DojoGroup.md) |  | [optional] [readonly] 
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**Role** | Pointer to [**map[string]Role**](Role.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedProductGroupListPrefetch

`func NewPaginatedProductGroupListPrefetch() *PaginatedProductGroupListPrefetch`

NewPaginatedProductGroupListPrefetch instantiates a new PaginatedProductGroupListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductGroupListPrefetchWithDefaults

`func NewPaginatedProductGroupListPrefetchWithDefaults() *PaginatedProductGroupListPrefetch`

NewPaginatedProductGroupListPrefetchWithDefaults instantiates a new PaginatedProductGroupListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *PaginatedProductGroupListPrefetch) GetGroup() map[string]DojoGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PaginatedProductGroupListPrefetch) GetGroupOk() (*map[string]DojoGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PaginatedProductGroupListPrefetch) SetGroup(v map[string]DojoGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PaginatedProductGroupListPrefetch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetProduct

`func (o *PaginatedProductGroupListPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PaginatedProductGroupListPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PaginatedProductGroupListPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PaginatedProductGroupListPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRole

`func (o *PaginatedProductGroupListPrefetch) GetRole() map[string]Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PaginatedProductGroupListPrefetch) GetRoleOk() (*map[string]Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PaginatedProductGroupListPrefetch) SetRole(v map[string]Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *PaginatedProductGroupListPrefetch) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


