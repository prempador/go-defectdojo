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

// checks if the PatchedCredentialMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedCredentialMappingRequest{}

// PatchedCredentialMappingRequest struct for PatchedCredentialMappingRequest
type PatchedCredentialMappingRequest struct {
	IsAuthnProvider *bool `json:"is_authn_provider,omitempty"`
	Url NullableString `json:"url,omitempty"`
	CredId *int32 `json:"cred_id,omitempty"`
	Product NullableInt32 `json:"product,omitempty"`
	Finding NullableInt32 `json:"finding,omitempty"`
	Engagement NullableInt32 `json:"engagement,omitempty"`
	Test NullableInt32 `json:"test,omitempty"`
}

// NewPatchedCredentialMappingRequest instantiates a new PatchedCredentialMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedCredentialMappingRequest() *PatchedCredentialMappingRequest {
	this := PatchedCredentialMappingRequest{}
	return &this
}

// NewPatchedCredentialMappingRequestWithDefaults instantiates a new PatchedCredentialMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedCredentialMappingRequestWithDefaults() *PatchedCredentialMappingRequest {
	this := PatchedCredentialMappingRequest{}
	return &this
}

// GetIsAuthnProvider returns the IsAuthnProvider field value if set, zero value otherwise.
func (o *PatchedCredentialMappingRequest) GetIsAuthnProvider() bool {
	if o == nil || IsNil(o.IsAuthnProvider) {
		var ret bool
		return ret
	}
	return *o.IsAuthnProvider
}

// GetIsAuthnProviderOk returns a tuple with the IsAuthnProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialMappingRequest) GetIsAuthnProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthnProvider) {
		return nil, false
	}
	return o.IsAuthnProvider, true
}

// HasIsAuthnProvider returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasIsAuthnProvider() bool {
	if o != nil && !IsNil(o.IsAuthnProvider) {
		return true
	}

	return false
}

// SetIsAuthnProvider gets a reference to the given bool and assigns it to the IsAuthnProvider field.
func (o *PatchedCredentialMappingRequest) SetIsAuthnProvider(v bool) {
	o.IsAuthnProvider = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialMappingRequest) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialMappingRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PatchedCredentialMappingRequest) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PatchedCredentialMappingRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PatchedCredentialMappingRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetCredId returns the CredId field value if set, zero value otherwise.
func (o *PatchedCredentialMappingRequest) GetCredId() int32 {
	if o == nil || IsNil(o.CredId) {
		var ret int32
		return ret
	}
	return *o.CredId
}

// GetCredIdOk returns a tuple with the CredId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedCredentialMappingRequest) GetCredIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CredId) {
		return nil, false
	}
	return o.CredId, true
}

// HasCredId returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasCredId() bool {
	if o != nil && !IsNil(o.CredId) {
		return true
	}

	return false
}

// SetCredId gets a reference to the given int32 and assigns it to the CredId field.
func (o *PatchedCredentialMappingRequest) SetCredId(v int32) {
	o.CredId = &v
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialMappingRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialMappingRequest) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *PatchedCredentialMappingRequest) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *PatchedCredentialMappingRequest) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *PatchedCredentialMappingRequest) UnsetProduct() {
	o.Product.Unset()
}

// GetFinding returns the Finding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialMappingRequest) GetFinding() int32 {
	if o == nil || IsNil(o.Finding.Get()) {
		var ret int32
		return ret
	}
	return *o.Finding.Get()
}

// GetFindingOk returns a tuple with the Finding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialMappingRequest) GetFindingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Finding.Get(), o.Finding.IsSet()
}

// HasFinding returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasFinding() bool {
	if o != nil && o.Finding.IsSet() {
		return true
	}

	return false
}

// SetFinding gets a reference to the given NullableInt32 and assigns it to the Finding field.
func (o *PatchedCredentialMappingRequest) SetFinding(v int32) {
	o.Finding.Set(&v)
}
// SetFindingNil sets the value for Finding to be an explicit nil
func (o *PatchedCredentialMappingRequest) SetFindingNil() {
	o.Finding.Set(nil)
}

// UnsetFinding ensures that no value is present for Finding, not even an explicit nil
func (o *PatchedCredentialMappingRequest) UnsetFinding() {
	o.Finding.Unset()
}

// GetEngagement returns the Engagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialMappingRequest) GetEngagement() int32 {
	if o == nil || IsNil(o.Engagement.Get()) {
		var ret int32
		return ret
	}
	return *o.Engagement.Get()
}

// GetEngagementOk returns a tuple with the Engagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialMappingRequest) GetEngagementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engagement.Get(), o.Engagement.IsSet()
}

// HasEngagement returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasEngagement() bool {
	if o != nil && o.Engagement.IsSet() {
		return true
	}

	return false
}

// SetEngagement gets a reference to the given NullableInt32 and assigns it to the Engagement field.
func (o *PatchedCredentialMappingRequest) SetEngagement(v int32) {
	o.Engagement.Set(&v)
}
// SetEngagementNil sets the value for Engagement to be an explicit nil
func (o *PatchedCredentialMappingRequest) SetEngagementNil() {
	o.Engagement.Set(nil)
}

// UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
func (o *PatchedCredentialMappingRequest) UnsetEngagement() {
	o.Engagement.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedCredentialMappingRequest) GetTest() int32 {
	if o == nil || IsNil(o.Test.Get()) {
		var ret int32
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedCredentialMappingRequest) GetTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *PatchedCredentialMappingRequest) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableInt32 and assigns it to the Test field.
func (o *PatchedCredentialMappingRequest) SetTest(v int32) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *PatchedCredentialMappingRequest) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *PatchedCredentialMappingRequest) UnsetTest() {
	o.Test.Unset()
}

func (o PatchedCredentialMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedCredentialMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsAuthnProvider) {
		toSerialize["is_authn_provider"] = o.IsAuthnProvider
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.CredId) {
		toSerialize["cred_id"] = o.CredId
	}
	if o.Product.IsSet() {
		toSerialize["product"] = o.Product.Get()
	}
	if o.Finding.IsSet() {
		toSerialize["finding"] = o.Finding.Get()
	}
	if o.Engagement.IsSet() {
		toSerialize["engagement"] = o.Engagement.Get()
	}
	if o.Test.IsSet() {
		toSerialize["test"] = o.Test.Get()
	}
	return toSerialize, nil
}

type NullablePatchedCredentialMappingRequest struct {
	value *PatchedCredentialMappingRequest
	isSet bool
}

func (v NullablePatchedCredentialMappingRequest) Get() *PatchedCredentialMappingRequest {
	return v.value
}

func (v *NullablePatchedCredentialMappingRequest) Set(val *PatchedCredentialMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedCredentialMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedCredentialMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedCredentialMappingRequest(val *PatchedCredentialMappingRequest) *NullablePatchedCredentialMappingRequest {
	return &NullablePatchedCredentialMappingRequest{value: val, isSet: true}
}

func (v NullablePatchedCredentialMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedCredentialMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


