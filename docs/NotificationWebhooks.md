# NotificationWebhooks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | Pointer to **string** | Name of the incoming webhook | [optional] 
**Url** | Pointer to **string** | The full URL of the incoming webhook | [optional] 
**HeaderName** | Pointer to **NullableString** | Name of the header required for interacting with Webhook endpoint | [optional] 
**HeaderValue** | Pointer to **NullableString** | Content of the header required for interacting with Webhook endpoint | [optional] 
**Status** | **string** | Status of the incoming webhook | [readonly] 
**FirstError** | **NullableTime** | If endpoint is active, when error happened first time | [readonly] 
**LastError** | **NullableTime** | If endpoint is active, when error happened last time | [readonly] 
**Note** | **NullableString** | Description of the latest error | [readonly] 
**Owner** | Pointer to **NullableInt32** | Owner/receiver of notification, if empty processed as system notification | [optional] 

## Methods

### NewNotificationWebhooks

`func NewNotificationWebhooks(id int32, status string, firstError NullableTime, lastError NullableTime, note NullableString, ) *NotificationWebhooks`

NewNotificationWebhooks instantiates a new NotificationWebhooks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWebhooksWithDefaults

`func NewNotificationWebhooksWithDefaults() *NotificationWebhooks`

NewNotificationWebhooksWithDefaults instantiates a new NotificationWebhooks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationWebhooks) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationWebhooks) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationWebhooks) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *NotificationWebhooks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationWebhooks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationWebhooks) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationWebhooks) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationWebhooks) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationWebhooks) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationWebhooks) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationWebhooks) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHeaderName

`func (o *NotificationWebhooks) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *NotificationWebhooks) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *NotificationWebhooks) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *NotificationWebhooks) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### SetHeaderNameNil

`func (o *NotificationWebhooks) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *NotificationWebhooks) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
### GetHeaderValue

`func (o *NotificationWebhooks) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *NotificationWebhooks) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *NotificationWebhooks) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *NotificationWebhooks) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### SetHeaderValueNil

`func (o *NotificationWebhooks) SetHeaderValueNil(b bool)`

 SetHeaderValueNil sets the value for HeaderValue to be an explicit nil

### UnsetHeaderValue
`func (o *NotificationWebhooks) UnsetHeaderValue()`

UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
### GetStatus

`func (o *NotificationWebhooks) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationWebhooks) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationWebhooks) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFirstError

`func (o *NotificationWebhooks) GetFirstError() time.Time`

GetFirstError returns the FirstError field if non-nil, zero value otherwise.

### GetFirstErrorOk

`func (o *NotificationWebhooks) GetFirstErrorOk() (*time.Time, bool)`

GetFirstErrorOk returns a tuple with the FirstError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstError

`func (o *NotificationWebhooks) SetFirstError(v time.Time)`

SetFirstError sets FirstError field to given value.


### SetFirstErrorNil

`func (o *NotificationWebhooks) SetFirstErrorNil(b bool)`

 SetFirstErrorNil sets the value for FirstError to be an explicit nil

### UnsetFirstError
`func (o *NotificationWebhooks) UnsetFirstError()`

UnsetFirstError ensures that no value is present for FirstError, not even an explicit nil
### GetLastError

`func (o *NotificationWebhooks) GetLastError() time.Time`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *NotificationWebhooks) GetLastErrorOk() (*time.Time, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *NotificationWebhooks) SetLastError(v time.Time)`

SetLastError sets LastError field to given value.


### SetLastErrorNil

`func (o *NotificationWebhooks) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *NotificationWebhooks) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetNote

`func (o *NotificationWebhooks) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *NotificationWebhooks) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *NotificationWebhooks) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *NotificationWebhooks) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *NotificationWebhooks) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetOwner

`func (o *NotificationWebhooks) GetOwner() int32`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationWebhooks) GetOwnerOk() (*int32, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationWebhooks) SetOwner(v int32)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NotificationWebhooks) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *NotificationWebhooks) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *NotificationWebhooks) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


