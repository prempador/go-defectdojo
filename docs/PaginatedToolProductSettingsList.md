# PaginatedToolProductSettingsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ToolProductSettings**](ToolProductSettings.md) |  | [optional] 

## Methods

### NewPaginatedToolProductSettingsList

`func NewPaginatedToolProductSettingsList() *PaginatedToolProductSettingsList`

NewPaginatedToolProductSettingsList instantiates a new PaginatedToolProductSettingsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedToolProductSettingsListWithDefaults

`func NewPaginatedToolProductSettingsListWithDefaults() *PaginatedToolProductSettingsList`

NewPaginatedToolProductSettingsListWithDefaults instantiates a new PaginatedToolProductSettingsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedToolProductSettingsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedToolProductSettingsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedToolProductSettingsList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedToolProductSettingsList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedToolProductSettingsList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedToolProductSettingsList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedToolProductSettingsList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedToolProductSettingsList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedToolProductSettingsList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedToolProductSettingsList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedToolProductSettingsList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedToolProductSettingsList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedToolProductSettingsList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedToolProductSettingsList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedToolProductSettingsList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedToolProductSettingsList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedToolProductSettingsList) GetResults() []ToolProductSettings`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedToolProductSettingsList) GetResultsOk() (*[]ToolProductSettings, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedToolProductSettingsList) SetResults(v []ToolProductSettings)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedToolProductSettingsList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


