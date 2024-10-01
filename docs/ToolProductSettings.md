# ToolProductSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**SettingUrl** | **string** |  | 
**Product** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**ToolProjectId** | Pointer to **NullableString** |  | [optional] 
**ToolConfiguration** | **int32** |  | 
**Notes** | **[]int32** |  | [readonly] 
**Prefetch** | Pointer to [**PaginatedToolProductSettingsListPrefetch**](PaginatedToolProductSettingsListPrefetch.md) |  | [optional] 

## Methods

### NewToolProductSettings

`func NewToolProductSettings(id int32, settingUrl string, product int32, name string, toolConfiguration int32, notes []int32, ) *ToolProductSettings`

NewToolProductSettings instantiates a new ToolProductSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolProductSettingsWithDefaults

`func NewToolProductSettingsWithDefaults() *ToolProductSettings`

NewToolProductSettingsWithDefaults instantiates a new ToolProductSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ToolProductSettings) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolProductSettings) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolProductSettings) SetId(v int32)`

SetId sets Id field to given value.


### GetSettingUrl

`func (o *ToolProductSettings) GetSettingUrl() string`

GetSettingUrl returns the SettingUrl field if non-nil, zero value otherwise.

### GetSettingUrlOk

`func (o *ToolProductSettings) GetSettingUrlOk() (*string, bool)`

GetSettingUrlOk returns a tuple with the SettingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingUrl

`func (o *ToolProductSettings) SetSettingUrl(v string)`

SetSettingUrl sets SettingUrl field to given value.


### GetProduct

`func (o *ToolProductSettings) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ToolProductSettings) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ToolProductSettings) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetName

`func (o *ToolProductSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolProductSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolProductSettings) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolProductSettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolProductSettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolProductSettings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolProductSettings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ToolProductSettings) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ToolProductSettings) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *ToolProductSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolProductSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolProductSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolProductSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ToolProductSettings) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ToolProductSettings) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetToolProjectId

`func (o *ToolProductSettings) GetToolProjectId() string`

GetToolProjectId returns the ToolProjectId field if non-nil, zero value otherwise.

### GetToolProjectIdOk

`func (o *ToolProductSettings) GetToolProjectIdOk() (*string, bool)`

GetToolProjectIdOk returns a tuple with the ToolProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolProjectId

`func (o *ToolProductSettings) SetToolProjectId(v string)`

SetToolProjectId sets ToolProjectId field to given value.

### HasToolProjectId

`func (o *ToolProductSettings) HasToolProjectId() bool`

HasToolProjectId returns a boolean if a field has been set.

### SetToolProjectIdNil

`func (o *ToolProductSettings) SetToolProjectIdNil(b bool)`

 SetToolProjectIdNil sets the value for ToolProjectId to be an explicit nil

### UnsetToolProjectId
`func (o *ToolProductSettings) UnsetToolProjectId()`

UnsetToolProjectId ensures that no value is present for ToolProjectId, not even an explicit nil
### GetToolConfiguration

`func (o *ToolProductSettings) GetToolConfiguration() int32`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *ToolProductSettings) GetToolConfigurationOk() (*int32, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *ToolProductSettings) SetToolConfiguration(v int32)`

SetToolConfiguration sets ToolConfiguration field to given value.


### GetNotes

`func (o *ToolProductSettings) GetNotes() []int32`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ToolProductSettings) GetNotesOk() (*[]int32, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ToolProductSettings) SetNotes(v []int32)`

SetNotes sets Notes field to given value.


### GetPrefetch

`func (o *ToolProductSettings) GetPrefetch() PaginatedToolProductSettingsListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ToolProductSettings) GetPrefetchOk() (*PaginatedToolProductSettingsListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ToolProductSettings) SetPrefetch(v PaginatedToolProductSettingsListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ToolProductSettings) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


