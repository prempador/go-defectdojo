# ProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**Description** | **string** |  | 
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
**ProdType** | **int32** |  | 
**SlaConfiguration** | Pointer to **int32** |  | [optional] 
**Regulations** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewProductRequest

`func NewProductRequest(name string, description string, prodType int32, ) *ProductRequest`

NewProductRequest instantiates a new ProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRequestWithDefaults

`func NewProductRequestWithDefaults() *ProductRequest`

NewProductRequestWithDefaults instantiates a new ProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ProductRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *ProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProdNumericGrade

`func (o *ProductRequest) GetProdNumericGrade() int32`

GetProdNumericGrade returns the ProdNumericGrade field if non-nil, zero value otherwise.

### GetProdNumericGradeOk

`func (o *ProductRequest) GetProdNumericGradeOk() (*int32, bool)`

GetProdNumericGradeOk returns a tuple with the ProdNumericGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdNumericGrade

`func (o *ProductRequest) SetProdNumericGrade(v int32)`

SetProdNumericGrade sets ProdNumericGrade field to given value.

### HasProdNumericGrade

`func (o *ProductRequest) HasProdNumericGrade() bool`

HasProdNumericGrade returns a boolean if a field has been set.

### SetProdNumericGradeNil

`func (o *ProductRequest) SetProdNumericGradeNil(b bool)`

 SetProdNumericGradeNil sets the value for ProdNumericGrade to be an explicit nil

### UnsetProdNumericGrade
`func (o *ProductRequest) UnsetProdNumericGrade()`

UnsetProdNumericGrade ensures that no value is present for ProdNumericGrade, not even an explicit nil
### GetBusinessCriticality

`func (o *ProductRequest) GetBusinessCriticality() string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *ProductRequest) GetBusinessCriticalityOk() (*string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *ProductRequest) SetBusinessCriticality(v string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *ProductRequest) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### SetBusinessCriticalityNil

`func (o *ProductRequest) SetBusinessCriticalityNil(b bool)`

 SetBusinessCriticalityNil sets the value for BusinessCriticality to be an explicit nil

### UnsetBusinessCriticality
`func (o *ProductRequest) UnsetBusinessCriticality()`

UnsetBusinessCriticality ensures that no value is present for BusinessCriticality, not even an explicit nil
### GetPlatform

`func (o *ProductRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ProductRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ProductRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ProductRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ProductRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ProductRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLifecycle

`func (o *ProductRequest) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *ProductRequest) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *ProductRequest) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *ProductRequest) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### SetLifecycleNil

`func (o *ProductRequest) SetLifecycleNil(b bool)`

 SetLifecycleNil sets the value for Lifecycle to be an explicit nil

### UnsetLifecycle
`func (o *ProductRequest) UnsetLifecycle()`

UnsetLifecycle ensures that no value is present for Lifecycle, not even an explicit nil
### GetOrigin

`func (o *ProductRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ProductRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ProductRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ProductRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ProductRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ProductRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetUserRecords

`func (o *ProductRequest) GetUserRecords() int32`

GetUserRecords returns the UserRecords field if non-nil, zero value otherwise.

### GetUserRecordsOk

`func (o *ProductRequest) GetUserRecordsOk() (*int32, bool)`

GetUserRecordsOk returns a tuple with the UserRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRecords

`func (o *ProductRequest) SetUserRecords(v int32)`

SetUserRecords sets UserRecords field to given value.

### HasUserRecords

`func (o *ProductRequest) HasUserRecords() bool`

HasUserRecords returns a boolean if a field has been set.

### SetUserRecordsNil

`func (o *ProductRequest) SetUserRecordsNil(b bool)`

 SetUserRecordsNil sets the value for UserRecords to be an explicit nil

### UnsetUserRecords
`func (o *ProductRequest) UnsetUserRecords()`

UnsetUserRecords ensures that no value is present for UserRecords, not even an explicit nil
### GetRevenue

`func (o *ProductRequest) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *ProductRequest) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *ProductRequest) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *ProductRequest) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *ProductRequest) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *ProductRequest) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetExternalAudience

`func (o *ProductRequest) GetExternalAudience() bool`

GetExternalAudience returns the ExternalAudience field if non-nil, zero value otherwise.

### GetExternalAudienceOk

`func (o *ProductRequest) GetExternalAudienceOk() (*bool, bool)`

GetExternalAudienceOk returns a tuple with the ExternalAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAudience

`func (o *ProductRequest) SetExternalAudience(v bool)`

SetExternalAudience sets ExternalAudience field to given value.

### HasExternalAudience

`func (o *ProductRequest) HasExternalAudience() bool`

HasExternalAudience returns a boolean if a field has been set.

### GetInternetAccessible

`func (o *ProductRequest) GetInternetAccessible() bool`

GetInternetAccessible returns the InternetAccessible field if non-nil, zero value otherwise.

### GetInternetAccessibleOk

`func (o *ProductRequest) GetInternetAccessibleOk() (*bool, bool)`

GetInternetAccessibleOk returns a tuple with the InternetAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccessible

`func (o *ProductRequest) SetInternetAccessible(v bool)`

SetInternetAccessible sets InternetAccessible field to given value.

### HasInternetAccessible

`func (o *ProductRequest) HasInternetAccessible() bool`

HasInternetAccessible returns a boolean if a field has been set.

### GetEnableProductTagInheritance

`func (o *ProductRequest) GetEnableProductTagInheritance() bool`

GetEnableProductTagInheritance returns the EnableProductTagInheritance field if non-nil, zero value otherwise.

### GetEnableProductTagInheritanceOk

`func (o *ProductRequest) GetEnableProductTagInheritanceOk() (*bool, bool)`

GetEnableProductTagInheritanceOk returns a tuple with the EnableProductTagInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductTagInheritance

`func (o *ProductRequest) SetEnableProductTagInheritance(v bool)`

SetEnableProductTagInheritance sets EnableProductTagInheritance field to given value.

### HasEnableProductTagInheritance

`func (o *ProductRequest) HasEnableProductTagInheritance() bool`

HasEnableProductTagInheritance returns a boolean if a field has been set.

### GetEnableSimpleRiskAcceptance

`func (o *ProductRequest) GetEnableSimpleRiskAcceptance() bool`

GetEnableSimpleRiskAcceptance returns the EnableSimpleRiskAcceptance field if non-nil, zero value otherwise.

### GetEnableSimpleRiskAcceptanceOk

`func (o *ProductRequest) GetEnableSimpleRiskAcceptanceOk() (*bool, bool)`

GetEnableSimpleRiskAcceptanceOk returns a tuple with the EnableSimpleRiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSimpleRiskAcceptance

`func (o *ProductRequest) SetEnableSimpleRiskAcceptance(v bool)`

SetEnableSimpleRiskAcceptance sets EnableSimpleRiskAcceptance field to given value.

### HasEnableSimpleRiskAcceptance

`func (o *ProductRequest) HasEnableSimpleRiskAcceptance() bool`

HasEnableSimpleRiskAcceptance returns a boolean if a field has been set.

### GetEnableFullRiskAcceptance

`func (o *ProductRequest) GetEnableFullRiskAcceptance() bool`

GetEnableFullRiskAcceptance returns the EnableFullRiskAcceptance field if non-nil, zero value otherwise.

### GetEnableFullRiskAcceptanceOk

`func (o *ProductRequest) GetEnableFullRiskAcceptanceOk() (*bool, bool)`

GetEnableFullRiskAcceptanceOk returns a tuple with the EnableFullRiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFullRiskAcceptance

`func (o *ProductRequest) SetEnableFullRiskAcceptance(v bool)`

SetEnableFullRiskAcceptance sets EnableFullRiskAcceptance field to given value.

### HasEnableFullRiskAcceptance

`func (o *ProductRequest) HasEnableFullRiskAcceptance() bool`

HasEnableFullRiskAcceptance returns a boolean if a field has been set.

### GetDisableSlaBreachNotifications

`func (o *ProductRequest) GetDisableSlaBreachNotifications() bool`

GetDisableSlaBreachNotifications returns the DisableSlaBreachNotifications field if non-nil, zero value otherwise.

### GetDisableSlaBreachNotificationsOk

`func (o *ProductRequest) GetDisableSlaBreachNotificationsOk() (*bool, bool)`

GetDisableSlaBreachNotificationsOk returns a tuple with the DisableSlaBreachNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSlaBreachNotifications

`func (o *ProductRequest) SetDisableSlaBreachNotifications(v bool)`

SetDisableSlaBreachNotifications sets DisableSlaBreachNotifications field to given value.

### HasDisableSlaBreachNotifications

`func (o *ProductRequest) HasDisableSlaBreachNotifications() bool`

HasDisableSlaBreachNotifications returns a boolean if a field has been set.

### GetProductManager

`func (o *ProductRequest) GetProductManager() int32`

GetProductManager returns the ProductManager field if non-nil, zero value otherwise.

### GetProductManagerOk

`func (o *ProductRequest) GetProductManagerOk() (*int32, bool)`

GetProductManagerOk returns a tuple with the ProductManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductManager

`func (o *ProductRequest) SetProductManager(v int32)`

SetProductManager sets ProductManager field to given value.

### HasProductManager

`func (o *ProductRequest) HasProductManager() bool`

HasProductManager returns a boolean if a field has been set.

### SetProductManagerNil

`func (o *ProductRequest) SetProductManagerNil(b bool)`

 SetProductManagerNil sets the value for ProductManager to be an explicit nil

### UnsetProductManager
`func (o *ProductRequest) UnsetProductManager()`

UnsetProductManager ensures that no value is present for ProductManager, not even an explicit nil
### GetTechnicalContact

`func (o *ProductRequest) GetTechnicalContact() int32`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *ProductRequest) GetTechnicalContactOk() (*int32, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *ProductRequest) SetTechnicalContact(v int32)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *ProductRequest) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### SetTechnicalContactNil

`func (o *ProductRequest) SetTechnicalContactNil(b bool)`

 SetTechnicalContactNil sets the value for TechnicalContact to be an explicit nil

### UnsetTechnicalContact
`func (o *ProductRequest) UnsetTechnicalContact()`

UnsetTechnicalContact ensures that no value is present for TechnicalContact, not even an explicit nil
### GetTeamManager

`func (o *ProductRequest) GetTeamManager() int32`

GetTeamManager returns the TeamManager field if non-nil, zero value otherwise.

### GetTeamManagerOk

`func (o *ProductRequest) GetTeamManagerOk() (*int32, bool)`

GetTeamManagerOk returns a tuple with the TeamManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManager

`func (o *ProductRequest) SetTeamManager(v int32)`

SetTeamManager sets TeamManager field to given value.

### HasTeamManager

`func (o *ProductRequest) HasTeamManager() bool`

HasTeamManager returns a boolean if a field has been set.

### SetTeamManagerNil

`func (o *ProductRequest) SetTeamManagerNil(b bool)`

 SetTeamManagerNil sets the value for TeamManager to be an explicit nil

### UnsetTeamManager
`func (o *ProductRequest) UnsetTeamManager()`

UnsetTeamManager ensures that no value is present for TeamManager, not even an explicit nil
### GetProdType

`func (o *ProductRequest) GetProdType() int32`

GetProdType returns the ProdType field if non-nil, zero value otherwise.

### GetProdTypeOk

`func (o *ProductRequest) GetProdTypeOk() (*int32, bool)`

GetProdTypeOk returns a tuple with the ProdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdType

`func (o *ProductRequest) SetProdType(v int32)`

SetProdType sets ProdType field to given value.


### GetSlaConfiguration

`func (o *ProductRequest) GetSlaConfiguration() int32`

GetSlaConfiguration returns the SlaConfiguration field if non-nil, zero value otherwise.

### GetSlaConfigurationOk

`func (o *ProductRequest) GetSlaConfigurationOk() (*int32, bool)`

GetSlaConfigurationOk returns a tuple with the SlaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaConfiguration

`func (o *ProductRequest) SetSlaConfiguration(v int32)`

SetSlaConfiguration sets SlaConfiguration field to given value.

### HasSlaConfiguration

`func (o *ProductRequest) HasSlaConfiguration() bool`

HasSlaConfiguration returns a boolean if a field has been set.

### GetRegulations

`func (o *ProductRequest) GetRegulations() []int32`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *ProductRequest) GetRegulationsOk() (*[]int32, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *ProductRequest) SetRegulations(v []int32)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *ProductRequest) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


