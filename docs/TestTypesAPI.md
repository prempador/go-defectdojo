# \TestTypesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestTypesCreate**](TestTypesAPI.md#TestTypesCreate) | **Post** /api/v2/test_types/ | 
[**TestTypesList**](TestTypesAPI.md#TestTypesList) | **Get** /api/v2/test_types/ | 
[**TestTypesPartialUpdate**](TestTypesAPI.md#TestTypesPartialUpdate) | **Patch** /api/v2/test_types/{id}/ | 
[**TestTypesRetrieve**](TestTypesAPI.md#TestTypesRetrieve) | **Get** /api/v2/test_types/{id}/ | 
[**TestTypesUpdate**](TestTypesAPI.md#TestTypesUpdate) | **Put** /api/v2/test_types/{id}/ | 



## TestTypesCreate

> TestType TestTypesCreate(ctx).TestTypeRequest(testTypeRequest).Execute()



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
	testTypeRequest := *openapiclient.NewTestTypeRequest("Name_example") // TestTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestTypesAPI.TestTypesCreate(context.Background()).TestTypeRequest(testTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestTypesAPI.TestTypesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTypesCreate`: TestType
	fmt.Fprintf(os.Stdout, "Response from `TestTypesAPI.TestTypesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestTypesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testTypeRequest** | [**TestTypeRequest**](TestTypeRequest.md) |  | 

### Return type

[**TestType**](TestType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestTypesList

> PaginatedTestTypeList TestTypesList(ctx).Limit(limit).Name(name).Offset(offset).Execute()



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
	name := "name_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestTypesAPI.TestTypesList(context.Background()).Limit(limit).Name(name).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestTypesAPI.TestTypesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTypesList`: PaginatedTestTypeList
	fmt.Fprintf(os.Stdout, "Response from `TestTypesAPI.TestTypesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestTypesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedTestTypeList**](PaginatedTestTypeList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestTypesPartialUpdate

> TestType TestTypesPartialUpdate(ctx, id).PatchedTestTypeRequest(patchedTestTypeRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ type.
	patchedTestTypeRequest := *openapiclient.NewPatchedTestTypeRequest() // PatchedTestTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestTypesAPI.TestTypesPartialUpdate(context.Background(), id).PatchedTestTypeRequest(patchedTestTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestTypesAPI.TestTypesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTypesPartialUpdate`: TestType
	fmt.Fprintf(os.Stdout, "Response from `TestTypesAPI.TestTypesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestTypesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTestTypeRequest** | [**PatchedTestTypeRequest**](PatchedTestTypeRequest.md) |  | 

### Return type

[**TestType**](TestType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestTypesRetrieve

> TestType TestTypesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestTypesAPI.TestTypesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestTypesAPI.TestTypesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTypesRetrieve`: TestType
	fmt.Fprintf(os.Stdout, "Response from `TestTypesAPI.TestTypesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestTypesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestType**](TestType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestTypesUpdate

> TestType TestTypesUpdate(ctx, id).TestTypeRequest(testTypeRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test_ type.
	testTypeRequest := *openapiclient.NewTestTypeRequest("Name_example") // TestTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestTypesAPI.TestTypesUpdate(context.Background(), id).TestTypeRequest(testTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestTypesAPI.TestTypesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestTypesUpdate`: TestType
	fmt.Fprintf(os.Stdout, "Response from `TestTypesAPI.TestTypesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestTypesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testTypeRequest** | [**TestTypeRequest**](TestTypeRequest.md) |  | 

### Return type

[**TestType**](TestType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

