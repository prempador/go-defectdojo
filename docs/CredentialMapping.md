# CredentialMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**IsAuthnProvider** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**CredId** | **int32** |  | 
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**Test** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCredentialMapping

`func NewCredentialMapping(id int32, credId int32, ) *CredentialMapping`

NewCredentialMapping instantiates a new CredentialMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialMappingWithDefaults

`func NewCredentialMappingWithDefaults() *CredentialMapping`

NewCredentialMappingWithDefaults instantiates a new CredentialMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialMapping) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialMapping) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialMapping) SetId(v int32)`

SetId sets Id field to given value.


### GetIsAuthnProvider

`func (o *CredentialMapping) GetIsAuthnProvider() bool`

GetIsAuthnProvider returns the IsAuthnProvider field if non-nil, zero value otherwise.

### GetIsAuthnProviderOk

`func (o *CredentialMapping) GetIsAuthnProviderOk() (*bool, bool)`

GetIsAuthnProviderOk returns a tuple with the IsAuthnProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthnProvider

`func (o *CredentialMapping) SetIsAuthnProvider(v bool)`

SetIsAuthnProvider sets IsAuthnProvider field to given value.

### HasIsAuthnProvider

`func (o *CredentialMapping) HasIsAuthnProvider() bool`

HasIsAuthnProvider returns a boolean if a field has been set.

### GetUrl

`func (o *CredentialMapping) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CredentialMapping) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CredentialMapping) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CredentialMapping) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CredentialMapping) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CredentialMapping) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCredId

`func (o *CredentialMapping) GetCredId() int32`

GetCredId returns the CredId field if non-nil, zero value otherwise.

### GetCredIdOk

`func (o *CredentialMapping) GetCredIdOk() (*int32, bool)`

GetCredIdOk returns a tuple with the CredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredId

`func (o *CredentialMapping) SetCredId(v int32)`

SetCredId sets CredId field to given value.


### GetProduct

`func (o *CredentialMapping) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CredentialMapping) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CredentialMapping) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CredentialMapping) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *CredentialMapping) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *CredentialMapping) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetFinding

`func (o *CredentialMapping) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *CredentialMapping) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *CredentialMapping) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *CredentialMapping) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *CredentialMapping) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *CredentialMapping) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetEngagement

`func (o *CredentialMapping) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *CredentialMapping) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *CredentialMapping) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *CredentialMapping) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *CredentialMapping) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *CredentialMapping) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetTest

`func (o *CredentialMapping) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CredentialMapping) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CredentialMapping) SetTest(v int32)`

SetTest sets Test field to given value.

### HasTest

`func (o *CredentialMapping) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *CredentialMapping) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *CredentialMapping) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


