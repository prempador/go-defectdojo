# CredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Username** | **string** |  | 
**Role** | **string** |  | 
**Authentication** | Pointer to **string** | * &#x60;Form&#x60; - Form Authentication * &#x60;SSO&#x60; - SSO Redirect | [optional] 
**HttpAuthentication** | Pointer to **NullableString** | * &#x60;Basic&#x60; - Basic * &#x60;NTLM&#x60; - NTLM | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**LoginRegex** | Pointer to **NullableString** |  | [optional] 
**LogoutRegex** | Pointer to **NullableString** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Environment** | **int32** |  | 

## Methods

### NewCredentialRequest

`func NewCredentialRequest(name string, username string, role string, url string, environment int32, ) *CredentialRequest`

NewCredentialRequest instantiates a new CredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialRequestWithDefaults

`func NewCredentialRequestWithDefaults() *CredentialRequest`

NewCredentialRequestWithDefaults instantiates a new CredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CredentialRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *CredentialRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRole

`func (o *CredentialRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CredentialRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CredentialRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetAuthentication

`func (o *CredentialRequest) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *CredentialRequest) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *CredentialRequest) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *CredentialRequest) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetHttpAuthentication

`func (o *CredentialRequest) GetHttpAuthentication() string`

GetHttpAuthentication returns the HttpAuthentication field if non-nil, zero value otherwise.

### GetHttpAuthenticationOk

`func (o *CredentialRequest) GetHttpAuthenticationOk() (*string, bool)`

GetHttpAuthenticationOk returns a tuple with the HttpAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthentication

`func (o *CredentialRequest) SetHttpAuthentication(v string)`

SetHttpAuthentication sets HttpAuthentication field to given value.

### HasHttpAuthentication

`func (o *CredentialRequest) HasHttpAuthentication() bool`

HasHttpAuthentication returns a boolean if a field has been set.

### SetHttpAuthenticationNil

`func (o *CredentialRequest) SetHttpAuthenticationNil(b bool)`

 SetHttpAuthenticationNil sets the value for HttpAuthentication to be an explicit nil

### UnsetHttpAuthentication
`func (o *CredentialRequest) UnsetHttpAuthentication()`

UnsetHttpAuthentication ensures that no value is present for HttpAuthentication, not even an explicit nil
### GetDescription

`func (o *CredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CredentialRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CredentialRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *CredentialRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CredentialRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CredentialRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLoginRegex

`func (o *CredentialRequest) GetLoginRegex() string`

GetLoginRegex returns the LoginRegex field if non-nil, zero value otherwise.

### GetLoginRegexOk

`func (o *CredentialRequest) GetLoginRegexOk() (*string, bool)`

GetLoginRegexOk returns a tuple with the LoginRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRegex

`func (o *CredentialRequest) SetLoginRegex(v string)`

SetLoginRegex sets LoginRegex field to given value.

### HasLoginRegex

`func (o *CredentialRequest) HasLoginRegex() bool`

HasLoginRegex returns a boolean if a field has been set.

### SetLoginRegexNil

`func (o *CredentialRequest) SetLoginRegexNil(b bool)`

 SetLoginRegexNil sets the value for LoginRegex to be an explicit nil

### UnsetLoginRegex
`func (o *CredentialRequest) UnsetLoginRegex()`

UnsetLoginRegex ensures that no value is present for LoginRegex, not even an explicit nil
### GetLogoutRegex

`func (o *CredentialRequest) GetLogoutRegex() string`

GetLogoutRegex returns the LogoutRegex field if non-nil, zero value otherwise.

### GetLogoutRegexOk

`func (o *CredentialRequest) GetLogoutRegexOk() (*string, bool)`

GetLogoutRegexOk returns a tuple with the LogoutRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRegex

`func (o *CredentialRequest) SetLogoutRegex(v string)`

SetLogoutRegex sets LogoutRegex field to given value.

### HasLogoutRegex

`func (o *CredentialRequest) HasLogoutRegex() bool`

HasLogoutRegex returns a boolean if a field has been set.

### SetLogoutRegexNil

`func (o *CredentialRequest) SetLogoutRegexNil(b bool)`

 SetLogoutRegexNil sets the value for LogoutRegex to be an explicit nil

### UnsetLogoutRegex
`func (o *CredentialRequest) UnsetLogoutRegex()`

UnsetLogoutRegex ensures that no value is present for LogoutRegex, not even an explicit nil
### GetIsValid

`func (o *CredentialRequest) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *CredentialRequest) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *CredentialRequest) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *CredentialRequest) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialRequest) GetEnvironment() int32`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialRequest) GetEnvironmentOk() (*int32, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialRequest) SetEnvironment(v int32)`

SetEnvironment sets Environment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


