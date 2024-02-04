# \QuestionnaireAnswersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuestionnaireAnswersList**](QuestionnaireAnswersAPI.md#QuestionnaireAnswersList) | **Get** /api/v2/questionnaire_answers/ | 
[**QuestionnaireAnswersRetrieve**](QuestionnaireAnswersAPI.md#QuestionnaireAnswersRetrieve) | **Get** /api/v2/questionnaire_answers/{id}/ | 



## QuestionnaireAnswersList

> PaginatedQuestionnaireAnswerList QuestionnaireAnswersList(ctx).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.QuestionnaireAnswersAPI.QuestionnaireAnswersList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswersAPI.QuestionnaireAnswersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireAnswersList`: PaginatedQuestionnaireAnswerList
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswersAPI.QuestionnaireAnswersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireAnswersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedQuestionnaireAnswerList**](PaginatedQuestionnaireAnswerList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionnaireAnswersRetrieve

> QuestionnaireAnswer QuestionnaireAnswersRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this answer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionnaireAnswersAPI.QuestionnaireAnswersRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnswersAPI.QuestionnaireAnswersRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireAnswersRetrieve`: QuestionnaireAnswer
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnswersAPI.QuestionnaireAnswersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this answer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireAnswersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireAnswer**](QuestionnaireAnswer.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

