# \JiraInstancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JiraInstancesCreate**](JiraInstancesAPI.md#JiraInstancesCreate) | **Post** /api/v2/jira_instances/ | 
[**JiraInstancesDeletePreviewList**](JiraInstancesAPI.md#JiraInstancesDeletePreviewList) | **Get** /api/v2/jira_instances/{id}/delete_preview/ | 
[**JiraInstancesDestroy**](JiraInstancesAPI.md#JiraInstancesDestroy) | **Delete** /api/v2/jira_instances/{id}/ | 
[**JiraInstancesList**](JiraInstancesAPI.md#JiraInstancesList) | **Get** /api/v2/jira_instances/ | 
[**JiraInstancesPartialUpdate**](JiraInstancesAPI.md#JiraInstancesPartialUpdate) | **Patch** /api/v2/jira_instances/{id}/ | 
[**JiraInstancesRetrieve**](JiraInstancesAPI.md#JiraInstancesRetrieve) | **Get** /api/v2/jira_instances/{id}/ | 
[**JiraInstancesUpdate**](JiraInstancesAPI.md#JiraInstancesUpdate) | **Put** /api/v2/jira_instances/{id}/ | 



## JiraInstancesCreate

> JIRAInstance JiraInstancesCreate(ctx).JIRAInstanceRequest(jIRAInstanceRequest).Execute()



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
	resp, r, err := apiClient.JiraInstancesAPI.JiraInstancesCreate(context.Background()).JIRAInstanceRequest(jIRAInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraInstancesCreate`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraInstancesAPI.JiraInstancesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraInstancesCreateRequest struct via the builder pattern


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


## JiraInstancesDeletePreviewList

> PaginatedDeletePreviewList JiraInstancesDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.JiraInstancesAPI.JiraInstancesDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraInstancesDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `JiraInstancesAPI.JiraInstancesDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraInstancesDeletePreviewListRequest struct via the builder pattern


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


## JiraInstancesDestroy

> JiraInstancesDestroy(ctx, id).Execute()



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
	r, err := apiClient.JiraInstancesAPI.JiraInstancesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiJiraInstancesDestroyRequest struct via the builder pattern


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


## JiraInstancesList

> PaginatedJIRAInstanceList JiraInstancesList(ctx).Id(id).Limit(limit).Offset(offset).Url(url).Execute()



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
	resp, r, err := apiClient.JiraInstancesAPI.JiraInstancesList(context.Background()).Id(id).Limit(limit).Offset(offset).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraInstancesList`: PaginatedJIRAInstanceList
	fmt.Fprintf(os.Stdout, "Response from `JiraInstancesAPI.JiraInstancesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraInstancesListRequest struct via the builder pattern


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


## JiraInstancesPartialUpdate

> JIRAInstance JiraInstancesPartialUpdate(ctx, id).PatchedJIRAInstanceRequest(patchedJIRAInstanceRequest).Execute()



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
	resp, r, err := apiClient.JiraInstancesAPI.JiraInstancesPartialUpdate(context.Background(), id).PatchedJIRAInstanceRequest(patchedJIRAInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraInstancesPartialUpdate`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraInstancesAPI.JiraInstancesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraInstancesPartialUpdateRequest struct via the builder pattern


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


## JiraInstancesRetrieve

> JIRAInstance JiraInstancesRetrieve(ctx, id).Execute()



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
	resp, r, err := apiClient.JiraInstancesAPI.JiraInstancesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraInstancesRetrieve`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraInstancesAPI.JiraInstancesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraInstancesRetrieveRequest struct via the builder pattern


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


## JiraInstancesUpdate

> JIRAInstance JiraInstancesUpdate(ctx, id).JIRAInstanceRequest(jIRAInstanceRequest).Execute()



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
	resp, r, err := apiClient.JiraInstancesAPI.JiraInstancesUpdate(context.Background(), id).JIRAInstanceRequest(jIRAInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraInstancesAPI.JiraInstancesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraInstancesUpdate`: JIRAInstance
	fmt.Fprintf(os.Stdout, "Response from `JiraInstancesAPI.JiraInstancesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraInstancesUpdateRequest struct via the builder pattern


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

