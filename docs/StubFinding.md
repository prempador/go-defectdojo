# StubFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Title** | **string** |  | 
**Date** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Test** | **int32** |  | [readonly] 
**Reporter** | **int32** |  | [readonly] 

## Methods

### NewStubFinding

`func NewStubFinding(id int32, title string, test int32, reporter int32, ) *StubFinding`

NewStubFinding instantiates a new StubFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStubFindingWithDefaults

`func NewStubFindingWithDefaults() *StubFinding`

NewStubFindingWithDefaults instantiates a new StubFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StubFinding) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StubFinding) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StubFinding) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *StubFinding) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StubFinding) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StubFinding) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDate

`func (o *StubFinding) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StubFinding) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StubFinding) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *StubFinding) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSeverity

`func (o *StubFinding) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StubFinding) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StubFinding) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *StubFinding) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *StubFinding) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *StubFinding) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetDescription

`func (o *StubFinding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StubFinding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StubFinding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StubFinding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StubFinding) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StubFinding) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTest

`func (o *StubFinding) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *StubFinding) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *StubFinding) SetTest(v int32)`

SetTest sets Test field to given value.


### GetReporter

`func (o *StubFinding) GetReporter() int32`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *StubFinding) GetReporterOk() (*int32, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *StubFinding) SetReporter(v int32)`

SetReporter sets Reporter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


