# SonarqubeIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Key** | **string** | SonarQube issue key | 
**Status** | **string** | SonarQube issue status | 
**Type** | **string** | SonarQube issue type | 

## Methods

### NewSonarqubeIssue

`func NewSonarqubeIssue(id int32, key string, status string, type_ string, ) *SonarqubeIssue`

NewSonarqubeIssue instantiates a new SonarqubeIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSonarqubeIssueWithDefaults

`func NewSonarqubeIssueWithDefaults() *SonarqubeIssue`

NewSonarqubeIssueWithDefaults instantiates a new SonarqubeIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SonarqubeIssue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SonarqubeIssue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SonarqubeIssue) SetId(v int32)`

SetId sets Id field to given value.


### GetKey

`func (o *SonarqubeIssue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SonarqubeIssue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SonarqubeIssue) SetKey(v string)`

SetKey sets Key field to given value.


### GetStatus

`func (o *SonarqubeIssue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SonarqubeIssue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SonarqubeIssue) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *SonarqubeIssue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SonarqubeIssue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SonarqubeIssue) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


