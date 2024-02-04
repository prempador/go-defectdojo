# ProductMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **int32** |  | 
**User** | **int32** |  | 
**Role** | **int32** |  | 

## Methods

### NewProductMemberRequest

`func NewProductMemberRequest(product int32, user int32, role int32, ) *ProductMemberRequest`

NewProductMemberRequest instantiates a new ProductMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductMemberRequestWithDefaults

`func NewProductMemberRequestWithDefaults() *ProductMemberRequest`

NewProductMemberRequestWithDefaults instantiates a new ProductMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ProductMemberRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ProductMemberRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ProductMemberRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.


### GetUser

`func (o *ProductMemberRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProductMemberRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProductMemberRequest) SetUser(v int32)`

SetUser sets User field to given value.


### GetRole

`func (o *ProductMemberRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProductMemberRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProductMemberRequest) SetRole(v int32)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


