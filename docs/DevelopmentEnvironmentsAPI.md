# \DevelopmentEnvironmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevelopmentEnvironmentsCreate**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsCreate) | **Post** /api/v2/development_environments/ | 
[**DevelopmentEnvironmentsDeletePreviewList**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsDeletePreviewList) | **Get** /api/v2/development_environments/{id}/delete_preview/ | 
[**DevelopmentEnvironmentsDestroy**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsDestroy) | **Delete** /api/v2/development_environments/{id}/ | 
[**DevelopmentEnvironmentsList**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsList) | **Get** /api/v2/development_environments/ | 
[**DevelopmentEnvironmentsPartialUpdate**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsPartialUpdate) | **Patch** /api/v2/development_environments/{id}/ | 
[**DevelopmentEnvironmentsRetrieve**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsRetrieve) | **Get** /api/v2/development_environments/{id}/ | 
[**DevelopmentEnvironmentsUpdate**](DevelopmentEnvironmentsAPI.md#DevelopmentEnvironmentsUpdate) | **Put** /api/v2/development_environments/{id}/ | 



## DevelopmentEnvironmentsCreate

> DevelopmentEnvironment DevelopmentEnvironmentsCreate(ctx).DevelopmentEnvironmentRequest(developmentEnvironmentRequest).Execute()



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
	developmentEnvironmentRequest := *openapiclient.NewDevelopmentEnvironmentRequest("Name_example") // DevelopmentEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsCreate(context.Background()).DevelopmentEnvironmentRequest(developmentEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevelopmentEnvironmentsCreate`: DevelopmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **developmentEnvironmentRequest** | [**DevelopmentEnvironmentRequest**](DevelopmentEnvironmentRequest.md) |  | 

### Return type

[**DevelopmentEnvironment**](DevelopmentEnvironment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevelopmentEnvironmentsDeletePreviewList

> PaginatedDeletePreviewList DevelopmentEnvironmentsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this development_ environment.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevelopmentEnvironmentsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this development_ environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsDeletePreviewListRequest struct via the builder pattern


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


## DevelopmentEnvironmentsDestroy

> DevelopmentEnvironmentsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this development_ environment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this development_ environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsDestroyRequest struct via the builder pattern


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


## DevelopmentEnvironmentsList

> PaginatedDevelopmentEnvironmentList DevelopmentEnvironmentsList(ctx).Limit(limit).Offset(offset).Execute()



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
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevelopmentEnvironmentsList`: PaginatedDevelopmentEnvironmentList
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedDevelopmentEnvironmentList**](PaginatedDevelopmentEnvironmentList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevelopmentEnvironmentsPartialUpdate

> DevelopmentEnvironment DevelopmentEnvironmentsPartialUpdate(ctx, id).PatchedDevelopmentEnvironmentRequest(patchedDevelopmentEnvironmentRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this development_ environment.
	patchedDevelopmentEnvironmentRequest := *openapiclient.NewPatchedDevelopmentEnvironmentRequest() // PatchedDevelopmentEnvironmentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsPartialUpdate(context.Background(), id).PatchedDevelopmentEnvironmentRequest(patchedDevelopmentEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevelopmentEnvironmentsPartialUpdate`: DevelopmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this development_ environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDevelopmentEnvironmentRequest** | [**PatchedDevelopmentEnvironmentRequest**](PatchedDevelopmentEnvironmentRequest.md) |  | 

### Return type

[**DevelopmentEnvironment**](DevelopmentEnvironment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevelopmentEnvironmentsRetrieve

> DevelopmentEnvironment DevelopmentEnvironmentsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this development_ environment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevelopmentEnvironmentsRetrieve`: DevelopmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this development_ environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DevelopmentEnvironment**](DevelopmentEnvironment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevelopmentEnvironmentsUpdate

> DevelopmentEnvironment DevelopmentEnvironmentsUpdate(ctx, id).DevelopmentEnvironmentRequest(developmentEnvironmentRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this development_ environment.
	developmentEnvironmentRequest := *openapiclient.NewDevelopmentEnvironmentRequest("Name_example") // DevelopmentEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsUpdate(context.Background(), id).DevelopmentEnvironmentRequest(developmentEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevelopmentEnvironmentsUpdate`: DevelopmentEnvironment
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentEnvironmentsAPI.DevelopmentEnvironmentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this development_ environment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevelopmentEnvironmentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **developmentEnvironmentRequest** | [**DevelopmentEnvironmentRequest**](DevelopmentEnvironmentRequest.md) |  | 

### Return type

[**DevelopmentEnvironment**](DevelopmentEnvironment.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

