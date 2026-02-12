# NotificationWebhooksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the incoming webhook | [optional] 
**Url** | Pointer to **string** | The full URL of the incoming webhook | [optional] 
**HeaderName** | Pointer to **NullableString** | Name of the header required for interacting with Webhook endpoint | [optional] 
**HeaderValue** | Pointer to **NullableString** | Content of the header required for interacting with Webhook endpoint | [optional] 
**Owner** | Pointer to **NullableInt32** | Owner/receiver of notification, if empty processed as system notification | [optional] 

## Methods

### NewNotificationWebhooksRequest

`func NewNotificationWebhooksRequest() *NotificationWebhooksRequest`

NewNotificationWebhooksRequest instantiates a new NotificationWebhooksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWebhooksRequestWithDefaults

`func NewNotificationWebhooksRequestWithDefaults() *NotificationWebhooksRequest`

NewNotificationWebhooksRequestWithDefaults instantiates a new NotificationWebhooksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationWebhooksRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationWebhooksRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationWebhooksRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationWebhooksRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationWebhooksRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationWebhooksRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationWebhooksRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationWebhooksRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaderName

`func (o *NotificationWebhooksRequest) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *NotificationWebhooksRequest) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *NotificationWebhooksRequest) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *NotificationWebhooksRequest) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### SetHeaderNameNil

`func (o *NotificationWebhooksRequest) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *NotificationWebhooksRequest) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
### GetHeaderValue

`func (o *NotificationWebhooksRequest) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *NotificationWebhooksRequest) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *NotificationWebhooksRequest) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *NotificationWebhooksRequest) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### SetHeaderValueNil

`func (o *NotificationWebhooksRequest) SetHeaderValueNil(b bool)`

 SetHeaderValueNil sets the value for HeaderValue to be an explicit nil

### UnsetHeaderValue
`func (o *NotificationWebhooksRequest) UnsetHeaderValue()`

UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
### GetOwner

`func (o *NotificationWebhooksRequest) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationWebhooksRequest) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationWebhooksRequest) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NotificationWebhooksRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *NotificationWebhooksRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *NotificationWebhooksRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


