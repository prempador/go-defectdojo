# \MetadataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataBatchCreate**](MetadataAPI.md#MetadataBatchCreate) | **Post** /api/v2/metadata/batch/ | 
[**MetadataBatchPartialUpdate**](MetadataAPI.md#MetadataBatchPartialUpdate) | **Patch** /api/v2/metadata/batch/ | 
[**MetadataCreate**](MetadataAPI.md#MetadataCreate) | **Post** /api/v2/metadata/ | 
[**MetadataDeletePreviewList**](MetadataAPI.md#MetadataDeletePreviewList) | **Get** /api/v2/metadata/{id}/delete_preview/ | 
[**MetadataDestroy**](MetadataAPI.md#MetadataDestroy) | **Delete** /api/v2/metadata/{id}/ | 
[**MetadataList**](MetadataAPI.md#MetadataList) | **Get** /api/v2/metadata/ | 
[**MetadataPartialUpdate**](MetadataAPI.md#MetadataPartialUpdate) | **Patch** /api/v2/metadata/{id}/ | 
[**MetadataRetrieve**](MetadataAPI.md#MetadataRetrieve) | **Get** /api/v2/metadata/{id}/ | 
[**MetadataUpdate**](MetadataAPI.md#MetadataUpdate) | **Put** /api/v2/metadata/{id}/ | 



## MetadataBatchCreate

> MetaMain MetadataBatchCreate(ctx).MetaMainRequest(metaMainRequest).Execute()



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
	metaMainRequest := *openapiclient.NewMetaMainRequest([]openapiclient.MetadataRequest{*openapiclient.NewMetadataRequest("Name_example", "Value_example")}) // MetaMainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataBatchCreate(context.Background()).MetaMainRequest(metaMainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataBatchCreate`: MetaMain
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metaMainRequest** | [**MetaMainRequest**](MetaMainRequest.md) |  | 

### Return type

[**MetaMain**](MetaMain.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataBatchPartialUpdate

> MetaMain MetadataBatchPartialUpdate(ctx).PatchedMetaMainRequest(patchedMetaMainRequest).Execute()



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
	patchedMetaMainRequest := *openapiclient.NewPatchedMetaMainRequest() // PatchedMetaMainRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataBatchPartialUpdate(context.Background()).PatchedMetaMainRequest(patchedMetaMainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataBatchPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataBatchPartialUpdate`: MetaMain
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataBatchPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataBatchPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchedMetaMainRequest** | [**PatchedMetaMainRequest**](PatchedMetaMainRequest.md) |  | 

### Return type

[**MetaMain**](MetaMain.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataCreate

> Meta MetadataCreate(ctx).MetaRequest(metaRequest).Execute()



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
	metaRequest := *openapiclient.NewMetaRequest("Name_example", "Value_example") // MetaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataCreate(context.Background()).MetaRequest(metaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataCreate`: Meta
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metaRequest** | [**MetaRequest**](MetaRequest.md) |  | 

### Return type

[**Meta**](Meta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataDeletePreviewList

> PaginatedDeletePreviewList MetadataDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo meta.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo meta. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataDeletePreviewListRequest struct via the builder pattern


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


## MetadataDestroy

> MetadataDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo meta.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetadataAPI.MetadataDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo meta. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataDestroyRequest struct via the builder pattern


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


## MetadataList

> PaginatedMetaList MetadataList(ctx).Endpoint(endpoint).Finding(finding).Id(id).Limit(limit).Name(name).NameCaseInsensitive(nameCaseInsensitive).Offset(offset).Product(product).Value(value).ValueCaseInsensitive(valueCaseInsensitive).Execute()



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
	finding := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	nameCaseInsensitive := "nameCaseInsensitive_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	product := int32(56) // int32 |  (optional)
	value := "value_example" // string |  (optional)
	valueCaseInsensitive := "valueCaseInsensitive_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataList(context.Background()).Endpoint(endpoint).Finding(finding).Id(id).Limit(limit).Name(name).NameCaseInsensitive(nameCaseInsensitive).Offset(offset).Product(product).Value(value).ValueCaseInsensitive(valueCaseInsensitive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataList`: PaginatedMetaList
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpoint** | **int32** |  | 
 **finding** | **int32** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **nameCaseInsensitive** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **product** | **int32** |  | 
 **value** | **string** |  | 
 **valueCaseInsensitive** | **string** |  | 

### Return type

[**PaginatedMetaList**](PaginatedMetaList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataPartialUpdate

> Meta MetadataPartialUpdate(ctx, id).PatchedMetaRequest(patchedMetaRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo meta.
	patchedMetaRequest := *openapiclient.NewPatchedMetaRequest() // PatchedMetaRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataPartialUpdate(context.Background(), id).PatchedMetaRequest(patchedMetaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataPartialUpdate`: Meta
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo meta. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedMetaRequest** | [**PatchedMetaRequest**](PatchedMetaRequest.md) |  | 

### Return type

[**Meta**](Meta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataRetrieve

> Meta MetadataRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo meta.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataRetrieve`: Meta
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo meta. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Meta**](Meta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataUpdate

> Meta MetadataUpdate(ctx, id).MetaRequest(metaRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo meta.
	metaRequest := *openapiclient.NewMetaRequest("Name_example", "Value_example") // MetaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetadataAPI.MetadataUpdate(context.Background(), id).MetaRequest(metaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.MetadataUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetadataUpdate`: Meta
	fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.MetadataUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo meta. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetadataUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metaRequest** | [**MetaRequest**](MetaRequest.md) |  | 

### Return type

[**Meta**](Meta.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

