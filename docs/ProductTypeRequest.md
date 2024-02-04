# ProductTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CriticalProduct** | Pointer to **bool** |  | [optional] 
**KeyProduct** | Pointer to **bool** |  | [optional] 

## Methods

### NewProductTypeRequest

`func NewProductTypeRequest(name string, ) *ProductTypeRequest`

NewProductTypeRequest instantiates a new ProductTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeRequestWithDefaults

`func NewProductTypeRequestWithDefaults() *ProductTypeRequest`

NewProductTypeRequestWithDefaults instantiates a new ProductTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductTypeRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductTypeRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCriticalProduct

`func (o *ProductTypeRequest) GetCriticalProduct() bool`

GetCriticalProduct returns the CriticalProduct field if non-nil, zero value otherwise.

### GetCriticalProductOk

`func (o *ProductTypeRequest) GetCriticalProductOk() (*bool, bool)`

GetCriticalProductOk returns a tuple with the CriticalProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalProduct

`func (o *ProductTypeRequest) SetCriticalProduct(v bool)`

SetCriticalProduct sets CriticalProduct field to given value.

### HasCriticalProduct

`func (o *ProductTypeRequest) HasCriticalProduct() bool`

HasCriticalProduct returns a boolean if a field has been set.

### GetKeyProduct

`func (o *ProductTypeRequest) GetKeyProduct() bool`

GetKeyProduct returns the KeyProduct field if non-nil, zero value otherwise.

### GetKeyProductOk

`func (o *ProductTypeRequest) GetKeyProductOk() (*bool, bool)`

GetKeyProductOk returns a tuple with the KeyProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProduct

`func (o *ProductTypeRequest) SetKeyProduct(v bool)`

SetKeyProduct sets KeyProduct field to given value.

### HasKeyProduct

`func (o *ProductTypeRequest) HasKeyProduct() bool`

HasKeyProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


