# FindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**PushToJira** | Pointer to **bool** |  | [optional] [default to false]
**VulnerabilityIds** | Pointer to [**[]VulnerabilityIdRequest**](VulnerabilityIdRequest.md) |  | [optional] 
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
**Active** | Pointer to **bool** | Denotes if this flaw is active or not. | [optional] 
**Verified** | Pointer to **bool** | Denotes if this flaw has been manually verified by the tester. | [optional] 
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

### NewFindingRequest

`func NewFindingRequest(title string, severity string, description string, numericalSeverity string, ) *FindingRequest`

NewFindingRequest instantiates a new FindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingRequestWithDefaults

`func NewFindingRequestWithDefaults() *FindingRequest`

NewFindingRequestWithDefaults instantiates a new FindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *FindingRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindingRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindingRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindingRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPushToJira

`func (o *FindingRequest) GetPushToJira() bool`

GetPushToJira returns the PushToJira field if non-nil, zero value otherwise.

### GetPushToJiraOk

`func (o *FindingRequest) GetPushToJiraOk() (*bool, bool)`

GetPushToJiraOk returns a tuple with the PushToJira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToJira

`func (o *FindingRequest) SetPushToJira(v bool)`

SetPushToJira sets PushToJira field to given value.

### HasPushToJira

`func (o *FindingRequest) HasPushToJira() bool`

HasPushToJira returns a boolean if a field has been set.

### GetVulnerabilityIds

`func (o *FindingRequest) GetVulnerabilityIds() []VulnerabilityIdRequest`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *FindingRequest) GetVulnerabilityIdsOk() (*[]VulnerabilityIdRequest, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *FindingRequest) SetVulnerabilityIds(v []VulnerabilityIdRequest)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *FindingRequest) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetReporter

`func (o *FindingRequest) GetReporter() int32`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *FindingRequest) GetReporterOk() (*int32, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *FindingRequest) SetReporter(v int32)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *FindingRequest) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetTitle

`func (o *FindingRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDate

`func (o *FindingRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FindingRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FindingRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FindingRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSlaStartDate

`func (o *FindingRequest) GetSlaStartDate() string`

GetSlaStartDate returns the SlaStartDate field if non-nil, zero value otherwise.

### GetSlaStartDateOk

`func (o *FindingRequest) GetSlaStartDateOk() (*string, bool)`

GetSlaStartDateOk returns a tuple with the SlaStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStartDate

`func (o *FindingRequest) SetSlaStartDate(v string)`

SetSlaStartDate sets SlaStartDate field to given value.

### HasSlaStartDate

`func (o *FindingRequest) HasSlaStartDate() bool`

HasSlaStartDate returns a boolean if a field has been set.

### SetSlaStartDateNil

`func (o *FindingRequest) SetSlaStartDateNil(b bool)`

 SetSlaStartDateNil sets the value for SlaStartDate to be an explicit nil

### UnsetSlaStartDate
`func (o *FindingRequest) UnsetSlaStartDate()`

UnsetSlaStartDate ensures that no value is present for SlaStartDate, not even an explicit nil
### GetSlaExpirationDate

`func (o *FindingRequest) GetSlaExpirationDate() string`

GetSlaExpirationDate returns the SlaExpirationDate field if non-nil, zero value otherwise.

### GetSlaExpirationDateOk

`func (o *FindingRequest) GetSlaExpirationDateOk() (*string, bool)`

GetSlaExpirationDateOk returns a tuple with the SlaExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaExpirationDate

`func (o *FindingRequest) SetSlaExpirationDate(v string)`

SetSlaExpirationDate sets SlaExpirationDate field to given value.

### HasSlaExpirationDate

`func (o *FindingRequest) HasSlaExpirationDate() bool`

HasSlaExpirationDate returns a boolean if a field has been set.

### SetSlaExpirationDateNil

`func (o *FindingRequest) SetSlaExpirationDateNil(b bool)`

 SetSlaExpirationDateNil sets the value for SlaExpirationDate to be an explicit nil

### UnsetSlaExpirationDate
`func (o *FindingRequest) UnsetSlaExpirationDate()`

UnsetSlaExpirationDate ensures that no value is present for SlaExpirationDate, not even an explicit nil
### GetCwe

`func (o *FindingRequest) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *FindingRequest) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *FindingRequest) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *FindingRequest) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *FindingRequest) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *FindingRequest) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetEpssScore

`func (o *FindingRequest) GetEpssScore() float64`

GetEpssScore returns the EpssScore field if non-nil, zero value otherwise.

### GetEpssScoreOk

`func (o *FindingRequest) GetEpssScoreOk() (*float64, bool)`

GetEpssScoreOk returns a tuple with the EpssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpssScore

`func (o *FindingRequest) SetEpssScore(v float64)`

SetEpssScore sets EpssScore field to given value.

### HasEpssScore

`func (o *FindingRequest) HasEpssScore() bool`

HasEpssScore returns a boolean if a field has been set.

### SetEpssScoreNil

`func (o *FindingRequest) SetEpssScoreNil(b bool)`

 SetEpssScoreNil sets the value for EpssScore to be an explicit nil

### UnsetEpssScore
`func (o *FindingRequest) UnsetEpssScore()`

UnsetEpssScore ensures that no value is present for EpssScore, not even an explicit nil
### GetEpssPercentile

`func (o *FindingRequest) GetEpssPercentile() float64`

GetEpssPercentile returns the EpssPercentile field if non-nil, zero value otherwise.

### GetEpssPercentileOk

`func (o *FindingRequest) GetEpssPercentileOk() (*float64, bool)`

GetEpssPercentileOk returns a tuple with the EpssPercentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpssPercentile

`func (o *FindingRequest) SetEpssPercentile(v float64)`

SetEpssPercentile sets EpssPercentile field to given value.

### HasEpssPercentile

`func (o *FindingRequest) HasEpssPercentile() bool`

HasEpssPercentile returns a boolean if a field has been set.

### SetEpssPercentileNil

`func (o *FindingRequest) SetEpssPercentileNil(b bool)`

 SetEpssPercentileNil sets the value for EpssPercentile to be an explicit nil

### UnsetEpssPercentile
`func (o *FindingRequest) UnsetEpssPercentile()`

UnsetEpssPercentile ensures that no value is present for EpssPercentile, not even an explicit nil
### GetCvssv3

`func (o *FindingRequest) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *FindingRequest) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *FindingRequest) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *FindingRequest) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *FindingRequest) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *FindingRequest) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetCvssv3Score

`func (o *FindingRequest) GetCvssv3Score() float64`

GetCvssv3Score returns the Cvssv3Score field if non-nil, zero value otherwise.

### GetCvssv3ScoreOk

`func (o *FindingRequest) GetCvssv3ScoreOk() (*float64, bool)`

GetCvssv3ScoreOk returns a tuple with the Cvssv3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3Score

`func (o *FindingRequest) SetCvssv3Score(v float64)`

SetCvssv3Score sets Cvssv3Score field to given value.

### HasCvssv3Score

`func (o *FindingRequest) HasCvssv3Score() bool`

HasCvssv3Score returns a boolean if a field has been set.

### SetCvssv3ScoreNil

`func (o *FindingRequest) SetCvssv3ScoreNil(b bool)`

 SetCvssv3ScoreNil sets the value for Cvssv3Score to be an explicit nil

### UnsetCvssv3Score
`func (o *FindingRequest) UnsetCvssv3Score()`

UnsetCvssv3Score ensures that no value is present for Cvssv3Score, not even an explicit nil
### GetSeverity

`func (o *FindingRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetDescription

`func (o *FindingRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMitigation

`func (o *FindingRequest) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *FindingRequest) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *FindingRequest) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *FindingRequest) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *FindingRequest) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *FindingRequest) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *FindingRequest) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *FindingRequest) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *FindingRequest) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *FindingRequest) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *FindingRequest) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *FindingRequest) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetStepsToReproduce

`func (o *FindingRequest) GetStepsToReproduce() string`

GetStepsToReproduce returns the StepsToReproduce field if non-nil, zero value otherwise.

### GetStepsToReproduceOk

`func (o *FindingRequest) GetStepsToReproduceOk() (*string, bool)`

GetStepsToReproduceOk returns a tuple with the StepsToReproduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsToReproduce

`func (o *FindingRequest) SetStepsToReproduce(v string)`

SetStepsToReproduce sets StepsToReproduce field to given value.

### HasStepsToReproduce

`func (o *FindingRequest) HasStepsToReproduce() bool`

HasStepsToReproduce returns a boolean if a field has been set.

### SetStepsToReproduceNil

`func (o *FindingRequest) SetStepsToReproduceNil(b bool)`

 SetStepsToReproduceNil sets the value for StepsToReproduce to be an explicit nil

### UnsetStepsToReproduce
`func (o *FindingRequest) UnsetStepsToReproduce()`

UnsetStepsToReproduce ensures that no value is present for StepsToReproduce, not even an explicit nil
### GetSeverityJustification

`func (o *FindingRequest) GetSeverityJustification() string`

GetSeverityJustification returns the SeverityJustification field if non-nil, zero value otherwise.

### GetSeverityJustificationOk

`func (o *FindingRequest) GetSeverityJustificationOk() (*string, bool)`

GetSeverityJustificationOk returns a tuple with the SeverityJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityJustification

`func (o *FindingRequest) SetSeverityJustification(v string)`

SetSeverityJustification sets SeverityJustification field to given value.

### HasSeverityJustification

`func (o *FindingRequest) HasSeverityJustification() bool`

HasSeverityJustification returns a boolean if a field has been set.

### SetSeverityJustificationNil

`func (o *FindingRequest) SetSeverityJustificationNil(b bool)`

 SetSeverityJustificationNil sets the value for SeverityJustification to be an explicit nil

### UnsetSeverityJustification
`func (o *FindingRequest) UnsetSeverityJustification()`

UnsetSeverityJustification ensures that no value is present for SeverityJustification, not even an explicit nil
### GetReferences

`func (o *FindingRequest) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FindingRequest) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FindingRequest) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *FindingRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *FindingRequest) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *FindingRequest) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetActive

`func (o *FindingRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *FindingRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *FindingRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *FindingRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVerified

`func (o *FindingRequest) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *FindingRequest) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *FindingRequest) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *FindingRequest) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetFalseP

`func (o *FindingRequest) GetFalseP() bool`

GetFalseP returns the FalseP field if non-nil, zero value otherwise.

### GetFalsePOk

`func (o *FindingRequest) GetFalsePOk() (*bool, bool)`

GetFalsePOk returns a tuple with the FalseP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseP

`func (o *FindingRequest) SetFalseP(v bool)`

SetFalseP sets FalseP field to given value.

### HasFalseP

`func (o *FindingRequest) HasFalseP() bool`

HasFalseP returns a boolean if a field has been set.

### GetDuplicate

`func (o *FindingRequest) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *FindingRequest) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *FindingRequest) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *FindingRequest) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetOutOfScope

`func (o *FindingRequest) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *FindingRequest) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *FindingRequest) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *FindingRequest) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *FindingRequest) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *FindingRequest) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *FindingRequest) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *FindingRequest) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetUnderReview

`func (o *FindingRequest) GetUnderReview() bool`

GetUnderReview returns the UnderReview field if non-nil, zero value otherwise.

### GetUnderReviewOk

`func (o *FindingRequest) GetUnderReviewOk() (*bool, bool)`

GetUnderReviewOk returns a tuple with the UnderReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview

`func (o *FindingRequest) SetUnderReview(v bool)`

SetUnderReview sets UnderReview field to given value.

### HasUnderReview

`func (o *FindingRequest) HasUnderReview() bool`

HasUnderReview returns a boolean if a field has been set.

### GetUnderDefectReview

`func (o *FindingRequest) GetUnderDefectReview() bool`

GetUnderDefectReview returns the UnderDefectReview field if non-nil, zero value otherwise.

### GetUnderDefectReviewOk

`func (o *FindingRequest) GetUnderDefectReviewOk() (*bool, bool)`

GetUnderDefectReviewOk returns a tuple with the UnderDefectReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderDefectReview

`func (o *FindingRequest) SetUnderDefectReview(v bool)`

SetUnderDefectReview sets UnderDefectReview field to given value.

### HasUnderDefectReview

`func (o *FindingRequest) HasUnderDefectReview() bool`

HasUnderDefectReview returns a boolean if a field has been set.

### GetIsMitigated

`func (o *FindingRequest) GetIsMitigated() bool`

GetIsMitigated returns the IsMitigated field if non-nil, zero value otherwise.

### GetIsMitigatedOk

`func (o *FindingRequest) GetIsMitigatedOk() (*bool, bool)`

GetIsMitigatedOk returns a tuple with the IsMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMitigated

`func (o *FindingRequest) SetIsMitigated(v bool)`

SetIsMitigated sets IsMitigated field to given value.

### HasIsMitigated

`func (o *FindingRequest) HasIsMitigated() bool`

HasIsMitigated returns a boolean if a field has been set.

### GetNumericalSeverity

`func (o *FindingRequest) GetNumericalSeverity() string`

GetNumericalSeverity returns the NumericalSeverity field if non-nil, zero value otherwise.

### GetNumericalSeverityOk

`func (o *FindingRequest) GetNumericalSeverityOk() (*string, bool)`

GetNumericalSeverityOk returns a tuple with the NumericalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericalSeverity

`func (o *FindingRequest) SetNumericalSeverity(v string)`

SetNumericalSeverity sets NumericalSeverity field to given value.


### GetLine

`func (o *FindingRequest) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *FindingRequest) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *FindingRequest) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *FindingRequest) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *FindingRequest) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *FindingRequest) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetFilePath

`func (o *FindingRequest) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FindingRequest) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FindingRequest) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FindingRequest) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *FindingRequest) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *FindingRequest) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetComponentName

`func (o *FindingRequest) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *FindingRequest) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *FindingRequest) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *FindingRequest) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### SetComponentNameNil

`func (o *FindingRequest) SetComponentNameNil(b bool)`

 SetComponentNameNil sets the value for ComponentName to be an explicit nil

### UnsetComponentName
`func (o *FindingRequest) UnsetComponentName()`

UnsetComponentName ensures that no value is present for ComponentName, not even an explicit nil
### GetComponentVersion

`func (o *FindingRequest) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *FindingRequest) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *FindingRequest) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *FindingRequest) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *FindingRequest) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *FindingRequest) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetStaticFinding

`func (o *FindingRequest) GetStaticFinding() bool`

GetStaticFinding returns the StaticFinding field if non-nil, zero value otherwise.

### GetStaticFindingOk

`func (o *FindingRequest) GetStaticFindingOk() (*bool, bool)`

GetStaticFindingOk returns a tuple with the StaticFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFinding

`func (o *FindingRequest) SetStaticFinding(v bool)`

SetStaticFinding sets StaticFinding field to given value.

### HasStaticFinding

`func (o *FindingRequest) HasStaticFinding() bool`

HasStaticFinding returns a boolean if a field has been set.

### GetDynamicFinding

`func (o *FindingRequest) GetDynamicFinding() bool`

GetDynamicFinding returns the DynamicFinding field if non-nil, zero value otherwise.

### GetDynamicFindingOk

`func (o *FindingRequest) GetDynamicFindingOk() (*bool, bool)`

GetDynamicFindingOk returns a tuple with the DynamicFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicFinding

`func (o *FindingRequest) SetDynamicFinding(v bool)`

SetDynamicFinding sets DynamicFinding field to given value.

### HasDynamicFinding

`func (o *FindingRequest) HasDynamicFinding() bool`

HasDynamicFinding returns a boolean if a field has been set.

### GetUniqueIdFromTool

`func (o *FindingRequest) GetUniqueIdFromTool() string`

GetUniqueIdFromTool returns the UniqueIdFromTool field if non-nil, zero value otherwise.

### GetUniqueIdFromToolOk

`func (o *FindingRequest) GetUniqueIdFromToolOk() (*string, bool)`

GetUniqueIdFromToolOk returns a tuple with the UniqueIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdFromTool

`func (o *FindingRequest) SetUniqueIdFromTool(v string)`

SetUniqueIdFromTool sets UniqueIdFromTool field to given value.

### HasUniqueIdFromTool

`func (o *FindingRequest) HasUniqueIdFromTool() bool`

HasUniqueIdFromTool returns a boolean if a field has been set.

### SetUniqueIdFromToolNil

`func (o *FindingRequest) SetUniqueIdFromToolNil(b bool)`

 SetUniqueIdFromToolNil sets the value for UniqueIdFromTool to be an explicit nil

### UnsetUniqueIdFromTool
`func (o *FindingRequest) UnsetUniqueIdFromTool()`

UnsetUniqueIdFromTool ensures that no value is present for UniqueIdFromTool, not even an explicit nil
### GetVulnIdFromTool

`func (o *FindingRequest) GetVulnIdFromTool() string`

GetVulnIdFromTool returns the VulnIdFromTool field if non-nil, zero value otherwise.

### GetVulnIdFromToolOk

`func (o *FindingRequest) GetVulnIdFromToolOk() (*string, bool)`

GetVulnIdFromToolOk returns a tuple with the VulnIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnIdFromTool

`func (o *FindingRequest) SetVulnIdFromTool(v string)`

SetVulnIdFromTool sets VulnIdFromTool field to given value.

### HasVulnIdFromTool

`func (o *FindingRequest) HasVulnIdFromTool() bool`

HasVulnIdFromTool returns a boolean if a field has been set.

### SetVulnIdFromToolNil

`func (o *FindingRequest) SetVulnIdFromToolNil(b bool)`

 SetVulnIdFromToolNil sets the value for VulnIdFromTool to be an explicit nil

### UnsetVulnIdFromTool
`func (o *FindingRequest) UnsetVulnIdFromTool()`

UnsetVulnIdFromTool ensures that no value is present for VulnIdFromTool, not even an explicit nil
### GetSastSourceObject

`func (o *FindingRequest) GetSastSourceObject() string`

GetSastSourceObject returns the SastSourceObject field if non-nil, zero value otherwise.

### GetSastSourceObjectOk

`func (o *FindingRequest) GetSastSourceObjectOk() (*string, bool)`

GetSastSourceObjectOk returns a tuple with the SastSourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceObject

`func (o *FindingRequest) SetSastSourceObject(v string)`

SetSastSourceObject sets SastSourceObject field to given value.

### HasSastSourceObject

`func (o *FindingRequest) HasSastSourceObject() bool`

HasSastSourceObject returns a boolean if a field has been set.

### SetSastSourceObjectNil

`func (o *FindingRequest) SetSastSourceObjectNil(b bool)`

 SetSastSourceObjectNil sets the value for SastSourceObject to be an explicit nil

### UnsetSastSourceObject
`func (o *FindingRequest) UnsetSastSourceObject()`

UnsetSastSourceObject ensures that no value is present for SastSourceObject, not even an explicit nil
### GetSastSinkObject

`func (o *FindingRequest) GetSastSinkObject() string`

GetSastSinkObject returns the SastSinkObject field if non-nil, zero value otherwise.

### GetSastSinkObjectOk

`func (o *FindingRequest) GetSastSinkObjectOk() (*string, bool)`

GetSastSinkObjectOk returns a tuple with the SastSinkObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSinkObject

`func (o *FindingRequest) SetSastSinkObject(v string)`

SetSastSinkObject sets SastSinkObject field to given value.

### HasSastSinkObject

`func (o *FindingRequest) HasSastSinkObject() bool`

HasSastSinkObject returns a boolean if a field has been set.

### SetSastSinkObjectNil

`func (o *FindingRequest) SetSastSinkObjectNil(b bool)`

 SetSastSinkObjectNil sets the value for SastSinkObject to be an explicit nil

### UnsetSastSinkObject
`func (o *FindingRequest) UnsetSastSinkObject()`

UnsetSastSinkObject ensures that no value is present for SastSinkObject, not even an explicit nil
### GetSastSourceLine

`func (o *FindingRequest) GetSastSourceLine() int32`

GetSastSourceLine returns the SastSourceLine field if non-nil, zero value otherwise.

### GetSastSourceLineOk

`func (o *FindingRequest) GetSastSourceLineOk() (*int32, bool)`

GetSastSourceLineOk returns a tuple with the SastSourceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceLine

`func (o *FindingRequest) SetSastSourceLine(v int32)`

SetSastSourceLine sets SastSourceLine field to given value.

### HasSastSourceLine

`func (o *FindingRequest) HasSastSourceLine() bool`

HasSastSourceLine returns a boolean if a field has been set.

### SetSastSourceLineNil

`func (o *FindingRequest) SetSastSourceLineNil(b bool)`

 SetSastSourceLineNil sets the value for SastSourceLine to be an explicit nil

### UnsetSastSourceLine
`func (o *FindingRequest) UnsetSastSourceLine()`

UnsetSastSourceLine ensures that no value is present for SastSourceLine, not even an explicit nil
### GetSastSourceFilePath

`func (o *FindingRequest) GetSastSourceFilePath() string`

GetSastSourceFilePath returns the SastSourceFilePath field if non-nil, zero value otherwise.

### GetSastSourceFilePathOk

`func (o *FindingRequest) GetSastSourceFilePathOk() (*string, bool)`

GetSastSourceFilePathOk returns a tuple with the SastSourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceFilePath

`func (o *FindingRequest) SetSastSourceFilePath(v string)`

SetSastSourceFilePath sets SastSourceFilePath field to given value.

### HasSastSourceFilePath

`func (o *FindingRequest) HasSastSourceFilePath() bool`

HasSastSourceFilePath returns a boolean if a field has been set.

### SetSastSourceFilePathNil

`func (o *FindingRequest) SetSastSourceFilePathNil(b bool)`

 SetSastSourceFilePathNil sets the value for SastSourceFilePath to be an explicit nil

### UnsetSastSourceFilePath
`func (o *FindingRequest) UnsetSastSourceFilePath()`

UnsetSastSourceFilePath ensures that no value is present for SastSourceFilePath, not even an explicit nil
### GetNbOccurences

`func (o *FindingRequest) GetNbOccurences() int32`

GetNbOccurences returns the NbOccurences field if non-nil, zero value otherwise.

### GetNbOccurencesOk

`func (o *FindingRequest) GetNbOccurencesOk() (*int32, bool)`

GetNbOccurencesOk returns a tuple with the NbOccurences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOccurences

`func (o *FindingRequest) SetNbOccurences(v int32)`

SetNbOccurences sets NbOccurences field to given value.

### HasNbOccurences

`func (o *FindingRequest) HasNbOccurences() bool`

HasNbOccurences returns a boolean if a field has been set.

### SetNbOccurencesNil

`func (o *FindingRequest) SetNbOccurencesNil(b bool)`

 SetNbOccurencesNil sets the value for NbOccurences to be an explicit nil

### UnsetNbOccurences
`func (o *FindingRequest) UnsetNbOccurences()`

UnsetNbOccurences ensures that no value is present for NbOccurences, not even an explicit nil
### GetPublishDate

`func (o *FindingRequest) GetPublishDate() string`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *FindingRequest) GetPublishDateOk() (*string, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *FindingRequest) SetPublishDate(v string)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *FindingRequest) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### SetPublishDateNil

`func (o *FindingRequest) SetPublishDateNil(b bool)`

 SetPublishDateNil sets the value for PublishDate to be an explicit nil

### UnsetPublishDate
`func (o *FindingRequest) UnsetPublishDate()`

UnsetPublishDate ensures that no value is present for PublishDate, not even an explicit nil
### GetService

`func (o *FindingRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *FindingRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *FindingRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *FindingRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *FindingRequest) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *FindingRequest) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetPlannedRemediationDate

`func (o *FindingRequest) GetPlannedRemediationDate() string`

GetPlannedRemediationDate returns the PlannedRemediationDate field if non-nil, zero value otherwise.

### GetPlannedRemediationDateOk

`func (o *FindingRequest) GetPlannedRemediationDateOk() (*string, bool)`

GetPlannedRemediationDateOk returns a tuple with the PlannedRemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationDate

`func (o *FindingRequest) SetPlannedRemediationDate(v string)`

SetPlannedRemediationDate sets PlannedRemediationDate field to given value.

### HasPlannedRemediationDate

`func (o *FindingRequest) HasPlannedRemediationDate() bool`

HasPlannedRemediationDate returns a boolean if a field has been set.

### SetPlannedRemediationDateNil

`func (o *FindingRequest) SetPlannedRemediationDateNil(b bool)`

 SetPlannedRemediationDateNil sets the value for PlannedRemediationDate to be an explicit nil

### UnsetPlannedRemediationDate
`func (o *FindingRequest) UnsetPlannedRemediationDate()`

UnsetPlannedRemediationDate ensures that no value is present for PlannedRemediationDate, not even an explicit nil
### GetPlannedRemediationVersion

`func (o *FindingRequest) GetPlannedRemediationVersion() string`

GetPlannedRemediationVersion returns the PlannedRemediationVersion field if non-nil, zero value otherwise.

### GetPlannedRemediationVersionOk

`func (o *FindingRequest) GetPlannedRemediationVersionOk() (*string, bool)`

GetPlannedRemediationVersionOk returns a tuple with the PlannedRemediationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationVersion

`func (o *FindingRequest) SetPlannedRemediationVersion(v string)`

SetPlannedRemediationVersion sets PlannedRemediationVersion field to given value.

### HasPlannedRemediationVersion

`func (o *FindingRequest) HasPlannedRemediationVersion() bool`

HasPlannedRemediationVersion returns a boolean if a field has been set.

### SetPlannedRemediationVersionNil

`func (o *FindingRequest) SetPlannedRemediationVersionNil(b bool)`

 SetPlannedRemediationVersionNil sets the value for PlannedRemediationVersion to be an explicit nil

### UnsetPlannedRemediationVersion
`func (o *FindingRequest) UnsetPlannedRemediationVersion()`

UnsetPlannedRemediationVersion ensures that no value is present for PlannedRemediationVersion, not even an explicit nil
### GetEffortForFixing

`func (o *FindingRequest) GetEffortForFixing() string`

GetEffortForFixing returns the EffortForFixing field if non-nil, zero value otherwise.

### GetEffortForFixingOk

`func (o *FindingRequest) GetEffortForFixingOk() (*string, bool)`

GetEffortForFixingOk returns a tuple with the EffortForFixing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortForFixing

`func (o *FindingRequest) SetEffortForFixing(v string)`

SetEffortForFixing sets EffortForFixing field to given value.

### HasEffortForFixing

`func (o *FindingRequest) HasEffortForFixing() bool`

HasEffortForFixing returns a boolean if a field has been set.

### SetEffortForFixingNil

`func (o *FindingRequest) SetEffortForFixingNil(b bool)`

 SetEffortForFixingNil sets the value for EffortForFixing to be an explicit nil

### UnsetEffortForFixing
`func (o *FindingRequest) UnsetEffortForFixing()`

UnsetEffortForFixing ensures that no value is present for EffortForFixing, not even an explicit nil
### GetReviewRequestedBy

`func (o *FindingRequest) GetReviewRequestedBy() int32`

GetReviewRequestedBy returns the ReviewRequestedBy field if non-nil, zero value otherwise.

### GetReviewRequestedByOk

`func (o *FindingRequest) GetReviewRequestedByOk() (*int32, bool)`

GetReviewRequestedByOk returns a tuple with the ReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequestedBy

`func (o *FindingRequest) SetReviewRequestedBy(v int32)`

SetReviewRequestedBy sets ReviewRequestedBy field to given value.

### HasReviewRequestedBy

`func (o *FindingRequest) HasReviewRequestedBy() bool`

HasReviewRequestedBy returns a boolean if a field has been set.

### SetReviewRequestedByNil

`func (o *FindingRequest) SetReviewRequestedByNil(b bool)`

 SetReviewRequestedByNil sets the value for ReviewRequestedBy to be an explicit nil

### UnsetReviewRequestedBy
`func (o *FindingRequest) UnsetReviewRequestedBy()`

UnsetReviewRequestedBy ensures that no value is present for ReviewRequestedBy, not even an explicit nil
### GetDefectReviewRequestedBy

`func (o *FindingRequest) GetDefectReviewRequestedBy() int32`

GetDefectReviewRequestedBy returns the DefectReviewRequestedBy field if non-nil, zero value otherwise.

### GetDefectReviewRequestedByOk

`func (o *FindingRequest) GetDefectReviewRequestedByOk() (*int32, bool)`

GetDefectReviewRequestedByOk returns a tuple with the DefectReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectReviewRequestedBy

`func (o *FindingRequest) SetDefectReviewRequestedBy(v int32)`

SetDefectReviewRequestedBy sets DefectReviewRequestedBy field to given value.

### HasDefectReviewRequestedBy

`func (o *FindingRequest) HasDefectReviewRequestedBy() bool`

HasDefectReviewRequestedBy returns a boolean if a field has been set.

### SetDefectReviewRequestedByNil

`func (o *FindingRequest) SetDefectReviewRequestedByNil(b bool)`

 SetDefectReviewRequestedByNil sets the value for DefectReviewRequestedBy to be an explicit nil

### UnsetDefectReviewRequestedBy
`func (o *FindingRequest) UnsetDefectReviewRequestedBy()`

UnsetDefectReviewRequestedBy ensures that no value is present for DefectReviewRequestedBy, not even an explicit nil
### GetSonarqubeIssue

`func (o *FindingRequest) GetSonarqubeIssue() int32`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *FindingRequest) GetSonarqubeIssueOk() (*int32, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *FindingRequest) SetSonarqubeIssue(v int32)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.

### HasSonarqubeIssue

`func (o *FindingRequest) HasSonarqubeIssue() bool`

HasSonarqubeIssue returns a boolean if a field has been set.

### SetSonarqubeIssueNil

`func (o *FindingRequest) SetSonarqubeIssueNil(b bool)`

 SetSonarqubeIssueNil sets the value for SonarqubeIssue to be an explicit nil

### UnsetSonarqubeIssue
`func (o *FindingRequest) UnsetSonarqubeIssue()`

UnsetSonarqubeIssue ensures that no value is present for SonarqubeIssue, not even an explicit nil
### GetReviewers

`func (o *FindingRequest) GetReviewers() []int32`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *FindingRequest) GetReviewersOk() (*[]int32, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *FindingRequest) SetReviewers(v []int32)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *FindingRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


