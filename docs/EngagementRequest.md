# EngagementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** | Version of the product the engagement tested. | [optional] 
**FirstContacted** | Pointer to **NullableString** |  | [optional] 
**TargetStart** | **string** |  | 
**TargetEnd** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Tracker** | Pointer to **NullableString** | Link to epic or ticket system with changes to version. | [optional] 
**TestStrategy** | Pointer to **NullableString** |  | [optional] 
**ThreatModel** | Pointer to **bool** |  | [optional] 
**ApiTest** | Pointer to **bool** |  | [optional] 
**PenTest** | Pointer to **bool** |  | [optional] 
**CheckList** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **NullableString** | * &#x60;Not Started&#x60; - Not Started * &#x60;Blocked&#x60; - Blocked * &#x60;Cancelled&#x60; - Cancelled * &#x60;Completed&#x60; - Completed * &#x60;In Progress&#x60; - In Progress * &#x60;On Hold&#x60; - On Hold * &#x60;Waiting for Resource&#x60; - Waiting for Resource | [optional] 
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

## Methods

### NewEngagementRequest

`func NewEngagementRequest(targetStart string, targetEnd string, product int32, ) *EngagementRequest`

NewEngagementRequest instantiates a new EngagementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementRequestWithDefaults

`func NewEngagementRequestWithDefaults() *EngagementRequest`

NewEngagementRequestWithDefaults instantiates a new EngagementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *EngagementRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EngagementRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EngagementRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EngagementRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *EngagementRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngagementRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngagementRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngagementRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EngagementRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EngagementRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EngagementRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EngagementRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EngagementRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EngagementRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EngagementRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EngagementRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVersion

`func (o *EngagementRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EngagementRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EngagementRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EngagementRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EngagementRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EngagementRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetFirstContacted

`func (o *EngagementRequest) GetFirstContacted() string`

GetFirstContacted returns the FirstContacted field if non-nil, zero value otherwise.

### GetFirstContactedOk

`func (o *EngagementRequest) GetFirstContactedOk() (*string, bool)`

GetFirstContactedOk returns a tuple with the FirstContacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstContacted

`func (o *EngagementRequest) SetFirstContacted(v string)`

SetFirstContacted sets FirstContacted field to given value.

### HasFirstContacted

`func (o *EngagementRequest) HasFirstContacted() bool`

HasFirstContacted returns a boolean if a field has been set.

### SetFirstContactedNil

`func (o *EngagementRequest) SetFirstContactedNil(b bool)`

 SetFirstContactedNil sets the value for FirstContacted to be an explicit nil

### UnsetFirstContacted
`func (o *EngagementRequest) UnsetFirstContacted()`

UnsetFirstContacted ensures that no value is present for FirstContacted, not even an explicit nil
### GetTargetStart

`func (o *EngagementRequest) GetTargetStart() string`

GetTargetStart returns the TargetStart field if non-nil, zero value otherwise.

### GetTargetStartOk

`func (o *EngagementRequest) GetTargetStartOk() (*string, bool)`

GetTargetStartOk returns a tuple with the TargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStart

`func (o *EngagementRequest) SetTargetStart(v string)`

SetTargetStart sets TargetStart field to given value.


### GetTargetEnd

`func (o *EngagementRequest) GetTargetEnd() string`

GetTargetEnd returns the TargetEnd field if non-nil, zero value otherwise.

### GetTargetEndOk

`func (o *EngagementRequest) GetTargetEndOk() (*string, bool)`

GetTargetEndOk returns a tuple with the TargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnd

`func (o *EngagementRequest) SetTargetEnd(v string)`

SetTargetEnd sets TargetEnd field to given value.


### GetReason

`func (o *EngagementRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EngagementRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EngagementRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EngagementRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *EngagementRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *EngagementRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetTracker

`func (o *EngagementRequest) GetTracker() string`

GetTracker returns the Tracker field if non-nil, zero value otherwise.

### GetTrackerOk

`func (o *EngagementRequest) GetTrackerOk() (*string, bool)`

GetTrackerOk returns a tuple with the Tracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracker

`func (o *EngagementRequest) SetTracker(v string)`

SetTracker sets Tracker field to given value.

### HasTracker

`func (o *EngagementRequest) HasTracker() bool`

HasTracker returns a boolean if a field has been set.

### SetTrackerNil

`func (o *EngagementRequest) SetTrackerNil(b bool)`

 SetTrackerNil sets the value for Tracker to be an explicit nil

### UnsetTracker
`func (o *EngagementRequest) UnsetTracker()`

UnsetTracker ensures that no value is present for Tracker, not even an explicit nil
### GetTestStrategy

`func (o *EngagementRequest) GetTestStrategy() string`

GetTestStrategy returns the TestStrategy field if non-nil, zero value otherwise.

### GetTestStrategyOk

`func (o *EngagementRequest) GetTestStrategyOk() (*string, bool)`

GetTestStrategyOk returns a tuple with the TestStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStrategy

`func (o *EngagementRequest) SetTestStrategy(v string)`

SetTestStrategy sets TestStrategy field to given value.

### HasTestStrategy

`func (o *EngagementRequest) HasTestStrategy() bool`

HasTestStrategy returns a boolean if a field has been set.

### SetTestStrategyNil

`func (o *EngagementRequest) SetTestStrategyNil(b bool)`

 SetTestStrategyNil sets the value for TestStrategy to be an explicit nil

### UnsetTestStrategy
`func (o *EngagementRequest) UnsetTestStrategy()`

UnsetTestStrategy ensures that no value is present for TestStrategy, not even an explicit nil
### GetThreatModel

`func (o *EngagementRequest) GetThreatModel() bool`

GetThreatModel returns the ThreatModel field if non-nil, zero value otherwise.

### GetThreatModelOk

`func (o *EngagementRequest) GetThreatModelOk() (*bool, bool)`

GetThreatModelOk returns a tuple with the ThreatModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatModel

`func (o *EngagementRequest) SetThreatModel(v bool)`

SetThreatModel sets ThreatModel field to given value.

### HasThreatModel

`func (o *EngagementRequest) HasThreatModel() bool`

HasThreatModel returns a boolean if a field has been set.

### GetApiTest

`func (o *EngagementRequest) GetApiTest() bool`

GetApiTest returns the ApiTest field if non-nil, zero value otherwise.

### GetApiTestOk

`func (o *EngagementRequest) GetApiTestOk() (*bool, bool)`

GetApiTestOk returns a tuple with the ApiTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTest

`func (o *EngagementRequest) SetApiTest(v bool)`

SetApiTest sets ApiTest field to given value.

### HasApiTest

`func (o *EngagementRequest) HasApiTest() bool`

HasApiTest returns a boolean if a field has been set.

### GetPenTest

`func (o *EngagementRequest) GetPenTest() bool`

GetPenTest returns the PenTest field if non-nil, zero value otherwise.

### GetPenTestOk

`func (o *EngagementRequest) GetPenTestOk() (*bool, bool)`

GetPenTestOk returns a tuple with the PenTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenTest

`func (o *EngagementRequest) SetPenTest(v bool)`

SetPenTest sets PenTest field to given value.

### HasPenTest

`func (o *EngagementRequest) HasPenTest() bool`

HasPenTest returns a boolean if a field has been set.

### GetCheckList

`func (o *EngagementRequest) GetCheckList() bool`

GetCheckList returns the CheckList field if non-nil, zero value otherwise.

### GetCheckListOk

`func (o *EngagementRequest) GetCheckListOk() (*bool, bool)`

GetCheckListOk returns a tuple with the CheckList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckList

`func (o *EngagementRequest) SetCheckList(v bool)`

SetCheckList sets CheckList field to given value.

### HasCheckList

`func (o *EngagementRequest) HasCheckList() bool`

HasCheckList returns a boolean if a field has been set.

### GetStatus

`func (o *EngagementRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngagementRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngagementRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngagementRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EngagementRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EngagementRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEngagementType

`func (o *EngagementRequest) GetEngagementType() string`

GetEngagementType returns the EngagementType field if non-nil, zero value otherwise.

### GetEngagementTypeOk

`func (o *EngagementRequest) GetEngagementTypeOk() (*string, bool)`

GetEngagementTypeOk returns a tuple with the EngagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementType

`func (o *EngagementRequest) SetEngagementType(v string)`

SetEngagementType sets EngagementType field to given value.

### HasEngagementType

`func (o *EngagementRequest) HasEngagementType() bool`

HasEngagementType returns a boolean if a field has been set.

### SetEngagementTypeNil

`func (o *EngagementRequest) SetEngagementTypeNil(b bool)`

 SetEngagementTypeNil sets the value for EngagementType to be an explicit nil

### UnsetEngagementType
`func (o *EngagementRequest) UnsetEngagementType()`

UnsetEngagementType ensures that no value is present for EngagementType, not even an explicit nil
### GetBuildId

`func (o *EngagementRequest) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *EngagementRequest) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *EngagementRequest) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *EngagementRequest) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *EngagementRequest) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *EngagementRequest) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *EngagementRequest) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *EngagementRequest) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *EngagementRequest) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *EngagementRequest) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *EngagementRequest) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *EngagementRequest) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *EngagementRequest) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *EngagementRequest) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *EngagementRequest) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *EngagementRequest) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *EngagementRequest) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *EngagementRequest) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetSourceCodeManagementUri

`func (o *EngagementRequest) GetSourceCodeManagementUri() string`

GetSourceCodeManagementUri returns the SourceCodeManagementUri field if non-nil, zero value otherwise.

### GetSourceCodeManagementUriOk

`func (o *EngagementRequest) GetSourceCodeManagementUriOk() (*string, bool)`

GetSourceCodeManagementUriOk returns a tuple with the SourceCodeManagementUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeManagementUri

`func (o *EngagementRequest) SetSourceCodeManagementUri(v string)`

SetSourceCodeManagementUri sets SourceCodeManagementUri field to given value.

### HasSourceCodeManagementUri

`func (o *EngagementRequest) HasSourceCodeManagementUri() bool`

HasSourceCodeManagementUri returns a boolean if a field has been set.

### SetSourceCodeManagementUriNil

`func (o *EngagementRequest) SetSourceCodeManagementUriNil(b bool)`

 SetSourceCodeManagementUriNil sets the value for SourceCodeManagementUri to be an explicit nil

### UnsetSourceCodeManagementUri
`func (o *EngagementRequest) UnsetSourceCodeManagementUri()`

UnsetSourceCodeManagementUri ensures that no value is present for SourceCodeManagementUri, not even an explicit nil
### GetDeduplicationOnEngagement

`func (o *EngagementRequest) GetDeduplicationOnEngagement() bool`

GetDeduplicationOnEngagement returns the DeduplicationOnEngagement field if non-nil, zero value otherwise.

### GetDeduplicationOnEngagementOk

`func (o *EngagementRequest) GetDeduplicationOnEngagementOk() (*bool, bool)`

GetDeduplicationOnEngagementOk returns a tuple with the DeduplicationOnEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicationOnEngagement

`func (o *EngagementRequest) SetDeduplicationOnEngagement(v bool)`

SetDeduplicationOnEngagement sets DeduplicationOnEngagement field to given value.

### HasDeduplicationOnEngagement

`func (o *EngagementRequest) HasDeduplicationOnEngagement() bool`

HasDeduplicationOnEngagement returns a boolean if a field has been set.

### GetLead

`func (o *EngagementRequest) GetLead() int32`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *EngagementRequest) GetLeadOk() (*int32, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *EngagementRequest) SetLead(v int32)`

SetLead sets Lead field to given value.

### HasLead

`func (o *EngagementRequest) HasLead() bool`

HasLead returns a boolean if a field has been set.

### SetLeadNil

`func (o *EngagementRequest) SetLeadNil(b bool)`

 SetLeadNil sets the value for Lead to be an explicit nil

### UnsetLead
`func (o *EngagementRequest) UnsetLead()`

UnsetLead ensures that no value is present for Lead, not even an explicit nil
### GetRequester

`func (o *EngagementRequest) GetRequester() int32`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *EngagementRequest) GetRequesterOk() (*int32, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *EngagementRequest) SetRequester(v int32)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *EngagementRequest) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *EngagementRequest) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *EngagementRequest) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetPreset

`func (o *EngagementRequest) GetPreset() int32`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *EngagementRequest) GetPresetOk() (*int32, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *EngagementRequest) SetPreset(v int32)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *EngagementRequest) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### SetPresetNil

`func (o *EngagementRequest) SetPresetNil(b bool)`

 SetPresetNil sets the value for Preset to be an explicit nil

### UnsetPreset
`func (o *EngagementRequest) UnsetPreset()`

UnsetPreset ensures that no value is present for Preset, not even an explicit nil
### GetReportType

`func (o *EngagementRequest) GetReportType() int32`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *EngagementRequest) GetReportTypeOk() (*int32, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *EngagementRequest) SetReportType(v int32)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *EngagementRequest) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### SetReportTypeNil

`func (o *EngagementRequest) SetReportTypeNil(b bool)`

 SetReportTypeNil sets the value for ReportType to be an explicit nil

### UnsetReportType
`func (o *EngagementRequest) UnsetReportType()`

UnsetReportType ensures that no value is present for ReportType, not even an explicit nil
### GetProduct

`func (o *EngagementRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EngagementRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EngagementRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetBuildServer

`func (o *EngagementRequest) GetBuildServer() int32`

GetBuildServer returns the BuildServer field if non-nil, zero value otherwise.

### GetBuildServerOk

`func (o *EngagementRequest) GetBuildServerOk() (*int32, bool)`

GetBuildServerOk returns a tuple with the BuildServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildServer

`func (o *EngagementRequest) SetBuildServer(v int32)`

SetBuildServer sets BuildServer field to given value.

### HasBuildServer

`func (o *EngagementRequest) HasBuildServer() bool`

HasBuildServer returns a boolean if a field has been set.

### SetBuildServerNil

`func (o *EngagementRequest) SetBuildServerNil(b bool)`

 SetBuildServerNil sets the value for BuildServer to be an explicit nil

### UnsetBuildServer
`func (o *EngagementRequest) UnsetBuildServer()`

UnsetBuildServer ensures that no value is present for BuildServer, not even an explicit nil
### GetSourceCodeManagementServer

`func (o *EngagementRequest) GetSourceCodeManagementServer() int32`

GetSourceCodeManagementServer returns the SourceCodeManagementServer field if non-nil, zero value otherwise.

### GetSourceCodeManagementServerOk

`func (o *EngagementRequest) GetSourceCodeManagementServerOk() (*int32, bool)`

GetSourceCodeManagementServerOk returns a tuple with the SourceCodeManagementServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeManagementServer

`func (o *EngagementRequest) SetSourceCodeManagementServer(v int32)`

SetSourceCodeManagementServer sets SourceCodeManagementServer field to given value.

### HasSourceCodeManagementServer

`func (o *EngagementRequest) HasSourceCodeManagementServer() bool`

HasSourceCodeManagementServer returns a boolean if a field has been set.

### SetSourceCodeManagementServerNil

`func (o *EngagementRequest) SetSourceCodeManagementServerNil(b bool)`

 SetSourceCodeManagementServerNil sets the value for SourceCodeManagementServer to be an explicit nil

### UnsetSourceCodeManagementServer
`func (o *EngagementRequest) UnsetSourceCodeManagementServer()`

UnsetSourceCodeManagementServer ensures that no value is present for SourceCodeManagementServer, not even an explicit nil
### GetOrchestrationEngine

`func (o *EngagementRequest) GetOrchestrationEngine() int32`

GetOrchestrationEngine returns the OrchestrationEngine field if non-nil, zero value otherwise.

### GetOrchestrationEngineOk

`func (o *EngagementRequest) GetOrchestrationEngineOk() (*int32, bool)`

GetOrchestrationEngineOk returns a tuple with the OrchestrationEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchestrationEngine

`func (o *EngagementRequest) SetOrchestrationEngine(v int32)`

SetOrchestrationEngine sets OrchestrationEngine field to given value.

### HasOrchestrationEngine

`func (o *EngagementRequest) HasOrchestrationEngine() bool`

HasOrchestrationEngine returns a boolean if a field has been set.

### SetOrchestrationEngineNil

`func (o *EngagementRequest) SetOrchestrationEngineNil(b bool)`

 SetOrchestrationEngineNil sets the value for OrchestrationEngine to be an explicit nil

### UnsetOrchestrationEngine
`func (o *EngagementRequest) UnsetOrchestrationEngine()`

UnsetOrchestrationEngine ensures that no value is present for OrchestrationEngine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


