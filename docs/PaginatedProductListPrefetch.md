# PaginatedProductListPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationGroups** | Pointer to [**map[string]DojoGroup**](DojoGroup.md) |  | [optional] [readonly] 
**Members** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**ProdType** | Pointer to [**map[string]ProductType**](ProductType.md) |  | [optional] [readonly] 
**ProductManager** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**Regulations** | Pointer to [**map[string]Regulation**](Regulation.md) |  | [optional] [readonly] 
**SlaConfiguration** | Pointer to [**map[string]SLAConfiguration**](SLAConfiguration.md) |  | [optional] [readonly] 
**TeamManager** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 
**TechnicalContact** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewPaginatedProductListPrefetch

`func NewPaginatedProductListPrefetch() *PaginatedProductListPrefetch`

NewPaginatedProductListPrefetch instantiates a new PaginatedProductListPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedProductListPrefetchWithDefaults

`func NewPaginatedProductListPrefetchWithDefaults() *PaginatedProductListPrefetch`

NewPaginatedProductListPrefetchWithDefaults instantiates a new PaginatedProductListPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationGroups

`func (o *PaginatedProductListPrefetch) GetAuthorizationGroups() map[string]DojoGroup`

GetAuthorizationGroups returns the AuthorizationGroups field if non-nil, zero value otherwise.

### GetAuthorizationGroupsOk

`func (o *PaginatedProductListPrefetch) GetAuthorizationGroupsOk() (*map[string]DojoGroup, bool)`

GetAuthorizationGroupsOk returns a tuple with the AuthorizationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroups

`func (o *PaginatedProductListPrefetch) SetAuthorizationGroups(v map[string]DojoGroup)`

SetAuthorizationGroups sets AuthorizationGroups field to given value.

### HasAuthorizationGroups

`func (o *PaginatedProductListPrefetch) HasAuthorizationGroups() bool`

HasAuthorizationGroups returns a boolean if a field has been set.

### GetMembers

`func (o *PaginatedProductListPrefetch) GetMembers() map[string]UserStub`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *PaginatedProductListPrefetch) GetMembersOk() (*map[string]UserStub, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *PaginatedProductListPrefetch) SetMembers(v map[string]UserStub)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *PaginatedProductListPrefetch) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetProdType

`func (o *PaginatedProductListPrefetch) GetProdType() map[string]ProductType`

GetProdType returns the ProdType field if non-nil, zero value otherwise.

### GetProdTypeOk

`func (o *PaginatedProductListPrefetch) GetProdTypeOk() (*map[string]ProductType, bool)`

GetProdTypeOk returns a tuple with the ProdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProdType

`func (o *PaginatedProductListPrefetch) SetProdType(v map[string]ProductType)`

SetProdType sets ProdType field to given value.

### HasProdType

`func (o *PaginatedProductListPrefetch) HasProdType() bool`

HasProdType returns a boolean if a field has been set.

### GetProductManager

`func (o *PaginatedProductListPrefetch) GetProductManager() map[string]UserStub`

GetProductManager returns the ProductManager field if non-nil, zero value otherwise.

### GetProductManagerOk

`func (o *PaginatedProductListPrefetch) GetProductManagerOk() (*map[string]UserStub, bool)`

GetProductManagerOk returns a tuple with the ProductManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductManager

`func (o *PaginatedProductListPrefetch) SetProductManager(v map[string]UserStub)`

SetProductManager sets ProductManager field to given value.

### HasProductManager

`func (o *PaginatedProductListPrefetch) HasProductManager() bool`

HasProductManager returns a boolean if a field has been set.

### GetRegulations

`func (o *PaginatedProductListPrefetch) GetRegulations() map[string]Regulation`

GetRegulations returns the Regulations field if non-nil, zero value otherwise.

### GetRegulationsOk

`func (o *PaginatedProductListPrefetch) GetRegulationsOk() (*map[string]Regulation, bool)`

GetRegulationsOk returns a tuple with the Regulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulations

`func (o *PaginatedProductListPrefetch) SetRegulations(v map[string]Regulation)`

SetRegulations sets Regulations field to given value.

### HasRegulations

`func (o *PaginatedProductListPrefetch) HasRegulations() bool`

HasRegulations returns a boolean if a field has been set.

### GetSlaConfiguration

`func (o *PaginatedProductListPrefetch) GetSlaConfiguration() map[string]SLAConfiguration`

GetSlaConfiguration returns the SlaConfiguration field if non-nil, zero value otherwise.

### GetSlaConfigurationOk

`func (o *PaginatedProductListPrefetch) GetSlaConfigurationOk() (*map[string]SLAConfiguration, bool)`

GetSlaConfigurationOk returns a tuple with the SlaConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaConfiguration

`func (o *PaginatedProductListPrefetch) SetSlaConfiguration(v map[string]SLAConfiguration)`

SetSlaConfiguration sets SlaConfiguration field to given value.

### HasSlaConfiguration

`func (o *PaginatedProductListPrefetch) HasSlaConfiguration() bool`

HasSlaConfiguration returns a boolean if a field has been set.

### GetTeamManager

`func (o *PaginatedProductListPrefetch) GetTeamManager() map[string]UserStub`

GetTeamManager returns the TeamManager field if non-nil, zero value otherwise.

### GetTeamManagerOk

`func (o *PaginatedProductListPrefetch) GetTeamManagerOk() (*map[string]UserStub, bool)`

GetTeamManagerOk returns a tuple with the TeamManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamManager

`func (o *PaginatedProductListPrefetch) SetTeamManager(v map[string]UserStub)`

SetTeamManager sets TeamManager field to given value.

### HasTeamManager

`func (o *PaginatedProductListPrefetch) HasTeamManager() bool`

HasTeamManager returns a boolean if a field has been set.

### GetTechnicalContact

`func (o *PaginatedProductListPrefetch) GetTechnicalContact() map[string]UserStub`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *PaginatedProductListPrefetch) GetTechnicalContactOk() (*map[string]UserStub, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *PaginatedProductListPrefetch) SetTechnicalContact(v map[string]UserStub)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *PaginatedProductListPrefetch) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


