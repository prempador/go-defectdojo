# NotificationsPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 
**User** | Pointer to [**map[string]UserStub**](UserStub.md) |  | [optional] [readonly] 

## Methods

### NewNotificationsPrefetch

`func NewNotificationsPrefetch() *NotificationsPrefetch`

NewNotificationsPrefetch instantiates a new NotificationsPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPrefetchWithDefaults

`func NewNotificationsPrefetchWithDefaults() *NotificationsPrefetch`

NewNotificationsPrefetchWithDefaults instantiates a new NotificationsPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *NotificationsPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *NotificationsPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *NotificationsPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *NotificationsPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetUser

`func (o *NotificationsPrefetch) GetUser() map[string]UserStub`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NotificationsPrefetch) GetUserOk() (*map[string]UserStub, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NotificationsPrefetch) SetUser(v map[string]UserStub)`

SetUser sets User field to given value.

### HasUser

`func (o *NotificationsPrefetch) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


