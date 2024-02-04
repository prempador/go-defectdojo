# PaginatedProductTypeGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ProductTypeGroup**](ProductTypeGroup.md) |  | [optional] 
**Prefetch** | Pointer to [**PaginatedProductTypeGroupListPrefetch**](PaginatedProductTypeGroupListPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedProductTypeGroupList

`func NewPaginatedProductTypeGroupList() *PaginatedProductTypeGroupList`

NewPaginatedProductTypeGroupList instantiates a new PaginatedProductTypeGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductTypeGroupListWithDefaults

`func NewPaginatedProductTypeGroupListWithDefaults() *PaginatedProductTypeGroupList`

NewPaginatedProductTypeGroupListWithDefaults instantiates a new PaginatedProductTypeGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedProductTypeGroupList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedProductTypeGroupList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedProductTypeGroupList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedProductTypeGroupList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedProductTypeGroupList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedProductTypeGroupList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedProductTypeGroupList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedProductTypeGroupList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedProductTypeGroupList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedProductTypeGroupList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedProductTypeGroupList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedProductTypeGroupList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedProductTypeGroupList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedProductTypeGroupList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedProductTypeGroupList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedProductTypeGroupList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedProductTypeGroupList) GetResults() []ProductTypeGroup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedProductTypeGroupList) GetResultsOk() (*[]ProductTypeGroup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedProductTypeGroupList) SetResults(v []ProductTypeGroup)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedProductTypeGroupList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPrefetch

`func (o *PaginatedProductTypeGroupList) GetPrefetch() PaginatedProductTypeGroupListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedProductTypeGroupList) GetPrefetchOk() (*PaginatedProductTypeGroupListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedProductTypeGroupList) SetPrefetch(v PaginatedProductTypeGroupListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedProductTypeGroupList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


