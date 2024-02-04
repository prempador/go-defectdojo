# \DojoGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DojoGroupsCreate**](DojoGroupsAPI.md#DojoGroupsCreate) | **Post** /api/v2/dojo_groups/ | 
[**DojoGroupsDeletePreviewList**](DojoGroupsAPI.md#DojoGroupsDeletePreviewList) | **Get** /api/v2/dojo_groups/{id}/delete_preview/ | 
[**DojoGroupsDestroy**](DojoGroupsAPI.md#DojoGroupsDestroy) | **Delete** /api/v2/dojo_groups/{id}/ | 
[**DojoGroupsList**](DojoGroupsAPI.md#DojoGroupsList) | **Get** /api/v2/dojo_groups/ | 
[**DojoGroupsPartialUpdate**](DojoGroupsAPI.md#DojoGroupsPartialUpdate) | **Patch** /api/v2/dojo_groups/{id}/ | 
[**DojoGroupsRetrieve**](DojoGroupsAPI.md#DojoGroupsRetrieve) | **Get** /api/v2/dojo_groups/{id}/ | 
[**DojoGroupsUpdate**](DojoGroupsAPI.md#DojoGroupsUpdate) | **Put** /api/v2/dojo_groups/{id}/ | 



## DojoGroupsCreate

> DojoGroup DojoGroupsCreate(ctx).DojoGroupRequest(dojoGroupRequest).Execute()



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
	dojoGroupRequest := *openapiclient.NewDojoGroupRequest("Name_example") // DojoGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupsAPI.DojoGroupsCreate(context.Background()).DojoGroupRequest(dojoGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupsCreate`: DojoGroup
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupsAPI.DojoGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dojoGroupRequest** | [**DojoGroupRequest**](DojoGroupRequest.md) |  | 

### Return type

[**DojoGroup**](DojoGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupsDeletePreviewList

> PaginatedDeletePreviewList DojoGroupsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupsAPI.DojoGroupsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupsAPI.DojoGroupsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsDeletePreviewListRequest struct via the builder pattern


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


## DojoGroupsDestroy

> DojoGroupsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DojoGroupsAPI.DojoGroupsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsDestroyRequest struct via the builder pattern


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


## DojoGroupsList

> PaginatedDojoGroupList DojoGroupsList(ctx).Id(id).Limit(limit).Name(name).Offset(offset).Prefetch(prefetch).SocialProvider(socialProvider).Execute()



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
	name := "name_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	socialProvider := "socialProvider_example" // string | Group imported from a social provider.  * `AzureAD` - AzureAD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupsAPI.DojoGroupsList(context.Background()).Id(id).Limit(limit).Name(name).Offset(offset).Prefetch(prefetch).SocialProvider(socialProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupsList`: PaginatedDojoGroupList
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupsAPI.DojoGroupsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **socialProvider** | **string** | Group imported from a social provider.  * &#x60;AzureAD&#x60; - AzureAD | 

### Return type

[**PaginatedDojoGroupList**](PaginatedDojoGroupList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupsPartialUpdate

> DojoGroup DojoGroupsPartialUpdate(ctx, id).PatchedDojoGroupRequest(patchedDojoGroupRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group.
	patchedDojoGroupRequest := *openapiclient.NewPatchedDojoGroupRequest() // PatchedDojoGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupsAPI.DojoGroupsPartialUpdate(context.Background(), id).PatchedDojoGroupRequest(patchedDojoGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupsPartialUpdate`: DojoGroup
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupsAPI.DojoGroupsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedDojoGroupRequest** | [**PatchedDojoGroupRequest**](PatchedDojoGroupRequest.md) |  | 

### Return type

[**DojoGroup**](DojoGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupsRetrieve

> DojoGroup DojoGroupsRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupsAPI.DojoGroupsRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupsRetrieve`: DojoGroup
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupsAPI.DojoGroupsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**DojoGroup**](DojoGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupsUpdate

> DojoGroup DojoGroupsUpdate(ctx, id).DojoGroupRequest(dojoGroupRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group.
	dojoGroupRequest := *openapiclient.NewDojoGroupRequest("Name_example") // DojoGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupsAPI.DojoGroupsUpdate(context.Background(), id).DojoGroupRequest(dojoGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupsAPI.DojoGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupsUpdate`: DojoGroup
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupsAPI.DojoGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dojoGroupRequest** | [**DojoGroupRequest**](DojoGroupRequest.md) |  | 

### Return type

[**DojoGroup**](DojoGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

