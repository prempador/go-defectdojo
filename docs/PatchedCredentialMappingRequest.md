# PatchedCredentialMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthnProvider** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**CredId** | Pointer to **int32** |  | [optional] 
**Product** | Pointer to **NullableInt32** |  | [optional] 
**Finding** | Pointer to **NullableInt32** |  | [optional] 
**Engagement** | Pointer to **NullableInt32** |  | [optional] 
**Test** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPatchedCredentialMappingRequest

`func NewPatchedCredentialMappingRequest() *PatchedCredentialMappingRequest`

NewPatchedCredentialMappingRequest instantiates a new PatchedCredentialMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCredentialMappingRequestWithDefaults

`func NewPatchedCredentialMappingRequestWithDefaults() *PatchedCredentialMappingRequest`

NewPatchedCredentialMappingRequestWithDefaults instantiates a new PatchedCredentialMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthnProvider

`func (o *PatchedCredentialMappingRequest) GetIsAuthnProvider() bool`

GetIsAuthnProvider returns the IsAuthnProvider field if non-nil, zero value otherwise.

### GetIsAuthnProviderOk

`func (o *PatchedCredentialMappingRequest) GetIsAuthnProviderOk() (*bool, bool)`

GetIsAuthnProviderOk returns a tuple with the IsAuthnProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthnProvider

`func (o *PatchedCredentialMappingRequest) SetIsAuthnProvider(v bool)`

SetIsAuthnProvider sets IsAuthnProvider field to given value.

### HasIsAuthnProvider

`func (o *PatchedCredentialMappingRequest) HasIsAuthnProvider() bool`

HasIsAuthnProvider returns a boolean if a field has been set.

### GetUrl

`func (o *PatchedCredentialMappingRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchedCredentialMappingRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchedCredentialMappingRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchedCredentialMappingRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PatchedCredentialMappingRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PatchedCredentialMappingRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCredId

`func (o *PatchedCredentialMappingRequest) GetCredId() int32`

GetCredId returns the CredId field if non-nil, zero value otherwise.

### GetCredIdOk

`func (o *PatchedCredentialMappingRequest) GetCredIdOk() (*int32, bool)`

GetCredIdOk returns a tuple with the CredId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredId

`func (o *PatchedCredentialMappingRequest) SetCredId(v int32)`

SetCredId sets CredId field to given value.

### HasCredId

`func (o *PatchedCredentialMappingRequest) HasCredId() bool`

HasCredId returns a boolean if a field has been set.

### GetProduct

`func (o *PatchedCredentialMappingRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedCredentialMappingRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedCredentialMappingRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedCredentialMappingRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *PatchedCredentialMappingRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *PatchedCredentialMappingRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetFinding

`func (o *PatchedCredentialMappingRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *PatchedCredentialMappingRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *PatchedCredentialMappingRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *PatchedCredentialMappingRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *PatchedCredentialMappingRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *PatchedCredentialMappingRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil
### GetEngagement

`func (o *PatchedCredentialMappingRequest) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *PatchedCredentialMappingRequest) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *PatchedCredentialMappingRequest) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *PatchedCredentialMappingRequest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### SetEngagementNil

`func (o *PatchedCredentialMappingRequest) SetEngagementNil(b bool)`

 SetEngagementNil sets the value for Engagement to be an explicit nil

### UnsetEngagement
`func (o *PatchedCredentialMappingRequest) UnsetEngagement()`

UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
### GetTest

`func (o *PatchedCredentialMappingRequest) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *PatchedCredentialMappingRequest) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *PatchedCredentialMappingRequest) SetTest(v int32)`

SetTest sets Test field to given value.

### HasTest

`func (o *PatchedCredentialMappingRequest) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *PatchedCredentialMappingRequest) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *PatchedCredentialMappingRequest) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


