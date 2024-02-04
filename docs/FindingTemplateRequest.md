# FindingTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**VulnerabilityIds** | Pointer to [**[]VulnerabilityIdTemplateRequest**](VulnerabilityIdTemplateRequest.md) |  | [optional] 
**Title** | **string** |  | 
**Cwe** | Pointer to **NullableInt32** |  | [optional] 
**Cvssv3** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Mitigation** | Pointer to **NullableString** |  | [optional] 
**Impact** | Pointer to **NullableString** |  | [optional] 
**References** | Pointer to **NullableString** |  | [optional] 
**TemplateMatch** | Pointer to **bool** | Enables this template for matching remediation advice. Match will be applied to all active, verified findings by CWE. | [optional] 
**TemplateMatchTitle** | Pointer to **bool** | Matches by title text (contains search) and CWE. | [optional] 

## Methods

### NewFindingTemplateRequest

`func NewFindingTemplateRequest(title string, ) *FindingTemplateRequest`

NewFindingTemplateRequest instantiates a new FindingTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingTemplateRequestWithDefaults

`func NewFindingTemplateRequestWithDefaults() *FindingTemplateRequest`

NewFindingTemplateRequestWithDefaults instantiates a new FindingTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *FindingTemplateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FindingTemplateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FindingTemplateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FindingTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVulnerabilityIds

`func (o *FindingTemplateRequest) GetVulnerabilityIds() []VulnerabilityIdTemplateRequest`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *FindingTemplateRequest) GetVulnerabilityIdsOk() (*[]VulnerabilityIdTemplateRequest, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *FindingTemplateRequest) SetVulnerabilityIds(v []VulnerabilityIdTemplateRequest)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *FindingTemplateRequest) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetTitle

`func (o *FindingTemplateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingTemplateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingTemplateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetCwe

`func (o *FindingTemplateRequest) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *FindingTemplateRequest) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *FindingTemplateRequest) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *FindingTemplateRequest) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *FindingTemplateRequest) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *FindingTemplateRequest) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetCvssv3

`func (o *FindingTemplateRequest) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *FindingTemplateRequest) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *FindingTemplateRequest) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *FindingTemplateRequest) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *FindingTemplateRequest) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *FindingTemplateRequest) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetSeverity

`func (o *FindingTemplateRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *FindingTemplateRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *FindingTemplateRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *FindingTemplateRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *FindingTemplateRequest) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *FindingTemplateRequest) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetDescription

`func (o *FindingTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindingTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FindingTemplateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FindingTemplateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMitigation

`func (o *FindingTemplateRequest) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *FindingTemplateRequest) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *FindingTemplateRequest) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *FindingTemplateRequest) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *FindingTemplateRequest) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *FindingTemplateRequest) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *FindingTemplateRequest) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *FindingTemplateRequest) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *FindingTemplateRequest) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *FindingTemplateRequest) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *FindingTemplateRequest) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *FindingTemplateRequest) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetReferences

`func (o *FindingTemplateRequest) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *FindingTemplateRequest) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *FindingTemplateRequest) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *FindingTemplateRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *FindingTemplateRequest) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *FindingTemplateRequest) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetTemplateMatch

`func (o *FindingTemplateRequest) GetTemplateMatch() bool`

GetTemplateMatch returns the TemplateMatch field if non-nil, zero value otherwise.

### GetTemplateMatchOk

`func (o *FindingTemplateRequest) GetTemplateMatchOk() (*bool, bool)`

GetTemplateMatchOk returns a tuple with the TemplateMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMatch

`func (o *FindingTemplateRequest) SetTemplateMatch(v bool)`

SetTemplateMatch sets TemplateMatch field to given value.

### HasTemplateMatch

`func (o *FindingTemplateRequest) HasTemplateMatch() bool`

HasTemplateMatch returns a boolean if a field has been set.

### GetTemplateMatchTitle

`func (o *FindingTemplateRequest) GetTemplateMatchTitle() bool`

GetTemplateMatchTitle returns the TemplateMatchTitle field if non-nil, zero value otherwise.

### GetTemplateMatchTitleOk

`func (o *FindingTemplateRequest) GetTemplateMatchTitleOk() (*bool, bool)`

GetTemplateMatchTitleOk returns a tuple with the TemplateMatchTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMatchTitle

`func (o *FindingTemplateRequest) SetTemplateMatchTitle(v bool)`

SetTemplateMatchTitle sets TemplateMatchTitle field to given value.

### HasTemplateMatchTitle

`func (o *FindingTemplateRequest) HasTemplateMatchTitle() bool`

HasTemplateMatchTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


