# UserContactInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Title** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number must be entered in the format: &#39;+999999999&#39;. Up to 15 digits allowed. | [optional] 
**CellNumber** | Pointer to **string** | Phone number must be entered in the format: &#39;+999999999&#39;. Up to 15 digits allowed. | [optional] 
**TwitterUsername** | Pointer to **NullableString** |  | [optional] 
**GithubUsername** | Pointer to **NullableString** |  | [optional] 
**SlackUsername** | Pointer to **NullableString** | Email address associated with your slack account | [optional] 
**SlackUserId** | Pointer to **NullableString** |  | [optional] 
**BlockExecution** | Pointer to **bool** | Instead of async deduping a finding the findings will be deduped synchronously and will &#39;block&#39; the user until completion. | [optional] 
**ForcePasswordReset** | Pointer to **bool** | Forces this user to reset their password on next login. | [optional] 
**User** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedUserContactInfoListPrefetch**](PaginatedUserContactInfoListPrefetch.md) |  | [optional] 

## Methods

### NewUserContactInfo

`func NewUserContactInfo(id int32, user int32, ) *UserContactInfo`

NewUserContactInfo instantiates a new UserContactInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserContactInfoWithDefaults

`func NewUserContactInfoWithDefaults() *UserContactInfo`

NewUserContactInfoWithDefaults instantiates a new UserContactInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserContactInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserContactInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserContactInfo) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *UserContactInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserContactInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserContactInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserContactInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UserContactInfo) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UserContactInfo) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPhoneNumber

`func (o *UserContactInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserContactInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserContactInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserContactInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCellNumber

`func (o *UserContactInfo) GetCellNumber() string`

GetCellNumber returns the CellNumber field if non-nil, zero value otherwise.

### GetCellNumberOk

`func (o *UserContactInfo) GetCellNumberOk() (*string, bool)`

GetCellNumberOk returns a tuple with the CellNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellNumber

`func (o *UserContactInfo) SetCellNumber(v string)`

SetCellNumber sets CellNumber field to given value.

### HasCellNumber

`func (o *UserContactInfo) HasCellNumber() bool`

HasCellNumber returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *UserContactInfo) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *UserContactInfo) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *UserContactInfo) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *UserContactInfo) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *UserContactInfo) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *UserContactInfo) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetGithubUsername

`func (o *UserContactInfo) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *UserContactInfo) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *UserContactInfo) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *UserContactInfo) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *UserContactInfo) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *UserContactInfo) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
### GetSlackUsername

`func (o *UserContactInfo) GetSlackUsername() string`

GetSlackUsername returns the SlackUsername field if non-nil, zero value otherwise.

### GetSlackUsernameOk

`func (o *UserContactInfo) GetSlackUsernameOk() (*string, bool)`

GetSlackUsernameOk returns a tuple with the SlackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUsername

`func (o *UserContactInfo) SetSlackUsername(v string)`

SetSlackUsername sets SlackUsername field to given value.

### HasSlackUsername

`func (o *UserContactInfo) HasSlackUsername() bool`

HasSlackUsername returns a boolean if a field has been set.

### SetSlackUsernameNil

`func (o *UserContactInfo) SetSlackUsernameNil(b bool)`

 SetSlackUsernameNil sets the value for SlackUsername to be an explicit nil

### UnsetSlackUsername
`func (o *UserContactInfo) UnsetSlackUsername()`

UnsetSlackUsername ensures that no value is present for SlackUsername, not even an explicit nil
### GetSlackUserId

`func (o *UserContactInfo) GetSlackUserId() string`

GetSlackUserId returns the SlackUserId field if non-nil, zero value otherwise.

### GetSlackUserIdOk

`func (o *UserContactInfo) GetSlackUserIdOk() (*string, bool)`

GetSlackUserIdOk returns a tuple with the SlackUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUserId

`func (o *UserContactInfo) SetSlackUserId(v string)`

SetSlackUserId sets SlackUserId field to given value.

### HasSlackUserId

`func (o *UserContactInfo) HasSlackUserId() bool`

HasSlackUserId returns a boolean if a field has been set.

### SetSlackUserIdNil

`func (o *UserContactInfo) SetSlackUserIdNil(b bool)`

 SetSlackUserIdNil sets the value for SlackUserId to be an explicit nil

### UnsetSlackUserId
`func (o *UserContactInfo) UnsetSlackUserId()`

UnsetSlackUserId ensures that no value is present for SlackUserId, not even an explicit nil
### GetBlockExecution

`func (o *UserContactInfo) GetBlockExecution() bool`

GetBlockExecution returns the BlockExecution field if non-nil, zero value otherwise.

### GetBlockExecutionOk

`func (o *UserContactInfo) GetBlockExecutionOk() (*bool, bool)`

GetBlockExecutionOk returns a tuple with the BlockExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExecution

`func (o *UserContactInfo) SetBlockExecution(v bool)`

SetBlockExecution sets BlockExecution field to given value.

### HasBlockExecution

`func (o *UserContactInfo) HasBlockExecution() bool`

HasBlockExecution returns a boolean if a field has been set.

### GetForcePasswordReset

`func (o *UserContactInfo) GetForcePasswordReset() bool`

GetForcePasswordReset returns the ForcePasswordReset field if non-nil, zero value otherwise.

### GetForcePasswordResetOk

`func (o *UserContactInfo) GetForcePasswordResetOk() (*bool, bool)`

GetForcePasswordResetOk returns a tuple with the ForcePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordReset

`func (o *UserContactInfo) SetForcePasswordReset(v bool)`

SetForcePasswordReset sets ForcePasswordReset field to given value.

### HasForcePasswordReset

`func (o *UserContactInfo) HasForcePasswordReset() bool`

HasForcePasswordReset returns a boolean if a field has been set.

### GetUser

`func (o *UserContactInfo) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserContactInfo) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserContactInfo) SetUser(v int32)`

SetUser sets User field to given value.


### GetPrefetch

`func (o *UserContactInfo) GetPrefetch() PaginatedUserContactInfoListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *UserContactInfo) GetPrefetchOk() (*PaginatedUserContactInfoListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *UserContactInfo) SetPrefetch(v PaginatedUserContactInfoListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *UserContactInfo) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


