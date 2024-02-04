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
	"reflect"
)


type UserContactInfosAPI interface {

	/*
	UserContactInfosCreate Method for UserContactInfosCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUserContactInfosCreateRequest
	*/
	UserContactInfosCreate(ctx context.Context) ApiUserContactInfosCreateRequest

	// UserContactInfosCreateExecute executes the request
	//  @return UserContactInfo
	UserContactInfosCreateExecute(r ApiUserContactInfosCreateRequest) (*UserContactInfo, *http.Response, error)

	/*
	UserContactInfosDeletePreviewList Method for UserContactInfosDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this user contact info.
	@return ApiUserContactInfosDeletePreviewListRequest
	*/
	UserContactInfosDeletePreviewList(ctx context.Context, id int32) ApiUserContactInfosDeletePreviewListRequest

	// UserContactInfosDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	UserContactInfosDeletePreviewListExecute(r ApiUserContactInfosDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	UserContactInfosDestroy Method for UserContactInfosDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this user contact info.
	@return ApiUserContactInfosDestroyRequest
	*/
	UserContactInfosDestroy(ctx context.Context, id int32) ApiUserContactInfosDestroyRequest

	// UserContactInfosDestroyExecute executes the request
	UserContactInfosDestroyExecute(r ApiUserContactInfosDestroyRequest) (*http.Response, error)

	/*
	UserContactInfosList Method for UserContactInfosList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUserContactInfosListRequest
	*/
	UserContactInfosList(ctx context.Context) ApiUserContactInfosListRequest

	// UserContactInfosListExecute executes the request
	//  @return PaginatedUserContactInfoList
	UserContactInfosListExecute(r ApiUserContactInfosListRequest) (*PaginatedUserContactInfoList, *http.Response, error)

	/*
	UserContactInfosPartialUpdate Method for UserContactInfosPartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this user contact info.
	@return ApiUserContactInfosPartialUpdateRequest
	*/
	UserContactInfosPartialUpdate(ctx context.Context, id int32) ApiUserContactInfosPartialUpdateRequest

	// UserContactInfosPartialUpdateExecute executes the request
	//  @return UserContactInfo
	UserContactInfosPartialUpdateExecute(r ApiUserContactInfosPartialUpdateRequest) (*UserContactInfo, *http.Response, error)

	/*
	UserContactInfosRetrieve Method for UserContactInfosRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this user contact info.
	@return ApiUserContactInfosRetrieveRequest
	*/
	UserContactInfosRetrieve(ctx context.Context, id int32) ApiUserContactInfosRetrieveRequest

	// UserContactInfosRetrieveExecute executes the request
	//  @return UserContactInfo
	UserContactInfosRetrieveExecute(r ApiUserContactInfosRetrieveRequest) (*UserContactInfo, *http.Response, error)

	/*
	UserContactInfosUpdate Method for UserContactInfosUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this user contact info.
	@return ApiUserContactInfosUpdateRequest
	*/
	UserContactInfosUpdate(ctx context.Context, id int32) ApiUserContactInfosUpdateRequest

	// UserContactInfosUpdateExecute executes the request
	//  @return UserContactInfo
	UserContactInfosUpdateExecute(r ApiUserContactInfosUpdateRequest) (*UserContactInfo, *http.Response, error)
}

// UserContactInfosAPIService UserContactInfosAPI service
type UserContactInfosAPIService service

type ApiUserContactInfosCreateRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	userContactInfoRequest *UserContactInfoRequest
}

func (r ApiUserContactInfosCreateRequest) UserContactInfoRequest(userContactInfoRequest UserContactInfoRequest) ApiUserContactInfosCreateRequest {
	r.userContactInfoRequest = &userContactInfoRequest
	return r
}

func (r ApiUserContactInfosCreateRequest) Execute() (*UserContactInfo, *http.Response, error) {
	return r.ApiService.UserContactInfosCreateExecute(r)
}

/*
UserContactInfosCreate Method for UserContactInfosCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserContactInfosCreateRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosCreate(ctx context.Context) ApiUserContactInfosCreateRequest {
	return ApiUserContactInfosCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserContactInfo
func (a *UserContactInfosAPIService) UserContactInfosCreateExecute(r ApiUserContactInfosCreateRequest) (*UserContactInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserContactInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userContactInfoRequest == nil {
		return localVarReturnValue, nil, reportError("userContactInfoRequest is required and must be specified")
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
	localVarPostBody = r.userContactInfoRequest
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

type ApiUserContactInfosDeletePreviewListRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiUserContactInfosDeletePreviewListRequest) Limit(limit int32) ApiUserContactInfosDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiUserContactInfosDeletePreviewListRequest) Offset(offset int32) ApiUserContactInfosDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiUserContactInfosDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.UserContactInfosDeletePreviewListExecute(r)
}

/*
UserContactInfosDeletePreviewList Method for UserContactInfosDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this user contact info.
 @return ApiUserContactInfosDeletePreviewListRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosDeletePreviewList(ctx context.Context, id int32) ApiUserContactInfosDeletePreviewListRequest {
	return ApiUserContactInfosDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *UserContactInfosAPIService) UserContactInfosDeletePreviewListExecute(r ApiUserContactInfosDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/{id}/delete_preview/"
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

type ApiUserContactInfosDestroyRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	id int32
}

func (r ApiUserContactInfosDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.UserContactInfosDestroyExecute(r)
}

/*
UserContactInfosDestroy Method for UserContactInfosDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this user contact info.
 @return ApiUserContactInfosDestroyRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosDestroy(ctx context.Context, id int32) ApiUserContactInfosDestroyRequest {
	return ApiUserContactInfosDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *UserContactInfosAPIService) UserContactInfosDestroyExecute(r ApiUserContactInfosDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/{id}/"
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

type ApiUserContactInfosListRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	blockExecution *bool
	cellNumber *string
	forcePasswordReset *bool
	githubUsername *string
	limit *int32
	offset *int32
	phoneNumber *string
	prefetch *[]string
	slackUserId *string
	slackUsername *string
	title *string
	twitterUsername *string
	user *int32
}

func (r ApiUserContactInfosListRequest) BlockExecution(blockExecution bool) ApiUserContactInfosListRequest {
	r.blockExecution = &blockExecution
	return r
}

func (r ApiUserContactInfosListRequest) CellNumber(cellNumber string) ApiUserContactInfosListRequest {
	r.cellNumber = &cellNumber
	return r
}

func (r ApiUserContactInfosListRequest) ForcePasswordReset(forcePasswordReset bool) ApiUserContactInfosListRequest {
	r.forcePasswordReset = &forcePasswordReset
	return r
}

func (r ApiUserContactInfosListRequest) GithubUsername(githubUsername string) ApiUserContactInfosListRequest {
	r.githubUsername = &githubUsername
	return r
}

// Number of results to return per page.
func (r ApiUserContactInfosListRequest) Limit(limit int32) ApiUserContactInfosListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiUserContactInfosListRequest) Offset(offset int32) ApiUserContactInfosListRequest {
	r.offset = &offset
	return r
}

func (r ApiUserContactInfosListRequest) PhoneNumber(phoneNumber string) ApiUserContactInfosListRequest {
	r.phoneNumber = &phoneNumber
	return r
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiUserContactInfosListRequest) Prefetch(prefetch []string) ApiUserContactInfosListRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiUserContactInfosListRequest) SlackUserId(slackUserId string) ApiUserContactInfosListRequest {
	r.slackUserId = &slackUserId
	return r
}

func (r ApiUserContactInfosListRequest) SlackUsername(slackUsername string) ApiUserContactInfosListRequest {
	r.slackUsername = &slackUsername
	return r
}

func (r ApiUserContactInfosListRequest) Title(title string) ApiUserContactInfosListRequest {
	r.title = &title
	return r
}

func (r ApiUserContactInfosListRequest) TwitterUsername(twitterUsername string) ApiUserContactInfosListRequest {
	r.twitterUsername = &twitterUsername
	return r
}

func (r ApiUserContactInfosListRequest) User(user int32) ApiUserContactInfosListRequest {
	r.user = &user
	return r
}

func (r ApiUserContactInfosListRequest) Execute() (*PaginatedUserContactInfoList, *http.Response, error) {
	return r.ApiService.UserContactInfosListExecute(r)
}

/*
UserContactInfosList Method for UserContactInfosList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserContactInfosListRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosList(ctx context.Context) ApiUserContactInfosListRequest {
	return ApiUserContactInfosListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedUserContactInfoList
func (a *UserContactInfosAPIService) UserContactInfosListExecute(r ApiUserContactInfosListRequest) (*PaginatedUserContactInfoList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedUserContactInfoList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.blockExecution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "block_execution", r.blockExecution, "")
	}
	if r.cellNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cell_number", r.cellNumber, "")
	}
	if r.forcePasswordReset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force_password_reset", r.forcePasswordReset, "")
	}
	if r.githubUsername != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "github_username", r.githubUsername, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.phoneNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "phone_number", r.phoneNumber, "")
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
	if r.slackUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "slack_user_id", r.slackUserId, "")
	}
	if r.slackUsername != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "slack_username", r.slackUsername, "")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title", r.title, "")
	}
	if r.twitterUsername != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "twitter_username", r.twitterUsername, "")
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

type ApiUserContactInfosPartialUpdateRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	id int32
	patchedUserContactInfoRequest *PatchedUserContactInfoRequest
}

func (r ApiUserContactInfosPartialUpdateRequest) PatchedUserContactInfoRequest(patchedUserContactInfoRequest PatchedUserContactInfoRequest) ApiUserContactInfosPartialUpdateRequest {
	r.patchedUserContactInfoRequest = &patchedUserContactInfoRequest
	return r
}

func (r ApiUserContactInfosPartialUpdateRequest) Execute() (*UserContactInfo, *http.Response, error) {
	return r.ApiService.UserContactInfosPartialUpdateExecute(r)
}

/*
UserContactInfosPartialUpdate Method for UserContactInfosPartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this user contact info.
 @return ApiUserContactInfosPartialUpdateRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosPartialUpdate(ctx context.Context, id int32) ApiUserContactInfosPartialUpdateRequest {
	return ApiUserContactInfosPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserContactInfo
func (a *UserContactInfosAPIService) UserContactInfosPartialUpdateExecute(r ApiUserContactInfosPartialUpdateRequest) (*UserContactInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserContactInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/{id}/"
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
	localVarPostBody = r.patchedUserContactInfoRequest
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

type ApiUserContactInfosRetrieveRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	id int32
	prefetch *[]string
}

// List of fields for which to prefetch model instances and add those to the response
func (r ApiUserContactInfosRetrieveRequest) Prefetch(prefetch []string) ApiUserContactInfosRetrieveRequest {
	r.prefetch = &prefetch
	return r
}

func (r ApiUserContactInfosRetrieveRequest) Execute() (*UserContactInfo, *http.Response, error) {
	return r.ApiService.UserContactInfosRetrieveExecute(r)
}

/*
UserContactInfosRetrieve Method for UserContactInfosRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this user contact info.
 @return ApiUserContactInfosRetrieveRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosRetrieve(ctx context.Context, id int32) ApiUserContactInfosRetrieveRequest {
	return ApiUserContactInfosRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserContactInfo
func (a *UserContactInfosAPIService) UserContactInfosRetrieveExecute(r ApiUserContactInfosRetrieveRequest) (*UserContactInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserContactInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/{id}/"
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

type ApiUserContactInfosUpdateRequest struct {
	ctx context.Context
	ApiService UserContactInfosAPI
	id int32
	userContactInfoRequest *UserContactInfoRequest
}

func (r ApiUserContactInfosUpdateRequest) UserContactInfoRequest(userContactInfoRequest UserContactInfoRequest) ApiUserContactInfosUpdateRequest {
	r.userContactInfoRequest = &userContactInfoRequest
	return r
}

func (r ApiUserContactInfosUpdateRequest) Execute() (*UserContactInfo, *http.Response, error) {
	return r.ApiService.UserContactInfosUpdateExecute(r)
}

/*
UserContactInfosUpdate Method for UserContactInfosUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this user contact info.
 @return ApiUserContactInfosUpdateRequest
*/
func (a *UserContactInfosAPIService) UserContactInfosUpdate(ctx context.Context, id int32) ApiUserContactInfosUpdateRequest {
	return ApiUserContactInfosUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UserContactInfo
func (a *UserContactInfosAPIService) UserContactInfosUpdateExecute(r ApiUserContactInfosUpdateRequest) (*UserContactInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserContactInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserContactInfosAPIService.UserContactInfosUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/user_contact_infos/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userContactInfoRequest == nil {
		return localVarReturnValue, nil, reportError("userContactInfoRequest is required and must be specified")
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
	localVarPostBody = r.userContactInfoRequest
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
