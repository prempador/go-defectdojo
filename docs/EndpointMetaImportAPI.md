# \EndpointMetaImportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointMetaImportCreate**](EndpointMetaImportAPI.md#EndpointMetaImportCreate) | **Post** /api/v2/endpoint_meta_import/ | 



## EndpointMetaImportCreate

> EndpointMetaImporter EndpointMetaImportCreate(ctx).File(file).CreateEndpoints(createEndpoints).CreateTags(createTags).CreateDojoMeta(createDojoMeta).ProductName(productName).Product(product).Execute()





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
	file := os.NewFile(1234, "some_file") // *os.File | 
	createEndpoints := true // bool |  (optional) (default to true)
	createTags := true // bool |  (optional) (default to true)
	createDojoMeta := true // bool |  (optional) (default to false)
	productName := "productName_example" // string |  (optional)
	product := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointMetaImportAPI.EndpointMetaImportCreate(context.Background()).File(file).CreateEndpoints(createEndpoints).CreateTags(createTags).CreateDojoMeta(createDojoMeta).ProductName(productName).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointMetaImportAPI.EndpointMetaImportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndpointMetaImportCreate`: EndpointMetaImporter
	fmt.Fprintf(os.Stdout, "Response from `EndpointMetaImportAPI.EndpointMetaImportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndpointMetaImportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **createEndpoints** | **bool** |  | [default to true]
 **createTags** | **bool** |  | [default to true]
 **createDojoMeta** | **bool** |  | [default to false]
 **productName** | **string** |  | 
 **product** | **int32** |  | 

### Return type

[**EndpointMetaImporter**](EndpointMetaImporter.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

