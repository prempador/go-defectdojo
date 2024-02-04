# TestCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Engagement** | **int32** |  | 
**Notes** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ScanType** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**TargetStart** | **time.Time** |  | 
**TargetEnd** | **time.Time** |  | 
**EstimatedTime** | **NullableString** |  | [readonly] 
**ActualTime** | **NullableString** |  | [readonly] 
**PercentComplete** | Pointer to **NullableInt32** |  | [optional] 
**Updated** | **NullableTime** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**Version** | Pointer to **NullableString** |  | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID that was tested, a reimport may update this field. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash tested, a reimport may update this field. | [optional] 
**BranchTag** | Pointer to **NullableString** | Tag or branch that was tested, a reimport may update this field. | [optional] 
**Lead** | Pointer to **NullableInt32** |  | [optional] 
**TestType** | **int32** |  | 
**Environment** | Pointer to **NullableInt32** |  | [optional] 
**ApiScanConfiguration** | Pointer to **NullableInt32** |  | [optional] 
**Files** | **[]int32** |  | [readonly] 

## Methods

### NewTestCreate

`func NewTestCreate(id int32, engagement int32, targetStart time.Time, targetEnd time.Time, estimatedTime NullableString, actualTime NullableString, updated NullableTime, created NullableTime, testType int32, files []int32, ) *TestCreate`

NewTestCreate instantiates a new TestCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestCreateWithDefaults

`func NewTestCreateWithDefaults() *TestCreate`

NewTestCreateWithDefaults instantiates a new TestCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestCreate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestCreate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestCreate) SetId(v int32)`

SetId sets Id field to given value.


### GetEngagement

`func (o *TestCreate) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *TestCreate) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *TestCreate) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.


### GetNotes

`func (o *TestCreate) GetNotes() []*int32`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TestCreate) GetNotesOk() (*[]*int32, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TestCreate) SetNotes(v []*int32)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TestCreate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *TestCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetScanType

`func (o *TestCreate) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *TestCreate) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *TestCreate) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *TestCreate) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### SetScanTypeNil

`func (o *TestCreate) SetScanTypeNil(b bool)`

 SetScanTypeNil sets the value for ScanType to be an explicit nil

### UnsetScanType
`func (o *TestCreate) UnsetScanType()`

UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
### GetTitle

`func (o *TestCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestCreate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestCreate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TestCreate) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TestCreate) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *TestCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTargetStart

`func (o *TestCreate) GetTargetStart() time.Time`

GetTargetStart returns the TargetStart field if non-nil, zero value otherwise.

### GetTargetStartOk

`func (o *TestCreate) GetTargetStartOk() (*time.Time, bool)`

GetTargetStartOk returns a tuple with the TargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStart

`func (o *TestCreate) SetTargetStart(v time.Time)`

SetTargetStart sets TargetStart field to given value.


### GetTargetEnd

`func (o *TestCreate) GetTargetEnd() time.Time`

GetTargetEnd returns the TargetEnd field if non-nil, zero value otherwise.

### GetTargetEndOk

`func (o *TestCreate) GetTargetEndOk() (*time.Time, bool)`

GetTargetEndOk returns a tuple with the TargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnd

`func (o *TestCreate) SetTargetEnd(v time.Time)`

SetTargetEnd sets TargetEnd field to given value.


### GetEstimatedTime

`func (o *TestCreate) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *TestCreate) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *TestCreate) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.


### SetEstimatedTimeNil

`func (o *TestCreate) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *TestCreate) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
### GetActualTime

`func (o *TestCreate) GetActualTime() string`

GetActualTime returns the ActualTime field if non-nil, zero value otherwise.

### GetActualTimeOk

`func (o *TestCreate) GetActualTimeOk() (*string, bool)`

GetActualTimeOk returns a tuple with the ActualTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTime

`func (o *TestCreate) SetActualTime(v string)`

SetActualTime sets ActualTime field to given value.


### SetActualTimeNil

`func (o *TestCreate) SetActualTimeNil(b bool)`

 SetActualTimeNil sets the value for ActualTime to be an explicit nil

### UnsetActualTime
`func (o *TestCreate) UnsetActualTime()`

UnsetActualTime ensures that no value is present for ActualTime, not even an explicit nil
### GetPercentComplete

`func (o *TestCreate) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *TestCreate) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *TestCreate) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *TestCreate) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *TestCreate) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *TestCreate) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetUpdated

`func (o *TestCreate) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TestCreate) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TestCreate) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### SetUpdatedNil

`func (o *TestCreate) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *TestCreate) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetCreated

`func (o *TestCreate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestCreate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestCreate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *TestCreate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *TestCreate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetVersion

`func (o *TestCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TestCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TestCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TestCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TestCreate) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TestCreate) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildId

`func (o *TestCreate) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *TestCreate) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *TestCreate) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *TestCreate) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *TestCreate) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *TestCreate) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *TestCreate) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *TestCreate) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *TestCreate) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *TestCreate) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *TestCreate) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *TestCreate) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *TestCreate) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *TestCreate) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *TestCreate) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *TestCreate) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *TestCreate) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *TestCreate) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetLead

`func (o *TestCreate) GetLead() int32`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *TestCreate) GetLeadOk() (*int32, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *TestCreate) SetLead(v int32)`

SetLead sets Lead field to given value.

### HasLead

`func (o *TestCreate) HasLead() bool`

HasLead returns a boolean if a field has been set.

### SetLeadNil

`func (o *TestCreate) SetLeadNil(b bool)`

 SetLeadNil sets the value for Lead to be an explicit nil

### UnsetLead
`func (o *TestCreate) UnsetLead()`

UnsetLead ensures that no value is present for Lead, not even an explicit nil
### GetTestType

`func (o *TestCreate) GetTestType() int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *TestCreate) GetTestTypeOk() (*int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *TestCreate) SetTestType(v int32)`

SetTestType sets TestType field to given value.


### GetEnvironment

`func (o *TestCreate) GetEnvironment() int32`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *TestCreate) GetEnvironmentOk() (*int32, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *TestCreate) SetEnvironment(v int32)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *TestCreate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *TestCreate) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *TestCreate) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetApiScanConfiguration

`func (o *TestCreate) GetApiScanConfiguration() int32`

GetApiScanConfiguration returns the ApiScanConfiguration field if non-nil, zero value otherwise.

### GetApiScanConfigurationOk

`func (o *TestCreate) GetApiScanConfigurationOk() (*int32, bool)`

GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiScanConfiguration

`func (o *TestCreate) SetApiScanConfiguration(v int32)`

SetApiScanConfiguration sets ApiScanConfiguration field to given value.

### HasApiScanConfiguration

`func (o *TestCreate) HasApiScanConfiguration() bool`

HasApiScanConfiguration returns a boolean if a field has been set.

### SetApiScanConfigurationNil

`func (o *TestCreate) SetApiScanConfigurationNil(b bool)`

 SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil

### UnsetApiScanConfiguration
`func (o *TestCreate) UnsetApiScanConfiguration()`

UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil
### GetFiles

`func (o *TestCreate) GetFiles() []int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TestCreate) GetFilesOk() (*[]int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TestCreate) SetFiles(v []int32)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


