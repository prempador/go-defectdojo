# \NetworkLocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkLocationsCreate**](NetworkLocationsAPI.md#NetworkLocationsCreate) | **Post** /api/v2/network_locations/ | 
[**NetworkLocationsDeletePreviewList**](NetworkLocationsAPI.md#NetworkLocationsDeletePreviewList) | **Get** /api/v2/network_locations/{id}/delete_preview/ | 
[**NetworkLocationsDestroy**](NetworkLocationsAPI.md#NetworkLocationsDestroy) | **Delete** /api/v2/network_locations/{id}/ | 
[**NetworkLocationsList**](NetworkLocationsAPI.md#NetworkLocationsList) | **Get** /api/v2/network_locations/ | 
[**NetworkLocationsPartialUpdate**](NetworkLocationsAPI.md#NetworkLocationsPartialUpdate) | **Patch** /api/v2/network_locations/{id}/ | 
[**NetworkLocationsRetrieve**](NetworkLocationsAPI.md#NetworkLocationsRetrieve) | **Get** /api/v2/network_locations/{id}/ | 
[**NetworkLocationsUpdate**](NetworkLocationsAPI.md#NetworkLocationsUpdate) | **Put** /api/v2/network_locations/{id}/ | 



## NetworkLocationsCreate

> NetworkLocations NetworkLocationsCreate(ctx).NetworkLocationsRequest(networkLocationsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	networkLocationsRequest := *openapiclient.NewNetworkLocationsRequest("Location_example") // NetworkLocationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLocationsAPI.NetworkLocationsCreate(context.Background()).NetworkLocationsRequest(networkLocationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkLocationsCreate`: NetworkLocations
	fmt.Fprintf(os.Stdout, "Response from `NetworkLocationsAPI.NetworkLocationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkLocationsRequest** | [**NetworkLocationsRequest**](NetworkLocationsRequest.md) |  | 

### Return type

[**NetworkLocations**](NetworkLocations.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkLocationsDeletePreviewList

> PaginatedDeletePreviewList NetworkLocationsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this network_ locations.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLocationsAPI.NetworkLocationsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkLocationsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `NetworkLocationsAPI.NetworkLocationsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this network_ locations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsDeletePreviewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDeletePreviewList**](PaginatedDeletePreviewList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkLocationsDestroy

> NetworkLocationsDestroy(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this network_ locations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkLocationsAPI.NetworkLocationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this network_ locations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkLocationsList

> PaginatedNetworkLocationsList NetworkLocationsList(ctx).Id(id).Limit(limit).Location(location).Offset(offset).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	location := "location_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLocationsAPI.NetworkLocationsList(context.Background()).Id(id).Limit(limit).Location(location).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkLocationsList`: PaginatedNetworkLocationsList
	fmt.Fprintf(os.Stdout, "Response from `NetworkLocationsAPI.NetworkLocationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **location** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedNetworkLocationsList**](PaginatedNetworkLocationsList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkLocationsPartialUpdate

> NetworkLocations NetworkLocationsPartialUpdate(ctx, id).PatchedNetworkLocationsRequest(patchedNetworkLocationsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this network_ locations.
	patchedNetworkLocationsRequest := *openapiclient.NewPatchedNetworkLocationsRequest() // PatchedNetworkLocationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLocationsAPI.NetworkLocationsPartialUpdate(context.Background(), id).PatchedNetworkLocationsRequest(patchedNetworkLocationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkLocationsPartialUpdate`: NetworkLocations
	fmt.Fprintf(os.Stdout, "Response from `NetworkLocationsAPI.NetworkLocationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this network_ locations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNetworkLocationsRequest** | [**PatchedNetworkLocationsRequest**](PatchedNetworkLocationsRequest.md) |  | 

### Return type

[**NetworkLocations**](NetworkLocations.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkLocationsRetrieve

> NetworkLocations NetworkLocationsRetrieve(ctx, id).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this network_ locations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLocationsAPI.NetworkLocationsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkLocationsRetrieve`: NetworkLocations
	fmt.Fprintf(os.Stdout, "Response from `NetworkLocationsAPI.NetworkLocationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this network_ locations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkLocations**](NetworkLocations.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkLocationsUpdate

> NetworkLocations NetworkLocationsUpdate(ctx, id).NetworkLocationsRequest(networkLocationsRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	id := int32(56) // int32 | A unique integer value identifying this network_ locations.
	networkLocationsRequest := *openapiclient.NewNetworkLocationsRequest("Location_example") // NetworkLocationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkLocationsAPI.NetworkLocationsUpdate(context.Background(), id).NetworkLocationsRequest(networkLocationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkLocationsAPI.NetworkLocationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkLocationsUpdate`: NetworkLocations
	fmt.Fprintf(os.Stdout, "Response from `NetworkLocationsAPI.NetworkLocationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this network_ locations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkLocationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkLocationsRequest** | [**NetworkLocationsRequest**](NetworkLocationsRequest.md) |  | 

### Return type

[**NetworkLocations**](NetworkLocations.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

