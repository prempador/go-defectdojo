# AddNewNoteOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | **string** |  | 
**Private** | Pointer to **bool** |  | [optional] 
**NoteType** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAddNewNoteOptionRequest

`func NewAddNewNoteOptionRequest(entry string, ) *AddNewNoteOptionRequest`

NewAddNewNoteOptionRequest instantiates a new AddNewNoteOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNewNoteOptionRequestWithDefaults

`func NewAddNewNoteOptionRequestWithDefaults() *AddNewNoteOptionRequest`

NewAddNewNoteOptionRequestWithDefaults instantiates a new AddNewNoteOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *AddNewNoteOptionRequest) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *AddNewNoteOptionRequest) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *AddNewNoteOptionRequest) SetEntry(v string)`

SetEntry sets Entry field to given value.


### GetPrivate

`func (o *AddNewNoteOptionRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AddNewNoteOptionRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AddNewNoteOptionRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *AddNewNoteOptionRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetNoteType

`func (o *AddNewNoteOptionRequest) GetNoteType() int32`

GetNoteType returns the NoteType field if non-nil, zero value otherwise.

### GetNoteTypeOk

`func (o *AddNewNoteOptionRequest) GetNoteTypeOk() (*int32, bool)`

GetNoteTypeOk returns a tuple with the NoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteType

`func (o *AddNewNoteOptionRequest) SetNoteType(v int32)`

SetNoteType sets NoteType field to given value.

### HasNoteType

`func (o *AddNewNoteOptionRequest) HasNoteType() bool`

HasNoteType returns a boolean if a field has been set.

### SetNoteTypeNil

`func (o *AddNewNoteOptionRequest) SetNoteTypeNil(b bool)`

 SetNoteTypeNil sets the value for NoteType to be an explicit nil

### UnsetNoteType
`func (o *AddNewNoteOptionRequest) UnsetNoteType()`

UnsetNoteType ensures that no value is present for NoteType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


