/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedNotificationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedNotificationsRequest{}

// PatchedNotificationsRequest struct for PatchedNotificationsRequest
type PatchedNotificationsRequest struct {
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
	SlaBreachCombined []string `json:"sla_breach_combined,omitempty"`
	RiskAcceptanceExpiration []string `json:"risk_acceptance_expiration,omitempty"`
	Template *bool `json:"template,omitempty"`
	// Triggered whenever an (re-)import has been done (even if that created/updated/closed no findings).  * `slack` - slack * `msteams` - msteams * `mail` - mail * `alert` - alert
	ScanAddedEmpty *string `json:"scan_added_empty,omitempty"`
}

// NewPatchedNotificationsRequest instantiates a new PatchedNotificationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNotificationsRequest() *PatchedNotificationsRequest {
	this := PatchedNotificationsRequest{}
	var template bool = false
	this.Template = &template
	return &this
}

// NewPatchedNotificationsRequestWithDefaults instantiates a new PatchedNotificationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNotificationsRequestWithDefaults() *PatchedNotificationsRequest {
	this := PatchedNotificationsRequest{}
	var template bool = false
	this.Template = &template
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedNotificationsRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedNotificationsRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *PatchedNotificationsRequest) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *PatchedNotificationsRequest) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *PatchedNotificationsRequest) UnsetProduct() {
	o.Product.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedNotificationsRequest) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedNotificationsRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PatchedNotificationsRequest) SetUser(v int32) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *PatchedNotificationsRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PatchedNotificationsRequest) UnsetUser() {
	o.User.Unset()
}

// GetProductTypeAdded returns the ProductTypeAdded field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetProductTypeAdded() []string {
	if o == nil || IsNil(o.ProductTypeAdded) {
		var ret []string
		return ret
	}
	return o.ProductTypeAdded
}

// GetProductTypeAddedOk returns a tuple with the ProductTypeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetProductTypeAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypeAdded) {
		return nil, false
	}
	return o.ProductTypeAdded, true
}

// HasProductTypeAdded returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasProductTypeAdded() bool {
	if o != nil && !IsNil(o.ProductTypeAdded) {
		return true
	}

	return false
}

// SetProductTypeAdded gets a reference to the given []string and assigns it to the ProductTypeAdded field.
func (o *PatchedNotificationsRequest) SetProductTypeAdded(v []string) {
	o.ProductTypeAdded = v
}

// GetProductAdded returns the ProductAdded field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetProductAdded() []string {
	if o == nil || IsNil(o.ProductAdded) {
		var ret []string
		return ret
	}
	return o.ProductAdded
}

// GetProductAddedOk returns a tuple with the ProductAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetProductAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductAdded) {
		return nil, false
	}
	return o.ProductAdded, true
}

// HasProductAdded returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasProductAdded() bool {
	if o != nil && !IsNil(o.ProductAdded) {
		return true
	}

	return false
}

// SetProductAdded gets a reference to the given []string and assigns it to the ProductAdded field.
func (o *PatchedNotificationsRequest) SetProductAdded(v []string) {
	o.ProductAdded = v
}

// GetEngagementAdded returns the EngagementAdded field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetEngagementAdded() []string {
	if o == nil || IsNil(o.EngagementAdded) {
		var ret []string
		return ret
	}
	return o.EngagementAdded
}

// GetEngagementAddedOk returns a tuple with the EngagementAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetEngagementAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.EngagementAdded) {
		return nil, false
	}
	return o.EngagementAdded, true
}

// HasEngagementAdded returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasEngagementAdded() bool {
	if o != nil && !IsNil(o.EngagementAdded) {
		return true
	}

	return false
}

// SetEngagementAdded gets a reference to the given []string and assigns it to the EngagementAdded field.
func (o *PatchedNotificationsRequest) SetEngagementAdded(v []string) {
	o.EngagementAdded = v
}

// GetTestAdded returns the TestAdded field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetTestAdded() []string {
	if o == nil || IsNil(o.TestAdded) {
		var ret []string
		return ret
	}
	return o.TestAdded
}

// GetTestAddedOk returns a tuple with the TestAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetTestAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.TestAdded) {
		return nil, false
	}
	return o.TestAdded, true
}

// HasTestAdded returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasTestAdded() bool {
	if o != nil && !IsNil(o.TestAdded) {
		return true
	}

	return false
}

// SetTestAdded gets a reference to the given []string and assigns it to the TestAdded field.
func (o *PatchedNotificationsRequest) SetTestAdded(v []string) {
	o.TestAdded = v
}

// GetScanAdded returns the ScanAdded field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetScanAdded() []string {
	if o == nil || IsNil(o.ScanAdded) {
		var ret []string
		return ret
	}
	return o.ScanAdded
}

// GetScanAddedOk returns a tuple with the ScanAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetScanAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanAdded) {
		return nil, false
	}
	return o.ScanAdded, true
}

// HasScanAdded returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasScanAdded() bool {
	if o != nil && !IsNil(o.ScanAdded) {
		return true
	}

	return false
}

// SetScanAdded gets a reference to the given []string and assigns it to the ScanAdded field.
func (o *PatchedNotificationsRequest) SetScanAdded(v []string) {
	o.ScanAdded = v
}

// GetJiraUpdate returns the JiraUpdate field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetJiraUpdate() []string {
	if o == nil || IsNil(o.JiraUpdate) {
		var ret []string
		return ret
	}
	return o.JiraUpdate
}

// GetJiraUpdateOk returns a tuple with the JiraUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetJiraUpdateOk() ([]string, bool) {
	if o == nil || IsNil(o.JiraUpdate) {
		return nil, false
	}
	return o.JiraUpdate, true
}

// HasJiraUpdate returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasJiraUpdate() bool {
	if o != nil && !IsNil(o.JiraUpdate) {
		return true
	}

	return false
}

// SetJiraUpdate gets a reference to the given []string and assigns it to the JiraUpdate field.
func (o *PatchedNotificationsRequest) SetJiraUpdate(v []string) {
	o.JiraUpdate = v
}

// GetUpcomingEngagement returns the UpcomingEngagement field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetUpcomingEngagement() []string {
	if o == nil || IsNil(o.UpcomingEngagement) {
		var ret []string
		return ret
	}
	return o.UpcomingEngagement
}

// GetUpcomingEngagementOk returns a tuple with the UpcomingEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetUpcomingEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.UpcomingEngagement) {
		return nil, false
	}
	return o.UpcomingEngagement, true
}

// HasUpcomingEngagement returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasUpcomingEngagement() bool {
	if o != nil && !IsNil(o.UpcomingEngagement) {
		return true
	}

	return false
}

// SetUpcomingEngagement gets a reference to the given []string and assigns it to the UpcomingEngagement field.
func (o *PatchedNotificationsRequest) SetUpcomingEngagement(v []string) {
	o.UpcomingEngagement = v
}

// GetStaleEngagement returns the StaleEngagement field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetStaleEngagement() []string {
	if o == nil || IsNil(o.StaleEngagement) {
		var ret []string
		return ret
	}
	return o.StaleEngagement
}

// GetStaleEngagementOk returns a tuple with the StaleEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetStaleEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.StaleEngagement) {
		return nil, false
	}
	return o.StaleEngagement, true
}

// HasStaleEngagement returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasStaleEngagement() bool {
	if o != nil && !IsNil(o.StaleEngagement) {
		return true
	}

	return false
}

// SetStaleEngagement gets a reference to the given []string and assigns it to the StaleEngagement field.
func (o *PatchedNotificationsRequest) SetStaleEngagement(v []string) {
	o.StaleEngagement = v
}

// GetAutoCloseEngagement returns the AutoCloseEngagement field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetAutoCloseEngagement() []string {
	if o == nil || IsNil(o.AutoCloseEngagement) {
		var ret []string
		return ret
	}
	return o.AutoCloseEngagement
}

// GetAutoCloseEngagementOk returns a tuple with the AutoCloseEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetAutoCloseEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoCloseEngagement) {
		return nil, false
	}
	return o.AutoCloseEngagement, true
}

// HasAutoCloseEngagement returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasAutoCloseEngagement() bool {
	if o != nil && !IsNil(o.AutoCloseEngagement) {
		return true
	}

	return false
}

// SetAutoCloseEngagement gets a reference to the given []string and assigns it to the AutoCloseEngagement field.
func (o *PatchedNotificationsRequest) SetAutoCloseEngagement(v []string) {
	o.AutoCloseEngagement = v
}

// GetCloseEngagement returns the CloseEngagement field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetCloseEngagement() []string {
	if o == nil || IsNil(o.CloseEngagement) {
		var ret []string
		return ret
	}
	return o.CloseEngagement
}

// GetCloseEngagementOk returns a tuple with the CloseEngagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetCloseEngagementOk() ([]string, bool) {
	if o == nil || IsNil(o.CloseEngagement) {
		return nil, false
	}
	return o.CloseEngagement, true
}

// HasCloseEngagement returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasCloseEngagement() bool {
	if o != nil && !IsNil(o.CloseEngagement) {
		return true
	}

	return false
}

// SetCloseEngagement gets a reference to the given []string and assigns it to the CloseEngagement field.
func (o *PatchedNotificationsRequest) SetCloseEngagement(v []string) {
	o.CloseEngagement = v
}

// GetUserMentioned returns the UserMentioned field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetUserMentioned() []string {
	if o == nil || IsNil(o.UserMentioned) {
		var ret []string
		return ret
	}
	return o.UserMentioned
}

// GetUserMentionedOk returns a tuple with the UserMentioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetUserMentionedOk() ([]string, bool) {
	if o == nil || IsNil(o.UserMentioned) {
		return nil, false
	}
	return o.UserMentioned, true
}

// HasUserMentioned returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasUserMentioned() bool {
	if o != nil && !IsNil(o.UserMentioned) {
		return true
	}

	return false
}

// SetUserMentioned gets a reference to the given []string and assigns it to the UserMentioned field.
func (o *PatchedNotificationsRequest) SetUserMentioned(v []string) {
	o.UserMentioned = v
}

// GetCodeReview returns the CodeReview field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetCodeReview() []string {
	if o == nil || IsNil(o.CodeReview) {
		var ret []string
		return ret
	}
	return o.CodeReview
}

// GetCodeReviewOk returns a tuple with the CodeReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetCodeReviewOk() ([]string, bool) {
	if o == nil || IsNil(o.CodeReview) {
		return nil, false
	}
	return o.CodeReview, true
}

// HasCodeReview returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasCodeReview() bool {
	if o != nil && !IsNil(o.CodeReview) {
		return true
	}

	return false
}

// SetCodeReview gets a reference to the given []string and assigns it to the CodeReview field.
func (o *PatchedNotificationsRequest) SetCodeReview(v []string) {
	o.CodeReview = v
}

// GetReviewRequested returns the ReviewRequested field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetReviewRequested() []string {
	if o == nil || IsNil(o.ReviewRequested) {
		var ret []string
		return ret
	}
	return o.ReviewRequested
}

// GetReviewRequestedOk returns a tuple with the ReviewRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetReviewRequestedOk() ([]string, bool) {
	if o == nil || IsNil(o.ReviewRequested) {
		return nil, false
	}
	return o.ReviewRequested, true
}

// HasReviewRequested returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasReviewRequested() bool {
	if o != nil && !IsNil(o.ReviewRequested) {
		return true
	}

	return false
}

// SetReviewRequested gets a reference to the given []string and assigns it to the ReviewRequested field.
func (o *PatchedNotificationsRequest) SetReviewRequested(v []string) {
	o.ReviewRequested = v
}

// GetOther returns the Other field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetOther() []string {
	if o == nil || IsNil(o.Other) {
		var ret []string
		return ret
	}
	return o.Other
}

// GetOtherOk returns a tuple with the Other field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetOtherOk() ([]string, bool) {
	if o == nil || IsNil(o.Other) {
		return nil, false
	}
	return o.Other, true
}

// HasOther returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasOther() bool {
	if o != nil && !IsNil(o.Other) {
		return true
	}

	return false
}

// SetOther gets a reference to the given []string and assigns it to the Other field.
func (o *PatchedNotificationsRequest) SetOther(v []string) {
	o.Other = v
}

// GetSlaBreach returns the SlaBreach field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetSlaBreach() []string {
	if o == nil || IsNil(o.SlaBreach) {
		var ret []string
		return ret
	}
	return o.SlaBreach
}

// GetSlaBreachOk returns a tuple with the SlaBreach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetSlaBreachOk() ([]string, bool) {
	if o == nil || IsNil(o.SlaBreach) {
		return nil, false
	}
	return o.SlaBreach, true
}

// HasSlaBreach returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasSlaBreach() bool {
	if o != nil && !IsNil(o.SlaBreach) {
		return true
	}

	return false
}

// SetSlaBreach gets a reference to the given []string and assigns it to the SlaBreach field.
func (o *PatchedNotificationsRequest) SetSlaBreach(v []string) {
	o.SlaBreach = v
}

// GetSlaBreachCombined returns the SlaBreachCombined field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetSlaBreachCombined() []string {
	if o == nil || IsNil(o.SlaBreachCombined) {
		var ret []string
		return ret
	}
	return o.SlaBreachCombined
}

// GetSlaBreachCombinedOk returns a tuple with the SlaBreachCombined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetSlaBreachCombinedOk() ([]string, bool) {
	if o == nil || IsNil(o.SlaBreachCombined) {
		return nil, false
	}
	return o.SlaBreachCombined, true
}

// HasSlaBreachCombined returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasSlaBreachCombined() bool {
	if o != nil && !IsNil(o.SlaBreachCombined) {
		return true
	}

	return false
}

// SetSlaBreachCombined gets a reference to the given []string and assigns it to the SlaBreachCombined field.
func (o *PatchedNotificationsRequest) SetSlaBreachCombined(v []string) {
	o.SlaBreachCombined = v
}

// GetRiskAcceptanceExpiration returns the RiskAcceptanceExpiration field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetRiskAcceptanceExpiration() []string {
	if o == nil || IsNil(o.RiskAcceptanceExpiration) {
		var ret []string
		return ret
	}
	return o.RiskAcceptanceExpiration
}

// GetRiskAcceptanceExpirationOk returns a tuple with the RiskAcceptanceExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetRiskAcceptanceExpirationOk() ([]string, bool) {
	if o == nil || IsNil(o.RiskAcceptanceExpiration) {
		return nil, false
	}
	return o.RiskAcceptanceExpiration, true
}

// HasRiskAcceptanceExpiration returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasRiskAcceptanceExpiration() bool {
	if o != nil && !IsNil(o.RiskAcceptanceExpiration) {
		return true
	}

	return false
}

// SetRiskAcceptanceExpiration gets a reference to the given []string and assigns it to the RiskAcceptanceExpiration field.
func (o *PatchedNotificationsRequest) SetRiskAcceptanceExpiration(v []string) {
	o.RiskAcceptanceExpiration = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetTemplate() bool {
	if o == nil || IsNil(o.Template) {
		var ret bool
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetTemplateOk() (*bool, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given bool and assigns it to the Template field.
func (o *PatchedNotificationsRequest) SetTemplate(v bool) {
	o.Template = &v
}

// GetScanAddedEmpty returns the ScanAddedEmpty field value if set, zero value otherwise.
func (o *PatchedNotificationsRequest) GetScanAddedEmpty() string {
	if o == nil || IsNil(o.ScanAddedEmpty) {
		var ret string
		return ret
	}
	return *o.ScanAddedEmpty
}

// GetScanAddedEmptyOk returns a tuple with the ScanAddedEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNotificationsRequest) GetScanAddedEmptyOk() (*string, bool) {
	if o == nil || IsNil(o.ScanAddedEmpty) {
		return nil, false
	}
	return o.ScanAddedEmpty, true
}

// HasScanAddedEmpty returns a boolean if a field has been set.
func (o *PatchedNotificationsRequest) HasScanAddedEmpty() bool {
	if o != nil && !IsNil(o.ScanAddedEmpty) {
		return true
	}

	return false
}

// SetScanAddedEmpty gets a reference to the given string and assigns it to the ScanAddedEmpty field.
func (o *PatchedNotificationsRequest) SetScanAddedEmpty(v string) {
	o.ScanAddedEmpty = &v
}

func (o PatchedNotificationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedNotificationsRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SlaBreachCombined) {
		toSerialize["sla_breach_combined"] = o.SlaBreachCombined
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

type NullablePatchedNotificationsRequest struct {
	value *PatchedNotificationsRequest
	isSet bool
}

func (v NullablePatchedNotificationsRequest) Get() *PatchedNotificationsRequest {
	return v.value
}

func (v *NullablePatchedNotificationsRequest) Set(val *PatchedNotificationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNotificationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNotificationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNotificationsRequest(val *PatchedNotificationsRequest) *NullablePatchedNotificationsRequest {
	return &NullablePatchedNotificationsRequest{value: val, isSet: true}
}

func (v NullablePatchedNotificationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNotificationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


