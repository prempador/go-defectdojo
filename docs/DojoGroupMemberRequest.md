# DojoGroupMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **int32** |  | 
**User** | **int32** |  | 
**Role** | **int32** | This role determines the permissions of the user to manage the group. | 

## Methods

### NewDojoGroupMemberRequest

`func NewDojoGroupMemberRequest(group int32, user int32, role int32, ) *DojoGroupMemberRequest`

NewDojoGroupMemberRequest instantiates a new DojoGroupMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDojoGroupMemberRequestWithDefaults

`func NewDojoGroupMemberRequestWithDefaults() *DojoGroupMemberRequest`

NewDojoGroupMemberRequestWithDefaults instantiates a new DojoGroupMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DojoGroupMemberRequest) GetGroup() int32`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DojoGroupMemberRequest) GetGroupOk() (*int32, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DojoGroupMemberRequest) SetGroup(v int32)`

SetGroup sets Group field to given value.


### GetUser

`func (o *DojoGroupMemberRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DojoGroupMemberRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DojoGroupMemberRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetRole

`func (o *DojoGroupMemberRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DojoGroupMemberRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DojoGroupMemberRequest) SetRole(v int32)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


