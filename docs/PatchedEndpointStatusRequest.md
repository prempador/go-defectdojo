# PatchedEndpointStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Mitigated** | Pointer to **bool** |  | [optional] 
**FalsePositive** | Pointer to **bool** |  | [optional] 
**OutOfScope** | Pointer to **bool** |  | [optional] 
**RiskAccepted** | Pointer to **bool** |  | [optional] 
**MitigatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Endpoint** | Pointer to **int32** |  | [optional] 
**Finding** | Pointer to **int32** |  | [optional] 

## Methods

### NewPatchedEndpointStatusRequest

`func NewPatchedEndpointStatusRequest() *PatchedEndpointStatusRequest`

NewPatchedEndpointStatusRequest instantiates a new PatchedEndpointStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedEndpointStatusRequestWithDefaults

`func NewPatchedEndpointStatusRequestWithDefaults() *PatchedEndpointStatusRequest`

NewPatchedEndpointStatusRequestWithDefaults instantiates a new PatchedEndpointStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PatchedEndpointStatusRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PatchedEndpointStatusRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PatchedEndpointStatusRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PatchedEndpointStatusRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMitigated

`func (o *PatchedEndpointStatusRequest) GetMitigated() bool`

GetMitigated returns the Mitigated field if non-nil, zero value otherwise.

### GetMitigatedOk

`func (o *PatchedEndpointStatusRequest) GetMitigatedOk() (*bool, bool)`

GetMitigatedOk returns a tuple with the Mitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigated

`func (o *PatchedEndpointStatusRequest) SetMitigated(v bool)`

SetMitigated sets Mitigated field to given value.

### HasMitigated

`func (o *PatchedEndpointStatusRequest) HasMitigated() bool`

HasMitigated returns a boolean if a field has been set.

### GetFalsePositive

`func (o *PatchedEndpointStatusRequest) GetFalsePositive() bool`

GetFalsePositive returns the FalsePositive field if non-nil, zero value otherwise.

### GetFalsePositiveOk

`func (o *PatchedEndpointStatusRequest) GetFalsePositiveOk() (*bool, bool)`

GetFalsePositiveOk returns a tuple with the FalsePositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositive

`func (o *PatchedEndpointStatusRequest) SetFalsePositive(v bool)`

SetFalsePositive sets FalsePositive field to given value.

### HasFalsePositive

`func (o *PatchedEndpointStatusRequest) HasFalsePositive() bool`

HasFalsePositive returns a boolean if a field has been set.

### GetOutOfScope

`func (o *PatchedEndpointStatusRequest) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *PatchedEndpointStatusRequest) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *PatchedEndpointStatusRequest) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *PatchedEndpointStatusRequest) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetRiskAccepted

`func (o *PatchedEndpointStatusRequest) GetRiskAccepted() bool`

GetRiskAccepted returns the RiskAccepted field if non-nil, zero value otherwise.

### GetRiskAcceptedOk

`func (o *PatchedEndpointStatusRequest) GetRiskAcceptedOk() (*bool, bool)`

GetRiskAcceptedOk returns a tuple with the RiskAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAccepted

`func (o *PatchedEndpointStatusRequest) SetRiskAccepted(v bool)`

SetRiskAccepted sets RiskAccepted field to given value.

### HasRiskAccepted

`func (o *PatchedEndpointStatusRequest) HasRiskAccepted() bool`

HasRiskAccepted returns a boolean if a field has been set.

### GetMitigatedBy

`func (o *PatchedEndpointStatusRequest) GetMitigatedBy() int32`

GetMitigatedBy returns the MitigatedBy field if non-nil, zero value otherwise.

### GetMitigatedByOk

`func (o *PatchedEndpointStatusRequest) GetMitigatedByOk() (*int32, bool)`

GetMitigatedByOk returns a tuple with the MitigatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigatedBy

`func (o *PatchedEndpointStatusRequest) SetMitigatedBy(v int32)`

SetMitigatedBy sets MitigatedBy field to given value.

### HasMitigatedBy

`func (o *PatchedEndpointStatusRequest) HasMitigatedBy() bool`

HasMitigatedBy returns a boolean if a field has been set.

### SetMitigatedByNil

`func (o *PatchedEndpointStatusRequest) SetMitigatedByNil(b bool)`

 SetMitigatedByNil sets the value for MitigatedBy to be an explicit nil

### UnsetMitigatedBy
`func (o *PatchedEndpointStatusRequest) UnsetMitigatedBy()`

UnsetMitigatedBy ensures that no value is present for MitigatedBy, not even an explicit nil
### GetEndpoint

`func (o *PatchedEndpointStatusRequest) GetEndpoint() int32`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PatchedEndpointStatusRequest) GetEndpointOk() (*int32, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PatchedEndpointStatusRequest) SetEndpoint(v int32)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PatchedEndpointStatusRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetFinding

`func (o *PatchedEndpointStatusRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *PatchedEndpointStatusRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *PatchedEndpointStatusRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *PatchedEndpointStatusRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


