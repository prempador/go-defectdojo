# \EndpointStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointStatusCreate**](EndpointStatusAPI.md#EndpointStatusCreate) | **Post** /api/v2/endpoint_status/ | 
[**EndpointStatusDeletePreviewList**](EndpointStatusAPI.md#EndpointStatusDeletePreviewList) | **Get** /api/v2/endpoint_status/{id}/delete_preview/ | 
[**EndpointStatusDestroy**](EndpointStatusAPI.md#EndpointStatusDestroy) | **Delete** /api/v2/endpoint_status/{id}/ | 
[**EndpointStatusList**](EndpointStatusAPI.md#EndpointStatusList) | **Get** /api/v2/endpoint_status/ | 
[**EndpointStatusPartialUpdate**](EndpointStatusAPI.md#EndpointStatusPartialUpdate) | **Patch** /api/v2/endpoint_status/{id}/ | 
[**EndpointStatusRetrieve**](EndpointStatusAPI.md#EndpointStatusRetrieve) | **Get** /api/v2/endpoint_status/{id}/ | 
[**EndpointStatusUpdate**](EndpointStatusAPI.md#EndpointStatusUpdate) | **Put** /api/v2/endpoint_status/{id}/ | 



## EndpointStatusCreate

> EndpointStatus EndpointStatusCreate(ctx).EndpointStatusRequest(endpointStatusRequest).Execute()



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
	endpointStatusRequest := *openapiclient.NewEndpointStatusRequest(int32(123), int32(123)) // EndpointStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointStatusAPI.EndpointStatusCreate(context.Background()).EndpointStatusRequest(endpointStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointStatusCreate`: EndpointStatus
	fmt.Fprintf(os.Stdout, "Response from `EndpointStatusAPI.EndpointStatusCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointStatusRequest** | [**EndpointStatusRequest**](EndpointStatusRequest.md) |  | 

### Return type

[**EndpointStatus**](EndpointStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointStatusDeletePreviewList

> PaginatedDeletePreviewList EndpointStatusDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint_ status.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointStatusAPI.EndpointStatusDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointStatusDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `EndpointStatusAPI.EndpointStatusDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint_ status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusDeletePreviewListRequest struct via the builder pattern


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


## EndpointStatusDestroy

> EndpointStatusDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint_ status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointStatusAPI.EndpointStatusDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint_ status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusDestroyRequest struct via the builder pattern


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


## EndpointStatusList

> PaginatedEndpointStatusList EndpointStatusList(ctx).Endpoint(endpoint).FalsePositive(falsePositive).Finding(finding).Limit(limit).Mitigated(mitigated).MitigatedBy(mitigatedBy).Offset(offset).OutOfScope(outOfScope).RiskAccepted(riskAccepted).Execute()



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
	endpoint := int32(56) // int32 |  (optional)
	falsePositive := true // bool |  (optional)
	finding := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	mitigated := true // bool |  (optional)
	mitigatedBy := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	outOfScope := true // bool |  (optional)
	riskAccepted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointStatusAPI.EndpointStatusList(context.Background()).Endpoint(endpoint).FalsePositive(falsePositive).Finding(finding).Limit(limit).Mitigated(mitigated).MitigatedBy(mitigatedBy).Offset(offset).OutOfScope(outOfScope).RiskAccepted(riskAccepted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointStatusList`: PaginatedEndpointStatusList
	fmt.Fprintf(os.Stdout, "Response from `EndpointStatusAPI.EndpointStatusList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **int32** |  | 
 **falsePositive** | **bool** |  | 
 **finding** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **mitigated** | **bool** |  | 
 **mitigatedBy** | **int32** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **outOfScope** | **bool** |  | 
 **riskAccepted** | **bool** |  | 

### Return type

[**PaginatedEndpointStatusList**](PaginatedEndpointStatusList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointStatusPartialUpdate

> EndpointStatus EndpointStatusPartialUpdate(ctx, id).PatchedEndpointStatusRequest(patchedEndpointStatusRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint_ status.
	patchedEndpointStatusRequest := *openapiclient.NewPatchedEndpointStatusRequest() // PatchedEndpointStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointStatusAPI.EndpointStatusPartialUpdate(context.Background(), id).PatchedEndpointStatusRequest(patchedEndpointStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointStatusPartialUpdate`: EndpointStatus
	fmt.Fprintf(os.Stdout, "Response from `EndpointStatusAPI.EndpointStatusPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint_ status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEndpointStatusRequest** | [**PatchedEndpointStatusRequest**](PatchedEndpointStatusRequest.md) |  | 

### Return type

[**EndpointStatus**](EndpointStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointStatusRetrieve

> EndpointStatus EndpointStatusRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint_ status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointStatusAPI.EndpointStatusRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointStatusRetrieve`: EndpointStatus
	fmt.Fprintf(os.Stdout, "Response from `EndpointStatusAPI.EndpointStatusRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint_ status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EndpointStatus**](EndpointStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointStatusUpdate

> EndpointStatus EndpointStatusUpdate(ctx, id).EndpointStatusRequest(endpointStatusRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint_ status.
	endpointStatusRequest := *openapiclient.NewEndpointStatusRequest(int32(123), int32(123)) // EndpointStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointStatusAPI.EndpointStatusUpdate(context.Background(), id).EndpointStatusRequest(endpointStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointStatusAPI.EndpointStatusUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointStatusUpdate`: EndpointStatus
	fmt.Fprintf(os.Stdout, "Response from `EndpointStatusAPI.EndpointStatusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint_ status. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointStatusRequest** | [**EndpointStatusRequest**](EndpointStatusRequest.md) |  | 

### Return type

[**EndpointStatus**](EndpointStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

