# \RiskAcceptanceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskAcceptanceCreate**](RiskAcceptanceAPI.md#RiskAcceptanceCreate) | **Post** /api/v2/risk_acceptance/ | 
[**RiskAcceptanceDeletePreviewList**](RiskAcceptanceAPI.md#RiskAcceptanceDeletePreviewList) | **Get** /api/v2/risk_acceptance/{id}/delete_preview/ | 
[**RiskAcceptanceDestroy**](RiskAcceptanceAPI.md#RiskAcceptanceDestroy) | **Delete** /api/v2/risk_acceptance/{id}/ | 
[**RiskAcceptanceDownloadProofRetrieve**](RiskAcceptanceAPI.md#RiskAcceptanceDownloadProofRetrieve) | **Get** /api/v2/risk_acceptance/{id}/download_proof/ | 
[**RiskAcceptanceList**](RiskAcceptanceAPI.md#RiskAcceptanceList) | **Get** /api/v2/risk_acceptance/ | 
[**RiskAcceptancePartialUpdate**](RiskAcceptanceAPI.md#RiskAcceptancePartialUpdate) | **Patch** /api/v2/risk_acceptance/{id}/ | 
[**RiskAcceptanceRetrieve**](RiskAcceptanceAPI.md#RiskAcceptanceRetrieve) | **Get** /api/v2/risk_acceptance/{id}/ | 
[**RiskAcceptanceUpdate**](RiskAcceptanceAPI.md#RiskAcceptanceUpdate) | **Put** /api/v2/risk_acceptance/{id}/ | 



## RiskAcceptanceCreate

> RiskAcceptance RiskAcceptanceCreate(ctx).RiskAcceptanceRequest(riskAcceptanceRequest).Execute()



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
	riskAcceptanceRequest := *openapiclient.NewRiskAcceptanceRequest("Name_example", int32(123), []int32{int32(123)}) // RiskAcceptanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceCreate(context.Background()).RiskAcceptanceRequest(riskAcceptanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptanceCreate`: RiskAcceptance
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptanceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **riskAcceptanceRequest** | [**RiskAcceptanceRequest**](RiskAcceptanceRequest.md) |  | 

### Return type

[**RiskAcceptance**](RiskAcceptance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskAcceptanceDeletePreviewList

> PaginatedDeletePreviewList RiskAcceptanceDeletePreviewList(ctx, id).Limit(limit).Offset(offset).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this risk_ acceptance.
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceDeletePreviewList(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceDeletePreviewList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptanceDeletePreviewList`: PaginatedDeletePreviewList
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptanceDeletePreviewList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this risk_ acceptance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceDeletePreviewListRequest struct via the builder pattern


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


## RiskAcceptanceDestroy

> RiskAcceptanceDestroy(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this risk_ acceptance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceDestroy(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this risk_ acceptance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceDestroyRequest struct via the builder pattern


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


## RiskAcceptanceDownloadProofRetrieve

> RiskAcceptanceProof RiskAcceptanceDownloadProofRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this risk_ acceptance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceDownloadProofRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceDownloadProofRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptanceDownloadProofRetrieve`: RiskAcceptanceProof
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptanceDownloadProofRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this risk_ acceptance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceDownloadProofRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskAcceptanceProof**](RiskAcceptanceProof.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskAcceptanceList

> PaginatedRiskAcceptanceList RiskAcceptanceList(ctx).AcceptedBy(acceptedBy).AcceptedFindings(acceptedFindings).Decision(decision).DecisionDetails(decisionDetails).ExpirationDate(expirationDate).ExpirationDateHandled(expirationDateHandled).ExpirationDateWarned(expirationDateWarned).Limit(limit).Name(name).Notes(notes).O(o).Offset(offset).Owner(owner).ReactivateExpired(reactivateExpired).Recommendation(recommendation).RecommendationDetails(recommendationDetails).RestartSlaExpired(restartSlaExpired).Execute()



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
	acceptedBy := "acceptedBy_example" // string |  (optional)
	acceptedFindings := []int32{int32(123)} // []int32 |  (optional)
	decision := "decision_example" // string | Risk treatment decision by risk owner  * `A` - Accept (The risk is acknowledged, yet remains) * `V` - Avoid (Do not engage with whatever creates the risk) * `M` - Mitigate (The risk still exists, yet compensating controls make it less of a threat) * `F` - Fix (The risk is eradicated) * `T` - Transfer (The risk is transferred to a 3rd party) (optional)
	decisionDetails := "decisionDetails_example" // string |  (optional)
	expirationDate := time.Now() // time.Time |  (optional)
	expirationDateHandled := time.Now() // time.Time |  (optional)
	expirationDateWarned := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 | Number of results to return per page. (optional)
	name := "name_example" // string |  (optional)
	notes := []int32{int32(123)} // []int32 |  (optional)
	o := []string{"O_example"} // []string | Ordering  * `name` - Name * `-name` - Name (descending) (optional)
	offset := int32(56) // int32 | The initial index from which to return the results. (optional)
	owner := int32(56) // int32 |  (optional)
	reactivateExpired := true // bool |  (optional)
	recommendation := "recommendation_example" // string | Recommendation from the security team.  * `A` - Accept (The risk is acknowledged, yet remains) * `V` - Avoid (Do not engage with whatever creates the risk) * `M` - Mitigate (The risk still exists, yet compensating controls make it less of a threat) * `F` - Fix (The risk is eradicated) * `T` - Transfer (The risk is transferred to a 3rd party) (optional)
	recommendationDetails := "recommendationDetails_example" // string |  (optional)
	restartSlaExpired := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceList(context.Background()).AcceptedBy(acceptedBy).AcceptedFindings(acceptedFindings).Decision(decision).DecisionDetails(decisionDetails).ExpirationDate(expirationDate).ExpirationDateHandled(expirationDateHandled).ExpirationDateWarned(expirationDateWarned).Limit(limit).Name(name).Notes(notes).O(o).Offset(offset).Owner(owner).ReactivateExpired(reactivateExpired).Recommendation(recommendation).RecommendationDetails(recommendationDetails).RestartSlaExpired(restartSlaExpired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptanceList`: PaginatedRiskAcceptanceList
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptanceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptedBy** | **string** |  | 
 **acceptedFindings** | **[]int32** |  | 
 **decision** | **string** | Risk treatment decision by risk owner  * &#x60;A&#x60; - Accept (The risk is acknowledged, yet remains) * &#x60;V&#x60; - Avoid (Do not engage with whatever creates the risk) * &#x60;M&#x60; - Mitigate (The risk still exists, yet compensating controls make it less of a threat) * &#x60;F&#x60; - Fix (The risk is eradicated) * &#x60;T&#x60; - Transfer (The risk is transferred to a 3rd party) | 
 **decisionDetails** | **string** |  | 
 **expirationDate** | **time.Time** |  | 
 **expirationDateHandled** | **time.Time** |  | 
 **expirationDateWarned** | **time.Time** |  | 
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** |  | 
 **notes** | **[]int32** |  | 
 **o** | **[]string** | Ordering  * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **owner** | **int32** |  | 
 **reactivateExpired** | **bool** |  | 
 **recommendation** | **string** | Recommendation from the security team.  * &#x60;A&#x60; - Accept (The risk is acknowledged, yet remains) * &#x60;V&#x60; - Avoid (Do not engage with whatever creates the risk) * &#x60;M&#x60; - Mitigate (The risk still exists, yet compensating controls make it less of a threat) * &#x60;F&#x60; - Fix (The risk is eradicated) * &#x60;T&#x60; - Transfer (The risk is transferred to a 3rd party) | 
 **recommendationDetails** | **string** |  | 
 **restartSlaExpired** | **bool** |  | 

### Return type

[**PaginatedRiskAcceptanceList**](PaginatedRiskAcceptanceList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskAcceptancePartialUpdate

> RiskAcceptance RiskAcceptancePartialUpdate(ctx, id).PatchedRiskAcceptanceRequest(patchedRiskAcceptanceRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this risk_ acceptance.
	patchedRiskAcceptanceRequest := *openapiclient.NewPatchedRiskAcceptanceRequest() // PatchedRiskAcceptanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptancePartialUpdate(context.Background(), id).PatchedRiskAcceptanceRequest(patchedRiskAcceptanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptancePartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptancePartialUpdate`: RiskAcceptance
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptancePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this risk_ acceptance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptancePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedRiskAcceptanceRequest** | [**PatchedRiskAcceptanceRequest**](PatchedRiskAcceptanceRequest.md) |  | 

### Return type

[**RiskAcceptance**](RiskAcceptance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskAcceptanceRetrieve

> RiskAcceptance RiskAcceptanceRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this risk_ acceptance.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptanceRetrieve`: RiskAcceptance
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptanceRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this risk_ acceptance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskAcceptance**](RiskAcceptance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskAcceptanceUpdate

> RiskAcceptance RiskAcceptanceUpdate(ctx, id).RiskAcceptanceRequest(riskAcceptanceRequest).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this risk_ acceptance.
	riskAcceptanceRequest := *openapiclient.NewRiskAcceptanceRequest("Name_example", int32(123), []int32{int32(123)}) // RiskAcceptanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskAcceptanceAPI.RiskAcceptanceUpdate(context.Background(), id).RiskAcceptanceRequest(riskAcceptanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskAcceptanceAPI.RiskAcceptanceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskAcceptanceUpdate`: RiskAcceptance
	fmt.Fprintf(os.Stdout, "Response from `RiskAcceptanceAPI.RiskAcceptanceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this risk_ acceptance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskAcceptanceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskAcceptanceRequest** | [**RiskAcceptanceRequest**](RiskAcceptanceRequest.md) |  | 

### Return type

[**RiskAcceptance**](RiskAcceptance.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

