# \TechnologiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TechnologiesCreate**](TechnologiesAPI.md#TechnologiesCreate) | **Post** /api/v2/technologies/ | 
[**TechnologiesDeletePreviewList**](TechnologiesAPI.md#TechnologiesDeletePreviewList) | **Get** /api/v2/technologies/{id}/delete_preview/ | 
[**TechnologiesDestroy**](TechnologiesAPI.md#TechnologiesDestroy) | **Delete** /api/v2/technologies/{id}/ | 
[**TechnologiesList**](TechnologiesAPI.md#TechnologiesList) | **Get** /api/v2/technologies/ | 
[**TechnologiesPartialUpdate**](TechnologiesAPI.md#TechnologiesPartialUpdate) | **Patch** /api/v2/technologies/{id}/ | 
[**TechnologiesRetrieve**](TechnologiesAPI.md#TechnologiesRetrieve) | **Get** /api/v2/technologies/{id}/ | 
[**TechnologiesUpdate**](TechnologiesAPI.md#TechnologiesUpdate) | **Put** /api/v2/technologies/{id}/ | 



## TechnologiesCreate

> AppAnalysis TechnologiesCreate(ctx).AppAnalysisRequest(appAnalysisRequest).Execute()



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
	appAnalysisRequest := *openapiclient.NewAppAnalysisRequest("Name_example", int32(123), int32(123)) // AppAnalysisRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechnologiesAPI.TechnologiesCreate(context.Background()).AppAnalysisRequest(appAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechnologiesCreate`: AppAnalysis
	fmt.Fprintf(os.Stdout, "Response from `TechnologiesAPI.TechnologiesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appAnalysisRequest** | [**AppAnalysisRequest**](AppAnalysisRequest.md) |  | 

### Return type

[**AppAnalysis**](AppAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechnologiesDeletePreviewList

> PaginatedDeletePreviewList TechnologiesDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this app_ analysis.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechnologiesAPI.TechnologiesDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechnologiesDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `TechnologiesAPI.TechnologiesDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this app_ analysis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesDeletePreviewListRequest struct via the builder pattern


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


## TechnologiesDestroy

> TechnologiesDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this app_ analysis.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TechnologiesAPI.TechnologiesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this app_ analysis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesDestroyRequest struct via the builder pattern


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


## TechnologiesList

> PaginatedAppAnalysisList TechnologiesList(ctx).Limit(limit).Name(name).NotTag(notTag).NotTags(notTags).Offset(offset).Product(product).Tag(tag).Tags(tags).User(user).Version(version).Execute()



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
	name := "name_example" // string |  (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	product := int32(56) // int32 |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	user := int32(56) // int32 |  (optional)
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechnologiesAPI.TechnologiesList(context.Background()).Limit(limit).Name(name).NotTag(notTag).NotTags(notTags).Offset(offset).Product(product).Tag(tag).Tags(tags).User(user).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechnologiesList`: PaginatedAppAnalysisList
	fmt.Fprintf(os.Stdout, "Response from `TechnologiesAPI.TechnologiesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **product** | **int32** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **user** | **int32** |  | 
 **version** | **string** |  | 

### Return type

[**PaginatedAppAnalysisList**](PaginatedAppAnalysisList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechnologiesPartialUpdate

> AppAnalysis TechnologiesPartialUpdate(ctx, id).PatchedAppAnalysisRequest(patchedAppAnalysisRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this app_ analysis.
	patchedAppAnalysisRequest := *openapiclient.NewPatchedAppAnalysisRequest() // PatchedAppAnalysisRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechnologiesAPI.TechnologiesPartialUpdate(context.Background(), id).PatchedAppAnalysisRequest(patchedAppAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechnologiesPartialUpdate`: AppAnalysis
	fmt.Fprintf(os.Stdout, "Response from `TechnologiesAPI.TechnologiesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this app_ analysis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAppAnalysisRequest** | [**PatchedAppAnalysisRequest**](PatchedAppAnalysisRequest.md) |  | 

### Return type

[**AppAnalysis**](AppAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechnologiesRetrieve

> AppAnalysis TechnologiesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this app_ analysis.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechnologiesAPI.TechnologiesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechnologiesRetrieve`: AppAnalysis
	fmt.Fprintf(os.Stdout, "Response from `TechnologiesAPI.TechnologiesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this app_ analysis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppAnalysis**](AppAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TechnologiesUpdate

> AppAnalysis TechnologiesUpdate(ctx, id).AppAnalysisRequest(appAnalysisRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this app_ analysis.
	appAnalysisRequest := *openapiclient.NewAppAnalysisRequest("Name_example", int32(123), int32(123)) // AppAnalysisRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TechnologiesAPI.TechnologiesUpdate(context.Background(), id).AppAnalysisRequest(appAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TechnologiesAPI.TechnologiesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TechnologiesUpdate`: AppAnalysis
	fmt.Fprintf(os.Stdout, "Response from `TechnologiesAPI.TechnologiesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this app_ analysis. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTechnologiesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appAnalysisRequest** | [**AppAnalysisRequest**](AppAnalysisRequest.md) |  | 

### Return type

[**AppAnalysis**](AppAnalysis.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

