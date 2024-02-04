# ReportGenerate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutiveSummary** | [**NullableExecutiveSummary**](ExecutiveSummary.md) |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | [readonly] 
**Product** | [**Product**](Product.md) |  | [readonly] 
**Engagement** | [**Engagement**](Engagement.md) |  | [readonly] 
**ReportName** | **string** |  | 
**ReportInfo** | **string** |  | 
**Test** | [**Test**](Test.md) |  | [readonly] 
**Endpoint** | [**Endpoint**](Endpoint.md) |  | [readonly] 
**Endpoints** | [**[]Endpoint**](Endpoint.md) |  | [readonly] 
**Findings** | [**[]Finding**](Finding.md) |  | [readonly] 
**User** | [**UserStub**](UserStub.md) |  | [readonly] 
**TeamName** | **string** |  | 
**Title** | **string** |  | 
**UserId** | **int32** |  | 
**Host** | **string** |  | 
**FindingNotes** | Pointer to [**[]FindingToNotes**](FindingToNotes.md) |  | [optional] 

## Methods

### NewReportGenerate

`func NewReportGenerate(executiveSummary NullableExecutiveSummary, productType ProductType, product Product, engagement Engagement, reportName string, reportInfo string, test Test, endpoint Endpoint, endpoints []Endpoint, findings []Finding, user UserStub, teamName string, title string, userId int32, host string, ) *ReportGenerate`

NewReportGenerate instantiates a new ReportGenerate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportGenerateWithDefaults

`func NewReportGenerateWithDefaults() *ReportGenerate`

NewReportGenerateWithDefaults instantiates a new ReportGenerate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutiveSummary

`func (o *ReportGenerate) GetExecutiveSummary() ExecutiveSummary`

GetExecutiveSummary returns the ExecutiveSummary field if non-nil, zero value otherwise.

### GetExecutiveSummaryOk

`func (o *ReportGenerate) GetExecutiveSummaryOk() (*ExecutiveSummary, bool)`

GetExecutiveSummaryOk returns a tuple with the ExecutiveSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveSummary

`func (o *ReportGenerate) SetExecutiveSummary(v ExecutiveSummary)`

SetExecutiveSummary sets ExecutiveSummary field to given value.


### SetExecutiveSummaryNil

`func (o *ReportGenerate) SetExecutiveSummaryNil(b bool)`

 SetExecutiveSummaryNil sets the value for ExecutiveSummary to be an explicit nil

### UnsetExecutiveSummary
`func (o *ReportGenerate) UnsetExecutiveSummary()`

UnsetExecutiveSummary ensures that no value is present for ExecutiveSummary, not even an explicit nil
### GetProductType

`func (o *ReportGenerate) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ReportGenerate) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ReportGenerate) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetProduct

`func (o *ReportGenerate) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ReportGenerate) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ReportGenerate) SetProduct(v Product)`

SetProduct sets Product field to given value.


### GetEngagement

`func (o *ReportGenerate) GetEngagement() Engagement`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *ReportGenerate) GetEngagementOk() (*Engagement, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *ReportGenerate) SetEngagement(v Engagement)`

SetEngagement sets Engagement field to given value.


### GetReportName

`func (o *ReportGenerate) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *ReportGenerate) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *ReportGenerate) SetReportName(v string)`

SetReportName sets ReportName field to given value.


### GetReportInfo

`func (o *ReportGenerate) GetReportInfo() string`

GetReportInfo returns the ReportInfo field if non-nil, zero value otherwise.

### GetReportInfoOk

`func (o *ReportGenerate) GetReportInfoOk() (*string, bool)`

GetReportInfoOk returns a tuple with the ReportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInfo

`func (o *ReportGenerate) SetReportInfo(v string)`

SetReportInfo sets ReportInfo field to given value.


### GetTest

`func (o *ReportGenerate) GetTest() Test`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *ReportGenerate) GetTestOk() (*Test, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *ReportGenerate) SetTest(v Test)`

SetTest sets Test field to given value.


### GetEndpoint

`func (o *ReportGenerate) GetEndpoint() Endpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ReportGenerate) GetEndpointOk() (*Endpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ReportGenerate) SetEndpoint(v Endpoint)`

SetEndpoint sets Endpoint field to given value.


### GetEndpoints

`func (o *ReportGenerate) GetEndpoints() []Endpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ReportGenerate) GetEndpointsOk() (*[]Endpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ReportGenerate) SetEndpoints(v []Endpoint)`

SetEndpoints sets Endpoints field to given value.


### GetFindings

`func (o *ReportGenerate) GetFindings() []Finding`

GetFindings returns the Findings field if non-nil, zero value otherwise.

### GetFindingsOk

`func (o *ReportGenerate) GetFindingsOk() (*[]Finding, bool)`

GetFindingsOk returns a tuple with the Findings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindings

`func (o *ReportGenerate) SetFindings(v []Finding)`

SetFindings sets Findings field to given value.


### GetUser

`func (o *ReportGenerate) GetUser() UserStub`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ReportGenerate) GetUserOk() (*UserStub, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ReportGenerate) SetUser(v UserStub)`

SetUser sets User field to given value.


### GetTeamName

`func (o *ReportGenerate) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ReportGenerate) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ReportGenerate) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetTitle

`func (o *ReportGenerate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportGenerate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportGenerate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUserId

`func (o *ReportGenerate) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ReportGenerate) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ReportGenerate) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetHost

`func (o *ReportGenerate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ReportGenerate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ReportGenerate) SetHost(v string)`

SetHost sets Host field to given value.


### GetFindingNotes

`func (o *ReportGenerate) GetFindingNotes() []FindingToNotes`

GetFindingNotes returns the FindingNotes field if non-nil, zero value otherwise.

### GetFindingNotesOk

`func (o *ReportGenerate) GetFindingNotesOk() (*[]FindingToNotes, bool)`

GetFindingNotesOk returns a tuple with the FindingNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFindingNotes

`func (o *ReportGenerate) SetFindingNotes(v []FindingToNotes)`

SetFindingNotes sets FindingNotes field to given value.

### HasFindingNotes

`func (o *ReportGenerate) HasFindingNotes() bool`

HasFindingNotes returns a boolean if a field has been set.

### SetFindingNotesNil

`func (o *ReportGenerate) SetFindingNotesNil(b bool)`

 SetFindingNotesNil sets the value for FindingNotes to be an explicit nil

### UnsetFindingNotes
`func (o *ReportGenerate) UnsetFindingNotes()`

UnsetFindingNotes ensures that no value is present for FindingNotes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


