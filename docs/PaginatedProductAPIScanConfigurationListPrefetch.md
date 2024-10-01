# PaginatedProductAPIScanConfigurationListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**ToolConfiguration** | Pointer to [**map[string]ToolConfiguration**](ToolConfiguration.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedProductAPIScanConfigurationListPrefetch

`func NewPaginatedProductAPIScanConfigurationListPrefetch() *PaginatedProductAPIScanConfigurationListPrefetch`

NewPaginatedProductAPIScanConfigurationListPrefetch instantiates a new PaginatedProductAPIScanConfigurationListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductAPIScanConfigurationListPrefetchWithDefaults

`func NewPaginatedProductAPIScanConfigurationListPrefetchWithDefaults() *PaginatedProductAPIScanConfigurationListPrefetch`

NewPaginatedProductAPIScanConfigurationListPrefetchWithDefaults instantiates a new PaginatedProductAPIScanConfigurationListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetToolConfiguration

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) GetToolConfiguration() map[string]ToolConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) GetToolConfigurationOk() (*map[string]ToolConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) SetToolConfiguration(v map[string]ToolConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *PaginatedProductAPIScanConfigurationListPrefetch) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


