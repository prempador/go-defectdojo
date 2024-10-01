# ProductAPIScanConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**ServiceKey1** | Pointer to **NullableString** |  | [optional] 
**ServiceKey2** | Pointer to **NullableString** |  | [optional] 
**ServiceKey3** | Pointer to **NullableString** |  | [optional] 
**Product** | **int32** |  | 
**ToolConfiguration** | **int32** |  | 
**Prefetch** | Pointer to [**PaginatedProductAPIScanConfigurationListPrefetch**](PaginatedProductAPIScanConfigurationListPrefetch.md) |  | [optional] 

## Methods

### NewProductAPIScanConfiguration

`func NewProductAPIScanConfiguration(id int32, product int32, toolConfiguration int32, ) *ProductAPIScanConfiguration`

NewProductAPIScanConfiguration instantiates a new ProductAPIScanConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductAPIScanConfigurationWithDefaults

`func NewProductAPIScanConfigurationWithDefaults() *ProductAPIScanConfiguration`

NewProductAPIScanConfigurationWithDefaults instantiates a new ProductAPIScanConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductAPIScanConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductAPIScanConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductAPIScanConfiguration) SetId(v int32)`

SetId sets Id field to given value.


### GetServiceKey1

`func (o *ProductAPIScanConfiguration) GetServiceKey1() string`

GetServiceKey1 returns the ServiceKey1 field if non-nil, zero value otherwise.

### GetServiceKey1Ok

`func (o *ProductAPIScanConfiguration) GetServiceKey1Ok() (*string, bool)`

GetServiceKey1Ok returns a tuple with the ServiceKey1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey1

`func (o *ProductAPIScanConfiguration) SetServiceKey1(v string)`

SetServiceKey1 sets ServiceKey1 field to given value.

### HasServiceKey1

`func (o *ProductAPIScanConfiguration) HasServiceKey1() bool`

HasServiceKey1 returns a boolean if a field has been set.

### SetServiceKey1Nil

`func (o *ProductAPIScanConfiguration) SetServiceKey1Nil(b bool)`

 SetServiceKey1Nil sets the value for ServiceKey1 to be an explicit nil

### UnsetServiceKey1
`func (o *ProductAPIScanConfiguration) UnsetServiceKey1()`

UnsetServiceKey1 ensures that no value is present for ServiceKey1, not even an explicit nil
### GetServiceKey2

`func (o *ProductAPIScanConfiguration) GetServiceKey2() string`

GetServiceKey2 returns the ServiceKey2 field if non-nil, zero value otherwise.

### GetServiceKey2Ok

`func (o *ProductAPIScanConfiguration) GetServiceKey2Ok() (*string, bool)`

GetServiceKey2Ok returns a tuple with the ServiceKey2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey2

`func (o *ProductAPIScanConfiguration) SetServiceKey2(v string)`

SetServiceKey2 sets ServiceKey2 field to given value.

### HasServiceKey2

`func (o *ProductAPIScanConfiguration) HasServiceKey2() bool`

HasServiceKey2 returns a boolean if a field has been set.

### SetServiceKey2Nil

`func (o *ProductAPIScanConfiguration) SetServiceKey2Nil(b bool)`

 SetServiceKey2Nil sets the value for ServiceKey2 to be an explicit nil

### UnsetServiceKey2
`func (o *ProductAPIScanConfiguration) UnsetServiceKey2()`

UnsetServiceKey2 ensures that no value is present for ServiceKey2, not even an explicit nil
### GetServiceKey3

`func (o *ProductAPIScanConfiguration) GetServiceKey3() string`

GetServiceKey3 returns the ServiceKey3 field if non-nil, zero value otherwise.

### GetServiceKey3Ok

`func (o *ProductAPIScanConfiguration) GetServiceKey3Ok() (*string, bool)`

GetServiceKey3Ok returns a tuple with the ServiceKey3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey3

`func (o *ProductAPIScanConfiguration) SetServiceKey3(v string)`

SetServiceKey3 sets ServiceKey3 field to given value.

### HasServiceKey3

`func (o *ProductAPIScanConfiguration) HasServiceKey3() bool`

HasServiceKey3 returns a boolean if a field has been set.

### SetServiceKey3Nil

`func (o *ProductAPIScanConfiguration) SetServiceKey3Nil(b bool)`

 SetServiceKey3Nil sets the value for ServiceKey3 to be an explicit nil

### UnsetServiceKey3
`func (o *ProductAPIScanConfiguration) UnsetServiceKey3()`

UnsetServiceKey3 ensures that no value is present for ServiceKey3, not even an explicit nil
### GetProduct

`func (o *ProductAPIScanConfiguration) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductAPIScanConfiguration) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductAPIScanConfiguration) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetToolConfiguration

`func (o *ProductAPIScanConfiguration) GetToolConfiguration() int32`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *ProductAPIScanConfiguration) GetToolConfigurationOk() (*int32, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *ProductAPIScanConfiguration) SetToolConfiguration(v int32)`

SetToolConfiguration sets ToolConfiguration field to given value.


### GetPrefetch

`func (o *ProductAPIScanConfiguration) GetPrefetch() PaginatedProductAPIScanConfigurationListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *ProductAPIScanConfiguration) GetPrefetchOk() (*PaginatedProductAPIScanConfigurationListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *ProductAPIScanConfiguration) SetPrefetch(v PaginatedProductAPIScanConfigurationListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *ProductAPIScanConfiguration) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


