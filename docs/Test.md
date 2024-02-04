# Test

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TestTypeName** | **string** |  | [readonly] 
**FindingGroups** | [**[]FindingGroup**](FindingGroup.md) |  | [readonly] 
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
**Engagement** | **int32** |  | [readonly] 
**Lead** | Pointer to **NullableInt32** |  | [optional] 
**TestType** | **int32** |  | 
**Environment** | Pointer to **NullableInt32** |  | [optional] 
**ApiScanConfiguration** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | [**[]Note**](Note.md) |  | [readonly] 
**Files** | [**[]File**](File.md) |  | [readonly] 

## Methods

### NewTest

`func NewTest(id int32, testTypeName string, findingGroups []FindingGroup, targetStart time.Time, targetEnd time.Time, estimatedTime NullableString, actualTime NullableString, updated NullableTime, created NullableTime, engagement int32, testType int32, notes []Note, files []File, ) *Test`

NewTest instantiates a new Test object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestWithDefaults

`func NewTestWithDefaults() *Test`

NewTestWithDefaults instantiates a new Test object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Test) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Test) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Test) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *Test) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Test) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Test) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Test) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTestTypeName

`func (o *Test) GetTestTypeName() string`

GetTestTypeName returns the TestTypeName field if non-nil, zero value otherwise.

### GetTestTypeNameOk

`func (o *Test) GetTestTypeNameOk() (*string, bool)`

GetTestTypeNameOk returns a tuple with the TestTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTypeName

`func (o *Test) SetTestTypeName(v string)`

SetTestTypeName sets TestTypeName field to given value.


### GetFindingGroups

`func (o *Test) GetFindingGroups() []FindingGroup`

GetFindingGroups returns the FindingGroups field if non-nil, zero value otherwise.

### GetFindingGroupsOk

`func (o *Test) GetFindingGroupsOk() (*[]FindingGroup, bool)`

GetFindingGroupsOk returns a tuple with the FindingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingGroups

`func (o *Test) SetFindingGroups(v []FindingGroup)`

SetFindingGroups sets FindingGroups field to given value.


### GetScanType

`func (o *Test) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *Test) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *Test) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *Test) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### SetScanTypeNil

`func (o *Test) SetScanTypeNil(b bool)`

 SetScanTypeNil sets the value for ScanType to be an explicit nil

### UnsetScanType
`func (o *Test) UnsetScanType()`

UnsetScanType ensures that no value is present for ScanType, not even an explicit nil
### GetTitle

`func (o *Test) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Test) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Test) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Test) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Test) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Test) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *Test) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Test) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Test) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Test) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Test) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Test) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTargetStart

`func (o *Test) GetTargetStart() time.Time`

GetTargetStart returns the TargetStart field if non-nil, zero value otherwise.

### GetTargetStartOk

`func (o *Test) GetTargetStartOk() (*time.Time, bool)`

GetTargetStartOk returns a tuple with the TargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStart

`func (o *Test) SetTargetStart(v time.Time)`

SetTargetStart sets TargetStart field to given value.


### GetTargetEnd

`func (o *Test) GetTargetEnd() time.Time`

GetTargetEnd returns the TargetEnd field if non-nil, zero value otherwise.

### GetTargetEndOk

`func (o *Test) GetTargetEndOk() (*time.Time, bool)`

GetTargetEndOk returns a tuple with the TargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnd

`func (o *Test) SetTargetEnd(v time.Time)`

SetTargetEnd sets TargetEnd field to given value.


### GetEstimatedTime

`func (o *Test) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *Test) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *Test) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.


### SetEstimatedTimeNil

`func (o *Test) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *Test) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
### GetActualTime

`func (o *Test) GetActualTime() string`

GetActualTime returns the ActualTime field if non-nil, zero value otherwise.

### GetActualTimeOk

`func (o *Test) GetActualTimeOk() (*string, bool)`

GetActualTimeOk returns a tuple with the ActualTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualTime

`func (o *Test) SetActualTime(v string)`

SetActualTime sets ActualTime field to given value.


### SetActualTimeNil

`func (o *Test) SetActualTimeNil(b bool)`

 SetActualTimeNil sets the value for ActualTime to be an explicit nil

### UnsetActualTime
`func (o *Test) UnsetActualTime()`

UnsetActualTime ensures that no value is present for ActualTime, not even an explicit nil
### GetPercentComplete

`func (o *Test) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *Test) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *Test) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *Test) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *Test) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *Test) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetUpdated

`func (o *Test) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Test) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Test) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### SetUpdatedNil

`func (o *Test) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *Test) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetCreated

`func (o *Test) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Test) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Test) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Test) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Test) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetVersion

`func (o *Test) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Test) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Test) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Test) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Test) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Test) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildId

`func (o *Test) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *Test) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *Test) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *Test) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *Test) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *Test) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *Test) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *Test) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *Test) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *Test) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *Test) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *Test) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *Test) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *Test) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *Test) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *Test) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *Test) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *Test) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetEngagement

`func (o *Test) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *Test) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *Test) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.


### GetLead

`func (o *Test) GetLead() int32`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *Test) GetLeadOk() (*int32, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *Test) SetLead(v int32)`

SetLead sets Lead field to given value.

### HasLead

`func (o *Test) HasLead() bool`

HasLead returns a boolean if a field has been set.

### SetLeadNil

`func (o *Test) SetLeadNil(b bool)`

 SetLeadNil sets the value for Lead to be an explicit nil

### UnsetLead
`func (o *Test) UnsetLead()`

UnsetLead ensures that no value is present for Lead, not even an explicit nil
### GetTestType

`func (o *Test) GetTestType() int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *Test) GetTestTypeOk() (*int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *Test) SetTestType(v int32)`

SetTestType sets TestType field to given value.


### GetEnvironment

`func (o *Test) GetEnvironment() int32`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Test) GetEnvironmentOk() (*int32, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Test) SetEnvironment(v int32)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Test) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *Test) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *Test) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetApiScanConfiguration

`func (o *Test) GetApiScanConfiguration() int32`

GetApiScanConfiguration returns the ApiScanConfiguration field if non-nil, zero value otherwise.

### GetApiScanConfigurationOk

`func (o *Test) GetApiScanConfigurationOk() (*int32, bool)`

GetApiScanConfigurationOk returns a tuple with the ApiScanConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiScanConfiguration

`func (o *Test) SetApiScanConfiguration(v int32)`

SetApiScanConfiguration sets ApiScanConfiguration field to given value.

### HasApiScanConfiguration

`func (o *Test) HasApiScanConfiguration() bool`

HasApiScanConfiguration returns a boolean if a field has been set.

### SetApiScanConfigurationNil

`func (o *Test) SetApiScanConfigurationNil(b bool)`

 SetApiScanConfigurationNil sets the value for ApiScanConfiguration to be an explicit nil

### UnsetApiScanConfiguration
`func (o *Test) UnsetApiScanConfiguration()`

UnsetApiScanConfiguration ensures that no value is present for ApiScanConfiguration, not even an explicit nil
### GetNotes

`func (o *Test) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Test) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Test) SetNotes(v []Note)`

SetNotes sets Notes field to given value.


### GetFiles

`func (o *Test) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Test) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Test) SetFiles(v []File)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


