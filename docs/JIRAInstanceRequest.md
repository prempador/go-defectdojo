# JIRAInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationName** | Pointer to **string** | Enter a name to give to this configuration | [optional] 
**Url** | **string** | For more information how to configure Jira, read the DefectDojo documentation. | 
**Username** | **string** |  | 
**Password** | **string** |  | 
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

### NewJIRAInstanceRequest

`func NewJIRAInstanceRequest(url string, username string, password string, epicNameId int32, openStatusKey int32, closeStatusKey int32, infoMappingSeverity string, lowMappingSeverity string, mediumMappingSeverity string, highMappingSeverity string, criticalMappingSeverity string, ) *JIRAInstanceRequest`

NewJIRAInstanceRequest instantiates a new JIRAInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJIRAInstanceRequestWithDefaults

`func NewJIRAInstanceRequestWithDefaults() *JIRAInstanceRequest`

NewJIRAInstanceRequestWithDefaults instantiates a new JIRAInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationName

`func (o *JIRAInstanceRequest) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *JIRAInstanceRequest) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *JIRAInstanceRequest) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.

### HasConfigurationName

`func (o *JIRAInstanceRequest) HasConfigurationName() bool`

HasConfigurationName returns a boolean if a field has been set.

### GetUrl

`func (o *JIRAInstanceRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JIRAInstanceRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JIRAInstanceRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsername

`func (o *JIRAInstanceRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *JIRAInstanceRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *JIRAInstanceRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *JIRAInstanceRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JIRAInstanceRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JIRAInstanceRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDefaultIssueType

`func (o *JIRAInstanceRequest) GetDefaultIssueType() string`

GetDefaultIssueType returns the DefaultIssueType field if non-nil, zero value otherwise.

### GetDefaultIssueTypeOk

`func (o *JIRAInstanceRequest) GetDefaultIssueTypeOk() (*string, bool)`

GetDefaultIssueTypeOk returns a tuple with the DefaultIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIssueType

`func (o *JIRAInstanceRequest) SetDefaultIssueType(v string)`

SetDefaultIssueType sets DefaultIssueType field to given value.

### HasDefaultIssueType

`func (o *JIRAInstanceRequest) HasDefaultIssueType() bool`

HasDefaultIssueType returns a boolean if a field has been set.

### GetIssueTemplateDir

`func (o *JIRAInstanceRequest) GetIssueTemplateDir() string`

GetIssueTemplateDir returns the IssueTemplateDir field if non-nil, zero value otherwise.

### GetIssueTemplateDirOk

`func (o *JIRAInstanceRequest) GetIssueTemplateDirOk() (*string, bool)`

GetIssueTemplateDirOk returns a tuple with the IssueTemplateDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTemplateDir

`func (o *JIRAInstanceRequest) SetIssueTemplateDir(v string)`

SetIssueTemplateDir sets IssueTemplateDir field to given value.

### HasIssueTemplateDir

`func (o *JIRAInstanceRequest) HasIssueTemplateDir() bool`

HasIssueTemplateDir returns a boolean if a field has been set.

### SetIssueTemplateDirNil

`func (o *JIRAInstanceRequest) SetIssueTemplateDirNil(b bool)`

 SetIssueTemplateDirNil sets the value for IssueTemplateDir to be an explicit nil

### UnsetIssueTemplateDir
`func (o *JIRAInstanceRequest) UnsetIssueTemplateDir()`

UnsetIssueTemplateDir ensures that no value is present for IssueTemplateDir, not even an explicit nil
### GetEpicNameId

`func (o *JIRAInstanceRequest) GetEpicNameId() int32`

GetEpicNameId returns the EpicNameId field if non-nil, zero value otherwise.

### GetEpicNameIdOk

`func (o *JIRAInstanceRequest) GetEpicNameIdOk() (*int32, bool)`

GetEpicNameIdOk returns a tuple with the EpicNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicNameId

`func (o *JIRAInstanceRequest) SetEpicNameId(v int32)`

SetEpicNameId sets EpicNameId field to given value.


### GetOpenStatusKey

`func (o *JIRAInstanceRequest) GetOpenStatusKey() int32`

GetOpenStatusKey returns the OpenStatusKey field if non-nil, zero value otherwise.

### GetOpenStatusKeyOk

`func (o *JIRAInstanceRequest) GetOpenStatusKeyOk() (*int32, bool)`

GetOpenStatusKeyOk returns a tuple with the OpenStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStatusKey

`func (o *JIRAInstanceRequest) SetOpenStatusKey(v int32)`

SetOpenStatusKey sets OpenStatusKey field to given value.


### GetCloseStatusKey

`func (o *JIRAInstanceRequest) GetCloseStatusKey() int32`

GetCloseStatusKey returns the CloseStatusKey field if non-nil, zero value otherwise.

### GetCloseStatusKeyOk

`func (o *JIRAInstanceRequest) GetCloseStatusKeyOk() (*int32, bool)`

GetCloseStatusKeyOk returns a tuple with the CloseStatusKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseStatusKey

`func (o *JIRAInstanceRequest) SetCloseStatusKey(v int32)`

SetCloseStatusKey sets CloseStatusKey field to given value.


### GetInfoMappingSeverity

`func (o *JIRAInstanceRequest) GetInfoMappingSeverity() string`

GetInfoMappingSeverity returns the InfoMappingSeverity field if non-nil, zero value otherwise.

### GetInfoMappingSeverityOk

`func (o *JIRAInstanceRequest) GetInfoMappingSeverityOk() (*string, bool)`

GetInfoMappingSeverityOk returns a tuple with the InfoMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoMappingSeverity

`func (o *JIRAInstanceRequest) SetInfoMappingSeverity(v string)`

SetInfoMappingSeverity sets InfoMappingSeverity field to given value.


### GetLowMappingSeverity

`func (o *JIRAInstanceRequest) GetLowMappingSeverity() string`

GetLowMappingSeverity returns the LowMappingSeverity field if non-nil, zero value otherwise.

### GetLowMappingSeverityOk

`func (o *JIRAInstanceRequest) GetLowMappingSeverityOk() (*string, bool)`

GetLowMappingSeverityOk returns a tuple with the LowMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowMappingSeverity

`func (o *JIRAInstanceRequest) SetLowMappingSeverity(v string)`

SetLowMappingSeverity sets LowMappingSeverity field to given value.


### GetMediumMappingSeverity

`func (o *JIRAInstanceRequest) GetMediumMappingSeverity() string`

GetMediumMappingSeverity returns the MediumMappingSeverity field if non-nil, zero value otherwise.

### GetMediumMappingSeverityOk

`func (o *JIRAInstanceRequest) GetMediumMappingSeverityOk() (*string, bool)`

GetMediumMappingSeverityOk returns a tuple with the MediumMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumMappingSeverity

`func (o *JIRAInstanceRequest) SetMediumMappingSeverity(v string)`

SetMediumMappingSeverity sets MediumMappingSeverity field to given value.


### GetHighMappingSeverity

`func (o *JIRAInstanceRequest) GetHighMappingSeverity() string`

GetHighMappingSeverity returns the HighMappingSeverity field if non-nil, zero value otherwise.

### GetHighMappingSeverityOk

`func (o *JIRAInstanceRequest) GetHighMappingSeverityOk() (*string, bool)`

GetHighMappingSeverityOk returns a tuple with the HighMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighMappingSeverity

`func (o *JIRAInstanceRequest) SetHighMappingSeverity(v string)`

SetHighMappingSeverity sets HighMappingSeverity field to given value.


### GetCriticalMappingSeverity

`func (o *JIRAInstanceRequest) GetCriticalMappingSeverity() string`

GetCriticalMappingSeverity returns the CriticalMappingSeverity field if non-nil, zero value otherwise.

### GetCriticalMappingSeverityOk

`func (o *JIRAInstanceRequest) GetCriticalMappingSeverityOk() (*string, bool)`

GetCriticalMappingSeverityOk returns a tuple with the CriticalMappingSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalMappingSeverity

`func (o *JIRAInstanceRequest) SetCriticalMappingSeverity(v string)`

SetCriticalMappingSeverity sets CriticalMappingSeverity field to given value.


### GetFindingText

`func (o *JIRAInstanceRequest) GetFindingText() string`

GetFindingText returns the FindingText field if non-nil, zero value otherwise.

### GetFindingTextOk

`func (o *JIRAInstanceRequest) GetFindingTextOk() (*string, bool)`

GetFindingTextOk returns a tuple with the FindingText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingText

`func (o *JIRAInstanceRequest) SetFindingText(v string)`

SetFindingText sets FindingText field to given value.

### HasFindingText

`func (o *JIRAInstanceRequest) HasFindingText() bool`

HasFindingText returns a boolean if a field has been set.

### SetFindingTextNil

`func (o *JIRAInstanceRequest) SetFindingTextNil(b bool)`

 SetFindingTextNil sets the value for FindingText to be an explicit nil

### UnsetFindingText
`func (o *JIRAInstanceRequest) UnsetFindingText()`

UnsetFindingText ensures that no value is present for FindingText, not even an explicit nil
### GetAcceptedMappingResolution

`func (o *JIRAInstanceRequest) GetAcceptedMappingResolution() string`

GetAcceptedMappingResolution returns the AcceptedMappingResolution field if non-nil, zero value otherwise.

### GetAcceptedMappingResolutionOk

`func (o *JIRAInstanceRequest) GetAcceptedMappingResolutionOk() (*string, bool)`

GetAcceptedMappingResolutionOk returns a tuple with the AcceptedMappingResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedMappingResolution

`func (o *JIRAInstanceRequest) SetAcceptedMappingResolution(v string)`

SetAcceptedMappingResolution sets AcceptedMappingResolution field to given value.

### HasAcceptedMappingResolution

`func (o *JIRAInstanceRequest) HasAcceptedMappingResolution() bool`

HasAcceptedMappingResolution returns a boolean if a field has been set.

### SetAcceptedMappingResolutionNil

`func (o *JIRAInstanceRequest) SetAcceptedMappingResolutionNil(b bool)`

 SetAcceptedMappingResolutionNil sets the value for AcceptedMappingResolution to be an explicit nil

### UnsetAcceptedMappingResolution
`func (o *JIRAInstanceRequest) UnsetAcceptedMappingResolution()`

UnsetAcceptedMappingResolution ensures that no value is present for AcceptedMappingResolution, not even an explicit nil
### GetFalsePositiveMappingResolution

`func (o *JIRAInstanceRequest) GetFalsePositiveMappingResolution() string`

GetFalsePositiveMappingResolution returns the FalsePositiveMappingResolution field if non-nil, zero value otherwise.

### GetFalsePositiveMappingResolutionOk

`func (o *JIRAInstanceRequest) GetFalsePositiveMappingResolutionOk() (*string, bool)`

GetFalsePositiveMappingResolutionOk returns a tuple with the FalsePositiveMappingResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositiveMappingResolution

`func (o *JIRAInstanceRequest) SetFalsePositiveMappingResolution(v string)`

SetFalsePositiveMappingResolution sets FalsePositiveMappingResolution field to given value.

### HasFalsePositiveMappingResolution

`func (o *JIRAInstanceRequest) HasFalsePositiveMappingResolution() bool`

HasFalsePositiveMappingResolution returns a boolean if a field has been set.

### SetFalsePositiveMappingResolutionNil

`func (o *JIRAInstanceRequest) SetFalsePositiveMappingResolutionNil(b bool)`

 SetFalsePositiveMappingResolutionNil sets the value for FalsePositiveMappingResolution to be an explicit nil

### UnsetFalsePositiveMappingResolution
`func (o *JIRAInstanceRequest) UnsetFalsePositiveMappingResolution()`

UnsetFalsePositiveMappingResolution ensures that no value is present for FalsePositiveMappingResolution, not even an explicit nil
### GetGlobalJiraSlaNotification

`func (o *JIRAInstanceRequest) GetGlobalJiraSlaNotification() bool`

GetGlobalJiraSlaNotification returns the GlobalJiraSlaNotification field if non-nil, zero value otherwise.

### GetGlobalJiraSlaNotificationOk

`func (o *JIRAInstanceRequest) GetGlobalJiraSlaNotificationOk() (*bool, bool)`

GetGlobalJiraSlaNotificationOk returns a tuple with the GlobalJiraSlaNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalJiraSlaNotification

`func (o *JIRAInstanceRequest) SetGlobalJiraSlaNotification(v bool)`

SetGlobalJiraSlaNotification sets GlobalJiraSlaNotification field to given value.

### HasGlobalJiraSlaNotification

`func (o *JIRAInstanceRequest) HasGlobalJiraSlaNotification() bool`

HasGlobalJiraSlaNotification returns a boolean if a field has been set.

### GetFindingJiraSync

`func (o *JIRAInstanceRequest) GetFindingJiraSync() bool`

GetFindingJiraSync returns the FindingJiraSync field if non-nil, zero value otherwise.

### GetFindingJiraSyncOk

`func (o *JIRAInstanceRequest) GetFindingJiraSyncOk() (*bool, bool)`

GetFindingJiraSyncOk returns a tuple with the FindingJiraSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingJiraSync

`func (o *JIRAInstanceRequest) SetFindingJiraSync(v bool)`

SetFindingJiraSync sets FindingJiraSync field to given value.

### HasFindingJiraSync

`func (o *JIRAInstanceRequest) HasFindingJiraSync() bool`

HasFindingJiraSync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


