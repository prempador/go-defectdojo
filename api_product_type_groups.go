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


type ProductTypeGroupsAPI interface {

	/*
	ProductTypeGroupsCreate Method for ProductTypeGroupsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductTypeGroupsCreateRequest
	*/
	ProductTypeGroupsCreate(ctx context.Context) ApiProductTypeGroupsCreateRequest

	// ProductTypeGroupsCreateExecute executes the request
	//  @return ProductTypeGroup
	ProductTypeGroupsCreateExecute(r ApiProductTypeGroupsCreateRequest) (*ProductTypeGroup, *http.Response, error)

	/*
	ProductTypeGroupsDeletePreviewList Method for ProductTypeGroupsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ type_ group.
	@return ApiProductTypeGroupsDeletePreviewListRequest
	*/
	ProductTypeGroupsDeletePreviewList(ctx context.Context, id int32) ApiProductTypeGroupsDeletePreviewListRequest

	// ProductTypeGroupsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	ProductTypeGroupsDeletePreviewListExecute(r ApiProductTypeGroupsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	ProductTypeGroupsDestroy Method for ProductTypeGroupsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ type_ group.
	@return ApiProductTypeGroupsDestroyRequest
	*/
	ProductTypeGroupsDestroy(ctx context.Context, id int32) ApiProductTypeGroupsDestroyRequest

	// ProductTypeGroupsDestroyExecute executes the request
	ProductTypeGroupsDestroyExecute(r ApiProductTypeGroupsDestroyRequest) (*http.Response, error)

	/*
	ProductTypeGroupsList Method for ProductTypeGroupsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductTypeGroupsListRequest
	*/
	ProductTypeGroupsList(ctx context.Context) ApiProductTypeGroupsListRequest

	// ProductTypeGroupsListExecute executes the request
	//  @return PaginatedProductTypeGroupList
	ProductTypeGroupsListExecute(r ApiProductTypeGroupsListRequest) (*PaginatedProductTypeGroupList, *http.Response, error)

	/*
	ProductTypeGroupsRetrieve Method for ProductTypeGroupsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ type_ group.
	@return ApiProductTypeGroupsRetrieveRequest
	*/
	ProductTypeGroupsRetrieve(ctx context.Context, id int32) ApiProductTypeGroupsRetrieveRequest

	// ProductTypeGroupsRetrieveExecute executes the request
	//  @return ProductTypeGroup
	ProductTypeGroupsRetrieveExecute(r ApiProductTypeGroupsRetrieveRequest) (*ProductTypeGroup, *http.Response, error)

	/*
	ProductTypeGroupsUpdate Method for ProductTypeGroupsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this product_ type_ group.
	@return ApiProductTypeGroupsUpdateRequest
	*/
	ProductTypeGroupsUpdate(ctx context.Context, id int32) ApiProductTypeGroupsUpdateRequest

	// ProductTypeGroupsUpdateExecute executes the request
	//  @return ProductTypeGroup
	ProductTypeGroupsUpdateExecute(r ApiProductTypeGroupsUpdateRequest) (*ProductTypeGroup, *http.Response, error)
}

// ProductTypeGroupsAPIService ProductTypeGroupsAPI service
type ProductTypeGroupsAPIService service

type ApiProductTypeGroupsCreateRequest struct {
	ctx context.Context
	ApiService ProductTypeGroupsAPI
	productTypeGroupRequest *ProductTypeGroupRequest
}

func (r ApiProductTypeGroupsCreateRequest) ProductTypeGroupRequest(productTypeGroupRequest ProductTypeGroupRequest) ApiProductTypeGroupsCreateRequest {
	r.productTypeGroupRequest = &productTypeGroupRequest
	return r
}

func (r ApiProductTypeGroupsCreateRequest) Execute() (*ProductTypeGroup, *http.Response, error) {
	return r.ApiService.ProductTypeGroupsCreateExecute(r)
}

/*
ProductTypeGroupsCreate Method for ProductTypeGroupsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProductTypeGroupsCreateRequest
*/
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsCreate(ctx context.Context) ApiProductTypeGroupsCreateRequest {
	return ApiProductTypeGroupsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductTypeGroup
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsCreateExecute(r ApiProductTypeGroupsCreateRequest) (*ProductTypeGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductTypeGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTypeGroupsAPIService.ProductTypeGroupsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_type_groups/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productTypeGroupRequest == nil {
		return localVarReturnValue, nil, reportError("productTypeGroupRequest is required and must be specified")
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
	localVarPostBody = r.productTypeGroupRequest
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

type ApiProductTypeGroupsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService ProductTypeGroupsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiProductTypeGroupsDeletePreviewListRequest) Limit(limit int32) ApiProductTypeGroupsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiProductTypeGroupsDeletePreviewListRequest) Offset(offset int32) ApiProductTypeGroupsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiProductTypeGroupsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.ProductTypeGroupsDeletePreviewListExecute(r)
}

/*
ProductTypeGroupsDeletePreviewList Method for ProductTypeGroupsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ type_ group.
 @return ApiProductTypeGroupsDeletePreviewListRequest
*/
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsDeletePreviewList(ctx context.Context, id int32) ApiProductTypeGroupsDeletePreviewListRequest {
	return ApiProductTypeGroupsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsDeletePreviewListExecute(r ApiProductTypeGroupsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTypeGroupsAPIService.ProductTypeGroupsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_type_groups/{id}/delete_preview/"
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

type ApiProductTypeGroupsDestroyRequest struct {
	ctx context.Context
	ApiService ProductTypeGroupsAPI
	id int32
}

func (r ApiProductTypeGroupsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProductTypeGroupsDestroyExecute(r)
}

/*
ProductTypeGroupsDestroy Method for ProductTypeGroupsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ type_ group.
 @return ApiProductTypeGroupsDestroyRequest
*/
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsDestroy(ctx context.Context, id int32) ApiProductTypeGroupsDestroyRequest {
	return ApiProductTypeGroupsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsDestroyExecute(r ApiProductTypeGroupsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTypeGroupsAPIService.ProductTypeGroupsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_type_groups/{id}/"
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

type ApiProductTypeGroupsListRequest struct {
	ctx context.Context
	ApiService ProductTypeGroupsAPI
	groupId *int32
	id *int32
	limit *int32
	offset *int32
	prefetch *[]string
	productTypeId *int32
}

func (r ApiProductTypeGroupsListRequest) GroupId(groupId int32) ApiProductTypeGroupsListRequest {
	r.groupId = &groupId
	return r
}

func (r ApiProductTypeGroupsListRequest) Id(id int32) ApiProductTypeGroupsListRequest {
	r.id = &id
	return r
}

// Number of results to return per page.
func (r ApiProductTypeGroupsListRequest) Limit(limit int32) ApiProductTypeGroupsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiProductTypeGroupsListRequest) Offset(offset int32) ApiProductTypeGroupsListRequest {
	r.offset = &offset
	return r
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiProductTypeGroupsListRequest) Prefetch(prefetch []string) ApiProductTypeGroupsListRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiProductTypeGroupsListRequest) ProductTypeId(productTypeId int32) ApiProductTypeGroupsListRequest {
	r.productTypeId = &productTypeId
	return r
}

func (r ApiProductTypeGroupsListRequest) Execute() (*PaginatedProductTypeGroupList, *http.Response, error) {
	return r.ApiService.ProductTypeGroupsListExecute(r)
}

/*
ProductTypeGroupsList Method for ProductTypeGroupsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProductTypeGroupsListRequest
*/
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsList(ctx context.Context) ApiProductTypeGroupsListRequest {
	return ApiProductTypeGroupsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedProductTypeGroupList
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsListExecute(r ApiProductTypeGroupsListRequest) (*PaginatedProductTypeGroupList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedProductTypeGroupList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTypeGroupsAPIService.ProductTypeGroupsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_type_groups/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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
	if r.productTypeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_type_id", r.productTypeId, "form", "")
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

type ApiProductTypeGroupsRetrieveRequest struct {
	ctx context.Context
	ApiService ProductTypeGroupsAPI
	id int32
	prefetch *[]string
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiProductTypeGroupsRetrieveRequest) Prefetch(prefetch []string) ApiProductTypeGroupsRetrieveRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiProductTypeGroupsRetrieveRequest) Execute() (*ProductTypeGroup, *http.Response, error) {
	return r.ApiService.ProductTypeGroupsRetrieveExecute(r)
}

/*
ProductTypeGroupsRetrieve Method for ProductTypeGroupsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ type_ group.
 @return ApiProductTypeGroupsRetrieveRequest
*/
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsRetrieve(ctx context.Context, id int32) ApiProductTypeGroupsRetrieveRequest {
	return ApiProductTypeGroupsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductTypeGroup
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsRetrieveExecute(r ApiProductTypeGroupsRetrieveRequest) (*ProductTypeGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductTypeGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTypeGroupsAPIService.ProductTypeGroupsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_type_groups/{id}/"
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

type ApiProductTypeGroupsUpdateRequest struct {
	ctx context.Context
	ApiService ProductTypeGroupsAPI
	id int32
	productTypeGroupRequest *ProductTypeGroupRequest
}

func (r ApiProductTypeGroupsUpdateRequest) ProductTypeGroupRequest(productTypeGroupRequest ProductTypeGroupRequest) ApiProductTypeGroupsUpdateRequest {
	r.productTypeGroupRequest = &productTypeGroupRequest
	return r
}

func (r ApiProductTypeGroupsUpdateRequest) Execute() (*ProductTypeGroup, *http.Response, error) {
	return r.ApiService.ProductTypeGroupsUpdateExecute(r)
}

/*
ProductTypeGroupsUpdate Method for ProductTypeGroupsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this product_ type_ group.
 @return ApiProductTypeGroupsUpdateRequest
*/
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsUpdate(ctx context.Context, id int32) ApiProductTypeGroupsUpdateRequest {
	return ApiProductTypeGroupsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ProductTypeGroup
func (a *ProductTypeGroupsAPIService) ProductTypeGroupsUpdateExecute(r ApiProductTypeGroupsUpdateRequest) (*ProductTypeGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductTypeGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductTypeGroupsAPIService.ProductTypeGroupsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/product_type_groups/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.productTypeGroupRequest == nil {
		return localVarReturnValue, nil, reportError("productTypeGroupRequest is required and must be specified")
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
	localVarPostBody = r.productTypeGroupRequest
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
