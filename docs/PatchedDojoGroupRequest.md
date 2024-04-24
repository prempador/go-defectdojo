# PatchedDojoGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationPermissions** | Pointer to **[]int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SocialProvider** | Pointer to **NullableString** | Group imported from a social provider.  * &#x60;AzureAD&#x60; - AzureAD * &#x60;Remote&#x60; - Remote | [optional] 

## Methods

### NewPatchedDojoGroupRequest

`func NewPatchedDojoGroupRequest() *PatchedDojoGroupRequest`

NewPatchedDojoGroupRequest instantiates a new PatchedDojoGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDojoGroupRequestWithDefaults

`func NewPatchedDojoGroupRequestWithDefaults() *PatchedDojoGroupRequest`

NewPatchedDojoGroupRequestWithDefaults instantiates a new PatchedDojoGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationPermissions

`func (o *PatchedDojoGroupRequest) GetConfigurationPermissions() []*int32`

GetConfigurationPermissions returns the ConfigurationPermissions field if non-nil, zero value otherwise.

### GetConfigurationPermissionsOk

`func (o *PatchedDojoGroupRequest) GetConfigurationPermissionsOk() (*[]*int32, bool)`

GetConfigurationPermissionsOk returns a tuple with the ConfigurationPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPermissions

`func (o *PatchedDojoGroupRequest) SetConfigurationPermissions(v []*int32)`

SetConfigurationPermissions sets ConfigurationPermissions field to given value.

### HasConfigurationPermissions

`func (o *PatchedDojoGroupRequest) HasConfigurationPermissions() bool`

HasConfigurationPermissions returns a boolean if a field has been set.

### GetName

`func (o *PatchedDojoGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDojoGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDojoGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDojoGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedDojoGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDojoGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDojoGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDojoGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedDojoGroupRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedDojoGroupRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSocialProvider

`func (o *PatchedDojoGroupRequest) GetSocialProvider() string`

GetSocialProvider returns the SocialProvider field if non-nil, zero value otherwise.

### GetSocialProviderOk

`func (o *PatchedDojoGroupRequest) GetSocialProviderOk() (*string, bool)`

GetSocialProviderOk returns a tuple with the SocialProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProvider

`func (o *PatchedDojoGroupRequest) SetSocialProvider(v string)`

SetSocialProvider sets SocialProvider field to given value.

### HasSocialProvider

`func (o *PatchedDojoGroupRequest) HasSocialProvider() bool`

HasSocialProvider returns a boolean if a field has been set.

### SetSocialProviderNil

`func (o *PatchedDojoGroupRequest) SetSocialProviderNil(b bool)`

 SetSocialProviderNil sets the value for SocialProvider to be an explicit nil

### UnsetSocialProvider
`func (o *PatchedDojoGroupRequest) UnsetSocialProvider()`

UnsetSocialProvider ensures that no value is present for SocialProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


