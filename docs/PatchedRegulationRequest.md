# PatchedRegulationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the regulation. | [optional] 
**Acronym** | Pointer to **string** | A shortened representation of the name. | [optional] 
**Category** | Pointer to **string** | The subject of the regulation.  * &#x60;privacy&#x60; - Privacy * &#x60;finance&#x60; - Finance * &#x60;education&#x60; - Education * &#x60;medical&#x60; - Medical * &#x60;corporate&#x60; - Corporate * &#x60;other&#x60; - Other | [optional] 
**Jurisdiction** | Pointer to **string** | The territory over which the regulation applies. | [optional] 
**Description** | Pointer to **string** | Information about the regulation&#39;s purpose. | [optional] 
**Reference** | Pointer to **string** | An external URL for more information. | [optional] 

## Methods

### NewPatchedRegulationRequest

`func NewPatchedRegulationRequest() *PatchedRegulationRequest`

NewPatchedRegulationRequest instantiates a new PatchedRegulationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRegulationRequestWithDefaults

`func NewPatchedRegulationRequestWithDefaults() *PatchedRegulationRequest`

NewPatchedRegulationRequestWithDefaults instantiates a new PatchedRegulationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedRegulationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedRegulationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedRegulationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedRegulationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAcronym

`func (o *PatchedRegulationRequest) GetAcronym() string`

GetAcronym returns the Acronym field if non-nil, zero value otherwise.

### GetAcronymOk

`func (o *PatchedRegulationRequest) GetAcronymOk() (*string, bool)`

GetAcronymOk returns a tuple with the Acronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcronym

`func (o *PatchedRegulationRequest) SetAcronym(v string)`

SetAcronym sets Acronym field to given value.

### HasAcronym

`func (o *PatchedRegulationRequest) HasAcronym() bool`

HasAcronym returns a boolean if a field has been set.

### GetCategory

`func (o *PatchedRegulationRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PatchedRegulationRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PatchedRegulationRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PatchedRegulationRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetJurisdiction

`func (o *PatchedRegulationRequest) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *PatchedRegulationRequest) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *PatchedRegulationRequest) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *PatchedRegulationRequest) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedRegulationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedRegulationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedRegulationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedRegulationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *PatchedRegulationRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PatchedRegulationRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PatchedRegulationRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PatchedRegulationRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


