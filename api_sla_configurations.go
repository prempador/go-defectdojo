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


type SlaConfigurationsAPI interface {

	/*
	SlaConfigurationsCreate Method for SlaConfigurationsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSlaConfigurationsCreateRequest
	*/
	SlaConfigurationsCreate(ctx context.Context) ApiSlaConfigurationsCreateRequest

	// SlaConfigurationsCreateExecute executes the request
	//  @return SLAConfiguration
	SlaConfigurationsCreateExecute(r ApiSlaConfigurationsCreateRequest) (*SLAConfiguration, *http.Response, error)

	/*
	SlaConfigurationsDeletePreviewList Method for SlaConfigurationsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sl a_ configuration.
	@return ApiSlaConfigurationsDeletePreviewListRequest
	*/
	SlaConfigurationsDeletePreviewList(ctx context.Context, id int32) ApiSlaConfigurationsDeletePreviewListRequest

	// SlaConfigurationsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	SlaConfigurationsDeletePreviewListExecute(r ApiSlaConfigurationsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	SlaConfigurationsDestroy Method for SlaConfigurationsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sl a_ configuration.
	@return ApiSlaConfigurationsDestroyRequest
	*/
	SlaConfigurationsDestroy(ctx context.Context, id int32) ApiSlaConfigurationsDestroyRequest

	// SlaConfigurationsDestroyExecute executes the request
	SlaConfigurationsDestroyExecute(r ApiSlaConfigurationsDestroyRequest) (*http.Response, error)

	/*
	SlaConfigurationsList Method for SlaConfigurationsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSlaConfigurationsListRequest
	*/
	SlaConfigurationsList(ctx context.Context) ApiSlaConfigurationsListRequest

	// SlaConfigurationsListExecute executes the request
	//  @return PaginatedSLAConfigurationList
	SlaConfigurationsListExecute(r ApiSlaConfigurationsListRequest) (*PaginatedSLAConfigurationList, *http.Response, error)

	/*
	SlaConfigurationsPartialUpdate Method for SlaConfigurationsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sl a_ configuration.
	@return ApiSlaConfigurationsPartialUpdateRequest
	*/
	SlaConfigurationsPartialUpdate(ctx context.Context, id int32) ApiSlaConfigurationsPartialUpdateRequest

	// SlaConfigurationsPartialUpdateExecute executes the request
	//  @return SLAConfiguration
	SlaConfigurationsPartialUpdateExecute(r ApiSlaConfigurationsPartialUpdateRequest) (*SLAConfiguration, *http.Response, error)

	/*
	SlaConfigurationsRetrieve Method for SlaConfigurationsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sl a_ configuration.
	@return ApiSlaConfigurationsRetrieveRequest
	*/
	SlaConfigurationsRetrieve(ctx context.Context, id int32) ApiSlaConfigurationsRetrieveRequest

	// SlaConfigurationsRetrieveExecute executes the request
	//  @return SLAConfiguration
	SlaConfigurationsRetrieveExecute(r ApiSlaConfigurationsRetrieveRequest) (*SLAConfiguration, *http.Response, error)

	/*
	SlaConfigurationsUpdate Method for SlaConfigurationsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this sl a_ configuration.
	@return ApiSlaConfigurationsUpdateRequest
	*/
	SlaConfigurationsUpdate(ctx context.Context, id int32) ApiSlaConfigurationsUpdateRequest

	// SlaConfigurationsUpdateExecute executes the request
	//  @return SLAConfiguration
	SlaConfigurationsUpdateExecute(r ApiSlaConfigurationsUpdateRequest) (*SLAConfiguration, *http.Response, error)
}

// SlaConfigurationsAPIService SlaConfigurationsAPI service
type SlaConfigurationsAPIService service

type ApiSlaConfigurationsCreateRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	sLAConfigurationRequest *SLAConfigurationRequest
}

func (r ApiSlaConfigurationsCreateRequest) SLAConfigurationRequest(sLAConfigurationRequest SLAConfigurationRequest) ApiSlaConfigurationsCreateRequest {
	r.sLAConfigurationRequest = &sLAConfigurationRequest
	return r
}

func (r ApiSlaConfigurationsCreateRequest) Execute() (*SLAConfiguration, *http.Response, error) {
	return r.ApiService.SlaConfigurationsCreateExecute(r)
}

/*
SlaConfigurationsCreate Method for SlaConfigurationsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSlaConfigurationsCreateRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsCreate(ctx context.Context) ApiSlaConfigurationsCreateRequest {
	return ApiSlaConfigurationsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SLAConfiguration
func (a *SlaConfigurationsAPIService) SlaConfigurationsCreateExecute(r ApiSlaConfigurationsCreateRequest) (*SLAConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SLAConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sLAConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("sLAConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.sLAConfigurationRequest
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

type ApiSlaConfigurationsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiSlaConfigurationsDeletePreviewListRequest) Limit(limit int32) ApiSlaConfigurationsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiSlaConfigurationsDeletePreviewListRequest) Offset(offset int32) ApiSlaConfigurationsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiSlaConfigurationsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.SlaConfigurationsDeletePreviewListExecute(r)
}

/*
SlaConfigurationsDeletePreviewList Method for SlaConfigurationsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sl a_ configuration.
 @return ApiSlaConfigurationsDeletePreviewListRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsDeletePreviewList(ctx context.Context, id int32) ApiSlaConfigurationsDeletePreviewListRequest {
	return ApiSlaConfigurationsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *SlaConfigurationsAPIService) SlaConfigurationsDeletePreviewListExecute(r ApiSlaConfigurationsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/{id}/delete_preview/"
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

type ApiSlaConfigurationsDestroyRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	id int32
}

func (r ApiSlaConfigurationsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.SlaConfigurationsDestroyExecute(r)
}

/*
SlaConfigurationsDestroy Method for SlaConfigurationsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sl a_ configuration.
 @return ApiSlaConfigurationsDestroyRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsDestroy(ctx context.Context, id int32) ApiSlaConfigurationsDestroyRequest {
	return ApiSlaConfigurationsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SlaConfigurationsAPIService) SlaConfigurationsDestroyExecute(r ApiSlaConfigurationsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/{id}/"
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

type ApiSlaConfigurationsListRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiSlaConfigurationsListRequest) Limit(limit int32) ApiSlaConfigurationsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiSlaConfigurationsListRequest) Offset(offset int32) ApiSlaConfigurationsListRequest {
	r.offset = &offset
	return r
}

func (r ApiSlaConfigurationsListRequest) Execute() (*PaginatedSLAConfigurationList, *http.Response, error) {
	return r.ApiService.SlaConfigurationsListExecute(r)
}

/*
SlaConfigurationsList Method for SlaConfigurationsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSlaConfigurationsListRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsList(ctx context.Context) ApiSlaConfigurationsListRequest {
	return ApiSlaConfigurationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedSLAConfigurationList
func (a *SlaConfigurationsAPIService) SlaConfigurationsListExecute(r ApiSlaConfigurationsListRequest) (*PaginatedSLAConfigurationList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedSLAConfigurationList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/"

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

type ApiSlaConfigurationsPartialUpdateRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	id int32
	patchedSLAConfigurationRequest *PatchedSLAConfigurationRequest
}

func (r ApiSlaConfigurationsPartialUpdateRequest) PatchedSLAConfigurationRequest(patchedSLAConfigurationRequest PatchedSLAConfigurationRequest) ApiSlaConfigurationsPartialUpdateRequest {
	r.patchedSLAConfigurationRequest = &patchedSLAConfigurationRequest
	return r
}

func (r ApiSlaConfigurationsPartialUpdateRequest) Execute() (*SLAConfiguration, *http.Response, error) {
	return r.ApiService.SlaConfigurationsPartialUpdateExecute(r)
}

/*
SlaConfigurationsPartialUpdate Method for SlaConfigurationsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sl a_ configuration.
 @return ApiSlaConfigurationsPartialUpdateRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsPartialUpdate(ctx context.Context, id int32) ApiSlaConfigurationsPartialUpdateRequest {
	return ApiSlaConfigurationsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SLAConfiguration
func (a *SlaConfigurationsAPIService) SlaConfigurationsPartialUpdateExecute(r ApiSlaConfigurationsPartialUpdateRequest) (*SLAConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SLAConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/{id}/"
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
	localVarPostBody = r.patchedSLAConfigurationRequest
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

type ApiSlaConfigurationsRetrieveRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	id int32
}

func (r ApiSlaConfigurationsRetrieveRequest) Execute() (*SLAConfiguration, *http.Response, error) {
	return r.ApiService.SlaConfigurationsRetrieveExecute(r)
}

/*
SlaConfigurationsRetrieve Method for SlaConfigurationsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sl a_ configuration.
 @return ApiSlaConfigurationsRetrieveRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsRetrieve(ctx context.Context, id int32) ApiSlaConfigurationsRetrieveRequest {
	return ApiSlaConfigurationsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SLAConfiguration
func (a *SlaConfigurationsAPIService) SlaConfigurationsRetrieveExecute(r ApiSlaConfigurationsRetrieveRequest) (*SLAConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SLAConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/{id}/"
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

type ApiSlaConfigurationsUpdateRequest struct {
	ctx context.Context
	ApiService SlaConfigurationsAPI
	id int32
	sLAConfigurationRequest *SLAConfigurationRequest
}

func (r ApiSlaConfigurationsUpdateRequest) SLAConfigurationRequest(sLAConfigurationRequest SLAConfigurationRequest) ApiSlaConfigurationsUpdateRequest {
	r.sLAConfigurationRequest = &sLAConfigurationRequest
	return r
}

func (r ApiSlaConfigurationsUpdateRequest) Execute() (*SLAConfiguration, *http.Response, error) {
	return r.ApiService.SlaConfigurationsUpdateExecute(r)
}

/*
SlaConfigurationsUpdate Method for SlaConfigurationsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this sl a_ configuration.
 @return ApiSlaConfigurationsUpdateRequest
*/
func (a *SlaConfigurationsAPIService) SlaConfigurationsUpdate(ctx context.Context, id int32) ApiSlaConfigurationsUpdateRequest {
	return ApiSlaConfigurationsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SLAConfiguration
func (a *SlaConfigurationsAPIService) SlaConfigurationsUpdateExecute(r ApiSlaConfigurationsUpdateRequest) (*SLAConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SLAConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SlaConfigurationsAPIService.SlaConfigurationsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/sla_configurations/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sLAConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("sLAConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.sLAConfigurationRequest
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
