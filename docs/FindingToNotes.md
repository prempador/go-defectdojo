# FindingToNotes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FindingId** | **NullableInt32** |  | 
**Notes** | [**[]Note**](Note.md) |  | 

## Methods

### NewFindingToNotes

`func NewFindingToNotes(findingId NullableInt32, notes []Note, ) *FindingToNotes`

NewFindingToNotes instantiates a new FindingToNotes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingToNotesWithDefaults

`func NewFindingToNotesWithDefaults() *FindingToNotes`

NewFindingToNotesWithDefaults instantiates a new FindingToNotes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFindingId

`func (o *FindingToNotes) GetFindingId() int32`

GetFindingId returns the FindingId field if non-nil, zero value otherwise.

### GetFindingIdOk

`func (o *FindingToNotes) GetFindingIdOk() (*int32, bool)`

GetFindingIdOk returns a tuple with the FindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingId

`func (o *FindingToNotes) SetFindingId(v int32)`

SetFindingId sets FindingId field to given value.


### SetFindingIdNil

`func (o *FindingToNotes) SetFindingIdNil(b bool)`

 SetFindingIdNil sets the value for FindingId to be an explicit nil

### UnsetFindingId
`func (o *FindingToNotes) UnsetFindingId()`

UnsetFindingId ensures that no value is present for FindingId, not even an explicit nil
### GetNotes

`func (o *FindingToNotes) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *FindingToNotes) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *FindingToNotes) SetNotes(v []Note)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


