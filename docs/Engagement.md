# Engagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** | Version of the product the engagement tested. | [optional] 
**FirstContacted** | Pointer to **NullableString** |  | [optional] 
**TargetStart** | **string** |  | 
**TargetEnd** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Updated** | **NullableTime** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**Active** | **bool** |  | [readonly] 
**Tracker** | Pointer to **NullableString** | Link to epic or ticket system with changes to version. | [optional] 
**TestStrategy** | Pointer to **NullableString** |  | [optional] 
**ThreatModel** | Pointer to **bool** |  | [optional] 
**ApiTest** | Pointer to **bool** |  | [optional] 
**PenTest** | Pointer to **bool** |  | [optional] 
**CheckList** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **NullableString** | * &#x60;Not Started&#x60; - Not Started * &#x60;Blocked&#x60; - Blocked * &#x60;Cancelled&#x60; - Cancelled * &#x60;Completed&#x60; - Completed * &#x60;In Progress&#x60; - In Progress * &#x60;On Hold&#x60; - On Hold * &#x60;Waiting for Resource&#x60; - Waiting for Resource | [optional] 
**Progress** | **string** |  | [readonly] 
**TmodelPath** | **NullableString** |  | [readonly] 
**DoneTesting** | **bool** |  | [readonly] 
**EngagementType** | Pointer to **NullableString** | * &#x60;Interactive&#x60; - Interactive * &#x60;CI/CD&#x60; - CI/CD | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID of the product the engagement tested. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash from repo | [optional] 
**BranchTag** | Pointer to **NullableString** | Tag or branch of the product the engagement tested. | [optional] 
**SourceCodeManagementUri** | Pointer to **NullableString** | Resource link to source code | [optional] 
**DeduplicationOnEngagement** | Pointer to **bool** | If enabled deduplication will only mark a finding in this engagement as duplicate of another finding if both findings are in this engagement. If disabled, deduplication is on the product level. | [optional] 
**Lead** | Pointer to **NullableInt32** |  | [optional] 
**Requester** | Pointer to **NullableInt32** |  | [optional] 
**Preset** | Pointer to **NullableInt32** | Settings and notes for performing this engagement. | [optional] 
**ReportType** | Pointer to **NullableInt32** |  | [optional] 
**Product** | **int32** |  | 
**BuildServer** | Pointer to **NullableInt32** | Build server responsible for CI/CD test | [optional] 
**SourceCodeManagementServer** | Pointer to **NullableInt32** | Source code server for CI/CD test | [optional] 
**OrchestrationEngine** | Pointer to **NullableInt32** | Orchestration service responsible for CI/CD test | [optional] 
**Notes** | [**[]Note**](Note.md) |  | [readonly] 
**Files** | [**[]File**](File.md) |  | [readonly] 
**RiskAcceptance** | **[]int32** |  | [readonly] 

## Methods

### NewEngagement

`func NewEngagement(id int32, targetStart string, targetEnd string, updated NullableTime, created NullableTime, active bool, progress string, tmodelPath NullableString, doneTesting bool, product int32, notes []Note, files []File, riskAcceptance []int32, ) *Engagement`

NewEngagement instantiates a new Engagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementWithDefaults

`func NewEngagementWithDefaults() *Engagement`

NewEngagementWithDefaults instantiates a new Engagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Engagement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Engagement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Engagement) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *Engagement) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Engagement) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Engagement) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Engagement) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *Engagement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Engagement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Engagement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Engagement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Engagement) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Engagement) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Engagement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Engagement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Engagement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Engagement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Engagement) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Engagement) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *Engagement) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Engagement) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Engagement) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Engagement) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Engagement) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Engagement) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFirstContacted

`func (o *Engagement) GetFirstContacted() string`

GetFirstContacted returns the FirstContacted field if non-nil, zero value otherwise.

### GetFirstContactedOk

`func (o *Engagement) GetFirstContactedOk() (*string, bool)`

GetFirstContactedOk returns a tuple with the FirstContacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstContacted

`func (o *Engagement) SetFirstContacted(v string)`

SetFirstContacted sets FirstContacted field to given value.

### HasFirstContacted

`func (o *Engagement) HasFirstContacted() bool`

HasFirstContacted returns a boolean if a field has been set.

### SetFirstContactedNil

`func (o *Engagement) SetFirstContactedNil(b bool)`

 SetFirstContactedNil sets the value for FirstContacted to be an explicit nil

### UnsetFirstContacted
`func (o *Engagement) UnsetFirstContacted()`

UnsetFirstContacted ensures that no value is present for FirstContacted, not even an explicit nil
### GetTargetStart

`func (o *Engagement) GetTargetStart() string`

GetTargetStart returns the TargetStart field if non-nil, zero value otherwise.

### GetTargetStartOk

`func (o *Engagement) GetTargetStartOk() (*string, bool)`

GetTargetStartOk returns a tuple with the TargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStart

`func (o *Engagement) SetTargetStart(v string)`

SetTargetStart sets TargetStart field to given value.


### GetTargetEnd

`func (o *Engagement) GetTargetEnd() string`

GetTargetEnd returns the TargetEnd field if non-nil, zero value otherwise.

### GetTargetEndOk

`func (o *Engagement) GetTargetEndOk() (*string, bool)`

GetTargetEndOk returns a tuple with the TargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnd

`func (o *Engagement) SetTargetEnd(v string)`

SetTargetEnd sets TargetEnd field to given value.


### GetReason

`func (o *Engagement) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Engagement) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Engagement) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Engagement) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *Engagement) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Engagement) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetUpdated

`func (o *Engagement) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Engagement) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Engagement) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### SetUpdatedNil

`func (o *Engagement) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *Engagement) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetCreated

`func (o *Engagement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Engagement) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Engagement) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Engagement) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Engagement) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetActive

`func (o *Engagement) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Engagement) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Engagement) SetActive(v bool)`

SetActive sets Active field to given value.


### GetTracker

`func (o *Engagement) GetTracker() string`

GetTracker returns the Tracker field if non-nil, zero value otherwise.

### GetTrackerOk

`func (o *Engagement) GetTrackerOk() (*string, bool)`

GetTrackerOk returns a tuple with the Tracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracker

`func (o *Engagement) SetTracker(v string)`

SetTracker sets Tracker field to given value.

### HasTracker

`func (o *Engagement) HasTracker() bool`

HasTracker returns a boolean if a field has been set.

### SetTrackerNil

`func (o *Engagement) SetTrackerNil(b bool)`

 SetTrackerNil sets the value for Tracker to be an explicit nil

### UnsetTracker
`func (o *Engagement) UnsetTracker()`

UnsetTracker ensures that no value is present for Tracker, not even an explicit nil
### GetTestStrategy

`func (o *Engagement) GetTestStrategy() string`

GetTestStrategy returns the TestStrategy field if non-nil, zero value otherwise.

### GetTestStrategyOk

`func (o *Engagement) GetTestStrategyOk() (*string, bool)`

GetTestStrategyOk returns a tuple with the TestStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStrategy

`func (o *Engagement) SetTestStrategy(v string)`

SetTestStrategy sets TestStrategy field to given value.

### HasTestStrategy

`func (o *Engagement) HasTestStrategy() bool`

HasTestStrategy returns a boolean if a field has been set.

### SetTestStrategyNil

`func (o *Engagement) SetTestStrategyNil(b bool)`

 SetTestStrategyNil sets the value for TestStrategy to be an explicit nil

### UnsetTestStrategy
`func (o *Engagement) UnsetTestStrategy()`

UnsetTestStrategy ensures that no value is present for TestStrategy, not even an explicit nil
### GetThreatModel

`func (o *Engagement) GetThreatModel() bool`

GetThreatModel returns the ThreatModel field if non-nil, zero value otherwise.

### GetThreatModelOk

`func (o *Engagement) GetThreatModelOk() (*bool, bool)`

GetThreatModelOk returns a tuple with the ThreatModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatModel

`func (o *Engagement) SetThreatModel(v bool)`

SetThreatModel sets ThreatModel field to given value.

### HasThreatModel

`func (o *Engagement) HasThreatModel() bool`

HasThreatModel returns a boolean if a field has been set.

### GetApiTest

`func (o *Engagement) GetApiTest() bool`

GetApiTest returns the ApiTest field if non-nil, zero value otherwise.

### GetApiTestOk

`func (o *Engagement) GetApiTestOk() (*bool, bool)`

GetApiTestOk returns a tuple with the ApiTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTest

`func (o *Engagement) SetApiTest(v bool)`

SetApiTest sets ApiTest field to given value.

### HasApiTest

`func (o *Engagement) HasApiTest() bool`

HasApiTest returns a boolean if a field has been set.

### GetPenTest

`func (o *Engagement) GetPenTest() bool`

GetPenTest returns the PenTest field if non-nil, zero value otherwise.

### GetPenTestOk

`func (o *Engagement) GetPenTestOk() (*bool, bool)`

GetPenTestOk returns a tuple with the PenTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenTest

`func (o *Engagement) SetPenTest(v bool)`

SetPenTest sets PenTest field to given value.

### HasPenTest

`func (o *Engagement) HasPenTest() bool`

HasPenTest returns a boolean if a field has been set.

### GetCheckList

`func (o *Engagement) GetCheckList() bool`

GetCheckList returns the CheckList field if non-nil, zero value otherwise.

### GetCheckListOk

`func (o *Engagement) GetCheckListOk() (*bool, bool)`

GetCheckListOk returns a tuple with the CheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckList

`func (o *Engagement) SetCheckList(v bool)`

SetCheckList sets CheckList field to given value.

### HasCheckList

`func (o *Engagement) HasCheckList() bool`

HasCheckList returns a boolean if a field has been set.

### GetStatus

`func (o *Engagement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Engagement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Engagement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Engagement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Engagement) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Engagement) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetProgress

`func (o *Engagement) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *Engagement) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *Engagement) SetProgress(v string)`

SetProgress sets Progress field to given value.


### GetTmodelPath

`func (o *Engagement) GetTmodelPath() string`

GetTmodelPath returns the TmodelPath field if non-nil, zero value otherwise.

### GetTmodelPathOk

`func (o *Engagement) GetTmodelPathOk() (*string, bool)`

GetTmodelPathOk returns a tuple with the TmodelPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmodelPath

`func (o *Engagement) SetTmodelPath(v string)`

SetTmodelPath sets TmodelPath field to given value.


### SetTmodelPathNil

`func (o *Engagement) SetTmodelPathNil(b bool)`

 SetTmodelPathNil sets the value for TmodelPath to be an explicit nil

### UnsetTmodelPath
`func (o *Engagement) UnsetTmodelPath()`

UnsetTmodelPath ensures that no value is present for TmodelPath, not even an explicit nil
### GetDoneTesting

`func (o *Engagement) GetDoneTesting() bool`

GetDoneTesting returns the DoneTesting field if non-nil, zero value otherwise.

### GetDoneTestingOk

`func (o *Engagement) GetDoneTestingOk() (*bool, bool)`

GetDoneTestingOk returns a tuple with the DoneTesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoneTesting

`func (o *Engagement) SetDoneTesting(v bool)`

SetDoneTesting sets DoneTesting field to given value.


### GetEngagementType

`func (o *Engagement) GetEngagementType() string`

GetEngagementType returns the EngagementType field if non-nil, zero value otherwise.

### GetEngagementTypeOk

`func (o *Engagement) GetEngagementTypeOk() (*string, bool)`

GetEngagementTypeOk returns a tuple with the EngagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementType

`func (o *Engagement) SetEngagementType(v string)`

SetEngagementType sets EngagementType field to given value.

### HasEngagementType

`func (o *Engagement) HasEngagementType() bool`

HasEngagementType returns a boolean if a field has been set.

### SetEngagementTypeNil

`func (o *Engagement) SetEngagementTypeNil(b bool)`

 SetEngagementTypeNil sets the value for EngagementType to be an explicit nil

### UnsetEngagementType
`func (o *Engagement) UnsetEngagementType()`

UnsetEngagementType ensures that no value is present for EngagementType, not even an explicit nil
### GetBuildId

`func (o *Engagement) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *Engagement) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *Engagement) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *Engagement) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *Engagement) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *Engagement) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *Engagement) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *Engagement) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *Engagement) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *Engagement) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *Engagement) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *Engagement) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *Engagement) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *Engagement) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *Engagement) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *Engagement) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *Engagement) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *Engagement) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetSourceCodeManagementUri

`func (o *Engagement) GetSourceCodeManagementUri() string`

GetSourceCodeManagementUri returns the SourceCodeManagementUri field if non-nil, zero value otherwise.

### GetSourceCodeManagementUriOk

`func (o *Engagement) GetSourceCodeManagementUriOk() (*string, bool)`

GetSourceCodeManagementUriOk returns a tuple with the SourceCodeManagementUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeManagementUri

`func (o *Engagement) SetSourceCodeManagementUri(v string)`

SetSourceCodeManagementUri sets SourceCodeManagementUri field to given value.

### HasSourceCodeManagementUri

`func (o *Engagement) HasSourceCodeManagementUri() bool`

HasSourceCodeManagementUri returns a boolean if a field has been set.

### SetSourceCodeManagementUriNil

`func (o *Engagement) SetSourceCodeManagementUriNil(b bool)`

 SetSourceCodeManagementUriNil sets the value for SourceCodeManagementUri to be an explicit nil

### UnsetSourceCodeManagementUri
`func (o *Engagement) UnsetSourceCodeManagementUri()`

UnsetSourceCodeManagementUri ensures that no value is present for SourceCodeManagementUri, not even an explicit nil
### GetDeduplicationOnEngagement

`func (o *Engagement) GetDeduplicationOnEngagement() bool`

GetDeduplicationOnEngagement returns the DeduplicationOnEngagement field if non-nil, zero value otherwise.

### GetDeduplicationOnEngagementOk

`func (o *Engagement) GetDeduplicationOnEngagementOk() (*bool, bool)`

GetDeduplicationOnEngagementOk returns a tuple with the DeduplicationOnEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationOnEngagement

`func (o *Engagement) SetDeduplicationOnEngagement(v bool)`

SetDeduplicationOnEngagement sets DeduplicationOnEngagement field to given value.

### HasDeduplicationOnEngagement

`func (o *Engagement) HasDeduplicationOnEngagement() bool`

HasDeduplicationOnEngagement returns a boolean if a field has been set.

### GetLead

`func (o *Engagement) GetLead() int32`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *Engagement) GetLeadOk() (*int32, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *Engagement) SetLead(v int32)`

SetLead sets Lead field to given value.

### HasLead

`func (o *Engagement) HasLead() bool`

HasLead returns a boolean if a field has been set.

### SetLeadNil

`func (o *Engagement) SetLeadNil(b bool)`

 SetLeadNil sets the value for Lead to be an explicit nil

### UnsetLead
`func (o *Engagement) UnsetLead()`

UnsetLead ensures that no value is present for Lead, not even an explicit nil
### GetRequester

`func (o *Engagement) GetRequester() int32`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *Engagement) GetRequesterOk() (*int32, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *Engagement) SetRequester(v int32)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *Engagement) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *Engagement) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *Engagement) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetPreset

`func (o *Engagement) GetPreset() int32`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *Engagement) GetPresetOk() (*int32, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *Engagement) SetPreset(v int32)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *Engagement) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### SetPresetNil

`func (o *Engagement) SetPresetNil(b bool)`

 SetPresetNil sets the value for Preset to be an explicit nil

### UnsetPreset
`func (o *Engagement) UnsetPreset()`

UnsetPreset ensures that no value is present for Preset, not even an explicit nil
### GetReportType

`func (o *Engagement) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *Engagement) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *Engagement) SetReportType(v int32)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *Engagement) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### SetReportTypeNil

`func (o *Engagement) SetReportTypeNil(b bool)`

 SetReportTypeNil sets the value for ReportType to be an explicit nil

### UnsetReportType
`func (o *Engagement) UnsetReportType()`

UnsetReportType ensures that no value is present for ReportType, not even an explicit nil
### GetProduct

`func (o *Engagement) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Engagement) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Engagement) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetBuildServer

`func (o *Engagement) GetBuildServer() int32`

GetBuildServer returns the BuildServer field if non-nil, zero value otherwise.

### GetBuildServerOk

`func (o *Engagement) GetBuildServerOk() (*int32, bool)`

GetBuildServerOk returns a tuple with the BuildServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildServer

`func (o *Engagement) SetBuildServer(v int32)`

SetBuildServer sets BuildServer field to given value.

### HasBuildServer

`func (o *Engagement) HasBuildServer() bool`

HasBuildServer returns a boolean if a field has been set.

### SetBuildServerNil

`func (o *Engagement) SetBuildServerNil(b bool)`

 SetBuildServerNil sets the value for BuildServer to be an explicit nil

### UnsetBuildServer
`func (o *Engagement) UnsetBuildServer()`

UnsetBuildServer ensures that no value is present for BuildServer, not even an explicit nil
### GetSourceCodeManagementServer

`func (o *Engagement) GetSourceCodeManagementServer() int32`

GetSourceCodeManagementServer returns the SourceCodeManagementServer field if non-nil, zero value otherwise.

### GetSourceCodeManagementServerOk

`func (o *Engagement) GetSourceCodeManagementServerOk() (*int32, bool)`

GetSourceCodeManagementServerOk returns a tuple with the SourceCodeManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeManagementServer

`func (o *Engagement) SetSourceCodeManagementServer(v int32)`

SetSourceCodeManagementServer sets SourceCodeManagementServer field to given value.

### HasSourceCodeManagementServer

`func (o *Engagement) HasSourceCodeManagementServer() bool`

HasSourceCodeManagementServer returns a boolean if a field has been set.

### SetSourceCodeManagementServerNil

`func (o *Engagement) SetSourceCodeManagementServerNil(b bool)`

 SetSourceCodeManagementServerNil sets the value for SourceCodeManagementServer to be an explicit nil

### UnsetSourceCodeManagementServer
`func (o *Engagement) UnsetSourceCodeManagementServer()`

UnsetSourceCodeManagementServer ensures that no value is present for SourceCodeManagementServer, not even an explicit nil
### GetOrchestrationEngine

`func (o *Engagement) GetOrchestrationEngine() int32`

GetOrchestrationEngine returns the OrchestrationEngine field if non-nil, zero value otherwise.

### GetOrchestrationEngineOk

`func (o *Engagement) GetOrchestrationEngineOk() (*int32, bool)`

GetOrchestrationEngineOk returns a tuple with the OrchestrationEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationEngine

`func (o *Engagement) SetOrchestrationEngine(v int32)`

SetOrchestrationEngine sets OrchestrationEngine field to given value.

### HasOrchestrationEngine

`func (o *Engagement) HasOrchestrationEngine() bool`

HasOrchestrationEngine returns a boolean if a field has been set.

### SetOrchestrationEngineNil

`func (o *Engagement) SetOrchestrationEngineNil(b bool)`

 SetOrchestrationEngineNil sets the value for OrchestrationEngine to be an explicit nil

### UnsetOrchestrationEngine
`func (o *Engagement) UnsetOrchestrationEngine()`

UnsetOrchestrationEngine ensures that no value is present for OrchestrationEngine, not even an explicit nil
### GetNotes

`func (o *Engagement) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Engagement) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Engagement) SetNotes(v []Note)`

SetNotes sets Notes field to given value.


### GetFiles

`func (o *Engagement) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Engagement) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Engagement) SetFiles(v []File)`

SetFiles sets Files field to given value.


### GetRiskAcceptance

`func (o *Engagement) GetRiskAcceptance() []int32`

GetRiskAcceptance returns the RiskAcceptance field if non-nil, zero value otherwise.

### GetRiskAcceptanceOk

`func (o *Engagement) GetRiskAcceptanceOk() (*[]int32, bool)`

GetRiskAcceptanceOk returns a tuple with the RiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptance

`func (o *Engagement) SetRiskAcceptance(v []int32)`

SetRiskAcceptance sets RiskAcceptance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


