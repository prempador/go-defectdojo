/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.30.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type QuestionnaireQuestionsAPI interface {

	/*
	QuestionnaireQuestionsList Method for QuestionnaireQuestionsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiQuestionnaireQuestionsListRequest
	*/
	QuestionnaireQuestionsList(ctx context.Context) ApiQuestionnaireQuestionsListRequest

	// QuestionnaireQuestionsListExecute executes the request
	//  @return PaginatedQuestionnaireQuestionList
	QuestionnaireQuestionsListExecute(r ApiQuestionnaireQuestionsListRequest) (*PaginatedQuestionnaireQuestionList, *http.Response, error)

	/*
	QuestionnaireQuestionsRetrieve Method for QuestionnaireQuestionsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this question.
	@return ApiQuestionnaireQuestionsRetrieveRequest
	*/
	QuestionnaireQuestionsRetrieve(ctx context.Context, id int32) ApiQuestionnaireQuestionsRetrieveRequest

	// QuestionnaireQuestionsRetrieveExecute executes the request
	//  @return QuestionnaireQuestion
	QuestionnaireQuestionsRetrieveExecute(r ApiQuestionnaireQuestionsRetrieveRequest) (*QuestionnaireQuestion, *http.Response, error)
}

// QuestionnaireQuestionsAPIService QuestionnaireQuestionsAPI service
type QuestionnaireQuestionsAPIService service

type ApiQuestionnaireQuestionsListRequest struct {
	ctx context.Context
	ApiService QuestionnaireQuestionsAPI
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiQuestionnaireQuestionsListRequest) Limit(limit int32) ApiQuestionnaireQuestionsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiQuestionnaireQuestionsListRequest) Offset(offset int32) ApiQuestionnaireQuestionsListRequest {
	r.offset = &offset
	return r
}

func (r ApiQuestionnaireQuestionsListRequest) Execute() (*PaginatedQuestionnaireQuestionList, *http.Response, error) {
	return r.ApiService.QuestionnaireQuestionsListExecute(r)
}

/*
QuestionnaireQuestionsList Method for QuestionnaireQuestionsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQuestionnaireQuestionsListRequest
*/
func (a *QuestionnaireQuestionsAPIService) QuestionnaireQuestionsList(ctx context.Context) ApiQuestionnaireQuestionsListRequest {
	return ApiQuestionnaireQuestionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedQuestionnaireQuestionList
func (a *QuestionnaireQuestionsAPIService) QuestionnaireQuestionsListExecute(r ApiQuestionnaireQuestionsListRequest) (*PaginatedQuestionnaireQuestionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedQuestionnaireQuestionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestionnaireQuestionsAPIService.QuestionnaireQuestionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/questionnaire_questions/"

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

type ApiQuestionnaireQuestionsRetrieveRequest struct {
	ctx context.Context
	ApiService QuestionnaireQuestionsAPI
	id int32
}

func (r ApiQuestionnaireQuestionsRetrieveRequest) Execute() (*QuestionnaireQuestion, *http.Response, error) {
	return r.ApiService.QuestionnaireQuestionsRetrieveExecute(r)
}

/*
QuestionnaireQuestionsRetrieve Method for QuestionnaireQuestionsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this question.
 @return ApiQuestionnaireQuestionsRetrieveRequest
*/
func (a *QuestionnaireQuestionsAPIService) QuestionnaireQuestionsRetrieve(ctx context.Context, id int32) ApiQuestionnaireQuestionsRetrieveRequest {
	return ApiQuestionnaireQuestionsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return QuestionnaireQuestion
func (a *QuestionnaireQuestionsAPIService) QuestionnaireQuestionsRetrieveExecute(r ApiQuestionnaireQuestionsRetrieveRequest) (*QuestionnaireQuestion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuestionnaireQuestion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuestionnaireQuestionsAPIService.QuestionnaireQuestionsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/questionnaire_questions/{id}/"
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
