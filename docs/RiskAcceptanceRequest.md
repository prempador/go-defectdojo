# RiskAcceptanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Descriptive name which in the future may also be used to group risk acceptances together across engagements and products | 
**RecommendationDetails** | Pointer to **NullableString** | Explanation of security recommendation | [optional] 
**DecisionDetails** | Pointer to **NullableString** | If a compensating control exists to mitigate the finding or reduce risk, then list the compensating control(s). | [optional] 
**AcceptedBy** | Pointer to **NullableString** | The person that accepts the risk, can be outside of DefectDojo. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | When the risk acceptance expires, the findings will be reactivated (unless disabled below). | [optional] 
**ExpirationDateWarned** | Pointer to **NullableTime** | (readonly) Date at which notice about the risk acceptance expiration was sent. | [optional] 
**ExpirationDateHandled** | Pointer to **NullableTime** | (readonly) When the risk acceptance expiration was handled (manually or by the daily job). | [optional] 
**ReactivateExpired** | Pointer to **bool** | Reactivate findings when risk acceptance expires? | [optional] 
**RestartSlaExpired** | Pointer to **bool** | When enabled, the SLA for findings is restarted when the risk acceptance expires. | [optional] 
**Owner** | **int32** | User in DefectDojo owning this acceptance. Only the owner and staff users can edit the risk acceptance. | 
**AcceptedFindings** | **[]int32** |  | 

## Methods

### NewRiskAcceptanceRequest

`func NewRiskAcceptanceRequest(name string, owner int32, acceptedFindings []int32, ) *RiskAcceptanceRequest`

NewRiskAcceptanceRequest instantiates a new RiskAcceptanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAcceptanceRequestWithDefaults

`func NewRiskAcceptanceRequestWithDefaults() *RiskAcceptanceRequest`

NewRiskAcceptanceRequestWithDefaults instantiates a new RiskAcceptanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RiskAcceptanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskAcceptanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskAcceptanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRecommendationDetails

`func (o *RiskAcceptanceRequest) GetRecommendationDetails() string`

GetRecommendationDetails returns the RecommendationDetails field if non-nil, zero value otherwise.

### GetRecommendationDetailsOk

`func (o *RiskAcceptanceRequest) GetRecommendationDetailsOk() (*string, bool)`

GetRecommendationDetailsOk returns a tuple with the RecommendationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationDetails

`func (o *RiskAcceptanceRequest) SetRecommendationDetails(v string)`

SetRecommendationDetails sets RecommendationDetails field to given value.

### HasRecommendationDetails

`func (o *RiskAcceptanceRequest) HasRecommendationDetails() bool`

HasRecommendationDetails returns a boolean if a field has been set.

### SetRecommendationDetailsNil

`func (o *RiskAcceptanceRequest) SetRecommendationDetailsNil(b bool)`

 SetRecommendationDetailsNil sets the value for RecommendationDetails to be an explicit nil

### UnsetRecommendationDetails
`func (o *RiskAcceptanceRequest) UnsetRecommendationDetails()`

UnsetRecommendationDetails ensures that no value is present for RecommendationDetails, not even an explicit nil
### GetDecisionDetails

`func (o *RiskAcceptanceRequest) GetDecisionDetails() string`

GetDecisionDetails returns the DecisionDetails field if non-nil, zero value otherwise.

### GetDecisionDetailsOk

`func (o *RiskAcceptanceRequest) GetDecisionDetailsOk() (*string, bool)`

GetDecisionDetailsOk returns a tuple with the DecisionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDetails

`func (o *RiskAcceptanceRequest) SetDecisionDetails(v string)`

SetDecisionDetails sets DecisionDetails field to given value.

### HasDecisionDetails

`func (o *RiskAcceptanceRequest) HasDecisionDetails() bool`

HasDecisionDetails returns a boolean if a field has been set.

### SetDecisionDetailsNil

`func (o *RiskAcceptanceRequest) SetDecisionDetailsNil(b bool)`

 SetDecisionDetailsNil sets the value for DecisionDetails to be an explicit nil

### UnsetDecisionDetails
`func (o *RiskAcceptanceRequest) UnsetDecisionDetails()`

UnsetDecisionDetails ensures that no value is present for DecisionDetails, not even an explicit nil
### GetAcceptedBy

`func (o *RiskAcceptanceRequest) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *RiskAcceptanceRequest) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *RiskAcceptanceRequest) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *RiskAcceptanceRequest) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### SetAcceptedByNil

`func (o *RiskAcceptanceRequest) SetAcceptedByNil(b bool)`

 SetAcceptedByNil sets the value for AcceptedBy to be an explicit nil

### UnsetAcceptedBy
`func (o *RiskAcceptanceRequest) UnsetAcceptedBy()`

UnsetAcceptedBy ensures that no value is present for AcceptedBy, not even an explicit nil
### GetExpirationDate

`func (o *RiskAcceptanceRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RiskAcceptanceRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RiskAcceptanceRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RiskAcceptanceRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *RiskAcceptanceRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *RiskAcceptanceRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetExpirationDateWarned

`func (o *RiskAcceptanceRequest) GetExpirationDateWarned() time.Time`

GetExpirationDateWarned returns the ExpirationDateWarned field if non-nil, zero value otherwise.

### GetExpirationDateWarnedOk

`func (o *RiskAcceptanceRequest) GetExpirationDateWarnedOk() (*time.Time, bool)`

GetExpirationDateWarnedOk returns a tuple with the ExpirationDateWarned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateWarned

`func (o *RiskAcceptanceRequest) SetExpirationDateWarned(v time.Time)`

SetExpirationDateWarned sets ExpirationDateWarned field to given value.

### HasExpirationDateWarned

`func (o *RiskAcceptanceRequest) HasExpirationDateWarned() bool`

HasExpirationDateWarned returns a boolean if a field has been set.

### SetExpirationDateWarnedNil

`func (o *RiskAcceptanceRequest) SetExpirationDateWarnedNil(b bool)`

 SetExpirationDateWarnedNil sets the value for ExpirationDateWarned to be an explicit nil

### UnsetExpirationDateWarned
`func (o *RiskAcceptanceRequest) UnsetExpirationDateWarned()`

UnsetExpirationDateWarned ensures that no value is present for ExpirationDateWarned, not even an explicit nil
### GetExpirationDateHandled

`func (o *RiskAcceptanceRequest) GetExpirationDateHandled() time.Time`

GetExpirationDateHandled returns the ExpirationDateHandled field if non-nil, zero value otherwise.

### GetExpirationDateHandledOk

`func (o *RiskAcceptanceRequest) GetExpirationDateHandledOk() (*time.Time, bool)`

GetExpirationDateHandledOk returns a tuple with the ExpirationDateHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateHandled

`func (o *RiskAcceptanceRequest) SetExpirationDateHandled(v time.Time)`

SetExpirationDateHandled sets ExpirationDateHandled field to given value.

### HasExpirationDateHandled

`func (o *RiskAcceptanceRequest) HasExpirationDateHandled() bool`

HasExpirationDateHandled returns a boolean if a field has been set.

### SetExpirationDateHandledNil

`func (o *RiskAcceptanceRequest) SetExpirationDateHandledNil(b bool)`

 SetExpirationDateHandledNil sets the value for ExpirationDateHandled to be an explicit nil

### UnsetExpirationDateHandled
`func (o *RiskAcceptanceRequest) UnsetExpirationDateHandled()`

UnsetExpirationDateHandled ensures that no value is present for ExpirationDateHandled, not even an explicit nil
### GetReactivateExpired

`func (o *RiskAcceptanceRequest) GetReactivateExpired() bool`

GetReactivateExpired returns the ReactivateExpired field if non-nil, zero value otherwise.

### GetReactivateExpiredOk

`func (o *RiskAcceptanceRequest) GetReactivateExpiredOk() (*bool, bool)`

GetReactivateExpiredOk returns a tuple with the ReactivateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivateExpired

`func (o *RiskAcceptanceRequest) SetReactivateExpired(v bool)`

SetReactivateExpired sets ReactivateExpired field to given value.

### HasReactivateExpired

`func (o *RiskAcceptanceRequest) HasReactivateExpired() bool`

HasReactivateExpired returns a boolean if a field has been set.

### GetRestartSlaExpired

`func (o *RiskAcceptanceRequest) GetRestartSlaExpired() bool`

GetRestartSlaExpired returns the RestartSlaExpired field if non-nil, zero value otherwise.

### GetRestartSlaExpiredOk

`func (o *RiskAcceptanceRequest) GetRestartSlaExpiredOk() (*bool, bool)`

GetRestartSlaExpiredOk returns a tuple with the RestartSlaExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSlaExpired

`func (o *RiskAcceptanceRequest) SetRestartSlaExpired(v bool)`

SetRestartSlaExpired sets RestartSlaExpired field to given value.

### HasRestartSlaExpired

`func (o *RiskAcceptanceRequest) HasRestartSlaExpired() bool`

HasRestartSlaExpired returns a boolean if a field has been set.

### GetOwner

`func (o *RiskAcceptanceRequest) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RiskAcceptanceRequest) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RiskAcceptanceRequest) SetOwner(v int32)`

SetOwner sets Owner field to given value.


### GetAcceptedFindings

`func (o *RiskAcceptanceRequest) GetAcceptedFindings() []int32`

GetAcceptedFindings returns the AcceptedFindings field if non-nil, zero value otherwise.

### GetAcceptedFindingsOk

`func (o *RiskAcceptanceRequest) GetAcceptedFindingsOk() (*[]int32, bool)`

GetAcceptedFindingsOk returns a tuple with the AcceptedFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFindings

`func (o *RiskAcceptanceRequest) SetAcceptedFindings(v []int32)`

SetAcceptedFindings sets AcceptedFindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


