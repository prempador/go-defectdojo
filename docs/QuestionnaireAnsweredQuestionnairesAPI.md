# \QuestionnaireAnsweredQuestionnairesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuestionnaireAnsweredQuestionnairesList**](QuestionnaireAnsweredQuestionnairesAPI.md#QuestionnaireAnsweredQuestionnairesList) | **Get** /api/v2/questionnaire_answered_questionnaires/ | 
[**QuestionnaireAnsweredQuestionnairesRetrieve**](QuestionnaireAnsweredQuestionnairesAPI.md#QuestionnaireAnsweredQuestionnairesRetrieve) | **Get** /api/v2/questionnaire_answered_questionnaires/{id}/ | 



## QuestionnaireAnsweredQuestionnairesList

> PaginatedQuestionnaireAnsweredSurveyList QuestionnaireAnsweredQuestionnairesList(ctx).Limit(limit).Offset(offset).Prefetch(prefetch).Execute()



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
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionnaireAnsweredQuestionnairesAPI.QuestionnaireAnsweredQuestionnairesList(context.Background()).Limit(limit).Offset(offset).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnsweredQuestionnairesAPI.QuestionnaireAnsweredQuestionnairesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireAnsweredQuestionnairesList`: PaginatedQuestionnaireAnsweredSurveyList
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnsweredQuestionnairesAPI.QuestionnaireAnsweredQuestionnairesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireAnsweredQuestionnairesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**PaginatedQuestionnaireAnsweredSurveyList**](PaginatedQuestionnaireAnsweredSurveyList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionnaireAnsweredQuestionnairesRetrieve

> QuestionnaireAnsweredSurvey QuestionnaireAnsweredQuestionnairesRetrieve(ctx, id).Prefetch(prefetch).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this Answered Engagement Survey.
	prefetch := []string{"Prefetch_example"} // []string | List of fields for which to prefetch model instances and add those to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionnaireAnsweredQuestionnairesAPI.QuestionnaireAnsweredQuestionnairesRetrieve(context.Background(), id).Prefetch(prefetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireAnsweredQuestionnairesAPI.QuestionnaireAnsweredQuestionnairesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireAnsweredQuestionnairesRetrieve`: QuestionnaireAnsweredSurvey
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireAnsweredQuestionnairesAPI.QuestionnaireAnsweredQuestionnairesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Answered Engagement Survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireAnsweredQuestionnairesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefetch** | **[]string** | List of fields for which to prefetch model instances and add those to the response | 

### Return type

[**QuestionnaireAnsweredSurvey**](QuestionnaireAnsweredSurvey.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

