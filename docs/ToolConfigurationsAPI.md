# \ToolConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ToolConfigurationsCreate**](ToolConfigurationsAPI.md#ToolConfigurationsCreate) | **Post** /api/v2/tool_configurations/ | 
[**ToolConfigurationsDeletePreviewList**](ToolConfigurationsAPI.md#ToolConfigurationsDeletePreviewList) | **Get** /api/v2/tool_configurations/{id}/delete_preview/ | 
[**ToolConfigurationsDestroy**](ToolConfigurationsAPI.md#ToolConfigurationsDestroy) | **Delete** /api/v2/tool_configurations/{id}/ | 
[**ToolConfigurationsList**](ToolConfigurationsAPI.md#ToolConfigurationsList) | **Get** /api/v2/tool_configurations/ | 
[**ToolConfigurationsPartialUpdate**](ToolConfigurationsAPI.md#ToolConfigurationsPartialUpdate) | **Patch** /api/v2/tool_configurations/{id}/ | 
[**ToolConfigurationsRetrieve**](ToolConfigurationsAPI.md#ToolConfigurationsRetrieve) | **Get** /api/v2/tool_configurations/{id}/ | 
[**ToolConfigurationsUpdate**](ToolConfigurationsAPI.md#ToolConfigurationsUpdate) | **Put** /api/v2/tool_configurations/{id}/ | 



## ToolConfigurationsCreate

> ToolConfiguration ToolConfigurationsCreate(ctx).ToolConfigurationRequest(toolConfigurationRequest).Execute()



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
	toolConfigurationRequest := *openapiclient.NewToolConfigurationRequest("Name_example", int32(123)) // ToolConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsCreate(context.Background()).ToolConfigurationRequest(toolConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolConfigurationsCreate`: ToolConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ToolConfigurationsAPI.ToolConfigurationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **toolConfigurationRequest** | [**ToolConfigurationRequest**](ToolConfigurationRequest.md) |  | 

### Return type

[**ToolConfiguration**](ToolConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolConfigurationsDeletePreviewList

> PaginatedDeletePreviewList ToolConfigurationsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ configuration.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolConfigurationsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `ToolConfigurationsAPI.ToolConfigurationsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsDeletePreviewListRequest struct via the builder pattern


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


## ToolConfigurationsDestroy

> ToolConfigurationsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsDestroyRequest struct via the builder pattern


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


## ToolConfigurationsList

> PaginatedToolConfigurationList ToolConfigurationsList(ctx).AuthenticationType(authenticationType).Id(id).Limit(limit).Name(name).Offset(offset).Prefetch(prefetch).ToolType(toolType).Url(url).Execute()



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
	authenticationType := "authenticationType_example" // string | * `API` - API Key * `Password` - Username/Password * `SSH` - SSH (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	toolType := int32(56) // int32 |  (optional)
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsList(context.Background()).AuthenticationType(authenticationType).Id(id).Limit(limit).Name(name).Offset(offset).Prefetch(prefetch).ToolType(toolType).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolConfigurationsList`: PaginatedToolConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `ToolConfigurationsAPI.ToolConfigurationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationType** | **string** | * &#x60;API&#x60; - API Key * &#x60;Password&#x60; - Username/Password * &#x60;SSH&#x60; - SSH | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **toolType** | **int32** |  | 
 **url** | **string** |  | 

### Return type

[**PaginatedToolConfigurationList**](PaginatedToolConfigurationList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolConfigurationsPartialUpdate

> ToolConfiguration ToolConfigurationsPartialUpdate(ctx, id).PatchedToolConfigurationRequest(patchedToolConfigurationRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ configuration.
	patchedToolConfigurationRequest := *openapiclient.NewPatchedToolConfigurationRequest() // PatchedToolConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsPartialUpdate(context.Background(), id).PatchedToolConfigurationRequest(patchedToolConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolConfigurationsPartialUpdate`: ToolConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ToolConfigurationsAPI.ToolConfigurationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedToolConfigurationRequest** | [**PatchedToolConfigurationRequest**](PatchedToolConfigurationRequest.md) |  | 

### Return type

[**ToolConfiguration**](ToolConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolConfigurationsRetrieve

> ToolConfiguration ToolConfigurationsRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ configuration.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolConfigurationsRetrieve`: ToolConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ToolConfigurationsAPI.ToolConfigurationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**ToolConfiguration**](ToolConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToolConfigurationsUpdate

> ToolConfiguration ToolConfigurationsUpdate(ctx, id).ToolConfigurationRequest(toolConfigurationRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this tool_ configuration.
	toolConfigurationRequest := *openapiclient.NewToolConfigurationRequest("Name_example", int32(123)) // ToolConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolConfigurationsAPI.ToolConfigurationsUpdate(context.Background(), id).ToolConfigurationRequest(toolConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolConfigurationsAPI.ToolConfigurationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToolConfigurationsUpdate`: ToolConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ToolConfigurationsAPI.ToolConfigurationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this tool_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToolConfigurationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toolConfigurationRequest** | [**ToolConfigurationRequest**](ToolConfigurationRequest.md) |  | 

### Return type

[**ToolConfiguration**](ToolConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

