/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchedLanguageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedLanguageRequest{}

// PatchedLanguageRequest struct for PatchedLanguageRequest
type PatchedLanguageRequest struct {
	Files NullableInt32 `json:"files,omitempty"`
	Blank NullableInt32 `json:"blank,omitempty"`
	Comment NullableInt32 `json:"comment,omitempty"`
	Code NullableInt32 `json:"code,omitempty"`
	Language *int32 `json:"language,omitempty"`
	Product *int32 `json:"product,omitempty"`
	User NullableInt32 `json:"user,omitempty"`
}

// NewPatchedLanguageRequest instantiates a new PatchedLanguageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedLanguageRequest() *PatchedLanguageRequest {
	this := PatchedLanguageRequest{}
	return &this
}

// NewPatchedLanguageRequestWithDefaults instantiates a new PatchedLanguageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedLanguageRequestWithDefaults() *PatchedLanguageRequest {
	this := PatchedLanguageRequest{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLanguageRequest) GetFiles() int32 {
	if o == nil || IsNil(o.Files.Get()) {
		var ret int32
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLanguageRequest) GetFilesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasFiles() bool {
	if o != nil && o.Files.IsSet() {
		return true
	}

	return false
}

// SetFiles gets a reference to the given NullableInt32 and assigns it to the Files field.
func (o *PatchedLanguageRequest) SetFiles(v int32) {
	o.Files.Set(&v)
}
// SetFilesNil sets the value for Files to be an explicit nil
func (o *PatchedLanguageRequest) SetFilesNil() {
	o.Files.Set(nil)
}

// UnsetFiles ensures that no value is present for Files, not even an explicit nil
func (o *PatchedLanguageRequest) UnsetFiles() {
	o.Files.Unset()
}

// GetBlank returns the Blank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLanguageRequest) GetBlank() int32 {
	if o == nil || IsNil(o.Blank.Get()) {
		var ret int32
		return ret
	}
	return *o.Blank.Get()
}

// GetBlankOk returns a tuple with the Blank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLanguageRequest) GetBlankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blank.Get(), o.Blank.IsSet()
}

// HasBlank returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasBlank() bool {
	if o != nil && o.Blank.IsSet() {
		return true
	}

	return false
}

// SetBlank gets a reference to the given NullableInt32 and assigns it to the Blank field.
func (o *PatchedLanguageRequest) SetBlank(v int32) {
	o.Blank.Set(&v)
}
// SetBlankNil sets the value for Blank to be an explicit nil
func (o *PatchedLanguageRequest) SetBlankNil() {
	o.Blank.Set(nil)
}

// UnsetBlank ensures that no value is present for Blank, not even an explicit nil
func (o *PatchedLanguageRequest) UnsetBlank() {
	o.Blank.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLanguageRequest) GetComment() int32 {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret int32
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLanguageRequest) GetCommentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableInt32 and assigns it to the Comment field.
func (o *PatchedLanguageRequest) SetComment(v int32) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *PatchedLanguageRequest) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *PatchedLanguageRequest) UnsetComment() {
	o.Comment.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLanguageRequest) GetCode() int32 {
	if o == nil || IsNil(o.Code.Get()) {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLanguageRequest) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *PatchedLanguageRequest) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *PatchedLanguageRequest) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *PatchedLanguageRequest) UnsetCode() {
	o.Code.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PatchedLanguageRequest) GetLanguage() int32 {
	if o == nil || IsNil(o.Language) {
		var ret int32
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLanguageRequest) GetLanguageOk() (*int32, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given int32 and assigns it to the Language field.
func (o *PatchedLanguageRequest) SetLanguage(v int32) {
	o.Language = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PatchedLanguageRequest) GetProduct() int32 {
	if o == nil || IsNil(o.Product) {
		var ret int32
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedLanguageRequest) GetProductOk() (*int32, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given int32 and assigns it to the Product field.
func (o *PatchedLanguageRequest) SetProduct(v int32) {
	o.Product = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedLanguageRequest) GetUser() int32 {
	if o == nil || IsNil(o.User.Get()) {
		var ret int32
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedLanguageRequest) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *PatchedLanguageRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableInt32 and assigns it to the User field.
func (o *PatchedLanguageRequest) SetUser(v int32) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *PatchedLanguageRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *PatchedLanguageRequest) UnsetUser() {
	o.User.Unset()
}

func (o PatchedLanguageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedLanguageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

type NullablePatchedLanguageRequest struct {
	value *PatchedLanguageRequest
	isSet bool
}

func (v NullablePatchedLanguageRequest) Get() *PatchedLanguageRequest {
	return v.value
}

func (v *NullablePatchedLanguageRequest) Set(val *PatchedLanguageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedLanguageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedLanguageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedLanguageRequest(val *PatchedLanguageRequest) *NullablePatchedLanguageRequest {
	return &NullablePatchedLanguageRequest{value: val, isSet: true}
}

func (v NullablePatchedLanguageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedLanguageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


