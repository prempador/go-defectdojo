# \SonarqubeIssuesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SonarqubeIssuesCreate**](SonarqubeIssuesAPI.md#SonarqubeIssuesCreate) | **Post** /api/v2/sonarqube_issues/ | 
[**SonarqubeIssuesDeletePreviewList**](SonarqubeIssuesAPI.md#SonarqubeIssuesDeletePreviewList) | **Get** /api/v2/sonarqube_issues/{id}/delete_preview/ | 
[**SonarqubeIssuesDestroy**](SonarqubeIssuesAPI.md#SonarqubeIssuesDestroy) | **Delete** /api/v2/sonarqube_issues/{id}/ | 
[**SonarqubeIssuesList**](SonarqubeIssuesAPI.md#SonarqubeIssuesList) | **Get** /api/v2/sonarqube_issues/ | 
[**SonarqubeIssuesPartialUpdate**](SonarqubeIssuesAPI.md#SonarqubeIssuesPartialUpdate) | **Patch** /api/v2/sonarqube_issues/{id}/ | 
[**SonarqubeIssuesRetrieve**](SonarqubeIssuesAPI.md#SonarqubeIssuesRetrieve) | **Get** /api/v2/sonarqube_issues/{id}/ | 
[**SonarqubeIssuesUpdate**](SonarqubeIssuesAPI.md#SonarqubeIssuesUpdate) | **Put** /api/v2/sonarqube_issues/{id}/ | 



## SonarqubeIssuesCreate

> SonarqubeIssue SonarqubeIssuesCreate(ctx).SonarqubeIssueRequest(sonarqubeIssueRequest).Execute()



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
	sonarqubeIssueRequest := *openapiclient.NewSonarqubeIssueRequest("Key_example", "Status_example", "Type_example") // SonarqubeIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesCreate(context.Background()).SonarqubeIssueRequest(sonarqubeIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeIssuesCreate`: SonarqubeIssue
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeIssuesAPI.SonarqubeIssuesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sonarqubeIssueRequest** | [**SonarqubeIssueRequest**](SonarqubeIssueRequest.md) |  | 

### Return type

[**SonarqubeIssue**](SonarqubeIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeIssuesDeletePreviewList

> PaginatedDeletePreviewList SonarqubeIssuesDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeIssuesDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeIssuesAPI.SonarqubeIssuesDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesDeletePreviewListRequest struct via the builder pattern


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


## SonarqubeIssuesDestroy

> SonarqubeIssuesDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesDestroyRequest struct via the builder pattern


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


## SonarqubeIssuesList

> PaginatedSonarqubeIssueList SonarqubeIssuesList(ctx).Id(id).Key(key).Limit(limit).Offset(offset).Status(status).Type_(type_).Execute()



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
	key := "key_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	status := "status_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesList(context.Background()).Id(id).Key(key).Limit(limit).Offset(offset).Status(status).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeIssuesList`: PaginatedSonarqubeIssueList
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeIssuesAPI.SonarqubeIssuesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **key** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **status** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**PaginatedSonarqubeIssueList**](PaginatedSonarqubeIssueList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeIssuesPartialUpdate

> SonarqubeIssue SonarqubeIssuesPartialUpdate(ctx, id).PatchedSonarqubeIssueRequest(patchedSonarqubeIssueRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue.
	patchedSonarqubeIssueRequest := *openapiclient.NewPatchedSonarqubeIssueRequest() // PatchedSonarqubeIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesPartialUpdate(context.Background(), id).PatchedSonarqubeIssueRequest(patchedSonarqubeIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeIssuesPartialUpdate`: SonarqubeIssue
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeIssuesAPI.SonarqubeIssuesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSonarqubeIssueRequest** | [**PatchedSonarqubeIssueRequest**](PatchedSonarqubeIssueRequest.md) |  | 

### Return type

[**SonarqubeIssue**](SonarqubeIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeIssuesRetrieve

> SonarqubeIssue SonarqubeIssuesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeIssuesRetrieve`: SonarqubeIssue
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeIssuesAPI.SonarqubeIssuesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SonarqubeIssue**](SonarqubeIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeIssuesUpdate

> SonarqubeIssue SonarqubeIssuesUpdate(ctx, id).SonarqubeIssueRequest(sonarqubeIssueRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue.
	sonarqubeIssueRequest := *openapiclient.NewSonarqubeIssueRequest("Key_example", "Status_example", "Type_example") // SonarqubeIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeIssuesAPI.SonarqubeIssuesUpdate(context.Background(), id).SonarqubeIssueRequest(sonarqubeIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeIssuesAPI.SonarqubeIssuesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeIssuesUpdate`: SonarqubeIssue
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeIssuesAPI.SonarqubeIssuesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeIssuesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sonarqubeIssueRequest** | [**SonarqubeIssueRequest**](SonarqubeIssueRequest.md) |  | 

### Return type

[**SonarqubeIssue**](SonarqubeIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

