# FindingGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Test** | **int32** |  | 
**JiraIssue** | [**NullableJIRAIssue**](JIRAIssue.md) |  | [readonly] 

## Methods

### NewFindingGroup

`func NewFindingGroup(id int32, name string, test int32, jiraIssue NullableJIRAIssue, ) *FindingGroup`

NewFindingGroup instantiates a new FindingGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingGroupWithDefaults

`func NewFindingGroupWithDefaults() *FindingGroup`

NewFindingGroupWithDefaults instantiates a new FindingGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FindingGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FindingGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindingGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindingGroup) SetName(v string)`

SetName sets Name field to given value.


### GetTest

`func (o *FindingGroup) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *FindingGroup) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *FindingGroup) SetTest(v int32)`

SetTest sets Test field to given value.


### GetJiraIssue

`func (o *FindingGroup) GetJiraIssue() JIRAIssue`

GetJiraIssue returns the JiraIssue field if non-nil, zero value otherwise.

### GetJiraIssueOk

`func (o *FindingGroup) GetJiraIssueOk() (*JIRAIssue, bool)`

GetJiraIssueOk returns a tuple with the JiraIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraIssue

`func (o *FindingGroup) SetJiraIssue(v JIRAIssue)`

SetJiraIssue sets JiraIssue field to given value.


### SetJiraIssueNil

`func (o *FindingGroup) SetJiraIssueNil(b bool)`

 SetJiraIssueNil sets the value for JiraIssue to be an explicit nil

### UnsetJiraIssue
`func (o *FindingGroup) UnsetJiraIssue()`

UnsetJiraIssue ensures that no value is present for JiraIssue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


