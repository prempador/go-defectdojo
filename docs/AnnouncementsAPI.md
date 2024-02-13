# \AnnouncementsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnouncementsCreate**](AnnouncementsAPI.md#AnnouncementsCreate) | **Post** /api/v2/announcements/ | 
[**AnnouncementsDeletePreviewList**](AnnouncementsAPI.md#AnnouncementsDeletePreviewList) | **Get** /api/v2/announcements/{id}/delete_preview/ | 
[**AnnouncementsDestroy**](AnnouncementsAPI.md#AnnouncementsDestroy) | **Delete** /api/v2/announcements/{id}/ | 
[**AnnouncementsList**](AnnouncementsAPI.md#AnnouncementsList) | **Get** /api/v2/announcements/ | 
[**AnnouncementsPartialUpdate**](AnnouncementsAPI.md#AnnouncementsPartialUpdate) | **Patch** /api/v2/announcements/{id}/ | 
[**AnnouncementsRetrieve**](AnnouncementsAPI.md#AnnouncementsRetrieve) | **Get** /api/v2/announcements/{id}/ | 
[**AnnouncementsUpdate**](AnnouncementsAPI.md#AnnouncementsUpdate) | **Put** /api/v2/announcements/{id}/ | 



## AnnouncementsCreate

> Announcement AnnouncementsCreate(ctx).AnnouncementRequest(announcementRequest).Execute()



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
	announcementRequest := *openapiclient.NewAnnouncementRequest() // AnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.AnnouncementsCreate(context.Background()).AnnouncementRequest(announcementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnouncementsCreate`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.AnnouncementsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **announcementRequest** | [**AnnouncementRequest**](AnnouncementRequest.md) |  | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnouncementsDeletePreviewList

> PaginatedDeletePreviewList AnnouncementsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this announcement.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.AnnouncementsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnouncementsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.AnnouncementsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this announcement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsDeletePreviewListRequest struct via the builder pattern


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


## AnnouncementsDestroy

> AnnouncementsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this announcement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnouncementsAPI.AnnouncementsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this announcement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsDestroyRequest struct via the builder pattern


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


## AnnouncementsList

> PaginatedAnnouncementList AnnouncementsList(ctx).Dismissable(dismissable).Limit(limit).Message(message).Offset(offset).Style(style).Execute()



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
	dismissable := true // bool |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	message := "message_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	style := "style_example" // string | The style of banner to display. (info, success, warning, danger)  * `info` - Info * `success` - Success * `warning` - Warning * `danger` - Danger (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.AnnouncementsList(context.Background()).Dismissable(dismissable).Limit(limit).Message(message).Offset(offset).Style(style).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnouncementsList`: PaginatedAnnouncementList
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.AnnouncementsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dismissable** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **message** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **style** | **string** | The style of banner to display. (info, success, warning, danger)  * &#x60;info&#x60; - Info * &#x60;success&#x60; - Success * &#x60;warning&#x60; - Warning * &#x60;danger&#x60; - Danger | 

### Return type

[**PaginatedAnnouncementList**](PaginatedAnnouncementList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnouncementsPartialUpdate

> Announcement AnnouncementsPartialUpdate(ctx, id).PatchedAnnouncementRequest(patchedAnnouncementRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this announcement.
	patchedAnnouncementRequest := *openapiclient.NewPatchedAnnouncementRequest() // PatchedAnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.AnnouncementsPartialUpdate(context.Background(), id).PatchedAnnouncementRequest(patchedAnnouncementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnouncementsPartialUpdate`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.AnnouncementsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this announcement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedAnnouncementRequest** | [**PatchedAnnouncementRequest**](PatchedAnnouncementRequest.md) |  | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnouncementsRetrieve

> Announcement AnnouncementsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this announcement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.AnnouncementsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnouncementsRetrieve`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.AnnouncementsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this announcement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Announcement**](Announcement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnouncementsUpdate

> Announcement AnnouncementsUpdate(ctx, id).AnnouncementRequest(announcementRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this announcement.
	announcementRequest := *openapiclient.NewAnnouncementRequest() // AnnouncementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementsAPI.AnnouncementsUpdate(context.Background(), id).AnnouncementRequest(announcementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementsAPI.AnnouncementsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnnouncementsUpdate`: Announcement
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementsAPI.AnnouncementsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this announcement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnnouncementsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **announcementRequest** | [**AnnouncementRequest**](AnnouncementRequest.md) |  | 

### Return type

[**Announcement**](Announcement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

