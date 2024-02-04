# PaginatedProductTypeMemberList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ProductTypeMember**](ProductTypeMember.md) |  | [optional] 
**Prefetch** | Pointer to [**PaginatedProductTypeMemberListPrefetch**](PaginatedProductTypeMemberListPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedProductTypeMemberList

`func NewPaginatedProductTypeMemberList() *PaginatedProductTypeMemberList`

NewPaginatedProductTypeMemberList instantiates a new PaginatedProductTypeMemberList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductTypeMemberListWithDefaults

`func NewPaginatedProductTypeMemberListWithDefaults() *PaginatedProductTypeMemberList`

NewPaginatedProductTypeMemberListWithDefaults instantiates a new PaginatedProductTypeMemberList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedProductTypeMemberList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedProductTypeMemberList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedProductTypeMemberList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedProductTypeMemberList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedProductTypeMemberList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedProductTypeMemberList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedProductTypeMemberList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedProductTypeMemberList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedProductTypeMemberList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedProductTypeMemberList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedProductTypeMemberList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedProductTypeMemberList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedProductTypeMemberList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedProductTypeMemberList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedProductTypeMemberList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedProductTypeMemberList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedProductTypeMemberList) GetResults() []ProductTypeMember`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedProductTypeMemberList) GetResultsOk() (*[]ProductTypeMember, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedProductTypeMemberList) SetResults(v []ProductTypeMember)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedProductTypeMemberList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPrefetch

`func (o *PaginatedProductTypeMemberList) GetPrefetch() PaginatedProductTypeMemberListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedProductTypeMemberList) GetPrefetchOk() (*PaginatedProductTypeMemberListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedProductTypeMemberList) SetPrefetch(v PaginatedProductTypeMemberListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedProductTypeMemberList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


