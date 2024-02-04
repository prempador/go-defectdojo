# QuestionnaireEngagementSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Questions** | **[]string** |  | [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewQuestionnaireEngagementSurvey

`func NewQuestionnaireEngagementSurvey(id int32, questions []string, ) *QuestionnaireEngagementSurvey`

NewQuestionnaireEngagementSurvey instantiates a new QuestionnaireEngagementSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireEngagementSurveyWithDefaults

`func NewQuestionnaireEngagementSurveyWithDefaults() *QuestionnaireEngagementSurvey`

NewQuestionnaireEngagementSurveyWithDefaults instantiates a new QuestionnaireEngagementSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireEngagementSurvey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireEngagementSurvey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireEngagementSurvey) SetId(v int32)`

SetId sets Id field to given value.


### GetQuestions

`func (o *QuestionnaireEngagementSurvey) GetQuestions() []string`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *QuestionnaireEngagementSurvey) GetQuestionsOk() (*[]string, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *QuestionnaireEngagementSurvey) SetQuestions(v []string)`

SetQuestions sets Questions field to given value.


### GetName

`func (o *QuestionnaireEngagementSurvey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuestionnaireEngagementSurvey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuestionnaireEngagementSurvey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuestionnaireEngagementSurvey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *QuestionnaireEngagementSurvey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuestionnaireEngagementSurvey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuestionnaireEngagementSurvey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuestionnaireEngagementSurvey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireEngagementSurvey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireEngagementSurvey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireEngagementSurvey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireEngagementSurvey) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


