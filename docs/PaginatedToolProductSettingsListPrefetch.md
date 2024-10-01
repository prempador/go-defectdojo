# PaginatedToolProductSettingsListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to [**map[string]Note**](Note.md) |  | [optional] [readonly] 
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**ToolConfiguration** | Pointer to [**map[string]ToolConfiguration**](ToolConfiguration.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedToolProductSettingsListPrefetch

`func NewPaginatedToolProductSettingsListPrefetch() *PaginatedToolProductSettingsListPrefetch`

NewPaginatedToolProductSettingsListPrefetch instantiates a new PaginatedToolProductSettingsListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedToolProductSettingsListPrefetchWithDefaults

`func NewPaginatedToolProductSettingsListPrefetchWithDefaults() *PaginatedToolProductSettingsListPrefetch`

NewPaginatedToolProductSettingsListPrefetchWithDefaults instantiates a new PaginatedToolProductSettingsListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *PaginatedToolProductSettingsListPrefetch) GetNotes() map[string]Note`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PaginatedToolProductSettingsListPrefetch) GetNotesOk() (*map[string]Note, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PaginatedToolProductSettingsListPrefetch) SetNotes(v map[string]Note)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PaginatedToolProductSettingsListPrefetch) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProduct

`func (o *PaginatedToolProductSettingsListPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PaginatedToolProductSettingsListPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PaginatedToolProductSettingsListPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PaginatedToolProductSettingsListPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetToolConfiguration

`func (o *PaginatedToolProductSettingsListPrefetch) GetToolConfiguration() map[string]ToolConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *PaginatedToolProductSettingsListPrefetch) GetToolConfigurationOk() (*map[string]ToolConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *PaginatedToolProductSettingsListPrefetch) SetToolConfiguration(v map[string]ToolConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *PaginatedToolProductSettingsListPrefetch) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


