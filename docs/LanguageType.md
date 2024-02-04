# LanguageType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Language** | **string** |  | 
**Color** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLanguageType

`func NewLanguageType(id int32, language string, ) *LanguageType`

NewLanguageType instantiates a new LanguageType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageTypeWithDefaults

`func NewLanguageTypeWithDefaults() *LanguageType`

NewLanguageTypeWithDefaults instantiates a new LanguageType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanguageType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanguageType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanguageType) SetId(v int32)`

SetId sets Id field to given value.


### GetLanguage

`func (o *LanguageType) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LanguageType) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LanguageType) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetColor

`func (o *LanguageType) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LanguageType) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LanguageType) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *LanguageType) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *LanguageType) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *LanguageType) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


