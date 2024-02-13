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

// checks if the EngagementPresets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EngagementPresets{}

// EngagementPresets struct for EngagementPresets
type EngagementPresets struct {
	Id int32 `json:"id"`
	// Brief description of preset.
	Title *string `json:"title,omitempty"`
	// Description of what needs to be tested or setting up environment for testing
	Notes NullableString `json:"notes,omitempty"`
	// Scope of Engagement testing, IP's/Resources/URL's)
	Scope *string `json:"scope,omitempty"`
	Created time.Time `json:"created"`
	Product int32 `json:"product"`
	TestType []int32 `json:"test_type,omitempty"`
	NetworkLocations []int32 `json:"network_locations,omitempty"`
}

type _EngagementPresets EngagementPresets

// NewEngagementPresets instantiates a new EngagementPresets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngagementPresets(id int32, created time.Time, product int32) *EngagementPresets {
	this := EngagementPresets{}
	this.Id = id
	this.Created = created
	this.Product = product
	return &this
}

// NewEngagementPresetsWithDefaults instantiates a new EngagementPresets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngagementPresetsWithDefaults() *EngagementPresets {
	this := EngagementPresets{}
	return &this
}

// GetId returns the Id field value
func (o *EngagementPresets) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EngagementPresets) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EngagementPresets) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EngagementPresets) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EngagementPresets) SetTitle(v string) {
	o.Title = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngagementPresets) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngagementPresets) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *EngagementPresets) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *EngagementPresets) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *EngagementPresets) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *EngagementPresets) UnsetNotes() {
	o.Notes.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *EngagementPresets) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *EngagementPresets) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *EngagementPresets) SetScope(v string) {
	o.Scope = &v
}

// GetCreated returns the Created field value
func (o *EngagementPresets) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *EngagementPresets) SetCreated(v time.Time) {
	o.Created = v
}

// GetProduct returns the Product field value
func (o *EngagementPresets) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *EngagementPresets) SetProduct(v int32) {
	o.Product = v
}

// GetTestType returns the TestType field value if set, zero value otherwise.
func (o *EngagementPresets) GetTestType() []int32 {
	if o == nil || IsNil(o.TestType) {
		var ret []int32
		return ret
	}
	return o.TestType
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetTestTypeOk() ([]int32, bool) {
	if o == nil || IsNil(o.TestType) {
		return nil, false
	}
	return o.TestType, true
}

// HasTestType returns a boolean if a field has been set.
func (o *EngagementPresets) HasTestType() bool {
	if o != nil && !IsNil(o.TestType) {
		return true
	}

	return false
}

// SetTestType gets a reference to the given []int32 and assigns it to the TestType field.
func (o *EngagementPresets) SetTestType(v []int32) {
	o.TestType = v
}

// GetNetworkLocations returns the NetworkLocations field value if set, zero value otherwise.
func (o *EngagementPresets) GetNetworkLocations() []int32 {
	if o == nil || IsNil(o.NetworkLocations) {
		var ret []int32
		return ret
	}
	return o.NetworkLocations
}

// GetNetworkLocationsOk returns a tuple with the NetworkLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngagementPresets) GetNetworkLocationsOk() ([]int32, bool) {
	if o == nil || IsNil(o.NetworkLocations) {
		return nil, false
	}
	return o.NetworkLocations, true
}

// HasNetworkLocations returns a boolean if a field has been set.
func (o *EngagementPresets) HasNetworkLocations() bool {
	if o != nil && !IsNil(o.NetworkLocations) {
		return true
	}

	return false
}

// SetNetworkLocations gets a reference to the given []int32 and assigns it to the NetworkLocations field.
func (o *EngagementPresets) SetNetworkLocations(v []int32) {
	o.NetworkLocations = v
}

func (o EngagementPresets) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EngagementPresets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["created"] = o.Created
	toSerialize["product"] = o.Product
	if !IsNil(o.TestType) {
		toSerialize["test_type"] = o.TestType
	}
	if !IsNil(o.NetworkLocations) {
		toSerialize["network_locations"] = o.NetworkLocations
	}
	return toSerialize, nil
}

func (o *EngagementPresets) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created",
		"product",
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

	varEngagementPresets := _EngagementPresets{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEngagementPresets)

	if err != nil {
		return err
	}

	*o = EngagementPresets(varEngagementPresets)

	return err
}

type NullableEngagementPresets struct {
	value *EngagementPresets
	isSet bool
}

func (v NullableEngagementPresets) Get() *EngagementPresets {
	return v.value
}

func (v *NullableEngagementPresets) Set(val *EngagementPresets) {
	v.value = val
	v.isSet = true
}

func (v NullableEngagementPresets) IsSet() bool {
	return v.isSet
}

func (v *NullableEngagementPresets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngagementPresets(val *EngagementPresets) *NullableEngagementPresets {
	return &NullableEngagementPresets{value: val, isSet: true}
}

func (v NullableEngagementPresets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngagementPresets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


