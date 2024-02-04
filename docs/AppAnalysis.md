# AppAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Confidence** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**WebsiteFound** | Pointer to **NullableString** |  | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Product** | **int32** |  | 
**User** | **int32** |  | 

## Methods

### NewAppAnalysis

`func NewAppAnalysis(id int32, name string, created time.Time, product int32, user int32, ) *AppAnalysis`

NewAppAnalysis instantiates a new AppAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAnalysisWithDefaults

`func NewAppAnalysisWithDefaults() *AppAnalysis`

NewAppAnalysisWithDefaults instantiates a new AppAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppAnalysis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAnalysis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAnalysis) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *AppAnalysis) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AppAnalysis) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AppAnalysis) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AppAnalysis) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *AppAnalysis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppAnalysis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppAnalysis) SetName(v string)`

SetName sets Name field to given value.


### GetConfidence

`func (o *AppAnalysis) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AppAnalysis) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AppAnalysis) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AppAnalysis) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *AppAnalysis) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *AppAnalysis) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetVersion

`func (o *AppAnalysis) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppAnalysis) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppAnalysis) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppAnalysis) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AppAnalysis) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AppAnalysis) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIcon

`func (o *AppAnalysis) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AppAnalysis) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AppAnalysis) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AppAnalysis) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AppAnalysis) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AppAnalysis) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetWebsite

`func (o *AppAnalysis) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AppAnalysis) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AppAnalysis) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AppAnalysis) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *AppAnalysis) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *AppAnalysis) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetWebsiteFound

`func (o *AppAnalysis) GetWebsiteFound() string`

GetWebsiteFound returns the WebsiteFound field if non-nil, zero value otherwise.

### GetWebsiteFoundOk

`func (o *AppAnalysis) GetWebsiteFoundOk() (*string, bool)`

GetWebsiteFoundOk returns a tuple with the WebsiteFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteFound

`func (o *AppAnalysis) SetWebsiteFound(v string)`

SetWebsiteFound sets WebsiteFound field to given value.

### HasWebsiteFound

`func (o *AppAnalysis) HasWebsiteFound() bool`

HasWebsiteFound returns a boolean if a field has been set.

### SetWebsiteFoundNil

`func (o *AppAnalysis) SetWebsiteFoundNil(b bool)`

 SetWebsiteFoundNil sets the value for WebsiteFound to be an explicit nil

### UnsetWebsiteFound
`func (o *AppAnalysis) UnsetWebsiteFound()`

UnsetWebsiteFound ensures that no value is present for WebsiteFound, not even an explicit nil
### GetCreated

`func (o *AppAnalysis) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppAnalysis) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppAnalysis) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProduct

`func (o *AppAnalysis) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AppAnalysis) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AppAnalysis) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetUser

`func (o *AppAnalysis) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppAnalysis) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppAnalysis) SetUser(v int32)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


