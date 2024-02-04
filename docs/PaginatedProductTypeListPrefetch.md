# PaginatedProductTypeListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationGroups** | Pointer to [**map[string]DojoGroup**](DojoGroup.md) |  | [optional] [readonly] 
**Members** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedProductTypeListPrefetch

`func NewPaginatedProductTypeListPrefetch() *PaginatedProductTypeListPrefetch`

NewPaginatedProductTypeListPrefetch instantiates a new PaginatedProductTypeListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductTypeListPrefetchWithDefaults

`func NewPaginatedProductTypeListPrefetchWithDefaults() *PaginatedProductTypeListPrefetch`

NewPaginatedProductTypeListPrefetchWithDefaults instantiates a new PaginatedProductTypeListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationGroups

`func (o *PaginatedProductTypeListPrefetch) GetAuthorizationGroups() map[string]DojoGroup`

GetAuthorizationGroups returns the AuthorizationGroups field if non-nil, zero value otherwise.

### GetAuthorizationGroupsOk

`func (o *PaginatedProductTypeListPrefetch) GetAuthorizationGroupsOk() (*map[string]DojoGroup, bool)`

GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroups

`func (o *PaginatedProductTypeListPrefetch) SetAuthorizationGroups(v map[string]DojoGroup)`

SetAuthorizationGroups sets AuthorizationGroups field to given value.

### HasAuthorizationGroups

`func (o *PaginatedProductTypeListPrefetch) HasAuthorizationGroups() bool`

HasAuthorizationGroups returns a boolean if a field has been set.

### GetMembers

`func (o *PaginatedProductTypeListPrefetch) GetMembers() map[string]UserStub`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PaginatedProductTypeListPrefetch) GetMembersOk() (*map[string]UserStub, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PaginatedProductTypeListPrefetch) SetMembers(v map[string]UserStub)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PaginatedProductTypeListPrefetch) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


