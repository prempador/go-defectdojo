# FindingTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Title** | Pointer to **NullableString** |  | [optional] 
**TestType** | Pointer to [**FindingTestType**](FindingTestType.md) |  | [optional] 
**Engagement** | Pointer to [**FindingEngagement**](FindingEngagement.md) |  | [optional] 
**Environment** | Pointer to [**FindingEnvironment**](FindingEnvironment.md) |  | [optional] 
**BranchTag** | Pointer to **NullableString** | Tag or branch that was tested, a reimport may update this field. | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID that was tested, a reimport may update this field. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash tested, a reimport may update this field. | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFindingTest

`func NewFindingTest(id int32, ) *FindingTest`

NewFindingTest instantiates a new FindingTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingTestWithDefaults

`func NewFindingTestWithDefaults() *FindingTest`

NewFindingTestWithDefaults instantiates a new FindingTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FindingTest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingTest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingTest) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *FindingTest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FindingTest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FindingTest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FindingTest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FindingTest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FindingTest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTestType

`func (o *FindingTest) GetTestType() FindingTestType`

GetTestType returns the TestType field if non-nil, zero value otherwise.

### GetTestTypeOk

`func (o *FindingTest) GetTestTypeOk() (*FindingTestType, bool)`

GetTestTypeOk returns a tuple with the TestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestType

`func (o *FindingTest) SetTestType(v FindingTestType)`

SetTestType sets TestType field to given value.

### HasTestType

`func (o *FindingTest) HasTestType() bool`

HasTestType returns a boolean if a field has been set.

### GetEngagement

`func (o *FindingTest) GetEngagement() FindingEngagement`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *FindingTest) GetEngagementOk() (*FindingEngagement, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *FindingTest) SetEngagement(v FindingEngagement)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *FindingTest) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetEnvironment

`func (o *FindingTest) GetEnvironment() FindingEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FindingTest) GetEnvironmentOk() (*FindingEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FindingTest) SetEnvironment(v FindingEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FindingTest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetBranchTag

`func (o *FindingTest) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *FindingTest) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *FindingTest) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *FindingTest) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *FindingTest) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *FindingTest) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetBuildId

`func (o *FindingTest) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *FindingTest) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *FindingTest) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *FindingTest) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *FindingTest) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *FindingTest) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *FindingTest) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *FindingTest) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *FindingTest) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *FindingTest) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *FindingTest) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *FindingTest) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetVersion

`func (o *FindingTest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FindingTest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FindingTest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FindingTest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *FindingTest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *FindingTest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


