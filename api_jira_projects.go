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


type JiraProjectsAPI interface {

	/*
	JiraProjectsCreate Method for JiraProjectsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiJiraProjectsCreateRequest
	*/
	JiraProjectsCreate(ctx context.Context) ApiJiraProjectsCreateRequest

	// JiraProjectsCreateExecute executes the request
	//  @return JIRAProject
	JiraProjectsCreateExecute(r ApiJiraProjectsCreateRequest) (*JIRAProject, *http.Response, error)

	/*
	JiraProjectsDeletePreviewList Method for JiraProjectsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ project.
	@return ApiJiraProjectsDeletePreviewListRequest
	*/
	JiraProjectsDeletePreviewList(ctx context.Context, id int32) ApiJiraProjectsDeletePreviewListRequest

	// JiraProjectsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	JiraProjectsDeletePreviewListExecute(r ApiJiraProjectsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	JiraProjectsDestroy Method for JiraProjectsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ project.
	@return ApiJiraProjectsDestroyRequest
	*/
	JiraProjectsDestroy(ctx context.Context, id int32) ApiJiraProjectsDestroyRequest

	// JiraProjectsDestroyExecute executes the request
	JiraProjectsDestroyExecute(r ApiJiraProjectsDestroyRequest) (*http.Response, error)

	/*
	JiraProjectsList Method for JiraProjectsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiJiraProjectsListRequest
	*/
	JiraProjectsList(ctx context.Context) ApiJiraProjectsListRequest

	// JiraProjectsListExecute executes the request
	//  @return PaginatedJIRAProjectList
	JiraProjectsListExecute(r ApiJiraProjectsListRequest) (*PaginatedJIRAProjectList, *http.Response, error)

	/*
	JiraProjectsPartialUpdate Method for JiraProjectsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ project.
	@return ApiJiraProjectsPartialUpdateRequest
	*/
	JiraProjectsPartialUpdate(ctx context.Context, id int32) ApiJiraProjectsPartialUpdateRequest

	// JiraProjectsPartialUpdateExecute executes the request
	//  @return JIRAProject
	JiraProjectsPartialUpdateExecute(r ApiJiraProjectsPartialUpdateRequest) (*JIRAProject, *http.Response, error)

	/*
	JiraProjectsRetrieve Method for JiraProjectsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ project.
	@return ApiJiraProjectsRetrieveRequest
	*/
	JiraProjectsRetrieve(ctx context.Context, id int32) ApiJiraProjectsRetrieveRequest

	// JiraProjectsRetrieveExecute executes the request
	//  @return JIRAProject
	JiraProjectsRetrieveExecute(r ApiJiraProjectsRetrieveRequest) (*JIRAProject, *http.Response, error)

	/*
	JiraProjectsUpdate Method for JiraProjectsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this jir a_ project.
	@return ApiJiraProjectsUpdateRequest
	*/
	JiraProjectsUpdate(ctx context.Context, id int32) ApiJiraProjectsUpdateRequest

	// JiraProjectsUpdateExecute executes the request
	//  @return JIRAProject
	JiraProjectsUpdateExecute(r ApiJiraProjectsUpdateRequest) (*JIRAProject, *http.Response, error)
}

// JiraProjectsAPIService JiraProjectsAPI service
type JiraProjectsAPIService service

type ApiJiraProjectsCreateRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	jIRAProjectRequest *JIRAProjectRequest
}

func (r ApiJiraProjectsCreateRequest) JIRAProjectRequest(jIRAProjectRequest JIRAProjectRequest) ApiJiraProjectsCreateRequest {
	r.jIRAProjectRequest = &jIRAProjectRequest
	return r
}

func (r ApiJiraProjectsCreateRequest) Execute() (*JIRAProject, *http.Response, error) {
	return r.ApiService.JiraProjectsCreateExecute(r)
}

/*
JiraProjectsCreate Method for JiraProjectsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiJiraProjectsCreateRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsCreate(ctx context.Context) ApiJiraProjectsCreateRequest {
	return ApiJiraProjectsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return JIRAProject
func (a *JiraProjectsAPIService) JiraProjectsCreateExecute(r ApiJiraProjectsCreateRequest) (*JIRAProject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/"

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
	localVarPostBody = r.jIRAProjectRequest
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

type ApiJiraProjectsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiJiraProjectsDeletePreviewListRequest) Limit(limit int32) ApiJiraProjectsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiJiraProjectsDeletePreviewListRequest) Offset(offset int32) ApiJiraProjectsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiJiraProjectsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.JiraProjectsDeletePreviewListExecute(r)
}

/*
JiraProjectsDeletePreviewList Method for JiraProjectsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ project.
 @return ApiJiraProjectsDeletePreviewListRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsDeletePreviewList(ctx context.Context, id int32) ApiJiraProjectsDeletePreviewListRequest {
	return ApiJiraProjectsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *JiraProjectsAPIService) JiraProjectsDeletePreviewListExecute(r ApiJiraProjectsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/{id}/delete_preview/"
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

type ApiJiraProjectsDestroyRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	id int32
}

func (r ApiJiraProjectsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.JiraProjectsDestroyExecute(r)
}

/*
JiraProjectsDestroy Method for JiraProjectsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ project.
 @return ApiJiraProjectsDestroyRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsDestroy(ctx context.Context, id int32) ApiJiraProjectsDestroyRequest {
	return ApiJiraProjectsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *JiraProjectsAPIService) JiraProjectsDestroyExecute(r ApiJiraProjectsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/{id}/"
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

type ApiJiraProjectsListRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	component *string
	enableEngagementEpicMapping *bool
	engagement *int32
	id *int32
	jiraInstance *int32
	limit *int32
	offset *int32
	product *int32
	projectKey *string
	pushAllIssues *bool
	pushNotes *bool
}

func (r ApiJiraProjectsListRequest) Component(component string) ApiJiraProjectsListRequest {
	r.component = &component
	return r
}

func (r ApiJiraProjectsListRequest) EnableEngagementEpicMapping(enableEngagementEpicMapping bool) ApiJiraProjectsListRequest {
	r.enableEngagementEpicMapping = &enableEngagementEpicMapping
	return r
}

func (r ApiJiraProjectsListRequest) Engagement(engagement int32) ApiJiraProjectsListRequest {
	r.engagement = &engagement
	return r
}

func (r ApiJiraProjectsListRequest) Id(id int32) ApiJiraProjectsListRequest {
	r.id = &id
	return r
}

func (r ApiJiraProjectsListRequest) JiraInstance(jiraInstance int32) ApiJiraProjectsListRequest {
	r.jiraInstance = &jiraInstance
	return r
}

// Number of results to return per page.
func (r ApiJiraProjectsListRequest) Limit(limit int32) ApiJiraProjectsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiJiraProjectsListRequest) Offset(offset int32) ApiJiraProjectsListRequest {
	r.offset = &offset
	return r
}

func (r ApiJiraProjectsListRequest) Product(product int32) ApiJiraProjectsListRequest {
	r.product = &product
	return r
}

func (r ApiJiraProjectsListRequest) ProjectKey(projectKey string) ApiJiraProjectsListRequest {
	r.projectKey = &projectKey
	return r
}

func (r ApiJiraProjectsListRequest) PushAllIssues(pushAllIssues bool) ApiJiraProjectsListRequest {
	r.pushAllIssues = &pushAllIssues
	return r
}

func (r ApiJiraProjectsListRequest) PushNotes(pushNotes bool) ApiJiraProjectsListRequest {
	r.pushNotes = &pushNotes
	return r
}

func (r ApiJiraProjectsListRequest) Execute() (*PaginatedJIRAProjectList, *http.Response, error) {
	return r.ApiService.JiraProjectsListExecute(r)
}

/*
JiraProjectsList Method for JiraProjectsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiJiraProjectsListRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsList(ctx context.Context) ApiJiraProjectsListRequest {
	return ApiJiraProjectsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedJIRAProjectList
func (a *JiraProjectsAPIService) JiraProjectsListExecute(r ApiJiraProjectsListRequest) (*PaginatedJIRAProjectList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedJIRAProjectList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.component != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "component", r.component, "")
	}
	if r.enableEngagementEpicMapping != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enable_engagement_epic_mapping", r.enableEngagementEpicMapping, "")
	}
	if r.engagement != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "engagement", r.engagement, "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.jiraInstance != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "jira_instance", r.jiraInstance, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.product != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product", r.product, "")
	}
	if r.projectKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "project_key", r.projectKey, "")
	}
	if r.pushAllIssues != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "push_all_issues", r.pushAllIssues, "")
	}
	if r.pushNotes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "push_notes", r.pushNotes, "")
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

type ApiJiraProjectsPartialUpdateRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	id int32
	patchedJIRAProjectRequest *PatchedJIRAProjectRequest
}

func (r ApiJiraProjectsPartialUpdateRequest) PatchedJIRAProjectRequest(patchedJIRAProjectRequest PatchedJIRAProjectRequest) ApiJiraProjectsPartialUpdateRequest {
	r.patchedJIRAProjectRequest = &patchedJIRAProjectRequest
	return r
}

func (r ApiJiraProjectsPartialUpdateRequest) Execute() (*JIRAProject, *http.Response, error) {
	return r.ApiService.JiraProjectsPartialUpdateExecute(r)
}

/*
JiraProjectsPartialUpdate Method for JiraProjectsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ project.
 @return ApiJiraProjectsPartialUpdateRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsPartialUpdate(ctx context.Context, id int32) ApiJiraProjectsPartialUpdateRequest {
	return ApiJiraProjectsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return JIRAProject
func (a *JiraProjectsAPIService) JiraProjectsPartialUpdateExecute(r ApiJiraProjectsPartialUpdateRequest) (*JIRAProject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/{id}/"
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
	localVarPostBody = r.patchedJIRAProjectRequest
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

type ApiJiraProjectsRetrieveRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	id int32
}

func (r ApiJiraProjectsRetrieveRequest) Execute() (*JIRAProject, *http.Response, error) {
	return r.ApiService.JiraProjectsRetrieveExecute(r)
}

/*
JiraProjectsRetrieve Method for JiraProjectsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ project.
 @return ApiJiraProjectsRetrieveRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsRetrieve(ctx context.Context, id int32) ApiJiraProjectsRetrieveRequest {
	return ApiJiraProjectsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return JIRAProject
func (a *JiraProjectsAPIService) JiraProjectsRetrieveExecute(r ApiJiraProjectsRetrieveRequest) (*JIRAProject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/{id}/"
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

type ApiJiraProjectsUpdateRequest struct {
	ctx context.Context
	ApiService JiraProjectsAPI
	id int32
	jIRAProjectRequest *JIRAProjectRequest
}

func (r ApiJiraProjectsUpdateRequest) JIRAProjectRequest(jIRAProjectRequest JIRAProjectRequest) ApiJiraProjectsUpdateRequest {
	r.jIRAProjectRequest = &jIRAProjectRequest
	return r
}

func (r ApiJiraProjectsUpdateRequest) Execute() (*JIRAProject, *http.Response, error) {
	return r.ApiService.JiraProjectsUpdateExecute(r)
}

/*
JiraProjectsUpdate Method for JiraProjectsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this jir a_ project.
 @return ApiJiraProjectsUpdateRequest
*/
func (a *JiraProjectsAPIService) JiraProjectsUpdate(ctx context.Context, id int32) ApiJiraProjectsUpdateRequest {
	return ApiJiraProjectsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return JIRAProject
func (a *JiraProjectsAPIService) JiraProjectsUpdateExecute(r ApiJiraProjectsUpdateRequest) (*JIRAProject, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *JIRAProject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JiraProjectsAPIService.JiraProjectsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/jira_projects/{id}/"
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
	localVarPostBody = r.jIRAProjectRequest
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
