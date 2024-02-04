# EngagementPresetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Brief description of preset. | [optional] 
**Notes** | Pointer to **NullableString** | Description of what needs to be tested or setting up environment for testing | [optional] 
**Scope** | Pointer to **string** | Scope of Engagement testing, IP&#39;s/Resources/URL&#39;s) | [optional] 
**Product** | **int32** |  | 
**TestType** | Pointer to **[]int32** |  | [optional] 
**NetworkLocations** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEngagementPresetsRequest

`func NewEngagementPresetsRequest(product int32, ) *EngagementPresetsRequest`

NewEngagementPresetsRequest instantiates a new EngagementPresetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementPresetsRequestWithDefaults

`func NewEngagementPresetsRequestWithDefaults() *EngagementPresetsRequest`

NewEngagementPresetsRequestWithDefaults instantiates a new EngagementPresetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *EngagementPresetsRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementPresetsRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementPresetsRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EngagementPresetsRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetNotes

`func (o *EngagementPresetsRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EngagementPresetsRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EngagementPresetsRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EngagementPresetsRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *EngagementPresetsRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *EngagementPresetsRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetScope

`func (o *EngagementPresetsRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EngagementPresetsRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EngagementPresetsRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EngagementPresetsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetProduct

`func (o *EngagementPresetsRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EngagementPresetsRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EngagementPresetsRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetTestType

`func (o *EngagementPresetsRequest) GetTestType() []int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *EngagementPresetsRequest) GetTestTypeOk() (*[]int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *EngagementPresetsRequest) SetTestType(v []int32)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *EngagementPresetsRequest) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### GetNetworkLocations

`func (o *EngagementPresetsRequest) GetNetworkLocations() []int32`

GetNetworkLocations returns the NetworkLocations field if non-nil, zero value otherwise.

### GetNetworkLocationsOk

`func (o *EngagementPresetsRequest) GetNetworkLocationsOk() (*[]int32, bool)`

GetNetworkLocationsOk returns a tuple with the NetworkLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLocations

`func (o *EngagementPresetsRequest) SetNetworkLocations(v []int32)`

SetNetworkLocations sets NetworkLocations field to given value.

### HasNetworkLocations

`func (o *EngagementPresetsRequest) HasNetworkLocations() bool`

HasNetworkLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


