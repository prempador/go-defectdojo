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


type SonarqubeIssuesAPI interface {

	/*
	SonarqubeIssuesCreate Method for SonarqubeIssuesCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSonarqubeIssuesCreateRequest
	*/
	SonarqubeIssuesCreate(ctx context.Context) ApiSonarqubeIssuesCreateRequest

	// SonarqubeIssuesCreateExecute executes the request
	//  @return SonarqubeIssue
	SonarqubeIssuesCreateExecute(r ApiSonarqubeIssuesCreateRequest) (*SonarqubeIssue, *http.Response, error)

	/*
	SonarqubeIssuesDeletePreviewList Method for SonarqubeIssuesDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue.
	@return ApiSonarqubeIssuesDeletePreviewListRequest
	*/
	SonarqubeIssuesDeletePreviewList(ctx context.Context, id int32) ApiSonarqubeIssuesDeletePreviewListRequest

	// SonarqubeIssuesDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	SonarqubeIssuesDeletePreviewListExecute(r ApiSonarqubeIssuesDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	SonarqubeIssuesDestroy Method for SonarqubeIssuesDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue.
	@return ApiSonarqubeIssuesDestroyRequest
	*/
	SonarqubeIssuesDestroy(ctx context.Context, id int32) ApiSonarqubeIssuesDestroyRequest

	// SonarqubeIssuesDestroyExecute executes the request
	SonarqubeIssuesDestroyExecute(r ApiSonarqubeIssuesDestroyRequest) (*http.Response, error)

	/*
	SonarqubeIssuesList Method for SonarqubeIssuesList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSonarqubeIssuesListRequest
	*/
	SonarqubeIssuesList(ctx context.Context) ApiSonarqubeIssuesListRequest

	// SonarqubeIssuesListExecute executes the request
	//  @return PaginatedSonarqubeIssueList
	SonarqubeIssuesListExecute(r ApiSonarqubeIssuesListRequest) (*PaginatedSonarqubeIssueList, *http.Response, error)

	/*
	SonarqubeIssuesPartialUpdate Method for SonarqubeIssuesPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue.
	@return ApiSonarqubeIssuesPartialUpdateRequest
	*/
	SonarqubeIssuesPartialUpdate(ctx context.Context, id int32) ApiSonarqubeIssuesPartialUpdateRequest

	// SonarqubeIssuesPartialUpdateExecute executes the request
	//  @return SonarqubeIssue
	SonarqubeIssuesPartialUpdateExecute(r ApiSonarqubeIssuesPartialUpdateRequest) (*SonarqubeIssue, *http.Response, error)

	/*
	SonarqubeIssuesRetrieve Method for SonarqubeIssuesRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue.
	@return ApiSonarqubeIssuesRetrieveRequest
	*/
	SonarqubeIssuesRetrieve(ctx context.Context, id int32) ApiSonarqubeIssuesRetrieveRequest

	// SonarqubeIssuesRetrieveExecute executes the request
	//  @return SonarqubeIssue
	SonarqubeIssuesRetrieveExecute(r ApiSonarqubeIssuesRetrieveRequest) (*SonarqubeIssue, *http.Response, error)

	/*
	SonarqubeIssuesUpdate Method for SonarqubeIssuesUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sonarqube_ issue.
	@return ApiSonarqubeIssuesUpdateRequest
	*/
	SonarqubeIssuesUpdate(ctx context.Context, id int32) ApiSonarqubeIssuesUpdateRequest

	// SonarqubeIssuesUpdateExecute executes the request
	//  @return SonarqubeIssue
	SonarqubeIssuesUpdateExecute(r ApiSonarqubeIssuesUpdateRequest) (*SonarqubeIssue, *http.Response, error)
}

// SonarqubeIssuesAPIService SonarqubeIssuesAPI service
type SonarqubeIssuesAPIService service

type ApiSonarqubeIssuesCreateRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	sonarqubeIssueRequest *SonarqubeIssueRequest
}

func (r ApiSonarqubeIssuesCreateRequest) SonarqubeIssueRequest(sonarqubeIssueRequest SonarqubeIssueRequest) ApiSonarqubeIssuesCreateRequest {
	r.sonarqubeIssueRequest = &sonarqubeIssueRequest
	return r
}

func (r ApiSonarqubeIssuesCreateRequest) Execute() (*SonarqubeIssue, *http.Response, error) {
	return r.ApiService.SonarqubeIssuesCreateExecute(r)
}

/*
SonarqubeIssuesCreate Method for SonarqubeIssuesCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSonarqubeIssuesCreateRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesCreate(ctx context.Context) ApiSonarqubeIssuesCreateRequest {
	return ApiSonarqubeIssuesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SonarqubeIssue
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesCreateExecute(r ApiSonarqubeIssuesCreateRequest) (*SonarqubeIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sonarqubeIssueRequest == nil {
		return localVarReturnValue, nil, reportError("sonarqubeIssueRequest is required and must be specified")
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
	localVarPostBody = r.sonarqubeIssueRequest
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

type ApiSonarqubeIssuesDeletePreviewListRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiSonarqubeIssuesDeletePreviewListRequest) Limit(limit int32) ApiSonarqubeIssuesDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiSonarqubeIssuesDeletePreviewListRequest) Offset(offset int32) ApiSonarqubeIssuesDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiSonarqubeIssuesDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.SonarqubeIssuesDeletePreviewListExecute(r)
}

/*
SonarqubeIssuesDeletePreviewList Method for SonarqubeIssuesDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue.
 @return ApiSonarqubeIssuesDeletePreviewListRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesDeletePreviewList(ctx context.Context, id int32) ApiSonarqubeIssuesDeletePreviewListRequest {
	return ApiSonarqubeIssuesDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesDeletePreviewListExecute(r ApiSonarqubeIssuesDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/{id}/delete_preview/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiSonarqubeIssuesDestroyRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	id int32
}

func (r ApiSonarqubeIssuesDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.SonarqubeIssuesDestroyExecute(r)
}

/*
SonarqubeIssuesDestroy Method for SonarqubeIssuesDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue.
 @return ApiSonarqubeIssuesDestroyRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesDestroy(ctx context.Context, id int32) ApiSonarqubeIssuesDestroyRequest {
	return ApiSonarqubeIssuesDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesDestroyExecute(r ApiSonarqubeIssuesDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/{id}/"
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

type ApiSonarqubeIssuesListRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	id *int32
	key *string
	limit *int32
	offset *int32
	status *string
	type_ *string
}

func (r ApiSonarqubeIssuesListRequest) Id(id int32) ApiSonarqubeIssuesListRequest {
	r.id = &id
	return r
}

func (r ApiSonarqubeIssuesListRequest) Key(key string) ApiSonarqubeIssuesListRequest {
	r.key = &key
	return r
}

// Number of results to return per page.
func (r ApiSonarqubeIssuesListRequest) Limit(limit int32) ApiSonarqubeIssuesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiSonarqubeIssuesListRequest) Offset(offset int32) ApiSonarqubeIssuesListRequest {
	r.offset = &offset
	return r
}

func (r ApiSonarqubeIssuesListRequest) Status(status string) ApiSonarqubeIssuesListRequest {
	r.status = &status
	return r
}

func (r ApiSonarqubeIssuesListRequest) Type_(type_ string) ApiSonarqubeIssuesListRequest {
	r.type_ = &type_
	return r
}

func (r ApiSonarqubeIssuesListRequest) Execute() (*PaginatedSonarqubeIssueList, *http.Response, error) {
	return r.ApiService.SonarqubeIssuesListExecute(r)
}

/*
SonarqubeIssuesList Method for SonarqubeIssuesList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSonarqubeIssuesListRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesList(ctx context.Context) ApiSonarqubeIssuesListRequest {
	return ApiSonarqubeIssuesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedSonarqubeIssueList
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesListExecute(r ApiSonarqubeIssuesListRequest) (*PaginatedSonarqubeIssueList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedSonarqubeIssueList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.key != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "key", r.key, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
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

type ApiSonarqubeIssuesPartialUpdateRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	id int32
	patchedSonarqubeIssueRequest *PatchedSonarqubeIssueRequest
}

func (r ApiSonarqubeIssuesPartialUpdateRequest) PatchedSonarqubeIssueRequest(patchedSonarqubeIssueRequest PatchedSonarqubeIssueRequest) ApiSonarqubeIssuesPartialUpdateRequest {
	r.patchedSonarqubeIssueRequest = &patchedSonarqubeIssueRequest
	return r
}

func (r ApiSonarqubeIssuesPartialUpdateRequest) Execute() (*SonarqubeIssue, *http.Response, error) {
	return r.ApiService.SonarqubeIssuesPartialUpdateExecute(r)
}

/*
SonarqubeIssuesPartialUpdate Method for SonarqubeIssuesPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue.
 @return ApiSonarqubeIssuesPartialUpdateRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesPartialUpdate(ctx context.Context, id int32) ApiSonarqubeIssuesPartialUpdateRequest {
	return ApiSonarqubeIssuesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SonarqubeIssue
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesPartialUpdateExecute(r ApiSonarqubeIssuesPartialUpdateRequest) (*SonarqubeIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/{id}/"
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
	localVarPostBody = r.patchedSonarqubeIssueRequest
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

type ApiSonarqubeIssuesRetrieveRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	id int32
}

func (r ApiSonarqubeIssuesRetrieveRequest) Execute() (*SonarqubeIssue, *http.Response, error) {
	return r.ApiService.SonarqubeIssuesRetrieveExecute(r)
}

/*
SonarqubeIssuesRetrieve Method for SonarqubeIssuesRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue.
 @return ApiSonarqubeIssuesRetrieveRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesRetrieve(ctx context.Context, id int32) ApiSonarqubeIssuesRetrieveRequest {
	return ApiSonarqubeIssuesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SonarqubeIssue
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesRetrieveExecute(r ApiSonarqubeIssuesRetrieveRequest) (*SonarqubeIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/{id}/"
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

type ApiSonarqubeIssuesUpdateRequest struct {
	ctx context.Context
	ApiService SonarqubeIssuesAPI
	id int32
	sonarqubeIssueRequest *SonarqubeIssueRequest
}

func (r ApiSonarqubeIssuesUpdateRequest) SonarqubeIssueRequest(sonarqubeIssueRequest SonarqubeIssueRequest) ApiSonarqubeIssuesUpdateRequest {
	r.sonarqubeIssueRequest = &sonarqubeIssueRequest
	return r
}

func (r ApiSonarqubeIssuesUpdateRequest) Execute() (*SonarqubeIssue, *http.Response, error) {
	return r.ApiService.SonarqubeIssuesUpdateExecute(r)
}

/*
SonarqubeIssuesUpdate Method for SonarqubeIssuesUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sonarqube_ issue.
 @return ApiSonarqubeIssuesUpdateRequest
*/
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesUpdate(ctx context.Context, id int32) ApiSonarqubeIssuesUpdateRequest {
	return ApiSonarqubeIssuesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SonarqubeIssue
func (a *SonarqubeIssuesAPIService) SonarqubeIssuesUpdateExecute(r ApiSonarqubeIssuesUpdateRequest) (*SonarqubeIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SonarqubeIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SonarqubeIssuesAPIService.SonarqubeIssuesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sonarqube_issues/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sonarqubeIssueRequest == nil {
		return localVarReturnValue, nil, reportError("sonarqubeIssueRequest is required and must be specified")
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
	localVarPostBody = r.sonarqubeIssueRequest
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