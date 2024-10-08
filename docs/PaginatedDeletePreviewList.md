# PaginatedDeletePreviewList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]DeletePreview**](DeletePreview.md) |  | 

## Methods

### NewPaginatedDeletePreviewList

`func NewPaginatedDeletePreviewList(count int32, results []DeletePreview, ) *PaginatedDeletePreviewList`

NewPaginatedDeletePreviewList instantiates a new PaginatedDeletePreviewList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDeletePreviewListWithDefaults

`func NewPaginatedDeletePreviewListWithDefaults() *PaginatedDeletePreviewList`

NewPaginatedDeletePreviewListWithDefaults instantiates a new PaginatedDeletePreviewList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedDeletePreviewList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedDeletePreviewList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedDeletePreviewList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedDeletePreviewList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedDeletePreviewList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedDeletePreviewList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedDeletePreviewList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedDeletePreviewList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedDeletePreviewList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedDeletePreviewList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedDeletePreviewList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedDeletePreviewList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedDeletePreviewList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedDeletePreviewList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedDeletePreviewList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedDeletePreviewList) GetResults() []DeletePreview`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedDeletePreviewList) GetResultsOk() (*[]DeletePreview, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedDeletePreviewList) SetResults(v []DeletePreview)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


