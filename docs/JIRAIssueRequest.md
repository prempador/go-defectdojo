# JIRAIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JiraId** | **string** |  | 
**JiraKey** | **string** |  | 
**JiraCreation** | Pointer to **NullableTime** | The date a Jira issue was created from this finding. | [optional] 
**JiraChange** | Pointer to **NullableTime** | The date the linked Jira issue was last modified. | [optional] 
**JiraProject** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**FindingGroup** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewJIRAIssueRequest

`func NewJIRAIssueRequest(jiraId string, jiraKey string, ) *JIRAIssueRequest`

NewJIRAIssueRequest instantiates a new JIRAIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJIRAIssueRequestWithDefaults

`func NewJIRAIssueRequestWithDefaults() *JIRAIssueRequest`

NewJIRAIssueRequestWithDefaults instantiates a new JIRAIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJiraId

`func (o *JIRAIssueRequest) GetJiraId() string`

GetJiraId returns the JiraId field if non-nil, zero value otherwise.

### GetJiraIdOk

`func (o *JIRAIssueRequest) GetJiraIdOk() (*string, bool)`

GetJiraIdOk returns a tuple with the JiraId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraId

`func (o *JIRAIssueRequest) SetJiraId(v string)`

SetJiraId sets JiraId field to given value.


### GetJiraKey

`func (o *JIRAIssueRequest) GetJiraKey() string`

GetJiraKey returns the JiraKey field if non-nil, zero value otherwise.

### GetJiraKeyOk

`func (o *JIRAIssueRequest) GetJiraKeyOk() (*string, bool)`

GetJiraKeyOk returns a tuple with the JiraKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraKey

`func (o *JIRAIssueRequest) SetJiraKey(v string)`

SetJiraKey sets JiraKey field to given value.


### GetJiraCreation

`func (o *JIRAIssueRequest) GetJiraCreation() time.Time`

GetJiraCreation returns the JiraCreation field if non-nil, zero value otherwise.

### GetJiraCreationOk

`func (o *JIRAIssueRequest) GetJiraCreationOk() (*time.Time, bool)`

GetJiraCreationOk returns a tuple with the JiraCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraCreation

`func (o *JIRAIssueRequest) SetJiraCreation(v time.Time)`

SetJiraCreation sets JiraCreation field to given value.

### HasJiraCreation

`func (o *JIRAIssueRequest) HasJiraCreation() bool`

HasJiraCreation returns a boolean if a field has been set.

### SetJiraCreationNil

`func (o *JIRAIssueRequest) SetJiraCreationNil(b bool)`

 SetJiraCreationNil sets the value for JiraCreation to be an explicit nil

### UnsetJiraCreation
`func (o *JIRAIssueRequest) UnsetJiraCreation()`

UnsetJiraCreation ensures that no value is present for JiraCreation, not even an explicit nil
### GetJiraChange

`func (o *JIRAIssueRequest) GetJiraChange() time.Time`

GetJiraChange returns the JiraChange field if non-nil, zero value otherwise.

### GetJiraChangeOk

`func (o *JIRAIssueRequest) GetJiraChangeOk() (*time.Time, bool)`

GetJiraChangeOk returns a tuple with the JiraChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraChange

`func (o *JIRAIssueRequest) SetJiraChange(v time.Time)`

SetJiraChange sets JiraChange field to given value.

### HasJiraChange

`func (o *JIRAIssueRequest) HasJiraChange() bool`

HasJiraChange returns a boolean if a field has been set.

### SetJiraChangeNil

`func (o *JIRAIssueRequest) SetJiraChangeNil(b bool)`

 SetJiraChangeNil sets the value for JiraChange to be an explicit nil

### UnsetJiraChange
`func (o *JIRAIssueRequest) UnsetJiraChange()`

UnsetJiraChange ensures that no value is present for JiraChange, not even an explicit nil
### GetJiraProject

`func (o *JIRAIssueRequest) GetJiraProject() int32`

GetJiraProject returns the JiraProject field if non-nil, zero value otherwise.

### GetJiraProjectOk

`func (o *JIRAIssueRequest) GetJiraProjectOk() (*int32, bool)`

GetJiraProjectOk returns a tuple with the JiraProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraProject

`func (o *JIRAIssueRequest) SetJiraProject(v int32)`

SetJiraProject sets JiraProject field to given value.

### HasJiraProject

`func (o *JIRAIssueRequest) HasJiraProject() bool`

HasJiraProject returns a boolean if a field has been set.

### SetJiraProjectNil

`func (o *JIRAIssueRequest) SetJiraProjectNil(b bool)`

 SetJiraProjectNil sets the value for JiraProject to be an explicit nil

### UnsetJiraProject
`func (o *JIRAIssueRequest) UnsetJiraProject()`

UnsetJiraProject ensures that no value is present for JiraProject, not even an explicit nil
### GetFinding

`func (o *JIRAIssueRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *JIRAIssueRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *JIRAIssueRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *JIRAIssueRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *JIRAIssueRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *JIRAIssueRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetEngagement

`func (o *JIRAIssueRequest) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *JIRAIssueRequest) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *JIRAIssueRequest) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *JIRAIssueRequest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *JIRAIssueRequest) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *JIRAIssueRequest) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetFindingGroup

`func (o *JIRAIssueRequest) GetFindingGroup() int32`

GetFindingGroup returns the FindingGroup field if non-nil, zero value otherwise.

### GetFindingGroupOk

`func (o *JIRAIssueRequest) GetFindingGroupOk() (*int32, bool)`

GetFindingGroupOk returns a tuple with the FindingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingGroup

`func (o *JIRAIssueRequest) SetFindingGroup(v int32)`

SetFindingGroup sets FindingGroup field to given value.

### HasFindingGroup

`func (o *JIRAIssueRequest) HasFindingGroup() bool`

HasFindingGroup returns a boolean if a field has been set.

### SetFindingGroupNil

`func (o *JIRAIssueRequest) SetFindingGroupNil(b bool)`

 SetFindingGroupNil sets the value for FindingGroup to be an explicit nil

### UnsetFindingGroup
`func (o *JIRAIssueRequest) UnsetFindingGroup()`

UnsetFindingGroup ensures that no value is present for FindingGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


