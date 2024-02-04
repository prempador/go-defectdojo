# \UserContactInfosAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserContactInfosCreate**](UserContactInfosAPI.md#UserContactInfosCreate) | **Post** /api/v2/user_contact_infos/ | 
[**UserContactInfosDeletePreviewList**](UserContactInfosAPI.md#UserContactInfosDeletePreviewList) | **Get** /api/v2/user_contact_infos/{id}/delete_preview/ | 
[**UserContactInfosDestroy**](UserContactInfosAPI.md#UserContactInfosDestroy) | **Delete** /api/v2/user_contact_infos/{id}/ | 
[**UserContactInfosList**](UserContactInfosAPI.md#UserContactInfosList) | **Get** /api/v2/user_contact_infos/ | 
[**UserContactInfosPartialUpdate**](UserContactInfosAPI.md#UserContactInfosPartialUpdate) | **Patch** /api/v2/user_contact_infos/{id}/ | 
[**UserContactInfosRetrieve**](UserContactInfosAPI.md#UserContactInfosRetrieve) | **Get** /api/v2/user_contact_infos/{id}/ | 
[**UserContactInfosUpdate**](UserContactInfosAPI.md#UserContactInfosUpdate) | **Put** /api/v2/user_contact_infos/{id}/ | 



## UserContactInfosCreate

> UserContactInfo UserContactInfosCreate(ctx).UserContactInfoRequest(userContactInfoRequest).Execute()



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
	userContactInfoRequest := *openapiclient.NewUserContactInfoRequest(int32(123)) // UserContactInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserContactInfosAPI.UserContactInfosCreate(context.Background()).UserContactInfoRequest(userContactInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserContactInfosCreate`: UserContactInfo
	fmt.Fprintf(os.Stdout, "Response from `UserContactInfosAPI.UserContactInfosCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userContactInfoRequest** | [**UserContactInfoRequest**](UserContactInfoRequest.md) |  | 

### Return type

[**UserContactInfo**](UserContactInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserContactInfosDeletePreviewList

> PaginatedDeletePreviewList UserContactInfosDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this user contact info.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserContactInfosAPI.UserContactInfosDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserContactInfosDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `UserContactInfosAPI.UserContactInfosDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user contact info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosDeletePreviewListRequest struct via the builder pattern


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


## UserContactInfosDestroy

> UserContactInfosDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this user contact info.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserContactInfosAPI.UserContactInfosDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user contact info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosDestroyRequest struct via the builder pattern


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


## UserContactInfosList

> PaginatedUserContactInfoList UserContactInfosList(ctx).BlockExecution(blockExecution).CellNumber(cellNumber).ForcePasswordReset(forcePasswordReset).GithubUsername(githubUsername).Limit(limit).Offset(offset).PhoneNumber(phoneNumber).Prefetch(prefetch).SlackUserId(slackUserId).SlackUsername(slackUsername).Title(title).TwitterUsername(twitterUsername).User(user).Execute()



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
	blockExecution := true // bool |  (optional)
	cellNumber := "cellNumber_example" // string |  (optional)
	forcePasswordReset := true // bool |  (optional)
	githubUsername := "githubUsername_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	phoneNumber := "phoneNumber_example" // string |  (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	slackUserId := "slackUserId_example" // string |  (optional)
	slackUsername := "slackUsername_example" // string |  (optional)
	title := "title_example" // string |  (optional)
	twitterUsername := "twitterUsername_example" // string |  (optional)
	user := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserContactInfosAPI.UserContactInfosList(context.Background()).BlockExecution(blockExecution).CellNumber(cellNumber).ForcePasswordReset(forcePasswordReset).GithubUsername(githubUsername).Limit(limit).Offset(offset).PhoneNumber(phoneNumber).Prefetch(prefetch).SlackUserId(slackUserId).SlackUsername(slackUsername).Title(title).TwitterUsername(twitterUsername).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserContactInfosList`: PaginatedUserContactInfoList
	fmt.Fprintf(os.Stdout, "Response from `UserContactInfosAPI.UserContactInfosList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockExecution** | **bool** |  | 
 **cellNumber** | **string** |  | 
 **forcePasswordReset** | **bool** |  | 
 **githubUsername** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **phoneNumber** | **string** |  | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **slackUserId** | **string** |  | 
 **slackUsername** | **string** |  | 
 **title** | **string** |  | 
 **twitterUsername** | **string** |  | 
 **user** | **int32** |  | 

### Return type

[**PaginatedUserContactInfoList**](PaginatedUserContactInfoList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserContactInfosPartialUpdate

> UserContactInfo UserContactInfosPartialUpdate(ctx, id).PatchedUserContactInfoRequest(patchedUserContactInfoRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this user contact info.
	patchedUserContactInfoRequest := *openapiclient.NewPatchedUserContactInfoRequest() // PatchedUserContactInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserContactInfosAPI.UserContactInfosPartialUpdate(context.Background(), id).PatchedUserContactInfoRequest(patchedUserContactInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserContactInfosPartialUpdate`: UserContactInfo
	fmt.Fprintf(os.Stdout, "Response from `UserContactInfosAPI.UserContactInfosPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user contact info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUserContactInfoRequest** | [**PatchedUserContactInfoRequest**](PatchedUserContactInfoRequest.md) |  | 

### Return type

[**UserContactInfo**](UserContactInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserContactInfosRetrieve

> UserContactInfo UserContactInfosRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this user contact info.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserContactInfosAPI.UserContactInfosRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserContactInfosRetrieve`: UserContactInfo
	fmt.Fprintf(os.Stdout, "Response from `UserContactInfosAPI.UserContactInfosRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user contact info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**UserContactInfo**](UserContactInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserContactInfosUpdate

> UserContactInfo UserContactInfosUpdate(ctx, id).UserContactInfoRequest(userContactInfoRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this user contact info.
	userContactInfoRequest := *openapiclient.NewUserContactInfoRequest(int32(123)) // UserContactInfoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserContactInfosAPI.UserContactInfosUpdate(context.Background(), id).UserContactInfoRequest(userContactInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserContactInfosAPI.UserContactInfosUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserContactInfosUpdate`: UserContactInfo
	fmt.Fprintf(os.Stdout, "Response from `UserContactInfosAPI.UserContactInfosUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this user contact info. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserContactInfosUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userContactInfoRequest** | [**UserContactInfoRequest**](UserContactInfoRequest.md) |  | 

### Return type

[**UserContactInfo**](UserContactInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

