# SonarqubeIssueTransitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FindingStatus** | **string** |  | 
**SonarqubeStatus** | **string** |  | 
**Transitions** | **string** |  | 
**SonarqubeIssue** | **int32** |  | 

## Methods

### NewSonarqubeIssueTransitionRequest

`func NewSonarqubeIssueTransitionRequest(findingStatus string, sonarqubeStatus string, transitions string, sonarqubeIssue int32, ) *SonarqubeIssueTransitionRequest`

NewSonarqubeIssueTransitionRequest instantiates a new SonarqubeIssueTransitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSonarqubeIssueTransitionRequestWithDefaults

`func NewSonarqubeIssueTransitionRequestWithDefaults() *SonarqubeIssueTransitionRequest`

NewSonarqubeIssueTransitionRequestWithDefaults instantiates a new SonarqubeIssueTransitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindingStatus

`func (o *SonarqubeIssueTransitionRequest) GetFindingStatus() string`

GetFindingStatus returns the FindingStatus field if non-nil, zero value otherwise.

### GetFindingStatusOk

`func (o *SonarqubeIssueTransitionRequest) GetFindingStatusOk() (*string, bool)`

GetFindingStatusOk returns a tuple with the FindingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingStatus

`func (o *SonarqubeIssueTransitionRequest) SetFindingStatus(v string)`

SetFindingStatus sets FindingStatus field to given value.


### GetSonarqubeStatus

`func (o *SonarqubeIssueTransitionRequest) GetSonarqubeStatus() string`

GetSonarqubeStatus returns the SonarqubeStatus field if non-nil, zero value otherwise.

### GetSonarqubeStatusOk

`func (o *SonarqubeIssueTransitionRequest) GetSonarqubeStatusOk() (*string, bool)`

GetSonarqubeStatusOk returns a tuple with the SonarqubeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeStatus

`func (o *SonarqubeIssueTransitionRequest) SetSonarqubeStatus(v string)`

SetSonarqubeStatus sets SonarqubeStatus field to given value.


### GetTransitions

`func (o *SonarqubeIssueTransitionRequest) GetTransitions() string`

GetTransitions returns the Transitions field if non-nil, zero value otherwise.

### GetTransitionsOk

`func (o *SonarqubeIssueTransitionRequest) GetTransitionsOk() (*string, bool)`

GetTransitionsOk returns a tuple with the Transitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitions

`func (o *SonarqubeIssueTransitionRequest) SetTransitions(v string)`

SetTransitions sets Transitions field to given value.


### GetSonarqubeIssue

`func (o *SonarqubeIssueTransitionRequest) GetSonarqubeIssue() int32`

GetSonarqubeIssue returns the SonarqubeIssue field if non-nil, zero value otherwise.

### GetSonarqubeIssueOk

`func (o *SonarqubeIssueTransitionRequest) GetSonarqubeIssueOk() (*int32, bool)`

GetSonarqubeIssueOk returns a tuple with the SonarqubeIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonarqubeIssue

`func (o *SonarqubeIssueTransitionRequest) SetSonarqubeIssue(v int32)`

SetSonarqubeIssue sets SonarqubeIssue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


