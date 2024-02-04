# LanguagePrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**map[string]LanguageType**](LanguageType.md) |  | [optional] [readonly] 
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**User** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewLanguagePrefetch

`func NewLanguagePrefetch() *LanguagePrefetch`

NewLanguagePrefetch instantiates a new LanguagePrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguagePrefetchWithDefaults

`func NewLanguagePrefetchWithDefaults() *LanguagePrefetch`

NewLanguagePrefetchWithDefaults instantiates a new LanguagePrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *LanguagePrefetch) GetLanguage() map[string]LanguageType`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LanguagePrefetch) GetLanguageOk() (*map[string]LanguageType, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LanguagePrefetch) SetLanguage(v map[string]LanguageType)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *LanguagePrefetch) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetProduct

`func (o *LanguagePrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *LanguagePrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *LanguagePrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *LanguagePrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetUser

`func (o *LanguagePrefetch) GetUser() map[string]UserStub`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LanguagePrefetch) GetUserOk() (*map[string]UserStub, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LanguagePrefetch) SetUser(v map[string]UserStub)`

SetUser sets User field to given value.

### HasUser

`func (o *LanguagePrefetch) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


