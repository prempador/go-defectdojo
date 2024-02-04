/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the JIRAProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JIRAProject{}

// JIRAProject struct for JIRAProject
type JIRAProject struct {
	Id int32 `json:"id"`
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

type _JIRAProject JIRAProject

// NewJIRAProject instantiates a new JIRAProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJIRAProject(id int32) *JIRAProject {
	this := JIRAProject{}
	this.Id = id
	return &this
}

// NewJIRAProjectWithDefaults instantiates a new JIRAProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJIRAProjectWithDefaults() *JIRAProject {
	this := JIRAProject{}
	return &this
}

// GetId returns the Id field value
func (o *JIRAProject) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JIRAProject) SetId(v int32) {
	o.Id = v
}

// GetProjectKey returns the ProjectKey field value if set, zero value otherwise.
func (o *JIRAProject) GetProjectKey() string {
	if o == nil || IsNil(o.ProjectKey) {
		var ret string
		return ret
	}
	return *o.ProjectKey
}

// GetProjectKeyOk returns a tuple with the ProjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetProjectKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectKey) {
		return nil, false
	}
	return o.ProjectKey, true
}

// HasProjectKey returns a boolean if a field has been set.
func (o *JIRAProject) HasProjectKey() bool {
	if o != nil && !IsNil(o.ProjectKey) {
		return true
	}

	return false
}

// SetProjectKey gets a reference to the given string and assigns it to the ProjectKey field.
func (o *JIRAProject) SetProjectKey(v string) {
	o.ProjectKey = &v
}

// GetIssueTemplateDir returns the IssueTemplateDir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetIssueTemplateDir() string {
	if o == nil || IsNil(o.IssueTemplateDir.Get()) {
		var ret string
		return ret
	}
	return *o.IssueTemplateDir.Get()
}

// GetIssueTemplateDirOk returns a tuple with the IssueTemplateDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetIssueTemplateDirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssueTemplateDir.Get(), o.IssueTemplateDir.IsSet()
}

// HasIssueTemplateDir returns a boolean if a field has been set.
func (o *JIRAProject) HasIssueTemplateDir() bool {
	if o != nil && o.IssueTemplateDir.IsSet() {
		return true
	}

	return false
}

// SetIssueTemplateDir gets a reference to the given NullableString and assigns it to the IssueTemplateDir field.
func (o *JIRAProject) SetIssueTemplateDir(v string) {
	o.IssueTemplateDir.Set(&v)
}
// SetIssueTemplateDirNil sets the value for IssueTemplateDir to be an explicit nil
func (o *JIRAProject) SetIssueTemplateDirNil() {
	o.IssueTemplateDir.Set(nil)
}

// UnsetIssueTemplateDir ensures that no value is present for IssueTemplateDir, not even an explicit nil
func (o *JIRAProject) UnsetIssueTemplateDir() {
	o.IssueTemplateDir.Unset()
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *JIRAProject) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *JIRAProject) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *JIRAProject) SetComponent(v string) {
	o.Component = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetCustomFields() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetCustomFieldsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return &o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *JIRAProject) HasCustomFields() bool {
	if o != nil && IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given interface{} and assigns it to the CustomFields field.
func (o *JIRAProject) SetCustomFields(v interface{}) {
	o.CustomFields = v
}

// GetDefaultAssignee returns the DefaultAssignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetDefaultAssignee() string {
	if o == nil || IsNil(o.DefaultAssignee.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultAssignee.Get()
}

// GetDefaultAssigneeOk returns a tuple with the DefaultAssignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetDefaultAssigneeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultAssignee.Get(), o.DefaultAssignee.IsSet()
}

// HasDefaultAssignee returns a boolean if a field has been set.
func (o *JIRAProject) HasDefaultAssignee() bool {
	if o != nil && o.DefaultAssignee.IsSet() {
		return true
	}

	return false
}

// SetDefaultAssignee gets a reference to the given NullableString and assigns it to the DefaultAssignee field.
func (o *JIRAProject) SetDefaultAssignee(v string) {
	o.DefaultAssignee.Set(&v)
}
// SetDefaultAssigneeNil sets the value for DefaultAssignee to be an explicit nil
func (o *JIRAProject) SetDefaultAssigneeNil() {
	o.DefaultAssignee.Set(nil)
}

// UnsetDefaultAssignee ensures that no value is present for DefaultAssignee, not even an explicit nil
func (o *JIRAProject) UnsetDefaultAssignee() {
	o.DefaultAssignee.Unset()
}

// GetJiraLabels returns the JiraLabels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetJiraLabels() string {
	if o == nil || IsNil(o.JiraLabels.Get()) {
		var ret string
		return ret
	}
	return *o.JiraLabels.Get()
}

// GetJiraLabelsOk returns a tuple with the JiraLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetJiraLabelsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraLabels.Get(), o.JiraLabels.IsSet()
}

// HasJiraLabels returns a boolean if a field has been set.
func (o *JIRAProject) HasJiraLabels() bool {
	if o != nil && o.JiraLabels.IsSet() {
		return true
	}

	return false
}

// SetJiraLabels gets a reference to the given NullableString and assigns it to the JiraLabels field.
func (o *JIRAProject) SetJiraLabels(v string) {
	o.JiraLabels.Set(&v)
}
// SetJiraLabelsNil sets the value for JiraLabels to be an explicit nil
func (o *JIRAProject) SetJiraLabelsNil() {
	o.JiraLabels.Set(nil)
}

// UnsetJiraLabels ensures that no value is present for JiraLabels, not even an explicit nil
func (o *JIRAProject) UnsetJiraLabels() {
	o.JiraLabels.Unset()
}

// GetAddVulnerabilityIdToJiraLabel returns the AddVulnerabilityIdToJiraLabel field value if set, zero value otherwise.
func (o *JIRAProject) GetAddVulnerabilityIdToJiraLabel() bool {
	if o == nil || IsNil(o.AddVulnerabilityIdToJiraLabel) {
		var ret bool
		return ret
	}
	return *o.AddVulnerabilityIdToJiraLabel
}

// GetAddVulnerabilityIdToJiraLabelOk returns a tuple with the AddVulnerabilityIdToJiraLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetAddVulnerabilityIdToJiraLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.AddVulnerabilityIdToJiraLabel) {
		return nil, false
	}
	return o.AddVulnerabilityIdToJiraLabel, true
}

// HasAddVulnerabilityIdToJiraLabel returns a boolean if a field has been set.
func (o *JIRAProject) HasAddVulnerabilityIdToJiraLabel() bool {
	if o != nil && !IsNil(o.AddVulnerabilityIdToJiraLabel) {
		return true
	}

	return false
}

// SetAddVulnerabilityIdToJiraLabel gets a reference to the given bool and assigns it to the AddVulnerabilityIdToJiraLabel field.
func (o *JIRAProject) SetAddVulnerabilityIdToJiraLabel(v bool) {
	o.AddVulnerabilityIdToJiraLabel = &v
}

// GetPushAllIssues returns the PushAllIssues field value if set, zero value otherwise.
func (o *JIRAProject) GetPushAllIssues() bool {
	if o == nil || IsNil(o.PushAllIssues) {
		var ret bool
		return ret
	}
	return *o.PushAllIssues
}

// GetPushAllIssuesOk returns a tuple with the PushAllIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetPushAllIssuesOk() (*bool, bool) {
	if o == nil || IsNil(o.PushAllIssues) {
		return nil, false
	}
	return o.PushAllIssues, true
}

// HasPushAllIssues returns a boolean if a field has been set.
func (o *JIRAProject) HasPushAllIssues() bool {
	if o != nil && !IsNil(o.PushAllIssues) {
		return true
	}

	return false
}

// SetPushAllIssues gets a reference to the given bool and assigns it to the PushAllIssues field.
func (o *JIRAProject) SetPushAllIssues(v bool) {
	o.PushAllIssues = &v
}

// GetEnableEngagementEpicMapping returns the EnableEngagementEpicMapping field value if set, zero value otherwise.
func (o *JIRAProject) GetEnableEngagementEpicMapping() bool {
	if o == nil || IsNil(o.EnableEngagementEpicMapping) {
		var ret bool
		return ret
	}
	return *o.EnableEngagementEpicMapping
}

// GetEnableEngagementEpicMappingOk returns a tuple with the EnableEngagementEpicMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetEnableEngagementEpicMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEngagementEpicMapping) {
		return nil, false
	}
	return o.EnableEngagementEpicMapping, true
}

// HasEnableEngagementEpicMapping returns a boolean if a field has been set.
func (o *JIRAProject) HasEnableEngagementEpicMapping() bool {
	if o != nil && !IsNil(o.EnableEngagementEpicMapping) {
		return true
	}

	return false
}

// SetEnableEngagementEpicMapping gets a reference to the given bool and assigns it to the EnableEngagementEpicMapping field.
func (o *JIRAProject) SetEnableEngagementEpicMapping(v bool) {
	o.EnableEngagementEpicMapping = &v
}

// GetPushNotes returns the PushNotes field value if set, zero value otherwise.
func (o *JIRAProject) GetPushNotes() bool {
	if o == nil || IsNil(o.PushNotes) {
		var ret bool
		return ret
	}
	return *o.PushNotes
}

// GetPushNotesOk returns a tuple with the PushNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetPushNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.PushNotes) {
		return nil, false
	}
	return o.PushNotes, true
}

// HasPushNotes returns a boolean if a field has been set.
func (o *JIRAProject) HasPushNotes() bool {
	if o != nil && !IsNil(o.PushNotes) {
		return true
	}

	return false
}

// SetPushNotes gets a reference to the given bool and assigns it to the PushNotes field.
func (o *JIRAProject) SetPushNotes(v bool) {
	o.PushNotes = &v
}

// GetProductJiraSlaNotification returns the ProductJiraSlaNotification field value if set, zero value otherwise.
func (o *JIRAProject) GetProductJiraSlaNotification() bool {
	if o == nil || IsNil(o.ProductJiraSlaNotification) {
		var ret bool
		return ret
	}
	return *o.ProductJiraSlaNotification
}

// GetProductJiraSlaNotificationOk returns a tuple with the ProductJiraSlaNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetProductJiraSlaNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.ProductJiraSlaNotification) {
		return nil, false
	}
	return o.ProductJiraSlaNotification, true
}

// HasProductJiraSlaNotification returns a boolean if a field has been set.
func (o *JIRAProject) HasProductJiraSlaNotification() bool {
	if o != nil && !IsNil(o.ProductJiraSlaNotification) {
		return true
	}

	return false
}

// SetProductJiraSlaNotification gets a reference to the given bool and assigns it to the ProductJiraSlaNotification field.
func (o *JIRAProject) SetProductJiraSlaNotification(v bool) {
	o.ProductJiraSlaNotification = &v
}

// GetRiskAcceptanceExpirationNotification returns the RiskAcceptanceExpirationNotification field value if set, zero value otherwise.
func (o *JIRAProject) GetRiskAcceptanceExpirationNotification() bool {
	if o == nil || IsNil(o.RiskAcceptanceExpirationNotification) {
		var ret bool
		return ret
	}
	return *o.RiskAcceptanceExpirationNotification
}

// GetRiskAcceptanceExpirationNotificationOk returns a tuple with the RiskAcceptanceExpirationNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JIRAProject) GetRiskAcceptanceExpirationNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RiskAcceptanceExpirationNotification) {
		return nil, false
	}
	return o.RiskAcceptanceExpirationNotification, true
}

// HasRiskAcceptanceExpirationNotification returns a boolean if a field has been set.
func (o *JIRAProject) HasRiskAcceptanceExpirationNotification() bool {
	if o != nil && !IsNil(o.RiskAcceptanceExpirationNotification) {
		return true
	}

	return false
}

// SetRiskAcceptanceExpirationNotification gets a reference to the given bool and assigns it to the RiskAcceptanceExpirationNotification field.
func (o *JIRAProject) SetRiskAcceptanceExpirationNotification(v bool) {
	o.RiskAcceptanceExpirationNotification = &v
}

// GetJiraInstance returns the JiraInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetJiraInstance() int32 {
	if o == nil || IsNil(o.JiraInstance.Get()) {
		var ret int32
		return ret
	}
	return *o.JiraInstance.Get()
}

// GetJiraInstanceOk returns a tuple with the JiraInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetJiraInstanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.JiraInstance.Get(), o.JiraInstance.IsSet()
}

// HasJiraInstance returns a boolean if a field has been set.
func (o *JIRAProject) HasJiraInstance() bool {
	if o != nil && o.JiraInstance.IsSet() {
		return true
	}

	return false
}

// SetJiraInstance gets a reference to the given NullableInt32 and assigns it to the JiraInstance field.
func (o *JIRAProject) SetJiraInstance(v int32) {
	o.JiraInstance.Set(&v)
}
// SetJiraInstanceNil sets the value for JiraInstance to be an explicit nil
func (o *JIRAProject) SetJiraInstanceNil() {
	o.JiraInstance.Set(nil)
}

// UnsetJiraInstance ensures that no value is present for JiraInstance, not even an explicit nil
func (o *JIRAProject) UnsetJiraInstance() {
	o.JiraInstance.Unset()
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *JIRAProject) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *JIRAProject) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *JIRAProject) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *JIRAProject) UnsetProduct() {
	o.Product.Unset()
}

// GetEngagement returns the Engagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JIRAProject) GetEngagement() int32 {
	if o == nil || IsNil(o.Engagement.Get()) {
		var ret int32
		return ret
	}
	return *o.Engagement.Get()
}

// GetEngagementOk returns a tuple with the Engagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JIRAProject) GetEngagementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engagement.Get(), o.Engagement.IsSet()
}

// HasEngagement returns a boolean if a field has been set.
func (o *JIRAProject) HasEngagement() bool {
	if o != nil && o.Engagement.IsSet() {
		return true
	}

	return false
}

// SetEngagement gets a reference to the given NullableInt32 and assigns it to the Engagement field.
func (o *JIRAProject) SetEngagement(v int32) {
	o.Engagement.Set(&v)
}
// SetEngagementNil sets the value for Engagement to be an explicit nil
func (o *JIRAProject) SetEngagementNil() {
	o.Engagement.Set(nil)
}

// UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
func (o *JIRAProject) UnsetEngagement() {
	o.Engagement.Unset()
}

func (o JIRAProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JIRAProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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

func (o *JIRAProject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varJIRAProject := _JIRAProject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJIRAProject)

	if err != nil {
		return err
	}

	*o = JIRAProject(varJIRAProject)

	return err
}

type NullableJIRAProject struct {
	value *JIRAProject
	isSet bool
}

func (v NullableJIRAProject) Get() *JIRAProject {
	return v.value
}

func (v *NullableJIRAProject) Set(val *JIRAProject) {
	v.value = val
	v.isSet = true
}

func (v NullableJIRAProject) IsSet() bool {
	return v.isSet
}

func (v *NullableJIRAProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJIRAProject(val *JIRAProject) *NullableJIRAProject {
	return &NullableJIRAProject{value: val, isSet: true}
}

func (v NullableJIRAProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJIRAProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


