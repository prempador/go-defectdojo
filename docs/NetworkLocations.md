# NetworkLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Location** | **string** | Location of network testing: Examples: VPN, Internet or Internal. | 

## Methods

### NewNetworkLocations

`func NewNetworkLocations(id int32, location string, ) *NetworkLocations`

NewNetworkLocations instantiates a new NetworkLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLocationsWithDefaults

`func NewNetworkLocationsWithDefaults() *NetworkLocations`

NewNetworkLocationsWithDefaults instantiates a new NetworkLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkLocations) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLocations) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLocations) SetId(v int32)`

SetId sets Id field to given value.


### GetLocation

`func (o *NetworkLocations) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkLocations) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkLocations) SetLocation(v string)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


