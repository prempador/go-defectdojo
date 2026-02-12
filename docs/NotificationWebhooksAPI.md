# \NotificationWebhooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationWebhooksCreate**](NotificationWebhooksAPI.md#NotificationWebhooksCreate) | **Post** /api/v2/notification_webhooks/ | 
[**NotificationWebhooksDeletePreviewList**](NotificationWebhooksAPI.md#NotificationWebhooksDeletePreviewList) | **Get** /api/v2/notification_webhooks/{id}/delete_preview/ | 
[**NotificationWebhooksDestroy**](NotificationWebhooksAPI.md#NotificationWebhooksDestroy) | **Delete** /api/v2/notification_webhooks/{id}/ | 
[**NotificationWebhooksList**](NotificationWebhooksAPI.md#NotificationWebhooksList) | **Get** /api/v2/notification_webhooks/ | 
[**NotificationWebhooksPartialUpdate**](NotificationWebhooksAPI.md#NotificationWebhooksPartialUpdate) | **Patch** /api/v2/notification_webhooks/{id}/ | 
[**NotificationWebhooksRetrieve**](NotificationWebhooksAPI.md#NotificationWebhooksRetrieve) | **Get** /api/v2/notification_webhooks/{id}/ | 
[**NotificationWebhooksUpdate**](NotificationWebhooksAPI.md#NotificationWebhooksUpdate) | **Put** /api/v2/notification_webhooks/{id}/ | 



## NotificationWebhooksCreate

> NotificationWebhooks NotificationWebhooksCreate(ctx).NotificationWebhooksRequest(notificationWebhooksRequest).Execute()



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
	notificationWebhooksRequest := *openapiclient.NewNotificationWebhooksRequest() // NotificationWebhooksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksCreate(context.Background()).NotificationWebhooksRequest(notificationWebhooksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationWebhooksCreate`: NotificationWebhooks
	fmt.Fprintf(os.Stdout, "Response from `NotificationWebhooksAPI.NotificationWebhooksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationWebhooksRequest** | [**NotificationWebhooksRequest**](NotificationWebhooksRequest.md) |  | 

### Return type

[**NotificationWebhooks**](NotificationWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationWebhooksDeletePreviewList

> PaginatedDeletePreviewList NotificationWebhooksDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this notification_ webhooks.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationWebhooksDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `NotificationWebhooksAPI.NotificationWebhooksDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this notification_ webhooks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksDeletePreviewListRequest struct via the builder pattern


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


## NotificationWebhooksDestroy

> NotificationWebhooksDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this notification_ webhooks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this notification_ webhooks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksDestroyRequest struct via the builder pattern


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


## NotificationWebhooksList

> PaginatedNotificationWebhooksList NotificationWebhooksList(ctx).FirstError(firstError).HeaderName(headerName).HeaderValue(headerValue).LastError(lastError).Limit(limit).Name(name).Note(note).Offset(offset).Owner(owner).Status(status).Url(url).Execute()



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
	firstError := time.Now() // time.Time |  (optional)
	headerName := "headerName_example" // string |  (optional)
	headerValue := "headerValue_example" // string |  (optional)
	lastError := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	note := "note_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	owner := int32(56) // int32 |  (optional)
	status := "status_example" // string | Status of the incoming webhook  * `active` - Active * `active_tmp` - Active but 5xx (or similar) error detected * `inactive_tmp` - Temporary inactive because of 5xx (or similar) error * `inactive_permanent` - Permanently inactive (optional)
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksList(context.Background()).FirstError(firstError).HeaderName(headerName).HeaderValue(headerValue).LastError(lastError).Limit(limit).Name(name).Note(note).Offset(offset).Owner(owner).Status(status).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationWebhooksList`: PaginatedNotificationWebhooksList
	fmt.Fprintf(os.Stdout, "Response from `NotificationWebhooksAPI.NotificationWebhooksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firstError** | **time.Time** |  | 
 **headerName** | **string** |  | 
 **headerValue** | **string** |  | 
 **lastError** | **time.Time** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **note** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **owner** | **int32** |  | 
 **status** | **string** | Status of the incoming webhook  * &#x60;active&#x60; - Active * &#x60;active_tmp&#x60; - Active but 5xx (or similar) error detected * &#x60;inactive_tmp&#x60; - Temporary inactive because of 5xx (or similar) error * &#x60;inactive_permanent&#x60; - Permanently inactive | 
 **url** | **string** |  | 

### Return type

[**PaginatedNotificationWebhooksList**](PaginatedNotificationWebhooksList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationWebhooksPartialUpdate

> NotificationWebhooks NotificationWebhooksPartialUpdate(ctx, id).PatchedNotificationWebhooksRequest(patchedNotificationWebhooksRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this notification_ webhooks.
	patchedNotificationWebhooksRequest := *openapiclient.NewPatchedNotificationWebhooksRequest() // PatchedNotificationWebhooksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksPartialUpdate(context.Background(), id).PatchedNotificationWebhooksRequest(patchedNotificationWebhooksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationWebhooksPartialUpdate`: NotificationWebhooks
	fmt.Fprintf(os.Stdout, "Response from `NotificationWebhooksAPI.NotificationWebhooksPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this notification_ webhooks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationWebhooksRequest** | [**PatchedNotificationWebhooksRequest**](PatchedNotificationWebhooksRequest.md) |  | 

### Return type

[**NotificationWebhooks**](NotificationWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationWebhooksRetrieve

> NotificationWebhooks NotificationWebhooksRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this notification_ webhooks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationWebhooksRetrieve`: NotificationWebhooks
	fmt.Fprintf(os.Stdout, "Response from `NotificationWebhooksAPI.NotificationWebhooksRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this notification_ webhooks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationWebhooks**](NotificationWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationWebhooksUpdate

> NotificationWebhooks NotificationWebhooksUpdate(ctx, id).NotificationWebhooksRequest(notificationWebhooksRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this notification_ webhooks.
	notificationWebhooksRequest := *openapiclient.NewNotificationWebhooksRequest() // NotificationWebhooksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationWebhooksAPI.NotificationWebhooksUpdate(context.Background(), id).NotificationWebhooksRequest(notificationWebhooksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationWebhooksAPI.NotificationWebhooksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationWebhooksUpdate`: NotificationWebhooks
	fmt.Fprintf(os.Stdout, "Response from `NotificationWebhooksAPI.NotificationWebhooksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this notification_ webhooks. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationWebhooksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationWebhooksRequest** | [**NotificationWebhooksRequest**](NotificationWebhooksRequest.md) |  | 

### Return type

[**NotificationWebhooks**](NotificationWebhooks.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

