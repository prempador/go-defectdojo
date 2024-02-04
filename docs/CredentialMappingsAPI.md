# \CredentialMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CredentialMappingsCreate**](CredentialMappingsAPI.md#CredentialMappingsCreate) | **Post** /api/v2/credential_mappings/ | 
[**CredentialMappingsDeletePreviewList**](CredentialMappingsAPI.md#CredentialMappingsDeletePreviewList) | **Get** /api/v2/credential_mappings/{id}/delete_preview/ | 
[**CredentialMappingsDestroy**](CredentialMappingsAPI.md#CredentialMappingsDestroy) | **Delete** /api/v2/credential_mappings/{id}/ | 
[**CredentialMappingsList**](CredentialMappingsAPI.md#CredentialMappingsList) | **Get** /api/v2/credential_mappings/ | 
[**CredentialMappingsPartialUpdate**](CredentialMappingsAPI.md#CredentialMappingsPartialUpdate) | **Patch** /api/v2/credential_mappings/{id}/ | 
[**CredentialMappingsRetrieve**](CredentialMappingsAPI.md#CredentialMappingsRetrieve) | **Get** /api/v2/credential_mappings/{id}/ | 
[**CredentialMappingsUpdate**](CredentialMappingsAPI.md#CredentialMappingsUpdate) | **Put** /api/v2/credential_mappings/{id}/ | 



## CredentialMappingsCreate

> CredentialMapping CredentialMappingsCreate(ctx).CredentialMappingRequest(credentialMappingRequest).Execute()



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
	credentialMappingRequest := *openapiclient.NewCredentialMappingRequest(int32(123)) // CredentialMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialMappingsAPI.CredentialMappingsCreate(context.Background()).CredentialMappingRequest(credentialMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialMappingsCreate`: CredentialMapping
	fmt.Fprintf(os.Stdout, "Response from `CredentialMappingsAPI.CredentialMappingsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentialMappingRequest** | [**CredentialMappingRequest**](CredentialMappingRequest.md) |  | 

### Return type

[**CredentialMapping**](CredentialMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialMappingsDeletePreviewList

> PaginatedDeletePreviewList CredentialMappingsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this cred_ mapping.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialMappingsAPI.CredentialMappingsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialMappingsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `CredentialMappingsAPI.CredentialMappingsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this cred_ mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsDeletePreviewListRequest struct via the builder pattern


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


## CredentialMappingsDestroy

> CredentialMappingsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this cred_ mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CredentialMappingsAPI.CredentialMappingsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this cred_ mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsDestroyRequest struct via the builder pattern


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


## CredentialMappingsList

> PaginatedCredentialMappingList CredentialMappingsList(ctx).CredId(credId).Engagement(engagement).Finding(finding).IsAuthnProvider(isAuthnProvider).Limit(limit).Offset(offset).Product(product).Test(test).Url(url).Execute()



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
	credId := int32(56) // int32 |  (optional)
	engagement := int32(56) // int32 |  (optional)
	finding := int32(56) // int32 |  (optional)
	isAuthnProvider := true // bool |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	product := int32(56) // int32 |  (optional)
	test := int32(56) // int32 |  (optional)
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialMappingsAPI.CredentialMappingsList(context.Background()).CredId(credId).Engagement(engagement).Finding(finding).IsAuthnProvider(isAuthnProvider).Limit(limit).Offset(offset).Product(product).Test(test).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialMappingsList`: PaginatedCredentialMappingList
	fmt.Fprintf(os.Stdout, "Response from `CredentialMappingsAPI.CredentialMappingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credId** | **int32** |  | 
 **engagement** | **int32** |  | 
 **finding** | **int32** |  | 
 **isAuthnProvider** | **bool** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **product** | **int32** |  | 
 **test** | **int32** |  | 
 **url** | **string** |  | 

### Return type

[**PaginatedCredentialMappingList**](PaginatedCredentialMappingList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialMappingsPartialUpdate

> CredentialMapping CredentialMappingsPartialUpdate(ctx, id).PatchedCredentialMappingRequest(patchedCredentialMappingRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this cred_ mapping.
	patchedCredentialMappingRequest := *openapiclient.NewPatchedCredentialMappingRequest() // PatchedCredentialMappingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialMappingsAPI.CredentialMappingsPartialUpdate(context.Background(), id).PatchedCredentialMappingRequest(patchedCredentialMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialMappingsPartialUpdate`: CredentialMapping
	fmt.Fprintf(os.Stdout, "Response from `CredentialMappingsAPI.CredentialMappingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this cred_ mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedCredentialMappingRequest** | [**PatchedCredentialMappingRequest**](PatchedCredentialMappingRequest.md) |  | 

### Return type

[**CredentialMapping**](CredentialMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialMappingsRetrieve

> CredentialMapping CredentialMappingsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this cred_ mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialMappingsAPI.CredentialMappingsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialMappingsRetrieve`: CredentialMapping
	fmt.Fprintf(os.Stdout, "Response from `CredentialMappingsAPI.CredentialMappingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this cred_ mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialMapping**](CredentialMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialMappingsUpdate

> CredentialMapping CredentialMappingsUpdate(ctx, id).CredentialMappingRequest(credentialMappingRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this cred_ mapping.
	credentialMappingRequest := *openapiclient.NewCredentialMappingRequest(int32(123)) // CredentialMappingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CredentialMappingsAPI.CredentialMappingsUpdate(context.Background(), id).CredentialMappingRequest(credentialMappingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialMappingsAPI.CredentialMappingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialMappingsUpdate`: CredentialMapping
	fmt.Fprintf(os.Stdout, "Response from `CredentialMappingsAPI.CredentialMappingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this cred_ mapping. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCredentialMappingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialMappingRequest** | [**CredentialMappingRequest**](CredentialMappingRequest.md) |  | 

### Return type

[**CredentialMapping**](CredentialMapping.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

