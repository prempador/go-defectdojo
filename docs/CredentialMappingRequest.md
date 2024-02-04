# CredentialMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthnProvider** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**CredId** | **int32** |  | 
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**Test** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCredentialMappingRequest

`func NewCredentialMappingRequest(credId int32, ) *CredentialMappingRequest`

NewCredentialMappingRequest instantiates a new CredentialMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialMappingRequestWithDefaults

`func NewCredentialMappingRequestWithDefaults() *CredentialMappingRequest`

NewCredentialMappingRequestWithDefaults instantiates a new CredentialMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthnProvider

`func (o *CredentialMappingRequest) GetIsAuthnProvider() bool`

GetIsAuthnProvider returns the IsAuthnProvider field if non-nil, zero value otherwise.

### GetIsAuthnProviderOk

`func (o *CredentialMappingRequest) GetIsAuthnProviderOk() (*bool, bool)`

GetIsAuthnProviderOk returns a tuple with the IsAuthnProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthnProvider

`func (o *CredentialMappingRequest) SetIsAuthnProvider(v bool)`

SetIsAuthnProvider sets IsAuthnProvider field to given value.

### HasIsAuthnProvider

`func (o *CredentialMappingRequest) HasIsAuthnProvider() bool`

HasIsAuthnProvider returns a boolean if a field has been set.

### GetUrl

`func (o *CredentialMappingRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CredentialMappingRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CredentialMappingRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CredentialMappingRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CredentialMappingRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CredentialMappingRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCredId

`func (o *CredentialMappingRequest) GetCredId() int32`

GetCredId returns the CredId field if non-nil, zero value otherwise.

### GetCredIdOk

`func (o *CredentialMappingRequest) GetCredIdOk() (*int32, bool)`

GetCredIdOk returns a tuple with the CredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredId

`func (o *CredentialMappingRequest) SetCredId(v int32)`

SetCredId sets CredId field to given value.


### GetProduct

`func (o *CredentialMappingRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CredentialMappingRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CredentialMappingRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CredentialMappingRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *CredentialMappingRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *CredentialMappingRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetFinding

`func (o *CredentialMappingRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *CredentialMappingRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *CredentialMappingRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *CredentialMappingRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *CredentialMappingRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *CredentialMappingRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetEngagement

`func (o *CredentialMappingRequest) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *CredentialMappingRequest) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *CredentialMappingRequest) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *CredentialMappingRequest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *CredentialMappingRequest) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *CredentialMappingRequest) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetTest

`func (o *CredentialMappingRequest) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *CredentialMappingRequest) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *CredentialMappingRequest) SetTest(v int32)`

SetTest sets Test field to given value.

### HasTest

`func (o *CredentialMappingRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *CredentialMappingRequest) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *CredentialMappingRequest) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


