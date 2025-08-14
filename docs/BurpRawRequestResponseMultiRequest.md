# BurpRawRequestResponseMultiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurpRequestBase64** | **string** |  | 
**BurpResponseBase64** | **string** |  | 
**Finding** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBurpRawRequestResponseMultiRequest

`func NewBurpRawRequestResponseMultiRequest(burpRequestBase64 string, burpResponseBase64 string, ) *BurpRawRequestResponseMultiRequest`

NewBurpRawRequestResponseMultiRequest instantiates a new BurpRawRequestResponseMultiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBurpRawRequestResponseMultiRequestWithDefaults

`func NewBurpRawRequestResponseMultiRequestWithDefaults() *BurpRawRequestResponseMultiRequest`

NewBurpRawRequestResponseMultiRequestWithDefaults instantiates a new BurpRawRequestResponseMultiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurpRequestBase64

`func (o *BurpRawRequestResponseMultiRequest) GetBurpRequestBase64() string`

GetBurpRequestBase64 returns the BurpRequestBase64 field if non-nil, zero value otherwise.

### GetBurpRequestBase64Ok

`func (o *BurpRawRequestResponseMultiRequest) GetBurpRequestBase64Ok() (*string, bool)`

GetBurpRequestBase64Ok returns a tuple with the BurpRequestBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurpRequestBase64

`func (o *BurpRawRequestResponseMultiRequest) SetBurpRequestBase64(v string)`

SetBurpRequestBase64 sets BurpRequestBase64 field to given value.


### GetBurpResponseBase64

`func (o *BurpRawRequestResponseMultiRequest) GetBurpResponseBase64() string`

GetBurpResponseBase64 returns the BurpResponseBase64 field if non-nil, zero value otherwise.

### GetBurpResponseBase64Ok

`func (o *BurpRawRequestResponseMultiRequest) GetBurpResponseBase64Ok() (*string, bool)`

GetBurpResponseBase64Ok returns a tuple with the BurpResponseBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurpResponseBase64

`func (o *BurpRawRequestResponseMultiRequest) SetBurpResponseBase64(v string)`

SetBurpResponseBase64 sets BurpResponseBase64 field to given value.


### GetFinding

`func (o *BurpRawRequestResponseMultiRequest) GetFinding() int32`

GetFinding returns the Finding field if non-nil, zero value otherwise.

### GetFindingOk

`func (o *BurpRawRequestResponseMultiRequest) GetFindingOk() (*int32, bool)`

GetFindingOk returns a tuple with the Finding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinding

`func (o *BurpRawRequestResponseMultiRequest) SetFinding(v int32)`

SetFinding sets Finding field to given value.

### HasFinding

`func (o *BurpRawRequestResponseMultiRequest) HasFinding() bool`

HasFinding returns a boolean if a field has been set.

### SetFindingNil

`func (o *BurpRawRequestResponseMultiRequest) SetFindingNil(b bool)`

 SetFindingNil sets the value for Finding to be an explicit nil

### UnsetFinding
`func (o *BurpRawRequestResponseMultiRequest) UnsetFinding()`

UnsetFinding ensures that no value is present for Finding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


