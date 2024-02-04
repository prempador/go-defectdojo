# PaginatedDojoGroupMemberList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]DojoGroupMember**](DojoGroupMember.md) |  | [optional] 
**Prefetch** | Pointer to [**DojoGroupMemberPrefetch**](DojoGroupMemberPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedDojoGroupMemberList

`func NewPaginatedDojoGroupMemberList() *PaginatedDojoGroupMemberList`

NewPaginatedDojoGroupMemberList instantiates a new PaginatedDojoGroupMemberList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDojoGroupMemberListWithDefaults

`func NewPaginatedDojoGroupMemberListWithDefaults() *PaginatedDojoGroupMemberList`

NewPaginatedDojoGroupMemberListWithDefaults instantiates a new PaginatedDojoGroupMemberList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedDojoGroupMemberList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedDojoGroupMemberList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedDojoGroupMemberList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedDojoGroupMemberList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedDojoGroupMemberList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedDojoGroupMemberList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedDojoGroupMemberList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedDojoGroupMemberList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedDojoGroupMemberList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedDojoGroupMemberList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedDojoGroupMemberList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedDojoGroupMemberList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedDojoGroupMemberList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedDojoGroupMemberList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedDojoGroupMemberList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedDojoGroupMemberList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedDojoGroupMemberList) GetResults() []DojoGroupMember`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedDojoGroupMemberList) GetResultsOk() (*[]DojoGroupMember, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedDojoGroupMemberList) SetResults(v []DojoGroupMember)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedDojoGroupMemberList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPrefetch

`func (o *PaginatedDojoGroupMemberList) GetPrefetch() DojoGroupMemberPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedDojoGroupMemberList) GetPrefetchOk() (*DojoGroupMemberPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedDojoGroupMemberList) SetPrefetch(v DojoGroupMemberPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedDojoGroupMemberList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


