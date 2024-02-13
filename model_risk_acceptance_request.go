/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RiskAcceptanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskAcceptanceRequest{}

// RiskAcceptanceRequest struct for RiskAcceptanceRequest
type RiskAcceptanceRequest struct {
	// Descriptive name which in the future may also be used to group risk acceptances together across engagements and products
	Name string `json:"name"`
	// Explanation of security recommendation
	RecommendationDetails NullableString `json:"recommendation_details,omitempty"`
	// If a compensating control exists to mitigate the finding or reduce risk, then list the compensating control(s).
	DecisionDetails NullableString `json:"decision_details,omitempty"`
	// The person that accepts the risk, can be outside of DefectDojo.
	AcceptedBy NullableString `json:"accepted_by,omitempty"`
	// When the risk acceptance expires, the findings will be reactivated (unless disabled below).
	ExpirationDate NullableTime `json:"expiration_date,omitempty"`
	// (readonly) Date at which notice about the risk acceptance expiration was sent.
	ExpirationDateWarned NullableTime `json:"expiration_date_warned,omitempty"`
	// (readonly) When the risk acceptance expiration was handled (manually or by the daily job).
	ExpirationDateHandled NullableTime `json:"expiration_date_handled,omitempty"`
	// Reactivate findings when risk acceptance expires?
	ReactivateExpired *bool `json:"reactivate_expired,omitempty"`
	// When enabled, the SLA for findings is restarted when the risk acceptance expires.
	RestartSlaExpired *bool `json:"restart_sla_expired,omitempty"`
	// User in DefectDojo owning this acceptance. Only the owner and staff users can edit the risk acceptance.
	Owner int32 `json:"owner"`
	AcceptedFindings []int32 `json:"accepted_findings"`
}

type _RiskAcceptanceRequest RiskAcceptanceRequest

// NewRiskAcceptanceRequest instantiates a new RiskAcceptanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAcceptanceRequest(name string, owner int32, acceptedFindings []int32) *RiskAcceptanceRequest {
	this := RiskAcceptanceRequest{}
	this.Name = name
	this.Owner = owner
	this.AcceptedFindings = acceptedFindings
	return &this
}

// NewRiskAcceptanceRequestWithDefaults instantiates a new RiskAcceptanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAcceptanceRequestWithDefaults() *RiskAcceptanceRequest {
	this := RiskAcceptanceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *RiskAcceptanceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskAcceptanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskAcceptanceRequest) SetName(v string) {
	o.Name = v
}

// GetRecommendationDetails returns the RecommendationDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskAcceptanceRequest) GetRecommendationDetails() string {
	if o == nil || IsNil(o.RecommendationDetails.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationDetails.Get()
}

// GetRecommendationDetailsOk returns a tuple with the RecommendationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskAcceptanceRequest) GetRecommendationDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationDetails.Get(), o.RecommendationDetails.IsSet()
}

// HasRecommendationDetails returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasRecommendationDetails() bool {
	if o != nil && o.RecommendationDetails.IsSet() {
		return true
	}

	return false
}

// SetRecommendationDetails gets a reference to the given NullableString and assigns it to the RecommendationDetails field.
func (o *RiskAcceptanceRequest) SetRecommendationDetails(v string) {
	o.RecommendationDetails.Set(&v)
}
// SetRecommendationDetailsNil sets the value for RecommendationDetails to be an explicit nil
func (o *RiskAcceptanceRequest) SetRecommendationDetailsNil() {
	o.RecommendationDetails.Set(nil)
}

// UnsetRecommendationDetails ensures that no value is present for RecommendationDetails, not even an explicit nil
func (o *RiskAcceptanceRequest) UnsetRecommendationDetails() {
	o.RecommendationDetails.Unset()
}

// GetDecisionDetails returns the DecisionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskAcceptanceRequest) GetDecisionDetails() string {
	if o == nil || IsNil(o.DecisionDetails.Get()) {
		var ret string
		return ret
	}
	return *o.DecisionDetails.Get()
}

// GetDecisionDetailsOk returns a tuple with the DecisionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskAcceptanceRequest) GetDecisionDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionDetails.Get(), o.DecisionDetails.IsSet()
}

// HasDecisionDetails returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasDecisionDetails() bool {
	if o != nil && o.DecisionDetails.IsSet() {
		return true
	}

	return false
}

// SetDecisionDetails gets a reference to the given NullableString and assigns it to the DecisionDetails field.
func (o *RiskAcceptanceRequest) SetDecisionDetails(v string) {
	o.DecisionDetails.Set(&v)
}
// SetDecisionDetailsNil sets the value for DecisionDetails to be an explicit nil
func (o *RiskAcceptanceRequest) SetDecisionDetailsNil() {
	o.DecisionDetails.Set(nil)
}

// UnsetDecisionDetails ensures that no value is present for DecisionDetails, not even an explicit nil
func (o *RiskAcceptanceRequest) UnsetDecisionDetails() {
	o.DecisionDetails.Unset()
}

// GetAcceptedBy returns the AcceptedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskAcceptanceRequest) GetAcceptedBy() string {
	if o == nil || IsNil(o.AcceptedBy.Get()) {
		var ret string
		return ret
	}
	return *o.AcceptedBy.Get()
}

// GetAcceptedByOk returns a tuple with the AcceptedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskAcceptanceRequest) GetAcceptedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedBy.Get(), o.AcceptedBy.IsSet()
}

// HasAcceptedBy returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasAcceptedBy() bool {
	if o != nil && o.AcceptedBy.IsSet() {
		return true
	}

	return false
}

// SetAcceptedBy gets a reference to the given NullableString and assigns it to the AcceptedBy field.
func (o *RiskAcceptanceRequest) SetAcceptedBy(v string) {
	o.AcceptedBy.Set(&v)
}
// SetAcceptedByNil sets the value for AcceptedBy to be an explicit nil
func (o *RiskAcceptanceRequest) SetAcceptedByNil() {
	o.AcceptedBy.Set(nil)
}

// UnsetAcceptedBy ensures that no value is present for AcceptedBy, not even an explicit nil
func (o *RiskAcceptanceRequest) UnsetAcceptedBy() {
	o.AcceptedBy.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskAcceptanceRequest) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskAcceptanceRequest) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *RiskAcceptanceRequest) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *RiskAcceptanceRequest) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *RiskAcceptanceRequest) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetExpirationDateWarned returns the ExpirationDateWarned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskAcceptanceRequest) GetExpirationDateWarned() time.Time {
	if o == nil || IsNil(o.ExpirationDateWarned.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateWarned.Get()
}

// GetExpirationDateWarnedOk returns a tuple with the ExpirationDateWarned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskAcceptanceRequest) GetExpirationDateWarnedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDateWarned.Get(), o.ExpirationDateWarned.IsSet()
}

// HasExpirationDateWarned returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasExpirationDateWarned() bool {
	if o != nil && o.ExpirationDateWarned.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateWarned gets a reference to the given NullableTime and assigns it to the ExpirationDateWarned field.
func (o *RiskAcceptanceRequest) SetExpirationDateWarned(v time.Time) {
	o.ExpirationDateWarned.Set(&v)
}
// SetExpirationDateWarnedNil sets the value for ExpirationDateWarned to be an explicit nil
func (o *RiskAcceptanceRequest) SetExpirationDateWarnedNil() {
	o.ExpirationDateWarned.Set(nil)
}

// UnsetExpirationDateWarned ensures that no value is present for ExpirationDateWarned, not even an explicit nil
func (o *RiskAcceptanceRequest) UnsetExpirationDateWarned() {
	o.ExpirationDateWarned.Unset()
}

// GetExpirationDateHandled returns the ExpirationDateHandled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RiskAcceptanceRequest) GetExpirationDateHandled() time.Time {
	if o == nil || IsNil(o.ExpirationDateHandled.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateHandled.Get()
}

// GetExpirationDateHandledOk returns a tuple with the ExpirationDateHandled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RiskAcceptanceRequest) GetExpirationDateHandledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDateHandled.Get(), o.ExpirationDateHandled.IsSet()
}

// HasExpirationDateHandled returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasExpirationDateHandled() bool {
	if o != nil && o.ExpirationDateHandled.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateHandled gets a reference to the given NullableTime and assigns it to the ExpirationDateHandled field.
func (o *RiskAcceptanceRequest) SetExpirationDateHandled(v time.Time) {
	o.ExpirationDateHandled.Set(&v)
}
// SetExpirationDateHandledNil sets the value for ExpirationDateHandled to be an explicit nil
func (o *RiskAcceptanceRequest) SetExpirationDateHandledNil() {
	o.ExpirationDateHandled.Set(nil)
}

// UnsetExpirationDateHandled ensures that no value is present for ExpirationDateHandled, not even an explicit nil
func (o *RiskAcceptanceRequest) UnsetExpirationDateHandled() {
	o.ExpirationDateHandled.Unset()
}

// GetReactivateExpired returns the ReactivateExpired field value if set, zero value otherwise.
func (o *RiskAcceptanceRequest) GetReactivateExpired() bool {
	if o == nil || IsNil(o.ReactivateExpired) {
		var ret bool
		return ret
	}
	return *o.ReactivateExpired
}

// GetReactivateExpiredOk returns a tuple with the ReactivateExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAcceptanceRequest) GetReactivateExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReactivateExpired) {
		return nil, false
	}
	return o.ReactivateExpired, true
}

// HasReactivateExpired returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasReactivateExpired() bool {
	if o != nil && !IsNil(o.ReactivateExpired) {
		return true
	}

	return false
}

// SetReactivateExpired gets a reference to the given bool and assigns it to the ReactivateExpired field.
func (o *RiskAcceptanceRequest) SetReactivateExpired(v bool) {
	o.ReactivateExpired = &v
}

// GetRestartSlaExpired returns the RestartSlaExpired field value if set, zero value otherwise.
func (o *RiskAcceptanceRequest) GetRestartSlaExpired() bool {
	if o == nil || IsNil(o.RestartSlaExpired) {
		var ret bool
		return ret
	}
	return *o.RestartSlaExpired
}

// GetRestartSlaExpiredOk returns a tuple with the RestartSlaExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskAcceptanceRequest) GetRestartSlaExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RestartSlaExpired) {
		return nil, false
	}
	return o.RestartSlaExpired, true
}

// HasRestartSlaExpired returns a boolean if a field has been set.
func (o *RiskAcceptanceRequest) HasRestartSlaExpired() bool {
	if o != nil && !IsNil(o.RestartSlaExpired) {
		return true
	}

	return false
}

// SetRestartSlaExpired gets a reference to the given bool and assigns it to the RestartSlaExpired field.
func (o *RiskAcceptanceRequest) SetRestartSlaExpired(v bool) {
	o.RestartSlaExpired = &v
}

// GetOwner returns the Owner field value
func (o *RiskAcceptanceRequest) GetOwner() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *RiskAcceptanceRequest) GetOwnerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *RiskAcceptanceRequest) SetOwner(v int32) {
	o.Owner = v
}

// GetAcceptedFindings returns the AcceptedFindings field value
func (o *RiskAcceptanceRequest) GetAcceptedFindings() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AcceptedFindings
}

// GetAcceptedFindingsOk returns a tuple with the AcceptedFindings field value
// and a boolean to check if the value has been set.
func (o *RiskAcceptanceRequest) GetAcceptedFindingsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedFindings, true
}

// SetAcceptedFindings sets field value
func (o *RiskAcceptanceRequest) SetAcceptedFindings(v []int32) {
	o.AcceptedFindings = v
}

func (o RiskAcceptanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskAcceptanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.RecommendationDetails.IsSet() {
		toSerialize["recommendation_details"] = o.RecommendationDetails.Get()
	}
	if o.DecisionDetails.IsSet() {
		toSerialize["decision_details"] = o.DecisionDetails.Get()
	}
	if o.AcceptedBy.IsSet() {
		toSerialize["accepted_by"] = o.AcceptedBy.Get()
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if o.ExpirationDateWarned.IsSet() {
		toSerialize["expiration_date_warned"] = o.ExpirationDateWarned.Get()
	}
	if o.ExpirationDateHandled.IsSet() {
		toSerialize["expiration_date_handled"] = o.ExpirationDateHandled.Get()
	}
	if !IsNil(o.ReactivateExpired) {
		toSerialize["reactivate_expired"] = o.ReactivateExpired
	}
	if !IsNil(o.RestartSlaExpired) {
		toSerialize["restart_sla_expired"] = o.RestartSlaExpired
	}
	toSerialize["owner"] = o.Owner
	toSerialize["accepted_findings"] = o.AcceptedFindings
	return toSerialize, nil
}

func (o *RiskAcceptanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"owner",
		"accepted_findings",
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

	varRiskAcceptanceRequest := _RiskAcceptanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskAcceptanceRequest)

	if err != nil {
		return err
	}

	*o = RiskAcceptanceRequest(varRiskAcceptanceRequest)

	return err
}

type NullableRiskAcceptanceRequest struct {
	value *RiskAcceptanceRequest
	isSet bool
}

func (v NullableRiskAcceptanceRequest) Get() *RiskAcceptanceRequest {
	return v.value
}

func (v *NullableRiskAcceptanceRequest) Set(val *RiskAcceptanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAcceptanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAcceptanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAcceptanceRequest(val *RiskAcceptanceRequest) *NullableRiskAcceptanceRequest {
	return &NullableRiskAcceptanceRequest{value: val, isSet: true}
}

func (v NullableRiskAcceptanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAcceptanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


