# EndpointStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Mitigated** | Pointer to **bool** |  | [optional] 
**FalsePositive** | Pointer to **bool** |  | [optional] 
**OutOfScope** | Pointer to **bool** |  | [optional] 
**RiskAccepted** | Pointer to **bool** |  | [optional] 
**MitigatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | **int32** |  | 
**Finding** | **int32** |  | 

## Methods

### NewEndpointStatusRequest

`func NewEndpointStatusRequest(endpoint int32, finding int32, ) *EndpointStatusRequest`

NewEndpointStatusRequest instantiates a new EndpointStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointStatusRequestWithDefaults

`func NewEndpointStatusRequestWithDefaults() *EndpointStatusRequest`

NewEndpointStatusRequestWithDefaults instantiates a new EndpointStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *EndpointStatusRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EndpointStatusRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EndpointStatusRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *EndpointStatusRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMitigated

`func (o *EndpointStatusRequest) GetMitigated() bool`

GetMitigated returns the Mitigated field if non-nil, zero value otherwise.

### GetMitigatedOk

`func (o *EndpointStatusRequest) GetMitigatedOk() (*bool, bool)`

GetMitigatedOk returns a tuple with the Mitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigated

`func (o *EndpointStatusRequest) SetMitigated(v bool)`

SetMitigated sets Mitigated field to given value.

### HasMitigated

`func (o *EndpointStatusRequest) HasMitigated() bool`

HasMitigated returns a boolean if a field has been set.

### GetFalsePositive

`func (o *EndpointStatusRequest) GetFalsePositive() bool`

GetFalsePositive returns the FalsePositive field if non-nil, zero value otherwise.

### GetFalsePositiveOk

`func (o *EndpointStatusRequest) GetFalsePositiveOk() (*bool, bool)`

GetFalsePositiveOk returns a tuple with the FalsePositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositive

`func (o *EndpointStatusRequest) SetFalsePositive(v bool)`

SetFalsePositive sets FalsePositive field to given value.

### HasFalsePositive

`func (o *EndpointStatusRequest) HasFalsePositive() bool`

HasFalsePositive returns a boolean if a field has been set.

### GetOutOfScope

`func (o *EndpointStatusRequest) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *EndpointStatusRequest) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *EndpointStatusRequest) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *EndpointStatusRequest) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *EndpointStatusRequest) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *EndpointStatusRequest) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *EndpointStatusRequest) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *EndpointStatusRequest) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetMitigatedBy

`func (o *EndpointStatusRequest) GetMitigatedBy() int32`

GetMitigatedBy returns the MitigatedBy field if non-nil, zero value otherwise.

### GetMitigatedByOk

`func (o *EndpointStatusRequest) GetMitigatedByOk() (*int32, bool)`

GetMitigatedByOk returns a tuple with the MitigatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedBy

`func (o *EndpointStatusRequest) SetMitigatedBy(v int32)`

SetMitigatedBy sets MitigatedBy field to given value.

### HasMitigatedBy

`func (o *EndpointStatusRequest) HasMitigatedBy() bool`

HasMitigatedBy returns a boolean if a field has been set.

### SetMitigatedByNil

`func (o *EndpointStatusRequest) SetMitigatedByNil(b bool)`

 SetMitigatedByNil sets the value for MitigatedBy to be an explicit nil

### UnsetMitigatedBy
`func (o *EndpointStatusRequest) UnsetMitigatedBy()`

UnsetMitigatedBy ensures that no value is present for MitigatedBy, not even an explicit nil
### GetEndpoint

`func (o *EndpointStatusRequest) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *EndpointStatusRequest) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *EndpointStatusRequest) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.


### GetFinding

`func (o *EndpointStatusRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *EndpointStatusRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *EndpointStatusRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


