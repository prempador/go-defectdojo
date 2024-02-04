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


type CredentialMappingsAPI interface {

	/*
	CredentialMappingsCreate Method for CredentialMappingsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCredentialMappingsCreateRequest
	*/
	CredentialMappingsCreate(ctx context.Context) ApiCredentialMappingsCreateRequest

	// CredentialMappingsCreateExecute executes the request
	//  @return CredentialMapping
	CredentialMappingsCreateExecute(r ApiCredentialMappingsCreateRequest) (*CredentialMapping, *http.Response, error)

	/*
	CredentialMappingsDeletePreviewList Method for CredentialMappingsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this cred_ mapping.
	@return ApiCredentialMappingsDeletePreviewListRequest
	*/
	CredentialMappingsDeletePreviewList(ctx context.Context, id int32) ApiCredentialMappingsDeletePreviewListRequest

	// CredentialMappingsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	CredentialMappingsDeletePreviewListExecute(r ApiCredentialMappingsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	CredentialMappingsDestroy Method for CredentialMappingsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this cred_ mapping.
	@return ApiCredentialMappingsDestroyRequest
	*/
	CredentialMappingsDestroy(ctx context.Context, id int32) ApiCredentialMappingsDestroyRequest

	// CredentialMappingsDestroyExecute executes the request
	CredentialMappingsDestroyExecute(r ApiCredentialMappingsDestroyRequest) (*http.Response, error)

	/*
	CredentialMappingsList Method for CredentialMappingsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCredentialMappingsListRequest
	*/
	CredentialMappingsList(ctx context.Context) ApiCredentialMappingsListRequest

	// CredentialMappingsListExecute executes the request
	//  @return PaginatedCredentialMappingList
	CredentialMappingsListExecute(r ApiCredentialMappingsListRequest) (*PaginatedCredentialMappingList, *http.Response, error)

	/*
	CredentialMappingsPartialUpdate Method for CredentialMappingsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this cred_ mapping.
	@return ApiCredentialMappingsPartialUpdateRequest
	*/
	CredentialMappingsPartialUpdate(ctx context.Context, id int32) ApiCredentialMappingsPartialUpdateRequest

	// CredentialMappingsPartialUpdateExecute executes the request
	//  @return CredentialMapping
	CredentialMappingsPartialUpdateExecute(r ApiCredentialMappingsPartialUpdateRequest) (*CredentialMapping, *http.Response, error)

	/*
	CredentialMappingsRetrieve Method for CredentialMappingsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this cred_ mapping.
	@return ApiCredentialMappingsRetrieveRequest
	*/
	CredentialMappingsRetrieve(ctx context.Context, id int32) ApiCredentialMappingsRetrieveRequest

	// CredentialMappingsRetrieveExecute executes the request
	//  @return CredentialMapping
	CredentialMappingsRetrieveExecute(r ApiCredentialMappingsRetrieveRequest) (*CredentialMapping, *http.Response, error)

	/*
	CredentialMappingsUpdate Method for CredentialMappingsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this cred_ mapping.
	@return ApiCredentialMappingsUpdateRequest
	*/
	CredentialMappingsUpdate(ctx context.Context, id int32) ApiCredentialMappingsUpdateRequest

	// CredentialMappingsUpdateExecute executes the request
	//  @return CredentialMapping
	CredentialMappingsUpdateExecute(r ApiCredentialMappingsUpdateRequest) (*CredentialMapping, *http.Response, error)
}

// CredentialMappingsAPIService CredentialMappingsAPI service
type CredentialMappingsAPIService service

type ApiCredentialMappingsCreateRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	credentialMappingRequest *CredentialMappingRequest
}

func (r ApiCredentialMappingsCreateRequest) CredentialMappingRequest(credentialMappingRequest CredentialMappingRequest) ApiCredentialMappingsCreateRequest {
	r.credentialMappingRequest = &credentialMappingRequest
	return r
}

func (r ApiCredentialMappingsCreateRequest) Execute() (*CredentialMapping, *http.Response, error) {
	return r.ApiService.CredentialMappingsCreateExecute(r)
}

/*
CredentialMappingsCreate Method for CredentialMappingsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCredentialMappingsCreateRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsCreate(ctx context.Context) ApiCredentialMappingsCreateRequest {
	return ApiCredentialMappingsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CredentialMapping
func (a *CredentialMappingsAPIService) CredentialMappingsCreateExecute(r ApiCredentialMappingsCreateRequest) (*CredentialMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CredentialMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.credentialMappingRequest == nil {
		return localVarReturnValue, nil, reportError("credentialMappingRequest is required and must be specified")
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
	localVarPostBody = r.credentialMappingRequest
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

type ApiCredentialMappingsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiCredentialMappingsDeletePreviewListRequest) Limit(limit int32) ApiCredentialMappingsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiCredentialMappingsDeletePreviewListRequest) Offset(offset int32) ApiCredentialMappingsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiCredentialMappingsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.CredentialMappingsDeletePreviewListExecute(r)
}

/*
CredentialMappingsDeletePreviewList Method for CredentialMappingsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this cred_ mapping.
 @return ApiCredentialMappingsDeletePreviewListRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsDeletePreviewList(ctx context.Context, id int32) ApiCredentialMappingsDeletePreviewListRequest {
	return ApiCredentialMappingsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *CredentialMappingsAPIService) CredentialMappingsDeletePreviewListExecute(r ApiCredentialMappingsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/{id}/delete_preview/"
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

type ApiCredentialMappingsDestroyRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	id int32
}

func (r ApiCredentialMappingsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.CredentialMappingsDestroyExecute(r)
}

/*
CredentialMappingsDestroy Method for CredentialMappingsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this cred_ mapping.
 @return ApiCredentialMappingsDestroyRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsDestroy(ctx context.Context, id int32) ApiCredentialMappingsDestroyRequest {
	return ApiCredentialMappingsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CredentialMappingsAPIService) CredentialMappingsDestroyExecute(r ApiCredentialMappingsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/{id}/"
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

type ApiCredentialMappingsListRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	credId *int32
	engagement *int32
	finding *int32
	isAuthnProvider *bool
	limit *int32
	offset *int32
	product *int32
	test *int32
	url *string
}

func (r ApiCredentialMappingsListRequest) CredId(credId int32) ApiCredentialMappingsListRequest {
	r.credId = &credId
	return r
}

func (r ApiCredentialMappingsListRequest) Engagement(engagement int32) ApiCredentialMappingsListRequest {
	r.engagement = &engagement
	return r
}

func (r ApiCredentialMappingsListRequest) Finding(finding int32) ApiCredentialMappingsListRequest {
	r.finding = &finding
	return r
}

func (r ApiCredentialMappingsListRequest) IsAuthnProvider(isAuthnProvider bool) ApiCredentialMappingsListRequest {
	r.isAuthnProvider = &isAuthnProvider
	return r
}

// Number of results to return per page.
func (r ApiCredentialMappingsListRequest) Limit(limit int32) ApiCredentialMappingsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiCredentialMappingsListRequest) Offset(offset int32) ApiCredentialMappingsListRequest {
	r.offset = &offset
	return r
}

func (r ApiCredentialMappingsListRequest) Product(product int32) ApiCredentialMappingsListRequest {
	r.product = &product
	return r
}

func (r ApiCredentialMappingsListRequest) Test(test int32) ApiCredentialMappingsListRequest {
	r.test = &test
	return r
}

func (r ApiCredentialMappingsListRequest) Url(url string) ApiCredentialMappingsListRequest {
	r.url = &url
	return r
}

func (r ApiCredentialMappingsListRequest) Execute() (*PaginatedCredentialMappingList, *http.Response, error) {
	return r.ApiService.CredentialMappingsListExecute(r)
}

/*
CredentialMappingsList Method for CredentialMappingsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCredentialMappingsListRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsList(ctx context.Context) ApiCredentialMappingsListRequest {
	return ApiCredentialMappingsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedCredentialMappingList
func (a *CredentialMappingsAPIService) CredentialMappingsListExecute(r ApiCredentialMappingsListRequest) (*PaginatedCredentialMappingList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCredentialMappingList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.credId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cred_id", r.credId, "")
	}
	if r.engagement != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "engagement", r.engagement, "")
	}
	if r.finding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "finding", r.finding, "")
	}
	if r.isAuthnProvider != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_authn_provider", r.isAuthnProvider, "")
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
	if r.test != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "test", r.test, "")
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

type ApiCredentialMappingsPartialUpdateRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	id int32
	patchedCredentialMappingRequest *PatchedCredentialMappingRequest
}

func (r ApiCredentialMappingsPartialUpdateRequest) PatchedCredentialMappingRequest(patchedCredentialMappingRequest PatchedCredentialMappingRequest) ApiCredentialMappingsPartialUpdateRequest {
	r.patchedCredentialMappingRequest = &patchedCredentialMappingRequest
	return r
}

func (r ApiCredentialMappingsPartialUpdateRequest) Execute() (*CredentialMapping, *http.Response, error) {
	return r.ApiService.CredentialMappingsPartialUpdateExecute(r)
}

/*
CredentialMappingsPartialUpdate Method for CredentialMappingsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this cred_ mapping.
 @return ApiCredentialMappingsPartialUpdateRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsPartialUpdate(ctx context.Context, id int32) ApiCredentialMappingsPartialUpdateRequest {
	return ApiCredentialMappingsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CredentialMapping
func (a *CredentialMappingsAPIService) CredentialMappingsPartialUpdateExecute(r ApiCredentialMappingsPartialUpdateRequest) (*CredentialMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CredentialMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/{id}/"
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
	localVarPostBody = r.patchedCredentialMappingRequest
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

type ApiCredentialMappingsRetrieveRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	id int32
}

func (r ApiCredentialMappingsRetrieveRequest) Execute() (*CredentialMapping, *http.Response, error) {
	return r.ApiService.CredentialMappingsRetrieveExecute(r)
}

/*
CredentialMappingsRetrieve Method for CredentialMappingsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this cred_ mapping.
 @return ApiCredentialMappingsRetrieveRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsRetrieve(ctx context.Context, id int32) ApiCredentialMappingsRetrieveRequest {
	return ApiCredentialMappingsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CredentialMapping
func (a *CredentialMappingsAPIService) CredentialMappingsRetrieveExecute(r ApiCredentialMappingsRetrieveRequest) (*CredentialMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CredentialMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/{id}/"
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

type ApiCredentialMappingsUpdateRequest struct {
	ctx context.Context
	ApiService CredentialMappingsAPI
	id int32
	credentialMappingRequest *CredentialMappingRequest
}

func (r ApiCredentialMappingsUpdateRequest) CredentialMappingRequest(credentialMappingRequest CredentialMappingRequest) ApiCredentialMappingsUpdateRequest {
	r.credentialMappingRequest = &credentialMappingRequest
	return r
}

func (r ApiCredentialMappingsUpdateRequest) Execute() (*CredentialMapping, *http.Response, error) {
	return r.ApiService.CredentialMappingsUpdateExecute(r)
}

/*
CredentialMappingsUpdate Method for CredentialMappingsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this cred_ mapping.
 @return ApiCredentialMappingsUpdateRequest
*/
func (a *CredentialMappingsAPIService) CredentialMappingsUpdate(ctx context.Context, id int32) ApiCredentialMappingsUpdateRequest {
	return ApiCredentialMappingsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CredentialMapping
func (a *CredentialMappingsAPIService) CredentialMappingsUpdateExecute(r ApiCredentialMappingsUpdateRequest) (*CredentialMapping, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CredentialMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialMappingsAPIService.CredentialMappingsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/credential_mappings/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.credentialMappingRequest == nil {
		return localVarReturnValue, nil, reportError("credentialMappingRequest is required and must be specified")
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
	localVarPostBody = r.credentialMappingRequest
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