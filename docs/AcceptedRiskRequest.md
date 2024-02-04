# AcceptedRiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VulnerabilityId** | **string** | An id of a vulnerability in a security advisory associated with this finding. Can be a Common Vulnerabilities and Exposure (CVE) or from other sources. | 
**Justification** | **string** | Justification for accepting findings with this vulnerability id | 
**AcceptedBy** | **string** | Name or email of person who accepts the risk | 

## Methods

### NewAcceptedRiskRequest

`func NewAcceptedRiskRequest(vulnerabilityId string, justification string, acceptedBy string, ) *AcceptedRiskRequest`

NewAcceptedRiskRequest instantiates a new AcceptedRiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedRiskRequestWithDefaults

`func NewAcceptedRiskRequestWithDefaults() *AcceptedRiskRequest`

NewAcceptedRiskRequestWithDefaults instantiates a new AcceptedRiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnerabilityId

`func (o *AcceptedRiskRequest) GetVulnerabilityId() string`

GetVulnerabilityId returns the VulnerabilityId field if non-nil, zero value otherwise.

### GetVulnerabilityIdOk

`func (o *AcceptedRiskRequest) GetVulnerabilityIdOk() (*string, bool)`

GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityId

`func (o *AcceptedRiskRequest) SetVulnerabilityId(v string)`

SetVulnerabilityId sets VulnerabilityId field to given value.


### GetJustification

`func (o *AcceptedRiskRequest) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *AcceptedRiskRequest) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *AcceptedRiskRequest) SetJustification(v string)`

SetJustification sets Justification field to given value.


### GetAcceptedBy

`func (o *AcceptedRiskRequest) GetAcceptedBy() string`

GetAcceptedBy returns the AcceptedBy field if non-nil, zero value otherwise.

### GetAcceptedByOk

`func (o *AcceptedRiskRequest) GetAcceptedByOk() (*string, bool)`

GetAcceptedByOk returns a tuple with the AcceptedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedBy

`func (o *AcceptedRiskRequest) SetAcceptedBy(v string)`

SetAcceptedBy sets AcceptedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


