# JIRAIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**JiraId** | **string** |  | 
**JiraKey** | **string** |  | 
**JiraCreation** | Pointer to **NullableTime** | The date a Jira issue was created from this finding. | [optional] 
**JiraChange** | Pointer to **NullableTime** | The date the linked Jira issue was last modified. | [optional] 
**JiraProject** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**FindingGroup** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewJIRAIssue

`func NewJIRAIssue(id int32, url string, jiraId string, jiraKey string, ) *JIRAIssue`

NewJIRAIssue instantiates a new JIRAIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJIRAIssueWithDefaults

`func NewJIRAIssueWithDefaults() *JIRAIssue`

NewJIRAIssueWithDefaults instantiates a new JIRAIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JIRAIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JIRAIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JIRAIssue) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *JIRAIssue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JIRAIssue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JIRAIssue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetJiraId

`func (o *JIRAIssue) GetJiraId() string`

GetJiraId returns the JiraId field if non-nil, zero value otherwise.

### GetJiraIdOk

`func (o *JIRAIssue) GetJiraIdOk() (*string, bool)`

GetJiraIdOk returns a tuple with the JiraId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraId

`func (o *JIRAIssue) SetJiraId(v string)`

SetJiraId sets JiraId field to given value.


### GetJiraKey

`func (o *JIRAIssue) GetJiraKey() string`

GetJiraKey returns the JiraKey field if non-nil, zero value otherwise.

### GetJiraKeyOk

`func (o *JIRAIssue) GetJiraKeyOk() (*string, bool)`

GetJiraKeyOk returns a tuple with the JiraKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraKey

`func (o *JIRAIssue) SetJiraKey(v string)`

SetJiraKey sets JiraKey field to given value.


### GetJiraCreation

`func (o *JIRAIssue) GetJiraCreation() time.Time`

GetJiraCreation returns the JiraCreation field if non-nil, zero value otherwise.

### GetJiraCreationOk

`func (o *JIRAIssue) GetJiraCreationOk() (*time.Time, bool)`

GetJiraCreationOk returns a tuple with the JiraCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraCreation

`func (o *JIRAIssue) SetJiraCreation(v time.Time)`

SetJiraCreation sets JiraCreation field to given value.

### HasJiraCreation

`func (o *JIRAIssue) HasJiraCreation() bool`

HasJiraCreation returns a boolean if a field has been set.

### SetJiraCreationNil

`func (o *JIRAIssue) SetJiraCreationNil(b bool)`

 SetJiraCreationNil sets the value for JiraCreation to be an explicit nil

### UnsetJiraCreation
`func (o *JIRAIssue) UnsetJiraCreation()`

UnsetJiraCreation ensures that no value is present for JiraCreation, not even an explicit nil
### GetJiraChange

`func (o *JIRAIssue) GetJiraChange() time.Time`

GetJiraChange returns the JiraChange field if non-nil, zero value otherwise.

### GetJiraChangeOk

`func (o *JIRAIssue) GetJiraChangeOk() (*time.Time, bool)`

GetJiraChangeOk returns a tuple with the JiraChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraChange

`func (o *JIRAIssue) SetJiraChange(v time.Time)`

SetJiraChange sets JiraChange field to given value.

### HasJiraChange

`func (o *JIRAIssue) HasJiraChange() bool`

HasJiraChange returns a boolean if a field has been set.

### SetJiraChangeNil

`func (o *JIRAIssue) SetJiraChangeNil(b bool)`

 SetJiraChangeNil sets the value for JiraChange to be an explicit nil

### UnsetJiraChange
`func (o *JIRAIssue) UnsetJiraChange()`

UnsetJiraChange ensures that no value is present for JiraChange, not even an explicit nil
### GetJiraProject

`func (o *JIRAIssue) GetJiraProject() int32`

GetJiraProject returns the JiraProject field if non-nil, zero value otherwise.

### GetJiraProjectOk

`func (o *JIRAIssue) GetJiraProjectOk() (*int32, bool)`

GetJiraProjectOk returns a tuple with the JiraProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraProject

`func (o *JIRAIssue) SetJiraProject(v int32)`

SetJiraProject sets JiraProject field to given value.

### HasJiraProject

`func (o *JIRAIssue) HasJiraProject() bool`

HasJiraProject returns a boolean if a field has been set.

### SetJiraProjectNil

`func (o *JIRAIssue) SetJiraProjectNil(b bool)`

 SetJiraProjectNil sets the value for JiraProject to be an explicit nil

### UnsetJiraProject
`func (o *JIRAIssue) UnsetJiraProject()`

UnsetJiraProject ensures that no value is present for JiraProject, not even an explicit nil
### GetFinding

`func (o *JIRAIssue) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *JIRAIssue) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *JIRAIssue) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *JIRAIssue) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *JIRAIssue) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *JIRAIssue) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetEngagement

`func (o *JIRAIssue) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *JIRAIssue) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *JIRAIssue) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *JIRAIssue) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *JIRAIssue) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *JIRAIssue) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetFindingGroup

`func (o *JIRAIssue) GetFindingGroup() int32`

GetFindingGroup returns the FindingGroup field if non-nil, zero value otherwise.

### GetFindingGroupOk

`func (o *JIRAIssue) GetFindingGroupOk() (*int32, bool)`

GetFindingGroupOk returns a tuple with the FindingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingGroup

`func (o *JIRAIssue) SetFindingGroup(v int32)`

SetFindingGroup sets FindingGroup field to given value.

### HasFindingGroup

`func (o *JIRAIssue) HasFindingGroup() bool`

HasFindingGroup returns a boolean if a field has been set.

### SetFindingGroupNil

`func (o *JIRAIssue) SetFindingGroupNil(b bool)`

 SetFindingGroupNil sets the value for FindingGroup to be an explicit nil

### UnsetFindingGroup
`func (o *JIRAIssue) UnsetFindingGroup()`

UnsetFindingGroup ensures that no value is present for FindingGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


