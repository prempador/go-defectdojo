# PatchedCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** | * &#x60;Form&#x60; - Form Authentication * &#x60;SSO&#x60; - SSO Redirect | [optional] 
**HttpAuthentication** | Pointer to **NullableString** | * &#x60;Basic&#x60; - Basic * &#x60;NTLM&#x60; - NTLM | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**LoginRegex** | Pointer to **NullableString** |  | [optional] 
**LogoutRegex** | Pointer to **NullableString** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Environment** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedCredentialRequest

`func NewPatchedCredentialRequest() *PatchedCredentialRequest`

NewPatchedCredentialRequest instantiates a new PatchedCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCredentialRequestWithDefaults

`func NewPatchedCredentialRequestWithDefaults() *PatchedCredentialRequest`

NewPatchedCredentialRequestWithDefaults instantiates a new PatchedCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedCredentialRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCredentialRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCredentialRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCredentialRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedCredentialRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedCredentialRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedCredentialRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedCredentialRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRole

`func (o *PatchedCredentialRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedCredentialRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedCredentialRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedCredentialRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAuthentication

`func (o *PatchedCredentialRequest) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *PatchedCredentialRequest) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *PatchedCredentialRequest) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *PatchedCredentialRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetHttpAuthentication

`func (o *PatchedCredentialRequest) GetHttpAuthentication() string`

GetHttpAuthentication returns the HttpAuthentication field if non-nil, zero value otherwise.

### GetHttpAuthenticationOk

`func (o *PatchedCredentialRequest) GetHttpAuthenticationOk() (*string, bool)`

GetHttpAuthenticationOk returns a tuple with the HttpAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthentication

`func (o *PatchedCredentialRequest) SetHttpAuthentication(v string)`

SetHttpAuthentication sets HttpAuthentication field to given value.

### HasHttpAuthentication

`func (o *PatchedCredentialRequest) HasHttpAuthentication() bool`

HasHttpAuthentication returns a boolean if a field has been set.

### SetHttpAuthenticationNil

`func (o *PatchedCredentialRequest) SetHttpAuthenticationNil(b bool)`

 SetHttpAuthenticationNil sets the value for HttpAuthentication to be an explicit nil

### UnsetHttpAuthentication
`func (o *PatchedCredentialRequest) UnsetHttpAuthentication()`

UnsetHttpAuthentication ensures that no value is present for HttpAuthentication, not even an explicit nil
### GetDescription

`func (o *PatchedCredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCredentialRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedCredentialRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedCredentialRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *PatchedCredentialRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedCredentialRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedCredentialRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedCredentialRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLoginRegex

`func (o *PatchedCredentialRequest) GetLoginRegex() string`

GetLoginRegex returns the LoginRegex field if non-nil, zero value otherwise.

### GetLoginRegexOk

`func (o *PatchedCredentialRequest) GetLoginRegexOk() (*string, bool)`

GetLoginRegexOk returns a tuple with the LoginRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRegex

`func (o *PatchedCredentialRequest) SetLoginRegex(v string)`

SetLoginRegex sets LoginRegex field to given value.

### HasLoginRegex

`func (o *PatchedCredentialRequest) HasLoginRegex() bool`

HasLoginRegex returns a boolean if a field has been set.

### SetLoginRegexNil

`func (o *PatchedCredentialRequest) SetLoginRegexNil(b bool)`

 SetLoginRegexNil sets the value for LoginRegex to be an explicit nil

### UnsetLoginRegex
`func (o *PatchedCredentialRequest) UnsetLoginRegex()`

UnsetLoginRegex ensures that no value is present for LoginRegex, not even an explicit nil
### GetLogoutRegex

`func (o *PatchedCredentialRequest) GetLogoutRegex() string`

GetLogoutRegex returns the LogoutRegex field if non-nil, zero value otherwise.

### GetLogoutRegexOk

`func (o *PatchedCredentialRequest) GetLogoutRegexOk() (*string, bool)`

GetLogoutRegexOk returns a tuple with the LogoutRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRegex

`func (o *PatchedCredentialRequest) SetLogoutRegex(v string)`

SetLogoutRegex sets LogoutRegex field to given value.

### HasLogoutRegex

`func (o *PatchedCredentialRequest) HasLogoutRegex() bool`

HasLogoutRegex returns a boolean if a field has been set.

### SetLogoutRegexNil

`func (o *PatchedCredentialRequest) SetLogoutRegexNil(b bool)`

 SetLogoutRegexNil sets the value for LogoutRegex to be an explicit nil

### UnsetLogoutRegex
`func (o *PatchedCredentialRequest) UnsetLogoutRegex()`

UnsetLogoutRegex ensures that no value is present for LogoutRegex, not even an explicit nil
### GetIsValid

`func (o *PatchedCredentialRequest) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *PatchedCredentialRequest) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *PatchedCredentialRequest) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *PatchedCredentialRequest) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetEnvironment

`func (o *PatchedCredentialRequest) GetEnvironment() int32`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PatchedCredentialRequest) GetEnvironmentOk() (*int32, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PatchedCredentialRequest) SetEnvironment(v int32)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PatchedCredentialRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


