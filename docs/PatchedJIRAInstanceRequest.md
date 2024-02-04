# PatchedJIRAInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationName** | Pointer to **string** | Enter a name to give to this configuration | [optional] 
**Url** | Pointer to **string** | For more information how to configure Jira, read the DefectDojo documentation. | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**DefaultIssueType** | Pointer to **string** | You can define extra issue types in settings.py  * &#x60;Task&#x60; - Task * &#x60;Story&#x60; - Story * &#x60;Epic&#x60; - Epic * &#x60;Spike&#x60; - Spike * &#x60;Bug&#x60; - Bug * &#x60;Security&#x60; - Security | [optional] 
**IssueTemplateDir** | Pointer to **NullableString** | Choose the folder containing the Django templates used to render the JIRA issue description. These are stored in dojo/templates/issue-trackers. Leave empty to use the default jira_full templates. | [optional] 
**EpicNameId** | Pointer to **int32** | To obtain the &#39;Epic name id&#39; visit https://&lt;YOUR JIRA URL&gt;/rest/api/2/field and search for Epic Name. Copy the number out of cf[number] and paste it here. | [optional] 
**OpenStatusKey** | Pointer to **int32** | Transition ID to Re-Open JIRA issues, visit https://&lt;YOUR JIRA URL&gt;/rest/api/latest/issue/&lt;ANY VALID ISSUE KEY&gt;/transitions?expand&#x3D;transitions.fields to find the ID for your JIRA instance | [optional] 
**CloseStatusKey** | Pointer to **int32** | Transition ID to Close JIRA issues, visit https://&lt;YOUR JIRA URL&gt;/rest/api/latest/issue/&lt;ANY VALID ISSUE KEY&gt;/transitions?expand&#x3D;transitions.fields to find the ID for your JIRA instance | [optional] 
**InfoMappingSeverity** | Pointer to **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Info | [optional] 
**LowMappingSeverity** | Pointer to **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Low | [optional] 
**MediumMappingSeverity** | Pointer to **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Medium | [optional] 
**HighMappingSeverity** | Pointer to **string** | Maps to the &#39;Priority&#39; field in Jira. For example: High | [optional] 
**CriticalMappingSeverity** | Pointer to **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Critical | [optional] 
**FindingText** | Pointer to **NullableString** | Additional text that will be added to the finding in Jira. For example including how the finding was created or who to contact for more information. | [optional] 
**AcceptedMappingResolution** | Pointer to **NullableString** | JIRA resolution names (comma-separated values) that maps to an Accepted Finding | [optional] 
**FalsePositiveMappingResolution** | Pointer to **NullableString** | JIRA resolution names (comma-separated values) that maps to a False Positive Finding | [optional] 
**GlobalJiraSlaNotification** | Pointer to **bool** | This setting can be overidden at the Product level | [optional] 
**FindingJiraSync** | Pointer to **bool** | If enabled, this will sync changes to a Finding automatically to JIRA | [optional] 

## Methods

### NewPatchedJIRAInstanceRequest

`func NewPatchedJIRAInstanceRequest() *PatchedJIRAInstanceRequest`

NewPatchedJIRAInstanceRequest instantiates a new PatchedJIRAInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJIRAInstanceRequestWithDefaults

`func NewPatchedJIRAInstanceRequestWithDefaults() *PatchedJIRAInstanceRequest`

NewPatchedJIRAInstanceRequestWithDefaults instantiates a new PatchedJIRAInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationName

`func (o *PatchedJIRAInstanceRequest) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *PatchedJIRAInstanceRequest) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *PatchedJIRAInstanceRequest) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *PatchedJIRAInstanceRequest) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedJIRAInstanceRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedJIRAInstanceRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedJIRAInstanceRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedJIRAInstanceRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedJIRAInstanceRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedJIRAInstanceRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedJIRAInstanceRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedJIRAInstanceRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PatchedJIRAInstanceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedJIRAInstanceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedJIRAInstanceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedJIRAInstanceRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDefaultIssueType

`func (o *PatchedJIRAInstanceRequest) GetDefaultIssueType() string`

GetDefaultIssueType returns the DefaultIssueType field if non-nil, zero value otherwise.

### GetDefaultIssueTypeOk

`func (o *PatchedJIRAInstanceRequest) GetDefaultIssueTypeOk() (*string, bool)`

GetDefaultIssueTypeOk returns a tuple with the DefaultIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueType

`func (o *PatchedJIRAInstanceRequest) SetDefaultIssueType(v string)`

SetDefaultIssueType sets DefaultIssueType field to given value.

### HasDefaultIssueType

`func (o *PatchedJIRAInstanceRequest) HasDefaultIssueType() bool`

HasDefaultIssueType returns a boolean if a field has been set.

### GetIssueTemplateDir

`func (o *PatchedJIRAInstanceRequest) GetIssueTemplateDir() string`

GetIssueTemplateDir returns the IssueTemplateDir field if non-nil, zero value otherwise.

### GetIssueTemplateDirOk

`func (o *PatchedJIRAInstanceRequest) GetIssueTemplateDirOk() (*string, bool)`

GetIssueTemplateDirOk returns a tuple with the IssueTemplateDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTemplateDir

`func (o *PatchedJIRAInstanceRequest) SetIssueTemplateDir(v string)`

SetIssueTemplateDir sets IssueTemplateDir field to given value.

### HasIssueTemplateDir

`func (o *PatchedJIRAInstanceRequest) HasIssueTemplateDir() bool`

HasIssueTemplateDir returns a boolean if a field has been set.

### SetIssueTemplateDirNil

`func (o *PatchedJIRAInstanceRequest) SetIssueTemplateDirNil(b bool)`

 SetIssueTemplateDirNil sets the value for IssueTemplateDir to be an explicit nil

### UnsetIssueTemplateDir
`func (o *PatchedJIRAInstanceRequest) UnsetIssueTemplateDir()`

UnsetIssueTemplateDir ensures that no value is present for IssueTemplateDir, not even an explicit nil
### GetEpicNameId

`func (o *PatchedJIRAInstanceRequest) GetEpicNameId() int32`

GetEpicNameId returns the EpicNameId field if non-nil, zero value otherwise.

### GetEpicNameIdOk

`func (o *PatchedJIRAInstanceRequest) GetEpicNameIdOk() (*int32, bool)`

GetEpicNameIdOk returns a tuple with the EpicNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicNameId

`func (o *PatchedJIRAInstanceRequest) SetEpicNameId(v int32)`

SetEpicNameId sets EpicNameId field to given value.

### HasEpicNameId

`func (o *PatchedJIRAInstanceRequest) HasEpicNameId() bool`

HasEpicNameId returns a boolean if a field has been set.

### GetOpenStatusKey

`func (o *PatchedJIRAInstanceRequest) GetOpenStatusKey() int32`

GetOpenStatusKey returns the OpenStatusKey field if non-nil, zero value otherwise.

### GetOpenStatusKeyOk

`func (o *PatchedJIRAInstanceRequest) GetOpenStatusKeyOk() (*int32, bool)`

GetOpenStatusKeyOk returns a tuple with the OpenStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatusKey

`func (o *PatchedJIRAInstanceRequest) SetOpenStatusKey(v int32)`

SetOpenStatusKey sets OpenStatusKey field to given value.

### HasOpenStatusKey

`func (o *PatchedJIRAInstanceRequest) HasOpenStatusKey() bool`

HasOpenStatusKey returns a boolean if a field has been set.

### GetCloseStatusKey

`func (o *PatchedJIRAInstanceRequest) GetCloseStatusKey() int32`

GetCloseStatusKey returns the CloseStatusKey field if non-nil, zero value otherwise.

### GetCloseStatusKeyOk

`func (o *PatchedJIRAInstanceRequest) GetCloseStatusKeyOk() (*int32, bool)`

GetCloseStatusKeyOk returns a tuple with the CloseStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseStatusKey

`func (o *PatchedJIRAInstanceRequest) SetCloseStatusKey(v int32)`

SetCloseStatusKey sets CloseStatusKey field to given value.

### HasCloseStatusKey

`func (o *PatchedJIRAInstanceRequest) HasCloseStatusKey() bool`

HasCloseStatusKey returns a boolean if a field has been set.

### GetInfoMappingSeverity

`func (o *PatchedJIRAInstanceRequest) GetInfoMappingSeverity() string`

GetInfoMappingSeverity returns the InfoMappingSeverity field if non-nil, zero value otherwise.

### GetInfoMappingSeverityOk

`func (o *PatchedJIRAInstanceRequest) GetInfoMappingSeverityOk() (*string, bool)`

GetInfoMappingSeverityOk returns a tuple with the InfoMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMappingSeverity

`func (o *PatchedJIRAInstanceRequest) SetInfoMappingSeverity(v string)`

SetInfoMappingSeverity sets InfoMappingSeverity field to given value.

### HasInfoMappingSeverity

`func (o *PatchedJIRAInstanceRequest) HasInfoMappingSeverity() bool`

HasInfoMappingSeverity returns a boolean if a field has been set.

### GetLowMappingSeverity

`func (o *PatchedJIRAInstanceRequest) GetLowMappingSeverity() string`

GetLowMappingSeverity returns the LowMappingSeverity field if non-nil, zero value otherwise.

### GetLowMappingSeverityOk

`func (o *PatchedJIRAInstanceRequest) GetLowMappingSeverityOk() (*string, bool)`

GetLowMappingSeverityOk returns a tuple with the LowMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowMappingSeverity

`func (o *PatchedJIRAInstanceRequest) SetLowMappingSeverity(v string)`

SetLowMappingSeverity sets LowMappingSeverity field to given value.

### HasLowMappingSeverity

`func (o *PatchedJIRAInstanceRequest) HasLowMappingSeverity() bool`

HasLowMappingSeverity returns a boolean if a field has been set.

### GetMediumMappingSeverity

`func (o *PatchedJIRAInstanceRequest) GetMediumMappingSeverity() string`

GetMediumMappingSeverity returns the MediumMappingSeverity field if non-nil, zero value otherwise.

### GetMediumMappingSeverityOk

`func (o *PatchedJIRAInstanceRequest) GetMediumMappingSeverityOk() (*string, bool)`

GetMediumMappingSeverityOk returns a tuple with the MediumMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumMappingSeverity

`func (o *PatchedJIRAInstanceRequest) SetMediumMappingSeverity(v string)`

SetMediumMappingSeverity sets MediumMappingSeverity field to given value.

### HasMediumMappingSeverity

`func (o *PatchedJIRAInstanceRequest) HasMediumMappingSeverity() bool`

HasMediumMappingSeverity returns a boolean if a field has been set.

### GetHighMappingSeverity

`func (o *PatchedJIRAInstanceRequest) GetHighMappingSeverity() string`

GetHighMappingSeverity returns the HighMappingSeverity field if non-nil, zero value otherwise.

### GetHighMappingSeverityOk

`func (o *PatchedJIRAInstanceRequest) GetHighMappingSeverityOk() (*string, bool)`

GetHighMappingSeverityOk returns a tuple with the HighMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighMappingSeverity

`func (o *PatchedJIRAInstanceRequest) SetHighMappingSeverity(v string)`

SetHighMappingSeverity sets HighMappingSeverity field to given value.

### HasHighMappingSeverity

`func (o *PatchedJIRAInstanceRequest) HasHighMappingSeverity() bool`

HasHighMappingSeverity returns a boolean if a field has been set.

### GetCriticalMappingSeverity

`func (o *PatchedJIRAInstanceRequest) GetCriticalMappingSeverity() string`

GetCriticalMappingSeverity returns the CriticalMappingSeverity field if non-nil, zero value otherwise.

### GetCriticalMappingSeverityOk

`func (o *PatchedJIRAInstanceRequest) GetCriticalMappingSeverityOk() (*string, bool)`

GetCriticalMappingSeverityOk returns a tuple with the CriticalMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalMappingSeverity

`func (o *PatchedJIRAInstanceRequest) SetCriticalMappingSeverity(v string)`

SetCriticalMappingSeverity sets CriticalMappingSeverity field to given value.

### HasCriticalMappingSeverity

`func (o *PatchedJIRAInstanceRequest) HasCriticalMappingSeverity() bool`

HasCriticalMappingSeverity returns a boolean if a field has been set.

### GetFindingText

`func (o *PatchedJIRAInstanceRequest) GetFindingText() string`

GetFindingText returns the FindingText field if non-nil, zero value otherwise.

### GetFindingTextOk

`func (o *PatchedJIRAInstanceRequest) GetFindingTextOk() (*string, bool)`

GetFindingTextOk returns a tuple with the FindingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingText

`func (o *PatchedJIRAInstanceRequest) SetFindingText(v string)`

SetFindingText sets FindingText field to given value.

### HasFindingText

`func (o *PatchedJIRAInstanceRequest) HasFindingText() bool`

HasFindingText returns a boolean if a field has been set.

### SetFindingTextNil

`func (o *PatchedJIRAInstanceRequest) SetFindingTextNil(b bool)`

 SetFindingTextNil sets the value for FindingText to be an explicit nil

### UnsetFindingText
`func (o *PatchedJIRAInstanceRequest) UnsetFindingText()`

UnsetFindingText ensures that no value is present for FindingText, not even an explicit nil
### GetAcceptedMappingResolution

`func (o *PatchedJIRAInstanceRequest) GetAcceptedMappingResolution() string`

GetAcceptedMappingResolution returns the AcceptedMappingResolution field if non-nil, zero value otherwise.

### GetAcceptedMappingResolutionOk

`func (o *PatchedJIRAInstanceRequest) GetAcceptedMappingResolutionOk() (*string, bool)`

GetAcceptedMappingResolutionOk returns a tuple with the AcceptedMappingResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedMappingResolution

`func (o *PatchedJIRAInstanceRequest) SetAcceptedMappingResolution(v string)`

SetAcceptedMappingResolution sets AcceptedMappingResolution field to given value.

### HasAcceptedMappingResolution

`func (o *PatchedJIRAInstanceRequest) HasAcceptedMappingResolution() bool`

HasAcceptedMappingResolution returns a boolean if a field has been set.

### SetAcceptedMappingResolutionNil

`func (o *PatchedJIRAInstanceRequest) SetAcceptedMappingResolutionNil(b bool)`

 SetAcceptedMappingResolutionNil sets the value for AcceptedMappingResolution to be an explicit nil

### UnsetAcceptedMappingResolution
`func (o *PatchedJIRAInstanceRequest) UnsetAcceptedMappingResolution()`

UnsetAcceptedMappingResolution ensures that no value is present for AcceptedMappingResolution, not even an explicit nil
### GetFalsePositiveMappingResolution

`func (o *PatchedJIRAInstanceRequest) GetFalsePositiveMappingResolution() string`

GetFalsePositiveMappingResolution returns the FalsePositiveMappingResolution field if non-nil, zero value otherwise.

### GetFalsePositiveMappingResolutionOk

`func (o *PatchedJIRAInstanceRequest) GetFalsePositiveMappingResolutionOk() (*string, bool)`

GetFalsePositiveMappingResolutionOk returns a tuple with the FalsePositiveMappingResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositiveMappingResolution

`func (o *PatchedJIRAInstanceRequest) SetFalsePositiveMappingResolution(v string)`

SetFalsePositiveMappingResolution sets FalsePositiveMappingResolution field to given value.

### HasFalsePositiveMappingResolution

`func (o *PatchedJIRAInstanceRequest) HasFalsePositiveMappingResolution() bool`

HasFalsePositiveMappingResolution returns a boolean if a field has been set.

### SetFalsePositiveMappingResolutionNil

`func (o *PatchedJIRAInstanceRequest) SetFalsePositiveMappingResolutionNil(b bool)`

 SetFalsePositiveMappingResolutionNil sets the value for FalsePositiveMappingResolution to be an explicit nil

### UnsetFalsePositiveMappingResolution
`func (o *PatchedJIRAInstanceRequest) UnsetFalsePositiveMappingResolution()`

UnsetFalsePositiveMappingResolution ensures that no value is present for FalsePositiveMappingResolution, not even an explicit nil
### GetGlobalJiraSlaNotification

`func (o *PatchedJIRAInstanceRequest) GetGlobalJiraSlaNotification() bool`

GetGlobalJiraSlaNotification returns the GlobalJiraSlaNotification field if non-nil, zero value otherwise.

### GetGlobalJiraSlaNotificationOk

`func (o *PatchedJIRAInstanceRequest) GetGlobalJiraSlaNotificationOk() (*bool, bool)`

GetGlobalJiraSlaNotificationOk returns a tuple with the GlobalJiraSlaNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalJiraSlaNotification

`func (o *PatchedJIRAInstanceRequest) SetGlobalJiraSlaNotification(v bool)`

SetGlobalJiraSlaNotification sets GlobalJiraSlaNotification field to given value.

### HasGlobalJiraSlaNotification

`func (o *PatchedJIRAInstanceRequest) HasGlobalJiraSlaNotification() bool`

HasGlobalJiraSlaNotification returns a boolean if a field has been set.

### GetFindingJiraSync

`func (o *PatchedJIRAInstanceRequest) GetFindingJiraSync() bool`

GetFindingJiraSync returns the FindingJiraSync field if non-nil, zero value otherwise.

### GetFindingJiraSyncOk

`func (o *PatchedJIRAInstanceRequest) GetFindingJiraSyncOk() (*bool, bool)`

GetFindingJiraSyncOk returns a tuple with the FindingJiraSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingJiraSync

`func (o *PatchedJIRAInstanceRequest) SetFindingJiraSync(v bool)`

SetFindingJiraSync sets FindingJiraSync field to given value.

### HasFindingJiraSync

`func (o *PatchedJIRAInstanceRequest) HasFindingJiraSync() bool`

HasFindingJiraSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


