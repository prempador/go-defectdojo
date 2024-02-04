/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package defectdojo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type QuestionnaireAnsweredQuestionnairesAPI interface {

	/*
	QuestionnaireAnsweredQuestionnairesList Method for QuestionnaireAnsweredQuestionnairesList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQuestionnaireAnsweredQuestionnairesListRequest
	*/
	QuestionnaireAnsweredQuestionnairesList(ctx context.Context) ApiQuestionnaireAnsweredQuestionnairesListRequest

	// QuestionnaireAnsweredQuestionnairesListExecute executes the request
	//  @return PaginatedQuestionnaireAnsweredSurveyList
	QuestionnaireAnsweredQuestionnairesListExecute(r ApiQuestionnaireAnsweredQuestionnairesListRequest) (*PaginatedQuestionnaireAnsweredSurveyList, *http.Response, error)

	/*
	QuestionnaireAnsweredQuestionnairesRetrieve Method for QuestionnaireAnsweredQuestionnairesRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this Answered Engagement Survey.
	@return ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest
	*/
	QuestionnaireAnsweredQuestionnairesRetrieve(ctx context.Context, id int32) ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest

	// QuestionnaireAnsweredQuestionnairesRetrieveExecute executes the request
	//  @return QuestionnaireAnsweredSurvey
	QuestionnaireAnsweredQuestionnairesRetrieveExecute(r ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest) (*QuestionnaireAnsweredSurvey, *http.Response, error)
}

// QuestionnaireAnsweredQuestionnairesAPIService QuestionnaireAnsweredQuestionnairesAPI service
type QuestionnaireAnsweredQuestionnairesAPIService service

type ApiQuestionnaireAnsweredQuestionnairesListRequest struct {
	ctx context.Context
	ApiService QuestionnaireAnsweredQuestionnairesAPI
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiQuestionnaireAnsweredQuestionnairesListRequest) Limit(limit int32) ApiQuestionnaireAnsweredQuestionnairesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiQuestionnaireAnsweredQuestionnairesListRequest) Offset(offset int32) ApiQuestionnaireAnsweredQuestionnairesListRequest {
	r.offset = &offset
	return r
}

func (r ApiQuestionnaireAnsweredQuestionnairesListRequest) Execute() (*PaginatedQuestionnaireAnsweredSurveyList, *http.Response, error) {
	return r.ApiService.QuestionnaireAnsweredQuestionnairesListExecute(r)
}

/*
QuestionnaireAnsweredQuestionnairesList Method for QuestionnaireAnsweredQuestionnairesList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQuestionnaireAnsweredQuestionnairesListRequest
*/
func (a *QuestionnaireAnsweredQuestionnairesAPIService) QuestionnaireAnsweredQuestionnairesList(ctx context.Context) ApiQuestionnaireAnsweredQuestionnairesListRequest {
	return ApiQuestionnaireAnsweredQuestionnairesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedQuestionnaireAnsweredSurveyList
func (a *QuestionnaireAnsweredQuestionnairesAPIService) QuestionnaireAnsweredQuestionnairesListExecute(r ApiQuestionnaireAnsweredQuestionnairesListRequest) (*PaginatedQuestionnaireAnsweredSurveyList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedQuestionnaireAnsweredSurveyList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestionnaireAnsweredQuestionnairesAPIService.QuestionnaireAnsweredQuestionnairesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/questionnaire_answered_questionnaires/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest struct {
	ctx context.Context
	ApiService QuestionnaireAnsweredQuestionnairesAPI
	id int32
}

func (r ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest) Execute() (*QuestionnaireAnsweredSurvey, *http.Response, error) {
	return r.ApiService.QuestionnaireAnsweredQuestionnairesRetrieveExecute(r)
}

/*
QuestionnaireAnsweredQuestionnairesRetrieve Method for QuestionnaireAnsweredQuestionnairesRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this Answered Engagement Survey.
 @return ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest
*/
func (a *QuestionnaireAnsweredQuestionnairesAPIService) QuestionnaireAnsweredQuestionnairesRetrieve(ctx context.Context, id int32) ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest {
	return ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return QuestionnaireAnsweredSurvey
func (a *QuestionnaireAnsweredQuestionnairesAPIService) QuestionnaireAnsweredQuestionnairesRetrieveExecute(r ApiQuestionnaireAnsweredQuestionnairesRetrieveRequest) (*QuestionnaireAnsweredSurvey, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuestionnaireAnsweredSurvey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestionnaireAnsweredQuestionnairesAPIService.QuestionnaireAnsweredQuestionnairesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/questionnaire_answered_questionnaires/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
