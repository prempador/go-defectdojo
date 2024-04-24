# PatchedJIRAProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectKey** | Pointer to **string** |  | [optional] 
**IssueTemplateDir** | Pointer to **NullableString** | Choose the folder containing the Django templates used to render the JIRA issue description. These are stored in dojo/templates/issue-trackers. Leave empty to use the default jira_full templates. | [optional] 
**Component** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **interface{}** | JIRA custom field JSON mapping of Id to value, e.g. {\&quot;customfield_10122\&quot;: [{\&quot;name\&quot;: \&quot;8.0.1\&quot;}]} | [optional] 
**DefaultAssignee** | Pointer to **NullableString** | JIRA default assignee (name). If left blank then it defaults to whatever is configured in JIRA. | [optional] 
**JiraLabels** | Pointer to **NullableString** | JIRA issue labels space seperated | [optional] 
**AddVulnerabilityIdToJiraLabel** | Pointer to **bool** |  | [optional] 
**PushAllIssues** | Pointer to **bool** | Automatically maintain parity with JIRA. Always create and update JIRA tickets for findings in this Product. | [optional] 
**EnableEngagementEpicMapping** | Pointer to **bool** |  | [optional] 
**EpicIssueTypeName** | Pointer to **string** | The name of the of structure that represents an Epic | [optional] 
**PushNotes** | Pointer to **bool** |  | [optional] 
**ProductJiraSlaNotification** | Pointer to **bool** |  | [optional] 
**RiskAcceptanceExpirationNotification** | Pointer to **bool** |  | [optional] 
**JiraInstance** | Pointer to **NullableInt32** |  | [optional] 
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPatchedJIRAProjectRequest

`func NewPatchedJIRAProjectRequest() *PatchedJIRAProjectRequest`

NewPatchedJIRAProjectRequest instantiates a new PatchedJIRAProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJIRAProjectRequestWithDefaults

`func NewPatchedJIRAProjectRequestWithDefaults() *PatchedJIRAProjectRequest`

NewPatchedJIRAProjectRequestWithDefaults instantiates a new PatchedJIRAProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectKey

`func (o *PatchedJIRAProjectRequest) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *PatchedJIRAProjectRequest) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *PatchedJIRAProjectRequest) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *PatchedJIRAProjectRequest) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetIssueTemplateDir

`func (o *PatchedJIRAProjectRequest) GetIssueTemplateDir() string`

GetIssueTemplateDir returns the IssueTemplateDir field if non-nil, zero value otherwise.

### GetIssueTemplateDirOk

`func (o *PatchedJIRAProjectRequest) GetIssueTemplateDirOk() (*string, bool)`

GetIssueTemplateDirOk returns a tuple with the IssueTemplateDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTemplateDir

`func (o *PatchedJIRAProjectRequest) SetIssueTemplateDir(v string)`

SetIssueTemplateDir sets IssueTemplateDir field to given value.

### HasIssueTemplateDir

`func (o *PatchedJIRAProjectRequest) HasIssueTemplateDir() bool`

HasIssueTemplateDir returns a boolean if a field has been set.

### SetIssueTemplateDirNil

`func (o *PatchedJIRAProjectRequest) SetIssueTemplateDirNil(b bool)`

 SetIssueTemplateDirNil sets the value for IssueTemplateDir to be an explicit nil

### UnsetIssueTemplateDir
`func (o *PatchedJIRAProjectRequest) UnsetIssueTemplateDir()`

UnsetIssueTemplateDir ensures that no value is present for IssueTemplateDir, not even an explicit nil
### GetComponent

`func (o *PatchedJIRAProjectRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PatchedJIRAProjectRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PatchedJIRAProjectRequest) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PatchedJIRAProjectRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedJIRAProjectRequest) GetCustomFields() interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedJIRAProjectRequest) GetCustomFieldsOk() (*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedJIRAProjectRequest) SetCustomFields(v interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedJIRAProjectRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *PatchedJIRAProjectRequest) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *PatchedJIRAProjectRequest) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDefaultAssignee

`func (o *PatchedJIRAProjectRequest) GetDefaultAssignee() string`

GetDefaultAssignee returns the DefaultAssignee field if non-nil, zero value otherwise.

### GetDefaultAssigneeOk

`func (o *PatchedJIRAProjectRequest) GetDefaultAssigneeOk() (*string, bool)`

GetDefaultAssigneeOk returns a tuple with the DefaultAssignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAssignee

`func (o *PatchedJIRAProjectRequest) SetDefaultAssignee(v string)`

SetDefaultAssignee sets DefaultAssignee field to given value.

### HasDefaultAssignee

`func (o *PatchedJIRAProjectRequest) HasDefaultAssignee() bool`

HasDefaultAssignee returns a boolean if a field has been set.

### SetDefaultAssigneeNil

`func (o *PatchedJIRAProjectRequest) SetDefaultAssigneeNil(b bool)`

 SetDefaultAssigneeNil sets the value for DefaultAssignee to be an explicit nil

### UnsetDefaultAssignee
`func (o *PatchedJIRAProjectRequest) UnsetDefaultAssignee()`

UnsetDefaultAssignee ensures that no value is present for DefaultAssignee, not even an explicit nil
### GetJiraLabels

`func (o *PatchedJIRAProjectRequest) GetJiraLabels() string`

GetJiraLabels returns the JiraLabels field if non-nil, zero value otherwise.

### GetJiraLabelsOk

`func (o *PatchedJIRAProjectRequest) GetJiraLabelsOk() (*string, bool)`

GetJiraLabelsOk returns a tuple with the JiraLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraLabels

`func (o *PatchedJIRAProjectRequest) SetJiraLabels(v string)`

SetJiraLabels sets JiraLabels field to given value.

### HasJiraLabels

`func (o *PatchedJIRAProjectRequest) HasJiraLabels() bool`

HasJiraLabels returns a boolean if a field has been set.

### SetJiraLabelsNil

`func (o *PatchedJIRAProjectRequest) SetJiraLabelsNil(b bool)`

 SetJiraLabelsNil sets the value for JiraLabels to be an explicit nil

### UnsetJiraLabels
`func (o *PatchedJIRAProjectRequest) UnsetJiraLabels()`

UnsetJiraLabels ensures that no value is present for JiraLabels, not even an explicit nil
### GetAddVulnerabilityIdToJiraLabel

`func (o *PatchedJIRAProjectRequest) GetAddVulnerabilityIdToJiraLabel() bool`

GetAddVulnerabilityIdToJiraLabel returns the AddVulnerabilityIdToJiraLabel field if non-nil, zero value otherwise.

### GetAddVulnerabilityIdToJiraLabelOk

`func (o *PatchedJIRAProjectRequest) GetAddVulnerabilityIdToJiraLabelOk() (*bool, bool)`

GetAddVulnerabilityIdToJiraLabelOk returns a tuple with the AddVulnerabilityIdToJiraLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVulnerabilityIdToJiraLabel

`func (o *PatchedJIRAProjectRequest) SetAddVulnerabilityIdToJiraLabel(v bool)`

SetAddVulnerabilityIdToJiraLabel sets AddVulnerabilityIdToJiraLabel field to given value.

### HasAddVulnerabilityIdToJiraLabel

`func (o *PatchedJIRAProjectRequest) HasAddVulnerabilityIdToJiraLabel() bool`

HasAddVulnerabilityIdToJiraLabel returns a boolean if a field has been set.

### GetPushAllIssues

`func (o *PatchedJIRAProjectRequest) GetPushAllIssues() bool`

GetPushAllIssues returns the PushAllIssues field if non-nil, zero value otherwise.

### GetPushAllIssuesOk

`func (o *PatchedJIRAProjectRequest) GetPushAllIssuesOk() (*bool, bool)`

GetPushAllIssuesOk returns a tuple with the PushAllIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushAllIssues

`func (o *PatchedJIRAProjectRequest) SetPushAllIssues(v bool)`

SetPushAllIssues sets PushAllIssues field to given value.

### HasPushAllIssues

`func (o *PatchedJIRAProjectRequest) HasPushAllIssues() bool`

HasPushAllIssues returns a boolean if a field has been set.

### GetEnableEngagementEpicMapping

`func (o *PatchedJIRAProjectRequest) GetEnableEngagementEpicMapping() bool`

GetEnableEngagementEpicMapping returns the EnableEngagementEpicMapping field if non-nil, zero value otherwise.

### GetEnableEngagementEpicMappingOk

`func (o *PatchedJIRAProjectRequest) GetEnableEngagementEpicMappingOk() (*bool, bool)`

GetEnableEngagementEpicMappingOk returns a tuple with the EnableEngagementEpicMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEngagementEpicMapping

`func (o *PatchedJIRAProjectRequest) SetEnableEngagementEpicMapping(v bool)`

SetEnableEngagementEpicMapping sets EnableEngagementEpicMapping field to given value.

### HasEnableEngagementEpicMapping

`func (o *PatchedJIRAProjectRequest) HasEnableEngagementEpicMapping() bool`

HasEnableEngagementEpicMapping returns a boolean if a field has been set.

### GetEpicIssueTypeName

`func (o *PatchedJIRAProjectRequest) GetEpicIssueTypeName() string`

GetEpicIssueTypeName returns the EpicIssueTypeName field if non-nil, zero value otherwise.

### GetEpicIssueTypeNameOk

`func (o *PatchedJIRAProjectRequest) GetEpicIssueTypeNameOk() (*string, bool)`

GetEpicIssueTypeNameOk returns a tuple with the EpicIssueTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicIssueTypeName

`func (o *PatchedJIRAProjectRequest) SetEpicIssueTypeName(v string)`

SetEpicIssueTypeName sets EpicIssueTypeName field to given value.

### HasEpicIssueTypeName

`func (o *PatchedJIRAProjectRequest) HasEpicIssueTypeName() bool`

HasEpicIssueTypeName returns a boolean if a field has been set.

### GetPushNotes

`func (o *PatchedJIRAProjectRequest) GetPushNotes() bool`

GetPushNotes returns the PushNotes field if non-nil, zero value otherwise.

### GetPushNotesOk

`func (o *PatchedJIRAProjectRequest) GetPushNotesOk() (*bool, bool)`

GetPushNotesOk returns a tuple with the PushNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushNotes

`func (o *PatchedJIRAProjectRequest) SetPushNotes(v bool)`

SetPushNotes sets PushNotes field to given value.

### HasPushNotes

`func (o *PatchedJIRAProjectRequest) HasPushNotes() bool`

HasPushNotes returns a boolean if a field has been set.

### GetProductJiraSlaNotification

`func (o *PatchedJIRAProjectRequest) GetProductJiraSlaNotification() bool`

GetProductJiraSlaNotification returns the ProductJiraSlaNotification field if non-nil, zero value otherwise.

### GetProductJiraSlaNotificationOk

`func (o *PatchedJIRAProjectRequest) GetProductJiraSlaNotificationOk() (*bool, bool)`

GetProductJiraSlaNotificationOk returns a tuple with the ProductJiraSlaNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductJiraSlaNotification

`func (o *PatchedJIRAProjectRequest) SetProductJiraSlaNotification(v bool)`

SetProductJiraSlaNotification sets ProductJiraSlaNotification field to given value.

### HasProductJiraSlaNotification

`func (o *PatchedJIRAProjectRequest) HasProductJiraSlaNotification() bool`

HasProductJiraSlaNotification returns a boolean if a field has been set.

### GetRiskAcceptanceExpirationNotification

`func (o *PatchedJIRAProjectRequest) GetRiskAcceptanceExpirationNotification() bool`

GetRiskAcceptanceExpirationNotification returns the RiskAcceptanceExpirationNotification field if non-nil, zero value otherwise.

### GetRiskAcceptanceExpirationNotificationOk

`func (o *PatchedJIRAProjectRequest) GetRiskAcceptanceExpirationNotificationOk() (*bool, bool)`

GetRiskAcceptanceExpirationNotificationOk returns a tuple with the RiskAcceptanceExpirationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceExpirationNotification

`func (o *PatchedJIRAProjectRequest) SetRiskAcceptanceExpirationNotification(v bool)`

SetRiskAcceptanceExpirationNotification sets RiskAcceptanceExpirationNotification field to given value.

### HasRiskAcceptanceExpirationNotification

`func (o *PatchedJIRAProjectRequest) HasRiskAcceptanceExpirationNotification() bool`

HasRiskAcceptanceExpirationNotification returns a boolean if a field has been set.

### GetJiraInstance

`func (o *PatchedJIRAProjectRequest) GetJiraInstance() int32`

GetJiraInstance returns the JiraInstance field if non-nil, zero value otherwise.

### GetJiraInstanceOk

`func (o *PatchedJIRAProjectRequest) GetJiraInstanceOk() (*int32, bool)`

GetJiraInstanceOk returns a tuple with the JiraInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraInstance

`func (o *PatchedJIRAProjectRequest) SetJiraInstance(v int32)`

SetJiraInstance sets JiraInstance field to given value.

### HasJiraInstance

`func (o *PatchedJIRAProjectRequest) HasJiraInstance() bool`

HasJiraInstance returns a boolean if a field has been set.

### SetJiraInstanceNil

`func (o *PatchedJIRAProjectRequest) SetJiraInstanceNil(b bool)`

 SetJiraInstanceNil sets the value for JiraInstance to be an explicit nil

### UnsetJiraInstance
`func (o *PatchedJIRAProjectRequest) UnsetJiraInstance()`

UnsetJiraInstance ensures that no value is present for JiraInstance, not even an explicit nil
### GetProduct

`func (o *PatchedJIRAProjectRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedJIRAProjectRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedJIRAProjectRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedJIRAProjectRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *PatchedJIRAProjectRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *PatchedJIRAProjectRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetEngagement

`func (o *PatchedJIRAProjectRequest) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *PatchedJIRAProjectRequest) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *PatchedJIRAProjectRequest) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *PatchedJIRAProjectRequest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *PatchedJIRAProjectRequest) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *PatchedJIRAProjectRequest) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


