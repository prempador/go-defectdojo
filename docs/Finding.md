# Finding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**RequestResponse** | [**BurpRawRequestResponse**](BurpRawRequestResponse.md) |  | [readonly] 
**AcceptedRisks** | [**[]RiskAcceptance**](RiskAcceptance.md) |  | [readonly] 
**PushToJira** | Pointer to **bool** |  | [optional] [default to false]
**Age** | **int32** |  | [readonly] 
**SlaDaysRemaining** | **int32** |  | [readonly] 
**FindingMeta** | [**[]FindingMeta**](FindingMeta.md) |  | [readonly] 
**RelatedFields** | [**FindingRelatedFields**](FindingRelatedFields.md) |  | [readonly] 
**JiraCreation** | **time.Time** |  | [readonly] 
**JiraChange** | **time.Time** |  | [readonly] 
**DisplayStatus** | **string** |  | [readonly] 
**FindingGroups** | [**[]FindingGroup**](FindingGroup.md) |  | [readonly] 
**VulnerabilityIds** | Pointer to [**[]VulnerabilityId**](VulnerabilityId.md) |  | [optional] 
**Reporter** | Pointer to **int32** |  | [optional] 
**Title** | **string** | A short description of the flaw. | 
**Date** | Pointer to **string** | The date the flaw was discovered. | [optional] 
**SlaStartDate** | Pointer to **NullableString** | (readonly)The date used as start date for SLA calculation. Set by expiring risk acceptances. Empty by default, causing a fallback to &#39;date&#39;. | [optional] 
**SlaExpirationDate** | Pointer to **NullableString** | (readonly)The date SLA expires for this finding. Empty by default, causing a fallback to &#39;date&#39;. | [optional] 
**Cwe** | Pointer to **NullableInt32** | The CWE number associated with this flaw. | [optional] 
**Cvssv3** | Pointer to **NullableString** | Common Vulnerability Scoring System version 3 (CVSSv3) score associated with this flaw. | [optional] 
**Cvssv3Score** | Pointer to **NullableFloat64** | Numerical CVSSv3 score for the vulnerability. If the vector is given, the score is updated while saving the finding | [optional] 
**Url** | **NullableString** | External reference that provides more information about this flaw. | [readonly] 
**Severity** | **string** | The severity level of this flaw (Critical, High, Medium, Low, Informational). | 
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
**LastStatusUpdate** | **NullableTime** | Timestamp of latest status update (change in status related fields). | [readonly] 
**UnderDefectReview** | Pointer to **bool** | Denotes if this finding is under defect review. | [optional] 
**IsMitigated** | Pointer to **bool** | Denotes if this flaw has been fixed. | [optional] 
**ThreadId** | **int32** |  | [readonly] 
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
**Test** | **int32** | The test that is associated with this flaw. | [readonly] 
**DuplicateFinding** | **NullableInt32** | Link to the original finding if this finding is a duplicate. | [readonly] 
**ReviewRequestedBy** | Pointer to **NullableInt32** | Documents who requested a review for this finding. | [optional] 
**DefectReviewRequestedBy** | Pointer to **NullableInt32** | Documents who requested a defect review for this flaw. | [optional] 
**MitigatedBy** | **NullableInt32** | Documents who has marked this flaw as fixed. | [readonly] 
**LastReviewedBy** | **NullableInt32** | Provides the person who last reviewed the flaw. | [readonly] 
**SonarqubeIssue** | Pointer to **NullableInt32** | The SonarQube issue associated with this finding. | [optional] 
**Endpoints** | **[]int32** | The hosts within the product that are susceptible to this flaw. + The status of the endpoint associated with this flaw (Vulnerable, Mitigated, ...). | [readonly] 
**Reviewers** | Pointer to **[]int32** | Documents who reviewed the flaw. | [optional] 
**Notes** | [**[]Note**](Note.md) |  | [readonly] 
**Files** | **[]int32** | Files(s) related to the flaw. | [readonly] 
**FoundBy** | **[]int32** | The name of the scanner that identified the flaw. | [readonly] 
**Prefetch** | Pointer to [**FindingPrefetch**](FindingPrefetch.md) |  | [optional] 

## Methods

### NewFinding

`func NewFinding(id int32, requestResponse BurpRawRequestResponse, acceptedRisks []RiskAcceptance, age int32, slaDaysRemaining int32, findingMeta []FindingMeta, relatedFields FindingRelatedFields, jiraCreation time.Time, jiraChange time.Time, displayStatus string, findingGroups []FindingGroup, title string, url NullableString, severity string, description string, lastStatusUpdate NullableTime, threadId int32, mitigated NullableTime, numericalSeverity string, lastReviewed NullableTime, param NullableString, payload NullableString, hashCode NullableString, created NullableTime, scannerConfidence NullableInt32, test int32, duplicateFinding NullableInt32, mitigatedBy NullableInt32, lastReviewedBy NullableInt32, endpoints []int32, notes []Note, files []int32, foundBy []int32, ) *Finding`

NewFinding instantiates a new Finding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingWithDefaults

`func NewFindingWithDefaults() *Finding`

NewFindingWithDefaults instantiates a new Finding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Finding) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Finding) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Finding) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *Finding) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Finding) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Finding) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Finding) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRequestResponse

`func (o *Finding) GetRequestResponse() BurpRawRequestResponse`

GetRequestResponse returns the RequestResponse field if non-nil, zero value otherwise.

### GetRequestResponseOk

`func (o *Finding) GetRequestResponseOk() (*BurpRawRequestResponse, bool)`

GetRequestResponseOk returns a tuple with the RequestResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestResponse

`func (o *Finding) SetRequestResponse(v BurpRawRequestResponse)`

SetRequestResponse sets RequestResponse field to given value.


### GetAcceptedRisks

`func (o *Finding) GetAcceptedRisks() []RiskAcceptance`

GetAcceptedRisks returns the AcceptedRisks field if non-nil, zero value otherwise.

### GetAcceptedRisksOk

`func (o *Finding) GetAcceptedRisksOk() (*[]RiskAcceptance, bool)`

GetAcceptedRisksOk returns a tuple with the AcceptedRisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedRisks

`func (o *Finding) SetAcceptedRisks(v []RiskAcceptance)`

SetAcceptedRisks sets AcceptedRisks field to given value.


### GetPushToJira

`func (o *Finding) GetPushToJira() bool`

GetPushToJira returns the PushToJira field if non-nil, zero value otherwise.

### GetPushToJiraOk

`func (o *Finding) GetPushToJiraOk() (*bool, bool)`

GetPushToJiraOk returns a tuple with the PushToJira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToJira

`func (o *Finding) SetPushToJira(v bool)`

SetPushToJira sets PushToJira field to given value.

### HasPushToJira

`func (o *Finding) HasPushToJira() bool`

HasPushToJira returns a boolean if a field has been set.

### GetAge

`func (o *Finding) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Finding) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Finding) SetAge(v int32)`

SetAge sets Age field to given value.


### GetSlaDaysRemaining

`func (o *Finding) GetSlaDaysRemaining() int32`

GetSlaDaysRemaining returns the SlaDaysRemaining field if non-nil, zero value otherwise.

### GetSlaDaysRemainingOk

`func (o *Finding) GetSlaDaysRemainingOk() (*int32, bool)`

GetSlaDaysRemainingOk returns a tuple with the SlaDaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaDaysRemaining

`func (o *Finding) SetSlaDaysRemaining(v int32)`

SetSlaDaysRemaining sets SlaDaysRemaining field to given value.


### GetFindingMeta

`func (o *Finding) GetFindingMeta() []FindingMeta`

GetFindingMeta returns the FindingMeta field if non-nil, zero value otherwise.

### GetFindingMetaOk

`func (o *Finding) GetFindingMetaOk() (*[]FindingMeta, bool)`

GetFindingMetaOk returns a tuple with the FindingMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingMeta

`func (o *Finding) SetFindingMeta(v []FindingMeta)`

SetFindingMeta sets FindingMeta field to given value.


### GetRelatedFields

`func (o *Finding) GetRelatedFields() FindingRelatedFields`

GetRelatedFields returns the RelatedFields field if non-nil, zero value otherwise.

### GetRelatedFieldsOk

`func (o *Finding) GetRelatedFieldsOk() (*FindingRelatedFields, bool)`

GetRelatedFieldsOk returns a tuple with the RelatedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedFields

`func (o *Finding) SetRelatedFields(v FindingRelatedFields)`

SetRelatedFields sets RelatedFields field to given value.


### GetJiraCreation

`func (o *Finding) GetJiraCreation() time.Time`

GetJiraCreation returns the JiraCreation field if non-nil, zero value otherwise.

### GetJiraCreationOk

`func (o *Finding) GetJiraCreationOk() (*time.Time, bool)`

GetJiraCreationOk returns a tuple with the JiraCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraCreation

`func (o *Finding) SetJiraCreation(v time.Time)`

SetJiraCreation sets JiraCreation field to given value.


### GetJiraChange

`func (o *Finding) GetJiraChange() time.Time`

GetJiraChange returns the JiraChange field if non-nil, zero value otherwise.

### GetJiraChangeOk

`func (o *Finding) GetJiraChangeOk() (*time.Time, bool)`

GetJiraChangeOk returns a tuple with the JiraChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraChange

`func (o *Finding) SetJiraChange(v time.Time)`

SetJiraChange sets JiraChange field to given value.


### GetDisplayStatus

`func (o *Finding) GetDisplayStatus() string`

GetDisplayStatus returns the DisplayStatus field if non-nil, zero value otherwise.

### GetDisplayStatusOk

`func (o *Finding) GetDisplayStatusOk() (*string, bool)`

GetDisplayStatusOk returns a tuple with the DisplayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStatus

`func (o *Finding) SetDisplayStatus(v string)`

SetDisplayStatus sets DisplayStatus field to given value.


### GetFindingGroups

`func (o *Finding) GetFindingGroups() []FindingGroup`

GetFindingGroups returns the FindingGroups field if non-nil, zero value otherwise.

### GetFindingGroupsOk

`func (o *Finding) GetFindingGroupsOk() (*[]FindingGroup, bool)`

GetFindingGroupsOk returns a tuple with the FindingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingGroups

`func (o *Finding) SetFindingGroups(v []FindingGroup)`

SetFindingGroups sets FindingGroups field to given value.


### GetVulnerabilityIds

`func (o *Finding) GetVulnerabilityIds() []VulnerabilityId`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *Finding) GetVulnerabilityIdsOk() (*[]VulnerabilityId, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *Finding) SetVulnerabilityIds(v []VulnerabilityId)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *Finding) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetReporter

`func (o *Finding) GetReporter() int32`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *Finding) GetReporterOk() (*int32, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *Finding) SetReporter(v int32)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *Finding) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetTitle

`func (o *Finding) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Finding) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Finding) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDate

`func (o *Finding) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Finding) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Finding) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Finding) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSlaStartDate

`func (o *Finding) GetSlaStartDate() string`

GetSlaStartDate returns the SlaStartDate field if non-nil, zero value otherwise.

### GetSlaStartDateOk

`func (o *Finding) GetSlaStartDateOk() (*string, bool)`

GetSlaStartDateOk returns a tuple with the SlaStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaStartDate

`func (o *Finding) SetSlaStartDate(v string)`

SetSlaStartDate sets SlaStartDate field to given value.

### HasSlaStartDate

`func (o *Finding) HasSlaStartDate() bool`

HasSlaStartDate returns a boolean if a field has been set.

### SetSlaStartDateNil

`func (o *Finding) SetSlaStartDateNil(b bool)`

 SetSlaStartDateNil sets the value for SlaStartDate to be an explicit nil

### UnsetSlaStartDate
`func (o *Finding) UnsetSlaStartDate()`

UnsetSlaStartDate ensures that no value is present for SlaStartDate, not even an explicit nil
### GetSlaExpirationDate

`func (o *Finding) GetSlaExpirationDate() string`

GetSlaExpirationDate returns the SlaExpirationDate field if non-nil, zero value otherwise.

### GetSlaExpirationDateOk

`func (o *Finding) GetSlaExpirationDateOk() (*string, bool)`

GetSlaExpirationDateOk returns a tuple with the SlaExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaExpirationDate

`func (o *Finding) SetSlaExpirationDate(v string)`

SetSlaExpirationDate sets SlaExpirationDate field to given value.

### HasSlaExpirationDate

`func (o *Finding) HasSlaExpirationDate() bool`

HasSlaExpirationDate returns a boolean if a field has been set.

### SetSlaExpirationDateNil

`func (o *Finding) SetSlaExpirationDateNil(b bool)`

 SetSlaExpirationDateNil sets the value for SlaExpirationDate to be an explicit nil

### UnsetSlaExpirationDate
`func (o *Finding) UnsetSlaExpirationDate()`

UnsetSlaExpirationDate ensures that no value is present for SlaExpirationDate, not even an explicit nil
### GetCwe

`func (o *Finding) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *Finding) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *Finding) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *Finding) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *Finding) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *Finding) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetCvssv3

`func (o *Finding) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *Finding) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *Finding) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *Finding) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *Finding) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *Finding) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetCvssv3Score

`func (o *Finding) GetCvssv3Score() float64`

GetCvssv3Score returns the Cvssv3Score field if non-nil, zero value otherwise.

### GetCvssv3ScoreOk

`func (o *Finding) GetCvssv3ScoreOk() (*float64, bool)`

GetCvssv3ScoreOk returns a tuple with the Cvssv3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3Score

`func (o *Finding) SetCvssv3Score(v float64)`

SetCvssv3Score sets Cvssv3Score field to given value.

### HasCvssv3Score

`func (o *Finding) HasCvssv3Score() bool`

HasCvssv3Score returns a boolean if a field has been set.

### SetCvssv3ScoreNil

`func (o *Finding) SetCvssv3ScoreNil(b bool)`

 SetCvssv3ScoreNil sets the value for Cvssv3Score to be an explicit nil

### UnsetCvssv3Score
`func (o *Finding) UnsetCvssv3Score()`

UnsetCvssv3Score ensures that no value is present for Cvssv3Score, not even an explicit nil
### GetUrl

`func (o *Finding) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Finding) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Finding) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *Finding) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Finding) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSeverity

`func (o *Finding) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Finding) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Finding) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetDescription

`func (o *Finding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Finding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Finding) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMitigation

`func (o *Finding) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *Finding) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *Finding) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *Finding) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *Finding) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *Finding) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *Finding) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *Finding) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *Finding) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *Finding) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *Finding) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *Finding) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetStepsToReproduce

`func (o *Finding) GetStepsToReproduce() string`

GetStepsToReproduce returns the StepsToReproduce field if non-nil, zero value otherwise.

### GetStepsToReproduceOk

`func (o *Finding) GetStepsToReproduceOk() (*string, bool)`

GetStepsToReproduceOk returns a tuple with the StepsToReproduce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsToReproduce

`func (o *Finding) SetStepsToReproduce(v string)`

SetStepsToReproduce sets StepsToReproduce field to given value.

### HasStepsToReproduce

`func (o *Finding) HasStepsToReproduce() bool`

HasStepsToReproduce returns a boolean if a field has been set.

### SetStepsToReproduceNil

`func (o *Finding) SetStepsToReproduceNil(b bool)`

 SetStepsToReproduceNil sets the value for StepsToReproduce to be an explicit nil

### UnsetStepsToReproduce
`func (o *Finding) UnsetStepsToReproduce()`

UnsetStepsToReproduce ensures that no value is present for StepsToReproduce, not even an explicit nil
### GetSeverityJustification

`func (o *Finding) GetSeverityJustification() string`

GetSeverityJustification returns the SeverityJustification field if non-nil, zero value otherwise.

### GetSeverityJustificationOk

`func (o *Finding) GetSeverityJustificationOk() (*string, bool)`

GetSeverityJustificationOk returns a tuple with the SeverityJustification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityJustification

`func (o *Finding) SetSeverityJustification(v string)`

SetSeverityJustification sets SeverityJustification field to given value.

### HasSeverityJustification

`func (o *Finding) HasSeverityJustification() bool`

HasSeverityJustification returns a boolean if a field has been set.

### SetSeverityJustificationNil

`func (o *Finding) SetSeverityJustificationNil(b bool)`

 SetSeverityJustificationNil sets the value for SeverityJustification to be an explicit nil

### UnsetSeverityJustification
`func (o *Finding) UnsetSeverityJustification()`

UnsetSeverityJustification ensures that no value is present for SeverityJustification, not even an explicit nil
### GetReferences

`func (o *Finding) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Finding) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Finding) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Finding) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *Finding) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *Finding) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetActive

`func (o *Finding) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Finding) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Finding) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Finding) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVerified

`func (o *Finding) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Finding) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Finding) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Finding) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetFalseP

`func (o *Finding) GetFalseP() bool`

GetFalseP returns the FalseP field if non-nil, zero value otherwise.

### GetFalsePOk

`func (o *Finding) GetFalsePOk() (*bool, bool)`

GetFalsePOk returns a tuple with the FalseP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseP

`func (o *Finding) SetFalseP(v bool)`

SetFalseP sets FalseP field to given value.

### HasFalseP

`func (o *Finding) HasFalseP() bool`

HasFalseP returns a boolean if a field has been set.

### GetDuplicate

`func (o *Finding) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *Finding) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *Finding) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *Finding) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.

### GetOutOfScope

`func (o *Finding) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *Finding) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *Finding) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *Finding) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *Finding) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *Finding) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *Finding) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *Finding) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetUnderReview

`func (o *Finding) GetUnderReview() bool`

GetUnderReview returns the UnderReview field if non-nil, zero value otherwise.

### GetUnderReviewOk

`func (o *Finding) GetUnderReviewOk() (*bool, bool)`

GetUnderReviewOk returns a tuple with the UnderReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderReview

`func (o *Finding) SetUnderReview(v bool)`

SetUnderReview sets UnderReview field to given value.

### HasUnderReview

`func (o *Finding) HasUnderReview() bool`

HasUnderReview returns a boolean if a field has been set.

### GetLastStatusUpdate

`func (o *Finding) GetLastStatusUpdate() time.Time`

GetLastStatusUpdate returns the LastStatusUpdate field if non-nil, zero value otherwise.

### GetLastStatusUpdateOk

`func (o *Finding) GetLastStatusUpdateOk() (*time.Time, bool)`

GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdate

`func (o *Finding) SetLastStatusUpdate(v time.Time)`

SetLastStatusUpdate sets LastStatusUpdate field to given value.


### SetLastStatusUpdateNil

`func (o *Finding) SetLastStatusUpdateNil(b bool)`

 SetLastStatusUpdateNil sets the value for LastStatusUpdate to be an explicit nil

### UnsetLastStatusUpdate
`func (o *Finding) UnsetLastStatusUpdate()`

UnsetLastStatusUpdate ensures that no value is present for LastStatusUpdate, not even an explicit nil
### GetUnderDefectReview

`func (o *Finding) GetUnderDefectReview() bool`

GetUnderDefectReview returns the UnderDefectReview field if non-nil, zero value otherwise.

### GetUnderDefectReviewOk

`func (o *Finding) GetUnderDefectReviewOk() (*bool, bool)`

GetUnderDefectReviewOk returns a tuple with the UnderDefectReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderDefectReview

`func (o *Finding) SetUnderDefectReview(v bool)`

SetUnderDefectReview sets UnderDefectReview field to given value.

### HasUnderDefectReview

`func (o *Finding) HasUnderDefectReview() bool`

HasUnderDefectReview returns a boolean if a field has been set.

### GetIsMitigated

`func (o *Finding) GetIsMitigated() bool`

GetIsMitigated returns the IsMitigated field if non-nil, zero value otherwise.

### GetIsMitigatedOk

`func (o *Finding) GetIsMitigatedOk() (*bool, bool)`

GetIsMitigatedOk returns a tuple with the IsMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMitigated

`func (o *Finding) SetIsMitigated(v bool)`

SetIsMitigated sets IsMitigated field to given value.

### HasIsMitigated

`func (o *Finding) HasIsMitigated() bool`

HasIsMitigated returns a boolean if a field has been set.

### GetThreadId

`func (o *Finding) GetThreadId() int32`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *Finding) GetThreadIdOk() (*int32, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *Finding) SetThreadId(v int32)`

SetThreadId sets ThreadId field to given value.


### GetMitigated

`func (o *Finding) GetMitigated() time.Time`

GetMitigated returns the Mitigated field if non-nil, zero value otherwise.

### GetMitigatedOk

`func (o *Finding) GetMitigatedOk() (*time.Time, bool)`

GetMitigatedOk returns a tuple with the Mitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigated

`func (o *Finding) SetMitigated(v time.Time)`

SetMitigated sets Mitigated field to given value.


### SetMitigatedNil

`func (o *Finding) SetMitigatedNil(b bool)`

 SetMitigatedNil sets the value for Mitigated to be an explicit nil

### UnsetMitigated
`func (o *Finding) UnsetMitigated()`

UnsetMitigated ensures that no value is present for Mitigated, not even an explicit nil
### GetNumericalSeverity

`func (o *Finding) GetNumericalSeverity() string`

GetNumericalSeverity returns the NumericalSeverity field if non-nil, zero value otherwise.

### GetNumericalSeverityOk

`func (o *Finding) GetNumericalSeverityOk() (*string, bool)`

GetNumericalSeverityOk returns a tuple with the NumericalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericalSeverity

`func (o *Finding) SetNumericalSeverity(v string)`

SetNumericalSeverity sets NumericalSeverity field to given value.


### GetLastReviewed

`func (o *Finding) GetLastReviewed() time.Time`

GetLastReviewed returns the LastReviewed field if non-nil, zero value otherwise.

### GetLastReviewedOk

`func (o *Finding) GetLastReviewedOk() (*time.Time, bool)`

GetLastReviewedOk returns a tuple with the LastReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewed

`func (o *Finding) SetLastReviewed(v time.Time)`

SetLastReviewed sets LastReviewed field to given value.


### SetLastReviewedNil

`func (o *Finding) SetLastReviewedNil(b bool)`

 SetLastReviewedNil sets the value for LastReviewed to be an explicit nil

### UnsetLastReviewed
`func (o *Finding) UnsetLastReviewed()`

UnsetLastReviewed ensures that no value is present for LastReviewed, not even an explicit nil
### GetParam

`func (o *Finding) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *Finding) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *Finding) SetParam(v string)`

SetParam sets Param field to given value.


### SetParamNil

`func (o *Finding) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *Finding) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetPayload

`func (o *Finding) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Finding) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Finding) SetPayload(v string)`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *Finding) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Finding) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetHashCode

`func (o *Finding) GetHashCode() string`

GetHashCode returns the HashCode field if non-nil, zero value otherwise.

### GetHashCodeOk

`func (o *Finding) GetHashCodeOk() (*string, bool)`

GetHashCodeOk returns a tuple with the HashCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashCode

`func (o *Finding) SetHashCode(v string)`

SetHashCode sets HashCode field to given value.


### SetHashCodeNil

`func (o *Finding) SetHashCodeNil(b bool)`

 SetHashCodeNil sets the value for HashCode to be an explicit nil

### UnsetHashCode
`func (o *Finding) UnsetHashCode()`

UnsetHashCode ensures that no value is present for HashCode, not even an explicit nil
### GetLine

`func (o *Finding) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Finding) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Finding) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *Finding) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLineNil

`func (o *Finding) SetLineNil(b bool)`

 SetLineNil sets the value for Line to be an explicit nil

### UnsetLine
`func (o *Finding) UnsetLine()`

UnsetLine ensures that no value is present for Line, not even an explicit nil
### GetFilePath

`func (o *Finding) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *Finding) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *Finding) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *Finding) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *Finding) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *Finding) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetComponentName

`func (o *Finding) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *Finding) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *Finding) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *Finding) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### SetComponentNameNil

`func (o *Finding) SetComponentNameNil(b bool)`

 SetComponentNameNil sets the value for ComponentName to be an explicit nil

### UnsetComponentName
`func (o *Finding) UnsetComponentName()`

UnsetComponentName ensures that no value is present for ComponentName, not even an explicit nil
### GetComponentVersion

`func (o *Finding) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *Finding) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *Finding) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *Finding) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *Finding) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *Finding) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetStaticFinding

`func (o *Finding) GetStaticFinding() bool`

GetStaticFinding returns the StaticFinding field if non-nil, zero value otherwise.

### GetStaticFindingOk

`func (o *Finding) GetStaticFindingOk() (*bool, bool)`

GetStaticFindingOk returns a tuple with the StaticFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticFinding

`func (o *Finding) SetStaticFinding(v bool)`

SetStaticFinding sets StaticFinding field to given value.

### HasStaticFinding

`func (o *Finding) HasStaticFinding() bool`

HasStaticFinding returns a boolean if a field has been set.

### GetDynamicFinding

`func (o *Finding) GetDynamicFinding() bool`

GetDynamicFinding returns the DynamicFinding field if non-nil, zero value otherwise.

### GetDynamicFindingOk

`func (o *Finding) GetDynamicFindingOk() (*bool, bool)`

GetDynamicFindingOk returns a tuple with the DynamicFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicFinding

`func (o *Finding) SetDynamicFinding(v bool)`

SetDynamicFinding sets DynamicFinding field to given value.

### HasDynamicFinding

`func (o *Finding) HasDynamicFinding() bool`

HasDynamicFinding returns a boolean if a field has been set.

### GetCreated

`func (o *Finding) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Finding) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Finding) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Finding) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Finding) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetScannerConfidence

`func (o *Finding) GetScannerConfidence() int32`

GetScannerConfidence returns the ScannerConfidence field if non-nil, zero value otherwise.

### GetScannerConfidenceOk

`func (o *Finding) GetScannerConfidenceOk() (*int32, bool)`

GetScannerConfidenceOk returns a tuple with the ScannerConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannerConfidence

`func (o *Finding) SetScannerConfidence(v int32)`

SetScannerConfidence sets ScannerConfidence field to given value.


### SetScannerConfidenceNil

`func (o *Finding) SetScannerConfidenceNil(b bool)`

 SetScannerConfidenceNil sets the value for ScannerConfidence to be an explicit nil

### UnsetScannerConfidence
`func (o *Finding) UnsetScannerConfidence()`

UnsetScannerConfidence ensures that no value is present for ScannerConfidence, not even an explicit nil
### GetUniqueIdFromTool

`func (o *Finding) GetUniqueIdFromTool() string`

GetUniqueIdFromTool returns the UniqueIdFromTool field if non-nil, zero value otherwise.

### GetUniqueIdFromToolOk

`func (o *Finding) GetUniqueIdFromToolOk() (*string, bool)`

GetUniqueIdFromToolOk returns a tuple with the UniqueIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdFromTool

`func (o *Finding) SetUniqueIdFromTool(v string)`

SetUniqueIdFromTool sets UniqueIdFromTool field to given value.

### HasUniqueIdFromTool

`func (o *Finding) HasUniqueIdFromTool() bool`

HasUniqueIdFromTool returns a boolean if a field has been set.

### SetUniqueIdFromToolNil

`func (o *Finding) SetUniqueIdFromToolNil(b bool)`

 SetUniqueIdFromToolNil sets the value for UniqueIdFromTool to be an explicit nil

### UnsetUniqueIdFromTool
`func (o *Finding) UnsetUniqueIdFromTool()`

UnsetUniqueIdFromTool ensures that no value is present for UniqueIdFromTool, not even an explicit nil
### GetVulnIdFromTool

`func (o *Finding) GetVulnIdFromTool() string`

GetVulnIdFromTool returns the VulnIdFromTool field if non-nil, zero value otherwise.

### GetVulnIdFromToolOk

`func (o *Finding) GetVulnIdFromToolOk() (*string, bool)`

GetVulnIdFromToolOk returns a tuple with the VulnIdFromTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnIdFromTool

`func (o *Finding) SetVulnIdFromTool(v string)`

SetVulnIdFromTool sets VulnIdFromTool field to given value.

### HasVulnIdFromTool

`func (o *Finding) HasVulnIdFromTool() bool`

HasVulnIdFromTool returns a boolean if a field has been set.

### SetVulnIdFromToolNil

`func (o *Finding) SetVulnIdFromToolNil(b bool)`

 SetVulnIdFromToolNil sets the value for VulnIdFromTool to be an explicit nil

### UnsetVulnIdFromTool
`func (o *Finding) UnsetVulnIdFromTool()`

UnsetVulnIdFromTool ensures that no value is present for VulnIdFromTool, not even an explicit nil
### GetSastSourceObject

`func (o *Finding) GetSastSourceObject() string`

GetSastSourceObject returns the SastSourceObject field if non-nil, zero value otherwise.

### GetSastSourceObjectOk

`func (o *Finding) GetSastSourceObjectOk() (*string, bool)`

GetSastSourceObjectOk returns a tuple with the SastSourceObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceObject

`func (o *Finding) SetSastSourceObject(v string)`

SetSastSourceObject sets SastSourceObject field to given value.

### HasSastSourceObject

`func (o *Finding) HasSastSourceObject() bool`

HasSastSourceObject returns a boolean if a field has been set.

### SetSastSourceObjectNil

`func (o *Finding) SetSastSourceObjectNil(b bool)`

 SetSastSourceObjectNil sets the value for SastSourceObject to be an explicit nil

### UnsetSastSourceObject
`func (o *Finding) UnsetSastSourceObject()`

UnsetSastSourceObject ensures that no value is present for SastSourceObject, not even an explicit nil
### GetSastSinkObject

`func (o *Finding) GetSastSinkObject() string`

GetSastSinkObject returns the SastSinkObject field if non-nil, zero value otherwise.

### GetSastSinkObjectOk

`func (o *Finding) GetSastSinkObjectOk() (*string, bool)`

GetSastSinkObjectOk returns a tuple with the SastSinkObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSinkObject

`func (o *Finding) SetSastSinkObject(v string)`

SetSastSinkObject sets SastSinkObject field to given value.

### HasSastSinkObject

`func (o *Finding) HasSastSinkObject() bool`

HasSastSinkObject returns a boolean if a field has been set.

### SetSastSinkObjectNil

`func (o *Finding) SetSastSinkObjectNil(b bool)`

 SetSastSinkObjectNil sets the value for SastSinkObject to be an explicit nil

### UnsetSastSinkObject
`func (o *Finding) UnsetSastSinkObject()`

UnsetSastSinkObject ensures that no value is present for SastSinkObject, not even an explicit nil
### GetSastSourceLine

`func (o *Finding) GetSastSourceLine() int32`

GetSastSourceLine returns the SastSourceLine field if non-nil, zero value otherwise.

### GetSastSourceLineOk

`func (o *Finding) GetSastSourceLineOk() (*int32, bool)`

GetSastSourceLineOk returns a tuple with the SastSourceLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceLine

`func (o *Finding) SetSastSourceLine(v int32)`

SetSastSourceLine sets SastSourceLine field to given value.

### HasSastSourceLine

`func (o *Finding) HasSastSourceLine() bool`

HasSastSourceLine returns a boolean if a field has been set.

### SetSastSourceLineNil

`func (o *Finding) SetSastSourceLineNil(b bool)`

 SetSastSourceLineNil sets the value for SastSourceLine to be an explicit nil

### UnsetSastSourceLine
`func (o *Finding) UnsetSastSourceLine()`

UnsetSastSourceLine ensures that no value is present for SastSourceLine, not even an explicit nil
### GetSastSourceFilePath

`func (o *Finding) GetSastSourceFilePath() string`

GetSastSourceFilePath returns the SastSourceFilePath field if non-nil, zero value otherwise.

### GetSastSourceFilePathOk

`func (o *Finding) GetSastSourceFilePathOk() (*string, bool)`

GetSastSourceFilePathOk returns a tuple with the SastSourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastSourceFilePath

`func (o *Finding) SetSastSourceFilePath(v string)`

SetSastSourceFilePath sets SastSourceFilePath field to given value.

### HasSastSourceFilePath

`func (o *Finding) HasSastSourceFilePath() bool`

HasSastSourceFilePath returns a boolean if a field has been set.

### SetSastSourceFilePathNil

`func (o *Finding) SetSastSourceFilePathNil(b bool)`

 SetSastSourceFilePathNil sets the value for SastSourceFilePath to be an explicit nil

### UnsetSastSourceFilePath
`func (o *Finding) UnsetSastSourceFilePath()`

UnsetSastSourceFilePath ensures that no value is present for SastSourceFilePath, not even an explicit nil
### GetNbOccurences

`func (o *Finding) GetNbOccurences() int32`

GetNbOccurences returns the NbOccurences field if non-nil, zero value otherwise.

### GetNbOccurencesOk

`func (o *Finding) GetNbOccurencesOk() (*int32, bool)`

GetNbOccurencesOk returns a tuple with the NbOccurences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOccurences

`func (o *Finding) SetNbOccurences(v int32)`

SetNbOccurences sets NbOccurences field to given value.

### HasNbOccurences

`func (o *Finding) HasNbOccurences() bool`

HasNbOccurences returns a boolean if a field has been set.

### SetNbOccurencesNil

`func (o *Finding) SetNbOccurencesNil(b bool)`

 SetNbOccurencesNil sets the value for NbOccurences to be an explicit nil

### UnsetNbOccurences
`func (o *Finding) UnsetNbOccurences()`

UnsetNbOccurences ensures that no value is present for NbOccurences, not even an explicit nil
### GetPublishDate

`func (o *Finding) GetPublishDate() string`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *Finding) GetPublishDateOk() (*string, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *Finding) SetPublishDate(v string)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *Finding) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### SetPublishDateNil

`func (o *Finding) SetPublishDateNil(b bool)`

 SetPublishDateNil sets the value for PublishDate to be an explicit nil

### UnsetPublishDate
`func (o *Finding) UnsetPublishDate()`

UnsetPublishDate ensures that no value is present for PublishDate, not even an explicit nil
### GetService

`func (o *Finding) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Finding) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Finding) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *Finding) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *Finding) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *Finding) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetPlannedRemediationDate

`func (o *Finding) GetPlannedRemediationDate() string`

GetPlannedRemediationDate returns the PlannedRemediationDate field if non-nil, zero value otherwise.

### GetPlannedRemediationDateOk

`func (o *Finding) GetPlannedRemediationDateOk() (*string, bool)`

GetPlannedRemediationDateOk returns a tuple with the PlannedRemediationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationDate

`func (o *Finding) SetPlannedRemediationDate(v string)`

SetPlannedRemediationDate sets PlannedRemediationDate field to given value.

### HasPlannedRemediationDate

`func (o *Finding) HasPlannedRemediationDate() bool`

HasPlannedRemediationDate returns a boolean if a field has been set.

### SetPlannedRemediationDateNil

`func (o *Finding) SetPlannedRemediationDateNil(b bool)`

 SetPlannedRemediationDateNil sets the value for PlannedRemediationDate to be an explicit nil

### UnsetPlannedRemediationDate
`func (o *Finding) UnsetPlannedRemediationDate()`

UnsetPlannedRemediationDate ensures that no value is present for PlannedRemediationDate, not even an explicit nil
### GetPlannedRemediationVersion

`func (o *Finding) GetPlannedRemediationVersion() string`

GetPlannedRemediationVersion returns the PlannedRemediationVersion field if non-nil, zero value otherwise.

### GetPlannedRemediationVersionOk

`func (o *Finding) GetPlannedRemediationVersionOk() (*string, bool)`

GetPlannedRemediationVersionOk returns a tuple with the PlannedRemediationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedRemediationVersion

`func (o *Finding) SetPlannedRemediationVersion(v string)`

SetPlannedRemediationVersion sets PlannedRemediationVersion field to given value.

### HasPlannedRemediationVersion

`func (o *Finding) HasPlannedRemediationVersion() bool`

HasPlannedRemediationVersion returns a boolean if a field has been set.

### SetPlannedRemediationVersionNil

`func (o *Finding) SetPlannedRemediationVersionNil(b bool)`

 SetPlannedRemediationVersionNil sets the value for PlannedRemediationVersion to be an explicit nil

### UnsetPlannedRemediationVersion
`func (o *Finding) UnsetPlannedRemediationVersion()`

UnsetPlannedRemediationVersion ensures that no value is present for PlannedRemediationVersion, not even an explicit nil
### GetEffortForFixing

`func (o *Finding) GetEffortForFixing() string`

GetEffortForFixing returns the EffortForFixing field if non-nil, zero value otherwise.

### GetEffortForFixingOk

`func (o *Finding) GetEffortForFixingOk() (*string, bool)`

GetEffortForFixingOk returns a tuple with the EffortForFixing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffortForFixing

`func (o *Finding) SetEffortForFixing(v string)`

SetEffortForFixing sets EffortForFixing field to given value.

### HasEffortForFixing

`func (o *Finding) HasEffortForFixing() bool`

HasEffortForFixing returns a boolean if a field has been set.

### SetEffortForFixingNil

`func (o *Finding) SetEffortForFixingNil(b bool)`

 SetEffortForFixingNil sets the value for EffortForFixing to be an explicit nil

### UnsetEffortForFixing
`func (o *Finding) UnsetEffortForFixing()`

UnsetEffortForFixing ensures that no value is present for EffortForFixing, not even an explicit nil
### GetTest

`func (o *Finding) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *Finding) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *Finding) SetTest(v int32)`

SetTest sets Test field to given value.


### GetDuplicateFinding

`func (o *Finding) GetDuplicateFinding() int32`

GetDuplicateFinding returns the DuplicateFinding field if non-nil, zero value otherwise.

### GetDuplicateFindingOk

`func (o *Finding) GetDuplicateFindingOk() (*int32, bool)`

GetDuplicateFindingOk returns a tuple with the DuplicateFinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateFinding

`func (o *Finding) SetDuplicateFinding(v int32)`

SetDuplicateFinding sets DuplicateFinding field to given value.


### SetDuplicateFindingNil

`func (o *Finding) SetDuplicateFindingNil(b bool)`

 SetDuplicateFindingNil sets the value for DuplicateFinding to be an explicit nil

### UnsetDuplicateFinding
`func (o *Finding) UnsetDuplicateFinding()`

UnsetDuplicateFinding ensures that no value is present for DuplicateFinding, not even an explicit nil
### GetReviewRequestedBy

`func (o *Finding) GetReviewRequestedBy() int32`

GetReviewRequestedBy returns the ReviewRequestedBy field if non-nil, zero value otherwise.

### GetReviewRequestedByOk

`func (o *Finding) GetReviewRequestedByOk() (*int32, bool)`

GetReviewRequestedByOk returns a tuple with the ReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequestedBy

`func (o *Finding) SetReviewRequestedBy(v int32)`

SetReviewRequestedBy sets ReviewRequestedBy field to given value.

### HasReviewRequestedBy

`func (o *Finding) HasReviewRequestedBy() bool`

HasReviewRequestedBy returns a boolean if a field has been set.

### SetReviewRequestedByNil

`func (o *Finding) SetReviewRequestedByNil(b bool)`

 SetReviewRequestedByNil sets the value for ReviewRequestedBy to be an explicit nil

### UnsetReviewRequestedBy
`func (o *Finding) UnsetReviewRequestedBy()`

UnsetReviewRequestedBy ensures that no value is present for ReviewRequestedBy, not even an explicit nil
### GetDefectReviewRequestedBy

`func (o *Finding) GetDefectReviewRequestedBy() int32`

GetDefectReviewRequestedBy returns the DefectReviewRequestedBy field if non-nil, zero value otherwise.

### GetDefectReviewRequestedByOk

`func (o *Finding) GetDefectReviewRequestedByOk() (*int32, bool)`

GetDefectReviewRequestedByOk returns a tuple with the DefectReviewRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectReviewRequestedBy

`func (o *Finding) SetDefectReviewRequestedBy(v int32)`

SetDefectReviewRequestedBy sets DefectReviewRequestedBy field to given value.

### HasDefectReviewRequestedBy

`func (o *Finding) HasDefectReviewRequestedBy() bool`

HasDefectReviewRequestedBy returns a boolean if a field has been set.

### SetDefectReviewRequestedByNil

`func (o *Finding) SetDefectReviewRequestedByNil(b bool)`

 SetDefectReviewRequestedByNil sets the value for DefectReviewRequestedBy to be an explicit nil

### UnsetDefectReviewRequestedBy
`func (o *Finding) UnsetDefectReviewRequestedBy()`

UnsetDefectReviewRequestedBy ensures that no value is present for DefectReviewRequestedBy, not even an explicit nil
### GetMitigatedBy

`func (o *Finding) GetMitigatedBy() int32`

GetMitigatedBy returns the MitigatedBy field if non-nil, zero value otherwise.

### GetMitigatedByOk

`func (o *Finding) GetMitigatedByOk() (*int32, bool)`

GetMitigatedByOk returns a tuple with the MitigatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedBy

`func (o *Finding) SetMitigatedBy(v int32)`

SetMitigatedBy sets MitigatedBy field to given value.


### SetMitigatedByNil

`func (o *Finding) SetMitigatedByNil(b bool)`

 SetMitigatedByNil sets the value for MitigatedBy to be an explicit nil

### UnsetMitigatedBy
`func (o *Finding) UnsetMitigatedBy()`

UnsetMitigatedBy ensures that no value is present for MitigatedBy, not even an explicit nil
### GetLastReviewedBy

`func (o *Finding) GetLastReviewedBy() int32`

GetLastReviewedBy returns the LastReviewedBy field if non-nil, zero value otherwise.

### GetLastReviewedByOk

`func (o *Finding) GetLastReviewedByOk() (*int32, bool)`

GetLastReviewedByOk returns a tuple with the LastReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewedBy

`func (o *Finding) SetLastReviewedBy(v int32)`

SetLastReviewedBy sets LastReviewedBy field to given value.


### SetLastReviewedByNil

`func (o *Finding) SetLastReviewedByNil(b bool)`

 SetLastReviewedByNil sets the value for LastReviewedBy to be an explicit nil

### UnsetLastReviewedBy
`func (o *Finding) UnsetLastReviewedBy()`

UnsetLastReviewedBy ensures that no value is present for LastReviewedBy, not even an explicit nil
### GetSonarqubeIssue

`func (o *Finding) GetSonarqubeIssue() int32`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *Finding) GetSonarqubeIssueOk() (*int32, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *Finding) SetSonarqubeIssue(v int32)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.

### HasSonarqubeIssue

`func (o *Finding) HasSonarqubeIssue() bool`

HasSonarqubeIssue returns a boolean if a field has been set.

### SetSonarqubeIssueNil

`func (o *Finding) SetSonarqubeIssueNil(b bool)`

 SetSonarqubeIssueNil sets the value for SonarqubeIssue to be an explicit nil

### UnsetSonarqubeIssue
`func (o *Finding) UnsetSonarqubeIssue()`

UnsetSonarqubeIssue ensures that no value is present for SonarqubeIssue, not even an explicit nil
### GetEndpoints

`func (o *Finding) GetEndpoints() []int32`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *Finding) GetEndpointsOk() (*[]int32, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *Finding) SetEndpoints(v []int32)`

SetEndpoints sets Endpoints field to given value.


### GetReviewers

`func (o *Finding) GetReviewers() []int32`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *Finding) GetReviewersOk() (*[]int32, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *Finding) SetReviewers(v []int32)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *Finding) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetNotes

`func (o *Finding) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Finding) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Finding) SetNotes(v []Note)`

SetNotes sets Notes field to given value.


### GetFiles

`func (o *Finding) GetFiles() []int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Finding) GetFilesOk() (*[]int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Finding) SetFiles(v []int32)`

SetFiles sets Files field to given value.


### GetFoundBy

`func (o *Finding) GetFoundBy() []int32`

GetFoundBy returns the FoundBy field if non-nil, zero value otherwise.

### GetFoundByOk

`func (o *Finding) GetFoundByOk() (*[]int32, bool)`

GetFoundByOk returns a tuple with the FoundBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundBy

`func (o *Finding) SetFoundBy(v []int32)`

SetFoundBy sets FoundBy field to given value.


### GetPrefetch

`func (o *Finding) GetPrefetch() FindingPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *Finding) GetPrefetchOk() (*FindingPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *Finding) SetPrefetch(v FindingPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *Finding) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


