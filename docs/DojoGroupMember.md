# DojoGroupMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Group** | **int32** |  | 
**User** | **int32** |  | 
**Role** | **int32** | This role determines the permissions of the user to manage the group. | 
**Prefetch** | Pointer to [**DojoGroupMemberPrefetch**](DojoGroupMemberPrefetch.md) |  | [optional] 

## Methods

### NewDojoGroupMember

`func NewDojoGroupMember(id int32, group int32, user int32, role int32, ) *DojoGroupMember`

NewDojoGroupMember instantiates a new DojoGroupMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDojoGroupMemberWithDefaults

`func NewDojoGroupMemberWithDefaults() *DojoGroupMember`

NewDojoGroupMemberWithDefaults instantiates a new DojoGroupMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DojoGroupMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DojoGroupMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DojoGroupMember) SetId(v int32)`

SetId sets Id field to given value.


### GetGroup

`func (o *DojoGroupMember) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DojoGroupMember) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DojoGroupMember) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetUser

`func (o *DojoGroupMember) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DojoGroupMember) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DojoGroupMember) SetUser(v int32)`

SetUser sets User field to given value.


### GetRole

`func (o *DojoGroupMember) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DojoGroupMember) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DojoGroupMember) SetRole(v int32)`

SetRole sets Role field to given value.


### GetPrefetch

`func (o *DojoGroupMember) GetPrefetch() DojoGroupMemberPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *DojoGroupMember) GetPrefetchOk() (*DojoGroupMemberPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *DojoGroupMember) SetPrefetch(v DojoGroupMemberPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *DojoGroupMember) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


