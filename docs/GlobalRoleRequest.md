# GlobalRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableInt32** |  | [optional] 
**Group** | Pointer to **NullableInt32** |  | [optional] 
**Role** | Pointer to **NullableInt32** | The global role will be applied to all product types and products. | [optional] 

## Methods

### NewGlobalRoleRequest

`func NewGlobalRoleRequest() *GlobalRoleRequest`

NewGlobalRoleRequest instantiates a new GlobalRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalRoleRequestWithDefaults

`func NewGlobalRoleRequestWithDefaults() *GlobalRoleRequest`

NewGlobalRoleRequestWithDefaults instantiates a new GlobalRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GlobalRoleRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GlobalRoleRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GlobalRoleRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *GlobalRoleRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GlobalRoleRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GlobalRoleRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetGroup

`func (o *GlobalRoleRequest) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GlobalRoleRequest) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GlobalRoleRequest) SetGroup(v int32)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GlobalRoleRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *GlobalRoleRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *GlobalRoleRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetRole

`func (o *GlobalRoleRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GlobalRoleRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GlobalRoleRequest) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *GlobalRoleRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *GlobalRoleRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *GlobalRoleRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


