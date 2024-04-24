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


type ToolConfigurationsAPI interface {

	/*
	ToolConfigurationsCreate Method for ToolConfigurationsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiToolConfigurationsCreateRequest
	*/
	ToolConfigurationsCreate(ctx context.Context) ApiToolConfigurationsCreateRequest

	// ToolConfigurationsCreateExecute executes the request
	//  @return ToolConfiguration
	ToolConfigurationsCreateExecute(r ApiToolConfigurationsCreateRequest) (*ToolConfiguration, *http.Response, error)

	/*
	ToolConfigurationsDeletePreviewList Method for ToolConfigurationsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ configuration.
	@return ApiToolConfigurationsDeletePreviewListRequest
	*/
	ToolConfigurationsDeletePreviewList(ctx context.Context, id int32) ApiToolConfigurationsDeletePreviewListRequest

	// ToolConfigurationsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	ToolConfigurationsDeletePreviewListExecute(r ApiToolConfigurationsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	ToolConfigurationsDestroy Method for ToolConfigurationsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ configuration.
	@return ApiToolConfigurationsDestroyRequest
	*/
	ToolConfigurationsDestroy(ctx context.Context, id int32) ApiToolConfigurationsDestroyRequest

	// ToolConfigurationsDestroyExecute executes the request
	ToolConfigurationsDestroyExecute(r ApiToolConfigurationsDestroyRequest) (*http.Response, error)

	/*
	ToolConfigurationsList Method for ToolConfigurationsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiToolConfigurationsListRequest
	*/
	ToolConfigurationsList(ctx context.Context) ApiToolConfigurationsListRequest

	// ToolConfigurationsListExecute executes the request
	//  @return PaginatedToolConfigurationList
	ToolConfigurationsListExecute(r ApiToolConfigurationsListRequest) (*PaginatedToolConfigurationList, *http.Response, error)

	/*
	ToolConfigurationsPartialUpdate Method for ToolConfigurationsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ configuration.
	@return ApiToolConfigurationsPartialUpdateRequest
	*/
	ToolConfigurationsPartialUpdate(ctx context.Context, id int32) ApiToolConfigurationsPartialUpdateRequest

	// ToolConfigurationsPartialUpdateExecute executes the request
	//  @return ToolConfiguration
	ToolConfigurationsPartialUpdateExecute(r ApiToolConfigurationsPartialUpdateRequest) (*ToolConfiguration, *http.Response, error)

	/*
	ToolConfigurationsRetrieve Method for ToolConfigurationsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ configuration.
	@return ApiToolConfigurationsRetrieveRequest
	*/
	ToolConfigurationsRetrieve(ctx context.Context, id int32) ApiToolConfigurationsRetrieveRequest

	// ToolConfigurationsRetrieveExecute executes the request
	//  @return ToolConfiguration
	ToolConfigurationsRetrieveExecute(r ApiToolConfigurationsRetrieveRequest) (*ToolConfiguration, *http.Response, error)

	/*
	ToolConfigurationsUpdate Method for ToolConfigurationsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ configuration.
	@return ApiToolConfigurationsUpdateRequest
	*/
	ToolConfigurationsUpdate(ctx context.Context, id int32) ApiToolConfigurationsUpdateRequest

	// ToolConfigurationsUpdateExecute executes the request
	//  @return ToolConfiguration
	ToolConfigurationsUpdateExecute(r ApiToolConfigurationsUpdateRequest) (*ToolConfiguration, *http.Response, error)
}

// ToolConfigurationsAPIService ToolConfigurationsAPI service
type ToolConfigurationsAPIService service

type ApiToolConfigurationsCreateRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	toolConfigurationRequest *ToolConfigurationRequest
}

func (r ApiToolConfigurationsCreateRequest) ToolConfigurationRequest(toolConfigurationRequest ToolConfigurationRequest) ApiToolConfigurationsCreateRequest {
	r.toolConfigurationRequest = &toolConfigurationRequest
	return r
}

func (r ApiToolConfigurationsCreateRequest) Execute() (*ToolConfiguration, *http.Response, error) {
	return r.ApiService.ToolConfigurationsCreateExecute(r)
}

/*
ToolConfigurationsCreate Method for ToolConfigurationsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolConfigurationsCreateRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsCreate(ctx context.Context) ApiToolConfigurationsCreateRequest {
	return ApiToolConfigurationsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolConfiguration
func (a *ToolConfigurationsAPIService) ToolConfigurationsCreateExecute(r ApiToolConfigurationsCreateRequest) (*ToolConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("toolConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.toolConfigurationRequest
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

type ApiToolConfigurationsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiToolConfigurationsDeletePreviewListRequest) Limit(limit int32) ApiToolConfigurationsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiToolConfigurationsDeletePreviewListRequest) Offset(offset int32) ApiToolConfigurationsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiToolConfigurationsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.ToolConfigurationsDeletePreviewListExecute(r)
}

/*
ToolConfigurationsDeletePreviewList Method for ToolConfigurationsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ configuration.
 @return ApiToolConfigurationsDeletePreviewListRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsDeletePreviewList(ctx context.Context, id int32) ApiToolConfigurationsDeletePreviewListRequest {
	return ApiToolConfigurationsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *ToolConfigurationsAPIService) ToolConfigurationsDeletePreviewListExecute(r ApiToolConfigurationsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/{id}/delete_preview/"
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

type ApiToolConfigurationsDestroyRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	id int32
}

func (r ApiToolConfigurationsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.ToolConfigurationsDestroyExecute(r)
}

/*
ToolConfigurationsDestroy Method for ToolConfigurationsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ configuration.
 @return ApiToolConfigurationsDestroyRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsDestroy(ctx context.Context, id int32) ApiToolConfigurationsDestroyRequest {
	return ApiToolConfigurationsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ToolConfigurationsAPIService) ToolConfigurationsDestroyExecute(r ApiToolConfigurationsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/{id}/"
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

type ApiToolConfigurationsListRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	authenticationType *string
	id *int32
	limit *int32
	name *string
	offset *int32
	toolType *int32
	url *string
}

// * &#x60;API&#x60; - API Key * &#x60;Password&#x60; - Username/Password * &#x60;SSH&#x60; - SSH
func (r ApiToolConfigurationsListRequest) AuthenticationType(authenticationType string) ApiToolConfigurationsListRequest {
	r.authenticationType = &authenticationType
	return r
}

func (r ApiToolConfigurationsListRequest) Id(id int32) ApiToolConfigurationsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiToolConfigurationsListRequest) Limit(limit int32) ApiToolConfigurationsListRequest {
	r.limit = &limit
	return r
}

func (r ApiToolConfigurationsListRequest) Name(name string) ApiToolConfigurationsListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r ApiToolConfigurationsListRequest) Offset(offset int32) ApiToolConfigurationsListRequest {
	r.offset = &offset
	return r
}

func (r ApiToolConfigurationsListRequest) ToolType(toolType int32) ApiToolConfigurationsListRequest {
	r.toolType = &toolType
	return r
}

func (r ApiToolConfigurationsListRequest) Url(url string) ApiToolConfigurationsListRequest {
	r.url = &url
	return r
}

func (r ApiToolConfigurationsListRequest) Execute() (*PaginatedToolConfigurationList, *http.Response, error) {
	return r.ApiService.ToolConfigurationsListExecute(r)
}

/*
ToolConfigurationsList Method for ToolConfigurationsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolConfigurationsListRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsList(ctx context.Context) ApiToolConfigurationsListRequest {
	return ApiToolConfigurationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedToolConfigurationList
func (a *ToolConfigurationsAPIService) ToolConfigurationsListExecute(r ApiToolConfigurationsListRequest) (*PaginatedToolConfigurationList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedToolConfigurationList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.authenticationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "authentication_type", r.authenticationType, "")
	}
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
	if r.toolType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tool_type", r.toolType, "")
	}
	if r.url != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "")
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

type ApiToolConfigurationsPartialUpdateRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	id int32
	patchedToolConfigurationRequest *PatchedToolConfigurationRequest
}

func (r ApiToolConfigurationsPartialUpdateRequest) PatchedToolConfigurationRequest(patchedToolConfigurationRequest PatchedToolConfigurationRequest) ApiToolConfigurationsPartialUpdateRequest {
	r.patchedToolConfigurationRequest = &patchedToolConfigurationRequest
	return r
}

func (r ApiToolConfigurationsPartialUpdateRequest) Execute() (*ToolConfiguration, *http.Response, error) {
	return r.ApiService.ToolConfigurationsPartialUpdateExecute(r)
}

/*
ToolConfigurationsPartialUpdate Method for ToolConfigurationsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ configuration.
 @return ApiToolConfigurationsPartialUpdateRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsPartialUpdate(ctx context.Context, id int32) ApiToolConfigurationsPartialUpdateRequest {
	return ApiToolConfigurationsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ToolConfiguration
func (a *ToolConfigurationsAPIService) ToolConfigurationsPartialUpdateExecute(r ApiToolConfigurationsPartialUpdateRequest) (*ToolConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/{id}/"
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
	localVarPostBody = r.patchedToolConfigurationRequest
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

type ApiToolConfigurationsRetrieveRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	id int32
}

func (r ApiToolConfigurationsRetrieveRequest) Execute() (*ToolConfiguration, *http.Response, error) {
	return r.ApiService.ToolConfigurationsRetrieveExecute(r)
}

/*
ToolConfigurationsRetrieve Method for ToolConfigurationsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ configuration.
 @return ApiToolConfigurationsRetrieveRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsRetrieve(ctx context.Context, id int32) ApiToolConfigurationsRetrieveRequest {
	return ApiToolConfigurationsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ToolConfiguration
func (a *ToolConfigurationsAPIService) ToolConfigurationsRetrieveExecute(r ApiToolConfigurationsRetrieveRequest) (*ToolConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/{id}/"
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

type ApiToolConfigurationsUpdateRequest struct {
	ctx context.Context
	ApiService ToolConfigurationsAPI
	id int32
	toolConfigurationRequest *ToolConfigurationRequest
}

func (r ApiToolConfigurationsUpdateRequest) ToolConfigurationRequest(toolConfigurationRequest ToolConfigurationRequest) ApiToolConfigurationsUpdateRequest {
	r.toolConfigurationRequest = &toolConfigurationRequest
	return r
}

func (r ApiToolConfigurationsUpdateRequest) Execute() (*ToolConfiguration, *http.Response, error) {
	return r.ApiService.ToolConfigurationsUpdateExecute(r)
}

/*
ToolConfigurationsUpdate Method for ToolConfigurationsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ configuration.
 @return ApiToolConfigurationsUpdateRequest
*/
func (a *ToolConfigurationsAPIService) ToolConfigurationsUpdate(ctx context.Context, id int32) ApiToolConfigurationsUpdateRequest {
	return ApiToolConfigurationsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ToolConfiguration
func (a *ToolConfigurationsAPIService) ToolConfigurationsUpdateExecute(r ApiToolConfigurationsUpdateRequest) (*ToolConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolConfigurationsAPIService.ToolConfigurationsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_configurations/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("toolConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.toolConfigurationRequest
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
