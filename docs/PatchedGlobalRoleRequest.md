# PatchedGlobalRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableInt32** |  | [optional] 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Role** | Pointer to **NullableInt32** | The global role will be applied to all product types and products. | [optional] 

## Methods

### NewPatchedGlobalRoleRequest

`func NewPatchedGlobalRoleRequest() *PatchedGlobalRoleRequest`

NewPatchedGlobalRoleRequest instantiates a new PatchedGlobalRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedGlobalRoleRequestWithDefaults

`func NewPatchedGlobalRoleRequestWithDefaults() *PatchedGlobalRoleRequest`

NewPatchedGlobalRoleRequestWithDefaults instantiates a new PatchedGlobalRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PatchedGlobalRoleRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedGlobalRoleRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedGlobalRoleRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedGlobalRoleRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedGlobalRoleRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedGlobalRoleRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGroup

`func (o *PatchedGlobalRoleRequest) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedGlobalRoleRequest) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedGlobalRoleRequest) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedGlobalRoleRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedGlobalRoleRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedGlobalRoleRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetRole

`func (o *PatchedGlobalRoleRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedGlobalRoleRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedGlobalRoleRequest) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedGlobalRoleRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedGlobalRoleRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedGlobalRoleRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


