# StubFindingCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Test** | **int32** |  | 
**Title** | **string** |  | 
**Date** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Reporter** | **int32** |  | [readonly] 

## Methods

### NewStubFindingCreate

`func NewStubFindingCreate(id int32, test int32, title string, reporter int32, ) *StubFindingCreate`

NewStubFindingCreate instantiates a new StubFindingCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStubFindingCreateWithDefaults

`func NewStubFindingCreateWithDefaults() *StubFindingCreate`

NewStubFindingCreateWithDefaults instantiates a new StubFindingCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StubFindingCreate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StubFindingCreate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StubFindingCreate) SetId(v int32)`

SetId sets Id field to given value.


### GetTest

`func (o *StubFindingCreate) GetTest() int32`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *StubFindingCreate) GetTestOk() (*int32, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *StubFindingCreate) SetTest(v int32)`

SetTest sets Test field to given value.


### GetTitle

`func (o *StubFindingCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StubFindingCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StubFindingCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDate

`func (o *StubFindingCreate) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StubFindingCreate) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StubFindingCreate) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *StubFindingCreate) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSeverity

`func (o *StubFindingCreate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *StubFindingCreate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *StubFindingCreate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *StubFindingCreate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *StubFindingCreate) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *StubFindingCreate) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetDescription

`func (o *StubFindingCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StubFindingCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StubFindingCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StubFindingCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StubFindingCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StubFindingCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReporter

`func (o *StubFindingCreate) GetReporter() int32`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *StubFindingCreate) GetReporterOk() (*int32, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *StubFindingCreate) SetReporter(v int32)`

SetReporter sets Reporter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


