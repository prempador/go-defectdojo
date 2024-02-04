# \NoteTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NoteTypeCreate**](NoteTypeAPI.md#NoteTypeCreate) | **Post** /api/v2/note_type/ | 
[**NoteTypeDeletePreviewList**](NoteTypeAPI.md#NoteTypeDeletePreviewList) | **Get** /api/v2/note_type/{id}/delete_preview/ | 
[**NoteTypeDestroy**](NoteTypeAPI.md#NoteTypeDestroy) | **Delete** /api/v2/note_type/{id}/ | 
[**NoteTypeList**](NoteTypeAPI.md#NoteTypeList) | **Get** /api/v2/note_type/ | 
[**NoteTypePartialUpdate**](NoteTypeAPI.md#NoteTypePartialUpdate) | **Patch** /api/v2/note_type/{id}/ | 
[**NoteTypeRetrieve**](NoteTypeAPI.md#NoteTypeRetrieve) | **Get** /api/v2/note_type/{id}/ | 
[**NoteTypeUpdate**](NoteTypeAPI.md#NoteTypeUpdate) | **Put** /api/v2/note_type/{id}/ | 



## NoteTypeCreate

> NoteType NoteTypeCreate(ctx).NoteTypeRequest(noteTypeRequest).Execute()



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
	noteTypeRequest := *openapiclient.NewNoteTypeRequest("Name_example", "Description_example") // NoteTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteTypeAPI.NoteTypeCreate(context.Background()).NoteTypeRequest(noteTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoteTypeCreate`: NoteType
	fmt.Fprintf(os.Stdout, "Response from `NoteTypeAPI.NoteTypeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noteTypeRequest** | [**NoteTypeRequest**](NoteTypeRequest.md) |  | 

### Return type

[**NoteType**](NoteType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteTypeDeletePreviewList

> PaginatedDeletePreviewList NoteTypeDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this note_ type.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteTypeAPI.NoteTypeDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypeDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoteTypeDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `NoteTypeAPI.NoteTypeDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this note_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypeDeletePreviewListRequest struct via the builder pattern


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


## NoteTypeDestroy

> NoteTypeDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this note_ type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NoteTypeAPI.NoteTypeDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypeDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this note_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypeDestroyRequest struct via the builder pattern


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


## NoteTypeList

> PaginatedNoteTypeList NoteTypeList(ctx).Description(description).Id(id).IsActive(isActive).IsMandatory(isMandatory).IsSingle(isSingle).Limit(limit).Name(name).Offset(offset).Execute()



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
	description := "description_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	isActive := true // bool |  (optional)
	isMandatory := true // bool |  (optional)
	isSingle := true // bool |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteTypeAPI.NoteTypeList(context.Background()).Description(description).Id(id).IsActive(isActive).IsMandatory(isMandatory).IsSingle(isSingle).Limit(limit).Name(name).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoteTypeList`: PaginatedNoteTypeList
	fmt.Fprintf(os.Stdout, "Response from `NoteTypeAPI.NoteTypeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** |  | 
 **id** | **int32** |  | 
 **isActive** | **bool** |  | 
 **isMandatory** | **bool** |  | 
 **isSingle** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedNoteTypeList**](PaginatedNoteTypeList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteTypePartialUpdate

> NoteType NoteTypePartialUpdate(ctx, id).PatchedNoteTypeRequest(patchedNoteTypeRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this note_ type.
	patchedNoteTypeRequest := *openapiclient.NewPatchedNoteTypeRequest() // PatchedNoteTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteTypeAPI.NoteTypePartialUpdate(context.Background(), id).PatchedNoteTypeRequest(patchedNoteTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoteTypePartialUpdate`: NoteType
	fmt.Fprintf(os.Stdout, "Response from `NoteTypeAPI.NoteTypePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this note_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNoteTypeRequest** | [**PatchedNoteTypeRequest**](PatchedNoteTypeRequest.md) |  | 

### Return type

[**NoteType**](NoteType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteTypeRetrieve

> NoteType NoteTypeRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this note_ type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteTypeAPI.NoteTypeRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypeRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoteTypeRetrieve`: NoteType
	fmt.Fprintf(os.Stdout, "Response from `NoteTypeAPI.NoteTypeRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this note_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypeRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NoteType**](NoteType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteTypeUpdate

> NoteType NoteTypeUpdate(ctx, id).NoteTypeRequest(noteTypeRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this note_ type.
	noteTypeRequest := *openapiclient.NewNoteTypeRequest("Name_example", "Description_example") // NoteTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteTypeAPI.NoteTypeUpdate(context.Background(), id).NoteTypeRequest(noteTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteTypeAPI.NoteTypeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoteTypeUpdate`: NoteType
	fmt.Fprintf(os.Stdout, "Response from `NoteTypeAPI.NoteTypeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this note_ type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoteTypeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteTypeRequest** | [**NoteTypeRequest**](NoteTypeRequest.md) |  | 

### Return type

[**NoteType**](NoteType.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

