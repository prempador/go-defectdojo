# ToolConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**AuthenticationType** | Pointer to **NullableString** | * &#x60;API&#x60; - API Key * &#x60;Password&#x60; - Username/Password * &#x60;SSH&#x60; - SSH | [optional] 
**Extras** | Pointer to **NullableString** | Additional definitions that will be consumed by scanner | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**AuthTitle** | Pointer to **NullableString** |  | [optional] 
**Ssh** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**ToolType** | **int32** |  | 

## Methods

### NewToolConfigurationRequest

`func NewToolConfigurationRequest(name string, toolType int32, ) *ToolConfigurationRequest`

NewToolConfigurationRequest instantiates a new ToolConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolConfigurationRequestWithDefaults

`func NewToolConfigurationRequestWithDefaults() *ToolConfigurationRequest`

NewToolConfigurationRequestWithDefaults instantiates a new ToolConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolConfigurationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolConfigurationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolConfigurationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolConfigurationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ToolConfigurationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ToolConfigurationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *ToolConfigurationRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolConfigurationRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolConfigurationRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolConfigurationRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ToolConfigurationRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ToolConfigurationRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAuthenticationType

`func (o *ToolConfigurationRequest) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ToolConfigurationRequest) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ToolConfigurationRequest) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *ToolConfigurationRequest) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *ToolConfigurationRequest) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *ToolConfigurationRequest) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetExtras

`func (o *ToolConfigurationRequest) GetExtras() string`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ToolConfigurationRequest) GetExtrasOk() (*string, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ToolConfigurationRequest) SetExtras(v string)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ToolConfigurationRequest) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *ToolConfigurationRequest) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *ToolConfigurationRequest) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetUsername

`func (o *ToolConfigurationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ToolConfigurationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ToolConfigurationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ToolConfigurationRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ToolConfigurationRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ToolConfigurationRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ToolConfigurationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ToolConfigurationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ToolConfigurationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ToolConfigurationRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ToolConfigurationRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ToolConfigurationRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetAuthTitle

`func (o *ToolConfigurationRequest) GetAuthTitle() string`

GetAuthTitle returns the AuthTitle field if non-nil, zero value otherwise.

### GetAuthTitleOk

`func (o *ToolConfigurationRequest) GetAuthTitleOk() (*string, bool)`

GetAuthTitleOk returns a tuple with the AuthTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTitle

`func (o *ToolConfigurationRequest) SetAuthTitle(v string)`

SetAuthTitle sets AuthTitle field to given value.

### HasAuthTitle

`func (o *ToolConfigurationRequest) HasAuthTitle() bool`

HasAuthTitle returns a boolean if a field has been set.

### SetAuthTitleNil

`func (o *ToolConfigurationRequest) SetAuthTitleNil(b bool)`

 SetAuthTitleNil sets the value for AuthTitle to be an explicit nil

### UnsetAuthTitle
`func (o *ToolConfigurationRequest) UnsetAuthTitle()`

UnsetAuthTitle ensures that no value is present for AuthTitle, not even an explicit nil
### GetSsh

`func (o *ToolConfigurationRequest) GetSsh() string`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *ToolConfigurationRequest) GetSshOk() (*string, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *ToolConfigurationRequest) SetSsh(v string)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *ToolConfigurationRequest) HasSsh() bool`

HasSsh returns a boolean if a field has been set.

### SetSshNil

`func (o *ToolConfigurationRequest) SetSshNil(b bool)`

 SetSshNil sets the value for Ssh to be an explicit nil

### UnsetSsh
`func (o *ToolConfigurationRequest) UnsetSsh()`

UnsetSsh ensures that no value is present for Ssh, not even an explicit nil
### GetApiKey

`func (o *ToolConfigurationRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ToolConfigurationRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ToolConfigurationRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ToolConfigurationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *ToolConfigurationRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *ToolConfigurationRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetToolType

`func (o *ToolConfigurationRequest) GetToolType() int32`

GetToolType returns the ToolType field if non-nil, zero value otherwise.

### GetToolTypeOk

`func (o *ToolConfigurationRequest) GetToolTypeOk() (*int32, bool)`

GetToolTypeOk returns a tuple with the ToolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolType

`func (o *ToolConfigurationRequest) SetToolType(v int32)`

SetToolType sets ToolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


