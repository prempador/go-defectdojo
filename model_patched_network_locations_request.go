/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"encoding/json"
)

// checks if the PatchedNetworkLocationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedNetworkLocationsRequest{}

// PatchedNetworkLocationsRequest struct for PatchedNetworkLocationsRequest
type PatchedNetworkLocationsRequest struct {
	// Location of network testing: Examples: VPN, Internet or Internal.
	Location *string `json:"location,omitempty"`
}

// NewPatchedNetworkLocationsRequest instantiates a new PatchedNetworkLocationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedNetworkLocationsRequest() *PatchedNetworkLocationsRequest {
	this := PatchedNetworkLocationsRequest{}
	return &this
}

// NewPatchedNetworkLocationsRequestWithDefaults instantiates a new PatchedNetworkLocationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedNetworkLocationsRequestWithDefaults() *PatchedNetworkLocationsRequest {
	this := PatchedNetworkLocationsRequest{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PatchedNetworkLocationsRequest) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedNetworkLocationsRequest) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedNetworkLocationsRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PatchedNetworkLocationsRequest) SetLocation(v string) {
	o.Location = &v
}

func (o PatchedNetworkLocationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedNetworkLocationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

type NullablePatchedNetworkLocationsRequest struct {
	value *PatchedNetworkLocationsRequest
	isSet bool
}

func (v NullablePatchedNetworkLocationsRequest) Get() *PatchedNetworkLocationsRequest {
	return v.value
}

func (v *NullablePatchedNetworkLocationsRequest) Set(val *PatchedNetworkLocationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedNetworkLocationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedNetworkLocationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedNetworkLocationsRequest(val *PatchedNetworkLocationsRequest) *NullablePatchedNetworkLocationsRequest {
	return &NullablePatchedNetworkLocationsRequest{value: val, isSet: true}
}

func (v NullablePatchedNetworkLocationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedNetworkLocationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


