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


type NoteTypeAPI interface {

	/*
	NoteTypeCreate Method for NoteTypeCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiNoteTypeCreateRequest
	*/
	NoteTypeCreate(ctx context.Context) ApiNoteTypeCreateRequest

	// NoteTypeCreateExecute executes the request
	//  @return NoteType
	NoteTypeCreateExecute(r ApiNoteTypeCreateRequest) (*NoteType, *http.Response, error)

	/*
	NoteTypeDeletePreviewList Method for NoteTypeDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this note_ type.
	@return ApiNoteTypeDeletePreviewListRequest
	*/
	NoteTypeDeletePreviewList(ctx context.Context, id int32) ApiNoteTypeDeletePreviewListRequest

	// NoteTypeDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	NoteTypeDeletePreviewListExecute(r ApiNoteTypeDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	NoteTypeDestroy Method for NoteTypeDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this note_ type.
	@return ApiNoteTypeDestroyRequest
	*/
	NoteTypeDestroy(ctx context.Context, id int32) ApiNoteTypeDestroyRequest

	// NoteTypeDestroyExecute executes the request
	NoteTypeDestroyExecute(r ApiNoteTypeDestroyRequest) (*http.Response, error)

	/*
	NoteTypeList Method for NoteTypeList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiNoteTypeListRequest
	*/
	NoteTypeList(ctx context.Context) ApiNoteTypeListRequest

	// NoteTypeListExecute executes the request
	//  @return PaginatedNoteTypeList
	NoteTypeListExecute(r ApiNoteTypeListRequest) (*PaginatedNoteTypeList, *http.Response, error)

	/*
	NoteTypePartialUpdate Method for NoteTypePartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this note_ type.
	@return ApiNoteTypePartialUpdateRequest
	*/
	NoteTypePartialUpdate(ctx context.Context, id int32) ApiNoteTypePartialUpdateRequest

	// NoteTypePartialUpdateExecute executes the request
	//  @return NoteType
	NoteTypePartialUpdateExecute(r ApiNoteTypePartialUpdateRequest) (*NoteType, *http.Response, error)

	/*
	NoteTypeRetrieve Method for NoteTypeRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this note_ type.
	@return ApiNoteTypeRetrieveRequest
	*/
	NoteTypeRetrieve(ctx context.Context, id int32) ApiNoteTypeRetrieveRequest

	// NoteTypeRetrieveExecute executes the request
	//  @return NoteType
	NoteTypeRetrieveExecute(r ApiNoteTypeRetrieveRequest) (*NoteType, *http.Response, error)

	/*
	NoteTypeUpdate Method for NoteTypeUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this note_ type.
	@return ApiNoteTypeUpdateRequest
	*/
	NoteTypeUpdate(ctx context.Context, id int32) ApiNoteTypeUpdateRequest

	// NoteTypeUpdateExecute executes the request
	//  @return NoteType
	NoteTypeUpdateExecute(r ApiNoteTypeUpdateRequest) (*NoteType, *http.Response, error)
}

// NoteTypeAPIService NoteTypeAPI service
type NoteTypeAPIService service

type ApiNoteTypeCreateRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	noteTypeRequest *NoteTypeRequest
}

func (r ApiNoteTypeCreateRequest) NoteTypeRequest(noteTypeRequest NoteTypeRequest) ApiNoteTypeCreateRequest {
	r.noteTypeRequest = &noteTypeRequest
	return r
}

func (r ApiNoteTypeCreateRequest) Execute() (*NoteType, *http.Response, error) {
	return r.ApiService.NoteTypeCreateExecute(r)
}

/*
NoteTypeCreate Method for NoteTypeCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNoteTypeCreateRequest
*/
func (a *NoteTypeAPIService) NoteTypeCreate(ctx context.Context) ApiNoteTypeCreateRequest {
	return ApiNoteTypeCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NoteType
func (a *NoteTypeAPIService) NoteTypeCreateExecute(r ApiNoteTypeCreateRequest) (*NoteType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypeCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.noteTypeRequest == nil {
		return localVarReturnValue, nil, reportError("noteTypeRequest is required and must be specified")
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
	localVarPostBody = r.noteTypeRequest
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

type ApiNoteTypeDeletePreviewListRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiNoteTypeDeletePreviewListRequest) Limit(limit int32) ApiNoteTypeDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiNoteTypeDeletePreviewListRequest) Offset(offset int32) ApiNoteTypeDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiNoteTypeDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.NoteTypeDeletePreviewListExecute(r)
}

/*
NoteTypeDeletePreviewList Method for NoteTypeDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this note_ type.
 @return ApiNoteTypeDeletePreviewListRequest
*/
func (a *NoteTypeAPIService) NoteTypeDeletePreviewList(ctx context.Context, id int32) ApiNoteTypeDeletePreviewListRequest {
	return ApiNoteTypeDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *NoteTypeAPIService) NoteTypeDeletePreviewListExecute(r ApiNoteTypeDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypeDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/{id}/delete_preview/"
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

type ApiNoteTypeDestroyRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	id int32
}

func (r ApiNoteTypeDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.NoteTypeDestroyExecute(r)
}

/*
NoteTypeDestroy Method for NoteTypeDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this note_ type.
 @return ApiNoteTypeDestroyRequest
*/
func (a *NoteTypeAPIService) NoteTypeDestroy(ctx context.Context, id int32) ApiNoteTypeDestroyRequest {
	return ApiNoteTypeDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *NoteTypeAPIService) NoteTypeDestroyExecute(r ApiNoteTypeDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypeDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/{id}/"
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

type ApiNoteTypeListRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	description *string
	id *int32
	isActive *bool
	isMandatory *bool
	isSingle *bool
	limit *int32
	name *string
	offset *int32
}

func (r ApiNoteTypeListRequest) Description(description string) ApiNoteTypeListRequest {
	r.description = &description
	return r
}

func (r ApiNoteTypeListRequest) Id(id int32) ApiNoteTypeListRequest {
	r.id = &id
	return r
}

func (r ApiNoteTypeListRequest) IsActive(isActive bool) ApiNoteTypeListRequest {
	r.isActive = &isActive
	return r
}

func (r ApiNoteTypeListRequest) IsMandatory(isMandatory bool) ApiNoteTypeListRequest {
	r.isMandatory = &isMandatory
	return r
}

func (r ApiNoteTypeListRequest) IsSingle(isSingle bool) ApiNoteTypeListRequest {
	r.isSingle = &isSingle
	return r
}

// Number of results to return per page.
func (r ApiNoteTypeListRequest) Limit(limit int32) ApiNoteTypeListRequest {
	r.limit = &limit
	return r
}

func (r ApiNoteTypeListRequest) Name(name string) ApiNoteTypeListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r ApiNoteTypeListRequest) Offset(offset int32) ApiNoteTypeListRequest {
	r.offset = &offset
	return r
}

func (r ApiNoteTypeListRequest) Execute() (*PaginatedNoteTypeList, *http.Response, error) {
	return r.ApiService.NoteTypeListExecute(r)
}

/*
NoteTypeList Method for NoteTypeList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNoteTypeListRequest
*/
func (a *NoteTypeAPIService) NoteTypeList(ctx context.Context) ApiNoteTypeListRequest {
	return ApiNoteTypeListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedNoteTypeList
func (a *NoteTypeAPIService) NoteTypeListExecute(r ApiNoteTypeListRequest) (*PaginatedNoteTypeList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedNoteTypeList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypeList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "description", r.description, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.isActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_active", r.isActive, "form", "")
	}
	if r.isMandatory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_mandatory", r.isMandatory, "form", "")
	}
	if r.isSingle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_single", r.isSingle, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
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

type ApiNoteTypePartialUpdateRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	id int32
	patchedNoteTypeRequest *PatchedNoteTypeRequest
}

func (r ApiNoteTypePartialUpdateRequest) PatchedNoteTypeRequest(patchedNoteTypeRequest PatchedNoteTypeRequest) ApiNoteTypePartialUpdateRequest {
	r.patchedNoteTypeRequest = &patchedNoteTypeRequest
	return r
}

func (r ApiNoteTypePartialUpdateRequest) Execute() (*NoteType, *http.Response, error) {
	return r.ApiService.NoteTypePartialUpdateExecute(r)
}

/*
NoteTypePartialUpdate Method for NoteTypePartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this note_ type.
 @return ApiNoteTypePartialUpdateRequest
*/
func (a *NoteTypeAPIService) NoteTypePartialUpdate(ctx context.Context, id int32) ApiNoteTypePartialUpdateRequest {
	return ApiNoteTypePartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NoteType
func (a *NoteTypeAPIService) NoteTypePartialUpdateExecute(r ApiNoteTypePartialUpdateRequest) (*NoteType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypePartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/{id}/"
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
	localVarPostBody = r.patchedNoteTypeRequest
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

type ApiNoteTypeRetrieveRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	id int32
}

func (r ApiNoteTypeRetrieveRequest) Execute() (*NoteType, *http.Response, error) {
	return r.ApiService.NoteTypeRetrieveExecute(r)
}

/*
NoteTypeRetrieve Method for NoteTypeRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this note_ type.
 @return ApiNoteTypeRetrieveRequest
*/
func (a *NoteTypeAPIService) NoteTypeRetrieve(ctx context.Context, id int32) ApiNoteTypeRetrieveRequest {
	return ApiNoteTypeRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NoteType
func (a *NoteTypeAPIService) NoteTypeRetrieveExecute(r ApiNoteTypeRetrieveRequest) (*NoteType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypeRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/{id}/"
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

type ApiNoteTypeUpdateRequest struct {
	ctx context.Context
	ApiService NoteTypeAPI
	id int32
	noteTypeRequest *NoteTypeRequest
}

func (r ApiNoteTypeUpdateRequest) NoteTypeRequest(noteTypeRequest NoteTypeRequest) ApiNoteTypeUpdateRequest {
	r.noteTypeRequest = &noteTypeRequest
	return r
}

func (r ApiNoteTypeUpdateRequest) Execute() (*NoteType, *http.Response, error) {
	return r.ApiService.NoteTypeUpdateExecute(r)
}

/*
NoteTypeUpdate Method for NoteTypeUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this note_ type.
 @return ApiNoteTypeUpdateRequest
*/
func (a *NoteTypeAPIService) NoteTypeUpdate(ctx context.Context, id int32) ApiNoteTypeUpdateRequest {
	return ApiNoteTypeUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NoteType
func (a *NoteTypeAPIService) NoteTypeUpdateExecute(r ApiNoteTypeUpdateRequest) (*NoteType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NoteType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NoteTypeAPIService.NoteTypeUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/note_type/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.noteTypeRequest == nil {
		return localVarReturnValue, nil, reportError("noteTypeRequest is required and must be specified")
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
	localVarPostBody = r.noteTypeRequest
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
