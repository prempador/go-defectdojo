/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NotificationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsRequest{}

// NotificationsRequest struct for NotificationsRequest
type NotificationsRequest struct {
	Product NullableInt32 `json:"product,omitempty"`
	User NullableInt32 `json:"user,omitempty"`
	ProductTypeAdded []string `json:"product_type_added,omitempty"`
	ProductAdded []string `json:"product_added,omitempty"`
	EngagementAdded []string `json:"engagement_added,omitempty"`
	TestAdded []string `json:"test_added,omitempty"`
	ScanAdded []string `json:"scan_added,omitempty"`
	JiraUpdate []string `json:"jira_update,omitempty"`
	UpcomingEngagement []string `json:"upcoming_engagement,omitempty"`
	StaleEngagement []string `json:"stale_engagement,omitempty"`
	AutoCloseEngagement []string `json:"auto_close_engagement,omitempty"`
	CloseEngagement []string `json:"close_engagement,omitempty"`
	UserMentioned []string `json:"user_mentioned,omitempty"`
	CodeReview []string `json:"code_review,omitempty"`
	ReviewRequested []string `json:"review_requested,omitempty"`
	Other []string `json:"other,omitempty"`
	SlaBreach []string `json:"sla_breach,omitempty"`
	RiskAcceptanceExpiration []string `json:"risk_acceptance_expiration,omitempty"`
	Template *bool `json:"template,omitempty"`
	// Triggered whenever an (re-)import has been done (even if that created/updated/closed no findings).  * `slack` - slack * `msteams` - msteams * `mail` - mail * `alert` - alert
	ScanAddedEmpty *string `json:"scan_added_empty,omitempty"`
}

// NewNotificationsRequest instantiates a new NotificationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsRequest() *NotificationsRequest {
	this := NotificationsRequest{}
	var template bool = false
	this.Template = &template
	return &this
}

// NewNotificationsRequestWithDefaults instantiates a new NotificationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsRequestWithDefaults() *NotificationsRequest {
	this := NotificationsRequest{}
	var template bool = false
	this.Template = &template
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationsRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationsRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *NotificationsRequest) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *NotificationsRequest) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *NotificationsRequest) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *NotificationsRequest) UnsetProduct() {
	o.Product.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationsRequest) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationsRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *NotificationsRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *NotificationsRequest) SetUser(v int32) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *NotificationsRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *NotificationsRequest) UnsetUser() {
	o.User.Unset()
}

// GetProductTypeAdded returns the ProductTypeAdded field value if set, zero value otherwise.
func (o *NotificationsRequest) GetProductTypeAdded() []string {
	if o == nil || IsNil(o.ProductTypeAdded) {
		var ret []string
		return ret
	}
	return o.ProductTypeAdded
}

// GetProductTypeAddedOk returns a tuple with the ProductTypeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetProductTypeAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypeAdded) {
		return nil, false
	}
	return o.ProductTypeAdded, true
}

// HasProductTypeAdded returns a boolean if a field has been set.
func (o *NotificationsRequest) HasProductTypeAdded() bool {
	if o != nil && !IsNil(o.ProductTypeAdded) {
		return true
	}

	return false
}

// SetProductTypeAdded gets a reference to the given []string and assigns it to the ProductTypeAdded field.
func (o *NotificationsRequest) SetProductTypeAdded(v []string) {
	o.ProductTypeAdded = v
}

// GetProductAdded returns the ProductAdded field value if set, zero value otherwise.
func (o *NotificationsRequest) GetProductAdded() []string {
	if o == nil || IsNil(o.ProductAdded) {
		var ret []string
		return ret
	}
	return o.ProductAdded
}

// GetProductAddedOk returns a tuple with the ProductAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetProductAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductAdded) {
		return nil, false
	}
	return o.ProductAdded, true
}

// HasProductAdded returns a boolean if a field has been set.
func (o *NotificationsRequest) HasProductAdded() bool {
	if o != nil && !IsNil(o.ProductAdded) {
		return true
	}

	return false
}

// SetProductAdded gets a reference to the given []string and assigns it to the ProductAdded field.
func (o *NotificationsRequest) SetProductAdded(v []string) {
	o.ProductAdded = v
}

// GetEngagementAdded returns the EngagementAdded field value if set, zero value otherwise.
func (o *NotificationsRequest) GetEngagementAdded() []string {
	if o == nil || IsNil(o.EngagementAdded) {
		var ret []string
		return ret
	}
	return o.EngagementAdded
}

// GetEngagementAddedOk returns a tuple with the EngagementAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetEngagementAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.EngagementAdded) {
		return nil, false
	}
	return o.EngagementAdded, true
}

// HasEngagementAdded returns a boolean if a field has been set.
func (o *NotificationsRequest) HasEngagementAdded() bool {
	if o != nil && !IsNil(o.EngagementAdded) {
		return true
	}

	return false
}

// SetEngagementAdded gets a reference to the given []string and assigns it to the EngagementAdded field.
func (o *NotificationsRequest) SetEngagementAdded(v []string) {
	o.EngagementAdded = v
}

// GetTestAdded returns the TestAdded field value if set, zero value otherwise.
func (o *NotificationsRequest) GetTestAdded() []string {
	if o == nil || IsNil(o.TestAdded) {
		var ret []string
		return ret
	}
	return o.TestAdded
}

// GetTestAddedOk returns a tuple with the TestAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetTestAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.TestAdded) {
		return nil, false
	}
	return o.TestAdded, true
}

// HasTestAdded returns a boolean if a field has been set.
func (o *NotificationsRequest) HasTestAdded() bool {
	if o != nil && !IsNil(o.TestAdded) {
		return true
	}

	return false
}

// SetTestAdded gets a reference to the given []string and assigns it to the TestAdded field.
func (o *NotificationsRequest) SetTestAdded(v []string) {
	o.TestAdded = v
}

// GetScanAdded returns the ScanAdded field value if set, zero value otherwise.
func (o *NotificationsRequest) GetScanAdded() []string {
	if o == nil || IsNil(o.ScanAdded) {
		var ret []string
		return ret
	}
	return o.ScanAdded
}

// GetScanAddedOk returns a tuple with the ScanAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetScanAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanAdded) {
		return nil, false
	}
	return o.ScanAdded, true
}

// HasScanAdded returns a boolean if a field has been set.
func (o *NotificationsRequest) HasScanAdded() bool {
	if o != nil && !IsNil(o.ScanAdded) {
		return true
	}

	return false
}

// SetScanAdded gets a reference to the given []string and assigns it to the ScanAdded field.
func (o *NotificationsRequest) SetScanAdded(v []string) {
	o.ScanAdded = v
}

// GetJiraUpdate returns the JiraUpdate field value if set, zero value otherwise.
func (o *NotificationsRequest) GetJiraUpdate() []string {
	if o == nil || IsNil(o.JiraUpdate) {
		var ret []string
		return ret
	}
	return o.JiraUpdate
}

// GetJiraUpdateOk returns a tuple with the JiraUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetJiraUpdateOk() ([]string, bool) {
	if o == nil || IsNil(o.JiraUpdate) {
		return nil, false
	}
	return o.JiraUpdate, true
}

// HasJiraUpdate returns a boolean if a field has been set.
func (o *NotificationsRequest) HasJiraUpdate() bool {
	if o != nil && !IsNil(o.JiraUpdate) {
		return true
	}

	return false
}

// SetJiraUpdate gets a reference to the given []string and assigns it to the JiraUpdate field.
func (o *NotificationsRequest) SetJiraUpdate(v []string) {
	o.JiraUpdate = v
}

// GetUpcomingEngagement returns the UpcomingEngagement field value if set, zero value otherwise.
func (o *NotificationsRequest) GetUpcomingEngagement() []string {
	if o == nil || IsNil(o.UpcomingEngagement) {
		var ret []string
		return ret
	}
	return o.UpcomingEngagement
}

// GetUpcomingEngagementOk returns a tuple with the UpcomingEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetUpcomingEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.UpcomingEngagement) {
		return nil, false
	}
	return o.UpcomingEngagement, true
}

// HasUpcomingEngagement returns a boolean if a field has been set.
func (o *NotificationsRequest) HasUpcomingEngagement() bool {
	if o != nil && !IsNil(o.UpcomingEngagement) {
		return true
	}

	return false
}

// SetUpcomingEngagement gets a reference to the given []string and assigns it to the UpcomingEngagement field.
func (o *NotificationsRequest) SetUpcomingEngagement(v []string) {
	o.UpcomingEngagement = v
}

// GetStaleEngagement returns the StaleEngagement field value if set, zero value otherwise.
func (o *NotificationsRequest) GetStaleEngagement() []string {
	if o == nil || IsNil(o.StaleEngagement) {
		var ret []string
		return ret
	}
	return o.StaleEngagement
}

// GetStaleEngagementOk returns a tuple with the StaleEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetStaleEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.StaleEngagement) {
		return nil, false
	}
	return o.StaleEngagement, true
}

// HasStaleEngagement returns a boolean if a field has been set.
func (o *NotificationsRequest) HasStaleEngagement() bool {
	if o != nil && !IsNil(o.StaleEngagement) {
		return true
	}

	return false
}

// SetStaleEngagement gets a reference to the given []string and assigns it to the StaleEngagement field.
func (o *NotificationsRequest) SetStaleEngagement(v []string) {
	o.StaleEngagement = v
}

// GetAutoCloseEngagement returns the AutoCloseEngagement field value if set, zero value otherwise.
func (o *NotificationsRequest) GetAutoCloseEngagement() []string {
	if o == nil || IsNil(o.AutoCloseEngagement) {
		var ret []string
		return ret
	}
	return o.AutoCloseEngagement
}

// GetAutoCloseEngagementOk returns a tuple with the AutoCloseEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetAutoCloseEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoCloseEngagement) {
		return nil, false
	}
	return o.AutoCloseEngagement, true
}

// HasAutoCloseEngagement returns a boolean if a field has been set.
func (o *NotificationsRequest) HasAutoCloseEngagement() bool {
	if o != nil && !IsNil(o.AutoCloseEngagement) {
		return true
	}

	return false
}

// SetAutoCloseEngagement gets a reference to the given []string and assigns it to the AutoCloseEngagement field.
func (o *NotificationsRequest) SetAutoCloseEngagement(v []string) {
	o.AutoCloseEngagement = v
}

// GetCloseEngagement returns the CloseEngagement field value if set, zero value otherwise.
func (o *NotificationsRequest) GetCloseEngagement() []string {
	if o == nil || IsNil(o.CloseEngagement) {
		var ret []string
		return ret
	}
	return o.CloseEngagement
}

// GetCloseEngagementOk returns a tuple with the CloseEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetCloseEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.CloseEngagement) {
		return nil, false
	}
	return o.CloseEngagement, true
}

// HasCloseEngagement returns a boolean if a field has been set.
func (o *NotificationsRequest) HasCloseEngagement() bool {
	if o != nil && !IsNil(o.CloseEngagement) {
		return true
	}

	return false
}

// SetCloseEngagement gets a reference to the given []string and assigns it to the CloseEngagement field.
func (o *NotificationsRequest) SetCloseEngagement(v []string) {
	o.CloseEngagement = v
}

// GetUserMentioned returns the UserMentioned field value if set, zero value otherwise.
func (o *NotificationsRequest) GetUserMentioned() []string {
	if o == nil || IsNil(o.UserMentioned) {
		var ret []string
		return ret
	}
	return o.UserMentioned
}

// GetUserMentionedOk returns a tuple with the UserMentioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetUserMentionedOk() ([]string, bool) {
	if o == nil || IsNil(o.UserMentioned) {
		return nil, false
	}
	return o.UserMentioned, true
}

// HasUserMentioned returns a boolean if a field has been set.
func (o *NotificationsRequest) HasUserMentioned() bool {
	if o != nil && !IsNil(o.UserMentioned) {
		return true
	}

	return false
}

// SetUserMentioned gets a reference to the given []string and assigns it to the UserMentioned field.
func (o *NotificationsRequest) SetUserMentioned(v []string) {
	o.UserMentioned = v
}

// GetCodeReview returns the CodeReview field value if set, zero value otherwise.
func (o *NotificationsRequest) GetCodeReview() []string {
	if o == nil || IsNil(o.CodeReview) {
		var ret []string
		return ret
	}
	return o.CodeReview
}

// GetCodeReviewOk returns a tuple with the CodeReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetCodeReviewOk() ([]string, bool) {
	if o == nil || IsNil(o.CodeReview) {
		return nil, false
	}
	return o.CodeReview, true
}

// HasCodeReview returns a boolean if a field has been set.
func (o *NotificationsRequest) HasCodeReview() bool {
	if o != nil && !IsNil(o.CodeReview) {
		return true
	}

	return false
}

// SetCodeReview gets a reference to the given []string and assigns it to the CodeReview field.
func (o *NotificationsRequest) SetCodeReview(v []string) {
	o.CodeReview = v
}

// GetReviewRequested returns the ReviewRequested field value if set, zero value otherwise.
func (o *NotificationsRequest) GetReviewRequested() []string {
	if o == nil || IsNil(o.ReviewRequested) {
		var ret []string
		return ret
	}
	return o.ReviewRequested
}

// GetReviewRequestedOk returns a tuple with the ReviewRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetReviewRequestedOk() ([]string, bool) {
	if o == nil || IsNil(o.ReviewRequested) {
		return nil, false
	}
	return o.ReviewRequested, true
}

// HasReviewRequested returns a boolean if a field has been set.
func (o *NotificationsRequest) HasReviewRequested() bool {
	if o != nil && !IsNil(o.ReviewRequested) {
		return true
	}

	return false
}

// SetReviewRequested gets a reference to the given []string and assigns it to the ReviewRequested field.
func (o *NotificationsRequest) SetReviewRequested(v []string) {
	o.ReviewRequested = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *NotificationsRequest) GetOther() []string {
	if o == nil || IsNil(o.Other) {
		var ret []string
		return ret
	}
	return o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetOtherOk() ([]string, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *NotificationsRequest) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given []string and assigns it to the Other field.
func (o *NotificationsRequest) SetOther(v []string) {
	o.Other = v
}

// GetSlaBreach returns the SlaBreach field value if set, zero value otherwise.
func (o *NotificationsRequest) GetSlaBreach() []string {
	if o == nil || IsNil(o.SlaBreach) {
		var ret []string
		return ret
	}
	return o.SlaBreach
}

// GetSlaBreachOk returns a tuple with the SlaBreach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetSlaBreachOk() ([]string, bool) {
	if o == nil || IsNil(o.SlaBreach) {
		return nil, false
	}
	return o.SlaBreach, true
}

// HasSlaBreach returns a boolean if a field has been set.
func (o *NotificationsRequest) HasSlaBreach() bool {
	if o != nil && !IsNil(o.SlaBreach) {
		return true
	}

	return false
}

// SetSlaBreach gets a reference to the given []string and assigns it to the SlaBreach field.
func (o *NotificationsRequest) SetSlaBreach(v []string) {
	o.SlaBreach = v
}

// GetRiskAcceptanceExpiration returns the RiskAcceptanceExpiration field value if set, zero value otherwise.
func (o *NotificationsRequest) GetRiskAcceptanceExpiration() []string {
	if o == nil || IsNil(o.RiskAcceptanceExpiration) {
		var ret []string
		return ret
	}
	return o.RiskAcceptanceExpiration
}

// GetRiskAcceptanceExpirationOk returns a tuple with the RiskAcceptanceExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetRiskAcceptanceExpirationOk() ([]string, bool) {
	if o == nil || IsNil(o.RiskAcceptanceExpiration) {
		return nil, false
	}
	return o.RiskAcceptanceExpiration, true
}

// HasRiskAcceptanceExpiration returns a boolean if a field has been set.
func (o *NotificationsRequest) HasRiskAcceptanceExpiration() bool {
	if o != nil && !IsNil(o.RiskAcceptanceExpiration) {
		return true
	}

	return false
}

// SetRiskAcceptanceExpiration gets a reference to the given []string and assigns it to the RiskAcceptanceExpiration field.
func (o *NotificationsRequest) SetRiskAcceptanceExpiration(v []string) {
	o.RiskAcceptanceExpiration = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *NotificationsRequest) GetTemplate() bool {
	if o == nil || IsNil(o.Template) {
		var ret bool
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *NotificationsRequest) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given bool and assigns it to the Template field.
func (o *NotificationsRequest) SetTemplate(v bool) {
	o.Template = &v
}

// GetScanAddedEmpty returns the ScanAddedEmpty field value if set, zero value otherwise.
func (o *NotificationsRequest) GetScanAddedEmpty() string {
	if o == nil || IsNil(o.ScanAddedEmpty) {
		var ret string
		return ret
	}
	return *o.ScanAddedEmpty
}

// GetScanAddedEmptyOk returns a tuple with the ScanAddedEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsRequest) GetScanAddedEmptyOk() (*string, bool) {
	if o == nil || IsNil(o.ScanAddedEmpty) {
		return nil, false
	}
	return o.ScanAddedEmpty, true
}

// HasScanAddedEmpty returns a boolean if a field has been set.
func (o *NotificationsRequest) HasScanAddedEmpty() bool {
	if o != nil && !IsNil(o.ScanAddedEmpty) {
		return true
	}

	return false
}

// SetScanAddedEmpty gets a reference to the given string and assigns it to the ScanAddedEmpty field.
func (o *NotificationsRequest) SetScanAddedEmpty(v string) {
	o.ScanAddedEmpty = &v
}

func (o NotificationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if !IsNil(o.ProductTypeAdded) {
		toSerialize["product_type_added"] = o.ProductTypeAdded
	}
	if !IsNil(o.ProductAdded) {
		toSerialize["product_added"] = o.ProductAdded
	}
	if !IsNil(o.EngagementAdded) {
		toSerialize["engagement_added"] = o.EngagementAdded
	}
	if !IsNil(o.TestAdded) {
		toSerialize["test_added"] = o.TestAdded
	}
	if !IsNil(o.ScanAdded) {
		toSerialize["scan_added"] = o.ScanAdded
	}
	if !IsNil(o.JiraUpdate) {
		toSerialize["jira_update"] = o.JiraUpdate
	}
	if !IsNil(o.UpcomingEngagement) {
		toSerialize["upcoming_engagement"] = o.UpcomingEngagement
	}
	if !IsNil(o.StaleEngagement) {
		toSerialize["stale_engagement"] = o.StaleEngagement
	}
	if !IsNil(o.AutoCloseEngagement) {
		toSerialize["auto_close_engagement"] = o.AutoCloseEngagement
	}
	if !IsNil(o.CloseEngagement) {
		toSerialize["close_engagement"] = o.CloseEngagement
	}
	if !IsNil(o.UserMentioned) {
		toSerialize["user_mentioned"] = o.UserMentioned
	}
	if !IsNil(o.CodeReview) {
		toSerialize["code_review"] = o.CodeReview
	}
	if !IsNil(o.ReviewRequested) {
		toSerialize["review_requested"] = o.ReviewRequested
	}
	if !IsNil(o.Other) {
		toSerialize["other"] = o.Other
	}
	if !IsNil(o.SlaBreach) {
		toSerialize["sla_breach"] = o.SlaBreach
	}
	if !IsNil(o.RiskAcceptanceExpiration) {
		toSerialize["risk_acceptance_expiration"] = o.RiskAcceptanceExpiration
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.ScanAddedEmpty) {
		toSerialize["scan_added_empty"] = o.ScanAddedEmpty
	}
	return toSerialize, nil
}

type NullableNotificationsRequest struct {
	value *NotificationsRequest
	isSet bool
}

func (v NullableNotificationsRequest) Get() *NotificationsRequest {
	return v.value
}

func (v *NullableNotificationsRequest) Set(val *NotificationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsRequest(val *NotificationsRequest) *NullableNotificationsRequest {
	return &NullableNotificationsRequest{value: val, isSet: true}
}

func (v NullableNotificationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

