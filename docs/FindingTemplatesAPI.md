# \FindingTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindingTemplatesCreate**](FindingTemplatesAPI.md#FindingTemplatesCreate) | **Post** /api/v2/finding_templates/ | 
[**FindingTemplatesDeletePreviewList**](FindingTemplatesAPI.md#FindingTemplatesDeletePreviewList) | **Get** /api/v2/finding_templates/{id}/delete_preview/ | 
[**FindingTemplatesDestroy**](FindingTemplatesAPI.md#FindingTemplatesDestroy) | **Delete** /api/v2/finding_templates/{id}/ | 
[**FindingTemplatesList**](FindingTemplatesAPI.md#FindingTemplatesList) | **Get** /api/v2/finding_templates/ | 
[**FindingTemplatesPartialUpdate**](FindingTemplatesAPI.md#FindingTemplatesPartialUpdate) | **Patch** /api/v2/finding_templates/{id}/ | 
[**FindingTemplatesRetrieve**](FindingTemplatesAPI.md#FindingTemplatesRetrieve) | **Get** /api/v2/finding_templates/{id}/ | 
[**FindingTemplatesUpdate**](FindingTemplatesAPI.md#FindingTemplatesUpdate) | **Put** /api/v2/finding_templates/{id}/ | 



## FindingTemplatesCreate

> FindingTemplate FindingTemplatesCreate(ctx).FindingTemplateRequest(findingTemplateRequest).Execute()



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
	findingTemplateRequest := *openapiclient.NewFindingTemplateRequest("Title_example") // FindingTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingTemplatesAPI.FindingTemplatesCreate(context.Background()).FindingTemplateRequest(findingTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingTemplatesCreate`: FindingTemplate
	fmt.Fprintf(os.Stdout, "Response from `FindingTemplatesAPI.FindingTemplatesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **findingTemplateRequest** | [**FindingTemplateRequest**](FindingTemplateRequest.md) |  | 

### Return type

[**FindingTemplate**](FindingTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingTemplatesDeletePreviewList

> PaginatedDeletePreviewList FindingTemplatesDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this finding_ template.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingTemplatesAPI.FindingTemplatesDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingTemplatesDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `FindingTemplatesAPI.FindingTemplatesDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding_ template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesDeletePreviewListRequest struct via the builder pattern


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


## FindingTemplatesDestroy

> FindingTemplatesDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this finding_ template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FindingTemplatesAPI.FindingTemplatesDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding_ template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesDestroyRequest struct via the builder pattern


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


## FindingTemplatesList

> PaginatedFindingTemplateList FindingTemplatesList(ctx).Cwe(cwe).Description(description).Id(id).Limit(limit).Mitigation(mitigation).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).Severity(severity).Tag(tag).Tags(tags).Title(title).Execute()



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
	cwe := int32(56) // int32 |  (optional)
	description := "description_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	mitigation := "mitigation_example" // string |  (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	o := []string{"O_example"} // []string | Ordering  * `title` - Title * `-title` - Title (descending) * `cwe` - Cwe * `-cwe` - Cwe (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	severity := "severity_example" // string |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	title := "title_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingTemplatesAPI.FindingTemplatesList(context.Background()).Cwe(cwe).Description(description).Id(id).Limit(limit).Mitigation(mitigation).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).Severity(severity).Tag(tag).Tags(tags).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingTemplatesList`: PaginatedFindingTemplateList
	fmt.Fprintf(os.Stdout, "Response from `FindingTemplatesAPI.FindingTemplatesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cwe** | **int32** |  | 
 **description** | **string** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **mitigation** | **string** |  | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **o** | **[]string** | Ordering  * &#x60;title&#x60; - Title * &#x60;-title&#x60; - Title (descending) * &#x60;cwe&#x60; - Cwe * &#x60;-cwe&#x60; - Cwe (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **severity** | **string** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **title** | **string** |  | 

### Return type

[**PaginatedFindingTemplateList**](PaginatedFindingTemplateList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingTemplatesPartialUpdate

> FindingTemplate FindingTemplatesPartialUpdate(ctx, id).PatchedFindingTemplateRequest(patchedFindingTemplateRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this finding_ template.
	patchedFindingTemplateRequest := *openapiclient.NewPatchedFindingTemplateRequest() // PatchedFindingTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingTemplatesAPI.FindingTemplatesPartialUpdate(context.Background(), id).PatchedFindingTemplateRequest(patchedFindingTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingTemplatesPartialUpdate`: FindingTemplate
	fmt.Fprintf(os.Stdout, "Response from `FindingTemplatesAPI.FindingTemplatesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding_ template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedFindingTemplateRequest** | [**PatchedFindingTemplateRequest**](PatchedFindingTemplateRequest.md) |  | 

### Return type

[**FindingTemplate**](FindingTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingTemplatesRetrieve

> FindingTemplate FindingTemplatesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this finding_ template.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingTemplatesAPI.FindingTemplatesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingTemplatesRetrieve`: FindingTemplate
	fmt.Fprintf(os.Stdout, "Response from `FindingTemplatesAPI.FindingTemplatesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding_ template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FindingTemplate**](FindingTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindingTemplatesUpdate

> FindingTemplate FindingTemplatesUpdate(ctx, id).FindingTemplateRequest(findingTemplateRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this finding_ template.
	findingTemplateRequest := *openapiclient.NewFindingTemplateRequest("Title_example") // FindingTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FindingTemplatesAPI.FindingTemplatesUpdate(context.Background(), id).FindingTemplateRequest(findingTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FindingTemplatesAPI.FindingTemplatesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindingTemplatesUpdate`: FindingTemplate
	fmt.Fprintf(os.Stdout, "Response from `FindingTemplatesAPI.FindingTemplatesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this finding_ template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindingTemplatesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **findingTemplateRequest** | [**FindingTemplateRequest**](FindingTemplateRequest.md) |  | 

### Return type

[**FindingTemplate**](FindingTemplate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

