# NoteType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**IsSingle** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsMandatory** | Pointer to **bool** |  | [optional] 

## Methods

### NewNoteType

`func NewNoteType(id int32, name string, description string, ) *NoteType`

NewNoteType instantiates a new NoteType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteTypeWithDefaults

`func NewNoteTypeWithDefaults() *NoteType`

NewNoteTypeWithDefaults instantiates a new NoteType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NoteType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NoteType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NoteType) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *NoteType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NoteType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoteType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoteType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsSingle

`func (o *NoteType) GetIsSingle() bool`

GetIsSingle returns the IsSingle field if non-nil, zero value otherwise.

### GetIsSingleOk

`func (o *NoteType) GetIsSingleOk() (*bool, bool)`

GetIsSingleOk returns a tuple with the IsSingle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingle

`func (o *NoteType) SetIsSingle(v bool)`

SetIsSingle sets IsSingle field to given value.

### HasIsSingle

`func (o *NoteType) HasIsSingle() bool`

HasIsSingle returns a boolean if a field has been set.

### GetIsActive

`func (o *NoteType) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NoteType) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NoteType) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *NoteType) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsMandatory

`func (o *NoteType) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *NoteType) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *NoteType) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *NoteType) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


