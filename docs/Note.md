# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Author** | [**UserStub**](UserStub.md) |  | [readonly] 
**Editor** | [**NullableUserStub**](UserStub.md) |  | [readonly] 
**History** | [**[]NoteHistory**](NoteHistory.md) |  | [readonly] 
**NoteType** | [**NoteType**](NoteType.md) |  | [readonly] 
**Entry** | **string** |  | 
**Date** | **time.Time** |  | [readonly] 
**Private** | Pointer to **bool** |  | [optional] 
**Edited** | Pointer to **bool** |  | [optional] 
**EditTime** | **NullableTime** |  | [readonly] 

## Methods

### NewNote

`func NewNote(id int32, author UserStub, editor NullableUserStub, history []NoteHistory, noteType NoteType, entry string, date time.Time, editTime NullableTime, ) *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Note) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Note) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Note) SetId(v int32)`

SetId sets Id field to given value.


### GetAuthor

`func (o *Note) GetAuthor() UserStub`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Note) GetAuthorOk() (*UserStub, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Note) SetAuthor(v UserStub)`

SetAuthor sets Author field to given value.


### GetEditor

`func (o *Note) GetEditor() UserStub`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *Note) GetEditorOk() (*UserStub, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *Note) SetEditor(v UserStub)`

SetEditor sets Editor field to given value.


### SetEditorNil

`func (o *Note) SetEditorNil(b bool)`

 SetEditorNil sets the value for Editor to be an explicit nil

### UnsetEditor
`func (o *Note) UnsetEditor()`

UnsetEditor ensures that no value is present for Editor, not even an explicit nil
### GetHistory

`func (o *Note) GetHistory() []NoteHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *Note) GetHistoryOk() (*[]NoteHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *Note) SetHistory(v []NoteHistory)`

SetHistory sets History field to given value.


### GetNoteType

`func (o *Note) GetNoteType() NoteType`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *Note) GetNoteTypeOk() (*NoteType, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *Note) SetNoteType(v NoteType)`

SetNoteType sets NoteType field to given value.


### GetEntry

`func (o *Note) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *Note) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *Note) SetEntry(v string)`

SetEntry sets Entry field to given value.


### GetDate

`func (o *Note) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Note) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Note) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetPrivate

`func (o *Note) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Note) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Note) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Note) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetEdited

`func (o *Note) GetEdited() bool`

GetEdited returns the Edited field if non-nil, zero value otherwise.

### GetEditedOk

`func (o *Note) GetEditedOk() (*bool, bool)`

GetEditedOk returns a tuple with the Edited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdited

`func (o *Note) SetEdited(v bool)`

SetEdited sets Edited field to given value.

### HasEdited

`func (o *Note) HasEdited() bool`

HasEdited returns a boolean if a field has been set.

### GetEditTime

`func (o *Note) GetEditTime() time.Time`

GetEditTime returns the EditTime field if non-nil, zero value otherwise.

### GetEditTimeOk

`func (o *Note) GetEditTimeOk() (*time.Time, bool)`

GetEditTimeOk returns a tuple with the EditTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditTime

`func (o *Note) SetEditTime(v time.Time)`

SetEditTime sets EditTime field to given value.


### SetEditTimeNil

`func (o *Note) SetEditTimeNil(b bool)`

 SetEditTimeNil sets the value for EditTime to be an explicit nil

### UnsetEditTime
`func (o *Note) UnsetEditTime()`

UnsetEditTime ensures that no value is present for EditTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


