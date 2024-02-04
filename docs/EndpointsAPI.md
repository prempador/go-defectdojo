# \EndpointsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointsCreate**](EndpointsAPI.md#EndpointsCreate) | **Post** /api/v2/endpoints/ | 
[**EndpointsDeletePreviewList**](EndpointsAPI.md#EndpointsDeletePreviewList) | **Get** /api/v2/endpoints/{id}/delete_preview/ | 
[**EndpointsDestroy**](EndpointsAPI.md#EndpointsDestroy) | **Delete** /api/v2/endpoints/{id}/ | 
[**EndpointsGenerateReportCreate**](EndpointsAPI.md#EndpointsGenerateReportCreate) | **Post** /api/v2/endpoints/{id}/generate_report/ | 
[**EndpointsList**](EndpointsAPI.md#EndpointsList) | **Get** /api/v2/endpoints/ | 
[**EndpointsPartialUpdate**](EndpointsAPI.md#EndpointsPartialUpdate) | **Patch** /api/v2/endpoints/{id}/ | 
[**EndpointsRetrieve**](EndpointsAPI.md#EndpointsRetrieve) | **Get** /api/v2/endpoints/{id}/ | 
[**EndpointsUpdate**](EndpointsAPI.md#EndpointsUpdate) | **Put** /api/v2/endpoints/{id}/ | 



## EndpointsCreate

> Endpoint EndpointsCreate(ctx).EndpointRequest(endpointRequest).Execute()



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
	endpointRequest := *openapiclient.NewEndpointRequest() // EndpointRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsCreate(context.Background()).EndpointRequest(endpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsCreate`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointRequest** | [**EndpointRequest**](EndpointRequest.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointsDeletePreviewList

> PaginatedDeletePreviewList EndpointsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsDeletePreviewListRequest struct via the builder pattern


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


## EndpointsDestroy

> EndpointsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EndpointsAPI.EndpointsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsDestroyRequest struct via the builder pattern


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


## EndpointsGenerateReportCreate

> ReportGenerate EndpointsGenerateReportCreate(ctx, id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint.
	reportGenerateOptionRequest := *openapiclient.NewReportGenerateOptionRequest() // ReportGenerateOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsGenerateReportCreate(context.Background(), id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsGenerateReportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsGenerateReportCreate`: ReportGenerate
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsGenerateReportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsGenerateReportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportGenerateOptionRequest** | [**ReportGenerateOptionRequest**](ReportGenerateOptionRequest.md) |  | 

### Return type

[**ReportGenerate**](ReportGenerate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointsList

> PaginatedEndpointList EndpointsList(ctx).Fragment(fragment).HasTags(hasTags).Host(host).Id(id).Limit(limit).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).Path(path).Port(port).Product(product).Protocol(protocol).Query(query).Tag(tag).Tags(tags).Userinfo(userinfo).Execute()



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
	fragment := "fragment_example" // string |  (optional)
	hasTags := true // bool | Has tags (optional)
	host := "host_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	o := []string{"O_example"} // []string | Ordering  * `host` - Host * `-host` - Host (descending) * `product` - Product * `-product` - Product (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	path := "path_example" // string |  (optional)
	port := int32(56) // int32 |  (optional)
	product := int32(56) // int32 |  (optional)
	protocol := "protocol_example" // string |  (optional)
	query := "query_example" // string |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	userinfo := "userinfo_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsList(context.Background()).Fragment(fragment).HasTags(hasTags).Host(host).Id(id).Limit(limit).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).Path(path).Port(port).Product(product).Protocol(protocol).Query(query).Tag(tag).Tags(tags).Userinfo(userinfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsList`: PaginatedEndpointList
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fragment** | **string** |  | 
 **hasTags** | **bool** | Has tags | 
 **host** | **string** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **o** | **[]string** | Ordering  * &#x60;host&#x60; - Host * &#x60;-host&#x60; - Host (descending) * &#x60;product&#x60; - Product * &#x60;-product&#x60; - Product (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **path** | **string** |  | 
 **port** | **int32** |  | 
 **product** | **int32** |  | 
 **protocol** | **string** |  | 
 **query** | **string** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **userinfo** | **string** |  | 

### Return type

[**PaginatedEndpointList**](PaginatedEndpointList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointsPartialUpdate

> Endpoint EndpointsPartialUpdate(ctx, id).PatchedEndpointRequest(patchedEndpointRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint.
	patchedEndpointRequest := *openapiclient.NewPatchedEndpointRequest() // PatchedEndpointRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsPartialUpdate(context.Background(), id).PatchedEndpointRequest(patchedEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsPartialUpdate`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEndpointRequest** | [**PatchedEndpointRequest**](PatchedEndpointRequest.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointsRetrieve

> Endpoint EndpointsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsRetrieve`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndpointsUpdate

> Endpoint EndpointsUpdate(ctx, id).EndpointRequest(endpointRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this endpoint.
	endpointRequest := *openapiclient.NewEndpointRequest() // EndpointRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.EndpointsUpdate(context.Background(), id).EndpointRequest(endpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.EndpointsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointsUpdate`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.EndpointsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndpointsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointRequest** | [**EndpointRequest**](EndpointRequest.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

