# \SlaConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlaConfigurationsCreate**](SlaConfigurationsAPI.md#SlaConfigurationsCreate) | **Post** /api/v2/sla_configurations/ | 
[**SlaConfigurationsDeletePreviewList**](SlaConfigurationsAPI.md#SlaConfigurationsDeletePreviewList) | **Get** /api/v2/sla_configurations/{id}/delete_preview/ | 
[**SlaConfigurationsDestroy**](SlaConfigurationsAPI.md#SlaConfigurationsDestroy) | **Delete** /api/v2/sla_configurations/{id}/ | 
[**SlaConfigurationsList**](SlaConfigurationsAPI.md#SlaConfigurationsList) | **Get** /api/v2/sla_configurations/ | 
[**SlaConfigurationsPartialUpdate**](SlaConfigurationsAPI.md#SlaConfigurationsPartialUpdate) | **Patch** /api/v2/sla_configurations/{id}/ | 
[**SlaConfigurationsRetrieve**](SlaConfigurationsAPI.md#SlaConfigurationsRetrieve) | **Get** /api/v2/sla_configurations/{id}/ | 
[**SlaConfigurationsUpdate**](SlaConfigurationsAPI.md#SlaConfigurationsUpdate) | **Put** /api/v2/sla_configurations/{id}/ | 



## SlaConfigurationsCreate

> SLAConfiguration SlaConfigurationsCreate(ctx).SLAConfigurationRequest(sLAConfigurationRequest).Execute()



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
	sLAConfigurationRequest := *openapiclient.NewSLAConfigurationRequest("Name_example") // SLAConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsCreate(context.Background()).SLAConfigurationRequest(sLAConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlaConfigurationsCreate`: SLAConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SlaConfigurationsAPI.SlaConfigurationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sLAConfigurationRequest** | [**SLAConfigurationRequest**](SLAConfigurationRequest.md) |  | 

### Return type

[**SLAConfiguration**](SLAConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaConfigurationsDeletePreviewList

> PaginatedDeletePreviewList SlaConfigurationsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sl a_ configuration.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlaConfigurationsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `SlaConfigurationsAPI.SlaConfigurationsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sl a_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsDeletePreviewListRequest struct via the builder pattern


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


## SlaConfigurationsDestroy

> SlaConfigurationsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sl a_ configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sl a_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsDestroyRequest struct via the builder pattern


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


## SlaConfigurationsList

> PaginatedSLAConfigurationList SlaConfigurationsList(ctx).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlaConfigurationsList`: PaginatedSLAConfigurationList
	fmt.Fprintf(os.Stdout, "Response from `SlaConfigurationsAPI.SlaConfigurationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedSLAConfigurationList**](PaginatedSLAConfigurationList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaConfigurationsPartialUpdate

> SLAConfiguration SlaConfigurationsPartialUpdate(ctx, id).PatchedSLAConfigurationRequest(patchedSLAConfigurationRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sl a_ configuration.
	patchedSLAConfigurationRequest := *openapiclient.NewPatchedSLAConfigurationRequest() // PatchedSLAConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsPartialUpdate(context.Background(), id).PatchedSLAConfigurationRequest(patchedSLAConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlaConfigurationsPartialUpdate`: SLAConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SlaConfigurationsAPI.SlaConfigurationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sl a_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSLAConfigurationRequest** | [**PatchedSLAConfigurationRequest**](PatchedSLAConfigurationRequest.md) |  | 

### Return type

[**SLAConfiguration**](SLAConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaConfigurationsRetrieve

> SLAConfiguration SlaConfigurationsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sl a_ configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlaConfigurationsRetrieve`: SLAConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SlaConfigurationsAPI.SlaConfigurationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sl a_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SLAConfiguration**](SLAConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlaConfigurationsUpdate

> SLAConfiguration SlaConfigurationsUpdate(ctx, id).SLAConfigurationRequest(sLAConfigurationRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sl a_ configuration.
	sLAConfigurationRequest := *openapiclient.NewSLAConfigurationRequest("Name_example") // SLAConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlaConfigurationsAPI.SlaConfigurationsUpdate(context.Background(), id).SLAConfigurationRequest(sLAConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlaConfigurationsAPI.SlaConfigurationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlaConfigurationsUpdate`: SLAConfiguration
	fmt.Fprintf(os.Stdout, "Response from `SlaConfigurationsAPI.SlaConfigurationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sl a_ configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlaConfigurationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sLAConfigurationRequest** | [**SLAConfigurationRequest**](SLAConfigurationRequest.md) |  | 

### Return type

[**SLAConfiguration**](SLAConfiguration.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

