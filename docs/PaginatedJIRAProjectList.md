# PaginatedJIRAProjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]JIRAProject**](JIRAProject.md) |  | 
**Prefetch** | Pointer to [**JIRAProjectPrefetch**](JIRAProjectPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedJIRAProjectList

`func NewPaginatedJIRAProjectList(count int32, results []JIRAProject, ) *PaginatedJIRAProjectList`

NewPaginatedJIRAProjectList instantiates a new PaginatedJIRAProjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedJIRAProjectListWithDefaults

`func NewPaginatedJIRAProjectListWithDefaults() *PaginatedJIRAProjectList`

NewPaginatedJIRAProjectListWithDefaults instantiates a new PaginatedJIRAProjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedJIRAProjectList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedJIRAProjectList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedJIRAProjectList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedJIRAProjectList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedJIRAProjectList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedJIRAProjectList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedJIRAProjectList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedJIRAProjectList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedJIRAProjectList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedJIRAProjectList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedJIRAProjectList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedJIRAProjectList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedJIRAProjectList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedJIRAProjectList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedJIRAProjectList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedJIRAProjectList) GetResults() []JIRAProject`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedJIRAProjectList) GetResultsOk() (*[]JIRAProject, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedJIRAProjectList) SetResults(v []JIRAProject)`

SetResults sets Results field to given value.


### GetPrefetch

`func (o *PaginatedJIRAProjectList) GetPrefetch() JIRAProjectPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedJIRAProjectList) GetPrefetchOk() (*JIRAProjectPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedJIRAProjectList) SetPrefetch(v JIRAProjectPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedJIRAProjectList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


