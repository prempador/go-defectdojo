/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the NetworkLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLocations{}

// NetworkLocations struct for NetworkLocations
type NetworkLocations struct {
	Id int32 `json:"id"`
	// Location of network testing: Examples: VPN, Internet or Internal.
	Location string `json:"location"`
}

type _NetworkLocations NetworkLocations

// NewNetworkLocations instantiates a new NetworkLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLocations(id int32, location string) *NetworkLocations {
	this := NetworkLocations{}
	this.Id = id
	this.Location = location
	return &this
}

// NewNetworkLocationsWithDefaults instantiates a new NetworkLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLocationsWithDefaults() *NetworkLocations {
	this := NetworkLocations{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkLocations) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkLocations) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkLocations) SetId(v int32) {
	o.Id = v
}

// GetLocation returns the Location field value
func (o *NetworkLocations) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *NetworkLocations) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *NetworkLocations) SetLocation(v string) {
	o.Location = v
}

func (o NetworkLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

func (o *NetworkLocations) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"location",
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

	varNetworkLocations := _NetworkLocations{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkLocations)

	if err != nil {
		return err
	}

	*o = NetworkLocations(varNetworkLocations)

	return err
}

type NullableNetworkLocations struct {
	value *NetworkLocations
	isSet bool
}

func (v NullableNetworkLocations) Get() *NetworkLocations {
	return v.value
}

func (v *NullableNetworkLocations) Set(val *NetworkLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLocations(val *NetworkLocations) *NullableNetworkLocations {
	return &NullableNetworkLocations{value: val, isSet: true}
}

func (v NullableNetworkLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


