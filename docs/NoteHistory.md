# NoteHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**CurrentEditor** | [**UserStub**](UserStub.md) |  | [readonly] 
**NoteType** | [**NoteType**](NoteType.md) |  | [readonly] 
**Data** | **string** |  | 
**Time** | **NullableTime** |  | [readonly] 

## Methods

### NewNoteHistory

`func NewNoteHistory(id int32, currentEditor UserStub, noteType NoteType, data string, time NullableTime, ) *NoteHistory`

NewNoteHistory instantiates a new NoteHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteHistoryWithDefaults

`func NewNoteHistoryWithDefaults() *NoteHistory`

NewNoteHistoryWithDefaults instantiates a new NoteHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteHistory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteHistory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteHistory) SetId(v int32)`

SetId sets Id field to given value.


### GetCurrentEditor

`func (o *NoteHistory) GetCurrentEditor() UserStub`

GetCurrentEditor returns the CurrentEditor field if non-nil, zero value otherwise.

### GetCurrentEditorOk

`func (o *NoteHistory) GetCurrentEditorOk() (*UserStub, bool)`

GetCurrentEditorOk returns a tuple with the CurrentEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentEditor

`func (o *NoteHistory) SetCurrentEditor(v UserStub)`

SetCurrentEditor sets CurrentEditor field to given value.


### GetNoteType

`func (o *NoteHistory) GetNoteType() NoteType`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *NoteHistory) GetNoteTypeOk() (*NoteType, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *NoteHistory) SetNoteType(v NoteType)`

SetNoteType sets NoteType field to given value.


### GetData

`func (o *NoteHistory) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NoteHistory) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NoteHistory) SetData(v string)`

SetData sets Data field to given value.


### GetTime

`func (o *NoteHistory) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *NoteHistory) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *NoteHistory) SetTime(v time.Time)`

SetTime sets Time field to given value.


### SetTimeNil

`func (o *NoteHistory) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *NoteHistory) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


