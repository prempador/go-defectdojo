# TestImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**TestImportFindingActionSet** | [**[]TestImportFindingAction**](TestImportFindingAction.md) |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**ImportSettings** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**BuildId** | Pointer to **NullableString** | Build ID that was tested, a reimport may update this field. | [optional] 
**CommitHash** | Pointer to **NullableString** | Commit hash tested, a reimport may update this field. | [optional] 
**BranchTag** | Pointer to **NullableString** | Tag or branch that was tested, a reimport may update this field. | [optional] 
**Test** | **int32** |  | [readonly] 
**FindingsAffected** | **[]int32** |  | [readonly] 
**Prefetch** | Pointer to [**PaginatedTestImportListPrefetch**](PaginatedTestImportListPrefetch.md) |  | [optional] 

## Methods

### NewTestImport

`func NewTestImport(id int32, testImportFindingActionSet []TestImportFindingAction, created time.Time, modified time.Time, test int32, findingsAffected []int32, ) *TestImport`

NewTestImport instantiates a new TestImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestImportWithDefaults

`func NewTestImportWithDefaults() *TestImport`

NewTestImportWithDefaults instantiates a new TestImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestImport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestImport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestImport) SetId(v int32)`

SetId sets Id field to given value.


### GetTestImportFindingActionSet

`func (o *TestImport) GetTestImportFindingActionSet() []TestImportFindingAction`

GetTestImportFindingActionSet returns the TestImportFindingActionSet field if non-nil, zero value otherwise.

### GetTestImportFindingActionSetOk

`func (o *TestImport) GetTestImportFindingActionSetOk() (*[]TestImportFindingAction, bool)`

GetTestImportFindingActionSetOk returns a tuple with the TestImportFindingActionSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestImportFindingActionSet

`func (o *TestImport) SetTestImportFindingActionSet(v []TestImportFindingAction)`

SetTestImportFindingActionSet sets TestImportFindingActionSet field to given value.


### GetCreated

`func (o *TestImport) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestImport) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestImport) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TestImport) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TestImport) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TestImport) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetImportSettings

`func (o *TestImport) GetImportSettings() interface{}`

GetImportSettings returns the ImportSettings field if non-nil, zero value otherwise.

### GetImportSettingsOk

`func (o *TestImport) GetImportSettingsOk() (*interface{}, bool)`

GetImportSettingsOk returns a tuple with the ImportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportSettings

`func (o *TestImport) SetImportSettings(v interface{})`

SetImportSettings sets ImportSettings field to given value.

### HasImportSettings

`func (o *TestImport) HasImportSettings() bool`

HasImportSettings returns a boolean if a field has been set.

### SetImportSettingsNil

`func (o *TestImport) SetImportSettingsNil(b bool)`

 SetImportSettingsNil sets the value for ImportSettings to be an explicit nil

### UnsetImportSettings
`func (o *TestImport) UnsetImportSettings()`

UnsetImportSettings ensures that no value is present for ImportSettings, not even an explicit nil
### GetType

`func (o *TestImport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestImport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestImport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TestImport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *TestImport) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TestImport) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TestImport) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TestImport) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TestImport) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TestImport) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetBuildId

`func (o *TestImport) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *TestImport) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *TestImport) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *TestImport) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *TestImport) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *TestImport) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetCommitHash

`func (o *TestImport) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *TestImport) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *TestImport) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *TestImport) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### SetCommitHashNil

`func (o *TestImport) SetCommitHashNil(b bool)`

 SetCommitHashNil sets the value for CommitHash to be an explicit nil

### UnsetCommitHash
`func (o *TestImport) UnsetCommitHash()`

UnsetCommitHash ensures that no value is present for CommitHash, not even an explicit nil
### GetBranchTag

`func (o *TestImport) GetBranchTag() string`

GetBranchTag returns the BranchTag field if non-nil, zero value otherwise.

### GetBranchTagOk

`func (o *TestImport) GetBranchTagOk() (*string, bool)`

GetBranchTagOk returns a tuple with the BranchTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchTag

`func (o *TestImport) SetBranchTag(v string)`

SetBranchTag sets BranchTag field to given value.

### HasBranchTag

`func (o *TestImport) HasBranchTag() bool`

HasBranchTag returns a boolean if a field has been set.

### SetBranchTagNil

`func (o *TestImport) SetBranchTagNil(b bool)`

 SetBranchTagNil sets the value for BranchTag to be an explicit nil

### UnsetBranchTag
`func (o *TestImport) UnsetBranchTag()`

UnsetBranchTag ensures that no value is present for BranchTag, not even an explicit nil
### GetTest

`func (o *TestImport) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *TestImport) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *TestImport) SetTest(v int32)`

SetTest sets Test field to given value.


### GetFindingsAffected

`func (o *TestImport) GetFindingsAffected() []int32`

GetFindingsAffected returns the FindingsAffected field if non-nil, zero value otherwise.

### GetFindingsAffectedOk

`func (o *TestImport) GetFindingsAffectedOk() (*[]int32, bool)`

GetFindingsAffectedOk returns a tuple with the FindingsAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingsAffected

`func (o *TestImport) SetFindingsAffected(v []int32)`

SetFindingsAffected sets FindingsAffected field to given value.


### GetPrefetch

`func (o *TestImport) GetPrefetch() PaginatedTestImportListPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *TestImport) GetPrefetchOk() (*PaginatedTestImportListPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *TestImport) SetPrefetch(v PaginatedTestImportListPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *TestImport) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


