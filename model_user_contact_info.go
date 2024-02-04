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

// checks if the UserContactInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserContactInfo{}

// UserContactInfo struct for UserContactInfo
type UserContactInfo struct {
	Id int32 `json:"id"`
	Title NullableString `json:"title,omitempty"`
	// Phone number must be entered in the format: '+999999999'. Up to 15 digits allowed.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// Phone number must be entered in the format: '+999999999'. Up to 15 digits allowed.
	CellNumber *string `json:"cell_number,omitempty"`
	TwitterUsername NullableString `json:"twitter_username,omitempty"`
	GithubUsername NullableString `json:"github_username,omitempty"`
	// Email address associated with your slack account
	SlackUsername NullableString `json:"slack_username,omitempty"`
	SlackUserId NullableString `json:"slack_user_id,omitempty"`
	// Instead of async deduping a finding the findings will be deduped synchronously and will 'block' the user until completion.
	BlockExecution *bool `json:"block_execution,omitempty"`
	// Forces this user to reset their password on next login.
	ForcePasswordReset *bool `json:"force_password_reset,omitempty"`
	User int32 `json:"user"`
	Prefetch *PaginatedUserContactInfoListPrefetch `json:"prefetch,omitempty"`
}

type _UserContactInfo UserContactInfo

// NewUserContactInfo instantiates a new UserContactInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserContactInfo(id int32, user int32) *UserContactInfo {
	this := UserContactInfo{}
	this.Id = id
	this.User = user
	return &this
}

// NewUserContactInfoWithDefaults instantiates a new UserContactInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserContactInfoWithDefaults() *UserContactInfo {
	this := UserContactInfo{}
	return &this
}

// GetId returns the Id field value
func (o *UserContactInfo) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserContactInfo) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserContactInfo) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserContactInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UserContactInfo) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UserContactInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UserContactInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UserContactInfo) UnsetTitle() {
	o.Title.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserContactInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserContactInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserContactInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCellNumber returns the CellNumber field value if set, zero value otherwise.
func (o *UserContactInfo) GetCellNumber() string {
	if o == nil || IsNil(o.CellNumber) {
		var ret string
		return ret
	}
	return *o.CellNumber
}

// GetCellNumberOk returns a tuple with the CellNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetCellNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CellNumber) {
		return nil, false
	}
	return o.CellNumber, true
}

// HasCellNumber returns a boolean if a field has been set.
func (o *UserContactInfo) HasCellNumber() bool {
	if o != nil && !IsNil(o.CellNumber) {
		return true
	}

	return false
}

// SetCellNumber gets a reference to the given string and assigns it to the CellNumber field.
func (o *UserContactInfo) SetCellNumber(v string) {
	o.CellNumber = &v
}

// GetTwitterUsername returns the TwitterUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserContactInfo) GetTwitterUsername() string {
	if o == nil || IsNil(o.TwitterUsername.Get()) {
		var ret string
		return ret
	}
	return *o.TwitterUsername.Get()
}

// GetTwitterUsernameOk returns a tuple with the TwitterUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserContactInfo) GetTwitterUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwitterUsername.Get(), o.TwitterUsername.IsSet()
}

// HasTwitterUsername returns a boolean if a field has been set.
func (o *UserContactInfo) HasTwitterUsername() bool {
	if o != nil && o.TwitterUsername.IsSet() {
		return true
	}

	return false
}

// SetTwitterUsername gets a reference to the given NullableString and assigns it to the TwitterUsername field.
func (o *UserContactInfo) SetTwitterUsername(v string) {
	o.TwitterUsername.Set(&v)
}
// SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil
func (o *UserContactInfo) SetTwitterUsernameNil() {
	o.TwitterUsername.Set(nil)
}

// UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
func (o *UserContactInfo) UnsetTwitterUsername() {
	o.TwitterUsername.Unset()
}

// GetGithubUsername returns the GithubUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserContactInfo) GetGithubUsername() string {
	if o == nil || IsNil(o.GithubUsername.Get()) {
		var ret string
		return ret
	}
	return *o.GithubUsername.Get()
}

// GetGithubUsernameOk returns a tuple with the GithubUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserContactInfo) GetGithubUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GithubUsername.Get(), o.GithubUsername.IsSet()
}

// HasGithubUsername returns a boolean if a field has been set.
func (o *UserContactInfo) HasGithubUsername() bool {
	if o != nil && o.GithubUsername.IsSet() {
		return true
	}

	return false
}

// SetGithubUsername gets a reference to the given NullableString and assigns it to the GithubUsername field.
func (o *UserContactInfo) SetGithubUsername(v string) {
	o.GithubUsername.Set(&v)
}
// SetGithubUsernameNil sets the value for GithubUsername to be an explicit nil
func (o *UserContactInfo) SetGithubUsernameNil() {
	o.GithubUsername.Set(nil)
}

// UnsetGithubUsername ensures that no value is present for GithubUsername, not even an explicit nil
func (o *UserContactInfo) UnsetGithubUsername() {
	o.GithubUsername.Unset()
}

// GetSlackUsername returns the SlackUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserContactInfo) GetSlackUsername() string {
	if o == nil || IsNil(o.SlackUsername.Get()) {
		var ret string
		return ret
	}
	return *o.SlackUsername.Get()
}

// GetSlackUsernameOk returns a tuple with the SlackUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserContactInfo) GetSlackUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackUsername.Get(), o.SlackUsername.IsSet()
}

// HasSlackUsername returns a boolean if a field has been set.
func (o *UserContactInfo) HasSlackUsername() bool {
	if o != nil && o.SlackUsername.IsSet() {
		return true
	}

	return false
}

// SetSlackUsername gets a reference to the given NullableString and assigns it to the SlackUsername field.
func (o *UserContactInfo) SetSlackUsername(v string) {
	o.SlackUsername.Set(&v)
}
// SetSlackUsernameNil sets the value for SlackUsername to be an explicit nil
func (o *UserContactInfo) SetSlackUsernameNil() {
	o.SlackUsername.Set(nil)
}

// UnsetSlackUsername ensures that no value is present for SlackUsername, not even an explicit nil
func (o *UserContactInfo) UnsetSlackUsername() {
	o.SlackUsername.Unset()
}

// GetSlackUserId returns the SlackUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserContactInfo) GetSlackUserId() string {
	if o == nil || IsNil(o.SlackUserId.Get()) {
		var ret string
		return ret
	}
	return *o.SlackUserId.Get()
}

// GetSlackUserIdOk returns a tuple with the SlackUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserContactInfo) GetSlackUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackUserId.Get(), o.SlackUserId.IsSet()
}

// HasSlackUserId returns a boolean if a field has been set.
func (o *UserContactInfo) HasSlackUserId() bool {
	if o != nil && o.SlackUserId.IsSet() {
		return true
	}

	return false
}

// SetSlackUserId gets a reference to the given NullableString and assigns it to the SlackUserId field.
func (o *UserContactInfo) SetSlackUserId(v string) {
	o.SlackUserId.Set(&v)
}
// SetSlackUserIdNil sets the value for SlackUserId to be an explicit nil
func (o *UserContactInfo) SetSlackUserIdNil() {
	o.SlackUserId.Set(nil)
}

// UnsetSlackUserId ensures that no value is present for SlackUserId, not even an explicit nil
func (o *UserContactInfo) UnsetSlackUserId() {
	o.SlackUserId.Unset()
}

// GetBlockExecution returns the BlockExecution field value if set, zero value otherwise.
func (o *UserContactInfo) GetBlockExecution() bool {
	if o == nil || IsNil(o.BlockExecution) {
		var ret bool
		return ret
	}
	return *o.BlockExecution
}

// GetBlockExecutionOk returns a tuple with the BlockExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetBlockExecutionOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockExecution) {
		return nil, false
	}
	return o.BlockExecution, true
}

// HasBlockExecution returns a boolean if a field has been set.
func (o *UserContactInfo) HasBlockExecution() bool {
	if o != nil && !IsNil(o.BlockExecution) {
		return true
	}

	return false
}

// SetBlockExecution gets a reference to the given bool and assigns it to the BlockExecution field.
func (o *UserContactInfo) SetBlockExecution(v bool) {
	o.BlockExecution = &v
}

// GetForcePasswordReset returns the ForcePasswordReset field value if set, zero value otherwise.
func (o *UserContactInfo) GetForcePasswordReset() bool {
	if o == nil || IsNil(o.ForcePasswordReset) {
		var ret bool
		return ret
	}
	return *o.ForcePasswordReset
}

// GetForcePasswordResetOk returns a tuple with the ForcePasswordReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetForcePasswordResetOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePasswordReset) {
		return nil, false
	}
	return o.ForcePasswordReset, true
}

// HasForcePasswordReset returns a boolean if a field has been set.
func (o *UserContactInfo) HasForcePasswordReset() bool {
	if o != nil && !IsNil(o.ForcePasswordReset) {
		return true
	}

	return false
}

// SetForcePasswordReset gets a reference to the given bool and assigns it to the ForcePasswordReset field.
func (o *UserContactInfo) SetForcePasswordReset(v bool) {
	o.ForcePasswordReset = &v
}

// GetUser returns the User field value
func (o *UserContactInfo) GetUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *UserContactInfo) SetUser(v int32) {
	o.User = v
}

// GetPrefetch returns the Prefetch field value if set, zero value otherwise.
func (o *UserContactInfo) GetPrefetch() PaginatedUserContactInfoListPrefetch {
	if o == nil || IsNil(o.Prefetch) {
		var ret PaginatedUserContactInfoListPrefetch
		return ret
	}
	return *o.Prefetch
}

// GetPrefetchOk returns a tuple with the Prefetch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserContactInfo) GetPrefetchOk() (*PaginatedUserContactInfoListPrefetch, bool) {
	if o == nil || IsNil(o.Prefetch) {
		return nil, false
	}
	return o.Prefetch, true
}

// HasPrefetch returns a boolean if a field has been set.
func (o *UserContactInfo) HasPrefetch() bool {
	if o != nil && !IsNil(o.Prefetch) {
		return true
	}

	return false
}

// SetPrefetch gets a reference to the given PaginatedUserContactInfoListPrefetch and assigns it to the Prefetch field.
func (o *UserContactInfo) SetPrefetch(v PaginatedUserContactInfoListPrefetch) {
	o.Prefetch = &v
}

func (o UserContactInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserContactInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.CellNumber) {
		toSerialize["cell_number"] = o.CellNumber
	}
	if o.TwitterUsername.IsSet() {
		toSerialize["twitter_username"] = o.TwitterUsername.Get()
	}
	if o.GithubUsername.IsSet() {
		toSerialize["github_username"] = o.GithubUsername.Get()
	}
	if o.SlackUsername.IsSet() {
		toSerialize["slack_username"] = o.SlackUsername.Get()
	}
	if o.SlackUserId.IsSet() {
		toSerialize["slack_user_id"] = o.SlackUserId.Get()
	}
	if !IsNil(o.BlockExecution) {
		toSerialize["block_execution"] = o.BlockExecution
	}
	if !IsNil(o.ForcePasswordReset) {
		toSerialize["force_password_reset"] = o.ForcePasswordReset
	}
	toSerialize["user"] = o.User
	if !IsNil(o.Prefetch) {
		toSerialize["prefetch"] = o.Prefetch
	}
	return toSerialize, nil
}

func (o *UserContactInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user",
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

	varUserContactInfo := _UserContactInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserContactInfo)

	if err != nil {
		return err
	}

	*o = UserContactInfo(varUserContactInfo)

	return err
}

type NullableUserContactInfo struct {
	value *UserContactInfo
	isSet bool
}

func (v NullableUserContactInfo) Get() *UserContactInfo {
	return v.value
}

func (v *NullableUserContactInfo) Set(val *UserContactInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserContactInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserContactInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserContactInfo(val *UserContactInfo) *NullableUserContactInfo {
	return &NullableUserContactInfo{value: val, isSet: true}
}

func (v NullableUserContactInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserContactInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

