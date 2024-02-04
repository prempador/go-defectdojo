# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**FindingsCount** | **int32** |  | [readonly] 
**FindingsList** | **[]int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ProductMeta** | [**[]ProductMeta**](ProductMeta.md) |  | [readonly] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Created** | **NullableTime** |  | [readonly] 
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
**Members** | **[]int32** |  | [readonly] 
**AuthorizationGroups** | **[]int32** |  | [readonly] 
**Regulations** | Pointer to **[]int32** |  | [optional] 
**Prefetch** | Pointer to [**PaginatedProductListPrefetch**](PaginatedProductListPrefetch.md) |  | [optional] 

## Methods

### NewProduct

`func NewProduct(id int32, findingsCount int32, findingsList []int32, productMeta []ProductMeta, name string, description string, created NullableTime, prodType int32, members []int32, authorizationGroups []int32, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v int32)`

SetId sets Id field to given value.


### GetFindingsCount

`func (o *Product) GetFindingsCount() int32`

GetFindingsCount returns the FindingsCount field if non-nil, zero value otherwise.

### GetFindingsCountOk

`func (o *Product) GetFindingsCountOk() (*int32, bool)`

GetFindingsCountOk returns a tuple with the FindingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingsCount

`func (o *Product) SetFindingsCount(v int32)`

SetFindingsCount sets FindingsCount field to given value.


### GetFindingsList

`func (o *Product) GetFindingsList() []int32`

GetFindingsList returns the FindingsList field if non-nil, zero value otherwise.

### GetFindingsListOk

`func (o *Product) GetFindingsListOk() (*[]int32, bool)`

GetFindingsListOk returns a tuple with the FindingsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingsList

`func (o *Product) SetFindingsList(v []int32)`

SetFindingsList sets FindingsList field to given value.


### GetTags

`func (o *Product) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Product) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Product) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Product) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProductMeta

`func (o *Product) GetProductMeta() []ProductMeta`

GetProductMeta returns the ProductMeta field if non-nil, zero value otherwise.

### GetProductMetaOk

`func (o *Product) GetProductMetaOk() (*[]ProductMeta, bool)`

GetProductMetaOk returns a tuple with the ProductMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMeta

`func (o *Product) SetProductMeta(v []ProductMeta)`

SetProductMeta sets ProductMeta field to given value.


### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *Product) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Product) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Product) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Product) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Product) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetProdNumericGrade

`func (o *Product) GetProdNumericGrade() int32`

GetProdNumericGrade returns the ProdNumericGrade field if non-nil, zero value otherwise.

### GetProdNumericGradeOk

`func (o *Product) GetProdNumericGradeOk() (*int32, bool)`

GetProdNumericGradeOk returns a tuple with the ProdNumericGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdNumericGrade

`func (o *Product) SetProdNumericGrade(v int32)`

SetProdNumericGrade sets ProdNumericGrade field to given value.

### HasProdNumericGrade

`func (o *Product) HasProdNumericGrade() bool`

HasProdNumericGrade returns a boolean if a field has been set.

### SetProdNumericGradeNil

`func (o *Product) SetProdNumericGradeNil(b bool)`

 SetProdNumericGradeNil sets the value for ProdNumericGrade to be an explicit nil

### UnsetProdNumericGrade
`func (o *Product) UnsetProdNumericGrade()`

UnsetProdNumericGrade ensures that no value is present for ProdNumericGrade, not even an explicit nil
### GetBusinessCriticality

`func (o *Product) GetBusinessCriticality() string`

GetBusinessCriticality returns the BusinessCriticality field if non-nil, zero value otherwise.

### GetBusinessCriticalityOk

`func (o *Product) GetBusinessCriticalityOk() (*string, bool)`

GetBusinessCriticalityOk returns a tuple with the BusinessCriticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCriticality

`func (o *Product) SetBusinessCriticality(v string)`

SetBusinessCriticality sets BusinessCriticality field to given value.

### HasBusinessCriticality

`func (o *Product) HasBusinessCriticality() bool`

HasBusinessCriticality returns a boolean if a field has been set.

### SetBusinessCriticalityNil

`func (o *Product) SetBusinessCriticalityNil(b bool)`

 SetBusinessCriticalityNil sets the value for BusinessCriticality to be an explicit nil

### UnsetBusinessCriticality
`func (o *Product) UnsetBusinessCriticality()`

UnsetBusinessCriticality ensures that no value is present for BusinessCriticality, not even an explicit nil
### GetPlatform

`func (o *Product) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Product) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Product) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Product) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Product) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Product) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLifecycle

`func (o *Product) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *Product) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *Product) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *Product) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### SetLifecycleNil

`func (o *Product) SetLifecycleNil(b bool)`

 SetLifecycleNil sets the value for Lifecycle to be an explicit nil

### UnsetLifecycle
`func (o *Product) UnsetLifecycle()`

UnsetLifecycle ensures that no value is present for Lifecycle, not even an explicit nil
### GetOrigin

`func (o *Product) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Product) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Product) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Product) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *Product) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *Product) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetUserRecords

`func (o *Product) GetUserRecords() int32`

GetUserRecords returns the UserRecords field if non-nil, zero value otherwise.

### GetUserRecordsOk

`func (o *Product) GetUserRecordsOk() (*int32, bool)`

GetUserRecordsOk returns a tuple with the UserRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRecords

`func (o *Product) SetUserRecords(v int32)`

SetUserRecords sets UserRecords field to given value.

### HasUserRecords

`func (o *Product) HasUserRecords() bool`

HasUserRecords returns a boolean if a field has been set.

### SetUserRecordsNil

`func (o *Product) SetUserRecordsNil(b bool)`

 SetUserRecordsNil sets the value for UserRecords to be an explicit nil

### UnsetUserRecords
`func (o *Product) UnsetUserRecords()`

UnsetUserRecords ensures that no value is present for UserRecords, not even an explicit nil
### GetRevenue

`func (o *Product) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *Product) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *Product) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *Product) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *Product) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *Product) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetExternalAudience

`func (o *Product) GetExternalAudience() bool`

GetExternalAudience returns the ExternalAudience field if non-nil, zero value otherwise.

### GetExternalAudienceOk

`func (o *Product) GetExternalAudienceOk() (*bool, bool)`

GetExternalAudienceOk returns a tuple with the ExternalAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAudience

`func (o *Product) SetExternalAudience(v bool)`

SetExternalAudience sets ExternalAudience field to given value.

### HasExternalAudience

`func (o *Product) HasExternalAudience() bool`

HasExternalAudience returns a boolean if a field has been set.

### GetInternetAccessible

`func (o *Product) GetInternetAccessible() bool`

GetInternetAccessible returns the InternetAccessible field if non-nil, zero value otherwise.

### GetInternetAccessibleOk

`func (o *Product) GetInternetAccessibleOk() (*bool, bool)`

GetInternetAccessibleOk returns a tuple with the InternetAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetAccessible

`func (o *Product) SetInternetAccessible(v bool)`

SetInternetAccessible sets InternetAccessible field to given value.

### HasInternetAccessible

`func (o *Product) HasInternetAccessible() bool`

HasInternetAccessible returns a boolean if a field has been set.

### GetEnableProductTagInheritance

`func (o *Product) GetEnableProductTagInheritance() bool`

GetEnableProductTagInheritance returns the EnableProductTagInheritance field if non-nil, zero value otherwise.

### GetEnableProductTagInheritanceOk

`func (o *Product) GetEnableProductTagInheritanceOk() (*bool, bool)`

GetEnableProductTagInheritanceOk returns a tuple with the EnableProductTagInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProductTagInheritance

`func (o *Product) SetEnableProductTagInheritance(v bool)`

SetEnableProductTagInheritance sets EnableProductTagInheritance field to given value.

### HasEnableProductTagInheritance

`func (o *Product) HasEnableProductTagInheritance() bool`

HasEnableProductTagInheritance returns a boolean if a field has been set.

### GetEnableSimpleRiskAcceptance

`func (o *Product) GetEnableSimpleRiskAcceptance() bool`

GetEnableSimpleRiskAcceptance returns the EnableSimpleRiskAcceptance field if non-nil, zero value otherwise.

### GetEnableSimpleRiskAcceptanceOk

`func (o *Product) GetEnableSimpleRiskAcceptanceOk() (*bool, bool)`

GetEnableSimpleRiskAcceptanceOk returns a tuple with the EnableSimpleRiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSimpleRiskAcceptance

`func (o *Product) SetEnableSimpleRiskAcceptance(v bool)`

SetEnableSimpleRiskAcceptance sets EnableSimpleRiskAcceptance field to given value.

### HasEnableSimpleRiskAcceptance

`func (o *Product) HasEnableSimpleRiskAcceptance() bool`

HasEnableSimpleRiskAcceptance returns a boolean if a field has been set.

### GetEnableFullRiskAcceptance

`func (o *Product) GetEnableFullRiskAcceptance() bool`

GetEnableFullRiskAcceptance returns the EnableFullRiskAcceptance field if non-nil, zero value otherwise.

### GetEnableFullRiskAcceptanceOk

`func (o *Product) GetEnableFullRiskAcceptanceOk() (*bool, bool)`

GetEnableFullRiskAcceptanceOk returns a tuple with the EnableFullRiskAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFullRiskAcceptance

`func (o *Product) SetEnableFullRiskAcceptance(v bool)`

SetEnableFullRiskAcceptance sets EnableFullRiskAcceptance field to given value.

### HasEnableFullRiskAcceptance

`func (o *Product) HasEnableFullRiskAcceptance() bool`

HasEnableFullRiskAcceptance returns a boolean if a field has been set.

### GetDisableSlaBreachNotifications

`func (o *Product) GetDisableSlaBreachNotifications() bool`

GetDisableSlaBreachNotifications returns the DisableSlaBreachNotifications field if non-nil, zero value otherwise.

### GetDisableSlaBreachNotificationsOk

`func (o *Product) GetDisableSlaBreachNotificationsOk() (*bool, bool)`

GetDisableSlaBreachNotificationsOk returns a tuple with the DisableSlaBreachNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSlaBreachNotifications

`func (o *Product) SetDisableSlaBreachNotifications(v bool)`

SetDisableSlaBreachNotifications sets DisableSlaBreachNotifications field to given value.

### HasDisableSlaBreachNotifications

`func (o *Product) HasDisableSlaBreachNotifications() bool`

HasDisableSlaBreachNotifications returns a boolean if a field has been set.

### GetProductManager

`func (o *Product) GetProductManager() int32`

GetProductManager returns the ProductManager field if non-nil, zero value otherwise.

### GetProductManagerOk

`func (o *Product) GetProductManagerOk() (*int32, bool)`

GetProductManagerOk returns a tuple with the ProductManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductManager

`func (o *Product) SetProductManager(v int32)`

SetProductManager sets ProductManager field to given value.

### HasProductManager

`func (o *Product) HasProductManager() bool`

HasProductManager returns a boolean if a field has been set.

### SetProductManagerNil

`func (o *Product) SetProductManagerNil(b bool)`

 SetProductManagerNil sets the value for ProductManager to be an explicit nil

### UnsetProductManager
`func (o *Product) UnsetProductManager()`

UnsetProductManager ensures that no value is present for ProductManager, not even an explicit nil
### GetTechnicalContact

`func (o *Product) GetTechnicalContact() int32`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *Product) GetTechnicalContactOk() (*int32, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *Product) SetTechnicalContact(v int32)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *Product) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### SetTechnicalContactNil

`func (o *Product) SetTechnicalContactNil(b bool)`

 SetTechnicalContactNil sets the value for TechnicalContact to be an explicit nil

### UnsetTechnicalContact
`func (o *Product) UnsetTechnicalContact()`

UnsetTechnicalContact ensures that no value is present for TechnicalContact, not even an explicit nil
### GetTeamManager

`func (o *Product) GetTeamManager() int32`

GetTeamManager returns the TeamManager field if non-nil, zero value otherwise.

### GetTeamManagerOk

`func (o *Product) GetTeamManagerOk() (*int32, bool)`

GetTeamManagerOk returns a tuple with the TeamManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManager

`func (o *Product) SetTeamManager(v int32)`

SetTeamManager sets TeamManager field to given value.

### HasTeamManager

`func (o *Product) HasTeamManager() bool`

HasTeamManager returns a boolean if a field has been set.

### SetTeamManagerNil

`func (o *Product) SetTeamManagerNil(b bool)`

 SetTeamManagerNil sets the value for TeamManager to be an explicit nil

### UnsetTeamManager
`func (o *Product) UnsetTeamManager()`

UnsetTeamManager ensures that no value is present for TeamManager, not even an explicit nil
### GetProdType

`func (o *Product) GetProdType() int32`

GetProdType returns the ProdType field if non-nil, zero value otherwise.

### GetProdTypeOk

`func (o *Product) GetProdTypeOk() (*int32, bool)`

GetProdTypeOk returns a tuple with the ProdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdType

`func (o *Product) SetProdType(v int32)`

SetProdType sets ProdType field to given value.


### GetSlaConfiguration

`func (o *Product) GetSlaConfiguration() int32`

GetSlaConfiguration returns the SlaConfiguration field if non-nil, zero value otherwise.

### GetSlaConfigurationOk

`func (o *Product) GetSlaConfigurationOk() (*int32, bool)`

GetSlaConfigurationOk returns a tuple with the SlaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaConfiguration

`func (o *Product) SetSlaConfiguration(v int32)`

SetSlaConfiguration sets SlaConfiguration field to given value.

### HasSlaConfiguration

`func (o *Product) HasSlaConfiguration() bool`

HasSlaConfiguration returns a boolean if a field has been set.

### GetMembers

`func (o *Product) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Product) GetMembersOk() (*[]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Product) SetMembers(v []int32)`

SetMembers sets Members field to given value.


### GetAuthorizationGroups

`func (o *Product) GetAuthorizationGroups() []int32`

GetAuthorizationGroups returns the AuthorizationGroups field if non-nil, zero value otherwise.

### GetAuthorizationGroupsOk

`func (o *Product) GetAuthorizationGroupsOk() (*[]int32, bool)`

GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroups

`func (o *Product) SetAuthorizationGroups(v []int32)`

SetAuthorizationGroups sets AuthorizationGroups field to given value.


### GetRegulations

`func (o *Product) GetRegulations() []int32`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *Product) GetRegulationsOk() (*[]int32, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *Product) SetRegulations(v []int32)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *Product) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.

### GetPrefetch

`func (o *Product) GetPrefetch() PaginatedProductListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *Product) GetPrefetchOk() (*PaginatedProductListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *Product) SetPrefetch(v PaginatedProductListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *Product) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


