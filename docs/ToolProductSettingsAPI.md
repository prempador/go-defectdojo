# \ToolProductSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ToolProductSettingsCreate**](ToolProductSettingsAPI.md#ToolProductSettingsCreate) | **Post** /api/v2/tool_product_settings/ | 
[**ToolProductSettingsDeletePreviewList**](ToolProductSettingsAPI.md#ToolProductSettingsDeletePreviewList) | **Get** /api/v2/tool_product_settings/{id}/delete_preview/ | 
[**ToolProductSettingsDestroy**](ToolProductSettingsAPI.md#ToolProductSettingsDestroy) | **Delete** /api/v2/tool_product_settings/{id}/ | 
[**ToolProductSettingsList**](ToolProductSettingsAPI.md#ToolProductSettingsList) | **Get** /api/v2/tool_product_settings/ | 
[**ToolProductSettingsPartialUpdate**](ToolProductSettingsAPI.md#ToolProductSettingsPartialUpdate) | **Patch** /api/v2/tool_product_settings/{id}/ | 
[**ToolProductSettingsRetrieve**](ToolProductSettingsAPI.md#ToolProductSettingsRetrieve) | **Get** /api/v2/tool_product_settings/{id}/ | 
[**ToolProductSettingsUpdate**](ToolProductSettingsAPI.md#ToolProductSettingsUpdate) | **Put** /api/v2/tool_product_settings/{id}/ | 



## ToolProductSettingsCreate

> ToolProductSettings ToolProductSettingsCreate(ctx).ToolProductSettingsRequest(toolProductSettingsRequest).Execute()



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
	toolProductSettingsRequest := *openapiclient.NewToolProductSettingsRequest("SettingUrl_example", int32(123), "Name_example", int32(123)) // ToolProductSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsCreate(context.Background()).ToolProductSettingsRequest(toolProductSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolProductSettingsCreate`: ToolProductSettings
	fmt.Fprintf(os.Stdout, "Response from `ToolProductSettingsAPI.ToolProductSettingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolProductSettingsRequest** | [**ToolProductSettingsRequest**](ToolProductSettingsRequest.md) |  | 

### Return type

[**ToolProductSettings**](ToolProductSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolProductSettingsDeletePreviewList

> PaginatedDeletePreviewList ToolProductSettingsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ product_ settings.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolProductSettingsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `ToolProductSettingsAPI.ToolProductSettingsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ product_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsDeletePreviewListRequest struct via the builder pattern


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


## ToolProductSettingsDestroy

> ToolProductSettingsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ product_ settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ product_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsDestroyRequest struct via the builder pattern


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


## ToolProductSettingsList

> PaginatedToolProductSettingsList ToolProductSettingsList(ctx).Id(id).Limit(limit).Name(name).Offset(offset).Product(product).ToolConfiguration(toolConfiguration).ToolProjectId(toolProjectId).Url(url).Execute()



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
	name := "name_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	product := int32(56) // int32 |  (optional)
	toolConfiguration := int32(56) // int32 |  (optional)
	toolProjectId := "toolProjectId_example" // string |  (optional)
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsList(context.Background()).Id(id).Limit(limit).Name(name).Offset(offset).Product(product).ToolConfiguration(toolConfiguration).ToolProjectId(toolProjectId).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolProductSettingsList`: PaginatedToolProductSettingsList
	fmt.Fprintf(os.Stdout, "Response from `ToolProductSettingsAPI.ToolProductSettingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **product** | **int32** |  | 
 **toolConfiguration** | **int32** |  | 
 **toolProjectId** | **string** |  | 
 **url** | **string** |  | 

### Return type

[**PaginatedToolProductSettingsList**](PaginatedToolProductSettingsList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolProductSettingsPartialUpdate

> ToolProductSettings ToolProductSettingsPartialUpdate(ctx, id).PatchedToolProductSettingsRequest(patchedToolProductSettingsRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ product_ settings.
	patchedToolProductSettingsRequest := *openapiclient.NewPatchedToolProductSettingsRequest() // PatchedToolProductSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsPartialUpdate(context.Background(), id).PatchedToolProductSettingsRequest(patchedToolProductSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolProductSettingsPartialUpdate`: ToolProductSettings
	fmt.Fprintf(os.Stdout, "Response from `ToolProductSettingsAPI.ToolProductSettingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ product_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedToolProductSettingsRequest** | [**PatchedToolProductSettingsRequest**](PatchedToolProductSettingsRequest.md) |  | 

### Return type

[**ToolProductSettings**](ToolProductSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolProductSettingsRetrieve

> ToolProductSettings ToolProductSettingsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ product_ settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolProductSettingsRetrieve`: ToolProductSettings
	fmt.Fprintf(os.Stdout, "Response from `ToolProductSettingsAPI.ToolProductSettingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ product_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ToolProductSettings**](ToolProductSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolProductSettingsUpdate

> ToolProductSettings ToolProductSettingsUpdate(ctx, id).ToolProductSettingsRequest(toolProductSettingsRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ product_ settings.
	toolProductSettingsRequest := *openapiclient.NewToolProductSettingsRequest("SettingUrl_example", int32(123), "Name_example", int32(123)) // ToolProductSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolProductSettingsAPI.ToolProductSettingsUpdate(context.Background(), id).ToolProductSettingsRequest(toolProductSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolProductSettingsAPI.ToolProductSettingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolProductSettingsUpdate`: ToolProductSettings
	fmt.Fprintf(os.Stdout, "Response from `ToolProductSettingsAPI.ToolProductSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ product_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolProductSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolProductSettingsRequest** | [**ToolProductSettingsRequest**](ToolProductSettingsRequest.md) |  | 

### Return type

[**ToolProductSettings**](ToolProductSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

