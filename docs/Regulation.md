# Regulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Name** | **string** | The name of the regulation. | 
**Acronym** | **string** | A shortened representation of the name. | 
**Category** | **string** | The subject of the regulation.  * &#x60;privacy&#x60; - Privacy * &#x60;finance&#x60; - Finance * &#x60;education&#x60; - Education * &#x60;medical&#x60; - Medical * &#x60;corporate&#x60; - Corporate * &#x60;other&#x60; - Other | 
**Jurisdiction** | **string** | The territory over which the regulation applies. | 
**Description** | Pointer to **string** | Information about the regulation&#39;s purpose. | [optional] 
**Reference** | Pointer to **string** | An external URL for more information. | [optional] 

## Methods

### NewRegulation

`func NewRegulation(id int32, name string, acronym string, category string, jurisdiction string, ) *Regulation`

NewRegulation instantiates a new Regulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegulationWithDefaults

`func NewRegulationWithDefaults() *Regulation`

NewRegulationWithDefaults instantiates a new Regulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Regulation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Regulation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Regulation) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Regulation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Regulation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Regulation) SetName(v string)`

SetName sets Name field to given value.


### GetAcronym

`func (o *Regulation) GetAcronym() string`

GetAcronym returns the Acronym field if non-nil, zero value otherwise.

### GetAcronymOk

`func (o *Regulation) GetAcronymOk() (*string, bool)`

GetAcronymOk returns a tuple with the Acronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcronym

`func (o *Regulation) SetAcronym(v string)`

SetAcronym sets Acronym field to given value.


### GetCategory

`func (o *Regulation) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Regulation) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Regulation) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetJurisdiction

`func (o *Regulation) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *Regulation) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *Regulation) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.


### GetDescription

`func (o *Regulation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Regulation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Regulation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Regulation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *Regulation) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Regulation) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Regulation) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Regulation) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


