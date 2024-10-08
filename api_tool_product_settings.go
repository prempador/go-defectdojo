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
	"reflect"
)


type ToolProductSettingsAPI interface {

	/*
	ToolProductSettingsCreate Method for ToolProductSettingsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiToolProductSettingsCreateRequest
	*/
	ToolProductSettingsCreate(ctx context.Context) ApiToolProductSettingsCreateRequest

	// ToolProductSettingsCreateExecute executes the request
	//  @return ToolProductSettings
	ToolProductSettingsCreateExecute(r ApiToolProductSettingsCreateRequest) (*ToolProductSettings, *http.Response, error)

	/*
	ToolProductSettingsDeletePreviewList Method for ToolProductSettingsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ product_ settings.
	@return ApiToolProductSettingsDeletePreviewListRequest
	*/
	ToolProductSettingsDeletePreviewList(ctx context.Context, id int32) ApiToolProductSettingsDeletePreviewListRequest

	// ToolProductSettingsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	ToolProductSettingsDeletePreviewListExecute(r ApiToolProductSettingsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	ToolProductSettingsDestroy Method for ToolProductSettingsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ product_ settings.
	@return ApiToolProductSettingsDestroyRequest
	*/
	ToolProductSettingsDestroy(ctx context.Context, id int32) ApiToolProductSettingsDestroyRequest

	// ToolProductSettingsDestroyExecute executes the request
	ToolProductSettingsDestroyExecute(r ApiToolProductSettingsDestroyRequest) (*http.Response, error)

	/*
	ToolProductSettingsList Method for ToolProductSettingsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiToolProductSettingsListRequest
	*/
	ToolProductSettingsList(ctx context.Context) ApiToolProductSettingsListRequest

	// ToolProductSettingsListExecute executes the request
	//  @return PaginatedToolProductSettingsList
	ToolProductSettingsListExecute(r ApiToolProductSettingsListRequest) (*PaginatedToolProductSettingsList, *http.Response, error)

	/*
	ToolProductSettingsPartialUpdate Method for ToolProductSettingsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ product_ settings.
	@return ApiToolProductSettingsPartialUpdateRequest
	*/
	ToolProductSettingsPartialUpdate(ctx context.Context, id int32) ApiToolProductSettingsPartialUpdateRequest

	// ToolProductSettingsPartialUpdateExecute executes the request
	//  @return ToolProductSettings
	ToolProductSettingsPartialUpdateExecute(r ApiToolProductSettingsPartialUpdateRequest) (*ToolProductSettings, *http.Response, error)

	/*
	ToolProductSettingsRetrieve Method for ToolProductSettingsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ product_ settings.
	@return ApiToolProductSettingsRetrieveRequest
	*/
	ToolProductSettingsRetrieve(ctx context.Context, id int32) ApiToolProductSettingsRetrieveRequest

	// ToolProductSettingsRetrieveExecute executes the request
	//  @return ToolProductSettings
	ToolProductSettingsRetrieveExecute(r ApiToolProductSettingsRetrieveRequest) (*ToolProductSettings, *http.Response, error)

	/*
	ToolProductSettingsUpdate Method for ToolProductSettingsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this tool_ product_ settings.
	@return ApiToolProductSettingsUpdateRequest
	*/
	ToolProductSettingsUpdate(ctx context.Context, id int32) ApiToolProductSettingsUpdateRequest

	// ToolProductSettingsUpdateExecute executes the request
	//  @return ToolProductSettings
	ToolProductSettingsUpdateExecute(r ApiToolProductSettingsUpdateRequest) (*ToolProductSettings, *http.Response, error)
}

// ToolProductSettingsAPIService ToolProductSettingsAPI service
type ToolProductSettingsAPIService service

type ApiToolProductSettingsCreateRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	toolProductSettingsRequest *ToolProductSettingsRequest
}

func (r ApiToolProductSettingsCreateRequest) ToolProductSettingsRequest(toolProductSettingsRequest ToolProductSettingsRequest) ApiToolProductSettingsCreateRequest {
	r.toolProductSettingsRequest = &toolProductSettingsRequest
	return r
}

func (r ApiToolProductSettingsCreateRequest) Execute() (*ToolProductSettings, *http.Response, error) {
	return r.ApiService.ToolProductSettingsCreateExecute(r)
}

/*
ToolProductSettingsCreate Method for ToolProductSettingsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolProductSettingsCreateRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsCreate(ctx context.Context) ApiToolProductSettingsCreateRequest {
	return ApiToolProductSettingsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ToolProductSettings
func (a *ToolProductSettingsAPIService) ToolProductSettingsCreateExecute(r ApiToolProductSettingsCreateRequest) (*ToolProductSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolProductSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolProductSettingsRequest == nil {
		return localVarReturnValue, nil, reportError("toolProductSettingsRequest is required and must be specified")
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
	localVarPostBody = r.toolProductSettingsRequest
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

type ApiToolProductSettingsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiToolProductSettingsDeletePreviewListRequest) Limit(limit int32) ApiToolProductSettingsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiToolProductSettingsDeletePreviewListRequest) Offset(offset int32) ApiToolProductSettingsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiToolProductSettingsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.ToolProductSettingsDeletePreviewListExecute(r)
}

/*
ToolProductSettingsDeletePreviewList Method for ToolProductSettingsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ product_ settings.
 @return ApiToolProductSettingsDeletePreviewListRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsDeletePreviewList(ctx context.Context, id int32) ApiToolProductSettingsDeletePreviewListRequest {
	return ApiToolProductSettingsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *ToolProductSettingsAPIService) ToolProductSettingsDeletePreviewListExecute(r ApiToolProductSettingsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/{id}/delete_preview/"
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

type ApiToolProductSettingsDestroyRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	id int32
}

func (r ApiToolProductSettingsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.ToolProductSettingsDestroyExecute(r)
}

/*
ToolProductSettingsDestroy Method for ToolProductSettingsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ product_ settings.
 @return ApiToolProductSettingsDestroyRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsDestroy(ctx context.Context, id int32) ApiToolProductSettingsDestroyRequest {
	return ApiToolProductSettingsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ToolProductSettingsAPIService) ToolProductSettingsDestroyExecute(r ApiToolProductSettingsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/{id}/"
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

type ApiToolProductSettingsListRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	id *int32
	limit *int32
	name *string
	offset *int32
	prefetch *[]string
	product *int32
	toolConfiguration *int32
	toolProjectId *string
	url *string
}

func (r ApiToolProductSettingsListRequest) Id(id int32) ApiToolProductSettingsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiToolProductSettingsListRequest) Limit(limit int32) ApiToolProductSettingsListRequest {
	r.limit = &limit
	return r
}

func (r ApiToolProductSettingsListRequest) Name(name string) ApiToolProductSettingsListRequest {
	r.name = &name
	return r
}

// The initial index from which to return the results.
func (r ApiToolProductSettingsListRequest) Offset(offset int32) ApiToolProductSettingsListRequest {
	r.offset = &offset
	return r
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiToolProductSettingsListRequest) Prefetch(prefetch []string) ApiToolProductSettingsListRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiToolProductSettingsListRequest) Product(product int32) ApiToolProductSettingsListRequest {
	r.product = &product
	return r
}

func (r ApiToolProductSettingsListRequest) ToolConfiguration(toolConfiguration int32) ApiToolProductSettingsListRequest {
	r.toolConfiguration = &toolConfiguration
	return r
}

func (r ApiToolProductSettingsListRequest) ToolProjectId(toolProjectId string) ApiToolProductSettingsListRequest {
	r.toolProjectId = &toolProjectId
	return r
}

func (r ApiToolProductSettingsListRequest) Url(url string) ApiToolProductSettingsListRequest {
	r.url = &url
	return r
}

func (r ApiToolProductSettingsListRequest) Execute() (*PaginatedToolProductSettingsList, *http.Response, error) {
	return r.ApiService.ToolProductSettingsListExecute(r)
}

/*
ToolProductSettingsList Method for ToolProductSettingsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiToolProductSettingsListRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsList(ctx context.Context) ApiToolProductSettingsListRequest {
	return ApiToolProductSettingsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedToolProductSettingsList
func (a *ToolProductSettingsAPIService) ToolProductSettingsListExecute(r ApiToolProductSettingsListRequest) (*PaginatedToolProductSettingsList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedToolProductSettingsList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
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
	if r.prefetch != nil {
		t := *r.prefetch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", t, "form", "multi")
		}
	}
	if r.product != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product", r.product, "form", "")
	}
	if r.toolConfiguration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tool_configuration", r.toolConfiguration, "form", "")
	}
	if r.toolProjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tool_project_id", r.toolProjectId, "form", "")
	}
	if r.url != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "form", "")
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

type ApiToolProductSettingsPartialUpdateRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	id int32
	patchedToolProductSettingsRequest *PatchedToolProductSettingsRequest
}

func (r ApiToolProductSettingsPartialUpdateRequest) PatchedToolProductSettingsRequest(patchedToolProductSettingsRequest PatchedToolProductSettingsRequest) ApiToolProductSettingsPartialUpdateRequest {
	r.patchedToolProductSettingsRequest = &patchedToolProductSettingsRequest
	return r
}

func (r ApiToolProductSettingsPartialUpdateRequest) Execute() (*ToolProductSettings, *http.Response, error) {
	return r.ApiService.ToolProductSettingsPartialUpdateExecute(r)
}

/*
ToolProductSettingsPartialUpdate Method for ToolProductSettingsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ product_ settings.
 @return ApiToolProductSettingsPartialUpdateRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsPartialUpdate(ctx context.Context, id int32) ApiToolProductSettingsPartialUpdateRequest {
	return ApiToolProductSettingsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ToolProductSettings
func (a *ToolProductSettingsAPIService) ToolProductSettingsPartialUpdateExecute(r ApiToolProductSettingsPartialUpdateRequest) (*ToolProductSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolProductSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/{id}/"
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
	localVarPostBody = r.patchedToolProductSettingsRequest
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

type ApiToolProductSettingsRetrieveRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	id int32
	prefetch *[]string
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiToolProductSettingsRetrieveRequest) Prefetch(prefetch []string) ApiToolProductSettingsRetrieveRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiToolProductSettingsRetrieveRequest) Execute() (*ToolProductSettings, *http.Response, error) {
	return r.ApiService.ToolProductSettingsRetrieveExecute(r)
}

/*
ToolProductSettingsRetrieve Method for ToolProductSettingsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ product_ settings.
 @return ApiToolProductSettingsRetrieveRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsRetrieve(ctx context.Context, id int32) ApiToolProductSettingsRetrieveRequest {
	return ApiToolProductSettingsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ToolProductSettings
func (a *ToolProductSettingsAPIService) ToolProductSettingsRetrieveExecute(r ApiToolProductSettingsRetrieveRequest) (*ToolProductSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolProductSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.prefetch != nil {
		t := *r.prefetch
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "prefetch", t, "form", "multi")
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

type ApiToolProductSettingsUpdateRequest struct {
	ctx context.Context
	ApiService ToolProductSettingsAPI
	id int32
	toolProductSettingsRequest *ToolProductSettingsRequest
}

func (r ApiToolProductSettingsUpdateRequest) ToolProductSettingsRequest(toolProductSettingsRequest ToolProductSettingsRequest) ApiToolProductSettingsUpdateRequest {
	r.toolProductSettingsRequest = &toolProductSettingsRequest
	return r
}

func (r ApiToolProductSettingsUpdateRequest) Execute() (*ToolProductSettings, *http.Response, error) {
	return r.ApiService.ToolProductSettingsUpdateExecute(r)
}

/*
ToolProductSettingsUpdate Method for ToolProductSettingsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this tool_ product_ settings.
 @return ApiToolProductSettingsUpdateRequest
*/
func (a *ToolProductSettingsAPIService) ToolProductSettingsUpdate(ctx context.Context, id int32) ApiToolProductSettingsUpdateRequest {
	return ApiToolProductSettingsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ToolProductSettings
func (a *ToolProductSettingsAPIService) ToolProductSettingsUpdateExecute(r ApiToolProductSettingsUpdateRequest) (*ToolProductSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ToolProductSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToolProductSettingsAPIService.ToolProductSettingsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tool_product_settings/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.toolProductSettingsRequest == nil {
		return localVarReturnValue, nil, reportError("toolProductSettingsRequest is required and must be specified")
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
	localVarPostBody = r.toolProductSettingsRequest
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
