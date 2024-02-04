# \ProductsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsCreate**](ProductsAPI.md#ProductsCreate) | **Post** /api/v2/products/ | 
[**ProductsDeletePreviewList**](ProductsAPI.md#ProductsDeletePreviewList) | **Get** /api/v2/products/{id}/delete_preview/ | 
[**ProductsDestroy**](ProductsAPI.md#ProductsDestroy) | **Delete** /api/v2/products/{id}/ | 
[**ProductsGenerateReportCreate**](ProductsAPI.md#ProductsGenerateReportCreate) | **Post** /api/v2/products/{id}/generate_report/ | 
[**ProductsList**](ProductsAPI.md#ProductsList) | **Get** /api/v2/products/ | 
[**ProductsPartialUpdate**](ProductsAPI.md#ProductsPartialUpdate) | **Patch** /api/v2/products/{id}/ | 
[**ProductsRetrieve**](ProductsAPI.md#ProductsRetrieve) | **Get** /api/v2/products/{id}/ | 
[**ProductsUpdate**](ProductsAPI.md#ProductsUpdate) | **Put** /api/v2/products/{id}/ | 



## ProductsCreate

> Product ProductsCreate(ctx).ProductRequest(productRequest).Execute()



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
	productRequest := *openapiclient.NewProductRequest("Name_example", "Description_example", int32(123)) // ProductRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsCreate(context.Background()).ProductRequest(productRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsCreate`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productRequest** | [**ProductRequest**](ProductRequest.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsDeletePreviewList

> PaginatedDeletePreviewList ProductsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsDeletePreviewListRequest struct via the builder pattern


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


## ProductsDestroy

> ProductsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductsAPI.ProductsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsDestroyRequest struct via the builder pattern


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


## ProductsGenerateReportCreate

> ReportGenerate ProductsGenerateReportCreate(ctx, id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product.
	reportGenerateOptionRequest := *openapiclient.NewReportGenerateOptionRequest() // ReportGenerateOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsGenerateReportCreate(context.Background(), id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsGenerateReportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsGenerateReportCreate`: ReportGenerate
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsGenerateReportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsGenerateReportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportGenerateOptionRequest** | [**ReportGenerateOptionRequest**](ReportGenerateOptionRequest.md) |  | 

### Return type

[**ReportGenerate**](ReportGenerate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsList

> PaginatedProductList ProductsList(ctx).BusinessCriticality(businessCriticality).Created(created).Description(description).ExternalAudience(externalAudience).HasTags(hasTags).Id(id).InternetAccessible(internetAccessible).Lifecycle(lifecycle).Limit(limit).Name(name).NameExact(nameExact).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).Origin(origin).OutsideOfSla(outsideOfSla).Platform(platform).Prefetch(prefetch).ProdNumericGrade(prodNumericGrade).ProdType(prodType).ProductManager(productManager).Regulations(regulations).Revenue(revenue).Tag(tag).Tags(tags).TeamManager(teamManager).TechnicalContact(technicalContact).Tid(tid).Updated(updated).UserRecords(userRecords).Execute()



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
	businessCriticality := "businessCriticality_example" // string |  (optional)
	created := time.Now() // time.Time | * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	description := "description_example" // string |  (optional)
	externalAudience := true // bool |  (optional)
	hasTags := true // bool | Has tags (optional)
	id := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	internetAccessible := true // bool |  (optional)
	lifecycle := "lifecycle_example" // string |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	nameExact := "nameExact_example" // string |  (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on product (optional)
	o := []string{"O_example"} // []string | Ordering  * `id` - Id * `-id` - Id (descending) * `tid` - Tid * `-tid` - Tid (descending) * `name` - Name * `-name` - Name (descending) * `created` - Created * `-created` - Created (descending) * `prod_numeric_grade` - Prod numeric grade * `-prod_numeric_grade` - Prod numeric grade (descending) * `business_criticality` - Business criticality * `-business_criticality` - Business criticality (descending) * `platform` - Platform * `-platform` - Platform (descending) * `lifecycle` - Lifecycle * `-lifecycle` - Lifecycle (descending) * `origin` - Origin * `-origin` - Origin (descending) * `revenue` - Revenue * `-revenue` - Revenue (descending) * `external_audience` - External audience * `-external_audience` - External audience (descending) * `internet_accessible` - Internet accessible * `-internet_accessible` - Internet accessible (descending) * `product_manager` - Product manager * `-product_manager` - Product manager (descending) * `product_manager__first_name` - Product manager  first name * `-product_manager__first_name` - Product manager  first name (descending) * `product_manager__last_name` - Product manager  last name * `-product_manager__last_name` - Product manager  last name (descending) * `technical_contact` - Technical contact * `-technical_contact` - Technical contact (descending) * `technical_contact__first_name` - Technical contact  first name * `-technical_contact__first_name` - Technical contact  first name (descending) * `technical_contact__last_name` - Technical contact  last name * `-technical_contact__last_name` - Technical contact  last name (descending) * `team_manager` - Team manager * `-team_manager` - Team manager (descending) * `team_manager__first_name` - Team manager  first name * `-team_manager__first_name` - Team manager  first name (descending) * `team_manager__last_name` - Team manager  last name * `-team_manager__last_name` - Team manager  last name (descending) * `prod_type` - Prod type * `-prod_type` - Prod type (descending) * `prod_type__name` - Prod type  name * `-prod_type__name` - Prod type  name (descending) * `updated` - Updated * `-updated` - Updated (descending) * `user_records` - User records * `-user_records` - User records (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	origin := "origin_example" // string |  (optional)
	outsideOfSla := float32(8.14) // float32 |  (optional)
	platform := "platform_example" // string |  (optional)
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)
	prodNumericGrade := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	prodType := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	productManager := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	regulations := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	revenue := float32(8.14) // float32 |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	teamManager := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	technicalContact := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	tid := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	updated := time.Now() // time.Time | * `None` - Any date * `1` - Today * `2` - Past 7 days * `3` - Past 30 days * `4` - Past 90 days * `5` - Current month * `6` - Current year * `7` - Past year (optional)
	userRecords := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsList(context.Background()).BusinessCriticality(businessCriticality).Created(created).Description(description).ExternalAudience(externalAudience).HasTags(hasTags).Id(id).InternetAccessible(internetAccessible).Lifecycle(lifecycle).Limit(limit).Name(name).NameExact(nameExact).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).Origin(origin).OutsideOfSla(outsideOfSla).Platform(platform).Prefetch(prefetch).ProdNumericGrade(prodNumericGrade).ProdType(prodType).ProductManager(productManager).Regulations(regulations).Revenue(revenue).Tag(tag).Tags(tags).TeamManager(teamManager).TechnicalContact(technicalContact).Tid(tid).Updated(updated).UserRecords(userRecords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsList`: PaginatedProductList
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessCriticality** | **string** |  | 
 **created** | **time.Time** | * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **description** | **string** |  | 
 **externalAudience** | **bool** |  | 
 **hasTags** | **bool** | Has tags | 
 **id** | **[]int32** | Multiple values may be separated by commas. | 
 **internetAccessible** | **bool** |  | 
 **lifecycle** | **string** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **nameExact** | **string** |  | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on product | 
 **o** | **[]string** | Ordering  * &#x60;id&#x60; - Id * &#x60;-id&#x60; - Id (descending) * &#x60;tid&#x60; - Tid * &#x60;-tid&#x60; - Tid (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;created&#x60; - Created * &#x60;-created&#x60; - Created (descending) * &#x60;prod_numeric_grade&#x60; - Prod numeric grade * &#x60;-prod_numeric_grade&#x60; - Prod numeric grade (descending) * &#x60;business_criticality&#x60; - Business criticality * &#x60;-business_criticality&#x60; - Business criticality (descending) * &#x60;platform&#x60; - Platform * &#x60;-platform&#x60; - Platform (descending) * &#x60;lifecycle&#x60; - Lifecycle * &#x60;-lifecycle&#x60; - Lifecycle (descending) * &#x60;origin&#x60; - Origin * &#x60;-origin&#x60; - Origin (descending) * &#x60;revenue&#x60; - Revenue * &#x60;-revenue&#x60; - Revenue (descending) * &#x60;external_audience&#x60; - External audience * &#x60;-external_audience&#x60; - External audience (descending) * &#x60;internet_accessible&#x60; - Internet accessible * &#x60;-internet_accessible&#x60; - Internet accessible (descending) * &#x60;product_manager&#x60; - Product manager * &#x60;-product_manager&#x60; - Product manager (descending) * &#x60;product_manager__first_name&#x60; - Product manager  first name * &#x60;-product_manager__first_name&#x60; - Product manager  first name (descending) * &#x60;product_manager__last_name&#x60; - Product manager  last name * &#x60;-product_manager__last_name&#x60; - Product manager  last name (descending) * &#x60;technical_contact&#x60; - Technical contact * &#x60;-technical_contact&#x60; - Technical contact (descending) * &#x60;technical_contact__first_name&#x60; - Technical contact  first name * &#x60;-technical_contact__first_name&#x60; - Technical contact  first name (descending) * &#x60;technical_contact__last_name&#x60; - Technical contact  last name * &#x60;-technical_contact__last_name&#x60; - Technical contact  last name (descending) * &#x60;team_manager&#x60; - Team manager * &#x60;-team_manager&#x60; - Team manager (descending) * &#x60;team_manager__first_name&#x60; - Team manager  first name * &#x60;-team_manager__first_name&#x60; - Team manager  first name (descending) * &#x60;team_manager__last_name&#x60; - Team manager  last name * &#x60;-team_manager__last_name&#x60; - Team manager  last name (descending) * &#x60;prod_type&#x60; - Prod type * &#x60;-prod_type&#x60; - Prod type (descending) * &#x60;prod_type__name&#x60; - Prod type  name * &#x60;-prod_type__name&#x60; - Prod type  name (descending) * &#x60;updated&#x60; - Updated * &#x60;-updated&#x60; - Updated (descending) * &#x60;user_records&#x60; - User records * &#x60;-user_records&#x60; - User records (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **origin** | **string** |  | 
 **outsideOfSla** | **float32** |  | 
 **platform** | **string** |  | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 
 **prodNumericGrade** | **[]int32** | Multiple values may be separated by commas. | 
 **prodType** | **[]int32** | Multiple values may be separated by commas. | 
 **productManager** | **[]int32** | Multiple values may be separated by commas. | 
 **regulations** | **[]int32** | Multiple values may be separated by commas. | 
 **revenue** | **float32** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **teamManager** | **[]int32** | Multiple values may be separated by commas. | 
 **technicalContact** | **[]int32** | Multiple values may be separated by commas. | 
 **tid** | **[]int32** | Multiple values may be separated by commas. | 
 **updated** | **time.Time** | * &#x60;None&#x60; - Any date * &#x60;1&#x60; - Today * &#x60;2&#x60; - Past 7 days * &#x60;3&#x60; - Past 30 days * &#x60;4&#x60; - Past 90 days * &#x60;5&#x60; - Current month * &#x60;6&#x60; - Current year * &#x60;7&#x60; - Past year | 
 **userRecords** | **[]int32** | Multiple values may be separated by commas. | 

### Return type

[**PaginatedProductList**](PaginatedProductList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsPartialUpdate

> Product ProductsPartialUpdate(ctx, id).PatchedProductRequest(patchedProductRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product.
	patchedProductRequest := *openapiclient.NewPatchedProductRequest() // PatchedProductRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsPartialUpdate(context.Background(), id).PatchedProductRequest(patchedProductRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsPartialUpdate`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedProductRequest** | [**PatchedProductRequest**](PatchedProductRequest.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsRetrieve

> Product ProductsRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsRetrieve`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**Product**](Product.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsUpdate

> Product ProductsUpdate(ctx, id).ProductRequest(productRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this product.
	productRequest := *openapiclient.NewProductRequest("Name_example", "Description_example", int32(123)) // ProductRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsUpdate(context.Background(), id).ProductRequest(productRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsUpdate`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this product. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productRequest** | [**ProductRequest**](ProductRequest.md) |  | 

### Return type

[**Product**](Product.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

