# UserContactInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewUserContactInfoRequest

`func NewUserContactInfoRequest(user int32, ) *UserContactInfoRequest`

NewUserContactInfoRequest instantiates a new UserContactInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserContactInfoRequestWithDefaults

`func NewUserContactInfoRequestWithDefaults() *UserContactInfoRequest`

NewUserContactInfoRequestWithDefaults instantiates a new UserContactInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UserContactInfoRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserContactInfoRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserContactInfoRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserContactInfoRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UserContactInfoRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UserContactInfoRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPhoneNumber

`func (o *UserContactInfoRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserContactInfoRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserContactInfoRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserContactInfoRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCellNumber

`func (o *UserContactInfoRequest) GetCellNumber() string`

GetCellNumber returns the CellNumber field if non-nil, zero value otherwise.

### GetCellNumberOk

`func (o *UserContactInfoRequest) GetCellNumberOk() (*string, bool)`

GetCellNumberOk returns a tuple with the CellNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellNumber

`func (o *UserContactInfoRequest) SetCellNumber(v string)`

SetCellNumber sets CellNumber field to given value.

### HasCellNumber

`func (o *UserContactInfoRequest) HasCellNumber() bool`

HasCellNumber returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *UserContactInfoRequest) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *UserContactInfoRequest) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *UserContactInfoRequest) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *UserContactInfoRequest) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *UserContactInfoRequest) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *UserContactInfoRequest) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetGithubUsername

`func (o *UserContactInfoRequest) GetGithubUsername() string`

GetGithubUsername returns the GithubUsername field if non-nil, zero value otherwise.

### GetGithubUsernameOk

`func (o *UserContactInfoRequest) GetGithubUsernameOk() (*string, bool)`

GetGithubUsernameOk returns a tuple with the GithubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUsername

`func (o *UserContactInfoRequest) SetGithubUsername(v string)`

SetGithubUsername sets GithubUsername field to given value.

### HasGithubUsername

`func (o *UserContactInfoRequest) HasGithubUsername() bool`

HasGithubUsername returns a boolean if a field has been set.

### SetGithubUsernameNil

`func (o *UserContactInfoRequest) SetGithubUsernameNil(b bool)`

 SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil

### UnsetGithubUsername
`func (o *UserContactInfoRequest) UnsetGithubUsername()`

UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
### GetSlackUsername

`func (o *UserContactInfoRequest) GetSlackUsername() string`

GetSlackUsername returns the SlackUsername field if non-nil, zero value otherwise.

### GetSlackUsernameOk

`func (o *UserContactInfoRequest) GetSlackUsernameOk() (*string, bool)`

GetSlackUsernameOk returns a tuple with the SlackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUsername

`func (o *UserContactInfoRequest) SetSlackUsername(v string)`

SetSlackUsername sets SlackUsername field to given value.

### HasSlackUsername

`func (o *UserContactInfoRequest) HasSlackUsername() bool`

HasSlackUsername returns a boolean if a field has been set.

### SetSlackUsernameNil

`func (o *UserContactInfoRequest) SetSlackUsernameNil(b bool)`

 SetSlackUsernameNil sets the value for SlackUsername to be an explicit nil

### UnsetSlackUsername
`func (o *UserContactInfoRequest) UnsetSlackUsername()`

UnsetSlackUsername ensures that no value is present for SlackUsername, not even an explicit nil
### GetSlackUserId

`func (o *UserContactInfoRequest) GetSlackUserId() string`

GetSlackUserId returns the SlackUserId field if non-nil, zero value otherwise.

### GetSlackUserIdOk

`func (o *UserContactInfoRequest) GetSlackUserIdOk() (*string, bool)`

GetSlackUserIdOk returns a tuple with the SlackUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUserId

`func (o *UserContactInfoRequest) SetSlackUserId(v string)`

SetSlackUserId sets SlackUserId field to given value.

### HasSlackUserId

`func (o *UserContactInfoRequest) HasSlackUserId() bool`

HasSlackUserId returns a boolean if a field has been set.

### SetSlackUserIdNil

`func (o *UserContactInfoRequest) SetSlackUserIdNil(b bool)`

 SetSlackUserIdNil sets the value for SlackUserId to be an explicit nil

### UnsetSlackUserId
`func (o *UserContactInfoRequest) UnsetSlackUserId()`

UnsetSlackUserId ensures that no value is present for SlackUserId, not even an explicit nil
### GetBlockExecution

`func (o *UserContactInfoRequest) GetBlockExecution() bool`

GetBlockExecution returns the BlockExecution field if non-nil, zero value otherwise.

### GetBlockExecutionOk

`func (o *UserContactInfoRequest) GetBlockExecutionOk() (*bool, bool)`

GetBlockExecutionOk returns a tuple with the BlockExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockExecution

`func (o *UserContactInfoRequest) SetBlockExecution(v bool)`

SetBlockExecution sets BlockExecution field to given value.

### HasBlockExecution

`func (o *UserContactInfoRequest) HasBlockExecution() bool`

HasBlockExecution returns a boolean if a field has been set.

### GetForcePasswordReset

`func (o *UserContactInfoRequest) GetForcePasswordReset() bool`

GetForcePasswordReset returns the ForcePasswordReset field if non-nil, zero value otherwise.

### GetForcePasswordResetOk

`func (o *UserContactInfoRequest) GetForcePasswordResetOk() (*bool, bool)`

GetForcePasswordResetOk returns a tuple with the ForcePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcePasswordReset

`func (o *UserContactInfoRequest) SetForcePasswordReset(v bool)`

SetForcePasswordReset sets ForcePasswordReset field to given value.

### HasForcePasswordReset

`func (o *UserContactInfoRequest) HasForcePasswordReset() bool`

HasForcePasswordReset returns a boolean if a field has been set.

### GetUser

`func (o *UserContactInfoRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserContactInfoRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserContactInfoRequest) SetUser(v int32)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


