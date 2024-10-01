# \TestImportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestImportsCreate**](TestImportsAPI.md#TestImportsCreate) | **Post** /api/v2/test_imports/ | 
[**TestImportsDeletePreviewList**](TestImportsAPI.md#TestImportsDeletePreviewList) | **Get** /api/v2/test_imports/{id}/delete_preview/ | 
[**TestImportsDestroy**](TestImportsAPI.md#TestImportsDestroy) | **Delete** /api/v2/test_imports/{id}/ | 
[**TestImportsList**](TestImportsAPI.md#TestImportsList) | **Get** /api/v2/test_imports/ | 
[**TestImportsPartialUpdate**](TestImportsAPI.md#TestImportsPartialUpdate) | **Patch** /api/v2/test_imports/{id}/ | 
[**TestImportsRetrieve**](TestImportsAPI.md#TestImportsRetrieve) | **Get** /api/v2/test_imports/{id}/ | 
[**TestImportsUpdate**](TestImportsAPI.md#TestImportsUpdate) | **Put** /api/v2/test_imports/{id}/ | 



## TestImportsCreate

> TestImport TestImportsCreate(ctx).TestImportRequest(testImportRequest).Execute()



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
	testImportRequest := *openapiclient.NewTestImportRequest() // TestImportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestImportsAPI.TestImportsCreate(context.Background()).TestImportRequest(testImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestImportsCreate`: TestImport
	fmt.Fprintf(os.Stdout, "Response from `TestImportsAPI.TestImportsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testImportRequest** | [**TestImportRequest**](TestImportRequest.md) |  | 

### Return type

[**TestImport**](TestImport.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestImportsDeletePreviewList

> PaginatedDeletePreviewList TestImportsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ import.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestImportsAPI.TestImportsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestImportsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `TestImportsAPI.TestImportsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsDeletePreviewListRequest struct via the builder pattern


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


## TestImportsDestroy

> TestImportsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ import.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestImportsAPI.TestImportsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsDestroyRequest struct via the builder pattern


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


## TestImportsList

> PaginatedTestImportList TestImportsList(ctx).BranchTag(branchTag).BuildId(buildId).CommitHash(commitHash).FindingsAffected(findingsAffected).Limit(limit).Offset(offset).Test(test).TestImportFindingActionAction(testImportFindingActionAction).TestImportFindingActionCreated(testImportFindingActionCreated).TestImportFindingActionFinding(testImportFindingActionFinding).Version(version).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/prempador/go-defectdojo"
)

func main() {
	branchTag := "branchTag_example" // string |  (optional)
	buildId := "buildId_example" // string |  (optional)
	commitHash := "commitHash_example" // string |  (optional)
	findingsAffected := []int32{int32(123)} // []int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	test := int32(56) // int32 |  (optional)
	testImportFindingActionAction := "testImportFindingActionAction_example" // string | * `N` - created * `C` - closed * `R` - reactivated * `U` - left untouched (optional)
	testImportFindingActionCreated := time.Now() // time.Time |  (optional)
	testImportFindingActionFinding := int32(56) // int32 |  (optional)
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestImportsAPI.TestImportsList(context.Background()).BranchTag(branchTag).BuildId(buildId).CommitHash(commitHash).FindingsAffected(findingsAffected).Limit(limit).Offset(offset).Test(test).TestImportFindingActionAction(testImportFindingActionAction).TestImportFindingActionCreated(testImportFindingActionCreated).TestImportFindingActionFinding(testImportFindingActionFinding).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestImportsList`: PaginatedTestImportList
	fmt.Fprintf(os.Stdout, "Response from `TestImportsAPI.TestImportsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **branchTag** | **string** |  | 
 **buildId** | **string** |  | 
 **commitHash** | **string** |  | 
 **findingsAffected** | **[]int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **test** | **int32** |  | 
 **testImportFindingActionAction** | **string** | * &#x60;N&#x60; - created * &#x60;C&#x60; - closed * &#x60;R&#x60; - reactivated * &#x60;U&#x60; - left untouched | 
 **testImportFindingActionCreated** | **time.Time** |  | 
 **testImportFindingActionFinding** | **int32** |  | 
 **version** | **string** |  | 

### Return type

[**PaginatedTestImportList**](PaginatedTestImportList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestImportsPartialUpdate

> TestImport TestImportsPartialUpdate(ctx, id).PatchedTestImportRequest(patchedTestImportRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ import.
	patchedTestImportRequest := *openapiclient.NewPatchedTestImportRequest() // PatchedTestImportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestImportsAPI.TestImportsPartialUpdate(context.Background(), id).PatchedTestImportRequest(patchedTestImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestImportsPartialUpdate`: TestImport
	fmt.Fprintf(os.Stdout, "Response from `TestImportsAPI.TestImportsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTestImportRequest** | [**PatchedTestImportRequest**](PatchedTestImportRequest.md) |  | 

### Return type

[**TestImport**](TestImport.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestImportsRetrieve

> TestImport TestImportsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ import.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestImportsAPI.TestImportsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestImportsRetrieve`: TestImport
	fmt.Fprintf(os.Stdout, "Response from `TestImportsAPI.TestImportsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestImport**](TestImport.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestImportsUpdate

> TestImport TestImportsUpdate(ctx, id).TestImportRequest(testImportRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ import.
	testImportRequest := *openapiclient.NewTestImportRequest() // TestImportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestImportsAPI.TestImportsUpdate(context.Background(), id).TestImportRequest(testImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestImportsAPI.TestImportsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestImportsUpdate`: TestImport
	fmt.Fprintf(os.Stdout, "Response from `TestImportsAPI.TestImportsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ import. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestImportsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testImportRequest** | [**TestImportRequest**](TestImportRequest.md) |  | 

### Return type

[**TestImport**](TestImport.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

