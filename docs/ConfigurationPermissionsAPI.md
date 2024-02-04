# \ConfigurationPermissionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationPermissionsList**](ConfigurationPermissionsAPI.md#ConfigurationPermissionsList) | **Get** /api/v2/configuration_permissions/ | 
[**ConfigurationPermissionsRetrieve**](ConfigurationPermissionsAPI.md#ConfigurationPermissionsRetrieve) | **Get** /api/v2/configuration_permissions/{id}/ | 



## ConfigurationPermissionsList

> PaginatedConfigurationPermissionList ConfigurationPermissionsList(ctx).Codename(codename).Id(id).Limit(limit).Name(name).Offset(offset).Execute()



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
	codename := "codename_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPermissionsAPI.ConfigurationPermissionsList(context.Background()).Codename(codename).Id(id).Limit(limit).Name(name).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPermissionsAPI.ConfigurationPermissionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurationPermissionsList`: PaginatedConfigurationPermissionList
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPermissionsAPI.ConfigurationPermissionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationPermissionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codename** | **string** |  | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedConfigurationPermissionList**](PaginatedConfigurationPermissionList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationPermissionsRetrieve

> ConfigurationPermission ConfigurationPermissionsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPermissionsAPI.ConfigurationPermissionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPermissionsAPI.ConfigurationPermissionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurationPermissionsRetrieve`: ConfigurationPermission
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPermissionsAPI.ConfigurationPermissionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationPermissionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationPermission**](ConfigurationPermission.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

