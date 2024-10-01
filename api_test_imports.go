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
	"time"
)


type TestImportsAPI interface {

	/*
	TestImportsCreate Method for TestImportsCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTestImportsCreateRequest
	*/
	TestImportsCreate(ctx context.Context) ApiTestImportsCreateRequest

	// TestImportsCreateExecute executes the request
	//  @return TestImport
	TestImportsCreateExecute(r ApiTestImportsCreateRequest) (*TestImport, *http.Response, error)

	/*
	TestImportsDeletePreviewList Method for TestImportsDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this test_ import.
	@return ApiTestImportsDeletePreviewListRequest
	*/
	TestImportsDeletePreviewList(ctx context.Context, id int32) ApiTestImportsDeletePreviewListRequest

	// TestImportsDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	TestImportsDeletePreviewListExecute(r ApiTestImportsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	TestImportsDestroy Method for TestImportsDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this test_ import.
	@return ApiTestImportsDestroyRequest
	*/
	TestImportsDestroy(ctx context.Context, id int32) ApiTestImportsDestroyRequest

	// TestImportsDestroyExecute executes the request
	TestImportsDestroyExecute(r ApiTestImportsDestroyRequest) (*http.Response, error)

	/*
	TestImportsList Method for TestImportsList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTestImportsListRequest
	*/
	TestImportsList(ctx context.Context) ApiTestImportsListRequest

	// TestImportsListExecute executes the request
	//  @return PaginatedTestImportList
	TestImportsListExecute(r ApiTestImportsListRequest) (*PaginatedTestImportList, *http.Response, error)

	/*
	TestImportsPartialUpdate Method for TestImportsPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this test_ import.
	@return ApiTestImportsPartialUpdateRequest
	*/
	TestImportsPartialUpdate(ctx context.Context, id int32) ApiTestImportsPartialUpdateRequest

	// TestImportsPartialUpdateExecute executes the request
	//  @return TestImport
	TestImportsPartialUpdateExecute(r ApiTestImportsPartialUpdateRequest) (*TestImport, *http.Response, error)

	/*
	TestImportsRetrieve Method for TestImportsRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this test_ import.
	@return ApiTestImportsRetrieveRequest
	*/
	TestImportsRetrieve(ctx context.Context, id int32) ApiTestImportsRetrieveRequest

	// TestImportsRetrieveExecute executes the request
	//  @return TestImport
	TestImportsRetrieveExecute(r ApiTestImportsRetrieveRequest) (*TestImport, *http.Response, error)

	/*
	TestImportsUpdate Method for TestImportsUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this test_ import.
	@return ApiTestImportsUpdateRequest
	*/
	TestImportsUpdate(ctx context.Context, id int32) ApiTestImportsUpdateRequest

	// TestImportsUpdateExecute executes the request
	//  @return TestImport
	TestImportsUpdateExecute(r ApiTestImportsUpdateRequest) (*TestImport, *http.Response, error)
}

// TestImportsAPIService TestImportsAPI service
type TestImportsAPIService service

type ApiTestImportsCreateRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	testImportRequest *TestImportRequest
}

func (r ApiTestImportsCreateRequest) TestImportRequest(testImportRequest TestImportRequest) ApiTestImportsCreateRequest {
	r.testImportRequest = &testImportRequest
	return r
}

func (r ApiTestImportsCreateRequest) Execute() (*TestImport, *http.Response, error) {
	return r.ApiService.TestImportsCreateExecute(r)
}

/*
TestImportsCreate Method for TestImportsCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTestImportsCreateRequest
*/
func (a *TestImportsAPIService) TestImportsCreate(ctx context.Context) ApiTestImportsCreateRequest {
	return ApiTestImportsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TestImport
func (a *TestImportsAPIService) TestImportsCreateExecute(r ApiTestImportsCreateRequest) (*TestImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/"

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
	localVarPostBody = r.testImportRequest
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

type ApiTestImportsDeletePreviewListRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiTestImportsDeletePreviewListRequest) Limit(limit int32) ApiTestImportsDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiTestImportsDeletePreviewListRequest) Offset(offset int32) ApiTestImportsDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiTestImportsDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.TestImportsDeletePreviewListExecute(r)
}

/*
TestImportsDeletePreviewList Method for TestImportsDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this test_ import.
 @return ApiTestImportsDeletePreviewListRequest
*/
func (a *TestImportsAPIService) TestImportsDeletePreviewList(ctx context.Context, id int32) ApiTestImportsDeletePreviewListRequest {
	return ApiTestImportsDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *TestImportsAPIService) TestImportsDeletePreviewListExecute(r ApiTestImportsDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/{id}/delete_preview/"
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

type ApiTestImportsDestroyRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	id int32
}

func (r ApiTestImportsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.TestImportsDestroyExecute(r)
}

/*
TestImportsDestroy Method for TestImportsDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this test_ import.
 @return ApiTestImportsDestroyRequest
*/
func (a *TestImportsAPIService) TestImportsDestroy(ctx context.Context, id int32) ApiTestImportsDestroyRequest {
	return ApiTestImportsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TestImportsAPIService) TestImportsDestroyExecute(r ApiTestImportsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/{id}/"
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

type ApiTestImportsListRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	branchTag *string
	buildId *string
	commitHash *string
	findingsAffected *[]int32
	limit *int32
	offset *int32
	test *int32
	testImportFindingActionAction *string
	testImportFindingActionCreated *time.Time
	testImportFindingActionFinding *int32
	version *string
}

func (r ApiTestImportsListRequest) BranchTag(branchTag string) ApiTestImportsListRequest {
	r.branchTag = &branchTag
	return r
}

func (r ApiTestImportsListRequest) BuildId(buildId string) ApiTestImportsListRequest {
	r.buildId = &buildId
	return r
}

func (r ApiTestImportsListRequest) CommitHash(commitHash string) ApiTestImportsListRequest {
	r.commitHash = &commitHash
	return r
}

func (r ApiTestImportsListRequest) FindingsAffected(findingsAffected []int32) ApiTestImportsListRequest {
	r.findingsAffected = &findingsAffected
	return r
}

// Number of results to return per page.
func (r ApiTestImportsListRequest) Limit(limit int32) ApiTestImportsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiTestImportsListRequest) Offset(offset int32) ApiTestImportsListRequest {
	r.offset = &offset
	return r
}

func (r ApiTestImportsListRequest) Test(test int32) ApiTestImportsListRequest {
	r.test = &test
	return r
}

// * &#x60;N&#x60; - created * &#x60;C&#x60; - closed * &#x60;R&#x60; - reactivated * &#x60;U&#x60; - left untouched
func (r ApiTestImportsListRequest) TestImportFindingActionAction(testImportFindingActionAction string) ApiTestImportsListRequest {
	r.testImportFindingActionAction = &testImportFindingActionAction
	return r
}

func (r ApiTestImportsListRequest) TestImportFindingActionCreated(testImportFindingActionCreated time.Time) ApiTestImportsListRequest {
	r.testImportFindingActionCreated = &testImportFindingActionCreated
	return r
}

func (r ApiTestImportsListRequest) TestImportFindingActionFinding(testImportFindingActionFinding int32) ApiTestImportsListRequest {
	r.testImportFindingActionFinding = &testImportFindingActionFinding
	return r
}

func (r ApiTestImportsListRequest) Version(version string) ApiTestImportsListRequest {
	r.version = &version
	return r
}

func (r ApiTestImportsListRequest) Execute() (*PaginatedTestImportList, *http.Response, error) {
	return r.ApiService.TestImportsListExecute(r)
}

/*
TestImportsList Method for TestImportsList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTestImportsListRequest
*/
func (a *TestImportsAPIService) TestImportsList(ctx context.Context) ApiTestImportsListRequest {
	return ApiTestImportsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedTestImportList
func (a *TestImportsAPIService) TestImportsListExecute(r ApiTestImportsListRequest) (*PaginatedTestImportList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedTestImportList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.branchTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "branch_tag", r.branchTag, "form", "")
	}
	if r.buildId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "build_id", r.buildId, "form", "")
	}
	if r.commitHash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "commit_hash", r.commitHash, "form", "")
	}
	if r.findingsAffected != nil {
		t := *r.findingsAffected
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "findings_affected", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "findings_affected", t, "form", "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.test != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "test", r.test, "form", "")
	}
	if r.testImportFindingActionAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "test_import_finding_action__action", r.testImportFindingActionAction, "form", "")
	}
	if r.testImportFindingActionCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "test_import_finding_action__created", r.testImportFindingActionCreated, "form", "")
	}
	if r.testImportFindingActionFinding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "test_import_finding_action__finding", r.testImportFindingActionFinding, "form", "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
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

type ApiTestImportsPartialUpdateRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	id int32
	patchedTestImportRequest *PatchedTestImportRequest
}

func (r ApiTestImportsPartialUpdateRequest) PatchedTestImportRequest(patchedTestImportRequest PatchedTestImportRequest) ApiTestImportsPartialUpdateRequest {
	r.patchedTestImportRequest = &patchedTestImportRequest
	return r
}

func (r ApiTestImportsPartialUpdateRequest) Execute() (*TestImport, *http.Response, error) {
	return r.ApiService.TestImportsPartialUpdateExecute(r)
}

/*
TestImportsPartialUpdate Method for TestImportsPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this test_ import.
 @return ApiTestImportsPartialUpdateRequest
*/
func (a *TestImportsAPIService) TestImportsPartialUpdate(ctx context.Context, id int32) ApiTestImportsPartialUpdateRequest {
	return ApiTestImportsPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TestImport
func (a *TestImportsAPIService) TestImportsPartialUpdateExecute(r ApiTestImportsPartialUpdateRequest) (*TestImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/{id}/"
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
	localVarPostBody = r.patchedTestImportRequest
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

type ApiTestImportsRetrieveRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	id int32
}

func (r ApiTestImportsRetrieveRequest) Execute() (*TestImport, *http.Response, error) {
	return r.ApiService.TestImportsRetrieveExecute(r)
}

/*
TestImportsRetrieve Method for TestImportsRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this test_ import.
 @return ApiTestImportsRetrieveRequest
*/
func (a *TestImportsAPIService) TestImportsRetrieve(ctx context.Context, id int32) ApiTestImportsRetrieveRequest {
	return ApiTestImportsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TestImport
func (a *TestImportsAPIService) TestImportsRetrieveExecute(r ApiTestImportsRetrieveRequest) (*TestImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/{id}/"
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

type ApiTestImportsUpdateRequest struct {
	ctx context.Context
	ApiService TestImportsAPI
	id int32
	testImportRequest *TestImportRequest
}

func (r ApiTestImportsUpdateRequest) TestImportRequest(testImportRequest TestImportRequest) ApiTestImportsUpdateRequest {
	r.testImportRequest = &testImportRequest
	return r
}

func (r ApiTestImportsUpdateRequest) Execute() (*TestImport, *http.Response, error) {
	return r.ApiService.TestImportsUpdateExecute(r)
}

/*
TestImportsUpdate Method for TestImportsUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this test_ import.
 @return ApiTestImportsUpdateRequest
*/
func (a *TestImportsAPIService) TestImportsUpdate(ctx context.Context, id int32) ApiTestImportsUpdateRequest {
	return ApiTestImportsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TestImport
func (a *TestImportsAPIService) TestImportsUpdateExecute(r ApiTestImportsUpdateRequest) (*TestImport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestImportsAPIService.TestImportsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/test_imports/{id}/"
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
	localVarPostBody = r.testImportRequest
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
