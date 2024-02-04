# PatchedJIRAIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JiraId** | Pointer to **string** |  | [optional] 
**JiraKey** | Pointer to **string** |  | [optional] 
**JiraCreation** | Pointer to **NullableTime** | The date a Jira issue was created from this finding. | [optional] 
**JiraChange** | Pointer to **NullableTime** | The date the linked Jira issue was last modified. | [optional] 
**JiraProject** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**FindingGroup** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPatchedJIRAIssueRequest

`func NewPatchedJIRAIssueRequest() *PatchedJIRAIssueRequest`

NewPatchedJIRAIssueRequest instantiates a new PatchedJIRAIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedJIRAIssueRequestWithDefaults

`func NewPatchedJIRAIssueRequestWithDefaults() *PatchedJIRAIssueRequest`

NewPatchedJIRAIssueRequestWithDefaults instantiates a new PatchedJIRAIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJiraId

`func (o *PatchedJIRAIssueRequest) GetJiraId() string`

GetJiraId returns the JiraId field if non-nil, zero value otherwise.

### GetJiraIdOk

`func (o *PatchedJIRAIssueRequest) GetJiraIdOk() (*string, bool)`

GetJiraIdOk returns a tuple with the JiraId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraId

`func (o *PatchedJIRAIssueRequest) SetJiraId(v string)`

SetJiraId sets JiraId field to given value.

### HasJiraId

`func (o *PatchedJIRAIssueRequest) HasJiraId() bool`

HasJiraId returns a boolean if a field has been set.

### GetJiraKey

`func (o *PatchedJIRAIssueRequest) GetJiraKey() string`

GetJiraKey returns the JiraKey field if non-nil, zero value otherwise.

### GetJiraKeyOk

`func (o *PatchedJIRAIssueRequest) GetJiraKeyOk() (*string, bool)`

GetJiraKeyOk returns a tuple with the JiraKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraKey

`func (o *PatchedJIRAIssueRequest) SetJiraKey(v string)`

SetJiraKey sets JiraKey field to given value.

### HasJiraKey

`func (o *PatchedJIRAIssueRequest) HasJiraKey() bool`

HasJiraKey returns a boolean if a field has been set.

### GetJiraCreation

`func (o *PatchedJIRAIssueRequest) GetJiraCreation() time.Time`

GetJiraCreation returns the JiraCreation field if non-nil, zero value otherwise.

### GetJiraCreationOk

`func (o *PatchedJIRAIssueRequest) GetJiraCreationOk() (*time.Time, bool)`

GetJiraCreationOk returns a tuple with the JiraCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraCreation

`func (o *PatchedJIRAIssueRequest) SetJiraCreation(v time.Time)`

SetJiraCreation sets JiraCreation field to given value.

### HasJiraCreation

`func (o *PatchedJIRAIssueRequest) HasJiraCreation() bool`

HasJiraCreation returns a boolean if a field has been set.

### SetJiraCreationNil

`func (o *PatchedJIRAIssueRequest) SetJiraCreationNil(b bool)`

 SetJiraCreationNil sets the value for JiraCreation to be an explicit nil

### UnsetJiraCreation
`func (o *PatchedJIRAIssueRequest) UnsetJiraCreation()`

UnsetJiraCreation ensures that no value is present for JiraCreation, not even an explicit nil
### GetJiraChange

`func (o *PatchedJIRAIssueRequest) GetJiraChange() time.Time`

GetJiraChange returns the JiraChange field if non-nil, zero value otherwise.

### GetJiraChangeOk

`func (o *PatchedJIRAIssueRequest) GetJiraChangeOk() (*time.Time, bool)`

GetJiraChangeOk returns a tuple with the JiraChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraChange

`func (o *PatchedJIRAIssueRequest) SetJiraChange(v time.Time)`

SetJiraChange sets JiraChange field to given value.

### HasJiraChange

`func (o *PatchedJIRAIssueRequest) HasJiraChange() bool`

HasJiraChange returns a boolean if a field has been set.

### SetJiraChangeNil

`func (o *PatchedJIRAIssueRequest) SetJiraChangeNil(b bool)`

 SetJiraChangeNil sets the value for JiraChange to be an explicit nil

### UnsetJiraChange
`func (o *PatchedJIRAIssueRequest) UnsetJiraChange()`

UnsetJiraChange ensures that no value is present for JiraChange, not even an explicit nil
### GetJiraProject

`func (o *PatchedJIRAIssueRequest) GetJiraProject() int32`

GetJiraProject returns the JiraProject field if non-nil, zero value otherwise.

### GetJiraProjectOk

`func (o *PatchedJIRAIssueRequest) GetJiraProjectOk() (*int32, bool)`

GetJiraProjectOk returns a tuple with the JiraProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraProject

`func (o *PatchedJIRAIssueRequest) SetJiraProject(v int32)`

SetJiraProject sets JiraProject field to given value.

### HasJiraProject

`func (o *PatchedJIRAIssueRequest) HasJiraProject() bool`

HasJiraProject returns a boolean if a field has been set.

### SetJiraProjectNil

`func (o *PatchedJIRAIssueRequest) SetJiraProjectNil(b bool)`

 SetJiraProjectNil sets the value for JiraProject to be an explicit nil

### UnsetJiraProject
`func (o *PatchedJIRAIssueRequest) UnsetJiraProject()`

UnsetJiraProject ensures that no value is present for JiraProject, not even an explicit nil
### GetFinding

`func (o *PatchedJIRAIssueRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *PatchedJIRAIssueRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *PatchedJIRAIssueRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *PatchedJIRAIssueRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *PatchedJIRAIssueRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *PatchedJIRAIssueRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetEngagement

`func (o *PatchedJIRAIssueRequest) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *PatchedJIRAIssueRequest) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *PatchedJIRAIssueRequest) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *PatchedJIRAIssueRequest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *PatchedJIRAIssueRequest) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *PatchedJIRAIssueRequest) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetFindingGroup

`func (o *PatchedJIRAIssueRequest) GetFindingGroup() int32`

GetFindingGroup returns the FindingGroup field if non-nil, zero value otherwise.

### GetFindingGroupOk

`func (o *PatchedJIRAIssueRequest) GetFindingGroupOk() (*int32, bool)`

GetFindingGroupOk returns a tuple with the FindingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingGroup

`func (o *PatchedJIRAIssueRequest) SetFindingGroup(v int32)`

SetFindingGroup sets FindingGroup field to given value.

### HasFindingGroup

`func (o *PatchedJIRAIssueRequest) HasFindingGroup() bool`

HasFindingGroup returns a boolean if a field has been set.

### SetFindingGroupNil

`func (o *PatchedJIRAIssueRequest) SetFindingGroupNil(b bool)`

 SetFindingGroupNil sets the value for FindingGroup to be an explicit nil

### UnsetFindingGroup
`func (o *PatchedJIRAIssueRequest) UnsetFindingGroup()`

UnsetFindingGroup ensures that no value is present for FindingGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


