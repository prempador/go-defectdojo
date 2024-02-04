# EngagementPresets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Title** | Pointer to **string** | Brief description of preset. | [optional] 
**Notes** | Pointer to **NullableString** | Description of what needs to be tested or setting up environment for testing | [optional] 
**Scope** | Pointer to **string** | Scope of Engagement testing, IP&#39;s/Resources/URL&#39;s) | [optional] 
**Created** | **time.Time** |  | [readonly] 
**Product** | **int32** |  | 
**TestType** | Pointer to **[]int32** |  | [optional] 
**NetworkLocations** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEngagementPresets

`func NewEngagementPresets(id int32, created time.Time, product int32, ) *EngagementPresets`

NewEngagementPresets instantiates a new EngagementPresets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngagementPresetsWithDefaults

`func NewEngagementPresetsWithDefaults() *EngagementPresets`

NewEngagementPresetsWithDefaults instantiates a new EngagementPresets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngagementPresets) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngagementPresets) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngagementPresets) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *EngagementPresets) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EngagementPresets) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EngagementPresets) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EngagementPresets) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetNotes

`func (o *EngagementPresets) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EngagementPresets) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EngagementPresets) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EngagementPresets) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *EngagementPresets) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *EngagementPresets) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetScope

`func (o *EngagementPresets) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EngagementPresets) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EngagementPresets) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EngagementPresets) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetCreated

`func (o *EngagementPresets) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EngagementPresets) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EngagementPresets) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetProduct

`func (o *EngagementPresets) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *EngagementPresets) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *EngagementPresets) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetTestType

`func (o *EngagementPresets) GetTestType() []int32`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *EngagementPresets) GetTestTypeOk() (*[]int32, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *EngagementPresets) SetTestType(v []int32)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *EngagementPresets) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### GetNetworkLocations

`func (o *EngagementPresets) GetNetworkLocations() []int32`

GetNetworkLocations returns the NetworkLocations field if non-nil, zero value otherwise.

### GetNetworkLocationsOk

`func (o *EngagementPresets) GetNetworkLocationsOk() (*[]int32, bool)`

GetNetworkLocationsOk returns a tuple with the NetworkLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLocations

`func (o *EngagementPresets) SetNetworkLocations(v []int32)`

SetNetworkLocations sets NetworkLocations field to given value.

### HasNetworkLocations

`func (o *EngagementPresets) HasNetworkLocations() bool`

HasNetworkLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


