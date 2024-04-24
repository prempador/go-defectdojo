/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.33.5
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


type JiraFindingMappingsAPI interface {

	/*
	JiraFindingMappingsCreate Method for JiraFindingMappingsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiJiraFindingMappingsCreateRequest
	*/
	JiraFindingMappingsCreate(ctx context.Context) ApiJiraFindingMappingsCreateRequest

	// JiraFindingMappingsCreateExecute executes the request
	//  @return JIRAIssue
	JiraFindingMappingsCreateExecute(r ApiJiraFindingMappingsCreateRequest) (*JIRAIssue, *http.Response, error)

	/*
	JiraFindingMappingsDeletePreviewList Method for JiraFindingMappingsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ issue.
	@return ApiJiraFindingMappingsDeletePreviewListRequest
	*/
	JiraFindingMappingsDeletePreviewList(ctx context.Context, id int32) ApiJiraFindingMappingsDeletePreviewListRequest

	// JiraFindingMappingsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	JiraFindingMappingsDeletePreviewListExecute(r ApiJiraFindingMappingsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	JiraFindingMappingsDestroy Method for JiraFindingMappingsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ issue.
	@return ApiJiraFindingMappingsDestroyRequest
	*/
	JiraFindingMappingsDestroy(ctx context.Context, id int32) ApiJiraFindingMappingsDestroyRequest

	// JiraFindingMappingsDestroyExecute executes the request
	JiraFindingMappingsDestroyExecute(r ApiJiraFindingMappingsDestroyRequest) (*http.Response, error)

	/*
	JiraFindingMappingsList Method for JiraFindingMappingsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiJiraFindingMappingsListRequest
	*/
	JiraFindingMappingsList(ctx context.Context) ApiJiraFindingMappingsListRequest

	// JiraFindingMappingsListExecute executes the request
	//  @return PaginatedJIRAIssueList
	JiraFindingMappingsListExecute(r ApiJiraFindingMappingsListRequest) (*PaginatedJIRAIssueList, *http.Response, error)

	/*
	JiraFindingMappingsPartialUpdate Method for JiraFindingMappingsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ issue.
	@return ApiJiraFindingMappingsPartialUpdateRequest
	*/
	JiraFindingMappingsPartialUpdate(ctx context.Context, id int32) ApiJiraFindingMappingsPartialUpdateRequest

	// JiraFindingMappingsPartialUpdateExecute executes the request
	//  @return JIRAIssue
	JiraFindingMappingsPartialUpdateExecute(r ApiJiraFindingMappingsPartialUpdateRequest) (*JIRAIssue, *http.Response, error)

	/*
	JiraFindingMappingsRetrieve Method for JiraFindingMappingsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ issue.
	@return ApiJiraFindingMappingsRetrieveRequest
	*/
	JiraFindingMappingsRetrieve(ctx context.Context, id int32) ApiJiraFindingMappingsRetrieveRequest

	// JiraFindingMappingsRetrieveExecute executes the request
	//  @return JIRAIssue
	JiraFindingMappingsRetrieveExecute(r ApiJiraFindingMappingsRetrieveRequest) (*JIRAIssue, *http.Response, error)

	/*
	JiraFindingMappingsUpdate Method for JiraFindingMappingsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ issue.
	@return ApiJiraFindingMappingsUpdateRequest
	*/
	JiraFindingMappingsUpdate(ctx context.Context, id int32) ApiJiraFindingMappingsUpdateRequest

	// JiraFindingMappingsUpdateExecute executes the request
	//  @return JIRAIssue
	JiraFindingMappingsUpdateExecute(r ApiJiraFindingMappingsUpdateRequest) (*JIRAIssue, *http.Response, error)
}

// JiraFindingMappingsAPIService JiraFindingMappingsAPI service
type JiraFindingMappingsAPIService service

type ApiJiraFindingMappingsCreateRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	jIRAIssueRequest *JIRAIssueRequest
}

func (r ApiJiraFindingMappingsCreateRequest) JIRAIssueRequest(jIRAIssueRequest JIRAIssueRequest) ApiJiraFindingMappingsCreateRequest {
	r.jIRAIssueRequest = &jIRAIssueRequest
	return r
}

func (r ApiJiraFindingMappingsCreateRequest) Execute() (*JIRAIssue, *http.Response, error) {
	return r.ApiService.JiraFindingMappingsCreateExecute(r)
}

/*
JiraFindingMappingsCreate Method for JiraFindingMappingsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiJiraFindingMappingsCreateRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsCreate(ctx context.Context) ApiJiraFindingMappingsCreateRequest {
	return ApiJiraFindingMappingsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JIRAIssue
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsCreateExecute(r ApiJiraFindingMappingsCreateRequest) (*JIRAIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.jIRAIssueRequest == nil {
		return localVarReturnValue, nil, reportError("jIRAIssueRequest is required and must be specified")
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
	localVarPostBody = r.jIRAIssueRequest
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

type ApiJiraFindingMappingsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiJiraFindingMappingsDeletePreviewListRequest) Limit(limit int32) ApiJiraFindingMappingsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiJiraFindingMappingsDeletePreviewListRequest) Offset(offset int32) ApiJiraFindingMappingsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiJiraFindingMappingsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.JiraFindingMappingsDeletePreviewListExecute(r)
}

/*
JiraFindingMappingsDeletePreviewList Method for JiraFindingMappingsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ issue.
 @return ApiJiraFindingMappingsDeletePreviewListRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsDeletePreviewList(ctx context.Context, id int32) ApiJiraFindingMappingsDeletePreviewListRequest {
	return ApiJiraFindingMappingsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsDeletePreviewListExecute(r ApiJiraFindingMappingsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/{id}/delete_preview/"
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

type ApiJiraFindingMappingsDestroyRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	id int32
}

func (r ApiJiraFindingMappingsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.JiraFindingMappingsDestroyExecute(r)
}

/*
JiraFindingMappingsDestroy Method for JiraFindingMappingsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ issue.
 @return ApiJiraFindingMappingsDestroyRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsDestroy(ctx context.Context, id int32) ApiJiraFindingMappingsDestroyRequest {
	return ApiJiraFindingMappingsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsDestroyExecute(r ApiJiraFindingMappingsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/{id}/"
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

type ApiJiraFindingMappingsListRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	engagement *int32
	finding *int32
	findingGroup *int32
	id *int32
	jiraId *string
	jiraKey *string
	limit *int32
	offset *int32
}

func (r ApiJiraFindingMappingsListRequest) Engagement(engagement int32) ApiJiraFindingMappingsListRequest {
	r.engagement = &engagement
	return r
}

func (r ApiJiraFindingMappingsListRequest) Finding(finding int32) ApiJiraFindingMappingsListRequest {
	r.finding = &finding
	return r
}

func (r ApiJiraFindingMappingsListRequest) FindingGroup(findingGroup int32) ApiJiraFindingMappingsListRequest {
	r.findingGroup = &findingGroup
	return r
}

func (r ApiJiraFindingMappingsListRequest) Id(id int32) ApiJiraFindingMappingsListRequest {
	r.id = &id
	return r
}

func (r ApiJiraFindingMappingsListRequest) JiraId(jiraId string) ApiJiraFindingMappingsListRequest {
	r.jiraId = &jiraId
	return r
}

func (r ApiJiraFindingMappingsListRequest) JiraKey(jiraKey string) ApiJiraFindingMappingsListRequest {
	r.jiraKey = &jiraKey
	return r
}

// Number of results to return per page.
func (r ApiJiraFindingMappingsListRequest) Limit(limit int32) ApiJiraFindingMappingsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiJiraFindingMappingsListRequest) Offset(offset int32) ApiJiraFindingMappingsListRequest {
	r.offset = &offset
	return r
}

func (r ApiJiraFindingMappingsListRequest) Execute() (*PaginatedJIRAIssueList, *http.Response, error) {
	return r.ApiService.JiraFindingMappingsListExecute(r)
}

/*
JiraFindingMappingsList Method for JiraFindingMappingsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiJiraFindingMappingsListRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsList(ctx context.Context) ApiJiraFindingMappingsListRequest {
	return ApiJiraFindingMappingsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedJIRAIssueList
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsListExecute(r ApiJiraFindingMappingsListRequest) (*PaginatedJIRAIssueList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedJIRAIssueList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.engagement != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "engagement", r.engagement, "")
	}
	if r.finding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "finding", r.finding, "")
	}
	if r.findingGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "finding_group", r.findingGroup, "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.jiraId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jira_id", r.jiraId, "")
	}
	if r.jiraKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jira_key", r.jiraKey, "")
	}
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

type ApiJiraFindingMappingsPartialUpdateRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	id int32
	patchedJIRAIssueRequest *PatchedJIRAIssueRequest
}

func (r ApiJiraFindingMappingsPartialUpdateRequest) PatchedJIRAIssueRequest(patchedJIRAIssueRequest PatchedJIRAIssueRequest) ApiJiraFindingMappingsPartialUpdateRequest {
	r.patchedJIRAIssueRequest = &patchedJIRAIssueRequest
	return r
}

func (r ApiJiraFindingMappingsPartialUpdateRequest) Execute() (*JIRAIssue, *http.Response, error) {
	return r.ApiService.JiraFindingMappingsPartialUpdateExecute(r)
}

/*
JiraFindingMappingsPartialUpdate Method for JiraFindingMappingsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ issue.
 @return ApiJiraFindingMappingsPartialUpdateRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsPartialUpdate(ctx context.Context, id int32) ApiJiraFindingMappingsPartialUpdateRequest {
	return ApiJiraFindingMappingsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return JIRAIssue
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsPartialUpdateExecute(r ApiJiraFindingMappingsPartialUpdateRequest) (*JIRAIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/{id}/"
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
	localVarPostBody = r.patchedJIRAIssueRequest
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

type ApiJiraFindingMappingsRetrieveRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	id int32
}

func (r ApiJiraFindingMappingsRetrieveRequest) Execute() (*JIRAIssue, *http.Response, error) {
	return r.ApiService.JiraFindingMappingsRetrieveExecute(r)
}

/*
JiraFindingMappingsRetrieve Method for JiraFindingMappingsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ issue.
 @return ApiJiraFindingMappingsRetrieveRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsRetrieve(ctx context.Context, id int32) ApiJiraFindingMappingsRetrieveRequest {
	return ApiJiraFindingMappingsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return JIRAIssue
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsRetrieveExecute(r ApiJiraFindingMappingsRetrieveRequest) (*JIRAIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/{id}/"
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

type ApiJiraFindingMappingsUpdateRequest struct {
	ctx context.Context
	ApiService JiraFindingMappingsAPI
	id int32
	jIRAIssueRequest *JIRAIssueRequest
}

func (r ApiJiraFindingMappingsUpdateRequest) JIRAIssueRequest(jIRAIssueRequest JIRAIssueRequest) ApiJiraFindingMappingsUpdateRequest {
	r.jIRAIssueRequest = &jIRAIssueRequest
	return r
}

func (r ApiJiraFindingMappingsUpdateRequest) Execute() (*JIRAIssue, *http.Response, error) {
	return r.ApiService.JiraFindingMappingsUpdateExecute(r)
}

/*
JiraFindingMappingsUpdate Method for JiraFindingMappingsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ issue.
 @return ApiJiraFindingMappingsUpdateRequest
*/
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsUpdate(ctx context.Context, id int32) ApiJiraFindingMappingsUpdateRequest {
	return ApiJiraFindingMappingsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return JIRAIssue
func (a *JiraFindingMappingsAPIService) JiraFindingMappingsUpdateExecute(r ApiJiraFindingMappingsUpdateRequest) (*JIRAIssue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAIssue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraFindingMappingsAPIService.JiraFindingMappingsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_finding_mappings/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.jIRAIssueRequest == nil {
		return localVarReturnValue, nil, reportError("jIRAIssueRequest is required and must be specified")
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
	localVarPostBody = r.jIRAIssueRequest
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
