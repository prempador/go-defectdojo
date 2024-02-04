# NetworkLocationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location of network testing: Examples: VPN, Internet or Internal. | 

## Methods

### NewNetworkLocationsRequest

`func NewNetworkLocationsRequest(location string, ) *NetworkLocationsRequest`

NewNetworkLocationsRequest instantiates a new NetworkLocationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLocationsRequestWithDefaults

`func NewNetworkLocationsRequestWithDefaults() *NetworkLocationsRequest`

NewNetworkLocationsRequestWithDefaults instantiates a new NetworkLocationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *NetworkLocationsRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkLocationsRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkLocationsRequest) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


