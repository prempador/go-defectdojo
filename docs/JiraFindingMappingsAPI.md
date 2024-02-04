# \JiraFindingMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JiraFindingMappingsCreate**](JiraFindingMappingsAPI.md#JiraFindingMappingsCreate) | **Post** /api/v2/jira_finding_mappings/ | 
[**JiraFindingMappingsDeletePreviewList**](JiraFindingMappingsAPI.md#JiraFindingMappingsDeletePreviewList) | **Get** /api/v2/jira_finding_mappings/{id}/delete_preview/ | 
[**JiraFindingMappingsDestroy**](JiraFindingMappingsAPI.md#JiraFindingMappingsDestroy) | **Delete** /api/v2/jira_finding_mappings/{id}/ | 
[**JiraFindingMappingsList**](JiraFindingMappingsAPI.md#JiraFindingMappingsList) | **Get** /api/v2/jira_finding_mappings/ | 
[**JiraFindingMappingsPartialUpdate**](JiraFindingMappingsAPI.md#JiraFindingMappingsPartialUpdate) | **Patch** /api/v2/jira_finding_mappings/{id}/ | 
[**JiraFindingMappingsRetrieve**](JiraFindingMappingsAPI.md#JiraFindingMappingsRetrieve) | **Get** /api/v2/jira_finding_mappings/{id}/ | 
[**JiraFindingMappingsUpdate**](JiraFindingMappingsAPI.md#JiraFindingMappingsUpdate) | **Put** /api/v2/jira_finding_mappings/{id}/ | 



## JiraFindingMappingsCreate

> JIRAIssue JiraFindingMappingsCreate(ctx).JIRAIssueRequest(jIRAIssueRequest).Execute()



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
	jIRAIssueRequest := *openapiclient.NewJIRAIssueRequest("JiraId_example", "JiraKey_example") // JIRAIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsCreate(context.Background()).JIRAIssueRequest(jIRAIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraFindingMappingsCreate`: JIRAIssue
	fmt.Fprintf(os.Stdout, "Response from `JiraFindingMappingsAPI.JiraFindingMappingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jIRAIssueRequest** | [**JIRAIssueRequest**](JIRAIssueRequest.md) |  | 

### Return type

[**JIRAIssue**](JIRAIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraFindingMappingsDeletePreviewList

> PaginatedDeletePreviewList JiraFindingMappingsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ issue.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraFindingMappingsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `JiraFindingMappingsAPI.JiraFindingMappingsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsDeletePreviewListRequest struct via the builder pattern


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


## JiraFindingMappingsDestroy

> JiraFindingMappingsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ issue.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsDestroyRequest struct via the builder pattern


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


## JiraFindingMappingsList

> PaginatedJIRAIssueList JiraFindingMappingsList(ctx).Engagement(engagement).Finding(finding).FindingGroup(findingGroup).Id(id).JiraId(jiraId).JiraKey(jiraKey).Limit(limit).Offset(offset).Execute()



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
	engagement := int32(56) // int32 |  (optional)
	finding := int32(56) // int32 |  (optional)
	findingGroup := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	jiraId := "jiraId_example" // string |  (optional)
	jiraKey := "jiraKey_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsList(context.Background()).Engagement(engagement).Finding(finding).FindingGroup(findingGroup).Id(id).JiraId(jiraId).JiraKey(jiraKey).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraFindingMappingsList`: PaginatedJIRAIssueList
	fmt.Fprintf(os.Stdout, "Response from `JiraFindingMappingsAPI.JiraFindingMappingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engagement** | **int32** |  | 
 **finding** | **int32** |  | 
 **findingGroup** | **int32** |  | 
 **id** | **int32** |  | 
 **jiraId** | **string** |  | 
 **jiraKey** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedJIRAIssueList**](PaginatedJIRAIssueList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraFindingMappingsPartialUpdate

> JIRAIssue JiraFindingMappingsPartialUpdate(ctx, id).PatchedJIRAIssueRequest(patchedJIRAIssueRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ issue.
	patchedJIRAIssueRequest := *openapiclient.NewPatchedJIRAIssueRequest() // PatchedJIRAIssueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsPartialUpdate(context.Background(), id).PatchedJIRAIssueRequest(patchedJIRAIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraFindingMappingsPartialUpdate`: JIRAIssue
	fmt.Fprintf(os.Stdout, "Response from `JiraFindingMappingsAPI.JiraFindingMappingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedJIRAIssueRequest** | [**PatchedJIRAIssueRequest**](PatchedJIRAIssueRequest.md) |  | 

### Return type

[**JIRAIssue**](JIRAIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraFindingMappingsRetrieve

> JIRAIssue JiraFindingMappingsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ issue.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraFindingMappingsRetrieve`: JIRAIssue
	fmt.Fprintf(os.Stdout, "Response from `JiraFindingMappingsAPI.JiraFindingMappingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JIRAIssue**](JIRAIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraFindingMappingsUpdate

> JIRAIssue JiraFindingMappingsUpdate(ctx, id).JIRAIssueRequest(jIRAIssueRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ issue.
	jIRAIssueRequest := *openapiclient.NewJIRAIssueRequest("JiraId_example", "JiraKey_example") // JIRAIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraFindingMappingsAPI.JiraFindingMappingsUpdate(context.Background(), id).JIRAIssueRequest(jIRAIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraFindingMappingsAPI.JiraFindingMappingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraFindingMappingsUpdate`: JIRAIssue
	fmt.Fprintf(os.Stdout, "Response from `JiraFindingMappingsAPI.JiraFindingMappingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraFindingMappingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jIRAIssueRequest** | [**JIRAIssueRequest**](JIRAIssueRequest.md) |  | 

### Return type

[**JIRAIssue**](JIRAIssue.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

