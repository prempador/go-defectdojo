# DojoGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ConfigurationPermissions** | Pointer to **[]int32** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**SocialProvider** | Pointer to **NullableString** | Group imported from a social provider.  * &#x60;AzureAD&#x60; - AzureAD | [optional] 
**Users** | **[]int32** |  | [readonly] 
**Prefetch** | Pointer to [**DojoGroupPrefetch**](DojoGroupPrefetch.md) |  | [optional] 

## Methods

### NewDojoGroup

`func NewDojoGroup(id int32, name string, users []int32, ) *DojoGroup`

NewDojoGroup instantiates a new DojoGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDojoGroupWithDefaults

`func NewDojoGroupWithDefaults() *DojoGroup`

NewDojoGroupWithDefaults instantiates a new DojoGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DojoGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DojoGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DojoGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetConfigurationPermissions

`func (o *DojoGroup) GetConfigurationPermissions() []*int32`

GetConfigurationPermissions returns the ConfigurationPermissions field if non-nil, zero value otherwise.

### GetConfigurationPermissionsOk

`func (o *DojoGroup) GetConfigurationPermissionsOk() (*[]*int32, bool)`

GetConfigurationPermissionsOk returns a tuple with the ConfigurationPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPermissions

`func (o *DojoGroup) SetConfigurationPermissions(v []*int32)`

SetConfigurationPermissions sets ConfigurationPermissions field to given value.

### HasConfigurationPermissions

`func (o *DojoGroup) HasConfigurationPermissions() bool`

HasConfigurationPermissions returns a boolean if a field has been set.

### GetName

`func (o *DojoGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DojoGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DojoGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DojoGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DojoGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DojoGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DojoGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DojoGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DojoGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSocialProvider

`func (o *DojoGroup) GetSocialProvider() string`

GetSocialProvider returns the SocialProvider field if non-nil, zero value otherwise.

### GetSocialProviderOk

`func (o *DojoGroup) GetSocialProviderOk() (*string, bool)`

GetSocialProviderOk returns a tuple with the SocialProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProvider

`func (o *DojoGroup) SetSocialProvider(v string)`

SetSocialProvider sets SocialProvider field to given value.

### HasSocialProvider

`func (o *DojoGroup) HasSocialProvider() bool`

HasSocialProvider returns a boolean if a field has been set.

### SetSocialProviderNil

`func (o *DojoGroup) SetSocialProviderNil(b bool)`

 SetSocialProviderNil sets the value for SocialProvider to be an explicit nil

### UnsetSocialProvider
`func (o *DojoGroup) UnsetSocialProvider()`

UnsetSocialProvider ensures that no value is present for SocialProvider, not even an explicit nil
### GetUsers

`func (o *DojoGroup) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DojoGroup) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DojoGroup) SetUsers(v []int32)`

SetUsers sets Users field to given value.


### GetPrefetch

`func (o *DojoGroup) GetPrefetch() DojoGroupPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *DojoGroup) GetPrefetchOk() (*DojoGroupPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *DojoGroup) SetPrefetch(v DojoGroupPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *DojoGroup) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


