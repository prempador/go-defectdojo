# PatchedToolProductSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettingUrl** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**ToolProjectId** | Pointer to **NullableString** |  | [optional] 
**ToolConfiguration** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedToolProductSettingsRequest

`func NewPatchedToolProductSettingsRequest() *PatchedToolProductSettingsRequest`

NewPatchedToolProductSettingsRequest instantiates a new PatchedToolProductSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedToolProductSettingsRequestWithDefaults

`func NewPatchedToolProductSettingsRequestWithDefaults() *PatchedToolProductSettingsRequest`

NewPatchedToolProductSettingsRequestWithDefaults instantiates a new PatchedToolProductSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettingUrl

`func (o *PatchedToolProductSettingsRequest) GetSettingUrl() string`

GetSettingUrl returns the SettingUrl field if non-nil, zero value otherwise.

### GetSettingUrlOk

`func (o *PatchedToolProductSettingsRequest) GetSettingUrlOk() (*string, bool)`

GetSettingUrlOk returns a tuple with the SettingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingUrl

`func (o *PatchedToolProductSettingsRequest) SetSettingUrl(v string)`

SetSettingUrl sets SettingUrl field to given value.

### HasSettingUrl

`func (o *PatchedToolProductSettingsRequest) HasSettingUrl() bool`

HasSettingUrl returns a boolean if a field has been set.

### GetProduct

`func (o *PatchedToolProductSettingsRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedToolProductSettingsRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedToolProductSettingsRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedToolProductSettingsRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetName

`func (o *PatchedToolProductSettingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedToolProductSettingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedToolProductSettingsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedToolProductSettingsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedToolProductSettingsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedToolProductSettingsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedToolProductSettingsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedToolProductSettingsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedToolProductSettingsRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedToolProductSettingsRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUrl

`func (o *PatchedToolProductSettingsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedToolProductSettingsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedToolProductSettingsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedToolProductSettingsRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PatchedToolProductSettingsRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PatchedToolProductSettingsRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetToolProjectId

`func (o *PatchedToolProductSettingsRequest) GetToolProjectId() string`

GetToolProjectId returns the ToolProjectId field if non-nil, zero value otherwise.

### GetToolProjectIdOk

`func (o *PatchedToolProductSettingsRequest) GetToolProjectIdOk() (*string, bool)`

GetToolProjectIdOk returns a tuple with the ToolProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolProjectId

`func (o *PatchedToolProductSettingsRequest) SetToolProjectId(v string)`

SetToolProjectId sets ToolProjectId field to given value.

### HasToolProjectId

`func (o *PatchedToolProductSettingsRequest) HasToolProjectId() bool`

HasToolProjectId returns a boolean if a field has been set.

### SetToolProjectIdNil

`func (o *PatchedToolProductSettingsRequest) SetToolProjectIdNil(b bool)`

 SetToolProjectIdNil sets the value for ToolProjectId to be an explicit nil

### UnsetToolProjectId
`func (o *PatchedToolProductSettingsRequest) UnsetToolProjectId()`

UnsetToolProjectId ensures that no value is present for ToolProjectId, not even an explicit nil
### GetToolConfiguration

`func (o *PatchedToolProductSettingsRequest) GetToolConfiguration() int32`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *PatchedToolProductSettingsRequest) GetToolConfigurationOk() (*int32, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *PatchedToolProductSettingsRequest) SetToolConfiguration(v int32)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *PatchedToolProductSettingsRequest) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


