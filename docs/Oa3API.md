# \Oa3API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oa3SchemaRetrieve**](Oa3API.md#Oa3SchemaRetrieve) | **Get** /api/v2/oa3/schema/ | 



## Oa3SchemaRetrieve

> map[string]interface{} Oa3SchemaRetrieve(ctx).Format(format).Lang(lang).Execute()





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
	format := "format_example" // string |  (optional)
	lang := "lang_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Oa3API.Oa3SchemaRetrieve(context.Background()).Format(format).Lang(lang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Oa3API.Oa3SchemaRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Oa3SchemaRetrieve`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Oa3API.Oa3SchemaRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOa3SchemaRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **lang** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/vnd.oai.openapi, application/yaml, application/vnd.oai.openapi+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

