# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**User**](User.md) |  | 
**UserContactInfo** | [**UserContactInfo**](UserContactInfo.md) |  | 
**GlobalRole** | [**GlobalRole**](GlobalRole.md) |  | 
**DojoGroupMember** | [**[]DojoGroupMember**](DojoGroupMember.md) |  | 
**ProductTypeMember** | [**[]ProductTypeMember**](ProductTypeMember.md) |  | 
**ProductMember** | [**[]ProductMember**](ProductMember.md) |  | 

## Methods

### NewUserProfile

`func NewUserProfile(user User, userContactInfo UserContactInfo, globalRole GlobalRole, dojoGroupMember []DojoGroupMember, productTypeMember []ProductTypeMember, productMember []ProductMember, ) *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserProfile) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserProfile) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserProfile) SetUser(v User)`

SetUser sets User field to given value.


### GetUserContactInfo

`func (o *UserProfile) GetUserContactInfo() UserContactInfo`

GetUserContactInfo returns the UserContactInfo field if non-nil, zero value otherwise.

### GetUserContactInfoOk

`func (o *UserProfile) GetUserContactInfoOk() (*UserContactInfo, bool)`

GetUserContactInfoOk returns a tuple with the UserContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserContactInfo

`func (o *UserProfile) SetUserContactInfo(v UserContactInfo)`

SetUserContactInfo sets UserContactInfo field to given value.


### GetGlobalRole

`func (o *UserProfile) GetGlobalRole() GlobalRole`

GetGlobalRole returns the GlobalRole field if non-nil, zero value otherwise.

### GetGlobalRoleOk

`func (o *UserProfile) GetGlobalRoleOk() (*GlobalRole, bool)`

GetGlobalRoleOk returns a tuple with the GlobalRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRole

`func (o *UserProfile) SetGlobalRole(v GlobalRole)`

SetGlobalRole sets GlobalRole field to given value.


### GetDojoGroupMember

`func (o *UserProfile) GetDojoGroupMember() []DojoGroupMember`

GetDojoGroupMember returns the DojoGroupMember field if non-nil, zero value otherwise.

### GetDojoGroupMemberOk

`func (o *UserProfile) GetDojoGroupMemberOk() (*[]DojoGroupMember, bool)`

GetDojoGroupMemberOk returns a tuple with the DojoGroupMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDojoGroupMember

`func (o *UserProfile) SetDojoGroupMember(v []DojoGroupMember)`

SetDojoGroupMember sets DojoGroupMember field to given value.


### GetProductTypeMember

`func (o *UserProfile) GetProductTypeMember() []ProductTypeMember`

GetProductTypeMember returns the ProductTypeMember field if non-nil, zero value otherwise.

### GetProductTypeMemberOk

`func (o *UserProfile) GetProductTypeMemberOk() (*[]ProductTypeMember, bool)`

GetProductTypeMemberOk returns a tuple with the ProductTypeMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeMember

`func (o *UserProfile) SetProductTypeMember(v []ProductTypeMember)`

SetProductTypeMember sets ProductTypeMember field to given value.


### GetProductMember

`func (o *UserProfile) GetProductMember() []ProductMember`

GetProductMember returns the ProductMember field if non-nil, zero value otherwise.

### GetProductMemberOk

`func (o *UserProfile) GetProductMemberOk() (*[]ProductMember, bool)`

GetProductMemberOk returns a tuple with the ProductMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMember

`func (o *UserProfile) SetProductMember(v []ProductMember)`

SetProductMember sets ProductMember field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


