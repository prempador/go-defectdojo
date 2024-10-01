# Notifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
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
**SlaBreachCombined** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**RiskAcceptanceExpiration** | Pointer to **[]string** |  | [optional] [default to ["alert"]]
**Template** | Pointer to **bool** |  | [optional] [default to false]
**ScanAddedEmpty** | Pointer to **string** | Triggered whenever an (re-)import has been done (even if that created/updated/closed no findings).  * &#x60;slack&#x60; - slack * &#x60;msteams&#x60; - msteams * &#x60;mail&#x60; - mail * &#x60;alert&#x60; - alert | [optional] 
**Prefetch** | Pointer to [**AppAnalysisPrefetch**](AppAnalysisPrefetch.md) |  | [optional] 

## Methods

### NewNotifications

`func NewNotifications(id int32, ) *Notifications`

NewNotifications instantiates a new Notifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsWithDefaults

`func NewNotificationsWithDefaults() *Notifications`

NewNotificationsWithDefaults instantiates a new Notifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Notifications) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notifications) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notifications) SetId(v int32)`

SetId sets Id field to given value.


### GetProduct

`func (o *Notifications) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Notifications) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Notifications) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Notifications) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *Notifications) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *Notifications) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetUser

`func (o *Notifications) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Notifications) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Notifications) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *Notifications) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *Notifications) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Notifications) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetProductTypeAdded

`func (o *Notifications) GetProductTypeAdded() []string`

GetProductTypeAdded returns the ProductTypeAdded field if non-nil, zero value otherwise.

### GetProductTypeAddedOk

`func (o *Notifications) GetProductTypeAddedOk() (*[]string, bool)`

GetProductTypeAddedOk returns a tuple with the ProductTypeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeAdded

`func (o *Notifications) SetProductTypeAdded(v []string)`

SetProductTypeAdded sets ProductTypeAdded field to given value.

### HasProductTypeAdded

`func (o *Notifications) HasProductTypeAdded() bool`

HasProductTypeAdded returns a boolean if a field has been set.

### GetProductAdded

`func (o *Notifications) GetProductAdded() []string`

GetProductAdded returns the ProductAdded field if non-nil, zero value otherwise.

### GetProductAddedOk

`func (o *Notifications) GetProductAddedOk() (*[]string, bool)`

GetProductAddedOk returns a tuple with the ProductAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAdded

`func (o *Notifications) SetProductAdded(v []string)`

SetProductAdded sets ProductAdded field to given value.

### HasProductAdded

`func (o *Notifications) HasProductAdded() bool`

HasProductAdded returns a boolean if a field has been set.

### GetEngagementAdded

`func (o *Notifications) GetEngagementAdded() []string`

GetEngagementAdded returns the EngagementAdded field if non-nil, zero value otherwise.

### GetEngagementAddedOk

`func (o *Notifications) GetEngagementAddedOk() (*[]string, bool)`

GetEngagementAddedOk returns a tuple with the EngagementAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementAdded

`func (o *Notifications) SetEngagementAdded(v []string)`

SetEngagementAdded sets EngagementAdded field to given value.

### HasEngagementAdded

`func (o *Notifications) HasEngagementAdded() bool`

HasEngagementAdded returns a boolean if a field has been set.

### GetTestAdded

`func (o *Notifications) GetTestAdded() []string`

GetTestAdded returns the TestAdded field if non-nil, zero value otherwise.

### GetTestAddedOk

`func (o *Notifications) GetTestAddedOk() (*[]string, bool)`

GetTestAddedOk returns a tuple with the TestAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAdded

`func (o *Notifications) SetTestAdded(v []string)`

SetTestAdded sets TestAdded field to given value.

### HasTestAdded

`func (o *Notifications) HasTestAdded() bool`

HasTestAdded returns a boolean if a field has been set.

### GetScanAdded

`func (o *Notifications) GetScanAdded() []string`

GetScanAdded returns the ScanAdded field if non-nil, zero value otherwise.

### GetScanAddedOk

`func (o *Notifications) GetScanAddedOk() (*[]string, bool)`

GetScanAddedOk returns a tuple with the ScanAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAdded

`func (o *Notifications) SetScanAdded(v []string)`

SetScanAdded sets ScanAdded field to given value.

### HasScanAdded

`func (o *Notifications) HasScanAdded() bool`

HasScanAdded returns a boolean if a field has been set.

### GetJiraUpdate

`func (o *Notifications) GetJiraUpdate() []string`

GetJiraUpdate returns the JiraUpdate field if non-nil, zero value otherwise.

### GetJiraUpdateOk

`func (o *Notifications) GetJiraUpdateOk() (*[]string, bool)`

GetJiraUpdateOk returns a tuple with the JiraUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraUpdate

`func (o *Notifications) SetJiraUpdate(v []string)`

SetJiraUpdate sets JiraUpdate field to given value.

### HasJiraUpdate

`func (o *Notifications) HasJiraUpdate() bool`

HasJiraUpdate returns a boolean if a field has been set.

### GetUpcomingEngagement

`func (o *Notifications) GetUpcomingEngagement() []string`

GetUpcomingEngagement returns the UpcomingEngagement field if non-nil, zero value otherwise.

### GetUpcomingEngagementOk

`func (o *Notifications) GetUpcomingEngagementOk() (*[]string, bool)`

GetUpcomingEngagementOk returns a tuple with the UpcomingEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingEngagement

`func (o *Notifications) SetUpcomingEngagement(v []string)`

SetUpcomingEngagement sets UpcomingEngagement field to given value.

### HasUpcomingEngagement

`func (o *Notifications) HasUpcomingEngagement() bool`

HasUpcomingEngagement returns a boolean if a field has been set.

### GetStaleEngagement

`func (o *Notifications) GetStaleEngagement() []string`

GetStaleEngagement returns the StaleEngagement field if non-nil, zero value otherwise.

### GetStaleEngagementOk

`func (o *Notifications) GetStaleEngagementOk() (*[]string, bool)`

GetStaleEngagementOk returns a tuple with the StaleEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleEngagement

`func (o *Notifications) SetStaleEngagement(v []string)`

SetStaleEngagement sets StaleEngagement field to given value.

### HasStaleEngagement

`func (o *Notifications) HasStaleEngagement() bool`

HasStaleEngagement returns a boolean if a field has been set.

### GetAutoCloseEngagement

`func (o *Notifications) GetAutoCloseEngagement() []string`

GetAutoCloseEngagement returns the AutoCloseEngagement field if non-nil, zero value otherwise.

### GetAutoCloseEngagementOk

`func (o *Notifications) GetAutoCloseEngagementOk() (*[]string, bool)`

GetAutoCloseEngagementOk returns a tuple with the AutoCloseEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCloseEngagement

`func (o *Notifications) SetAutoCloseEngagement(v []string)`

SetAutoCloseEngagement sets AutoCloseEngagement field to given value.

### HasAutoCloseEngagement

`func (o *Notifications) HasAutoCloseEngagement() bool`

HasAutoCloseEngagement returns a boolean if a field has been set.

### GetCloseEngagement

`func (o *Notifications) GetCloseEngagement() []string`

GetCloseEngagement returns the CloseEngagement field if non-nil, zero value otherwise.

### GetCloseEngagementOk

`func (o *Notifications) GetCloseEngagementOk() (*[]string, bool)`

GetCloseEngagementOk returns a tuple with the CloseEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseEngagement

`func (o *Notifications) SetCloseEngagement(v []string)`

SetCloseEngagement sets CloseEngagement field to given value.

### HasCloseEngagement

`func (o *Notifications) HasCloseEngagement() bool`

HasCloseEngagement returns a boolean if a field has been set.

### GetUserMentioned

`func (o *Notifications) GetUserMentioned() []string`

GetUserMentioned returns the UserMentioned field if non-nil, zero value otherwise.

### GetUserMentionedOk

`func (o *Notifications) GetUserMentionedOk() (*[]string, bool)`

GetUserMentionedOk returns a tuple with the UserMentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMentioned

`func (o *Notifications) SetUserMentioned(v []string)`

SetUserMentioned sets UserMentioned field to given value.

### HasUserMentioned

`func (o *Notifications) HasUserMentioned() bool`

HasUserMentioned returns a boolean if a field has been set.

### GetCodeReview

`func (o *Notifications) GetCodeReview() []string`

GetCodeReview returns the CodeReview field if non-nil, zero value otherwise.

### GetCodeReviewOk

`func (o *Notifications) GetCodeReviewOk() (*[]string, bool)`

GetCodeReviewOk returns a tuple with the CodeReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeReview

`func (o *Notifications) SetCodeReview(v []string)`

SetCodeReview sets CodeReview field to given value.

### HasCodeReview

`func (o *Notifications) HasCodeReview() bool`

HasCodeReview returns a boolean if a field has been set.

### GetReviewRequested

`func (o *Notifications) GetReviewRequested() []string`

GetReviewRequested returns the ReviewRequested field if non-nil, zero value otherwise.

### GetReviewRequestedOk

`func (o *Notifications) GetReviewRequestedOk() (*[]string, bool)`

GetReviewRequestedOk returns a tuple with the ReviewRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequested

`func (o *Notifications) SetReviewRequested(v []string)`

SetReviewRequested sets ReviewRequested field to given value.

### HasReviewRequested

`func (o *Notifications) HasReviewRequested() bool`

HasReviewRequested returns a boolean if a field has been set.

### GetOther

`func (o *Notifications) GetOther() []string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *Notifications) GetOtherOk() (*[]string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *Notifications) SetOther(v []string)`

SetOther sets Other field to given value.

### HasOther

`func (o *Notifications) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetSlaBreach

`func (o *Notifications) GetSlaBreach() []string`

GetSlaBreach returns the SlaBreach field if non-nil, zero value otherwise.

### GetSlaBreachOk

`func (o *Notifications) GetSlaBreachOk() (*[]string, bool)`

GetSlaBreachOk returns a tuple with the SlaBreach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaBreach

`func (o *Notifications) SetSlaBreach(v []string)`

SetSlaBreach sets SlaBreach field to given value.

### HasSlaBreach

`func (o *Notifications) HasSlaBreach() bool`

HasSlaBreach returns a boolean if a field has been set.

### GetSlaBreachCombined

`func (o *Notifications) GetSlaBreachCombined() []string`

GetSlaBreachCombined returns the SlaBreachCombined field if non-nil, zero value otherwise.

### GetSlaBreachCombinedOk

`func (o *Notifications) GetSlaBreachCombinedOk() (*[]string, bool)`

GetSlaBreachCombinedOk returns a tuple with the SlaBreachCombined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaBreachCombined

`func (o *Notifications) SetSlaBreachCombined(v []string)`

SetSlaBreachCombined sets SlaBreachCombined field to given value.

### HasSlaBreachCombined

`func (o *Notifications) HasSlaBreachCombined() bool`

HasSlaBreachCombined returns a boolean if a field has been set.

### GetRiskAcceptanceExpiration

`func (o *Notifications) GetRiskAcceptanceExpiration() []string`

GetRiskAcceptanceExpiration returns the RiskAcceptanceExpiration field if non-nil, zero value otherwise.

### GetRiskAcceptanceExpirationOk

`func (o *Notifications) GetRiskAcceptanceExpirationOk() (*[]string, bool)`

GetRiskAcceptanceExpirationOk returns a tuple with the RiskAcceptanceExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceExpiration

`func (o *Notifications) SetRiskAcceptanceExpiration(v []string)`

SetRiskAcceptanceExpiration sets RiskAcceptanceExpiration field to given value.

### HasRiskAcceptanceExpiration

`func (o *Notifications) HasRiskAcceptanceExpiration() bool`

HasRiskAcceptanceExpiration returns a boolean if a field has been set.

### GetTemplate

`func (o *Notifications) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Notifications) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Notifications) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Notifications) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetScanAddedEmpty

`func (o *Notifications) GetScanAddedEmpty() string`

GetScanAddedEmpty returns the ScanAddedEmpty field if non-nil, zero value otherwise.

### GetScanAddedEmptyOk

`func (o *Notifications) GetScanAddedEmptyOk() (*string, bool)`

GetScanAddedEmptyOk returns a tuple with the ScanAddedEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAddedEmpty

`func (o *Notifications) SetScanAddedEmpty(v string)`

SetScanAddedEmpty sets ScanAddedEmpty field to given value.

### HasScanAddedEmpty

`func (o *Notifications) HasScanAddedEmpty() bool`

HasScanAddedEmpty returns a boolean if a field has been set.

### GetPrefetch

`func (o *Notifications) GetPrefetch() AppAnalysisPrefetch`

GetPrefetch returns the Prefetch field if non-nil, zero value otherwise.

### GetPrefetchOk

`func (o *Notifications) GetPrefetchOk() (*AppAnalysisPrefetch, bool)`

GetPrefetchOk returns a tuple with the Prefetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetch

`func (o *Notifications) SetPrefetch(v AppAnalysisPrefetch)`

SetPrefetch sets Prefetch field to given value.

### HasPrefetch

`func (o *Notifications) HasPrefetch() bool`

HasPrefetch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


