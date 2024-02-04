# PatchedToolConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**AuthenticationType** | Pointer to **NullableString** | * &#x60;API&#x60; - API Key * &#x60;Password&#x60; - Username/Password * &#x60;SSH&#x60; - SSH | [optional] 
**Extras** | Pointer to **NullableString** | Additional definitions that will be consumed by scanner | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**AuthTitle** | Pointer to **NullableString** |  | [optional] 
**Ssh** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**ToolType** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedToolConfigurationRequest

`func NewPatchedToolConfigurationRequest() *PatchedToolConfigurationRequest`

NewPatchedToolConfigurationRequest instantiates a new PatchedToolConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedToolConfigurationRequestWithDefaults

`func NewPatchedToolConfigurationRequestWithDefaults() *PatchedToolConfigurationRequest`

NewPatchedToolConfigurationRequestWithDefaults instantiates a new PatchedToolConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedToolConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedToolConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedToolConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedToolConfigurationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedToolConfigurationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedToolConfigurationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedToolConfigurationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedToolConfigurationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedToolConfigurationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedToolConfigurationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *PatchedToolConfigurationRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedToolConfigurationRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedToolConfigurationRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedToolConfigurationRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PatchedToolConfigurationRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PatchedToolConfigurationRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAuthenticationType

`func (o *PatchedToolConfigurationRequest) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *PatchedToolConfigurationRequest) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *PatchedToolConfigurationRequest) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *PatchedToolConfigurationRequest) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *PatchedToolConfigurationRequest) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *PatchedToolConfigurationRequest) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetExtras

`func (o *PatchedToolConfigurationRequest) GetExtras() string`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *PatchedToolConfigurationRequest) GetExtrasOk() (*string, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *PatchedToolConfigurationRequest) SetExtras(v string)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *PatchedToolConfigurationRequest) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *PatchedToolConfigurationRequest) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *PatchedToolConfigurationRequest) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetUsername

`func (o *PatchedToolConfigurationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedToolConfigurationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedToolConfigurationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedToolConfigurationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PatchedToolConfigurationRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PatchedToolConfigurationRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PatchedToolConfigurationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedToolConfigurationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedToolConfigurationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedToolConfigurationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchedToolConfigurationRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchedToolConfigurationRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetAuthTitle

`func (o *PatchedToolConfigurationRequest) GetAuthTitle() string`

GetAuthTitle returns the AuthTitle field if non-nil, zero value otherwise.

### GetAuthTitleOk

`func (o *PatchedToolConfigurationRequest) GetAuthTitleOk() (*string, bool)`

GetAuthTitleOk returns a tuple with the AuthTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTitle

`func (o *PatchedToolConfigurationRequest) SetAuthTitle(v string)`

SetAuthTitle sets AuthTitle field to given value.

### HasAuthTitle

`func (o *PatchedToolConfigurationRequest) HasAuthTitle() bool`

HasAuthTitle returns a boolean if a field has been set.

### SetAuthTitleNil

`func (o *PatchedToolConfigurationRequest) SetAuthTitleNil(b bool)`

 SetAuthTitleNil sets the value for AuthTitle to be an explicit nil

### UnsetAuthTitle
`func (o *PatchedToolConfigurationRequest) UnsetAuthTitle()`

UnsetAuthTitle ensures that no value is present for AuthTitle, not even an explicit nil
### GetSsh

`func (o *PatchedToolConfigurationRequest) GetSsh() string`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *PatchedToolConfigurationRequest) GetSshOk() (*string, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *PatchedToolConfigurationRequest) SetSsh(v string)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *PatchedToolConfigurationRequest) HasSsh() bool`

HasSsh returns a boolean if a field has been set.

### SetSshNil

`func (o *PatchedToolConfigurationRequest) SetSshNil(b bool)`

 SetSshNil sets the value for Ssh to be an explicit nil

### UnsetSsh
`func (o *PatchedToolConfigurationRequest) UnsetSsh()`

UnsetSsh ensures that no value is present for Ssh, not even an explicit nil
### GetApiKey

`func (o *PatchedToolConfigurationRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PatchedToolConfigurationRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PatchedToolConfigurationRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PatchedToolConfigurationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *PatchedToolConfigurationRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *PatchedToolConfigurationRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetToolType

`func (o *PatchedToolConfigurationRequest) GetToolType() int32`

GetToolType returns the ToolType field if non-nil, zero value otherwise.

### GetToolTypeOk

`func (o *PatchedToolConfigurationRequest) GetToolTypeOk() (*int32, bool)`

GetToolTypeOk returns a tuple with the ToolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolType

`func (o *PatchedToolConfigurationRequest) SetToolType(v int32)`

SetToolType sets ToolType field to given value.

### HasToolType

`func (o *PatchedToolConfigurationRequest) HasToolType() bool`

HasToolType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


