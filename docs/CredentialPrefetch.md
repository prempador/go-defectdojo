# CredentialPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**map[string]FindingEnvironment**](FindingEnvironment.md) |  | [optional] [readonly] 
**Notes** | Pointer to [**map[string]Note**](Note.md) |  | [optional] [readonly] 

## Methods

### NewCredentialPrefetch

`func NewCredentialPrefetch() *CredentialPrefetch`

NewCredentialPrefetch instantiates a new CredentialPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialPrefetchWithDefaults

`func NewCredentialPrefetchWithDefaults() *CredentialPrefetch`

NewCredentialPrefetchWithDefaults instantiates a new CredentialPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CredentialPrefetch) GetEnvironment() map[string]FindingEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialPrefetch) GetEnvironmentOk() (*map[string]FindingEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialPrefetch) SetEnvironment(v map[string]FindingEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialPrefetch) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetNotes

`func (o *CredentialPrefetch) GetNotes() map[string]Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CredentialPrefetch) GetNotesOk() (*map[string]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CredentialPrefetch) SetNotes(v map[string]Note)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CredentialPrefetch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


