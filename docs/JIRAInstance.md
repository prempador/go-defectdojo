# JIRAInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ConfigurationName** | Pointer to **string** | Enter a name to give to this configuration | [optional] 
**Url** | **string** | For more information how to configure Jira, read the DefectDojo documentation. | 
**Username** | **string** |  | 
**DefaultIssueType** | Pointer to **string** | You can define extra issue types in settings.py  * &#x60;Task&#x60; - Task * &#x60;Story&#x60; - Story * &#x60;Epic&#x60; - Epic * &#x60;Spike&#x60; - Spike * &#x60;Bug&#x60; - Bug * &#x60;Security&#x60; - Security | [optional] 
**IssueTemplateDir** | Pointer to **NullableString** | Choose the folder containing the Django templates used to render the JIRA issue description. These are stored in dojo/templates/issue-trackers. Leave empty to use the default jira_full templates. | [optional] 
**EpicNameId** | **int32** | To obtain the &#39;Epic name id&#39; visit https://&lt;YOUR JIRA URL&gt;/rest/api/2/field and search for Epic Name. Copy the number out of cf[number] and paste it here. | 
**OpenStatusKey** | **int32** | Transition ID to Re-Open JIRA issues, visit https://&lt;YOUR JIRA URL&gt;/rest/api/latest/issue/&lt;ANY VALID ISSUE KEY&gt;/transitions?expand&#x3D;transitions.fields to find the ID for your JIRA instance | 
**CloseStatusKey** | **int32** | Transition ID to Close JIRA issues, visit https://&lt;YOUR JIRA URL&gt;/rest/api/latest/issue/&lt;ANY VALID ISSUE KEY&gt;/transitions?expand&#x3D;transitions.fields to find the ID for your JIRA instance | 
**InfoMappingSeverity** | **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Info | 
**LowMappingSeverity** | **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Low | 
**MediumMappingSeverity** | **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Medium | 
**HighMappingSeverity** | **string** | Maps to the &#39;Priority&#39; field in Jira. For example: High | 
**CriticalMappingSeverity** | **string** | Maps to the &#39;Priority&#39; field in Jira. For example: Critical | 
**FindingText** | Pointer to **NullableString** | Additional text that will be added to the finding in Jira. For example including how the finding was created or who to contact for more information. | [optional] 
**AcceptedMappingResolution** | Pointer to **NullableString** | JIRA resolution names (comma-separated values) that maps to an Accepted Finding | [optional] 
**FalsePositiveMappingResolution** | Pointer to **NullableString** | JIRA resolution names (comma-separated values) that maps to a False Positive Finding | [optional] 
**GlobalJiraSlaNotification** | Pointer to **bool** | This setting can be overidden at the Product level | [optional] 
**FindingJiraSync** | Pointer to **bool** | If enabled, this will sync changes to a Finding automatically to JIRA | [optional] 

## Methods

### NewJIRAInstance

`func NewJIRAInstance(id int32, url string, username string, epicNameId int32, openStatusKey int32, closeStatusKey int32, infoMappingSeverity string, lowMappingSeverity string, mediumMappingSeverity string, highMappingSeverity string, criticalMappingSeverity string, ) *JIRAInstance`

NewJIRAInstance instantiates a new JIRAInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJIRAInstanceWithDefaults

`func NewJIRAInstanceWithDefaults() *JIRAInstance`

NewJIRAInstanceWithDefaults instantiates a new JIRAInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JIRAInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JIRAInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JIRAInstance) SetId(v int32)`

SetId sets Id field to given value.


### GetConfigurationName

`func (o *JIRAInstance) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *JIRAInstance) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *JIRAInstance) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *JIRAInstance) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### GetUrl

`func (o *JIRAInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JIRAInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JIRAInstance) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *JIRAInstance) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *JIRAInstance) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *JIRAInstance) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDefaultIssueType

`func (o *JIRAInstance) GetDefaultIssueType() string`

GetDefaultIssueType returns the DefaultIssueType field if non-nil, zero value otherwise.

### GetDefaultIssueTypeOk

`func (o *JIRAInstance) GetDefaultIssueTypeOk() (*string, bool)`

GetDefaultIssueTypeOk returns a tuple with the DefaultIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueType

`func (o *JIRAInstance) SetDefaultIssueType(v string)`

SetDefaultIssueType sets DefaultIssueType field to given value.

### HasDefaultIssueType

`func (o *JIRAInstance) HasDefaultIssueType() bool`

HasDefaultIssueType returns a boolean if a field has been set.

### GetIssueTemplateDir

`func (o *JIRAInstance) GetIssueTemplateDir() string`

GetIssueTemplateDir returns the IssueTemplateDir field if non-nil, zero value otherwise.

### GetIssueTemplateDirOk

`func (o *JIRAInstance) GetIssueTemplateDirOk() (*string, bool)`

GetIssueTemplateDirOk returns a tuple with the IssueTemplateDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTemplateDir

`func (o *JIRAInstance) SetIssueTemplateDir(v string)`

SetIssueTemplateDir sets IssueTemplateDir field to given value.

### HasIssueTemplateDir

`func (o *JIRAInstance) HasIssueTemplateDir() bool`

HasIssueTemplateDir returns a boolean if a field has been set.

### SetIssueTemplateDirNil

`func (o *JIRAInstance) SetIssueTemplateDirNil(b bool)`

 SetIssueTemplateDirNil sets the value for IssueTemplateDir to be an explicit nil

### UnsetIssueTemplateDir
`func (o *JIRAInstance) UnsetIssueTemplateDir()`

UnsetIssueTemplateDir ensures that no value is present for IssueTemplateDir, not even an explicit nil
### GetEpicNameId

`func (o *JIRAInstance) GetEpicNameId() int32`

GetEpicNameId returns the EpicNameId field if non-nil, zero value otherwise.

### GetEpicNameIdOk

`func (o *JIRAInstance) GetEpicNameIdOk() (*int32, bool)`

GetEpicNameIdOk returns a tuple with the EpicNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicNameId

`func (o *JIRAInstance) SetEpicNameId(v int32)`

SetEpicNameId sets EpicNameId field to given value.


### GetOpenStatusKey

`func (o *JIRAInstance) GetOpenStatusKey() int32`

GetOpenStatusKey returns the OpenStatusKey field if non-nil, zero value otherwise.

### GetOpenStatusKeyOk

`func (o *JIRAInstance) GetOpenStatusKeyOk() (*int32, bool)`

GetOpenStatusKeyOk returns a tuple with the OpenStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatusKey

`func (o *JIRAInstance) SetOpenStatusKey(v int32)`

SetOpenStatusKey sets OpenStatusKey field to given value.


### GetCloseStatusKey

`func (o *JIRAInstance) GetCloseStatusKey() int32`

GetCloseStatusKey returns the CloseStatusKey field if non-nil, zero value otherwise.

### GetCloseStatusKeyOk

`func (o *JIRAInstance) GetCloseStatusKeyOk() (*int32, bool)`

GetCloseStatusKeyOk returns a tuple with the CloseStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseStatusKey

`func (o *JIRAInstance) SetCloseStatusKey(v int32)`

SetCloseStatusKey sets CloseStatusKey field to given value.


### GetInfoMappingSeverity

`func (o *JIRAInstance) GetInfoMappingSeverity() string`

GetInfoMappingSeverity returns the InfoMappingSeverity field if non-nil, zero value otherwise.

### GetInfoMappingSeverityOk

`func (o *JIRAInstance) GetInfoMappingSeverityOk() (*string, bool)`

GetInfoMappingSeverityOk returns a tuple with the InfoMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMappingSeverity

`func (o *JIRAInstance) SetInfoMappingSeverity(v string)`

SetInfoMappingSeverity sets InfoMappingSeverity field to given value.


### GetLowMappingSeverity

`func (o *JIRAInstance) GetLowMappingSeverity() string`

GetLowMappingSeverity returns the LowMappingSeverity field if non-nil, zero value otherwise.

### GetLowMappingSeverityOk

`func (o *JIRAInstance) GetLowMappingSeverityOk() (*string, bool)`

GetLowMappingSeverityOk returns a tuple with the LowMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowMappingSeverity

`func (o *JIRAInstance) SetLowMappingSeverity(v string)`

SetLowMappingSeverity sets LowMappingSeverity field to given value.


### GetMediumMappingSeverity

`func (o *JIRAInstance) GetMediumMappingSeverity() string`

GetMediumMappingSeverity returns the MediumMappingSeverity field if non-nil, zero value otherwise.

### GetMediumMappingSeverityOk

`func (o *JIRAInstance) GetMediumMappingSeverityOk() (*string, bool)`

GetMediumMappingSeverityOk returns a tuple with the MediumMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumMappingSeverity

`func (o *JIRAInstance) SetMediumMappingSeverity(v string)`

SetMediumMappingSeverity sets MediumMappingSeverity field to given value.


### GetHighMappingSeverity

`func (o *JIRAInstance) GetHighMappingSeverity() string`

GetHighMappingSeverity returns the HighMappingSeverity field if non-nil, zero value otherwise.

### GetHighMappingSeverityOk

`func (o *JIRAInstance) GetHighMappingSeverityOk() (*string, bool)`

GetHighMappingSeverityOk returns a tuple with the HighMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighMappingSeverity

`func (o *JIRAInstance) SetHighMappingSeverity(v string)`

SetHighMappingSeverity sets HighMappingSeverity field to given value.


### GetCriticalMappingSeverity

`func (o *JIRAInstance) GetCriticalMappingSeverity() string`

GetCriticalMappingSeverity returns the CriticalMappingSeverity field if non-nil, zero value otherwise.

### GetCriticalMappingSeverityOk

`func (o *JIRAInstance) GetCriticalMappingSeverityOk() (*string, bool)`

GetCriticalMappingSeverityOk returns a tuple with the CriticalMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalMappingSeverity

`func (o *JIRAInstance) SetCriticalMappingSeverity(v string)`

SetCriticalMappingSeverity sets CriticalMappingSeverity field to given value.


### GetFindingText

`func (o *JIRAInstance) GetFindingText() string`

GetFindingText returns the FindingText field if non-nil, zero value otherwise.

### GetFindingTextOk

`func (o *JIRAInstance) GetFindingTextOk() (*string, bool)`

GetFindingTextOk returns a tuple with the FindingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingText

`func (o *JIRAInstance) SetFindingText(v string)`

SetFindingText sets FindingText field to given value.

### HasFindingText

`func (o *JIRAInstance) HasFindingText() bool`

HasFindingText returns a boolean if a field has been set.

### SetFindingTextNil

`func (o *JIRAInstance) SetFindingTextNil(b bool)`

 SetFindingTextNil sets the value for FindingText to be an explicit nil

### UnsetFindingText
`func (o *JIRAInstance) UnsetFindingText()`

UnsetFindingText ensures that no value is present for FindingText, not even an explicit nil
### GetAcceptedMappingResolution

`func (o *JIRAInstance) GetAcceptedMappingResolution() string`

GetAcceptedMappingResolution returns the AcceptedMappingResolution field if non-nil, zero value otherwise.

### GetAcceptedMappingResolutionOk

`func (o *JIRAInstance) GetAcceptedMappingResolutionOk() (*string, bool)`

GetAcceptedMappingResolutionOk returns a tuple with the AcceptedMappingResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedMappingResolution

`func (o *JIRAInstance) SetAcceptedMappingResolution(v string)`

SetAcceptedMappingResolution sets AcceptedMappingResolution field to given value.

### HasAcceptedMappingResolution

`func (o *JIRAInstance) HasAcceptedMappingResolution() bool`

HasAcceptedMappingResolution returns a boolean if a field has been set.

### SetAcceptedMappingResolutionNil

`func (o *JIRAInstance) SetAcceptedMappingResolutionNil(b bool)`

 SetAcceptedMappingResolutionNil sets the value for AcceptedMappingResolution to be an explicit nil

### UnsetAcceptedMappingResolution
`func (o *JIRAInstance) UnsetAcceptedMappingResolution()`

UnsetAcceptedMappingResolution ensures that no value is present for AcceptedMappingResolution, not even an explicit nil
### GetFalsePositiveMappingResolution

`func (o *JIRAInstance) GetFalsePositiveMappingResolution() string`

GetFalsePositiveMappingResolution returns the FalsePositiveMappingResolution field if non-nil, zero value otherwise.

### GetFalsePositiveMappingResolutionOk

`func (o *JIRAInstance) GetFalsePositiveMappingResolutionOk() (*string, bool)`

GetFalsePositiveMappingResolutionOk returns a tuple with the FalsePositiveMappingResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositiveMappingResolution

`func (o *JIRAInstance) SetFalsePositiveMappingResolution(v string)`

SetFalsePositiveMappingResolution sets FalsePositiveMappingResolution field to given value.

### HasFalsePositiveMappingResolution

`func (o *JIRAInstance) HasFalsePositiveMappingResolution() bool`

HasFalsePositiveMappingResolution returns a boolean if a field has been set.

### SetFalsePositiveMappingResolutionNil

`func (o *JIRAInstance) SetFalsePositiveMappingResolutionNil(b bool)`

 SetFalsePositiveMappingResolutionNil sets the value for FalsePositiveMappingResolution to be an explicit nil

### UnsetFalsePositiveMappingResolution
`func (o *JIRAInstance) UnsetFalsePositiveMappingResolution()`

UnsetFalsePositiveMappingResolution ensures that no value is present for FalsePositiveMappingResolution, not even an explicit nil
### GetGlobalJiraSlaNotification

`func (o *JIRAInstance) GetGlobalJiraSlaNotification() bool`

GetGlobalJiraSlaNotification returns the GlobalJiraSlaNotification field if non-nil, zero value otherwise.

### GetGlobalJiraSlaNotificationOk

`func (o *JIRAInstance) GetGlobalJiraSlaNotificationOk() (*bool, bool)`

GetGlobalJiraSlaNotificationOk returns a tuple with the GlobalJiraSlaNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalJiraSlaNotification

`func (o *JIRAInstance) SetGlobalJiraSlaNotification(v bool)`

SetGlobalJiraSlaNotification sets GlobalJiraSlaNotification field to given value.

### HasGlobalJiraSlaNotification

`func (o *JIRAInstance) HasGlobalJiraSlaNotification() bool`

HasGlobalJiraSlaNotification returns a boolean if a field has been set.

### GetFindingJiraSync

`func (o *JIRAInstance) GetFindingJiraSync() bool`

GetFindingJiraSync returns the FindingJiraSync field if non-nil, zero value otherwise.

### GetFindingJiraSyncOk

`func (o *JIRAInstance) GetFindingJiraSyncOk() (*bool, bool)`

GetFindingJiraSyncOk returns a tuple with the FindingJiraSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingJiraSync

`func (o *JIRAInstance) SetFindingJiraSync(v bool)`

SetFindingJiraSync sets FindingJiraSync field to given value.

### HasFindingJiraSync

`func (o *JIRAInstance) HasFindingJiraSync() bool`

HasFindingJiraSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


