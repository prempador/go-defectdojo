/*
Defect Dojo API v2

Defect Dojo - Open Source vulnerability Management made easy. Prefetch related parameters/responses not yet in the schema.

API version: 2.31.1
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


type FindingTemplatesAPI interface {

	/*
	FindingTemplatesCreate Method for FindingTemplatesCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindingTemplatesCreateRequest
	*/
	FindingTemplatesCreate(ctx context.Context) ApiFindingTemplatesCreateRequest

	// FindingTemplatesCreateExecute executes the request
	//  @return FindingTemplate
	FindingTemplatesCreateExecute(r ApiFindingTemplatesCreateRequest) (*FindingTemplate, *http.Response, error)

	/*
	FindingTemplatesDeletePreviewList Method for FindingTemplatesDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this finding_ template.
	@return ApiFindingTemplatesDeletePreviewListRequest
	*/
	FindingTemplatesDeletePreviewList(ctx context.Context, id int32) ApiFindingTemplatesDeletePreviewListRequest

	// FindingTemplatesDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	FindingTemplatesDeletePreviewListExecute(r ApiFindingTemplatesDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	FindingTemplatesDestroy Method for FindingTemplatesDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this finding_ template.
	@return ApiFindingTemplatesDestroyRequest
	*/
	FindingTemplatesDestroy(ctx context.Context, id int32) ApiFindingTemplatesDestroyRequest

	// FindingTemplatesDestroyExecute executes the request
	FindingTemplatesDestroyExecute(r ApiFindingTemplatesDestroyRequest) (*http.Response, error)

	/*
	FindingTemplatesList Method for FindingTemplatesList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiFindingTemplatesListRequest
	*/
	FindingTemplatesList(ctx context.Context) ApiFindingTemplatesListRequest

	// FindingTemplatesListExecute executes the request
	//  @return PaginatedFindingTemplateList
	FindingTemplatesListExecute(r ApiFindingTemplatesListRequest) (*PaginatedFindingTemplateList, *http.Response, error)

	/*
	FindingTemplatesPartialUpdate Method for FindingTemplatesPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this finding_ template.
	@return ApiFindingTemplatesPartialUpdateRequest
	*/
	FindingTemplatesPartialUpdate(ctx context.Context, id int32) ApiFindingTemplatesPartialUpdateRequest

	// FindingTemplatesPartialUpdateExecute executes the request
	//  @return FindingTemplate
	FindingTemplatesPartialUpdateExecute(r ApiFindingTemplatesPartialUpdateRequest) (*FindingTemplate, *http.Response, error)

	/*
	FindingTemplatesRetrieve Method for FindingTemplatesRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this finding_ template.
	@return ApiFindingTemplatesRetrieveRequest
	*/
	FindingTemplatesRetrieve(ctx context.Context, id int32) ApiFindingTemplatesRetrieveRequest

	// FindingTemplatesRetrieveExecute executes the request
	//  @return FindingTemplate
	FindingTemplatesRetrieveExecute(r ApiFindingTemplatesRetrieveRequest) (*FindingTemplate, *http.Response, error)

	/*
	FindingTemplatesUpdate Method for FindingTemplatesUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this finding_ template.
	@return ApiFindingTemplatesUpdateRequest
	*/
	FindingTemplatesUpdate(ctx context.Context, id int32) ApiFindingTemplatesUpdateRequest

	// FindingTemplatesUpdateExecute executes the request
	//  @return FindingTemplate
	FindingTemplatesUpdateExecute(r ApiFindingTemplatesUpdateRequest) (*FindingTemplate, *http.Response, error)
}

// FindingTemplatesAPIService FindingTemplatesAPI service
type FindingTemplatesAPIService service

type ApiFindingTemplatesCreateRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	findingTemplateRequest *FindingTemplateRequest
}

func (r ApiFindingTemplatesCreateRequest) FindingTemplateRequest(findingTemplateRequest FindingTemplateRequest) ApiFindingTemplatesCreateRequest {
	r.findingTemplateRequest = &findingTemplateRequest
	return r
}

func (r ApiFindingTemplatesCreateRequest) Execute() (*FindingTemplate, *http.Response, error) {
	return r.ApiService.FindingTemplatesCreateExecute(r)
}

/*
FindingTemplatesCreate Method for FindingTemplatesCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindingTemplatesCreateRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesCreate(ctx context.Context) ApiFindingTemplatesCreateRequest {
	return ApiFindingTemplatesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FindingTemplate
func (a *FindingTemplatesAPIService) FindingTemplatesCreateExecute(r ApiFindingTemplatesCreateRequest) (*FindingTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FindingTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.findingTemplateRequest == nil {
		return localVarReturnValue, nil, reportError("findingTemplateRequest is required and must be specified")
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
	localVarPostBody = r.findingTemplateRequest
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

type ApiFindingTemplatesDeletePreviewListRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiFindingTemplatesDeletePreviewListRequest) Limit(limit int32) ApiFindingTemplatesDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiFindingTemplatesDeletePreviewListRequest) Offset(offset int32) ApiFindingTemplatesDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiFindingTemplatesDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.FindingTemplatesDeletePreviewListExecute(r)
}

/*
FindingTemplatesDeletePreviewList Method for FindingTemplatesDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this finding_ template.
 @return ApiFindingTemplatesDeletePreviewListRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesDeletePreviewList(ctx context.Context, id int32) ApiFindingTemplatesDeletePreviewListRequest {
	return ApiFindingTemplatesDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *FindingTemplatesAPIService) FindingTemplatesDeletePreviewListExecute(r ApiFindingTemplatesDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/{id}/delete_preview/"
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

type ApiFindingTemplatesDestroyRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	id int32
}

func (r ApiFindingTemplatesDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.FindingTemplatesDestroyExecute(r)
}

/*
FindingTemplatesDestroy Method for FindingTemplatesDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this finding_ template.
 @return ApiFindingTemplatesDestroyRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesDestroy(ctx context.Context, id int32) ApiFindingTemplatesDestroyRequest {
	return ApiFindingTemplatesDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *FindingTemplatesAPIService) FindingTemplatesDestroyExecute(r ApiFindingTemplatesDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/{id}/"
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

type ApiFindingTemplatesListRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	cwe *int32
	description *string
	id *int32
	limit *int32
	mitigation *string
	notTag *string
	notTags *[]string
	o *[]string
	offset *int32
	severity *string
	tag *string
	tags *[]string
	title *string
}

func (r ApiFindingTemplatesListRequest) Cwe(cwe int32) ApiFindingTemplatesListRequest {
	r.cwe = &cwe
	return r
}

func (r ApiFindingTemplatesListRequest) Description(description string) ApiFindingTemplatesListRequest {
	r.description = &description
	return r
}

func (r ApiFindingTemplatesListRequest) Id(id int32) ApiFindingTemplatesListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiFindingTemplatesListRequest) Limit(limit int32) ApiFindingTemplatesListRequest {
	r.limit = &limit
	return r
}

func (r ApiFindingTemplatesListRequest) Mitigation(mitigation string) ApiFindingTemplatesListRequest {
	r.mitigation = &mitigation
	return r
}

// Not Tag name contains
func (r ApiFindingTemplatesListRequest) NotTag(notTag string) ApiFindingTemplatesListRequest {
	r.notTag = &notTag
	return r
}

// Comma separated list of exact tags not present on model
func (r ApiFindingTemplatesListRequest) NotTags(notTags []string) ApiFindingTemplatesListRequest {
	r.notTags = &notTags
	return r
}

// Ordering  * &#x60;title&#x60; - Title * &#x60;-title&#x60; - Title (descending) * &#x60;cwe&#x60; - Cwe * &#x60;-cwe&#x60; - Cwe (descending)
func (r ApiFindingTemplatesListRequest) O(o []string) ApiFindingTemplatesListRequest {
	r.o = &o
	return r
}

// The initial index from which to return the results.
func (r ApiFindingTemplatesListRequest) Offset(offset int32) ApiFindingTemplatesListRequest {
	r.offset = &offset
	return r
}

func (r ApiFindingTemplatesListRequest) Severity(severity string) ApiFindingTemplatesListRequest {
	r.severity = &severity
	return r
}

// Tag name contains
func (r ApiFindingTemplatesListRequest) Tag(tag string) ApiFindingTemplatesListRequest {
	r.tag = &tag
	return r
}

// Comma separated list of exact tags
func (r ApiFindingTemplatesListRequest) Tags(tags []string) ApiFindingTemplatesListRequest {
	r.tags = &tags
	return r
}

func (r ApiFindingTemplatesListRequest) Title(title string) ApiFindingTemplatesListRequest {
	r.title = &title
	return r
}

func (r ApiFindingTemplatesListRequest) Execute() (*PaginatedFindingTemplateList, *http.Response, error) {
	return r.ApiService.FindingTemplatesListExecute(r)
}

/*
FindingTemplatesList Method for FindingTemplatesList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindingTemplatesListRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesList(ctx context.Context) ApiFindingTemplatesListRequest {
	return ApiFindingTemplatesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedFindingTemplateList
func (a *FindingTemplatesAPIService) FindingTemplatesListExecute(r ApiFindingTemplatesListRequest) (*PaginatedFindingTemplateList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedFindingTemplateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cwe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cwe", r.cwe, "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description", r.description, "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.mitigation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mitigation", r.mitigation, "")
	}
	if r.notTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "not_tag", r.notTag, "")
	}
	if r.notTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "not_tags", r.notTags, "csv")
	}
	if r.o != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "o", r.o, "csv")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.severity != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "severity", r.severity, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "csv")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title", r.title, "")
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

type ApiFindingTemplatesPartialUpdateRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	id int32
	patchedFindingTemplateRequest *PatchedFindingTemplateRequest
}

func (r ApiFindingTemplatesPartialUpdateRequest) PatchedFindingTemplateRequest(patchedFindingTemplateRequest PatchedFindingTemplateRequest) ApiFindingTemplatesPartialUpdateRequest {
	r.patchedFindingTemplateRequest = &patchedFindingTemplateRequest
	return r
}

func (r ApiFindingTemplatesPartialUpdateRequest) Execute() (*FindingTemplate, *http.Response, error) {
	return r.ApiService.FindingTemplatesPartialUpdateExecute(r)
}

/*
FindingTemplatesPartialUpdate Method for FindingTemplatesPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this finding_ template.
 @return ApiFindingTemplatesPartialUpdateRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesPartialUpdate(ctx context.Context, id int32) ApiFindingTemplatesPartialUpdateRequest {
	return ApiFindingTemplatesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FindingTemplate
func (a *FindingTemplatesAPIService) FindingTemplatesPartialUpdateExecute(r ApiFindingTemplatesPartialUpdateRequest) (*FindingTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FindingTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/{id}/"
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
	localVarPostBody = r.patchedFindingTemplateRequest
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

type ApiFindingTemplatesRetrieveRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	id int32
}

func (r ApiFindingTemplatesRetrieveRequest) Execute() (*FindingTemplate, *http.Response, error) {
	return r.ApiService.FindingTemplatesRetrieveExecute(r)
}

/*
FindingTemplatesRetrieve Method for FindingTemplatesRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this finding_ template.
 @return ApiFindingTemplatesRetrieveRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesRetrieve(ctx context.Context, id int32) ApiFindingTemplatesRetrieveRequest {
	return ApiFindingTemplatesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FindingTemplate
func (a *FindingTemplatesAPIService) FindingTemplatesRetrieveExecute(r ApiFindingTemplatesRetrieveRequest) (*FindingTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FindingTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/{id}/"
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

type ApiFindingTemplatesUpdateRequest struct {
	ctx context.Context
	ApiService FindingTemplatesAPI
	id int32
	findingTemplateRequest *FindingTemplateRequest
}

func (r ApiFindingTemplatesUpdateRequest) FindingTemplateRequest(findingTemplateRequest FindingTemplateRequest) ApiFindingTemplatesUpdateRequest {
	r.findingTemplateRequest = &findingTemplateRequest
	return r
}

func (r ApiFindingTemplatesUpdateRequest) Execute() (*FindingTemplate, *http.Response, error) {
	return r.ApiService.FindingTemplatesUpdateExecute(r)
}

/*
FindingTemplatesUpdate Method for FindingTemplatesUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this finding_ template.
 @return ApiFindingTemplatesUpdateRequest
*/
func (a *FindingTemplatesAPIService) FindingTemplatesUpdate(ctx context.Context, id int32) ApiFindingTemplatesUpdateRequest {
	return ApiFindingTemplatesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FindingTemplate
func (a *FindingTemplatesAPIService) FindingTemplatesUpdateExecute(r ApiFindingTemplatesUpdateRequest) (*FindingTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FindingTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FindingTemplatesAPIService.FindingTemplatesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/finding_templates/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.findingTemplateRequest == nil {
		return localVarReturnValue, nil, reportError("findingTemplateRequest is required and must be specified")
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
	localVarPostBody = r.findingTemplateRequest
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
