# TestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**ScanType** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TargetStart** | **time.Time** |  | 
**TargetEnd** | **time.Time** |  | 
**PercentComplete** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID that was tested, a reimport may update this field. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash tested, a reimport may update this field. | [optional] 
**BranchTag** | Pointer to **NullableString** | Tag or branch that was tested, a reimport may update this field. | [optional] 
**Lead** | Pointer to **NullableInt32** |  | [optional] 
**TestType** | **int32** |  | 
**Environment** | Pointer to **NullableInt32** |  | [optional] 
**ApiScanConfiguration** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewTestRequest

`func NewTestRequest(targetStart time.Time, targetEnd time.Time, testType int32, ) *TestRequest`

NewTestRequest instantiates a new TestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRequestWithDefaults

`func NewTestRequestWithDefaults() *TestRequest`

NewTestRequestWithDefaults instantiates a new TestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TestRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetScanType

`func (o *TestRequest) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *TestRequest) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *TestRequest) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *TestRequest) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### SetScanTypeNil

`func (o *TestRequest) SetScanTypeNil(b bool)`

 SetScanTypeNil sets the value for ScanType to be an explicit nil

### UnsetScanType
`func (o *TestRequest) UnsetScanType()`

UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
### GetTitle

`func (o *TestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TestRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TestRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *TestRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTargetStart

`func (o *TestRequest) GetTargetStart() time.Time`

GetTargetStart returns the TargetStart field if non-nil, zero value otherwise.

### GetTargetStartOk

`func (o *TestRequest) GetTargetStartOk() (*time.Time, bool)`

GetTargetStartOk returns a tuple with the TargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStart

`func (o *TestRequest) SetTargetStart(v time.Time)`

SetTargetStart sets TargetStart field to given value.


### GetTargetEnd

`func (o *TestRequest) GetTargetEnd() time.Time`

GetTargetEnd returns the TargetEnd field if non-nil, zero value otherwise.

### GetTargetEndOk

`func (o *TestRequest) GetTargetEndOk() (*time.Time, bool)`

GetTargetEndOk returns a tuple with the TargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnd

`func (o *TestRequest) SetTargetEnd(v time.Time)`

SetTargetEnd sets TargetEnd field to given value.


### GetPercentComplete

`func (o *TestRequest) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TestRequest) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TestRequest) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *TestRequest) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *TestRequest) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *TestRequest) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetVersion

`func (o *TestRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TestRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TestRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TestRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TestRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TestRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildId

`func (o *TestRequest) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *TestRequest) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *TestRequest) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *TestRequest) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *TestRequest) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *TestRequest) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *TestRequest) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *TestRequest) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *TestRequest) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *TestRequest) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *TestRequest) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *TestRequest) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *TestRequest) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *TestRequest) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *TestRequest) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *TestRequest) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *TestRequest) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *TestRequest) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetLead

`func (o *TestRequest) GetLead() int32`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *TestRequest) GetLeadOk() (*int32, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *TestRequest) SetLead(v int32)`

SetLead sets Lead field to given value.

### HasLead

`func (o *TestRequest) HasLead() bool`

HasLead returns a boolean if a field has been set.

### SetLeadNil

`func (o *TestRequest) SetLeadNil(b bool)`

 SetLeadNil sets the value for Lead to be an explicit nil

### UnsetLead
`func (o *TestRequest) UnsetLead()`

UnsetLead ensures that no value is present for Lead, not even an explicit nil
### GetTestType

`func (o *TestRequest) GetTestType() int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *TestRequest) GetTestTypeOk() (*int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *TestRequest) SetTestType(v int32)`

SetTestType sets TestType field to given value.


### GetEnvironment

`func (o *TestRequest) GetEnvironment() int32`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TestRequest) GetEnvironmentOk() (*int32, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TestRequest) SetEnvironment(v int32)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TestRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *TestRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *TestRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetApiScanConfiguration

`func (o *TestRequest) GetApiScanConfiguration() int32`

GetApiScanConfiguration returns the ApiScanConfiguration field if non-nil, zero value otherwise.

### GetApiScanConfigurationOk

`func (o *TestRequest) GetApiScanConfigurationOk() (*int32, bool)`

GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiScanConfiguration

`func (o *TestRequest) SetApiScanConfiguration(v int32)`

SetApiScanConfiguration sets ApiScanConfiguration field to given value.

### HasApiScanConfiguration

`func (o *TestRequest) HasApiScanConfiguration() bool`

HasApiScanConfiguration returns a boolean if a field has been set.

### SetApiScanConfigurationNil

`func (o *TestRequest) SetApiScanConfigurationNil(b bool)`

 SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil

### UnsetApiScanConfiguration
`func (o *TestRequest) UnsetApiScanConfiguration()`

UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


