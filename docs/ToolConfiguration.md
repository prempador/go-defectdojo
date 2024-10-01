# ToolConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**AuthenticationType** | Pointer to **NullableString** | * &#x60;API&#x60; - API Key * &#x60;Password&#x60; - Username/Password * &#x60;SSH&#x60; - SSH | [optional] 
**Extras** | Pointer to **NullableString** | Additional definitions that will be consumed by scanner | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**AuthTitle** | Pointer to **NullableString** |  | [optional] 
**ToolType** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedToolConfigurationListPrefetch**](PaginatedToolConfigurationListPrefetch.md) |  | [optional] 

## Methods

### NewToolConfiguration

`func NewToolConfiguration(id int32, name string, toolType int32, ) *ToolConfiguration`

NewToolConfiguration instantiates a new ToolConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolConfigurationWithDefaults

`func NewToolConfigurationWithDefaults() *ToolConfiguration`

NewToolConfigurationWithDefaults instantiates a new ToolConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolConfiguration) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ToolConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ToolConfiguration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ToolConfiguration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *ToolConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ToolConfiguration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ToolConfiguration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAuthenticationType

`func (o *ToolConfiguration) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *ToolConfiguration) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *ToolConfiguration) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *ToolConfiguration) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### SetAuthenticationTypeNil

`func (o *ToolConfiguration) SetAuthenticationTypeNil(b bool)`

 SetAuthenticationTypeNil sets the value for AuthenticationType to be an explicit nil

### UnsetAuthenticationType
`func (o *ToolConfiguration) UnsetAuthenticationType()`

UnsetAuthenticationType ensures that no value is present for AuthenticationType, not even an explicit nil
### GetExtras

`func (o *ToolConfiguration) GetExtras() string`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ToolConfiguration) GetExtrasOk() (*string, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ToolConfiguration) SetExtras(v string)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ToolConfiguration) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *ToolConfiguration) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *ToolConfiguration) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetUsername

`func (o *ToolConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ToolConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ToolConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ToolConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ToolConfiguration) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ToolConfiguration) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAuthTitle

`func (o *ToolConfiguration) GetAuthTitle() string`

GetAuthTitle returns the AuthTitle field if non-nil, zero value otherwise.

### GetAuthTitleOk

`func (o *ToolConfiguration) GetAuthTitleOk() (*string, bool)`

GetAuthTitleOk returns a tuple with the AuthTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTitle

`func (o *ToolConfiguration) SetAuthTitle(v string)`

SetAuthTitle sets AuthTitle field to given value.

### HasAuthTitle

`func (o *ToolConfiguration) HasAuthTitle() bool`

HasAuthTitle returns a boolean if a field has been set.

### SetAuthTitleNil

`func (o *ToolConfiguration) SetAuthTitleNil(b bool)`

 SetAuthTitleNil sets the value for AuthTitle to be an explicit nil

### UnsetAuthTitle
`func (o *ToolConfiguration) UnsetAuthTitle()`

UnsetAuthTitle ensures that no value is present for AuthTitle, not even an explicit nil
### GetToolType

`func (o *ToolConfiguration) GetToolType() int32`

GetToolType returns the ToolType field if non-nil, zero value otherwise.

### GetToolTypeOk

`func (o *ToolConfiguration) GetToolTypeOk() (*int32, bool)`

GetToolTypeOk returns a tuple with the ToolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolType

`func (o *ToolConfiguration) SetToolType(v int32)`

SetToolType sets ToolType field to given value.


### GetPrefetch

`func (o *ToolConfiguration) GetPrefetch() PaginatedToolConfigurationListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ToolConfiguration) GetPrefetchOk() (*PaginatedToolConfigurationListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ToolConfiguration) SetPrefetch(v PaginatedToolConfigurationListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ToolConfiguration) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


