# FindingCloseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMitigated** | Pointer to **bool** |  | [optional] 
**Mitigated** | Pointer to **time.Time** |  | [optional] 
**FalseP** | Pointer to **bool** |  | [optional] 
**OutOfScope** | Pointer to **bool** |  | [optional] 
**Duplicate** | Pointer to **bool** |  | [optional] 

## Methods

### NewFindingCloseRequest

`func NewFindingCloseRequest() *FindingCloseRequest`

NewFindingCloseRequest instantiates a new FindingCloseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingCloseRequestWithDefaults

`func NewFindingCloseRequestWithDefaults() *FindingCloseRequest`

NewFindingCloseRequestWithDefaults instantiates a new FindingCloseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMitigated

`func (o *FindingCloseRequest) GetIsMitigated() bool`

GetIsMitigated returns the IsMitigated field if non-nil, zero value otherwise.

### GetIsMitigatedOk

`func (o *FindingCloseRequest) GetIsMitigatedOk() (*bool, bool)`

GetIsMitigatedOk returns a tuple with the IsMitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMitigated

`func (o *FindingCloseRequest) SetIsMitigated(v bool)`

SetIsMitigated sets IsMitigated field to given value.

### HasIsMitigated

`func (o *FindingCloseRequest) HasIsMitigated() bool`

HasIsMitigated returns a boolean if a field has been set.

### GetMitigated

`func (o *FindingCloseRequest) GetMitigated() time.Time`

GetMitigated returns the Mitigated field if non-nil, zero value otherwise.

### GetMitigatedOk

`func (o *FindingCloseRequest) GetMitigatedOk() (*time.Time, bool)`

GetMitigatedOk returns a tuple with the Mitigated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigated

`func (o *FindingCloseRequest) SetMitigated(v time.Time)`

SetMitigated sets Mitigated field to given value.

### HasMitigated

`func (o *FindingCloseRequest) HasMitigated() bool`

HasMitigated returns a boolean if a field has been set.

### GetFalseP

`func (o *FindingCloseRequest) GetFalseP() bool`

GetFalseP returns the FalseP field if non-nil, zero value otherwise.

### GetFalsePOk

`func (o *FindingCloseRequest) GetFalsePOk() (*bool, bool)`

GetFalsePOk returns a tuple with the FalseP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseP

`func (o *FindingCloseRequest) SetFalseP(v bool)`

SetFalseP sets FalseP field to given value.

### HasFalseP

`func (o *FindingCloseRequest) HasFalseP() bool`

HasFalseP returns a boolean if a field has been set.

### GetOutOfScope

`func (o *FindingCloseRequest) GetOutOfScope() bool`

GetOutOfScope returns the OutOfScope field if non-nil, zero value otherwise.

### GetOutOfScopeOk

`func (o *FindingCloseRequest) GetOutOfScopeOk() (*bool, bool)`

GetOutOfScopeOk returns a tuple with the OutOfScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfScope

`func (o *FindingCloseRequest) SetOutOfScope(v bool)`

SetOutOfScope sets OutOfScope field to given value.

### HasOutOfScope

`func (o *FindingCloseRequest) HasOutOfScope() bool`

HasOutOfScope returns a boolean if a field has been set.

### GetDuplicate

`func (o *FindingCloseRequest) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *FindingCloseRequest) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *FindingCloseRequest) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.

### HasDuplicate

`func (o *FindingCloseRequest) HasDuplicate() bool`

HasDuplicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


