# PatchedProductTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CriticalProduct** | Pointer to **bool** |  | [optional] 
**KeyProduct** | Pointer to **bool** |  | [optional] 

## Methods

### NewPatchedProductTypeRequest

`func NewPatchedProductTypeRequest() *PatchedProductTypeRequest`

NewPatchedProductTypeRequest instantiates a new PatchedProductTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProductTypeRequestWithDefaults

`func NewPatchedProductTypeRequestWithDefaults() *PatchedProductTypeRequest`

NewPatchedProductTypeRequestWithDefaults instantiates a new PatchedProductTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedProductTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProductTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProductTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProductTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedProductTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedProductTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedProductTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedProductTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedProductTypeRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedProductTypeRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCriticalProduct

`func (o *PatchedProductTypeRequest) GetCriticalProduct() bool`

GetCriticalProduct returns the CriticalProduct field if non-nil, zero value otherwise.

### GetCriticalProductOk

`func (o *PatchedProductTypeRequest) GetCriticalProductOk() (*bool, bool)`

GetCriticalProductOk returns a tuple with the CriticalProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalProduct

`func (o *PatchedProductTypeRequest) SetCriticalProduct(v bool)`

SetCriticalProduct sets CriticalProduct field to given value.

### HasCriticalProduct

`func (o *PatchedProductTypeRequest) HasCriticalProduct() bool`

HasCriticalProduct returns a boolean if a field has been set.

### GetKeyProduct

`func (o *PatchedProductTypeRequest) GetKeyProduct() bool`

GetKeyProduct returns the KeyProduct field if non-nil, zero value otherwise.

### GetKeyProductOk

`func (o *PatchedProductTypeRequest) GetKeyProductOk() (*bool, bool)`

GetKeyProductOk returns a tuple with the KeyProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProduct

`func (o *PatchedProductTypeRequest) SetKeyProduct(v bool)`

SetKeyProduct sets KeyProduct field to given value.

### HasKeyProduct

`func (o *PatchedProductTypeRequest) HasKeyProduct() bool`

HasKeyProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


