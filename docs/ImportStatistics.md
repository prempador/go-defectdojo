# ImportStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to [**SeverityStatusStatistics**](SeverityStatusStatistics.md) | Finding statistics as stored in Defect Dojo before the import | [optional] 
**Delta** | Pointer to [**DeltaStatistics**](DeltaStatistics.md) | Finding statistics of modifications made by the reimport. Only available when TRACK_IMPORT_HISTORY hass not disabled. | [optional] 
**After** | [**SeverityStatusStatistics**](SeverityStatusStatistics.md) | Finding statistics as stored in Defect Dojo after the import | 

## Methods

### NewImportStatistics

`func NewImportStatistics(after SeverityStatusStatistics, ) *ImportStatistics`

NewImportStatistics instantiates a new ImportStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportStatisticsWithDefaults

`func NewImportStatisticsWithDefaults() *ImportStatistics`

NewImportStatisticsWithDefaults instantiates a new ImportStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *ImportStatistics) GetBefore() SeverityStatusStatistics`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ImportStatistics) GetBeforeOk() (*SeverityStatusStatistics, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ImportStatistics) SetBefore(v SeverityStatusStatistics)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ImportStatistics) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetDelta

`func (o *ImportStatistics) GetDelta() DeltaStatistics`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *ImportStatistics) GetDeltaOk() (*DeltaStatistics, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *ImportStatistics) SetDelta(v DeltaStatistics)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *ImportStatistics) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetAfter

`func (o *ImportStatistics) GetAfter() SeverityStatusStatistics`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ImportStatistics) GetAfterOk() (*SeverityStatusStatistics, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ImportStatistics) SetAfter(v SeverityStatusStatistics)`

SetAfter sets After field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


