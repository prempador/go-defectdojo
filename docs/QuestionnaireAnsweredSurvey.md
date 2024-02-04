# QuestionnaireAnsweredSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Completed** | Pointer to **bool** |  | [optional] 
**AnsweredOn** | Pointer to **NullableString** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**Survey** | **int32** |  | 
**Assignee** | Pointer to **NullableInt32** |  | [optional] 
**Responder** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewQuestionnaireAnsweredSurvey

`func NewQuestionnaireAnsweredSurvey(id int32, survey int32, ) *QuestionnaireAnsweredSurvey`

NewQuestionnaireAnsweredSurvey instantiates a new QuestionnaireAnsweredSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireAnsweredSurveyWithDefaults

`func NewQuestionnaireAnsweredSurveyWithDefaults() *QuestionnaireAnsweredSurvey`

NewQuestionnaireAnsweredSurveyWithDefaults instantiates a new QuestionnaireAnsweredSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireAnsweredSurvey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireAnsweredSurvey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireAnsweredSurvey) SetId(v int32)`

SetId sets Id field to given value.


### GetCompleted

`func (o *QuestionnaireAnsweredSurvey) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *QuestionnaireAnsweredSurvey) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *QuestionnaireAnsweredSurvey) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *QuestionnaireAnsweredSurvey) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetAnsweredOn

`func (o *QuestionnaireAnsweredSurvey) GetAnsweredOn() string`

GetAnsweredOn returns the AnsweredOn field if non-nil, zero value otherwise.

### GetAnsweredOnOk

`func (o *QuestionnaireAnsweredSurvey) GetAnsweredOnOk() (*string, bool)`

GetAnsweredOnOk returns a tuple with the AnsweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredOn

`func (o *QuestionnaireAnsweredSurvey) SetAnsweredOn(v string)`

SetAnsweredOn sets AnsweredOn field to given value.

### HasAnsweredOn

`func (o *QuestionnaireAnsweredSurvey) HasAnsweredOn() bool`

HasAnsweredOn returns a boolean if a field has been set.

### SetAnsweredOnNil

`func (o *QuestionnaireAnsweredSurvey) SetAnsweredOnNil(b bool)`

 SetAnsweredOnNil sets the value for AnsweredOn to be an explicit nil

### UnsetAnsweredOn
`func (o *QuestionnaireAnsweredSurvey) UnsetAnsweredOn()`

UnsetAnsweredOn ensures that no value is present for AnsweredOn, not even an explicit nil
### GetEngagement

`func (o *QuestionnaireAnsweredSurvey) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *QuestionnaireAnsweredSurvey) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *QuestionnaireAnsweredSurvey) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *QuestionnaireAnsweredSurvey) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *QuestionnaireAnsweredSurvey) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *QuestionnaireAnsweredSurvey) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetSurvey

`func (o *QuestionnaireAnsweredSurvey) GetSurvey() int32`

GetSurvey returns the Survey field if non-nil, zero value otherwise.

### GetSurveyOk

`func (o *QuestionnaireAnsweredSurvey) GetSurveyOk() (*int32, bool)`

GetSurveyOk returns a tuple with the Survey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvey

`func (o *QuestionnaireAnsweredSurvey) SetSurvey(v int32)`

SetSurvey sets Survey field to given value.


### GetAssignee

`func (o *QuestionnaireAnsweredSurvey) GetAssignee() int32`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *QuestionnaireAnsweredSurvey) GetAssigneeOk() (*int32, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *QuestionnaireAnsweredSurvey) SetAssignee(v int32)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *QuestionnaireAnsweredSurvey) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *QuestionnaireAnsweredSurvey) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *QuestionnaireAnsweredSurvey) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetResponder

`func (o *QuestionnaireAnsweredSurvey) GetResponder() int32`

GetResponder returns the Responder field if non-nil, zero value otherwise.

### GetResponderOk

`func (o *QuestionnaireAnsweredSurvey) GetResponderOk() (*int32, bool)`

GetResponderOk returns a tuple with the Responder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponder

`func (o *QuestionnaireAnsweredSurvey) SetResponder(v int32)`

SetResponder sets Responder field to given value.

### HasResponder

`func (o *QuestionnaireAnsweredSurvey) HasResponder() bool`

HasResponder returns a boolean if a field has been set.

### SetResponderNil

`func (o *QuestionnaireAnsweredSurvey) SetResponderNil(b bool)`

 SetResponderNil sets the value for Responder to be an explicit nil

### UnsetResponder
`func (o *QuestionnaireAnsweredSurvey) UnsetResponder()`

UnsetResponder ensures that no value is present for Responder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


