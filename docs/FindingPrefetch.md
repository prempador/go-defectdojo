# FindingPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**AuthorIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**ConfigIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**CryptoIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**DataIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**DefectReviewRequestedBy** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**DuplicateFinding** | Pointer to [**map[string]Finding**](Finding.md) |  | [optional] [readonly] 
**EndpointSet** | Pointer to [**map[string]Endpoint**](Endpoint.md) |  | [optional] [readonly] 
**Endpoints** | Pointer to [**map[string]Endpoint**](Endpoint.md) |  | [optional] [readonly] 
**Files** | Pointer to [**map[string]File**](File.md) |  | [optional] [readonly] 
**FindingGroupSet** | Pointer to [**map[string]FindingGroup**](FindingGroup.md) |  | [optional] [readonly] 
**FoundBy** | Pointer to [**map[string]TestType**](TestType.md) |  | [optional] [readonly] 
**LastReviewedBy** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**MitigatedBy** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**Notes** | Pointer to [**map[string]Note**](Note.md) |  | [optional] [readonly] 
**OtherIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**Reporter** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**ReviewRequestedBy** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**Reviewers** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**RiskAcceptanceSet** | Pointer to [**map[string]RiskAcceptance**](RiskAcceptance.md) |  | [optional] [readonly] 
**SensitiveIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**SessionIssues** | Pointer to [**map[string]EngagementCheckList**](EngagementCheckList.md) |  | [optional] [readonly] 
**SonarqubeIssue** | Pointer to [**map[string]SonarqubeIssue**](SonarqubeIssue.md) |  | [optional] [readonly] 
**Test** | Pointer to [**map[string]Test**](Test.md) |  | [optional] [readonly] 
**TestImportSet** | Pointer to [**map[string]TestImport**](TestImport.md) |  | [optional] [readonly] 

## Methods

### NewFindingPrefetch

`func NewFindingPrefetch() *FindingPrefetch`

NewFindingPrefetch instantiates a new FindingPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingPrefetchWithDefaults

`func NewFindingPrefetchWithDefaults() *FindingPrefetch`

NewFindingPrefetchWithDefaults instantiates a new FindingPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthIssues

`func (o *FindingPrefetch) GetAuthIssues() map[string]EngagementCheckList`

GetAuthIssues returns the AuthIssues field if non-nil, zero value otherwise.

### GetAuthIssuesOk

`func (o *FindingPrefetch) GetAuthIssuesOk() (*map[string]EngagementCheckList, bool)`

GetAuthIssuesOk returns a tuple with the AuthIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthIssues

`func (o *FindingPrefetch) SetAuthIssues(v map[string]EngagementCheckList)`

SetAuthIssues sets AuthIssues field to given value.

### HasAuthIssues

`func (o *FindingPrefetch) HasAuthIssues() bool`

HasAuthIssues returns a boolean if a field has been set.

### GetAuthorIssues

`func (o *FindingPrefetch) GetAuthorIssues() map[string]EngagementCheckList`

GetAuthorIssues returns the AuthorIssues field if non-nil, zero value otherwise.

### GetAuthorIssuesOk

`func (o *FindingPrefetch) GetAuthorIssuesOk() (*map[string]EngagementCheckList, bool)`

GetAuthorIssuesOk returns a tuple with the AuthorIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIssues

`func (o *FindingPrefetch) SetAuthorIssues(v map[string]EngagementCheckList)`

SetAuthorIssues sets AuthorIssues field to given value.

### HasAuthorIssues

`func (o *FindingPrefetch) HasAuthorIssues() bool`

HasAuthorIssues returns a boolean if a field has been set.

### GetConfigIssues

`func (o *FindingPrefetch) GetConfigIssues() map[string]EngagementCheckList`

GetConfigIssues returns the ConfigIssues field if non-nil, zero value otherwise.

### GetConfigIssuesOk

`func (o *FindingPrefetch) GetConfigIssuesOk() (*map[string]EngagementCheckList, bool)`

GetConfigIssuesOk returns a tuple with the ConfigIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigIssues

`func (o *FindingPrefetch) SetConfigIssues(v map[string]EngagementCheckList)`

SetConfigIssues sets ConfigIssues field to given value.

### HasConfigIssues

`func (o *FindingPrefetch) HasConfigIssues() bool`

HasConfigIssues returns a boolean if a field has been set.

### GetCryptoIssues

`func (o *FindingPrefetch) GetCryptoIssues() map[string]EngagementCheckList`

GetCryptoIssues returns the CryptoIssues field if non-nil, zero value otherwise.

### GetCryptoIssuesOk

`func (o *FindingPrefetch) GetCryptoIssuesOk() (*map[string]EngagementCheckList, bool)`

GetCryptoIssuesOk returns a tuple with the CryptoIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoIssues

`func (o *FindingPrefetch) SetCryptoIssues(v map[string]EngagementCheckList)`

SetCryptoIssues sets CryptoIssues field to given value.

### HasCryptoIssues

`func (o *FindingPrefetch) HasCryptoIssues() bool`

HasCryptoIssues returns a boolean if a field has been set.

### GetDataIssues

`func (o *FindingPrefetch) GetDataIssues() map[string]EngagementCheckList`

GetDataIssues returns the DataIssues field if non-nil, zero value otherwise.

### GetDataIssuesOk

`func (o *FindingPrefetch) GetDataIssuesOk() (*map[string]EngagementCheckList, bool)`

GetDataIssuesOk returns a tuple with the DataIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIssues

`func (o *FindingPrefetch) SetDataIssues(v map[string]EngagementCheckList)`

SetDataIssues sets DataIssues field to given value.

### HasDataIssues

`func (o *FindingPrefetch) HasDataIssues() bool`

HasDataIssues returns a boolean if a field has been set.

### GetDefectReviewRequestedBy

`func (o *FindingPrefetch) GetDefectReviewRequestedBy() map[string]UserStub`

GetDefectReviewRequestedBy returns the DefectReviewRequestedBy field if non-nil, zero value otherwise.

### GetDefectReviewRequestedByOk

`func (o *FindingPrefetch) GetDefectReviewRequestedByOk() (*map[string]UserStub, bool)`

GetDefectReviewRequestedByOk returns a tuple with the DefectReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectReviewRequestedBy

`func (o *FindingPrefetch) SetDefectReviewRequestedBy(v map[string]UserStub)`

SetDefectReviewRequestedBy sets DefectReviewRequestedBy field to given value.

### HasDefectReviewRequestedBy

`func (o *FindingPrefetch) HasDefectReviewRequestedBy() bool`

HasDefectReviewRequestedBy returns a boolean if a field has been set.

### GetDuplicateFinding

`func (o *FindingPrefetch) GetDuplicateFinding() map[string]Finding`

GetDuplicateFinding returns the DuplicateFinding field if non-nil, zero value otherwise.

### GetDuplicateFindingOk

`func (o *FindingPrefetch) GetDuplicateFindingOk() (*map[string]Finding, bool)`

GetDuplicateFindingOk returns a tuple with the DuplicateFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateFinding

`func (o *FindingPrefetch) SetDuplicateFinding(v map[string]Finding)`

SetDuplicateFinding sets DuplicateFinding field to given value.

### HasDuplicateFinding

`func (o *FindingPrefetch) HasDuplicateFinding() bool`

HasDuplicateFinding returns a boolean if a field has been set.

### GetEndpointSet

`func (o *FindingPrefetch) GetEndpointSet() map[string]Endpoint`

GetEndpointSet returns the EndpointSet field if non-nil, zero value otherwise.

### GetEndpointSetOk

`func (o *FindingPrefetch) GetEndpointSetOk() (*map[string]Endpoint, bool)`

GetEndpointSetOk returns a tuple with the EndpointSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSet

`func (o *FindingPrefetch) SetEndpointSet(v map[string]Endpoint)`

SetEndpointSet sets EndpointSet field to given value.

### HasEndpointSet

`func (o *FindingPrefetch) HasEndpointSet() bool`

HasEndpointSet returns a boolean if a field has been set.

### GetEndpoints

`func (o *FindingPrefetch) GetEndpoints() map[string]Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *FindingPrefetch) GetEndpointsOk() (*map[string]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *FindingPrefetch) SetEndpoints(v map[string]Endpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *FindingPrefetch) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetFiles

`func (o *FindingPrefetch) GetFiles() map[string]File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FindingPrefetch) GetFilesOk() (*map[string]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FindingPrefetch) SetFiles(v map[string]File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FindingPrefetch) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetFindingGroupSet

`func (o *FindingPrefetch) GetFindingGroupSet() map[string]FindingGroup`

GetFindingGroupSet returns the FindingGroupSet field if non-nil, zero value otherwise.

### GetFindingGroupSetOk

`func (o *FindingPrefetch) GetFindingGroupSetOk() (*map[string]FindingGroup, bool)`

GetFindingGroupSetOk returns a tuple with the FindingGroupSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingGroupSet

`func (o *FindingPrefetch) SetFindingGroupSet(v map[string]FindingGroup)`

SetFindingGroupSet sets FindingGroupSet field to given value.

### HasFindingGroupSet

`func (o *FindingPrefetch) HasFindingGroupSet() bool`

HasFindingGroupSet returns a boolean if a field has been set.

### GetFoundBy

`func (o *FindingPrefetch) GetFoundBy() map[string]TestType`

GetFoundBy returns the FoundBy field if non-nil, zero value otherwise.

### GetFoundByOk

`func (o *FindingPrefetch) GetFoundByOk() (*map[string]TestType, bool)`

GetFoundByOk returns a tuple with the FoundBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundBy

`func (o *FindingPrefetch) SetFoundBy(v map[string]TestType)`

SetFoundBy sets FoundBy field to given value.

### HasFoundBy

`func (o *FindingPrefetch) HasFoundBy() bool`

HasFoundBy returns a boolean if a field has been set.

### GetLastReviewedBy

`func (o *FindingPrefetch) GetLastReviewedBy() map[string]UserStub`

GetLastReviewedBy returns the LastReviewedBy field if non-nil, zero value otherwise.

### GetLastReviewedByOk

`func (o *FindingPrefetch) GetLastReviewedByOk() (*map[string]UserStub, bool)`

GetLastReviewedByOk returns a tuple with the LastReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewedBy

`func (o *FindingPrefetch) SetLastReviewedBy(v map[string]UserStub)`

SetLastReviewedBy sets LastReviewedBy field to given value.

### HasLastReviewedBy

`func (o *FindingPrefetch) HasLastReviewedBy() bool`

HasLastReviewedBy returns a boolean if a field has been set.

### GetMitigatedBy

`func (o *FindingPrefetch) GetMitigatedBy() map[string]UserStub`

GetMitigatedBy returns the MitigatedBy field if non-nil, zero value otherwise.

### GetMitigatedByOk

`func (o *FindingPrefetch) GetMitigatedByOk() (*map[string]UserStub, bool)`

GetMitigatedByOk returns a tuple with the MitigatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedBy

`func (o *FindingPrefetch) SetMitigatedBy(v map[string]UserStub)`

SetMitigatedBy sets MitigatedBy field to given value.

### HasMitigatedBy

`func (o *FindingPrefetch) HasMitigatedBy() bool`

HasMitigatedBy returns a boolean if a field has been set.

### GetNotes

`func (o *FindingPrefetch) GetNotes() map[string]Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FindingPrefetch) GetNotesOk() (*map[string]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FindingPrefetch) SetNotes(v map[string]Note)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *FindingPrefetch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOtherIssues

`func (o *FindingPrefetch) GetOtherIssues() map[string]EngagementCheckList`

GetOtherIssues returns the OtherIssues field if non-nil, zero value otherwise.

### GetOtherIssuesOk

`func (o *FindingPrefetch) GetOtherIssuesOk() (*map[string]EngagementCheckList, bool)`

GetOtherIssuesOk returns a tuple with the OtherIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherIssues

`func (o *FindingPrefetch) SetOtherIssues(v map[string]EngagementCheckList)`

SetOtherIssues sets OtherIssues field to given value.

### HasOtherIssues

`func (o *FindingPrefetch) HasOtherIssues() bool`

HasOtherIssues returns a boolean if a field has been set.

### GetReporter

`func (o *FindingPrefetch) GetReporter() map[string]UserStub`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *FindingPrefetch) GetReporterOk() (*map[string]UserStub, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *FindingPrefetch) SetReporter(v map[string]UserStub)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *FindingPrefetch) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetReviewRequestedBy

`func (o *FindingPrefetch) GetReviewRequestedBy() map[string]UserStub`

GetReviewRequestedBy returns the ReviewRequestedBy field if non-nil, zero value otherwise.

### GetReviewRequestedByOk

`func (o *FindingPrefetch) GetReviewRequestedByOk() (*map[string]UserStub, bool)`

GetReviewRequestedByOk returns a tuple with the ReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequestedBy

`func (o *FindingPrefetch) SetReviewRequestedBy(v map[string]UserStub)`

SetReviewRequestedBy sets ReviewRequestedBy field to given value.

### HasReviewRequestedBy

`func (o *FindingPrefetch) HasReviewRequestedBy() bool`

HasReviewRequestedBy returns a boolean if a field has been set.

### GetReviewers

`func (o *FindingPrefetch) GetReviewers() map[string]UserStub`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *FindingPrefetch) GetReviewersOk() (*map[string]UserStub, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *FindingPrefetch) SetReviewers(v map[string]UserStub)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *FindingPrefetch) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetRiskAcceptanceSet

`func (o *FindingPrefetch) GetRiskAcceptanceSet() map[string]RiskAcceptance`

GetRiskAcceptanceSet returns the RiskAcceptanceSet field if non-nil, zero value otherwise.

### GetRiskAcceptanceSetOk

`func (o *FindingPrefetch) GetRiskAcceptanceSetOk() (*map[string]RiskAcceptance, bool)`

GetRiskAcceptanceSetOk returns a tuple with the RiskAcceptanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceSet

`func (o *FindingPrefetch) SetRiskAcceptanceSet(v map[string]RiskAcceptance)`

SetRiskAcceptanceSet sets RiskAcceptanceSet field to given value.

### HasRiskAcceptanceSet

`func (o *FindingPrefetch) HasRiskAcceptanceSet() bool`

HasRiskAcceptanceSet returns a boolean if a field has been set.

### GetSensitiveIssues

`func (o *FindingPrefetch) GetSensitiveIssues() map[string]EngagementCheckList`

GetSensitiveIssues returns the SensitiveIssues field if non-nil, zero value otherwise.

### GetSensitiveIssuesOk

`func (o *FindingPrefetch) GetSensitiveIssuesOk() (*map[string]EngagementCheckList, bool)`

GetSensitiveIssuesOk returns a tuple with the SensitiveIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveIssues

`func (o *FindingPrefetch) SetSensitiveIssues(v map[string]EngagementCheckList)`

SetSensitiveIssues sets SensitiveIssues field to given value.

### HasSensitiveIssues

`func (o *FindingPrefetch) HasSensitiveIssues() bool`

HasSensitiveIssues returns a boolean if a field has been set.

### GetSessionIssues

`func (o *FindingPrefetch) GetSessionIssues() map[string]EngagementCheckList`

GetSessionIssues returns the SessionIssues field if non-nil, zero value otherwise.

### GetSessionIssuesOk

`func (o *FindingPrefetch) GetSessionIssuesOk() (*map[string]EngagementCheckList, bool)`

GetSessionIssuesOk returns a tuple with the SessionIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIssues

`func (o *FindingPrefetch) SetSessionIssues(v map[string]EngagementCheckList)`

SetSessionIssues sets SessionIssues field to given value.

### HasSessionIssues

`func (o *FindingPrefetch) HasSessionIssues() bool`

HasSessionIssues returns a boolean if a field has been set.

### GetSonarqubeIssue

`func (o *FindingPrefetch) GetSonarqubeIssue() map[string]SonarqubeIssue`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *FindingPrefetch) GetSonarqubeIssueOk() (*map[string]SonarqubeIssue, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *FindingPrefetch) SetSonarqubeIssue(v map[string]SonarqubeIssue)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.

### HasSonarqubeIssue

`func (o *FindingPrefetch) HasSonarqubeIssue() bool`

HasSonarqubeIssue returns a boolean if a field has been set.

### GetTest

`func (o *FindingPrefetch) GetTest() map[string]Test`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *FindingPrefetch) GetTestOk() (*map[string]Test, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *FindingPrefetch) SetTest(v map[string]Test)`

SetTest sets Test field to given value.

### HasTest

`func (o *FindingPrefetch) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetTestImportSet

`func (o *FindingPrefetch) GetTestImportSet() map[string]TestImport`

GetTestImportSet returns the TestImportSet field if non-nil, zero value otherwise.

### GetTestImportSetOk

`func (o *FindingPrefetch) GetTestImportSetOk() (*map[string]TestImport, bool)`

GetTestImportSetOk returns a tuple with the TestImportSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestImportSet

`func (o *FindingPrefetch) SetTestImportSet(v map[string]TestImport)`

SetTestImportSet sets TestImportSet field to given value.

### HasTestImportSet

`func (o *FindingPrefetch) HasTestImportSet() bool`

HasTestImportSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


