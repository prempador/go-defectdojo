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

// checks if the TestImportFindingAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestImportFindingAction{}

// TestImportFindingAction struct for TestImportFindingAction
type TestImportFindingAction struct {
	Id int32 `json:"id"`
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	// * `N` - created * `C` - closed * `R` - reactivated * `U` - left untouched
	Action NullableString `json:"action,omitempty"`
	TestImport int32 `json:"test_import"`
	Finding int32 `json:"finding"`
	AdditionalProperties map[string]interface{}
}

type _TestImportFindingAction TestImportFindingAction

// NewTestImportFindingAction instantiates a new TestImportFindingAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestImportFindingAction(id int32, created time.Time, modified time.Time, testImport int32, finding int32) *TestImportFindingAction {
	this := TestImportFindingAction{}
	this.Id = id
	this.Created = created
	this.Modified = modified
	this.TestImport = testImport
	this.Finding = finding
	return &this
}

// NewTestImportFindingActionWithDefaults instantiates a new TestImportFindingAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestImportFindingActionWithDefaults() *TestImportFindingAction {
	this := TestImportFindingAction{}
	return &this
}

// GetId returns the Id field value
func (o *TestImportFindingAction) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestImportFindingAction) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestImportFindingAction) SetId(v int32) {
	o.Id = v
}


// GetCreated returns the Created field value
func (o *TestImportFindingAction) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TestImportFindingAction) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TestImportFindingAction) SetCreated(v time.Time) {
	o.Created = v
}


// GetModified returns the Modified field value
func (o *TestImportFindingAction) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *TestImportFindingAction) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *TestImportFindingAction) SetModified(v time.Time) {
	o.Modified = v
}


// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TestImportFindingAction) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TestImportFindingAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *TestImportFindingAction) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *TestImportFindingAction) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *TestImportFindingAction) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *TestImportFindingAction) UnsetAction() {
	o.Action.Unset()
}

// GetTestImport returns the TestImport field value
func (o *TestImportFindingAction) GetTestImport() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestImport
}

// GetTestImportOk returns a tuple with the TestImport field value
// and a boolean to check if the value has been set.
func (o *TestImportFindingAction) GetTestImportOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestImport, true
}

// SetTestImport sets field value
func (o *TestImportFindingAction) SetTestImport(v int32) {
	o.TestImport = v
}


// GetFinding returns the Finding field value
func (o *TestImportFindingAction) GetFinding() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Finding
}

// GetFindingOk returns a tuple with the Finding field value
// and a boolean to check if the value has been set.
func (o *TestImportFindingAction) GetFindingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finding, true
}

// SetFinding sets field value
func (o *TestImportFindingAction) SetFinding(v int32) {
	o.Finding = v
}


func (o TestImportFindingAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestImportFindingAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["modified"] = o.Modified
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	toSerialize["test_import"] = o.TestImport
	toSerialize["finding"] = o.Finding

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestImportFindingAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created",
		"modified",
		"test_import",
		"finding",
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
	varTestImportFindingAction := _TestImportFindingAction{}

	err = json.Unmarshal(data, &varTestImportFindingAction)

	if err != nil {
		return err
	}

	*o = TestImportFindingAction(varTestImportFindingAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "action")
		delete(additionalProperties, "test_import")
		delete(additionalProperties, "finding")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestImportFindingAction struct {
	value *TestImportFindingAction
	isSet bool
}

func (v NullableTestImportFindingAction) Get() *TestImportFindingAction {
	return v.value
}

func (v *NullableTestImportFindingAction) Set(val *TestImportFindingAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTestImportFindingAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTestImportFindingAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestImportFindingAction(val *TestImportFindingAction) *NullableTestImportFindingAction {
	return &NullableTestImportFindingAction{value: val, isSet: true}
}

func (v NullableTestImportFindingAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestImportFindingAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


