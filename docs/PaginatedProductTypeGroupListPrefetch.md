# PaginatedProductTypeGroupListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**map[string]DojoGroup**](DojoGroup.md) |  | [optional] [readonly] 
**ProductType** | Pointer to [**map[string]ProductType**](ProductType.md) |  | [optional] [readonly] 
**Role** | Pointer to [**map[string]Role**](Role.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedProductTypeGroupListPrefetch

`func NewPaginatedProductTypeGroupListPrefetch() *PaginatedProductTypeGroupListPrefetch`

NewPaginatedProductTypeGroupListPrefetch instantiates a new PaginatedProductTypeGroupListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductTypeGroupListPrefetchWithDefaults

`func NewPaginatedProductTypeGroupListPrefetchWithDefaults() *PaginatedProductTypeGroupListPrefetch`

NewPaginatedProductTypeGroupListPrefetchWithDefaults instantiates a new PaginatedProductTypeGroupListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *PaginatedProductTypeGroupListPrefetch) GetGroup() map[string]DojoGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PaginatedProductTypeGroupListPrefetch) GetGroupOk() (*map[string]DojoGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PaginatedProductTypeGroupListPrefetch) SetGroup(v map[string]DojoGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PaginatedProductTypeGroupListPrefetch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetProductType

`func (o *PaginatedProductTypeGroupListPrefetch) GetProductType() map[string]ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PaginatedProductTypeGroupListPrefetch) GetProductTypeOk() (*map[string]ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PaginatedProductTypeGroupListPrefetch) SetProductType(v map[string]ProductType)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *PaginatedProductTypeGroupListPrefetch) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetRole

`func (o *PaginatedProductTypeGroupListPrefetch) GetRole() map[string]Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PaginatedProductTypeGroupListPrefetch) GetRoleOk() (*map[string]Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PaginatedProductTypeGroupListPrefetch) SetRole(v map[string]Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *PaginatedProductTypeGroupListPrefetch) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


