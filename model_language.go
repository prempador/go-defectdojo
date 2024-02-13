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

// checks if the Language type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Language{}

// Language struct for Language
type Language struct {
	Id int32 `json:"id"`
	Files NullableInt32 `json:"files,omitempty"`
	Blank NullableInt32 `json:"blank,omitempty"`
	Comment NullableInt32 `json:"comment,omitempty"`
	Code NullableInt32 `json:"code,omitempty"`
	Created time.Time `json:"created"`
	Language int32 `json:"language"`
	Product int32 `json:"product"`
	User NullableInt32 `json:"user,omitempty"`
	Prefetch *LanguagePrefetch `json:"prefetch,omitempty"`
}

type _Language Language

// NewLanguage instantiates a new Language object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguage(id int32, created time.Time, language int32, product int32) *Language {
	this := Language{}
	this.Id = id
	this.Created = created
	this.Language = language
	this.Product = product
	return &this
}

// NewLanguageWithDefaults instantiates a new Language object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageWithDefaults() *Language {
	this := Language{}
	return &this
}

// GetId returns the Id field value
func (o *Language) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Language) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Language) SetId(v int32) {
	o.Id = v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetFiles() int32 {
	if o == nil || IsNil(o.Files.Get()) {
		var ret int32
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetFilesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a field has been set.
func (o *Language) HasFiles() bool {
	if o != nil && o.Files.IsSet() {
		return true
	}

	return false
}

// SetFiles gets a reference to the given NullableInt32 and assigns it to the Files field.
func (o *Language) SetFiles(v int32) {
	o.Files.Set(&v)
}
// SetFilesNil sets the value for Files to be an explicit nil
func (o *Language) SetFilesNil() {
	o.Files.Set(nil)
}

// UnsetFiles ensures that no value is present for Files, not even an explicit nil
func (o *Language) UnsetFiles() {
	o.Files.Unset()
}

// GetBlank returns the Blank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetBlank() int32 {
	if o == nil || IsNil(o.Blank.Get()) {
		var ret int32
		return ret
	}
	return *o.Blank.Get()
}

// GetBlankOk returns a tuple with the Blank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetBlankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blank.Get(), o.Blank.IsSet()
}

// HasBlank returns a boolean if a field has been set.
func (o *Language) HasBlank() bool {
	if o != nil && o.Blank.IsSet() {
		return true
	}

	return false
}

// SetBlank gets a reference to the given NullableInt32 and assigns it to the Blank field.
func (o *Language) SetBlank(v int32) {
	o.Blank.Set(&v)
}
// SetBlankNil sets the value for Blank to be an explicit nil
func (o *Language) SetBlankNil() {
	o.Blank.Set(nil)
}

// UnsetBlank ensures that no value is present for Blank, not even an explicit nil
func (o *Language) UnsetBlank() {
	o.Blank.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetComment() int32 {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret int32
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetCommentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *Language) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableInt32 and assigns it to the Comment field.
func (o *Language) SetComment(v int32) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *Language) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *Language) UnsetComment() {
	o.Comment.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetCode() int32 {
	if o == nil || IsNil(o.Code.Get()) {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *Language) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *Language) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *Language) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *Language) UnsetCode() {
	o.Code.Unset()
}

// GetCreated returns the Created field value
func (o *Language) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Language) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Language) SetCreated(v time.Time) {
	o.Created = v
}

// GetLanguage returns the Language field value
func (o *Language) GetLanguage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *Language) GetLanguageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *Language) SetLanguage(v int32) {
	o.Language = v
}

// GetProduct returns the Product field value
func (o *Language) GetProduct() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *Language) GetProductOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *Language) SetProduct(v int32) {
	o.Product = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Language) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Language) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *Language) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *Language) SetUser(v int32) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *Language) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *Language) UnsetUser() {
	o.User.Unset()
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *Language) GetPrefetch() LanguagePrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret LanguagePrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Language) GetPrefetchOk() (*LanguagePrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *Language) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given LanguagePrefetch and assigns it to the Prefetch field.
func (o *Language) SetPrefetch(v LanguagePrefetch) {
	o.Prefetch = &v
}

func (o Language) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Language) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Files.IsSet() {
		toSerialize["files"] = o.Files.Get()
	}
	if o.Blank.IsSet() {
		toSerialize["blank"] = o.Blank.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	toSerialize["created"] = o.Created
	toSerialize["language"] = o.Language
	toSerialize["product"] = o.Product
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}
	return toSerialize, nil
}

func (o *Language) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created",
		"language",
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

	varLanguage := _Language{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLanguage)

	if err != nil {
		return err
	}

	*o = Language(varLanguage)

	return err
}

type NullableLanguage struct {
	value *Language
	isSet bool
}

func (v NullableLanguage) Get() *Language {
	return v.value
}

func (v *NullableLanguage) Set(val *Language) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguage(val *Language) *NullableLanguage {
	return &NullableLanguage{value: val, isSet: true}
}

func (v NullableLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


