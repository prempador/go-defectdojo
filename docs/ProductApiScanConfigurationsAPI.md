# \ProductApiScanConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductApiScanConfigurationsCreate**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsCreate) | **Post** /api/v2/product_api_scan_configurations/ | 
[**ProductApiScanConfigurationsDeletePreviewList**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsDeletePreviewList) | **Get** /api/v2/product_api_scan_configurations/{id}/delete_preview/ | 
[**ProductApiScanConfigurationsDestroy**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsDestroy) | **Delete** /api/v2/product_api_scan_configurations/{id}/ | 
[**ProductApiScanConfigurationsList**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsList) | **Get** /api/v2/product_api_scan_configurations/ | 
[**ProductApiScanConfigurationsPartialUpdate**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsPartialUpdate) | **Patch** /api/v2/product_api_scan_configurations/{id}/ | 
[**ProductApiScanConfigurationsRetrieve**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsRetrieve) | **Get** /api/v2/product_api_scan_configurations/{id}/ | 
[**ProductApiScanConfigurationsUpdate**](ProductApiScanConfigurationsAPI.md#ProductApiScanConfigurationsUpdate) | **Put** /api/v2/product_api_scan_configurations/{id}/ | 



## ProductApiScanConfigurationsCreate

> ProductAPIScanConfiguration ProductApiScanConfigurationsCreate(ctx).ProductAPIScanConfigurationRequest(productAPIScanConfigurationRequest).Execute()



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
	productAPIScanConfigurationRequest := *openapiclient.NewProductAPIScanConfigurationRequest(int32(123), int32(123)) // ProductAPIScanConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsCreate(context.Background()).ProductAPIScanConfigurationRequest(productAPIScanConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductApiScanConfigurationsCreate`: ProductAPIScanConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productAPIScanConfigurationRequest** | [**ProductAPIScanConfigurationRequest**](ProductAPIScanConfigurationRequest.md) |  | 

### Return type

[**ProductAPIScanConfiguration**](ProductAPIScanConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductApiScanConfigurationsDeletePreviewList

> PaginatedDeletePreviewList ProductApiScanConfigurationsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ap i_ scan_ configuration.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductApiScanConfigurationsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ap i_ scan_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsDeletePreviewListRequest struct via the builder pattern


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


## ProductApiScanConfigurationsDestroy

> ProductApiScanConfigurationsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ap i_ scan_ configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ap i_ scan_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsDestroyRequest struct via the builder pattern


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


## ProductApiScanConfigurationsList

> PaginatedProductAPIScanConfigurationList ProductApiScanConfigurationsList(ctx).Id(id).Limit(limit).Offset(offset).Product(product).ServiceKey1(serviceKey1).ServiceKey2(serviceKey2).ServiceKey3(serviceKey3).ToolConfiguration(toolConfiguration).Execute()



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
	product := int32(56) // int32 |  (optional)
	serviceKey1 := "serviceKey1_example" // string |  (optional)
	serviceKey2 := "serviceKey2_example" // string |  (optional)
	serviceKey3 := "serviceKey3_example" // string |  (optional)
	toolConfiguration := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsList(context.Background()).Id(id).Limit(limit).Offset(offset).Product(product).ServiceKey1(serviceKey1).ServiceKey2(serviceKey2).ServiceKey3(serviceKey3).ToolConfiguration(toolConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductApiScanConfigurationsList`: PaginatedProductAPIScanConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **product** | **int32** |  | 
 **serviceKey1** | **string** |  | 
 **serviceKey2** | **string** |  | 
 **serviceKey3** | **string** |  | 
 **toolConfiguration** | **int32** |  | 

### Return type

[**PaginatedProductAPIScanConfigurationList**](PaginatedProductAPIScanConfigurationList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductApiScanConfigurationsPartialUpdate

> ProductAPIScanConfiguration ProductApiScanConfigurationsPartialUpdate(ctx, id).PatchedProductAPIScanConfigurationRequest(patchedProductAPIScanConfigurationRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ap i_ scan_ configuration.
	patchedProductAPIScanConfigurationRequest := *openapiclient.NewPatchedProductAPIScanConfigurationRequest() // PatchedProductAPIScanConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsPartialUpdate(context.Background(), id).PatchedProductAPIScanConfigurationRequest(patchedProductAPIScanConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductApiScanConfigurationsPartialUpdate`: ProductAPIScanConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ap i_ scan_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedProductAPIScanConfigurationRequest** | [**PatchedProductAPIScanConfigurationRequest**](PatchedProductAPIScanConfigurationRequest.md) |  | 

### Return type

[**ProductAPIScanConfiguration**](ProductAPIScanConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductApiScanConfigurationsRetrieve

> ProductAPIScanConfiguration ProductApiScanConfigurationsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ap i_ scan_ configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductApiScanConfigurationsRetrieve`: ProductAPIScanConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ap i_ scan_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductAPIScanConfiguration**](ProductAPIScanConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductApiScanConfigurationsUpdate

> ProductAPIScanConfiguration ProductApiScanConfigurationsUpdate(ctx, id).ProductAPIScanConfigurationRequest(productAPIScanConfigurationRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product_ap i_ scan_ configuration.
	productAPIScanConfigurationRequest := *openapiclient.NewProductAPIScanConfigurationRequest(int32(123), int32(123)) // ProductAPIScanConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsUpdate(context.Background(), id).ProductAPIScanConfigurationRequest(productAPIScanConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductApiScanConfigurationsUpdate`: ProductAPIScanConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ProductApiScanConfigurationsAPI.ProductApiScanConfigurationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product_ap i_ scan_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductApiScanConfigurationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productAPIScanConfigurationRequest** | [**ProductAPIScanConfigurationRequest**](ProductAPIScanConfigurationRequest.md) |  | 

### Return type

[**ProductAPIScanConfiguration**](ProductAPIScanConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

