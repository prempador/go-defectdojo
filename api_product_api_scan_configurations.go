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


type ProductApiScanConfigurationsAPI interface {

	/*
	ProductApiScanConfigurationsCreate Method for ProductApiScanConfigurationsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductApiScanConfigurationsCreateRequest
	*/
	ProductApiScanConfigurationsCreate(ctx context.Context) ApiProductApiScanConfigurationsCreateRequest

	// ProductApiScanConfigurationsCreateExecute executes the request
	//  @return ProductAPIScanConfiguration
	ProductApiScanConfigurationsCreateExecute(r ApiProductApiScanConfigurationsCreateRequest) (*ProductAPIScanConfiguration, *http.Response, error)

	/*
	ProductApiScanConfigurationsDeletePreviewList Method for ProductApiScanConfigurationsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ap i_ scan_ configuration.
	@return ApiProductApiScanConfigurationsDeletePreviewListRequest
	*/
	ProductApiScanConfigurationsDeletePreviewList(ctx context.Context, id int32) ApiProductApiScanConfigurationsDeletePreviewListRequest

	// ProductApiScanConfigurationsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	ProductApiScanConfigurationsDeletePreviewListExecute(r ApiProductApiScanConfigurationsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	ProductApiScanConfigurationsDestroy Method for ProductApiScanConfigurationsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ap i_ scan_ configuration.
	@return ApiProductApiScanConfigurationsDestroyRequest
	*/
	ProductApiScanConfigurationsDestroy(ctx context.Context, id int32) ApiProductApiScanConfigurationsDestroyRequest

	// ProductApiScanConfigurationsDestroyExecute executes the request
	ProductApiScanConfigurationsDestroyExecute(r ApiProductApiScanConfigurationsDestroyRequest) (*http.Response, error)

	/*
	ProductApiScanConfigurationsList Method for ProductApiScanConfigurationsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductApiScanConfigurationsListRequest
	*/
	ProductApiScanConfigurationsList(ctx context.Context) ApiProductApiScanConfigurationsListRequest

	// ProductApiScanConfigurationsListExecute executes the request
	//  @return PaginatedProductAPIScanConfigurationList
	ProductApiScanConfigurationsListExecute(r ApiProductApiScanConfigurationsListRequest) (*PaginatedProductAPIScanConfigurationList, *http.Response, error)

	/*
	ProductApiScanConfigurationsPartialUpdate Method for ProductApiScanConfigurationsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ap i_ scan_ configuration.
	@return ApiProductApiScanConfigurationsPartialUpdateRequest
	*/
	ProductApiScanConfigurationsPartialUpdate(ctx context.Context, id int32) ApiProductApiScanConfigurationsPartialUpdateRequest

	// ProductApiScanConfigurationsPartialUpdateExecute executes the request
	//  @return ProductAPIScanConfiguration
	ProductApiScanConfigurationsPartialUpdateExecute(r ApiProductApiScanConfigurationsPartialUpdateRequest) (*ProductAPIScanConfiguration, *http.Response, error)

	/*
	ProductApiScanConfigurationsRetrieve Method for ProductApiScanConfigurationsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ap i_ scan_ configuration.
	@return ApiProductApiScanConfigurationsRetrieveRequest
	*/
	ProductApiScanConfigurationsRetrieve(ctx context.Context, id int32) ApiProductApiScanConfigurationsRetrieveRequest

	// ProductApiScanConfigurationsRetrieveExecute executes the request
	//  @return ProductAPIScanConfiguration
	ProductApiScanConfigurationsRetrieveExecute(r ApiProductApiScanConfigurationsRetrieveRequest) (*ProductAPIScanConfiguration, *http.Response, error)

	/*
	ProductApiScanConfigurationsUpdate Method for ProductApiScanConfigurationsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ap i_ scan_ configuration.
	@return ApiProductApiScanConfigurationsUpdateRequest
	*/
	ProductApiScanConfigurationsUpdate(ctx context.Context, id int32) ApiProductApiScanConfigurationsUpdateRequest

	// ProductApiScanConfigurationsUpdateExecute executes the request
	//  @return ProductAPIScanConfiguration
	ProductApiScanConfigurationsUpdateExecute(r ApiProductApiScanConfigurationsUpdateRequest) (*ProductAPIScanConfiguration, *http.Response, error)
}

// ProductApiScanConfigurationsAPIService ProductApiScanConfigurationsAPI service
type ProductApiScanConfigurationsAPIService service

type ApiProductApiScanConfigurationsCreateRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	productAPIScanConfigurationRequest *ProductAPIScanConfigurationRequest
}

func (r ApiProductApiScanConfigurationsCreateRequest) ProductAPIScanConfigurationRequest(productAPIScanConfigurationRequest ProductAPIScanConfigurationRequest) ApiProductApiScanConfigurationsCreateRequest {
	r.productAPIScanConfigurationRequest = &productAPIScanConfigurationRequest
	return r
}

func (r ApiProductApiScanConfigurationsCreateRequest) Execute() (*ProductAPIScanConfiguration, *http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsCreateExecute(r)
}

/*
ProductApiScanConfigurationsCreate Method for ProductApiScanConfigurationsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProductApiScanConfigurationsCreateRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsCreate(ctx context.Context) ApiProductApiScanConfigurationsCreateRequest {
	return ApiProductApiScanConfigurationsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductAPIScanConfiguration
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsCreateExecute(r ApiProductApiScanConfigurationsCreateRequest) (*ProductAPIScanConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductAPIScanConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productAPIScanConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("productAPIScanConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.productAPIScanConfigurationRequest
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

type ApiProductApiScanConfigurationsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiProductApiScanConfigurationsDeletePreviewListRequest) Limit(limit int32) ApiProductApiScanConfigurationsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiProductApiScanConfigurationsDeletePreviewListRequest) Offset(offset int32) ApiProductApiScanConfigurationsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiProductApiScanConfigurationsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsDeletePreviewListExecute(r)
}

/*
ProductApiScanConfigurationsDeletePreviewList Method for ProductApiScanConfigurationsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ap i_ scan_ configuration.
 @return ApiProductApiScanConfigurationsDeletePreviewListRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsDeletePreviewList(ctx context.Context, id int32) ApiProductApiScanConfigurationsDeletePreviewListRequest {
	return ApiProductApiScanConfigurationsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsDeletePreviewListExecute(r ApiProductApiScanConfigurationsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/{id}/delete_preview/"
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

type ApiProductApiScanConfigurationsDestroyRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	id int32
}

func (r ApiProductApiScanConfigurationsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsDestroyExecute(r)
}

/*
ProductApiScanConfigurationsDestroy Method for ProductApiScanConfigurationsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ap i_ scan_ configuration.
 @return ApiProductApiScanConfigurationsDestroyRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsDestroy(ctx context.Context, id int32) ApiProductApiScanConfigurationsDestroyRequest {
	return ApiProductApiScanConfigurationsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsDestroyExecute(r ApiProductApiScanConfigurationsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/{id}/"
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

type ApiProductApiScanConfigurationsListRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	id *int32
	limit *int32
	offset *int32
	product *int32
	serviceKey1 *string
	serviceKey2 *string
	serviceKey3 *string
	toolConfiguration *int32
}

func (r ApiProductApiScanConfigurationsListRequest) Id(id int32) ApiProductApiScanConfigurationsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiProductApiScanConfigurationsListRequest) Limit(limit int32) ApiProductApiScanConfigurationsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiProductApiScanConfigurationsListRequest) Offset(offset int32) ApiProductApiScanConfigurationsListRequest {
	r.offset = &offset
	return r
}

func (r ApiProductApiScanConfigurationsListRequest) Product(product int32) ApiProductApiScanConfigurationsListRequest {
	r.product = &product
	return r
}

func (r ApiProductApiScanConfigurationsListRequest) ServiceKey1(serviceKey1 string) ApiProductApiScanConfigurationsListRequest {
	r.serviceKey1 = &serviceKey1
	return r
}

func (r ApiProductApiScanConfigurationsListRequest) ServiceKey2(serviceKey2 string) ApiProductApiScanConfigurationsListRequest {
	r.serviceKey2 = &serviceKey2
	return r
}

func (r ApiProductApiScanConfigurationsListRequest) ServiceKey3(serviceKey3 string) ApiProductApiScanConfigurationsListRequest {
	r.serviceKey3 = &serviceKey3
	return r
}

func (r ApiProductApiScanConfigurationsListRequest) ToolConfiguration(toolConfiguration int32) ApiProductApiScanConfigurationsListRequest {
	r.toolConfiguration = &toolConfiguration
	return r
}

func (r ApiProductApiScanConfigurationsListRequest) Execute() (*PaginatedProductAPIScanConfigurationList, *http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsListExecute(r)
}

/*
ProductApiScanConfigurationsList Method for ProductApiScanConfigurationsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProductApiScanConfigurationsListRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsList(ctx context.Context) ApiProductApiScanConfigurationsListRequest {
	return ApiProductApiScanConfigurationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedProductAPIScanConfigurationList
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsListExecute(r ApiProductApiScanConfigurationsListRequest) (*PaginatedProductAPIScanConfigurationList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedProductAPIScanConfigurationList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/"

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
	if r.serviceKey1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "service_key_1", r.serviceKey1, "")
	}
	if r.serviceKey2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "service_key_2", r.serviceKey2, "")
	}
	if r.serviceKey3 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "service_key_3", r.serviceKey3, "")
	}
	if r.toolConfiguration != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tool_configuration", r.toolConfiguration, "")
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

type ApiProductApiScanConfigurationsPartialUpdateRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	id int32
	patchedProductAPIScanConfigurationRequest *PatchedProductAPIScanConfigurationRequest
}

func (r ApiProductApiScanConfigurationsPartialUpdateRequest) PatchedProductAPIScanConfigurationRequest(patchedProductAPIScanConfigurationRequest PatchedProductAPIScanConfigurationRequest) ApiProductApiScanConfigurationsPartialUpdateRequest {
	r.patchedProductAPIScanConfigurationRequest = &patchedProductAPIScanConfigurationRequest
	return r
}

func (r ApiProductApiScanConfigurationsPartialUpdateRequest) Execute() (*ProductAPIScanConfiguration, *http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsPartialUpdateExecute(r)
}

/*
ProductApiScanConfigurationsPartialUpdate Method for ProductApiScanConfigurationsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ap i_ scan_ configuration.
 @return ApiProductApiScanConfigurationsPartialUpdateRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsPartialUpdate(ctx context.Context, id int32) ApiProductApiScanConfigurationsPartialUpdateRequest {
	return ApiProductApiScanConfigurationsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductAPIScanConfiguration
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsPartialUpdateExecute(r ApiProductApiScanConfigurationsPartialUpdateRequest) (*ProductAPIScanConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductAPIScanConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/{id}/"
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
	localVarPostBody = r.patchedProductAPIScanConfigurationRequest
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

type ApiProductApiScanConfigurationsRetrieveRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	id int32
}

func (r ApiProductApiScanConfigurationsRetrieveRequest) Execute() (*ProductAPIScanConfiguration, *http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsRetrieveExecute(r)
}

/*
ProductApiScanConfigurationsRetrieve Method for ProductApiScanConfigurationsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ap i_ scan_ configuration.
 @return ApiProductApiScanConfigurationsRetrieveRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsRetrieve(ctx context.Context, id int32) ApiProductApiScanConfigurationsRetrieveRequest {
	return ApiProductApiScanConfigurationsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductAPIScanConfiguration
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsRetrieveExecute(r ApiProductApiScanConfigurationsRetrieveRequest) (*ProductAPIScanConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductAPIScanConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/{id}/"
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

type ApiProductApiScanConfigurationsUpdateRequest struct {
	ctx context.Context
	ApiService ProductApiScanConfigurationsAPI
	id int32
	productAPIScanConfigurationRequest *ProductAPIScanConfigurationRequest
}

func (r ApiProductApiScanConfigurationsUpdateRequest) ProductAPIScanConfigurationRequest(productAPIScanConfigurationRequest ProductAPIScanConfigurationRequest) ApiProductApiScanConfigurationsUpdateRequest {
	r.productAPIScanConfigurationRequest = &productAPIScanConfigurationRequest
	return r
}

func (r ApiProductApiScanConfigurationsUpdateRequest) Execute() (*ProductAPIScanConfiguration, *http.Response, error) {
	return r.ApiService.ProductApiScanConfigurationsUpdateExecute(r)
}

/*
ProductApiScanConfigurationsUpdate Method for ProductApiScanConfigurationsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ap i_ scan_ configuration.
 @return ApiProductApiScanConfigurationsUpdateRequest
*/
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsUpdate(ctx context.Context, id int32) ApiProductApiScanConfigurationsUpdateRequest {
	return ApiProductApiScanConfigurationsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductAPIScanConfiguration
func (a *ProductApiScanConfigurationsAPIService) ProductApiScanConfigurationsUpdateExecute(r ApiProductApiScanConfigurationsUpdateRequest) (*ProductAPIScanConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductAPIScanConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductApiScanConfigurationsAPIService.ProductApiScanConfigurationsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_api_scan_configurations/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productAPIScanConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("productAPIScanConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.productAPIScanConfigurationRequest
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
