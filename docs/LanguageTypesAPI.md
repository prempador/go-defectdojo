# \LanguageTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LanguageTypesCreate**](LanguageTypesAPI.md#LanguageTypesCreate) | **Post** /api/v2/language_types/ | 
[**LanguageTypesDeletePreviewList**](LanguageTypesAPI.md#LanguageTypesDeletePreviewList) | **Get** /api/v2/language_types/{id}/delete_preview/ | 
[**LanguageTypesDestroy**](LanguageTypesAPI.md#LanguageTypesDestroy) | **Delete** /api/v2/language_types/{id}/ | 
[**LanguageTypesList**](LanguageTypesAPI.md#LanguageTypesList) | **Get** /api/v2/language_types/ | 
[**LanguageTypesPartialUpdate**](LanguageTypesAPI.md#LanguageTypesPartialUpdate) | **Patch** /api/v2/language_types/{id}/ | 
[**LanguageTypesRetrieve**](LanguageTypesAPI.md#LanguageTypesRetrieve) | **Get** /api/v2/language_types/{id}/ | 
[**LanguageTypesUpdate**](LanguageTypesAPI.md#LanguageTypesUpdate) | **Put** /api/v2/language_types/{id}/ | 



## LanguageTypesCreate

> LanguageType LanguageTypesCreate(ctx).LanguageTypeRequest(languageTypeRequest).Execute()



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
	languageTypeRequest := *openapiclient.NewLanguageTypeRequest("Language_example") // LanguageTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageTypesAPI.LanguageTypesCreate(context.Background()).LanguageTypeRequest(languageTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguageTypesCreate`: LanguageType
	fmt.Fprintf(os.Stdout, "Response from `LanguageTypesAPI.LanguageTypesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageTypeRequest** | [**LanguageTypeRequest**](LanguageTypeRequest.md) |  | 

### Return type

[**LanguageType**](LanguageType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LanguageTypesDeletePreviewList

> PaginatedDeletePreviewList LanguageTypesDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this language_ type.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageTypesAPI.LanguageTypesDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguageTypesDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `LanguageTypesAPI.LanguageTypesDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this language_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesDeletePreviewListRequest struct via the builder pattern


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


## LanguageTypesDestroy

> LanguageTypesDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this language_ type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LanguageTypesAPI.LanguageTypesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this language_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesDestroyRequest struct via the builder pattern


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


## LanguageTypesList

> PaginatedLanguageTypeList LanguageTypesList(ctx).Color(color).Id(id).Language(language).Limit(limit).Offset(offset).Execute()



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
	color := "color_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	language := "language_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageTypesAPI.LanguageTypesList(context.Background()).Color(color).Id(id).Language(language).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguageTypesList`: PaginatedLanguageTypeList
	fmt.Fprintf(os.Stdout, "Response from `LanguageTypesAPI.LanguageTypesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **color** | **string** |  | 
 **id** | **int32** |  | 
 **language** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedLanguageTypeList**](PaginatedLanguageTypeList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LanguageTypesPartialUpdate

> LanguageType LanguageTypesPartialUpdate(ctx, id).PatchedLanguageTypeRequest(patchedLanguageTypeRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this language_ type.
	patchedLanguageTypeRequest := *openapiclient.NewPatchedLanguageTypeRequest() // PatchedLanguageTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageTypesAPI.LanguageTypesPartialUpdate(context.Background(), id).PatchedLanguageTypeRequest(patchedLanguageTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguageTypesPartialUpdate`: LanguageType
	fmt.Fprintf(os.Stdout, "Response from `LanguageTypesAPI.LanguageTypesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this language_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedLanguageTypeRequest** | [**PatchedLanguageTypeRequest**](PatchedLanguageTypeRequest.md) |  | 

### Return type

[**LanguageType**](LanguageType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LanguageTypesRetrieve

> LanguageType LanguageTypesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this language_ type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageTypesAPI.LanguageTypesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguageTypesRetrieve`: LanguageType
	fmt.Fprintf(os.Stdout, "Response from `LanguageTypesAPI.LanguageTypesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this language_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LanguageType**](LanguageType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LanguageTypesUpdate

> LanguageType LanguageTypesUpdate(ctx, id).LanguageTypeRequest(languageTypeRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this language_ type.
	languageTypeRequest := *openapiclient.NewLanguageTypeRequest("Language_example") // LanguageTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanguageTypesAPI.LanguageTypesUpdate(context.Background(), id).LanguageTypeRequest(languageTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanguageTypesAPI.LanguageTypesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguageTypesUpdate`: LanguageType
	fmt.Fprintf(os.Stdout, "Response from `LanguageTypesAPI.LanguageTypesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this language_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLanguageTypesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageTypeRequest** | [**LanguageTypeRequest**](LanguageTypeRequest.md) |  | 

### Return type

[**LanguageType**](LanguageType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

