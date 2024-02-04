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

// checks if the ReportGenerateOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportGenerateOptionRequest{}

// ReportGenerateOptionRequest struct for ReportGenerateOptionRequest
type ReportGenerateOptionRequest struct {
	IncludeFindingNotes *bool `json:"include_finding_notes,omitempty"`
	IncludeFindingImages *bool `json:"include_finding_images,omitempty"`
	IncludeExecutiveSummary *bool `json:"include_executive_summary,omitempty"`
	IncludeTableOfContents *bool `json:"include_table_of_contents,omitempty"`
}

// NewReportGenerateOptionRequest instantiates a new ReportGenerateOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportGenerateOptionRequest() *ReportGenerateOptionRequest {
	this := ReportGenerateOptionRequest{}
	var includeFindingNotes bool = false
	this.IncludeFindingNotes = &includeFindingNotes
	var includeFindingImages bool = false
	this.IncludeFindingImages = &includeFindingImages
	var includeExecutiveSummary bool = false
	this.IncludeExecutiveSummary = &includeExecutiveSummary
	var includeTableOfContents bool = false
	this.IncludeTableOfContents = &includeTableOfContents
	return &this
}

// NewReportGenerateOptionRequestWithDefaults instantiates a new ReportGenerateOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportGenerateOptionRequestWithDefaults() *ReportGenerateOptionRequest {
	this := ReportGenerateOptionRequest{}
	var includeFindingNotes bool = false
	this.IncludeFindingNotes = &includeFindingNotes
	var includeFindingImages bool = false
	this.IncludeFindingImages = &includeFindingImages
	var includeExecutiveSummary bool = false
	this.IncludeExecutiveSummary = &includeExecutiveSummary
	var includeTableOfContents bool = false
	this.IncludeTableOfContents = &includeTableOfContents
	return &this
}

// GetIncludeFindingNotes returns the IncludeFindingNotes field value if set, zero value otherwise.
func (o *ReportGenerateOptionRequest) GetIncludeFindingNotes() bool {
	if o == nil || IsNil(o.IncludeFindingNotes) {
		var ret bool
		return ret
	}
	return *o.IncludeFindingNotes
}

// GetIncludeFindingNotesOk returns a tuple with the IncludeFindingNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerateOptionRequest) GetIncludeFindingNotesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeFindingNotes) {
		return nil, false
	}
	return o.IncludeFindingNotes, true
}

// HasIncludeFindingNotes returns a boolean if a field has been set.
func (o *ReportGenerateOptionRequest) HasIncludeFindingNotes() bool {
	if o != nil && !IsNil(o.IncludeFindingNotes) {
		return true
	}

	return false
}

// SetIncludeFindingNotes gets a reference to the given bool and assigns it to the IncludeFindingNotes field.
func (o *ReportGenerateOptionRequest) SetIncludeFindingNotes(v bool) {
	o.IncludeFindingNotes = &v
}

// GetIncludeFindingImages returns the IncludeFindingImages field value if set, zero value otherwise.
func (o *ReportGenerateOptionRequest) GetIncludeFindingImages() bool {
	if o == nil || IsNil(o.IncludeFindingImages) {
		var ret bool
		return ret
	}
	return *o.IncludeFindingImages
}

// GetIncludeFindingImagesOk returns a tuple with the IncludeFindingImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerateOptionRequest) GetIncludeFindingImagesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeFindingImages) {
		return nil, false
	}
	return o.IncludeFindingImages, true
}

// HasIncludeFindingImages returns a boolean if a field has been set.
func (o *ReportGenerateOptionRequest) HasIncludeFindingImages() bool {
	if o != nil && !IsNil(o.IncludeFindingImages) {
		return true
	}

	return false
}

// SetIncludeFindingImages gets a reference to the given bool and assigns it to the IncludeFindingImages field.
func (o *ReportGenerateOptionRequest) SetIncludeFindingImages(v bool) {
	o.IncludeFindingImages = &v
}

// GetIncludeExecutiveSummary returns the IncludeExecutiveSummary field value if set, zero value otherwise.
func (o *ReportGenerateOptionRequest) GetIncludeExecutiveSummary() bool {
	if o == nil || IsNil(o.IncludeExecutiveSummary) {
		var ret bool
		return ret
	}
	return *o.IncludeExecutiveSummary
}

// GetIncludeExecutiveSummaryOk returns a tuple with the IncludeExecutiveSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerateOptionRequest) GetIncludeExecutiveSummaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExecutiveSummary) {
		return nil, false
	}
	return o.IncludeExecutiveSummary, true
}

// HasIncludeExecutiveSummary returns a boolean if a field has been set.
func (o *ReportGenerateOptionRequest) HasIncludeExecutiveSummary() bool {
	if o != nil && !IsNil(o.IncludeExecutiveSummary) {
		return true
	}

	return false
}

// SetIncludeExecutiveSummary gets a reference to the given bool and assigns it to the IncludeExecutiveSummary field.
func (o *ReportGenerateOptionRequest) SetIncludeExecutiveSummary(v bool) {
	o.IncludeExecutiveSummary = &v
}

// GetIncludeTableOfContents returns the IncludeTableOfContents field value if set, zero value otherwise.
func (o *ReportGenerateOptionRequest) GetIncludeTableOfContents() bool {
	if o == nil || IsNil(o.IncludeTableOfContents) {
		var ret bool
		return ret
	}
	return *o.IncludeTableOfContents
}

// GetIncludeTableOfContentsOk returns a tuple with the IncludeTableOfContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGenerateOptionRequest) GetIncludeTableOfContentsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeTableOfContents) {
		return nil, false
	}
	return o.IncludeTableOfContents, true
}

// HasIncludeTableOfContents returns a boolean if a field has been set.
func (o *ReportGenerateOptionRequest) HasIncludeTableOfContents() bool {
	if o != nil && !IsNil(o.IncludeTableOfContents) {
		return true
	}

	return false
}

// SetIncludeTableOfContents gets a reference to the given bool and assigns it to the IncludeTableOfContents field.
func (o *ReportGenerateOptionRequest) SetIncludeTableOfContents(v bool) {
	o.IncludeTableOfContents = &v
}

func (o ReportGenerateOptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportGenerateOptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IncludeFindingNotes) {
		toSerialize["include_finding_notes"] = o.IncludeFindingNotes
	}
	if !IsNil(o.IncludeFindingImages) {
		toSerialize["include_finding_images"] = o.IncludeFindingImages
	}
	if !IsNil(o.IncludeExecutiveSummary) {
		toSerialize["include_executive_summary"] = o.IncludeExecutiveSummary
	}
	if !IsNil(o.IncludeTableOfContents) {
		toSerialize["include_table_of_contents"] = o.IncludeTableOfContents
	}
	return toSerialize, nil
}

type NullableReportGenerateOptionRequest struct {
	value *ReportGenerateOptionRequest
	isSet bool
}

func (v NullableReportGenerateOptionRequest) Get() *ReportGenerateOptionRequest {
	return v.value
}

func (v *NullableReportGenerateOptionRequest) Set(val *ReportGenerateOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportGenerateOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportGenerateOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportGenerateOptionRequest(val *ReportGenerateOptionRequest) *NullableReportGenerateOptionRequest {
	return &NullableReportGenerateOptionRequest{value: val, isSet: true}
}

func (v NullableReportGenerateOptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportGenerateOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

