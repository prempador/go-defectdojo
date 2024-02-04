# \SonarqubeTransitionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SonarqubeTransitionsCreate**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsCreate) | **Post** /api/v2/sonarqube_transitions/ | 
[**SonarqubeTransitionsDeletePreviewList**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsDeletePreviewList) | **Get** /api/v2/sonarqube_transitions/{id}/delete_preview/ | 
[**SonarqubeTransitionsDestroy**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsDestroy) | **Delete** /api/v2/sonarqube_transitions/{id}/ | 
[**SonarqubeTransitionsList**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsList) | **Get** /api/v2/sonarqube_transitions/ | 
[**SonarqubeTransitionsPartialUpdate**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsPartialUpdate) | **Patch** /api/v2/sonarqube_transitions/{id}/ | 
[**SonarqubeTransitionsRetrieve**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsRetrieve) | **Get** /api/v2/sonarqube_transitions/{id}/ | 
[**SonarqubeTransitionsUpdate**](SonarqubeTransitionsAPI.md#SonarqubeTransitionsUpdate) | **Put** /api/v2/sonarqube_transitions/{id}/ | 



## SonarqubeTransitionsCreate

> SonarqubeIssueTransition SonarqubeTransitionsCreate(ctx).SonarqubeIssueTransitionRequest(sonarqubeIssueTransitionRequest).Execute()



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
	sonarqubeIssueTransitionRequest := *openapiclient.NewSonarqubeIssueTransitionRequest("FindingStatus_example", "SonarqubeStatus_example", "Transitions_example", int32(123)) // SonarqubeIssueTransitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsCreate(context.Background()).SonarqubeIssueTransitionRequest(sonarqubeIssueTransitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeTransitionsCreate`: SonarqubeIssueTransition
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeTransitionsAPI.SonarqubeTransitionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sonarqubeIssueTransitionRequest** | [**SonarqubeIssueTransitionRequest**](SonarqubeIssueTransitionRequest.md) |  | 

### Return type

[**SonarqubeIssueTransition**](SonarqubeIssueTransition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeTransitionsDeletePreviewList

> PaginatedDeletePreviewList SonarqubeTransitionsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue_ transition.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeTransitionsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeTransitionsAPI.SonarqubeTransitionsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue_ transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsDeletePreviewListRequest struct via the builder pattern


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


## SonarqubeTransitionsDestroy

> SonarqubeTransitionsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue_ transition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue_ transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsDestroyRequest struct via the builder pattern


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


## SonarqubeTransitionsList

> PaginatedSonarqubeIssueTransitionList SonarqubeTransitionsList(ctx).FindingStatus(findingStatus).Id(id).Limit(limit).Offset(offset).SonarqubeIssue(sonarqubeIssue).SonarqubeStatus(sonarqubeStatus).Transitions(transitions).Execute()



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
	findingStatus := "findingStatus_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	sonarqubeIssue := int32(56) // int32 |  (optional)
	sonarqubeStatus := "sonarqubeStatus_example" // string |  (optional)
	transitions := "transitions_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsList(context.Background()).FindingStatus(findingStatus).Id(id).Limit(limit).Offset(offset).SonarqubeIssue(sonarqubeIssue).SonarqubeStatus(sonarqubeStatus).Transitions(transitions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeTransitionsList`: PaginatedSonarqubeIssueTransitionList
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeTransitionsAPI.SonarqubeTransitionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findingStatus** | **string** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **sonarqubeIssue** | **int32** |  | 
 **sonarqubeStatus** | **string** |  | 
 **transitions** | **string** |  | 

### Return type

[**PaginatedSonarqubeIssueTransitionList**](PaginatedSonarqubeIssueTransitionList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeTransitionsPartialUpdate

> SonarqubeIssueTransition SonarqubeTransitionsPartialUpdate(ctx, id).PatchedSonarqubeIssueTransitionRequest(patchedSonarqubeIssueTransitionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue_ transition.
	patchedSonarqubeIssueTransitionRequest := *openapiclient.NewPatchedSonarqubeIssueTransitionRequest() // PatchedSonarqubeIssueTransitionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsPartialUpdate(context.Background(), id).PatchedSonarqubeIssueTransitionRequest(patchedSonarqubeIssueTransitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeTransitionsPartialUpdate`: SonarqubeIssueTransition
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeTransitionsAPI.SonarqubeTransitionsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue_ transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSonarqubeIssueTransitionRequest** | [**PatchedSonarqubeIssueTransitionRequest**](PatchedSonarqubeIssueTransitionRequest.md) |  | 

### Return type

[**SonarqubeIssueTransition**](SonarqubeIssueTransition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeTransitionsRetrieve

> SonarqubeIssueTransition SonarqubeTransitionsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue_ transition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeTransitionsRetrieve`: SonarqubeIssueTransition
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeTransitionsAPI.SonarqubeTransitionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue_ transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SonarqubeIssueTransition**](SonarqubeIssueTransition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SonarqubeTransitionsUpdate

> SonarqubeIssueTransition SonarqubeTransitionsUpdate(ctx, id).SonarqubeIssueTransitionRequest(sonarqubeIssueTransitionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this sonarqube_ issue_ transition.
	sonarqubeIssueTransitionRequest := *openapiclient.NewSonarqubeIssueTransitionRequest("FindingStatus_example", "SonarqubeStatus_example", "Transitions_example", int32(123)) // SonarqubeIssueTransitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SonarqubeTransitionsAPI.SonarqubeTransitionsUpdate(context.Background(), id).SonarqubeIssueTransitionRequest(sonarqubeIssueTransitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SonarqubeTransitionsAPI.SonarqubeTransitionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SonarqubeTransitionsUpdate`: SonarqubeIssueTransition
	fmt.Fprintf(os.Stdout, "Response from `SonarqubeTransitionsAPI.SonarqubeTransitionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this sonarqube_ issue_ transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSonarqubeTransitionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sonarqubeIssueTransitionRequest** | [**SonarqubeIssueTransitionRequest**](SonarqubeIssueTransitionRequest.md) |  | 

### Return type

[**SonarqubeIssueTransition**](SonarqubeIssueTransition.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

