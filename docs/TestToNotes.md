# TestToNotes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestId** | **NullableInt32** |  | 
**Notes** | [**[]Note**](Note.md) |  | 

## Methods

### NewTestToNotes

`func NewTestToNotes(testId NullableInt32, notes []Note, ) *TestToNotes`

NewTestToNotes instantiates a new TestToNotes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestToNotesWithDefaults

`func NewTestToNotesWithDefaults() *TestToNotes`

NewTestToNotesWithDefaults instantiates a new TestToNotes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestId

`func (o *TestToNotes) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *TestToNotes) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *TestToNotes) SetTestId(v int32)`

SetTestId sets TestId field to given value.


### SetTestIdNil

`func (o *TestToNotes) SetTestIdNil(b bool)`

 SetTestIdNil sets the value for TestId to be an explicit nil

### UnsetTestId
`func (o *TestToNotes) UnsetTestId()`

UnsetTestId ensures that no value is present for TestId, not even an explicit nil
### GetNotes

`func (o *TestToNotes) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TestToNotes) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TestToNotes) SetNotes(v []Note)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


