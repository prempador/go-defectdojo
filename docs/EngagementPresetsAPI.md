# \EngagementPresetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngagementPresetsCreate**](EngagementPresetsAPI.md#EngagementPresetsCreate) | **Post** /api/v2/engagement_presets/ | 
[**EngagementPresetsDeletePreviewList**](EngagementPresetsAPI.md#EngagementPresetsDeletePreviewList) | **Get** /api/v2/engagement_presets/{id}/delete_preview/ | 
[**EngagementPresetsDestroy**](EngagementPresetsAPI.md#EngagementPresetsDestroy) | **Delete** /api/v2/engagement_presets/{id}/ | 
[**EngagementPresetsList**](EngagementPresetsAPI.md#EngagementPresetsList) | **Get** /api/v2/engagement_presets/ | 
[**EngagementPresetsPartialUpdate**](EngagementPresetsAPI.md#EngagementPresetsPartialUpdate) | **Patch** /api/v2/engagement_presets/{id}/ | 
[**EngagementPresetsRetrieve**](EngagementPresetsAPI.md#EngagementPresetsRetrieve) | **Get** /api/v2/engagement_presets/{id}/ | 
[**EngagementPresetsUpdate**](EngagementPresetsAPI.md#EngagementPresetsUpdate) | **Put** /api/v2/engagement_presets/{id}/ | 



## EngagementPresetsCreate

> EngagementPresets EngagementPresetsCreate(ctx).EngagementPresetsRequest(engagementPresetsRequest).Execute()



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
	engagementPresetsRequest := *openapiclient.NewEngagementPresetsRequest(int32(123)) // EngagementPresetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementPresetsAPI.EngagementPresetsCreate(context.Background()).EngagementPresetsRequest(engagementPresetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementPresetsCreate`: EngagementPresets
	fmt.Fprintf(os.Stdout, "Response from `EngagementPresetsAPI.EngagementPresetsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engagementPresetsRequest** | [**EngagementPresetsRequest**](EngagementPresetsRequest.md) |  | 

### Return type

[**EngagementPresets**](EngagementPresets.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementPresetsDeletePreviewList

> PaginatedDeletePreviewList EngagementPresetsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement_ presets.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementPresetsAPI.EngagementPresetsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementPresetsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `EngagementPresetsAPI.EngagementPresetsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement_ presets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsDeletePreviewListRequest struct via the builder pattern


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


## EngagementPresetsDestroy

> EngagementPresetsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement_ presets.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementPresetsAPI.EngagementPresetsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement_ presets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsDestroyRequest struct via the builder pattern


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


## EngagementPresetsList

> PaginatedEngagementPresetsList EngagementPresetsList(ctx).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).Product(product).Title(title).Execute()



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
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	product := int32(56) // int32 |  (optional)
	title := "title_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementPresetsAPI.EngagementPresetsList(context.Background()).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).Product(product).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementPresetsList`: PaginatedEngagementPresetsList
	fmt.Fprintf(os.Stdout, "Response from `EngagementPresetsAPI.EngagementPresetsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **product** | **int32** |  | 
 **title** | **string** |  | 

### Return type

[**PaginatedEngagementPresetsList**](PaginatedEngagementPresetsList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementPresetsPartialUpdate

> EngagementPresets EngagementPresetsPartialUpdate(ctx, id).PatchedEngagementPresetsRequest(patchedEngagementPresetsRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement_ presets.
	patchedEngagementPresetsRequest := *openapiclient.NewPatchedEngagementPresetsRequest() // PatchedEngagementPresetsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementPresetsAPI.EngagementPresetsPartialUpdate(context.Background(), id).PatchedEngagementPresetsRequest(patchedEngagementPresetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementPresetsPartialUpdate`: EngagementPresets
	fmt.Fprintf(os.Stdout, "Response from `EngagementPresetsAPI.EngagementPresetsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement_ presets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEngagementPresetsRequest** | [**PatchedEngagementPresetsRequest**](PatchedEngagementPresetsRequest.md) |  | 

### Return type

[**EngagementPresets**](EngagementPresets.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementPresetsRetrieve

> EngagementPresets EngagementPresetsRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement_ presets.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementPresetsAPI.EngagementPresetsRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementPresetsRetrieve`: EngagementPresets
	fmt.Fprintf(os.Stdout, "Response from `EngagementPresetsAPI.EngagementPresetsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement_ presets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**EngagementPresets**](EngagementPresets.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementPresetsUpdate

> EngagementPresets EngagementPresetsUpdate(ctx, id).EngagementPresetsRequest(engagementPresetsRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement_ presets.
	engagementPresetsRequest := *openapiclient.NewEngagementPresetsRequest(int32(123)) // EngagementPresetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementPresetsAPI.EngagementPresetsUpdate(context.Background(), id).EngagementPresetsRequest(engagementPresetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementPresetsAPI.EngagementPresetsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementPresetsUpdate`: EngagementPresets
	fmt.Fprintf(os.Stdout, "Response from `EngagementPresetsAPI.EngagementPresetsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement_ presets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementPresetsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engagementPresetsRequest** | [**EngagementPresetsRequest**](EngagementPresetsRequest.md) |  | 

### Return type

[**EngagementPresets**](EngagementPresets.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

