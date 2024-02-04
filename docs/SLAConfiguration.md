# SLAConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** | A unique name for the set of SLAs. | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Critical** | Pointer to **int32** | number of days to remediate a critical finding. | [optional] 
**High** | Pointer to **int32** | number of days to remediate a high finding. | [optional] 
**Medium** | Pointer to **int32** | number of days to remediate a medium finding. | [optional] 
**Low** | Pointer to **int32** | number of days to remediate a low finding. | [optional] 

## Methods

### NewSLAConfiguration

`func NewSLAConfiguration(id int32, name string, ) *SLAConfiguration`

NewSLAConfiguration instantiates a new SLAConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLAConfigurationWithDefaults

`func NewSLAConfigurationWithDefaults() *SLAConfiguration`

NewSLAConfigurationWithDefaults instantiates a new SLAConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SLAConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLAConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLAConfiguration) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SLAConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLAConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLAConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SLAConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLAConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLAConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLAConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SLAConfiguration) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SLAConfiguration) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCritical

`func (o *SLAConfiguration) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *SLAConfiguration) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *SLAConfiguration) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *SLAConfiguration) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHigh

`func (o *SLAConfiguration) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *SLAConfiguration) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *SLAConfiguration) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *SLAConfiguration) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetMedium

`func (o *SLAConfiguration) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *SLAConfiguration) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *SLAConfiguration) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *SLAConfiguration) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetLow

`func (o *SLAConfiguration) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *SLAConfiguration) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *SLAConfiguration) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *SLAConfiguration) HasLow() bool`

HasLow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


