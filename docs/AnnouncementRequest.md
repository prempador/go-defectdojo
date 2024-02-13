# AnnouncementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | This dismissable message will be displayed on all pages for authenticated users. It can contain basic html tags, for example &lt;a href&#x3D;&#39;https://www.fred.com&#39; style&#x3D;&#39;color: #337ab7;&#39; target&#x3D;&#39;_blank&#39;&gt;https://example.com&lt;/a&gt; | [optional] 
**Style** | Pointer to **string** | The style of banner to display. (info, success, warning, danger)  * &#x60;info&#x60; - Info * &#x60;success&#x60; - Success * &#x60;warning&#x60; - Warning * &#x60;danger&#x60; - Danger | [optional] 
**Dismissable** | Pointer to **bool** | Ticking this box allows users to dismiss the current announcement | [optional] 

## Methods

### NewAnnouncementRequest

`func NewAnnouncementRequest() *AnnouncementRequest`

NewAnnouncementRequest instantiates a new AnnouncementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementRequestWithDefaults

`func NewAnnouncementRequestWithDefaults() *AnnouncementRequest`

NewAnnouncementRequestWithDefaults instantiates a new AnnouncementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AnnouncementRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnnouncementRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnnouncementRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AnnouncementRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStyle

`func (o *AnnouncementRequest) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *AnnouncementRequest) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *AnnouncementRequest) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *AnnouncementRequest) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetDismissable

`func (o *AnnouncementRequest) GetDismissable() bool`

GetDismissable returns the Dismissable field if non-nil, zero value otherwise.

### GetDismissableOk

`func (o *AnnouncementRequest) GetDismissableOk() (*bool, bool)`

GetDismissableOk returns a tuple with the Dismissable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissable

`func (o *AnnouncementRequest) SetDismissable(v bool)`

SetDismissable sets Dismissable field to given value.

### HasDismissable

`func (o *AnnouncementRequest) HasDismissable() bool`

HasDismissable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


