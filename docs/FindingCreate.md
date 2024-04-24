# FindingCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Notes** | **[]int32** |  | [readonly] 
**Test** | **int32** |  | 
**ThreadId** | Pointer to **int32** |  | [optional] [default to 0]
**FoundBy** | **[]int32** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**PushToJira** | Pointer to **bool** |  | [optional] [default to false]
**VulnerabilityIds** | Pointer to [**[]VulnerabilityId**](VulnerabilityId.md) |  | [optional] 
**Reporter** | Pointer to **int32** |  | [optional] 
**Title** | **string** | A short description of the flaw. | 
**Date** | Pointer to **string** | The date the flaw was discovered. | [optional] 
**SlaStartDate** | Pointer to **NullableString** | (readonly)The date used as start date for SLA calculation. Set by expiring risk acceptances. Empty by default, causing a fallback to &#39;date&#39;. | [optional] 
**SlaExpirationDate** | Pointer to **NullableString** | (readonly)The date SLA expires for this finding. Empty by default, causing a fallback to &#39;date&#39;. | [optional] 
**Cwe** | Pointer to **NullableInt32** | The CWE number associated with this flaw. | [optional] 
**EpssScore** | Pointer to **NullableFloat64** | EPSS score for the CVE. Describes how likely it is the vulnerability will be exploited in the next 30 days. | [optional] 
**EpssPercentile** | Pointer to **NullableFloat64** | EPSS percentile for the CVE. Describes how many CVEs are scored at or below this one. | [optional] 
**Cvssv3** | Pointer to **NullableString** | Common Vulnerability Scoring System version 3 (CVSSv3) score associated with this flaw. | [optional] 
**Cvssv3Score** | Pointer to **NullableFloat64** | Numerical CVSSv3 score for the vulnerability. If the vector is given, the score is updated while saving the finding. The value must be between 0-10. | [optional] 
**Severity** | **string** | The severity level of this flaw (Critical, High, Medium, Low, Info). | 
**Description** | **string** | Longer more descriptive information about the flaw. | 
**Mitigation** | Pointer to **NullableString** | Text describing how to best fix the flaw. | [optional] 
**Impact** | Pointer to **NullableString** | Text describing the impact this flaw has on systems, products, enterprise, etc. | [optional] 
**StepsToReproduce** | Pointer to **NullableString** | Text describing the steps that must be followed in order to reproduce the flaw / bug. | [optional] 
**SeverityJustification** | Pointer to **NullableString** | Text describing why a certain severity was associated with this flaw. | [optional] 
**References** | Pointer to **NullableString** | The external documentation available for this flaw. | [optional] 
**Active** | **bool** | Denotes if this flaw is active or not. | 
**Verified** | **bool** | Denotes if this flaw has been manually verified by the tester. | 
**FalseP** | Pointer to **bool** | Denotes if this flaw has been deemed a false positive by the tester. | [optional] 
**Duplicate** | Pointer to **bool** | Denotes if this flaw is a duplicate of other flaws reported. | [optional] 
**OutOfScope** | Pointer to **bool** | Denotes if this flaw falls outside the scope of the test and/or engagement. | [optional] 
**RiskAccepted** | Pointer to **bool** | Denotes if this finding has been marked as an accepted risk. | [optional] 
**UnderReview** | Pointer to **bool** | Denotes is this flaw is currently being reviewed. | [optional] 
**LastStatusUpdate** | **NullableTime** | Timestamp of latest status update (change in status related fields). | [readonly] 
**UnderDefectReview** | Pointer to **bool** | Denotes if this finding is under defect review. | [optional] 
**IsMitigated** | Pointer to **bool** | Denotes if this flaw has been fixed. | [optional] 
**Mitigated** | **NullableTime** | Denotes if this flaw has been fixed by storing the date it was fixed. | [readonly] 
**NumericalSeverity** | **string** | The numerical representation of the severity (S0, S1, S2, S3, S4). | 
**LastReviewed** | **NullableTime** | Provides the date the flaw was last &#39;touched&#39; by a tester. | [readonly] 
**Param** | **NullableString** | Parameter used to trigger the issue (DAST). | [readonly] 
**Payload** | **NullableString** | Payload used to attack the service / application and trigger the bug / problem. | [readonly] 
**HashCode** | **NullableString** | A hash over a configurable set of fields that is used for findings deduplication. | [readonly] 
**Line** | Pointer to **NullableInt32** | Source line number of the attack vector. | [optional] 
**FilePath** | Pointer to **NullableString** | Identified file(s) containing the flaw. | [optional] 
**ComponentName** | Pointer to **NullableString** | Name of the affected component (library name, part of a system, ...). | [optional] 
**ComponentVersion** | Pointer to **NullableString** | Version of the affected component. | [optional] 
**StaticFinding** | Pointer to **bool** | Flaw has been detected from a Static Application Security Testing tool (SAST). | [optional] 
**DynamicFinding** | Pointer to **bool** | Flaw has been detected from a Dynamic Application Security Testing tool (DAST). | [optional] 
**Created** | **NullableTime** | The date the finding was created inside DefectDojo. | [readonly] 
**ScannerConfidence** | **NullableInt32** | Confidence level of vulnerability which is supplied by the scanner. | [readonly] 
**UniqueIdFromTool** | Pointer to **NullableString** | Vulnerability technical id from the source tool. Allows to track unique vulnerabilities. | [optional] 
**VulnIdFromTool** | Pointer to **NullableString** | Non-unique technical id from the source tool associated with the vulnerability type. | [optional] 
**SastSourceObject** | Pointer to **NullableString** | Source object (variable, function...) of the attack vector. | [optional] 
**SastSinkObject** | Pointer to **NullableString** | Sink object (variable, function...) of the attack vector. | [optional] 
**SastSourceLine** | Pointer to **NullableInt32** | Source line number of the attack vector. | [optional] 
**SastSourceFilePath** | Pointer to **NullableString** | Source file path of the attack vector. | [optional] 
**NbOccurences** | Pointer to **NullableInt32** | Number of occurences in the source tool when several vulnerabilites were found and aggregated by the scanner. | [optional] 
**PublishDate** | Pointer to **NullableString** | Date when this vulnerability was made publicly available. | [optional] 
**Service** | Pointer to **NullableString** | A service is a self-contained piece of functionality within a Product. This is an optional field which is used in deduplication of findings when set. | [optional] 
**PlannedRemediationDate** | Pointer to **NullableString** | The date the flaw is expected to be remediated. | [optional] 
**PlannedRemediationVersion** | Pointer to **NullableString** | The target version when the vulnerability should be fixed / remediated | [optional] 
**EffortForFixing** | Pointer to **NullableString** | Effort for fixing / remediating the vulnerability (Low, Medium, High) | [optional] 
**DuplicateFinding** | **NullableInt32** | Link to the original finding if this finding is a duplicate. | [readonly] 
**ReviewRequestedBy** | Pointer to **NullableInt32** | Documents who requested a review for this finding. | [optional] 
**DefectReviewRequestedBy** | Pointer to **NullableInt32** | Documents who requested a defect review for this flaw. | [optional] 
**MitigatedBy** | **NullableInt32** | Documents who has marked this flaw as fixed. | [readonly] 
**LastReviewedBy** | **NullableInt32** | Provides the person who last reviewed the flaw. | [readonly] 
**SonarqubeIssue** | Pointer to **NullableInt32** | The SonarQube issue associated with this finding. | [optional] 
**Endpoints** | **[]int32** | The hosts within the product that are susceptible to this flaw. + The status of the endpoint associated with this flaw (Vulnerable, Mitigated, ...). | [readonly] 
**Reviewers** | Pointer to **[]int32** | Documents who reviewed the flaw. | [optional] 
**Files** | **[]int32** | Files(s) related to the flaw. | [readonly] 

## Methods

### NewFindingCreate

`func NewFindingCreate(id int32, notes []*int32, test int32, foundBy []int32, title string, severity string, description string, active bool, verified bool, lastStatusUpdate NullableTime, mitigated NullableTime, numericalSeverity string, lastReviewed NullableTime, param NullableString, payload NullableString, hashCode NullableString, created NullableTime, scannerConfidence NullableInt32, duplicateFinding NullableInt32, mitigatedBy NullableInt32, lastReviewedBy NullableInt32, endpoints []int32, files []int32, ) *FindingCreate`

NewFindingCreate instantiates a new FindingCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingCreateWithDefaults

`func NewFindingCreateWithDefaults() *FindingCreate`

NewFindingCreateWithDefaults instantiates a new FindingCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FindingCreate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingCreate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingCreate) SetId(v int32)`

SetId sets Id field to given value.


### GetNotes

`func (o *FindingCreate) GetNotes() []*int32`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FindingCreate) GetNotesOk() (*[]*int32, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FindingCreate) SetNotes(v []*int32)`

SetNotes sets Notes field to given value.


### GetTest

`func (o *FindingCreate) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *FindingCreate) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *FindingCreate) SetTest(v int32)`

SetTest sets Test field to given value.


### GetThreadId

`func (o *FindingCreate) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *FindingCreate) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *FindingCreate) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *FindingCreate) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetFoundBy

`func (o *FindingCreate) GetFoundBy() []int32`

GetFoundBy returns the FoundBy field if non-nil, zero value otherwise.

### GetFoundByOk

`func (o *FindingCreate) GetFoundByOk() (*[]int32, bool)`

GetFoundByOk returns a tuple with the FoundBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundBy

`func (o *FindingCreate) SetFoundBy(v []int32)`

SetFoundBy sets FoundBy field to given value.


### GetUrl

`func (o *FindingCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FindingCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FindingCreate) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FindingCreate) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FindingCreate) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FindingCreate) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTags

`func (o *FindingCreate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindingCreate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindingCreate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindingCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPushToJira

`func (o *FindingCreate) GetPushToJira() bool`

GetPushToJira returns the PushToJira field if non-nil, zero value otherwise.

### GetPushToJiraOk

`func (o *FindingCreate) GetPushToJiraOk() (*bool, bool)`

GetPushToJiraOk returns a tuple with the PushToJira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToJira

`func (o *FindingCreate) SetPushToJira(v bool)`

SetPushToJira sets PushToJira field to given value.

### HasPushToJira

`func (o *FindingCreate) HasPushToJira() bool`

HasPushToJira returns a boolean if a field has been set.

### GetVulnerabilityIds

`func (o *FindingCreate) GetVulnerabilityIds() []VulnerabilityId`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *FindingCreate) GetVulnerabilityIdsOk() (*[]VulnerabilityId, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *FindingCreate) SetVulnerabilityIds(v []VulnerabilityId)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *FindingCreate) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetReporter

`func (o *FindingCreate) GetReporter() int32`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *FindingCreate) GetReporterOk() (*int32, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *FindingCreate) SetReporter(v int32)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *FindingCreate) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetTitle

`func (o *FindingCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDate

`func (o *FindingCreate) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FindingCreate) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FindingCreate) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FindingCreate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSlaStartDate

`func (o *FindingCreate) GetSlaStartDate() string`

GetSlaStartDate returns the SlaStartDate field if non-nil, zero value otherwise.

### GetSlaStartDateOk

`func (o *FindingCreate) GetSlaStartDateOk() (*string, bool)`

GetSlaStartDateOk returns a tuple with the SlaStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStartDate

`func (o *FindingCreate) SetSlaStartDate(v string)`

SetSlaStartDate sets SlaStartDate field to given value.

### HasSlaStartDate

`func (o *FindingCreate) HasSlaStartDate() bool`

HasSlaStartDate returns a boolean if a field has been set.

### SetSlaStartDateNil

`func (o *FindingCreate) SetSlaStartDateNil(b bool)`

 SetSlaStartDateNil sets the value for SlaStartDate to be an explicit nil

### UnsetSlaStartDate
`func (o *FindingCreate) UnsetSlaStartDate()`

UnsetSlaStartDate ensures that no value is present for SlaStartDate, not even an explicit nil
### GetSlaExpirationDate

`func (o *FindingCreate) GetSlaExpirationDate() string`

GetSlaExpirationDate returns the SlaExpirationDate field if non-nil, zero value otherwise.

### GetSlaExpirationDateOk

`func (o *FindingCreate) GetSlaExpirationDateOk() (*string, bool)`

GetSlaExpirationDateOk returns a tuple with the SlaExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaExpirationDate

`func (o *FindingCreate) SetSlaExpirationDate(v string)`

SetSlaExpirationDate sets SlaExpirationDate field to given value.

### HasSlaExpirationDate

`func (o *FindingCreate) HasSlaExpirationDate() bool`

HasSlaExpirationDate returns a boolean if a field has been set.

### SetSlaExpirationDateNil

`func (o *FindingCreate) SetSlaExpirationDateNil(b bool)`

 SetSlaExpirationDateNil sets the value for SlaExpirationDate to be an explicit nil

### UnsetSlaExpirationDate
`func (o *FindingCreate) UnsetSlaExpirationDate()`

UnsetSlaExpirationDate ensures that no value is present for SlaExpirationDate, not even an explicit nil
### GetCwe

`func (o *FindingCreate) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *FindingCreate) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *FindingCreate) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *FindingCreate) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *FindingCreate) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *FindingCreate) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetEpssScore

`func (o *FindingCreate) GetEpssScore() float64`

GetEpssScore returns the EpssScore field if non-nil, zero value otherwise.

### GetEpssScoreOk

`func (o *FindingCreate) GetEpssScoreOk() (*float64, bool)`

GetEpssScoreOk returns a tuple with the EpssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpssScore

`func (o *FindingCreate) SetEpssScore(v float64)`

SetEpssScore sets EpssScore field to given value.

### HasEpssScore

`func (o *FindingCreate) HasEpssScore() bool`

HasEpssScore returns a boolean if a field has been set.

### SetEpssScoreNil

`func (o *FindingCreate) SetEpssScoreNil(b bool)`

 SetEpssScoreNil sets the value for EpssScore to be an explicit nil

### UnsetEpssScore
`func (o *FindingCreate) UnsetEpssScore()`

UnsetEpssScore ensures that no value is present for EpssScore, not even an explicit nil
### GetEpssPercentile

`func (o *FindingCreate) GetEpssPercentile() float64`

GetEpssPercentile returns the EpssPercentile field if non-nil, zero value otherwise.

### GetEpssPercentileOk

`func (o *FindingCreate) GetEpssPercentileOk() (*float64, bool)`

GetEpssPercentileOk returns a tuple with the EpssPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpssPercentile

`func (o *FindingCreate) SetEpssPercentile(v float64)`

SetEpssPercentile sets EpssPercentile field to given value.

### HasEpssPercentile

`func (o *FindingCreate) HasEpssPercentile() bool`

HasEpssPercentile returns a boolean if a field has been set.

### SetEpssPercentileNil

`func (o *FindingCreate) SetEpssPercentileNil(b bool)`

 SetEpssPercentileNil sets the value for EpssPercentile to be an explicit nil

### UnsetEpssPercentile
`func (o *FindingCreate) UnsetEpssPercentile()`

UnsetEpssPercentile ensures that no value is present for EpssPercentile, not even an explicit nil
### GetCvssv3

`func (o *FindingCreate) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *FindingCreate) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *FindingCreate) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *FindingCreate) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *FindingCreate) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *FindingCreate) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetCvssv3Score

`func (o *FindingCreate) GetCvssv3Score() float64`

GetCvssv3Score returns the Cvssv3Score field if non-nil, zero value otherwise.

### GetCvssv3ScoreOk

`func (o *FindingCreate) GetCvssv3ScoreOk() (*float64, bool)`

GetCvssv3ScoreOk returns a tuple with the Cvssv3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3Score

`func (o *FindingCreate) SetCvssv3Score(v float64)`

SetCvssv3Score sets Cvssv3Score field to given value.

### HasCvssv3Score

`func (o *FindingCreate) HasCvssv3Score() bool`

HasCvssv3Score returns a boolean if a field has been set.

### SetCvssv3ScoreNil

`func (o *FindingCreate) SetCvssv3ScoreNil(b bool)`

 SetCvssv3ScoreNil sets the value for Cvssv3Score to be an explicit nil

### UnsetCvssv3Score
`func (o *FindingCreate) UnsetCvssv3Score()`

UnsetCvssv3Score ensures that no value is present for Cvssv3Score, not even an explicit nil
### GetSeverity

`func (o *FindingCreate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingCreate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingCreate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetDescription

`func (o *FindingCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMitigation

`func (o *FindingCreate) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *FindingCreate) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *FindingCreate) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *FindingCreate) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *FindingCreate) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *FindingCreate) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *FindingCreate) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *FindingCreate) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *FindingCreate) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *FindingCreate) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *FindingCreate) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *FindingCreate) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetStepsToReproduce

`func (o *FindingCreate) GetStepsToReproduce() string`

GetStepsToReproduce returns the StepsToReproduce field if non-nil, zero value otherwise.

### GetStepsToReproduceOk

`func (o *FindingCreate) GetStepsToReproduceOk() (*string, bool)`

GetStepsToReproduceOk returns a tuple with the StepsToReproduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsToReproduce

`func (o *FindingCreate) SetStepsToReproduce(v string)`

SetStepsToReproduce sets StepsToReproduce field to given value.

### HasStepsToReproduce

`func (o *FindingCreate) HasStepsToReproduce() bool`

HasStepsToReproduce returns a boolean if a field has been set.

### SetStepsToReproduceNil

`func (o *FindingCreate) SetStepsToReproduceNil(b bool)`

 SetStepsToReproduceNil sets the value for StepsToReproduce to be an explicit nil

### UnsetStepsToReproduce
`func (o *FindingCreate) UnsetStepsToReproduce()`

UnsetStepsToReproduce ensures that no value is present for StepsToReproduce, not even an explicit nil
### GetSeverityJustification

`func (o *FindingCreate) GetSeverityJustification() string`

GetSeverityJustification returns the SeverityJustification field if non-nil, zero value otherwise.

### GetSeverityJustificationOk

`func (o *FindingCreate) GetSeverityJustificationOk() (*string, bool)`

GetSeverityJustificationOk returns a tuple with the SeverityJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityJustification

`func (o *FindingCreate) SetSeverityJustification(v string)`

SetSeverityJustification sets SeverityJustification field to given value.

### HasSeverityJustification

`func (o *FindingCreate) HasSeverityJustification() bool`

HasSeverityJustification returns a boolean if a field has been set.

### SetSeverityJustificationNil

`func (o *FindingCreate) SetSeverityJustificationNil(b bool)`

 SetSeverityJustificationNil sets the value for SeverityJustification to be an explicit nil

### UnsetSeverityJustification
`func (o *FindingCreate) UnsetSeverityJustification()`

UnsetSeverityJustification ensures that no value is present for SeverityJustification, not even an explicit nil
### GetReferences

`func (o *FindingCreate) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FindingCreate) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FindingCreate) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *FindingCreate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *FindingCreate) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *FindingCreate) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetActive

`func (o *FindingCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FindingCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FindingCreate) SetActive(v bool)`

SetActive sets Active field to given value.


### GetVerified

`func (o *FindingCreate) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *FindingCreate) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *FindingCreate) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetFalseP

`func (o *FindingCreate) GetFalseP() bool`

GetFalseP returns the FalseP field if non-nil, zero value otherwise.

### GetFalsePOk

`func (o *FindingCreate) GetFalsePOk() (*bool, bool)`

GetFalsePOk returns a tuple with the FalseP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseP

`func (o *FindingCreate) SetFalseP(v bool)`

SetFalseP sets FalseP field to given value.

### HasFalseP

`func (o *FindingCreate) HasFalseP() bool`

HasFalseP returns a boolean if a field has been set.

### GetDuplicate

`func (o *FindingCreate) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *FindingCreate) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *FindingCreate) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *FindingCreate) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetOutOfScope

`func (o *FindingCreate) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *FindingCreate) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *FindingCreate) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *FindingCreate) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *FindingCreate) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *FindingCreate) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *FindingCreate) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *FindingCreate) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetUnderReview

`func (o *FindingCreate) GetUnderReview() bool`

GetUnderReview returns the UnderReview field if non-nil, zero value otherwise.

### GetUnderReviewOk

`func (o *FindingCreate) GetUnderReviewOk() (*bool, bool)`

GetUnderReviewOk returns a tuple with the UnderReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview

`func (o *FindingCreate) SetUnderReview(v bool)`

SetUnderReview sets UnderReview field to given value.

### HasUnderReview

`func (o *FindingCreate) HasUnderReview() bool`

HasUnderReview returns a boolean if a field has been set.

### GetLastStatusUpdate

`func (o *FindingCreate) GetLastStatusUpdate() time.Time`

GetLastStatusUpdate returns the LastStatusUpdate field if non-nil, zero value otherwise.

### GetLastStatusUpdateOk

`func (o *FindingCreate) GetLastStatusUpdateOk() (*time.Time, bool)`

GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdate

`func (o *FindingCreate) SetLastStatusUpdate(v time.Time)`

SetLastStatusUpdate sets LastStatusUpdate field to given value.


### SetLastStatusUpdateNil

`func (o *FindingCreate) SetLastStatusUpdateNil(b bool)`

 SetLastStatusUpdateNil sets the value for LastStatusUpdate to be an explicit nil

### UnsetLastStatusUpdate
`func (o *FindingCreate) UnsetLastStatusUpdate()`

UnsetLastStatusUpdate ensures that no value is present for LastStatusUpdate, not even an explicit nil
### GetUnderDefectReview

`func (o *FindingCreate) GetUnderDefectReview() bool`

GetUnderDefectReview returns the UnderDefectReview field if non-nil, zero value otherwise.

### GetUnderDefectReviewOk

`func (o *FindingCreate) GetUnderDefectReviewOk() (*bool, bool)`

GetUnderDefectReviewOk returns a tuple with the UnderDefectReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderDefectReview

`func (o *FindingCreate) SetUnderDefectReview(v bool)`

SetUnderDefectReview sets UnderDefectReview field to given value.

### HasUnderDefectReview

`func (o *FindingCreate) HasUnderDefectReview() bool`

HasUnderDefectReview returns a boolean if a field has been set.

### GetIsMitigated

`func (o *FindingCreate) GetIsMitigated() bool`

GetIsMitigated returns the IsMitigated field if non-nil, zero value otherwise.

### GetIsMitigatedOk

`func (o *FindingCreate) GetIsMitigatedOk() (*bool, bool)`

GetIsMitigatedOk returns a tuple with the IsMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMitigated

`func (o *FindingCreate) SetIsMitigated(v bool)`

SetIsMitigated sets IsMitigated field to given value.

### HasIsMitigated

`func (o *FindingCreate) HasIsMitigated() bool`

HasIsMitigated returns a boolean if a field has been set.

### GetMitigated

`func (o *FindingCreate) GetMitigated() time.Time`

GetMitigated returns the Mitigated field if non-nil, zero value otherwise.

### GetMitigatedOk

`func (o *FindingCreate) GetMitigatedOk() (*time.Time, bool)`

GetMitigatedOk returns a tuple with the Mitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigated

`func (o *FindingCreate) SetMitigated(v time.Time)`

SetMitigated sets Mitigated field to given value.


### SetMitigatedNil

`func (o *FindingCreate) SetMitigatedNil(b bool)`

 SetMitigatedNil sets the value for Mitigated to be an explicit nil

### UnsetMitigated
`func (o *FindingCreate) UnsetMitigated()`

UnsetMitigated ensures that no value is present for Mitigated, not even an explicit nil
### GetNumericalSeverity

`func (o *FindingCreate) GetNumericalSeverity() string`

GetNumericalSeverity returns the NumericalSeverity field if non-nil, zero value otherwise.

### GetNumericalSeverityOk

`func (o *FindingCreate) GetNumericalSeverityOk() (*string, bool)`

GetNumericalSeverityOk returns a tuple with the NumericalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericalSeverity

`func (o *FindingCreate) SetNumericalSeverity(v string)`

SetNumericalSeverity sets NumericalSeverity field to given value.


### GetLastReviewed

`func (o *FindingCreate) GetLastReviewed() time.Time`

GetLastReviewed returns the LastReviewed field if non-nil, zero value otherwise.

### GetLastReviewedOk

`func (o *FindingCreate) GetLastReviewedOk() (*time.Time, bool)`

GetLastReviewedOk returns a tuple with the LastReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewed

`func (o *FindingCreate) SetLastReviewed(v time.Time)`

SetLastReviewed sets LastReviewed field to given value.


### SetLastReviewedNil

`func (o *FindingCreate) SetLastReviewedNil(b bool)`

 SetLastReviewedNil sets the value for LastReviewed to be an explicit nil

### UnsetLastReviewed
`func (o *FindingCreate) UnsetLastReviewed()`

UnsetLastReviewed ensures that no value is present for LastReviewed, not even an explicit nil
### GetParam

`func (o *FindingCreate) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *FindingCreate) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *FindingCreate) SetParam(v string)`

SetParam sets Param field to given value.


### SetParamNil

`func (o *FindingCreate) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *FindingCreate) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetPayload

`func (o *FindingCreate) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FindingCreate) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FindingCreate) SetPayload(v string)`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *FindingCreate) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *FindingCreate) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetHashCode

`func (o *FindingCreate) GetHashCode() string`

GetHashCode returns the HashCode field if non-nil, zero value otherwise.

### GetHashCodeOk

`func (o *FindingCreate) GetHashCodeOk() (*string, bool)`

GetHashCodeOk returns a tuple with the HashCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashCode

`func (o *FindingCreate) SetHashCode(v string)`

SetHashCode sets HashCode field to given value.


### SetHashCodeNil

`func (o *FindingCreate) SetHashCodeNil(b bool)`

 SetHashCodeNil sets the value for HashCode to be an explicit nil

### UnsetHashCode
`func (o *FindingCreate) UnsetHashCode()`

UnsetHashCode ensures that no value is present for HashCode, not even an explicit nil
### GetLine

`func (o *FindingCreate) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FindingCreate) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FindingCreate) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *FindingCreate) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *FindingCreate) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *FindingCreate) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetFilePath

`func (o *FindingCreate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FindingCreate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FindingCreate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FindingCreate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *FindingCreate) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *FindingCreate) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetComponentName

`func (o *FindingCreate) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *FindingCreate) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *FindingCreate) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *FindingCreate) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### SetComponentNameNil

`func (o *FindingCreate) SetComponentNameNil(b bool)`

 SetComponentNameNil sets the value for ComponentName to be an explicit nil

### UnsetComponentName
`func (o *FindingCreate) UnsetComponentName()`

UnsetComponentName ensures that no value is present for ComponentName, not even an explicit nil
### GetComponentVersion

`func (o *FindingCreate) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *FindingCreate) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *FindingCreate) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *FindingCreate) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *FindingCreate) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *FindingCreate) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetStaticFinding

`func (o *FindingCreate) GetStaticFinding() bool`

GetStaticFinding returns the StaticFinding field if non-nil, zero value otherwise.

### GetStaticFindingOk

`func (o *FindingCreate) GetStaticFindingOk() (*bool, bool)`

GetStaticFindingOk returns a tuple with the StaticFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFinding

`func (o *FindingCreate) SetStaticFinding(v bool)`

SetStaticFinding sets StaticFinding field to given value.

### HasStaticFinding

`func (o *FindingCreate) HasStaticFinding() bool`

HasStaticFinding returns a boolean if a field has been set.

### GetDynamicFinding

`func (o *FindingCreate) GetDynamicFinding() bool`

GetDynamicFinding returns the DynamicFinding field if non-nil, zero value otherwise.

### GetDynamicFindingOk

`func (o *FindingCreate) GetDynamicFindingOk() (*bool, bool)`

GetDynamicFindingOk returns a tuple with the DynamicFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicFinding

`func (o *FindingCreate) SetDynamicFinding(v bool)`

SetDynamicFinding sets DynamicFinding field to given value.

### HasDynamicFinding

`func (o *FindingCreate) HasDynamicFinding() bool`

HasDynamicFinding returns a boolean if a field has been set.

### GetCreated

`func (o *FindingCreate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FindingCreate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FindingCreate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *FindingCreate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FindingCreate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetScannerConfidence

`func (o *FindingCreate) GetScannerConfidence() int32`

GetScannerConfidence returns the ScannerConfidence field if non-nil, zero value otherwise.

### GetScannerConfidenceOk

`func (o *FindingCreate) GetScannerConfidenceOk() (*int32, bool)`

GetScannerConfidenceOk returns a tuple with the ScannerConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannerConfidence

`func (o *FindingCreate) SetScannerConfidence(v int32)`

SetScannerConfidence sets ScannerConfidence field to given value.


### SetScannerConfidenceNil

`func (o *FindingCreate) SetScannerConfidenceNil(b bool)`

 SetScannerConfidenceNil sets the value for ScannerConfidence to be an explicit nil

### UnsetScannerConfidence
`func (o *FindingCreate) UnsetScannerConfidence()`

UnsetScannerConfidence ensures that no value is present for ScannerConfidence, not even an explicit nil
### GetUniqueIdFromTool

`func (o *FindingCreate) GetUniqueIdFromTool() string`

GetUniqueIdFromTool returns the UniqueIdFromTool field if non-nil, zero value otherwise.

### GetUniqueIdFromToolOk

`func (o *FindingCreate) GetUniqueIdFromToolOk() (*string, bool)`

GetUniqueIdFromToolOk returns a tuple with the UniqueIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdFromTool

`func (o *FindingCreate) SetUniqueIdFromTool(v string)`

SetUniqueIdFromTool sets UniqueIdFromTool field to given value.

### HasUniqueIdFromTool

`func (o *FindingCreate) HasUniqueIdFromTool() bool`

HasUniqueIdFromTool returns a boolean if a field has been set.

### SetUniqueIdFromToolNil

`func (o *FindingCreate) SetUniqueIdFromToolNil(b bool)`

 SetUniqueIdFromToolNil sets the value for UniqueIdFromTool to be an explicit nil

### UnsetUniqueIdFromTool
`func (o *FindingCreate) UnsetUniqueIdFromTool()`

UnsetUniqueIdFromTool ensures that no value is present for UniqueIdFromTool, not even an explicit nil
### GetVulnIdFromTool

`func (o *FindingCreate) GetVulnIdFromTool() string`

GetVulnIdFromTool returns the VulnIdFromTool field if non-nil, zero value otherwise.

### GetVulnIdFromToolOk

`func (o *FindingCreate) GetVulnIdFromToolOk() (*string, bool)`

GetVulnIdFromToolOk returns a tuple with the VulnIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnIdFromTool

`func (o *FindingCreate) SetVulnIdFromTool(v string)`

SetVulnIdFromTool sets VulnIdFromTool field to given value.

### HasVulnIdFromTool

`func (o *FindingCreate) HasVulnIdFromTool() bool`

HasVulnIdFromTool returns a boolean if a field has been set.

### SetVulnIdFromToolNil

`func (o *FindingCreate) SetVulnIdFromToolNil(b bool)`

 SetVulnIdFromToolNil sets the value for VulnIdFromTool to be an explicit nil

### UnsetVulnIdFromTool
`func (o *FindingCreate) UnsetVulnIdFromTool()`

UnsetVulnIdFromTool ensures that no value is present for VulnIdFromTool, not even an explicit nil
### GetSastSourceObject

`func (o *FindingCreate) GetSastSourceObject() string`

GetSastSourceObject returns the SastSourceObject field if non-nil, zero value otherwise.

### GetSastSourceObjectOk

`func (o *FindingCreate) GetSastSourceObjectOk() (*string, bool)`

GetSastSourceObjectOk returns a tuple with the SastSourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceObject

`func (o *FindingCreate) SetSastSourceObject(v string)`

SetSastSourceObject sets SastSourceObject field to given value.

### HasSastSourceObject

`func (o *FindingCreate) HasSastSourceObject() bool`

HasSastSourceObject returns a boolean if a field has been set.

### SetSastSourceObjectNil

`func (o *FindingCreate) SetSastSourceObjectNil(b bool)`

 SetSastSourceObjectNil sets the value for SastSourceObject to be an explicit nil

### UnsetSastSourceObject
`func (o *FindingCreate) UnsetSastSourceObject()`

UnsetSastSourceObject ensures that no value is present for SastSourceObject, not even an explicit nil
### GetSastSinkObject

`func (o *FindingCreate) GetSastSinkObject() string`

GetSastSinkObject returns the SastSinkObject field if non-nil, zero value otherwise.

### GetSastSinkObjectOk

`func (o *FindingCreate) GetSastSinkObjectOk() (*string, bool)`

GetSastSinkObjectOk returns a tuple with the SastSinkObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSinkObject

`func (o *FindingCreate) SetSastSinkObject(v string)`

SetSastSinkObject sets SastSinkObject field to given value.

### HasSastSinkObject

`func (o *FindingCreate) HasSastSinkObject() bool`

HasSastSinkObject returns a boolean if a field has been set.

### SetSastSinkObjectNil

`func (o *FindingCreate) SetSastSinkObjectNil(b bool)`

 SetSastSinkObjectNil sets the value for SastSinkObject to be an explicit nil

### UnsetSastSinkObject
`func (o *FindingCreate) UnsetSastSinkObject()`

UnsetSastSinkObject ensures that no value is present for SastSinkObject, not even an explicit nil
### GetSastSourceLine

`func (o *FindingCreate) GetSastSourceLine() int32`

GetSastSourceLine returns the SastSourceLine field if non-nil, zero value otherwise.

### GetSastSourceLineOk

`func (o *FindingCreate) GetSastSourceLineOk() (*int32, bool)`

GetSastSourceLineOk returns a tuple with the SastSourceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceLine

`func (o *FindingCreate) SetSastSourceLine(v int32)`

SetSastSourceLine sets SastSourceLine field to given value.

### HasSastSourceLine

`func (o *FindingCreate) HasSastSourceLine() bool`

HasSastSourceLine returns a boolean if a field has been set.

### SetSastSourceLineNil

`func (o *FindingCreate) SetSastSourceLineNil(b bool)`

 SetSastSourceLineNil sets the value for SastSourceLine to be an explicit nil

### UnsetSastSourceLine
`func (o *FindingCreate) UnsetSastSourceLine()`

UnsetSastSourceLine ensures that no value is present for SastSourceLine, not even an explicit nil
### GetSastSourceFilePath

`func (o *FindingCreate) GetSastSourceFilePath() string`

GetSastSourceFilePath returns the SastSourceFilePath field if non-nil, zero value otherwise.

### GetSastSourceFilePathOk

`func (o *FindingCreate) GetSastSourceFilePathOk() (*string, bool)`

GetSastSourceFilePathOk returns a tuple with the SastSourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceFilePath

`func (o *FindingCreate) SetSastSourceFilePath(v string)`

SetSastSourceFilePath sets SastSourceFilePath field to given value.

### HasSastSourceFilePath

`func (o *FindingCreate) HasSastSourceFilePath() bool`

HasSastSourceFilePath returns a boolean if a field has been set.

### SetSastSourceFilePathNil

`func (o *FindingCreate) SetSastSourceFilePathNil(b bool)`

 SetSastSourceFilePathNil sets the value for SastSourceFilePath to be an explicit nil

### UnsetSastSourceFilePath
`func (o *FindingCreate) UnsetSastSourceFilePath()`

UnsetSastSourceFilePath ensures that no value is present for SastSourceFilePath, not even an explicit nil
### GetNbOccurences

`func (o *FindingCreate) GetNbOccurences() int32`

GetNbOccurences returns the NbOccurences field if non-nil, zero value otherwise.

### GetNbOccurencesOk

`func (o *FindingCreate) GetNbOccurencesOk() (*int32, bool)`

GetNbOccurencesOk returns a tuple with the NbOccurences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOccurences

`func (o *FindingCreate) SetNbOccurences(v int32)`

SetNbOccurences sets NbOccurences field to given value.

### HasNbOccurences

`func (o *FindingCreate) HasNbOccurences() bool`

HasNbOccurences returns a boolean if a field has been set.

### SetNbOccurencesNil

`func (o *FindingCreate) SetNbOccurencesNil(b bool)`

 SetNbOccurencesNil sets the value for NbOccurences to be an explicit nil

### UnsetNbOccurences
`func (o *FindingCreate) UnsetNbOccurences()`

UnsetNbOccurences ensures that no value is present for NbOccurences, not even an explicit nil
### GetPublishDate

`func (o *FindingCreate) GetPublishDate() string`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *FindingCreate) GetPublishDateOk() (*string, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *FindingCreate) SetPublishDate(v string)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *FindingCreate) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### SetPublishDateNil

`func (o *FindingCreate) SetPublishDateNil(b bool)`

 SetPublishDateNil sets the value for PublishDate to be an explicit nil

### UnsetPublishDate
`func (o *FindingCreate) UnsetPublishDate()`

UnsetPublishDate ensures that no value is present for PublishDate, not even an explicit nil
### GetService

`func (o *FindingCreate) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FindingCreate) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FindingCreate) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *FindingCreate) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *FindingCreate) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *FindingCreate) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetPlannedRemediationDate

`func (o *FindingCreate) GetPlannedRemediationDate() string`

GetPlannedRemediationDate returns the PlannedRemediationDate field if non-nil, zero value otherwise.

### GetPlannedRemediationDateOk

`func (o *FindingCreate) GetPlannedRemediationDateOk() (*string, bool)`

GetPlannedRemediationDateOk returns a tuple with the PlannedRemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationDate

`func (o *FindingCreate) SetPlannedRemediationDate(v string)`

SetPlannedRemediationDate sets PlannedRemediationDate field to given value.

### HasPlannedRemediationDate

`func (o *FindingCreate) HasPlannedRemediationDate() bool`

HasPlannedRemediationDate returns a boolean if a field has been set.

### SetPlannedRemediationDateNil

`func (o *FindingCreate) SetPlannedRemediationDateNil(b bool)`

 SetPlannedRemediationDateNil sets the value for PlannedRemediationDate to be an explicit nil

### UnsetPlannedRemediationDate
`func (o *FindingCreate) UnsetPlannedRemediationDate()`

UnsetPlannedRemediationDate ensures that no value is present for PlannedRemediationDate, not even an explicit nil
### GetPlannedRemediationVersion

`func (o *FindingCreate) GetPlannedRemediationVersion() string`

GetPlannedRemediationVersion returns the PlannedRemediationVersion field if non-nil, zero value otherwise.

### GetPlannedRemediationVersionOk

`func (o *FindingCreate) GetPlannedRemediationVersionOk() (*string, bool)`

GetPlannedRemediationVersionOk returns a tuple with the PlannedRemediationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationVersion

`func (o *FindingCreate) SetPlannedRemediationVersion(v string)`

SetPlannedRemediationVersion sets PlannedRemediationVersion field to given value.

### HasPlannedRemediationVersion

`func (o *FindingCreate) HasPlannedRemediationVersion() bool`

HasPlannedRemediationVersion returns a boolean if a field has been set.

### SetPlannedRemediationVersionNil

`func (o *FindingCreate) SetPlannedRemediationVersionNil(b bool)`

 SetPlannedRemediationVersionNil sets the value for PlannedRemediationVersion to be an explicit nil

### UnsetPlannedRemediationVersion
`func (o *FindingCreate) UnsetPlannedRemediationVersion()`

UnsetPlannedRemediationVersion ensures that no value is present for PlannedRemediationVersion, not even an explicit nil
### GetEffortForFixing

`func (o *FindingCreate) GetEffortForFixing() string`

GetEffortForFixing returns the EffortForFixing field if non-nil, zero value otherwise.

### GetEffortForFixingOk

`func (o *FindingCreate) GetEffortForFixingOk() (*string, bool)`

GetEffortForFixingOk returns a tuple with the EffortForFixing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortForFixing

`func (o *FindingCreate) SetEffortForFixing(v string)`

SetEffortForFixing sets EffortForFixing field to given value.

### HasEffortForFixing

`func (o *FindingCreate) HasEffortForFixing() bool`

HasEffortForFixing returns a boolean if a field has been set.

### SetEffortForFixingNil

`func (o *FindingCreate) SetEffortForFixingNil(b bool)`

 SetEffortForFixingNil sets the value for EffortForFixing to be an explicit nil

### UnsetEffortForFixing
`func (o *FindingCreate) UnsetEffortForFixing()`

UnsetEffortForFixing ensures that no value is present for EffortForFixing, not even an explicit nil
### GetDuplicateFinding

`func (o *FindingCreate) GetDuplicateFinding() int32`

GetDuplicateFinding returns the DuplicateFinding field if non-nil, zero value otherwise.

### GetDuplicateFindingOk

`func (o *FindingCreate) GetDuplicateFindingOk() (*int32, bool)`

GetDuplicateFindingOk returns a tuple with the DuplicateFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateFinding

`func (o *FindingCreate) SetDuplicateFinding(v int32)`

SetDuplicateFinding sets DuplicateFinding field to given value.


### SetDuplicateFindingNil

`func (o *FindingCreate) SetDuplicateFindingNil(b bool)`

 SetDuplicateFindingNil sets the value for DuplicateFinding to be an explicit nil

### UnsetDuplicateFinding
`func (o *FindingCreate) UnsetDuplicateFinding()`

UnsetDuplicateFinding ensures that no value is present for DuplicateFinding, not even an explicit nil
### GetReviewRequestedBy

`func (o *FindingCreate) GetReviewRequestedBy() int32`

GetReviewRequestedBy returns the ReviewRequestedBy field if non-nil, zero value otherwise.

### GetReviewRequestedByOk

`func (o *FindingCreate) GetReviewRequestedByOk() (*int32, bool)`

GetReviewRequestedByOk returns a tuple with the ReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequestedBy

`func (o *FindingCreate) SetReviewRequestedBy(v int32)`

SetReviewRequestedBy sets ReviewRequestedBy field to given value.

### HasReviewRequestedBy

`func (o *FindingCreate) HasReviewRequestedBy() bool`

HasReviewRequestedBy returns a boolean if a field has been set.

### SetReviewRequestedByNil

`func (o *FindingCreate) SetReviewRequestedByNil(b bool)`

 SetReviewRequestedByNil sets the value for ReviewRequestedBy to be an explicit nil

### UnsetReviewRequestedBy
`func (o *FindingCreate) UnsetReviewRequestedBy()`

UnsetReviewRequestedBy ensures that no value is present for ReviewRequestedBy, not even an explicit nil
### GetDefectReviewRequestedBy

`func (o *FindingCreate) GetDefectReviewRequestedBy() int32`

GetDefectReviewRequestedBy returns the DefectReviewRequestedBy field if non-nil, zero value otherwise.

### GetDefectReviewRequestedByOk

`func (o *FindingCreate) GetDefectReviewRequestedByOk() (*int32, bool)`

GetDefectReviewRequestedByOk returns a tuple with the DefectReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectReviewRequestedBy

`func (o *FindingCreate) SetDefectReviewRequestedBy(v int32)`

SetDefectReviewRequestedBy sets DefectReviewRequestedBy field to given value.

### HasDefectReviewRequestedBy

`func (o *FindingCreate) HasDefectReviewRequestedBy() bool`

HasDefectReviewRequestedBy returns a boolean if a field has been set.

### SetDefectReviewRequestedByNil

`func (o *FindingCreate) SetDefectReviewRequestedByNil(b bool)`

 SetDefectReviewRequestedByNil sets the value for DefectReviewRequestedBy to be an explicit nil

### UnsetDefectReviewRequestedBy
`func (o *FindingCreate) UnsetDefectReviewRequestedBy()`

UnsetDefectReviewRequestedBy ensures that no value is present for DefectReviewRequestedBy, not even an explicit nil
### GetMitigatedBy

`func (o *FindingCreate) GetMitigatedBy() int32`

GetMitigatedBy returns the MitigatedBy field if non-nil, zero value otherwise.

### GetMitigatedByOk

`func (o *FindingCreate) GetMitigatedByOk() (*int32, bool)`

GetMitigatedByOk returns a tuple with the MitigatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedBy

`func (o *FindingCreate) SetMitigatedBy(v int32)`

SetMitigatedBy sets MitigatedBy field to given value.


### SetMitigatedByNil

`func (o *FindingCreate) SetMitigatedByNil(b bool)`

 SetMitigatedByNil sets the value for MitigatedBy to be an explicit nil

### UnsetMitigatedBy
`func (o *FindingCreate) UnsetMitigatedBy()`

UnsetMitigatedBy ensures that no value is present for MitigatedBy, not even an explicit nil
### GetLastReviewedBy

`func (o *FindingCreate) GetLastReviewedBy() int32`

GetLastReviewedBy returns the LastReviewedBy field if non-nil, zero value otherwise.

### GetLastReviewedByOk

`func (o *FindingCreate) GetLastReviewedByOk() (*int32, bool)`

GetLastReviewedByOk returns a tuple with the LastReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewedBy

`func (o *FindingCreate) SetLastReviewedBy(v int32)`

SetLastReviewedBy sets LastReviewedBy field to given value.


### SetLastReviewedByNil

`func (o *FindingCreate) SetLastReviewedByNil(b bool)`

 SetLastReviewedByNil sets the value for LastReviewedBy to be an explicit nil

### UnsetLastReviewedBy
`func (o *FindingCreate) UnsetLastReviewedBy()`

UnsetLastReviewedBy ensures that no value is present for LastReviewedBy, not even an explicit nil
### GetSonarqubeIssue

`func (o *FindingCreate) GetSonarqubeIssue() int32`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *FindingCreate) GetSonarqubeIssueOk() (*int32, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *FindingCreate) SetSonarqubeIssue(v int32)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.

### HasSonarqubeIssue

`func (o *FindingCreate) HasSonarqubeIssue() bool`

HasSonarqubeIssue returns a boolean if a field has been set.

### SetSonarqubeIssueNil

`func (o *FindingCreate) SetSonarqubeIssueNil(b bool)`

 SetSonarqubeIssueNil sets the value for SonarqubeIssue to be an explicit nil

### UnsetSonarqubeIssue
`func (o *FindingCreate) UnsetSonarqubeIssue()`

UnsetSonarqubeIssue ensures that no value is present for SonarqubeIssue, not even an explicit nil
### GetEndpoints

`func (o *FindingCreate) GetEndpoints() []int32`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *FindingCreate) GetEndpointsOk() (*[]int32, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *FindingCreate) SetEndpoints(v []int32)`

SetEndpoints sets Endpoints field to given value.


### GetReviewers

`func (o *FindingCreate) GetReviewers() []int32`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *FindingCreate) GetReviewersOk() (*[]int32, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *FindingCreate) SetReviewers(v []int32)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *FindingCreate) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetFiles

`func (o *FindingCreate) GetFiles() []int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FindingCreate) GetFilesOk() (*[]int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FindingCreate) SetFiles(v []int32)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


