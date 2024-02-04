# \QuestionnaireGeneralQuestionnairesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuestionnaireGeneralQuestionnairesList**](QuestionnaireGeneralQuestionnairesAPI.md#QuestionnaireGeneralQuestionnairesList) | **Get** /api/v2/questionnaire_general_questionnaires/ | 
[**QuestionnaireGeneralQuestionnairesRetrieve**](QuestionnaireGeneralQuestionnairesAPI.md#QuestionnaireGeneralQuestionnairesRetrieve) | **Get** /api/v2/questionnaire_general_questionnaires/{id}/ | 



## QuestionnaireGeneralQuestionnairesList

> PaginatedQuestionnaireGeneralSurveyList QuestionnaireGeneralQuestionnairesList(ctx).Limit(limit).Offset(offset).Execute()



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
	resp, r, err := apiClient.QuestionnaireGeneralQuestionnairesAPI.QuestionnaireGeneralQuestionnairesList(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireGeneralQuestionnairesAPI.QuestionnaireGeneralQuestionnairesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireGeneralQuestionnairesList`: PaginatedQuestionnaireGeneralSurveyList
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireGeneralQuestionnairesAPI.QuestionnaireGeneralQuestionnairesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireGeneralQuestionnairesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 

### Return type

[**PaginatedQuestionnaireGeneralSurveyList**](PaginatedQuestionnaireGeneralSurveyList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuestionnaireGeneralQuestionnairesRetrieve

> QuestionnaireGeneralSurvey QuestionnaireGeneralQuestionnairesRetrieve(ctx, id).Execute()



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
	id := int32(56) // int32 | A unique integer value identifying this General Engagement Survey.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuestionnaireGeneralQuestionnairesAPI.QuestionnaireGeneralQuestionnairesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuestionnaireGeneralQuestionnairesAPI.QuestionnaireGeneralQuestionnairesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuestionnaireGeneralQuestionnairesRetrieve`: QuestionnaireGeneralSurvey
	fmt.Fprintf(os.Stdout, "Response from `QuestionnaireGeneralQuestionnairesAPI.QuestionnaireGeneralQuestionnairesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | A unique integer value identifying this General Engagement Survey. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuestionnaireGeneralQuestionnairesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuestionnaireGeneralSurvey**](QuestionnaireGeneralSurvey.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth), [tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

