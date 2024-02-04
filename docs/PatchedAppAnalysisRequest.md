# PatchedAppAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**WebsiteFound** | Pointer to **NullableString** |  | [optional] 
**Product** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedAppAnalysisRequest

`func NewPatchedAppAnalysisRequest() *PatchedAppAnalysisRequest`

NewPatchedAppAnalysisRequest instantiates a new PatchedAppAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedAppAnalysisRequestWithDefaults

`func NewPatchedAppAnalysisRequestWithDefaults() *PatchedAppAnalysisRequest`

NewPatchedAppAnalysisRequestWithDefaults instantiates a new PatchedAppAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedAppAnalysisRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedAppAnalysisRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedAppAnalysisRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedAppAnalysisRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *PatchedAppAnalysisRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedAppAnalysisRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedAppAnalysisRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedAppAnalysisRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfidence

`func (o *PatchedAppAnalysisRequest) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PatchedAppAnalysisRequest) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PatchedAppAnalysisRequest) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PatchedAppAnalysisRequest) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *PatchedAppAnalysisRequest) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *PatchedAppAnalysisRequest) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetVersion

`func (o *PatchedAppAnalysisRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedAppAnalysisRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedAppAnalysisRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedAppAnalysisRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PatchedAppAnalysisRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PatchedAppAnalysisRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetIcon

`func (o *PatchedAppAnalysisRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *PatchedAppAnalysisRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *PatchedAppAnalysisRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *PatchedAppAnalysisRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *PatchedAppAnalysisRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *PatchedAppAnalysisRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetWebsite

`func (o *PatchedAppAnalysisRequest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *PatchedAppAnalysisRequest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *PatchedAppAnalysisRequest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *PatchedAppAnalysisRequest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *PatchedAppAnalysisRequest) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *PatchedAppAnalysisRequest) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetWebsiteFound

`func (o *PatchedAppAnalysisRequest) GetWebsiteFound() string`

GetWebsiteFound returns the WebsiteFound field if non-nil, zero value otherwise.

### GetWebsiteFoundOk

`func (o *PatchedAppAnalysisRequest) GetWebsiteFoundOk() (*string, bool)`

GetWebsiteFoundOk returns a tuple with the WebsiteFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteFound

`func (o *PatchedAppAnalysisRequest) SetWebsiteFound(v string)`

SetWebsiteFound sets WebsiteFound field to given value.

### HasWebsiteFound

`func (o *PatchedAppAnalysisRequest) HasWebsiteFound() bool`

HasWebsiteFound returns a boolean if a field has been set.

### SetWebsiteFoundNil

`func (o *PatchedAppAnalysisRequest) SetWebsiteFoundNil(b bool)`

 SetWebsiteFoundNil sets the value for WebsiteFound to be an explicit nil

### UnsetWebsiteFound
`func (o *PatchedAppAnalysisRequest) UnsetWebsiteFound()`

UnsetWebsiteFound ensures that no value is present for WebsiteFound, not even an explicit nil
### GetProduct

`func (o *PatchedAppAnalysisRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedAppAnalysisRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedAppAnalysisRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedAppAnalysisRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetUser

`func (o *PatchedAppAnalysisRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedAppAnalysisRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedAppAnalysisRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedAppAnalysisRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


