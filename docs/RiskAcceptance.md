# RiskAcceptance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Recommendation** | **string** |  | [readonly] 
**Decision** | **string** |  | [readonly] 
**Path** | **string** |  | [readonly] 
**Name** | **string** | Descriptive name which in the future may also be used to group risk acceptances together across engagements and products | 
**RecommendationDetails** | Pointer to **NullableString** | Explanation of security recommendation | [optional] 
**DecisionDetails** | Pointer to **NullableString** | If a compensating control exists to mitigate the finding or reduce risk, then list the compensating control(s). | [optional] 
**AcceptedBy** | Pointer to **NullableString** | The person that accepts the risk, can be outside of DefectDojo. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | When the risk acceptance expires, the findings will be reactivated (unless disabled below). | [optional] 
**ExpirationDateWarned** | Pointer to **NullableTime** | (readonly) Date at which notice about the risk acceptance expiration was sent. | [optional] 
**ExpirationDateHandled** | Pointer to **NullableTime** | (readonly) When the risk acceptance expiration was handled (manually or by the daily job). | [optional] 
**ReactivateExpired** | Pointer to **bool** | Reactivate findings when risk acceptance expires? | [optional] 
**RestartSlaExpired** | Pointer to **bool** | When enabled, the SLA for findings is restarted when the risk acceptance expires. | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Updated** | **time.Time** |  | [readonly] 
**Owner** | **int32** | User in DefectDojo owning this acceptance. Only the owner and staff users can edit the risk acceptance. | 
**AcceptedFindings** | **[]int32** |  | 
**Notes** | **[]int32** |  | [readonly] 

## Methods

### NewRiskAcceptance

`func NewRiskAcceptance(id int32, recommendation string, decision string, path string, name string, created time.Time, updated time.Time, owner int32, acceptedFindings []int32, notes []int32, ) *RiskAcceptance`

NewRiskAcceptance instantiates a new RiskAcceptance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAcceptanceWithDefaults

`func NewRiskAcceptanceWithDefaults() *RiskAcceptance`

NewRiskAcceptanceWithDefaults instantiates a new RiskAcceptance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskAcceptance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskAcceptance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskAcceptance) SetId(v int32)`

SetId sets Id field to given value.


### GetRecommendation

`func (o *RiskAcceptance) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *RiskAcceptance) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *RiskAcceptance) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.


### GetDecision

`func (o *RiskAcceptance) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *RiskAcceptance) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *RiskAcceptance) SetDecision(v string)`

SetDecision sets Decision field to given value.


### GetPath

`func (o *RiskAcceptance) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RiskAcceptance) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RiskAcceptance) SetPath(v string)`

SetPath sets Path field to given value.


### GetName

`func (o *RiskAcceptance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskAcceptance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskAcceptance) SetName(v string)`

SetName sets Name field to given value.


### GetRecommendationDetails

`func (o *RiskAcceptance) GetRecommendationDetails() string`

GetRecommendationDetails returns the RecommendationDetails field if non-nil, zero value otherwise.

### GetRecommendationDetailsOk

`func (o *RiskAcceptance) GetRecommendationDetailsOk() (*string, bool)`

GetRecommendationDetailsOk returns a tuple with the RecommendationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationDetails

`func (o *RiskAcceptance) SetRecommendationDetails(v string)`

SetRecommendationDetails sets RecommendationDetails field to given value.

### HasRecommendationDetails

`func (o *RiskAcceptance) HasRecommendationDetails() bool`

HasRecommendationDetails returns a boolean if a field has been set.

### SetRecommendationDetailsNil

`func (o *RiskAcceptance) SetRecommendationDetailsNil(b bool)`

 SetRecommendationDetailsNil sets the value for RecommendationDetails to be an explicit nil

### UnsetRecommendationDetails
`func (o *RiskAcceptance) UnsetRecommendationDetails()`

UnsetRecommendationDetails ensures that no value is present for RecommendationDetails, not even an explicit nil
### GetDecisionDetails

`func (o *RiskAcceptance) GetDecisionDetails() string`

GetDecisionDetails returns the DecisionDetails field if non-nil, zero value otherwise.

### GetDecisionDetailsOk

`func (o *RiskAcceptance) GetDecisionDetailsOk() (*string, bool)`

GetDecisionDetailsOk returns a tuple with the DecisionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDetails

`func (o *RiskAcceptance) SetDecisionDetails(v string)`

SetDecisionDetails sets DecisionDetails field to given value.

### HasDecisionDetails

`func (o *RiskAcceptance) HasDecisionDetails() bool`

HasDecisionDetails returns a boolean if a field has been set.

### SetDecisionDetailsNil

`func (o *RiskAcceptance) SetDecisionDetailsNil(b bool)`

 SetDecisionDetailsNil sets the value for DecisionDetails to be an explicit nil

### UnsetDecisionDetails
`func (o *RiskAcceptance) UnsetDecisionDetails()`

UnsetDecisionDetails ensures that no value is present for DecisionDetails, not even an explicit nil
### GetAcceptedBy

`func (o *RiskAcceptance) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *RiskAcceptance) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *RiskAcceptance) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *RiskAcceptance) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### SetAcceptedByNil

`func (o *RiskAcceptance) SetAcceptedByNil(b bool)`

 SetAcceptedByNil sets the value for AcceptedBy to be an explicit nil

### UnsetAcceptedBy
`func (o *RiskAcceptance) UnsetAcceptedBy()`

UnsetAcceptedBy ensures that no value is present for AcceptedBy, not even an explicit nil
### GetExpirationDate

`func (o *RiskAcceptance) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RiskAcceptance) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RiskAcceptance) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RiskAcceptance) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *RiskAcceptance) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *RiskAcceptance) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetExpirationDateWarned

`func (o *RiskAcceptance) GetExpirationDateWarned() time.Time`

GetExpirationDateWarned returns the ExpirationDateWarned field if non-nil, zero value otherwise.

### GetExpirationDateWarnedOk

`func (o *RiskAcceptance) GetExpirationDateWarnedOk() (*time.Time, bool)`

GetExpirationDateWarnedOk returns a tuple with the ExpirationDateWarned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateWarned

`func (o *RiskAcceptance) SetExpirationDateWarned(v time.Time)`

SetExpirationDateWarned sets ExpirationDateWarned field to given value.

### HasExpirationDateWarned

`func (o *RiskAcceptance) HasExpirationDateWarned() bool`

HasExpirationDateWarned returns a boolean if a field has been set.

### SetExpirationDateWarnedNil

`func (o *RiskAcceptance) SetExpirationDateWarnedNil(b bool)`

 SetExpirationDateWarnedNil sets the value for ExpirationDateWarned to be an explicit nil

### UnsetExpirationDateWarned
`func (o *RiskAcceptance) UnsetExpirationDateWarned()`

UnsetExpirationDateWarned ensures that no value is present for ExpirationDateWarned, not even an explicit nil
### GetExpirationDateHandled

`func (o *RiskAcceptance) GetExpirationDateHandled() time.Time`

GetExpirationDateHandled returns the ExpirationDateHandled field if non-nil, zero value otherwise.

### GetExpirationDateHandledOk

`func (o *RiskAcceptance) GetExpirationDateHandledOk() (*time.Time, bool)`

GetExpirationDateHandledOk returns a tuple with the ExpirationDateHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateHandled

`func (o *RiskAcceptance) SetExpirationDateHandled(v time.Time)`

SetExpirationDateHandled sets ExpirationDateHandled field to given value.

### HasExpirationDateHandled

`func (o *RiskAcceptance) HasExpirationDateHandled() bool`

HasExpirationDateHandled returns a boolean if a field has been set.

### SetExpirationDateHandledNil

`func (o *RiskAcceptance) SetExpirationDateHandledNil(b bool)`

 SetExpirationDateHandledNil sets the value for ExpirationDateHandled to be an explicit nil

### UnsetExpirationDateHandled
`func (o *RiskAcceptance) UnsetExpirationDateHandled()`

UnsetExpirationDateHandled ensures that no value is present for ExpirationDateHandled, not even an explicit nil
### GetReactivateExpired

`func (o *RiskAcceptance) GetReactivateExpired() bool`

GetReactivateExpired returns the ReactivateExpired field if non-nil, zero value otherwise.

### GetReactivateExpiredOk

`func (o *RiskAcceptance) GetReactivateExpiredOk() (*bool, bool)`

GetReactivateExpiredOk returns a tuple with the ReactivateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivateExpired

`func (o *RiskAcceptance) SetReactivateExpired(v bool)`

SetReactivateExpired sets ReactivateExpired field to given value.

### HasReactivateExpired

`func (o *RiskAcceptance) HasReactivateExpired() bool`

HasReactivateExpired returns a boolean if a field has been set.

### GetRestartSlaExpired

`func (o *RiskAcceptance) GetRestartSlaExpired() bool`

GetRestartSlaExpired returns the RestartSlaExpired field if non-nil, zero value otherwise.

### GetRestartSlaExpiredOk

`func (o *RiskAcceptance) GetRestartSlaExpiredOk() (*bool, bool)`

GetRestartSlaExpiredOk returns a tuple with the RestartSlaExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSlaExpired

`func (o *RiskAcceptance) SetRestartSlaExpired(v bool)`

SetRestartSlaExpired sets RestartSlaExpired field to given value.

### HasRestartSlaExpired

`func (o *RiskAcceptance) HasRestartSlaExpired() bool`

HasRestartSlaExpired returns a boolean if a field has been set.

### GetCreated

`func (o *RiskAcceptance) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskAcceptance) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskAcceptance) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *RiskAcceptance) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RiskAcceptance) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RiskAcceptance) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetOwner

`func (o *RiskAcceptance) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RiskAcceptance) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RiskAcceptance) SetOwner(v int32)`

SetOwner sets Owner field to given value.


### GetAcceptedFindings

`func (o *RiskAcceptance) GetAcceptedFindings() []int32`

GetAcceptedFindings returns the AcceptedFindings field if non-nil, zero value otherwise.

### GetAcceptedFindingsOk

`func (o *RiskAcceptance) GetAcceptedFindingsOk() (*[]int32, bool)`

GetAcceptedFindingsOk returns a tuple with the AcceptedFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFindings

`func (o *RiskAcceptance) SetAcceptedFindings(v []int32)`

SetAcceptedFindings sets AcceptedFindings field to given value.


### GetNotes

`func (o *RiskAcceptance) GetNotes() []int32`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RiskAcceptance) GetNotesOk() (*[]int32, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RiskAcceptance) SetNotes(v []int32)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


