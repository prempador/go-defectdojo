# FindingTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**VulnerabilityIds** | Pointer to [**[]VulnerabilityIdTemplate**](VulnerabilityIdTemplate.md) |  | [optional] 
**Title** | **string** |  | 
**Cwe** | Pointer to **NullableInt32** |  | [optional] 
**Cvssv3** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Mitigation** | Pointer to **NullableString** |  | [optional] 
**Impact** | Pointer to **NullableString** |  | [optional] 
**References** | Pointer to **NullableString** |  | [optional] 
**LastUsed** | **NullableTime** |  | [readonly] 
**NumericalSeverity** | **NullableString** |  | [readonly] 
**TemplateMatch** | Pointer to **bool** | Enables this template for matching remediation advice. Match will be applied to all active, verified findings by CWE. | [optional] 
**TemplateMatchTitle** | Pointer to **bool** | Matches by title text (contains search) and CWE. | [optional] 

## Methods

### NewFindingTemplate

`func NewFindingTemplate(id int32, title string, lastUsed NullableTime, numericalSeverity NullableString, ) *FindingTemplate`

NewFindingTemplate instantiates a new FindingTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingTemplateWithDefaults

`func NewFindingTemplateWithDefaults() *FindingTemplate`

NewFindingTemplateWithDefaults instantiates a new FindingTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FindingTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *FindingTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindingTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindingTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindingTemplate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVulnerabilityIds

`func (o *FindingTemplate) GetVulnerabilityIds() []VulnerabilityIdTemplate`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *FindingTemplate) GetVulnerabilityIdsOk() (*[]VulnerabilityIdTemplate, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *FindingTemplate) SetVulnerabilityIds(v []VulnerabilityIdTemplate)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *FindingTemplate) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetTitle

`func (o *FindingTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCwe

`func (o *FindingTemplate) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *FindingTemplate) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *FindingTemplate) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *FindingTemplate) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *FindingTemplate) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *FindingTemplate) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetCvssv3

`func (o *FindingTemplate) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *FindingTemplate) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *FindingTemplate) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *FindingTemplate) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *FindingTemplate) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *FindingTemplate) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetSeverity

`func (o *FindingTemplate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingTemplate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingTemplate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FindingTemplate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *FindingTemplate) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *FindingTemplate) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetDescription

`func (o *FindingTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindingTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FindingTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FindingTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMitigation

`func (o *FindingTemplate) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *FindingTemplate) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *FindingTemplate) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *FindingTemplate) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *FindingTemplate) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *FindingTemplate) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *FindingTemplate) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *FindingTemplate) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *FindingTemplate) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *FindingTemplate) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *FindingTemplate) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *FindingTemplate) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetReferences

`func (o *FindingTemplate) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FindingTemplate) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FindingTemplate) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *FindingTemplate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *FindingTemplate) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *FindingTemplate) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetLastUsed

`func (o *FindingTemplate) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *FindingTemplate) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *FindingTemplate) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.


### SetLastUsedNil

`func (o *FindingTemplate) SetLastUsedNil(b bool)`

 SetLastUsedNil sets the value for LastUsed to be an explicit nil

### UnsetLastUsed
`func (o *FindingTemplate) UnsetLastUsed()`

UnsetLastUsed ensures that no value is present for LastUsed, not even an explicit nil
### GetNumericalSeverity

`func (o *FindingTemplate) GetNumericalSeverity() string`

GetNumericalSeverity returns the NumericalSeverity field if non-nil, zero value otherwise.

### GetNumericalSeverityOk

`func (o *FindingTemplate) GetNumericalSeverityOk() (*string, bool)`

GetNumericalSeverityOk returns a tuple with the NumericalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericalSeverity

`func (o *FindingTemplate) SetNumericalSeverity(v string)`

SetNumericalSeverity sets NumericalSeverity field to given value.


### SetNumericalSeverityNil

`func (o *FindingTemplate) SetNumericalSeverityNil(b bool)`

 SetNumericalSeverityNil sets the value for NumericalSeverity to be an explicit nil

### UnsetNumericalSeverity
`func (o *FindingTemplate) UnsetNumericalSeverity()`

UnsetNumericalSeverity ensures that no value is present for NumericalSeverity, not even an explicit nil
### GetTemplateMatch

`func (o *FindingTemplate) GetTemplateMatch() bool`

GetTemplateMatch returns the TemplateMatch field if non-nil, zero value otherwise.

### GetTemplateMatchOk

`func (o *FindingTemplate) GetTemplateMatchOk() (*bool, bool)`

GetTemplateMatchOk returns a tuple with the TemplateMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMatch

`func (o *FindingTemplate) SetTemplateMatch(v bool)`

SetTemplateMatch sets TemplateMatch field to given value.

### HasTemplateMatch

`func (o *FindingTemplate) HasTemplateMatch() bool`

HasTemplateMatch returns a boolean if a field has been set.

### GetTemplateMatchTitle

`func (o *FindingTemplate) GetTemplateMatchTitle() bool`

GetTemplateMatchTitle returns the TemplateMatchTitle field if non-nil, zero value otherwise.

### GetTemplateMatchTitleOk

`func (o *FindingTemplate) GetTemplateMatchTitleOk() (*bool, bool)`

GetTemplateMatchTitleOk returns a tuple with the TemplateMatchTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMatchTitle

`func (o *FindingTemplate) SetTemplateMatchTitle(v bool)`

SetTemplateMatchTitle sets TemplateMatchTitle field to given value.

### HasTemplateMatchTitle

`func (o *FindingTemplate) HasTemplateMatchTitle() bool`

HasTemplateMatchTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


