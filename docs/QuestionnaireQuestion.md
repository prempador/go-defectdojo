# QuestionnaireQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**Order** | Pointer to **int32** | The render order | [optional] 
**Optional** | Pointer to **bool** | If selected, user doesn&#39;t have to answer this question | [optional] 
**Text** | Pointer to **string** | The question text | [optional] 

## Methods

### NewQuestionnaireQuestion

`func NewQuestionnaireQuestion(id int32, created time.Time, modified time.Time, ) *QuestionnaireQuestion`

NewQuestionnaireQuestion instantiates a new QuestionnaireQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireQuestionWithDefaults

`func NewQuestionnaireQuestionWithDefaults() *QuestionnaireQuestion`

NewQuestionnaireQuestionWithDefaults instantiates a new QuestionnaireQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireQuestion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireQuestion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireQuestion) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *QuestionnaireQuestion) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *QuestionnaireQuestion) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *QuestionnaireQuestion) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *QuestionnaireQuestion) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *QuestionnaireQuestion) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *QuestionnaireQuestion) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetOrder

`func (o *QuestionnaireQuestion) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuestionnaireQuestion) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuestionnaireQuestion) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QuestionnaireQuestion) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOptional

`func (o *QuestionnaireQuestion) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *QuestionnaireQuestion) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *QuestionnaireQuestion) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *QuestionnaireQuestion) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetText

`func (o *QuestionnaireQuestion) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *QuestionnaireQuestion) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *QuestionnaireQuestion) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *QuestionnaireQuestion) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


