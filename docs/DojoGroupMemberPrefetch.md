# DojoGroupMemberPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**map[string]DojoGroup**](DojoGroup.md) |  | [optional] [readonly] 
**Role** | Pointer to [**map[string]Role**](Role.md) |  | [optional] [readonly] 
**User** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewDojoGroupMemberPrefetch

`func NewDojoGroupMemberPrefetch() *DojoGroupMemberPrefetch`

NewDojoGroupMemberPrefetch instantiates a new DojoGroupMemberPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDojoGroupMemberPrefetchWithDefaults

`func NewDojoGroupMemberPrefetchWithDefaults() *DojoGroupMemberPrefetch`

NewDojoGroupMemberPrefetchWithDefaults instantiates a new DojoGroupMemberPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DojoGroupMemberPrefetch) GetGroup() map[string]DojoGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DojoGroupMemberPrefetch) GetGroupOk() (*map[string]DojoGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DojoGroupMemberPrefetch) SetGroup(v map[string]DojoGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DojoGroupMemberPrefetch) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetRole

`func (o *DojoGroupMemberPrefetch) GetRole() map[string]Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DojoGroupMemberPrefetch) GetRoleOk() (*map[string]Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DojoGroupMemberPrefetch) SetRole(v map[string]Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *DojoGroupMemberPrefetch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *DojoGroupMemberPrefetch) GetUser() map[string]UserStub`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DojoGroupMemberPrefetch) GetUserOk() (*map[string]UserStub, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DojoGroupMemberPrefetch) SetUser(v map[string]UserStub)`

SetUser sets User field to given value.

### HasUser

`func (o *DojoGroupMemberPrefetch) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


