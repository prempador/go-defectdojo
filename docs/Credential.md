# Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
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
**Notes** | **[]int32** |  | [readonly] 

## Methods

### NewCredential

`func NewCredential(id int32, name string, username string, role string, url string, environment int32, notes []int32, ) *Credential`

NewCredential instantiates a new Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialWithDefaults

`func NewCredentialWithDefaults() *Credential`

NewCredentialWithDefaults instantiates a new Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Credential) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credential) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credential) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Credential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Credential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Credential) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *Credential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Credential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Credential) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRole

`func (o *Credential) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Credential) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Credential) SetRole(v string)`

SetRole sets Role field to given value.


### GetAuthentication

`func (o *Credential) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *Credential) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *Credential) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *Credential) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetHttpAuthentication

`func (o *Credential) GetHttpAuthentication() string`

GetHttpAuthentication returns the HttpAuthentication field if non-nil, zero value otherwise.

### GetHttpAuthenticationOk

`func (o *Credential) GetHttpAuthenticationOk() (*string, bool)`

GetHttpAuthenticationOk returns a tuple with the HttpAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthentication

`func (o *Credential) SetHttpAuthentication(v string)`

SetHttpAuthentication sets HttpAuthentication field to given value.

### HasHttpAuthentication

`func (o *Credential) HasHttpAuthentication() bool`

HasHttpAuthentication returns a boolean if a field has been set.

### SetHttpAuthenticationNil

`func (o *Credential) SetHttpAuthenticationNil(b bool)`

 SetHttpAuthenticationNil sets the value for HttpAuthentication to be an explicit nil

### UnsetHttpAuthentication
`func (o *Credential) UnsetHttpAuthentication()`

UnsetHttpAuthentication ensures that no value is present for HttpAuthentication, not even an explicit nil
### GetDescription

`func (o *Credential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Credential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Credential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Credential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Credential) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Credential) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *Credential) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Credential) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Credential) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLoginRegex

`func (o *Credential) GetLoginRegex() string`

GetLoginRegex returns the LoginRegex field if non-nil, zero value otherwise.

### GetLoginRegexOk

`func (o *Credential) GetLoginRegexOk() (*string, bool)`

GetLoginRegexOk returns a tuple with the LoginRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRegex

`func (o *Credential) SetLoginRegex(v string)`

SetLoginRegex sets LoginRegex field to given value.

### HasLoginRegex

`func (o *Credential) HasLoginRegex() bool`

HasLoginRegex returns a boolean if a field has been set.

### SetLoginRegexNil

`func (o *Credential) SetLoginRegexNil(b bool)`

 SetLoginRegexNil sets the value for LoginRegex to be an explicit nil

### UnsetLoginRegex
`func (o *Credential) UnsetLoginRegex()`

UnsetLoginRegex ensures that no value is present for LoginRegex, not even an explicit nil
### GetLogoutRegex

`func (o *Credential) GetLogoutRegex() string`

GetLogoutRegex returns the LogoutRegex field if non-nil, zero value otherwise.

### GetLogoutRegexOk

`func (o *Credential) GetLogoutRegexOk() (*string, bool)`

GetLogoutRegexOk returns a tuple with the LogoutRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRegex

`func (o *Credential) SetLogoutRegex(v string)`

SetLogoutRegex sets LogoutRegex field to given value.

### HasLogoutRegex

`func (o *Credential) HasLogoutRegex() bool`

HasLogoutRegex returns a boolean if a field has been set.

### SetLogoutRegexNil

`func (o *Credential) SetLogoutRegexNil(b bool)`

 SetLogoutRegexNil sets the value for LogoutRegex to be an explicit nil

### UnsetLogoutRegex
`func (o *Credential) UnsetLogoutRegex()`

UnsetLogoutRegex ensures that no value is present for LogoutRegex, not even an explicit nil
### GetIsValid

`func (o *Credential) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *Credential) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *Credential) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *Credential) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetEnvironment

`func (o *Credential) GetEnvironment() int32`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Credential) GetEnvironmentOk() (*int32, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Credential) SetEnvironment(v int32)`

SetEnvironment sets Environment field to given value.


### GetNotes

`func (o *Credential) GetNotes() []int32`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Credential) GetNotesOk() (*[]int32, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Credential) SetNotes(v []int32)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


