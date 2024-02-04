# TestToFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **NullableInt32** |  | 
**Files** | [**[]File**](File.md) |  | 

## Methods

### NewTestToFiles

`func NewTestToFiles(testId NullableInt32, files []File, ) *TestToFiles`

NewTestToFiles instantiates a new TestToFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestToFilesWithDefaults

`func NewTestToFilesWithDefaults() *TestToFiles`

NewTestToFilesWithDefaults instantiates a new TestToFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *TestToFiles) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestToFiles) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestToFiles) SetTestId(v int32)`

SetTestId sets TestId field to given value.


### SetTestIdNil

`func (o *TestToFiles) SetTestIdNil(b bool)`

 SetTestIdNil sets the value for TestId to be an explicit nil

### UnsetTestId
`func (o *TestToFiles) UnsetTestId()`

UnsetTestId ensures that no value is present for TestId, not even an explicit nil
### GetFiles

`func (o *TestToFiles) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TestToFiles) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TestToFiles) SetFiles(v []File)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


