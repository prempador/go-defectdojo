# EngagementToFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngagementId** | **NullableInt32** |  | 
**Files** | [**[]File**](File.md) |  | 

## Methods

### NewEngagementToFiles

`func NewEngagementToFiles(engagementId NullableInt32, files []File, ) *EngagementToFiles`

NewEngagementToFiles instantiates a new EngagementToFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementToFilesWithDefaults

`func NewEngagementToFilesWithDefaults() *EngagementToFiles`

NewEngagementToFilesWithDefaults instantiates a new EngagementToFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagementId

`func (o *EngagementToFiles) GetEngagementId() int32`

GetEngagementId returns the EngagementId field if non-nil, zero value otherwise.

### GetEngagementIdOk

`func (o *EngagementToFiles) GetEngagementIdOk() (*int32, bool)`

GetEngagementIdOk returns a tuple with the EngagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementId

`func (o *EngagementToFiles) SetEngagementId(v int32)`

SetEngagementId sets EngagementId field to given value.


### SetEngagementIdNil

`func (o *EngagementToFiles) SetEngagementIdNil(b bool)`

 SetEngagementIdNil sets the value for EngagementId to be an explicit nil

### UnsetEngagementId
`func (o *EngagementToFiles) UnsetEngagementId()`

UnsetEngagementId ensures that no value is present for EngagementId, not even an explicit nil
### GetFiles

`func (o *EngagementToFiles) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *EngagementToFiles) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *EngagementToFiles) SetFiles(v []File)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


