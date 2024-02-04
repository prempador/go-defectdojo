# PatchedNotificationsRequest

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

### NewPatchedNotificationsRequest

`func NewPatchedNotificationsRequest() *PatchedNotificationsRequest`

NewPatchedNotificationsRequest instantiates a new PatchedNotificationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotificationsRequestWithDefaults

`func NewPatchedNotificationsRequestWithDefaults() *PatchedNotificationsRequest`

NewPatchedNotificationsRequestWithDefaults instantiates a new PatchedNotificationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *PatchedNotificationsRequest) GetProduct() int32`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PatchedNotificationsRequest) GetProductOk() (*int32, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PatchedNotificationsRequest) SetProduct(v int32)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PatchedNotificationsRequest) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *PatchedNotificationsRequest) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *PatchedNotificationsRequest) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetUser

`func (o *PatchedNotificationsRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedNotificationsRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedNotificationsRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedNotificationsRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedNotificationsRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedNotificationsRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetProductTypeAdded

`func (o *PatchedNotificationsRequest) GetProductTypeAdded() []string`

GetProductTypeAdded returns the ProductTypeAdded field if non-nil, zero value otherwise.

### GetProductTypeAddedOk

`func (o *PatchedNotificationsRequest) GetProductTypeAddedOk() (*[]string, bool)`

GetProductTypeAddedOk returns a tuple with the ProductTypeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypeAdded

`func (o *PatchedNotificationsRequest) SetProductTypeAdded(v []string)`

SetProductTypeAdded sets ProductTypeAdded field to given value.

### HasProductTypeAdded

`func (o *PatchedNotificationsRequest) HasProductTypeAdded() bool`

HasProductTypeAdded returns a boolean if a field has been set.

### GetProductAdded

`func (o *PatchedNotificationsRequest) GetProductAdded() []string`

GetProductAdded returns the ProductAdded field if non-nil, zero value otherwise.

### GetProductAddedOk

`func (o *PatchedNotificationsRequest) GetProductAddedOk() (*[]string, bool)`

GetProductAddedOk returns a tuple with the ProductAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAdded

`func (o *PatchedNotificationsRequest) SetProductAdded(v []string)`

SetProductAdded sets ProductAdded field to given value.

### HasProductAdded

`func (o *PatchedNotificationsRequest) HasProductAdded() bool`

HasProductAdded returns a boolean if a field has been set.

### GetEngagementAdded

`func (o *PatchedNotificationsRequest) GetEngagementAdded() []string`

GetEngagementAdded returns the EngagementAdded field if non-nil, zero value otherwise.

### GetEngagementAddedOk

`func (o *PatchedNotificationsRequest) GetEngagementAddedOk() (*[]string, bool)`

GetEngagementAddedOk returns a tuple with the EngagementAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementAdded

`func (o *PatchedNotificationsRequest) SetEngagementAdded(v []string)`

SetEngagementAdded sets EngagementAdded field to given value.

### HasEngagementAdded

`func (o *PatchedNotificationsRequest) HasEngagementAdded() bool`

HasEngagementAdded returns a boolean if a field has been set.

### GetTestAdded

`func (o *PatchedNotificationsRequest) GetTestAdded() []string`

GetTestAdded returns the TestAdded field if non-nil, zero value otherwise.

### GetTestAddedOk

`func (o *PatchedNotificationsRequest) GetTestAddedOk() (*[]string, bool)`

GetTestAddedOk returns a tuple with the TestAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAdded

`func (o *PatchedNotificationsRequest) SetTestAdded(v []string)`

SetTestAdded sets TestAdded field to given value.

### HasTestAdded

`func (o *PatchedNotificationsRequest) HasTestAdded() bool`

HasTestAdded returns a boolean if a field has been set.

### GetScanAdded

`func (o *PatchedNotificationsRequest) GetScanAdded() []string`

GetScanAdded returns the ScanAdded field if non-nil, zero value otherwise.

### GetScanAddedOk

`func (o *PatchedNotificationsRequest) GetScanAddedOk() (*[]string, bool)`

GetScanAddedOk returns a tuple with the ScanAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAdded

`func (o *PatchedNotificationsRequest) SetScanAdded(v []string)`

SetScanAdded sets ScanAdded field to given value.

### HasScanAdded

`func (o *PatchedNotificationsRequest) HasScanAdded() bool`

HasScanAdded returns a boolean if a field has been set.

### GetJiraUpdate

`func (o *PatchedNotificationsRequest) GetJiraUpdate() []string`

GetJiraUpdate returns the JiraUpdate field if non-nil, zero value otherwise.

### GetJiraUpdateOk

`func (o *PatchedNotificationsRequest) GetJiraUpdateOk() (*[]string, bool)`

GetJiraUpdateOk returns a tuple with the JiraUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraUpdate

`func (o *PatchedNotificationsRequest) SetJiraUpdate(v []string)`

SetJiraUpdate sets JiraUpdate field to given value.

### HasJiraUpdate

`func (o *PatchedNotificationsRequest) HasJiraUpdate() bool`

HasJiraUpdate returns a boolean if a field has been set.

### GetUpcomingEngagement

`func (o *PatchedNotificationsRequest) GetUpcomingEngagement() []string`

GetUpcomingEngagement returns the UpcomingEngagement field if non-nil, zero value otherwise.

### GetUpcomingEngagementOk

`func (o *PatchedNotificationsRequest) GetUpcomingEngagementOk() (*[]string, bool)`

GetUpcomingEngagementOk returns a tuple with the UpcomingEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingEngagement

`func (o *PatchedNotificationsRequest) SetUpcomingEngagement(v []string)`

SetUpcomingEngagement sets UpcomingEngagement field to given value.

### HasUpcomingEngagement

`func (o *PatchedNotificationsRequest) HasUpcomingEngagement() bool`

HasUpcomingEngagement returns a boolean if a field has been set.

### GetStaleEngagement

`func (o *PatchedNotificationsRequest) GetStaleEngagement() []string`

GetStaleEngagement returns the StaleEngagement field if non-nil, zero value otherwise.

### GetStaleEngagementOk

`func (o *PatchedNotificationsRequest) GetStaleEngagementOk() (*[]string, bool)`

GetStaleEngagementOk returns a tuple with the StaleEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleEngagement

`func (o *PatchedNotificationsRequest) SetStaleEngagement(v []string)`

SetStaleEngagement sets StaleEngagement field to given value.

### HasStaleEngagement

`func (o *PatchedNotificationsRequest) HasStaleEngagement() bool`

HasStaleEngagement returns a boolean if a field has been set.

### GetAutoCloseEngagement

`func (o *PatchedNotificationsRequest) GetAutoCloseEngagement() []string`

GetAutoCloseEngagement returns the AutoCloseEngagement field if non-nil, zero value otherwise.

### GetAutoCloseEngagementOk

`func (o *PatchedNotificationsRequest) GetAutoCloseEngagementOk() (*[]string, bool)`

GetAutoCloseEngagementOk returns a tuple with the AutoCloseEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCloseEngagement

`func (o *PatchedNotificationsRequest) SetAutoCloseEngagement(v []string)`

SetAutoCloseEngagement sets AutoCloseEngagement field to given value.

### HasAutoCloseEngagement

`func (o *PatchedNotificationsRequest) HasAutoCloseEngagement() bool`

HasAutoCloseEngagement returns a boolean if a field has been set.

### GetCloseEngagement

`func (o *PatchedNotificationsRequest) GetCloseEngagement() []string`

GetCloseEngagement returns the CloseEngagement field if non-nil, zero value otherwise.

### GetCloseEngagementOk

`func (o *PatchedNotificationsRequest) GetCloseEngagementOk() (*[]string, bool)`

GetCloseEngagementOk returns a tuple with the CloseEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseEngagement

`func (o *PatchedNotificationsRequest) SetCloseEngagement(v []string)`

SetCloseEngagement sets CloseEngagement field to given value.

### HasCloseEngagement

`func (o *PatchedNotificationsRequest) HasCloseEngagement() bool`

HasCloseEngagement returns a boolean if a field has been set.

### GetUserMentioned

`func (o *PatchedNotificationsRequest) GetUserMentioned() []string`

GetUserMentioned returns the UserMentioned field if non-nil, zero value otherwise.

### GetUserMentionedOk

`func (o *PatchedNotificationsRequest) GetUserMentionedOk() (*[]string, bool)`

GetUserMentionedOk returns a tuple with the UserMentioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMentioned

`func (o *PatchedNotificationsRequest) SetUserMentioned(v []string)`

SetUserMentioned sets UserMentioned field to given value.

### HasUserMentioned

`func (o *PatchedNotificationsRequest) HasUserMentioned() bool`

HasUserMentioned returns a boolean if a field has been set.

### GetCodeReview

`func (o *PatchedNotificationsRequest) GetCodeReview() []string`

GetCodeReview returns the CodeReview field if non-nil, zero value otherwise.

### GetCodeReviewOk

`func (o *PatchedNotificationsRequest) GetCodeReviewOk() (*[]string, bool)`

GetCodeReviewOk returns a tuple with the CodeReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeReview

`func (o *PatchedNotificationsRequest) SetCodeReview(v []string)`

SetCodeReview sets CodeReview field to given value.

### HasCodeReview

`func (o *PatchedNotificationsRequest) HasCodeReview() bool`

HasCodeReview returns a boolean if a field has been set.

### GetReviewRequested

`func (o *PatchedNotificationsRequest) GetReviewRequested() []string`

GetReviewRequested returns the ReviewRequested field if non-nil, zero value otherwise.

### GetReviewRequestedOk

`func (o *PatchedNotificationsRequest) GetReviewRequestedOk() (*[]string, bool)`

GetReviewRequestedOk returns a tuple with the ReviewRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequested

`func (o *PatchedNotificationsRequest) SetReviewRequested(v []string)`

SetReviewRequested sets ReviewRequested field to given value.

### HasReviewRequested

`func (o *PatchedNotificationsRequest) HasReviewRequested() bool`

HasReviewRequested returns a boolean if a field has been set.

### GetOther

`func (o *PatchedNotificationsRequest) GetOther() []string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *PatchedNotificationsRequest) GetOtherOk() (*[]string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *PatchedNotificationsRequest) SetOther(v []string)`

SetOther sets Other field to given value.

### HasOther

`func (o *PatchedNotificationsRequest) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetSlaBreach

`func (o *PatchedNotificationsRequest) GetSlaBreach() []string`

GetSlaBreach returns the SlaBreach field if non-nil, zero value otherwise.

### GetSlaBreachOk

`func (o *PatchedNotificationsRequest) GetSlaBreachOk() (*[]string, bool)`

GetSlaBreachOk returns a tuple with the SlaBreach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaBreach

`func (o *PatchedNotificationsRequest) SetSlaBreach(v []string)`

SetSlaBreach sets SlaBreach field to given value.

### HasSlaBreach

`func (o *PatchedNotificationsRequest) HasSlaBreach() bool`

HasSlaBreach returns a boolean if a field has been set.

### GetRiskAcceptanceExpiration

`func (o *PatchedNotificationsRequest) GetRiskAcceptanceExpiration() []string`

GetRiskAcceptanceExpiration returns the RiskAcceptanceExpiration field if non-nil, zero value otherwise.

### GetRiskAcceptanceExpirationOk

`func (o *PatchedNotificationsRequest) GetRiskAcceptanceExpirationOk() (*[]string, bool)`

GetRiskAcceptanceExpirationOk returns a tuple with the RiskAcceptanceExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAcceptanceExpiration

`func (o *PatchedNotificationsRequest) SetRiskAcceptanceExpiration(v []string)`

SetRiskAcceptanceExpiration sets RiskAcceptanceExpiration field to given value.

### HasRiskAcceptanceExpiration

`func (o *PatchedNotificationsRequest) HasRiskAcceptanceExpiration() bool`

HasRiskAcceptanceExpiration returns a boolean if a field has been set.

### GetTemplate

`func (o *PatchedNotificationsRequest) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PatchedNotificationsRequest) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PatchedNotificationsRequest) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PatchedNotificationsRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetScanAddedEmpty

`func (o *PatchedNotificationsRequest) GetScanAddedEmpty() string`

GetScanAddedEmpty returns the ScanAddedEmpty field if non-nil, zero value otherwise.

### GetScanAddedEmptyOk

`func (o *PatchedNotificationsRequest) GetScanAddedEmptyOk() (*string, bool)`

GetScanAddedEmptyOk returns a tuple with the ScanAddedEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanAddedEmpty

`func (o *PatchedNotificationsRequest) SetScanAddedEmpty(v string)`

SetScanAddedEmpty sets ScanAddedEmpty field to given value.

### HasScanAddedEmpty

`func (o *PatchedNotificationsRequest) HasScanAddedEmpty() bool`

HasScanAddedEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


