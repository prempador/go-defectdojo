# \EngagementsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngagementsAcceptRisksCreate**](EngagementsAPI.md#EngagementsAcceptRisksCreate) | **Post** /api/v2/engagements/{id}/accept_risks/ | 
[**EngagementsCloseCreate**](EngagementsAPI.md#EngagementsCloseCreate) | **Post** /api/v2/engagements/{id}/close/ | 
[**EngagementsCompleteChecklistCreate**](EngagementsAPI.md#EngagementsCompleteChecklistCreate) | **Post** /api/v2/engagements/{id}/complete_checklist/ | 
[**EngagementsCompleteChecklistRetrieve**](EngagementsAPI.md#EngagementsCompleteChecklistRetrieve) | **Get** /api/v2/engagements/{id}/complete_checklist/ | 
[**EngagementsCreate**](EngagementsAPI.md#EngagementsCreate) | **Post** /api/v2/engagements/ | 
[**EngagementsDeletePreviewList**](EngagementsAPI.md#EngagementsDeletePreviewList) | **Get** /api/v2/engagements/{id}/delete_preview/ | 
[**EngagementsDestroy**](EngagementsAPI.md#EngagementsDestroy) | **Delete** /api/v2/engagements/{id}/ | 
[**EngagementsFilesCreate**](EngagementsAPI.md#EngagementsFilesCreate) | **Post** /api/v2/engagements/{id}/files/ | 
[**EngagementsFilesDownloadRetrieve**](EngagementsAPI.md#EngagementsFilesDownloadRetrieve) | **Get** /api/v2/engagements/{id}/files/download/{file_id}/ | 
[**EngagementsFilesRetrieve**](EngagementsAPI.md#EngagementsFilesRetrieve) | **Get** /api/v2/engagements/{id}/files/ | 
[**EngagementsGenerateReportCreate**](EngagementsAPI.md#EngagementsGenerateReportCreate) | **Post** /api/v2/engagements/{id}/generate_report/ | 
[**EngagementsList**](EngagementsAPI.md#EngagementsList) | **Get** /api/v2/engagements/ | 
[**EngagementsNotesCreate**](EngagementsAPI.md#EngagementsNotesCreate) | **Post** /api/v2/engagements/{id}/notes/ | 
[**EngagementsNotesRetrieve**](EngagementsAPI.md#EngagementsNotesRetrieve) | **Get** /api/v2/engagements/{id}/notes/ | 
[**EngagementsPartialUpdate**](EngagementsAPI.md#EngagementsPartialUpdate) | **Patch** /api/v2/engagements/{id}/ | 
[**EngagementsReopenCreate**](EngagementsAPI.md#EngagementsReopenCreate) | **Post** /api/v2/engagements/{id}/reopen/ | 
[**EngagementsRetrieve**](EngagementsAPI.md#EngagementsRetrieve) | **Get** /api/v2/engagements/{id}/ | 
[**EngagementsUpdate**](EngagementsAPI.md#EngagementsUpdate) | **Put** /api/v2/engagements/{id}/ | 



## EngagementsAcceptRisksCreate

> []RiskAcceptance EngagementsAcceptRisksCreate(ctx, id).AcceptedRiskRequest(acceptedRiskRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	acceptedRiskRequest := []openapiclient.AcceptedRiskRequest{*openapiclient.NewAcceptedRiskRequest("VulnerabilityId_example", "Justification_example", "AcceptedBy_example")} // []AcceptedRiskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsAcceptRisksCreate(context.Background(), id).AcceptedRiskRequest(acceptedRiskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsAcceptRisksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsAcceptRisksCreate`: []RiskAcceptance
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsAcceptRisksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsAcceptRisksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptedRiskRequest** | [**[]AcceptedRiskRequest**](AcceptedRiskRequest.md) |  | 

### Return type

[**[]RiskAcceptance**](RiskAcceptance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsCloseCreate

> EngagementsCloseCreate(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementsAPI.EngagementsCloseCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsCloseCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsCloseCreateRequest struct via the builder pattern


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


## EngagementsCompleteChecklistCreate

> EngagementCheckList EngagementsCompleteChecklistCreate(ctx, id).EngagementCheckListRequest(engagementCheckListRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	engagementCheckListRequest := *openapiclient.NewEngagementCheckListRequest() // EngagementCheckListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsCompleteChecklistCreate(context.Background(), id).EngagementCheckListRequest(engagementCheckListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsCompleteChecklistCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsCompleteChecklistCreate`: EngagementCheckList
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsCompleteChecklistCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsCompleteChecklistCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engagementCheckListRequest** | [**EngagementCheckListRequest**](EngagementCheckListRequest.md) |  | 

### Return type

[**EngagementCheckList**](EngagementCheckList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsCompleteChecklistRetrieve

> Engagement EngagementsCompleteChecklistRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsCompleteChecklistRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsCompleteChecklistRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsCompleteChecklistRetrieve`: Engagement
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsCompleteChecklistRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsCompleteChecklistRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Engagement**](Engagement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsCreate

> Engagement EngagementsCreate(ctx).EngagementRequest(engagementRequest).Execute()



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
	engagementRequest := *openapiclient.NewEngagementRequest(time.Now(), time.Now(), int32(123)) // EngagementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsCreate(context.Background()).EngagementRequest(engagementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsCreate`: Engagement
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **engagementRequest** | [**EngagementRequest**](EngagementRequest.md) |  | 

### Return type

[**Engagement**](Engagement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsDeletePreviewList

> PaginatedDeletePreviewList EngagementsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsDeletePreviewListRequest struct via the builder pattern


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


## EngagementsDestroy

> EngagementsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementsAPI.EngagementsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsDestroyRequest struct via the builder pattern


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


## EngagementsFilesCreate

> File EngagementsFilesCreate(ctx, id).Title(title).File(file).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	title := "title_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsFilesCreate(context.Background(), id).Title(title).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsFilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsFilesCreate`: File
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsFilesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsFilesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

[**File**](File.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsFilesDownloadRetrieve

> RawFile EngagementsFilesDownloadRetrieve(ctx, fileId, id).Execute()



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
	fileId := "fileId_example" // string | 
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsFilesDownloadRetrieve(context.Background(), fileId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsFilesDownloadRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsFilesDownloadRetrieve`: RawFile
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsFilesDownloadRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsFilesDownloadRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RawFile**](RawFile.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsFilesRetrieve

> EngagementToFiles EngagementsFilesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsFilesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsFilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsFilesRetrieve`: EngagementToFiles
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsFilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsFilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngagementToFiles**](EngagementToFiles.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsGenerateReportCreate

> ReportGenerate EngagementsGenerateReportCreate(ctx, id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	reportGenerateOptionRequest := *openapiclient.NewReportGenerateOptionRequest() // ReportGenerateOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsGenerateReportCreate(context.Background(), id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsGenerateReportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsGenerateReportCreate`: ReportGenerate
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsGenerateReportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsGenerateReportCreateRequest struct via the builder pattern


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


## EngagementsList

> PaginatedEngagementList EngagementsList(ctx).Active(active).ApiTest(apiTest).HasTags(hasTags).Id(id).Limit(limit).Name(name).NotProductTags(notProductTags).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).PenTest(penTest).Product(product).ProductProdType(productProdType).ProductTags(productTags).ReportType(reportType).Requester(requester).Status(status).Tag(tag).Tags(tags).TargetEnd(targetEnd).TargetStart(targetStart).ThreatModel(threatModel).Updated(updated).Version(version).Execute()



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
	active := true // bool |  (optional)
	apiTest := true // bool |  (optional)
	hasTags := true // bool | Has tags (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	notProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on product (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	o := []string{"O_example"} // []string | Ordering  * `name` - Engagement Name * `-name` - Engagement Name (descending) * `version` - Version * `-version` - Version (descending) * `target_start` - Target start * `-target_start` - Target start (descending) * `target_end` - Target end * `-target_end` - Target end (descending) * `status` - Status * `-status` - Status (descending) * `lead` - Lead * `-lead` - Lead (descending) * `created` - Created * `-created` - Created (descending) * `updated` - Updated * `-updated` - Updated (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	penTest := true // bool |  (optional)
	product := int32(56) // int32 |  (optional)
	productProdType := []int32{int32(123)} // []int32 | Multiple values may be separated by commas. (optional)
	productTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on product (optional)
	reportType := int32(56) // int32 |  (optional)
	requester := int32(56) // int32 |  (optional)
	status := "status_example" // string | * `Not Started` - Not Started * `Blocked` - Blocked * `Cancelled` - Cancelled * `Completed` - Completed * `In Progress` - In Progress * `On Hold` - On Hold * `Waiting for Resource` - Waiting for Resource (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	targetEnd := time.Now() // string |  (optional)
	targetStart := time.Now() // string |  (optional)
	threatModel := true // bool |  (optional)
	updated := time.Now() // time.Time |  (optional)
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsList(context.Background()).Active(active).ApiTest(apiTest).HasTags(hasTags).Id(id).Limit(limit).Name(name).NotProductTags(notProductTags).NotTag(notTag).NotTags(notTags).O(o).Offset(offset).PenTest(penTest).Product(product).ProductProdType(productProdType).ProductTags(productTags).ReportType(reportType).Requester(requester).Status(status).Tag(tag).Tags(tags).TargetEnd(targetEnd).TargetStart(targetStart).ThreatModel(threatModel).Updated(updated).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsList`: PaginatedEngagementList
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** |  | 
 **apiTest** | **bool** |  | 
 **hasTags** | **bool** | Has tags | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **notProductTags** | **[]string** | Comma separated list of exact tags not present on product | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **o** | **[]string** | Ordering  * &#x60;name&#x60; - Engagement Name * &#x60;-name&#x60; - Engagement Name (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;target_start&#x60; - Target start * &#x60;-target_start&#x60; - Target start (descending) * &#x60;target_end&#x60; - Target end * &#x60;-target_end&#x60; - Target end (descending) * &#x60;status&#x60; - Status * &#x60;-status&#x60; - Status (descending) * &#x60;lead&#x60; - Lead * &#x60;-lead&#x60; - Lead (descending) * &#x60;created&#x60; - Created * &#x60;-created&#x60; - Created (descending) * &#x60;updated&#x60; - Updated * &#x60;-updated&#x60; - Updated (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **penTest** | **bool** |  | 
 **product** | **int32** |  | 
 **productProdType** | **[]int32** | Multiple values may be separated by commas. | 
 **productTags** | **[]string** | Comma separated list of exact tags present on product | 
 **reportType** | **int32** |  | 
 **requester** | **int32** |  | 
 **status** | **string** | * &#x60;Not Started&#x60; - Not Started * &#x60;Blocked&#x60; - Blocked * &#x60;Cancelled&#x60; - Cancelled * &#x60;Completed&#x60; - Completed * &#x60;In Progress&#x60; - In Progress * &#x60;On Hold&#x60; - On Hold * &#x60;Waiting for Resource&#x60; - Waiting for Resource | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **targetEnd** | **string** |  | 
 **targetStart** | **string** |  | 
 **threatModel** | **bool** |  | 
 **updated** | **time.Time** |  | 
 **version** | **string** |  | 

### Return type

[**PaginatedEngagementList**](PaginatedEngagementList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsNotesCreate

> Note EngagementsNotesCreate(ctx, id).AddNewNoteOptionRequest(addNewNoteOptionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	addNewNoteOptionRequest := *openapiclient.NewAddNewNoteOptionRequest("Entry_example") // AddNewNoteOptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsNotesCreate(context.Background(), id).AddNewNoteOptionRequest(addNewNoteOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsNotesCreate`: Note
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addNewNoteOptionRequest** | [**AddNewNoteOptionRequest**](AddNewNoteOptionRequest.md) |  | 

### Return type

[**Note**](Note.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsNotesRetrieve

> EngagementToNotes EngagementsNotesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsNotesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsNotesRetrieve`: EngagementToNotes
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsNotesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EngagementToNotes**](EngagementToNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsPartialUpdate

> Engagement EngagementsPartialUpdate(ctx, id).PatchedEngagementRequest(patchedEngagementRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	patchedEngagementRequest := *openapiclient.NewPatchedEngagementRequest() // PatchedEngagementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsPartialUpdate(context.Background(), id).PatchedEngagementRequest(patchedEngagementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsPartialUpdate`: Engagement
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEngagementRequest** | [**PatchedEngagementRequest**](PatchedEngagementRequest.md) |  | 

### Return type

[**Engagement**](Engagement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsReopenCreate

> EngagementsReopenCreate(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EngagementsAPI.EngagementsReopenCreate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsReopenCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsReopenCreateRequest struct via the builder pattern


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


## EngagementsRetrieve

> Engagement EngagementsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsRetrieve`: Engagement
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Engagement**](Engagement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngagementsUpdate

> Engagement EngagementsUpdate(ctx, id).EngagementRequest(engagementRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this engagement.
	engagementRequest := *openapiclient.NewEngagementRequest(time.Now(), time.Now(), int32(123)) // EngagementRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementsAPI.EngagementsUpdate(context.Background(), id).EngagementRequest(engagementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementsAPI.EngagementsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngagementsUpdate`: Engagement
	fmt.Fprintf(os.Stdout, "Response from `EngagementsAPI.EngagementsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this engagement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEngagementsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **engagementRequest** | [**EngagementRequest**](EngagementRequest.md) |  | 

### Return type

[**Engagement**](Engagement.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

