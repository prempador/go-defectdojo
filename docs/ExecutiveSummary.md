# ExecutiveSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngagementName** | **string** |  | 
**EngagementTargetStart** | **string** |  | 
**EngagementTargetEnd** | **string** |  | 
**TestTypeName** | **string** |  | 
**TestTargetStart** | **time.Time** |  | 
**TestTargetEnd** | **time.Time** |  | 
**TestEnvironmentName** | **string** |  | 
**TestStrategyRef** | **string** |  | 
**TotalFindings** | **int32** |  | 

## Methods

### NewExecutiveSummary

`func NewExecutiveSummary(engagementName string, engagementTargetStart string, engagementTargetEnd string, testTypeName string, testTargetStart time.Time, testTargetEnd time.Time, testEnvironmentName string, testStrategyRef string, totalFindings int32, ) *ExecutiveSummary`

NewExecutiveSummary instantiates a new ExecutiveSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutiveSummaryWithDefaults

`func NewExecutiveSummaryWithDefaults() *ExecutiveSummary`

NewExecutiveSummaryWithDefaults instantiates a new ExecutiveSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagementName

`func (o *ExecutiveSummary) GetEngagementName() string`

GetEngagementName returns the EngagementName field if non-nil, zero value otherwise.

### GetEngagementNameOk

`func (o *ExecutiveSummary) GetEngagementNameOk() (*string, bool)`

GetEngagementNameOk returns a tuple with the EngagementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementName

`func (o *ExecutiveSummary) SetEngagementName(v string)`

SetEngagementName sets EngagementName field to given value.


### GetEngagementTargetStart

`func (o *ExecutiveSummary) GetEngagementTargetStart() string`

GetEngagementTargetStart returns the EngagementTargetStart field if non-nil, zero value otherwise.

### GetEngagementTargetStartOk

`func (o *ExecutiveSummary) GetEngagementTargetStartOk() (*string, bool)`

GetEngagementTargetStartOk returns a tuple with the EngagementTargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementTargetStart

`func (o *ExecutiveSummary) SetEngagementTargetStart(v string)`

SetEngagementTargetStart sets EngagementTargetStart field to given value.


### GetEngagementTargetEnd

`func (o *ExecutiveSummary) GetEngagementTargetEnd() string`

GetEngagementTargetEnd returns the EngagementTargetEnd field if non-nil, zero value otherwise.

### GetEngagementTargetEndOk

`func (o *ExecutiveSummary) GetEngagementTargetEndOk() (*string, bool)`

GetEngagementTargetEndOk returns a tuple with the EngagementTargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementTargetEnd

`func (o *ExecutiveSummary) SetEngagementTargetEnd(v string)`

SetEngagementTargetEnd sets EngagementTargetEnd field to given value.


### GetTestTypeName

`func (o *ExecutiveSummary) GetTestTypeName() string`

GetTestTypeName returns the TestTypeName field if non-nil, zero value otherwise.

### GetTestTypeNameOk

`func (o *ExecutiveSummary) GetTestTypeNameOk() (*string, bool)`

GetTestTypeNameOk returns a tuple with the TestTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTypeName

`func (o *ExecutiveSummary) SetTestTypeName(v string)`

SetTestTypeName sets TestTypeName field to given value.


### GetTestTargetStart

`func (o *ExecutiveSummary) GetTestTargetStart() time.Time`

GetTestTargetStart returns the TestTargetStart field if non-nil, zero value otherwise.

### GetTestTargetStartOk

`func (o *ExecutiveSummary) GetTestTargetStartOk() (*time.Time, bool)`

GetTestTargetStartOk returns a tuple with the TestTargetStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTargetStart

`func (o *ExecutiveSummary) SetTestTargetStart(v time.Time)`

SetTestTargetStart sets TestTargetStart field to given value.


### GetTestTargetEnd

`func (o *ExecutiveSummary) GetTestTargetEnd() time.Time`

GetTestTargetEnd returns the TestTargetEnd field if non-nil, zero value otherwise.

### GetTestTargetEndOk

`func (o *ExecutiveSummary) GetTestTargetEndOk() (*time.Time, bool)`

GetTestTargetEndOk returns a tuple with the TestTargetEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestTargetEnd

`func (o *ExecutiveSummary) SetTestTargetEnd(v time.Time)`

SetTestTargetEnd sets TestTargetEnd field to given value.


### GetTestEnvironmentName

`func (o *ExecutiveSummary) GetTestEnvironmentName() string`

GetTestEnvironmentName returns the TestEnvironmentName field if non-nil, zero value otherwise.

### GetTestEnvironmentNameOk

`func (o *ExecutiveSummary) GetTestEnvironmentNameOk() (*string, bool)`

GetTestEnvironmentNameOk returns a tuple with the TestEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestEnvironmentName

`func (o *ExecutiveSummary) SetTestEnvironmentName(v string)`

SetTestEnvironmentName sets TestEnvironmentName field to given value.


### GetTestStrategyRef

`func (o *ExecutiveSummary) GetTestStrategyRef() string`

GetTestStrategyRef returns the TestStrategyRef field if non-nil, zero value otherwise.

### GetTestStrategyRefOk

`func (o *ExecutiveSummary) GetTestStrategyRefOk() (*string, bool)`

GetTestStrategyRefOk returns a tuple with the TestStrategyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStrategyRef

`func (o *ExecutiveSummary) SetTestStrategyRef(v string)`

SetTestStrategyRef sets TestStrategyRef field to given value.


### GetTotalFindings

`func (o *ExecutiveSummary) GetTotalFindings() int32`

GetTotalFindings returns the TotalFindings field if non-nil, zero value otherwise.

### GetTotalFindingsOk

`func (o *ExecutiveSummary) GetTotalFindingsOk() (*int32, bool)`

GetTotalFindingsOk returns a tuple with the TotalFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFindings

`func (o *ExecutiveSummary) SetTotalFindings(v int32)`

SetTotalFindings sets TotalFindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


