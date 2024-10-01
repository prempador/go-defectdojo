# JIRAProjectPrefetch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Engagement** | Pointer to [**map[string]FindingEngagement**](FindingEngagement.md) |  | [optional] [readonly] 
**JiraInstance** | Pointer to [**map[string]JIRAInstance**](JIRAInstance.md) |  | [optional] [readonly] 
**Product** | Pointer to [**map[string]Product**](Product.md) |  | [optional] [readonly] 

## Methods

### NewJIRAProjectPrefetch

`func NewJIRAProjectPrefetch() *JIRAProjectPrefetch`

NewJIRAProjectPrefetch instantiates a new JIRAProjectPrefetch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJIRAProjectPrefetchWithDefaults

`func NewJIRAProjectPrefetchWithDefaults() *JIRAProjectPrefetch`

NewJIRAProjectPrefetchWithDefaults instantiates a new JIRAProjectPrefetch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagement

`func (o *JIRAProjectPrefetch) GetEngagement() map[string]FindingEngagement`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *JIRAProjectPrefetch) GetEngagementOk() (*map[string]FindingEngagement, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *JIRAProjectPrefetch) SetEngagement(v map[string]FindingEngagement)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *JIRAProjectPrefetch) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetJiraInstance

`func (o *JIRAProjectPrefetch) GetJiraInstance() map[string]JIRAInstance`

GetJiraInstance returns the JiraInstance field if non-nil, zero value otherwise.

### GetJiraInstanceOk

`func (o *JIRAProjectPrefetch) GetJiraInstanceOk() (*map[string]JIRAInstance, bool)`

GetJiraInstanceOk returns a tuple with the JiraInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraInstance

`func (o *JIRAProjectPrefetch) SetJiraInstance(v map[string]JIRAInstance)`

SetJiraInstance sets JiraInstance field to given value.

### HasJiraInstance

`func (o *JIRAProjectPrefetch) HasJiraInstance() bool`

HasJiraInstance returns a boolean if a field has been set.

### GetProduct

`func (o *JIRAProjectPrefetch) GetProduct() map[string]Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *JIRAProjectPrefetch) GetProductOk() (*map[string]Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *JIRAProjectPrefetch) SetProduct(v map[string]Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *JIRAProjectPrefetch) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


