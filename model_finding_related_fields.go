/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"fmt"
)

// checks if the FindingRelatedFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindingRelatedFields{}

// FindingRelatedFields struct for FindingRelatedFields
type FindingRelatedFields struct {
	Test FindingTest `json:"test"`
	Jira JIRAIssue `json:"jira"`
	AdditionalProperties map[string]interface{}
}

type _FindingRelatedFields FindingRelatedFields

// NewFindingRelatedFields instantiates a new FindingRelatedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindingRelatedFields(test FindingTest, jira JIRAIssue) *FindingRelatedFields {
	this := FindingRelatedFields{}
	this.Test = test
	this.Jira = jira
	return &this
}

// NewFindingRelatedFieldsWithDefaults instantiates a new FindingRelatedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindingRelatedFieldsWithDefaults() *FindingRelatedFields {
	this := FindingRelatedFields{}
	return &this
}

// GetTest returns the Test field value
func (o *FindingRelatedFields) GetTest() FindingTest {
	if o == nil {
		var ret FindingTest
		return ret
	}

	return o.Test
}

// GetTestOk returns a tuple with the Test field value
// and a boolean to check if the value has been set.
func (o *FindingRelatedFields) GetTestOk() (*FindingTest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Test, true
}

// SetTest sets field value
func (o *FindingRelatedFields) SetTest(v FindingTest) {
	o.Test = v
}


// GetJira returns the Jira field value
func (o *FindingRelatedFields) GetJira() JIRAIssue {
	if o == nil {
		var ret JIRAIssue
		return ret
	}

	return o.Jira
}

// GetJiraOk returns a tuple with the Jira field value
// and a boolean to check if the value has been set.
func (o *FindingRelatedFields) GetJiraOk() (*JIRAIssue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jira, true
}

// SetJira sets field value
func (o *FindingRelatedFields) SetJira(v JIRAIssue) {
	o.Jira = v
}


func (o FindingRelatedFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindingRelatedFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["test"] = o.Test
	toSerialize["jira"] = o.Jira

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FindingRelatedFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"test",
		"jira",
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
	varFindingRelatedFields := _FindingRelatedFields{}

	err = json.Unmarshal(data, &varFindingRelatedFields)

	if err != nil {
		return err
	}

	*o = FindingRelatedFields(varFindingRelatedFields)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "test")
		delete(additionalProperties, "jira")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFindingRelatedFields struct {
	value *FindingRelatedFields
	isSet bool
}

func (v NullableFindingRelatedFields) Get() *FindingRelatedFields {
	return v.value
}

func (v *NullableFindingRelatedFields) Set(val *FindingRelatedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableFindingRelatedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableFindingRelatedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindingRelatedFields(val *FindingRelatedFields) *NullableFindingRelatedFields {
	return &NullableFindingRelatedFields{value: val, isSet: true}
}

func (v NullableFindingRelatedFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindingRelatedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


