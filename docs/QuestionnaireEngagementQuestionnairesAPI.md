# \QuestionnaireEngagementQuestionnairesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuestionnaireEngagementQuestionnairesList**](QuestionnaireEngagementQuestionnairesAPI.md#QuestionnaireEngagementQuestionnairesList) | **Get** /api/v2/questionnaire_engagement_questionnaires/ | 
[**QuestionnaireEngagementQuestionnairesRetrieve**](QuestionnaireEngagementQuestionnairesAPI.md#QuestionnaireEngagementQuestionnairesRetrieve) | **Get** /api/v2/questionnaire_engagement_questionnaires/{id}/ | 



## QuestionnaireEngagementQuestionnairesList

> PaginatedQuestionnaireEngagementSurveyList QuestionnaireEngagementQuestionnairesList(ctx).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.QuestionnaireEngagementQuestionnairesAPI.QuestionnaireEngagementQuestionnairesList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireEngagementQuestionnairesAPI.QuestionnaireEngagementQuestionnairesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireEngagementQuestionnairesList`: PaginatedQuestionnaireEngagementSurveyList
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireEngagementQuestionnairesAPI.QuestionnaireEngagementQuestionnairesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireEngagementQuestionnairesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedQuestionnaireEngagementSurveyList**](PaginatedQuestionnaireEngagementSurveyList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionnaireEngagementQuestionnairesRetrieve

> QuestionnaireEngagementSurvey QuestionnaireEngagementQuestionnairesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this Engagement Survey.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionnaireEngagementQuestionnairesAPI.QuestionnaireEngagementQuestionnairesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireEngagementQuestionnairesAPI.QuestionnaireEngagementQuestionnairesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireEngagementQuestionnairesRetrieve`: QuestionnaireEngagementSurvey
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireEngagementQuestionnairesAPI.QuestionnaireEngagementQuestionnairesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this Engagement Survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireEngagementQuestionnairesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireEngagementSurvey**](QuestionnaireEngagementSurvey.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

