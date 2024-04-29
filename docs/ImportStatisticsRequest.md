# ImportStatisticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to [**SeverityStatusStatisticsRequest**](SeverityStatusStatisticsRequest.md) | Finding statistics as stored in Defect Dojo before the import | [optional] 
**Delta** | Pointer to [**DeltaStatisticsRequest**](DeltaStatisticsRequest.md) | Finding statistics of modifications made by the reimport. Only available when TRACK_IMPORT_HISTORY hass not disabled. | [optional] 
**After** | [**SeverityStatusStatisticsRequest**](SeverityStatusStatisticsRequest.md) | Finding statistics as stored in Defect Dojo after the import | 

## Methods

### NewImportStatisticsRequest

`func NewImportStatisticsRequest(after SeverityStatusStatisticsRequest, ) *ImportStatisticsRequest`

NewImportStatisticsRequest instantiates a new ImportStatisticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportStatisticsRequestWithDefaults

`func NewImportStatisticsRequestWithDefaults() *ImportStatisticsRequest`

NewImportStatisticsRequestWithDefaults instantiates a new ImportStatisticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *ImportStatisticsRequest) GetBefore() SeverityStatusStatisticsRequest`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ImportStatisticsRequest) GetBeforeOk() (*SeverityStatusStatisticsRequest, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ImportStatisticsRequest) SetBefore(v SeverityStatusStatisticsRequest)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ImportStatisticsRequest) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetDelta

`func (o *ImportStatisticsRequest) GetDelta() DeltaStatisticsRequest`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *ImportStatisticsRequest) GetDeltaOk() (*DeltaStatisticsRequest, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *ImportStatisticsRequest) SetDelta(v DeltaStatisticsRequest)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *ImportStatisticsRequest) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetAfter

`func (o *ImportStatisticsRequest) GetAfter() SeverityStatusStatisticsRequest`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ImportStatisticsRequest) GetAfterOk() (*SeverityStatusStatisticsRequest, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ImportStatisticsRequest) SetAfter(v SeverityStatusStatisticsRequest)`

SetAfter sets After field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


