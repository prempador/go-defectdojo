# QuestionnaireGeneralSurvey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Survey** | [**QuestionnaireEngagementSurvey**](QuestionnaireEngagementSurvey.md) |  | 
**NumResponses** | Pointer to **int32** |  | [optional] 
**Generated** | **NullableTime** |  | [readonly] 
**Expiration** | **time.Time** |  | 

## Methods

### NewQuestionnaireGeneralSurvey

`func NewQuestionnaireGeneralSurvey(id int32, survey QuestionnaireEngagementSurvey, generated NullableTime, expiration time.Time, ) *QuestionnaireGeneralSurvey`

NewQuestionnaireGeneralSurvey instantiates a new QuestionnaireGeneralSurvey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireGeneralSurveyWithDefaults

`func NewQuestionnaireGeneralSurveyWithDefaults() *QuestionnaireGeneralSurvey`

NewQuestionnaireGeneralSurveyWithDefaults instantiates a new QuestionnaireGeneralSurvey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireGeneralSurvey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireGeneralSurvey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireGeneralSurvey) SetId(v int32)`

SetId sets Id field to given value.


### GetSurvey

`func (o *QuestionnaireGeneralSurvey) GetSurvey() QuestionnaireEngagementSurvey`

GetSurvey returns the Survey field if non-nil, zero value otherwise.

### GetSurveyOk

`func (o *QuestionnaireGeneralSurvey) GetSurveyOk() (*QuestionnaireEngagementSurvey, bool)`

GetSurveyOk returns a tuple with the Survey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurvey

`func (o *QuestionnaireGeneralSurvey) SetSurvey(v QuestionnaireEngagementSurvey)`

SetSurvey sets Survey field to given value.


### GetNumResponses

`func (o *QuestionnaireGeneralSurvey) GetNumResponses() int32`

GetNumResponses returns the NumResponses field if non-nil, zero value otherwise.

### GetNumResponsesOk

`func (o *QuestionnaireGeneralSurvey) GetNumResponsesOk() (*int32, bool)`

GetNumResponsesOk returns a tuple with the NumResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumResponses

`func (o *QuestionnaireGeneralSurvey) SetNumResponses(v int32)`

SetNumResponses sets NumResponses field to given value.

### HasNumResponses

`func (o *QuestionnaireGeneralSurvey) HasNumResponses() bool`

HasNumResponses returns a boolean if a field has been set.

### GetGenerated

`func (o *QuestionnaireGeneralSurvey) GetGenerated() time.Time`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *QuestionnaireGeneralSurvey) GetGeneratedOk() (*time.Time, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *QuestionnaireGeneralSurvey) SetGenerated(v time.Time)`

SetGenerated sets Generated field to given value.


### SetGeneratedNil

`func (o *QuestionnaireGeneralSurvey) SetGeneratedNil(b bool)`

 SetGeneratedNil sets the value for Generated to be an explicit nil

### UnsetGenerated
`func (o *QuestionnaireGeneralSurvey) UnsetGenerated()`

UnsetGenerated ensures that no value is present for Generated, not even an explicit nil
### GetExpiration

`func (o *QuestionnaireGeneralSurvey) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *QuestionnaireGeneralSurvey) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *QuestionnaireGeneralSurvey) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


