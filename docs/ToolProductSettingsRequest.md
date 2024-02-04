# ToolProductSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingUrl** | **string** |  | 
**Product** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**ToolProjectId** | Pointer to **NullableString** |  | [optional] 
**ToolConfiguration** | **int32** |  | 

## Methods

### NewToolProductSettingsRequest

`func NewToolProductSettingsRequest(settingUrl string, product int32, name string, toolConfiguration int32, ) *ToolProductSettingsRequest`

NewToolProductSettingsRequest instantiates a new ToolProductSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolProductSettingsRequestWithDefaults

`func NewToolProductSettingsRequestWithDefaults() *ToolProductSettingsRequest`

NewToolProductSettingsRequestWithDefaults instantiates a new ToolProductSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingUrl

`func (o *ToolProductSettingsRequest) GetSettingUrl() string`

GetSettingUrl returns the SettingUrl field if non-nil, zero value otherwise.

### GetSettingUrlOk

`func (o *ToolProductSettingsRequest) GetSettingUrlOk() (*string, bool)`

GetSettingUrlOk returns a tuple with the SettingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingUrl

`func (o *ToolProductSettingsRequest) SetSettingUrl(v string)`

SetSettingUrl sets SettingUrl field to given value.


### GetProduct

`func (o *ToolProductSettingsRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ToolProductSettingsRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ToolProductSettingsRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetName

`func (o *ToolProductSettingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolProductSettingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolProductSettingsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolProductSettingsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolProductSettingsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolProductSettingsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolProductSettingsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ToolProductSettingsRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ToolProductSettingsRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *ToolProductSettingsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ToolProductSettingsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ToolProductSettingsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ToolProductSettingsRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ToolProductSettingsRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ToolProductSettingsRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetToolProjectId

`func (o *ToolProductSettingsRequest) GetToolProjectId() string`

GetToolProjectId returns the ToolProjectId field if non-nil, zero value otherwise.

### GetToolProjectIdOk

`func (o *ToolProductSettingsRequest) GetToolProjectIdOk() (*string, bool)`

GetToolProjectIdOk returns a tuple with the ToolProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolProjectId

`func (o *ToolProductSettingsRequest) SetToolProjectId(v string)`

SetToolProjectId sets ToolProjectId field to given value.

### HasToolProjectId

`func (o *ToolProductSettingsRequest) HasToolProjectId() bool`

HasToolProjectId returns a boolean if a field has been set.

### SetToolProjectIdNil

`func (o *ToolProductSettingsRequest) SetToolProjectIdNil(b bool)`

 SetToolProjectIdNil sets the value for ToolProjectId to be an explicit nil

### UnsetToolProjectId
`func (o *ToolProductSettingsRequest) UnsetToolProjectId()`

UnsetToolProjectId ensures that no value is present for ToolProjectId, not even an explicit nil
### GetToolConfiguration

`func (o *ToolProductSettingsRequest) GetToolConfiguration() int32`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *ToolProductSettingsRequest) GetToolConfigurationOk() (*int32, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *ToolProductSettingsRequest) SetToolConfiguration(v int32)`

SetToolConfiguration sets ToolConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


