# PaginatedNotificationsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]Notifications**](Notifications.md) |  | 
**Prefetch** | Pointer to [**AppAnalysisPrefetch**](AppAnalysisPrefetch.md) |  | [optional] 

## Methods

### NewPaginatedNotificationsList

`func NewPaginatedNotificationsList(count int32, results []Notifications, ) *PaginatedNotificationsList`

NewPaginatedNotificationsList instantiates a new PaginatedNotificationsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedNotificationsListWithDefaults

`func NewPaginatedNotificationsListWithDefaults() *PaginatedNotificationsList`

NewPaginatedNotificationsListWithDefaults instantiates a new PaginatedNotificationsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedNotificationsList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedNotificationsList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedNotificationsList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedNotificationsList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedNotificationsList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedNotificationsList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedNotificationsList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedNotificationsList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedNotificationsList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedNotificationsList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedNotificationsList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedNotificationsList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedNotificationsList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedNotificationsList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedNotificationsList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedNotificationsList) GetResults() []Notifications`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedNotificationsList) GetResultsOk() (*[]Notifications, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedNotificationsList) SetResults(v []Notifications)`

SetResults sets Results field to given value.


### GetPrefetch

`func (o *PaginatedNotificationsList) GetPrefetch() AppAnalysisPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *PaginatedNotificationsList) GetPrefetchOk() (*AppAnalysisPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *PaginatedNotificationsList) SetPrefetch(v AppAnalysisPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *PaginatedNotificationsList) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


