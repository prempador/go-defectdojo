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


type GlobalRolesAPI interface {

	/*
	GlobalRolesCreate Method for GlobalRolesCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalRolesCreateRequest
	*/
	GlobalRolesCreate(ctx context.Context) ApiGlobalRolesCreateRequest

	// GlobalRolesCreateExecute executes the request
	//  @return GlobalRole
	GlobalRolesCreateExecute(r ApiGlobalRolesCreateRequest) (*GlobalRole, *http.Response, error)

	/*
	GlobalRolesDeletePreviewList Method for GlobalRolesDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this global_ role.
	@return ApiGlobalRolesDeletePreviewListRequest
	*/
	GlobalRolesDeletePreviewList(ctx context.Context, id int32) ApiGlobalRolesDeletePreviewListRequest

	// GlobalRolesDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	GlobalRolesDeletePreviewListExecute(r ApiGlobalRolesDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	GlobalRolesDestroy Method for GlobalRolesDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this global_ role.
	@return ApiGlobalRolesDestroyRequest
	*/
	GlobalRolesDestroy(ctx context.Context, id int32) ApiGlobalRolesDestroyRequest

	// GlobalRolesDestroyExecute executes the request
	GlobalRolesDestroyExecute(r ApiGlobalRolesDestroyRequest) (*http.Response, error)

	/*
	GlobalRolesList Method for GlobalRolesList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalRolesListRequest
	*/
	GlobalRolesList(ctx context.Context) ApiGlobalRolesListRequest

	// GlobalRolesListExecute executes the request
	//  @return PaginatedGlobalRoleList
	GlobalRolesListExecute(r ApiGlobalRolesListRequest) (*PaginatedGlobalRoleList, *http.Response, error)

	/*
	GlobalRolesPartialUpdate Method for GlobalRolesPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this global_ role.
	@return ApiGlobalRolesPartialUpdateRequest
	*/
	GlobalRolesPartialUpdate(ctx context.Context, id int32) ApiGlobalRolesPartialUpdateRequest

	// GlobalRolesPartialUpdateExecute executes the request
	//  @return GlobalRole
	GlobalRolesPartialUpdateExecute(r ApiGlobalRolesPartialUpdateRequest) (*GlobalRole, *http.Response, error)

	/*
	GlobalRolesRetrieve Method for GlobalRolesRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this global_ role.
	@return ApiGlobalRolesRetrieveRequest
	*/
	GlobalRolesRetrieve(ctx context.Context, id int32) ApiGlobalRolesRetrieveRequest

	// GlobalRolesRetrieveExecute executes the request
	//  @return GlobalRole
	GlobalRolesRetrieveExecute(r ApiGlobalRolesRetrieveRequest) (*GlobalRole, *http.Response, error)

	/*
	GlobalRolesUpdate Method for GlobalRolesUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this global_ role.
	@return ApiGlobalRolesUpdateRequest
	*/
	GlobalRolesUpdate(ctx context.Context, id int32) ApiGlobalRolesUpdateRequest

	// GlobalRolesUpdateExecute executes the request
	//  @return GlobalRole
	GlobalRolesUpdateExecute(r ApiGlobalRolesUpdateRequest) (*GlobalRole, *http.Response, error)
}

// GlobalRolesAPIService GlobalRolesAPI service
type GlobalRolesAPIService service

type ApiGlobalRolesCreateRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	globalRoleRequest *GlobalRoleRequest
}

func (r ApiGlobalRolesCreateRequest) GlobalRoleRequest(globalRoleRequest GlobalRoleRequest) ApiGlobalRolesCreateRequest {
	r.globalRoleRequest = &globalRoleRequest
	return r
}

func (r ApiGlobalRolesCreateRequest) Execute() (*GlobalRole, *http.Response, error) {
	return r.ApiService.GlobalRolesCreateExecute(r)
}

/*
GlobalRolesCreate Method for GlobalRolesCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalRolesCreateRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesCreate(ctx context.Context) ApiGlobalRolesCreateRequest {
	return ApiGlobalRolesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GlobalRole
func (a *GlobalRolesAPIService) GlobalRolesCreateExecute(r ApiGlobalRolesCreateRequest) (*GlobalRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/"

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
	localVarPostBody = r.globalRoleRequest
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

type ApiGlobalRolesDeletePreviewListRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiGlobalRolesDeletePreviewListRequest) Limit(limit int32) ApiGlobalRolesDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiGlobalRolesDeletePreviewListRequest) Offset(offset int32) ApiGlobalRolesDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiGlobalRolesDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.GlobalRolesDeletePreviewListExecute(r)
}

/*
GlobalRolesDeletePreviewList Method for GlobalRolesDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this global_ role.
 @return ApiGlobalRolesDeletePreviewListRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesDeletePreviewList(ctx context.Context, id int32) ApiGlobalRolesDeletePreviewListRequest {
	return ApiGlobalRolesDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *GlobalRolesAPIService) GlobalRolesDeletePreviewListExecute(r ApiGlobalRolesDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/{id}/delete_preview/"
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

type ApiGlobalRolesDestroyRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	id int32
}

func (r ApiGlobalRolesDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.GlobalRolesDestroyExecute(r)
}

/*
GlobalRolesDestroy Method for GlobalRolesDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this global_ role.
 @return ApiGlobalRolesDestroyRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesDestroy(ctx context.Context, id int32) ApiGlobalRolesDestroyRequest {
	return ApiGlobalRolesDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GlobalRolesAPIService) GlobalRolesDestroyExecute(r ApiGlobalRolesDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/{id}/"
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

type ApiGlobalRolesListRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	group *int32
	id *int32
	limit *int32
	offset *int32
	role *int32
	user *int32
}

func (r ApiGlobalRolesListRequest) Group(group int32) ApiGlobalRolesListRequest {
	r.group = &group
	return r
}

func (r ApiGlobalRolesListRequest) Id(id int32) ApiGlobalRolesListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiGlobalRolesListRequest) Limit(limit int32) ApiGlobalRolesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiGlobalRolesListRequest) Offset(offset int32) ApiGlobalRolesListRequest {
	r.offset = &offset
	return r
}

func (r ApiGlobalRolesListRequest) Role(role int32) ApiGlobalRolesListRequest {
	r.role = &role
	return r
}

func (r ApiGlobalRolesListRequest) User(user int32) ApiGlobalRolesListRequest {
	r.user = &user
	return r
}

func (r ApiGlobalRolesListRequest) Execute() (*PaginatedGlobalRoleList, *http.Response, error) {
	return r.ApiService.GlobalRolesListExecute(r)
}

/*
GlobalRolesList Method for GlobalRolesList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalRolesListRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesList(ctx context.Context) ApiGlobalRolesListRequest {
	return ApiGlobalRolesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedGlobalRoleList
func (a *GlobalRolesAPIService) GlobalRolesListExecute(r ApiGlobalRolesListRequest) (*PaginatedGlobalRoleList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGlobalRoleList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.role != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "role", r.role, "")
	}
	if r.user != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user", r.user, "")
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

type ApiGlobalRolesPartialUpdateRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	id int32
	patchedGlobalRoleRequest *PatchedGlobalRoleRequest
}

func (r ApiGlobalRolesPartialUpdateRequest) PatchedGlobalRoleRequest(patchedGlobalRoleRequest PatchedGlobalRoleRequest) ApiGlobalRolesPartialUpdateRequest {
	r.patchedGlobalRoleRequest = &patchedGlobalRoleRequest
	return r
}

func (r ApiGlobalRolesPartialUpdateRequest) Execute() (*GlobalRole, *http.Response, error) {
	return r.ApiService.GlobalRolesPartialUpdateExecute(r)
}

/*
GlobalRolesPartialUpdate Method for GlobalRolesPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this global_ role.
 @return ApiGlobalRolesPartialUpdateRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesPartialUpdate(ctx context.Context, id int32) ApiGlobalRolesPartialUpdateRequest {
	return ApiGlobalRolesPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GlobalRole
func (a *GlobalRolesAPIService) GlobalRolesPartialUpdateExecute(r ApiGlobalRolesPartialUpdateRequest) (*GlobalRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/{id}/"
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
	localVarPostBody = r.patchedGlobalRoleRequest
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

type ApiGlobalRolesRetrieveRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	id int32
}

func (r ApiGlobalRolesRetrieveRequest) Execute() (*GlobalRole, *http.Response, error) {
	return r.ApiService.GlobalRolesRetrieveExecute(r)
}

/*
GlobalRolesRetrieve Method for GlobalRolesRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this global_ role.
 @return ApiGlobalRolesRetrieveRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesRetrieve(ctx context.Context, id int32) ApiGlobalRolesRetrieveRequest {
	return ApiGlobalRolesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GlobalRole
func (a *GlobalRolesAPIService) GlobalRolesRetrieveExecute(r ApiGlobalRolesRetrieveRequest) (*GlobalRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/{id}/"
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

type ApiGlobalRolesUpdateRequest struct {
	ctx context.Context
	ApiService GlobalRolesAPI
	id int32
	globalRoleRequest *GlobalRoleRequest
}

func (r ApiGlobalRolesUpdateRequest) GlobalRoleRequest(globalRoleRequest GlobalRoleRequest) ApiGlobalRolesUpdateRequest {
	r.globalRoleRequest = &globalRoleRequest
	return r
}

func (r ApiGlobalRolesUpdateRequest) Execute() (*GlobalRole, *http.Response, error) {
	return r.ApiService.GlobalRolesUpdateExecute(r)
}

/*
GlobalRolesUpdate Method for GlobalRolesUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this global_ role.
 @return ApiGlobalRolesUpdateRequest
*/
func (a *GlobalRolesAPIService) GlobalRolesUpdate(ctx context.Context, id int32) ApiGlobalRolesUpdateRequest {
	return ApiGlobalRolesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GlobalRole
func (a *GlobalRolesAPIService) GlobalRolesUpdateExecute(r ApiGlobalRolesUpdateRequest) (*GlobalRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRolesAPIService.GlobalRolesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/global_roles/{id}/"
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
	localVarPostBody = r.globalRoleRequest
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
