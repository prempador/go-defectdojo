# EndpointStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Date** | Pointer to **string** |  | [optional] 
**LastModified** | **NullableTime** |  | [readonly] 
**Mitigated** | Pointer to **bool** |  | [optional] 
**MitigatedTime** | **NullableTime** |  | [readonly] 
**FalsePositive** | Pointer to **bool** |  | [optional] 
**OutOfScope** | Pointer to **bool** |  | [optional] 
**RiskAccepted** | Pointer to **bool** |  | [optional] 
**MitigatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | **int32** |  | 
**Finding** | **int32** |  | 

## Methods

### NewEndpointStatus

`func NewEndpointStatus(id int32, lastModified NullableTime, mitigatedTime NullableTime, endpoint int32, finding int32, ) *EndpointStatus`

NewEndpointStatus instantiates a new EndpointStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointStatusWithDefaults

`func NewEndpointStatusWithDefaults() *EndpointStatus`

NewEndpointStatusWithDefaults instantiates a new EndpointStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointStatus) SetId(v int32)`

SetId sets Id field to given value.


### GetDate

`func (o *EndpointStatus) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EndpointStatus) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EndpointStatus) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EndpointStatus) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastModified

`func (o *EndpointStatus) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EndpointStatus) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EndpointStatus) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *EndpointStatus) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *EndpointStatus) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetMitigated

`func (o *EndpointStatus) GetMitigated() bool`

GetMitigated returns the Mitigated field if non-nil, zero value otherwise.

### GetMitigatedOk

`func (o *EndpointStatus) GetMitigatedOk() (*bool, bool)`

GetMitigatedOk returns a tuple with the Mitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigated

`func (o *EndpointStatus) SetMitigated(v bool)`

SetMitigated sets Mitigated field to given value.

### HasMitigated

`func (o *EndpointStatus) HasMitigated() bool`

HasMitigated returns a boolean if a field has been set.

### GetMitigatedTime

`func (o *EndpointStatus) GetMitigatedTime() time.Time`

GetMitigatedTime returns the MitigatedTime field if non-nil, zero value otherwise.

### GetMitigatedTimeOk

`func (o *EndpointStatus) GetMitigatedTimeOk() (*time.Time, bool)`

GetMitigatedTimeOk returns a tuple with the MitigatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedTime

`func (o *EndpointStatus) SetMitigatedTime(v time.Time)`

SetMitigatedTime sets MitigatedTime field to given value.


### SetMitigatedTimeNil

`func (o *EndpointStatus) SetMitigatedTimeNil(b bool)`

 SetMitigatedTimeNil sets the value for MitigatedTime to be an explicit nil

### UnsetMitigatedTime
`func (o *EndpointStatus) UnsetMitigatedTime()`

UnsetMitigatedTime ensures that no value is present for MitigatedTime, not even an explicit nil
### GetFalsePositive

`func (o *EndpointStatus) GetFalsePositive() bool`

GetFalsePositive returns the FalsePositive field if non-nil, zero value otherwise.

### GetFalsePositiveOk

`func (o *EndpointStatus) GetFalsePositiveOk() (*bool, bool)`

GetFalsePositiveOk returns a tuple with the FalsePositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositive

`func (o *EndpointStatus) SetFalsePositive(v bool)`

SetFalsePositive sets FalsePositive field to given value.

### HasFalsePositive

`func (o *EndpointStatus) HasFalsePositive() bool`

HasFalsePositive returns a boolean if a field has been set.

### GetOutOfScope

`func (o *EndpointStatus) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *EndpointStatus) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *EndpointStatus) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *EndpointStatus) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *EndpointStatus) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *EndpointStatus) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *EndpointStatus) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *EndpointStatus) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetMitigatedBy

`func (o *EndpointStatus) GetMitigatedBy() int32`

GetMitigatedBy returns the MitigatedBy field if non-nil, zero value otherwise.

### GetMitigatedByOk

`func (o *EndpointStatus) GetMitigatedByOk() (*int32, bool)`

GetMitigatedByOk returns a tuple with the MitigatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedBy

`func (o *EndpointStatus) SetMitigatedBy(v int32)`

SetMitigatedBy sets MitigatedBy field to given value.

### HasMitigatedBy

`func (o *EndpointStatus) HasMitigatedBy() bool`

HasMitigatedBy returns a boolean if a field has been set.

### SetMitigatedByNil

`func (o *EndpointStatus) SetMitigatedByNil(b bool)`

 SetMitigatedByNil sets the value for MitigatedBy to be an explicit nil

### UnsetMitigatedBy
`func (o *EndpointStatus) UnsetMitigatedBy()`

UnsetMitigatedBy ensures that no value is present for MitigatedBy, not even an explicit nil
### GetEndpoint

`func (o *EndpointStatus) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EndpointStatus) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EndpointStatus) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.


### GetFinding

`func (o *EndpointStatus) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *EndpointStatus) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *EndpointStatus) SetFinding(v int32)`

SetFinding sets Finding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


