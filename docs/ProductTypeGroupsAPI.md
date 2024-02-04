# \ProductTypeGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductTypeGroupsCreate**](ProductTypeGroupsAPI.md#ProductTypeGroupsCreate) | **Post** /api/v2/product_type_groups/ | 
[**ProductTypeGroupsDeletePreviewList**](ProductTypeGroupsAPI.md#ProductTypeGroupsDeletePreviewList) | **Get** /api/v2/product_type_groups/{id}/delete_preview/ | 
[**ProductTypeGroupsDestroy**](ProductTypeGroupsAPI.md#ProductTypeGroupsDestroy) | **Delete** /api/v2/product_type_groups/{id}/ | 
[**ProductTypeGroupsList**](ProductTypeGroupsAPI.md#ProductTypeGroupsList) | **Get** /api/v2/product_type_groups/ | 
[**ProductTypeGroupsRetrieve**](ProductTypeGroupsAPI.md#ProductTypeGroupsRetrieve) | **Get** /api/v2/product_type_groups/{id}/ | 
[**ProductTypeGroupsUpdate**](ProductTypeGroupsAPI.md#ProductTypeGroupsUpdate) | **Put** /api/v2/product_type_groups/{id}/ | 



## ProductTypeGroupsCreate

> ProductTypeGroup ProductTypeGroupsCreate(ctx).ProductTypeGroupRequest(productTypeGroupRequest).Execute()



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
	productTypeGroupRequest := *openapiclient.NewProductTypeGroupRequest(int32(123), int32(123), int32(123)) // ProductTypeGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeGroupsAPI.ProductTypeGroupsCreate(context.Background()).ProductTypeGroupRequest(productTypeGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeGroupsAPI.ProductTypeGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeGroupsCreate`: ProductTypeGroup
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeGroupsAPI.ProductTypeGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productTypeGroupRequest** | [**ProductTypeGroupRequest**](ProductTypeGroupRequest.md) |  | 

### Return type

[**ProductTypeGroup**](ProductTypeGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypeGroupsDeletePreviewList

> PaginatedDeletePreviewList ProductTypeGroupsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ group.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeGroupsAPI.ProductTypeGroupsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeGroupsAPI.ProductTypeGroupsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeGroupsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeGroupsAPI.ProductTypeGroupsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeGroupsDeletePreviewListRequest struct via the builder pattern


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


## ProductTypeGroupsDestroy

> ProductTypeGroupsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTypeGroupsAPI.ProductTypeGroupsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeGroupsAPI.ProductTypeGroupsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeGroupsDestroyRequest struct via the builder pattern


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


## ProductTypeGroupsList

> PaginatedProductTypeGroupList ProductTypeGroupsList(ctx).GroupId(groupId).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).ProductTypeId(productTypeId).Execute()



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
	groupId := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	productTypeId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeGroupsAPI.ProductTypeGroupsList(context.Background()).GroupId(groupId).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).ProductTypeId(productTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeGroupsAPI.ProductTypeGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeGroupsList`: PaginatedProductTypeGroupList
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeGroupsAPI.ProductTypeGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **productTypeId** | **int32** |  | 

### Return type

[**PaginatedProductTypeGroupList**](PaginatedProductTypeGroupList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypeGroupsRetrieve

> ProductTypeGroup ProductTypeGroupsRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ group.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeGroupsAPI.ProductTypeGroupsRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeGroupsAPI.ProductTypeGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeGroupsRetrieve`: ProductTypeGroup
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeGroupsAPI.ProductTypeGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**ProductTypeGroup**](ProductTypeGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypeGroupsUpdate

> ProductTypeGroup ProductTypeGroupsUpdate(ctx, id).ProductTypeGroupRequest(productTypeGroupRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ group.
	productTypeGroupRequest := *openapiclient.NewProductTypeGroupRequest(int32(123), int32(123), int32(123)) // ProductTypeGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeGroupsAPI.ProductTypeGroupsUpdate(context.Background(), id).ProductTypeGroupRequest(productTypeGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeGroupsAPI.ProductTypeGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeGroupsUpdate`: ProductTypeGroup
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeGroupsAPI.ProductTypeGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productTypeGroupRequest** | [**ProductTypeGroupRequest**](ProductTypeGroupRequest.md) |  | 

### Return type

[**ProductTypeGroup**](ProductTypeGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

