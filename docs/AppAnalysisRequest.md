# AppAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Confidence** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**WebsiteFound** | Pointer to **NullableString** |  | [optional] 
**Product** | **int32** |  | 
**User** | **int32** |  | 

## Methods

### NewAppAnalysisRequest

`func NewAppAnalysisRequest(name string, product int32, user int32, ) *AppAnalysisRequest`

NewAppAnalysisRequest instantiates a new AppAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAnalysisRequestWithDefaults

`func NewAppAnalysisRequestWithDefaults() *AppAnalysisRequest`

NewAppAnalysisRequestWithDefaults instantiates a new AppAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *AppAnalysisRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AppAnalysisRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AppAnalysisRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AppAnalysisRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *AppAnalysisRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppAnalysisRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppAnalysisRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfidence

`func (o *AppAnalysisRequest) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AppAnalysisRequest) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AppAnalysisRequest) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *AppAnalysisRequest) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *AppAnalysisRequest) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *AppAnalysisRequest) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetVersion

`func (o *AppAnalysisRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AppAnalysisRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AppAnalysisRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AppAnalysisRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *AppAnalysisRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AppAnalysisRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIcon

`func (o *AppAnalysisRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AppAnalysisRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AppAnalysisRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AppAnalysisRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AppAnalysisRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AppAnalysisRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetWebsite

`func (o *AppAnalysisRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AppAnalysisRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AppAnalysisRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AppAnalysisRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *AppAnalysisRequest) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *AppAnalysisRequest) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetWebsiteFound

`func (o *AppAnalysisRequest) GetWebsiteFound() string`

GetWebsiteFound returns the WebsiteFound field if non-nil, zero value otherwise.

### GetWebsiteFoundOk

`func (o *AppAnalysisRequest) GetWebsiteFoundOk() (*string, bool)`

GetWebsiteFoundOk returns a tuple with the WebsiteFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteFound

`func (o *AppAnalysisRequest) SetWebsiteFound(v string)`

SetWebsiteFound sets WebsiteFound field to given value.

### HasWebsiteFound

`func (o *AppAnalysisRequest) HasWebsiteFound() bool`

HasWebsiteFound returns a boolean if a field has been set.

### SetWebsiteFoundNil

`func (o *AppAnalysisRequest) SetWebsiteFoundNil(b bool)`

 SetWebsiteFoundNil sets the value for WebsiteFound to be an explicit nil

### UnsetWebsiteFound
`func (o *AppAnalysisRequest) UnsetWebsiteFound()`

UnsetWebsiteFound ensures that no value is present for WebsiteFound, not even an explicit nil
### GetProduct

`func (o *AppAnalysisRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AppAnalysisRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AppAnalysisRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetUser

`func (o *AppAnalysisRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AppAnalysisRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AppAnalysisRequest) SetUser(v int32)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


