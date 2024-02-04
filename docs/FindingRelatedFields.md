# FindingRelatedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | [**FindingTest**](FindingTest.md) |  | [readonly] 
**Jira** | [**JIRAIssue**](JIRAIssue.md) |  | [readonly] 

## Methods

### NewFindingRelatedFields

`func NewFindingRelatedFields(test FindingTest, jira JIRAIssue, ) *FindingRelatedFields`

NewFindingRelatedFields instantiates a new FindingRelatedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingRelatedFieldsWithDefaults

`func NewFindingRelatedFieldsWithDefaults() *FindingRelatedFields`

NewFindingRelatedFieldsWithDefaults instantiates a new FindingRelatedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTest

`func (o *FindingRelatedFields) GetTest() FindingTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *FindingRelatedFields) GetTestOk() (*FindingTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *FindingRelatedFields) SetTest(v FindingTest)`

SetTest sets Test field to given value.


### GetJira

`func (o *FindingRelatedFields) GetJira() JIRAIssue`

GetJira returns the Jira field if non-nil, zero value otherwise.

### GetJiraOk

`func (o *FindingRelatedFields) GetJiraOk() (*JIRAIssue, bool)`

GetJiraOk returns a tuple with the Jira field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJira

`func (o *FindingRelatedFields) SetJira(v JIRAIssue)`

SetJira sets Jira field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


