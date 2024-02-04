# Language

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Files** | Pointer to **NullableInt32** |  | [optional] 
**Blank** | Pointer to **NullableInt32** |  | [optional] 
**Comment** | Pointer to **NullableInt32** |  | [optional] 
**Code** | Pointer to **NullableInt32** |  | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Language** | **int32** |  | 
**Product** | **int32** |  | 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Prefetch** | Pointer to [**LanguagePrefetch**](LanguagePrefetch.md) |  | [optional] 

## Methods

### NewLanguage

`func NewLanguage(id int32, created time.Time, language int32, product int32, ) *Language`

NewLanguage instantiates a new Language object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageWithDefaults

`func NewLanguageWithDefaults() *Language`

NewLanguageWithDefaults instantiates a new Language object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Language) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Language) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Language) SetId(v int32)`

SetId sets Id field to given value.


### GetFiles

`func (o *Language) GetFiles() int32`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Language) GetFilesOk() (*int32, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Language) SetFiles(v int32)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Language) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *Language) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *Language) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetBlank

`func (o *Language) GetBlank() int32`

GetBlank returns the Blank field if non-nil, zero value otherwise.

### GetBlankOk

`func (o *Language) GetBlankOk() (*int32, bool)`

GetBlankOk returns a tuple with the Blank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlank

`func (o *Language) SetBlank(v int32)`

SetBlank sets Blank field to given value.

### HasBlank

`func (o *Language) HasBlank() bool`

HasBlank returns a boolean if a field has been set.

### SetBlankNil

`func (o *Language) SetBlankNil(b bool)`

 SetBlankNil sets the value for Blank to be an explicit nil

### UnsetBlank
`func (o *Language) UnsetBlank()`

UnsetBlank ensures that no value is present for Blank, not even an explicit nil
### GetComment

`func (o *Language) GetComment() int32`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Language) GetCommentOk() (*int32, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Language) SetComment(v int32)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Language) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Language) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Language) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCode

`func (o *Language) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Language) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Language) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Language) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Language) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Language) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCreated

`func (o *Language) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Language) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Language) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLanguage

`func (o *Language) GetLanguage() int32`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Language) GetLanguageOk() (*int32, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Language) SetLanguage(v int32)`

SetLanguage sets Language field to given value.


### GetProduct

`func (o *Language) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Language) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Language) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetUser

`func (o *Language) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Language) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Language) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *Language) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Language) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Language) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPrefetch

`func (o *Language) GetPrefetch() LanguagePrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *Language) GetPrefetchOk() (*LanguagePrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *Language) SetPrefetch(v LanguagePrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *Language) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


