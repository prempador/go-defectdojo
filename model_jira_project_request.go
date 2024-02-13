/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the JIRAProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JIRAProjectRequest{}

// JIRAProjectRequest struct for JIRAProjectRequest
type JIRAProjectRequest struct {
	ProjectKey *string `json:"project_key,omitempty"`
	// Choose the folder containing the Django templates used to render the JIRA issue description. These are stored in dojo/templates/issue-trackers. Leave empty to use the default jira_full templates.
	IssueTemplateDir NullableString `json:"issue_template_dir,omitempty"`
	Component *string `json:"component,omitempty"`
	// JIRA custom field JSON mapping of Id to value, e.g. {\"customfield_10122\": [{\"name\": \"8.0.1\"}]}
	CustomFields interface{} `json:"custom_fields,omitempty"`
	// JIRA default assignee (name). If left blank then it defaults to whatever is configured in JIRA.
	DefaultAssignee NullableString `json:"default_assignee,omitempty"`
	// JIRA issue labels space seperated
	JiraLabels NullableString `json:"jira_labels,omitempty"`
	AddVulnerabilityIdToJiraLabel *bool `json:"add_vulnerability_id_to_jira_label,omitempty"`
	// Automatically maintain parity with JIRA. Always create and update JIRA tickets for findings in this Product.
	PushAllIssues *bool `json:"push_all_issues,omitempty"`
	EnableEngagementEpicMapping *bool `json:"enable_engagement_epic_mapping,omitempty"`
	PushNotes *bool `json:"push_notes,omitempty"`
	ProductJiraSlaNotification *bool `json:"product_jira_sla_notification,omitempty"`
	RiskAcceptanceExpirationNotification *bool `json:"risk_acceptance_expiration_notification,omitempty"`
	JiraInstance NullableInt32 `json:"jira_instance,omitempty"`
	Product NullableInt32 `json:"product,omitempty"`
	Engagement NullableInt32 `json:"engagement,omitempty"`
}

// NewJIRAProjectRequest instantiates a new JIRAProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJIRAProjectRequest() *JIRAProjectRequest {
	this := JIRAProjectRequest{}
	return &this
}

// NewJIRAProjectRequestWithDefaults instantiates a new JIRAProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJIRAProjectRequestWithDefaults() *JIRAProjectRequest {
	this := JIRAProjectRequest{}
	return &this
}

// GetProjectKey returns the ProjectKey field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetProjectKey() string {
	if o == nil || IsNil(o.ProjectKey) {
		var ret string
		return ret
	}
	return *o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetProjectKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectKey) {
		return nil, false
	}
	return o.ProjectKey, true
}

// HasProjectKey returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasProjectKey() bool {
	if o != nil && !IsNil(o.ProjectKey) {
		return true
	}

	return false
}

// SetProjectKey gets a reference to the given string and assigns it to the ProjectKey field.
func (o *JIRAProjectRequest) SetProjectKey(v string) {
	o.ProjectKey = &v
}

// GetIssueTemplateDir returns the IssueTemplateDir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetIssueTemplateDir() string {
	if o == nil || IsNil(o.IssueTemplateDir.Get()) {
		var ret string
		return ret
	}
	return *o.IssueTemplateDir.Get()
}

// GetIssueTemplateDirOk returns a tuple with the IssueTemplateDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetIssueTemplateDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssueTemplateDir.Get(), o.IssueTemplateDir.IsSet()
}

// HasIssueTemplateDir returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasIssueTemplateDir() bool {
	if o != nil && o.IssueTemplateDir.IsSet() {
		return true
	}

	return false
}

// SetIssueTemplateDir gets a reference to the given NullableString and assigns it to the IssueTemplateDir field.
func (o *JIRAProjectRequest) SetIssueTemplateDir(v string) {
	o.IssueTemplateDir.Set(&v)
}
// SetIssueTemplateDirNil sets the value for IssueTemplateDir to be an explicit nil
func (o *JIRAProjectRequest) SetIssueTemplateDirNil() {
	o.IssueTemplateDir.Set(nil)
}

// UnsetIssueTemplateDir ensures that no value is present for IssueTemplateDir, not even an explicit nil
func (o *JIRAProjectRequest) UnsetIssueTemplateDir() {
	o.IssueTemplateDir.Unset()
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *JIRAProjectRequest) SetComponent(v string) {
	o.Component = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetCustomFields() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetCustomFieldsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return &o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasCustomFields() bool {
	if o != nil && IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given interface{} and assigns it to the CustomFields field.
func (o *JIRAProjectRequest) SetCustomFields(v interface{}) {
	o.CustomFields = v
}

// GetDefaultAssignee returns the DefaultAssignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetDefaultAssignee() string {
	if o == nil || IsNil(o.DefaultAssignee.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultAssignee.Get()
}

// GetDefaultAssigneeOk returns a tuple with the DefaultAssignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetDefaultAssigneeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultAssignee.Get(), o.DefaultAssignee.IsSet()
}

// HasDefaultAssignee returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasDefaultAssignee() bool {
	if o != nil && o.DefaultAssignee.IsSet() {
		return true
	}

	return false
}

// SetDefaultAssignee gets a reference to the given NullableString and assigns it to the DefaultAssignee field.
func (o *JIRAProjectRequest) SetDefaultAssignee(v string) {
	o.DefaultAssignee.Set(&v)
}
// SetDefaultAssigneeNil sets the value for DefaultAssignee to be an explicit nil
func (o *JIRAProjectRequest) SetDefaultAssigneeNil() {
	o.DefaultAssignee.Set(nil)
}

// UnsetDefaultAssignee ensures that no value is present for DefaultAssignee, not even an explicit nil
func (o *JIRAProjectRequest) UnsetDefaultAssignee() {
	o.DefaultAssignee.Unset()
}

// GetJiraLabels returns the JiraLabels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetJiraLabels() string {
	if o == nil || IsNil(o.JiraLabels.Get()) {
		var ret string
		return ret
	}
	return *o.JiraLabels.Get()
}

// GetJiraLabelsOk returns a tuple with the JiraLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetJiraLabelsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraLabels.Get(), o.JiraLabels.IsSet()
}

// HasJiraLabels returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasJiraLabels() bool {
	if o != nil && o.JiraLabels.IsSet() {
		return true
	}

	return false
}

// SetJiraLabels gets a reference to the given NullableString and assigns it to the JiraLabels field.
func (o *JIRAProjectRequest) SetJiraLabels(v string) {
	o.JiraLabels.Set(&v)
}
// SetJiraLabelsNil sets the value for JiraLabels to be an explicit nil
func (o *JIRAProjectRequest) SetJiraLabelsNil() {
	o.JiraLabels.Set(nil)
}

// UnsetJiraLabels ensures that no value is present for JiraLabels, not even an explicit nil
func (o *JIRAProjectRequest) UnsetJiraLabels() {
	o.JiraLabels.Unset()
}

// GetAddVulnerabilityIdToJiraLabel returns the AddVulnerabilityIdToJiraLabel field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetAddVulnerabilityIdToJiraLabel() bool {
	if o == nil || IsNil(o.AddVulnerabilityIdToJiraLabel) {
		var ret bool
		return ret
	}
	return *o.AddVulnerabilityIdToJiraLabel
}

// GetAddVulnerabilityIdToJiraLabelOk returns a tuple with the AddVulnerabilityIdToJiraLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetAddVulnerabilityIdToJiraLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.AddVulnerabilityIdToJiraLabel) {
		return nil, false
	}
	return o.AddVulnerabilityIdToJiraLabel, true
}

// HasAddVulnerabilityIdToJiraLabel returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasAddVulnerabilityIdToJiraLabel() bool {
	if o != nil && !IsNil(o.AddVulnerabilityIdToJiraLabel) {
		return true
	}

	return false
}

// SetAddVulnerabilityIdToJiraLabel gets a reference to the given bool and assigns it to the AddVulnerabilityIdToJiraLabel field.
func (o *JIRAProjectRequest) SetAddVulnerabilityIdToJiraLabel(v bool) {
	o.AddVulnerabilityIdToJiraLabel = &v
}

// GetPushAllIssues returns the PushAllIssues field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetPushAllIssues() bool {
	if o == nil || IsNil(o.PushAllIssues) {
		var ret bool
		return ret
	}
	return *o.PushAllIssues
}

// GetPushAllIssuesOk returns a tuple with the PushAllIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetPushAllIssuesOk() (*bool, bool) {
	if o == nil || IsNil(o.PushAllIssues) {
		return nil, false
	}
	return o.PushAllIssues, true
}

// HasPushAllIssues returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasPushAllIssues() bool {
	if o != nil && !IsNil(o.PushAllIssues) {
		return true
	}

	return false
}

// SetPushAllIssues gets a reference to the given bool and assigns it to the PushAllIssues field.
func (o *JIRAProjectRequest) SetPushAllIssues(v bool) {
	o.PushAllIssues = &v
}

// GetEnableEngagementEpicMapping returns the EnableEngagementEpicMapping field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetEnableEngagementEpicMapping() bool {
	if o == nil || IsNil(o.EnableEngagementEpicMapping) {
		var ret bool
		return ret
	}
	return *o.EnableEngagementEpicMapping
}

// GetEnableEngagementEpicMappingOk returns a tuple with the EnableEngagementEpicMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetEnableEngagementEpicMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEngagementEpicMapping) {
		return nil, false
	}
	return o.EnableEngagementEpicMapping, true
}

// HasEnableEngagementEpicMapping returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasEnableEngagementEpicMapping() bool {
	if o != nil && !IsNil(o.EnableEngagementEpicMapping) {
		return true
	}

	return false
}

// SetEnableEngagementEpicMapping gets a reference to the given bool and assigns it to the EnableEngagementEpicMapping field.
func (o *JIRAProjectRequest) SetEnableEngagementEpicMapping(v bool) {
	o.EnableEngagementEpicMapping = &v
}

// GetPushNotes returns the PushNotes field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetPushNotes() bool {
	if o == nil || IsNil(o.PushNotes) {
		var ret bool
		return ret
	}
	return *o.PushNotes
}

// GetPushNotesOk returns a tuple with the PushNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetPushNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.PushNotes) {
		return nil, false
	}
	return o.PushNotes, true
}

// HasPushNotes returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasPushNotes() bool {
	if o != nil && !IsNil(o.PushNotes) {
		return true
	}

	return false
}

// SetPushNotes gets a reference to the given bool and assigns it to the PushNotes field.
func (o *JIRAProjectRequest) SetPushNotes(v bool) {
	o.PushNotes = &v
}

// GetProductJiraSlaNotification returns the ProductJiraSlaNotification field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetProductJiraSlaNotification() bool {
	if o == nil || IsNil(o.ProductJiraSlaNotification) {
		var ret bool
		return ret
	}
	return *o.ProductJiraSlaNotification
}

// GetProductJiraSlaNotificationOk returns a tuple with the ProductJiraSlaNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetProductJiraSlaNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductJiraSlaNotification) {
		return nil, false
	}
	return o.ProductJiraSlaNotification, true
}

// HasProductJiraSlaNotification returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasProductJiraSlaNotification() bool {
	if o != nil && !IsNil(o.ProductJiraSlaNotification) {
		return true
	}

	return false
}

// SetProductJiraSlaNotification gets a reference to the given bool and assigns it to the ProductJiraSlaNotification field.
func (o *JIRAProjectRequest) SetProductJiraSlaNotification(v bool) {
	o.ProductJiraSlaNotification = &v
}

// GetRiskAcceptanceExpirationNotification returns the RiskAcceptanceExpirationNotification field value if set, zero value otherwise.
func (o *JIRAProjectRequest) GetRiskAcceptanceExpirationNotification() bool {
	if o == nil || IsNil(o.RiskAcceptanceExpirationNotification) {
		var ret bool
		return ret
	}
	return *o.RiskAcceptanceExpirationNotification
}

// GetRiskAcceptanceExpirationNotificationOk returns a tuple with the RiskAcceptanceExpirationNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProjectRequest) GetRiskAcceptanceExpirationNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RiskAcceptanceExpirationNotification) {
		return nil, false
	}
	return o.RiskAcceptanceExpirationNotification, true
}

// HasRiskAcceptanceExpirationNotification returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasRiskAcceptanceExpirationNotification() bool {
	if o != nil && !IsNil(o.RiskAcceptanceExpirationNotification) {
		return true
	}

	return false
}

// SetRiskAcceptanceExpirationNotification gets a reference to the given bool and assigns it to the RiskAcceptanceExpirationNotification field.
func (o *JIRAProjectRequest) SetRiskAcceptanceExpirationNotification(v bool) {
	o.RiskAcceptanceExpirationNotification = &v
}

// GetJiraInstance returns the JiraInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetJiraInstance() int32 {
	if o == nil || IsNil(o.JiraInstance.Get()) {
		var ret int32
		return ret
	}
	return *o.JiraInstance.Get()
}

// GetJiraInstanceOk returns a tuple with the JiraInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetJiraInstanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraInstance.Get(), o.JiraInstance.IsSet()
}

// HasJiraInstance returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasJiraInstance() bool {
	if o != nil && o.JiraInstance.IsSet() {
		return true
	}

	return false
}

// SetJiraInstance gets a reference to the given NullableInt32 and assigns it to the JiraInstance field.
func (o *JIRAProjectRequest) SetJiraInstance(v int32) {
	o.JiraInstance.Set(&v)
}
// SetJiraInstanceNil sets the value for JiraInstance to be an explicit nil
func (o *JIRAProjectRequest) SetJiraInstanceNil() {
	o.JiraInstance.Set(nil)
}

// UnsetJiraInstance ensures that no value is present for JiraInstance, not even an explicit nil
func (o *JIRAProjectRequest) UnsetJiraInstance() {
	o.JiraInstance.Unset()
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *JIRAProjectRequest) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *JIRAProjectRequest) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *JIRAProjectRequest) UnsetProduct() {
	o.Product.Unset()
}

// GetEngagement returns the Engagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProjectRequest) GetEngagement() int32 {
	if o == nil || IsNil(o.Engagement.Get()) {
		var ret int32
		return ret
	}
	return *o.Engagement.Get()
}

// GetEngagementOk returns a tuple with the Engagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProjectRequest) GetEngagementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engagement.Get(), o.Engagement.IsSet()
}

// HasEngagement returns a boolean if a field has been set.
func (o *JIRAProjectRequest) HasEngagement() bool {
	if o != nil && o.Engagement.IsSet() {
		return true
	}

	return false
}

// SetEngagement gets a reference to the given NullableInt32 and assigns it to the Engagement field.
func (o *JIRAProjectRequest) SetEngagement(v int32) {
	o.Engagement.Set(&v)
}
// SetEngagementNil sets the value for Engagement to be an explicit nil
func (o *JIRAProjectRequest) SetEngagementNil() {
	o.Engagement.Set(nil)
}

// UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
func (o *JIRAProjectRequest) UnsetEngagement() {
	o.Engagement.Unset()
}

func (o JIRAProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JIRAProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectKey) {
		toSerialize["project_key"] = o.ProjectKey
	}
	if o.IssueTemplateDir.IsSet() {
		toSerialize["issue_template_dir"] = o.IssueTemplateDir.Get()
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.DefaultAssignee.IsSet() {
		toSerialize["default_assignee"] = o.DefaultAssignee.Get()
	}
	if o.JiraLabels.IsSet() {
		toSerialize["jira_labels"] = o.JiraLabels.Get()
	}
	if !IsNil(o.AddVulnerabilityIdToJiraLabel) {
		toSerialize["add_vulnerability_id_to_jira_label"] = o.AddVulnerabilityIdToJiraLabel
	}
	if !IsNil(o.PushAllIssues) {
		toSerialize["push_all_issues"] = o.PushAllIssues
	}
	if !IsNil(o.EnableEngagementEpicMapping) {
		toSerialize["enable_engagement_epic_mapping"] = o.EnableEngagementEpicMapping
	}
	if !IsNil(o.PushNotes) {
		toSerialize["push_notes"] = o.PushNotes
	}
	if !IsNil(o.ProductJiraSlaNotification) {
		toSerialize["product_jira_sla_notification"] = o.ProductJiraSlaNotification
	}
	if !IsNil(o.RiskAcceptanceExpirationNotification) {
		toSerialize["risk_acceptance_expiration_notification"] = o.RiskAcceptanceExpirationNotification
	}
	if o.JiraInstance.IsSet() {
		toSerialize["jira_instance"] = o.JiraInstance.Get()
	}
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	if o.Engagement.IsSet() {
		toSerialize["engagement"] = o.Engagement.Get()
	}
	return toSerialize, nil
}

type NullableJIRAProjectRequest struct {
	value *JIRAProjectRequest
	isSet bool
}

func (v NullableJIRAProjectRequest) Get() *JIRAProjectRequest {
	return v.value
}

func (v *NullableJIRAProjectRequest) Set(val *JIRAProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJIRAProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJIRAProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJIRAProjectRequest(val *JIRAProjectRequest) *NullableJIRAProjectRequest {
	return &NullableJIRAProjectRequest{value: val, isSet: true}
}

func (v NullableJIRAProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJIRAProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


