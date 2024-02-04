# PaginatedTestImportListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FindingsAffected** | Pointer to [**map[string]Finding**](Finding.md) |  | [optional] [readonly] 
**Test** | Pointer to [**map[string]Test**](Test.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedTestImportListPrefetch

`func NewPaginatedTestImportListPrefetch() *PaginatedTestImportListPrefetch`

NewPaginatedTestImportListPrefetch instantiates a new PaginatedTestImportListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedTestImportListPrefetchWithDefaults

`func NewPaginatedTestImportListPrefetchWithDefaults() *PaginatedTestImportListPrefetch`

NewPaginatedTestImportListPrefetchWithDefaults instantiates a new PaginatedTestImportListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindingsAffected

`func (o *PaginatedTestImportListPrefetch) GetFindingsAffected() map[string]Finding`

GetFindingsAffected returns the FindingsAffected field if non-nil, zero value otherwise.

### GetFindingsAffectedOk

`func (o *PaginatedTestImportListPrefetch) GetFindingsAffectedOk() (*map[string]Finding, bool)`

GetFindingsAffectedOk returns a tuple with the FindingsAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingsAffected

`func (o *PaginatedTestImportListPrefetch) SetFindingsAffected(v map[string]Finding)`

SetFindingsAffected sets FindingsAffected field to given value.

### HasFindingsAffected

`func (o *PaginatedTestImportListPrefetch) HasFindingsAffected() bool`

HasFindingsAffected returns a boolean if a field has been set.

### GetTest

`func (o *PaginatedTestImportListPrefetch) GetTest() map[string]Test`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PaginatedTestImportListPrefetch) GetTestOk() (*map[string]Test, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PaginatedTestImportListPrefetch) SetTest(v map[string]Test)`

SetTest sets Test field to given value.

### HasTest

`func (o *PaginatedTestImportListPrefetch) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


