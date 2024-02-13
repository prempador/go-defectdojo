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
)

// checks if the PatchedRiskAcceptanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRiskAcceptanceRequest{}

// PatchedRiskAcceptanceRequest struct for PatchedRiskAcceptanceRequest
type PatchedRiskAcceptanceRequest struct {
	// Descriptive name which in the future may also be used to group risk acceptances together across engagements and products
	Name *string `json:"name,omitempty"`
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
	Owner *int32 `json:"owner,omitempty"`
	AcceptedFindings []int32 `json:"accepted_findings,omitempty"`
}

// NewPatchedRiskAcceptanceRequest instantiates a new PatchedRiskAcceptanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRiskAcceptanceRequest() *PatchedRiskAcceptanceRequest {
	this := PatchedRiskAcceptanceRequest{}
	return &this
}

// NewPatchedRiskAcceptanceRequestWithDefaults instantiates a new PatchedRiskAcceptanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRiskAcceptanceRequestWithDefaults() *PatchedRiskAcceptanceRequest {
	this := PatchedRiskAcceptanceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRiskAcceptanceRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRiskAcceptanceRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRiskAcceptanceRequest) SetName(v string) {
	o.Name = &v
}

// GetRecommendationDetails returns the RecommendationDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRiskAcceptanceRequest) GetRecommendationDetails() string {
	if o == nil || IsNil(o.RecommendationDetails.Get()) {
		var ret string
		return ret
	}
	return *o.RecommendationDetails.Get()
}

// GetRecommendationDetailsOk returns a tuple with the RecommendationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRiskAcceptanceRequest) GetRecommendationDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecommendationDetails.Get(), o.RecommendationDetails.IsSet()
}

// HasRecommendationDetails returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasRecommendationDetails() bool {
	if o != nil && o.RecommendationDetails.IsSet() {
		return true
	}

	return false
}

// SetRecommendationDetails gets a reference to the given NullableString and assigns it to the RecommendationDetails field.
func (o *PatchedRiskAcceptanceRequest) SetRecommendationDetails(v string) {
	o.RecommendationDetails.Set(&v)
}
// SetRecommendationDetailsNil sets the value for RecommendationDetails to be an explicit nil
func (o *PatchedRiskAcceptanceRequest) SetRecommendationDetailsNil() {
	o.RecommendationDetails.Set(nil)
}

// UnsetRecommendationDetails ensures that no value is present for RecommendationDetails, not even an explicit nil
func (o *PatchedRiskAcceptanceRequest) UnsetRecommendationDetails() {
	o.RecommendationDetails.Unset()
}

// GetDecisionDetails returns the DecisionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRiskAcceptanceRequest) GetDecisionDetails() string {
	if o == nil || IsNil(o.DecisionDetails.Get()) {
		var ret string
		return ret
	}
	return *o.DecisionDetails.Get()
}

// GetDecisionDetailsOk returns a tuple with the DecisionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRiskAcceptanceRequest) GetDecisionDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DecisionDetails.Get(), o.DecisionDetails.IsSet()
}

// HasDecisionDetails returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasDecisionDetails() bool {
	if o != nil && o.DecisionDetails.IsSet() {
		return true
	}

	return false
}

// SetDecisionDetails gets a reference to the given NullableString and assigns it to the DecisionDetails field.
func (o *PatchedRiskAcceptanceRequest) SetDecisionDetails(v string) {
	o.DecisionDetails.Set(&v)
}
// SetDecisionDetailsNil sets the value for DecisionDetails to be an explicit nil
func (o *PatchedRiskAcceptanceRequest) SetDecisionDetailsNil() {
	o.DecisionDetails.Set(nil)
}

// UnsetDecisionDetails ensures that no value is present for DecisionDetails, not even an explicit nil
func (o *PatchedRiskAcceptanceRequest) UnsetDecisionDetails() {
	o.DecisionDetails.Unset()
}

// GetAcceptedBy returns the AcceptedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRiskAcceptanceRequest) GetAcceptedBy() string {
	if o == nil || IsNil(o.AcceptedBy.Get()) {
		var ret string
		return ret
	}
	return *o.AcceptedBy.Get()
}

// GetAcceptedByOk returns a tuple with the AcceptedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRiskAcceptanceRequest) GetAcceptedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AcceptedBy.Get(), o.AcceptedBy.IsSet()
}

// HasAcceptedBy returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasAcceptedBy() bool {
	if o != nil && o.AcceptedBy.IsSet() {
		return true
	}

	return false
}

// SetAcceptedBy gets a reference to the given NullableString and assigns it to the AcceptedBy field.
func (o *PatchedRiskAcceptanceRequest) SetAcceptedBy(v string) {
	o.AcceptedBy.Set(&v)
}
// SetAcceptedByNil sets the value for AcceptedBy to be an explicit nil
func (o *PatchedRiskAcceptanceRequest) SetAcceptedByNil() {
	o.AcceptedBy.Set(nil)
}

// UnsetAcceptedBy ensures that no value is present for AcceptedBy, not even an explicit nil
func (o *PatchedRiskAcceptanceRequest) UnsetAcceptedBy() {
	o.AcceptedBy.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRiskAcceptanceRequest) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRiskAcceptanceRequest) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *PatchedRiskAcceptanceRequest) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *PatchedRiskAcceptanceRequest) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *PatchedRiskAcceptanceRequest) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetExpirationDateWarned returns the ExpirationDateWarned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRiskAcceptanceRequest) GetExpirationDateWarned() time.Time {
	if o == nil || IsNil(o.ExpirationDateWarned.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateWarned.Get()
}

// GetExpirationDateWarnedOk returns a tuple with the ExpirationDateWarned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRiskAcceptanceRequest) GetExpirationDateWarnedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDateWarned.Get(), o.ExpirationDateWarned.IsSet()
}

// HasExpirationDateWarned returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasExpirationDateWarned() bool {
	if o != nil && o.ExpirationDateWarned.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateWarned gets a reference to the given NullableTime and assigns it to the ExpirationDateWarned field.
func (o *PatchedRiskAcceptanceRequest) SetExpirationDateWarned(v time.Time) {
	o.ExpirationDateWarned.Set(&v)
}
// SetExpirationDateWarnedNil sets the value for ExpirationDateWarned to be an explicit nil
func (o *PatchedRiskAcceptanceRequest) SetExpirationDateWarnedNil() {
	o.ExpirationDateWarned.Set(nil)
}

// UnsetExpirationDateWarned ensures that no value is present for ExpirationDateWarned, not even an explicit nil
func (o *PatchedRiskAcceptanceRequest) UnsetExpirationDateWarned() {
	o.ExpirationDateWarned.Unset()
}

// GetExpirationDateHandled returns the ExpirationDateHandled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRiskAcceptanceRequest) GetExpirationDateHandled() time.Time {
	if o == nil || IsNil(o.ExpirationDateHandled.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateHandled.Get()
}

// GetExpirationDateHandledOk returns a tuple with the ExpirationDateHandled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRiskAcceptanceRequest) GetExpirationDateHandledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDateHandled.Get(), o.ExpirationDateHandled.IsSet()
}

// HasExpirationDateHandled returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasExpirationDateHandled() bool {
	if o != nil && o.ExpirationDateHandled.IsSet() {
		return true
	}

	return false
}

// SetExpirationDateHandled gets a reference to the given NullableTime and assigns it to the ExpirationDateHandled field.
func (o *PatchedRiskAcceptanceRequest) SetExpirationDateHandled(v time.Time) {
	o.ExpirationDateHandled.Set(&v)
}
// SetExpirationDateHandledNil sets the value for ExpirationDateHandled to be an explicit nil
func (o *PatchedRiskAcceptanceRequest) SetExpirationDateHandledNil() {
	o.ExpirationDateHandled.Set(nil)
}

// UnsetExpirationDateHandled ensures that no value is present for ExpirationDateHandled, not even an explicit nil
func (o *PatchedRiskAcceptanceRequest) UnsetExpirationDateHandled() {
	o.ExpirationDateHandled.Unset()
}

// GetReactivateExpired returns the ReactivateExpired field value if set, zero value otherwise.
func (o *PatchedRiskAcceptanceRequest) GetReactivateExpired() bool {
	if o == nil || IsNil(o.ReactivateExpired) {
		var ret bool
		return ret
	}
	return *o.ReactivateExpired
}

// GetReactivateExpiredOk returns a tuple with the ReactivateExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRiskAcceptanceRequest) GetReactivateExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ReactivateExpired) {
		return nil, false
	}
	return o.ReactivateExpired, true
}

// HasReactivateExpired returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasReactivateExpired() bool {
	if o != nil && !IsNil(o.ReactivateExpired) {
		return true
	}

	return false
}

// SetReactivateExpired gets a reference to the given bool and assigns it to the ReactivateExpired field.
func (o *PatchedRiskAcceptanceRequest) SetReactivateExpired(v bool) {
	o.ReactivateExpired = &v
}

// GetRestartSlaExpired returns the RestartSlaExpired field value if set, zero value otherwise.
func (o *PatchedRiskAcceptanceRequest) GetRestartSlaExpired() bool {
	if o == nil || IsNil(o.RestartSlaExpired) {
		var ret bool
		return ret
	}
	return *o.RestartSlaExpired
}

// GetRestartSlaExpiredOk returns a tuple with the RestartSlaExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRiskAcceptanceRequest) GetRestartSlaExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RestartSlaExpired) {
		return nil, false
	}
	return o.RestartSlaExpired, true
}

// HasRestartSlaExpired returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasRestartSlaExpired() bool {
	if o != nil && !IsNil(o.RestartSlaExpired) {
		return true
	}

	return false
}

// SetRestartSlaExpired gets a reference to the given bool and assigns it to the RestartSlaExpired field.
func (o *PatchedRiskAcceptanceRequest) SetRestartSlaExpired(v bool) {
	o.RestartSlaExpired = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PatchedRiskAcceptanceRequest) GetOwner() int32 {
	if o == nil || IsNil(o.Owner) {
		var ret int32
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRiskAcceptanceRequest) GetOwnerOk() (*int32, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given int32 and assigns it to the Owner field.
func (o *PatchedRiskAcceptanceRequest) SetOwner(v int32) {
	o.Owner = &v
}

// GetAcceptedFindings returns the AcceptedFindings field value if set, zero value otherwise.
func (o *PatchedRiskAcceptanceRequest) GetAcceptedFindings() []int32 {
	if o == nil || IsNil(o.AcceptedFindings) {
		var ret []int32
		return ret
	}
	return o.AcceptedFindings
}

// GetAcceptedFindingsOk returns a tuple with the AcceptedFindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRiskAcceptanceRequest) GetAcceptedFindingsOk() ([]int32, bool) {
	if o == nil || IsNil(o.AcceptedFindings) {
		return nil, false
	}
	return o.AcceptedFindings, true
}

// HasAcceptedFindings returns a boolean if a field has been set.
func (o *PatchedRiskAcceptanceRequest) HasAcceptedFindings() bool {
	if o != nil && !IsNil(o.AcceptedFindings) {
		return true
	}

	return false
}

// SetAcceptedFindings gets a reference to the given []int32 and assigns it to the AcceptedFindings field.
func (o *PatchedRiskAcceptanceRequest) SetAcceptedFindings(v []int32) {
	o.AcceptedFindings = v
}

func (o PatchedRiskAcceptanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRiskAcceptanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
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
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.AcceptedFindings) {
		toSerialize["accepted_findings"] = o.AcceptedFindings
	}
	return toSerialize, nil
}

type NullablePatchedRiskAcceptanceRequest struct {
	value *PatchedRiskAcceptanceRequest
	isSet bool
}

func (v NullablePatchedRiskAcceptanceRequest) Get() *PatchedRiskAcceptanceRequest {
	return v.value
}

func (v *NullablePatchedRiskAcceptanceRequest) Set(val *PatchedRiskAcceptanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRiskAcceptanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRiskAcceptanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRiskAcceptanceRequest(val *PatchedRiskAcceptanceRequest) *NullablePatchedRiskAcceptanceRequest {
	return &NullablePatchedRiskAcceptanceRequest{value: val, isSet: true}
}

func (v NullablePatchedRiskAcceptanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRiskAcceptanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


