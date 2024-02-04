# \StubFindingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StubFindingsCreate**](StubFindingsAPI.md#StubFindingsCreate) | **Post** /api/v2/stub_findings/ | 
[**StubFindingsDeletePreviewList**](StubFindingsAPI.md#StubFindingsDeletePreviewList) | **Get** /api/v2/stub_findings/{id}/delete_preview/ | 
[**StubFindingsDestroy**](StubFindingsAPI.md#StubFindingsDestroy) | **Delete** /api/v2/stub_findings/{id}/ | 
[**StubFindingsList**](StubFindingsAPI.md#StubFindingsList) | **Get** /api/v2/stub_findings/ | 
[**StubFindingsPartialUpdate**](StubFindingsAPI.md#StubFindingsPartialUpdate) | **Patch** /api/v2/stub_findings/{id}/ | 
[**StubFindingsRetrieve**](StubFindingsAPI.md#StubFindingsRetrieve) | **Get** /api/v2/stub_findings/{id}/ | 
[**StubFindingsUpdate**](StubFindingsAPI.md#StubFindingsUpdate) | **Put** /api/v2/stub_findings/{id}/ | 



## StubFindingsCreate

> StubFindingCreate StubFindingsCreate(ctx).StubFindingCreateRequest(stubFindingCreateRequest).Execute()



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
	stubFindingCreateRequest := *openapiclient.NewStubFindingCreateRequest(int32(123), "Title_example") // StubFindingCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StubFindingsAPI.StubFindingsCreate(context.Background()).StubFindingCreateRequest(stubFindingCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StubFindingsCreate`: StubFindingCreate
	fmt.Fprintf(os.Stdout, "Response from `StubFindingsAPI.StubFindingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubFindingCreateRequest** | [**StubFindingCreateRequest**](StubFindingCreateRequest.md) |  | 

### Return type

[**StubFindingCreate**](StubFindingCreate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StubFindingsDeletePreviewList

> PaginatedDeletePreviewList StubFindingsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this stub_ finding.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StubFindingsAPI.StubFindingsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StubFindingsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `StubFindingsAPI.StubFindingsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this stub_ finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsDeletePreviewListRequest struct via the builder pattern


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


## StubFindingsDestroy

> StubFindingsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this stub_ finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StubFindingsAPI.StubFindingsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this stub_ finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsDestroyRequest struct via the builder pattern


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


## StubFindingsList

> PaginatedStubFindingList StubFindingsList(ctx).Date(date).Description(description).Id(id).Limit(limit).Offset(offset).Severity(severity).Title(title).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	date := time.Now() // string |  (optional)
	description := "description_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	severity := "severity_example" // string |  (optional)
	title := "title_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StubFindingsAPI.StubFindingsList(context.Background()).Date(date).Description(description).Id(id).Limit(limit).Offset(offset).Severity(severity).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StubFindingsList`: PaginatedStubFindingList
	fmt.Fprintf(os.Stdout, "Response from `StubFindingsAPI.StubFindingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** |  | 
 **description** | **string** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **severity** | **string** |  | 
 **title** | **string** |  | 

### Return type

[**PaginatedStubFindingList**](PaginatedStubFindingList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StubFindingsPartialUpdate

> StubFinding StubFindingsPartialUpdate(ctx, id).PatchedStubFindingRequest(patchedStubFindingRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this stub_ finding.
	patchedStubFindingRequest := *openapiclient.NewPatchedStubFindingRequest() // PatchedStubFindingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StubFindingsAPI.StubFindingsPartialUpdate(context.Background(), id).PatchedStubFindingRequest(patchedStubFindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StubFindingsPartialUpdate`: StubFinding
	fmt.Fprintf(os.Stdout, "Response from `StubFindingsAPI.StubFindingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this stub_ finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedStubFindingRequest** | [**PatchedStubFindingRequest**](PatchedStubFindingRequest.md) |  | 

### Return type

[**StubFinding**](StubFinding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StubFindingsRetrieve

> StubFinding StubFindingsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this stub_ finding.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StubFindingsAPI.StubFindingsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StubFindingsRetrieve`: StubFinding
	fmt.Fprintf(os.Stdout, "Response from `StubFindingsAPI.StubFindingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this stub_ finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StubFinding**](StubFinding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StubFindingsUpdate

> StubFinding StubFindingsUpdate(ctx, id).StubFindingRequest(stubFindingRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this stub_ finding.
	stubFindingRequest := *openapiclient.NewStubFindingRequest("Title_example") // StubFindingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StubFindingsAPI.StubFindingsUpdate(context.Background(), id).StubFindingRequest(stubFindingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StubFindingsAPI.StubFindingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StubFindingsUpdate`: StubFinding
	fmt.Fprintf(os.Stdout, "Response from `StubFindingsAPI.StubFindingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this stub_ finding. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStubFindingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stubFindingRequest** | [**StubFindingRequest**](StubFindingRequest.md) |  | 

### Return type

[**StubFinding**](StubFinding.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

