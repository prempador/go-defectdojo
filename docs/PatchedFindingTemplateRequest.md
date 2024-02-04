# PatchedFindingTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**VulnerabilityIds** | Pointer to [**[]VulnerabilityIdTemplateRequest**](VulnerabilityIdTemplateRequest.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
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

### NewPatchedFindingTemplateRequest

`func NewPatchedFindingTemplateRequest() *PatchedFindingTemplateRequest`

NewPatchedFindingTemplateRequest instantiates a new PatchedFindingTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedFindingTemplateRequestWithDefaults

`func NewPatchedFindingTemplateRequestWithDefaults() *PatchedFindingTemplateRequest`

NewPatchedFindingTemplateRequestWithDefaults instantiates a new PatchedFindingTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedFindingTemplateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedFindingTemplateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedFindingTemplateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedFindingTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVulnerabilityIds

`func (o *PatchedFindingTemplateRequest) GetVulnerabilityIds() []VulnerabilityIdTemplateRequest`

GetVulnerabilityIds returns the VulnerabilityIds field if non-nil, zero value otherwise.

### GetVulnerabilityIdsOk

`func (o *PatchedFindingTemplateRequest) GetVulnerabilityIdsOk() (*[]VulnerabilityIdTemplateRequest, bool)`

GetVulnerabilityIdsOk returns a tuple with the VulnerabilityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIds

`func (o *PatchedFindingTemplateRequest) SetVulnerabilityIds(v []VulnerabilityIdTemplateRequest)`

SetVulnerabilityIds sets VulnerabilityIds field to given value.

### HasVulnerabilityIds

`func (o *PatchedFindingTemplateRequest) HasVulnerabilityIds() bool`

HasVulnerabilityIds returns a boolean if a field has been set.

### GetTitle

`func (o *PatchedFindingTemplateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatchedFindingTemplateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatchedFindingTemplateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatchedFindingTemplateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCwe

`func (o *PatchedFindingTemplateRequest) GetCwe() int32`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *PatchedFindingTemplateRequest) GetCweOk() (*int32, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *PatchedFindingTemplateRequest) SetCwe(v int32)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *PatchedFindingTemplateRequest) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### SetCweNil

`func (o *PatchedFindingTemplateRequest) SetCweNil(b bool)`

 SetCweNil sets the value for Cwe to be an explicit nil

### UnsetCwe
`func (o *PatchedFindingTemplateRequest) UnsetCwe()`

UnsetCwe ensures that no value is present for Cwe, not even an explicit nil
### GetCvssv3

`func (o *PatchedFindingTemplateRequest) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *PatchedFindingTemplateRequest) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *PatchedFindingTemplateRequest) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *PatchedFindingTemplateRequest) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### SetCvssv3Nil

`func (o *PatchedFindingTemplateRequest) SetCvssv3Nil(b bool)`

 SetCvssv3Nil sets the value for Cvssv3 to be an explicit nil

### UnsetCvssv3
`func (o *PatchedFindingTemplateRequest) UnsetCvssv3()`

UnsetCvssv3 ensures that no value is present for Cvssv3, not even an explicit nil
### GetSeverity

`func (o *PatchedFindingTemplateRequest) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PatchedFindingTemplateRequest) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PatchedFindingTemplateRequest) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *PatchedFindingTemplateRequest) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *PatchedFindingTemplateRequest) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *PatchedFindingTemplateRequest) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetDescription

`func (o *PatchedFindingTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedFindingTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedFindingTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedFindingTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedFindingTemplateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedFindingTemplateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMitigation

`func (o *PatchedFindingTemplateRequest) GetMitigation() string`

GetMitigation returns the Mitigation field if non-nil, zero value otherwise.

### GetMitigationOk

`func (o *PatchedFindingTemplateRequest) GetMitigationOk() (*string, bool)`

GetMitigationOk returns a tuple with the Mitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigation

`func (o *PatchedFindingTemplateRequest) SetMitigation(v string)`

SetMitigation sets Mitigation field to given value.

### HasMitigation

`func (o *PatchedFindingTemplateRequest) HasMitigation() bool`

HasMitigation returns a boolean if a field has been set.

### SetMitigationNil

`func (o *PatchedFindingTemplateRequest) SetMitigationNil(b bool)`

 SetMitigationNil sets the value for Mitigation to be an explicit nil

### UnsetMitigation
`func (o *PatchedFindingTemplateRequest) UnsetMitigation()`

UnsetMitigation ensures that no value is present for Mitigation, not even an explicit nil
### GetImpact

`func (o *PatchedFindingTemplateRequest) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *PatchedFindingTemplateRequest) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *PatchedFindingTemplateRequest) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *PatchedFindingTemplateRequest) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### SetImpactNil

`func (o *PatchedFindingTemplateRequest) SetImpactNil(b bool)`

 SetImpactNil sets the value for Impact to be an explicit nil

### UnsetImpact
`func (o *PatchedFindingTemplateRequest) UnsetImpact()`

UnsetImpact ensures that no value is present for Impact, not even an explicit nil
### GetReferences

`func (o *PatchedFindingTemplateRequest) GetReferences() string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PatchedFindingTemplateRequest) GetReferencesOk() (*string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PatchedFindingTemplateRequest) SetReferences(v string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PatchedFindingTemplateRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *PatchedFindingTemplateRequest) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *PatchedFindingTemplateRequest) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil
### GetTemplateMatch

`func (o *PatchedFindingTemplateRequest) GetTemplateMatch() bool`

GetTemplateMatch returns the TemplateMatch field if non-nil, zero value otherwise.

### GetTemplateMatchOk

`func (o *PatchedFindingTemplateRequest) GetTemplateMatchOk() (*bool, bool)`

GetTemplateMatchOk returns a tuple with the TemplateMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMatch

`func (o *PatchedFindingTemplateRequest) SetTemplateMatch(v bool)`

SetTemplateMatch sets TemplateMatch field to given value.

### HasTemplateMatch

`func (o *PatchedFindingTemplateRequest) HasTemplateMatch() bool`

HasTemplateMatch returns a boolean if a field has been set.

### GetTemplateMatchTitle

`func (o *PatchedFindingTemplateRequest) GetTemplateMatchTitle() bool`

GetTemplateMatchTitle returns the TemplateMatchTitle field if non-nil, zero value otherwise.

### GetTemplateMatchTitleOk

`func (o *PatchedFindingTemplateRequest) GetTemplateMatchTitleOk() (*bool, bool)`

GetTemplateMatchTitleOk returns a tuple with the TemplateMatchTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMatchTitle

`func (o *PatchedFindingTemplateRequest) SetTemplateMatchTitle(v bool)`

SetTemplateMatchTitle sets TemplateMatchTitle field to given value.

### HasTemplateMatchTitle

`func (o *PatchedFindingTemplateRequest) HasTemplateMatchTitle() bool`

HasTemplateMatchTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


