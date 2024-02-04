# TestImportFindingAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Created** | **time.Time** |  | [readonly] 
**Modified** | **time.Time** |  | [readonly] 
**Action** | Pointer to **NullableString** | * &#x60;N&#x60; - created * &#x60;C&#x60; - closed * &#x60;R&#x60; - reactivated * &#x60;U&#x60; - left untouched | [optional] 
**TestImport** | **int32** |  | [readonly] 
**Finding** | **int32** |  | [readonly] 

## Methods

### NewTestImportFindingAction

`func NewTestImportFindingAction(id int32, created time.Time, modified time.Time, testImport int32, finding int32, ) *TestImportFindingAction`

NewTestImportFindingAction instantiates a new TestImportFindingAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestImportFindingActionWithDefaults

`func NewTestImportFindingActionWithDefaults() *TestImportFindingAction`

NewTestImportFindingActionWithDefaults instantiates a new TestImportFindingAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestImportFindingAction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestImportFindingAction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestImportFindingAction) SetId(v int32)`

SetId sets Id field to given value.


### GetCreated

`func (o *TestImportFindingAction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TestImportFindingAction) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TestImportFindingAction) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *TestImportFindingAction) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TestImportFindingAction) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TestImportFindingAction) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAction

`func (o *TestImportFindingAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TestImportFindingAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TestImportFindingAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TestImportFindingAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *TestImportFindingAction) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *TestImportFindingAction) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetTestImport

`func (o *TestImportFindingAction) GetTestImport() int32`

GetTestImport returns the TestImport field if non-nil, zero value otherwise.

### GetTestImportOk

`func (o *TestImportFindingAction) GetTestImportOk() (*int32, bool)`

GetTestImportOk returns a tuple with the TestImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestImport

`func (o *TestImportFindingAction) SetTestImport(v int32)`

SetTestImport sets TestImport field to given value.


### GetFinding

`func (o *TestImportFindingAction) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *TestImportFindingAction) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *TestImportFindingAction) SetFinding(v int32)`

SetFinding sets Finding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


