# LanguageTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **string** |  | 
**Color** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLanguageTypeRequest

`func NewLanguageTypeRequest(language string, ) *LanguageTypeRequest`

NewLanguageTypeRequest instantiates a new LanguageTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageTypeRequestWithDefaults

`func NewLanguageTypeRequestWithDefaults() *LanguageTypeRequest`

NewLanguageTypeRequestWithDefaults instantiates a new LanguageTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *LanguageTypeRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LanguageTypeRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LanguageTypeRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetColor

`func (o *LanguageTypeRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LanguageTypeRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LanguageTypeRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *LanguageTypeRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *LanguageTypeRequest) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *LanguageTypeRequest) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


