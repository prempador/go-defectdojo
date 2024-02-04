# QuestionnaireAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**Question** | **int32** |  | 
**AnsweredSurvey** | **int32** |  | 

## Methods

### NewQuestionnaireAnswer

`func NewQuestionnaireAnswer(id int32, created time.Time, modified time.Time, question int32, answeredSurvey int32, ) *QuestionnaireAnswer`

NewQuestionnaireAnswer instantiates a new QuestionnaireAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireAnswerWithDefaults

`func NewQuestionnaireAnswerWithDefaults() *QuestionnaireAnswer`

NewQuestionnaireAnswerWithDefaults instantiates a new QuestionnaireAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireAnswer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireAnswer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireAnswer) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *QuestionnaireAnswer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *QuestionnaireAnswer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *QuestionnaireAnswer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *QuestionnaireAnswer) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *QuestionnaireAnswer) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *QuestionnaireAnswer) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetQuestion

`func (o *QuestionnaireAnswer) GetQuestion() int32`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *QuestionnaireAnswer) GetQuestionOk() (*int32, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *QuestionnaireAnswer) SetQuestion(v int32)`

SetQuestion sets Question field to given value.


### GetAnsweredSurvey

`func (o *QuestionnaireAnswer) GetAnsweredSurvey() int32`

GetAnsweredSurvey returns the AnsweredSurvey field if non-nil, zero value otherwise.

### GetAnsweredSurveyOk

`func (o *QuestionnaireAnswer) GetAnsweredSurveyOk() (*int32, bool)`

GetAnsweredSurveyOk returns a tuple with the AnsweredSurvey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredSurvey

`func (o *QuestionnaireAnswer) SetAnsweredSurvey(v int32)`

SetAnsweredSurvey sets AnsweredSurvey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


