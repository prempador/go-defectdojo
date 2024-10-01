# PatchedSLAConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for the set of SLAs. | [optional] 
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

### NewPatchedSLAConfigurationRequest

`func NewPatchedSLAConfigurationRequest() *PatchedSLAConfigurationRequest`

NewPatchedSLAConfigurationRequest instantiates a new PatchedSLAConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSLAConfigurationRequestWithDefaults

`func NewPatchedSLAConfigurationRequestWithDefaults() *PatchedSLAConfigurationRequest`

NewPatchedSLAConfigurationRequestWithDefaults instantiates a new PatchedSLAConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSLAConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSLAConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSLAConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSLAConfigurationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedSLAConfigurationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedSLAConfigurationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedSLAConfigurationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedSLAConfigurationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PatchedSLAConfigurationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PatchedSLAConfigurationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCritical

`func (o *PatchedSLAConfigurationRequest) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *PatchedSLAConfigurationRequest) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *PatchedSLAConfigurationRequest) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *PatchedSLAConfigurationRequest) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetEnforceCritical

`func (o *PatchedSLAConfigurationRequest) GetEnforceCritical() bool`

GetEnforceCritical returns the EnforceCritical field if non-nil, zero value otherwise.

### GetEnforceCriticalOk

`func (o *PatchedSLAConfigurationRequest) GetEnforceCriticalOk() (*bool, bool)`

GetEnforceCriticalOk returns a tuple with the EnforceCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceCritical

`func (o *PatchedSLAConfigurationRequest) SetEnforceCritical(v bool)`

SetEnforceCritical sets EnforceCritical field to given value.

### HasEnforceCritical

`func (o *PatchedSLAConfigurationRequest) HasEnforceCritical() bool`

HasEnforceCritical returns a boolean if a field has been set.

### GetHigh

`func (o *PatchedSLAConfigurationRequest) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *PatchedSLAConfigurationRequest) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *PatchedSLAConfigurationRequest) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *PatchedSLAConfigurationRequest) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetEnforceHigh

`func (o *PatchedSLAConfigurationRequest) GetEnforceHigh() bool`

GetEnforceHigh returns the EnforceHigh field if non-nil, zero value otherwise.

### GetEnforceHighOk

`func (o *PatchedSLAConfigurationRequest) GetEnforceHighOk() (*bool, bool)`

GetEnforceHighOk returns a tuple with the EnforceHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceHigh

`func (o *PatchedSLAConfigurationRequest) SetEnforceHigh(v bool)`

SetEnforceHigh sets EnforceHigh field to given value.

### HasEnforceHigh

`func (o *PatchedSLAConfigurationRequest) HasEnforceHigh() bool`

HasEnforceHigh returns a boolean if a field has been set.

### GetMedium

`func (o *PatchedSLAConfigurationRequest) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *PatchedSLAConfigurationRequest) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *PatchedSLAConfigurationRequest) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *PatchedSLAConfigurationRequest) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetEnforceMedium

`func (o *PatchedSLAConfigurationRequest) GetEnforceMedium() bool`

GetEnforceMedium returns the EnforceMedium field if non-nil, zero value otherwise.

### GetEnforceMediumOk

`func (o *PatchedSLAConfigurationRequest) GetEnforceMediumOk() (*bool, bool)`

GetEnforceMediumOk returns a tuple with the EnforceMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMedium

`func (o *PatchedSLAConfigurationRequest) SetEnforceMedium(v bool)`

SetEnforceMedium sets EnforceMedium field to given value.

### HasEnforceMedium

`func (o *PatchedSLAConfigurationRequest) HasEnforceMedium() bool`

HasEnforceMedium returns a boolean if a field has been set.

### GetLow

`func (o *PatchedSLAConfigurationRequest) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *PatchedSLAConfigurationRequest) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *PatchedSLAConfigurationRequest) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *PatchedSLAConfigurationRequest) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetEnforceLow

`func (o *PatchedSLAConfigurationRequest) GetEnforceLow() bool`

GetEnforceLow returns the EnforceLow field if non-nil, zero value otherwise.

### GetEnforceLowOk

`func (o *PatchedSLAConfigurationRequest) GetEnforceLowOk() (*bool, bool)`

GetEnforceLowOk returns a tuple with the EnforceLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceLow

`func (o *PatchedSLAConfigurationRequest) SetEnforceLow(v bool)`

SetEnforceLow sets EnforceLow field to given value.

### HasEnforceLow

`func (o *PatchedSLAConfigurationRequest) HasEnforceLow() bool`

HasEnforceLow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


