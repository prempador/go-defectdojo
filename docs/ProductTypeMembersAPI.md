# \ProductTypeMembersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductTypeMembersCreate**](ProductTypeMembersAPI.md#ProductTypeMembersCreate) | **Post** /api/v2/product_type_members/ | 
[**ProductTypeMembersDeletePreviewList**](ProductTypeMembersAPI.md#ProductTypeMembersDeletePreviewList) | **Get** /api/v2/product_type_members/{id}/delete_preview/ | 
[**ProductTypeMembersDestroy**](ProductTypeMembersAPI.md#ProductTypeMembersDestroy) | **Delete** /api/v2/product_type_members/{id}/ | 
[**ProductTypeMembersList**](ProductTypeMembersAPI.md#ProductTypeMembersList) | **Get** /api/v2/product_type_members/ | 
[**ProductTypeMembersRetrieve**](ProductTypeMembersAPI.md#ProductTypeMembersRetrieve) | **Get** /api/v2/product_type_members/{id}/ | 
[**ProductTypeMembersUpdate**](ProductTypeMembersAPI.md#ProductTypeMembersUpdate) | **Put** /api/v2/product_type_members/{id}/ | 



## ProductTypeMembersCreate

> ProductTypeMember ProductTypeMembersCreate(ctx).ProductTypeMemberRequest(productTypeMemberRequest).Execute()



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
	productTypeMemberRequest := *openapiclient.NewProductTypeMemberRequest(int32(123), int32(123), int32(123)) // ProductTypeMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeMembersAPI.ProductTypeMembersCreate(context.Background()).ProductTypeMemberRequest(productTypeMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeMembersAPI.ProductTypeMembersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeMembersCreate`: ProductTypeMember
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeMembersAPI.ProductTypeMembersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeMembersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productTypeMemberRequest** | [**ProductTypeMemberRequest**](ProductTypeMemberRequest.md) |  | 

### Return type

[**ProductTypeMember**](ProductTypeMember.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypeMembersDeletePreviewList

> PaginatedDeletePreviewList ProductTypeMembersDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ member.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeMembersAPI.ProductTypeMembersDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeMembersAPI.ProductTypeMembersDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeMembersDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeMembersAPI.ProductTypeMembersDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeMembersDeletePreviewListRequest struct via the builder pattern


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


## ProductTypeMembersDestroy

> ProductTypeMembersDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ member.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductTypeMembersAPI.ProductTypeMembersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeMembersAPI.ProductTypeMembersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeMembersDestroyRequest struct via the builder pattern


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


## ProductTypeMembersList

> PaginatedProductTypeMemberList ProductTypeMembersList(ctx).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).ProductTypeId(productTypeId).UserId(userId).Execute()



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
	productTypeId := int32(56) // int32 |  (optional)
	userId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeMembersAPI.ProductTypeMembersList(context.Background()).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).ProductTypeId(productTypeId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeMembersAPI.ProductTypeMembersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeMembersList`: PaginatedProductTypeMemberList
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeMembersAPI.ProductTypeMembersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeMembersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **productTypeId** | **int32** |  | 
 **userId** | **int32** |  | 

### Return type

[**PaginatedProductTypeMemberList**](PaginatedProductTypeMemberList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypeMembersRetrieve

> ProductTypeMember ProductTypeMembersRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ member.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeMembersAPI.ProductTypeMembersRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeMembersAPI.ProductTypeMembersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeMembersRetrieve`: ProductTypeMember
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeMembersAPI.ProductTypeMembersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeMembersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**ProductTypeMember**](ProductTypeMember.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypeMembersUpdate

> ProductTypeMember ProductTypeMembersUpdate(ctx, id).ProductTypeMemberRequest(productTypeMemberRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ type_ member.
	productTypeMemberRequest := *openapiclient.NewProductTypeMemberRequest(int32(123), int32(123), int32(123)) // ProductTypeMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductTypeMembersAPI.ProductTypeMembersUpdate(context.Background(), id).ProductTypeMemberRequest(productTypeMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductTypeMembersAPI.ProductTypeMembersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductTypeMembersUpdate`: ProductTypeMember
	fmt.Fprintf(os.Stdout, "Response from `ProductTypeMembersAPI.ProductTypeMembersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ type_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypeMembersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productTypeMemberRequest** | [**ProductTypeMemberRequest**](ProductTypeMemberRequest.md) |  | 

### Return type

[**ProductTypeMember**](ProductTypeMember.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

