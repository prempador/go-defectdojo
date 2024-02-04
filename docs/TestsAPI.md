# \TestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestsAcceptRisksCreate**](TestsAPI.md#TestsAcceptRisksCreate) | **Post** /api/v2/tests/{id}/accept_risks/ | 
[**TestsCreate**](TestsAPI.md#TestsCreate) | **Post** /api/v2/tests/ | 
[**TestsDeletePreviewList**](TestsAPI.md#TestsDeletePreviewList) | **Get** /api/v2/tests/{id}/delete_preview/ | 
[**TestsDestroy**](TestsAPI.md#TestsDestroy) | **Delete** /api/v2/tests/{id}/ | 
[**TestsFilesCreate**](TestsAPI.md#TestsFilesCreate) | **Post** /api/v2/tests/{id}/files/ | 
[**TestsFilesDownloadRetrieve**](TestsAPI.md#TestsFilesDownloadRetrieve) | **Get** /api/v2/tests/{id}/files/download/{file_id}/ | 
[**TestsFilesRetrieve**](TestsAPI.md#TestsFilesRetrieve) | **Get** /api/v2/tests/{id}/files/ | 
[**TestsGenerateReportCreate**](TestsAPI.md#TestsGenerateReportCreate) | **Post** /api/v2/tests/{id}/generate_report/ | 
[**TestsList**](TestsAPI.md#TestsList) | **Get** /api/v2/tests/ | 
[**TestsNotesCreate**](TestsAPI.md#TestsNotesCreate) | **Post** /api/v2/tests/{id}/notes/ | 
[**TestsNotesRetrieve**](TestsAPI.md#TestsNotesRetrieve) | **Get** /api/v2/tests/{id}/notes/ | 
[**TestsPartialUpdate**](TestsAPI.md#TestsPartialUpdate) | **Patch** /api/v2/tests/{id}/ | 
[**TestsRetrieve**](TestsAPI.md#TestsRetrieve) | **Get** /api/v2/tests/{id}/ | 
[**TestsUpdate**](TestsAPI.md#TestsUpdate) | **Put** /api/v2/tests/{id}/ | 



## TestsAcceptRisksCreate

> []RiskAcceptance TestsAcceptRisksCreate(ctx, id).AcceptedRiskRequest(acceptedRiskRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	acceptedRiskRequest := []openapiclient.AcceptedRiskRequest{*openapiclient.NewAcceptedRiskRequest("VulnerabilityId_example", "Justification_example", "AcceptedBy_example")} // []AcceptedRiskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsAcceptRisksCreate(context.Background(), id).AcceptedRiskRequest(acceptedRiskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsAcceptRisksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsAcceptRisksCreate`: []RiskAcceptance
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsAcceptRisksCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsAcceptRisksCreateRequest struct via the builder pattern


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


## TestsCreate

> TestCreate TestsCreate(ctx).TestCreateRequest(testCreateRequest).Execute()



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
	testCreateRequest := *openapiclient.NewTestCreateRequest(int32(123), time.Now(), time.Now(), int32(123)) // TestCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsCreate(context.Background()).TestCreateRequest(testCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsCreate`: TestCreate
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testCreateRequest** | [**TestCreateRequest**](TestCreateRequest.md) |  | 

### Return type

[**TestCreate**](TestCreate.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsDeletePreviewList

> PaginatedDeletePreviewList TestsDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsDeletePreviewListRequest struct via the builder pattern


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


## TestsDestroy

> TestsDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TestsAPI.TestsDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsDestroyRequest struct via the builder pattern


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


## TestsFilesCreate

> File TestsFilesCreate(ctx, id).Title(title).File(file).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	title := "title_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsFilesCreate(context.Background(), id).Title(title).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsFilesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsFilesCreate`: File
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsFilesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsFilesCreateRequest struct via the builder pattern


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


## TestsFilesDownloadRetrieve

> RawFile TestsFilesDownloadRetrieve(ctx, fileId, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsFilesDownloadRetrieve(context.Background(), fileId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsFilesDownloadRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsFilesDownloadRetrieve`: RawFile
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsFilesDownloadRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsFilesDownloadRetrieveRequest struct via the builder pattern


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


## TestsFilesRetrieve

> TestToFiles TestsFilesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsFilesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsFilesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsFilesRetrieve`: TestToFiles
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsFilesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsFilesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestToFiles**](TestToFiles.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsGenerateReportCreate

> ReportGenerate TestsGenerateReportCreate(ctx, id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	reportGenerateOptionRequest := *openapiclient.NewReportGenerateOptionRequest() // ReportGenerateOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsGenerateReportCreate(context.Background(), id).ReportGenerateOptionRequest(reportGenerateOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsGenerateReportCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsGenerateReportCreate`: ReportGenerate
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsGenerateReportCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsGenerateReportCreateRequest struct via the builder pattern


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


## TestsList

> PaginatedTestList TestsList(ctx).ActualTime(actualTime).ApiScanConfiguration(apiScanConfiguration).BranchTag(branchTag).BuildId(buildId).CommitHash(commitHash).Engagement(engagement).EngagementProductTags(engagementProductTags).EngagementTags(engagementTags).HasTags(hasTags).Id(id).Limit(limit).NotEngagementProductTags(notEngagementProductTags).NotEngagementTags(notEngagementTags).NotTag(notTag).NotTags(notTags).Notes(notes).O(o).Offset(offset).PercentComplete(percentComplete).ScanType(scanType).Tag(tag).Tags(tags).TargetEnd(targetEnd).TargetStart(targetStart).TestType(testType).Title(title).Version(version).Execute()



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
	actualTime := "actualTime_example" // string |  (optional)
	apiScanConfiguration := int32(56) // int32 |  (optional)
	branchTag := "branchTag_example" // string |  (optional)
	buildId := "buildId_example" // string |  (optional)
	commitHash := "commitHash_example" // string |  (optional)
	engagement := int32(56) // int32 |  (optional)
	engagementProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on product (optional)
	engagementTags := []string{"Inner_example"} // []string | Comma separated list of exact tags present on engagement (optional)
	hasTags := true // bool | Has tags (optional)
	id := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	notEngagementProductTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on product (optional)
	notEngagementTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on engagement (optional)
	notTag := "notTag_example" // string | Not Tag name contains (optional)
	notTags := []string{"Inner_example"} // []string | Comma separated list of exact tags not present on model (optional)
	notes := []int32{int32(123)} // []int32 |  (optional)
	o := []string{"O_example"} // []string | Ordering  * `title` - Title * `-title` - Title (descending) * `version` - Version * `-version` - Version (descending) * `target_start` - Target start * `-target_start` - Target start (descending) * `target_end` - Target end * `-target_end` - Target end (descending) * `test_type` - Test type * `-test_type` - Test type (descending) * `lead` - Lead * `-lead` - Lead (descending) * `branch_tag` - Branch tag * `-branch_tag` - Branch tag (descending) * `build_id` - Build id * `-build_id` - Build id (descending) * `commit_hash` - Commit hash * `-commit_hash` - Commit hash (descending) * `api_scan_configuration` - Api scan configuration * `-api_scan_configuration` - Api scan configuration (descending) * `engagement` - Engagement * `-engagement` - Engagement (descending) * `created` - Created * `-created` - Created (descending) * `updated` - Updated * `-updated` - Updated (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	percentComplete := int32(56) // int32 |  (optional)
	scanType := "scanType_example" // string |  (optional)
	tag := "tag_example" // string | Tag name contains (optional)
	tags := []string{"Inner_example"} // []string | Comma separated list of exact tags (optional)
	targetEnd := time.Now() // time.Time |  (optional)
	targetStart := time.Now() // time.Time |  (optional)
	testType := int32(56) // int32 |  (optional)
	title := "title_example" // string |  (optional)
	version := "version_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsList(context.Background()).ActualTime(actualTime).ApiScanConfiguration(apiScanConfiguration).BranchTag(branchTag).BuildId(buildId).CommitHash(commitHash).Engagement(engagement).EngagementProductTags(engagementProductTags).EngagementTags(engagementTags).HasTags(hasTags).Id(id).Limit(limit).NotEngagementProductTags(notEngagementProductTags).NotEngagementTags(notEngagementTags).NotTag(notTag).NotTags(notTags).Notes(notes).O(o).Offset(offset).PercentComplete(percentComplete).ScanType(scanType).Tag(tag).Tags(tags).TargetEnd(targetEnd).TargetStart(targetStart).TestType(testType).Title(title).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsList`: PaginatedTestList
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actualTime** | **string** |  | 
 **apiScanConfiguration** | **int32** |  | 
 **branchTag** | **string** |  | 
 **buildId** | **string** |  | 
 **commitHash** | **string** |  | 
 **engagement** | **int32** |  | 
 **engagementProductTags** | **[]string** | Comma separated list of exact tags present on product | 
 **engagementTags** | **[]string** | Comma separated list of exact tags present on engagement | 
 **hasTags** | **bool** | Has tags | 
 **id** | **int32** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **notEngagementProductTags** | **[]string** | Comma separated list of exact tags not present on product | 
 **notEngagementTags** | **[]string** | Comma separated list of exact tags not present on engagement | 
 **notTag** | **string** | Not Tag name contains | 
 **notTags** | **[]string** | Comma separated list of exact tags not present on model | 
 **notes** | **[]int32** |  | 
 **o** | **[]string** | Ordering  * &#x60;title&#x60; - Title * &#x60;-title&#x60; - Title (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;target_start&#x60; - Target start * &#x60;-target_start&#x60; - Target start (descending) * &#x60;target_end&#x60; - Target end * &#x60;-target_end&#x60; - Target end (descending) * &#x60;test_type&#x60; - Test type * &#x60;-test_type&#x60; - Test type (descending) * &#x60;lead&#x60; - Lead * &#x60;-lead&#x60; - Lead (descending) * &#x60;branch_tag&#x60; - Branch tag * &#x60;-branch_tag&#x60; - Branch tag (descending) * &#x60;build_id&#x60; - Build id * &#x60;-build_id&#x60; - Build id (descending) * &#x60;commit_hash&#x60; - Commit hash * &#x60;-commit_hash&#x60; - Commit hash (descending) * &#x60;api_scan_configuration&#x60; - Api scan configuration * &#x60;-api_scan_configuration&#x60; - Api scan configuration (descending) * &#x60;engagement&#x60; - Engagement * &#x60;-engagement&#x60; - Engagement (descending) * &#x60;created&#x60; - Created * &#x60;-created&#x60; - Created (descending) * &#x60;updated&#x60; - Updated * &#x60;-updated&#x60; - Updated (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **percentComplete** | **int32** |  | 
 **scanType** | **string** |  | 
 **tag** | **string** | Tag name contains | 
 **tags** | **[]string** | Comma separated list of exact tags | 
 **targetEnd** | **time.Time** |  | 
 **targetStart** | **time.Time** |  | 
 **testType** | **int32** |  | 
 **title** | **string** |  | 
 **version** | **string** |  | 

### Return type

[**PaginatedTestList**](PaginatedTestList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsNotesCreate

> Note TestsNotesCreate(ctx, id).AddNewNoteOptionRequest(addNewNoteOptionRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	addNewNoteOptionRequest := *openapiclient.NewAddNewNoteOptionRequest("Entry_example") // AddNewNoteOptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsNotesCreate(context.Background(), id).AddNewNoteOptionRequest(addNewNoteOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsNotesCreate`: Note
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsNotesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsNotesCreateRequest struct via the builder pattern


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


## TestsNotesRetrieve

> TestToNotes TestsNotesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsNotesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsNotesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsNotesRetrieve`: TestToNotes
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsNotesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsNotesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestToNotes**](TestToNotes.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsPartialUpdate

> Test TestsPartialUpdate(ctx, id).PatchedTestRequest(patchedTestRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	patchedTestRequest := *openapiclient.NewPatchedTestRequest() // PatchedTestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsPartialUpdate(context.Background(), id).PatchedTestRequest(patchedTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsPartialUpdate`: Test
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedTestRequest** | [**PatchedTestRequest**](PatchedTestRequest.md) |  | 

### Return type

[**Test**](Test.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsRetrieve

> Test TestsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsRetrieve`: Test
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Test**](Test.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestsUpdate

> Test TestsUpdate(ctx, id).TestRequest(testRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this test.
	testRequest := *openapiclient.NewTestRequest(time.Now(), time.Now(), int32(123)) // TestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TestsUpdate(context.Background(), id).TestRequest(testRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TestsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestsUpdate`: Test
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TestsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this test. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRequest** | [**TestRequest**](TestRequest.md) |  | 

### Return type

[**Test**](Test.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

