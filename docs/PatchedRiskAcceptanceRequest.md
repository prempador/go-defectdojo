# PatchedRiskAcceptanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Descriptive name which in the future may also be used to group risk acceptances together across engagements and products | [optional] 
**RecommendationDetails** | Pointer to **NullableString** | Explanation of security recommendation | [optional] 
**DecisionDetails** | Pointer to **NullableString** | If a compensating control exists to mitigate the finding or reduce risk, then list the compensating control(s). | [optional] 
**AcceptedBy** | Pointer to **NullableString** | The person that accepts the risk, can be outside of DefectDojo. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | When the risk acceptance expires, the findings will be reactivated (unless disabled below). | [optional] 
**ExpirationDateWarned** | Pointer to **NullableTime** | (readonly) Date at which notice about the risk acceptance expiration was sent. | [optional] 
**ExpirationDateHandled** | Pointer to **NullableTime** | (readonly) When the risk acceptance expiration was handled (manually or by the daily job). | [optional] 
**ReactivateExpired** | Pointer to **bool** | Reactivate findings when risk acceptance expires? | [optional] 
**RestartSlaExpired** | Pointer to **bool** | When enabled, the SLA for findings is restarted when the risk acceptance expires. | [optional] 
**Owner** | Pointer to **int32** | User in DefectDojo owning this acceptance. Only the owner and staff users can edit the risk acceptance. | [optional] 
**AcceptedFindings** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedRiskAcceptanceRequest

`func NewPatchedRiskAcceptanceRequest() *PatchedRiskAcceptanceRequest`

NewPatchedRiskAcceptanceRequest instantiates a new PatchedRiskAcceptanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRiskAcceptanceRequestWithDefaults

`func NewPatchedRiskAcceptanceRequestWithDefaults() *PatchedRiskAcceptanceRequest`

NewPatchedRiskAcceptanceRequestWithDefaults instantiates a new PatchedRiskAcceptanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRiskAcceptanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRiskAcceptanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRiskAcceptanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRiskAcceptanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecommendationDetails

`func (o *PatchedRiskAcceptanceRequest) GetRecommendationDetails() string`

GetRecommendationDetails returns the RecommendationDetails field if non-nil, zero value otherwise.

### GetRecommendationDetailsOk

`func (o *PatchedRiskAcceptanceRequest) GetRecommendationDetailsOk() (*string, bool)`

GetRecommendationDetailsOk returns a tuple with the RecommendationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationDetails

`func (o *PatchedRiskAcceptanceRequest) SetRecommendationDetails(v string)`

SetRecommendationDetails sets RecommendationDetails field to given value.

### HasRecommendationDetails

`func (o *PatchedRiskAcceptanceRequest) HasRecommendationDetails() bool`

HasRecommendationDetails returns a boolean if a field has been set.

### SetRecommendationDetailsNil

`func (o *PatchedRiskAcceptanceRequest) SetRecommendationDetailsNil(b bool)`

 SetRecommendationDetailsNil sets the value for RecommendationDetails to be an explicit nil

### UnsetRecommendationDetails
`func (o *PatchedRiskAcceptanceRequest) UnsetRecommendationDetails()`

UnsetRecommendationDetails ensures that no value is present for RecommendationDetails, not even an explicit nil
### GetDecisionDetails

`func (o *PatchedRiskAcceptanceRequest) GetDecisionDetails() string`

GetDecisionDetails returns the DecisionDetails field if non-nil, zero value otherwise.

### GetDecisionDetailsOk

`func (o *PatchedRiskAcceptanceRequest) GetDecisionDetailsOk() (*string, bool)`

GetDecisionDetailsOk returns a tuple with the DecisionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDetails

`func (o *PatchedRiskAcceptanceRequest) SetDecisionDetails(v string)`

SetDecisionDetails sets DecisionDetails field to given value.

### HasDecisionDetails

`func (o *PatchedRiskAcceptanceRequest) HasDecisionDetails() bool`

HasDecisionDetails returns a boolean if a field has been set.

### SetDecisionDetailsNil

`func (o *PatchedRiskAcceptanceRequest) SetDecisionDetailsNil(b bool)`

 SetDecisionDetailsNil sets the value for DecisionDetails to be an explicit nil

### UnsetDecisionDetails
`func (o *PatchedRiskAcceptanceRequest) UnsetDecisionDetails()`

UnsetDecisionDetails ensures that no value is present for DecisionDetails, not even an explicit nil
### GetAcceptedBy

`func (o *PatchedRiskAcceptanceRequest) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *PatchedRiskAcceptanceRequest) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *PatchedRiskAcceptanceRequest) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.

### HasAcceptedBy

`func (o *PatchedRiskAcceptanceRequest) HasAcceptedBy() bool`

HasAcceptedBy returns a boolean if a field has been set.

### SetAcceptedByNil

`func (o *PatchedRiskAcceptanceRequest) SetAcceptedByNil(b bool)`

 SetAcceptedByNil sets the value for AcceptedBy to be an explicit nil

### UnsetAcceptedBy
`func (o *PatchedRiskAcceptanceRequest) UnsetAcceptedBy()`

UnsetAcceptedBy ensures that no value is present for AcceptedBy, not even an explicit nil
### GetExpirationDate

`func (o *PatchedRiskAcceptanceRequest) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PatchedRiskAcceptanceRequest) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PatchedRiskAcceptanceRequest) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *PatchedRiskAcceptanceRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *PatchedRiskAcceptanceRequest) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *PatchedRiskAcceptanceRequest) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetExpirationDateWarned

`func (o *PatchedRiskAcceptanceRequest) GetExpirationDateWarned() time.Time`

GetExpirationDateWarned returns the ExpirationDateWarned field if non-nil, zero value otherwise.

### GetExpirationDateWarnedOk

`func (o *PatchedRiskAcceptanceRequest) GetExpirationDateWarnedOk() (*time.Time, bool)`

GetExpirationDateWarnedOk returns a tuple with the ExpirationDateWarned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateWarned

`func (o *PatchedRiskAcceptanceRequest) SetExpirationDateWarned(v time.Time)`

SetExpirationDateWarned sets ExpirationDateWarned field to given value.

### HasExpirationDateWarned

`func (o *PatchedRiskAcceptanceRequest) HasExpirationDateWarned() bool`

HasExpirationDateWarned returns a boolean if a field has been set.

### SetExpirationDateWarnedNil

`func (o *PatchedRiskAcceptanceRequest) SetExpirationDateWarnedNil(b bool)`

 SetExpirationDateWarnedNil sets the value for ExpirationDateWarned to be an explicit nil

### UnsetExpirationDateWarned
`func (o *PatchedRiskAcceptanceRequest) UnsetExpirationDateWarned()`

UnsetExpirationDateWarned ensures that no value is present for ExpirationDateWarned, not even an explicit nil
### GetExpirationDateHandled

`func (o *PatchedRiskAcceptanceRequest) GetExpirationDateHandled() time.Time`

GetExpirationDateHandled returns the ExpirationDateHandled field if non-nil, zero value otherwise.

### GetExpirationDateHandledOk

`func (o *PatchedRiskAcceptanceRequest) GetExpirationDateHandledOk() (*time.Time, bool)`

GetExpirationDateHandledOk returns a tuple with the ExpirationDateHandled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateHandled

`func (o *PatchedRiskAcceptanceRequest) SetExpirationDateHandled(v time.Time)`

SetExpirationDateHandled sets ExpirationDateHandled field to given value.

### HasExpirationDateHandled

`func (o *PatchedRiskAcceptanceRequest) HasExpirationDateHandled() bool`

HasExpirationDateHandled returns a boolean if a field has been set.

### SetExpirationDateHandledNil

`func (o *PatchedRiskAcceptanceRequest) SetExpirationDateHandledNil(b bool)`

 SetExpirationDateHandledNil sets the value for ExpirationDateHandled to be an explicit nil

### UnsetExpirationDateHandled
`func (o *PatchedRiskAcceptanceRequest) UnsetExpirationDateHandled()`

UnsetExpirationDateHandled ensures that no value is present for ExpirationDateHandled, not even an explicit nil
### GetReactivateExpired

`func (o *PatchedRiskAcceptanceRequest) GetReactivateExpired() bool`

GetReactivateExpired returns the ReactivateExpired field if non-nil, zero value otherwise.

### GetReactivateExpiredOk

`func (o *PatchedRiskAcceptanceRequest) GetReactivateExpiredOk() (*bool, bool)`

GetReactivateExpiredOk returns a tuple with the ReactivateExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactivateExpired

`func (o *PatchedRiskAcceptanceRequest) SetReactivateExpired(v bool)`

SetReactivateExpired sets ReactivateExpired field to given value.

### HasReactivateExpired

`func (o *PatchedRiskAcceptanceRequest) HasReactivateExpired() bool`

HasReactivateExpired returns a boolean if a field has been set.

### GetRestartSlaExpired

`func (o *PatchedRiskAcceptanceRequest) GetRestartSlaExpired() bool`

GetRestartSlaExpired returns the RestartSlaExpired field if non-nil, zero value otherwise.

### GetRestartSlaExpiredOk

`func (o *PatchedRiskAcceptanceRequest) GetRestartSlaExpiredOk() (*bool, bool)`

GetRestartSlaExpiredOk returns a tuple with the RestartSlaExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartSlaExpired

`func (o *PatchedRiskAcceptanceRequest) SetRestartSlaExpired(v bool)`

SetRestartSlaExpired sets RestartSlaExpired field to given value.

### HasRestartSlaExpired

`func (o *PatchedRiskAcceptanceRequest) HasRestartSlaExpired() bool`

HasRestartSlaExpired returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedRiskAcceptanceRequest) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedRiskAcceptanceRequest) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedRiskAcceptanceRequest) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedRiskAcceptanceRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAcceptedFindings

`func (o *PatchedRiskAcceptanceRequest) GetAcceptedFindings() []int32`

GetAcceptedFindings returns the AcceptedFindings field if non-nil, zero value otherwise.

### GetAcceptedFindingsOk

`func (o *PatchedRiskAcceptanceRequest) GetAcceptedFindingsOk() (*[]int32, bool)`

GetAcceptedFindingsOk returns a tuple with the AcceptedFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedFindings

`func (o *PatchedRiskAcceptanceRequest) SetAcceptedFindings(v []int32)`

SetAcceptedFindings sets AcceptedFindings field to given value.

### HasAcceptedFindings

`func (o *PatchedRiskAcceptanceRequest) HasAcceptedFindings() bool`

HasAcceptedFindings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


