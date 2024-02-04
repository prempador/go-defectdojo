# \JiraProjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JiraProjectsCreate**](JiraProjectsAPI.md#JiraProjectsCreate) | **Post** /api/v2/jira_projects/ | 
[**JiraProjectsDeletePreviewList**](JiraProjectsAPI.md#JiraProjectsDeletePreviewList) | **Get** /api/v2/jira_projects/{id}/delete_preview/ | 
[**JiraProjectsDestroy**](JiraProjectsAPI.md#JiraProjectsDestroy) | **Delete** /api/v2/jira_projects/{id}/ | 
[**JiraProjectsList**](JiraProjectsAPI.md#JiraProjectsList) | **Get** /api/v2/jira_projects/ | 
[**JiraProjectsPartialUpdate**](JiraProjectsAPI.md#JiraProjectsPartialUpdate) | **Patch** /api/v2/jira_projects/{id}/ | 
[**JiraProjectsRetrieve**](JiraProjectsAPI.md#JiraProjectsRetrieve) | **Get** /api/v2/jira_projects/{id}/ | 
[**JiraProjectsUpdate**](JiraProjectsAPI.md#JiraProjectsUpdate) | **Put** /api/v2/jira_projects/{id}/ | 



## JiraProjectsCreate

> JIRAProject JiraProjectsCreate(ctx).JIRAProjectRequest(jIRAProjectRequest).Execute()



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
	jIRAProjectRequest := *openapiclient.NewJIRAProjectRequest() // JIRAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProjectsAPI.JiraProjectsCreate(context.Background()).JIRAProjectRequest(jIRAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProjectsCreate`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProjectsAPI.JiraProjectsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jIRAProjectRequest** | [**JIRAProjectRequest**](JIRAProjectRequest.md) |  | 

### Return type

[**JIRAProject**](JIRAProject.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraProjectsDeletePreviewList

> PaginatedDeletePreviewList JiraProjectsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ project.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProjectsAPI.JiraProjectsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProjectsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `JiraProjectsAPI.JiraProjectsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsDeletePreviewListRequest struct via the builder pattern


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


## JiraProjectsDestroy

> JiraProjectsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JiraProjectsAPI.JiraProjectsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsDestroyRequest struct via the builder pattern


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


## JiraProjectsList

> PaginatedJIRAProjectList JiraProjectsList(ctx).Component(component).EnableEngagementEpicMapping(enableEngagementEpicMapping).Engagement(engagement).Id(id).JiraInstance(jiraInstance).Limit(limit).Offset(offset).Product(product).ProjectKey(projectKey).PushAllIssues(pushAllIssues).PushNotes(pushNotes).Execute()



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
	component := "component_example" // string |  (optional)
	enableEngagementEpicMapping := true // bool |  (optional)
	engagement := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	jiraInstance := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	product := int32(56) // int32 |  (optional)
	projectKey := "projectKey_example" // string |  (optional)
	pushAllIssues := true // bool |  (optional)
	pushNotes := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProjectsAPI.JiraProjectsList(context.Background()).Component(component).EnableEngagementEpicMapping(enableEngagementEpicMapping).Engagement(engagement).Id(id).JiraInstance(jiraInstance).Limit(limit).Offset(offset).Product(product).ProjectKey(projectKey).PushAllIssues(pushAllIssues).PushNotes(pushNotes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProjectsList`: PaginatedJIRAProjectList
	fmt.Fprintf(os.Stdout, "Response from `JiraProjectsAPI.JiraProjectsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **component** | **string** |  | 
 **enableEngagementEpicMapping** | **bool** |  | 
 **engagement** | **int32** |  | 
 **id** | **int32** |  | 
 **jiraInstance** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **product** | **int32** |  | 
 **projectKey** | **string** |  | 
 **pushAllIssues** | **bool** |  | 
 **pushNotes** | **bool** |  | 

### Return type

[**PaginatedJIRAProjectList**](PaginatedJIRAProjectList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraProjectsPartialUpdate

> JIRAProject JiraProjectsPartialUpdate(ctx, id).PatchedJIRAProjectRequest(patchedJIRAProjectRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ project.
	patchedJIRAProjectRequest := *openapiclient.NewPatchedJIRAProjectRequest() // PatchedJIRAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProjectsAPI.JiraProjectsPartialUpdate(context.Background(), id).PatchedJIRAProjectRequest(patchedJIRAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProjectsPartialUpdate`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProjectsAPI.JiraProjectsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedJIRAProjectRequest** | [**PatchedJIRAProjectRequest**](PatchedJIRAProjectRequest.md) |  | 

### Return type

[**JIRAProject**](JIRAProject.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraProjectsRetrieve

> JIRAProject JiraProjectsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProjectsAPI.JiraProjectsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProjectsRetrieve`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProjectsAPI.JiraProjectsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JIRAProject**](JIRAProject.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JiraProjectsUpdate

> JIRAProject JiraProjectsUpdate(ctx, id).JIRAProjectRequest(jIRAProjectRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this jir a_ project.
	jIRAProjectRequest := *openapiclient.NewJIRAProjectRequest() // JIRAProjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProjectsAPI.JiraProjectsUpdate(context.Background(), id).JIRAProjectRequest(jIRAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProjectsAPI.JiraProjectsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProjectsUpdate`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProjectsAPI.JiraProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jIRAProjectRequest** | [**JIRAProjectRequest**](JIRAProjectRequest.md) |  | 

### Return type

[**JIRAProject**](JIRAProject.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

