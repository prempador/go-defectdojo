# \JiraConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JiraConfigurationsCreate**](JiraConfigurationsAPI.md#JiraConfigurationsCreate) | **Post** /api/v2/jira_configurations/ | 
[**JiraConfigurationsDeletePreviewList**](JiraConfigurationsAPI.md#JiraConfigurationsDeletePreviewList) | **Get** /api/v2/jira_configurations/{id}/delete_preview/ | 
[**JiraConfigurationsDestroy**](JiraConfigurationsAPI.md#JiraConfigurationsDestroy) | **Delete** /api/v2/jira_configurations/{id}/ | 
[**JiraConfigurationsList**](JiraConfigurationsAPI.md#JiraConfigurationsList) | **Get** /api/v2/jira_configurations/ | 
[**JiraConfigurationsPartialUpdate**](JiraConfigurationsAPI.md#JiraConfigurationsPartialUpdate) | **Patch** /api/v2/jira_configurations/{id}/ | 
[**JiraConfigurationsRetrieve**](JiraConfigurationsAPI.md#JiraConfigurationsRetrieve) | **Get** /api/v2/jira_configurations/{id}/ | 
[**JiraConfigurationsUpdate**](JiraConfigurationsAPI.md#JiraConfigurationsUpdate) | **Put** /api/v2/jira_configurations/{id}/ | 



## JiraConfigurationsCreate

> JIRAInstance JiraConfigurationsCreate(ctx).JIRAInstanceRequest(jIRAInstanceRequest).Execute()



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
	jIRAInstanceRequest := *openapiclient.NewJIRAInstanceRequest("Url_example", "Username_example", "Password_example", int32(123), int32(123), int32(123), "InfoMappingSeverity_example", "LowMappingSeverity_example", "MediumMappingSeverity_example", "HighMappingSeverity_example", "CriticalMappingSeverity_example") // JIRAInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsCreate(context.Background()).JIRAInstanceRequest(jIRAInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraConfigurationsCreate`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraConfigurationsAPI.JiraConfigurationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jIRAInstanceRequest** | [**JIRAInstanceRequest**](JIRAInstanceRequest.md) |  | 

### Return type

[**JIRAInstance**](JIRAInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraConfigurationsDeletePreviewList

> PaginatedDeletePreviewList JiraConfigurationsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ instance.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraConfigurationsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `JiraConfigurationsAPI.JiraConfigurationsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsDeletePreviewListRequest struct via the builder pattern


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


## JiraConfigurationsDestroy

> JiraConfigurationsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsDestroyRequest struct via the builder pattern


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


## JiraConfigurationsList

> PaginatedJIRAInstanceList JiraConfigurationsList(ctx).Id(id).Limit(limit).Offset(offset).Url(url).Execute()



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
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsList(context.Background()).Id(id).Limit(limit).Offset(offset).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraConfigurationsList`: PaginatedJIRAInstanceList
	fmt.Fprintf(os.Stdout, "Response from `JiraConfigurationsAPI.JiraConfigurationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **url** | **string** |  | 

### Return type

[**PaginatedJIRAInstanceList**](PaginatedJIRAInstanceList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraConfigurationsPartialUpdate

> JIRAInstance JiraConfigurationsPartialUpdate(ctx, id).PatchedJIRAInstanceRequest(patchedJIRAInstanceRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ instance.
	patchedJIRAInstanceRequest := *openapiclient.NewPatchedJIRAInstanceRequest() // PatchedJIRAInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsPartialUpdate(context.Background(), id).PatchedJIRAInstanceRequest(patchedJIRAInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraConfigurationsPartialUpdate`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraConfigurationsAPI.JiraConfigurationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedJIRAInstanceRequest** | [**PatchedJIRAInstanceRequest**](PatchedJIRAInstanceRequest.md) |  | 

### Return type

[**JIRAInstance**](JIRAInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraConfigurationsRetrieve

> JIRAInstance JiraConfigurationsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ instance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraConfigurationsRetrieve`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraConfigurationsAPI.JiraConfigurationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JIRAInstance**](JIRAInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraConfigurationsUpdate

> JIRAInstance JiraConfigurationsUpdate(ctx, id).JIRAInstanceRequest(jIRAInstanceRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ instance.
	jIRAInstanceRequest := *openapiclient.NewJIRAInstanceRequest("Url_example", "Username_example", "Password_example", int32(123), int32(123), int32(123), "InfoMappingSeverity_example", "LowMappingSeverity_example", "MediumMappingSeverity_example", "HighMappingSeverity_example", "CriticalMappingSeverity_example") // JIRAInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraConfigurationsAPI.JiraConfigurationsUpdate(context.Background(), id).JIRAInstanceRequest(jIRAInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraConfigurationsAPI.JiraConfigurationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraConfigurationsUpdate`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraConfigurationsAPI.JiraConfigurationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraConfigurationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jIRAInstanceRequest** | [**JIRAInstanceRequest**](JIRAInstanceRequest.md) |  | 

### Return type

[**JIRAInstance**](JIRAInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

