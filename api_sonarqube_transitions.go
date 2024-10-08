/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.38.4
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


type SonarqubeTransitionsAPI interface {

	/*
	SonarqubeTransitionsCreate Method for SonarqubeTransitionsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSonarqubeTransitionsCreateRequest
	*/
	SonarqubeTransitionsCreate(ctx context.Context) ApiSonarqubeTransitionsCreateRequest

	// SonarqubeTransitionsCreateExecute executes the request
	//  @return SonarqubeIssueTransition
	SonarqubeTransitionsCreateExecute(r ApiSonarqubeTransitionsCreateRequest) (*SonarqubeIssueTransition, *http.Response, error)

	/*
	SonarqubeTransitionsDeletePreviewList Method for SonarqubeTransitionsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue_ transition.
	@return ApiSonarqubeTransitionsDeletePreviewListRequest
	*/
	SonarqubeTransitionsDeletePreviewList(ctx context.Context, id int32) ApiSonarqubeTransitionsDeletePreviewListRequest

	// SonarqubeTransitionsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	SonarqubeTransitionsDeletePreviewListExecute(r ApiSonarqubeTransitionsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	SonarqubeTransitionsDestroy Method for SonarqubeTransitionsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue_ transition.
	@return ApiSonarqubeTransitionsDestroyRequest
	*/
	SonarqubeTransitionsDestroy(ctx context.Context, id int32) ApiSonarqubeTransitionsDestroyRequest

	// SonarqubeTransitionsDestroyExecute executes the request
	SonarqubeTransitionsDestroyExecute(r ApiSonarqubeTransitionsDestroyRequest) (*http.Response, error)

	/*
	SonarqubeTransitionsList Method for SonarqubeTransitionsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSonarqubeTransitionsListRequest
	*/
	SonarqubeTransitionsList(ctx context.Context) ApiSonarqubeTransitionsListRequest

	// SonarqubeTransitionsListExecute executes the request
	//  @return PaginatedSonarqubeIssueTransitionList
	SonarqubeTransitionsListExecute(r ApiSonarqubeTransitionsListRequest) (*PaginatedSonarqubeIssueTransitionList, *http.Response, error)

	/*
	SonarqubeTransitionsPartialUpdate Method for SonarqubeTransitionsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue_ transition.
	@return ApiSonarqubeTransitionsPartialUpdateRequest
	*/
	SonarqubeTransitionsPartialUpdate(ctx context.Context, id int32) ApiSonarqubeTransitionsPartialUpdateRequest

	// SonarqubeTransitionsPartialUpdateExecute executes the request
	//  @return SonarqubeIssueTransition
	SonarqubeTransitionsPartialUpdateExecute(r ApiSonarqubeTransitionsPartialUpdateRequest) (*SonarqubeIssueTransition, *http.Response, error)

	/*
	SonarqubeTransitionsRetrieve Method for SonarqubeTransitionsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue_ transition.
	@return ApiSonarqubeTransitionsRetrieveRequest
	*/
	SonarqubeTransitionsRetrieve(ctx context.Context, id int32) ApiSonarqubeTransitionsRetrieveRequest

	// SonarqubeTransitionsRetrieveExecute executes the request
	//  @return SonarqubeIssueTransition
	SonarqubeTransitionsRetrieveExecute(r ApiSonarqubeTransitionsRetrieveRequest) (*SonarqubeIssueTransition, *http.Response, error)

	/*
	SonarqubeTransitionsUpdate Method for SonarqubeTransitionsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue_ transition.
	@return ApiSonarqubeTransitionsUpdateRequest
	*/
	SonarqubeTransitionsUpdate(ctx context.Context, id int32) ApiSonarqubeTransitionsUpdateRequest

	// SonarqubeTransitionsUpdateExecute executes the request
	//  @return SonarqubeIssueTransition
	SonarqubeTransitionsUpdateExecute(r ApiSonarqubeTransitionsUpdateRequest) (*SonarqubeIssueTransition, *http.Response, error)
}

// SonarqubeTransitionsAPIService SonarqubeTransitionsAPI service
type SonarqubeTransitionsAPIService service

type ApiSonarqubeTransitionsCreateRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	sonarqubeIssueTransitionRequest *SonarqubeIssueTransitionRequest
}

func (r ApiSonarqubeTransitionsCreateRequest) SonarqubeIssueTransitionRequest(sonarqubeIssueTransitionRequest SonarqubeIssueTransitionRequest) ApiSonarqubeTransitionsCreateRequest {
	r.sonarqubeIssueTransitionRequest = &sonarqubeIssueTransitionRequest
	return r
}

func (r ApiSonarqubeTransitionsCreateRequest) Execute() (*SonarqubeIssueTransition, *http.Response, error) {
	return r.ApiService.SonarqubeTransitionsCreateExecute(r)
}

/*
SonarqubeTransitionsCreate Method for SonarqubeTransitionsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSonarqubeTransitionsCreateRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsCreate(ctx context.Context) ApiSonarqubeTransitionsCreateRequest {
	return ApiSonarqubeTransitionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SonarqubeIssueTransition
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsCreateExecute(r ApiSonarqubeTransitionsCreateRequest) (*SonarqubeIssueTransition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssueTransition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sonarqubeIssueTransitionRequest == nil {
		return localVarReturnValue, nil, reportError("sonarqubeIssueTransitionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.sonarqubeIssueTransitionRequest
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

type ApiSonarqubeTransitionsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiSonarqubeTransitionsDeletePreviewListRequest) Limit(limit int32) ApiSonarqubeTransitionsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiSonarqubeTransitionsDeletePreviewListRequest) Offset(offset int32) ApiSonarqubeTransitionsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiSonarqubeTransitionsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.SonarqubeTransitionsDeletePreviewListExecute(r)
}

/*
SonarqubeTransitionsDeletePreviewList Method for SonarqubeTransitionsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue_ transition.
 @return ApiSonarqubeTransitionsDeletePreviewListRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsDeletePreviewList(ctx context.Context, id int32) ApiSonarqubeTransitionsDeletePreviewListRequest {
	return ApiSonarqubeTransitionsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsDeletePreviewListExecute(r ApiSonarqubeTransitionsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/{id}/delete_preview/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
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

type ApiSonarqubeTransitionsDestroyRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	id int32
}

func (r ApiSonarqubeTransitionsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.SonarqubeTransitionsDestroyExecute(r)
}

/*
SonarqubeTransitionsDestroy Method for SonarqubeTransitionsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue_ transition.
 @return ApiSonarqubeTransitionsDestroyRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsDestroy(ctx context.Context, id int32) ApiSonarqubeTransitionsDestroyRequest {
	return ApiSonarqubeTransitionsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsDestroyExecute(r ApiSonarqubeTransitionsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/{id}/"
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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSonarqubeTransitionsListRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	findingStatus *string
	id *int32
	limit *int32
	offset *int32
	sonarqubeIssue *int32
	sonarqubeStatus *string
	transitions *string
}

func (r ApiSonarqubeTransitionsListRequest) FindingStatus(findingStatus string) ApiSonarqubeTransitionsListRequest {
	r.findingStatus = &findingStatus
	return r
}

func (r ApiSonarqubeTransitionsListRequest) Id(id int32) ApiSonarqubeTransitionsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiSonarqubeTransitionsListRequest) Limit(limit int32) ApiSonarqubeTransitionsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiSonarqubeTransitionsListRequest) Offset(offset int32) ApiSonarqubeTransitionsListRequest {
	r.offset = &offset
	return r
}

func (r ApiSonarqubeTransitionsListRequest) SonarqubeIssue(sonarqubeIssue int32) ApiSonarqubeTransitionsListRequest {
	r.sonarqubeIssue = &sonarqubeIssue
	return r
}

func (r ApiSonarqubeTransitionsListRequest) SonarqubeStatus(sonarqubeStatus string) ApiSonarqubeTransitionsListRequest {
	r.sonarqubeStatus = &sonarqubeStatus
	return r
}

func (r ApiSonarqubeTransitionsListRequest) Transitions(transitions string) ApiSonarqubeTransitionsListRequest {
	r.transitions = &transitions
	return r
}

func (r ApiSonarqubeTransitionsListRequest) Execute() (*PaginatedSonarqubeIssueTransitionList, *http.Response, error) {
	return r.ApiService.SonarqubeTransitionsListExecute(r)
}

/*
SonarqubeTransitionsList Method for SonarqubeTransitionsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSonarqubeTransitionsListRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsList(ctx context.Context) ApiSonarqubeTransitionsListRequest {
	return ApiSonarqubeTransitionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedSonarqubeIssueTransitionList
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsListExecute(r ApiSonarqubeTransitionsListRequest) (*PaginatedSonarqubeIssueTransitionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedSonarqubeIssueTransitionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.findingStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "finding_status", r.findingStatus, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.sonarqubeIssue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sonarqube_issue", r.sonarqubeIssue, "form", "")
	}
	if r.sonarqubeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sonarqube_status", r.sonarqubeStatus, "form", "")
	}
	if r.transitions != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transitions", r.transitions, "form", "")
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

type ApiSonarqubeTransitionsPartialUpdateRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	id int32
	patchedSonarqubeIssueTransitionRequest *PatchedSonarqubeIssueTransitionRequest
}

func (r ApiSonarqubeTransitionsPartialUpdateRequest) PatchedSonarqubeIssueTransitionRequest(patchedSonarqubeIssueTransitionRequest PatchedSonarqubeIssueTransitionRequest) ApiSonarqubeTransitionsPartialUpdateRequest {
	r.patchedSonarqubeIssueTransitionRequest = &patchedSonarqubeIssueTransitionRequest
	return r
}

func (r ApiSonarqubeTransitionsPartialUpdateRequest) Execute() (*SonarqubeIssueTransition, *http.Response, error) {
	return r.ApiService.SonarqubeTransitionsPartialUpdateExecute(r)
}

/*
SonarqubeTransitionsPartialUpdate Method for SonarqubeTransitionsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue_ transition.
 @return ApiSonarqubeTransitionsPartialUpdateRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsPartialUpdate(ctx context.Context, id int32) ApiSonarqubeTransitionsPartialUpdateRequest {
	return ApiSonarqubeTransitionsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SonarqubeIssueTransition
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsPartialUpdateExecute(r ApiSonarqubeTransitionsPartialUpdateRequest) (*SonarqubeIssueTransition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssueTransition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.patchedSonarqubeIssueTransitionRequest
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

type ApiSonarqubeTransitionsRetrieveRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	id int32
}

func (r ApiSonarqubeTransitionsRetrieveRequest) Execute() (*SonarqubeIssueTransition, *http.Response, error) {
	return r.ApiService.SonarqubeTransitionsRetrieveExecute(r)
}

/*
SonarqubeTransitionsRetrieve Method for SonarqubeTransitionsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue_ transition.
 @return ApiSonarqubeTransitionsRetrieveRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsRetrieve(ctx context.Context, id int32) ApiSonarqubeTransitionsRetrieveRequest {
	return ApiSonarqubeTransitionsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SonarqubeIssueTransition
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsRetrieveExecute(r ApiSonarqubeTransitionsRetrieveRequest) (*SonarqubeIssueTransition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssueTransition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/{id}/"
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

type ApiSonarqubeTransitionsUpdateRequest struct {
	ctx context.Context
	ApiService SonarqubeTransitionsAPI
	id int32
	sonarqubeIssueTransitionRequest *SonarqubeIssueTransitionRequest
}

func (r ApiSonarqubeTransitionsUpdateRequest) SonarqubeIssueTransitionRequest(sonarqubeIssueTransitionRequest SonarqubeIssueTransitionRequest) ApiSonarqubeTransitionsUpdateRequest {
	r.sonarqubeIssueTransitionRequest = &sonarqubeIssueTransitionRequest
	return r
}

func (r ApiSonarqubeTransitionsUpdateRequest) Execute() (*SonarqubeIssueTransition, *http.Response, error) {
	return r.ApiService.SonarqubeTransitionsUpdateExecute(r)
}

/*
SonarqubeTransitionsUpdate Method for SonarqubeTransitionsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue_ transition.
 @return ApiSonarqubeTransitionsUpdateRequest
*/
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsUpdate(ctx context.Context, id int32) ApiSonarqubeTransitionsUpdateRequest {
	return ApiSonarqubeTransitionsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SonarqubeIssueTransition
func (a *SonarqubeTransitionsAPIService) SonarqubeTransitionsUpdateExecute(r ApiSonarqubeTransitionsUpdateRequest) (*SonarqubeIssueTransition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssueTransition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeTransitionsAPIService.SonarqubeTransitionsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_transitions/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sonarqubeIssueTransitionRequest == nil {
		return localVarReturnValue, nil, reportError("sonarqubeIssueTransitionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	// body params
	localVarPostBody = r.sonarqubeIssueTransitionRequest
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
