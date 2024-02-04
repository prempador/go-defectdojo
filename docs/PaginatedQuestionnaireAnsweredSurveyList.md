# PaginatedQuestionnaireAnsweredSurveyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]QuestionnaireAnsweredSurvey**](QuestionnaireAnsweredSurvey.md) |  | [optional] 

## Methods

### NewPaginatedQuestionnaireAnsweredSurveyList

`func NewPaginatedQuestionnaireAnsweredSurveyList() *PaginatedQuestionnaireAnsweredSurveyList`

NewPaginatedQuestionnaireAnsweredSurveyList instantiates a new PaginatedQuestionnaireAnsweredSurveyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedQuestionnaireAnsweredSurveyListWithDefaults

`func NewPaginatedQuestionnaireAnsweredSurveyListWithDefaults() *PaginatedQuestionnaireAnsweredSurveyList`

NewPaginatedQuestionnaireAnsweredSurveyListWithDefaults instantiates a new PaginatedQuestionnaireAnsweredSurveyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedQuestionnaireAnsweredSurveyList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedQuestionnaireAnsweredSurveyList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedQuestionnaireAnsweredSurveyList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedQuestionnaireAnsweredSurveyList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedQuestionnaireAnsweredSurveyList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedQuestionnaireAnsweredSurveyList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedQuestionnaireAnsweredSurveyList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedQuestionnaireAnsweredSurveyList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedQuestionnaireAnsweredSurveyList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedQuestionnaireAnsweredSurveyList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetResults() []QuestionnaireAnsweredSurvey`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedQuestionnaireAnsweredSurveyList) GetResultsOk() (*[]QuestionnaireAnsweredSurvey, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedQuestionnaireAnsweredSurveyList) SetResults(v []QuestionnaireAnsweredSurvey)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedQuestionnaireAnsweredSurveyList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


