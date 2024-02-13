# FindingCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | **int32** |  | 
**ThreadId** | Pointer to **int32** |  | [optional] [default to 0]
**FoundBy** | **[]int32** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**PushToJira** | Pointer to **bool** |  | [optional] [default to false]
**VulnerabilityIds** | Pointer to [**[]VulnerabilityIdRequest**](VulnerabilityIdRequest.md) |  | [optional] 
**Reporter** | Pointer to **int32** |  | [optional] 
**Title** | **string** | A short description of the flaw. | 
**Date** | Pointer to **string** | The date the flaw was discovered. | [optional] 
**SlaStartDate** | Pointer to **NullableString** | (readonly)The date used as start date for SLA calculation. Set by expiring risk acceptances. Empty by default, causing a fallback to &#39;date&#39;. | [optional] 
**SlaExpirationDate** | Pointer to **NullableString** | (readonly)The date SLA expires for this finding. Empty by default, causing a fallback to &#39;date&#39;. | [optional] 
**Cwe** | Pointer to **NullableInt32** | The CWE number associated with this flaw. | [optional] 
**Cvssv3** | Pointer to **NullableString** | Common Vulnerability Scoring System version 3 (CVSSv3) score associated with this flaw. | [optional] 
**Cvssv3Score** | Pointer to **NullableFloat64** | Numerical CVSSv3 score for the vulnerability. If the vector is given, the score is updated while saving the finding | [optional] 
**Severity** | **string** | The severity level of this flaw (Critical, High, Medium, Low, Informational). | 
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
**UnderDefectReview** | Pointer to **bool** | Denotes if this finding is under defect review. | [optional] 
**IsMitigated** | Pointer to **bool** | Denotes if this flaw has been fixed. | [optional] 
**NumericalSeverity** | **string** | The numerical representation of the severity (S0, S1, S2, S3, S4). | 
**Line** | Pointer to **NullableInt32** | Source line number of the attack vector. | [optional] 
**FilePath** | Pointer to **NullableString** | Identified file(s) containing the flaw. | [optional] 
**ComponentName** | Pointer to **NullableString** | Name of the affected component (library name, part of a system, ...). | [optional] 
**ComponentVersion** | Pointer to **NullableString** | Version of the affected component. | [optional] 
**StaticFinding** | Pointer to **bool** | Flaw has been detected from a Static Application Security Testing tool (SAST). | [optional] 
**DynamicFinding** | Pointer to **bool** | Flaw has been detected from a Dynamic Application Security Testing tool (DAST). | [optional] 
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
**ReviewRequestedBy** | Pointer to **NullableInt32** | Documents who requested a review for this finding. | [optional] 
**DefectReviewRequestedBy** | Pointer to **NullableInt32** | Documents who requested a defect review for this flaw. | [optional] 
**SonarqubeIssue** | Pointer to **NullableInt32** | The SonarQube issue associated with this finding. | [optional] 
**Reviewers** | Pointer to **[]int32** | Documents who reviewed the flaw. | [optional] 

## Methods

### NewFindingCreateRequest

`func NewFindingCreateRequest(test int32, foundBy []int32, title string, severity string, description string, active bool, verified bool, numericalSeverity string, ) *FindingCreateRequest`

NewFindingCreateRequest instantiates a new FindingCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingCreateRequestWithDefaults

`func NewFindingCreateRequestWithDefaults() *FindingCreateRequest`

NewFindingCreateRequestWithDefaults instantiates a new FindingCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *FindingCreateRequest) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *FindingCreateRequest) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *FindingCreateRequest) SetTest(v int32)`

SetTest sets Test field to given value.


### GetThreadId

`func (o *FindingCreateRequest) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *FindingCreateRequest) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *FindingCreateRequest) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *FindingCreateRequest) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetFoundBy

`func (o *FindingCreateRequest) GetFoundBy() []int32`

GetFoundBy returns the FoundBy field if non-nil, zero value otherwise.

### GetFoundByOk

`func (o *FindingCreateRequest) GetFoundByOk() (*[]int32, bool)`

GetFoundByOk returns a tuple with the FoundBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundBy

`func (o *FindingCreateRequest) SetFoundBy(v []int32)`

SetFoundBy sets FoundBy field to given value.


### GetUrl

`func (o *FindingCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FindingCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FindingCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FindingCreateRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FindingCreateRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FindingCreateRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTags

`func (o *FindingCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindingCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindingCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindingCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPushToJira

`func (o *FindingCreateRequest) GetPushToJira() bool`

GetPushToJira returns the PushToJira field if non-nil, zero value otherwise.

### GetPushToJiraOk

`func (o *FindingCreateRequest) GetPushToJiraOk() (*bool, bool)`

GetPushToJiraOk returns a tuple with the PushToJira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToJira

`func (o *FindingCreateRequest) SetPushToJira(v bool)`

SetPushToJira sets PushToJira field to given value.

### HasPushToJira

`func (o *FindingCreateRequest) HasPushToJira() bool`

HasPushToJira returns a boolean if a field has been set.

### GetVulnerabilityIds

`func (o *FindingCreateRequest) GetVulnerabilityIds() []VulnerabilityIdRequest`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *FindingCreateRequest) GetVulnerabilityIdsOk() (*[]VulnerabilityIdRequest, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *FindingCreateRequest) SetVulnerabilityIds(v []VulnerabilityIdRequest)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *FindingCreateRequest) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetReporter

`func (o *FindingCreateRequest) GetReporter() int32`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *FindingCreateRequest) GetReporterOk() (*int32, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *FindingCreateRequest) SetReporter(v int32)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *FindingCreateRequest) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetTitle

`func (o *FindingCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDate

`func (o *FindingCreateRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FindingCreateRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FindingCreateRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FindingCreateRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSlaStartDate

`func (o *FindingCreateRequest) GetSlaStartDate() string`

GetSlaStartDate returns the SlaStartDate field if non-nil, zero value otherwise.

### GetSlaStartDateOk

`func (o *FindingCreateRequest) GetSlaStartDateOk() (*string, bool)`

GetSlaStartDateOk returns a tuple with the SlaStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStartDate

`func (o *FindingCreateRequest) SetSlaStartDate(v string)`

SetSlaStartDate sets SlaStartDate field to given value.

### HasSlaStartDate

`func (o *FindingCreateRequest) HasSlaStartDate() bool`

HasSlaStartDate returns a boolean if a field has been set.

### SetSlaStartDateNil

`func (o *FindingCreateRequest) SetSlaStartDateNil(b bool)`

 SetSlaStartDateNil sets the value for SlaStartDate to be an explicit nil

### UnsetSlaStartDate
`func (o *FindingCreateRequest) UnsetSlaStartDate()`

UnsetSlaStartDate ensures that no value is present for SlaStartDate, not even an explicit nil
### GetSlaExpirationDate

`func (o *FindingCreateRequest) GetSlaExpirationDate() string`

GetSlaExpirationDate returns the SlaExpirationDate field if non-nil, zero value otherwise.

### GetSlaExpirationDateOk

`func (o *FindingCreateRequest) GetSlaExpirationDateOk() (*string, bool)`

GetSlaExpirationDateOk returns a tuple with the SlaExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaExpirationDate

`func (o *FindingCreateRequest) SetSlaExpirationDate(v string)`

SetSlaExpirationDate sets SlaExpirationDate field to given value.

### HasSlaExpirationDate

`func (o *FindingCreateRequest) HasSlaExpirationDate() bool`

HasSlaExpirationDate returns a boolean if a field has been set.

### SetSlaExpirationDateNil

`func (o *FindingCreateRequest) SetSlaExpirationDateNil(b bool)`

 SetSlaExpirationDateNil sets the value for SlaExpirationDate to be an explicit nil

### UnsetSlaExpirationDate
`func (o *FindingCreateRequest) UnsetSlaExpirationDate()`

UnsetSlaExpirationDate ensures that no value is present for SlaExpirationDate, not even an explicit nil
### GetCwe

`func (o *FindingCreateRequest) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *FindingCreateRequest) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *FindingCreateRequest) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *FindingCreateRequest) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *FindingCreateRequest) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *FindingCreateRequest) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetCvssv3

`func (o *FindingCreateRequest) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *FindingCreateRequest) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *FindingCreateRequest) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *FindingCreateRequest) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *FindingCreateRequest) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *FindingCreateRequest) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetCvssv3Score

`func (o *FindingCreateRequest) GetCvssv3Score() float64`

GetCvssv3Score returns the Cvssv3Score field if non-nil, zero value otherwise.

### GetCvssv3ScoreOk

`func (o *FindingCreateRequest) GetCvssv3ScoreOk() (*float64, bool)`

GetCvssv3ScoreOk returns a tuple with the Cvssv3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3Score

`func (o *FindingCreateRequest) SetCvssv3Score(v float64)`

SetCvssv3Score sets Cvssv3Score field to given value.

### HasCvssv3Score

`func (o *FindingCreateRequest) HasCvssv3Score() bool`

HasCvssv3Score returns a boolean if a field has been set.

### SetCvssv3ScoreNil

`func (o *FindingCreateRequest) SetCvssv3ScoreNil(b bool)`

 SetCvssv3ScoreNil sets the value for Cvssv3Score to be an explicit nil

### UnsetCvssv3Score
`func (o *FindingCreateRequest) UnsetCvssv3Score()`

UnsetCvssv3Score ensures that no value is present for Cvssv3Score, not even an explicit nil
### GetSeverity

`func (o *FindingCreateRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingCreateRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingCreateRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetDescription

`func (o *FindingCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMitigation

`func (o *FindingCreateRequest) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *FindingCreateRequest) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *FindingCreateRequest) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *FindingCreateRequest) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *FindingCreateRequest) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *FindingCreateRequest) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *FindingCreateRequest) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *FindingCreateRequest) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *FindingCreateRequest) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *FindingCreateRequest) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *FindingCreateRequest) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *FindingCreateRequest) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetStepsToReproduce

`func (o *FindingCreateRequest) GetStepsToReproduce() string`

GetStepsToReproduce returns the StepsToReproduce field if non-nil, zero value otherwise.

### GetStepsToReproduceOk

`func (o *FindingCreateRequest) GetStepsToReproduceOk() (*string, bool)`

GetStepsToReproduceOk returns a tuple with the StepsToReproduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsToReproduce

`func (o *FindingCreateRequest) SetStepsToReproduce(v string)`

SetStepsToReproduce sets StepsToReproduce field to given value.

### HasStepsToReproduce

`func (o *FindingCreateRequest) HasStepsToReproduce() bool`

HasStepsToReproduce returns a boolean if a field has been set.

### SetStepsToReproduceNil

`func (o *FindingCreateRequest) SetStepsToReproduceNil(b bool)`

 SetStepsToReproduceNil sets the value for StepsToReproduce to be an explicit nil

### UnsetStepsToReproduce
`func (o *FindingCreateRequest) UnsetStepsToReproduce()`

UnsetStepsToReproduce ensures that no value is present for StepsToReproduce, not even an explicit nil
### GetSeverityJustification

`func (o *FindingCreateRequest) GetSeverityJustification() string`

GetSeverityJustification returns the SeverityJustification field if non-nil, zero value otherwise.

### GetSeverityJustificationOk

`func (o *FindingCreateRequest) GetSeverityJustificationOk() (*string, bool)`

GetSeverityJustificationOk returns a tuple with the SeverityJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityJustification

`func (o *FindingCreateRequest) SetSeverityJustification(v string)`

SetSeverityJustification sets SeverityJustification field to given value.

### HasSeverityJustification

`func (o *FindingCreateRequest) HasSeverityJustification() bool`

HasSeverityJustification returns a boolean if a field has been set.

### SetSeverityJustificationNil

`func (o *FindingCreateRequest) SetSeverityJustificationNil(b bool)`

 SetSeverityJustificationNil sets the value for SeverityJustification to be an explicit nil

### UnsetSeverityJustification
`func (o *FindingCreateRequest) UnsetSeverityJustification()`

UnsetSeverityJustification ensures that no value is present for SeverityJustification, not even an explicit nil
### GetReferences

`func (o *FindingCreateRequest) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FindingCreateRequest) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FindingCreateRequest) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *FindingCreateRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *FindingCreateRequest) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *FindingCreateRequest) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetActive

`func (o *FindingCreateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FindingCreateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FindingCreateRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetVerified

`func (o *FindingCreateRequest) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *FindingCreateRequest) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *FindingCreateRequest) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetFalseP

`func (o *FindingCreateRequest) GetFalseP() bool`

GetFalseP returns the FalseP field if non-nil, zero value otherwise.

### GetFalsePOk

`func (o *FindingCreateRequest) GetFalsePOk() (*bool, bool)`

GetFalsePOk returns a tuple with the FalseP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseP

`func (o *FindingCreateRequest) SetFalseP(v bool)`

SetFalseP sets FalseP field to given value.

### HasFalseP

`func (o *FindingCreateRequest) HasFalseP() bool`

HasFalseP returns a boolean if a field has been set.

### GetDuplicate

`func (o *FindingCreateRequest) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *FindingCreateRequest) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *FindingCreateRequest) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *FindingCreateRequest) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetOutOfScope

`func (o *FindingCreateRequest) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *FindingCreateRequest) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *FindingCreateRequest) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *FindingCreateRequest) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *FindingCreateRequest) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *FindingCreateRequest) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *FindingCreateRequest) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *FindingCreateRequest) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetUnderReview

`func (o *FindingCreateRequest) GetUnderReview() bool`

GetUnderReview returns the UnderReview field if non-nil, zero value otherwise.

### GetUnderReviewOk

`func (o *FindingCreateRequest) GetUnderReviewOk() (*bool, bool)`

GetUnderReviewOk returns a tuple with the UnderReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview

`func (o *FindingCreateRequest) SetUnderReview(v bool)`

SetUnderReview sets UnderReview field to given value.

### HasUnderReview

`func (o *FindingCreateRequest) HasUnderReview() bool`

HasUnderReview returns a boolean if a field has been set.

### GetUnderDefectReview

`func (o *FindingCreateRequest) GetUnderDefectReview() bool`

GetUnderDefectReview returns the UnderDefectReview field if non-nil, zero value otherwise.

### GetUnderDefectReviewOk

`func (o *FindingCreateRequest) GetUnderDefectReviewOk() (*bool, bool)`

GetUnderDefectReviewOk returns a tuple with the UnderDefectReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderDefectReview

`func (o *FindingCreateRequest) SetUnderDefectReview(v bool)`

SetUnderDefectReview sets UnderDefectReview field to given value.

### HasUnderDefectReview

`func (o *FindingCreateRequest) HasUnderDefectReview() bool`

HasUnderDefectReview returns a boolean if a field has been set.

### GetIsMitigated

`func (o *FindingCreateRequest) GetIsMitigated() bool`

GetIsMitigated returns the IsMitigated field if non-nil, zero value otherwise.

### GetIsMitigatedOk

`func (o *FindingCreateRequest) GetIsMitigatedOk() (*bool, bool)`

GetIsMitigatedOk returns a tuple with the IsMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMitigated

`func (o *FindingCreateRequest) SetIsMitigated(v bool)`

SetIsMitigated sets IsMitigated field to given value.

### HasIsMitigated

`func (o *FindingCreateRequest) HasIsMitigated() bool`

HasIsMitigated returns a boolean if a field has been set.

### GetNumericalSeverity

`func (o *FindingCreateRequest) GetNumericalSeverity() string`

GetNumericalSeverity returns the NumericalSeverity field if non-nil, zero value otherwise.

### GetNumericalSeverityOk

`func (o *FindingCreateRequest) GetNumericalSeverityOk() (*string, bool)`

GetNumericalSeverityOk returns a tuple with the NumericalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericalSeverity

`func (o *FindingCreateRequest) SetNumericalSeverity(v string)`

SetNumericalSeverity sets NumericalSeverity field to given value.


### GetLine

`func (o *FindingCreateRequest) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FindingCreateRequest) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FindingCreateRequest) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *FindingCreateRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *FindingCreateRequest) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *FindingCreateRequest) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetFilePath

`func (o *FindingCreateRequest) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FindingCreateRequest) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FindingCreateRequest) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FindingCreateRequest) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *FindingCreateRequest) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *FindingCreateRequest) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetComponentName

`func (o *FindingCreateRequest) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *FindingCreateRequest) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *FindingCreateRequest) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *FindingCreateRequest) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### SetComponentNameNil

`func (o *FindingCreateRequest) SetComponentNameNil(b bool)`

 SetComponentNameNil sets the value for ComponentName to be an explicit nil

### UnsetComponentName
`func (o *FindingCreateRequest) UnsetComponentName()`

UnsetComponentName ensures that no value is present for ComponentName, not even an explicit nil
### GetComponentVersion

`func (o *FindingCreateRequest) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *FindingCreateRequest) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *FindingCreateRequest) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *FindingCreateRequest) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *FindingCreateRequest) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *FindingCreateRequest) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetStaticFinding

`func (o *FindingCreateRequest) GetStaticFinding() bool`

GetStaticFinding returns the StaticFinding field if non-nil, zero value otherwise.

### GetStaticFindingOk

`func (o *FindingCreateRequest) GetStaticFindingOk() (*bool, bool)`

GetStaticFindingOk returns a tuple with the StaticFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFinding

`func (o *FindingCreateRequest) SetStaticFinding(v bool)`

SetStaticFinding sets StaticFinding field to given value.

### HasStaticFinding

`func (o *FindingCreateRequest) HasStaticFinding() bool`

HasStaticFinding returns a boolean if a field has been set.

### GetDynamicFinding

`func (o *FindingCreateRequest) GetDynamicFinding() bool`

GetDynamicFinding returns the DynamicFinding field if non-nil, zero value otherwise.

### GetDynamicFindingOk

`func (o *FindingCreateRequest) GetDynamicFindingOk() (*bool, bool)`

GetDynamicFindingOk returns a tuple with the DynamicFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicFinding

`func (o *FindingCreateRequest) SetDynamicFinding(v bool)`

SetDynamicFinding sets DynamicFinding field to given value.

### HasDynamicFinding

`func (o *FindingCreateRequest) HasDynamicFinding() bool`

HasDynamicFinding returns a boolean if a field has been set.

### GetUniqueIdFromTool

`func (o *FindingCreateRequest) GetUniqueIdFromTool() string`

GetUniqueIdFromTool returns the UniqueIdFromTool field if non-nil, zero value otherwise.

### GetUniqueIdFromToolOk

`func (o *FindingCreateRequest) GetUniqueIdFromToolOk() (*string, bool)`

GetUniqueIdFromToolOk returns a tuple with the UniqueIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdFromTool

`func (o *FindingCreateRequest) SetUniqueIdFromTool(v string)`

SetUniqueIdFromTool sets UniqueIdFromTool field to given value.

### HasUniqueIdFromTool

`func (o *FindingCreateRequest) HasUniqueIdFromTool() bool`

HasUniqueIdFromTool returns a boolean if a field has been set.

### SetUniqueIdFromToolNil

`func (o *FindingCreateRequest) SetUniqueIdFromToolNil(b bool)`

 SetUniqueIdFromToolNil sets the value for UniqueIdFromTool to be an explicit nil

### UnsetUniqueIdFromTool
`func (o *FindingCreateRequest) UnsetUniqueIdFromTool()`

UnsetUniqueIdFromTool ensures that no value is present for UniqueIdFromTool, not even an explicit nil
### GetVulnIdFromTool

`func (o *FindingCreateRequest) GetVulnIdFromTool() string`

GetVulnIdFromTool returns the VulnIdFromTool field if non-nil, zero value otherwise.

### GetVulnIdFromToolOk

`func (o *FindingCreateRequest) GetVulnIdFromToolOk() (*string, bool)`

GetVulnIdFromToolOk returns a tuple with the VulnIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnIdFromTool

`func (o *FindingCreateRequest) SetVulnIdFromTool(v string)`

SetVulnIdFromTool sets VulnIdFromTool field to given value.

### HasVulnIdFromTool

`func (o *FindingCreateRequest) HasVulnIdFromTool() bool`

HasVulnIdFromTool returns a boolean if a field has been set.

### SetVulnIdFromToolNil

`func (o *FindingCreateRequest) SetVulnIdFromToolNil(b bool)`

 SetVulnIdFromToolNil sets the value for VulnIdFromTool to be an explicit nil

### UnsetVulnIdFromTool
`func (o *FindingCreateRequest) UnsetVulnIdFromTool()`

UnsetVulnIdFromTool ensures that no value is present for VulnIdFromTool, not even an explicit nil
### GetSastSourceObject

`func (o *FindingCreateRequest) GetSastSourceObject() string`

GetSastSourceObject returns the SastSourceObject field if non-nil, zero value otherwise.

### GetSastSourceObjectOk

`func (o *FindingCreateRequest) GetSastSourceObjectOk() (*string, bool)`

GetSastSourceObjectOk returns a tuple with the SastSourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceObject

`func (o *FindingCreateRequest) SetSastSourceObject(v string)`

SetSastSourceObject sets SastSourceObject field to given value.

### HasSastSourceObject

`func (o *FindingCreateRequest) HasSastSourceObject() bool`

HasSastSourceObject returns a boolean if a field has been set.

### SetSastSourceObjectNil

`func (o *FindingCreateRequest) SetSastSourceObjectNil(b bool)`

 SetSastSourceObjectNil sets the value for SastSourceObject to be an explicit nil

### UnsetSastSourceObject
`func (o *FindingCreateRequest) UnsetSastSourceObject()`

UnsetSastSourceObject ensures that no value is present for SastSourceObject, not even an explicit nil
### GetSastSinkObject

`func (o *FindingCreateRequest) GetSastSinkObject() string`

GetSastSinkObject returns the SastSinkObject field if non-nil, zero value otherwise.

### GetSastSinkObjectOk

`func (o *FindingCreateRequest) GetSastSinkObjectOk() (*string, bool)`

GetSastSinkObjectOk returns a tuple with the SastSinkObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSinkObject

`func (o *FindingCreateRequest) SetSastSinkObject(v string)`

SetSastSinkObject sets SastSinkObject field to given value.

### HasSastSinkObject

`func (o *FindingCreateRequest) HasSastSinkObject() bool`

HasSastSinkObject returns a boolean if a field has been set.

### SetSastSinkObjectNil

`func (o *FindingCreateRequest) SetSastSinkObjectNil(b bool)`

 SetSastSinkObjectNil sets the value for SastSinkObject to be an explicit nil

### UnsetSastSinkObject
`func (o *FindingCreateRequest) UnsetSastSinkObject()`

UnsetSastSinkObject ensures that no value is present for SastSinkObject, not even an explicit nil
### GetSastSourceLine

`func (o *FindingCreateRequest) GetSastSourceLine() int32`

GetSastSourceLine returns the SastSourceLine field if non-nil, zero value otherwise.

### GetSastSourceLineOk

`func (o *FindingCreateRequest) GetSastSourceLineOk() (*int32, bool)`

GetSastSourceLineOk returns a tuple with the SastSourceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceLine

`func (o *FindingCreateRequest) SetSastSourceLine(v int32)`

SetSastSourceLine sets SastSourceLine field to given value.

### HasSastSourceLine

`func (o *FindingCreateRequest) HasSastSourceLine() bool`

HasSastSourceLine returns a boolean if a field has been set.

### SetSastSourceLineNil

`func (o *FindingCreateRequest) SetSastSourceLineNil(b bool)`

 SetSastSourceLineNil sets the value for SastSourceLine to be an explicit nil

### UnsetSastSourceLine
`func (o *FindingCreateRequest) UnsetSastSourceLine()`

UnsetSastSourceLine ensures that no value is present for SastSourceLine, not even an explicit nil
### GetSastSourceFilePath

`func (o *FindingCreateRequest) GetSastSourceFilePath() string`

GetSastSourceFilePath returns the SastSourceFilePath field if non-nil, zero value otherwise.

### GetSastSourceFilePathOk

`func (o *FindingCreateRequest) GetSastSourceFilePathOk() (*string, bool)`

GetSastSourceFilePathOk returns a tuple with the SastSourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceFilePath

`func (o *FindingCreateRequest) SetSastSourceFilePath(v string)`

SetSastSourceFilePath sets SastSourceFilePath field to given value.

### HasSastSourceFilePath

`func (o *FindingCreateRequest) HasSastSourceFilePath() bool`

HasSastSourceFilePath returns a boolean if a field has been set.

### SetSastSourceFilePathNil

`func (o *FindingCreateRequest) SetSastSourceFilePathNil(b bool)`

 SetSastSourceFilePathNil sets the value for SastSourceFilePath to be an explicit nil

### UnsetSastSourceFilePath
`func (o *FindingCreateRequest) UnsetSastSourceFilePath()`

UnsetSastSourceFilePath ensures that no value is present for SastSourceFilePath, not even an explicit nil
### GetNbOccurences

`func (o *FindingCreateRequest) GetNbOccurences() int32`

GetNbOccurences returns the NbOccurences field if non-nil, zero value otherwise.

### GetNbOccurencesOk

`func (o *FindingCreateRequest) GetNbOccurencesOk() (*int32, bool)`

GetNbOccurencesOk returns a tuple with the NbOccurences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOccurences

`func (o *FindingCreateRequest) SetNbOccurences(v int32)`

SetNbOccurences sets NbOccurences field to given value.

### HasNbOccurences

`func (o *FindingCreateRequest) HasNbOccurences() bool`

HasNbOccurences returns a boolean if a field has been set.

### SetNbOccurencesNil

`func (o *FindingCreateRequest) SetNbOccurencesNil(b bool)`

 SetNbOccurencesNil sets the value for NbOccurences to be an explicit nil

### UnsetNbOccurences
`func (o *FindingCreateRequest) UnsetNbOccurences()`

UnsetNbOccurences ensures that no value is present for NbOccurences, not even an explicit nil
### GetPublishDate

`func (o *FindingCreateRequest) GetPublishDate() string`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *FindingCreateRequest) GetPublishDateOk() (*string, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *FindingCreateRequest) SetPublishDate(v string)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *FindingCreateRequest) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### SetPublishDateNil

`func (o *FindingCreateRequest) SetPublishDateNil(b bool)`

 SetPublishDateNil sets the value for PublishDate to be an explicit nil

### UnsetPublishDate
`func (o *FindingCreateRequest) UnsetPublishDate()`

UnsetPublishDate ensures that no value is present for PublishDate, not even an explicit nil
### GetService

`func (o *FindingCreateRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FindingCreateRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FindingCreateRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *FindingCreateRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *FindingCreateRequest) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *FindingCreateRequest) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetPlannedRemediationDate

`func (o *FindingCreateRequest) GetPlannedRemediationDate() string`

GetPlannedRemediationDate returns the PlannedRemediationDate field if non-nil, zero value otherwise.

### GetPlannedRemediationDateOk

`func (o *FindingCreateRequest) GetPlannedRemediationDateOk() (*string, bool)`

GetPlannedRemediationDateOk returns a tuple with the PlannedRemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationDate

`func (o *FindingCreateRequest) SetPlannedRemediationDate(v string)`

SetPlannedRemediationDate sets PlannedRemediationDate field to given value.

### HasPlannedRemediationDate

`func (o *FindingCreateRequest) HasPlannedRemediationDate() bool`

HasPlannedRemediationDate returns a boolean if a field has been set.

### SetPlannedRemediationDateNil

`func (o *FindingCreateRequest) SetPlannedRemediationDateNil(b bool)`

 SetPlannedRemediationDateNil sets the value for PlannedRemediationDate to be an explicit nil

### UnsetPlannedRemediationDate
`func (o *FindingCreateRequest) UnsetPlannedRemediationDate()`

UnsetPlannedRemediationDate ensures that no value is present for PlannedRemediationDate, not even an explicit nil
### GetPlannedRemediationVersion

`func (o *FindingCreateRequest) GetPlannedRemediationVersion() string`

GetPlannedRemediationVersion returns the PlannedRemediationVersion field if non-nil, zero value otherwise.

### GetPlannedRemediationVersionOk

`func (o *FindingCreateRequest) GetPlannedRemediationVersionOk() (*string, bool)`

GetPlannedRemediationVersionOk returns a tuple with the PlannedRemediationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationVersion

`func (o *FindingCreateRequest) SetPlannedRemediationVersion(v string)`

SetPlannedRemediationVersion sets PlannedRemediationVersion field to given value.

### HasPlannedRemediationVersion

`func (o *FindingCreateRequest) HasPlannedRemediationVersion() bool`

HasPlannedRemediationVersion returns a boolean if a field has been set.

### SetPlannedRemediationVersionNil

`func (o *FindingCreateRequest) SetPlannedRemediationVersionNil(b bool)`

 SetPlannedRemediationVersionNil sets the value for PlannedRemediationVersion to be an explicit nil

### UnsetPlannedRemediationVersion
`func (o *FindingCreateRequest) UnsetPlannedRemediationVersion()`

UnsetPlannedRemediationVersion ensures that no value is present for PlannedRemediationVersion, not even an explicit nil
### GetEffortForFixing

`func (o *FindingCreateRequest) GetEffortForFixing() string`

GetEffortForFixing returns the EffortForFixing field if non-nil, zero value otherwise.

### GetEffortForFixingOk

`func (o *FindingCreateRequest) GetEffortForFixingOk() (*string, bool)`

GetEffortForFixingOk returns a tuple with the EffortForFixing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortForFixing

`func (o *FindingCreateRequest) SetEffortForFixing(v string)`

SetEffortForFixing sets EffortForFixing field to given value.

### HasEffortForFixing

`func (o *FindingCreateRequest) HasEffortForFixing() bool`

HasEffortForFixing returns a boolean if a field has been set.

### SetEffortForFixingNil

`func (o *FindingCreateRequest) SetEffortForFixingNil(b bool)`

 SetEffortForFixingNil sets the value for EffortForFixing to be an explicit nil

### UnsetEffortForFixing
`func (o *FindingCreateRequest) UnsetEffortForFixing()`

UnsetEffortForFixing ensures that no value is present for EffortForFixing, not even an explicit nil
### GetReviewRequestedBy

`func (o *FindingCreateRequest) GetReviewRequestedBy() int32`

GetReviewRequestedBy returns the ReviewRequestedBy field if non-nil, zero value otherwise.

### GetReviewRequestedByOk

`func (o *FindingCreateRequest) GetReviewRequestedByOk() (*int32, bool)`

GetReviewRequestedByOk returns a tuple with the ReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequestedBy

`func (o *FindingCreateRequest) SetReviewRequestedBy(v int32)`

SetReviewRequestedBy sets ReviewRequestedBy field to given value.

### HasReviewRequestedBy

`func (o *FindingCreateRequest) HasReviewRequestedBy() bool`

HasReviewRequestedBy returns a boolean if a field has been set.

### SetReviewRequestedByNil

`func (o *FindingCreateRequest) SetReviewRequestedByNil(b bool)`

 SetReviewRequestedByNil sets the value for ReviewRequestedBy to be an explicit nil

### UnsetReviewRequestedBy
`func (o *FindingCreateRequest) UnsetReviewRequestedBy()`

UnsetReviewRequestedBy ensures that no value is present for ReviewRequestedBy, not even an explicit nil
### GetDefectReviewRequestedBy

`func (o *FindingCreateRequest) GetDefectReviewRequestedBy() int32`

GetDefectReviewRequestedBy returns the DefectReviewRequestedBy field if non-nil, zero value otherwise.

### GetDefectReviewRequestedByOk

`func (o *FindingCreateRequest) GetDefectReviewRequestedByOk() (*int32, bool)`

GetDefectReviewRequestedByOk returns a tuple with the DefectReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectReviewRequestedBy

`func (o *FindingCreateRequest) SetDefectReviewRequestedBy(v int32)`

SetDefectReviewRequestedBy sets DefectReviewRequestedBy field to given value.

### HasDefectReviewRequestedBy

`func (o *FindingCreateRequest) HasDefectReviewRequestedBy() bool`

HasDefectReviewRequestedBy returns a boolean if a field has been set.

### SetDefectReviewRequestedByNil

`func (o *FindingCreateRequest) SetDefectReviewRequestedByNil(b bool)`

 SetDefectReviewRequestedByNil sets the value for DefectReviewRequestedBy to be an explicit nil

### UnsetDefectReviewRequestedBy
`func (o *FindingCreateRequest) UnsetDefectReviewRequestedBy()`

UnsetDefectReviewRequestedBy ensures that no value is present for DefectReviewRequestedBy, not even an explicit nil
### GetSonarqubeIssue

`func (o *FindingCreateRequest) GetSonarqubeIssue() int32`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *FindingCreateRequest) GetSonarqubeIssueOk() (*int32, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *FindingCreateRequest) SetSonarqubeIssue(v int32)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.

### HasSonarqubeIssue

`func (o *FindingCreateRequest) HasSonarqubeIssue() bool`

HasSonarqubeIssue returns a boolean if a field has been set.

### SetSonarqubeIssueNil

`func (o *FindingCreateRequest) SetSonarqubeIssueNil(b bool)`

 SetSonarqubeIssueNil sets the value for SonarqubeIssue to be an explicit nil

### UnsetSonarqubeIssue
`func (o *FindingCreateRequest) UnsetSonarqubeIssue()`

UnsetSonarqubeIssue ensures that no value is present for SonarqubeIssue, not even an explicit nil
### GetReviewers

`func (o *FindingCreateRequest) GetReviewers() []int32`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *FindingCreateRequest) GetReviewersOk() (*[]int32, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *FindingCreateRequest) SetReviewers(v []int32)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *FindingCreateRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


