# GlobalRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Role** | Pointer to **NullableInt32** | The global role will be applied to all product types and products. | [optional] 
**Prefetch** | Pointer to [**DojoGroupMemberPrefetch**](DojoGroupMemberPrefetch.md) |  | [optional] 

## Methods

### NewGlobalRole

`func NewGlobalRole(id int32, ) *GlobalRole`

NewGlobalRole instantiates a new GlobalRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalRoleWithDefaults

`func NewGlobalRoleWithDefaults() *GlobalRole`

NewGlobalRoleWithDefaults instantiates a new GlobalRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalRole) SetId(v int32)`

SetId sets Id field to given value.


### GetUser

`func (o *GlobalRole) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GlobalRole) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GlobalRole) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *GlobalRole) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GlobalRole) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GlobalRole) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGroup

`func (o *GlobalRole) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GlobalRole) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GlobalRole) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GlobalRole) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *GlobalRole) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *GlobalRole) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetRole

`func (o *GlobalRole) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GlobalRole) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GlobalRole) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *GlobalRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *GlobalRole) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *GlobalRole) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPrefetch

`func (o *GlobalRole) GetPrefetch() DojoGroupMemberPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *GlobalRole) GetPrefetchOk() (*DojoGroupMemberPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *GlobalRole) SetPrefetch(v DojoGroupMemberPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *GlobalRole) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


