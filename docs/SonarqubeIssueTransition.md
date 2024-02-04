# SonarqubeIssueTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**FindingStatus** | **string** |  | 
**SonarqubeStatus** | **string** |  | 
**Transitions** | **string** |  | 
**SonarqubeIssue** | **int32** |  | 

## Methods

### NewSonarqubeIssueTransition

`func NewSonarqubeIssueTransition(id int32, created time.Time, findingStatus string, sonarqubeStatus string, transitions string, sonarqubeIssue int32, ) *SonarqubeIssueTransition`

NewSonarqubeIssueTransition instantiates a new SonarqubeIssueTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSonarqubeIssueTransitionWithDefaults

`func NewSonarqubeIssueTransitionWithDefaults() *SonarqubeIssueTransition`

NewSonarqubeIssueTransitionWithDefaults instantiates a new SonarqubeIssueTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SonarqubeIssueTransition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SonarqubeIssueTransition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SonarqubeIssueTransition) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *SonarqubeIssueTransition) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SonarqubeIssueTransition) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SonarqubeIssueTransition) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetFindingStatus

`func (o *SonarqubeIssueTransition) GetFindingStatus() string`

GetFindingStatus returns the FindingStatus field if non-nil, zero value otherwise.

### GetFindingStatusOk

`func (o *SonarqubeIssueTransition) GetFindingStatusOk() (*string, bool)`

GetFindingStatusOk returns a tuple with the FindingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingStatus

`func (o *SonarqubeIssueTransition) SetFindingStatus(v string)`

SetFindingStatus sets FindingStatus field to given value.


### GetSonarqubeStatus

`func (o *SonarqubeIssueTransition) GetSonarqubeStatus() string`

GetSonarqubeStatus returns the SonarqubeStatus field if non-nil, zero value otherwise.

### GetSonarqubeStatusOk

`func (o *SonarqubeIssueTransition) GetSonarqubeStatusOk() (*string, bool)`

GetSonarqubeStatusOk returns a tuple with the SonarqubeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeStatus

`func (o *SonarqubeIssueTransition) SetSonarqubeStatus(v string)`

SetSonarqubeStatus sets SonarqubeStatus field to given value.


### GetTransitions

`func (o *SonarqubeIssueTransition) GetTransitions() string`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *SonarqubeIssueTransition) GetTransitionsOk() (*string, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *SonarqubeIssueTransition) SetTransitions(v string)`

SetTransitions sets Transitions field to given value.


### GetSonarqubeIssue

`func (o *SonarqubeIssueTransition) GetSonarqubeIssue() int32`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *SonarqubeIssueTransition) GetSonarqubeIssueOk() (*int32, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *SonarqubeIssueTransition) SetSonarqubeIssue(v int32)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


