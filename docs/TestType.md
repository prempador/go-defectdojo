# TestType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**StaticTool** | Pointer to **bool** |  | [optional] 
**DynamicTool** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewTestType

`func NewTestType(id int32, name string, ) *TestType`

NewTestType instantiates a new TestType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestTypeWithDefaults

`func NewTestTypeWithDefaults() *TestType`

NewTestTypeWithDefaults instantiates a new TestType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestType) SetId(v int32)`

SetId sets Id field to given value.


### GetTags

`func (o *TestType) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestType) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestType) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestType) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *TestType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestType) SetName(v string)`

SetName sets Name field to given value.


### GetStaticTool

`func (o *TestType) GetStaticTool() bool`

GetStaticTool returns the StaticTool field if non-nil, zero value otherwise.

### GetStaticToolOk

`func (o *TestType) GetStaticToolOk() (*bool, bool)`

GetStaticToolOk returns a tuple with the StaticTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticTool

`func (o *TestType) SetStaticTool(v bool)`

SetStaticTool sets StaticTool field to given value.

### HasStaticTool

`func (o *TestType) HasStaticTool() bool`

HasStaticTool returns a boolean if a field has been set.

### GetDynamicTool

`func (o *TestType) GetDynamicTool() bool`

GetDynamicTool returns the DynamicTool field if non-nil, zero value otherwise.

### GetDynamicToolOk

`func (o *TestType) GetDynamicToolOk() (*bool, bool)`

GetDynamicToolOk returns a tuple with the DynamicTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicTool

`func (o *TestType) SetDynamicTool(v bool)`

SetDynamicTool sets DynamicTool field to given value.

### HasDynamicTool

`func (o *TestType) HasDynamicTool() bool`

HasDynamicTool returns a boolean if a field has been set.

### GetActive

`func (o *TestType) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TestType) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TestType) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TestType) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


