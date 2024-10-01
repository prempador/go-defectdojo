# PaginatedCredentialList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]Credential**](Credential.md) |  | 
**Prefetch** | Pointer to [**CredentialPrefetch**](CredentialPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedCredentialList

`func NewPaginatedCredentialList(count int32, results []Credential, ) *PaginatedCredentialList`

NewPaginatedCredentialList instantiates a new PaginatedCredentialList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCredentialListWithDefaults

`func NewPaginatedCredentialListWithDefaults() *PaginatedCredentialList`

NewPaginatedCredentialListWithDefaults instantiates a new PaginatedCredentialList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedCredentialList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedCredentialList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedCredentialList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedCredentialList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedCredentialList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedCredentialList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedCredentialList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedCredentialList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedCredentialList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedCredentialList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedCredentialList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedCredentialList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedCredentialList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedCredentialList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedCredentialList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedCredentialList) GetResults() []Credential`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCredentialList) GetResultsOk() (*[]Credential, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCredentialList) SetResults(v []Credential)`

SetResults sets Results field to given value.


### GetPrefetch

`func (o *PaginatedCredentialList) GetPrefetch() CredentialPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedCredentialList) GetPrefetchOk() (*CredentialPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedCredentialList) SetPrefetch(v CredentialPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedCredentialList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


