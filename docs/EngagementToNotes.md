# EngagementToNotes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngagementId** | **NullableInt32** |  | 
**Notes** | [**[]Note**](Note.md) |  | 

## Methods

### NewEngagementToNotes

`func NewEngagementToNotes(engagementId NullableInt32, notes []Note, ) *EngagementToNotes`

NewEngagementToNotes instantiates a new EngagementToNotes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementToNotesWithDefaults

`func NewEngagementToNotesWithDefaults() *EngagementToNotes`

NewEngagementToNotesWithDefaults instantiates a new EngagementToNotes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagementId

`func (o *EngagementToNotes) GetEngagementId() int32`

GetEngagementId returns the EngagementId field if non-nil, zero value otherwise.

### GetEngagementIdOk

`func (o *EngagementToNotes) GetEngagementIdOk() (*int32, bool)`

GetEngagementIdOk returns a tuple with the EngagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementId

`func (o *EngagementToNotes) SetEngagementId(v int32)`

SetEngagementId sets EngagementId field to given value.


### SetEngagementIdNil

`func (o *EngagementToNotes) SetEngagementIdNil(b bool)`

 SetEngagementIdNil sets the value for EngagementId to be an explicit nil

### UnsetEngagementId
`func (o *EngagementToNotes) UnsetEngagementId()`

UnsetEngagementId ensures that no value is present for EngagementId, not even an explicit nil
### GetNotes

`func (o *EngagementToNotes) GetNotes() []Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EngagementToNotes) GetNotesOk() (*[]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EngagementToNotes) SetNotes(v []Note)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


