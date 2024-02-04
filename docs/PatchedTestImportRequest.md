# PatchedTestImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportSettings** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID that was tested, a reimport may update this field. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash tested, a reimport may update this field. | [optional] 
**BranchTag** | Pointer to **NullableString** | Tag or branch that was tested, a reimport may update this field. | [optional] 

## Methods

### NewPatchedTestImportRequest

`func NewPatchedTestImportRequest() *PatchedTestImportRequest`

NewPatchedTestImportRequest instantiates a new PatchedTestImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTestImportRequestWithDefaults

`func NewPatchedTestImportRequestWithDefaults() *PatchedTestImportRequest`

NewPatchedTestImportRequestWithDefaults instantiates a new PatchedTestImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportSettings

`func (o *PatchedTestImportRequest) GetImportSettings() interface{}`

GetImportSettings returns the ImportSettings field if non-nil, zero value otherwise.

### GetImportSettingsOk

`func (o *PatchedTestImportRequest) GetImportSettingsOk() (*interface{}, bool)`

GetImportSettingsOk returns a tuple with the ImportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSettings

`func (o *PatchedTestImportRequest) SetImportSettings(v interface{})`

SetImportSettings sets ImportSettings field to given value.

### HasImportSettings

`func (o *PatchedTestImportRequest) HasImportSettings() bool`

HasImportSettings returns a boolean if a field has been set.

### SetImportSettingsNil

`func (o *PatchedTestImportRequest) SetImportSettingsNil(b bool)`

 SetImportSettingsNil sets the value for ImportSettings to be an explicit nil

### UnsetImportSettings
`func (o *PatchedTestImportRequest) UnsetImportSettings()`

UnsetImportSettings ensures that no value is present for ImportSettings, not even an explicit nil
### GetType

`func (o *PatchedTestImportRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedTestImportRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedTestImportRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedTestImportRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *PatchedTestImportRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PatchedTestImportRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PatchedTestImportRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PatchedTestImportRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PatchedTestImportRequest) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PatchedTestImportRequest) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildId

`func (o *PatchedTestImportRequest) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *PatchedTestImportRequest) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *PatchedTestImportRequest) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *PatchedTestImportRequest) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *PatchedTestImportRequest) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *PatchedTestImportRequest) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *PatchedTestImportRequest) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *PatchedTestImportRequest) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *PatchedTestImportRequest) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *PatchedTestImportRequest) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *PatchedTestImportRequest) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *PatchedTestImportRequest) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *PatchedTestImportRequest) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *PatchedTestImportRequest) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *PatchedTestImportRequest) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *PatchedTestImportRequest) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *PatchedTestImportRequest) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *PatchedTestImportRequest) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


