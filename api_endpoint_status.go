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


type EndpointStatusAPI interface {

	/*
	EndpointStatusCreate Method for EndpointStatusCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEndpointStatusCreateRequest
	*/
	EndpointStatusCreate(ctx context.Context) ApiEndpointStatusCreateRequest

	// EndpointStatusCreateExecute executes the request
	//  @return EndpointStatus
	EndpointStatusCreateExecute(r ApiEndpointStatusCreateRequest) (*EndpointStatus, *http.Response, error)

	/*
	EndpointStatusDeletePreviewList Method for EndpointStatusDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this endpoint_ status.
	@return ApiEndpointStatusDeletePreviewListRequest
	*/
	EndpointStatusDeletePreviewList(ctx context.Context, id int32) ApiEndpointStatusDeletePreviewListRequest

	// EndpointStatusDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	EndpointStatusDeletePreviewListExecute(r ApiEndpointStatusDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	EndpointStatusDestroy Method for EndpointStatusDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this endpoint_ status.
	@return ApiEndpointStatusDestroyRequest
	*/
	EndpointStatusDestroy(ctx context.Context, id int32) ApiEndpointStatusDestroyRequest

	// EndpointStatusDestroyExecute executes the request
	EndpointStatusDestroyExecute(r ApiEndpointStatusDestroyRequest) (*http.Response, error)

	/*
	EndpointStatusList Method for EndpointStatusList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEndpointStatusListRequest
	*/
	EndpointStatusList(ctx context.Context) ApiEndpointStatusListRequest

	// EndpointStatusListExecute executes the request
	//  @return PaginatedEndpointStatusList
	EndpointStatusListExecute(r ApiEndpointStatusListRequest) (*PaginatedEndpointStatusList, *http.Response, error)

	/*
	EndpointStatusPartialUpdate Method for EndpointStatusPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this endpoint_ status.
	@return ApiEndpointStatusPartialUpdateRequest
	*/
	EndpointStatusPartialUpdate(ctx context.Context, id int32) ApiEndpointStatusPartialUpdateRequest

	// EndpointStatusPartialUpdateExecute executes the request
	//  @return EndpointStatus
	EndpointStatusPartialUpdateExecute(r ApiEndpointStatusPartialUpdateRequest) (*EndpointStatus, *http.Response, error)

	/*
	EndpointStatusRetrieve Method for EndpointStatusRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this endpoint_ status.
	@return ApiEndpointStatusRetrieveRequest
	*/
	EndpointStatusRetrieve(ctx context.Context, id int32) ApiEndpointStatusRetrieveRequest

	// EndpointStatusRetrieveExecute executes the request
	//  @return EndpointStatus
	EndpointStatusRetrieveExecute(r ApiEndpointStatusRetrieveRequest) (*EndpointStatus, *http.Response, error)

	/*
	EndpointStatusUpdate Method for EndpointStatusUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this endpoint_ status.
	@return ApiEndpointStatusUpdateRequest
	*/
	EndpointStatusUpdate(ctx context.Context, id int32) ApiEndpointStatusUpdateRequest

	// EndpointStatusUpdateExecute executes the request
	//  @return EndpointStatus
	EndpointStatusUpdateExecute(r ApiEndpointStatusUpdateRequest) (*EndpointStatus, *http.Response, error)
}

// EndpointStatusAPIService EndpointStatusAPI service
type EndpointStatusAPIService service

type ApiEndpointStatusCreateRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	endpointStatusRequest *EndpointStatusRequest
}

func (r ApiEndpointStatusCreateRequest) EndpointStatusRequest(endpointStatusRequest EndpointStatusRequest) ApiEndpointStatusCreateRequest {
	r.endpointStatusRequest = &endpointStatusRequest
	return r
}

func (r ApiEndpointStatusCreateRequest) Execute() (*EndpointStatus, *http.Response, error) {
	return r.ApiService.EndpointStatusCreateExecute(r)
}

/*
EndpointStatusCreate Method for EndpointStatusCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEndpointStatusCreateRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusCreate(ctx context.Context) ApiEndpointStatusCreateRequest {
	return ApiEndpointStatusCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EndpointStatus
func (a *EndpointStatusAPIService) EndpointStatusCreateExecute(r ApiEndpointStatusCreateRequest) (*EndpointStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndpointStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.endpointStatusRequest == nil {
		return localVarReturnValue, nil, reportError("endpointStatusRequest is required and must be specified")
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
	localVarPostBody = r.endpointStatusRequest
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

type ApiEndpointStatusDeletePreviewListRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiEndpointStatusDeletePreviewListRequest) Limit(limit int32) ApiEndpointStatusDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiEndpointStatusDeletePreviewListRequest) Offset(offset int32) ApiEndpointStatusDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiEndpointStatusDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.EndpointStatusDeletePreviewListExecute(r)
}

/*
EndpointStatusDeletePreviewList Method for EndpointStatusDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this endpoint_ status.
 @return ApiEndpointStatusDeletePreviewListRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusDeletePreviewList(ctx context.Context, id int32) ApiEndpointStatusDeletePreviewListRequest {
	return ApiEndpointStatusDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *EndpointStatusAPIService) EndpointStatusDeletePreviewListExecute(r ApiEndpointStatusDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/{id}/delete_preview/"
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

type ApiEndpointStatusDestroyRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	id int32
}

func (r ApiEndpointStatusDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.EndpointStatusDestroyExecute(r)
}

/*
EndpointStatusDestroy Method for EndpointStatusDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this endpoint_ status.
 @return ApiEndpointStatusDestroyRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusDestroy(ctx context.Context, id int32) ApiEndpointStatusDestroyRequest {
	return ApiEndpointStatusDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *EndpointStatusAPIService) EndpointStatusDestroyExecute(r ApiEndpointStatusDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/{id}/"
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

type ApiEndpointStatusListRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	endpoint *int32
	falsePositive *bool
	finding *int32
	limit *int32
	mitigated *bool
	mitigatedBy *int32
	offset *int32
	outOfScope *bool
	riskAccepted *bool
}

func (r ApiEndpointStatusListRequest) Endpoint(endpoint int32) ApiEndpointStatusListRequest {
	r.endpoint = &endpoint
	return r
}

func (r ApiEndpointStatusListRequest) FalsePositive(falsePositive bool) ApiEndpointStatusListRequest {
	r.falsePositive = &falsePositive
	return r
}

func (r ApiEndpointStatusListRequest) Finding(finding int32) ApiEndpointStatusListRequest {
	r.finding = &finding
	return r
}

// Number of results to return per page.
func (r ApiEndpointStatusListRequest) Limit(limit int32) ApiEndpointStatusListRequest {
	r.limit = &limit
	return r
}

func (r ApiEndpointStatusListRequest) Mitigated(mitigated bool) ApiEndpointStatusListRequest {
	r.mitigated = &mitigated
	return r
}

func (r ApiEndpointStatusListRequest) MitigatedBy(mitigatedBy int32) ApiEndpointStatusListRequest {
	r.mitigatedBy = &mitigatedBy
	return r
}

// The initial index from which to return the results.
func (r ApiEndpointStatusListRequest) Offset(offset int32) ApiEndpointStatusListRequest {
	r.offset = &offset
	return r
}

func (r ApiEndpointStatusListRequest) OutOfScope(outOfScope bool) ApiEndpointStatusListRequest {
	r.outOfScope = &outOfScope
	return r
}

func (r ApiEndpointStatusListRequest) RiskAccepted(riskAccepted bool) ApiEndpointStatusListRequest {
	r.riskAccepted = &riskAccepted
	return r
}

func (r ApiEndpointStatusListRequest) Execute() (*PaginatedEndpointStatusList, *http.Response, error) {
	return r.ApiService.EndpointStatusListExecute(r)
}

/*
EndpointStatusList Method for EndpointStatusList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEndpointStatusListRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusList(ctx context.Context) ApiEndpointStatusListRequest {
	return ApiEndpointStatusListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedEndpointStatusList
func (a *EndpointStatusAPIService) EndpointStatusListExecute(r ApiEndpointStatusListRequest) (*PaginatedEndpointStatusList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedEndpointStatusList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.endpoint != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endpoint", r.endpoint, "form", "")
	}
	if r.falsePositive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "false_positive", r.falsePositive, "form", "")
	}
	if r.finding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "finding", r.finding, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.mitigated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mitigated", r.mitigated, "form", "")
	}
	if r.mitigatedBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mitigated_by", r.mitigatedBy, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.outOfScope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "out_of_scope", r.outOfScope, "form", "")
	}
	if r.riskAccepted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "risk_accepted", r.riskAccepted, "form", "")
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

type ApiEndpointStatusPartialUpdateRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	id int32
	patchedEndpointStatusRequest *PatchedEndpointStatusRequest
}

func (r ApiEndpointStatusPartialUpdateRequest) PatchedEndpointStatusRequest(patchedEndpointStatusRequest PatchedEndpointStatusRequest) ApiEndpointStatusPartialUpdateRequest {
	r.patchedEndpointStatusRequest = &patchedEndpointStatusRequest
	return r
}

func (r ApiEndpointStatusPartialUpdateRequest) Execute() (*EndpointStatus, *http.Response, error) {
	return r.ApiService.EndpointStatusPartialUpdateExecute(r)
}

/*
EndpointStatusPartialUpdate Method for EndpointStatusPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this endpoint_ status.
 @return ApiEndpointStatusPartialUpdateRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusPartialUpdate(ctx context.Context, id int32) ApiEndpointStatusPartialUpdateRequest {
	return ApiEndpointStatusPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EndpointStatus
func (a *EndpointStatusAPIService) EndpointStatusPartialUpdateExecute(r ApiEndpointStatusPartialUpdateRequest) (*EndpointStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndpointStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/{id}/"
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
	localVarPostBody = r.patchedEndpointStatusRequest
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

type ApiEndpointStatusRetrieveRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	id int32
}

func (r ApiEndpointStatusRetrieveRequest) Execute() (*EndpointStatus, *http.Response, error) {
	return r.ApiService.EndpointStatusRetrieveExecute(r)
}

/*
EndpointStatusRetrieve Method for EndpointStatusRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this endpoint_ status.
 @return ApiEndpointStatusRetrieveRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusRetrieve(ctx context.Context, id int32) ApiEndpointStatusRetrieveRequest {
	return ApiEndpointStatusRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EndpointStatus
func (a *EndpointStatusAPIService) EndpointStatusRetrieveExecute(r ApiEndpointStatusRetrieveRequest) (*EndpointStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndpointStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/{id}/"
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

type ApiEndpointStatusUpdateRequest struct {
	ctx context.Context
	ApiService EndpointStatusAPI
	id int32
	endpointStatusRequest *EndpointStatusRequest
}

func (r ApiEndpointStatusUpdateRequest) EndpointStatusRequest(endpointStatusRequest EndpointStatusRequest) ApiEndpointStatusUpdateRequest {
	r.endpointStatusRequest = &endpointStatusRequest
	return r
}

func (r ApiEndpointStatusUpdateRequest) Execute() (*EndpointStatus, *http.Response, error) {
	return r.ApiService.EndpointStatusUpdateExecute(r)
}

/*
EndpointStatusUpdate Method for EndpointStatusUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this endpoint_ status.
 @return ApiEndpointStatusUpdateRequest
*/
func (a *EndpointStatusAPIService) EndpointStatusUpdate(ctx context.Context, id int32) ApiEndpointStatusUpdateRequest {
	return ApiEndpointStatusUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EndpointStatus
func (a *EndpointStatusAPIService) EndpointStatusUpdateExecute(r ApiEndpointStatusUpdateRequest) (*EndpointStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndpointStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointStatusAPIService.EndpointStatusUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_status/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.endpointStatusRequest == nil {
		return localVarReturnValue, nil, reportError("endpointStatusRequest is required and must be specified")
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
	localVarPostBody = r.endpointStatusRequest
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
