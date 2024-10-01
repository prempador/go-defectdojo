# PaginatedAppAnalysisList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]AppAnalysis**](AppAnalysis.md) |  | 
**Prefetch** | Pointer to [**AppAnalysisPrefetch**](AppAnalysisPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedAppAnalysisList

`func NewPaginatedAppAnalysisList(count int32, results []AppAnalysis, ) *PaginatedAppAnalysisList`

NewPaginatedAppAnalysisList instantiates a new PaginatedAppAnalysisList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedAppAnalysisListWithDefaults

`func NewPaginatedAppAnalysisListWithDefaults() *PaginatedAppAnalysisList`

NewPaginatedAppAnalysisListWithDefaults instantiates a new PaginatedAppAnalysisList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedAppAnalysisList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedAppAnalysisList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedAppAnalysisList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedAppAnalysisList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedAppAnalysisList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedAppAnalysisList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedAppAnalysisList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedAppAnalysisList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedAppAnalysisList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedAppAnalysisList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedAppAnalysisList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedAppAnalysisList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedAppAnalysisList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedAppAnalysisList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedAppAnalysisList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedAppAnalysisList) GetResults() []AppAnalysis`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedAppAnalysisList) GetResultsOk() (*[]AppAnalysis, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedAppAnalysisList) SetResults(v []AppAnalysis)`

SetResults sets Results field to given value.


### GetPrefetch

`func (o *PaginatedAppAnalysisList) GetPrefetch() AppAnalysisPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedAppAnalysisList) GetPrefetchOk() (*AppAnalysisPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedAppAnalysisList) SetPrefetch(v AppAnalysisPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedAppAnalysisList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


