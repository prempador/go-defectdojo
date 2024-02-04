# NotificationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **NullableInt32** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**ProductTypeAdded** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**ProductAdded** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**EngagementAdded** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**TestAdded** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**ScanAdded** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**JiraUpdate** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**UpcomingEngagement** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**StaleEngagement** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**AutoCloseEngagement** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**CloseEngagement** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**UserMentioned** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**CodeReview** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**ReviewRequested** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**Other** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**SlaBreach** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**RiskAcceptanceExpiration** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**Template** | Pointer to **bool** |  | [optional] [default to false]
**ScanAddedEmpty** | Pointer to **string** | Triggered whenever an (re-)import has been done (even if that created/updated/closed no findings).  * &#x60;slack&#x60; - slack * &#x60;msteams&#x60; - msteams * &#x60;mail&#x60; - mail * &#x60;alert&#x60; - alert | [optional] 

## Methods

### NewNotificationsRequest

`func NewNotificationsRequest() *NotificationsRequest`

NewNotificationsRequest instantiates a new NotificationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsRequestWithDefaults

`func NewNotificationsRequestWithDefaults() *NotificationsRequest`

NewNotificationsRequestWithDefaults instantiates a new NotificationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *NotificationsRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *NotificationsRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *NotificationsRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *NotificationsRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *NotificationsRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *NotificationsRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetUser

`func (o *NotificationsRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NotificationsRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NotificationsRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *NotificationsRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *NotificationsRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *NotificationsRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetProductTypeAdded

`func (o *NotificationsRequest) GetProductTypeAdded() []string`

GetProductTypeAdded returns the ProductTypeAdded field if non-nil, zero value otherwise.

### GetProductTypeAddedOk

`func (o *NotificationsRequest) GetProductTypeAddedOk() (*[]string, bool)`

GetProductTypeAddedOk returns a tuple with the ProductTypeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeAdded

`func (o *NotificationsRequest) SetProductTypeAdded(v []string)`

SetProductTypeAdded sets ProductTypeAdded field to given value.

### HasProductTypeAdded

`func (o *NotificationsRequest) HasProductTypeAdded() bool`

HasProductTypeAdded returns a boolean if a field has been set.

### GetProductAdded

`func (o *NotificationsRequest) GetProductAdded() []string`

GetProductAdded returns the ProductAdded field if non-nil, zero value otherwise.

### GetProductAddedOk

`func (o *NotificationsRequest) GetProductAddedOk() (*[]string, bool)`

GetProductAddedOk returns a tuple with the ProductAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAdded

`func (o *NotificationsRequest) SetProductAdded(v []string)`

SetProductAdded sets ProductAdded field to given value.

### HasProductAdded

`func (o *NotificationsRequest) HasProductAdded() bool`

HasProductAdded returns a boolean if a field has been set.

### GetEngagementAdded

`func (o *NotificationsRequest) GetEngagementAdded() []string`

GetEngagementAdded returns the EngagementAdded field if non-nil, zero value otherwise.

### GetEngagementAddedOk

`func (o *NotificationsRequest) GetEngagementAddedOk() (*[]string, bool)`

GetEngagementAddedOk returns a tuple with the EngagementAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementAdded

`func (o *NotificationsRequest) SetEngagementAdded(v []string)`

SetEngagementAdded sets EngagementAdded field to given value.

### HasEngagementAdded

`func (o *NotificationsRequest) HasEngagementAdded() bool`

HasEngagementAdded returns a boolean if a field has been set.

### GetTestAdded

`func (o *NotificationsRequest) GetTestAdded() []string`

GetTestAdded returns the TestAdded field if non-nil, zero value otherwise.

### GetTestAddedOk

`func (o *NotificationsRequest) GetTestAddedOk() (*[]string, bool)`

GetTestAddedOk returns a tuple with the TestAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAdded

`func (o *NotificationsRequest) SetTestAdded(v []string)`

SetTestAdded sets TestAdded field to given value.

### HasTestAdded

`func (o *NotificationsRequest) HasTestAdded() bool`

HasTestAdded returns a boolean if a field has been set.

### GetScanAdded

`func (o *NotificationsRequest) GetScanAdded() []string`

GetScanAdded returns the ScanAdded field if non-nil, zero value otherwise.

### GetScanAddedOk

`func (o *NotificationsRequest) GetScanAddedOk() (*[]string, bool)`

GetScanAddedOk returns a tuple with the ScanAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAdded

`func (o *NotificationsRequest) SetScanAdded(v []string)`

SetScanAdded sets ScanAdded field to given value.

### HasScanAdded

`func (o *NotificationsRequest) HasScanAdded() bool`

HasScanAdded returns a boolean if a field has been set.

### GetJiraUpdate

`func (o *NotificationsRequest) GetJiraUpdate() []string`

GetJiraUpdate returns the JiraUpdate field if non-nil, zero value otherwise.

### GetJiraUpdateOk

`func (o *NotificationsRequest) GetJiraUpdateOk() (*[]string, bool)`

GetJiraUpdateOk returns a tuple with the JiraUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraUpdate

`func (o *NotificationsRequest) SetJiraUpdate(v []string)`

SetJiraUpdate sets JiraUpdate field to given value.

### HasJiraUpdate

`func (o *NotificationsRequest) HasJiraUpdate() bool`

HasJiraUpdate returns a boolean if a field has been set.

### GetUpcomingEngagement

`func (o *NotificationsRequest) GetUpcomingEngagement() []string`

GetUpcomingEngagement returns the UpcomingEngagement field if non-nil, zero value otherwise.

### GetUpcomingEngagementOk

`func (o *NotificationsRequest) GetUpcomingEngagementOk() (*[]string, bool)`

GetUpcomingEngagementOk returns a tuple with the UpcomingEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingEngagement

`func (o *NotificationsRequest) SetUpcomingEngagement(v []string)`

SetUpcomingEngagement sets UpcomingEngagement field to given value.

### HasUpcomingEngagement

`func (o *NotificationsRequest) HasUpcomingEngagement() bool`

HasUpcomingEngagement returns a boolean if a field has been set.

### GetStaleEngagement

`func (o *NotificationsRequest) GetStaleEngagement() []string`

GetStaleEngagement returns the StaleEngagement field if non-nil, zero value otherwise.

### GetStaleEngagementOk

`func (o *NotificationsRequest) GetStaleEngagementOk() (*[]string, bool)`

GetStaleEngagementOk returns a tuple with the StaleEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleEngagement

`func (o *NotificationsRequest) SetStaleEngagement(v []string)`

SetStaleEngagement sets StaleEngagement field to given value.

### HasStaleEngagement

`func (o *NotificationsRequest) HasStaleEngagement() bool`

HasStaleEngagement returns a boolean if a field has been set.

### GetAutoCloseEngagement

`func (o *NotificationsRequest) GetAutoCloseEngagement() []string`

GetAutoCloseEngagement returns the AutoCloseEngagement field if non-nil, zero value otherwise.

### GetAutoCloseEngagementOk

`func (o *NotificationsRequest) GetAutoCloseEngagementOk() (*[]string, bool)`

GetAutoCloseEngagementOk returns a tuple with the AutoCloseEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCloseEngagement

`func (o *NotificationsRequest) SetAutoCloseEngagement(v []string)`

SetAutoCloseEngagement sets AutoCloseEngagement field to given value.

### HasAutoCloseEngagement

`func (o *NotificationsRequest) HasAutoCloseEngagement() bool`

HasAutoCloseEngagement returns a boolean if a field has been set.

### GetCloseEngagement

`func (o *NotificationsRequest) GetCloseEngagement() []string`

GetCloseEngagement returns the CloseEngagement field if non-nil, zero value otherwise.

### GetCloseEngagementOk

`func (o *NotificationsRequest) GetCloseEngagementOk() (*[]string, bool)`

GetCloseEngagementOk returns a tuple with the CloseEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseEngagement

`func (o *NotificationsRequest) SetCloseEngagement(v []string)`

SetCloseEngagement sets CloseEngagement field to given value.

### HasCloseEngagement

`func (o *NotificationsRequest) HasCloseEngagement() bool`

HasCloseEngagement returns a boolean if a field has been set.

### GetUserMentioned

`func (o *NotificationsRequest) GetUserMentioned() []string`

GetUserMentioned returns the UserMentioned field if non-nil, zero value otherwise.

### GetUserMentionedOk

`func (o *NotificationsRequest) GetUserMentionedOk() (*[]string, bool)`

GetUserMentionedOk returns a tuple with the UserMentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMentioned

`func (o *NotificationsRequest) SetUserMentioned(v []string)`

SetUserMentioned sets UserMentioned field to given value.

### HasUserMentioned

`func (o *NotificationsRequest) HasUserMentioned() bool`

HasUserMentioned returns a boolean if a field has been set.

### GetCodeReview

`func (o *NotificationsRequest) GetCodeReview() []string`

GetCodeReview returns the CodeReview field if non-nil, zero value otherwise.

### GetCodeReviewOk

`func (o *NotificationsRequest) GetCodeReviewOk() (*[]string, bool)`

GetCodeReviewOk returns a tuple with the CodeReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeReview

`func (o *NotificationsRequest) SetCodeReview(v []string)`

SetCodeReview sets CodeReview field to given value.

### HasCodeReview

`func (o *NotificationsRequest) HasCodeReview() bool`

HasCodeReview returns a boolean if a field has been set.

### GetReviewRequested

`func (o *NotificationsRequest) GetReviewRequested() []string`

GetReviewRequested returns the ReviewRequested field if non-nil, zero value otherwise.

### GetReviewRequestedOk

`func (o *NotificationsRequest) GetReviewRequestedOk() (*[]string, bool)`

GetReviewRequestedOk returns a tuple with the ReviewRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequested

`func (o *NotificationsRequest) SetReviewRequested(v []string)`

SetReviewRequested sets ReviewRequested field to given value.

### HasReviewRequested

`func (o *NotificationsRequest) HasReviewRequested() bool`

HasReviewRequested returns a boolean if a field has been set.

### GetOther

`func (o *NotificationsRequest) GetOther() []string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *NotificationsRequest) GetOtherOk() (*[]string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *NotificationsRequest) SetOther(v []string)`

SetOther sets Other field to given value.

### HasOther

`func (o *NotificationsRequest) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetSlaBreach

`func (o *NotificationsRequest) GetSlaBreach() []string`

GetSlaBreach returns the SlaBreach field if non-nil, zero value otherwise.

### GetSlaBreachOk

`func (o *NotificationsRequest) GetSlaBreachOk() (*[]string, bool)`

GetSlaBreachOk returns a tuple with the SlaBreach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaBreach

`func (o *NotificationsRequest) SetSlaBreach(v []string)`

SetSlaBreach sets SlaBreach field to given value.

### HasSlaBreach

`func (o *NotificationsRequest) HasSlaBreach() bool`

HasSlaBreach returns a boolean if a field has been set.

### GetRiskAcceptanceExpiration

`func (o *NotificationsRequest) GetRiskAcceptanceExpiration() []string`

GetRiskAcceptanceExpiration returns the RiskAcceptanceExpiration field if non-nil, zero value otherwise.

### GetRiskAcceptanceExpirationOk

`func (o *NotificationsRequest) GetRiskAcceptanceExpirationOk() (*[]string, bool)`

GetRiskAcceptanceExpirationOk returns a tuple with the RiskAcceptanceExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceExpiration

`func (o *NotificationsRequest) SetRiskAcceptanceExpiration(v []string)`

SetRiskAcceptanceExpiration sets RiskAcceptanceExpiration field to given value.

### HasRiskAcceptanceExpiration

`func (o *NotificationsRequest) HasRiskAcceptanceExpiration() bool`

HasRiskAcceptanceExpiration returns a boolean if a field has been set.

### GetTemplate

`func (o *NotificationsRequest) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *NotificationsRequest) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *NotificationsRequest) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *NotificationsRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetScanAddedEmpty

`func (o *NotificationsRequest) GetScanAddedEmpty() string`

GetScanAddedEmpty returns the ScanAddedEmpty field if non-nil, zero value otherwise.

### GetScanAddedEmptyOk

`func (o *NotificationsRequest) GetScanAddedEmptyOk() (*string, bool)`

GetScanAddedEmptyOk returns a tuple with the ScanAddedEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAddedEmpty

`func (o *NotificationsRequest) SetScanAddedEmpty(v string)`

SetScanAddedEmpty sets ScanAddedEmpty field to given value.

### HasScanAddedEmpty

`func (o *NotificationsRequest) HasScanAddedEmpty() bool`

HasScanAddedEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


