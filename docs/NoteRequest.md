# NoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | **string** |  | 
**Private** | Pointer to **bool** |  | [optional] 
**Edited** | Pointer to **bool** |  | [optional] 

## Methods

### NewNoteRequest

`func NewNoteRequest(entry string, ) *NoteRequest`

NewNoteRequest instantiates a new NoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteRequestWithDefaults

`func NewNoteRequestWithDefaults() *NoteRequest`

NewNoteRequestWithDefaults instantiates a new NoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *NoteRequest) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *NoteRequest) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *NoteRequest) SetEntry(v string)`

SetEntry sets Entry field to given value.


### GetPrivate

`func (o *NoteRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *NoteRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *NoteRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *NoteRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetEdited

`func (o *NoteRequest) GetEdited() bool`

GetEdited returns the Edited field if non-nil, zero value otherwise.

### GetEditedOk

`func (o *NoteRequest) GetEditedOk() (*bool, bool)`

GetEditedOk returns a tuple with the Edited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdited

`func (o *NoteRequest) SetEdited(v bool)`

SetEdited sets Edited field to given value.

### HasEdited

`func (o *NoteRequest) HasEdited() bool`

HasEdited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


