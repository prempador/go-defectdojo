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
	"time"
)


type RiskAcceptanceAPI interface {

	/*
	RiskAcceptanceDeletePreviewList Method for RiskAcceptanceDeletePreviewList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this risk_ acceptance.
	@return ApiRiskAcceptanceDeletePreviewListRequest
	*/
	RiskAcceptanceDeletePreviewList(ctx context.Context, id int32) ApiRiskAcceptanceDeletePreviewListRequest

	// RiskAcceptanceDeletePreviewListExecute executes the request
	//  @return PaginatedDeletePreviewList
	RiskAcceptanceDeletePreviewListExecute(r ApiRiskAcceptanceDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error)

	/*
	RiskAcceptanceDestroy Method for RiskAcceptanceDestroy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this risk_ acceptance.
	@return ApiRiskAcceptanceDestroyRequest
	*/
	RiskAcceptanceDestroy(ctx context.Context, id int32) ApiRiskAcceptanceDestroyRequest

	// RiskAcceptanceDestroyExecute executes the request
	RiskAcceptanceDestroyExecute(r ApiRiskAcceptanceDestroyRequest) (*http.Response, error)

	/*
	RiskAcceptanceDownloadProofRetrieve Method for RiskAcceptanceDownloadProofRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this risk_ acceptance.
	@return ApiRiskAcceptanceDownloadProofRetrieveRequest
	*/
	RiskAcceptanceDownloadProofRetrieve(ctx context.Context, id int32) ApiRiskAcceptanceDownloadProofRetrieveRequest

	// RiskAcceptanceDownloadProofRetrieveExecute executes the request
	//  @return RiskAcceptanceProof
	RiskAcceptanceDownloadProofRetrieveExecute(r ApiRiskAcceptanceDownloadProofRetrieveRequest) (*RiskAcceptanceProof, *http.Response, error)

	/*
	RiskAcceptanceList Method for RiskAcceptanceList

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRiskAcceptanceListRequest
	*/
	RiskAcceptanceList(ctx context.Context) ApiRiskAcceptanceListRequest

	// RiskAcceptanceListExecute executes the request
	//  @return PaginatedRiskAcceptanceList
	RiskAcceptanceListExecute(r ApiRiskAcceptanceListRequest) (*PaginatedRiskAcceptanceList, *http.Response, error)

	/*
	RiskAcceptancePartialUpdate Method for RiskAcceptancePartialUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this risk_ acceptance.
	@return ApiRiskAcceptancePartialUpdateRequest
	*/
	RiskAcceptancePartialUpdate(ctx context.Context, id int32) ApiRiskAcceptancePartialUpdateRequest

	// RiskAcceptancePartialUpdateExecute executes the request
	//  @return RiskAcceptance
	RiskAcceptancePartialUpdateExecute(r ApiRiskAcceptancePartialUpdateRequest) (*RiskAcceptance, *http.Response, error)

	/*
	RiskAcceptanceRetrieve Method for RiskAcceptanceRetrieve

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this risk_ acceptance.
	@return ApiRiskAcceptanceRetrieveRequest
	*/
	RiskAcceptanceRetrieve(ctx context.Context, id int32) ApiRiskAcceptanceRetrieveRequest

	// RiskAcceptanceRetrieveExecute executes the request
	//  @return RiskAcceptance
	RiskAcceptanceRetrieveExecute(r ApiRiskAcceptanceRetrieveRequest) (*RiskAcceptance, *http.Response, error)

	/*
	RiskAcceptanceUpdate Method for RiskAcceptanceUpdate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id A unique integer value identifying this risk_ acceptance.
	@return ApiRiskAcceptanceUpdateRequest
	*/
	RiskAcceptanceUpdate(ctx context.Context, id int32) ApiRiskAcceptanceUpdateRequest

	// RiskAcceptanceUpdateExecute executes the request
	//  @return RiskAcceptance
	RiskAcceptanceUpdateExecute(r ApiRiskAcceptanceUpdateRequest) (*RiskAcceptance, *http.Response, error)
}

// RiskAcceptanceAPIService RiskAcceptanceAPI service
type RiskAcceptanceAPIService service

type ApiRiskAcceptanceDeletePreviewListRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	id int32
	limit *int32
	offset *int32
}

// Number of results to return per page.
func (r ApiRiskAcceptanceDeletePreviewListRequest) Limit(limit int32) ApiRiskAcceptanceDeletePreviewListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ApiRiskAcceptanceDeletePreviewListRequest) Offset(offset int32) ApiRiskAcceptanceDeletePreviewListRequest {
	r.offset = &offset
	return r
}

func (r ApiRiskAcceptanceDeletePreviewListRequest) Execute() (*PaginatedDeletePreviewList, *http.Response, error) {
	return r.ApiService.RiskAcceptanceDeletePreviewListExecute(r)
}

/*
RiskAcceptanceDeletePreviewList Method for RiskAcceptanceDeletePreviewList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this risk_ acceptance.
 @return ApiRiskAcceptanceDeletePreviewListRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptanceDeletePreviewList(ctx context.Context, id int32) ApiRiskAcceptanceDeletePreviewListRequest {
	return ApiRiskAcceptanceDeletePreviewListRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaginatedDeletePreviewList
func (a *RiskAcceptanceAPIService) RiskAcceptanceDeletePreviewListExecute(r ApiRiskAcceptanceDeletePreviewListRequest) (*PaginatedDeletePreviewList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedDeletePreviewList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptanceDeletePreviewList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/{id}/delete_preview/"
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

type ApiRiskAcceptanceDestroyRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	id int32
}

func (r ApiRiskAcceptanceDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.RiskAcceptanceDestroyExecute(r)
}

/*
RiskAcceptanceDestroy Method for RiskAcceptanceDestroy

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this risk_ acceptance.
 @return ApiRiskAcceptanceDestroyRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptanceDestroy(ctx context.Context, id int32) ApiRiskAcceptanceDestroyRequest {
	return ApiRiskAcceptanceDestroyRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *RiskAcceptanceAPIService) RiskAcceptanceDestroyExecute(r ApiRiskAcceptanceDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptanceDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/{id}/"
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

type ApiRiskAcceptanceDownloadProofRetrieveRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	id int32
}

func (r ApiRiskAcceptanceDownloadProofRetrieveRequest) Execute() (*RiskAcceptanceProof, *http.Response, error) {
	return r.ApiService.RiskAcceptanceDownloadProofRetrieveExecute(r)
}

/*
RiskAcceptanceDownloadProofRetrieve Method for RiskAcceptanceDownloadProofRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this risk_ acceptance.
 @return ApiRiskAcceptanceDownloadProofRetrieveRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptanceDownloadProofRetrieve(ctx context.Context, id int32) ApiRiskAcceptanceDownloadProofRetrieveRequest {
	return ApiRiskAcceptanceDownloadProofRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RiskAcceptanceProof
func (a *RiskAcceptanceAPIService) RiskAcceptanceDownloadProofRetrieveExecute(r ApiRiskAcceptanceDownloadProofRetrieveRequest) (*RiskAcceptanceProof, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RiskAcceptanceProof
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptanceDownloadProofRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/{id}/download_proof/"
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

type ApiRiskAcceptanceListRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	acceptedBy *string
	acceptedFindings *[]int32
	decision *string
	decisionDetails *string
	expirationDate *time.Time
	expirationDateHandled *time.Time
	expirationDateWarned *time.Time
	limit *int32
	name *string
	notes *[]int32
	o *[]string
	offset *int32
	owner *int32
	reactivateExpired *bool
	recommendation *string
	recommendationDetails *string
	restartSlaExpired *bool
}

func (r ApiRiskAcceptanceListRequest) AcceptedBy(acceptedBy string) ApiRiskAcceptanceListRequest {
	r.acceptedBy = &acceptedBy
	return r
}

func (r ApiRiskAcceptanceListRequest) AcceptedFindings(acceptedFindings []int32) ApiRiskAcceptanceListRequest {
	r.acceptedFindings = &acceptedFindings
	return r
}

// Risk treatment decision by risk owner  * &#x60;A&#x60; - Accept (The risk is acknowledged, yet remains) * &#x60;V&#x60; - Avoid (Do not engage with whatever creates the risk) * &#x60;M&#x60; - Mitigate (The risk still exists, yet compensating controls make it less of a threat) * &#x60;F&#x60; - Fix (The risk is eradicated) * &#x60;T&#x60; - Transfer (The risk is transferred to a 3rd party)
func (r ApiRiskAcceptanceListRequest) Decision(decision string) ApiRiskAcceptanceListRequest {
	r.decision = &decision
	return r
}

func (r ApiRiskAcceptanceListRequest) DecisionDetails(decisionDetails string) ApiRiskAcceptanceListRequest {
	r.decisionDetails = &decisionDetails
	return r
}

func (r ApiRiskAcceptanceListRequest) ExpirationDate(expirationDate time.Time) ApiRiskAcceptanceListRequest {
	r.expirationDate = &expirationDate
	return r
}

func (r ApiRiskAcceptanceListRequest) ExpirationDateHandled(expirationDateHandled time.Time) ApiRiskAcceptanceListRequest {
	r.expirationDateHandled = &expirationDateHandled
	return r
}

func (r ApiRiskAcceptanceListRequest) ExpirationDateWarned(expirationDateWarned time.Time) ApiRiskAcceptanceListRequest {
	r.expirationDateWarned = &expirationDateWarned
	return r
}

// Number of results to return per page.
func (r ApiRiskAcceptanceListRequest) Limit(limit int32) ApiRiskAcceptanceListRequest {
	r.limit = &limit
	return r
}

func (r ApiRiskAcceptanceListRequest) Name(name string) ApiRiskAcceptanceListRequest {
	r.name = &name
	return r
}

func (r ApiRiskAcceptanceListRequest) Notes(notes []int32) ApiRiskAcceptanceListRequest {
	r.notes = &notes
	return r
}

// Ordering  * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending)
func (r ApiRiskAcceptanceListRequest) O(o []string) ApiRiskAcceptanceListRequest {
	r.o = &o
	return r
}

// The initial index from which to return the results.
func (r ApiRiskAcceptanceListRequest) Offset(offset int32) ApiRiskAcceptanceListRequest {
	r.offset = &offset
	return r
}

func (r ApiRiskAcceptanceListRequest) Owner(owner int32) ApiRiskAcceptanceListRequest {
	r.owner = &owner
	return r
}

func (r ApiRiskAcceptanceListRequest) ReactivateExpired(reactivateExpired bool) ApiRiskAcceptanceListRequest {
	r.reactivateExpired = &reactivateExpired
	return r
}

// Recommendation from the security team.  * &#x60;A&#x60; - Accept (The risk is acknowledged, yet remains) * &#x60;V&#x60; - Avoid (Do not engage with whatever creates the risk) * &#x60;M&#x60; - Mitigate (The risk still exists, yet compensating controls make it less of a threat) * &#x60;F&#x60; - Fix (The risk is eradicated) * &#x60;T&#x60; - Transfer (The risk is transferred to a 3rd party)
func (r ApiRiskAcceptanceListRequest) Recommendation(recommendation string) ApiRiskAcceptanceListRequest {
	r.recommendation = &recommendation
	return r
}

func (r ApiRiskAcceptanceListRequest) RecommendationDetails(recommendationDetails string) ApiRiskAcceptanceListRequest {
	r.recommendationDetails = &recommendationDetails
	return r
}

func (r ApiRiskAcceptanceListRequest) RestartSlaExpired(restartSlaExpired bool) ApiRiskAcceptanceListRequest {
	r.restartSlaExpired = &restartSlaExpired
	return r
}

func (r ApiRiskAcceptanceListRequest) Execute() (*PaginatedRiskAcceptanceList, *http.Response, error) {
	return r.ApiService.RiskAcceptanceListExecute(r)
}

/*
RiskAcceptanceList Method for RiskAcceptanceList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRiskAcceptanceListRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptanceList(ctx context.Context) ApiRiskAcceptanceListRequest {
	return ApiRiskAcceptanceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedRiskAcceptanceList
func (a *RiskAcceptanceAPIService) RiskAcceptanceListExecute(r ApiRiskAcceptanceListRequest) (*PaginatedRiskAcceptanceList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRiskAcceptanceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptanceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.acceptedBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accepted_by", r.acceptedBy, "")
	}
	if r.acceptedFindings != nil {
		t := *r.acceptedFindings
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "accepted_findings", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "accepted_findings", t, "multi")
		}
	}
	if r.decision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "decision", r.decision, "")
	}
	if r.decisionDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "decision_details", r.decisionDetails, "")
	}
	if r.expirationDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expiration_date", r.expirationDate, "")
	}
	if r.expirationDateHandled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expiration_date_handled", r.expirationDateHandled, "")
	}
	if r.expirationDateWarned != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expiration_date_warned", r.expirationDateWarned, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.notes != nil {
		t := *r.notes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "notes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "notes", t, "multi")
		}
	}
	if r.o != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "o", r.o, "csv")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.owner != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "owner", r.owner, "")
	}
	if r.reactivateExpired != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reactivate_expired", r.reactivateExpired, "")
	}
	if r.recommendation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recommendation", r.recommendation, "")
	}
	if r.recommendationDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recommendation_details", r.recommendationDetails, "")
	}
	if r.restartSlaExpired != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "restart_sla_expired", r.restartSlaExpired, "")
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

type ApiRiskAcceptancePartialUpdateRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	id int32
	patchedRiskAcceptanceRequest *PatchedRiskAcceptanceRequest
}

func (r ApiRiskAcceptancePartialUpdateRequest) PatchedRiskAcceptanceRequest(patchedRiskAcceptanceRequest PatchedRiskAcceptanceRequest) ApiRiskAcceptancePartialUpdateRequest {
	r.patchedRiskAcceptanceRequest = &patchedRiskAcceptanceRequest
	return r
}

func (r ApiRiskAcceptancePartialUpdateRequest) Execute() (*RiskAcceptance, *http.Response, error) {
	return r.ApiService.RiskAcceptancePartialUpdateExecute(r)
}

/*
RiskAcceptancePartialUpdate Method for RiskAcceptancePartialUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this risk_ acceptance.
 @return ApiRiskAcceptancePartialUpdateRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptancePartialUpdate(ctx context.Context, id int32) ApiRiskAcceptancePartialUpdateRequest {
	return ApiRiskAcceptancePartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RiskAcceptance
func (a *RiskAcceptanceAPIService) RiskAcceptancePartialUpdateExecute(r ApiRiskAcceptancePartialUpdateRequest) (*RiskAcceptance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RiskAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptancePartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/{id}/"
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
	localVarPostBody = r.patchedRiskAcceptanceRequest
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

type ApiRiskAcceptanceRetrieveRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	id int32
}

func (r ApiRiskAcceptanceRetrieveRequest) Execute() (*RiskAcceptance, *http.Response, error) {
	return r.ApiService.RiskAcceptanceRetrieveExecute(r)
}

/*
RiskAcceptanceRetrieve Method for RiskAcceptanceRetrieve

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this risk_ acceptance.
 @return ApiRiskAcceptanceRetrieveRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptanceRetrieve(ctx context.Context, id int32) ApiRiskAcceptanceRetrieveRequest {
	return ApiRiskAcceptanceRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RiskAcceptance
func (a *RiskAcceptanceAPIService) RiskAcceptanceRetrieveExecute(r ApiRiskAcceptanceRetrieveRequest) (*RiskAcceptance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RiskAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptanceRetrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/{id}/"
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

type ApiRiskAcceptanceUpdateRequest struct {
	ctx context.Context
	ApiService RiskAcceptanceAPI
	id int32
	riskAcceptanceRequest *RiskAcceptanceRequest
}

func (r ApiRiskAcceptanceUpdateRequest) RiskAcceptanceRequest(riskAcceptanceRequest RiskAcceptanceRequest) ApiRiskAcceptanceUpdateRequest {
	r.riskAcceptanceRequest = &riskAcceptanceRequest
	return r
}

func (r ApiRiskAcceptanceUpdateRequest) Execute() (*RiskAcceptance, *http.Response, error) {
	return r.ApiService.RiskAcceptanceUpdateExecute(r)
}

/*
RiskAcceptanceUpdate Method for RiskAcceptanceUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id A unique integer value identifying this risk_ acceptance.
 @return ApiRiskAcceptanceUpdateRequest
*/
func (a *RiskAcceptanceAPIService) RiskAcceptanceUpdate(ctx context.Context, id int32) ApiRiskAcceptanceUpdateRequest {
	return ApiRiskAcceptanceUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RiskAcceptance
func (a *RiskAcceptanceAPIService) RiskAcceptanceUpdateExecute(r ApiRiskAcceptanceUpdateRequest) (*RiskAcceptance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RiskAcceptance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RiskAcceptanceAPIService.RiskAcceptanceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/risk_acceptance/{id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.riskAcceptanceRequest == nil {
		return localVarReturnValue, nil, reportError("riskAcceptanceRequest is required and must be specified")
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
	localVarPostBody = r.riskAcceptanceRequest
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
