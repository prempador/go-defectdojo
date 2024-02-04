# FindingToFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FindingId** | **NullableInt32** |  | 
**Files** | [**[]File**](File.md) |  | 

## Methods

### NewFindingToFiles

`func NewFindingToFiles(findingId NullableInt32, files []File, ) *FindingToFiles`

NewFindingToFiles instantiates a new FindingToFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingToFilesWithDefaults

`func NewFindingToFilesWithDefaults() *FindingToFiles`

NewFindingToFilesWithDefaults instantiates a new FindingToFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindingId

`func (o *FindingToFiles) GetFindingId() int32`

GetFindingId returns the FindingId field if non-nil, zero value otherwise.

### GetFindingIdOk

`func (o *FindingToFiles) GetFindingIdOk() (*int32, bool)`

GetFindingIdOk returns a tuple with the FindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingId

`func (o *FindingToFiles) SetFindingId(v int32)`

SetFindingId sets FindingId field to given value.


### SetFindingIdNil

`func (o *FindingToFiles) SetFindingIdNil(b bool)`

 SetFindingIdNil sets the value for FindingId to be an explicit nil

### UnsetFindingId
`func (o *FindingToFiles) UnsetFindingId()`

UnsetFindingId ensures that no value is present for FindingId, not even an explicit nil
### GetFiles

`func (o *FindingToFiles) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FindingToFiles) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FindingToFiles) SetFiles(v []File)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


