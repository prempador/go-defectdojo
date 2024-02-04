# DojoGroupPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductGroups** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**ProductTypeGroups** | Pointer to [**map[string]ProductType**](ProductType.md) |  | [optional] [readonly] 
**Users** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewDojoGroupPrefetch

`func NewDojoGroupPrefetch() *DojoGroupPrefetch`

NewDojoGroupPrefetch instantiates a new DojoGroupPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDojoGroupPrefetchWithDefaults

`func NewDojoGroupPrefetchWithDefaults() *DojoGroupPrefetch`

NewDojoGroupPrefetchWithDefaults instantiates a new DojoGroupPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductGroups

`func (o *DojoGroupPrefetch) GetProductGroups() map[string]Product`

GetProductGroups returns the ProductGroups field if non-nil, zero value otherwise.

### GetProductGroupsOk

`func (o *DojoGroupPrefetch) GetProductGroupsOk() (*map[string]Product, bool)`

GetProductGroupsOk returns a tuple with the ProductGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroups

`func (o *DojoGroupPrefetch) SetProductGroups(v map[string]Product)`

SetProductGroups sets ProductGroups field to given value.

### HasProductGroups

`func (o *DojoGroupPrefetch) HasProductGroups() bool`

HasProductGroups returns a boolean if a field has been set.

### GetProductTypeGroups

`func (o *DojoGroupPrefetch) GetProductTypeGroups() map[string]ProductType`

GetProductTypeGroups returns the ProductTypeGroups field if non-nil, zero value otherwise.

### GetProductTypeGroupsOk

`func (o *DojoGroupPrefetch) GetProductTypeGroupsOk() (*map[string]ProductType, bool)`

GetProductTypeGroupsOk returns a tuple with the ProductTypeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeGroups

`func (o *DojoGroupPrefetch) SetProductTypeGroups(v map[string]ProductType)`

SetProductTypeGroups sets ProductTypeGroups field to given value.

### HasProductTypeGroups

`func (o *DojoGroupPrefetch) HasProductTypeGroups() bool`

HasProductTypeGroups returns a boolean if a field has been set.

### GetUsers

`func (o *DojoGroupPrefetch) GetUsers() map[string]UserStub`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DojoGroupPrefetch) GetUsersOk() (*map[string]UserStub, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DojoGroupPrefetch) SetUsers(v map[string]UserStub)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DojoGroupPrefetch) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


