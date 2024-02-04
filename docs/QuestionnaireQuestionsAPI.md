# \QuestionnaireQuestionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuestionnaireQuestionsList**](QuestionnaireQuestionsAPI.md#QuestionnaireQuestionsList) | **Get** /api/v2/questionnaire_questions/ | 
[**QuestionnaireQuestionsRetrieve**](QuestionnaireQuestionsAPI.md#QuestionnaireQuestionsRetrieve) | **Get** /api/v2/questionnaire_questions/{id}/ | 



## QuestionnaireQuestionsList

> PaginatedQuestionnaireQuestionList QuestionnaireQuestionsList(ctx).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.QuestionnaireQuestionsAPI.QuestionnaireQuestionsList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireQuestionsAPI.QuestionnaireQuestionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireQuestionsList`: PaginatedQuestionnaireQuestionList
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireQuestionsAPI.QuestionnaireQuestionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireQuestionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedQuestionnaireQuestionList**](PaginatedQuestionnaireQuestionList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionnaireQuestionsRetrieve

> QuestionnaireQuestion QuestionnaireQuestionsRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this question.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionnaireQuestionsAPI.QuestionnaireQuestionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireQuestionsAPI.QuestionnaireQuestionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireQuestionsRetrieve`: QuestionnaireQuestion
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireQuestionsAPI.QuestionnaireQuestionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this question. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireQuestionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireQuestion**](QuestionnaireQuestion.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

