# PaginatedQuestionnaireAnsweredSurveyListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**Engagement** | Pointer to [**map[string]FindingEngagement**](FindingEngagement.md) |  | [optional] [readonly] 
**Responder** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**Survey** | Pointer to [**map[string]QuestionnaireEngagementSurvey**](QuestionnaireEngagementSurvey.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedQuestionnaireAnsweredSurveyListPrefetch

`func NewPaginatedQuestionnaireAnsweredSurveyListPrefetch() *PaginatedQuestionnaireAnsweredSurveyListPrefetch`

NewPaginatedQuestionnaireAnsweredSurveyListPrefetch instantiates a new PaginatedQuestionnaireAnsweredSurveyListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedQuestionnaireAnsweredSurveyListPrefetchWithDefaults

`func NewPaginatedQuestionnaireAnsweredSurveyListPrefetchWithDefaults() *PaginatedQuestionnaireAnsweredSurveyListPrefetch`

NewPaginatedQuestionnaireAnsweredSurveyListPrefetchWithDefaults instantiates a new PaginatedQuestionnaireAnsweredSurveyListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetAssignee() map[string]UserStub`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetAssigneeOk() (*map[string]UserStub, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetAssignee(v map[string]UserStub)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetEngagement

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetEngagement() map[string]FindingEngagement`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetEngagementOk() (*map[string]FindingEngagement, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetEngagement(v map[string]FindingEngagement)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetResponder

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetResponder() map[string]UserStub`

GetResponder returns the Responder field if non-nil, zero value otherwise.

### GetResponderOk

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetResponderOk() (*map[string]UserStub, bool)`

GetResponderOk returns a tuple with the Responder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponder

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetResponder(v map[string]UserStub)`

SetResponder sets Responder field to given value.

### HasResponder

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasResponder() bool`

HasResponder returns a boolean if a field has been set.

### GetSurvey

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetSurvey() map[string]QuestionnaireEngagementSurvey`

GetSurvey returns the Survey field if non-nil, zero value otherwise.

### GetSurveyOk

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) GetSurveyOk() (*map[string]QuestionnaireEngagementSurvey, bool)`

GetSurveyOk returns a tuple with the Survey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvey

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) SetSurvey(v map[string]QuestionnaireEngagementSurvey)`

SetSurvey sets Survey field to given value.

### HasSurvey

`func (o *PaginatedQuestionnaireAnsweredSurveyListPrefetch) HasSurvey() bool`

HasSurvey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


