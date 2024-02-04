# PatchedProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProdNumericGrade** | Pointer to **NullableInt32** |  | [optional] 
**BusinessCriticality** | Pointer to **NullableString** | * &#x60;very high&#x60; - Very High * &#x60;high&#x60; - High * &#x60;medium&#x60; - Medium * &#x60;low&#x60; - Low * &#x60;very low&#x60; - Very Low * &#x60;none&#x60; - None | [optional] 
**Platform** | Pointer to **NullableString** | * &#x60;web service&#x60; - API * &#x60;desktop&#x60; - Desktop * &#x60;iot&#x60; - Internet of Things * &#x60;mobile&#x60; - Mobile * &#x60;web&#x60; - Web | [optional] 
**Lifecycle** | Pointer to **NullableString** | * &#x60;construction&#x60; - Construction * &#x60;production&#x60; - Production * &#x60;retirement&#x60; - Retirement | [optional] 
**Origin** | Pointer to **NullableString** | * &#x60;third party library&#x60; - Third Party Library * &#x60;purchased&#x60; - Purchased * &#x60;contractor&#x60; - Contractor Developed * &#x60;internal&#x60; - Internally Developed * &#x60;open source&#x60; - Open Source * &#x60;outsourced&#x60; - Outsourced | [optional] 
**UserRecords** | Pointer to **NullableInt32** | Estimate the number of user records within the application. | [optional] 
**Revenue** | Pointer to **NullableFloat64** | Estimate the application&#39;s revenue. | [optional] 
**ExternalAudience** | Pointer to **bool** | Specify if the application is used by people outside the organization. | [optional] 
**InternetAccessible** | Pointer to **bool** | Specify if the application is accessible from the public internet. | [optional] 
**EnableProductTagInheritance** | Pointer to **bool** | Enables product tag inheritance. Any tags added on a product will automatically be added to all Engagements, Tests, and Findings | [optional] 
**EnableSimpleRiskAcceptance** | Pointer to **bool** | Allows simple risk acceptance by checking/unchecking a checkbox. | [optional] 
**EnableFullRiskAcceptance** | Pointer to **bool** | Allows full risk acceptance using a risk acceptance form, expiration date, uploaded proof, etc. | [optional] 
**DisableSlaBreachNotifications** | Pointer to **bool** | Disable SLA breach notifications if configured in the global settings | [optional] 
**ProductManager** | Pointer to **NullableInt32** |  | [optional] 
**TechnicalContact** | Pointer to **NullableInt32** |  | [optional] 
**TeamManager** | Pointer to **NullableInt32** |  | [optional] 
**ProdType** | Pointer to **int32** |  | [optional] 
**SlaConfiguration** | Pointer to **int32** |  | [optional] 
**Regulations** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedProductRequest

`func NewPatchedProductRequest() *PatchedProductRequest`

NewPatchedProductRequest instantiates a new PatchedProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProductRequestWithDefaults

`func NewPatchedProductRequestWithDefaults() *PatchedProductRequest`

NewPatchedProductRequestWithDefaults instantiates a new PatchedProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PatchedProductRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedProductRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedProductRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedProductRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *PatchedProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProductRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProductRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProdNumericGrade

`func (o *PatchedProductRequest) GetProdNumericGrade() int32`

GetProdNumericGrade returns the ProdNumericGrade field if non-nil, zero value otherwise.

### GetProdNumericGradeOk

`func (o *PatchedProductRequest) GetProdNumericGradeOk() (*int32, bool)`

GetProdNumericGradeOk returns a tuple with the ProdNumericGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdNumericGrade

`func (o *PatchedProductRequest) SetProdNumericGrade(v int32)`

SetProdNumericGrade sets ProdNumericGrade field to given value.

### HasProdNumericGrade

`func (o *PatchedProductRequest) HasProdNumericGrade() bool`

HasProdNumericGrade returns a boolean if a field has been set.

### SetProdNumericGradeNil

`func (o *PatchedProductRequest) SetProdNumericGradeNil(b bool)`

 SetProdNumericGradeNil sets the value for ProdNumericGrade to be an explicit nil

### UnsetProdNumericGrade
`func (o *PatchedProductRequest) UnsetProdNumericGrade()`

UnsetProdNumericGrade ensures that no value is present for ProdNumericGrade, not even an explicit nil
### GetBusinessCriticality

`func (o *PatchedProductRequest) GetBusinessCriticality() string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *PatchedProductRequest) GetBusinessCriticalityOk() (*string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *PatchedProductRequest) SetBusinessCriticality(v string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *PatchedProductRequest) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### SetBusinessCriticalityNil

`func (o *PatchedProductRequest) SetBusinessCriticalityNil(b bool)`

 SetBusinessCriticalityNil sets the value for BusinessCriticality to be an explicit nil

### UnsetBusinessCriticality
`func (o *PatchedProductRequest) UnsetBusinessCriticality()`

UnsetBusinessCriticality ensures that no value is present for BusinessCriticality, not even an explicit nil
### GetPlatform

`func (o *PatchedProductRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedProductRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedProductRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedProductRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedProductRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedProductRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLifecycle

`func (o *PatchedProductRequest) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *PatchedProductRequest) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *PatchedProductRequest) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *PatchedProductRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### SetLifecycleNil

`func (o *PatchedProductRequest) SetLifecycleNil(b bool)`

 SetLifecycleNil sets the value for Lifecycle to be an explicit nil

### UnsetLifecycle
`func (o *PatchedProductRequest) UnsetLifecycle()`

UnsetLifecycle ensures that no value is present for Lifecycle, not even an explicit nil
### GetOrigin

`func (o *PatchedProductRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PatchedProductRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PatchedProductRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PatchedProductRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *PatchedProductRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *PatchedProductRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetUserRecords

`func (o *PatchedProductRequest) GetUserRecords() int32`

GetUserRecords returns the UserRecords field if non-nil, zero value otherwise.

### GetUserRecordsOk

`func (o *PatchedProductRequest) GetUserRecordsOk() (*int32, bool)`

GetUserRecordsOk returns a tuple with the UserRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRecords

`func (o *PatchedProductRequest) SetUserRecords(v int32)`

SetUserRecords sets UserRecords field to given value.

### HasUserRecords

`func (o *PatchedProductRequest) HasUserRecords() bool`

HasUserRecords returns a boolean if a field has been set.

### SetUserRecordsNil

`func (o *PatchedProductRequest) SetUserRecordsNil(b bool)`

 SetUserRecordsNil sets the value for UserRecords to be an explicit nil

### UnsetUserRecords
`func (o *PatchedProductRequest) UnsetUserRecords()`

UnsetUserRecords ensures that no value is present for UserRecords, not even an explicit nil
### GetRevenue

`func (o *PatchedProductRequest) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *PatchedProductRequest) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *PatchedProductRequest) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *PatchedProductRequest) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *PatchedProductRequest) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *PatchedProductRequest) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetExternalAudience

`func (o *PatchedProductRequest) GetExternalAudience() bool`

GetExternalAudience returns the ExternalAudience field if non-nil, zero value otherwise.

### GetExternalAudienceOk

`func (o *PatchedProductRequest) GetExternalAudienceOk() (*bool, bool)`

GetExternalAudienceOk returns a tuple with the ExternalAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAudience

`func (o *PatchedProductRequest) SetExternalAudience(v bool)`

SetExternalAudience sets ExternalAudience field to given value.

### HasExternalAudience

`func (o *PatchedProductRequest) HasExternalAudience() bool`

HasExternalAudience returns a boolean if a field has been set.

### GetInternetAccessible

`func (o *PatchedProductRequest) GetInternetAccessible() bool`

GetInternetAccessible returns the InternetAccessible field if non-nil, zero value otherwise.

### GetInternetAccessibleOk

`func (o *PatchedProductRequest) GetInternetAccessibleOk() (*bool, bool)`

GetInternetAccessibleOk returns a tuple with the InternetAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccessible

`func (o *PatchedProductRequest) SetInternetAccessible(v bool)`

SetInternetAccessible sets InternetAccessible field to given value.

### HasInternetAccessible

`func (o *PatchedProductRequest) HasInternetAccessible() bool`

HasInternetAccessible returns a boolean if a field has been set.

### GetEnableProductTagInheritance

`func (o *PatchedProductRequest) GetEnableProductTagInheritance() bool`

GetEnableProductTagInheritance returns the EnableProductTagInheritance field if non-nil, zero value otherwise.

### GetEnableProductTagInheritanceOk

`func (o *PatchedProductRequest) GetEnableProductTagInheritanceOk() (*bool, bool)`

GetEnableProductTagInheritanceOk returns a tuple with the EnableProductTagInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductTagInheritance

`func (o *PatchedProductRequest) SetEnableProductTagInheritance(v bool)`

SetEnableProductTagInheritance sets EnableProductTagInheritance field to given value.

### HasEnableProductTagInheritance

`func (o *PatchedProductRequest) HasEnableProductTagInheritance() bool`

HasEnableProductTagInheritance returns a boolean if a field has been set.

### GetEnableSimpleRiskAcceptance

`func (o *PatchedProductRequest) GetEnableSimpleRiskAcceptance() bool`

GetEnableSimpleRiskAcceptance returns the EnableSimpleRiskAcceptance field if non-nil, zero value otherwise.

### GetEnableSimpleRiskAcceptanceOk

`func (o *PatchedProductRequest) GetEnableSimpleRiskAcceptanceOk() (*bool, bool)`

GetEnableSimpleRiskAcceptanceOk returns a tuple with the EnableSimpleRiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSimpleRiskAcceptance

`func (o *PatchedProductRequest) SetEnableSimpleRiskAcceptance(v bool)`

SetEnableSimpleRiskAcceptance sets EnableSimpleRiskAcceptance field to given value.

### HasEnableSimpleRiskAcceptance

`func (o *PatchedProductRequest) HasEnableSimpleRiskAcceptance() bool`

HasEnableSimpleRiskAcceptance returns a boolean if a field has been set.

### GetEnableFullRiskAcceptance

`func (o *PatchedProductRequest) GetEnableFullRiskAcceptance() bool`

GetEnableFullRiskAcceptance returns the EnableFullRiskAcceptance field if non-nil, zero value otherwise.

### GetEnableFullRiskAcceptanceOk

`func (o *PatchedProductRequest) GetEnableFullRiskAcceptanceOk() (*bool, bool)`

GetEnableFullRiskAcceptanceOk returns a tuple with the EnableFullRiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFullRiskAcceptance

`func (o *PatchedProductRequest) SetEnableFullRiskAcceptance(v bool)`

SetEnableFullRiskAcceptance sets EnableFullRiskAcceptance field to given value.

### HasEnableFullRiskAcceptance

`func (o *PatchedProductRequest) HasEnableFullRiskAcceptance() bool`

HasEnableFullRiskAcceptance returns a boolean if a field has been set.

### GetDisableSlaBreachNotifications

`func (o *PatchedProductRequest) GetDisableSlaBreachNotifications() bool`

GetDisableSlaBreachNotifications returns the DisableSlaBreachNotifications field if non-nil, zero value otherwise.

### GetDisableSlaBreachNotificationsOk

`func (o *PatchedProductRequest) GetDisableSlaBreachNotificationsOk() (*bool, bool)`

GetDisableSlaBreachNotificationsOk returns a tuple with the DisableSlaBreachNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSlaBreachNotifications

`func (o *PatchedProductRequest) SetDisableSlaBreachNotifications(v bool)`

SetDisableSlaBreachNotifications sets DisableSlaBreachNotifications field to given value.

### HasDisableSlaBreachNotifications

`func (o *PatchedProductRequest) HasDisableSlaBreachNotifications() bool`

HasDisableSlaBreachNotifications returns a boolean if a field has been set.

### GetProductManager

`func (o *PatchedProductRequest) GetProductManager() int32`

GetProductManager returns the ProductManager field if non-nil, zero value otherwise.

### GetProductManagerOk

`func (o *PatchedProductRequest) GetProductManagerOk() (*int32, bool)`

GetProductManagerOk returns a tuple with the ProductManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductManager

`func (o *PatchedProductRequest) SetProductManager(v int32)`

SetProductManager sets ProductManager field to given value.

### HasProductManager

`func (o *PatchedProductRequest) HasProductManager() bool`

HasProductManager returns a boolean if a field has been set.

### SetProductManagerNil

`func (o *PatchedProductRequest) SetProductManagerNil(b bool)`

 SetProductManagerNil sets the value for ProductManager to be an explicit nil

### UnsetProductManager
`func (o *PatchedProductRequest) UnsetProductManager()`

UnsetProductManager ensures that no value is present for ProductManager, not even an explicit nil
### GetTechnicalContact

`func (o *PatchedProductRequest) GetTechnicalContact() int32`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *PatchedProductRequest) GetTechnicalContactOk() (*int32, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *PatchedProductRequest) SetTechnicalContact(v int32)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *PatchedProductRequest) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### SetTechnicalContactNil

`func (o *PatchedProductRequest) SetTechnicalContactNil(b bool)`

 SetTechnicalContactNil sets the value for TechnicalContact to be an explicit nil

### UnsetTechnicalContact
`func (o *PatchedProductRequest) UnsetTechnicalContact()`

UnsetTechnicalContact ensures that no value is present for TechnicalContact, not even an explicit nil
### GetTeamManager

`func (o *PatchedProductRequest) GetTeamManager() int32`

GetTeamManager returns the TeamManager field if non-nil, zero value otherwise.

### GetTeamManagerOk

`func (o *PatchedProductRequest) GetTeamManagerOk() (*int32, bool)`

GetTeamManagerOk returns a tuple with the TeamManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManager

`func (o *PatchedProductRequest) SetTeamManager(v int32)`

SetTeamManager sets TeamManager field to given value.

### HasTeamManager

`func (o *PatchedProductRequest) HasTeamManager() bool`

HasTeamManager returns a boolean if a field has been set.

### SetTeamManagerNil

`func (o *PatchedProductRequest) SetTeamManagerNil(b bool)`

 SetTeamManagerNil sets the value for TeamManager to be an explicit nil

### UnsetTeamManager
`func (o *PatchedProductRequest) UnsetTeamManager()`

UnsetTeamManager ensures that no value is present for TeamManager, not even an explicit nil
### GetProdType

`func (o *PatchedProductRequest) GetProdType() int32`

GetProdType returns the ProdType field if non-nil, zero value otherwise.

### GetProdTypeOk

`func (o *PatchedProductRequest) GetProdTypeOk() (*int32, bool)`

GetProdTypeOk returns a tuple with the ProdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdType

`func (o *PatchedProductRequest) SetProdType(v int32)`

SetProdType sets ProdType field to given value.

### HasProdType

`func (o *PatchedProductRequest) HasProdType() bool`

HasProdType returns a boolean if a field has been set.

### GetSlaConfiguration

`func (o *PatchedProductRequest) GetSlaConfiguration() int32`

GetSlaConfiguration returns the SlaConfiguration field if non-nil, zero value otherwise.

### GetSlaConfigurationOk

`func (o *PatchedProductRequest) GetSlaConfigurationOk() (*int32, bool)`

GetSlaConfigurationOk returns a tuple with the SlaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaConfiguration

`func (o *PatchedProductRequest) SetSlaConfiguration(v int32)`

SetSlaConfiguration sets SlaConfiguration field to given value.

### HasSlaConfiguration

`func (o *PatchedProductRequest) HasSlaConfiguration() bool`

HasSlaConfiguration returns a boolean if a field has been set.

### GetRegulations

`func (o *PatchedProductRequest) GetRegulations() []int32`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *PatchedProductRequest) GetRegulationsOk() (*[]int32, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *PatchedProductRequest) SetRegulations(v []int32)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *PatchedProductRequest) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


