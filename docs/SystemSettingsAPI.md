# \SystemSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemSettingsList**](SystemSettingsAPI.md#SystemSettingsList) | **Get** /api/v2/system_settings/ | 
[**SystemSettingsPartialUpdate**](SystemSettingsAPI.md#SystemSettingsPartialUpdate) | **Patch** /api/v2/system_settings/{id}/ | 
[**SystemSettingsUpdate**](SystemSettingsAPI.md#SystemSettingsUpdate) | **Put** /api/v2/system_settings/{id}/ | 



## SystemSettingsList

> PaginatedSystemSettingsList SystemSettingsList(ctx).Limit(limit).Offset(offset).Execute()





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
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemSettingsAPI.SystemSettingsList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemSettingsAPI.SystemSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemSettingsList`: PaginatedSystemSettingsList
	fmt.Fprintf(os.Stdout, "Response from `SystemSettingsAPI.SystemSettingsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSystemSettingsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedSystemSettingsList**](PaginatedSystemSettingsList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemSettingsPartialUpdate

> SystemSettings SystemSettingsPartialUpdate(ctx, id).PatchedSystemSettingsRequest(patchedSystemSettingsRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this system_ settings.
	patchedSystemSettingsRequest := *openapiclient.NewPatchedSystemSettingsRequest() // PatchedSystemSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemSettingsAPI.SystemSettingsPartialUpdate(context.Background(), id).PatchedSystemSettingsRequest(patchedSystemSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemSettingsAPI.SystemSettingsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemSettingsPartialUpdate`: SystemSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemSettingsAPI.SystemSettingsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this system_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemSettingsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedSystemSettingsRequest** | [**PatchedSystemSettingsRequest**](PatchedSystemSettingsRequest.md) |  | 

### Return type

[**SystemSettings**](SystemSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemSettingsUpdate

> SystemSettings SystemSettingsUpdate(ctx, id).SystemSettingsRequest(systemSettingsRequest).Execute()





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
	id := int32(56) // int32 | A unique integer value identifying this system_ settings.
	systemSettingsRequest := *openapiclient.NewSystemSettingsRequest() // SystemSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemSettingsAPI.SystemSettingsUpdate(context.Background(), id).SystemSettingsRequest(systemSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemSettingsAPI.SystemSettingsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SystemSettingsUpdate`: SystemSettings
	fmt.Fprintf(os.Stdout, "Response from `SystemSettingsAPI.SystemSettingsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this system_ settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemSettingsRequest** | [**SystemSettingsRequest**](SystemSettingsRequest.md) |  | 

### Return type

[**SystemSettings**](SystemSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

