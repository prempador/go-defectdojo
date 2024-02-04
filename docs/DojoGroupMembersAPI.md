# \DojoGroupMembersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DojoGroupMembersCreate**](DojoGroupMembersAPI.md#DojoGroupMembersCreate) | **Post** /api/v2/dojo_group_members/ | 
[**DojoGroupMembersDeletePreviewList**](DojoGroupMembersAPI.md#DojoGroupMembersDeletePreviewList) | **Get** /api/v2/dojo_group_members/{id}/delete_preview/ | 
[**DojoGroupMembersDestroy**](DojoGroupMembersAPI.md#DojoGroupMembersDestroy) | **Delete** /api/v2/dojo_group_members/{id}/ | 
[**DojoGroupMembersList**](DojoGroupMembersAPI.md#DojoGroupMembersList) | **Get** /api/v2/dojo_group_members/ | 
[**DojoGroupMembersRetrieve**](DojoGroupMembersAPI.md#DojoGroupMembersRetrieve) | **Get** /api/v2/dojo_group_members/{id}/ | 
[**DojoGroupMembersUpdate**](DojoGroupMembersAPI.md#DojoGroupMembersUpdate) | **Put** /api/v2/dojo_group_members/{id}/ | 



## DojoGroupMembersCreate

> DojoGroupMember DojoGroupMembersCreate(ctx).DojoGroupMemberRequest(dojoGroupMemberRequest).Execute()



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
	dojoGroupMemberRequest := *openapiclient.NewDojoGroupMemberRequest(int32(123), int32(123), int32(123)) // DojoGroupMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupMembersAPI.DojoGroupMembersCreate(context.Background()).DojoGroupMemberRequest(dojoGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupMembersAPI.DojoGroupMembersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupMembersCreate`: DojoGroupMember
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupMembersAPI.DojoGroupMembersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupMembersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dojoGroupMemberRequest** | [**DojoGroupMemberRequest**](DojoGroupMemberRequest.md) |  | 

### Return type

[**DojoGroupMember**](DojoGroupMember.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupMembersDeletePreviewList

> PaginatedDeletePreviewList DojoGroupMembersDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group_ member.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupMembersAPI.DojoGroupMembersDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupMembersAPI.DojoGroupMembersDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupMembersDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupMembersAPI.DojoGroupMembersDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupMembersDeletePreviewListRequest struct via the builder pattern


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


## DojoGroupMembersDestroy

> DojoGroupMembersDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group_ member.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DojoGroupMembersAPI.DojoGroupMembersDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupMembersAPI.DojoGroupMembersDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupMembersDestroyRequest struct via the builder pattern


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


## DojoGroupMembersList

> PaginatedDojoGroupMemberList DojoGroupMembersList(ctx).GroupId(groupId).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).UserId(userId).Execute()



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
	groupId := int32(56) // int32 |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	userId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupMembersAPI.DojoGroupMembersList(context.Background()).GroupId(groupId).Id(id).Limit(limit).Offset(offset).Prefetch(prefetch).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupMembersAPI.DojoGroupMembersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupMembersList`: PaginatedDojoGroupMemberList
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupMembersAPI.DojoGroupMembersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupMembersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int32** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **userId** | **int32** |  | 

### Return type

[**PaginatedDojoGroupMemberList**](PaginatedDojoGroupMemberList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupMembersRetrieve

> DojoGroupMember DojoGroupMembersRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group_ member.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupMembersAPI.DojoGroupMembersRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupMembersAPI.DojoGroupMembersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupMembersRetrieve`: DojoGroupMember
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupMembersAPI.DojoGroupMembersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupMembersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**DojoGroupMember**](DojoGroupMember.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DojoGroupMembersUpdate

> DojoGroupMember DojoGroupMembersUpdate(ctx, id).DojoGroupMemberRequest(dojoGroupMemberRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this dojo_ group_ member.
	dojoGroupMemberRequest := *openapiclient.NewDojoGroupMemberRequest(int32(123), int32(123), int32(123)) // DojoGroupMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DojoGroupMembersAPI.DojoGroupMembersUpdate(context.Background(), id).DojoGroupMemberRequest(dojoGroupMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DojoGroupMembersAPI.DojoGroupMembersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DojoGroupMembersUpdate`: DojoGroupMember
	fmt.Fprintf(os.Stdout, "Response from `DojoGroupMembersAPI.DojoGroupMembersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this dojo_ group_ member. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDojoGroupMembersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dojoGroupMemberRequest** | [**DojoGroupMemberRequest**](DojoGroupMemberRequest.md) |  | 

### Return type

[**DojoGroupMember**](DojoGroupMember.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

