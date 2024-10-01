# SLAConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the set of SLAs. | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Critical** | Pointer to **int32** | The number of days to remediate a critical finding. | [optional] 
**EnforceCritical** | Pointer to **bool** | When enabled, critical findings will be assigned an SLA expiration date based on the critical finding SLA days within this SLA configuration. | [optional] 
**High** | Pointer to **int32** | The number of days to remediate a high finding. | [optional] 
**EnforceHigh** | Pointer to **bool** | When enabled, high findings will be assigned an SLA expiration date based on the high finding SLA days within this SLA configuration. | [optional] 
**Medium** | Pointer to **int32** | The number of days to remediate a medium finding. | [optional] 
**EnforceMedium** | Pointer to **bool** | When enabled, medium findings will be assigned an SLA expiration date based on the medium finding SLA days within this SLA configuration. | [optional] 
**Low** | Pointer to **int32** | The number of days to remediate a low finding. | [optional] 
**EnforceLow** | Pointer to **bool** | When enabled, low findings will be assigned an SLA expiration date based on the low finding SLA days within this SLA configuration. | [optional] 

## Methods

### NewSLAConfigurationRequest

`func NewSLAConfigurationRequest(name string, ) *SLAConfigurationRequest`

NewSLAConfigurationRequest instantiates a new SLAConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLAConfigurationRequestWithDefaults

`func NewSLAConfigurationRequestWithDefaults() *SLAConfigurationRequest`

NewSLAConfigurationRequestWithDefaults instantiates a new SLAConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SLAConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLAConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLAConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SLAConfigurationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLAConfigurationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLAConfigurationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLAConfigurationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SLAConfigurationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SLAConfigurationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCritical

`func (o *SLAConfigurationRequest) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *SLAConfigurationRequest) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *SLAConfigurationRequest) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *SLAConfigurationRequest) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetEnforceCritical

`func (o *SLAConfigurationRequest) GetEnforceCritical() bool`

GetEnforceCritical returns the EnforceCritical field if non-nil, zero value otherwise.

### GetEnforceCriticalOk

`func (o *SLAConfigurationRequest) GetEnforceCriticalOk() (*bool, bool)`

GetEnforceCriticalOk returns a tuple with the EnforceCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceCritical

`func (o *SLAConfigurationRequest) SetEnforceCritical(v bool)`

SetEnforceCritical sets EnforceCritical field to given value.

### HasEnforceCritical

`func (o *SLAConfigurationRequest) HasEnforceCritical() bool`

HasEnforceCritical returns a boolean if a field has been set.

### GetHigh

`func (o *SLAConfigurationRequest) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *SLAConfigurationRequest) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *SLAConfigurationRequest) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *SLAConfigurationRequest) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetEnforceHigh

`func (o *SLAConfigurationRequest) GetEnforceHigh() bool`

GetEnforceHigh returns the EnforceHigh field if non-nil, zero value otherwise.

### GetEnforceHighOk

`func (o *SLAConfigurationRequest) GetEnforceHighOk() (*bool, bool)`

GetEnforceHighOk returns a tuple with the EnforceHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceHigh

`func (o *SLAConfigurationRequest) SetEnforceHigh(v bool)`

SetEnforceHigh sets EnforceHigh field to given value.

### HasEnforceHigh

`func (o *SLAConfigurationRequest) HasEnforceHigh() bool`

HasEnforceHigh returns a boolean if a field has been set.

### GetMedium

`func (o *SLAConfigurationRequest) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *SLAConfigurationRequest) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *SLAConfigurationRequest) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *SLAConfigurationRequest) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetEnforceMedium

`func (o *SLAConfigurationRequest) GetEnforceMedium() bool`

GetEnforceMedium returns the EnforceMedium field if non-nil, zero value otherwise.

### GetEnforceMediumOk

`func (o *SLAConfigurationRequest) GetEnforceMediumOk() (*bool, bool)`

GetEnforceMediumOk returns a tuple with the EnforceMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMedium

`func (o *SLAConfigurationRequest) SetEnforceMedium(v bool)`

SetEnforceMedium sets EnforceMedium field to given value.

### HasEnforceMedium

`func (o *SLAConfigurationRequest) HasEnforceMedium() bool`

HasEnforceMedium returns a boolean if a field has been set.

### GetLow

`func (o *SLAConfigurationRequest) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *SLAConfigurationRequest) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *SLAConfigurationRequest) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *SLAConfigurationRequest) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetEnforceLow

`func (o *SLAConfigurationRequest) GetEnforceLow() bool`

GetEnforceLow returns the EnforceLow field if non-nil, zero value otherwise.

### GetEnforceLowOk

`func (o *SLAConfigurationRequest) GetEnforceLowOk() (*bool, bool)`

GetEnforceLowOk returns a tuple with the EnforceLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceLow

`func (o *SLAConfigurationRequest) SetEnforceLow(v bool)`

SetEnforceLow sets EnforceLow field to given value.

### HasEnforceLow

`func (o *SLAConfigurationRequest) HasEnforceLow() bool`

HasEnforceLow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


