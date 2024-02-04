# NoteTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**IsSingle** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsMandatory** | Pointer to **bool** |  | [optional] 

## Methods

### NewNoteTypeRequest

`func NewNoteTypeRequest(name string, description string, ) *NoteTypeRequest`

NewNoteTypeRequest instantiates a new NoteTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteTypeRequestWithDefaults

`func NewNoteTypeRequestWithDefaults() *NoteTypeRequest`

NewNoteTypeRequestWithDefaults instantiates a new NoteTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NoteTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NoteTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NoteTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NoteTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NoteTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NoteTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsSingle

`func (o *NoteTypeRequest) GetIsSingle() bool`

GetIsSingle returns the IsSingle field if non-nil, zero value otherwise.

### GetIsSingleOk

`func (o *NoteTypeRequest) GetIsSingleOk() (*bool, bool)`

GetIsSingleOk returns a tuple with the IsSingle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingle

`func (o *NoteTypeRequest) SetIsSingle(v bool)`

SetIsSingle sets IsSingle field to given value.

### HasIsSingle

`func (o *NoteTypeRequest) HasIsSingle() bool`

HasIsSingle returns a boolean if a field has been set.

### GetIsActive

`func (o *NoteTypeRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *NoteTypeRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *NoteTypeRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *NoteTypeRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsMandatory

`func (o *NoteTypeRequest) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *NoteTypeRequest) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *NoteTypeRequest) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *NoteTypeRequest) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


