# \RequestResponsePairsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestResponsePairsCreate**](RequestResponsePairsAPI.md#RequestResponsePairsCreate) | **Post** /api/v2/request_response_pairs/ | 
[**RequestResponsePairsDeletePreviewList**](RequestResponsePairsAPI.md#RequestResponsePairsDeletePreviewList) | **Get** /api/v2/request_response_pairs/{id}/delete_preview/ | 
[**RequestResponsePairsDestroy**](RequestResponsePairsAPI.md#RequestResponsePairsDestroy) | **Delete** /api/v2/request_response_pairs/{id}/ | 
[**RequestResponsePairsList**](RequestResponsePairsAPI.md#RequestResponsePairsList) | **Get** /api/v2/request_response_pairs/ | 
[**RequestResponsePairsPartialUpdate**](RequestResponsePairsAPI.md#RequestResponsePairsPartialUpdate) | **Patch** /api/v2/request_response_pairs/{id}/ | 
[**RequestResponsePairsRetrieve**](RequestResponsePairsAPI.md#RequestResponsePairsRetrieve) | **Get** /api/v2/request_response_pairs/{id}/ | 
[**RequestResponsePairsUpdate**](RequestResponsePairsAPI.md#RequestResponsePairsUpdate) | **Put** /api/v2/request_response_pairs/{id}/ | 



## RequestResponsePairsCreate

> BurpRawRequestResponseMulti RequestResponsePairsCreate(ctx).BurpRawRequestResponseMultiRequest(burpRawRequestResponseMultiRequest).Execute()



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
	burpRawRequestResponseMultiRequest := *openapiclient.NewBurpRawRequestResponseMultiRequest("BurpRequestBase64_example", "BurpResponseBase64_example") // BurpRawRequestResponseMultiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsCreate(context.Background()).BurpRawRequestResponseMultiRequest(burpRawRequestResponseMultiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResponsePairsCreate`: BurpRawRequestResponseMulti
	fmt.Fprintf(os.Stdout, "Response from `RequestResponsePairsAPI.RequestResponsePairsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **burpRawRequestResponseMultiRequest** | [**BurpRawRequestResponseMultiRequest**](BurpRawRequestResponseMultiRequest.md) |  | 

### Return type

[**BurpRawRequestResponseMulti**](BurpRawRequestResponseMulti.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestResponsePairsDeletePreviewList

> PaginatedDeletePreviewList RequestResponsePairsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this burp raw request response.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResponsePairsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `RequestResponsePairsAPI.RequestResponsePairsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this burp raw request response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsDeletePreviewListRequest struct via the builder pattern


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


## RequestResponsePairsDestroy

> RequestResponsePairsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this burp raw request response.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this burp raw request response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsDestroyRequest struct via the builder pattern


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


## RequestResponsePairsList

> PaginatedBurpRawRequestResponseMultiList RequestResponsePairsList(ctx).Finding(finding).Limit(limit).Offset(offset).Execute()



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
	finding := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsList(context.Background()).Finding(finding).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResponsePairsList`: PaginatedBurpRawRequestResponseMultiList
	fmt.Fprintf(os.Stdout, "Response from `RequestResponsePairsAPI.RequestResponsePairsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **finding** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedBurpRawRequestResponseMultiList**](PaginatedBurpRawRequestResponseMultiList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestResponsePairsPartialUpdate

> BurpRawRequestResponseMulti RequestResponsePairsPartialUpdate(ctx, id).PatchedBurpRawRequestResponseMultiRequest(patchedBurpRawRequestResponseMultiRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this burp raw request response.
	patchedBurpRawRequestResponseMultiRequest := *openapiclient.NewPatchedBurpRawRequestResponseMultiRequest() // PatchedBurpRawRequestResponseMultiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsPartialUpdate(context.Background(), id).PatchedBurpRawRequestResponseMultiRequest(patchedBurpRawRequestResponseMultiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResponsePairsPartialUpdate`: BurpRawRequestResponseMulti
	fmt.Fprintf(os.Stdout, "Response from `RequestResponsePairsAPI.RequestResponsePairsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this burp raw request response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedBurpRawRequestResponseMultiRequest** | [**PatchedBurpRawRequestResponseMultiRequest**](PatchedBurpRawRequestResponseMultiRequest.md) |  | 

### Return type

[**BurpRawRequestResponseMulti**](BurpRawRequestResponseMulti.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestResponsePairsRetrieve

> BurpRawRequestResponseMulti RequestResponsePairsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this burp raw request response.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResponsePairsRetrieve`: BurpRawRequestResponseMulti
	fmt.Fprintf(os.Stdout, "Response from `RequestResponsePairsAPI.RequestResponsePairsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this burp raw request response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BurpRawRequestResponseMulti**](BurpRawRequestResponseMulti.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestResponsePairsUpdate

> BurpRawRequestResponseMulti RequestResponsePairsUpdate(ctx, id).BurpRawRequestResponseMultiRequest(burpRawRequestResponseMultiRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this burp raw request response.
	burpRawRequestResponseMultiRequest := *openapiclient.NewBurpRawRequestResponseMultiRequest("BurpRequestBase64_example", "BurpResponseBase64_example") // BurpRawRequestResponseMultiRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestResponsePairsAPI.RequestResponsePairsUpdate(context.Background(), id).BurpRawRequestResponseMultiRequest(burpRawRequestResponseMultiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestResponsePairsAPI.RequestResponsePairsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestResponsePairsUpdate`: BurpRawRequestResponseMulti
	fmt.Fprintf(os.Stdout, "Response from `RequestResponsePairsAPI.RequestResponsePairsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this burp raw request response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestResponsePairsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **burpRawRequestResponseMultiRequest** | [**BurpRawRequestResponseMultiRequest**](BurpRawRequestResponseMultiRequest.md) |  | 

### Return type

[**BurpRawRequestResponseMulti**](BurpRawRequestResponseMulti.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

