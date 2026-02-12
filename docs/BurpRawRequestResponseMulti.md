# BurpRawRequestResponseMulti

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**BurpRequestBase64** | **string** |  | 
**BurpResponseBase64** | **string** |  | 
**Finding** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBurpRawRequestResponseMulti

`func NewBurpRawRequestResponseMulti(id int32, burpRequestBase64 string, burpResponseBase64 string, ) *BurpRawRequestResponseMulti`

NewBurpRawRequestResponseMulti instantiates a new BurpRawRequestResponseMulti object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurpRawRequestResponseMultiWithDefaults

`func NewBurpRawRequestResponseMultiWithDefaults() *BurpRawRequestResponseMulti`

NewBurpRawRequestResponseMultiWithDefaults instantiates a new BurpRawRequestResponseMulti object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BurpRawRequestResponseMulti) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BurpRawRequestResponseMulti) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BurpRawRequestResponseMulti) SetId(v int32)`

SetId sets Id field to given value.


### GetBurpRequestBase64

`func (o *BurpRawRequestResponseMulti) GetBurpRequestBase64() string`

GetBurpRequestBase64 returns the BurpRequestBase64 field if non-nil, zero value otherwise.

### GetBurpRequestBase64Ok

`func (o *BurpRawRequestResponseMulti) GetBurpRequestBase64Ok() (*string, bool)`

GetBurpRequestBase64Ok returns a tuple with the BurpRequestBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurpRequestBase64

`func (o *BurpRawRequestResponseMulti) SetBurpRequestBase64(v string)`

SetBurpRequestBase64 sets BurpRequestBase64 field to given value.


### GetBurpResponseBase64

`func (o *BurpRawRequestResponseMulti) GetBurpResponseBase64() string`

GetBurpResponseBase64 returns the BurpResponseBase64 field if non-nil, zero value otherwise.

### GetBurpResponseBase64Ok

`func (o *BurpRawRequestResponseMulti) GetBurpResponseBase64Ok() (*string, bool)`

GetBurpResponseBase64Ok returns a tuple with the BurpResponseBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurpResponseBase64

`func (o *BurpRawRequestResponseMulti) SetBurpResponseBase64(v string)`

SetBurpResponseBase64 sets BurpResponseBase64 field to given value.


### GetFinding

`func (o *BurpRawRequestResponseMulti) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *BurpRawRequestResponseMulti) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *BurpRawRequestResponseMulti) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *BurpRawRequestResponseMulti) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *BurpRawRequestResponseMulti) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *BurpRawRequestResponseMulti) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


