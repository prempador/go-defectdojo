# TestTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**StaticTool** | Pointer to **bool** |  | [optional] 
**DynamicTool** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewTestTypeRequest

`func NewTestTypeRequest(name string, ) *TestTypeRequest`

NewTestTypeRequest instantiates a new TestTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestTypeRequestWithDefaults

`func NewTestTypeRequestWithDefaults() *TestTypeRequest`

NewTestTypeRequestWithDefaults instantiates a new TestTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TestTypeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TestTypeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TestTypeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TestTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *TestTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStaticTool

`func (o *TestTypeRequest) GetStaticTool() bool`

GetStaticTool returns the StaticTool field if non-nil, zero value otherwise.

### GetStaticToolOk

`func (o *TestTypeRequest) GetStaticToolOk() (*bool, bool)`

GetStaticToolOk returns a tuple with the StaticTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticTool

`func (o *TestTypeRequest) SetStaticTool(v bool)`

SetStaticTool sets StaticTool field to given value.

### HasStaticTool

`func (o *TestTypeRequest) HasStaticTool() bool`

HasStaticTool returns a boolean if a field has been set.

### GetDynamicTool

`func (o *TestTypeRequest) GetDynamicTool() bool`

GetDynamicTool returns the DynamicTool field if non-nil, zero value otherwise.

### GetDynamicToolOk

`func (o *TestTypeRequest) GetDynamicToolOk() (*bool, bool)`

GetDynamicToolOk returns a tuple with the DynamicTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicTool

`func (o *TestTypeRequest) SetDynamicTool(v bool)`

SetDynamicTool sets DynamicTool field to given value.

### HasDynamicTool

`func (o *TestTypeRequest) HasDynamicTool() bool`

HasDynamicTool returns a boolean if a field has been set.

### GetActive

`func (o *TestTypeRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TestTypeRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TestTypeRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *TestTypeRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


