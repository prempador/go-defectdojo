# DojoGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationPermissions** | Pointer to **[]int32** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**SocialProvider** | Pointer to **NullableString** | Group imported from a social provider.  * &#x60;AzureAD&#x60; - AzureAD | [optional] 

## Methods

### NewDojoGroupRequest

`func NewDojoGroupRequest(name string, ) *DojoGroupRequest`

NewDojoGroupRequest instantiates a new DojoGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDojoGroupRequestWithDefaults

`func NewDojoGroupRequestWithDefaults() *DojoGroupRequest`

NewDojoGroupRequestWithDefaults instantiates a new DojoGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationPermissions

`func (o *DojoGroupRequest) GetConfigurationPermissions() []*int32`

GetConfigurationPermissions returns the ConfigurationPermissions field if non-nil, zero value otherwise.

### GetConfigurationPermissionsOk

`func (o *DojoGroupRequest) GetConfigurationPermissionsOk() (*[]*int32, bool)`

GetConfigurationPermissionsOk returns a tuple with the ConfigurationPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPermissions

`func (o *DojoGroupRequest) SetConfigurationPermissions(v []*int32)`

SetConfigurationPermissions sets ConfigurationPermissions field to given value.

### HasConfigurationPermissions

`func (o *DojoGroupRequest) HasConfigurationPermissions() bool`

HasConfigurationPermissions returns a boolean if a field has been set.

### GetName

`func (o *DojoGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DojoGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DojoGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DojoGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DojoGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DojoGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DojoGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DojoGroupRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DojoGroupRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSocialProvider

`func (o *DojoGroupRequest) GetSocialProvider() string`

GetSocialProvider returns the SocialProvider field if non-nil, zero value otherwise.

### GetSocialProviderOk

`func (o *DojoGroupRequest) GetSocialProviderOk() (*string, bool)`

GetSocialProviderOk returns a tuple with the SocialProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProvider

`func (o *DojoGroupRequest) SetSocialProvider(v string)`

SetSocialProvider sets SocialProvider field to given value.

### HasSocialProvider

`func (o *DojoGroupRequest) HasSocialProvider() bool`

HasSocialProvider returns a boolean if a field has been set.

### SetSocialProviderNil

`func (o *DojoGroupRequest) SetSocialProviderNil(b bool)`

 SetSocialProviderNil sets the value for SocialProvider to be an explicit nil

### UnsetSocialProvider
`func (o *DojoGroupRequest) UnsetSocialProvider()`

UnsetSocialProvider ensures that no value is present for SocialProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


