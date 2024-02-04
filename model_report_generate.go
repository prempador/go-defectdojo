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

// checks if the ReportGenerate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportGenerate{}

// ReportGenerate struct for ReportGenerate
type ReportGenerate struct {
	ExecutiveSummary NullableExecutiveSummary `json:"executive_summary"`
	ProductType ProductType `json:"product_type"`
	Product Product `json:"product"`
	Engagement Engagement `json:"engagement"`
	ReportName string `json:"report_name"`
	ReportInfo string `json:"report_info"`
	Test Test `json:"test"`
	Endpoint Endpoint `json:"endpoint"`
	Endpoints []Endpoint `json:"endpoints"`
	Findings []Finding `json:"findings"`
	User UserStub `json:"user"`
	TeamName string `json:"team_name"`
	Title string `json:"title"`
	UserId int32 `json:"user_id"`
	Host string `json:"host"`
	FindingNotes []FindingToNotes `json:"finding_notes,omitempty"`
}

type _ReportGenerate ReportGenerate

// NewReportGenerate instantiates a new ReportGenerate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportGenerate(executiveSummary NullableExecutiveSummary, productType ProductType, product Product, engagement Engagement, reportName string, reportInfo string, test Test, endpoint Endpoint, endpoints []Endpoint, findings []Finding, user UserStub, teamName string, title string, userId int32, host string) *ReportGenerate {
	this := ReportGenerate{}
	this.ExecutiveSummary = executiveSummary
	this.ProductType = productType
	this.Product = product
	this.Engagement = engagement
	this.ReportName = reportName
	this.ReportInfo = reportInfo
	this.Test = test
	this.Endpoint = endpoint
	this.Endpoints = endpoints
	this.Findings = findings
	this.User = user
	this.TeamName = teamName
	this.Title = title
	this.UserId = userId
	this.Host = host
	return &this
}

// NewReportGenerateWithDefaults instantiates a new ReportGenerate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportGenerateWithDefaults() *ReportGenerate {
	this := ReportGenerate{}
	return &this
}

// GetExecutiveSummary returns the ExecutiveSummary field value
// If the value is explicit nil, the zero value for ExecutiveSummary will be returned
func (o *ReportGenerate) GetExecutiveSummary() ExecutiveSummary {
	if o == nil || o.ExecutiveSummary.Get() == nil {
		var ret ExecutiveSummary
		return ret
	}

	return *o.ExecutiveSummary.Get()
}

// GetExecutiveSummaryOk returns a tuple with the ExecutiveSummary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportGenerate) GetExecutiveSummaryOk() (*ExecutiveSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutiveSummary.Get(), o.ExecutiveSummary.IsSet()
}

// SetExecutiveSummary sets field value
func (o *ReportGenerate) SetExecutiveSummary(v ExecutiveSummary) {
	o.ExecutiveSummary.Set(&v)
}

// GetProductType returns the ProductType field value
func (o *ReportGenerate) GetProductType() ProductType {
	if o == nil {
		var ret ProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetProductTypeOk() (*ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ReportGenerate) SetProductType(v ProductType) {
	o.ProductType = v
}

// GetProduct returns the Product field value
func (o *ReportGenerate) GetProduct() Product {
	if o == nil {
		var ret Product
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetProductOk() (*Product, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *ReportGenerate) SetProduct(v Product) {
	o.Product = v
}

// GetEngagement returns the Engagement field value
func (o *ReportGenerate) GetEngagement() Engagement {
	if o == nil {
		var ret Engagement
		return ret
	}

	return o.Engagement
}

// GetEngagementOk returns a tuple with the Engagement field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetEngagementOk() (*Engagement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engagement, true
}

// SetEngagement sets field value
func (o *ReportGenerate) SetEngagement(v Engagement) {
	o.Engagement = v
}

// GetReportName returns the ReportName field value
func (o *ReportGenerate) GetReportName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetReportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportName, true
}

// SetReportName sets field value
func (o *ReportGenerate) SetReportName(v string) {
	o.ReportName = v
}

// GetReportInfo returns the ReportInfo field value
func (o *ReportGenerate) GetReportInfo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportInfo
}

// GetReportInfoOk returns a tuple with the ReportInfo field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetReportInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportInfo, true
}

// SetReportInfo sets field value
func (o *ReportGenerate) SetReportInfo(v string) {
	o.ReportInfo = v
}

// GetTest returns the Test field value
func (o *ReportGenerate) GetTest() Test {
	if o == nil {
		var ret Test
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetTestOk() (*Test, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *ReportGenerate) SetTest(v Test) {
	o.Test = v
}

// GetEndpoint returns the Endpoint field value
func (o *ReportGenerate) GetEndpoint() Endpoint {
	if o == nil {
		var ret Endpoint
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetEndpointOk() (*Endpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *ReportGenerate) SetEndpoint(v Endpoint) {
	o.Endpoint = v
}

// GetEndpoints returns the Endpoints field value
func (o *ReportGenerate) GetEndpoints() []Endpoint {
	if o == nil {
		var ret []Endpoint
		return ret
	}

	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetEndpointsOk() ([]Endpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// SetEndpoints sets field value
func (o *ReportGenerate) SetEndpoints(v []Endpoint) {
	o.Endpoints = v
}

// GetFindings returns the Findings field value
func (o *ReportGenerate) GetFindings() []Finding {
	if o == nil {
		var ret []Finding
		return ret
	}

	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetFindingsOk() ([]Finding, bool) {
	if o == nil {
		return nil, false
	}
	return o.Findings, true
}

// SetFindings sets field value
func (o *ReportGenerate) SetFindings(v []Finding) {
	o.Findings = v
}

// GetUser returns the User field value
func (o *ReportGenerate) GetUser() UserStub {
	if o == nil {
		var ret UserStub
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetUserOk() (*UserStub, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ReportGenerate) SetUser(v UserStub) {
	o.User = v
}

// GetTeamName returns the TeamName field value
func (o *ReportGenerate) GetTeamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamName, true
}

// SetTeamName sets field value
func (o *ReportGenerate) SetTeamName(v string) {
	o.TeamName = v
}

// GetTitle returns the Title field value
func (o *ReportGenerate) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ReportGenerate) SetTitle(v string) {
	o.Title = v
}

// GetUserId returns the UserId field value
func (o *ReportGenerate) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ReportGenerate) SetUserId(v int32) {
	o.UserId = v
}

// GetHost returns the Host field value
func (o *ReportGenerate) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *ReportGenerate) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *ReportGenerate) SetHost(v string) {
	o.Host = v
}

// GetFindingNotes returns the FindingNotes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportGenerate) GetFindingNotes() []FindingToNotes {
	if o == nil {
		var ret []FindingToNotes
		return ret
	}
	return o.FindingNotes
}

// GetFindingNotesOk returns a tuple with the FindingNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportGenerate) GetFindingNotesOk() ([]FindingToNotes, bool) {
	if o == nil || IsNil(o.FindingNotes) {
		return nil, false
	}
	return o.FindingNotes, true
}

// HasFindingNotes returns a boolean if a field has been set.
func (o *ReportGenerate) HasFindingNotes() bool {
	if o != nil && IsNil(o.FindingNotes) {
		return true
	}

	return false
}

// SetFindingNotes gets a reference to the given []FindingToNotes and assigns it to the FindingNotes field.
func (o *ReportGenerate) SetFindingNotes(v []FindingToNotes) {
	o.FindingNotes = v
}

func (o ReportGenerate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportGenerate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["executive_summary"] = o.ExecutiveSummary.Get()
	toSerialize["product_type"] = o.ProductType
	toSerialize["product"] = o.Product
	toSerialize["engagement"] = o.Engagement
	toSerialize["report_name"] = o.ReportName
	toSerialize["report_info"] = o.ReportInfo
	toSerialize["test"] = o.Test
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["endpoints"] = o.Endpoints
	toSerialize["findings"] = o.Findings
	toSerialize["user"] = o.User
	toSerialize["team_name"] = o.TeamName
	toSerialize["title"] = o.Title
	toSerialize["user_id"] = o.UserId
	toSerialize["host"] = o.Host
	if o.FindingNotes != nil {
		toSerialize["finding_notes"] = o.FindingNotes
	}
	return toSerialize, nil
}

func (o *ReportGenerate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"executive_summary",
		"product_type",
		"product",
		"engagement",
		"report_name",
		"report_info",
		"test",
		"endpoint",
		"endpoints",
		"findings",
		"user",
		"team_name",
		"title",
		"user_id",
		"host",
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

	varReportGenerate := _ReportGenerate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportGenerate)

	if err != nil {
		return err
	}

	*o = ReportGenerate(varReportGenerate)

	return err
}

type NullableReportGenerate struct {
	value *ReportGenerate
	isSet bool
}

func (v NullableReportGenerate) Get() *ReportGenerate {
	return v.value
}

func (v *NullableReportGenerate) Set(val *ReportGenerate) {
	v.value = val
	v.isSet = true
}

func (v NullableReportGenerate) IsSet() bool {
	return v.isSet
}

func (v *NullableReportGenerate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportGenerate(val *ReportGenerate) *NullableReportGenerate {
	return &NullableReportGenerate{value: val, isSet: true}
}

func (v NullableReportGenerate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportGenerate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

