# ProductType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CriticalProduct** | Pointer to **bool** |  | [optional] 
**KeyProduct** | Pointer to **bool** |  | [optional] 
**Updated** | **NullableTime** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**Members** | **[]int32** |  | [readonly] 
**AuthorizationGroups** | **[]int32** |  | [readonly] 
**Prefetch** | Pointer to [**PaginatedProductTypeListPrefetch**](PaginatedProductTypeListPrefetch.md) |  | [optional] 

## Methods

### NewProductType

`func NewProductType(id int32, name string, updated NullableTime, created NullableTime, members []int32, authorizationGroups []int32, ) *ProductType`

NewProductType instantiates a new ProductType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeWithDefaults

`func NewProductTypeWithDefaults() *ProductType`

NewProductTypeWithDefaults instantiates a new ProductType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductType) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProductType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCriticalProduct

`func (o *ProductType) GetCriticalProduct() bool`

GetCriticalProduct returns the CriticalProduct field if non-nil, zero value otherwise.

### GetCriticalProductOk

`func (o *ProductType) GetCriticalProductOk() (*bool, bool)`

GetCriticalProductOk returns a tuple with the CriticalProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalProduct

`func (o *ProductType) SetCriticalProduct(v bool)`

SetCriticalProduct sets CriticalProduct field to given value.

### HasCriticalProduct

`func (o *ProductType) HasCriticalProduct() bool`

HasCriticalProduct returns a boolean if a field has been set.

### GetKeyProduct

`func (o *ProductType) GetKeyProduct() bool`

GetKeyProduct returns the KeyProduct field if non-nil, zero value otherwise.

### GetKeyProductOk

`func (o *ProductType) GetKeyProductOk() (*bool, bool)`

GetKeyProductOk returns a tuple with the KeyProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProduct

`func (o *ProductType) SetKeyProduct(v bool)`

SetKeyProduct sets KeyProduct field to given value.

### HasKeyProduct

`func (o *ProductType) HasKeyProduct() bool`

HasKeyProduct returns a boolean if a field has been set.

### GetUpdated

`func (o *ProductType) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProductType) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProductType) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### SetUpdatedNil

`func (o *ProductType) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *ProductType) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetCreated

`func (o *ProductType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProductType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProductType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ProductType) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ProductType) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetMembers

`func (o *ProductType) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ProductType) GetMembersOk() (*[]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ProductType) SetMembers(v []int32)`

SetMembers sets Members field to given value.


### GetAuthorizationGroups

`func (o *ProductType) GetAuthorizationGroups() []int32`

GetAuthorizationGroups returns the AuthorizationGroups field if non-nil, zero value otherwise.

### GetAuthorizationGroupsOk

`func (o *ProductType) GetAuthorizationGroupsOk() (*[]int32, bool)`

GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroups

`func (o *ProductType) SetAuthorizationGroups(v []int32)`

SetAuthorizationGroups sets AuthorizationGroups field to given value.


### GetPrefetch

`func (o *ProductType) GetPrefetch() PaginatedProductTypeListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ProductType) GetPrefetchOk() (*PaginatedProductTypeListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ProductType) SetPrefetch(v PaginatedProductTypeListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ProductType) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


