# FindingEngagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Product** | Pointer to [**FindingProduct**](FindingProduct.md) |  | [optional] 
**TargetStart** | **string** |  | 
**TargetEnd** | **string** |  | 
**BranchTag** | Pointer to **NullableString** | Tag or branch of the product the engagement tested. | [optional] 
**EngagementType** | Pointer to **NullableString** | * &#x60;Interactive&#x60; - Interactive * &#x60;CI/CD&#x60; - CI/CD | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID of the product the engagement tested. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash from repo | [optional] 
**Version** | Pointer to **NullableString** | Version of the product the engagement tested. | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**Updated** | **NullableTime** |  | [readonly] 

## Methods

### NewFindingEngagement

`func NewFindingEngagement(id int32, targetStart string, targetEnd string, created NullableTime, updated NullableTime, ) *FindingEngagement`

NewFindingEngagement instantiates a new FindingEngagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingEngagementWithDefaults

`func NewFindingEngagementWithDefaults() *FindingEngagement`

NewFindingEngagementWithDefaults instantiates a new FindingEngagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FindingEngagement) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingEngagement) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingEngagement) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FindingEngagement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FindingEngagement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FindingEngagement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FindingEngagement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FindingEngagement) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FindingEngagement) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *FindingEngagement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FindingEngagement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FindingEngagement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FindingEngagement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FindingEngagement) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FindingEngagement) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProduct

`func (o *FindingEngagement) GetProduct() FindingProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *FindingEngagement) GetProductOk() (*FindingProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *FindingEngagement) SetProduct(v FindingProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *FindingEngagement) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTargetStart

`func (o *FindingEngagement) GetTargetStart() string`

GetTargetStart returns the TargetStart field if non-nil, zero value otherwise.

### GetTargetStartOk

`func (o *FindingEngagement) GetTargetStartOk() (*string, bool)`

GetTargetStartOk returns a tuple with the TargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStart

`func (o *FindingEngagement) SetTargetStart(v string)`

SetTargetStart sets TargetStart field to given value.


### GetTargetEnd

`func (o *FindingEngagement) GetTargetEnd() string`

GetTargetEnd returns the TargetEnd field if non-nil, zero value otherwise.

### GetTargetEndOk

`func (o *FindingEngagement) GetTargetEndOk() (*string, bool)`

GetTargetEndOk returns a tuple with the TargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnd

`func (o *FindingEngagement) SetTargetEnd(v string)`

SetTargetEnd sets TargetEnd field to given value.


### GetBranchTag

`func (o *FindingEngagement) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *FindingEngagement) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *FindingEngagement) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *FindingEngagement) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *FindingEngagement) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *FindingEngagement) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetEngagementType

`func (o *FindingEngagement) GetEngagementType() string`

GetEngagementType returns the EngagementType field if non-nil, zero value otherwise.

### GetEngagementTypeOk

`func (o *FindingEngagement) GetEngagementTypeOk() (*string, bool)`

GetEngagementTypeOk returns a tuple with the EngagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementType

`func (o *FindingEngagement) SetEngagementType(v string)`

SetEngagementType sets EngagementType field to given value.

### HasEngagementType

`func (o *FindingEngagement) HasEngagementType() bool`

HasEngagementType returns a boolean if a field has been set.

### SetEngagementTypeNil

`func (o *FindingEngagement) SetEngagementTypeNil(b bool)`

 SetEngagementTypeNil sets the value for EngagementType to be an explicit nil

### UnsetEngagementType
`func (o *FindingEngagement) UnsetEngagementType()`

UnsetEngagementType ensures that no value is present for EngagementType, not even an explicit nil
### GetBuildId

`func (o *FindingEngagement) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *FindingEngagement) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *FindingEngagement) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *FindingEngagement) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *FindingEngagement) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *FindingEngagement) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *FindingEngagement) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *FindingEngagement) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *FindingEngagement) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *FindingEngagement) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *FindingEngagement) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *FindingEngagement) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetVersion

`func (o *FindingEngagement) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FindingEngagement) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FindingEngagement) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FindingEngagement) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FindingEngagement) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FindingEngagement) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCreated

`func (o *FindingEngagement) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FindingEngagement) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FindingEngagement) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *FindingEngagement) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FindingEngagement) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *FindingEngagement) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FindingEngagement) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FindingEngagement) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### SetUpdatedNil

`func (o *FindingEngagement) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *FindingEngagement) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


