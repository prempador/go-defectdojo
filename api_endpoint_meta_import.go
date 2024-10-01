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
	"os"
)


type EndpointMetaImportAPI interface {

	/*
	EndpointMetaImportCreate Method for EndpointMetaImportCreate

	Imports a CSV file into a product to propagate arbitrary meta and tags on endpoints.

By Names:
- Provide `product_name` of existing product

By ID:
- Provide the id of the product in the `product` parameter

In this scenario Defect Dojo will look up the product by the provided details.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEndpointMetaImportCreateRequest
	*/
	EndpointMetaImportCreate(ctx context.Context) ApiEndpointMetaImportCreateRequest

	// EndpointMetaImportCreateExecute executes the request
	//  @return EndpointMetaImporter
	EndpointMetaImportCreateExecute(r ApiEndpointMetaImportCreateRequest) (*EndpointMetaImporter, *http.Response, error)
}

// EndpointMetaImportAPIService EndpointMetaImportAPI service
type EndpointMetaImportAPIService service

type ApiEndpointMetaImportCreateRequest struct {
	ctx context.Context
	ApiService EndpointMetaImportAPI
	file *os.File
	createEndpoints *bool
	createTags *bool
	createDojoMeta *bool
	productName *string
	product *int32
}

func (r ApiEndpointMetaImportCreateRequest) File(file *os.File) ApiEndpointMetaImportCreateRequest {
	r.file = file
	return r
}

func (r ApiEndpointMetaImportCreateRequest) CreateEndpoints(createEndpoints bool) ApiEndpointMetaImportCreateRequest {
	r.createEndpoints = &createEndpoints
	return r
}

func (r ApiEndpointMetaImportCreateRequest) CreateTags(createTags bool) ApiEndpointMetaImportCreateRequest {
	r.createTags = &createTags
	return r
}

func (r ApiEndpointMetaImportCreateRequest) CreateDojoMeta(createDojoMeta bool) ApiEndpointMetaImportCreateRequest {
	r.createDojoMeta = &createDojoMeta
	return r
}

func (r ApiEndpointMetaImportCreateRequest) ProductName(productName string) ApiEndpointMetaImportCreateRequest {
	r.productName = &productName
	return r
}

func (r ApiEndpointMetaImportCreateRequest) Product(product int32) ApiEndpointMetaImportCreateRequest {
	r.product = &product
	return r
}

func (r ApiEndpointMetaImportCreateRequest) Execute() (*EndpointMetaImporter, *http.Response, error) {
	return r.ApiService.EndpointMetaImportCreateExecute(r)
}

/*
EndpointMetaImportCreate Method for EndpointMetaImportCreate

Imports a CSV file into a product to propagate arbitrary meta and tags on endpoints.

By Names:
- Provide `product_name` of existing product

By ID:
- Provide the id of the product in the `product` parameter

In this scenario Defect Dojo will look up the product by the provided details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEndpointMetaImportCreateRequest
*/
func (a *EndpointMetaImportAPIService) EndpointMetaImportCreate(ctx context.Context) ApiEndpointMetaImportCreateRequest {
	return ApiEndpointMetaImportCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EndpointMetaImporter
func (a *EndpointMetaImportAPIService) EndpointMetaImportCreateExecute(r ApiEndpointMetaImportCreateRequest) (*EndpointMetaImporter, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndpointMetaImporter
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EndpointMetaImportAPIService.EndpointMetaImportCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/endpoint_meta_import/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"
	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.createEndpoints != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "create_endpoints", r.createEndpoints, "", "")
	}
	if r.createTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "create_tags", r.createTags, "", "")
	}
	if r.createDojoMeta != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "create_dojo_meta", r.createDojoMeta, "", "")
	}
	if r.productName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "product_name", r.productName, "", "")
	}
	if r.product != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "product", r.product, "", "")
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
