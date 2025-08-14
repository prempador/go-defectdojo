# PaginatedBurpRawRequestResponseMultiList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]BurpRawRequestResponseMulti**](BurpRawRequestResponseMulti.md) |  | 

## Methods

### NewPaginatedBurpRawRequestResponseMultiList

`func NewPaginatedBurpRawRequestResponseMultiList(count int32, results []BurpRawRequestResponseMulti, ) *PaginatedBurpRawRequestResponseMultiList`

NewPaginatedBurpRawRequestResponseMultiList instantiates a new PaginatedBurpRawRequestResponseMultiList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedBurpRawRequestResponseMultiListWithDefaults

`func NewPaginatedBurpRawRequestResponseMultiListWithDefaults() *PaginatedBurpRawRequestResponseMultiList`

NewPaginatedBurpRawRequestResponseMultiListWithDefaults instantiates a new PaginatedBurpRawRequestResponseMultiList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedBurpRawRequestResponseMultiList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedBurpRawRequestResponseMultiList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedBurpRawRequestResponseMultiList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedBurpRawRequestResponseMultiList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedBurpRawRequestResponseMultiList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedBurpRawRequestResponseMultiList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedBurpRawRequestResponseMultiList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedBurpRawRequestResponseMultiList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedBurpRawRequestResponseMultiList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedBurpRawRequestResponseMultiList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedBurpRawRequestResponseMultiList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedBurpRawRequestResponseMultiList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedBurpRawRequestResponseMultiList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedBurpRawRequestResponseMultiList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedBurpRawRequestResponseMultiList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedBurpRawRequestResponseMultiList) GetResults() []BurpRawRequestResponseMulti`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedBurpRawRequestResponseMultiList) GetResultsOk() (*[]BurpRawRequestResponseMulti, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedBurpRawRequestResponseMultiList) SetResults(v []BurpRawRequestResponseMulti)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


