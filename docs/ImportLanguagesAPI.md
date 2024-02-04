# \ImportLanguagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportLanguagesCreate**](ImportLanguagesAPI.md#ImportLanguagesCreate) | **Post** /api/v2/import-languages/ | 



## ImportLanguagesCreate

> ImportLanguages ImportLanguagesCreate(ctx).Product(product).File(file).Execute()



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
	product := int32(56) // int32 | 
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImportLanguagesAPI.ImportLanguagesCreate(context.Background()).Product(product).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportLanguagesAPI.ImportLanguagesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportLanguagesCreate`: ImportLanguages
	fmt.Fprintf(os.Stdout, "Response from `ImportLanguagesAPI.ImportLanguagesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportLanguagesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **int32** |  | 
 **file** | ***os.File** |  | 

### Return type

[**ImportLanguages**](ImportLanguages.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

