/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CredentialMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialMapping{}

// CredentialMapping struct for CredentialMapping
type CredentialMapping struct {
	Id int32 `json:"id"`
	IsAuthnProvider *bool `json:"is_authn_provider,omitempty"`
	Url NullableString `json:"url,omitempty"`
	CredId int32 `json:"cred_id"`
	Product NullableInt32 `json:"product,omitempty"`
	Finding NullableInt32 `json:"finding,omitempty"`
	Engagement NullableInt32 `json:"engagement,omitempty"`
	Test NullableInt32 `json:"test,omitempty"`
}

type _CredentialMapping CredentialMapping

// NewCredentialMapping instantiates a new CredentialMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialMapping(id int32, credId int32) *CredentialMapping {
	this := CredentialMapping{}
	this.Id = id
	this.CredId = credId
	return &this
}

// NewCredentialMappingWithDefaults instantiates a new CredentialMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialMappingWithDefaults() *CredentialMapping {
	this := CredentialMapping{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialMapping) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialMapping) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CredentialMapping) SetId(v int32) {
	o.Id = v
}

// GetIsAuthnProvider returns the IsAuthnProvider field value if set, zero value otherwise.
func (o *CredentialMapping) GetIsAuthnProvider() bool {
	if o == nil || IsNil(o.IsAuthnProvider) {
		var ret bool
		return ret
	}
	return *o.IsAuthnProvider
}

// GetIsAuthnProviderOk returns a tuple with the IsAuthnProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialMapping) GetIsAuthnProviderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthnProvider) {
		return nil, false
	}
	return o.IsAuthnProvider, true
}

// HasIsAuthnProvider returns a boolean if a field has been set.
func (o *CredentialMapping) HasIsAuthnProvider() bool {
	if o != nil && !IsNil(o.IsAuthnProvider) {
		return true
	}

	return false
}

// SetIsAuthnProvider gets a reference to the given bool and assigns it to the IsAuthnProvider field.
func (o *CredentialMapping) SetIsAuthnProvider(v bool) {
	o.IsAuthnProvider = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialMapping) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialMapping) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CredentialMapping) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CredentialMapping) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CredentialMapping) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CredentialMapping) UnsetUrl() {
	o.Url.Unset()
}

// GetCredId returns the CredId field value
func (o *CredentialMapping) GetCredId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CredId
}

// GetCredIdOk returns a tuple with the CredId field value
// and a boolean to check if the value has been set.
func (o *CredentialMapping) GetCredIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CredId, true
}

// SetCredId sets field value
func (o *CredentialMapping) SetCredId(v int32) {
	o.CredId = v
}

// GetProduct returns the Product field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialMapping) GetProduct() int32 {
	if o == nil || IsNil(o.Product.Get()) {
		var ret int32
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialMapping) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// HasProduct returns a boolean if a field has been set.
func (o *CredentialMapping) HasProduct() bool {
	if o != nil && o.Product.IsSet() {
		return true
	}

	return false
}

// SetProduct gets a reference to the given NullableInt32 and assigns it to the Product field.
func (o *CredentialMapping) SetProduct(v int32) {
	o.Product.Set(&v)
}
// SetProductNil sets the value for Product to be an explicit nil
func (o *CredentialMapping) SetProductNil() {
	o.Product.Set(nil)
}

// UnsetProduct ensures that no value is present for Product, not even an explicit nil
func (o *CredentialMapping) UnsetProduct() {
	o.Product.Unset()
}

// GetFinding returns the Finding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialMapping) GetFinding() int32 {
	if o == nil || IsNil(o.Finding.Get()) {
		var ret int32
		return ret
	}
	return *o.Finding.Get()
}

// GetFindingOk returns a tuple with the Finding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialMapping) GetFindingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Finding.Get(), o.Finding.IsSet()
}

// HasFinding returns a boolean if a field has been set.
func (o *CredentialMapping) HasFinding() bool {
	if o != nil && o.Finding.IsSet() {
		return true
	}

	return false
}

// SetFinding gets a reference to the given NullableInt32 and assigns it to the Finding field.
func (o *CredentialMapping) SetFinding(v int32) {
	o.Finding.Set(&v)
}
// SetFindingNil sets the value for Finding to be an explicit nil
func (o *CredentialMapping) SetFindingNil() {
	o.Finding.Set(nil)
}

// UnsetFinding ensures that no value is present for Finding, not even an explicit nil
func (o *CredentialMapping) UnsetFinding() {
	o.Finding.Unset()
}

// GetEngagement returns the Engagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialMapping) GetEngagement() int32 {
	if o == nil || IsNil(o.Engagement.Get()) {
		var ret int32
		return ret
	}
	return *o.Engagement.Get()
}

// GetEngagementOk returns a tuple with the Engagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialMapping) GetEngagementOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Engagement.Get(), o.Engagement.IsSet()
}

// HasEngagement returns a boolean if a field has been set.
func (o *CredentialMapping) HasEngagement() bool {
	if o != nil && o.Engagement.IsSet() {
		return true
	}

	return false
}

// SetEngagement gets a reference to the given NullableInt32 and assigns it to the Engagement field.
func (o *CredentialMapping) SetEngagement(v int32) {
	o.Engagement.Set(&v)
}
// SetEngagementNil sets the value for Engagement to be an explicit nil
func (o *CredentialMapping) SetEngagementNil() {
	o.Engagement.Set(nil)
}

// UnsetEngagement ensures that no value is present for Engagement, not even an explicit nil
func (o *CredentialMapping) UnsetEngagement() {
	o.Engagement.Unset()
}

// GetTest returns the Test field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialMapping) GetTest() int32 {
	if o == nil || IsNil(o.Test.Get()) {
		var ret int32
		return ret
	}
	return *o.Test.Get()
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialMapping) GetTestOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Test.Get(), o.Test.IsSet()
}

// HasTest returns a boolean if a field has been set.
func (o *CredentialMapping) HasTest() bool {
	if o != nil && o.Test.IsSet() {
		return true
	}

	return false
}

// SetTest gets a reference to the given NullableInt32 and assigns it to the Test field.
func (o *CredentialMapping) SetTest(v int32) {
	o.Test.Set(&v)
}
// SetTestNil sets the value for Test to be an explicit nil
func (o *CredentialMapping) SetTestNil() {
	o.Test.Set(nil)
}

// UnsetTest ensures that no value is present for Test, not even an explicit nil
func (o *CredentialMapping) UnsetTest() {
	o.Test.Unset()
}

func (o CredentialMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.IsAuthnProvider) {
		toSerialize["is_authn_provider"] = o.IsAuthnProvider
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	toSerialize["cred_id"] = o.CredId
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

func (o *CredentialMapping) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"cred_id",
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

	varCredentialMapping := _CredentialMapping{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCredentialMapping)

	if err != nil {
		return err
	}

	*o = CredentialMapping(varCredentialMapping)

	return err
}

type NullableCredentialMapping struct {
	value *CredentialMapping
	isSet bool
}

func (v NullableCredentialMapping) Get() *CredentialMapping {
	return v.value
}

func (v *NullableCredentialMapping) Set(val *CredentialMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialMapping(val *CredentialMapping) *NullableCredentialMapping {
	return &NullableCredentialMapping{value: val, isSet: true}
}

func (v NullableCredentialMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


