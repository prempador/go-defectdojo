/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the AppAnalysis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAnalysis{}

// AppAnalysis struct for AppAnalysis
type AppAnalysis struct {
	Id int32 `json:"id"`
	Tags []string `json:"tags,omitempty"`
	Name string `json:"name"`
	Confidence NullableInt32 `json:"confidence,omitempty"`
	Version NullableString `json:"version,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	Website NullableString `json:"website,omitempty"`
	WebsiteFound NullableString `json:"website_found,omitempty"`
	Created time.Time `json:"created"`
	Product int32 `json:"product"`
	User int32 `json:"user"`
	Prefetch *AppAnalysisPrefetch `json:"prefetch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAnalysis AppAnalysis

// NewAppAnalysis instantiates a new AppAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAnalysis(id int32, name string, created time.Time, product int32, user int32) *AppAnalysis {
	this := AppAnalysis{}
	this.Id = id
	this.Name = name
	this.Created = created
	this.Product = product
	this.User = user
	return &this
}

// NewAppAnalysisWithDefaults instantiates a new AppAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAnalysisWithDefaults() *AppAnalysis {
	this := AppAnalysis{}
	return &this
}

// GetId returns the Id field value
func (o *AppAnalysis) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppAnalysis) SetId(v int32) {
	o.Id = v
}


// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AppAnalysis) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AppAnalysis) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AppAnalysis) SetTags(v []string) {
	o.Tags = v
}

// GetName returns the Name field value
func (o *AppAnalysis) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppAnalysis) SetName(v string) {
	o.Name = v
}


// GetConfidence returns the Confidence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAnalysis) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence.Get()) {
		var ret int32
		return ret
	}
	return *o.Confidence.Get()
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAnalysis) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Confidence.Get(), o.Confidence.IsSet()
}

// HasConfidence returns a boolean if a field has been set.
func (o *AppAnalysis) HasConfidence() bool {
	if o != nil && o.Confidence.IsSet() {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given NullableInt32 and assigns it to the Confidence field.
func (o *AppAnalysis) SetConfidence(v int32) {
	o.Confidence.Set(&v)
}
// SetConfidenceNil sets the value for Confidence to be an explicit nil
func (o *AppAnalysis) SetConfidenceNil() {
	o.Confidence.Set(nil)
}

// UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
func (o *AppAnalysis) UnsetConfidence() {
	o.Confidence.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAnalysis) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAnalysis) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *AppAnalysis) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *AppAnalysis) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *AppAnalysis) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *AppAnalysis) UnsetVersion() {
	o.Version.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAnalysis) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAnalysis) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *AppAnalysis) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *AppAnalysis) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *AppAnalysis) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *AppAnalysis) UnsetIcon() {
	o.Icon.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAnalysis) GetWebsite() string {
	if o == nil || IsNil(o.Website.Get()) {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAnalysis) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *AppAnalysis) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *AppAnalysis) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *AppAnalysis) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *AppAnalysis) UnsetWebsite() {
	o.Website.Unset()
}

// GetWebsiteFound returns the WebsiteFound field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAnalysis) GetWebsiteFound() string {
	if o == nil || IsNil(o.WebsiteFound.Get()) {
		var ret string
		return ret
	}
	return *o.WebsiteFound.Get()
}

// GetWebsiteFoundOk returns a tuple with the WebsiteFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAnalysis) GetWebsiteFoundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteFound.Get(), o.WebsiteFound.IsSet()
}

// HasWebsiteFound returns a boolean if a field has been set.
func (o *AppAnalysis) HasWebsiteFound() bool {
	if o != nil && o.WebsiteFound.IsSet() {
		return true
	}

	return false
}

// SetWebsiteFound gets a reference to the given NullableString and assigns it to the WebsiteFound field.
func (o *AppAnalysis) SetWebsiteFound(v string) {
	o.WebsiteFound.Set(&v)
}
// SetWebsiteFoundNil sets the value for WebsiteFound to be an explicit nil
func (o *AppAnalysis) SetWebsiteFoundNil() {
	o.WebsiteFound.Set(nil)
}

// UnsetWebsiteFound ensures that no value is present for WebsiteFound, not even an explicit nil
func (o *AppAnalysis) UnsetWebsiteFound() {
	o.WebsiteFound.Unset()
}

// GetCreated returns the Created field value
func (o *AppAnalysis) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AppAnalysis) SetCreated(v time.Time) {
	o.Created = v
}


// GetProduct returns the Product field value
func (o *AppAnalysis) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *AppAnalysis) SetProduct(v int32) {
	o.Product = v
}


// GetUser returns the User field value
func (o *AppAnalysis) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AppAnalysis) SetUser(v int32) {
	o.User = v
}


// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *AppAnalysis) GetPrefetch() AppAnalysisPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret AppAnalysisPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAnalysis) GetPrefetchOk() (*AppAnalysisPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *AppAnalysis) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given AppAnalysisPrefetch and assigns it to the Prefetch field.
func (o *AppAnalysis) SetPrefetch(v AppAnalysisPrefetch) {
	o.Prefetch = &v
}

func (o AppAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAnalysis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["name"] = o.Name
	if o.Confidence.IsSet() {
		toSerialize["confidence"] = o.Confidence.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if o.WebsiteFound.IsSet() {
		toSerialize["website_found"] = o.WebsiteFound.Get()
	}
	toSerialize["created"] = o.Created
	toSerialize["product"] = o.Product
	toSerialize["user"] = o.User
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAnalysis) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"created",
		"product",
		"user",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{} {
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == ""{
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil{
			return err
		}
	}
	varAppAnalysis := _AppAnalysis{}

	err = json.Unmarshal(data, &varAppAnalysis)

	if err != nil {
		return err
	}

	*o = AppAnalysis(varAppAnalysis)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "name")
		delete(additionalProperties, "confidence")
		delete(additionalProperties, "version")
		delete(additionalProperties, "icon")
		delete(additionalProperties, "website")
		delete(additionalProperties, "website_found")
		delete(additionalProperties, "created")
		delete(additionalProperties, "product")
		delete(additionalProperties, "user")
		delete(additionalProperties, "prefetch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAnalysis struct {
	value *AppAnalysis
	isSet bool
}

func (v NullableAppAnalysis) Get() *AppAnalysis {
	return v.value
}

func (v *NullableAppAnalysis) Set(val *AppAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAnalysis(val *AppAnalysis) *NullableAppAnalysis {
	return &NullableAppAnalysis{value: val, isSet: true}
}

func (v NullableAppAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


