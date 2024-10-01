# \JiraProductConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JiraProductConfigurationsCreate**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsCreate) | **Post** /api/v2/jira_product_configurations/ | 
[**JiraProductConfigurationsDeletePreviewList**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsDeletePreviewList) | **Get** /api/v2/jira_product_configurations/{id}/delete_preview/ | 
[**JiraProductConfigurationsDestroy**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsDestroy) | **Delete** /api/v2/jira_product_configurations/{id}/ | 
[**JiraProductConfigurationsList**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsList) | **Get** /api/v2/jira_product_configurations/ | 
[**JiraProductConfigurationsPartialUpdate**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsPartialUpdate) | **Patch** /api/v2/jira_product_configurations/{id}/ | 
[**JiraProductConfigurationsRetrieve**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsRetrieve) | **Get** /api/v2/jira_product_configurations/{id}/ | 
[**JiraProductConfigurationsUpdate**](JiraProductConfigurationsAPI.md#JiraProductConfigurationsUpdate) | **Put** /api/v2/jira_product_configurations/{id}/ | 



## JiraProductConfigurationsCreate

> JIRAProject JiraProductConfigurationsCreate(ctx).JIRAProjectRequest(jIRAProjectRequest).Execute()



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
	resp, r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsCreate(context.Background()).JIRAProjectRequest(jIRAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProductConfigurationsCreate`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProductConfigurationsAPI.JiraProductConfigurationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraProductConfigurationsCreateRequest struct via the builder pattern


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


## JiraProductConfigurationsDeletePreviewList

> PaginatedDeletePreviewList JiraProductConfigurationsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProductConfigurationsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `JiraProductConfigurationsAPI.JiraProductConfigurationsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProductConfigurationsDeletePreviewListRequest struct via the builder pattern


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


## JiraProductConfigurationsDestroy

> JiraProductConfigurationsDestroy(ctx, id).Execute()



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
	r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiJiraProductConfigurationsDestroyRequest struct via the builder pattern


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


## JiraProductConfigurationsList

> PaginatedJIRAProjectList JiraProductConfigurationsList(ctx).Component(component).EnableEngagementEpicMapping(enableEngagementEpicMapping).Engagement(engagement).Id(id).JiraInstance(jiraInstance).Limit(limit).Offset(offset).Prefetch(prefetch).Product(product).ProjectKey(projectKey).PushAllIssues(pushAllIssues).PushNotes(pushNotes).Execute()



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
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	product := int32(56) // int32 |  (optional)
	projectKey := "projectKey_example" // string |  (optional)
	pushAllIssues := true // bool |  (optional)
	pushNotes := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsList(context.Background()).Component(component).EnableEngagementEpicMapping(enableEngagementEpicMapping).Engagement(engagement).Id(id).JiraInstance(jiraInstance).Limit(limit).Offset(offset).Prefetch(prefetch).Product(product).ProjectKey(projectKey).PushAllIssues(pushAllIssues).PushNotes(pushNotes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProductConfigurationsList`: PaginatedJIRAProjectList
	fmt.Fprintf(os.Stdout, "Response from `JiraProductConfigurationsAPI.JiraProductConfigurationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJiraProductConfigurationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **component** | **string** |  | 
 **enableEngagementEpicMapping** | **bool** |  | 
 **engagement** | **int32** |  | 
 **id** | **int32** |  | 
 **jiraInstance** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
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


## JiraProductConfigurationsPartialUpdate

> JIRAProject JiraProductConfigurationsPartialUpdate(ctx, id).PatchedJIRAProjectRequest(patchedJIRAProjectRequest).Execute()



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
	resp, r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsPartialUpdate(context.Background(), id).PatchedJIRAProjectRequest(patchedJIRAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProductConfigurationsPartialUpdate`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProductConfigurationsAPI.JiraProductConfigurationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProductConfigurationsPartialUpdateRequest struct via the builder pattern


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


## JiraProductConfigurationsRetrieve

> JIRAProject JiraProductConfigurationsRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProductConfigurationsRetrieve`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProductConfigurationsAPI.JiraProductConfigurationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProductConfigurationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

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


## JiraProductConfigurationsUpdate

> JIRAProject JiraProductConfigurationsUpdate(ctx, id).JIRAProjectRequest(jIRAProjectRequest).Execute()



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
	resp, r, err := apiClient.JiraProductConfigurationsAPI.JiraProductConfigurationsUpdate(context.Background(), id).JIRAProjectRequest(jIRAProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraProductConfigurationsAPI.JiraProductConfigurationsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JiraProductConfigurationsUpdate`: JIRAProject
	fmt.Fprintf(os.Stdout, "Response from `JiraProductConfigurationsAPI.JiraProductConfigurationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this jir a_ project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJiraProductConfigurationsUpdateRequest struct via the builder pattern


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

