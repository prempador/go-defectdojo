# PaginatedUserContactInfoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]UserContactInfo**](UserContactInfo.md) |  | 
**Prefetch** | Pointer to [**PaginatedUserContactInfoListPrefetch**](PaginatedUserContactInfoListPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedUserContactInfoList

`func NewPaginatedUserContactInfoList(count int32, results []UserContactInfo, ) *PaginatedUserContactInfoList`

NewPaginatedUserContactInfoList instantiates a new PaginatedUserContactInfoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedUserContactInfoListWithDefaults

`func NewPaginatedUserContactInfoListWithDefaults() *PaginatedUserContactInfoList`

NewPaginatedUserContactInfoListWithDefaults instantiates a new PaginatedUserContactInfoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedUserContactInfoList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedUserContactInfoList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedUserContactInfoList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedUserContactInfoList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedUserContactInfoList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedUserContactInfoList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedUserContactInfoList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedUserContactInfoList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedUserContactInfoList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedUserContactInfoList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedUserContactInfoList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedUserContactInfoList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedUserContactInfoList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedUserContactInfoList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedUserContactInfoList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedUserContactInfoList) GetResults() []UserContactInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedUserContactInfoList) GetResultsOk() (*[]UserContactInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedUserContactInfoList) SetResults(v []UserContactInfo)`

SetResults sets Results field to given value.


### GetPrefetch

`func (o *PaginatedUserContactInfoList) GetPrefetch() PaginatedUserContactInfoListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedUserContactInfoList) GetPrefetchOk() (*PaginatedUserContactInfoListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedUserContactInfoList) SetPrefetch(v PaginatedUserContactInfoListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedUserContactInfoList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


