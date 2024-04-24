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
)


type ApiTokenAuthAPI interface {

	/*
	ApiTokenAuthCreate Method for ApiTokenAuthCreate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiTokenAuthCreateRequest
	*/
	ApiTokenAuthCreate(ctx context.Context) ApiApiTokenAuthCreateRequest

	// ApiTokenAuthCreateExecute executes the request
	//  @return AuthToken
	ApiTokenAuthCreateExecute(r ApiApiTokenAuthCreateRequest) (*AuthToken, *http.Response, error)
}

// ApiTokenAuthAPIService ApiTokenAuthAPI service
type ApiTokenAuthAPIService service

type ApiApiTokenAuthCreateRequest struct {
	ctx context.Context
	ApiService ApiTokenAuthAPI
	username *string
	password *string
}

func (r ApiApiTokenAuthCreateRequest) Username(username string) ApiApiTokenAuthCreateRequest {
	r.username = &username
	return r
}

func (r ApiApiTokenAuthCreateRequest) Password(password string) ApiApiTokenAuthCreateRequest {
	r.password = &password
	return r
}

func (r ApiApiTokenAuthCreateRequest) Execute() (*AuthToken, *http.Response, error) {
	return r.ApiService.ApiTokenAuthCreateExecute(r)
}

/*
ApiTokenAuthCreate Method for ApiTokenAuthCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiTokenAuthCreateRequest
*/
func (a *ApiTokenAuthAPIService) ApiTokenAuthCreate(ctx context.Context) ApiApiTokenAuthCreateRequest {
	return ApiApiTokenAuthCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthToken
func (a *ApiTokenAuthAPIService) ApiTokenAuthCreateExecute(r ApiApiTokenAuthCreateRequest) (*AuthToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiTokenAuthAPIService.ApiTokenAuthCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/api-token-auth/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.username == nil {
		return localVarReturnValue, nil, reportError("username is required and must be specified")
	}
	if strlen(*r.username) < 1 {
		return localVarReturnValue, nil, reportError("username must have at least 1 elements")
	}
	if r.password == nil {
		return localVarReturnValue, nil, reportError("password is required and must be specified")
	}
	if strlen(*r.password) < 1 {
		return localVarReturnValue, nil, reportError("password must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded", "multipart/form-data", "application/json"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "username", r.username, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "")
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
