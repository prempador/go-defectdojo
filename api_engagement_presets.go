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


type EngagementPresetsAPI interface {

	/*
	EngagementPresetsCreate Method for EngagementPresetsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEngagementPresetsCreateRequest
	*/
	EngagementPresetsCreate(ctx context.Context) ApiEngagementPresetsCreateRequest

	// EngagementPresetsCreateExecute executes the request
	//  @return EngagementPresets
	EngagementPresetsCreateExecute(r ApiEngagementPresetsCreateRequest) (*EngagementPresets, *http.Response, error)

	/*
	EngagementPresetsDeletePreviewList Method for EngagementPresetsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this engagement_ presets.
	@return ApiEngagementPresetsDeletePreviewListRequest
	*/
	EngagementPresetsDeletePreviewList(ctx context.Context, id int32) ApiEngagementPresetsDeletePreviewListRequest

	// EngagementPresetsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	EngagementPresetsDeletePreviewListExecute(r ApiEngagementPresetsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	EngagementPresetsDestroy Method for EngagementPresetsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this engagement_ presets.
	@return ApiEngagementPresetsDestroyRequest
	*/
	EngagementPresetsDestroy(ctx context.Context, id int32) ApiEngagementPresetsDestroyRequest

	// EngagementPresetsDestroyExecute executes the request
	EngagementPresetsDestroyExecute(r ApiEngagementPresetsDestroyRequest) (*http.Response, error)

	/*
	EngagementPresetsList Method for EngagementPresetsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEngagementPresetsListRequest
	*/
	EngagementPresetsList(ctx context.Context) ApiEngagementPresetsListRequest

	// EngagementPresetsListExecute executes the request
	//  @return PaginatedEngagementPresetsList
	EngagementPresetsListExecute(r ApiEngagementPresetsListRequest) (*PaginatedEngagementPresetsList, *http.Response, error)

	/*
	EngagementPresetsPartialUpdate Method for EngagementPresetsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this engagement_ presets.
	@return ApiEngagementPresetsPartialUpdateRequest
	*/
	EngagementPresetsPartialUpdate(ctx context.Context, id int32) ApiEngagementPresetsPartialUpdateRequest

	// EngagementPresetsPartialUpdateExecute executes the request
	//  @return EngagementPresets
	EngagementPresetsPartialUpdateExecute(r ApiEngagementPresetsPartialUpdateRequest) (*EngagementPresets, *http.Response, error)

	/*
	EngagementPresetsRetrieve Method for EngagementPresetsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this engagement_ presets.
	@return ApiEngagementPresetsRetrieveRequest
	*/
	EngagementPresetsRetrieve(ctx context.Context, id int32) ApiEngagementPresetsRetrieveRequest

	// EngagementPresetsRetrieveExecute executes the request
	//  @return EngagementPresets
	EngagementPresetsRetrieveExecute(r ApiEngagementPresetsRetrieveRequest) (*EngagementPresets, *http.Response, error)

	/*
	EngagementPresetsUpdate Method for EngagementPresetsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this engagement_ presets.
	@return ApiEngagementPresetsUpdateRequest
	*/
	EngagementPresetsUpdate(ctx context.Context, id int32) ApiEngagementPresetsUpdateRequest

	// EngagementPresetsUpdateExecute executes the request
	//  @return EngagementPresets
	EngagementPresetsUpdateExecute(r ApiEngagementPresetsUpdateRequest) (*EngagementPresets, *http.Response, error)
}

// EngagementPresetsAPIService EngagementPresetsAPI service
type EngagementPresetsAPIService service

type ApiEngagementPresetsCreateRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	engagementPresetsRequest *EngagementPresetsRequest
}

func (r ApiEngagementPresetsCreateRequest) EngagementPresetsRequest(engagementPresetsRequest EngagementPresetsRequest) ApiEngagementPresetsCreateRequest {
	r.engagementPresetsRequest = &engagementPresetsRequest
	return r
}

func (r ApiEngagementPresetsCreateRequest) Execute() (*EngagementPresets, *http.Response, error) {
	return r.ApiService.EngagementPresetsCreateExecute(r)
}

/*
EngagementPresetsCreate Method for EngagementPresetsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEngagementPresetsCreateRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsCreate(ctx context.Context) ApiEngagementPresetsCreateRequest {
	return ApiEngagementPresetsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EngagementPresets
func (a *EngagementPresetsAPIService) EngagementPresetsCreateExecute(r ApiEngagementPresetsCreateRequest) (*EngagementPresets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EngagementPresets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.engagementPresetsRequest == nil {
		return localVarReturnValue, nil, reportError("engagementPresetsRequest is required and must be specified")
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
	localVarPostBody = r.engagementPresetsRequest
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

type ApiEngagementPresetsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiEngagementPresetsDeletePreviewListRequest) Limit(limit int32) ApiEngagementPresetsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiEngagementPresetsDeletePreviewListRequest) Offset(offset int32) ApiEngagementPresetsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiEngagementPresetsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.EngagementPresetsDeletePreviewListExecute(r)
}

/*
EngagementPresetsDeletePreviewList Method for EngagementPresetsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this engagement_ presets.
 @return ApiEngagementPresetsDeletePreviewListRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsDeletePreviewList(ctx context.Context, id int32) ApiEngagementPresetsDeletePreviewListRequest {
	return ApiEngagementPresetsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *EngagementPresetsAPIService) EngagementPresetsDeletePreviewListExecute(r ApiEngagementPresetsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/{id}/delete_preview/"
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

type ApiEngagementPresetsDestroyRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	id int32
}

func (r ApiEngagementPresetsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.EngagementPresetsDestroyExecute(r)
}

/*
EngagementPresetsDestroy Method for EngagementPresetsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this engagement_ presets.
 @return ApiEngagementPresetsDestroyRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsDestroy(ctx context.Context, id int32) ApiEngagementPresetsDestroyRequest {
	return ApiEngagementPresetsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *EngagementPresetsAPIService) EngagementPresetsDestroyExecute(r ApiEngagementPresetsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/{id}/"
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

type ApiEngagementPresetsListRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	id *int32
	limit *int32
	offset *int32
	product *int32
	title *string
}

func (r ApiEngagementPresetsListRequest) Id(id int32) ApiEngagementPresetsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiEngagementPresetsListRequest) Limit(limit int32) ApiEngagementPresetsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiEngagementPresetsListRequest) Offset(offset int32) ApiEngagementPresetsListRequest {
	r.offset = &offset
	return r
}

func (r ApiEngagementPresetsListRequest) Product(product int32) ApiEngagementPresetsListRequest {
	r.product = &product
	return r
}

func (r ApiEngagementPresetsListRequest) Title(title string) ApiEngagementPresetsListRequest {
	r.title = &title
	return r
}

func (r ApiEngagementPresetsListRequest) Execute() (*PaginatedEngagementPresetsList, *http.Response, error) {
	return r.ApiService.EngagementPresetsListExecute(r)
}

/*
EngagementPresetsList Method for EngagementPresetsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEngagementPresetsListRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsList(ctx context.Context) ApiEngagementPresetsListRequest {
	return ApiEngagementPresetsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedEngagementPresetsList
func (a *EngagementPresetsAPIService) EngagementPresetsListExecute(r ApiEngagementPresetsListRequest) (*PaginatedEngagementPresetsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedEngagementPresetsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
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

type ApiEngagementPresetsPartialUpdateRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	id int32
	patchedEngagementPresetsRequest *PatchedEngagementPresetsRequest
}

func (r ApiEngagementPresetsPartialUpdateRequest) PatchedEngagementPresetsRequest(patchedEngagementPresetsRequest PatchedEngagementPresetsRequest) ApiEngagementPresetsPartialUpdateRequest {
	r.patchedEngagementPresetsRequest = &patchedEngagementPresetsRequest
	return r
}

func (r ApiEngagementPresetsPartialUpdateRequest) Execute() (*EngagementPresets, *http.Response, error) {
	return r.ApiService.EngagementPresetsPartialUpdateExecute(r)
}

/*
EngagementPresetsPartialUpdate Method for EngagementPresetsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this engagement_ presets.
 @return ApiEngagementPresetsPartialUpdateRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsPartialUpdate(ctx context.Context, id int32) ApiEngagementPresetsPartialUpdateRequest {
	return ApiEngagementPresetsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EngagementPresets
func (a *EngagementPresetsAPIService) EngagementPresetsPartialUpdateExecute(r ApiEngagementPresetsPartialUpdateRequest) (*EngagementPresets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EngagementPresets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/{id}/"
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
	localVarPostBody = r.patchedEngagementPresetsRequest
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

type ApiEngagementPresetsRetrieveRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	id int32
}

func (r ApiEngagementPresetsRetrieveRequest) Execute() (*EngagementPresets, *http.Response, error) {
	return r.ApiService.EngagementPresetsRetrieveExecute(r)
}

/*
EngagementPresetsRetrieve Method for EngagementPresetsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this engagement_ presets.
 @return ApiEngagementPresetsRetrieveRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsRetrieve(ctx context.Context, id int32) ApiEngagementPresetsRetrieveRequest {
	return ApiEngagementPresetsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EngagementPresets
func (a *EngagementPresetsAPIService) EngagementPresetsRetrieveExecute(r ApiEngagementPresetsRetrieveRequest) (*EngagementPresets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EngagementPresets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/{id}/"
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

type ApiEngagementPresetsUpdateRequest struct {
	ctx context.Context
	ApiService EngagementPresetsAPI
	id int32
	engagementPresetsRequest *EngagementPresetsRequest
}

func (r ApiEngagementPresetsUpdateRequest) EngagementPresetsRequest(engagementPresetsRequest EngagementPresetsRequest) ApiEngagementPresetsUpdateRequest {
	r.engagementPresetsRequest = &engagementPresetsRequest
	return r
}

func (r ApiEngagementPresetsUpdateRequest) Execute() (*EngagementPresets, *http.Response, error) {
	return r.ApiService.EngagementPresetsUpdateExecute(r)
}

/*
EngagementPresetsUpdate Method for EngagementPresetsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this engagement_ presets.
 @return ApiEngagementPresetsUpdateRequest
*/
func (a *EngagementPresetsAPIService) EngagementPresetsUpdate(ctx context.Context, id int32) ApiEngagementPresetsUpdateRequest {
	return ApiEngagementPresetsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EngagementPresets
func (a *EngagementPresetsAPIService) EngagementPresetsUpdateExecute(r ApiEngagementPresetsUpdateRequest) (*EngagementPresets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EngagementPresets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EngagementPresetsAPIService.EngagementPresetsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/engagement_presets/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.engagementPresetsRequest == nil {
		return localVarReturnValue, nil, reportError("engagementPresetsRequest is required and must be specified")
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
	localVarPostBody = r.engagementPresetsRequest
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
