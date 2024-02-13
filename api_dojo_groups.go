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
	"reflect"
)


type DojoGroupsAPI interface {

	/*
	DojoGroupsCreate Method for DojoGroupsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDojoGroupsCreateRequest
	*/
	DojoGroupsCreate(ctx context.Context) ApiDojoGroupsCreateRequest

	// DojoGroupsCreateExecute executes the request
	//  @return DojoGroup
	DojoGroupsCreateExecute(r ApiDojoGroupsCreateRequest) (*DojoGroup, *http.Response, error)

	/*
	DojoGroupsDeletePreviewList Method for DojoGroupsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this dojo_ group.
	@return ApiDojoGroupsDeletePreviewListRequest
	*/
	DojoGroupsDeletePreviewList(ctx context.Context, id int32) ApiDojoGroupsDeletePreviewListRequest

	// DojoGroupsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	DojoGroupsDeletePreviewListExecute(r ApiDojoGroupsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	DojoGroupsDestroy Method for DojoGroupsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this dojo_ group.
	@return ApiDojoGroupsDestroyRequest
	*/
	DojoGroupsDestroy(ctx context.Context, id int32) ApiDojoGroupsDestroyRequest

	// DojoGroupsDestroyExecute executes the request
	DojoGroupsDestroyExecute(r ApiDojoGroupsDestroyRequest) (*http.Response, error)

	/*
	DojoGroupsList Method for DojoGroupsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDojoGroupsListRequest
	*/
	DojoGroupsList(ctx context.Context) ApiDojoGroupsListRequest

	// DojoGroupsListExecute executes the request
	//  @return PaginatedDojoGroupList
	DojoGroupsListExecute(r ApiDojoGroupsListRequest) (*PaginatedDojoGroupList, *http.Response, error)

	/*
	DojoGroupsPartialUpdate Method for DojoGroupsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this dojo_ group.
	@return ApiDojoGroupsPartialUpdateRequest
	*/
	DojoGroupsPartialUpdate(ctx context.Context, id int32) ApiDojoGroupsPartialUpdateRequest

	// DojoGroupsPartialUpdateExecute executes the request
	//  @return DojoGroup
	DojoGroupsPartialUpdateExecute(r ApiDojoGroupsPartialUpdateRequest) (*DojoGroup, *http.Response, error)

	/*
	DojoGroupsRetrieve Method for DojoGroupsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this dojo_ group.
	@return ApiDojoGroupsRetrieveRequest
	*/
	DojoGroupsRetrieve(ctx context.Context, id int32) ApiDojoGroupsRetrieveRequest

	// DojoGroupsRetrieveExecute executes the request
	//  @return DojoGroup
	DojoGroupsRetrieveExecute(r ApiDojoGroupsRetrieveRequest) (*DojoGroup, *http.Response, error)

	/*
	DojoGroupsUpdate Method for DojoGroupsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this dojo_ group.
	@return ApiDojoGroupsUpdateRequest
	*/
	DojoGroupsUpdate(ctx context.Context, id int32) ApiDojoGroupsUpdateRequest

	// DojoGroupsUpdateExecute executes the request
	//  @return DojoGroup
	DojoGroupsUpdateExecute(r ApiDojoGroupsUpdateRequest) (*DojoGroup, *http.Response, error)
}

// DojoGroupsAPIService DojoGroupsAPI service
type DojoGroupsAPIService service

type ApiDojoGroupsCreateRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	dojoGroupRequest *DojoGroupRequest
}

func (r ApiDojoGroupsCreateRequest) DojoGroupRequest(dojoGroupRequest DojoGroupRequest) ApiDojoGroupsCreateRequest {
	r.dojoGroupRequest = &dojoGroupRequest
	return r
}

func (r ApiDojoGroupsCreateRequest) Execute() (*DojoGroup, *http.Response, error) {
	return r.ApiService.DojoGroupsCreateExecute(r)
}

/*
DojoGroupsCreate Method for DojoGroupsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDojoGroupsCreateRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsCreate(ctx context.Context) ApiDojoGroupsCreateRequest {
	return ApiDojoGroupsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DojoGroup
func (a *DojoGroupsAPIService) DojoGroupsCreateExecute(r ApiDojoGroupsCreateRequest) (*DojoGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DojoGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dojoGroupRequest == nil {
		return localVarReturnValue, nil, reportError("dojoGroupRequest is required and must be specified")
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
	localVarPostBody = r.dojoGroupRequest
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

type ApiDojoGroupsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiDojoGroupsDeletePreviewListRequest) Limit(limit int32) ApiDojoGroupsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiDojoGroupsDeletePreviewListRequest) Offset(offset int32) ApiDojoGroupsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiDojoGroupsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.DojoGroupsDeletePreviewListExecute(r)
}

/*
DojoGroupsDeletePreviewList Method for DojoGroupsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this dojo_ group.
 @return ApiDojoGroupsDeletePreviewListRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsDeletePreviewList(ctx context.Context, id int32) ApiDojoGroupsDeletePreviewListRequest {
	return ApiDojoGroupsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *DojoGroupsAPIService) DojoGroupsDeletePreviewListExecute(r ApiDojoGroupsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/{id}/delete_preview/"
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

type ApiDojoGroupsDestroyRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	id int32
}

func (r ApiDojoGroupsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DojoGroupsDestroyExecute(r)
}

/*
DojoGroupsDestroy Method for DojoGroupsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this dojo_ group.
 @return ApiDojoGroupsDestroyRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsDestroy(ctx context.Context, id int32) ApiDojoGroupsDestroyRequest {
	return ApiDojoGroupsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *DojoGroupsAPIService) DojoGroupsDestroyExecute(r ApiDojoGroupsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/{id}/"
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

type ApiDojoGroupsListRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	id *int32
	limit *int32
	name *string
	offset *int32
	prefetch *[]string
	socialProvider *string
}

func (r ApiDojoGroupsListRequest) Id(id int32) ApiDojoGroupsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiDojoGroupsListRequest) Limit(limit int32) ApiDojoGroupsListRequest {
	r.limit = &limit
	return r
}

func (r ApiDojoGroupsListRequest) Name(name string) ApiDojoGroupsListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r ApiDojoGroupsListRequest) Offset(offset int32) ApiDojoGroupsListRequest {
	r.offset = &offset
	return r
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiDojoGroupsListRequest) Prefetch(prefetch []string) ApiDojoGroupsListRequest {
	r.prefetch = &prefetch
	return r
}

// Group imported from a social provider.  * &#x60;AzureAD&#x60; - AzureAD
func (r ApiDojoGroupsListRequest) SocialProvider(socialProvider string) ApiDojoGroupsListRequest {
	r.socialProvider = &socialProvider
	return r
}

func (r ApiDojoGroupsListRequest) Execute() (*PaginatedDojoGroupList, *http.Response, error) {
	return r.ApiService.DojoGroupsListExecute(r)
}

/*
DojoGroupsList Method for DojoGroupsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDojoGroupsListRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsList(ctx context.Context) ApiDojoGroupsListRequest {
	return ApiDojoGroupsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedDojoGroupList
func (a *DojoGroupsAPIService) DojoGroupsListExecute(r ApiDojoGroupsListRequest) (*PaginatedDojoGroupList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDojoGroupList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.prefetch != nil {
		t := *r.prefetch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", t, "multi")
		}
	}
	if r.socialProvider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "social_provider", r.socialProvider, "")
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

type ApiDojoGroupsPartialUpdateRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	id int32
	patchedDojoGroupRequest *PatchedDojoGroupRequest
}

func (r ApiDojoGroupsPartialUpdateRequest) PatchedDojoGroupRequest(patchedDojoGroupRequest PatchedDojoGroupRequest) ApiDojoGroupsPartialUpdateRequest {
	r.patchedDojoGroupRequest = &patchedDojoGroupRequest
	return r
}

func (r ApiDojoGroupsPartialUpdateRequest) Execute() (*DojoGroup, *http.Response, error) {
	return r.ApiService.DojoGroupsPartialUpdateExecute(r)
}

/*
DojoGroupsPartialUpdate Method for DojoGroupsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this dojo_ group.
 @return ApiDojoGroupsPartialUpdateRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsPartialUpdate(ctx context.Context, id int32) ApiDojoGroupsPartialUpdateRequest {
	return ApiDojoGroupsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DojoGroup
func (a *DojoGroupsAPIService) DojoGroupsPartialUpdateExecute(r ApiDojoGroupsPartialUpdateRequest) (*DojoGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DojoGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/{id}/"
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
	localVarPostBody = r.patchedDojoGroupRequest
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

type ApiDojoGroupsRetrieveRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	id int32
	prefetch *[]string
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiDojoGroupsRetrieveRequest) Prefetch(prefetch []string) ApiDojoGroupsRetrieveRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiDojoGroupsRetrieveRequest) Execute() (*DojoGroup, *http.Response, error) {
	return r.ApiService.DojoGroupsRetrieveExecute(r)
}

/*
DojoGroupsRetrieve Method for DojoGroupsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this dojo_ group.
 @return ApiDojoGroupsRetrieveRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsRetrieve(ctx context.Context, id int32) ApiDojoGroupsRetrieveRequest {
	return ApiDojoGroupsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DojoGroup
func (a *DojoGroupsAPIService) DojoGroupsRetrieveExecute(r ApiDojoGroupsRetrieveRequest) (*DojoGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DojoGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.prefetch != nil {
		t := *r.prefetch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", t, "multi")
		}
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

type ApiDojoGroupsUpdateRequest struct {
	ctx context.Context
	ApiService DojoGroupsAPI
	id int32
	dojoGroupRequest *DojoGroupRequest
}

func (r ApiDojoGroupsUpdateRequest) DojoGroupRequest(dojoGroupRequest DojoGroupRequest) ApiDojoGroupsUpdateRequest {
	r.dojoGroupRequest = &dojoGroupRequest
	return r
}

func (r ApiDojoGroupsUpdateRequest) Execute() (*DojoGroup, *http.Response, error) {
	return r.ApiService.DojoGroupsUpdateExecute(r)
}

/*
DojoGroupsUpdate Method for DojoGroupsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this dojo_ group.
 @return ApiDojoGroupsUpdateRequest
*/
func (a *DojoGroupsAPIService) DojoGroupsUpdate(ctx context.Context, id int32) ApiDojoGroupsUpdateRequest {
	return ApiDojoGroupsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DojoGroup
func (a *DojoGroupsAPIService) DojoGroupsUpdateExecute(r ApiDojoGroupsUpdateRequest) (*DojoGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DojoGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DojoGroupsAPIService.DojoGroupsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/dojo_groups/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dojoGroupRequest == nil {
		return localVarReturnValue, nil, reportError("dojoGroupRequest is required and must be specified")
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
	localVarPostBody = r.dojoGroupRequest
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
