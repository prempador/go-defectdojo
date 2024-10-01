# PaginatedProductAPIScanConfigurationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]ProductAPIScanConfiguration**](ProductAPIScanConfiguration.md) |  | 
**Prefetch** | Pointer to [**PaginatedProductAPIScanConfigurationListPrefetch**](PaginatedProductAPIScanConfigurationListPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedProductAPIScanConfigurationList

`func NewPaginatedProductAPIScanConfigurationList(count int32, results []ProductAPIScanConfiguration, ) *PaginatedProductAPIScanConfigurationList`

NewPaginatedProductAPIScanConfigurationList instantiates a new PaginatedProductAPIScanConfigurationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductAPIScanConfigurationListWithDefaults

`func NewPaginatedProductAPIScanConfigurationListWithDefaults() *PaginatedProductAPIScanConfigurationList`

NewPaginatedProductAPIScanConfigurationListWithDefaults instantiates a new PaginatedProductAPIScanConfigurationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedProductAPIScanConfigurationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedProductAPIScanConfigurationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedProductAPIScanConfigurationList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedProductAPIScanConfigurationList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedProductAPIScanConfigurationList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedProductAPIScanConfigurationList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedProductAPIScanConfigurationList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedProductAPIScanConfigurationList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedProductAPIScanConfigurationList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedProductAPIScanConfigurationList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedProductAPIScanConfigurationList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedProductAPIScanConfigurationList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedProductAPIScanConfigurationList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedProductAPIScanConfigurationList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedProductAPIScanConfigurationList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedProductAPIScanConfigurationList) GetResults() []ProductAPIScanConfiguration`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedProductAPIScanConfigurationList) GetResultsOk() (*[]ProductAPIScanConfiguration, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedProductAPIScanConfigurationList) SetResults(v []ProductAPIScanConfiguration)`

SetResults sets Results field to given value.


### GetPrefetch

`func (o *PaginatedProductAPIScanConfigurationList) GetPrefetch() PaginatedProductAPIScanConfigurationListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedProductAPIScanConfigurationList) GetPrefetchOk() (*PaginatedProductAPIScanConfigurationListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedProductAPIScanConfigurationList) SetPrefetch(v PaginatedProductAPIScanConfigurationListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedProductAPIScanConfigurationList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


